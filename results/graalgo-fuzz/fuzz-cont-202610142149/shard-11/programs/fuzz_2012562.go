package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool; F2 uint32 }
type GenS2 struct { F0 uint; F1 bool; F2 int32 }

func main() {
    v1 := []string{"hello", "foo bar"}
    v2 := []int{-419, -905, 128}
    var v3 float32 = (0.0 / -2.25)
    var v4 uint16 = 255
    if (-1 != 1) {
        var v5 uint = 2
        v6 := GenS1{F0: v4, F1: true, F2: 0}
        _ = v5
        _ = v6
    } else {
        v3 = ((((1.5 + v3) / ((((v3 + v3) * (float32(0.1) / (float32(0.0) / v3))) / v3) * v3)) / float32(1.5)) / 0.1)
        v4 = v4
    }
    v7 := GenS1{F0: v4, F1: (-1 == 0), F2: 145}
    if (false && ((1 > 2) && true)) {
        for fi8 := 0; fi8 < 13; fi8++ {
            for fi9 := 0; fi9 < 20; fi9++ {
                if len(v1) > 0 {
                    v1[0] = "\n\t"
                }
                switch fi8 {
                case 31:
                    v3 += v3
                }
                v10 := GenS1{F0: (v4 / ((v4) | 1)), F1: false, F2: 4294967295}
                _ = fi9
                _ = v10
            }
            _ = fi8
        }
        v3 += (v3 - 0.1)
        for fi11 := 0; fi11 < 15; fi11++ {
            if false {
                var v12 int = (fi11 | -780)
                if len(v1) > 0 {
                    v1[0] = ("\n\t" + "a")
                }
                _ = v12
            }
            fmt.Println(v4)
            _ = fi11
        }
    } else {
        var v13 int64 = -1
        _ = v13
    }
    for fi14 := 0; fi14 < 7; fi14++ {
        v15 := GenS1{F0: v4, F1: false, F2: 0}
        switch "foo bar" {
        case "":
            fmt.Println(fi14)
        case "hello":
            if len(v2) > 0 {
                v2[0] = fi14
            }
        default:
            var v16 uint8 = 255
            _ = v16
        }
        _ = fi14
        _ = v15
    }
    v4 = v4
    v17 := GenS2{F0: 0, F1: (923 == 255), F2: -1}
    v4 -= 0
    v18 := GenS2{F0: 65535, F1: true, F2: -839}
    for fi19 := 0; fi19 < 20; fi19++ {
        for fi20 := 0; fi20 < 5; fi20++ {
            var v21 uint64 = 4294967295
            switch (fi19 / ((fi20) | 1)) {
            case -9223372036854775808:
                var v22 int8 = -97
                _ = v22
            }
            _ = fi20
            _ = v21
        }
        _ = fi19
    }
    var v23 string = ""
    switch "hello" {
    case "\n\t":
        var v24 int32 = -1
        _ = v24
    case "":
        v3 -= 1.5
    case "foo bar":
        v4 += 255
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v7
}
