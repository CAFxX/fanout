package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 uint64 }
type GenS2 struct { F0 string; F1 uint64 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 string; F1 int8 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint; F1 uint8 }
type GenS7 struct { F0 int8; F1 bool }
type GenS8 struct { F0 uint16; F1 uint; F2 uint32 }
func genFn0(p0 int16, p1 uint16, p2 int16) uint16 {
    fmt.Println(p1, p1, p1)
    var v7 uint = 42
    var v8 float64 = (0.0 - float64(1.5))
    p1 -= (p1 % ((uint16(v7)) | 1))
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return p1
}
func genFn1(p0 uint16, p1 int32, p2 int8) uint64 {
    p0 += p0
    switch "a" {
    case "":
        for fi9 := 0; fi9 < 11; fi9++ {
            var v10 string = "foo bar"
            _ = fi9
            _ = v10
        }
    case "a":
        p1 -= 42
    }
    if true {
        p0 -= genFn0(-362, genFn0(-32768, (p0 ^ p0), 962), 42)
        v11 := []int{}
        v11 = append(v11, 1)
        _ = v11
    }
    for fi12 := 0; fi12 < 5; fi12++ {
        v13 := GenS1{F0: 0, F1: 42, F2: 4294967295}
        v14 := GenS1{F0: 127, F1: 141, F2: 18446744073709551615}
        _ = fi12
        _ = v13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    v1 := []int{979, 4294967295}
    var v2 int32 = 896
    var v3 uint16 = 42
    var v4 int32 = (v2 << ((1) & 63))
    var v5 uint = 4294967295
    v6 := []string{}
    var v15 uint64 = 0
    _ = len(v6)
    for fi16 := 0; fi16 < 9; fi16++ {
        switch "a" {
        case "foo bar":
            v17 := GenS2{F0: "foo bar", F1: (^(v15))}
            _ = v17
        case "":
            if (v2 != (^(v4))) {
                if (false && (22 <= 334)) {
                    fmt.Println(v5)
                    for fi18 := 0; fi18 < 6; fi18++ {
                        v19 := GenS3{F0: uint64(v5)}
                        switch "" {
                        case "":
                            switch (fi16 << ((v5) & 63)) {
                            case 2147483647:
                                v1 = append(v1, (fi18 & (^(fi16))))
                            default:
                                var v20 int16 = 0
                                _ = v20
                            }
                        case "hello":
                            switch 127 {
                            case 0:
                                switch 127 {
                                case 0:
                                    _ = len(v1)
                                case 4294967295:
                                    v3 += (^(v3))
                                case -1:
                                    var v21 uint32 = 2147483647
                                    _ = v21
                                default:
                                    v15 = genFn1((v3 << ((v5) & 63)), (v4 * 1), 42)
                                }
                            default:
                                var v22 int = (-(fi16))
                                _ = v22
                            }
                        case "a":
                            if true {
                                if len(v6) >= 2 {
                                    v6 = v6[1:len(v6)]
                                }
                                if false {
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                    if ("\n\t" >= "a") {
                                        if (((3.14159 / (-2.25 + 0.0)) == 1e10) || (0.1 >= -2.25)) {
                                            var v23 uint = 18446744073709551615
                                            fmt.Println(v23, v15, v2)
                                            _ = v23
                                        }
                                        v4 = (-(v2))
                                        v2 -= v4
                                    } else {
                                        if false {
                                            switch ("a" + "foo bar") {
                                            case "hello":
                                                if (v5 == v5) {
                                                    var v24 float64 = -2.25
                                                    var v25 uint32 = 534
                                                    v2 -= v4
                                                    _ = v24
                                                    _ = v25
                                                } else {
                                                    var v26 uint32 = 736
                                                    _ = v26
                                                }
                                                fmt.Println(v5)
                                            case "\n\t":
                                                v27 := GenS4{F0: ("\n\t" + ("hello" + "")), F1: 43}
                                                _ = v27
                                            case "foo bar":
                                                v4 -= (v2 * -1)
                                            }
                                            fmt.Println(v15)
                                            _ = len(v1)
                                        } else {
                                            if len(v6) > 0 {
                                                v6[0] = "hello"
                                            }
                                            v28 := GenS4{F0: "foo bar", F1: 110}
                                            v1 = append(v1, fi16)
                                            _ = v28
                                        }
                                        if (498 == -1) {
                                            var v29 int64 = -9223372036854775808
                                            _ = len(v1)
                                            v4 -= -1
                                            if ((false || false) || true) {
                                                _ = len(v1)
                                                v1 = append(v1, fi16)
                                                var v30 uint64 = (v15 * (^(v15)))
                                                _ = v30
                                            }
                                            _ = v29
                                        } else {
                                            if true {
                                                if len(v6) > 0 {
                                                    v6[0] = "\n\t"
                                                }
                                                if true {
                                                    if len(v6) >= 2 {
                                                        v6 = v6[1:len(v6)]
                                                    }
                                                }
                                            } else {
                                                fmt.Println(fi16, v2)
                                                if false {
                                                    v5 = 65535
                                                    v31 := GenS5{F0: 76}
                                                    _ = v31
                                                } else {
                                                    v6 = append(v6, "a")
                                                    var v32 uint16 = v3
                                                    _ = v32
                                                }
                                                switch (("\n\t" + "") + "foo bar") {
                                                case "\n\t":
                                                    v33 := GenS6{F0: (v5 % ((4294967295) | 1)), F1: 128}
                                                    _ = v33
                                                default:
                                                    if ((^(v3)) <= (v3 >> ((42) & 63))) {
                                                        fmt.Println(v2, v5)
                                                        var v34 uint = 42
                                                        fmt.Println(v5)
                                                        _ = v34
                                                    } else {
                                                        if len(v6) > 0 {
                                                            v6[0] = "\n\t"
                                                        }
                                                        fmt.Println(fi16, fi18, fi18)
                                                        v35 := GenS7{F0: -1, F1: true}
                                                        _ = v35
                                                    }
                                                }
                                            }
                                            switch fi16 {
                                            case 4294967295:
                                                switch fi18 {
                                                case 9223372036854775807:
                                                    v5 = v5
                                                case 128:
                                                    var v36 uint32 = 2
                                                    _ = v36
                                                case 2147483647:
                                                    var v37 int = (^(fi16))
                                                    _ = v37
                                                default:
                                                    var v38 int16 = 1000
                                                    _ = v38
                                                }
                                            case -9223372036854775808:
                                                v4 = (v2 - (v2 >> ((v5) & 63)))
                                            }
                                        }
                                        v4 = 0
                                    }
                                } else {
                                    if len(v6) >= 2 {
                                        v6 = v6[1:len(v6)]
                                    }
                                }
                                var v39 float64 = 1e10
                                _ = v39
                            }
                        default:
                            if false {
                                v2 = 770
                                fmt.Println(v2)
                            }
                        }
                        _ = fi18
                        _ = v19
                    }
                }
            } else {
                v5 = (v5 * v5)
                v3 += genFn0(1, 414, 42)
                v40 := GenS8{F0: v3, F1: v5, F2: 0}
                _ = v40
            }
        case "hello":
            for fi41 := 0; fi41 < 18; fi41++ {
                v2 = 2147483647
                var v42 uint64 = uint64(v3)
                fmt.Println(v3, v4, v3)
                _ = fi41
                _ = v42
            }
        default:
            v5 = (v5 << (((v5 - (v5 | (^(v5))))) & 63))
        }
        switch fi16 {
        case 9223372036854775807:
            var v43 uint = (^(v5))
            _ = v43
        case -2147483648:
            v4 = -2147483648
        default:
            v2 = (v2 ^ v4)
        }
        var v44 int64 = 42
        _ = fi16
        _ = v44
    }
    var v45 uint = v5
    var v46 uint32 = 170
    v5 += 65535
    v3 = v3
    _ = len(v1)
    var v47 int = -2147483648
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v6
}
