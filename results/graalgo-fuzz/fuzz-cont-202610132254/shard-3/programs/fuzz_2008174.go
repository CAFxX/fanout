package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 int64; F2 int16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int8; F1 bool; F2 uint64 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint16; F1 uint8 }
func genFn0(p0 int) int8 {
    var v7 float32 = (0.0 * 0.1)
    v7 += -2.25
    for fi8 := 0; fi8 < 7; fi8++ {
        v9 := GenS1{F0: 4294967295}
        fmt.Println(v7)
        v10 := GenS2{F0: 793, F1: -43, F2: 0}
        _ = fi8
        _ = v10
        _ = v9
    }
    v11 := GenS1{F0: 345}
    _ = p0
    _ = v11
    _ = v7
    return -111
}

func main() {
    v1 := []string{"hello"}
    var v2 string = "foo bar"
    var v3 uint = 18446744073709551615
    var v4 uint = (v3 << (((v3 << ((v3) & 63))) & 63))
    var v5 uint8 = uint8(v4)
    v6 := make(map[string]bool)
    var v12 int = -17
    for fi13 := 0; fi13 < 13; fi13++ {
        fmt.Println(v3, fi13, v5)
        v3 += (v4 * v3)
        _ = fi13
    }
    if true {
        v14 := GenS3{F0: v4}
        if ((!true) || (4294967295 < 4294967295)) {
            fmt.Println(v2)
        }
        if true {
            fmt.Println(v3, v5, v4)
            if (!(true || true)) {
                switch (v12 ^ v12) {
                case 128:
                    for fi15 := 0; fi15 < 9; fi15++ {
                        v4 += (v3 - uint(v5))
                        _ = fi15
                    }
                case 272:
                    v16 := GenS1{F0: v3}
                    _ = v16
                case -2147483648:
                    for fi17 := 0; fi17 < 16; fi17++ {
                        if (32767 > 32767) {
                            fmt.Println(v5, fi17)
                        }
                        var v18 uint8 = 42
                        _ = fi17
                        _ = v18
                    }
                }
            } else {
                _ = len(v6)
                var v19 float64 = 0.0
                if (!true) {
                    v20 := v6["hello"]
                    _ = v20
                }
                _ = v19
            }
            v3 = (v4 & v4)
        }
        _ = v14
    } else {
        v12 += 0
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v3 = (^(v3))
    v5 -= v5
    var v21 float64 = ((1e10 / 1.0) / 3.14159)
    if (!false) {
        if (!false) {
            v22 := GenS3{F0: (v4 ^ (v4 | v3))}
            for fi23 := 0; fi23 < 11; fi23++ {
                var v24 bool = (false || false)
                var v25 float32 = (float32((v21 + v21)) * 1e10)
                fmt.Println(v21, v12)
                _ = fi23
                _ = v24
                _ = v25
            }
            v26 := GenS2{F0: 42, F1: -9223372036854775808, F2: 1000}
            _ = v22
            _ = v26
        }
        if true {
            switch "" {
            case "foo bar":
                var v27 int16 = 1
                _ = v27
            default:
                if false {
                    switch v12 {
                    case 9223372036854775807:
                        switch (v2 + ("foo bar" + "")) {
                        case "foo bar":
                            var v28 int = (v12 & v12)
                            _ = v28
                        case "":
                            v29 := v6[(v2 + v2)]
                            _ = v29
                        case "hello":
                            if len(v1) > 0 {
                                v1[0] = (v2 + "foo bar")
                            }
                        }
                    case -663:
                        fmt.Println(v2, v5)
                    default:
                        switch v12 {
                        case -9223372036854775808:
                            if (65535 > 2) {
                                v30 := GenS2{F0: 603, F1: -9223372036854775808, F2: 1000}
                                if (false || false) {
                                    v31 := v6[((v2 + (((("\n\t" + v2) + (v2 + "")) + "") + v2)) + "")]
                                    switch v2 {
                                    case "\n\t":
                                        v12 = 9223372036854775807
                                    default:
                                        switch "\n\t" {
                                        case "\n\t":
                                            fmt.Println(v5)
                                        case "":
                                            var v32 uint8 = (v5 >> ((255) & 63))
                                            _ = v32
                                        case "a":
                                            fmt.Println(v5)
                                        default:
                                            v33 := GenS2{F0: 10, F1: -316, F2: 28}
                                            _ = v33
                                        }
                                    }
                                    _ = v31
                                } else {
                                    v34 := GenS4{F0: 100, F1: true, F2: uint64(v3)}
                                    fmt.Println(v5, v5, v3)
                                    _ = v34
                                }
                                v35 := GenS2{F0: 65535, F1: 42, F2: -32768}
                                _ = v30
                                _ = v35
                            }
                        case 9:
                            fmt.Println(v21, v2, v12)
                        case 42:
                            switch (v12 << ((v3) & 63)) {
                            case 42:
                                var v36 int8 = 12
                                _ = v36
                            default:
                                if true {
                                    v5 -= 128
                                    v37 := GenS4{F0: 42, F1: (!((!false) && (false || (v4 >= v4)))), F2: 2}
                                    v2 = (("hello" + ("hello" + v2)) + v2)
                                    _ = v37
                                }
                            }
                        }
                    }
                } else {
                    var v38 int16 = int16(v12)
                    if true {
                        var v39 uint64 = 4294967295
                        if (42 < 0) {
                            var v40 float64 = 1.5
                            if ((3.14159 / v40) <= float64(float32(v21))) {
                                _ = len(v1)
                                if len(v1) > 0 {
                                    v1[0] = ""
                                }
                                v41 := GenS3{F0: v3}
                                _ = v41
                            } else {
                                if false {
                                    var v42 int16 = (v38 >> ((v3) & 63))
                                    for fi43 := 0; fi43 < 19; fi43++ {
                                        var v44 int16 = v42
                                        _ = fi43
                                        _ = v44
                                    }
                                    _ = v42
                                }
                                v12 = v12
                                switch v12 {
                                case 32767:
                                    v39 -= (v39 + 4294967295)
                                case 42:
                                    v1 = append(v1, v2)
                                case -9223372036854775808:
                                    var v45 float64 = float64((((((1e10 / -2.25) * 1e10) + (0.1 * -2.25)) * -2.25) / ((0.1 * 0.1) / ((1e10 - (-2.25 * 1e10)) + 3.14159))))
                                    _ = v45
                                default:
                                    var v46 bool = false
                                    _ = v46
                                }
                            }
                            _ = v40
                        } else {
                            switch 128 {
                            case 9223372036854775807:
                                var v47 uint16 = 32767
                                _ = v47
                            case -1:
                                switch (v2 + (((v2 + v2) + ((((("a" + v2) + (v2 + v2)) + "hello") + v2) + v2)) + v2)) {
                                case "hello":
                                    if true {
                                        v2 = v2
                                        var v48 int16 = (v38 + (v38 | (v38 | 1)))
                                        _ = v48
                                    } else {
                                        delete(v6, (v2 + "hello"))
                                    }
                                default:
                                    var v49 int = 642
                                    _ = v49
                                }
                            default:
                                if (((true && ((32767 >= (v38 << ((v3) & 63))) || (!false))) && (!false)) || true) {
                                    v12 -= (^(v12))
                                    if (true && true) {
                                        v21 += v21
                                    } else {
                                        v38 = (v38 | v38)
                                        v50 := GenS4{F0: 85, F1: true, F2: 2}
                                        _ = v50
                                    }
                                }
                            }
                        }
                        _ = v39
                    }
                    switch v2 {
                    case "hello":
                        v3 += v4
                    case "\n\t":
                        var v51 float64 = (1e10 / (1.5 * v21))
                        _ = v51
                    }
                    _ = v38
                }
            }
            if false {
                switch (v12 - (v12 ^ (v12 + (v12 << ((916) & 63))))) {
                case 2147483647:
                    fmt.Println(v2)
                }
            }
        }
    }
    switch v2 {
    case "\n\t":
        switch int(v3) {
        case 42:
            v52 := GenS5{F0: 125}
            _ = v52
        case 2147483647:
            v21 += float64(1e10)
        default:
            var v53 float32 = (1e10 - 3.14159)
            _ = v53
        }
    case "a":
        v54 := GenS6{F0: 65535, F1: v5}
        _ = v54
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
