package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 string, p1 string, p2 uint8) bool {
    var v4 int64 = -824
    fmt.Println(p2, v4, v4)
    switch p1 {
    case "a":
        var v5 int16 = -32768
        _ = v5
    case "":
        var v6 bool = (!false)
        _ = v6
    case "foo bar":
        for fi7 := 0; fi7 < 18; fi7++ {
            var v8 uint64 = 2
            _ = fi7
            _ = v8
        }
    default:
        fmt.Println(p1, p2, p2)
    }
    v9 := GenS1{F0: 22}
    fmt.Println(p0, p2, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v9
    return false
}

func main() {
    var v1 int32 = 2147483647
    var v2 int32 = v1
    var v3 uint64 = 349
    for fi10 := 0; fi10 < 18; fi10++ {
        fmt.Println(v2)
        _ = fi10
    }
    v11 := GenS1{F0: -35}
    v3 = 76
    var v12 uint = 18446744073709551615
    var v13 int32 = v1
    var v14 float32 = 0.1
    fmt.Println(v3, v3, v14)
    if (true || false) {
        var v15 bool = true
        for fi16 := 0; fi16 < 3; fi16++ {
            v13 = v13
            _ = fi16
        }
        _ = v15
    } else {
        for fi17 := 0; fi17 < 10; fi17++ {
            if true {
                var v18 uint32 = 1
                var v19 uint = (v12 & v12)
                v20 := []int{(fi17 / ((fi17) | 1)), fi17, -334}
                _ = v18
                _ = v19
                _ = v20
            }
            v14 = (float32(1.5) + v14)
            _ = fi17
        }
        var v21 int8 = -1
        var v22 int32 = v1
        _ = v21
        _ = v22
    }
    var v23 int = 4294967295
    v24 := []int{(-(v23)), v23, -9223372036854775808, v23}
    var v25 float32 = v14
    var v26 string = ("hello" + "hello")
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
}
