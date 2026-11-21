package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int16; F1 uint32 }
type GenS4 struct { F0 int8; F1 bool; F2 uint16 }

func main() {
    v1 := []uint32{751, 2, 2}
    var v2 int16 = 1
    var v3 uint64 = 0
    v4 := make(map[uint32]int)
    var v5 int8 = 100
    if ((42 != -64) || (!true)) {
        v6 := GenS1{F0: 42}
        v2 = -32768
        for fi7 := 0; fi7 < 18; fi7++ {
            if ((!(!(0.1 >= (3.14159 / 1e10)))) || (!(!true))) {
                if true {
                    for fi8 := 0; fi8 < 6; fi8++ {
                        var v9 string = "\n\t"
                        v3 = (v3 ^ v3)
                        var v10 int8 = 1
                        _ = fi8
                        _ = v10
                        _ = v9
                    }
                    if false {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    } else {
                        for fi11 := 0; fi11 < 20; fi11++ {
                            v2 = (v2 >> ((823) & 63))
                            _ = fi11
                        }
                    }
                } else {
                    v12 := GenS2{F0: 1}
                    _ = v12
                }
            } else {
                switch "\n\t" {
                case "foo bar":
                    if (v3 != (v3 << ((18446744073709551615) & 63))) {
                        v3 += v3
                        for fi13 := 0; fi13 < 15; fi13++ {
                            var v14 float32 = 0.1
                            v2 += (v2 % ((v2) | 1))
                            var v15 int = 128
                            _ = fi13
                            _ = v14
                            _ = v15
                        }
                        for fi16 := 0; fi16 < 1; fi16++ {
                            fmt.Println(v5, v5)
                            if false {
                                var v17 int8 = 127
                                switch 1 {
                                case 0:
                                    v17 += v17
                                case 32767:
                                    var v18 uint = 18446744073709551615
                                    _ = v18
                                default:
                                    v2 -= int16(v5)
                                }
                                _ = v17
                            } else {
                                var v19 int64 = -425
                                var v20 int64 = -9223372036854775808
                                v20 = (v20 << ((65535) & 63))
                                _ = v19
                                _ = v20
                            }
                            if (42 == (v3 | 42)) {
                                v21 := GenS3{F0: v2, F1: 2147483648}
                                if false {
                                    var v22 uint16 = 32767
                                    _ = v22
                                }
                                _ = v21
                            } else {
                                v2 += v2
                            }
                            _ = fi16
                        }
                    } else {
                        if len(v1) > 0 {
                            v1[0] = 65535
                        }
                        if false {
                            v5 -= (v5 ^ int8(v3))
                            v2 = (-(v2))
                            for fi23 := 0; fi23 < 19; fi23++ {
                                v2 = v2
                                _ = fi23
                            }
                        } else {
                            for fi24 := 0; fi24 < 17; fi24++ {
                                v25 := GenS1{F0: 8}
                                if (0 != 32768) {
                                    if true {
                                        switch ("hello" + ("\n\t" + "\n\t")) {
                                        case "hello":
                                            if false {
                                                if len(v1) > 0 {
                                                    v1[0] = 65535
                                                }
                                                var v26 uint64 = (^(v3))
                                                _ = v26
                                            } else {
                                                if (772 == 121) {
                                                    fmt.Println(fi7, v5, v3)
                                                } else {
                                                    switch "a" {
                                                    case "":
                                                        v5 -= v5
                                                    default:
                                                        fmt.Println(v3)
                                                    }
                                                    v27 := GenS1{F0: 127}
                                                    if false {
                                                        v28 := GenS4{F0: 127, F1: false, F2: 42}
                                                        v4[2147483648] = (-(fi7))
                                                        v3 += 2
                                                        v5 += 1
                                                        _ = v28
                                                    }
                                                    _ = v27
                                                }
                                            }
                                        default:
                                            var v29 uint32 = 1
                                            _ = v29
                                        }
                                    } else {
                                        if ((-(fi7)) <= fi24) {
                                            var v30 int8 = v5
                                            _ = v30
                                        }
                                    }
                                    fmt.Println(fi24, v5, fi7)
                                    v31 := v4[2]
                                    _ = v31
                                } else {
                                    var v32 float64 = 0.0
                                    _ = v32
                                }
                                _ = fi24
                                _ = v25
                            }
                        }
                    }
                case "\n\t":
                    v5 -= (^(v5))
                }
                var v33 int16 = (v2 * v2)
                _ = v33
            }
            _ = fi7
        }
        _ = v6
    } else {
        if (!false) {
            var v34 int = 127
            _ = v34
        }
        if (2 <= 463) {
            v5 = v5
        } else {
            if false {
                if len(v1) > 0 {
                    v1[0] = 1
                }
                switch "hello" {
                case "\n\t":
                    v3 -= v3
                case "hello":
                    fmt.Println(v2, v5, v5)
                }
            } else {
                v35 := GenS3{F0: v2, F1: uint32(v3)}
                v2 = (v2 | int16(v5))
                _ = v35
            }
            v2 = -1
        }
        v5 = v5
    }
    var v36 uint = 4294967295
    var v37 int8 = (v5 / ((100) | 1))
    v2 = -570
    var v38 string = "hello"
    var v39 int = 362
    var v40 uint = 1
    var v41 int16 = v2
    v41 -= (v41 + (v41 ^ -246))
    fmt.Println(v5, v3, v36)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v2 = v2
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v42 int32 = 306
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
}
