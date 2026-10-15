package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int; F1 uint32; F2 int16 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 int32; F1 uint64; F2 bool }
type GenS10 struct { F0 uint; F1 int64 }
func genFn0(p0 int32, p1 bool) string {
    v4 := GenS1{F0: 1}
    v5 := []int{9223372036854775807, -132, 127, 32768}
    v6 := GenS2{F0: 32767}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return "a"
}

func main() {
    v1 := []int{0, -9223372036854775808}
    var v2 int16 = 997
    var v3 float64 = 0.1
    for fi7 := 0; fi7 < 17; fi7++ {
        switch (fi7 & (fi7 << ((65535) & 63))) {
        case 257:
            for fi8 := 0; fi8 < 11; fi8++ {
                var v9 float32 = 1.5
                v3 = v3
                v1 = append(v1, fi7)
                _ = fi8
                _ = v9
            }
        case -9223372036854775808:
            if (!((v3 - v3) < (-2.25 + 0.0))) {
                fmt.Println(v3)
                var v10 float32 = -2.25
                _ = v10
            } else {
                var v11 uint32 = 2147483648
                var v12 bool = true
                if len(v1) > 0 {
                    v1[0] = 127
                }
                _ = v11
                _ = v12
            }
        default:
            if true {
                var v13 int8 = -45
                fmt.Println(v13, v3)
                _ = v13
            } else {
                if true {
                    switch "" {
                    case "hello":
                        var v14 uint8 = 42
                        _ = v14
                    case "a":
                        for fi15 := 0; fi15 < 6; fi15++ {
                            switch 4294967295 {
                            case 2147483647:
                                v16 := GenS3{F0: 2147483647}
                                _ = v16
                            case 764:
                                var v17 int32 = -2147483648
                                _ = v17
                            default:
                                var v18 int32 = -1
                                _ = v18
                            }
                            v3 += -2.25
                            _ = len(v1)
                            _ = fi15
                        }
                    case "foo bar":
                        var v19 uint32 = 2147483648
                        _ = v19
                    default:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    if false {
                        v20 := GenS4{F0: 180}
                        _ = v20
                    }
                    v21 := GenS5{F0: fi7, F1: 744, F2: (v2 ^ v2)}
                    _ = v21
                } else {
                    _ = len(v1)
                }
                v3 = (float64((-2.25 / 1.5)) - -2.25)
            }
        }
        _ = fi7
    }
    v3 = v3
    switch "foo bar" {
    case "a":
        fmt.Println(v2)
    case "hello":
        var v22 int16 = v2
        _ = v22
    case "":
        var v23 int16 = v2
        _ = v23
    default:
        for fi24 := 0; fi24 < 10; fi24++ {
            fmt.Println(fi24, v2, fi24)
            v3 += v3
            _ = fi24
        }
    }
    var v25 int64 = 9223372036854775807
    v25 = v25
    if (!true) {
        switch 174 {
        case 32767:
            v26 := GenS6{F0: false}
            _ = v26
        }
        var v27 int8 = 0
        _ = v27
    } else {
        v28 := GenS7{F0: v2}
        if (42 > 255) {
            var v29 int8 = 127
            v29 = v29
            var v30 float64 = -2.25
            _ = v29
            _ = v30
        }
        for fi31 := 0; fi31 < 15; fi31++ {
            switch fi31 {
            case 128:
                if false {
                    for fi32 := 0; fi32 < 20; fi32++ {
                        if false {
                            _ = len(v1)
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        } else {
                            v33 := GenS3{F0: 1}
                            v25 = 0
                            if true {
                                if (!false) {
                                    v34 := GenS6{F0: ((1 < 0) && (false && (v25 >= (v25 & v25))))}
                                    _ = v34
                                }
                            } else {
                                v35 := GenS2{F0: (v2 * -647)}
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                v25 = (v25 << ((42) & 63))
                                var v36 float32 = (0.1 * 1e10)
                                _ = v35
                                _ = v36
                            }
                            _ = v33
                        }
                        _ = fi32
                    }
                    var v37 uint32 = 65535
                    fmt.Println(v2, v37)
                    _ = v37
                } else {
                    switch "hello" {
                    case "hello":
                        if false {
                            fmt.Println(v25, fi31, fi31)
                        }
                    case "a":
                        v3 += 1e10
                    }
                    v38 := GenS8{F0: fi31}
                    switch (((("hello" + "a") + (("\n\t" + "foo bar") + "foo bar")) + "a") + ("foo bar" + ("" + (("hello" + ("hello" + "\n\t")) + "hello")))) {
                    case "foo bar":
                        var v39 bool = false
                        _ = v39
                    case "\n\t":
                        if ((fi31 - -891) < (fi31 | (fi31 & fi31))) {
                            v1 = append(v1, (-(fi31)))
                        } else {
                            fmt.Println(v2, v25)
                            v40 := GenS6{F0: (2 == 127)}
                            _ = v40
                        }
                    default:
                        for fi41 := 0; fi41 < 8; fi41++ {
                            var v42 int = (fi31 % ((32768) | 1))
                            var v43 uint8 = 255
                            _ = fi41
                            _ = v42
                            _ = v43
                        }
                    }
                    _ = v38
                }
            case 42:
                _ = len(v1)
            }
            v3 = (-2.25 - (v3 + float64(((0.1 + float32(1.5)) * 3.14159))))
            _ = fi31
        }
        _ = v28
    }
    if true {
        if (1 == 42) {
            var v44 float64 = v3
            var v45 uint8 = 12
            v45 = v45
            _ = v44
            _ = v45
        } else {
            for fi46 := 0; fi46 < 13; fi46++ {
                var v47 float64 = v3
                _ = fi46
                _ = v47
            }
        }
        fmt.Println(v3, v25)
        v48 := GenS8{F0: 1}
        _ = v48
    }
    if (((v3 - (v3 - v3)) / v3) > (v3 * v3)) {
        if (1 <= 0) {
            for fi49 := 0; fi49 < 3; fi49++ {
                var v50 int16 = (v2 / (((^(v2))) | 1))
                _ = fi49
                _ = v50
            }
            if false {
                if (42 <= 2147483647) {
                    switch 896 {
                    case 0:
                        var v51 uint8 = 255
                        _ = v51
                    case 32768:
                        v1 = append(v1, 42)
                    default:
                        _ = len(v1)
                    }
                    v2 = v2
                }
            }
        }
        fmt.Println(v2, v2, v3)
    }
    var v52 uint64 = 18446744073709551615
    v53 := GenS2{F0: v2}
    for fi54 := 0; fi54 < 16; fi54++ {
        if false {
            if true {
                for fi55 := 0; fi55 < 4; fi55++ {
                    v2 += (^(v2))
                    v2 = (v2 & (v2 << ((1) & 63)))
                    v52 += v52
                    _ = fi55
                }
                var v56 int = (fi54 % (((fi54 ^ fi54)) | 1))
                _ = v56
            } else {
                for fi57 := 0; fi57 < 17; fi57++ {
                    var v58 uint16 = 153
                    var v59 bool = false
                    _ = fi57
                    _ = v58
                    _ = v59
                }
                v60 := GenS9{F0: 0, F1: (v52 | v52), F2: true}
                _ = v60
            }
        } else {
            v61 := GenS10{F0: 2, F1: 42}
            _ = v61
        }
        v3 = 0.0
        v52 = v52
        _ = fi54
    }
    switch -2147483648 {
    case 1:
        for fi62 := 0; fi62 < 13; fi62++ {
            v3 = (v3 / (v3 + (float64(0.0) + ((3.14159 + v3) - v3))))
            _ = fi62
        }
    case 127:
        var v63 int32 = 0
        _ = v63
    case -305:
        v3 -= ((v3 / v3) * v3)
    }
    for fi64 := 0; fi64 < 14; fi64++ {
        if true {
            var v65 int64 = -846
            for fi66 := 0; fi66 < 9; fi66++ {
                switch fi64 {
                case 0:
                    if len(v1) > 0 {
                        v1[0] = (fi64 / ((int(v25)) | 1))
                    }
                }
                _ = fi66
            }
            _ = v65
        } else {
            _ = len(v1)
            v52 = 18446744073709551615
        }
        _ = fi64
    }
    for fi67 := 0; fi67 < 3; fi67++ {
        for fi68 := 0; fi68 < 3; fi68++ {
            if ((3.14159 + 3.14159) != 0.1) {
                fmt.Println(fi68, v25)
                var v69 uint64 = (v52 + v52)
                fmt.Println(v2, v2)
                _ = v69
            }
            _ = fi68
        }
        _ = fi67
    }
    fmt.Println(v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v52
    _ = v53
}
