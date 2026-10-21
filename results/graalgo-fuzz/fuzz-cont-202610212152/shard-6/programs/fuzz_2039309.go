package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 int16 }
type GenS2 struct { F0 uint16; F1 uint32 }
func genFn0() string {
    var v5 uint8 = 0
    var v6 int32 = -307
    var v7 bool = false
    _ = v5
    _ = v6
    _ = v7
    return ""
}

func main() {
    var v1 uint32 = 4294967295
    var v2 uint64 = 1
    v3 := []string{"", "hello"}
    var v4 int8 = 1
    fmt.Println(v2, v4)
    v8 := GenS1{F0: -59, F1: "a", F2: 1000}
    var v9 float32 = 0.1
    if (!(("foo bar" <= "a") || true)) {
        v4 -= 0
        for fi10 := 0; fi10 < 5; fi10++ {
            var v11 uint32 = v1
            switch "a" {
            case "":
                v12 := GenS2{F0: uint16(v2), F1: 1}
                _ = v12
            case "a":
                var v13 uint64 = (v2 % ((v2) | 1))
                _ = v13
            default:
                v2 = (v2 % (((v2 >> ((uint(v2)) & 63))) | 1))
            }
            var v14 int = fi10
            _ = fi10
            _ = v11
            _ = v14
        }
        var v15 float32 = v9
        _ = v15
    } else {
        var v16 string = ("a" + ("a" + "foo bar"))
        _ = v16
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v17 uint16 = 347
    v18 := GenS2{F0: v17, F1: (v1 | v1)}
    var v19 uint = 1
    for fi20 := 0; fi20 < 10; fi20++ {
        var v21 uint64 = v2
        var v22 uint32 = 42
        _ = fi20
        _ = v21
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
