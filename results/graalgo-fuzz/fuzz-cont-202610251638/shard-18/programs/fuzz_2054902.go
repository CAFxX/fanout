package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int; F1 string }
type GenS4 struct { F0 uint }

func main() {
    var v1 uint32 = 2147483647
    var v2 int64 = 9223372036854775807
    v3 := make(map[uint32]string)
    for fi4 := 0; fi4 < 16; fi4++ {
        if (!false) {
            fmt.Println(fi4)
            if (!(-2.25 < 0.1)) {
                v5 := []int{(^(fi4)), 127, (fi4 % ((-1) | 1))}
                if false {
                    var v6 string = "foo bar"
                    var v7 uint32 = v1
                    if (false || ((42 <= 2) || true)) {
                        for fi8 := 0; fi8 < 13; fi8++ {
                            v9 := GenS1{F0: 42, F1: 0}
                            _ = fi8
                            _ = v9
                        }
                        var v10 float32 = 0.0
                        v11 := GenS2{F0: 0}
                        _ = v10
                        _ = v11
                    }
                    _ = v6
                    _ = v7
                } else {
                    v12 := GenS3{F0: (fi4 ^ fi4), F1: "hello"}
                    _ = v12
                }
                switch fi4 {
                case 9223372036854775807:
                    if (42 != 65535) {
                        for fi13 := 0; fi13 < 14; fi13++ {
                            var v14 int64 = (v2 << ((255) & 63))
                            _ = fi13
                            _ = v14
                        }
                        switch -572 {
                        case -794:
                            v1 += (v1 ^ 0)
                        }
                    } else {
                        for fi15 := 0; fi15 < 8; fi15++ {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            _ = fi15
                        }
                        v2 += (v2 ^ v2)
                        v16 := GenS3{F0: (-(fi4)), F1: ("foo bar" + "")}
                        _ = v16
                    }
                case -2147483648:
                    for fi17 := 0; fi17 < 16; fi17++ {
                        switch "hello" {
                        case "foo bar":
                            if (((4294967295 != uint(v1)) || false) && (63 != 0)) {
                                var v18 int64 = -849
                                var v19 uint32 = (^(v1))
                                _ = v18
                                _ = v19
                            } else {
                                v1 = v1
                                if ((v2 <= (v2 | v2)) || true) {
                                    var v20 int = (fi17 - -577)
                                    _ = v20
                                }
                            }
                        default:
                            var v21 uint8 = 128
                            _ = v21
                        }
                        if ((0.1 + 0.1) >= 0.1) {
                            v22 := GenS3{F0: 9223372036854775807, F1: ""}
                            _ = v22
                        } else {
                            if ("a" >= "\n\t") {
                                v3[(v1 - v1)] = ""
                                var v23 int16 = -32768
                                v1 += 621
                                if true {
                                    var v24 uint32 = 42
                                    _ = len(v3)
                                    fmt.Println(v24, v1)
                                    var v25 uint32 = 371
                                    _ = v24
                                    _ = v25
                                }
                                v1 = (v1 * v1)
                                fmt.Println(v23)
                                _ = v23
                            }
                            var v26 bool = (255 < 42)
                            _ = v26
                        }
                        switch "\n\t" {
                        case "a":
                            switch (fi4 * fi17) {
                            case 9223372036854775807:
                                fmt.Println(fi4, fi17, v1)
                            default:
                                if true {
                                    var v27 int32 = -1
                                    switch fi4 {
                                    case 1:
                                        fmt.Println(v27)
                                    }
                                    _ = v27
                                } else {
                                    var v28 int16 = 788
                                    _ = v28
                                }
                            }
                        case "foo bar":
                            fmt.Println(fi17)
                        case "\n\t":
                            v29 := GenS2{F0: 219}
                            var v30 int64 = v2
                            v30 += 581
                            _ = v29
                            _ = v30
                        default:
                            v2 -= (v2 * (v2 ^ (-(v2))))
                        }
                        _ = fi17
                    }
                case -217:
                    v1 -= (^(v1))
                }
                _ = v5
            } else {
                switch "" {
                case "a":
                    var v31 int64 = (v2 ^ v2)
                    _ = v31
                case "hello":
                    var v32 int32 = int32(v2)
                    _ = v32
                case "foo bar":
                    v33 := []int{fi4, fi4, fi4, fi4}
                    _ = v33
                default:
                    v34 := v3[2147483648]
                    _ = v34
                }
                var v35 int64 = v2
                _ = v35
            }
        } else {
            switch fi4 {
            case 791:
                var v36 int8 = 127
                _ = v36
            }
            v1 = 2
            if true {
                var v37 int = (^(fi4))
                _ = v37
            } else {
                if ((false || (-205 >= -1)) || true) {
                    fmt.Println(v1, fi4, v1)
                }
                v2 = v2
            }
        }
        v2 += v2
        v2 = (v2 + (v2 / ((v2) | 1)))
        _ = fi4
    }
    v38 := []int{378, -1, 42}
    v39 := GenS4{F0: 68}
    v40 := v3[(v1 & (^(v1)))]
    v3[2147483648] = v40
    fmt.Println(v1, v40, v40)
    var v41 uint = 263
    switch v40 {
    case "hello":
        var v42 uint16 = 255
        _ = v42
    case "":
        var v43 string = (v40 + "hello")
        _ = v43
    case "\n\t":
        v2 = (v2 + v2)
    }
    switch v40 {
    case "a":
        var v44 uint = v41
        _ = v44
    case "":
        var v45 uint32 = 450
        _ = v45
    default:
        for fi46 := 0; fi46 < 11; fi46++ {
            if (v1 != (v1 ^ 65535)) {
                v2 = 1
            } else {
                v40 = (v40 + (v40 + (v40 + v40)))
                v2 -= (v2 ^ (-(v2)))
            }
            _ = fi46
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v40
    _ = v41
}
