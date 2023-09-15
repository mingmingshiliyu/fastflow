package main

import (
	"fmt"
	"github.com/shiningrush/fastflow/pkg/entity"
	"github.com/shiningrush/fastflow/pkg/mod"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"reflect"
	"time"

	"github.com/shiningrush/fastflow"
	mongoKeeper "github.com/shiningrush/fastflow/keeper/mongo"
	"github.com/shiningrush/fastflow/pkg/entity/run"
	mongoStore "github.com/shiningrush/fastflow/store/mongo"
)

func init() {
	mod.RegisterCanRunAction(PrintAction1{})
	mod.RegisterCanRunAction(PrintAction{})
}

type PrintAction struct {
}

// Name define the unique action identity, it will be used by Task
func (a PrintAction) Name() string {
	return "PrintAction"
}
func (a PrintAction) Run(ctx run.ExecuteContext, params interface{}) error {
	fmt.Println("action1123 start: ", time.Now())
	duration, _ := time.ParseDuration("5s")
	time.Sleep(duration)
	fmt.Println("action1123 end: ", time.Now())
	return nil
}

type PrintAction1 struct {
}

// Name define the unique action identity, it will be used by Task
func (a PrintAction1) Name() string {
	return "PrintAction1"
}
func (a PrintAction1) Run(ctx run.ExecuteContext, params interface{}) error {
	fmt.Println("print start: ", time.Now())
	return nil
}

func main() {
	// Register action
	//fastflow.RegisterAction([]run.Action{
	//	&PrintAction{},
	//})
	readDag("./test-dag.yaml")
	//dynamicFunction()
	//StructRegister(PrintAction{}, PrintAction1{})
	// init keeper, it used to e
	keeper := mongoKeeper.NewKeeper(&mongoKeeper.KeeperOption{
		Key: "worker-1",
		// if your mongo does not set user/pwd, you should remove it
		ConnStr:  "mongodb://81.68.119.162:27018/fastflow?authSource=admin",
		Database: "fastflow",
		Prefix:   "test",
	})
	if err := keeper.Init(); err != nil {
		log.Fatal(fmt.Errorf("init keeper failed: %w", err))
	}

	// init store
	st := mongoStore.NewStore(&mongoStore.StoreOption{
		// if your mongo does not set user/pwd, you should remove it
		ConnStr:  "mongodb://81.68.119.162:27018/fastflow?authSource=admin",
		Database: "fastflow",
		Prefix:   "test",
	})
	if err := st.Init(); err != nil {
		log.Fatal(fmt.Errorf("init store failed: %w", err))
	}
	go createDagAndInstance()

	// start fastflow
	if err := fastflow.Start(&fastflow.InitialOption{
		Keeper: keeper,
		Store:  st,
		// use yaml to define dag 结构体传递会影响相对路径，函数不会
		ReadDagFromDir: "./examples/simple/",
	}); err != nil {
		panic(fmt.Sprintf("init fastflow failed: %s", err))
	}

}

func createDagAndInstance() {
	// wait fast start completed
	time.Sleep(1 * time.Second)

	// run some dag instance
	//for i := 0; i < 2; i++ {
	_, err := mod.GetCommander().RunDag("test-dag", nil)
	if err != nil {
		log.Fatal(err)
	}
	time.Sleep(time.Second * 2)
	//}
}

func StructRegister(act ...run.Action) {
	params := make([]run.Action, 0)

	for _, stru := range act {
		name := reflect.TypeOf(stru).Name()
		//name1 := reflect.TypeOf(PrintAction1{}).Name()
		mod.RegisterCanRunAction(stru)
		action := mod.GetCanRunAction(name)
		//typeReflect[name1] = reflect.TypeOf(PrintAction1{})
		//of := reflect.TypeOf(PrintAction{})
		//tp := reflect.TypeOf(PrintAction1{})
		//instance2 := reflect.New(typeReflect[name1]).Elem().Interface().(run.Action)
		params = append(params, action)
		//params = append(params, instance2)
	}
	fastflow.RegisterAction(params)

}

func dynamicFunction() {
	//registered := make(map[string]run.Action)
	typeReflect := make(map[string]reflect.Type)
	name := reflect.TypeOf(PrintAction{}).Name()
	name1 := reflect.TypeOf(PrintAction1{}).Name()
	typeReflect[name] = reflect.TypeOf(PrintAction{})
	typeReflect[name1] = reflect.TypeOf(PrintAction1{})
	params := make([]run.Action, 0)
	//of := reflect.TypeOf(PrintAction{})
	//tp := reflect.TypeOf(PrintAction1{})
	instance1 := reflect.New(typeReflect[name]).Elem().Interface().(run.Action)
	instance2 := reflect.New(typeReflect[name1]).Elem().Interface().(run.Action)
	params = append(params, instance1)
	params = append(params, instance2)
	//for _, param := range params {
	//	param.Run(nil, nil)
	//}
	fastflow.RegisterAction(params)
}

// 定义一个接口
type Action interface {
	Execute()
}

// 子类1
type CustomAction1 struct {
	// 自定义字段
}

// 实现接口的方法
func (ca CustomAction1) Execute() {
	fmt.Println("Custom action 1 executed")
}

// 子类2
type CustomAction2 struct {
	// 自定义字段
}

// 实现接口的方法
func (ca CustomAction2) Execute() {
	fmt.Println("Custom action 2 executed")
}

// 子类3
type CustomAction3 struct {
	// 自定义字段
}

// 实现接口的方法
func (ca CustomAction3) Execute() {
	fmt.Println("Custom action 3 executed")
}
func main1() {
	params := make([]Action, 0)

	// 获取子类的reflect.Type
	type1 := reflect.TypeOf(CustomAction1{})
	type2 := reflect.TypeOf(CustomAction2{})
	type3 := reflect.TypeOf(CustomAction3{})

	// 使用反射创建子类实例并转换为Action接口类型
	instance1 := reflect.New(type1).Elem().Interface().(Action)
	instance2 := reflect.New(type2).Elem().Interface().(Action)
	instance3 := reflect.New(type3).Elem().Interface().(Action)

	// 将实例添加到params切片中
	params = append(params, instance1)
	params = append(params, instance2)
	params = append(params, instance3)

	// 调用执行方法
	for _, action := range params {
		action.Execute()
	}
}

func readDag(path string) error {
	file, err := os.ReadFile("/Users/wtst45x/GolandProjects/fastflow/examples/simple/test-dag.yaml")
	if err != nil {
		return err
	}
	dag := entity.Dag{}
	if err := yaml.Unmarshal(file, &dag); err != nil {
		return err
	}
	actions := []run.Action{}
	for _, task := range dag.Tasks {
		if action := mod.GetCanRunAction(task.ActionName); action != nil {
			actions = append(actions, action)
		} else {
			fmt.Sprintf("task: %v not register into fastflow", task.ActionName)
		}
	}
	fastflow.RegisterAction(actions)
	return nil
}
