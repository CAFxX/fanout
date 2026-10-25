package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16; F2 int32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint64 }

func main() {
    var v1 int = 4294967295
    var v2 int16 = 42
    v3 := make(map[int]uint32)
    var v4 string = ""
    v1 -= 9223372036854775807
    switch -787 {
    case 2147483647:
        v1 = 0
    case -9223372036854775808:
        if (4294967295 == 18446744073709551615) {
            v4 = "a"
            var v5 float32 = (3.14159 * -2.25)
            _ = v5
        }
    default:
        v1 += v1
    }
    v4 = v4
    v6 := []int{32768}
    var v7 int64 = 730
    if (!false) {
        for fi8 := 0; fi8 < 7; fi8++ {
            for fi9 := 0; fi9 < 15; fi9++ {
                v4 = "\n\t"
                _ = fi9
            }
            v1 += (fi8 >> ((448) & 63))
            v10 := GenS1{F0: 32767, F1: 42, F2: -2147483648}
            _ = fi8
            _ = v10
        }
        if (float32(0.0) >= 1e10) {
            for fi11 := 0; fi11 < 12; fi11++ {
                var v12 int = fi11
                _ = fi11
                _ = v12
            }
            if ((v2 == v2) && (((v7 | 0) != v7) || (2 < 44))) {
                v7 = (v7 - v7)
                for fi13 := 0; fi13 < 9; fi13++ {
                    v7 = (v7 & 42)
                    if true {
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        fmt.Println(v7, v7, v1)
                    }
                    if (!((fi13 ^ fi13) < (-(fi13)))) {
                        fmt.Println(fi13, v2, v4)
                        v14 := GenS2{F0: 42}
                        var v15 uint16 = 42
                        _ = v14
                        _ = v15
                    } else {
                        v1 += int(v7)
                    }
                    _ = fi13
                }
                var v16 bool = (411 < 186)
                _ = v16
            }
            if true {
                v4 = "foo bar"
                if false {
                    for fi17 := 0; fi17 < 7; fi17++ {
                        switch v4 {
                        case "a":
                            v18 := GenS1{F0: 255, F1: 255, F2: -1}
                            _ = v18
                        }
                        _ = fi17
                    }
                }
                switch (v4 + "\n\t") {
                case "hello":
                    fmt.Println(v4, v2)
                case "\n\t":
                    v2 = (v2 << ((255) & 63))
                case "foo bar":
                    var v19 int = v1
                    _ = v19
                }
            } else {
                fmt.Println(v2, v1)
            }
        }
    } else {
        if (80 != 255) {
            for fi20 := 0; fi20 < 12; fi20++ {
                var v21 bool = false
                _ = fi20
                _ = v21
            }
            var v22 string = ("\n\t" + "foo bar")
            var v23 uint16 = 463
            _ = v22
            _ = v23
        }
        fmt.Println(v4, v7)
    }
    v24 := GenS3{F0: 1}
    for fi25 := 0; fi25 < 10; fi25++ {
        if false {
            if (0 > 2) {
                var v26 int8 = 0
                _ = v26
            } else {
                _ = len(v3)
                if (-32768 != 32767) {
                    _ = len(v3)
                    var v27 uint8 = 127
                    _ = v27
                }
            }
            switch 1 {
            case 4294967295:
                switch v4 {
                case "":
                    switch (v1 ^ (v1 + (v1 >> ((0) & 63)))) {
                    case 525:
                        var v28 float32 = 1e10
                        _ = v28
                    case 560:
                        for fi29 := 0; fi29 < 9; fi29++ {
                            v7 = (v7 % ((v7) | 1))
                            _ = fi29
                        }
                    case -249:
                        v30 := v3[1]
                        _ = v30
                    }
                case "\n\t":
                    var v31 int8 = -128
                    _ = v31
                default:
                    v1 += (-(fi25))
                }
            default:
                v32 := GenS1{F0: 2, F1: 1, F2: -328}
                _ = v32
            }
            switch v1 {
            case -2147483648:
                switch 441 {
                case 32767:
                    var v33 int8 = 127
                    _ = v33
                case 1:
                    switch (^(v1)) {
                    case 1:
                        v1 -= (v1 + (-(fi25)))
                    case 32768:
                        delete(v3, v1)
                    default:
                        v34 := v3[(fi25 | (v1 + (fi25 | (v1 >> ((1) & 63)))))]
                        _ = v34
                    }
                default:
                    for fi35 := 0; fi35 < 10; fi35++ {
                        v2 += v2
                        v7 -= -1
                        var v36 float32 = ((0.0 / (float32(1e10) + -2.25)) - -2.25)
                        var v37 bool = (2147483647 <= 4294967295)
                        var v38 int32 = 1
                        _ = fi35
                        _ = v36
                        _ = v37
                        _ = v38
                    }
                }
            }
        } else {
            v2 -= (v2 / (((^(v2))) | 1))
            switch fi25 {
            case 42:
                delete(v3, (^(v1)))
            case -929:
                if false {
                    switch (v4 + "a") {
                    case "hello":
                        var v39 string = v4
                        _ = v39
                    case "\n\t":
                        v40 := GenS3{F0: 182}
                        _ = v40
                    default:
                        v2 -= 764
                    }
                } else {
                    var v41 uint8 = 121
                    var v42 int16 = (v2 % ((v2) | 1))
                    var v43 bool = true
                    _ = v41
                    _ = v42
                    _ = v43
                }
            default:
                var v44 int16 = (v2 & -32768)
                _ = v44
            }
            switch (v1 | fi25) {
            case 0:
                for fi45 := 0; fi45 < 12; fi45++ {
                    var v46 float32 = 0.1
                    var v47 float32 = v46
                    v1 += fi25
                    _ = fi45
                    _ = v46
                    _ = v47
                }
            case 2147483647:
                v7 = v7
            case 42:
                var v48 uint8 = 0
                _ = v48
            }
        }
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = fi25
    }
    fmt.Println(v2, v2, v2)
    if (true && (!true)) {
        var v49 int = (v1 / ((v1) | 1))
        v6 = append(v6, (v49 / (((v49 / ((v49) | 1))) | 1)))
        switch v4 {
        case "a":
            _ = len(v3)
        }
        _ = v49
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v6
    _ = v7
}
