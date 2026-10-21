package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int32; F1 string; F2 int16 }
type GenS4 struct { F0 int64; F1 int16 }

func main() {
    var v1 int = 0
    v2 := []uint32{2147483648, 2147483647}
    var v3 float32 = 0.0
    var v4 int16 = 0
    var v5 int16 = (v4 % ((1) | 1))
    var v6 float32 = (float32(3.14159) + v3)
    if (873 != 519) {
        fmt.Println(v5)
        for fi7 := 0; fi7 < 17; fi7++ {
            switch fi7 {
            case 1:
                fmt.Println(v6)
            case 32767:
                fmt.Println(v4, v3, fi7)
            }
            switch "foo bar" {
            case "\n\t":
                var v8 float64 = 0.1
                _ = v8
            default:
                v9 := GenS1{F0: 127, F1: -17}
                _ = v9
            }
            v1 -= v1
            _ = fi7
        }
        var v10 uint = 2
        _ = v10
    } else {
        var v11 int16 = v5
        switch v1 {
        case 0:
            v12 := GenS2{F0: ""}
            _ = v12
        }
        switch v1 {
        case 9223372036854775807:
            if true {
                var v13 int32 = -2147483648
                var v14 bool = ((1.5 - (v3 * v6)) < v6)
                v15 := GenS3{F0: (v13 & v13), F1: "", F2: v11}
                _ = v13
                _ = v14
                _ = v15
            }
        case 42:
            if true {
                var v16 bool = (!((127 == 1) || (v1 < v1)))
                _ = v16
            }
        case 128:
            v17 := GenS1{F0: 0, F1: 1}
            _ = v17
        default:
            var v18 int8 = 100
            _ = v18
        }
        _ = v11
    }
    if false {
        switch v1 {
        case -9223372036854775808:
            if (-280 != 42) {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
        case 387:
            v4 = (v5 - -32768)
        case 792:
            if (!true) {
                fmt.Println(v4)
                fmt.Println(v5, v1)
            } else {
                var v19 uint16 = 42
                switch v1 {
                case 0:
                    v2 = append(v2, 65535)
                }
                if false {
                    _ = len(v2)
                    var v20 int8 = -93
                    if (!((!false) || true)) {
                        var v21 float32 = (0.1 + ((v6 * (((1.5 - float32((0.1 / 3.14159))) - float32(1.5)) - 0.1)) + v3))
                        _ = v21
                    }
                    _ = v20
                } else {
                    if ("a" != "foo bar") {
                        var v22 float64 = 0.1
                        var v23 float64 = 1.5
                        switch 127 {
                        case 32768:
                            v23 -= (0.1 - float64(float32(v22)))
                        case 127:
                            var v24 uint8 = 255
                            _ = v24
                        }
                        _ = v22
                        _ = v23
                    } else {
                        var v25 int64 = -414
                        _ = v25
                    }
                }
                _ = v19
            }
        default:
            var v26 int16 = (v4 << ((4294967295) & 63))
            _ = v26
        }
    } else {
        var v27 uint = 42
        if len(v2) > 0 {
            v2[0] = 815
        }
        var v28 int32 = 42
        _ = v27
        _ = v28
    }
    fmt.Println(v6, v5)
    fmt.Println(v3, v6)
    var v29 int32 = -1
    for fi30 := 0; fi30 < 13; fi30++ {
        switch (fi30 >> ((4294967295) & 63)) {
        case -2147483648:
            if len(v2) > 0 {
                v2[0] = 61
            }
        case 128:
            v29 += 0
        }
        var v31 uint64 = 2
        _ = fi30
        _ = v31
    }
    switch v1 {
    case 9223372036854775807:
        for fi32 := 0; fi32 < 11; fi32++ {
            switch "foo bar" {
            case "a":
                switch "" {
                case "a":
                    v29 = -1
                case "":
                    v3 = float32(1.5)
                }
            case "\n\t":
                v5 += 42
            case "hello":
                v4 = 1000
            default:
                if false {
                    switch ("hello" + "\n\t") {
                    case "a":
                        for fi33 := 0; fi33 < 18; fi33++ {
                            v34 := GenS2{F0: "a"}
                            var v35 uint32 = 2147483647
                            v3 -= 1e10
                            var v36 int8 = 127
                            if ((fi33 + (-(fi32))) == int(v4)) {
                                v37 := GenS1{F0: 39, F1: v36}
                                fmt.Println(v5, v6, v35)
                                _ = v37
                            }
                            var v38 float64 = -2.25
                            var v39 uint64 = 18446744073709551615
                            v40 := GenS2{F0: "a"}
                            _ = fi33
                            _ = v34
                            _ = v35
                            _ = v36
                            _ = v38
                            _ = v39
                            _ = v40
                        }
                    case "\n\t":
                        v5 += (v4 & 661)
                    default:
                        var v41 int = v1
                        _ = v41
                    }
                }
            }
            _ = fi32
        }
    case -271:
        for fi42 := 0; fi42 < 17; fi42++ {
            if true {
                switch fi42 {
                case 1:
                    for fi43 := 0; fi43 < 3; fi43++ {
                        v44 := GenS2{F0: "a"}
                        var v45 float64 = float64((v6 + (v6 / 3.14159)))
                        _ = fi43
                        _ = v44
                        _ = v45
                    }
                case 128:
                    v6 += v3
                case 129:
                    switch (fi42 - (fi42 | (fi42 % ((1) | 1)))) {
                    case -391:
                        for fi46 := 0; fi46 < 7; fi46++ {
                            v29 = -2147483648
                            v2 = append(v2, 1)
                            v47 := GenS4{F0: 1, F1: v5}
                            _ = fi46
                            _ = v47
                        }
                    case 2147483647:
                        var v48 uint8 = 143
                        _ = v48
                    }
                default:
                    var v49 int16 = (-(v4))
                    _ = v49
                }
                v29 -= (v29 >> ((251) & 63))
            } else {
                var v50 int16 = v5
                _ = v50
            }
            v5 -= v5
            _ = len(v2)
            _ = fi42
        }
    case 1:
        var v51 int8 = 1
        _ = v51
    default:
        var v52 int8 = -5
        _ = v52
    }
    switch "hello" {
    case "a":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    case "foo bar":
        v6 = v6
    case "\n\t":
        var v53 uint64 = 65
        _ = v53
    default:
        v6 = ((3.14159 * 0.1) - (v6 * v6))
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
