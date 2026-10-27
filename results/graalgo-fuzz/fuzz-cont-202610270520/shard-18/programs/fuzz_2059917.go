package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32; F2 int }
type GenS2 struct { F0 uint64; F1 uint32; F2 uint16 }
func genFn0(p0 int16, p1 int32, p2 uint64) uint32 {
    v4 := GenS1{F0: true, F1: 1, F2: 128}
    p0 -= (p0 & p0)
    var v5 int = 32768
    p0 = 1000
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 136
}
func genFn1(p0 int8, p1 string) int {
    switch ((p1 + (p1 + (p1 + "a"))) + p1) {
    case "foo bar":
        if false {
            v6 := []int{128, -766, int(p0), 4294967295}
            _ = v6
        } else {
            p1 = ("a" + "foo bar")
            p0 = -128
            v7 := GenS1{F0: (((false || (p0 > -128)) || (-186 >= -598)) && (1 > 1)), F1: -1, F2: 9223372036854775807}
            _ = v7
        }
    default:
        fmt.Println(p0, p0)
    }
    fmt.Println(p1)
    _ = p0
    _ = p1
    return 0
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint = 1
    v3 := make(map[uint32]uint32)
    v8 := GenS1{F0: false, F1: 0, F2: 9223372036854775807}
    for fi9 := 0; fi9 < 20; fi9++ {
        _ = len(v1)
        if (fi9 <= 2147483647) {
            var v10 bool = true
            v2 = 4294967295
            _ = v10
        }
        v11 := GenS1{F0: true, F1: 1, F2: 295}
        _ = fi9
        _ = v11
    }
    v2 -= v2
    var v12 float64 = ((1e10 + (-2.25 * -2.25)) - 0.1)
    _ = len(v1)
    v2 += v2
    fmt.Println(v12)
    for fi13 := 0; fi13 < 14; fi13++ {
        v14 := []int{(fi13 + genFn1(-2, ("a" + (("a" + ((("" + ("hello" + ("\n\t" + "a"))) + "a") + "foo bar")) + "a")))), (fi13 & (^(fi13))), (fi13 + (fi13 - (fi13 & fi13)))}
        v12 += v12
        _ = fi13
        _ = v14
    }
    v2 = (v2 & 894)
    var v15 string = ("hello" + "foo bar")
    if (0.1 > 3.14159) {
        for fi16 := 0; fi16 < 15; fi16++ {
            var v17 int16 = 0
            _ = fi16
            _ = v17
        }
        v18 := GenS2{F0: 2, F1: 803, F2: 2}
        _ = v18
    } else {
        v2 -= v2
        for fi19 := 0; fi19 < 13; fi19++ {
            fmt.Println(v12)
            if true {
                v15 = (v15 + "hello")
            } else {
                switch (fi19 / ((fi19) | 1)) {
                case 2147483647:
                    fmt.Println(v12, fi19, v15)
                case 128:
                    v2 -= v2
                }
                v20 := []int{}
                delete(v1, 0)
                _ = v20
            }
            v12 -= v12
            _ = fi19
        }
    }
    v21 := GenS1{F0: true, F1: -901, F2: -1}
    v22 := GenS2{F0: 42, F1: 0, F2: 2}
    if false {
        if false {
            var v23 int = 2147483647
            fmt.Println(v12)
            fmt.Println(v23, v2)
            _ = v23
        }
        var v24 uint16 = 255
        _ = v24
    }
    var v25 bool = (32767 != 32767)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v21
    _ = v22
    _ = v25
    _ = v3
    _ = v8
}
