package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 int8; F1 uint }
func genFn0(p0 uint16, p1 uint16) int64 {
    p1 = p0
    switch ("\n\t" + "a") {
    case "\n\t":
        var v5 int = 459
        _ = v5
    default:
        var v6 string = ("" + "a")
        _ = v6
    }
    _ = p0
    _ = p1
    return 42
}
func genFn1(p0 bool) int64 {
    fmt.Println(p0)
    var v7 uint8 = 42
    _ = p0
    _ = v7
    return 421
}

func main() {
    v1 := []int{143}
    var v2 int16 = 0
    var v3 float32 = 0.0
    var v4 float64 = (3.14159 * 0.1)
    var v8 uint8 = 255
    v8 = v8
    if (4294967295 == 42) {
        if (v3 >= 0.1) {
            var v9 int = 9223372036854775807
            _ = v9
        }
        var v10 int64 = 292
        _ = v10
    }
    v3 = v3
    v11 := GenS1{F0: (1 >= (v8 + v8)), F1: 255}
    fmt.Println(v3)
    v2 -= v2
    for fi12 := 0; fi12 < 18; fi12++ {
        if (v8 <= (v8 & 172)) {
            var v13 uint8 = 255
            var v14 uint64 = 1
            if (814 <= 2) {
                switch "hello" {
                case "hello":
                    v15 := GenS1{F0: (v4 > float64((v3 / (v3 * (v3 / float32(1.5)))))), F1: 161}
                    _ = v15
                case "":
                    for fi16 := 0; fi16 < 17; fi16++ {
                        fmt.Println(v13)
                        _ = fi16
                    }
                }
            } else {
                for fi17 := 0; fi17 < 11; fi17++ {
                    fmt.Println(v3, fi12)
                    _ = fi17
                }
                v1 = append(v1, fi12)
                v8 += v13
            }
            _ = v13
            _ = v14
        } else {
            v18 := GenS1{F0: ((v3 / v3) <= 1e10), F1: 32767}
            v8 -= 0
            _ = v18
        }
        switch (fi12 / ((fi12) | 1)) {
        case -809:
            v1 = append(v1, -2147483648)
        case 2147483647:
            if false {
                v19 := GenS2{F0: ("hello" + "a"), F1: v2}
                _ = v19
            }
        case 9223372036854775807:
            v2 -= 1000
        }
        var v20 int32 = -2147483648
        _ = fi12
        _ = v20
    }
    var v21 uint16 = 65535
    fmt.Println(v3)
    v22 := GenS3{F0: 0, F1: 709}
    var v23 int64 = genFn1(true)
    var v24 int64 = 9223372036854775807
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v8
}
