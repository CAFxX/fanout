package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 bool; F1 uint64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int16; F1 int8 }

func main() {
    v1 := []uint32{42}
    var v2 int = 127
    v3 := make(map[string]int)
    var v4 int8 = 0
    for fi5 := 0; fi5 < 14; fi5++ {
        v6 := GenS1{F0: 128}
        for fi7 := 0; fi7 < 18; fi7++ {
            v4 -= v4
            var v8 int8 = v4
            if (42 <= (v8 << ((2) & 63))) {
                v9 := GenS1{F0: 118}
                v3[("foo bar" + "\n\t")] = 214
                if len(v1) > 0 {
                    v1[0] = 2
                }
                v1 = append(v1, 2147483648)
                _ = v9
            }
            _ = fi7
            _ = v8
        }
        if false {
            for fi10 := 0; fi10 < 6; fi10++ {
                v2 = 32767
                var v11 uint64 = 125
                fmt.Println(v11)
                _ = fi10
                _ = v11
            }
            var v12 uint16 = 1
            if len(v1) > 0 {
                v1[0] = uint32(v12)
            }
            _ = v12
        }
        _ = fi5
        _ = v6
    }
    if (32767 <= 1) {
        if true {
            fmt.Println(v2, v4)
            v13 := GenS2{F0: ("\n\t" < "hello"), F1: -337}
            var v14 int64 = 1
            _ = v13
            _ = v14
        } else {
            v2 = v2
        }
    } else {
        switch 2147483647 {
        case 128:
            if (!false) {
                if (!((32767 == 23) || (1 <= 948))) {
                    var v15 int8 = (v4 * v4)
                    _ = v15
                }
                var v16 bool = false
                _ = v16
            } else {
                for fi17 := 0; fi17 < 7; fi17++ {
                    var v18 bool = (0.1 >= float32(3.14159))
                    _ = fi17
                    _ = v18
                }
                if false {
                    v4 = 0
                    var v19 uint8 = 128
                    _ = v19
                } else {
                    v20 := GenS2{F0: false, F1: 0}
                    _ = v20
                }
                v21 := GenS3{F0: (1 > 667), F1: 957}
                _ = v21
            }
        case -2147483648:
            var v22 float64 = float64((float32(((3.14159 / 0.1) * ((float64(0.0) + 3.14159) + -2.25))) * 1e10))
            _ = v22
        case 1:
            v23 := GenS4{F0: 2}
            _ = v23
        }
        for fi24 := 0; fi24 < 15; fi24++ {
            for fi25 := 0; fi25 < 6; fi25++ {
                fmt.Println(v2, fi24, v2)
                _ = fi25
            }
            if false {
                for fi26 := 0; fi26 < 4; fi26++ {
                    if true {
                        var v27 int16 = 1
                        if len(v1) > 0 {
                            v1[0] = 2
                        }
                        _ = v27
                    } else {
                        v4 += (-(v4))
                        v4 -= (v4 - 100)
                    }
                    v2 -= (-(fi24))
                    if false {
                        switch fi26 {
                        case -360:
                            if (128 <= 42) {
                                v28 := GenS2{F0: true, F1: -1}
                                var v29 int = fi26
                                _ = v28
                                _ = v29
                            } else {
                                v4 = v4
                                switch "a" {
                                case "":
                                    fmt.Println(v4)
                                }
                                if (223 <= 128) {
                                    var v30 uint8 = 1
                                    var v31 int64 = 9223372036854775807
                                    v31 += v31
                                    _ = v30
                                    _ = v31
                                } else {
                                    v4 += (^(v4))
                                    fmt.Println(v2, fi26)
                                    var v32 uint64 = 42
                                    _ = v32
                                }
                                v2 += v2
                            }
                        case 0:
                            switch (fi26 % (((fi26 % ((fi26) | 1))) | 1)) {
                            case -707:
                                if true {
                                    var v33 string = ("" + ("\n\t" + ""))
                                    v4 -= (v4 / ((v4) | 1))
                                    _ = v33
                                } else {
                                    var v34 int = (-(fi26))
                                    switch "\n\t" {
                                    case "hello":
                                        var v35 uint32 = 373
                                        _ = v35
                                    case "\n\t":
                                        var v36 bool = true
                                        _ = v36
                                    case "":
                                        v2 = (fi26 + v34)
                                    default:
                                        var v37 bool = (1 <= -2147483648)
                                        _ = v37
                                    }
                                    _ = v34
                                }
                            case 127:
                                var v38 float32 = 3.14159
                                _ = v38
                            case 216:
                                if true {
                                    v2 = (-(fi24))
                                } else {
                                    v4 += 1
                                    v39 := GenS5{F0: 42, F1: (^(v4))}
                                    var v40 uint8 = 96
                                    v40 += (v40 / ((v40) | 1))
                                    _ = v39
                                    _ = v40
                                }
                            }
                        case 128:
                            var v41 int8 = (v4 ^ (-(v4)))
                            _ = v41
                        }
                    }
                    _ = fi26
                }
            }
            _ = fi24
        }
    }
    v4 = 0
    v2 = (v2 >> ((854) & 63))
    v2 += (^(v2))
    v1 = append(v1, 42)
    var v42 uint8 = 21
    if (float64(1e10) <= (3.14159 - 1e10)) {
        v4 = v4
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
}
