package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint32; F1 uint8 }
type GenS5 struct { F0 uint; F1 int }
func genFn0(p0 uint8, p1 int32) int32 {
    p0 = p0
    v4 := GenS1{F0: 127, F1: 187}
    for fi5 := 0; fi5 < 6; fi5++ {
        v6 := GenS1{F0: -56, F1: 348}
        p1 -= (p1 + p1)
        var v7 int8 = -128
        _ = fi5
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 18; fi8++ {
        for fi9 := 0; fi9 < 12; fi9++ {
            p0 = 42
            _ = fi9
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = v4
    return (p1 >> ((65535) & 63))
}

func main() {
    var v1 float32 = (0.0 * ((0.0 + -2.25) + (1.5 - (0.1 - ((-2.25 * ((0.1 + 3.14159) - ((0.0 / 1.0) / 0.1))) * ((1.5 - 0.1) / 1.0))))))
    var v2 uint8 = 2
    v3 := []uint32{953, 4294967295, 148, 2147483648}
    v1 += 3.14159
    switch "\n\t" {
    case "":
        var v10 bool = (1 == 255)
        _ = v10
    default:
        v11 := GenS1{F0: 0, F1: 0}
        _ = v11
    }
    fmt.Println(v2, v2)
    v3 = append(v3, 42)
    var v12 uint16 = 65535
    for fi13 := 0; fi13 < 16; fi13++ {
        v14 := GenS2{F0: ""}
        fmt.Println(v12, v12, v12)
        if false {
            v2 = 128
            var v15 uint16 = (^(v12))
            var v16 int64 = -9223372036854775808
            _ = v15
            _ = v16
        }
        _ = fi13
        _ = v14
    }
    v2 -= (^(v2))
    for fi17 := 0; fi17 < 5; fi17++ {
        var v18 int = (fi17 & 32767)
        _ = fi17
        _ = v18
    }
    switch "foo bar" {
    case "\n\t":
        fmt.Println(v2, v2, v2)
    default:
        if (2147483647 < 0) {
            for fi19 := 0; fi19 < 16; fi19++ {
                if false {
                    var v20 uint = 2
                    if len(v3) > 0 {
                        v3[0] = 42
                    }
                    fmt.Println(v12, v2)
                    _ = v20
                }
                v21 := GenS3{F0: "a"}
                _ = fi19
                _ = v21
            }
            var v22 float32 = -2.25
            if (!(255 == 18446744073709551615)) {
                var v23 int8 = 100
                if (!true) {
                    _ = len(v3)
                }
                fmt.Println(v23, v23, v23)
                _ = v23
            } else {
                v3 = append(v3, 0)
            }
            _ = v22
        } else {
            for fi24 := 0; fi24 < 13; fi24++ {
                switch ("a" + "foo bar") {
                case "hello":
                    var v25 float64 = 0.0
                    _ = v25
                case "\n\t":
                    v26 := GenS4{F0: 65535, F1: 128}
                    _ = v26
                }
                _ = fi24
            }
            for fi27 := 0; fi27 < 11; fi27++ {
                for fi28 := 0; fi28 < 16; fi28++ {
                    var v29 int64 = 42
                    v12 -= (v12 | 2)
                    _ = fi28
                    _ = v29
                }
                _ = fi27
            }
            for fi30 := 0; fi30 < 9; fi30++ {
                switch fi30 {
                case -1:
                    v31 := GenS5{F0: 2, F1: (fi30 - (fi30 + (^(fi30))))}
                    _ = v31
                case 42:
                    var v32 uint16 = (v12 * 42)
                    _ = v32
                }
                fmt.Println(v1)
                v33 := GenS1{F0: -21, F1: (^(v12))}
                _ = fi30
                _ = v33
            }
        }
    }
    switch 128 {
    case -2147483648:
        v12 += 304
    }
    var v34 float32 = v1
    var v35 int8 = 0
    v34 = v1
    var v36 uint32 = 2
    v36 = v36
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
}
