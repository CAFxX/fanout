package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 uint32; F1 bool; F2 string }
type GenS3 struct { F0 int; F1 uint8 }

func main() {
    var v1 bool = false
    var v2 uint = 42
    v3 := []uint32{2147483647, 0, 65535}
    var v4 float64 = 1.5
    var v5 string = ""
    v6 := make(map[string]uint32)
    var v7 int8 = -1
    if (v1 || (2 != 2)) {
        if v1 {
            var v8 uint32 = 2147483648
            _ = v8
        }
        for fi9 := 0; fi9 < 6; fi9++ {
            for fi10 := 0; fi10 < 6; fi10++ {
                v11 := GenS1{F0: 0, F1: (v5 + "hello")}
                _ = fi10
                _ = v11
            }
            _ = fi9
        }
    } else {
        var v12 uint16 = 574
        _ = v12
    }
    v13 := GenS1{F0: 291, F1: "hello"}
    v1 = v1
    v14 := GenS2{F0: 1, F1: v1, F2: ("" + "")}
    for fi15 := 0; fi15 < 11; fi15++ {
        if (false && (!v1)) {
            fmt.Println(v2)
        }
        _ = fi15
    }
    fmt.Println(v5)
    for fi16 := 0; fi16 < 13; fi16++ {
        switch v5 {
        case "":
            v6[v5] = 2147483648
        case "a":
            fmt.Println(v5)
        }
        if ("foo bar" <= v5) {
            for fi17 := 0; fi17 < 8; fi17++ {
                if true {
                    var v18 int16 = int16(fi17)
                    if (!true) {
                        v19 := GenS3{F0: (fi17 / ((fi16) | 1)), F1: 255}
                        if (2 >= 1) {
                            v2 = (v2 % ((v2) | 1))
                            var v20 float64 = v4
                            _ = v20
                        } else {
                            if len(v3) > 0 {
                                v3[0] = 0
                            }
                        }
                        _ = v19
                    }
                    _ = v18
                }
                fmt.Println(fi16, fi16, v5)
                v4 -= float64(1.5)
                var v21 int64 = 42
                _ = fi17
                _ = v21
            }
        } else {
            switch fi16 {
            case -2147483648:
                v5 = v5
            case 128:
                fmt.Println(fi16)
            default:
                switch v5 {
                case "\n\t":
                    for fi22 := 0; fi22 < 18; fi22++ {
                        if v1 {
                            if (-1 <= (fi16 / (((fi22 & fi22)) | 1))) {
                                var v23 bool = (((v4 / ((-2.25 * float64(-2.25)) / v4)) * float64(0.1)) != float64(3.14159))
                                fmt.Println(fi16, v23, v1)
                                _ = v23
                            } else {
                                var v24 bool = (!(!v1))
                                _ = v24
                            }
                        }
                        if v1 {
                            switch "\n\t" {
                            case "a":
                                var v25 uint = (v2 << (((v2 * (^(v2)))) & 63))
                                _ = v25
                            case "\n\t":
                                v26 := GenS3{F0: (fi16 & fi16), F1: 128}
                                _ = v26
                            case "hello":
                                var v27 uint = (v2 + (v2 % ((v2) | 1)))
                                _ = v27
                            default:
                                v4 -= 3.14159
                            }
                            var v28 uint = v2
                            switch (fi22 | (fi16 & fi22)) {
                            case -695:
                                var v29 float64 = ((((3.14159 / ((((v4 * (0.1 * v4)) + ((float64(3.14159) * float64(3.14159)) + v4)) + v4) - v4)) - v4) * v4) * v4)
                                _ = v29
                            }
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            fmt.Println(v4)
                            v6["hello"] = 76
                            _ = v28
                        }
                        _ = fi22
                    }
                default:
                    if (v1 || v1) {
                        fmt.Println(v7, v1)
                        for fi30 := 0; fi30 < 8; fi30++ {
                            if len(v3) > 0 {
                                v3[0] = 2147483647
                            }
                            var v31 int = -1
                            _ = fi30
                            _ = v31
                        }
                        if v1 {
                            v4 -= float64((0.1 + (1.5 * ((1e10 + (1.5 * (1e10 * (-2.25 - 0.0)))) * (float32((v4 - float64(3.14159))) - float32(v4))))))
                            var v32 string = v5
                            _ = v32
                        } else {
                            var v33 uint16 = 32767
                            fmt.Println(v1, v1, fi16)
                            _ = v33
                        }
                    }
                }
            }
            if (-128 <= (^(v7))) {
                v34 := GenS1{F0: 64, F1: ((((("hello" + ((v5 + v5) + v5)) + (("" + v5) + v5)) + v5) + v5) + v5)}
                _ = v34
            } else {
                if (2 >= 0) {
                    v1 = true
                    fmt.Println(v4, fi16)
                }
                fmt.Println(fi16, v5, v1)
            }
        }
        var v35 int16 = 32767
        _ = fi16
        _ = v35
    }
    var v36 uint64 = 1
    v2 += v2
    v37 := GenS2{F0: 71, F1: v1, F2: "\n\t"}
    v5 = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
