package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 uint64; F2 uint16 }
type GenS3 struct { F0 int16; F1 int16; F2 uint64 }
type GenS4 struct { F0 int32; F1 uint; F2 int16 }
type GenS5 struct { F0 uint16; F1 int64 }

func main() {
    var v1 int32 = -597
    v2 := make(map[int]uint32)
    v3 := []uint32{1, 474, 42, 2147483647}
    for fi4 := 0; fi4 < 8; fi4++ {
        v5 := GenS1{F0: 1}
        _ = fi4
        _ = v5
    }
    v1 -= v1
    v1 = (v1 - v1)
    if (2 > 0) {
        for fi6 := 0; fi6 < 13; fi6++ {
            var v7 bool = (2 >= 777)
            for fi8 := 0; fi8 < 4; fi8++ {
                var v9 uint16 = 255
                _ = fi8
                _ = v9
            }
            v1 = v1
            _ = fi6
            _ = v7
        }
    }
    v10 := GenS1{F0: 2}
    for fi11 := 0; fi11 < 5; fi11++ {
        if (true || true) {
            for fi12 := 0; fi12 < 1; fi12++ {
                var v13 float32 = float32(-2.25)
                _ = fi12
                _ = v13
            }
            switch "foo bar" {
            case "hello":
                var v14 int8 = 1
                _ = v14
            default:
                if (!true) {
                    if (1 != 900) {
                        v15 := v2[(fi11 * 9223372036854775807)]
                        _ = v15
                    } else {
                        fmt.Println(fi11, fi11)
                        v1 += (v1 * 42)
                        var v16 int8 = 42
                        _ = v16
                    }
                    var v17 int16 = 32767
                    _ = v17
                } else {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
            }
            delete(v2, fi11)
        } else {
            v1 = (v1 - 1)
            v1 = (v1 & v1)
            if (0 <= 42) {
                if false {
                    v1 += v1
                    if (255 != 2) {
                        var v18 float64 = (((float64((1e10 + (-2.25 / 1.0))) - 3.14159) * -2.25) + 0.1)
                        switch -375 {
                        case 2147483647:
                            for fi19 := 0; fi19 < 10; fi19++ {
                                v2[fi11] = 263
                                _ = fi19
                            }
                        case 480:
                            fmt.Println(v1)
                        }
                        var v20 int16 = -719
                        _ = v18
                        _ = v20
                    }
                    fmt.Println(fi11)
                } else {
                    v21 := GenS2{F0: 255, F1: 0, F2: 65535}
                    var v22 int = (fi11 & (fi11 / ((fi11) | 1)))
                    var v23 int8 = 42
                    _ = v21
                    _ = v22
                    _ = v23
                }
                for fi24 := 0; fi24 < 11; fi24++ {
                    var v25 int8 = 100
                    v1 += (v1 - -1)
                    var v26 int16 = -86
                    _ = fi24
                    _ = v25
                    _ = v26
                }
            } else {
                if (32767 != -337) {
                    if false {
                        if true {
                            for fi27 := 0; fi27 < 5; fi27++ {
                                v28 := GenS3{F0: 1, F1: -737, F2: 18446744073709551615}
                                _ = fi27
                                _ = v28
                            }
                            fmt.Println(v1, v1)
                        } else {
                            var v29 int8 = -98
                            _ = v29
                        }
                        for fi30 := 0; fi30 < 1; fi30++ {
                            v1 = (^(v1))
                            var v31 float64 = 1e10
                            var v32 bool = true
                            _ = fi30
                            _ = v31
                            _ = v32
                        }
                    }
                }
                var v33 int = (fi11 << ((0) & 63))
                _ = v33
            }
        }
        if ((^(fi11)) != (fi11 * 32767)) {
            switch "\n\t" {
            case "\n\t":
                var v34 int8 = 42
                _ = v34
            case "hello":
                if (65535 >= 2) {
                    var v35 float64 = (-2.25 - 1e10)
                    v36 := GenS4{F0: -1, F1: 2, F2: 0}
                    _ = v35
                    _ = v36
                }
            case "foo bar":
                var v37 float64 = ((0.1 * (1e10 / 1.0)) * -2.25)
                _ = v37
            }
        } else {
            var v38 int16 = 32767
            switch fi11 {
            case -9223372036854775808:
                switch ("foo bar" + "foo bar") {
                case "":
                    v1 = 2147483647
                case "foo bar":
                    var v39 int64 = 42
                    _ = v39
                default:
                    v1 -= (v1 / ((v1) | 1))
                }
            case 2147483647:
                if ((v1 << ((2) & 63)) < 1) {
                    var v40 int = -658
                    _ = v40
                } else {
                    for fi41 := 0; fi41 < 1; fi41++ {
                        v38 = v38
                        _ = fi41
                    }
                    var v42 int = -760
                    _ = v42
                }
            default:
                v43 := v2[(fi11 - fi11)]
                _ = v43
            }
            _ = v38
        }
        if ((fi11 & (fi11 / (((fi11 / ((2147483647) | 1))) | 1))) <= (fi11 << ((1) & 63))) {
            v1 += (v1 * v1)
            var v44 uint8 = 109
            _ = v44
        } else {
            _ = len(v3)
            var v45 bool = (1000 <= -515)
            _ = v45
        }
        _ = fi11
    }
    v46 := GenS5{F0: 42, F1: -1}
    var v47 float64 = 1.5
    v48 := v2[0]
    fmt.Println(v48, v1, v47)
    switch 42 {
    case -1:
        fmt.Println(v1)
    }
    var v49 bool = (222 == 127)
    var v50 int = 9223372036854775807
    v49 = ((v1 & v1) <= (v1 >> ((511) & 63)))
    if len(v3) > 0 {
        v3[0] = (^(v48))
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v50
}
