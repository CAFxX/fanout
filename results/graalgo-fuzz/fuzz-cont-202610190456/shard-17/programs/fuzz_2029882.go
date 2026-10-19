package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 int8 }
type GenS2 struct { F0 int64; F1 int16; F2 int32 }
type GenS3 struct { F0 string; F1 int16 }

func main() {
    v1 := make(map[int]uint32)
    v2 := []uint32{65535, 2147483648}
    var v3 float32 = 1e10
    var v4 string = "\n\t"
    var v5 float64 = 1e10
    v2 = append(v2, 65535)
    _ = len(v2)
    var v6 float32 = v3
    if false {
        for fi7 := 0; fi7 < 16; fi7++ {
            v6 -= float32(float64(1e10))
            v6 += 1.5
            _ = fi7
        }
        var v8 int16 = 1
        var v9 uint8 = 2
        _ = v8
        _ = v9
    } else {
        v10 := GenS1{F0: -733, F1: 0, F2: -128}
        var v11 uint = 4294967295
        switch v4 {
        case "\n\t":
            switch 146 {
            case -163:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            default:
                if true {
                    var v12 uint64 = 2
                    for fi13 := 0; fi13 < 19; fi13++ {
                        if true {
                            var v14 int8 = int8(fi13)
                            v15 := GenS1{F0: (fi13 << ((v11) & 63)), F1: (v14 ^ v14), F2: v14}
                            v4 = v4
                            _ = v14
                            _ = v15
                        }
                        v16 := GenS1{F0: (fi13 + (-(fi13))), F1: -128, F2: 108}
                        _ = fi13
                        _ = v16
                    }
                    v17 := GenS1{F0: 4294967295, F1: 127, F2: -128}
                    _ = v12
                    _ = v17
                } else {
                    for fi18 := 0; fi18 < 9; fi18++ {
                        for fi19 := 0; fi19 < 13; fi19++ {
                            var v20 uint32 = uint32(v11)
                            _ = len(v1)
                            var v21 int8 = -15
                            _ = fi19
                            _ = v20
                            _ = v21
                        }
                        fmt.Println(fi18)
                        _ = fi18
                    }
                    var v22 string = "foo bar"
                    _ = v22
                }
            }
        case "foo bar":
            switch 763 {
            case 506:
                if (0.1 < 0.0) {
                    v11 += (^(v11))
                    var v23 bool = true
                    _ = len(v1)
                    _ = v23
                } else {
                    v24 := GenS2{F0: 42, F1: -1, F2: -2147483648}
                    switch "" {
                    case "hello":
                        for fi25 := 0; fi25 < 1; fi25++ {
                            for fi26 := 0; fi26 < 3; fi26++ {
                                v4 = ((v4 + v4) + "")
                                switch fi26 {
                                case -473:
                                    if (814 <= 731) {
                                        switch "\n\t" {
                                        case "":
                                            switch ("a" + v4) {
                                            case "":
                                                fmt.Println(v3, v3, v3)
                                            case "a":
                                                var v27 float32 = (v3 + v6)
                                                _ = v27
                                            default:
                                                switch (fi26 - fi25) {
                                                case -9223372036854775808:
                                                    var v28 uint = (v11 | (v11 ^ 279))
                                                    _ = v28
                                                case -1:
                                                    var v29 int = fi25
                                                    switch fi25 {
                                                    case 2147483647:
                                                        v1[v29] = 4294967295
                                                    case 206:
                                                        switch "" {
                                                        case "hello":
                                                            var v30 uint = v11
                                                            _ = v30
                                                        case "":
                                                            v29 = fi26
                                                            _ = len(v2)
                                                        case "\n\t":
                                                            v1[128] = 65535
                                                        }
                                                    case 798:
                                                        var v31 uint16 = 544
                                                        _ = v31
                                                    }
                                                    v11 -= 4294967295
                                                    _ = v29
                                                }
                                            }
                                        }
                                        v3 = (v3 / v3)
                                        switch (((("foo bar" + ((("a" + ("a" + v4)) + "hello") + v4)) + v4) + v4) + ("a" + "hello")) {
                                        case "":
                                            switch "" {
                                            case "":
                                                var v32 string = v4
                                                _ = v32
                                            }
                                        default:
                                            if true {
                                                v33 := GenS2{F0: 1, F1: 32767, F2: -2147483648}
                                                v5 -= (0.0 * 1.5)
                                                var v34 uint = v11
                                                switch (fi25 ^ -932) {
                                                case -53:
                                                    if true {
                                                        var v35 float64 = ((3.14159 - ((float64(0.1) / v5) / float64((0.0 * float32(((v5 * ((v5 / v5) / (0.0 * ((v5 / 1.5) / v5)))) + float64(3.14159))))))) / 0.1)
                                                        v2 = append(v2, 2)
                                                        if (!(true || (!(343 == 645)))) {
                                                            v36 := GenS1{F0: (fi26 % ((fi25) | 1)), F1: 127, F2: 100}
                                                            var v37 int32 = -110
                                                            _ = v36
                                                            _ = v37
                                                        }
                                                        _ = v35
                                                    }
                                                case 42:
                                                    fmt.Println(v5, v5)
                                                case 0:
                                                    v38 := v1[(fi26 / ((fi26) | 1))]
                                                    _ = v38
                                                }
                                                v11 = (^(v34))
                                                v2 = append(v2, 65535)
                                                var v39 int32 = -1
                                                _ = v33
                                                _ = v34
                                                _ = v39
                                            } else {
                                                v3 += (0.0 * (1.5 * v6))
                                            }
                                        }
                                        v5 -= v5
                                        v6 -= (3.14159 * -2.25)
                                    } else {
                                        var v40 int64 = -551
                                        v2 = append(v2, 2)
                                        var v41 uint8 = 114
                                        _ = v40
                                        _ = v41
                                    }
                                default:
                                    v5 -= 1e10
                                }
                                _ = fi26
                            }
                            _ = fi25
                        }
                    case "\n\t":
                        var v42 uint64 = 582
                        _ = v42
                    default:
                        v1[1] = 42
                    }
                    _ = v24
                }
            default:
                var v43 int8 = -128
                _ = v43
            }
        }
        _ = v10
        _ = v11
    }
    fmt.Println(v5, v6, v4)
    for fi44 := 0; fi44 < 2; fi44++ {
        var v45 bool = false
        var v46 string = v4
        _ = fi44
        _ = v45
        _ = v46
    }
    var v47 uint = 2
    fmt.Println(v47, v3)
    for fi48 := 0; fi48 < 12; fi48++ {
        v49 := GenS3{F0: ((v4 + (v4 + "hello")) + v4), F1: 0}
        if (true && false) {
            fmt.Println(v5)
        } else {
            var v50 uint32 = 684
            switch v4 {
            case "\n\t":
                for fi51 := 0; fi51 < 7; fi51++ {
                    fmt.Println(v4)
                    var v52 bool = false
                    v53 := GenS2{F0: 0, F1: int16(v50), F2: int32(v50)}
                    _ = fi51
                    _ = v52
                    _ = v53
                }
            }
            _ = v50
        }
        _ = fi48
        _ = v49
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v4 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v6
}
