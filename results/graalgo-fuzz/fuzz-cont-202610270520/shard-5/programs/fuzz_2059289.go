package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
func genFn0(p0 bool) bool {
    if ("a" >= ("hello" + "foo bar")) {
        p0 = (147 == 143)
    } else {
        fmt.Println(p0)
        v6 := GenS1{F0: ("a" + ("" + "hello")), F1: 49}
        _ = v6
    }
    if p0 {
        for fi7 := 0; fi7 < 7; fi7++ {
            if false {
                v8 := []int{4294967295}
                _ = v8
            } else {
                fmt.Println(fi7)
                p0 = (520 >= 397)
            }
            var v9 int64 = 1
            var v10 string = "a"
            _ = fi7
            _ = v10
            _ = v9
        }
    }
    _ = p0
    return (949 < 42)
}
func genFn1() int {
    switch "" {
    case "hello":
        for fi11 := 0; fi11 < 17; fi11++ {
            v12 := GenS1{F0: "hello", F1: -441}
            _ = fi11
            _ = v12
        }
    case "a":
        v13 := []int{953}
        _ = v13
    }
    v14 := []int{155}
    var v15 int32 = -1
    _ = v14
    _ = v15
    return 517
}

func main() {
    var v1 int64 = -328
    var v2 bool = true
    v3 := make(map[int]bool)
    var v4 uint16 = 723
    var v5 int32 = 1
    fmt.Println(v1, v1)
    fmt.Println(v2, v2, v2)
    _ = len(v3)
    if v2 {
        var v16 int = 888
        _ = v16
    }
    var v17 uint64 = 0
    var v18 float64 = (1.5 / (3.14159 * (0.1 + 0.0)))
    v4 -= v4
    delete(v3, 106)
    fmt.Println(v4, v17, v2)
    var v19 int32 = v5
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
