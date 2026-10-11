package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint64; F1 uint16; F2 uint16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 bool; F1 uint8 }
type GenS5 struct { F0 bool; F1 int64; F2 int }
func genFn0(p0 uint32, p1 int32) uint16 {
    var v7 uint = 0
    p1 += (p1 << ((v7) & 63))
    _ = p0
    _ = p1
    _ = v7
    return 467
}
func genFn1(p0 string) uint8 {
    p0 = (p0 + p0)
    var v8 float32 = (float32(0.1) - 0.0)
    var v9 int32 = 2147483647
    v9 -= (-(v9))
    fmt.Println(v9)
    _ = p0
    _ = v8
    _ = v9
    return 2
}

func main() {
    var v1 int16 = -1
    v2 := []int{}
    var v3 uint = 534
    var v4 string = "a"
    var v5 int16 = 32767
    var v6 int32 = -2147483648
    for fi10 := 0; fi10 < 11; fi10++ {
        var v11 int = 0
        _ = fi10
        _ = v11
    }
    v6 = (v6 << ((v3) & 63))
    fmt.Println(v1)
    for fi12 := 0; fi12 < 15; fi12++ {
        v6 += v6
        if true {
            for fi13 := 0; fi13 < 2; fi13++ {
                v6 = (v6 >> ((2) & 63))
                _ = fi13
            }
            for fi14 := 0; fi14 < 3; fi14++ {
                v3 -= 42
                _ = fi14
            }
            var v15 float64 = -2.25
            _ = v15
        }
        _ = fi12
    }
    fmt.Println(v6, v6, v1)
    switch v4 {
    case "a":
        v16 := GenS1{F0: 51}
        _ = v16
    }
    fmt.Println(v3, v5)
    var v17 uint64 = 0
    for fi18 := 0; fi18 < 4; fi18++ {
        if len(v2) > 0 {
            v2[0] = int(v17)
        }
        v19 := GenS1{F0: 111}
        switch "" {
        case "foo bar":
            for fi20 := 0; fi20 < 9; fi20++ {
                v3 = 255
                v21 := GenS1{F0: 127}
                var v22 int = fi20
                var v23 string = ""
                switch (v22 & (-(v22))) {
                case -2147483648:
                    v2 = append(v2, 0)
                case 32768:
                    if false {
                        switch "a" {
                        case "a":
                            switch (v22 << ((525) & 63)) {
                            case -9223372036854775808:
                                switch "hello" {
                                case "\n\t":
                                    fmt.Println(v4, v5, v6)
                                case "foo bar":
                                    v24 := GenS2{F0: (v17 << (((v3 | (v3 / ((2) | 1)))) & 63)), F1: 0, F2: 65535}
                                    _ = v24
                                }
                            case 776:
                                fmt.Println(v17, v3)
                            case 128:
                                v4 = v4
                            }
                        }
                        switch v23 {
                        case "a":
                            fmt.Println(v4, v22)
                        }
                    } else {
                        var v25 uint32 = 2
                        v26 := GenS1{F0: genFn1((v4 + "hello"))}
                        _ = v25
                        _ = v26
                    }
                case -1:
                    v1 += v1
                default:
                    switch (-(v22)) {
                    case -613:
                        if (65535 == 166) {
                            v3 = v3
                        } else {
                            switch "\n\t" {
                            case "hello":
                                v27 := GenS1{F0: 248}
                                _ = v27
                            default:
                                var v28 float64 = 1.5
                                _ = v28
                            }
                            var v29 int8 = 119
                            _ = v29
                        }
                    case -2147483648:
                        v30 := GenS3{F0: ""}
                        _ = v30
                    default:
                        if true {
                            var v31 uint32 = 81
                            v6 -= int32(v3)
                            v6 -= (v6 & 2147483647)
                            _ = v31
                        } else {
                            v23 = ("hello" + "\n\t")
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            fmt.Println(v22)
                            v3 += v3
                        }
                    }
                    if false {
                        _ = len(v2)
                        fmt.Println(v3, v23)
                    } else {
                        if ((genFn0(2147483647, 42) != 32767) || (float64(1e10) != -2.25)) {
                            _ = len(v2)
                            switch 1 {
                            case 9223372036854775807:
                                switch (-(fi18)) {
                                case 1:
                                    var v32 uint16 = 1
                                    _ = v32
                                case 117:
                                    _ = len(v2)
                                default:
                                    fmt.Println(v1)
                                }
                            case 32768:
                                v4 = ((v4 + "foo bar") + v23)
                            default:
                                if false {
                                    v1 -= (v5 * v5)
                                    var v33 int32 = (v6 | -662)
                                    if (!((v4 + (((v23 + "hello") + v4) + "a")) <= v4)) {
                                        var v34 uint64 = (v17 << ((4294967295) & 63))
                                        _ = v34
                                    }
                                    _ = v33
                                }
                            }
                            var v35 uint64 = 2
                            var v36 uint16 = 32767
                            var v37 int8 = 0
                            _ = v35
                            _ = v36
                            _ = v37
                        } else {
                            v38 := GenS4{F0: false, F1: genFn1((("foo bar" + "foo bar") + "foo bar"))}
                            fmt.Println(v5, fi20, v23)
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = v38
                        }
                        v6 = (^(v6))
                    }
                }
                _ = fi20
                _ = v21
                _ = v22
                _ = v23
            }
        case "":
            switch "hello" {
            case "foo bar":
                fmt.Println(v1, v17)
            case "hello":
                v39 := GenS3{F0: ("hello" + (v4 + v4))}
                _ = v39
            }
        default:
            for fi40 := 0; fi40 < 13; fi40++ {
                v41 := GenS5{F0: (true && false), F1: 1, F2: fi40}
                _ = fi40
                _ = v41
            }
        }
        _ = fi18
        _ = v19
    }
    v6 += 2147483647
    v42 := GenS2{F0: (v17 | (v17 - 18446744073709551615)), F1: 838, F2: 32767}
    v43 := GenS1{F0: 255}
    fmt.Println(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
