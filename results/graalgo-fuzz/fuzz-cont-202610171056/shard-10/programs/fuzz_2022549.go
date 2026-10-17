package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
func genFn0(p0 bool) string {
    v6 := []int{}
    var v7 uint8 = 255
    fmt.Println(v7, p0)
    v7 = v7
    _ = p0
    _ = v6
    _ = v7
    return ("" + "foo bar")
}
func genFn1() uint32 {
    v8 := GenS1{F0: 776, F1: "foo bar"}
    if false {
        if (("\n\t" > "a") || false) {
            var v9 string = (((("a" + "foo bar") + ("hello" + "")) + "") + "foo bar")
            _ = v9
        }
    }
    v10 := GenS1{F0: 613, F1: "a"}
    _ = v10
    _ = v8
    return 1
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []string{(("a" + ("a" + ("\n\t" + "hello"))) + ("a" + ("\n\t" + ""))), "", "", "hello"}
    var v3 int16 = 42
    var v4 string = "hello"
    v5 := make(map[string]int)
    var v11 uint32 = 427
    var v12 int32 = 1
    var v13 int64 = 9223372036854775807
    var v14 uint8 = 42
    var v15 uint64 = 18446744073709551615
    var v16 int64 = (v13 >> ((1) & 63))
    var v17 uint32 = 0
    for fi18 := 0; fi18 < 3; fi18++ {
        if len(v2) > 0 {
            v2[0] = (v4 + v4)
        }
        if ((v15 - (^(v15))) != 2) {
            fmt.Println(v14)
        } else {
            _ = len(v5)
        }
        _ = fi18
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
