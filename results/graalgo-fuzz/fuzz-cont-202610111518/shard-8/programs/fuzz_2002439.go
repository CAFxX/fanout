package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 uint; F1 int; F2 uint32 }
type GenS4 struct { F0 int64; F1 uint32; F2 int64 }
type GenS5 struct { F0 uint8; F1 uint32 }
type GenS6 struct { F0 int16; F1 bool }
func genFn0(p0 string, p1 int8) int {
    p0 = p0
    switch p0 {
    case "foo bar":
        p0 = ("hello" + p0)
    default:
        fmt.Println(p0)
    }
    _ = p0
    _ = p1
    return 1
}

func main() {
    var v1 int64 = -9223372036854775808
    v2 := []string{"", "hello", "\n\t"}
    var v3 int64 = (v1 >> ((65535) & 63))
    if true {
        fmt.Println(v3)
    }
    var v4 uint8 = 2
    v2 = append(v2, "a")
    v5 := GenS1{F0: 955}
    for fi6 := 0; fi6 < 20; fi6++ {
        v1 += (v3 / (((-(v3))) | 1))
        _ = len(v2)
        _ = fi6
    }
    var v7 int16 = 1
    if ((int8(v7) < 26) && false) {
        for fi8 := 0; fi8 < 14; fi8++ {
            var v9 uint = 0
            v3 -= -886
            _ = fi8
            _ = v9
        }
        v10 := GenS2{F0: ((1.5 + 1.5) <= 0.0), F1: 0}
        _ = v10
    } else {
        switch 1 {
        case 353:
            switch "hello" {
            case "foo bar":
                if true {
                    v2 = append(v2, "foo bar")
                    switch 2147483647 {
                    case -9223372036854775808:
                        v4 -= (^(v4))
                    case 1:
                        for fi11 := 0; fi11 < 15; fi11++ {
                            if false {
                                for fi12 := 0; fi12 < 15; fi12++ {
                                    fmt.Println(v3)
                                    var v13 uint32 = 42
                                    _ = fi12
                                    _ = v13
                                }
                            } else {
                                for fi14 := 0; fi14 < 5; fi14++ {
                                    v15 := GenS2{F0: ("foo bar" != "hello"), F1: 480}
                                    _ = fi14
                                    _ = v15
                                }
                                if true {
                                    v7 = (^(v7))
                                    v3 = v3
                                    switch (fi11 ^ (fi11 << ((464) & 63))) {
                                    case 35:
                                        v2 = append(v2, "foo bar")
                                    case 32768:
                                        var v16 uint8 = 251
                                        _ = v16
                                    default:
                                        v17 := GenS3{F0: 510, F1: (fi11 - fi11), F2: 564}
                                        _ = v17
                                    }
                                } else {
                                    v2 = append(v2, "")
                                    var v18 float64 = 0.0
                                    v2 = append(v2, "\n\t")
                                    _ = v18
                                }
                            }
                            v4 = (v4 % ((v4) | 1))
                            _ = fi11
                        }
                    case 0:
                        fmt.Println(v7)
                    default:
                        if (4294967295 > 169) {
                            var v19 int32 = 1
                            for fi20 := 0; fi20 < 5; fi20++ {
                                var v21 int64 = v3
                                _ = fi20
                                _ = v21
                            }
                            if len(v2) > 0 {
                                v2[0] = "\n\t"
                            }
                            _ = v19
                        } else {
                            v7 = 42
                        }
                    }
                    v3 -= (v3 & (v3 | v3))
                } else {
                    v22 := GenS1{F0: 269}
                    for fi23 := 0; fi23 < 7; fi23++ {
                        for fi24 := 0; fi24 < 13; fi24++ {
                            if true {
                                var v25 float64 = 0.1
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                var v26 bool = true
                                _ = v25
                                _ = v26
                            }
                            if true {
                                v7 = (v7 | (v7 / (((^(v7))) | 1)))
                                if (!true) {
                                    v27 := GenS2{F0: (v7 <= int16(v4)), F1: 719}
                                    _ = v27
                                }
                            } else {
                                var v28 uint16 = 42
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                _ = v28
                            }
                            v1 += v3
                            _ = fi24
                        }
                        v29 := GenS3{F0: 42, F1: (fi23 ^ fi23), F2: 2}
                        _ = fi23
                        _ = v29
                    }
                    _ = v22
                }
            }
        case -764:
            v7 += (v7 * v7)
        case 42:
            v30 := GenS4{F0: (^(v1)), F1: 2147483647, F2: v1}
            _ = v30
        }
        var v31 string = ""
        _ = v31
    }
    switch 761 {
    case -2147483648:
        v3 = (v3 | (v1 - (v3 | (v1 - v3))))
    case -322:
        v32 := GenS5{F0: 127, F1: 2147483647}
        _ = v32
    }
    v7 += (v7 + (v7 / (((-(v7))) | 1)))
    if (992 < 0) {
        var v33 int = 32767
        _ = v33
    } else {
        fmt.Println(v1, v1)
        for fi34 := 0; fi34 < 19; fi34++ {
            fmt.Println(v1)
            for fi35 := 0; fi35 < 17; fi35++ {
                if ((((fi34 / ((fi35) | 1)) <= int(v7)) && false) || (true && (65535 <= uint32(v4)))) {
                    v3 -= (v1 ^ (v3 % (((^(v1))) | 1)))
                }
                var v36 uint8 = 127
                var v37 float64 = ((3.14159 * 0.0) / 0.1)
                _ = fi35
                _ = v36
                _ = v37
            }
            var v38 uint8 = 42
            _ = fi34
            _ = v38
        }
    }
    var v39 int64 = 42
    if ((1 < 2147483647) && (false && (true || true))) {
        v39 = (v39 << ((1) & 63))
    }
    v40 := GenS6{F0: (v7 - v7), F1: false}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v5
    _ = v7
}
