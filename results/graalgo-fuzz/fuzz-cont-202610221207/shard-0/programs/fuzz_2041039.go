package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 uint16; F1 int64; F2 uint16 }
type GenS3 struct { F0 uint64; F1 uint32; F2 uint8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 bool; F1 uint8 }
type GenS6 struct { F0 int64; F1 int32; F2 int16 }
func genFn0(p0 uint8, p1 uint8) uint64 {
    p0 += 0
    var v7 int16 = 424
    v7 = (v7 | (-(v7)))
    var v8 float64 = -2.25
    v9 := []int{9223372036854775807, -190, 127}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 400
}
func genFn1(p0 int8, p1 int, p2 uint32) uint {
    fmt.Println(p1, p2)
    var v10 bool = false
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 2
}

func main() {
    var v1 string = ("hello" + "a")
    v2 := make(map[string]bool)
    var v3 float32 = (0.1 + float32(3.14159))
    var v4 string = v1
    v5 := make(map[uint32]string)
    var v6 bool = (!true)
    v1 = ""
    v6 = false
    var v11 int16 = 0
    v12 := []int{4294967295}
    v13 := GenS1{F0: 0, F1: 42}
    v14 := GenS2{F0: 813, F1: 805, F2: 377}
    switch "hello" {
    case "\n\t":
        delete(v5, 1)
    }
    var v15 uint8 = 15
    v6 = false
    v3 = float32(0.0)
    var v16 int16 = 0
    for fi17 := 0; fi17 < 5; fi17++ {
        switch 4294967295 {
        case 302:
            v16 += (v11 * 0)
        case -60:
            fmt.Println(v16, v1)
        default:
            if v6 {
                for fi18 := 0; fi18 < 13; fi18++ {
                    v19 := GenS1{F0: 472, F1: 62}
                    var v20 int32 = -1
                    if (0 == 18446744073709551615) {
                        v11 = (^(v16))
                        if ((2147483647 <= 767) || ((-2.25 + 1e10) < v3)) {
                            switch fi18 {
                            case 0:
                                if true {
                                    v2[v4] = (v6 || true)
                                    v21 := v2[v4]
                                    v12 = append(v12, 127)
                                    _ = v21
                                } else {
                                    var v22 uint64 = genFn0(v15, v15)
                                    _ = v22
                                }
                            case 22:
                                v23 := GenS3{F0: 279, F1: 0, F2: v15}
                                _ = v23
                            }
                        } else {
                            fmt.Println(v1, v6, v16)
                            var v24 int64 = -9223372036854775808
                            switch v1 {
                            case "":
                                v1 = "hello"
                                v20 -= (^(v20))
                            case "hello":
                                switch (-(fi17)) {
                                case 128:
                                    if (2 <= 2) {
                                        var v25 uint = 2
                                        if ((fi18 | fi17) != fi17) {
                                            v2[((v1 + v1) + "\n\t")] = v6
                                        }
                                        var v26 uint = genFn1(1, (fi18 - (^(fi17))), 474)
                                        _ = v25
                                        _ = v26
                                    }
                                case -1:
                                    switch fi18 {
                                    case 127:
                                        v11 += (v16 - (v11 - -1))
                                    case 9223372036854775807:
                                        var v27 float64 = (float64(v3) / ((((-2.25 / 3.14159) * 3.14159) / 0.1) / 1.5))
                                        _ = v27
                                    case 758:
                                        _ = len(v5)
                                    }
                                }
                            case "a":
                                if false {
                                    v11 = v16
                                    if (v6 && false) {
                                        v3 = 1e10
                                    }
                                    switch -929 {
                                    case 32767:
                                        var v28 int32 = (^(v20))
                                        _ = v28
                                    case 128:
                                        var v29 float32 = v3
                                        _ = v29
                                    case 1:
                                        var v30 uint64 = 4294967295
                                        _ = v30
                                    }
                                } else {
                                    switch v1 {
                                    case "\n\t":
                                        v3 = v3
                                    case "":
                                        _ = len(v2)
                                    case "hello":
                                        v31 := GenS1{F0: (v20 % ((v20) | 1)), F1: 100}
                                        _ = v31
                                    }
                                    switch (fi18 ^ 4294967295) {
                                    case 694:
                                        if (v6 || v6) {
                                            fmt.Println(v24, v1)
                                        }
                                    }
                                }
                            default:
                                fmt.Println(v15)
                                if v6 {
                                    if (!(v20 >= v20)) {
                                        fmt.Println(fi17, v24)
                                        if (100 == -1) {
                                            if v6 {
                                                var v32 int = (fi18 | fi18)
                                                _ = v32
                                            }
                                            switch v1 {
                                            case "\n\t":
                                                v11 = 836
                                            }
                                        }
                                        var v33 int = (fi17 % (((^(fi17))) | 1))
                                        v34 := v2[v1]
                                        fmt.Println(v4)
                                        _ = v33
                                        _ = v34
                                    } else {
                                        if len(v12) >= 2 {
                                            v12 = v12[1:len(v12)]
                                        }
                                    }
                                }
                            }
                            v12 = append(v12, 648)
                            var v35 uint16 = 255
                            if ("a" > "hello") {
                                if (v11 == (v16 % ((v16) | 1))) {
                                    fmt.Println(v16, fi17, fi18)
                                    v20 -= (-(v20))
                                    var v36 int64 = (^(v24))
                                    _ = v36
                                }
                                v1 = v1
                                switch fi18 {
                                case -2147483648:
                                    v37 := GenS4{F0: 0}
                                    _ = v37
                                case 9223372036854775807:
                                    fmt.Println(v11, v6)
                                    v1 = ""
                                    var v38 uint64 = 1
                                    _ = v38
                                case -832:
                                    var v39 int16 = v16
                                    _ = v39
                                }
                                var v40 float64 = 3.14159
                                _ = v40
                            } else {
                                if (!(367 == 42)) {
                                    v41 := v5[0]
                                    fmt.Println(v20, v35)
                                    v20 += (v20 << ((uint(v15)) & 63))
                                    _ = v41
                                }
                                if v6 {
                                    switch fi18 {
                                    case 42:
                                        switch 32767 {
                                        case 32768:
                                            v42 := GenS5{F0: v6, F1: 1}
                                            _ = v42
                                        case 986:
                                            fmt.Println(v3, v3, fi17)
                                        case 32767:
                                            if v6 {
                                                var v43 string = ""
                                                v24 -= 9223372036854775807
                                                _ = v43
                                            } else {
                                                if len(v12) > 0 {
                                                    v12[0] = -896
                                                }
                                                var v44 uint16 = (v35 - v35)
                                                _ = v44
                                            }
                                        }
                                    default:
                                        if len(v12) > 0 {
                                            v12[0] = fi17
                                        }
                                    }
                                }
                                v45 := GenS1{F0: (^(v20)), F1: -1}
                                _ = v45
                            }
                            _ = v24
                            _ = v35
                        }
                        v15 += (v15 ^ v15)
                    }
                    _ = fi18
                    _ = v19
                    _ = v20
                }
            }
        }
        for fi46 := 0; fi46 < 4; fi46++ {
            if v6 {
                if v6 {
                    var v47 string = v1
                    _ = v47
                } else {
                    v48 := GenS6{F0: -9223372036854775808, F1: 506, F2: (v16 ^ 32767)}
                    var v49 int = fi17
                    _ = v48
                    _ = v49
                }
            } else {
                fmt.Println(v1, v15, v16)
                v3 += v3
                v16 = (v11 * v11)
            }
            v12 = append(v12, int(v15))
            _ = fi46
        }
        for fi50 := 0; fi50 < 12; fi50++ {
            v16 = v16
            _ = fi50
        }
        _ = fi17
    }
    if false {
        v4 = ("a" + ((("hello" + "a") + v1) + v4))
    } else {
        v11 -= (-(v11))
        var v51 int16 = v11
        _ = v51
    }
    if ((2 < 752) && v6) {
        v11 -= (v11 / (((v16 << ((851) & 63))) | 1))
        v52 := GenS3{F0: 4294967295, F1: 65535, F2: 162}
        var v53 uint16 = 179
        _ = v52
        _ = v53
    }
    switch (((v1 + ("hello" + "hello")) + v4) + v1) {
    case "foo bar":
        var v54 bool = v6
        _ = v54
    default:
        switch v1 {
        case "hello":
            v16 = v16
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
