package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 int16 }
func genFn0(p0 int8) int64 {
    for fi6 := 0; fi6 < 16; fi6++ {
        v7 := GenS1{F0: 2147483648, F1: ("foo bar" + "hello")}
        _ = fi6
        _ = v7
    }
    fmt.Println(p0)
    p0 = 42
    _ = p0
    return -400
}

func main() {
    var v1 uint = 18446744073709551615
    v2 := make(map[uint32]string)
    var v3 uint32 = 42
    var v4 uint64 = 42
    v5 := make(map[string]string)
    var v8 int8 = 100
    v9 := GenS1{F0: (v3 / ((v3) | 1)), F1: ("foo bar" + "a")}
    v10 := GenS1{F0: (v3 - (v3 << ((18446744073709551615) & 63))), F1: "\n\t"}
    switch "foo bar" {
    case "\n\t":
        v4 = v4
    default:
        var v11 uint8 = 1
        _ = v11
    }
    if (-665 == 42) {
        fmt.Println(v1)
    }
    if ((v4 <= (^(v4))) && false) {
        v12 := GenS1{F0: (v3 << (((v1 & v1)) & 63)), F1: ""}
        switch (("\n\t" + "hello") + "") {
        case "foo bar":
            if (!(!(1 <= 255))) {
                v13 := GenS2{F0: 0, F1: 254}
                v3 += (^(v3))
                v8 -= 115
                _ = v13
            } else {
                v14 := GenS1{F0: (v3 | 670), F1: "hello"}
                var v15 float32 = float32(float64(1e10))
                switch "foo bar" {
                case "foo bar":
                    for fi16 := 0; fi16 < 18; fi16++ {
                        fmt.Println(v15)
                        v4 -= 0
                        _ = fi16
                    }
                default:
                    var v17 uint8 = 66
                    _ = v17
                }
                _ = v14
                _ = v15
            }
        case "hello":
            var v18 float32 = 0.1
            _ = v18
        }
        for fi19 := 0; fi19 < 8; fi19++ {
            var v20 int32 = 42
            var v21 int8 = (v8 & v8)
            var v22 uint8 = 0
            _ = fi19
            _ = v20
            _ = v21
            _ = v22
        }
        _ = v12
    }
    var v23 uint32 = v3
    v8 -= v8
    switch -2147483648 {
    case 9223372036854775807:
        fmt.Println(v8, v8, v4)
    case 4294967295:
        v24 := v5["a"]
        _ = v24
    }
    var v25 uint16 = 739
    for fi26 := 0; fi26 < 15; fi26++ {
        if false {
            v8 += 100
            v27 := GenS2{F0: int32(v3), F1: uint8(v4)}
            switch 0 {
            case 4294967295:
                for fi28 := 0; fi28 < 9; fi28++ {
                    var v29 uint16 = (v25 + v25)
                    var v30 uint8 = 0
                    var v31 float32 = ((0.0 / 0.1) + (1.5 / 1e10))
                    _ = fi28
                    _ = v29
                    _ = v30
                    _ = v31
                }
            default:
                var v32 uint16 = (v25 % (((v25 + v25)) | 1))
                _ = v32
            }
            _ = v27
        } else {
            for fi33 := 0; fi33 < 6; fi33++ {
                v34 := GenS1{F0: 795, F1: ("foo bar" + "")}
                _ = fi33
                _ = v34
            }
        }
        var v35 int32 = 2147483647
        _ = fi26
        _ = v35
    }
    var v36 int64 = 9223372036854775807
    switch 9223372036854775807 {
    case -2147483648:
        v8 += 42
    case 842:
        v37 := GenS3{F0: 32767}
        _ = v37
    }
    if false {
        var v38 float32 = (((-2.25 / (1.5 * 3.14159)) * 0.1) - 1e10)
        _ = v38
    }
    fmt.Println(v25, v25, v4)
    var v39 uint16 = (v25 % ((uint16(v4)) | 1))
    var v40 uint32 = 2
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v23
    _ = v25
    _ = v3
    _ = v36
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v8
    _ = v9
}
