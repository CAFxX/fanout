package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 string }
type GenS3 struct { F0 string; F1 int32; F2 string }
func genFn0(p0 uint16, p1 int32, p2 uint) bool {
    if false {
        if (2 < 2147483647) {
            v5 := []int{}
            if len(v5) > 0 {
                v5[0] = 4294967295
            }
            _ = v5
        }
        v6 := []int{902}
        _ = v6
    } else {
        v7 := []int{-971}
        _ = v7
    }
    p1 = 2147483647
    v8 := GenS1{F0: 127}
    var v9 float64 = -2.25
    v10 := GenS1{F0: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v8
    _ = v9
    return (("" + "foo bar") <= "\n\t")
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 float32 = (3.14159 / float32(3.14159))
    v3 := []uint32{}
    var v4 bool = (true || false)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch 9223372036854775807 {
    case -740:
        v2 += -2.25
    case -1:
        _ = len(v3)
    case 241:
        v4 = v4
    default:
        v2 = (v2 - v2)
    }
    v2 += (-2.25 / (((3.14159 - ((1.5 + (v2 / v2)) * (v2 + 0.1))) - -2.25) * v2))
    v11 := GenS1{F0: 255}
    fmt.Println(v4)
    v4 = (327 > 9223372036854775807)
    v1 = (v1 & 18446744073709551615)
    fmt.Println(v2, v1)
    v3 = append(v3, uint32(v1))
    v12 := GenS2{F0: v4, F1: "\n\t"}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v1 -= v1
    if (1 < -1) {
        var v13 uint = 289
        for fi14 := 0; fi14 < 16; fi14++ {
            for fi15 := 0; fi15 < 10; fi15++ {
                var v16 string = "foo bar"
                fmt.Println(v4, fi14, fi14)
                var v17 bool = (true && (2147483647 == -2147483648))
                _ = fi15
                _ = v16
                _ = v17
            }
            switch "a" {
            case "\n\t":
                _ = len(v3)
            case "foo bar":
                for fi18 := 0; fi18 < 12; fi18++ {
                    switch (^(fi18)) {
                    case -1:
                        v2 = (-2.25 + float32(0.0))
                    case 128:
                        var v19 int32 = 2147483647
                        _ = v19
                    }
                    fmt.Println(fi14, v13)
                    v20 := GenS2{F0: (!false), F1: "hello"}
                    _ = fi18
                    _ = v20
                }
            case "hello":
                var v21 uint8 = uint8(v13)
                _ = v21
            default:
                var v22 float64 = 1.5
                _ = v22
            }
            v13 += v13
            _ = fi14
        }
        v23 := GenS3{F0: "a", F1: -2147483648, F2: ("hello" + ("" + ""))}
        _ = v13
        _ = v23
    } else {
        v1 = (v1 / ((18446744073709551615) | 1))
    }
    _ = len(v3)
    switch "hello" {
    case "foo bar":
        if true {
            v24 := GenS3{F0: "\n\t", F1: -2147483648, F2: "foo bar"}
            _ = v24
        } else {
            fmt.Println(v1)
        }
    case "":
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    default:
        var v25 string = "\n\t"
        _ = v25
    }
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
}
