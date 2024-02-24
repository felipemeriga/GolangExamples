package example

import (
	"testing"
)

// -------------------------------  Fibonacci  ----------------------------------
func Test_ExecuteFibo(t *testing.T) {
	ExecuteFibo()
}

func Test_Example(t *testing.T) {
	RemoveRightZeros()
}

// -------------------------------------------------------------------------------

// -------------------------------  Error Group  ----------------------------------
func Test_ErrorGroup(t *testing.T) {
	ErrorGroup()
}

// --------------------------------------------------------------------------------

// -------------------------------  Waiting Group  ----------------------------------
func Test_SimpleWaitingGroup(t *testing.T) {
	SimpleWaitingGroup()
}

func Test_WaitGroupExample(t *testing.T) {
	WaitGroupExample()
}

// --------------------------------------------------------------------------------

// -------------------------------  Reflection and Unmarshalling  ----------------------------------
func Test_FormUnmarshall(t *testing.T) {
	FormUnmarshall()
}

func Test_Reflection(t *testing.T) {
	Reflection()
}

// --------------------------------------------------------------------------------

// -------------------------------  Channel-----------------------------------------
func Test_Parallel(t *testing.T) {
	Parallel()
}

func Test_WaitingChannel(t *testing.T) {
	WaitingChannel()
}

func Test_WaitingTicker(t *testing.T) {
	WaitingTicker()
}

// --------------------------------------------------------------------------------

// -------------------------------Mutex----------------------------------------------
func Test_SafeCounterMutexExample(t *testing.T) {
	SafeCounterMutexExample()
}

func Test_MapRMutexExample(t *testing.T) {
	MapRMutexExample()
}

// ----------------------------------Reflection----------------------------------------------

func Test_ReflectionOnCharacter(t *testing.T) {
	ReflectionOnCharacter()
}

// -----------------------------------Pointers-----------------------------

func Test_PointersAndMemory(t *testing.T) {
	PointersAndMemory()
}

func Test_TestArray(t *testing.T) {
	TestArray()
}

// -----------------------------------Cache-----------------------------

func Test_TestCache(t *testing.T) {
	TestCache()
}

// ------- Time  ----------

func Test_Time(t *testing.T) {
	Time()
}

// ------- Application Restart  ----------

func Test_Restart_Application(t *testing.T) {
	RestartExample()
}
