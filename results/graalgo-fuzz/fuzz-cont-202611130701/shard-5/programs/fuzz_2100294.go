package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 int8 }
type GenS2 struct { F0 uint64; F1 uint32; F2 uint }
func genFn0(p0 int32) uint32 {
    fmt.Println(p0)
    p0 = (-(p0))
    fmt.Println(p0)
    _ = p0
    return 399
}

func main() {
    var v1 int32 = -463
    var v2 uint64 = 247
    var v3 int8 = -22
    var v4 string = ("foo bar" + "")
    v5 := make(map[uint32]string)
    v6 := []string{v4, v4, v4}
    fmt.Println(v4)
    for fi7 := 0; fi7 < 2; fi7++ {
        fmt.Println(v1, fi7, v1)
        var v8 int32 = v1
        _ = fi7
        _ = v8
    }
    v4 = "a"
    v1 = (v1 | v1)
    var v9 int32 = (v1 * -489)
    v9 = v9
    if (4294967295 > 0) {
        for fi10 := 0; fi10 < 20; fi10++ {
            var v11 int64 = 897
            if false {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
            }
            _ = fi10
            _ = v11
        }
        for fi12 := 0; fi12 < 19; fi12++ {
            var v13 int32 = 1
            _ = len(v6)
            fmt.Println(v3)
            _ = fi12
            _ = v13
        }
        v5[2147483647] = v4
    }
    v14 := GenS1{F0: (!(false && (!(true && (((float64((1.5 / float32((0.0 + 3.14159)))) / -2.25) >= 1e10) && (1 > 2)))))), F1: int64(v2), F2: (-(v3))}
    if false {
        fmt.Println(v9)
        switch -1 {
        case -592:
            v2 = v2
        default:
            v1 += v1
        }
        fmt.Println(v4)
    } else {
        v15 := GenS2{F0: (^(v2)), F1: 42, F2: 610}
        _ = v15
    }
    v1 = v9
    var v16 uint16 = 32767
    v17 := GenS2{F0: (v2 >> ((742) & 63)), F1: 42, F2: 2}
    var v18 bool = (("hello" + "hello") > v4)
    var v19 bool = ("foo bar" >= v4)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
