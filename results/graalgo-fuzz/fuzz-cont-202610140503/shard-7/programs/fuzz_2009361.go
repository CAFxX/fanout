package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string }
type GenS2 struct { F0 int; F1 uint64; F2 bool }
func genFn0(p0 int, p1 uint8, p2 uint64) int {
    var v5 uint32 = 1
    p1 -= (p1 % ((p1) | 1))
    v6 := []int{int(p1), p0}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return p0
}

func main() {
    v1 := []uint32{2}
    v2 := []uint32{}
    var v3 string = ("hello" + "hello")
    v4 := []int{}
    var v7 string = v3
    if true {
        var v8 bool = true
        for fi9 := 0; fi9 < 3; fi9++ {
            v10 := GenS1{F0: -922, F1: (v7 + v7)}
            for fi11 := 0; fi11 < 19; fi11++ {
                v12 := GenS1{F0: -187, F1: v3}
                _ = fi11
                _ = v12
            }
            _ = fi9
            _ = v10
        }
        _ = v8
    } else {
        for fi13 := 0; fi13 < 5; fi13++ {
            v1 = append(v1, 960)
            if len(v4) > 0 {
                v4[0] = genFn0(fi13, 0, 0)
            }
            _ = fi13
        }
    }
    v14 := GenS2{F0: -100, F1: 775, F2: (!(false || (0.0 != 1e10)))}
    var v15 bool = (v3 > v7)
    for fi16 := 0; fi16 < 4; fi16++ {
        v2 = append(v2, 42)
        switch "hello" {
        case "a":
            fmt.Println(v3, v15, fi16)
        case "\n\t":
            for fi17 := 0; fi17 < 3; fi17++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                if ((417 <= 264) && (!(v15 || v15))) {
                    v2 = append(v2, 2147483648)
                } else {
                    v15 = v15
                }
                _ = fi17
            }
        case "foo bar":
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = fi16
    }
    v7 = (((v3 + "hello") + "hello") + v7)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v18 string = ("a" + (v3 + v7))
    v3 = v3
    switch -346 {
    case 128:
        var v19 int32 = 1
        _ = v19
    default:
        var v20 float64 = (1.5 * ((float64((1.5 / 0.1)) / 1.5) - (((1.5 * 1.5) / float64(3.14159)) - (3.14159 - (1e10 * 1.5)))))
        _ = v20
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
