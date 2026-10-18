package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string; F1 int64; F2 int }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint32; F1 uint16; F2 uint32 }
type GenS5 struct { F0 uint32; F1 int8 }
func genFn0(p0 uint32, p1 uint8, p2 uint16) int64 {
    v6 := []int{}
    v7 := GenS1{F0: 1}
    p0 = (p0 ^ (p0 & 657))
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 606
}
func genFn1(p0 int, p1 uint8, p2 int32) int64 {
    var v8 int64 = 122
    v9 := GenS1{F0: 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return v8
}

func main() {
    var v1 bool = true
    v2 := make(map[int]uint32)
    v3 := []int{612, 32767, 42}
    v4 := make(map[string]string)
    var v5 uint = 255
    v10 := GenS1{F0: 0}
    if true {
        v1 = (65535 != 2)
        if false {
            var v11 int16 = 1
            v11 = v11
            _ = len(v2)
            _ = v11
        }
        if true {
            fmt.Println(v5, v1)
            v5 = v5
        }
    }
    var v12 uint16 = 2
    for fi13 := 0; fi13 < 11; fi13++ {
        for fi14 := 0; fi14 < 19; fi14++ {
            _ = len(v2)
            _ = fi14
        }
        _ = fi13
    }
    v1 = (42 <= -1)
    v15 := GenS1{F0: 127}
    var v16 int64 = 42
    fmt.Println(v16)
    for fi17 := 0; fi17 < 15; fi17++ {
        switch "a" {
        case "hello":
            if v1 {
                var v18 int16 = 32767
                v12 = 289
                _ = v18
            }
        case "":
            if len(v3) > 0 {
                v3[0] = (fi17 / (((fi17 << ((v5) & 63))) | 1))
            }
        case "\n\t":
            var v19 float32 = 0.0
            _ = v19
        }
        if false {
            if v1 {
                var v20 bool = ((32767 > -32768) || (100 > -1))
                _ = v20
            } else {
                _ = len(v3)
                for fi21 := 0; fi21 < 6; fi21++ {
                    if len(v3) > 0 {
                        v3[0] = -897
                    }
                    v22 := GenS1{F0: -1}
                    switch (fi17 ^ -994) {
                    case 440:
                        switch 42 {
                        case 2147483647:
                            if false {
                                var v23 uint32 = 0
                                v5 = (v5 & (v5 << (((v5 + (v5 << (((v5 >> (((v5 / ((255) | 1))) & 63))) & 63)))) & 63)))
                                _ = v23
                            } else {
                                v1 = v1
                                v24 := GenS1{F0: 1}
                                fmt.Println(v1, v1, v16)
                                v2[fi17] = 2
                                _ = v24
                            }
                        case -988:
                            var v25 float64 = (float64(float32(3.14159)) * -2.25)
                            _ = v25
                        }
                    case -209:
                        var v26 float32 = 3.14159
                        _ = v26
                    default:
                        v1 = v1
                    }
                    _ = fi21
                    _ = v22
                }
                switch int(v12) {
                case -91:
                    var v27 uint16 = v12
                    _ = v27
                case -2147483648:
                    if (42 >= 2147483648) {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        v28 := GenS2{F0: "foo bar", F1: -672, F2: 0}
                        v16 = 330
                        _ = v28
                    }
                default:
                    var v29 uint16 = (v12 & 0)
                    _ = v29
                }
            }
            v1 = v1
            switch fi17 {
            case -9223372036854775808:
                for fi30 := 0; fi30 < 14; fi30++ {
                    switch ((("a" + ("\n\t" + "")) + ("a" + "a")) + ("foo bar" + "")) {
                    case "foo bar":
                        v31 := GenS2{F0: "\n\t", F1: -1, F2: (-(fi30))}
                        _ = v31
                    case "a":
                        var v32 uint64 = 1
                        _ = v32
                    default:
                        fmt.Println(fi30, fi30, fi17)
                    }
                    v16 += (v16 - (v16 | v16))
                    switch "hello" {
                    case "hello":
                        fmt.Println(v16)
                    case "a":
                        v1 = ((int32(v16) < 620) || true)
                    case "\n\t":
                        v33 := GenS2{F0: ("" + "a"), F1: -1, F2: (fi17 >> ((2) & 63))}
                        _ = v33
                    default:
                        switch int(v16) {
                        case -9223372036854775808:
                            fmt.Println(fi30)
                        default:
                            var v34 uint64 = 2
                            _ = v34
                        }
                    }
                    if (736 >= 32767) {
                        fmt.Println(v5, v1)
                    } else {
                        _ = len(v3)
                    }
                    _ = fi30
                }
            case 4294967295:
                v35 := GenS3{F0: 1}
                _ = v35
            }
        } else {
            var v36 uint64 = 4294967295
            for fi37 := 0; fi37 < 5; fi37++ {
                switch fi17 {
                case 0:
                    var v38 uint32 = 2
                    _ = v38
                case 9223372036854775807:
                    switch ("a" + ("a" + "hello")) {
                    case "hello":
                        v5 = v5
                    case "a":
                        switch 42 {
                        case 0:
                            v39 := GenS4{F0: 1, F1: v12, F2: 627}
                            _ = v39
                        case 32767:
                            if (149 >= 796) {
                                v40 := GenS1{F0: 42}
                                v16 -= 0
                                _ = v40
                            }
                        default:
                            if (uint8(v12) != 127) {
                                if v1 {
                                    fmt.Println(v1)
                                    v41 := GenS5{F0: 787, F1: 29}
                                    _ = v41
                                } else {
                                    v1 = (1 != v36)
                                    var v42 uint64 = (v36 & v36)
                                    fmt.Println(v5, v5, fi37)
                                    _ = v42
                                }
                            }
                        }
                    }
                default:
                    v43 := GenS4{F0: 1, F1: (v12 | 669), F2: 335}
                    _ = v43
                }
                _ = fi37
            }
            v12 += (v12 * (^(v12)))
            _ = v36
        }
        for fi44 := 0; fi44 < 8; fi44++ {
            var v45 bool = v1
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            if true {
                var v46 int16 = -687
                fmt.Println(v12, v12, v46)
                v46 -= (v46 - v46)
                _ = v46
            }
            _ = fi44
            _ = v45
        }
        _ = fi17
    }
    var v47 float32 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v12
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
}
