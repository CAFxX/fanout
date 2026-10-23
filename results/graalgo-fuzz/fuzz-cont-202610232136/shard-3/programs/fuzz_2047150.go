package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 bool; F1 int8 }
type GenS4 struct { F0 uint; F1 int }
func genFn0(p0 uint32, p1 uint32, p2 uint16) int8 {
    fmt.Println(p2)
    v6 := []int{32768, 2147483647, 128}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 1
}
func genFn1(p0 int16) int64 {
    v7 := []int{127, 32768, 127}
    if true {
        for fi8 := 0; fi8 < 20; fi8++ {
            switch (fi8 - fi8) {
            case 32768:
                var v9 uint = 129
                _ = v9
            case -9223372036854775808:
                fmt.Println(fi8)
            case -2147483648:
                fmt.Println(p0)
            }
            _ = fi8
        }
        p0 += p0
    } else {
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        for fi10 := 0; fi10 < 17; fi10++ {
            p0 = p0
            _ = fi10
        }
    }
    p0 += (p0 | (p0 | (p0 << ((4294967295) & 63))))
    var v11 float32 = ((((1.5 + (((0.1 * float32(0.1)) - ((3.14159 + 1.5) / 3.14159)) / 1e10)) - ((0.0 * 0.0) + 3.14159)) + 0.1) - 3.14159)
    _ = p0
    _ = v11
    _ = v7
    return 42
}

