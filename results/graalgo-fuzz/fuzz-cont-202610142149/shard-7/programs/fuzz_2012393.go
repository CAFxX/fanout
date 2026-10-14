package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 uint16; F2 uint16 }
type GenS3 struct { F0 uint32; F1 uint16; F2 bool }
type GenS4 struct { F0 int8; F1 uint16 }
type GenS5 struct { F0 int; F1 int64; F2 string }
type GenS6 struct { F0 uint16; F1 string }
type GenS7 struct { F0 bool; F1 int64; F2 string }
func genFn0(p0 int8, p1 string) bool {
    var v4 int = 128
    v5 := GenS1{F0: 2}
    p0 = (^(p0))
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return ((p0 != (p0 * 93)) || true)
}

func main() {
    var v1 int8 = 1
    v2 := []int{}
    v3 := make(map[int]uint32)
    switch "foo bar" {
    case "foo bar":
        var v6 int16 = 0
        _ = v6
    case "":
        if true {
            var v7 uint16 = 255
            _ = v7
        } else {
            var v8 uint32 = 506
            _ = v8
        }
    case "hello":
        v1 = -106
    default:
        v3[32767] = 428
    }
    var v9 int32 = 376
    v10 := GenS2{F0: 0, F1: 32767, F2: 32767}
    var v11 uint64 = 0
    if (("a" <= ((("" + "foo bar") + (("hello" + "foo bar") + "a")) + ("\n\t" + "hello"))) && false) {
        v11 = 1
        if len(v2) > 0 {
            v2[0] = -9223372036854775808
        }
    }
    v12 := GenS2{F0: 0, F1: 0, F2: 43}
    v1 = (-(v1))
    for fi13 := 0; fi13 < 18; fi13++ {
        fmt.Println(v1, v9, v9)
        var v14 float32 = ((3.14159 * -2.25) / 0.1)
        _ = fi13
        _ = v14
    }
    if false {
        if true {
            var v15 int8 = v1
            v16 := GenS2{F0: 2147483647, F1: 654, F2: 2}
            switch 32767 {
            case 0:
                for fi17 := 0; fi17 < 12; fi17++ {
                    fmt.Println(fi17, v15, v11)
                    _ = fi17
                }
            case -1:
                v9 += (v9 % ((v9) | 1))
            case -994:
                v15 = v15
            default:
                switch -692 {
                case 255:
                    if ((v11 & (^(v11))) > (v11 - v11)) {
                        v3[4294967295] = 4294967295
                    }
                case 0:
                    var v18 bool = (true || false)
                    _ = v18
                case 4294967295:
                    v19 := v3[-883]
                    _ = v19
                }
            }
            _ = v15
            _ = v16
        } else {
            if (!false) {
                fmt.Println(v11)
                var v20 int16 = -1
                v1 += -65
                _ = v20
            } else {
                delete(v3, int(v1))
                v21 := GenS1{F0: 97}
                switch "" {
                case "\n\t":
                    v22 := GenS1{F0: 100}
                    _ = v22
                }
                _ = v21
            }
            v23 := GenS3{F0: 2147483647, F1: 42, F2: true}
            _ = v23
        }
    }
    if true {
        for fi24 := 0; fi24 < 4; fi24++ {
            switch (^(fi24)) {
            case -53:
                for fi25 := 0; fi25 < 10; fi25++ {
                    if true {
                        switch fi24 {
                        case -9223372036854775808:
                            if ((("\n\t" + "foo bar") != ("\n\t" + "")) && ("foo bar" != ("a" + "hello"))) {
                                delete(v3, (fi25 ^ fi24))
                                fmt.Println(fi24)
                            } else {
                                v2 = append(v2, (fi25 % ((42) | 1)))
                                if (2 < 2147483648) {
                                    var v26 int64 = 42
                                    v27 := v3[int(v1)]
                                    _ = v26
                                    _ = v27
                                }
                            }
                        case 42:
                            if (-1 <= (v1 + v1)) {
                                var v28 bool = false
                                v28 = v28
                                v9 = (v9 >> ((18446744073709551615) & 63))
                                v11 += v11
                                if false {
                                    v11 = (v11 ^ v11)
                                    if v28 {
                                        v29 := GenS2{F0: 699, F1: 32767, F2: 32767}
                                        _ = v29
                                    }
                                } else {
                                    v28 = v28
                                    switch "foo bar" {
                                    case "foo bar":
                                        v9 -= (v9 * v9)
                                    case "a":
                                        var v30 uint = 436
                                        _ = v30
                                    }
                                    var v31 float32 = 0.0
                                    _ = v31
                                }
                                if len(v2) > 0 {
                                    v2[0] = 42
                                }
                                if v28 {
                                    switch "foo bar" {
                                    case "hello":
                                        var v32 uint16 = 594
                                        _ = v32
                                    case "\n\t":
                                        v3[fi25] = 0
                                    case "":
                                        fmt.Println(v28, v9, fi24)
                                    default:
                                        switch "" {
                                        case "hello":
                                            v33 := GenS2{F0: 141, F1: 0, F2: 65535}
                                            _ = v33
                                        }
                                    }
                                    var v34 uint = 2
                                    v3[-1] = 65535
                                    _ = v34
                                } else {
                                    var v35 int = fi25
                                    fmt.Println(v1)
                                    v1 = v1
                                    _ = v35
                                }
                                switch ("\n\t" + (("" + "hello") + "")) {
                                case "foo bar":
                                    v2 = append(v2, -645)
                                case "\n\t":
                                    var v36 int32 = v9
                                    _ = v36
                                }
                                _ = v28
                            } else {
                                v11 += v11
                                fmt.Println(fi25)
                            }
                        default:
                            _ = len(v2)
                        }
                        var v37 int32 = v9
                        var v38 float32 = ((0.0 / 1.0) + (1.5 / (1e10 * -2.25)))
                        _ = v37
                        _ = v38
                    } else {
                        var v39 int8 = (v1 - v1)
                        if len(v2) > 0 {
                            v2[0] = (fi25 | (-(fi24)))
                        }
                        _ = v39
                    }
                    _ = fi25
                }
            case -9223372036854775808:
                fmt.Println(fi24)
            default:
                v40 := GenS2{F0: 65535, F1: 292, F2: 32767}
                _ = v40
            }
            var v41 int16 = 934
            for fi42 := 0; fi42 < 10; fi42++ {
                if false {
                    var v43 int64 = 0
                    v9 += (-(v9))
                    switch ("foo bar" + "\n\t") {
                    case "\n\t":
                        if (2 == 351) {
                            if (true || (false || (!false))) {
                                v41 = (v41 + v41)
                            } else {
                                v11 += v11
                                var v44 int32 = 679
                                _ = len(v3)
                                _ = v44
                            }
                            v9 = (v9 - v9)
                        } else {
                            var v45 uint32 = 2147483647
                            var v46 int64 = (v43 * v43)
                            v45 = (^(v45))
                            _ = v45
                            _ = v46
                        }
                    case "hello":
                        var v47 uint8 = 128
                        _ = v47
                    case "a":
                        switch fi42 {
                        case 257:
                            v9 -= (v9 ^ (v9 << ((4294967295) & 63)))
                        }
                    default:
                        v48 := GenS3{F0: 4294967295, F1: 0, F2: false}
                        _ = v48
                    }
                    _ = v43
                } else {
                    var v49 uint = 42
                    _ = v49
                }
                _ = fi42
            }
            _ = fi24
            _ = v41
        }
        v50 := GenS4{F0: -128, F1: 255}
        _ = v50
    }
    for fi51 := 0; fi51 < 12; fi51++ {
        for fi52 := 0; fi52 < 19; fi52++ {
            switch -721 {
            case 2147483647:
                var v53 uint64 = 790
                _ = v53
            case 127:
                _ = len(v3)
            case -2147483648:
                v54 := GenS5{F0: fi51, F1: int64(fi52), F2: (("a" + "\n\t") + "\n\t")}
                _ = v54
            default:
                if false {
                    var v55 float64 = 3.14159
                    delete(v3, -318)
                    v56 := GenS6{F0: 255, F1: "\n\t"}
                    _ = v55
                    _ = v56
                }
            }
            fmt.Println(fi51, v1, fi52)
            var v57 int64 = 42
            _ = fi52
            _ = v57
        }
        _ = fi51
    }
    v58 := GenS7{F0: (((((("a" + ("\n\t" + "foo bar")) + "hello") + "\n\t") + "hello") + ("a" + "hello")) != "\n\t"), F1: -9223372036854775808, F2: "foo bar"}
    v1 += 1
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v58
    _ = v9
}
