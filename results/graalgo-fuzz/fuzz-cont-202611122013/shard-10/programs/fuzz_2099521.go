package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64; F1 string; F2 int32 }
type GenS3 struct { F0 uint32 }

func main() {
    v1 := []uint32{0, 2, 42}
    v2 := make(map[uint32]uint32)
    var v3 uint64 = 2
    v4 := make(map[int]bool)
    v5 := make(map[uint32]bool)
    for fi6 := 0; fi6 < 10; fi6++ {
        if true {
            switch fi6 {
            case 631:
                v3 = 0
            case 645:
                fmt.Println(fi6)
            }
            var v7 float32 = (0.0 + 0.0)
            _ = v7
        }
        for fi8 := 0; fi8 < 2; fi8++ {
            v9 := GenS1{F0: 0}
            _ = fi8
            _ = v9
        }
        v3 += v3
        _ = fi6
    }
    v3 = 18446744073709551615
    var v10 int64 = 0
    if false {
        for fi11 := 0; fi11 < 19; fi11++ {
            v12 := GenS1{F0: 0}
            if (0 <= 2147483647) {
                var v13 float64 = float64(0.0)
                for fi14 := 0; fi14 < 5; fi14++ {
                    if true {
                        v3 = v3
                        v10 = (v10 << ((18446744073709551615) & 63))
                        v10 += v10
                    } else {
                        if false {
                            v15 := GenS1{F0: 398}
                            v1 = append(v1, 856)
                            var v16 uint = 42
                            _ = v15
                            _ = v16
                        }
                        if false {
                            v10 = (v10 + 9223372036854775807)
                            if (98 <= 1) {
                                var v17 int64 = (v10 / ((v10) | 1))
                                var v18 uint64 = (v3 >> ((2) & 63))
                                v19 := GenS1{F0: 65535}
                                _ = v17
                                _ = v18
                                _ = v19
                            }
                        } else {
                            v10 += (^(v10))
                        }
                        var v20 int64 = (v10 << ((282) & 63))
                        _ = v20
                    }
                    v21 := GenS1{F0: 65535}
                    _ = fi14
                    _ = v21
                }
                for fi22 := 0; fi22 < 17; fi22++ {
                    fmt.Println(v10)
                    switch "a" {
                    case "":
                        if true {
                            fmt.Println(fi11)
                            _ = len(v1)
                        }
                    case "hello":
                        switch fi22 {
                        case 42:
                            var v23 string = ("foo bar" + "a")
                            _ = v23
                        }
                    }
                    if (1 != 127) {
                        if (!(258 < 2)) {
                            var v24 int32 = 42
                            var v25 int16 = -760
                            var v26 float32 = (((((0.0 - float32(v13)) - 0.1) + 1.5) + (0.0 - 0.1)) * 0.1)
                            _ = v24
                            _ = v25
                            _ = v26
                        } else {
                            v10 = v10
                        }
                    } else {
                        v1 = append(v1, 400)
                        fmt.Println(fi11)
                    }
                    var v27 uint64 = 42
                    fmt.Println(v27, v3)
                    var v28 uint = 1
                    _ = fi22
                    _ = v27
                    _ = v28
                }
                _ = v13
            } else {
                switch (-(fi11)) {
                case 127:
                    v5[0] = (1e10 >= 0.1)
                case -2147483648:
                    var v29 int = (fi11 | 2147483647)
                    _ = v29
                case 32767:
                    v3 = (^(v3))
                default:
                    v3 -= (v3 / (((v3 * (v3 >> ((534) & 63)))) | 1))
                }
                switch fi11 {
                case 4294967295:
                    var v30 int64 = v10
                    _ = v30
                }
                if (!(!(-128 <= 0))) {
                    var v31 int16 = -32768
                    _ = v31
                } else {
                    switch (fi11 << ((2) & 63)) {
                    case 127:
                        v32 := v5[42]
                        _ = v32
                    case -1:
                        if (false || false) {
                            var v33 int8 = 8
                            _ = v33
                        }
                    default:
                        for fi34 := 0; fi34 < 11; fi34++ {
                            var v35 float32 = 0.1
                            var v36 float64 = -2.25
                            fmt.Println(fi34)
                            _ = fi34
                            _ = v35
                            _ = v36
                        }
                    }
                }
            }
            v1 = append(v1, 2)
            _ = fi11
            _ = v12
        }
        var v37 uint8 = 1
        _ = v37
    } else {
        var v38 uint8 = 127
        v39 := GenS2{F0: uint64(v38), F1: (("" + "foo bar") + ("\n\t" + ("a" + "\n\t"))), F2: 0}
        _ = v38
        _ = v39
    }
    if false {
        v2[42] = 227
        if (true || false) {
            fmt.Println(v10)
            for fi40 := 0; fi40 < 16; fi40++ {
                v41 := GenS2{F0: v3, F1: "hello", F2: 1}
                var v42 uint8 = 2
                _ = fi40
                _ = v41
                _ = v42
            }
            fmt.Println(v10, v3)
        }
    }
    v10 = (v10 & (v10 * (-(v10))))
    if ((v10 * (v10 >> ((4294967295) & 63))) < (v10 * (v10 % ((-9223372036854775808) | 1)))) {
        fmt.Println(v3, v3)
    }
    var v43 uint8 = 255
    v44 := GenS3{F0: 342}
    v45 := GenS2{F0: v3, F1: "foo bar", F2: -2147483648}
    var v46 bool = (!(127 != uint8(v3)))
    var v47 uint = 2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v47
    _ = v5
}
