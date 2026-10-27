package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 string }
type GenS2 struct { F0 int8; F1 uint16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint32; F1 uint64 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int8; F1 uint; F2 string }
type GenS7 struct { F0 int32; F1 string }
func genFn0() bool {
    v4 := []int{539, 4294967295, 9223372036854775807}
    if len(v4) > 0 {
        v4[0] = -94
    }
    v5 := GenS1{F0: 562, F1: 65535, F2: ""}
    var v6 uint32 = 94
    _ = v4
    _ = v5
    _ = v6
    return (-888 == -1)
}

func main() {
    v1 := []int{-236}
    var v2 int32 = 16
    var v3 int8 = 100
    if true {
        if (true || false) {
            for fi7 := 0; fi7 < 11; fi7++ {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v2 = v2
                _ = fi7
            }
            v8 := GenS2{F0: v3, F1: 193}
            _ = v8
        }
    } else {
        if true {
            var v9 uint8 = 2
            _ = v9
        } else {
            v2 = v2
            var v10 int64 = -1
            for fi11 := 0; fi11 < 3; fi11++ {
                _ = len(v1)
                _ = fi11
            }
            _ = v10
        }
        fmt.Println(v2)
    }
    fmt.Println(v3, v2)
    v2 = (v2 * 0)
    v12 := GenS2{F0: v3, F1: 262}
    v3 += (v3 + v3)
    for fi13 := 0; fi13 < 1; fi13++ {
        var v14 float32 = (0.0 / -2.25)
        v2 -= (v2 | v2)
        var v15 uint32 = 2147483648
        _ = fi13
        _ = v14
        _ = v15
    }
    if false {
        v16 := GenS1{F0: 4294967295, F1: 895, F2: "a"}
        _ = v16
    }
    for fi17 := 0; fi17 < 13; fi17++ {
        if (!true) {
            v18 := GenS3{F0: 0}
            var v19 float64 = 0.1
            switch "a" {
            case "a":
                v20 := GenS2{F0: (v3 - v3), F1: 2}
                _ = v20
            }
            _ = v18
            _ = v19
        }
        v3 = (v3 + 126)
        _ = fi17
    }
    var v21 string = ("" + "hello")
    var v22 uint8 = 193
    switch 127 {
    case -1:
        for fi23 := 0; fi23 < 5; fi23++ {
            switch v21 {
            case "foo bar":
                for fi24 := 0; fi24 < 9; fi24++ {
                    switch v21 {
                    case "\n\t":
                        _ = len(v1)
                    }
                    switch 42 {
                    case 4294967295:
                        if true {
                            if false {
                                if (2147483647 > 4294967295) {
                                    v1 = append(v1, 2147483647)
                                    var v25 bool = (2 == 4294967295)
                                    var v26 float32 = -2.25
                                    _ = v25
                                    _ = v26
                                }
                                var v27 int = (fi23 >> ((931) & 63))
                                _ = v27
                            }
                            v3 = 41
                        }
                    case 42:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    default:
                        var v28 uint8 = (^(v22))
                        _ = v28
                    }
                    _ = fi24
                }
            case "\n\t":
                v2 = v2
            default:
                for fi29 := 0; fi29 < 12; fi29++ {
                    fmt.Println(fi29, fi29, v22)
                    switch fi29 {
                    case 42:
                        v30 := GenS4{F0: 4294967295, F1: 42}
                        _ = v30
                    }
                    _ = fi29
                }
            }
            fmt.Println(v22)
            for fi31 := 0; fi31 < 14; fi31++ {
                v2 -= 0
                if false {
                    v1 = append(v1, 9223372036854775807)
                    switch fi23 {
                    case -686:
                        _ = len(v1)
                        v2 = 0
                        fmt.Println(v3, fi31)
                    case 0:
                        v21 = (v21 + v21)
                    case 32767:
                        var v32 bool = (0 < 4294967295)
                        _ = v32
                    default:
                        v1 = append(v1, 42)
                    }
                    fmt.Println(fi23, v3)
                } else {
                    v33 := GenS5{F0: 42}
                    var v34 uint32 = 167
                    _ = v33
                    _ = v34
                }
                _ = fi31
            }
            _ = fi23
        }
    case 32767:
        var v35 uint32 = 4294967295
        _ = v35
    case 127:
        v21 = ""
    }
    if (!(-2.25 >= ((0.1 + (1e10 / 1.0)) / (((1e10 - float64(3.14159)) - 3.14159) * -2.25)))) {
        for fi36 := 0; fi36 < 12; fi36++ {
            if (v2 < (v2 / ((-346) | 1))) {
                if false {
                    v22 += v22
                    v37 := GenS4{F0: 286, F1: 0}
                    _ = v37
                } else {
                    var v38 float64 = 0.1
                    v38 = (float64(((float32((v38 - 0.1)) + 1e10) / 1e10)) + (v38 + v38))
                    _ = v38
                }
                if true {
                    var v39 string = v21
                    if len(v1) > 0 {
                        v1[0] = fi36
                    }
                    _ = v39
                } else {
                    var v40 int = (fi36 | fi36)
                    _ = v40
                }
                if (int32(v22) != v2) {
                    var v41 uint = 42
                    v42 := GenS1{F0: fi36, F1: 686, F2: (v21 + v21)}
                    _ = v41
                    _ = v42
                } else {
                    if (65535 > 32767) {
                        var v43 int64 = 9223372036854775807
                        for fi44 := 0; fi44 < 2; fi44++ {
                            var v45 uint32 = 488
                            _ = fi44
                            _ = v45
                        }
                        var v46 uint = 42
                        _ = v43
                        _ = v46
                    } else {
                        fmt.Println(v21, v2)
                        for fi47 := 0; fi47 < 6; fi47++ {
                            v48 := GenS5{F0: 1}
                            _ = fi47
                            _ = v48
                        }
                        v49 := GenS2{F0: (^(v3)), F1: 42}
                        _ = v49
                    }
                }
            } else {
                v50 := GenS6{F0: v3, F1: 18446744073709551615, F2: (("hello" + "a") + v21)}
                var v51 int32 = -2147483648
                _ = v50
                _ = v51
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            switch (-(fi36)) {
            case 528:
                if (v2 <= -1) {
                    var v52 int32 = v2
                    var v53 int = fi36
                    _ = v52
                    _ = v53
                } else {
                    switch fi36 {
                    case 491:
                        switch ("a" + v21) {
                        case "\n\t":
                            if (false && (!false)) {
                                var v54 uint8 = (v22 & 128)
                                _ = v54
                            }
                        }
                    }
                    v55 := GenS7{F0: (^(v2)), F1: v21}
                    _ = v55
                }
            default:
                var v56 uint64 = 42
                _ = v56
            }
            _ = fi36
        }
        if (1 == 4294967295) {
            var v57 int64 = 42
            var v58 int8 = 42
            _ = v57
            _ = v58
        } else {
            var v59 int = 1
            _ = v59
        }
        v22 += 2
    }
    var v60 int = -335
    fmt.Println(v60, v60)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v12
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v60
}
