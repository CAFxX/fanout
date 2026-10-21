package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 bool }
type GenS2 struct { F0 int64 }

func main() {
    v1 := make(map[string]int)
    var v2 int32 = -1
    v3 := []string{"hello", ("" + "hello"), (("" + ((("" + ("" + "foo bar")) + "hello") + "a")) + "hello")}
    var v4 int32 = 989
    fmt.Println(v2, v2)
    fmt.Println(v2, v2)
    v4 = (v2 % ((v2) | 1))
    fmt.Println(v4)
    if (9223372036854775807 >= -1) {
        switch 4294967295 {
        case -643:
            var v5 int8 = 100
            _ = v5
        case -9223372036854775808:
            if len(v3) > 0 {
                v3[0] = "\n\t"
            }
        case -1:
            v6 := GenS1{F0: false, F1: false, F2: false}
            _ = v6
        default:
            var v7 uint16 = 65535
            _ = v7
        }
        v4 = 1
    }
    fmt.Println(v4)
    v8 := GenS1{F0: (!((0 < 1) && (722 != 2147483647))), F1: (!("\n\t" == "a")), F2: false}
    fmt.Println(v4)
    var v9 int8 = -1
    if (1 > 568) {
        v9 = v9
        fmt.Println(v4, v2)
        var v10 string = "\n\t"
        _ = v10
    }
    _ = len(v3)
    _ = len(v3)
    v11 := GenS2{F0: 684}
    fmt.Println(v2, v4, v2)
    for fi12 := 0; fi12 < 2; fi12++ {
        if (-501 > 0) {
            v13 := GenS1{F0: (77 == 1), F1: false, F2: (true && (159 == 597))}
            var v14 int32 = 0
            _ = v13
            _ = v14
        }
        _ = fi12
    }
    if false {
        fmt.Println(v4, v4)
        v4 = (v2 * -268)
        var v15 uint8 = 92
        _ = v15
    }
    for fi16 := 0; fi16 < 10; fi16++ {
        var v17 int64 = int64(v2)
        var v18 uint64 = 0
        _ = fi16
        _ = v17
        _ = v18
    }
    for fi19 := 0; fi19 < 5; fi19++ {
        var v20 float32 = 3.14159
        var v21 string = "a"
        v2 = (-(v2))
        _ = fi19
        _ = v20
        _ = v21
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
