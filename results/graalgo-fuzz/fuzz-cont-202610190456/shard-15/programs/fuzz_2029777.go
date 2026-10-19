package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64; F1 uint16; F2 uint8 }
type GenS3 struct { F0 uint32; F1 bool }
type GenS4 struct { F0 uint32; F1 uint8; F2 int64 }
type GenS5 struct { F0 bool; F1 uint; F2 int16 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int16; F1 uint8; F2 int64 }
func genFn0() uint64 {
    for fi7 := 0; fi7 < 1; fi7++ {
        var v8 uint64 = 18446744073709551615
        var v9 int32 = 1
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 uint32 = 2147483648
    _ = v10
    return 915
}
func genFn1(p0 bool, p1 int16, p2 uint32) uint {
    p1 = p1
    v11 := []int{32768}
    var v12 uint64 = 4294967295
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return 400
}

func main() {
    v1 := make(map[string]uint32)
    var v2 float32 = 3.14159
    v3 := []uint32{786}
    v4 := make(map[int]int)
    var v5 int64 = 9223372036854775807
    var v6 int64 = 0
    for fi13 := 0; fi13 < 9; fi13++ {
        switch "hello" {
        case "foo bar":
            var v14 int = fi13
            _ = v14
        case "hello":
            if (-397 < 1) {
                for fi15 := 0; fi15 < 20; fi15++ {
                    v6 -= -1
                    v3 = append(v3, 0)
                    _ = fi15
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                switch "" {
                case "a":
                    v5 -= (v6 % (((v5 | v5)) | 1))
                default:
                    v16 := GenS1{F0: fi13}
                    _ = v16
                }
            } else {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                v5 = v5
                var v17 int64 = v6
                _ = v17
            }
        case "\n\t":
            v18 := GenS2{F0: 4294967295, F1: 241, F2: 0}
            _ = v18
        default:
            v19 := GenS3{F0: 2147483648, F1: false}
            _ = v19
        }
        _ = fi13
    }
    v20 := GenS3{F0: 65535, F1: (-612 == 128)}
    switch 386 {
    case 9223372036854775807:
        v21 := GenS4{F0: 455, F1: 0, F2: v5}
        _ = v21
    default:
        var v22 uint8 = 127
        _ = v22
    }
    v3 = append(v3, 1)
    fmt.Println(v5)
    fmt.Println(v2)
    v23 := GenS5{F0: (false && true), F1: 571, F2: -32768}
    switch "hello" {
    case "":
        var v24 int8 = -82
        _ = v24
    case "\n\t":
        switch -284 {
        case -1:
            if false {
                if true {
                    _ = len(v4)
                    switch ("hello" + ("\n\t" + ("\n\t" + (("\n\t" + "\n\t") + "hello")))) {
                    case "a":
                        if false {
                            v3 = append(v3, 4294967295)
                            if ((-899 == 0) && (true || false)) {
                                fmt.Println(v6)
                            } else {
                                for fi25 := 0; fi25 < 10; fi25++ {
                                    for fi26 := 0; fi26 < 9; fi26++ {
                                        var v27 float32 = 0.1
                                        v3 = append(v3, 414)
                                        v6 -= (^(v5))
                                        if (!(v6 < (v6 - (v6 << ((65535) & 63))))) {
                                            _ = len(v1)
                                            v28 := GenS4{F0: 2, F1: 2, F2: (v5 % (((v5 * (-(v6)))) | 1))}
                                            v6 += v5
                                            _ = v28
                                        }
                                        _ = fi26
                                        _ = v27
                                    }
                                    v6 = v5
                                    _ = fi25
                                }
                            }
                            v2 = (v2 + v2)
                        }
                    }
                    v29 := GenS4{F0: 42, F1: 11, F2: (v5 >> ((65535) & 63))}
                    _ = v29
                }
                switch -9223372036854775808 {
                case -2147483648:
                    v30 := GenS2{F0: 403, F1: 32767, F2: 0}
                    _ = v30
                default:
                    var v31 int16 = 836
                    _ = v31
                }
                v6 = (v5 << ((450) & 63))
            }
        case 9223372036854775807:
            if len(v3) > 0 {
                v3[0] = 42
            }
        default:
            if true {
                if (42 <= 1000) {
                    var v32 uint64 = 4294967295
                    _ = v32
                } else {
                    var v33 int32 = -630
                    delete(v1, ("foo bar" + "\n\t"))
                    v34 := GenS5{F0: (!(!(!true))), F1: 1, F2: 0}
                    _ = v33
                    _ = v34
                }
                switch "hello" {
                case "":
                    v35 := GenS2{F0: 2, F1: 1, F2: 255}
                    _ = v35
                case "foo bar":
                    for fi36 := 0; fi36 < 20; fi36++ {
                        for fi37 := 0; fi37 < 4; fi37++ {
                            v4[(fi37 >> ((4294967295) & 63))] = (fi36 ^ (fi36 & fi37))
                            _ = fi37
                        }
                        _ = fi36
                    }
                case "\n\t":
                    v6 -= v5
                default:
                    v38 := GenS6{F0: 1}
                    _ = v38
                }
            }
        }
    default:
        if (v2 != ((1e10 / (float32(1e10) + (1e10 + (0.0 - v2)))) / v2)) {
            var v39 uint = 42
            _ = v39
        }
    }
    var v40 uint64 = 18446744073709551615
    v6 -= v6
    v41 := GenS7{F0: -32768, F1: 127, F2: (-(v6))}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
