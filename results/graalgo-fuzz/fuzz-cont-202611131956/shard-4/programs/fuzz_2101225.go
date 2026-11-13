package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    var v1 int32 = -2147483648
    var v2 int64 = 547
    var v3 int8 = 42
    var v4 bool = false
    v5 := GenS1{F0: int(v2)}
    v4 = true
    v6 := GenS1{F0: 2147483647}
    v3 = (v3 ^ (^(v3)))
    v7 := GenS1{F0: -563}
    fmt.Println(v1, v1, v3)
    var v8 uint32 = 65535
    for fi9 := 0; fi9 < 11; fi9++ {
        v4 = (-375 != 1)
        _ = fi9
    }
    for fi10 := 0; fi10 < 20; fi10++ {
        if ("a" < "a") {
            var v11 uint = 18446744073709551615
            v12 := GenS1{F0: int(v3)}
            var v13 float32 = (1.5 + 0.0)
            _ = v11
            _ = v12
            _ = v13
        } else {
            if v4 {
                fmt.Println(fi10, v3, v4)
                v3 = -1
            }
        }
        _ = fi10
    }
    v14 := []int{}
    if len(v14) > 0 {
        v14[0] = -2147483648
    }
    if v4 {
        switch "a" {
        case "foo bar":
            v2 += (v2 + v2)
        default:
            var v15 int8 = v3
            _ = v15
        }
    } else {
        if len(v14) > 0 {
            v14[0] = 32768
        }
    }
    var v16 float32 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
