package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 string; F2 uint64 }
func genFn0(p0 uint, p1 int16) int8 {
    v6 := GenS1{F0: "foo bar"}
    fmt.Println(p1)
    v7 := []int{859}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 82
}

func main() {
    var v1 string = "hello"
    v2 := make(map[uint32]bool)
    var v3 int16 = -1
    v4 := make(map[string]string)
    var v5 int8 = 127
    fmt.Println(v1)
    v3 = v3
    v3 = -1
    v1 = (((v1 + ("\n\t" + (v1 + ((v1 + (("hello" + ("foo bar" + ("hello" + v1))) + (v1 + "a"))) + ((("" + v1) + (v1 + v1)) + v1))))) + (v1 + v1)) + v1)
    v8 := GenS1{F0: ""}
    v9 := v4[(v1 + (v1 + (v1 + v1)))]
    var v10 uint = 42
    if (!(((2 > 0) && true) && true)) {
        v11 := GenS2{F0: true, F1: "a", F2: 1000}
        v5 = genFn0(v10, int16(v10))
        _ = v11
    }
    var v12 uint32 = 65535
    switch -2147483648 {
    case 4294967295:
        var v13 float64 = -2.25
        _ = v13
    case 45:
        var v14 uint16 = 2
        _ = v14
    }
    v1 = v1
    v1 = v1
    var v15 uint8 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
