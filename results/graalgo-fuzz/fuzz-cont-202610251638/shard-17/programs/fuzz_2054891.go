package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 uint16 }
func genFn0() uint8 {
    var v5 float64 = 1e10
    var v6 string = ""
    v7 := GenS1{F0: -1, F1: 0}
    v8 := []int{}
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 127
}

func main() {
    var v1 int64 = -523
    var v2 int64 = (v1 ^ 484)
    v3 := make(map[string]int)
    v4 := []uint32{0, 1}
    var v9 string = "a"
    v2 = 1
    var v10 uint32 = 224
    fmt.Println(v2)
    v11 := GenS2{F0: 2, F1: 2}
    fmt.Println(v9, v2, v2)
    if ((2147483647 == 748) || (v9 <= "hello")) {
        v12 := v3[v9]
        _ = len(v4)
        var v13 int32 = 468
        _ = v12
        _ = v13
    } else {
        switch (v9 + v9) {
        case "a":
            switch 322 {
            case 127:
                for fi14 := 0; fi14 < 5; fi14++ {
                    if (!(!((((1.5 + -2.25) - (1e10 / ((float64(-2.25) + 1e10) - 1.5))) > 1.5) && true))) {
                        fmt.Println(v2, v2)
                        v15 := GenS2{F0: 65535, F1: v10}
                        _ = v15
                    }
                    _ = fi14
                }
            case 9223372036854775807:
                switch (("\n\t" + (v9 + v9)) + (("hello" + v9) + v9)) {
                case "":
                    for fi16 := 0; fi16 < 4; fi16++ {
                        v17 := GenS1{F0: -980, F1: 42}
                        _ = fi16
                        _ = v17
                    }
                }
            case 32767:
                if false {
                    v2 -= v1
                } else {
                    var v18 float64 = (1e10 + 3.14159)
                    _ = v18
                }
            default:
                var v19 int64 = v2
                _ = v19
            }
        }
        var v20 float64 = (float64((1.5 * 1e10)) + 0.0)
        switch 127 {
        case 2147483647:
            var v21 uint16 = 65535
            _ = v21
        }
        _ = v20
    }
    v10 = 0
    for fi22 := 0; fi22 < 5; fi22++ {
        fmt.Println(v1)
        v2 = 9223372036854775807
        _ = fi22
    }
    if (!true) {
        for fi23 := 0; fi23 < 7; fi23++ {
            v2 += (-(v2))
            fmt.Println(v2, v9, v10)
            v10 = v10
            _ = fi23
        }
        var v24 int16 = -171
        _ = v24
    }
    v10 = (v10 ^ 2)
    switch 172 {
    case 127:
        v9 = v9
    case 32767:
        if (!("\n\t" == "")) {
            v9 = ("foo bar" + "foo bar")
        } else {
            var v25 uint8 = 0
            _ = v25
        }
    case 2147483647:
        var v26 uint64 = 0
        _ = v26
    }
    v27 := GenS3{F0: 32767}
    fmt.Println(v9, v2, v2)
    v28 := GenS2{F0: 888, F1: v10}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v9
}
