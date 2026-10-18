package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32; F1 bool }
func genFn0(p0 uint16) int16 {
    for fi7 := 0; fi7 < 17; fi7++ {
        fmt.Println(fi7, p0)
        var v8 uint64 = 4294967295
        switch "" {
        case "foo bar":
            p0 = (^(p0))
        case "hello":
            switch (^(fi7)) {
            case -799:
                var v9 int16 = 42
                _ = v9
            }
        case "a":
            var v10 int16 = 32767
            _ = v10
        }
        _ = fi7
        _ = v8
    }
    p0 = (p0 - 1)
    _ = p0
    return 1000
}
func genFn1(p0 uint, p1 uint64, p2 uint) int8 {
    p1 += uint64(p0)
    v11 := GenS1{F0: 95}
    v12 := []int{-1, 952}
    v13 := GenS1{F0: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    _ = v13
    return -128
}

func main() {
    var v1 float64 = float64(1e10)
    var v2 uint8 = 2
    var v3 uint64 = 1
    var v4 uint64 = (v3 % (((v3 ^ 2)) | 1))
    v5 := make(map[string]int)
    var v6 int32 = 2147483647
    delete(v5, ("a" + ("hello" + "a")))
    v5[""] = 103
    v14 := GenS1{F0: 2147483648}
    v15 := []int{-9223372036854775808, 128, 319, -290}
    for fi16 := 0; fi16 < 3; fi16++ {
        if (!false) {
            switch (-(fi16)) {
            case 937:
                v17 := GenS2{F0: 0, F1: false}
                _ = v17
            case 139:
                v18 := GenS2{F0: 2147483647, F1: false}
                _ = v18
            default:
                if true {
                    _ = len(v5)
                    _ = len(v5)
                }
            }
        }
        v1 = v1
        _ = fi16
    }
    for fi19 := 0; fi19 < 4; fi19++ {
        for fi20 := 0; fi20 < 17; fi20++ {
            var v21 uint = 1
            _ = fi20
            _ = v21
        }
        _ = fi19
    }
    fmt.Println(v4, v4)
    v22 := GenS1{F0: 813}
    if false {
        if (1.5 != (float32((v1 + v1)) + (1e10 / 1.0))) {
            v1 += v1
        }
        v6 = v6
        var v23 uint32 = 4294967295
        _ = v23
    }
    delete(v5, "a")
    var v24 uint = 0
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v22
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
