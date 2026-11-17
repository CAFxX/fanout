package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 uint32; F2 uint32 }
func genFn0(p0 uint64, p1 uint64, p2 int32) uint8 {
    fmt.Println(p0, p2, p2)
    v5 := GenS1{F0: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 210
}

func main() {
    var v1 int8 = 0
    v2 := make(map[string]string)
    var v3 uint16 = 1
    var v4 int32 = -19
    v6 := v2[("a" + ("a" + ("a" + "hello")))]
    var v7 uint64 = 1
    v7 = (v7 + (v7 ^ (v7 % (((^(v7))) | 1))))
    var v8 int = 242
    v3 = 514
    var v9 bool = (!(!false))
    v9 = (v9 || v9)
    var v10 string = (v6 + (("" + (v6 + (v6 + v6))) + "\n\t"))
    for fi11 := 0; fi11 < 20; fi11++ {
        var v12 uint64 = (v7 - v7)
        _ = fi11
        _ = v12
    }
    if ((true || ((!v9) || (0 > 1))) && (v9 && (!((false || false) || true)))) {
        v13 := GenS2{F0: ((1e10 + 0.0) < 3.14159), F1: 2147483647, F2: uint32(v7)}
        _ = v13
    } else {
        v10 = ""
        if false {
            var v14 int16 = 1000
            _ = v14
        }
    }
    v9 = ((((146 != 0) && ((219 > 255) && (-2.25 != 0.1))) || v9) || v9)
    for fi15 := 0; fi15 < 14; fi15++ {
        for fi16 := 0; fi16 < 7; fi16++ {
            v17 := GenS1{F0: 128}
            v8 = (^(v8))
            _ = fi16
            _ = v17
        }
        _ = fi15
    }
    v18 := []int{1, -9223372036854775808}
    var v19 float32 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
