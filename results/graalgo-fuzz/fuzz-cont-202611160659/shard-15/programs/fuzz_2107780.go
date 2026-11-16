package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint }
func genFn0() int8 {
    v5 := GenS1{F0: "foo bar"}
    var v6 uint32 = 0
    switch ("" + "\n\t") {
    case "a":
        fmt.Println(v6)
    default:
        v7 := GenS1{F0: "foo bar"}
        _ = v7
    }
    if (42 == 255) {
        fmt.Println(v6)
    } else {
        fmt.Println(v6)
        var v8 int64 = 0
        _ = v8
    }
    var v9 uint32 = (^(v6))
    _ = v5
    _ = v6
    _ = v9
    return 33
}

func main() {
    var v1 int32 = -1
    v2 := make(map[string]int)
    var v3 float32 = 0.0
    var v4 int32 = 718
    v10 := []int{4294967295, -806, 4294967295, 273}
    if (!true) {
        v11 := GenS1{F0: ""}
        _ = v11
    }
    if true {
        fmt.Println(v1, v4, v1)
    } else {
        v12 := GenS1{F0: ("a" + "hello")}
        v3 -= v3
        var v13 string = ""
        _ = v12
        _ = v13
    }
    var v14 int = 2147483647
    v15 := GenS2{F0: 0}
    v16 := GenS1{F0: ("\n\t" + "foo bar")}
    if (true && true) {
        var v17 uint8 = 42
        _ = v17
    }
    fmt.Println(v3, v1, v3)
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