func main() {
    var v1 int8 = 100
    var v2 bool = (-1 < 9223372036854775807)
    var v3 int16 = -1
    var v4 int64 = 1
    var v5 uint32 = 0
    var v12 uint16 = 255
    var v13 uint8 = 127
    var v14 int8 = (v1 + (v1 * (-(v1))))
    var v15 string = ("hello" + "a")
    var v16 int32 = 2147483647
    var v17 int32 = v16
    var v18 uint32 = (v5 ^ (v5 + 0))
    if ((!((1e10 - -2.25) != (1e10 * (-2.25 - (3.14159 - ((0.1 * -2.25) + ((0.0 / (((1.5 - 0.0) - 0.1) * 1e10)) * (1e10 + 3.14159)))))))) && ((v12 ^ (^(v12))) < (v12 / (((v12 / ((32767) | 1))) | 1)))) {
        fmt.Println(v13, v12)
    } else {
        v17 += 1
    }
    switch "a" {
    case "foo bar":
        v19 := GenS1{F0: (v14 >> ((uint(v12)) & 63))}
        _ = v19
    case "hello":
        if v2 {
            if (-274 > (v16 % ((v16) | 1))) {
                v17 -= -2147483648
                var v20 uint16 = (^(v12))
                if (!v2) {
                    if (genFn1(v3) <= (v4 | -1)) {
                        switch v15 {
                        case "a":
                            var v21 bool = (4294967295 >= -1)
                            _ = v21
                        case "\n\t":
                            fmt.Println(v12, v14)
                        case "":
                            v13 = 86
                        }
                    } else {
                        fmt.Println(v4, v2)
                        var v22 int16 = int16(v16)
                        _ = v22
                    }
                    v23 := []int{42, 401, 963, 42}
                    var v24 uint16 = v20
                    _ = v23
                    _ = v24
                } else {
                    if ((^(v20)) >= (v12 & (v12 << ((255) & 63)))) {
                        v25 := GenS1{F0: 100}
                        _ = v25
                    } else {
                        for fi26 := 0; fi26 < 1; fi26++ {
                            var v27 int32 = (v17 & v16)
                            if (v2 && v2) {
                                v28 := GenS1{F0: genFn0((v5 / ((v5) | 1)), (v18 >> ((4294967295) & 63)), (^(v20)))}
                                _ = v28
                            }
                            _ = fi26
                            _ = v27
                        }
                    }
                    v29 := GenS2{F0: 42}
                    _ = v29
                }
                _ = v20
            }
            v3 -= 1
            if ((v4 >> ((65535) & 63)) != int64(v5)) {
                var v30 uint8 = (v13 + 128)
                _ = v30
            } else {
                var v31 uint32 = 4294967295
                if v2 {
                    var v32 uint64 = 0
                    v33 := []int{}
                    _ = v32
                    _ = v33
                }
                _ = v31
            }
        }
    default:
        v34 := []int{0, 1}
        _ = v34
    }
    v16 = (v17 & (v16 + (^(v17))))
    v35 := []int{9223372036854775807}
    if (!(!v2)) {
        switch "hello" {
        case "foo bar":
            var v36 int16 = (^(v3))
            _ = v36
        case "a":
            var v37 int = -2147483648
            _ = v37
        }
    } else {
        var v38 uint32 = (v5 << ((65535) & 63))
        v17 = v16
        var v39 float32 = 3.14159
        _ = v38
        _ = v39
    }
    _ = len(v35)
    switch -676 {
    case -1:
        if (v2 && ((v13 - v13) != (v13 / ((v13) | 1)))) {
            if v2 {
                var v40 uint = 42
                _ = v40
            }
            fmt.Println(v5)
            var v41 uint64 = 839
            _ = v41
        }
    default:
        var v42 uint64 = 2
        _ = v42
    }
    if v2 {
        var v43 int16 = v3
        switch 128 {
        case 1:
            if true {
                v12 -= v12
                fmt.Println(v2, v43)
            } else {
                v4 = (v4 / (((v4 << ((255) & 63))) | 1))
                if ((v13 * (v13 / ((v13) | 1))) == (v13 | 1)) {
                    var v44 uint8 = (v13 | v13)
                    for fi45 := 0; fi45 < 19; fi45++ {
                        if ((v2 && ((v17 >> ((2) & 63)) != v17)) || v2) {
                            v14 = (v1 - v14)
                            v3 -= int16(v17)
                        } else {
                            switch 32767 {
                            case -1:
                                var v46 float32 = (0.1 * (1e10 - 0.0))
                                _ = v46
                            default:
                                v47 := GenS3{F0: true, F1: (v1 >> ((18446744073709551615) & 63))}
                                _ = v47
                            }
                            v3 -= (v3 & v3)
                        }
                        if (-709 < -1) {
                            if v2 {
                                fmt.Println(v3, v43, v5)
                                if v2 {
                                    var v48 int32 = (v17 + (v16 << ((776) & 63)))
                                    v49 := GenS1{F0: 42}
                                    if v2 {
                                        fmt.Println(v3)
                                        fmt.Println(v13, v15, v43)
                                    } else {
                                        var v50 int16 = (v43 << ((65535) & 63))
                                        if ((fi45 * -246) < (fi45 ^ int(v12))) {
                                            v13 += 220
                                        }
                                        _ = v50
                                    }
                                    _ = v48
                                    _ = v49
                                } else {
                                    if (true && v2) {
                                        if v2 {
                                            _ = len(v35)
                                        }
                                        if len(v35) >= 2 {
                                            v35 = v35[1:len(v35)]
                                        }
                                    }
                                }
                            } else {
                                for fi51 := 0; fi51 < 4; fi51++ {
                                    fmt.Println(fi45, v1, v43)
                                    _ = fi51
                                }
                                switch fi45 {
                                case -1:
                                    v52 := GenS1{F0: v1}
                                    _ = v52
                                case 127:
                                    if len(v35) > 0 {
                                        v35[0] = (^(fi45))
                                    }
                                case 9223372036854775807:
                                    if v2 {
                                        if v2 {
                                            switch v15 {
                                            case "foo bar":
                                                if v2 {
                                                    v53 := GenS3{F0: (1 <= 4294967295), F1: -128}
                                                    _ = v53
                                                }
                                            }
                                            if v2 {
                                                if false {
                                                    v54 := GenS3{F0: v2, F1: v1}
                                                    var v55 uint32 = (^(v5))
                                                    _ = v54
                                                    _ = v55
                                                }
                                                var v56 uint8 = (v13 >> ((65535) & 63))
                                                _ = v56
                                            }
                                        }
                                    } else {
                                        var v57 uint32 = (v5 / ((2) | 1))
                                        _ = v57
                                    }
                                }
                                var v58 bool = (18446744073709551615 < 18446744073709551615)
                                _ = v58
                            }
                            fmt.Println(v14, v15, v17)
                        } else {
                            var v59 uint = 2
                            _ = v59
                        }
                        _ = fi45
                    }
                    v60 := GenS4{F0: 42, F1: 0}
                    _ = v44
                    _ = v60
                } else {
                    v35 = append(v35, -944)
                    fmt.Println(v15, v13, v18)
                    fmt.Println(v17, v1)
                }
            }
        case -2147483648:
            for fi61 := 0; fi61 < 20; fi61++ {
                v12 = v12
                var v62 string = "foo bar"
                _ = fi61
                _ = v62
            }
        case 9223372036854775807:
            var v63 int = 951
            _ = v63
        default:
            v5 = v18
        }
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
}
