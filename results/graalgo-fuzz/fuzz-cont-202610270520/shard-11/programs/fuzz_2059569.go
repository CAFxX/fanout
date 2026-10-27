package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int32; F1 uint; F2 uint32 }
type GenS4 struct { F0 bool; F1 int16; F2 int8 }

func main() {
    v1 := []string{"a", "\n\t", ("foo bar" + "foo bar"), "a"}
    var v2 uint32 = 420
    var v3 uint8 = 191
    var v4 int16 = 32767
    v5 := []string{}
    fmt.Println(v4)
    switch "" {
    case "foo bar":
        switch 32767 {
        case 32768:
            fmt.Println(v2, v3, v4)
        default:
            var v6 uint8 = (v3 % ((v3) | 1))
            _ = v6
        }
    default:
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    }
    v7 := GenS1{F0: (18446744073709551615 > 634), F1: 42}
    v8 := GenS2{F0: 32768}
    if false {
        var v9 int32 = -2147483648
        switch "" {
        case "\n\t":
            switch 128 {
            case 42:
                var v10 uint8 = (v3 & 2)
                _ = v10
            }
        case "hello":
            if len(v5) > 0 {
                v5[0] = "a"
            }
        case "a":
            var v11 int8 = -1
            _ = v11
        default:
            if true {
                var v12 float64 = (float64((float32(-2.25) + 1e10)) * 1e10)
                var v13 int = -824
                fmt.Println(v12, v2)
                _ = v12
                _ = v13
            }
        }
        var v14 float32 = float32((1.5 - float64(float32(0.1))))
        _ = v14
        _ = v9
    } else {
        if (v3 < (v3 - (v3 >> ((2) & 63)))) {
            v4 = (v4 + v4)
            v15 := GenS3{F0: 42, F1: 660, F2: v2}
            var v16 float32 = 0.0
            _ = v15
            _ = v16
        } else {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = len(v1)
        }
        v2 -= 350
    }
    v2 += v2
    for fi17 := 0; fi17 < 14; fi17++ {
        switch (fi17 ^ fi17) {
        case 425:
            if len(v5) > 0 {
                v5[0] = ""
            }
        case 32767:
            for fi18 := 0; fi18 < 8; fi18++ {
                v4 = (v4 - v4)
                switch ("hello" + (("" + (("\n\t" + ("" + "hello")) + "\n\t")) + ("a" + ("" + "\n\t")))) {
                case "a":
                    v19 := GenS2{F0: fi17}
                    _ = v19
                case "":
                    v2 = uint32(v3)
                case "hello":
                    fmt.Println(fi17, fi18)
                }
                v2 = v2
                var v20 float32 = 0.0
                _ = fi18
                _ = v20
            }
        }
        _ = fi17
    }
    fmt.Println(v2, v2, v2)
    if ((3.14159 < float64((-2.25 + 1e10))) && (0.0 >= ((0.1 * 0.1) - ((1e10 + (0.0 / 0.1)) + 1e10)))) {
        v21 := GenS3{F0: -1, F1: 4294967295, F2: v2}
        var v22 uint32 = 2147483647
        _ = v21
        _ = v22
    }
    fmt.Println(v3)
    v23 := GenS2{F0: 532}
    if false {
        v24 := GenS3{F0: int32(v3), F1: 18446744073709551615, F2: v2}
        v2 += uint32(v3)
        _ = v24
    }
    v3 += v3
    if ("foo bar" == (("" + ("" + ("foo bar" + ""))) + "\n\t")) {
        var v25 uint16 = 65535
        v2 = 0
        if (false || true) {
            switch 32768 {
            case 849:
                var v26 int32 = 2147483647
                _ = v26
            }
        }
        _ = v25
    }
    v27 := GenS4{F0: true, F1: 1000, F2: 100}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
