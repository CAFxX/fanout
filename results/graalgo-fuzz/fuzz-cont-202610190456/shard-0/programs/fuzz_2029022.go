package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64; F1 int16; F2 int16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int64; F1 uint; F2 int16 }
type GenS8 struct { F0 int32; F1 bool; F2 int16 }
type GenS9 struct { F0 int16 }

func main() {
    var v1 string = ("" + ("foo bar" + "foo bar"))
    var v2 int16 = 42
    var v3 float64 = 0.1
    var v4 bool = false
    v5 := make(map[int]uint32)
    var v6 string = "foo bar"
    var v7 int16 = (v2 | (^(v2)))
    for fi8 := 0; fi8 < 4; fi8++ {
        v7 = v2
        _ = fi8
    }
    v4 = (v4 || (251 > 18446744073709551615))
    var v9 int8 = 100
    v1 = ("a" + v6)
    v10 := []int{-2147483648, 20}
    delete(v5, 146)
    if (((((float64((3.14159 * 0.1)) * v3) / 3.14159) + (((v3 * (v3 * 0.0)) - v3) + 3.14159)) / -2.25) == v3) {
        var v11 float64 = float64(((((0.1 / -2.25) / (3.14159 + (((0.1 - (1.5 * 0.0)) / ((0.1 - (1.5 * 0.0)) - -2.25)) - float32((v3 - ((1e10 - v3) - 1.5)))))) + (float32(v3) / (1e10 - 1.5))) - 3.14159))
        var v12 int64 = 9223372036854775807
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 9; fi13++ {
        v14 := v5[-324]
        v15 := GenS1{F0: v9}
        _ = fi13
        _ = v14
        _ = v15
    }
    var v16 uint = 0
    switch v1 {
    case "a":
        v6 = v1
    case "\n\t":
        for fi17 := 0; fi17 < 12; fi17++ {
            fmt.Println(v7, v3)
            for fi18 := 0; fi18 < 13; fi18++ {
                fmt.Println(v7, v3, v7)
                var v19 bool = (true || v4)
                switch 127 {
                case 4294967295:
                    var v20 int = -84
                    _ = v20
                case 127:
                    v21 := GenS2{F0: 1, F1: v7, F2: 0}
                    _ = v21
                default:
                    var v22 uint8 = 120
                    _ = v22
                }
                _ = fi18
                _ = v19
            }
            v5[(-(fi17))] = 0
            _ = fi17
        }
    case "foo bar":
        switch 32767 {
        case -703:
            fmt.Println(v7, v2)
        }
    default:
        delete(v5, 2147483647)
    }
    for fi23 := 0; fi23 < 9; fi23++ {
        if v4 {
            if v4 {
                if (42 <= 736) {
                    v1 = v1
                } else {
                    v4 = (v4 || (v4 || false))
                }
                v24 := GenS1{F0: v9}
                v2 = v2
                _ = v24
            } else {
                v5[2147483647] = 65535
            }
        } else {
            var v25 int8 = (v9 * (v9 % ((v9) | 1)))
            _ = v25
        }
        if (1 > uint16(v16)) {
            switch (-(fi23)) {
            case 0:
                for fi26 := 0; fi26 < 15; fi26++ {
                    fmt.Println(fi23)
                    v7 = (v2 + (v2 >> ((v16) & 63)))
                    var v27 int64 = -1
                    v28 := GenS3{F0: (v27 ^ v27)}
                    _ = fi26
                    _ = v27
                    _ = v28
                }
            case -2147483648:
                var v29 int32 = 984
                _ = v29
            case -1:
                v30 := GenS4{F0: 42}
                _ = v30
            default:
                v31 := GenS1{F0: v9}
                _ = v31
            }
            var v32 uint = (v16 >> ((4294967295) & 63))
            _ = v32
        } else {
            fmt.Println(v1)
            var v33 uint64 = 707
            for fi34 := 0; fi34 < 3; fi34++ {
                var v35 float32 = ((0.1 / 3.14159) * 3.14159)
                if v4 {
                    _ = len(v10)
                }
                _ = fi34
                _ = v35
            }
            _ = v33
        }
        _ = fi23
    }
    if v4 {
        if v4 {
            v36 := GenS5{F0: (v16 ^ (v16 & v16))}
            for fi37 := 0; fi37 < 16; fi37++ {
                var v38 string = v1
                v39 := GenS6{F0: 2}
                for fi40 := 0; fi40 < 2; fi40++ {
                    v6 = v1
                    switch v1 {
                    case "hello":
                        v5[(fi40 ^ 0)] = 0
                    case "foo bar":
                        v9 = int8(v16)
                    case "\n\t":
                        switch (fi37 / ((fi40) | 1)) {
                        case 127:
                            switch 32767 {
                            case 9223372036854775807:
                                if v4 {
                                    var v41 bool = (255 < 430)
                                    var v42 uint64 = 42
                                    fmt.Println(v7)
                                    delete(v5, fi37)
                                    _ = v41
                                    _ = v42
                                } else {
                                    var v43 bool = (0 <= v16)
                                    v44 := GenS4{F0: 816}
                                    _ = v43
                                    _ = v44
                                }
                            case 1:
                                if (-1 == 464) {
                                    var v45 int8 = 1
                                    v46 := GenS2{F0: 42, F1: (v7 - v7), F2: v7}
                                    var v47 float64 = (v3 / (0.1 * (v3 * (((v3 / 3.14159) + (v3 / ((1.5 - 0.0) * 0.1))) - (v3 - 1.5)))))
                                    switch (((v6 + "\n\t") + (v38 + v1)) + ("a" + v1)) {
                                    case "\n\t":
                                        fmt.Println(v38)
                                    case "a":
                                        v48 := GenS7{F0: -1, F1: (v16 | v16), F2: 693}
                                        _ = v48
                                    }
                                    v9 -= -128
                                    _ = v45
                                    _ = v46
                                    _ = v47
                                }
                            default:
                                _ = len(v10)
                            }
                        case -1:
                            if (!v4) {
                                v49 := GenS3{F0: -1}
                                _ = v49
                            }
                        case 9223372036854775807:
                            var v50 int64 = int64(v16)
                            _ = v50
                        }
                    }
                    fmt.Println(v3)
                    _ = fi40
                }
                _ = fi37
                _ = v38
                _ = v39
            }
            _ = v36
        } else {
            v9 -= v9
            var v51 int32 = 0
            _ = v51
        }
        var v52 int8 = (v9 & (-(v9)))
        for fi53 := 0; fi53 < 1; fi53++ {
            var v54 uint8 = 255
            _ = fi53
            _ = v54
        }
        _ = v52
    }
    delete(v5, 32768)
    switch -2147483648 {
    case -2147483648:
        var v55 int64 = 0
        _ = v55
    case -527:
        if (!v4) {
            if len(v10) > 0 {
                v10[0] = 9223372036854775807
            }
            var v56 float64 = (3.14159 + (0.1 * 1e10))
            for fi57 := 0; fi57 < 15; fi57++ {
                var v58 uint64 = 18446744073709551615
                if v4 {
                    if (!(255 > 32767)) {
                        switch "" {
                        case "foo bar":
                            v6 = "\n\t"
                        case "hello":
                            _ = len(v5)
                        case "\n\t":
                            var v59 int32 = -2147483648
                            _ = v59
                        }
                        v60 := v5[fi57]
                        v61 := GenS6{F0: (v58 - v58)}
                        _ = v60
                        _ = v61
                    } else {
                        if v4 {
                            _ = len(v5)
                        } else {
                            var v62 bool = true
                            _ = v62
                        }
                    }
                } else {
                    v63 := GenS8{F0: 2147483647, F1: (v4 || (32767 > 65535)), F2: v2}
                    _ = v63
                }
                v64 := GenS5{F0: (^(v16))}
                _ = fi57
                _ = v58
                _ = v64
            }
            _ = v56
        }
    case 128:
        if true {
            v65 := v5[824]
            _ = v65
        } else {
            v7 = (v7 + int16(v16))
            switch 4294967295 {
            case 9223372036854775807:
                v66 := GenS9{F0: v2}
                _ = v66
            case 0:
                var v67 uint = 65535
                _ = v67
            }
        }
    default:
        v68 := v5[131]
        _ = v68
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
