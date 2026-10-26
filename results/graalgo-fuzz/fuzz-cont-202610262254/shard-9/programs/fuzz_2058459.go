package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 uint8 }
type GenS2 struct { F0 int8; F1 int64; F2 int32 }
type GenS3 struct { F0 int16; F1 int16 }
type GenS4 struct { F0 int64; F1 int32; F2 uint64 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint; F1 int16; F2 int8 }
type GenS7 struct { F0 uint32; F1 uint64; F2 bool }
type GenS8 struct { F0 uint64; F1 int; F2 int8 }

func main() {
    var v1 int = -973
    var v2 uint32 = 204
    var v3 int8 = -76
    var v4 string = "foo bar"
    for fi5 := 0; fi5 < 12; fi5++ {
        v6 := GenS1{F0: (v3 - (v3 + int8(v2))), F1: 1, F2: 2}
        v7 := []int{}
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 uint32 = v2
    var v9 uint = 18446744073709551615
    var v10 uint16 = 1
    v11 := GenS2{F0: (-(v3)), F1: 9223372036854775807, F2: -243}
    v9 = (v9 + (v9 * (v9 - (^(v9)))))
    if (!true) {
        v12 := GenS3{F0: 42, F1: -32768}
        switch (v4 + v4) {
        case "a":
            if true {
                v2 = 2
                for fi13 := 0; fi13 < 9; fi13++ {
                    v4 = "hello"
                    switch "\n\t" {
                    case "hello":
                        if true {
                            var v14 uint64 = 801
                            if ((v2 >> (((v9 >> ((v9) & 63))) & 63)) <= v2) {
                                v10 = (v10 | v10)
                                v15 := []int{v1, (v1 * (fi13 | (v1 - v1))), fi13}
                                for fi16 := 0; fi16 < 9; fi16++ {
                                    v17 := GenS4{F0: -9223372036854775808, F1: 1, F2: (v14 - (v14 - v14))}
                                    var v18 int8 = (-(v3))
                                    _ = fi16
                                    _ = v17
                                    _ = v18
                                }
                                _ = v15
                            }
                            for fi19 := 0; fi19 < 15; fi19++ {
                                fmt.Println(v4)
                                if false {
                                    if false {
                                        v20 := GenS5{F0: (v10 ^ uint16(v9))}
                                        _ = v20
                                    }
                                } else {
                                    v21 := GenS6{F0: (v9 >> ((v9) & 63)), F1: 0, F2: int8(v9)}
                                    _ = v21
                                }
                                v9 -= (v9 ^ (v9 & uint(v14)))
                                _ = fi19
                            }
                            _ = v14
                        }
                    case "foo bar":
                        if (((1e10 - 0.0) * 0.1) == float64(3.14159)) {
                            var v22 uint16 = (^(v10))
                            for fi23 := 0; fi23 < 3; fi23++ {
                                fmt.Println(v4, v9, fi13)
                                v24 := []int{(-(fi23))}
                                switch (^(fi23)) {
                                case 1:
                                    v22 += v22
                                case 522:
                                    if true {
                                        _ = len(v24)
                                    }
                                default:
                                    var v25 uint64 = 4294967295
                                    switch 32768 {
                                    case -227:
                                        _ = len(v24)
                                    case 128:
                                        fmt.Println(v22, v1)
                                    default:
                                        v22 -= 65535
                                    }
                                    if len(v24) >= 2 {
                                        v24 = v24[1:len(v24)]
                                    }
                                    _ = v25
                                }
                                _ = fi23
                                _ = v24
                            }
                            switch v4 {
                            case "foo bar":
                                v1 = 32768
                            case "a":
                                for fi26 := 0; fi26 < 15; fi26++ {
                                    var v27 float32 = (float32(3.14159) - (1.5 + (float32(0.1) / (-2.25 * (float32(1.5) - (((0.1 + -2.25) + 1e10) * (((0.0 / float32(1.5)) + 0.0) + (0.1 - 0.0))))))))
                                    v22 -= v10
                                    _ = fi26
                                    _ = v27
                                }
                            case "hello":
                                v2 = v2
                            }
                            _ = v22
                        }
                    case "":
                        var v28 string = v4
                        _ = v28
                    default:
                        if true {
                            v29 := GenS5{F0: v10}
                            switch v4 {
                            case "\n\t":
                                for fi30 := 0; fi30 < 18; fi30++ {
                                    var v31 uint = (v9 / ((v9) | 1))
                                    _ = fi30
                                    _ = v31
                                }
                            case "":
                                var v32 int16 = 42
                                _ = v32
                            }
                            _ = v29
                        } else {
                            for fi33 := 0; fi33 < 3; fi33++ {
                                v34 := GenS7{F0: uint32(v9), F1: 42, F2: false}
                                _ = fi33
                                _ = v34
                            }
                        }
                    }
                    v2 += (v2 >> ((18446744073709551615) & 63))
                    _ = fi13
                }
                v8 -= v8
            }
        case "\n\t":
            v1 = v1
        case "foo bar":
            var v35 int8 = v3
            _ = v35
        }
        for fi36 := 0; fi36 < 6; fi36++ {
            var v37 uint64 = 4294967295
            _ = fi36
            _ = v37
        }
        _ = v12
    }
    if true {
        v38 := []int{127, int(v9), -9223372036854775808}
        var v39 string = (v4 + (v4 + ("a" + "foo bar")))
        var v40 int64 = 0
        _ = v38
        _ = v39
        _ = v40
    }
    var v41 uint16 = v10
    fmt.Println(v8, v9, v3)
    v10 += v41
    var v42 uint64 = 2
    for fi43 := 0; fi43 < 4; fi43++ {
        for fi44 := 0; fi44 < 5; fi44++ {
            if true {
                var v45 uint32 = (^(v2))
                var v46 float64 = float64((3.14159 + float32(1e10)))
                v47 := GenS8{F0: (v42 >> ((v9) & 63)), F1: fi43, F2: v3}
                _ = v45
                _ = v46
                _ = v47
            } else {
                var v48 bool = (32767 > -233)
                v8 += 4294967295
                _ = v48
            }
            _ = fi44
        }
        switch v1 {
        case 32767:
            v49 := []int{}
            _ = v49
        case -635:
            v50 := []int{111}
            _ = v50
        }
        _ = fi43
    }
    v51 := GenS7{F0: v8, F1: (v42 | (v42 & v42)), F2: (!false)}
    if (2 == 172) {
        for fi52 := 0; fi52 < 1; fi52++ {
            v1 = (^(v1))
            v2 -= uint32(v41)
            var v53 float32 = 3.14159
            _ = fi52
            _ = v53
        }
        v54 := GenS6{F0: 721, F1: 42, F2: (v3 | 1)}
        v55 := GenS6{F0: 2, F1: 42, F2: v3}
        _ = v54
        _ = v55
    } else {
        if false {
            var v56 int16 = -67
            var v57 uint32 = 2
            _ = v56
            _ = v57
        } else {
            v9 -= 18446744073709551615
        }
        if false {
            var v58 string = "a"
            var v59 uint16 = v41
            _ = v58
            _ = v59
        } else {
            if false {
                var v60 int32 = -2147483648
                v61 := []int{v1, (v1 & v1)}
                v8 = v2
                _ = v60
                _ = v61
            }
            fmt.Println(v42, v41)
        }
    }
    fmt.Println(v42, v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v51
    _ = v8
    _ = v9
}
