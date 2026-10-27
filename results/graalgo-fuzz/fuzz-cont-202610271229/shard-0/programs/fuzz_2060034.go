package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint64; F1 uint8; F2 uint32 }
type GenS5 struct { F0 int8; F1 uint32; F2 int8 }

func main() {
    var v1 uint64 = 771
    var v2 int8 = 83
    var v3 uint8 = 127
    var v4 string = "hello"
    var v5 int16 = 32767
    for fi6 := 0; fi6 < 3; fi6++ {
        switch int(v2) {
        case 0:
            v5 = (v5 * 42)
        case 9223372036854775807:
            if (1 <= int8(fi6)) {
                v7 := []int{}
                if false {
                    v2 = v2
                    v3 -= (v3 % (((v3 << ((42) & 63))) | 1))
                } else {
                    var v8 uint8 = v3
                    var v9 bool = false
                    _ = v8
                    _ = v9
                }
                v3 += v3
                _ = v7
            } else {
                var v10 uint = 65535
                switch v4 {
                case "a":
                    switch v4 {
                    case "foo bar":
                        if true {
                            if false {
                                var v11 uint8 = 128
                                _ = v11
                            }
                            for fi12 := 0; fi12 < 20; fi12++ {
                                if (!(!((1e10 / (((float64((0.0 + ((0.0 - (((-2.25 * 1e10) - float32(-2.25)) - (1.5 / 3.14159))) * float32((((1e10 / 3.14159) + (-2.25 + (3.14159 - ((0.1 + 0.0) - (float64(0.0) * 3.14159))))) / -2.25))))) - 3.14159) * 1.5) + float64(((-2.25 + float32(0.1)) * float32(float64(0.1)))))) == (0.1 + -2.25)))) {
                                    var v13 bool = false
                                    v5 = (v5 + (v5 - int16(v1)))
                                    fmt.Println(v3, v3, v1)
                                    fmt.Println(v5, v1)
                                    fmt.Println(v13)
                                    v5 = v5
                                    v5 = (v5 * int16(v3))
                                    var v14 bool = v13
                                    _ = v13
                                    _ = v14
                                } else {
                                    v15 := GenS1{F0: 926, F1: ("" + v4)}
                                    _ = v15
                                }
                                if (571 <= 65535) {
                                    v16 := []int{fi6}
                                    v4 = "a"
                                    _ = v16
                                } else {
                                    v3 = v3
                                    v5 = (v5 | (v5 ^ v5))
                                }
                                v17 := GenS1{F0: 42, F1: v4}
                                _ = fi12
                                _ = v17
                            }
                            var v18 int16 = (v5 % ((v5) | 1))
                            _ = v18
                        }
                    }
                case "hello":
                    fmt.Println(v1)
                case "\n\t":
                    v19 := GenS1{F0: 65535, F1: (((v4 + "a") + (v4 + (v4 + "\n\t"))) + ((("\n\t" + (v4 + v4)) + (((v4 + v4) + v4) + v4)) + (("a" + (((((v4 + "a") + "a") + ((v4 + v4) + v4)) + "foo bar") + v4)) + (v4 + ""))))}
                    _ = v19
                }
                v20 := GenS2{F0: (("" + ((("" + v4) + "foo bar") + "hello")) + v4)}
                _ = v10
                _ = v20
            }
        case 127:
            v4 = (v4 + (v4 + ""))
        default:
            v5 = 107
        }
        v21 := GenS1{F0: 69, F1: "hello"}
        _ = fi6
        _ = v21
    }
    v22 := []int{-9223372036854775808, 42, 0}
    for fi23 := 0; fi23 < 3; fi23++ {
        if false {
            if ((889 == 0) && ((fi23 % ((fi23) | 1)) >= fi23)) {
                v22 = append(v22, (fi23 - -288))
                var v24 uint = 668
                _ = v24
            }
            fmt.Println(v5)
        }
        var v25 float64 = (-2.25 - float64(1.5))
        _ = fi23
        _ = v25
    }
    switch 0 {
    case 32767:
        for fi26 := 0; fi26 < 15; fi26++ {
            var v27 uint8 = (v3 * v3)
            for fi28 := 0; fi28 < 2; fi28++ {
                var v29 uint = 4294967295
                switch int(v2) {
                case 127:
                    v30 := GenS3{F0: false}
                    _ = v30
                }
                var v31 uint = (v29 / ((v29) | 1))
                v29 = 2
                if len(v22) > 0 {
                    v22[0] = -2147483648
                }
                _ = fi28
                _ = v29
                _ = v31
            }
            var v32 int16 = (v5 << ((uint(v1)) & 63))
            _ = fi26
            _ = v27
            _ = v32
        }
    case -2147483648:
        for fi33 := 0; fi33 < 14; fi33++ {
            if (2147483647 <= int32(v5)) {
                var v34 int64 = 42
                _ = v34
            } else {
                var v35 uint = 255
                _ = v35
            }
            switch ((v4 + (v4 + "a")) + v4) {
            case "hello":
                var v36 int = int(v1)
                _ = v36
            case "a":
                var v37 uint16 = 0
                _ = v37
            case "foo bar":
                switch ((v4 + v4) + (v4 + (v4 + (("foo bar" + v4) + (v4 + "\n\t"))))) {
                case "hello":
                    v4 = ((v4 + ("hello" + v4)) + v4)
                case "\n\t":
                    v22 = append(v22, (fi33 ^ fi33))
                case "a":
                    v38 := GenS4{F0: (v1 | (v1 << ((65535) & 63))), F1: v3, F2: 882}
                    _ = v38
                }
            default:
                if len(v22) >= 2 {
                    v22 = v22[1:len(v22)]
                }
            }
            _ = fi33
        }
    case 32768:
        var v39 int32 = 2147483647
        _ = v39
    }
    v40 := GenS5{F0: v2, F1: 1, F2: v2}
    switch 68 {
    case -9223372036854775808:
        if true {
            for fi41 := 0; fi41 < 19; fi41++ {
                var v42 int = fi41
                _ = fi41
                _ = v42
            }
        } else {
            var v43 uint = 42
            _ = v43
        }
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v40
    _ = v5
}
