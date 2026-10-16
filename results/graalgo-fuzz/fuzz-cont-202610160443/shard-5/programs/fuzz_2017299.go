package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string; F1 uint16; F2 uint64 }
type GenS3 struct { F0 int; F1 int64 }
type GenS4 struct { F0 uint32; F1 string; F2 int64 }
func genFn0() uint {
    var v6 uint64 = 333
    v7 := GenS1{F0: 0}
    for fi8 := 0; fi8 < 13; fi8++ {
        for fi9 := 0; fi9 < 7; fi9++ {
            v6 += (v6 - 18446744073709551615)
            _ = fi9
        }
        switch "hello" {
        case "foo bar":
            switch "a" {
            case "":
                v6 += 585
            case "\n\t":
                fmt.Println(v6)
            case "a":
                v10 := []int{32768, 553}
                _ = v10
            }
        case "a":
            fmt.Println(v6, v6)
        }
        _ = fi8
    }
    if (true || (true || (false && false))) {
        for fi11 := 0; fi11 < 7; fi11++ {
            switch "" {
            case "":
                var v12 string = ("foo bar" + "")
                _ = v12
            case "hello":
                var v13 int16 = 42
                _ = v13
            default:
                if true {
                    v14 := GenS1{F0: 100}
                    if false {
                        var v15 uint64 = 4294967295
                        var v16 uint64 = v15
                        _ = v15
                        _ = v16
                    }
                    _ = v14
                } else {
                    if ("hello" <= "hello") {
                        var v17 float64 = ((1e10 + float64(((1.5 - 0.1) + float32((float64(1.5) * float64(3.14159)))))) * 0.0)
                        _ = v17
                    }
                    for fi18 := 0; fi18 < 8; fi18++ {
                        v19 := GenS1{F0: int8(v6)}
                        if false {
                            v6 -= v6
                            if false {
                                var v20 uint = 0
                                v6 += 4294967295
                                _ = v20
                            }
                            fmt.Println(fi11)
                            if ((fi18 & 32768) <= (fi11 & fi11)) {
                                if true {
                                    v21 := []int{(-(fi18)), fi18, fi11}
                                    if len(v21) >= 2 {
                                        v21 = v21[1:len(v21)]
                                    }
                                    _ = v21
                                } else {
                                    if true {
                                        var v22 string = ("foo bar" + ("foo bar" + ("hello" + ("foo bar" + ""))))
                                        switch (fi18 << ((42) & 63)) {
                                        case -902:
                                            v22 = v22
                                        case 2147483647:
                                            v6 -= v6
                                        case -282:
                                            var v23 uint64 = 2
                                            _ = v23
                                        }
                                        v22 = v22
                                        _ = v22
                                    }
                                    fmt.Println(v6, v6, fi18)
                                }
                                var v24 int32 = 1
                                _ = v24
                            }
                            var v25 float64 = 0.0
                            if (32767 >= 169) {
                                v26 := GenS2{F0: "a", F1: 65535, F2: (^(v6))}
                                _ = v26
                            } else {
                                if (!true) {
                                    var v27 float32 = float32(v25)
                                    v27 += v27
                                    _ = v27
                                }
                                var v28 int16 = -32768
                                fmt.Println(v25, fi11, v6)
                                _ = v28
                            }
                            _ = v25
                        } else {
                            v29 := GenS2{F0: "a", F1: 255, F2: (^(v6))}
                            v30 := []int{}
                            _ = v29
                            _ = v30
                        }
                        _ = fi18
                        _ = v19
                    }
                }
            }
            var v31 int = fi11
            var v32 uint8 = 115
            _ = fi11
            _ = v31
            _ = v32
        }
    }
    switch "foo bar" {
    case "\n\t":
        if false {
            for fi33 := 0; fi33 < 6; fi33++ {
                v6 += v6
                _ = fi33
            }
        }
    case "foo bar":
        v6 -= v6
    default:
        var v34 int16 = -340
        _ = v34
    }
    _ = v6
    _ = v7
    return 18446744073709551615
}
func genFn1(p0 int64, p1 int16) int {
    v35 := GenS3{F0: 32767, F1: p0}
    switch -268 {
    case 496:
        switch "" {
        case "":
            for fi36 := 0; fi36 < 19; fi36++ {
                p0 -= (p0 * p0)
                _ = fi36
            }
        }
    case 42:
        switch (("\n\t" + ("hello" + "foo bar")) + "\n\t") {
        case "foo bar":
            if (127 <= 2) {
                fmt.Println(p0, p1)
            } else {
                if (!false) {
                    v37 := GenS1{F0: -65}
                    _ = v37
                } else {
                    p1 = (p1 % ((1000) | 1))
                    if false {
                        v38 := []int{-2147483648}
                        p0 = -9223372036854775808
                        switch 127 {
                        case 304:
                            if false {
                                v39 := GenS3{F0: 4294967295, F1: p0}
                                _ = v39
                            }
                        default:
                            switch 0 {
                            case 42:
                                p1 -= 621
                            default:
                                v38 = append(v38, 32768)
                            }
                        }
                        _ = v38
                    } else {
                        var v40 string = "a"
                        _ = v40
                    }
                }
            }
        }
    case -888:
        p0 = -9223372036854775808
    }
    _ = p0
    _ = p1
    _ = v35
    return 127
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int8 = 1
    var v3 float64 = 1.5
    var v4 int32 = 42
    var v5 int32 = (-(v4))
    fmt.Println(v5, v3, v4)
    fmt.Println(v2, v5, v5)
    if (!true) {
        v41 := v1[236]
        _ = v41
    } else {
        fmt.Println(v4, v3, v5)
    }
    var v42 uint8 = 128
    v43 := GenS4{F0: 2147483647, F1: "foo bar", F2: -945}
    var v44 float32 = 3.14159
    v45 := GenS2{F0: "\n\t", F1: 1, F2: 980}
    v4 += v4
    switch genFn1(878, 32767) {
    case 32768:
        v46 := []int{-82, -549, genFn1(int64(v5), 1)}
        _ = v46
    case 42:
        if false {
            switch ("a" + "hello") {
            case "\n\t":
                var v47 uint16 = 1
                _ = v47
            case "":
                var v48 float32 = (0.1 / (1e10 * (1e10 / (1.5 * (((1e10 / ((0.1 / v44) - ((((v44 - float32((0.1 * v3))) / (0.0 / ((v44 - (0.0 / v44)) + (v44 / v44)))) / (((1e10 / v44) * (1e10 / (1e10 * 1.5))) - 3.14159)) * (1.5 - float32((((-2.25 + ((-2.25 + float64(v44)) / ((v3 + v3) * v3))) - 3.14159) - 0.1)))))) / float32(((v3 - v3) - (((0.0 / (v3 / v3)) - 0.0) / 3.14159)))) + v44)))))
                _ = v48
            default:
                v49 := GenS3{F0: 883, F1: -137}
                _ = v49
            }
            var v50 int8 = 127
            _ = v50
        } else {
            for fi51 := 0; fi51 < 4; fi51++ {
                fmt.Println(v44, v2, v44)
                _ = fi51
            }
            var v52 int8 = v2
            _ = v52
        }
    default:
        fmt.Println(v5, v2, v42)
    }
    var v53 int32 = (v5 ^ v5)
    v3 = (v3 + ((0.0 * v3) / 1.0))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v53
}
