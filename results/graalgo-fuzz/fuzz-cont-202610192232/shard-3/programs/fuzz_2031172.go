package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16 }
func genFn0(p0 uint32) int64 {
    var v5 string = (("hello" + (("" + ("a" + "hello")) + ("hello" + ("hello" + "\n\t")))) + "a")
    var v6 bool = ((((1e10 + float32((float64(1.5) / 1.0))) * (1e10 + 3.14159)) != (-2.25 - 1e10)) && (4294967295 <= 968))
    _ = p0
    _ = v5
    _ = v6
    return -1
}
func genFn1(p0 int8, p1 string, p2 int8) bool {
    fmt.Println(p1, p1)
    p2 += -128
    switch 128 {
    case 1:
        var v7 uint64 = 238
        _ = v7
    case 4294967295:
        for fi8 := 0; fi8 < 3; fi8++ {
            v9 := []int{fi8}
            p2 = p0
            var v10 uint8 = 79
            _ = fi8
            _ = v10
            _ = v9
        }
    default:
        p0 = -26
    }
    var v11 uint32 = 4294967295
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return true
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{(("a" + "a") + "\n\t")}
    v3 := []uint32{962, 398, 243, 4294967295}
    var v4 uint64 = 4294967295
    for fi12 := 0; fi12 < 13; fi12++ {
        var v13 uint8 = 223
        _ = fi12
        _ = v13
    }
    v14 := GenS1{F0: 32767, F1: 128}
    for fi15 := 0; fi15 < 3; fi15++ {
        for fi16 := 0; fi16 < 10; fi16++ {
            v17 := GenS1{F0: 42, F1: 127}
            _ = fi16
            _ = v17
        }
        _ = fi15
    }
    v18 := GenS2{F0: int8(v4)}
    v19 := GenS2{F0: -128}
    for fi20 := 0; fi20 < 13; fi20++ {
        v4 -= (v4 * 4294967295)
        if (("foo bar" != "foo bar") || (false || (false && false))) {
            fmt.Println(fi20, fi20)
            v4 = (v4 << ((855) & 63))
        } else {
            v4 -= (^(v4))
            for fi21 := 0; fi21 < 15; fi21++ {
                switch "\n\t" {
                case "hello":
                    v22 := GenS3{F0: 1}
                    _ = v22
                }
                _ = fi21
            }
        }
        _ = fi20
    }
    for fi23 := 0; fi23 < 9; fi23++ {
        if len(v2) > 0 {
            v2[0] = (((("\n\t" + "foo bar") + "") + "a") + "\n\t")
        }
        var v24 int32 = 974
        if true {
            if false {
                v4 += v4
            }
            var v25 uint32 = 65535
            _ = v25
        }
        _ = fi23
        _ = v24
    }
    for fi26 := 0; fi26 < 14; fi26++ {
        switch "hello" {
        case "\n\t":
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        case "hello":
            fmt.Println(fi26)
        case "foo bar":
            v27 := GenS1{F0: -1, F1: 42}
            _ = v27
        default:
            if false {
                v28 := GenS3{F0: -664}
                _ = v28
            } else {
                var v29 float32 = float32(0.0)
                for fi30 := 0; fi30 < 3; fi30++ {
                    v1[93] = (!false)
                    _ = fi30
                }
                _ = v29
            }
        }
        _ = fi26
    }
    if (2147483647 <= 42) {
        switch "a" {
        case "hello":
            var v31 int8 = 42
            _ = v31
        case "\n\t":
            var v32 uint32 = 42
            _ = v32
        default:
            v33 := GenS3{F0: -32768}
            _ = v33
        }
        var v34 float64 = 0.1
        _ = v34
    }
    var v35 float32 = 0.0
    var v36 int32 = -174
    _ = len(v1)
    _ = len(v2)
    switch "foo bar" {
    case "foo bar":
        v37 := GenS2{F0: 60}
        _ = v37
    default:
        var v38 string = ("hello" + (("foo bar" + "") + ""))
        _ = v38
    }
    v35 += ((float32((((3.14159 + -2.25) + float64(v35)) / 1.0)) + (v35 - (1e10 * v35))) + v35)
    v36 += v36
    var v39 float64 = (float64(v35) * (float64(v35) / -2.25))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v39
    _ = v4
}
