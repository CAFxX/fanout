package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32; F2 int32 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 bool) int32 {
    var v4 int32 = 2147483647
    var v5 uint64 = 885
    _ = p0
    _ = v4
    _ = v5
    return (v4 / ((v4) | 1))
}
func genFn1(p0 string) uint32 {
    for fi6 := 0; fi6 < 11; fi6++ {
        for fi7 := 0; fi7 < 17; fi7++ {
            if false {
                var v8 int32 = -891
                _ = v8
            }
            _ = fi7
        }
        v9 := []int{(fi6 | 1), -1}
        p0 = "\n\t"
        _ = fi6
        _ = v9
    }
    fmt.Println(p0)
    var v10 int64 = 1
    v11 := GenS1{F0: 1, F1: 481, F2: genFn0(false)}
    _ = p0
    _ = v10
    _ = v11
    return 4294967295
}

func main() {
    var v1 uint = 495
    var v2 uint64 = 2
    v3 := []int{}
    var v12 string = "\n\t"
    v12 = ("hello" + "")
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v12, v1, v1)
    fmt.Println(v12)
    v2 = (v2 / ((v2) | 1))
    for fi13 := 0; fi13 < 17; fi13++ {
        var v14 int8 = 103
        v15 := GenS2{F0: 1}
        v16 := GenS2{F0: 2}
        _ = fi13
        _ = v14
        _ = v15
        _ = v16
    }
    fmt.Println(v2, v12, v2)
    var v17 bool = (!true)
    v3 = append(v3, 454)
    switch (v12 + (("hello" + ("foo bar" + (v12 + "foo bar"))) + v12)) {
    case "":
        var v18 uint = 65535
        _ = v18
    default:
        var v19 int = 4294967295
        _ = v19
    }
    if (1.5 <= 1.5) {
        for fi20 := 0; fi20 < 14; fi20++ {
            var v21 int32 = -1
            v12 = (v12 + "")
            if v17 {
                v21 = genFn0((!v17))
            } else {
                v21 -= genFn0(v17)
                for fi22 := 0; fi22 < 1; fi22++ {
                    v2 -= v2
                    _ = fi22
                }
            }
            _ = fi20
            _ = v21
        }
        if (genFn1(v12) <= genFn1(v12)) {
            for fi23 := 0; fi23 < 2; fi23++ {
                var v24 uint8 = 255
                fmt.Println(v24, v12, v2)
                _ = fi23
                _ = v24
            }
        }
    } else {
        var v25 uint64 = (v2 | v2)
        var v26 int16 = 0
        _ = v25
        _ = v26
    }
    v27 := GenS2{F0: 255}
    for fi28 := 0; fi28 < 2; fi28++ {
        v12 = v12
        v3 = append(v3, 32767)
        v17 = (!v17)
        _ = fi28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v12
    _ = v17
    _ = v2
    _ = v27
    _ = v3
}
