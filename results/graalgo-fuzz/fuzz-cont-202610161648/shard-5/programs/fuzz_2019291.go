package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32; F1 int16; F2 string }
type GenS3 struct { F0 uint }

func main() {
    var v1 int64 = 1
    var v2 float32 = float32(3.14159)
    v3 := []uint32{2147483647}
    var v4 float32 = (v2 - v2)
    _ = len(v3)
    if true {
        if (!(2 == 42)) {
            if true {
                v2 = 1.5
                var v5 string = ("a" + "")
                _ = v5
            }
            for fi6 := 0; fi6 < 2; fi6++ {
                v7 := GenS1{F0: 94}
                _ = fi6
                _ = v7
            }
        } else {
            switch 2147483647 {
            case 1:
                v8 := GenS1{F0: 32767}
                _ = v8
            case -703:
                v1 -= v1
            default:
                v9 := GenS2{F0: 2147483648, F1: 1, F2: "\n\t"}
                _ = v9
            }
            v1 -= (v1 / ((v1) | 1))
            if len(v3) > 0 {
                v3[0] = 305
            }
        }
        _ = len(v3)
    } else {
        fmt.Println(v2, v4, v2)
        fmt.Println(v1, v4, v4)
    }
    v10 := GenS1{F0: 445}
    v4 = (v2 - (((v4 - (v2 + v2)) + 3.14159) / v2))
    v2 -= ((v4 + v4) + v4)
    for fi11 := 0; fi11 < 10; fi11++ {
        v4 = v4
        if true {
            v1 += (v1 + (v1 & 562))
            _ = len(v3)
        } else {
            _ = len(v3)
            for fi12 := 0; fi12 < 7; fi12++ {
                var v13 uint64 = 507
                _ = fi12
                _ = v13
            }
        }
        fmt.Println(v2, v2)
        _ = fi11
    }
    switch "hello" {
    case "hello":
        var v14 string = "a"
        _ = v14
    case "foo bar":
        var v15 string = "foo bar"
        _ = v15
    case "":
        for fi16 := 0; fi16 < 2; fi16++ {
            if true {
                var v17 int16 = -1
                _ = v17
            } else {
                for fi18 := 0; fi18 < 13; fi18++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    v4 = 0.1
                    if (0 == 0) {
                        var v19 bool = (4294967295 != 42)
                        fmt.Println(v1, v1, v4)
                        _ = v19
                    }
                    _ = fi18
                }
            }
            switch (fi16 << ((223) & 63)) {
            case -1:
                fmt.Println(v4)
            }
            _ = fi16
        }
    default:
        v1 = v1
    }
    for fi20 := 0; fi20 < 15; fi20++ {
        var v21 int16 = 42
        _ = fi20
        _ = v21
    }
    _ = len(v3)
    var v22 float64 = (-2.25 * 3.14159)
    for fi23 := 0; fi23 < 19; fi23++ {
        var v24 float64 = 0.1
        var v25 uint = 18446744073709551615
        for fi26 := 0; fi26 < 12; fi26++ {
            if false {
                if len(v3) > 0 {
                    v3[0] = 4294967295
                }
                v27 := GenS3{F0: v25}
                switch fi26 {
                case 0:
                    var v28 int16 = -32768
                    v29 := GenS2{F0: 4294967295, F1: 42, F2: ((("a" + ("a" + "hello")) + "foo bar") + ("hello" + "foo bar"))}
                    _ = v28
                    _ = v29
                case 4294967295:
                    v3 = append(v3, 0)
                default:
                    v1 += (-(v1))
                }
                _ = v27
            }
            _ = fi26
        }
        _ = fi23
        _ = v24
        _ = v25
    }
    for fi30 := 0; fi30 < 18; fi30++ {
        v1 -= (v1 | v1)
        v31 := GenS2{F0: 65535, F1: int16(v1), F2: ""}
        _ = fi30
        _ = v31
    }
    var v32 int16 = 340
    v33 := GenS3{F0: 2}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v3
    _ = v32
    _ = v33
    _ = v4
}
