package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int8; F1 bool }
type GenS3 struct { F0 int8; F1 int8 }
func genFn0() uint8 {
    var v4 int8 = -128
    var v5 uint8 = 0
    var v6 int32 = 2147483647
    if ((-76 == int8(v6)) && (1 == v5)) {
        for fi7 := 0; fi7 < 9; fi7++ {
            v5 -= 2
            var v8 float64 = 0.1
            var v9 uint16 = 42
            _ = fi7
            _ = v8
            _ = v9
        }
    } else {
        switch -399 {
        case 397:
            var v10 int = -568
            _ = v10
        }
    }
    _ = v4
    _ = v5
    _ = v6
    return 1
}

func main() {
    v1 := []int{127, 4294967295, 32767, 4294967295}
    var v2 int64 = -505
    v3 := []int{42, -9223372036854775808, 649}
    _ = len(v3)
    if true {
        var v11 uint64 = 212
        _ = v11
    }
    var v12 int = 127
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v2 = (v2 / ((v2) | 1))
    var v13 string = "\n\t"
    for fi14 := 0; fi14 < 7; fi14++ {
        if (false && (42 <= 42)) {
            v15 := GenS1{F0: 0}
            _ = v15
        }
        var v16 uint32 = 65535
        _ = fi14
        _ = v16
    }
    for fi17 := 0; fi17 < 20; fi17++ {
        if len(v3) > 0 {
            v3[0] = fi17
        }
        _ = fi17
    }
    var v18 float64 = 0.1
    if len(v1) > 0 {
        v1[0] = v12
    }
    var v19 int16 = 1000
    if len(v3) > 0 {
        v3[0] = v12
    }
    for fi20 := 0; fi20 < 3; fi20++ {
        switch v13 {
        case "":
            var v21 float64 = v18
            _ = v21
        case "a":
            var v22 bool = false
            _ = v22
        case "foo bar":
            v23 := GenS2{F0: 127, F1: (1 <= 0)}
            _ = v23
        default:
            v13 = "\n\t"
        }
        for fi24 := 0; fi24 < 7; fi24++ {
            switch (-(fi20)) {
            case -1:
                if false {
                    var v25 int32 = 562
                    _ = v25
                } else {
                    if (1 > 128) {
                        var v26 float64 = v18
                        var v27 uint64 = 187
                        _ = v26
                        _ = v27
                    }
                    v28 := GenS3{F0: -1, F1: -80}
                    if false {
                        var v29 int = v12
                        _ = v29
                    }
                    _ = v28
                }
            }
            _ = fi24
        }
        _ = fi20
    }
    var v30 int = (^(v12))
    _ = len(v1)
    fmt.Println(v13, v19, v2)
    v19 += v19
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v12
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v30
}
