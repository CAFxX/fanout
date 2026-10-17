package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 int16 }
type GenS2 struct { F0 uint64; F1 int; F2 bool }
type GenS3 struct { F0 uint32; F1 int8 }
type GenS4 struct { F0 int }
func genFn0(p0 int64, p1 uint64) bool {
    var v6 bool = false
    v6 = v6
    var v7 uint32 = 4294967295
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return v6
}
func genFn1() string {
    if false {
        var v8 int16 = 0
        var v9 float32 = (-2.25 * (((1e10 - ((3.14159 - float32(1e10)) / 1.0)) * (float32(((float64((-2.25 + 0.1)) + 3.14159) * ((3.14159 - (1e10 - 1e10)) + 1e10))) / 1e10)) * (3.14159 / -2.25)))
        _ = v8
        _ = v9
    }
    var v10 int = 32768
    _ = v10
    return "\n\t"
}

func main() {
    var v1 bool = ((0.1 == -2.25) && (0 < 255))
    v2 := []uint32{316}
    v3 := make(map[int]bool)
    v4 := make(map[int]uint32)
    var v5 float64 = 3.14159
    var v11 int32 = -1
    var v12 float64 = v5
    v13 := GenS1{F0: 249, F1: 588, F2: 1000}
    v14 := GenS2{F0: 926, F1: -614, F2: v1}
    fmt.Println(v5, v5)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v15 float64 = (((v5 / 1e10) * (float64(((1e10 - (0.0 - 3.14159)) + (1.5 * 3.14159))) - v5)) * v12)
    v16 := GenS2{F0: 1, F1: 9223372036854775807, F2: true}
    for fi17 := 0; fi17 < 14; fi17++ {
        v2 = append(v2, 1)
        if false {
            switch "a" {
            case "\n\t":
                var v18 uint32 = 0
                _ = v18
            }
            switch (-(fi17)) {
            case 127:
                for fi19 := 0; fi19 < 10; fi19++ {
                    switch (^(fi19)) {
                    case 9:
                        switch 704 {
                        case 128:
                            v20 := GenS3{F0: 373, F1: 42}
                            _ = v20
                        case 0:
                            v21 := GenS2{F0: 814, F1: 0, F2: ((v11 - 42) < 1)}
                            _ = v21
                        case -9223372036854775808:
                            v12 = v12
                        default:
                            var v22 int = (^(fi19))
                            _ = v22
                        }
                    case 128:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    switch (fi19 | -318) {
                    case -1:
                        if len(v2) > 0 {
                            v2[0] = 2
                        }
                    default:
                        if (v1 && (!true)) {
                            var v23 uint = 402
                            switch "hello" {
                            case "\n\t":
                                v23 = 0
                            case "hello":
                                v1 = true
                            case "":
                                v24 := v3[(fi17 / (((fi17 / ((524) | 1))) | 1))]
                                _ = v24
                            default:
                                var v25 int64 = 42
                                _ = v25
                            }
                            var v26 float32 = -2.25
                            if ((!((v23 * (v23 % ((v23) | 1))) <= 18446744073709551615)) || (("\n\t" <= "") && v1)) {
                                delete(v3, (fi17 % (((fi19 * fi17)) | 1)))
                            }
                            fmt.Println(v23, fi17)
                            var v27 float32 = v26
                            if (true && (!(false || false))) {
                                fmt.Println(fi17)
                                var v28 uint64 = 0
                                _ = v28
                            }
                            if ("" > (("hello" + ("foo bar" + "")) + "a")) {
                                var v29 int64 = 1
                                _ = v29
                            }
                            _ = v23
                            _ = v26
                            _ = v27
                        }
                    }
                    _ = fi19
                }
            }
        }
        _ = fi17
    }
    if (("hello" != "a") || false) {
        v30 := v4[9223372036854775807]
        switch "hello" {
        case "a":
            var v31 bool = (!((((v30 == v30) || (1 != 42)) || v1) && v1))
            _ = v31
        case "\n\t":
            fmt.Println(v5, v30)
        case "foo bar":
            if v1 {
                var v32 float32 = 3.14159
                v30 += (v30 / ((v30) | 1))
                if v1 {
                    switch (("hello" + "hello") + ("foo bar" + ("a" + (("hello" + "") + "a")))) {
                    case "a":
                        v11 = int32(v30)
                    default:
                        for fi33 := 0; fi33 < 15; fi33++ {
                            if (((^(v30)) != v30) || ((0.1 + 0.1) > v32)) {
                                v11 += (v11 - -800)
                            }
                            for fi34 := 0; fi34 < 17; fi34++ {
                                v4[fi34] = 4294967295
                                _ = fi34
                            }
                            _ = fi33
                        }
                    }
                } else {
                    delete(v3, -35)
                    var v35 int = -278
                    _ = v35
                }
                _ = v32
            } else {
                if ((281 <= 0) || ((-1 != 9223372036854775807) && v1)) {
                    v36 := GenS4{F0: -2147483648}
                    _ = v36
                }
                var v37 float64 = v12
                v15 = 1.5
                _ = v37
            }
        }
        var v38 string = "foo bar"
        _ = v30
        _ = v38
    }
    v39 := v3[-2147483648]
    fmt.Println(v11)
    var v40 int32 = 707
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
}
