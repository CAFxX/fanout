package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 int32 }
type GenS2 struct { F0 int64; F1 uint16; F2 uint32 }
type GenS3 struct { F0 uint; F1 int8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint32; F1 int64; F2 uint8 }
type GenS6 struct { F0 bool; F1 uint64; F2 uint16 }
type GenS7 struct { F0 int32 }
func genFn0(p0 string) int8 {
    p0 = p0
    fmt.Println(p0)
    p0 = "foo bar"
    _ = p0
    return -18
}

func main() {
    var v1 bool = true
    var v2 uint32 = 2
    var v3 float32 = 3.14159
    v4 := make(map[string]int)
    v5 := make(map[uint32]int)
    var v6 int64 = 0
    if (255 == 345) {
        var v7 uint = 2
        _ = v7
    }
    v8 := GenS1{F0: "\n\t", F1: v6, F2: 936}
    if (float64((v3 * v3)) >= 0.1) {
        for fi9 := 0; fi9 < 12; fi9++ {
            v3 = (((v3 * (v3 + v3)) + 3.14159) / ((v3 / v3) - v3))
            _ = fi9
        }
    } else {
        for fi10 := 0; fi10 < 15; fi10++ {
            v1 = (v6 > (v6 - -1))
            _ = fi10
        }
        switch "\n\t" {
        case "":
            switch ("\n\t" + ("" + "")) {
            case "\n\t":
                if false {
                    switch ("\n\t" + "\n\t") {
                    case "\n\t":
                        var v11 int32 = -2147483648
                        _ = v11
                    case "a":
                        if false {
                            v1 = (!v1)
                        } else {
                            var v12 uint64 = 1
                            v12 += 0
                            v13 := GenS2{F0: 42, F1: 42, F2: (v2 << ((285) & 63))}
                            _ = v12
                            _ = v13
                        }
                    }
                } else {
                    v3 += v3
                    var v14 bool = v1
                    _ = v14
                }
            case "foo bar":
                v3 += float32(float64((v3 / v3)))
            case "hello":
                var v15 int8 = int8(v6)
                _ = v15
            }
        case "hello":
            v16 := GenS2{F0: v6, F1: 431, F2: 2}
            _ = v16
        case "\n\t":
            switch 425 {
            case 9223372036854775807:
                var v17 int16 = -32768
                _ = v17
            case 705:
                var v18 uint32 = v2
                _ = v18
            default:
                v19 := []int{625, 4294967295, 2147483647}
                _ = v19
            }
        }
        fmt.Println(v1, v1, v3)
    }
    v20 := GenS2{F0: (v6 ^ -786), F1: 5, F2: 0}
    if (-2.25 <= (1e10 - ((0.1 * -2.25) * 0.0))) {
        v3 -= (v3 * (v3 + (float32(((0.0 + ((0.0 * (1e10 * 0.1)) - 0.0)) / (1e10 - (1.5 / 3.14159)))) / (v3 - 1e10))))
        if false {
            v21 := GenS3{F0: 18446744073709551615, F1: 0}
            var v22 float64 = (float64(0.0) / 1.0)
            _ = v21
            _ = v22
        } else {
            v23 := GenS2{F0: v6, F1: 32767, F2: v2}
            v24 := GenS3{F0: 2, F1: -36}
            v25 := GenS4{F0: false}
            _ = v23
            _ = v24
            _ = v25
        }
    } else {
        v3 = float32(float64(v3))
        if v1 {
            v26 := GenS5{F0: (v2 | 65535), F1: (-(v6)), F2: 255}
            var v27 uint64 = 858
            _ = v26
            _ = v27
        } else {
            fmt.Println(v1)
            if (42 <= 0) {
                switch 0 {
                case 9223372036854775807:
                    _ = len(v4)
                default:
                    for fi28 := 0; fi28 < 4; fi28++ {
                        if (4294967295 != 94) {
                            v1 = (32767 > 786)
                            v2 -= 0
                            if v1 {
                                var v29 uint16 = 913
                                v29 += (v29 >> ((255) & 63))
                                v30 := GenS4{F0: (1e10 != 0.1)}
                                _ = v29
                                _ = v30
                            } else {
                                if v1 {
                                    v31 := GenS4{F0: v1}
                                    _ = v31
                                }
                                v32 := GenS4{F0: v1}
                                _ = v32
                            }
                        } else {
                            var v33 int8 = -1
                            fmt.Println(v6)
                            _ = v33
                        }
                        _ = fi28
                    }
                }
                for fi34 := 0; fi34 < 18; fi34++ {
                    if false {
                        fmt.Println(v6, v6)
                    } else {
                        v35 := GenS6{F0: true, F1: 4294967295, F2: 42}
                        fmt.Println(v2, v1, v3)
                        switch "a" {
                        case "foo bar":
                            var v36 bool = v1
                            _ = v36
                        case "\n\t":
                            switch ("foo bar" + (("a" + (("\n\t" + "foo bar") + "hello")) + "")) {
                            case "a":
                                var v37 float32 = v3
                                _ = v37
                            case "foo bar":
                                var v38 int = fi34
                                _ = v38
                            }
                        case "":
                            switch fi34 {
                            case 127:
                                if v1 {
                                    if v1 {
                                        v39 := GenS7{F0: 42}
                                        _ = v39
                                    }
                                } else {
                                    delete(v4, ("foo bar" + "hello"))
                                    var v40 uint16 = 335
                                    switch "hello" {
                                    case "hello":
                                        v41 := GenS2{F0: v6, F1: (v40 * 2), F2: v2}
                                        _ = v41
                                    case "":
                                        var v42 int8 = 100
                                        _ = v42
                                    default:
                                        v43 := GenS5{F0: uint32(v40), F1: v6, F2: 255}
                                        _ = v43
                                    }
                                    _ = v40
                                }
                            }
                        default:
                            v6 -= v6
                        }
                        _ = v35
                    }
                    for fi44 := 0; fi44 < 5; fi44++ {
                        if false {
                            v1 = v1
                            v6 = (-(v6))
                            if v1 {
                                v1 = v1
                            }
                            if (!true) {
                                v2 += (v2 + (v2 / (((v2 | 196)) | 1)))
                            } else {
                                v6 -= (v6 << ((65535) & 63))
                                v3 = 0.1
                                switch ("a" + "foo bar") {
                                case "\n\t":
                                    v6 += 42
                                    v45 := []int{fi34, fi34, (fi34 * fi44), fi34}
                                    var v46 float32 = v3
                                    _ = v45
                                    _ = v46
                                default:
                                    v47 := []int{(-(fi44))}
                                    _ = v47
                                }
                                switch "" {
                                case "\n\t":
                                    v48 := GenS4{F0: (554 >= 400)}
                                    _ = v48
                                case "hello":
                                    v2 += 2
                                }
                            }
                        }
                        _ = fi44
                    }
                    _ = fi34
                }
            }
            v6 = -643
        }
        if false {
            switch ("foo bar" + "") {
            case "hello":
                v49 := v4["foo bar"]
                _ = v49
            case "a":
                var v50 int16 = -1
                _ = v50
            }
            fmt.Println(v2, v6, v2)
        }
    }
    if v1 {
        var v51 int8 = -1
        _ = v51
    }
    v52 := []int{}
    v53 := GenS7{F0: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v52:", len(v52), v52)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v53
    _ = v6
    _ = v8
}
