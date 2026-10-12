package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 int8 }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 uint }
func genFn0(p0 string, p1 int, p2 uint8) uint32 {
    v7 := []int{}
    for fi8 := 0; fi8 < 20; fi8++ {
        var v9 uint64 = 18446744073709551615
        if (0 > -752) {
            fmt.Println(p0, v9)
            v10 := GenS1{F0: true, F1: 0, F2: 1}
            _ = v10
        }
        fmt.Println(p1, v9, v9)
        _ = fi8
        _ = v9
    }
    var v11 uint = 1
    v12 := GenS2{F0: (v11 + v11), F1: (p2 - (p2 & p2))}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    _ = v7
    return 2
}
func genFn1() bool {
    v13 := GenS1{F0: false, F1: 65535, F2: 0}
    var v14 string = ("a" + "")
    if false {
        v14 = v14
        var v15 float32 = -2.25
        var v16 bool = (0 < 32768)
        _ = v15
        _ = v16
    } else {
        for fi17 := 0; fi17 < 7; fi17++ {
            v18 := GenS3{F0: 2}
            _ = fi17
            _ = v18
        }
        v14 = "a"
        for fi19 := 0; fi19 < 1; fi19++ {
            var v20 uint64 = 4294967295
            _ = fi19
            _ = v20
        }
    }
    _ = v13
    _ = v14
    return (!true)
}

func main() {
    var v1 int = 42
    v2 := []int{}
    var v3 int64 = 506
    var v4 float64 = -2.25
    v5 := []int{1, 2147483647, (v1 % (((v1 & (v1 % ((v1) | 1)))) | 1))}
    var v6 bool = true
    v6 = v6
    var v21 string = "a"
    if len(v5) > 0 {
        v5[0] = 127
    }
    if (42 <= 1000) {
        v4 = float64((1.5 - (((((1e10 / 1e10) / 1.0) - 1.5) - 0.1) * 0.0)))
        fmt.Println(v21, v21)
        v4 = (((0.1 + v4) / (v4 * float64((3.14159 + 1e10)))) - v4)
    } else {
        v3 = v3
        if v6 {
            v2 = append(v2, -324)
            var v22 string = "foo bar"
            v3 = v3
            _ = v22
        } else {
            fmt.Println(v3)
        }
    }
    var v23 bool = (!true)
    v2 = append(v2, (v1 / ((v1) | 1)))
    var v24 int32 = -1
    for fi25 := 0; fi25 < 2; fi25++ {
        for fi26 := 0; fi26 < 19; fi26++ {
            v5 = append(v5, (fi26 >> ((910) & 63)))
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v27 uint64 = 1
            _ = fi26
            _ = v27
        }
        _ = fi25
    }
    v5 = append(v5, (v1 << ((0) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
