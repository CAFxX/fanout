package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint32; F1 int16; F2 string }
func genFn0(p0 int32, p1 uint) uint32 {
    if (!(903 != 65535)) {
        p1 += p1
        p0 = (-(p0))
        v4 := GenS1{F0: 876, F1: 912, F2: 0}
        _ = v4
    } else {
        fmt.Println(p1)
        p1 = p1
    }
    var v5 int8 = -1
    v6 := []int{4294967295, 127, -1, -698}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 65535
}

func main() {
    v1 := []int{-930}
    v2 := []uint32{65535, 4294967295}
    v3 := []int{2147483647}
    for fi7 := 0; fi7 < 18; fi7++ {
        for fi8 := 0; fi8 < 16; fi8++ {
            var v9 uint16 = 144
            v9 += (^(v9))
            v9 = v9
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    if len(v1) > 0 {
        v1[0] = 32768
    }
    switch ("" + "foo bar") {
    case "foo bar":
        var v10 float32 = 3.14159
        _ = v10
    case "":
        var v11 int = 2147483647
        _ = v11
    default:
        if true {
            if ((45 != 180) && (2147483647 > 19)) {
                var v12 bool = (true && (!false))
                fmt.Println(v12)
                _ = v12
            } else {
                switch "hello" {
                case "\n\t":
                    var v13 float32 = (0.0 / -2.25)
                    _ = v13
                default:
                    var v14 float64 = 0.1
                    _ = v14
                }
            }
        } else {
            var v15 string = "foo bar"
            if false {
                for fi16 := 0; fi16 < 13; fi16++ {
                    v17 := GenS1{F0: 42, F1: 2147483647, F2: 150}
                    _ = len(v1)
                    var v18 int64 = -1
                    _ = fi16
                    _ = v17
                    _ = v18
                }
                v15 = v15
                fmt.Println(v15)
            }
            v19 := GenS2{F0: 255}
            _ = v15
            _ = v19
        }
    }
    var v20 float64 = (1e10 + 3.14159)
    for fi21 := 0; fi21 < 17; fi21++ {
        var v22 float32 = ((1.5 / -2.25) * 3.14159)
        v23 := GenS3{F0: 42, F1: 938, F2: ("" + "a")}
        _ = fi21
        _ = v22
        _ = v23
    }
    v24 := GenS2{F0: 42}
    if len(v2) > 0 {
        v2[0] = 2147483647
    }
    switch "hello" {
    case "":
        v20 += float64(1.5)
    case "hello":
        for fi25 := 0; fi25 < 18; fi25++ {
            v26 := GenS1{F0: 255, F1: 0, F2: 255}
            _ = fi25
            _ = v26
        }
    }
    var v27 float32 = 1.5
    v27 -= v27
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if true {
        v28 := GenS1{F0: 42, F1: 910, F2: 128}
        if (!(42 >= 0)) {
            var v29 int32 = 879
            var v30 string = "a"
            _ = v29
            _ = v30
        }
        _ = v28
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v24
    _ = v27
    _ = v3
}
