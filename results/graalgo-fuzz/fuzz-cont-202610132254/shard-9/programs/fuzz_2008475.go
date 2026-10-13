package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 int32; F1 uint64; F2 int16 }
type GenS3 struct { F0 int8; F1 uint16 }
func genFn0() string {
    var v5 uint16 = 65535
    switch 230 {
    case 9223372036854775807:
        var v6 int16 = 32767
        _ = v6
    case 32768:
        v7 := GenS1{F0: 988, F1: -1}
        _ = v7
    case -370:
        for fi8 := 0; fi8 < 16; fi8++ {
            for fi9 := 0; fi9 < 2; fi9++ {
                switch "foo bar" {
                case "":
                    v10 := []int{}
                    _ = v10
                case "hello":
                    v11 := GenS1{F0: 2, F1: 100}
                    v12 := []int{(fi9 & fi8), fi9}
                    v5 = 17
                    _ = v11
                    _ = v12
                case "foo bar":
                    if (!false) {
                        v5 = (v5 % ((0) | 1))
                        v5 = 1
                    }
                default:
                    if true {
                        v5 = v5
                    }
                }
                _ = fi9
            }
            v5 -= 0
            _ = fi8
        }
    default:
        var v13 int16 = 32767
        _ = v13
    }
    _ = v5
    return (("foo bar" + "a") + "foo bar")
}
func genFn1() uint32 {
    var v14 string = "\n\t"
    v14 = "hello"
    fmt.Println(v14)
    _ = v14
    return 2
}

func main() {
    var v1 int = 32768
    var v2 uint = 366
    var v3 int64 = 9223372036854775807
    var v4 float32 = ((1.5 - 0.0) / 1.0)
    var v15 uint16 = 255
    for fi16 := 0; fi16 < 4; fi16++ {
        v17 := GenS2{F0: -2147483648, F1: 42, F2: 1}
        v18 := []int{(fi16 - (fi16 >> ((298) & 63))), (v1 + fi16), -310}
        _ = fi16
        _ = v17
        _ = v18
    }
    v19 := []int{127, v1}
    v2 -= 65535
    v20 := GenS2{F0: -2147483648, F1: 4294967295, F2: 0}
    var v21 float32 = float32(3.14159)
    if false {
        v15 = uint16(v2)
        v22 := GenS3{F0: 127, F1: (v15 >> ((v2) & 63))}
        switch 32767 {
        case -9223372036854775808:
            if false {
                _ = len(v19)
            }
        default:
            var v23 float64 = 1.5
            _ = v23
        }
        _ = v22
    }
    fmt.Println(v2)
    v21 = (-2.25 * 1e10)
    v4 += v21
    if false {
        var v24 uint = v2
        v25 := GenS1{F0: 1, F1: 42}
        _ = v24
        _ = v25
    } else {
        v4 += ((v21 + ((v4 - float32(3.14159)) - 1e10)) + float32((float64(v4) * 1e10)))
    }
    v26 := GenS1{F0: 42, F1: 0}
    v4 -= 0.0
    v1 -= (v1 | (^(v1)))
    switch (v1 >> (((v2 | (v2 << ((v2) & 63)))) & 63)) {
    case 2147483647:
        var v27 uint8 = 1
        _ = v27
    }
    v3 -= (v3 * (v3 / ((v3) | 1)))
    if len(v19) >= 2 {
        v19 = v19[1:len(v19)]
    }
    var v28 uint = 960
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v26
    _ = v28
    _ = v3
    _ = v4
}
