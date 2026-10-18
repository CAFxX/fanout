package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16; F2 uint16 }
type GenS2 struct { F0 uint8; F1 int32 }
type GenS3 struct { F0 bool; F1 int64; F2 int32 }
type GenS4 struct { F0 int }

func main() {
    var v1 int = 127
    v2 := make(map[uint32]uint32)
    var v3 uint64 = 4294967295
    var v4 float32 = 3.14159
    v5 := []int{v1, -9223372036854775808, 447, (v1 << ((255) & 63))}
    var v6 bool = true
    for fi7 := 0; fi7 < 17; fi7++ {
        for fi8 := 0; fi8 < 1; fi8++ {
            var v9 float32 = float32(1e10)
            _ = fi8
            _ = v9
        }
        v10 := GenS1{F0: -53, F1: 840, F2: 2}
        _ = fi7
        _ = v10
    }
    for fi11 := 0; fi11 < 8; fi11++ {
        if v6 {
            if v6 {
                switch v1 {
                case 128:
                    fmt.Println(v4, v6)
                case 676:
                    var v12 bool = false
                    _ = v12
                case -919:
                    var v13 int32 = 1
                    _ = v13
                }
                switch (-(fi11)) {
                case 1:
                    v4 = v4
                case -557:
                    v14 := GenS1{F0: -56, F1: uint16(v3), F2: 343}
                    _ = v14
                case -521:
                    v4 += ((v4 / (v4 + float32(1.5))) + (v4 - (v4 / 1e10)))
                default:
                    for fi15 := 0; fi15 < 10; fi15++ {
                        v6 = (!v6)
                        var v16 uint64 = (v3 / ((v3) | 1))
                        var v17 uint8 = 0
                        var v18 int = fi11
                        _ = fi15
                        _ = v16
                        _ = v17
                        _ = v18
                    }
                }
                var v19 int = v1
                _ = v19
            }
            fmt.Println(v1)
            v20 := v2[481]
            _ = v20
        }
        v4 = 3.14159
        _ = fi11
    }
    v3 -= 1
    if (-1 >= -1) {
        v4 -= (v4 / v4)
        v1 -= v1
        var v21 float32 = 1.5
        _ = v21
    }
    if (!(int(v3) < (v1 << ((65535) & 63)))) {
        for fi22 := 0; fi22 < 18; fi22++ {
            switch "" {
            case "foo bar":
                var v23 int16 = 1
                _ = v23
            case "hello":
                for fi24 := 0; fi24 < 18; fi24++ {
                    switch (("" + "hello") + "hello") {
                    case "foo bar":
                        var v25 uint32 = 737
                        _ = v25
                    case "hello":
                        v4 -= v4
                    case "\n\t":
                        v1 += fi22
                    default:
                        var v26 float32 = -2.25
                        _ = v26
                    }
                    v27 := GenS1{F0: -1, F1: 1, F2: 495}
                    _ = fi24
                    _ = v27
                }
            case "\n\t":
                if len(v5) > 0 {
                    v5[0] = -2147483648
                }
            default:
                for fi28 := 0; fi28 < 6; fi28++ {
                    v6 = (!v6)
                    var v29 int32 = 2147483647
                    switch v1 {
                    case 4294967295:
                        var v30 int16 = 396
                        _ = v30
                    case 2147483647:
                        v3 = (v3 << ((65535) & 63))
                        var v31 int8 = 42
                        _ = v31
                    case 128:
                        v1 = (fi28 | (v1 & (^(fi22))))
                    default:
                        fmt.Println(v4, fi22, fi28)
                    }
                    _ = fi28
                    _ = v29
                }
            }
            _ = fi22
        }
        var v32 bool = (!(v6 && ((false || v6) && (-32768 > 32767))))
        _ = v32
    } else {
        v2[4294967295] = 42
        v33 := GenS1{F0: 0, F1: 65535, F2: 42}
        v1 += 691
        _ = v33
    }
    if v6 {
        fmt.Println(v1, v4)
    } else {
        fmt.Println(v4, v6)
        var v34 int16 = -1
        v35 := GenS1{F0: -1, F1: 892, F2: 65535}
        _ = v34
        _ = v35
    }
    _ = len(v2)
    for fi36 := 0; fi36 < 11; fi36++ {
        v37 := GenS2{F0: 127, F1: -1}
        _ = fi36
        _ = v37
    }
    v38 := v2[65535]
    if false {
        fmt.Println(v4, v4)
    }
    switch v1 {
    case -1:
        if (((((v4 + -2.25) * v4) + (float32((((1.5 + (1e10 + 0.1)) + 1.5) * 0.0)) + ((0.1 - ((v4 * v4) * (v4 + v4))) - ((float32(3.14159) - 3.14159) / v4)))) / v4) < 0.0) {
            if true {
                for fi39 := 0; fi39 < 10; fi39++ {
                    switch (v1 | fi39) {
                    case -9223372036854775808:
                        if (42 > 255) {
                            v40 := GenS3{F0: true, F1: -499, F2: 1}
                            v2[42] = v38
                            if (-114 >= -128) {
                                v41 := GenS3{F0: (!true), F1: -1, F2: 0}
                                var v42 bool = ("" >= "hello")
                                var v43 uint16 = 255
                                _ = v41
                                _ = v42
                                _ = v43
                            }
                            _ = v40
                        }
                    }
                    _ = fi39
                }
                v44 := GenS4{F0: v1}
                var v45 bool = v6
                _ = v44
                _ = v45
            }
        }
    }
    v46 := GenS3{F0: v6, F1: 910, F2: 1}
    v38 += uint32(v3)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v46
    _ = v5
    _ = v6
}
