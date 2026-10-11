package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 int64; F1 int64 }
type GenS3 struct { F0 int8; F1 uint64 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 string; F1 int16 }

func main() {
    var v1 uint8 = 42
    var v2 bool = false
    v3 := []uint32{}
    var v4 string = ("" + "a")
    v5 := make(map[int]uint32)
    for fi6 := 0; fi6 < 12; fi6++ {
        for fi7 := 0; fi7 < 10; fi7++ {
            var v8 int = (fi7 ^ -1)
            v8 += 45
            _ = fi7
            _ = v8
        }
        fmt.Println(v1)
        _ = fi6
    }
    v9 := GenS1{F0: 2, F1: 39}
    if v2 {
        v10 := GenS1{F0: 2, F1: v1}
        var v11 uint32 = 42
        fmt.Println(v1, v2)
        _ = v10
        _ = v11
    }
    var v12 float64 = 3.14159
    fmt.Println(v12, v4)
    for fi13 := 0; fi13 < 14; fi13++ {
        var v14 int64 = 42
        switch (-(fi13)) {
        case 128:
            v14 = v14
        case 42:
            if ((v4 == (v4 + "")) || (v2 && ((((-2.25 * float32(v12)) * ((1.5 * 0.0) / 3.14159)) - float32((float64((float32((float64(-2.25) * (1.5 / ((v12 + ((-2.25 / (v12 * (3.14159 * ((((v12 * 1.5) / float64(float32((v12 * (v12 / 1.0))))) / (float64(1.5) * v12)) - v12)))) * ((3.14159 / v12) / v12))) * v12)))) + 0.0)) - (v12 / 1e10)))) >= (1e10 * 0.0)))) {
                v15 := GenS1{F0: 2, F1: v1}
                _ = v15
            } else {
                switch (v4 + (((("" + "") + "\n\t") + ((v4 + ((v4 + (v4 + "a")) + (v4 + v4))) + ("a" + ((v4 + (v4 + (v4 + v4))) + v4)))) + ("a" + v4))) {
                case "\n\t":
                    if (v2 || v2) {
                        switch fi13 {
                        case -1:
                            _ = len(v3)
                        case 127:
                            delete(v5, fi13)
                        case 9223372036854775807:
                            v16 := v5[(fi13 ^ fi13)]
                            _ = v16
                        default:
                            if (42 != uint(v1)) {
                                v17 := GenS2{F0: (v14 - v14), F1: 0}
                                v14 = v14
                                if len(v3) > 0 {
                                    v3[0] = 1
                                }
                                _ = v17
                            } else {
                                v18 := GenS1{F0: 42, F1: v1}
                                var v19 int16 = 32767
                                if v2 {
                                    var v20 string = (("" + v4) + v4)
                                    if true {
                                        fmt.Println(v12, v1)
                                        var v21 uint8 = v1
                                        _ = v21
                                    } else {
                                        if (v2 || v2) {
                                            v3 = append(v3, 65535)
                                        } else {
                                            fmt.Println(v19)
                                            delete(v5, (-(fi13)))
                                            var v22 uint = 255
                                            _ = v22
                                        }
                                        v3 = append(v3, 2147483647)
                                        for fi23 := 0; fi23 < 11; fi23++ {
                                            fmt.Println(v14)
                                            var v24 int = fi23
                                            _ = fi23
                                            _ = v24
                                        }
                                    }
                                    fmt.Println(v12)
                                    _ = v20
                                }
                                _ = v18
                                _ = v19
                            }
                        }
                        v25 := GenS3{F0: int8(v1), F1: 1}
                        for fi26 := 0; fi26 < 3; fi26++ {
                            fmt.Println(v12)
                            _ = fi26
                        }
                        _ = v25
                    }
                case "hello":
                    var v27 int32 = int32(v1)
                    _ = v27
                }
                var v28 uint32 = 259
                _ = v28
            }
        }
        var v29 float32 = (float32(0.0) / 1.5)
        _ = fi13
        _ = v14
        _ = v29
    }
    fmt.Println(v4)
    switch ("a" + v4) {
    case "foo bar":
        if (!(2 != 1)) {
            var v30 uint8 = (v1 % ((v1) | 1))
            _ = v30
        } else {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            fmt.Println(v12)
        }
    case "":
        switch ("hello" + "\n\t") {
        case "a":
            for fi31 := 0; fi31 < 10; fi31++ {
                if ((3.14159 - 1.5) >= v12) {
                    if ((9223372036854775807 != -9223372036854775808) || (42 != 46)) {
                        v32 := GenS3{F0: 111, F1: 18446744073709551615}
                        _ = v32
                    } else {
                        if true {
                            var v33 uint64 = 4294967295
                            fmt.Println(v1, v1, fi31)
                            v3 = append(v3, 0)
                            _ = v33
                        }
                        v2 = false
                        switch (fi31 << ((572) & 63)) {
                        case 32767:
                            fmt.Println(v2)
                        default:
                            if ("foo bar" >= v4) {
                                for fi34 := 0; fi34 < 3; fi34++ {
                                    if (v2 || v2) {
                                        v3 = append(v3, 2)
                                        v12 = 0.1
                                        v5[(fi31 + fi34)] = 917
                                    }
                                    v35 := GenS4{F0: 42}
                                    var v36 int16 = 0
                                    _ = fi34
                                    _ = v35
                                    _ = v36
                                }
                            } else {
                                v12 -= ((v12 + ((0.0 - (v12 / v12)) - v12)) * 0.0)
                            }
                        }
                    }
                    var v37 string = v4
                    _ = v37
                }
                v5[fi31] = 679
                v4 = v4
                _ = fi31
            }
        case "\n\t":
            fmt.Println(v12)
        }
    }
    delete(v5, -2147483648)
    v4 = "hello"
    for fi38 := 0; fi38 < 7; fi38++ {
        if v2 {
            if v2 {
                var v39 float32 = (0.1 - 0.0)
                switch v4 {
                case "\n\t":
                    var v40 string = (("hello" + v4) + "hello")
                    _ = v40
                default:
                    v41 := GenS5{F0: ("hello" + (v4 + v4)), F1: 1}
                    _ = v41
                }
                var v42 uint64 = 4294967295
                _ = v39
                _ = v42
            }
            fmt.Println(v2)
        }
        _ = fi38
    }
    switch -2147483648 {
    case 712:
        v12 = (v12 - -2.25)
    }
    var v43 int64 = 875
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v9
}
