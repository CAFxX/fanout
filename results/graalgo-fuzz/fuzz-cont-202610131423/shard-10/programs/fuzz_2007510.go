package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 uint32; F2 int32 }
type GenS3 struct { F0 uint64; F1 int32 }

func main() {
    var v1 uint64 = 0
    var v2 string = ("foo bar" + "\n\t")
    var v3 int64 = -1
    var v4 uint = 4294967295
    if true {
        for fi5 := 0; fi5 < 6; fi5++ {
            var v6 bool = false
            v7 := GenS1{F0: -576}
            v8 := []int{1, fi5, fi5}
            _ = fi5
            _ = v6
            _ = v7
            _ = v8
        }
        var v9 int64 = (v3 ^ -501)
        switch (((("\n\t" + (v2 + (v2 + ((("" + v2) + ((v2 + "") + v2)) + v2)))) + (v2 + v2)) + v2) + (v2 + "hello")) {
        case "\n\t":
            v9 -= v3
        case "foo bar":
            v3 += -715
        }
        _ = v9
    } else {
        if (!(false || false)) {
            v10 := GenS1{F0: 409}
            _ = v10
        } else {
            var v11 uint32 = 42
            var v12 int16 = 1
            _ = v11
            _ = v12
        }
        v4 = 4294967295
        var v13 uint16 = 1
        _ = v13
    }
    if ((v3 % ((9223372036854775807) | 1)) < v3) {
        v2 = ""
        v1 = v1
        var v14 uint8 = 255
        _ = v14
    } else {
        for fi15 := 0; fi15 < 9; fi15++ {
            var v16 float64 = float64(3.14159)
            _ = fi15
            _ = v16
        }
        if (128 <= -684) {
            for fi17 := 0; fi17 < 3; fi17++ {
                if (-128 >= 0) {
                    v4 = (v4 / ((v4) | 1))
                } else {
                    var v18 string = v2
                    for fi19 := 0; fi19 < 10; fi19++ {
                        var v20 uint64 = v1
                        _ = fi19
                        _ = v20
                    }
                    _ = v18
                }
                if (2 <= 267) {
                    v21 := []int{1, (fi17 & fi17), -9223372036854775808}
                    v22 := GenS1{F0: 0}
                    _ = v21
                    _ = v22
                }
                fmt.Println(fi17, fi17)
                _ = fi17
            }
        } else {
            var v23 int16 = 42
            v4 -= v4
            _ = v23
        }
        switch 2147483647 {
        case 32768:
            if ((0.0 / (-2.25 * (float32((((3.14159 / ((1e10 / 1e10) - float64(1.5))) + (3.14159 + float64(((0.0 - ((((0.1 / 0.1) + (1e10 * 3.14159)) + 1e10) - 3.14159)) - float32(0.1))))) / 1.0)) - (0.0 * (0.1 + (3.14159 - ((3.14159 + 3.14159) * 3.14159))))))) > ((0.1 * (1e10 - -2.25)) - (3.14159 / -2.25))) {
                v24 := GenS2{F0: 42, F1: uint32(v1), F2: 0}
                v3 = v3
                var v25 float64 = (float64((0.1 - (0.0 - (0.1 / 1.0)))) + (3.14159 + 0.0))
                _ = v24
                _ = v25
            } else {
                for fi26 := 0; fi26 < 14; fi26++ {
                    var v27 float64 = 0.0
                    var v28 int8 = 1
                    _ = fi26
                    _ = v27
                    _ = v28
                }
            }
        case 32767:
            var v29 float64 = (0.1 - 1.5)
            _ = v29
        case 42:
            v2 = (v2 + "\n\t")
        default:
            var v30 int32 = 42
            _ = v30
        }
    }
    v4 += v4
    var v31 uint8 = 1
    v2 = v2
    v3 = (v3 ^ v3)
    var v32 string = ("" + v2)
    switch -435 {
    case -1:
        fmt.Println(v4, v32)
    case 318:
        var v33 uint8 = (v31 ^ (v31 >> ((v4) & 63)))
        _ = v33
    case 0:
        for fi34 := 0; fi34 < 20; fi34++ {
            fmt.Println(v1, v3, v2)
            v35 := []int{}
            if ((fi34 / (((fi34 * fi34)) | 1)) > -2147483648) {
                fmt.Println(v3)
                v36 := GenS3{F0: (v1 / (((v1 * (v1 + v1))) | 1)), F1: 2147483647}
                var v37 string = (("foo bar" + v32) + v2)
                _ = v36
                _ = v37
            }
            _ = fi34
            _ = v35
        }
    default:
        switch -1 {
        case 32768:
            fmt.Println(v2, v3)
        case 1:
            v38 := []int{}
            _ = v38
        }
    }
    var v39 uint8 = (v31 | 128)
    v3 = -1
    v40 := []int{}
    v39 = (v31 - (v31 * v31))
    fmt.Println(v4)
    v41 := GenS3{F0: (^(v1)), F1: -2147483648}
    v42 := GenS3{F0: 4294967295, F1: -990}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v40:", len(v40), v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
}
