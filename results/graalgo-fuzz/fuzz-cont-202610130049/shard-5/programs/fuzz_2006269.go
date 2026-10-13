package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64; F1 int8; F2 uint }
type GenS3 struct { F0 uint; F1 uint64 }

func main() {
    var v1 uint64 = 755
    var v2 int64 = 42
    var v3 uint16 = 255
    var v4 uint64 = (^(v1))
    for fi5 := 0; fi5 < 5; fi5++ {
        v6 := GenS1{F0: "\n\t"}
        v7 := GenS1{F0: "hello"}
        _ = fi5
        _ = v6
        _ = v7
    }
    v1 -= (^(v1))
    var v8 int16 = 901
    for fi9 := 0; fi9 < 3; fi9++ {
        v1 -= 810
        _ = fi9
    }
    for fi10 := 0; fi10 < 16; fi10++ {
        var v11 uint16 = v3
        _ = fi10
        _ = v11
    }
    v12 := GenS1{F0: ("" + "a")}
    v13 := []int{}
    v4 = v4
    switch "foo bar" {
    case "\n\t":
        v14 := GenS2{F0: v2, F1: 1, F2: 1}
        _ = v14
    case "hello":
        if (0 == 783) {
            var v15 uint16 = (v3 >> ((42) & 63))
            _ = v15
        }
    }
    switch 32767 {
    case 32768:
        if (18446744073709551615 <= 42) {
            switch "foo bar" {
            case "foo bar":
                var v16 uint8 = 42
                _ = v16
            case "\n\t":
                fmt.Println(v8, v8, v4)
            }
            v17 := GenS3{F0: 42, F1: (v1 - v4)}
            if len(v13) > 0 {
                v13[0] = 332
            }
            _ = v17
        }
    case -2147483648:
        switch 831 {
        case 0:
            v18 := GenS2{F0: v2, F1: -87, F2: 255}
            _ = v18
        case 9223372036854775807:
            v1 = v4
        default:
            _ = len(v13)
        }
    case -976:
        if len(v13) > 0 {
            v13[0] = -1
        }
    }
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    if (!true) {
        _ = len(v13)
        var v19 string = "foo bar"
        _ = v19
    } else {
        if len(v13) > 0 {
            v13[0] = 128
        }
        var v20 int8 = 127
        fmt.Println(v20, v1, v4)
        _ = v20
    }
    if (-738 > -9223372036854775808) {
        switch (("\n\t" + (("" + "") + "\n\t")) + "a") {
        case "":
            for fi21 := 0; fi21 < 17; fi21++ {
                v3 = (v3 & (v3 * v3))
                v1 += v4
                var v22 int64 = 54
                _ = fi21
                _ = v22
            }
        case "\n\t":
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
        default:
            var v23 bool = true
            _ = v23
        }
    }
    v3 = 7
    if false {
        var v24 float32 = float32((-2.25 * (0.0 - float64(((1e10 / 1e10) + ((3.14159 * (float32(((-2.25 * (((((1e10 - (float64(0.1) / -2.25)) * (1e10 - 0.1)) / 1e10) + -2.25) - (0.1 * float64(1.5)))) * 0.0)) / (float32(-2.25) * (-2.25 - 3.14159)))) - 1e10))))))
        _ = v24
    } else {
        if true {
            for fi25 := 0; fi25 < 9; fi25++ {
                for fi26 := 0; fi26 < 5; fi26++ {
                    v4 = 343
                    v27 := GenS3{F0: 18446744073709551615, F1: 2}
                    _ = fi26
                    _ = v27
                }
                fmt.Println(v3, v3, v8)
                switch (^(fi25)) {
                case -893:
                    fmt.Println(fi25, fi25, v3)
                default:
                    v8 += 42
                }
                _ = fi25
            }
        }
    }
    v3 -= (^(v3))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v8
}
