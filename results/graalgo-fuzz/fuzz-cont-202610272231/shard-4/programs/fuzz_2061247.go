package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 bool; F1 int32; F2 int32 }
type GenS3 struct { F0 uint32; F1 uint64; F2 bool }
type GenS4 struct { F0 int8; F1 int64; F2 int }

func main() {
    v1 := make(map[string]int)
    var v2 uint64 = 42
    var v3 int = -2147483648
    v4 := make(map[uint32]uint32)
    var v5 uint16 = 42
    var v6 bool = false
    v7 := []int{(v3 / ((32767) | 1)), (v3 * (v3 * -9223372036854775808)), v3}
    var v8 int8 = 24
    switch "foo bar" {
    case "\n\t":
        v9 := GenS1{F0: 42, F1: 2147483647}
        _ = v9
    case "foo bar":
        switch "hello" {
        case "hello":
            v10 := GenS1{F0: 1, F1: -1}
            _ = v10
        default:
            fmt.Println(v5, v2, v3)
        }
    case "":
        v6 = true
    default:
        var v11 int8 = 45
        _ = v11
    }
    var v12 uint8 = 127
    v12 -= 1
    fmt.Println(v6, v2)
    for fi13 := 0; fi13 < 11; fi13++ {
        fmt.Println(v5, fi13, v8)
        _ = fi13
    }
    fmt.Println(v2, v8)
    var v14 int64 = 1
    if (32767 <= 32767) {
        for fi15 := 0; fi15 < 12; fi15++ {
            v14 = v14
            if len(v7) > 0 {
                v7[0] = (fi15 & (fi15 / ((fi15) | 1)))
            }
            v5 -= 255
            _ = fi15
        }
        v16 := GenS2{F0: v6, F1: 0, F2: 42}
        v12 -= (v12 << ((255) & 63))
        _ = v16
    } else {
        v7 = append(v7, (v3 << ((687) & 63)))
        if (859 == -32768) {
            if false {
                switch (v3 / ((int(v14)) | 1)) {
                case 2147483647:
                    var v17 uint64 = v2
                    _ = v17
                default:
                    v3 = 280
                }
                v18 := v1[("\n\t" + "\n\t")]
                _ = v18
            } else {
                v2 = (v2 / (((v2 << ((65535) & 63))) | 1))
                if (v14 <= (v14 - (v14 << ((529) & 63)))) {
                    for fi19 := 0; fi19 < 3; fi19++ {
                        var v20 int32 = 359
                        fmt.Println(v8)
                        var v21 uint32 = 4294967295
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                    switch ("hello" + "hello") {
                    case "a":
                        delete(v4, 65535)
                    case "foo bar":
                        if (3.14159 <= 3.14159) {
                            fmt.Println(v12, v14)
                            if v6 {
                                if (0 <= 460) {
                                    var v22 int64 = (v14 + int64(v3))
                                    var v23 uint = 4294967295
                                    _ = v22
                                    _ = v23
                                } else {
                                    for fi24 := 0; fi24 < 8; fi24++ {
                                        fmt.Println(v8, v8)
                                        v3 = fi24
                                        if (4294967295 >= (^(v2))) {
                                            if (v6 || ((-(v8)) != (v8 + (^(v8))))) {
                                                if len(v7) >= 2 {
                                                    v7 = v7[1:len(v7)]
                                                }
                                                fmt.Println(v12)
                                            }
                                            v2 += 2
                                        } else {
                                            for fi25 := 0; fi25 < 7; fi25++ {
                                                switch ("a" + ((("a" + "\n\t") + ("hello" + ("" + (("foo bar" + "\n\t") + "a")))) + (("a" + "foo bar") + "a"))) {
                                                case "a":
                                                    switch "foo bar" {
                                                    case "foo bar":
                                                        var v26 float32 = (1.5 + 0.1)
                                                        _ = v26
                                                    case "a":
                                                        v27 := GenS2{F0: true, F1: -979, F2: -2147483648}
                                                        _ = v27
                                                    default:
                                                        var v28 uint = 435
                                                        _ = v28
                                                    }
                                                }
                                                v2 = (v2 & (v2 >> ((42) & 63)))
                                                _ = fi25
                                            }
                                            for fi29 := 0; fi29 < 5; fi29++ {
                                                _ = len(v7)
                                                _ = fi29
                                            }
                                        }
                                        _ = fi24
                                    }
                                }
                                for fi30 := 0; fi30 < 12; fi30++ {
                                    v31 := GenS3{F0: 993, F1: 351, F2: ((v5 / (((v5 + 125)) | 1)) != v5)}
                                    var v32 int8 = (^(v8))
                                    _ = fi30
                                    _ = v31
                                    _ = v32
                                }
                            }
                            v33 := GenS1{F0: (v12 % ((107) | 1)), F1: 0}
                            _ = v33
                        } else {
                            v12 = (v12 - (^(v12)))
                            var v34 int64 = v14
                            _ = v34
                        }
                    }
                } else {
                    if len(v7) >= 2 {
                        v7 = v7[1:len(v7)]
                    }
                    v8 = -128
                }
                var v35 uint64 = uint64(v5)
                _ = v35
            }
            v36 := GenS3{F0: 2147483647, F1: (v2 / ((v2) | 1)), F2: ((^(v8)) != 93)}
            var v37 bool = v6
            _ = v36
            _ = v37
        } else {
            _ = len(v7)
            var v38 float32 = 1e10
            _ = v38
        }
    }
    fmt.Println(v12)
    v3 += v3
    if v6 {
        fmt.Println(v8, v12, v5)
    } else {
        for fi39 := 0; fi39 < 7; fi39++ {
            var v40 uint = 1
            switch (fi39 | v3) {
            case 128:
                v41 := GenS3{F0: 851, F1: 42, F2: (0.1 > ((float64((((-2.25 - 1e10) * -2.25) / 3.14159)) + 1.5) / 1e10))}
                _ = v41
            case 9223372036854775807:
                var v42 uint8 = (v12 ^ v12)
                _ = v42
            case 4294967295:
                var v43 int16 = 1
                _ = v43
            default:
                for fi44 := 0; fi44 < 10; fi44++ {
                    if v6 {
                        v45 := v4[0]
                        var v46 uint = (v40 * (^(v40)))
                        _ = v45
                        _ = v46
                    } else {
                        v7 = append(v7, int(v40))
                        switch ((("hello" + "hello") + "") + "") {
                        case "\n\t":
                            var v47 float32 = ((1e10 * -2.25) + (1e10 + 1.5))
                            _ = v47
                        case "hello":
                            var v48 uint16 = 65535
                            _ = v48
                        case "a":
                            v6 = v6
                        default:
                            fmt.Println(fi39)
                        }
                    }
                    var v49 int32 = 2147483647
                    _ = fi44
                    _ = v49
                }
            }
            var v50 float64 = ((3.14159 + ((3.14159 / 3.14159) * 1e10)) - float64(1e10))
            _ = fi39
            _ = v40
            _ = v50
        }
        v51 := GenS4{F0: v8, F1: v14, F2: 2147483647}
        switch v3 {
        case 2147483647:
            var v52 uint = 1
            _ = v52
        }
        _ = v51
    }
    for fi53 := 0; fi53 < 18; fi53++ {
        v8 -= v8
        v54 := GenS2{F0: v6, F1: -746, F2: -2147483648}
        _ = fi53
        _ = v54
    }
    v5 = v5
    v3 = (-(v3))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
