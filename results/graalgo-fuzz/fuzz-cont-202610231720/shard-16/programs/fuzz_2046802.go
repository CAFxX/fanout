package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint; F1 int16 }
func genFn0(p0 int64, p1 int32) uint8 {
    v4 := []int{0, 1, 128, 250}
    if true {
        for fi5 := 0; fi5 < 20; fi5++ {
            if (45 <= 0) {
                var v6 string = ("" + "\n\t")
                _ = v6
            } else {
                v4 = append(v4, 127)
            }
            for fi7 := 0; fi7 < 10; fi7++ {
                var v8 int32 = p1
                _ = fi7
                _ = v8
            }
            if false {
                for fi9 := 0; fi9 < 3; fi9++ {
                    var v10 int32 = 1
                    p0 -= (p0 ^ (p0 + p0))
                    var v11 float32 = (0.0 + 1.5)
                    _ = fi9
                    _ = v10
                    _ = v11
                }
                v12 := GenS1{F0: 2}
                _ = v12
            } else {
                v13 := GenS1{F0: 220}
                p1 = -72
                if (32767 <= 2) {
                    var v14 string = (("\n\t" + ("a" + "a")) + "a")
                    _ = v14
                }
                _ = v13
            }
            _ = fi5
        }
        p1 = 808
    }
    fmt.Println(p0, p0)
    for fi15 := 0; fi15 < 17; fi15++ {
        p1 = -1
        var v16 uint = 514
        _ = fi15
        _ = v16
    }
    _ = p0
    _ = p1
    _ = v4
    return 127
}

func main() {
    v1 := make(map[int]string)
    v2 := []uint32{129, 0, 65535, 65535}
    var v3 uint = 1
    v1[-1] = "\n\t"
    fmt.Println(v3)
    if true {
        if len(v2) > 0 {
            v2[0] = 2147483647
        }
        var v17 bool = (-128 <= 100)
        switch 39 {
        case -472:
            if true {
                v3 += (v3 - (v3 >> ((v3) & 63)))
                var v18 bool = false
                _ = v18
            }
        default:
            _ = len(v1)
        }
        _ = v17
    }
    var v19 int32 = -2147483648
    var v20 float64 = float64(1e10)
    for fi21 := 0; fi21 < 11; fi21++ {
        v22 := GenS2{F0: (v3 ^ v3), F1: 1}
        v20 = v20
        fmt.Println(v3, v3, v19)
        _ = fi21
        _ = v22
    }
    fmt.Println(v19)
    fmt.Println(v3, v3, v20)
    for fi23 := 0; fi23 < 10; fi23++ {
        if false {
            v3 -= (v3 - v3)
        }
        _ = fi23
    }
    v20 -= 1e10
    switch 32768 {
    case 128:
        var v24 uint64 = 1
        _ = v24
    case 32767:
        var v25 uint32 = 2
        _ = v25
    default:
        var v26 string = (("foo bar" + "\n\t") + ("\n\t" + ("" + "foo bar")))
        _ = v26
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
