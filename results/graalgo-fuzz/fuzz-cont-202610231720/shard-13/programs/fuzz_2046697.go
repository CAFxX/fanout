package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int8; F1 int64 }

func main() {
    v1 := make(map[int]int)
    var v2 uint64 = 2
    var v3 float64 = (1e10 + 3.14159)
    v4 := v1[2147483647]
    v5 := []int{int(v2), 32767}
    _ = len(v1)
    if false {
        v6 := v1[v4]
        if (275 <= 0) {
            v3 = ((1e10 * (0.0 - 3.14159)) + v3)
            var v7 uint64 = v2
            switch ("hello" + "") {
            case "hello":
                var v8 float32 = float32(v3)
                _ = v8
            default:
                var v9 int32 = -830
                _ = v9
            }
            _ = v7
        } else {
            v10 := GenS1{F0: 2}
            switch "\n\t" {
            case "a":
                for fi11 := 0; fi11 < 19; fi11++ {
                    switch (fi11 / (((fi11 & (fi11 - 42))) | 1)) {
                    case -775:
                        if (!true) {
                            switch (("hello" + "foo bar") + "a") {
                            case "a":
                                v2 -= (v2 * (v2 | v2))
                            case "hello":
                                v12 := v1[(fi11 * (-(v6)))]
                                _ = v12
                            }
                            v6 -= fi11
                            var v13 uint32 = 2147483648
                            _ = v13
                        } else {
                            if len(v5) > 0 {
                                v5[0] = v4
                            }
                            v14 := GenS1{F0: 128}
                            _ = v14
                        }
                    default:
                        if (!(572 != int64(v6))) {
                            v15 := GenS2{F0: 325}
                            fmt.Println(v2, fi11)
                            _ = v15
                        }
                    }
                    _ = fi11
                }
            case "\n\t":
                var v16 int64 = -1
                _ = v16
            case "foo bar":
                v2 = v2
            }
            fmt.Println(v4)
            _ = v10
        }
        if (!(255 > 0)) {
            var v17 int64 = -1
            v3 += 0.0
            fmt.Println(v2, v3, v2)
            _ = v17
        } else {
            if (-1 != 0) {
                var v18 bool = true
                for fi19 := 0; fi19 < 5; fi19++ {
                    v2 += v2
                    v20 := GenS1{F0: 167}
                    switch "" {
                    case "":
                        if (42 < 2) {
                            fmt.Println(v3, v18)
                            if (!v18) {
                                var v21 int = v6
                                var v22 bool = v18
                                v23 := GenS3{F0: -42}
                                _ = v21
                                _ = v22
                                _ = v23
                            }
                            v2 += (v2 >> ((4294967295) & 63))
                        } else {
                            v6 += (-(v4))
                        }
                    case "foo bar":
                        v24 := v1[v6]
                        _ = v24
                    default:
                        v2 = v2
                    }
                    _ = fi19
                    _ = v20
                }
                for fi25 := 0; fi25 < 7; fi25++ {
                    switch ("hello" + "") {
                    case "hello":
                        var v26 int = v4
                        _ = v26
                    case "a":
                        var v27 uint = 4294967295
                        _ = v27
                    case "":
                        var v28 uint16 = 2
                        _ = v28
                    }
                    fmt.Println(v3, v4)
                    _ = len(v1)
                    _ = fi25
                }
                _ = v18
            } else {
                var v29 uint16 = 0
                v6 = v4
                _ = v29
            }
            fmt.Println(v3)
            if (!((439 != 2147483648) && (4294967295 != 715))) {
                if true {
                    v30 := GenS1{F0: 41}
                    _ = v30
                } else {
                    fmt.Println(v2)
                    for fi31 := 0; fi31 < 19; fi31++ {
                        v3 = v3
                        _ = fi31
                    }
                    switch 723 {
                    case 127:
                        if (2 != 32767) {
                            v3 = float64((1e10 * 0.1))
                            v32 := GenS4{F0: ((0 <= 42) && (!(42 >= 42)))}
                            switch (v6 & v4) {
                            case 1000:
                                v2 = (v2 ^ 619)
                            case 985:
                                var v33 float32 = (0.0 + ((3.14159 + 1.5) * 3.14159))
                                _ = v33
                            case 768:
                                v2 -= v2
                            default:
                                fmt.Println(v3, v3, v6)
                            }
                            _ = v32
                        }
                    case 32768:
                        if (v2 == v2) {
                            var v34 bool = true
                            v2 = (v2 << ((65535) & 63))
                            v35 := v1[v6]
                            _ = v34
                            _ = v35
                        }
                    default:
                        if ((2 >= 255) || (258 == 32767)) {
                            for fi36 := 0; fi36 < 3; fi36++ {
                                if len(v5) > 0 {
                                    v5[0] = fi36
                                }
                                if (1 == v2) {
                                    fmt.Println(v6)
                                    if (false || true) {
                                        switch 128 {
                                        case 2147483647:
                                            v3 = ((v3 / float64((3.14159 * ((((1.5 + float32(0.1)) / -2.25) - -2.25) / 0.1)))) + (3.14159 * (((float64(0.0) + float64(1e10)) * v3) + 3.14159)))
                                        case 9223372036854775807:
                                            for fi37 := 0; fi37 < 8; fi37++ {
                                                v6 += fi37
                                                var v38 string = "\n\t"
                                                var v39 uint8 = 1
                                                v2 = v2
                                                _ = fi37
                                                _ = v38
                                                _ = v39
                                            }
                                        case 0:
                                            for fi40 := 0; fi40 < 1; fi40++ {
                                                v4 -= -2147483648
                                                if (136 > 1000) {
                                                    fmt.Println(fi40, v6)
                                                    if (-358 >= 32767) {
                                                        var v41 uint64 = 4294967295
                                                        _ = v41
                                                    }
                                                } else {
                                                    v42 := GenS3{F0: 90}
                                                    v4 = fi36
                                                    if len(v5) > 0 {
                                                        v5[0] = v4
                                                    }
                                                    _ = v42
                                                }
                                                _ = fi40
                                            }
                                        default:
                                            v2 = v2
                                        }
                                    } else {
                                        for fi43 := 0; fi43 < 1; fi43++ {
                                            if (("hello" > "foo bar") || (0 >= 0)) {
                                                v1[v6] = fi43
                                                _ = len(v5)
                                            } else {
                                                v3 = (3.14159 * (v3 * ((v3 + (v3 * v3)) + (v3 - 1e10))))
                                            }
                                            fmt.Println(v3, v6, v3)
                                            _ = fi43
                                        }
                                        fmt.Println(v2, v2)
                                    }
                                    var v44 int32 = 19
                                    _ = v44
                                } else {
                                    var v45 uint64 = v2
                                    _ = v45
                                }
                                _ = fi36
                            }
                            v4 += (v6 / (((v4 ^ v6)) | 1))
                        }
                    }
                }
                var v46 uint = 18446744073709551615
                for fi47 := 0; fi47 < 3; fi47++ {
                    for fi48 := 0; fi48 < 18; fi48++ {
                        if true {
                            fmt.Println(v2, v6, v46)
                            if (false && false) {
                                v6 -= fi47
                            }
                        }
                        var v49 int16 = -1
                        var v50 int16 = (v49 & (v49 / (((v49 >> (((v46 & v46)) & 63))) | 1)))
                        _ = fi48
                        _ = v49
                        _ = v50
                    }
                    var v51 uint64 = v2
                    _ = fi47
                    _ = v51
                }
                _ = v46
            } else {
                delete(v1, 127)
                var v52 uint16 = 118
                _ = v52
            }
        }
        _ = v6
    }
    if (-2147483648 > -2147483648) {
        v5 = append(v5, -222)
        v53 := GenS5{F0: int16(v2)}
        _ = v53
    } else {
        var v54 float32 = (float32(v3) - 1.5)
        fmt.Println(v4, v3)
        if (65535 <= 463) {
            for fi55 := 0; fi55 < 9; fi55++ {
                for fi56 := 0; fi56 < 8; fi56++ {
                    v54 = 3.14159
                    v2 -= (v2 << ((967) & 63))
                    _ = fi56
                }
                var v57 bool = false
                var v58 int = (v4 % ((2147483647) | 1))
                _ = fi55
                _ = v57
                _ = v58
            }
        }
        _ = v54
    }
    for fi59 := 0; fi59 < 6; fi59++ {
        var v60 uint = 974
        var v61 float32 = 3.14159
        v62 := GenS6{F0: -35, F1: 548}
        _ = fi59
        _ = v60
        _ = v61
        _ = v62
    }
    fmt.Println(v2)
    v3 = v3
    if false {
        v3 += v3
        v63 := v1[127]
        _ = v63
    } else {
        var v64 string = "hello"
        _ = v64
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
