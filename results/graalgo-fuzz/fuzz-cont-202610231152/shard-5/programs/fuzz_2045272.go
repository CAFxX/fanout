package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint16; F1 int64 }
func genFn0(p0 int64, p1 string, p2 uint32) uint32 {
    fmt.Println(p0)
    v6 := []int{9223372036854775807, 42}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return p2
}

func main() {
    var v1 int = 9223372036854775807
    v2 := []string{""}
    var v3 uint32 = 4294967295
    var v4 uint = 121
    var v5 uint8 = 0
    var v7 uint16 = 42
    v5 = 1
    v7 = 42
    v8 := GenS1{F0: 32767, F1: 100}
    v9 := GenS1{F0: 0, F1: -19}
    if ((false && true) && false) {
        v2 = append(v2, (("a" + "foo bar") + "\n\t"))
    }
    var v10 string = "\n\t"
    if (18446744073709551615 >= 203) {
        v1 = (v1 % (((v1 << (((v4 >> (((v4 >> ((42) & 63))) & 63))) & 63))) | 1))
        var v11 uint64 = 4294967295
        fmt.Println(v10)
        _ = v11
    } else {
        var v12 uint = (v4 >> ((v4) & 63))
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v12
    }
    if true {
        for fi13 := 0; fi13 < 16; fi13++ {
            if (-1 <= -1) {
                v3 += (v3 & (v3 >> ((v4) & 63)))
                switch 127 {
                case 127:
                    v2 = append(v2, (("hello" + ("" + v10)) + v10))
                }
                if ((false && (v10 > v10)) && false) {
                    for fi14 := 0; fi14 < 7; fi14++ {
                        v3 += 817
                        v7 = (v7 | 106)
                        v2 = append(v2, "\n\t")
                        _ = fi14
                    }
                    var v15 string = (v10 + v10)
                    v16 := GenS2{F0: uint16(v4)}
                    _ = v15
                    _ = v16
                }
            } else {
                v2 = append(v2, v10)
                var v17 int64 = int64(v7)
                for fi18 := 0; fi18 < 17; fi18++ {
                    var v19 float64 = 1.5
                    _ = fi18
                    _ = v19
                }
                _ = v17
            }
            var v20 int16 = 42
            var v21 uint32 = (v3 * (v3 << ((v4) & 63)))
            _ = fi13
            _ = v20
            _ = v21
        }
        v2 = append(v2, "a")
        if false {
            v5 = v5
            var v22 int = v1
            v23 := GenS3{F0: 0}
            _ = v22
            _ = v23
        } else {
            if (!false) {
                v2 = append(v2, (v10 + v10))
                var v24 uint64 = 2
                fmt.Println(v7)
                _ = v24
            }
            v5 -= (^(v5))
        }
    } else {
        if (((((38 != 4294967295) && ((0 <= int16(v4)) || true)) || ((v3 | v3) >= v3)) && ((v4 - 65535) <= (v4 << (((v4 + 2)) & 63)))) || ((true || (true && true)) && (0 <= 636))) {
            v25 := GenS3{F0: 42}
            v26 := GenS4{F0: v3}
            _ = v25
            _ = v26
        }
    }
    fmt.Println(v5, v7)
    if (!false) {
        switch "foo bar" {
        case "\n\t":
            v5 += 2
        case "a":
            if (0 == int32(v1)) {
                var v27 uint32 = genFn0(-314, (("\n\t" + (("a" + v10) + v10)) + v10), (^(v3)))
                _ = v27
            } else {
                if (373 > 42) {
                    var v28 uint32 = v3
                    var v29 uint = (v4 >> ((1) & 63))
                    v30 := GenS1{F0: 32, F1: 17}
                    _ = v28
                    _ = v29
                    _ = v30
                } else {
                    for fi31 := 0; fi31 < 19; fi31++ {
                        var v32 uint = v4
                        v5 = v5
                        if true {
                            v32 += 1
                            switch v10 {
                            case "foo bar":
                                if false {
                                    if ((false && true) || true) {
                                        for fi33 := 0; fi33 < 11; fi33++ {
                                            v34 := GenS5{F0: uint16(v5), F1: 9223372036854775807}
                                            var v35 int16 = 0
                                            _ = fi33
                                            _ = v34
                                            _ = v35
                                        }
                                        if true {
                                            var v36 float32 = 0.1
                                            _ = v36
                                        }
                                    }
                                    v37 := GenS4{F0: genFn0(-1, (v10 + (v10 + "a")), (v3 >> ((v32) & 63)))}
                                    fmt.Println(v32)
                                    _ = v37
                                } else {
                                    switch fi31 {
                                    case 42:
                                        var v38 uint8 = 127
                                        _ = v38
                                    default:
                                        v2 = append(v2, "")
                                    }
                                    _ = len(v2)
                                    if len(v2) > 0 {
                                        v2[0] = ("" + (v10 + (("a" + "\n\t") + v10)))
                                    }
                                }
                            case "":
                                var v39 int8 = 102
                                _ = v39
                            }
                            var v40 uint = v4
                            _ = v40
                        }
                        _ = fi31
                        _ = v32
                    }
                    v5 = v5
                }
                v7 -= v7
                v41 := GenS1{F0: 0, F1: -25}
                _ = v41
            }
        default:
            v1 -= (v1 - v1)
        }
        v3 += genFn0(-1, (((v10 + "") + (v10 + (((v10 + "\n\t") + v10) + "a"))) + "foo bar"), (^(v3)))
        if len(v2) > 0 {
            v2[0] = ((v10 + ("\n\t" + "")) + "")
        }
    }
    _ = len(v2)
    if (0 != (^(v7))) {
        for fi42 := 0; fi42 < 5; fi42++ {
            var v43 float32 = 3.14159
            _ = fi42
            _ = v43
        }
    }
    for fi44 := 0; fi44 < 8; fi44++ {
        v45 := GenS3{F0: (v5 % (((v5 << (((v4 / ((v4) | 1))) & 63))) | 1))}
        _ = fi44
        _ = v45
    }
    for fi46 := 0; fi46 < 10; fi46++ {
        var v47 bool = false
        v2 = append(v2, (v10 + ""))
        _ = fi46
        _ = v47
    }
    switch v10 {
    case "":
        var v48 int32 = 1
        _ = v48
    case "foo bar":
        v7 = (v7 - v7)
    case "\n\t":
        if (1e10 == 1e10) {
            v49 := GenS5{F0: (^(v7)), F1: -524}
            _ = v49
        } else {
            switch v10 {
            case "foo bar":
                v7 -= v7
            }
            var v50 uint64 = 519
            _ = v50
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
