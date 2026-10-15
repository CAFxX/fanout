package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 uint64; F1 string; F2 int32 }
func genFn0(p0 int64, p1 uint32) bool {
    p0 -= p0
    v5 := []int{}
    _ = p0
    _ = p1
    _ = v5
    return true
}

func main() {
    var v1 uint8 = 255
    v2 := []int{565, -2147483648, 9223372036854775807, -429}
    var v3 string = (("hello" + "") + "\n\t")
    var v4 int8 = 0
    v6 := GenS1{F0: "foo bar", F1: -115}
    switch v3 {
    case "foo bar":
        v4 += -1
    default:
        var v7 bool = (float64(0.1) <= ((1.5 * 3.14159) + 0.0))
        _ = v7
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v8 int8 = v4
    switch v3 {
    case "":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case "foo bar":
        var v9 string = (v3 + "foo bar")
        _ = v9
    case "hello":
        switch "\n\t" {
        case "\n\t":
            switch 0 {
            case 42:
                v10 := GenS2{F0: 2147483647, F1: v1}
                _ = v10
            case -342:
                var v11 int16 = 42
                _ = v11
            case -179:
                var v12 uint = 0
                _ = v12
            default:
                if len(v2) > 0 {
                    v2[0] = 32768
                }
            }
        default:
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
    default:
        switch v3 {
        case "\n\t":
            var v13 float32 = (1e10 + (-2.25 * ((0.0 + 1e10) / -2.25)))
            _ = v13
        case "":
            v2 = append(v2, 32768)
        default:
            v3 = "\n\t"
        }
    }
    for fi14 := 0; fi14 < 4; fi14++ {
        switch 2147483647 {
        case 9223372036854775807:
            v15 := GenS3{F0: 569, F1: (v3 + (((v3 + "a") + v3) + ("" + (("" + "\n\t") + (v3 + (v3 + "a")))))), F2: 0}
            _ = v15
        case -9223372036854775808:
            fmt.Println(v3)
        case -803:
            fmt.Println(v8)
        default:
            for fi16 := 0; fi16 < 19; fi16++ {
                _ = len(v2)
                _ = len(v2)
                _ = fi16
            }
        }
        for fi17 := 0; fi17 < 20; fi17++ {
            v8 -= (^(v8))
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v18 uint32 = 4294967295
            _ = fi17
            _ = v18
        }
        v3 = "hello"
        _ = fi14
    }
    fmt.Println(v4, v1, v8)
    for fi19 := 0; fi19 < 5; fi19++ {
        for fi20 := 0; fi20 < 15; fi20++ {
            var v21 int16 = 32767
            if true {
                if false {
                    var v22 float32 = 0.1
                    _ = v22
                }
            }
            _ = fi20
            _ = v21
        }
        fmt.Println(v4, v4, v4)
        _ = fi19
    }
    var v23 float32 = 0.0
    fmt.Println(v1, v3, v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v6
    _ = v8
}
