package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 int64; F1 uint32; F2 uint64 }
type GenS3 struct { F0 int64 }
func genFn0(p0 uint, p1 bool) bool {
    var v4 uint = (^(p0))
    v4 += v4
    v4 = (v4 / ((v4) | 1))
    _ = p0
    _ = p1
    _ = v4
    return p1
}
func genFn1(p0 bool) int64 {
    var v5 uint8 = 0
    if true {
        var v6 int = 127
        _ = v6
    }
    _ = p0
    _ = v5
    return 0
}

func main() {
    var v1 uint8 = 128
    var v2 int16 = 304
    var v3 string = ""
    for fi7 := 0; fi7 < 11; fi7++ {
        fmt.Println(v1)
        _ = fi7
    }
    if true {
        for fi8 := 0; fi8 < 10; fi8++ {
            v2 -= 687
            _ = fi8
        }
        switch 32767 {
        case 127:
            var v9 int32 = 2147483647
            _ = v9
        case -401:
            if true {
                for fi10 := 0; fi10 < 7; fi10++ {
                    fmt.Println(v1, v3, v1)
                    v3 = (((v3 + v3) + (v3 + v3)) + "\n\t")
                    var v11 uint16 = 593
                    _ = fi10
                    _ = v11
                }
                v1 += 77
            } else {
                v2 = (v2 << ((65535) & 63))
                switch "\n\t" {
                case "hello":
                    fmt.Println(v3)
                default:
                    fmt.Println(v2)
                }
            }
        case 32768:
            if false {
                v12 := GenS1{F0: true, F1: 1}
                var v13 uint64 = 251
                v3 = "\n\t"
                _ = v12
                _ = v13
            } else {
                v2 -= v2
                v14 := []int{-671, 128, -1}
                v15 := GenS2{F0: 9223372036854775807, F1: 2, F2: uint64(v1)}
                _ = v14
                _ = v15
            }
        default:
            switch ("foo bar" + v3) {
            case "":
                v16 := GenS1{F0: false, F1: 42}
                _ = v16
            }
        }
    } else {
        var v17 int8 = -28
        _ = v17
    }
    switch (("\n\t" + "foo bar") + ("foo bar" + "a")) {
    case "foo bar":
        var v18 float32 = 1e10
        _ = v18
    }
    for fi19 := 0; fi19 < 11; fi19++ {
        var v20 int64 = 0
        for fi21 := 0; fi21 < 4; fi21++ {
            v22 := GenS2{F0: v20, F1: 2147483647, F2: 1}
            v2 = 0
            var v23 int32 = 42
            _ = fi21
            _ = v22
            _ = v23
        }
        v24 := GenS1{F0: ("a" <= v3), F1: -634}
        _ = fi19
        _ = v20
        _ = v24
    }
    v25 := []int{400, 4294967295, -498, 32768}
    fmt.Println(v2, v2, v1)
    for fi26 := 0; fi26 < 1; fi26++ {
        for fi27 := 0; fi27 < 12; fi27++ {
            switch v3 {
            case "foo bar":
                if false {
                    var v28 int32 = -1
                    var v29 uint8 = (v1 - v1)
                    v1 -= (v29 & (v29 >> ((329) & 63)))
                    _ = v28
                    _ = v29
                }
            case "a":
                fmt.Println(v3, v1)
            }
            if ("" >= v3) {
                v30 := GenS1{F0: (float64((float32(3.14159) / 3.14159)) >= 0.1), F1: int(v2)}
                _ = v30
            }
            _ = fi27
        }
        for fi31 := 0; fi31 < 16; fi31++ {
            fmt.Println(fi31, fi26, v3)
            switch "a" {
            case "":
                v32 := GenS3{F0: -443}
                _ = v32
            }
            _ = fi31
        }
        _ = fi26
    }
    var v33 int = 0
    fmt.Println(v33, v33, v1)
    for fi34 := 0; fi34 < 17; fi34++ {
        v3 = "\n\t"
        if true {
            v2 = v2
            v25 = append(v25, (v33 / ((728) | 1)))
            for fi35 := 0; fi35 < 3; fi35++ {
                if true {
                    var v36 float32 = ((-2.25 / 1.0) * 0.1)
                    var v37 float64 = (-2.25 / 1.5)
                    _ = v36
                    _ = v37
                }
                var v38 int64 = 0
                _ = fi35
                _ = v38
            }
        }
        _ = fi34
    }
    fmt.Println(v3, v3, v3)
    var v39 int8 = 0
    var v40 float64 = 1.5
    var v41 string = v3
    for fi42 := 0; fi42 < 16; fi42++ {
        v1 += v1
        var v43 int32 = 1
        _ = fi42
        _ = v43
    }
    v1 -= 127
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v33
    _ = v39
    _ = v40
    _ = v41
}
