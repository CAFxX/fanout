package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 int32 }
type GenS2 struct { F0 int16; F1 uint16; F2 bool }
type GenS3 struct { F0 bool; F1 int8; F2 int64 }

func main() {
    var v1 bool = (-1 <= -86)
    var v2 float32 = (0.0 - (3.14159 - 3.14159))
    var v3 int16 = 1
    var v4 uint = 65535
    var v5 int = 1
    v6 := make(map[uint32]int)
    var v7 string = ("a" + ("hello" + (("a" + ("" + "hello")) + "a")))
    v8 := []int{(v5 >> ((v4) & 63)), v5, v5}
    var v9 float64 = (1e10 / (0.0 - 1.5))
    switch "hello" {
    case "hello":
        _ = len(v6)
    case "foo bar":
        var v10 int32 = -20
        _ = v10
    case "":
        if (v1 || (1 > 14)) {
            var v11 float64 = (-2.25 * v9)
            _ = v11
        } else {
            if (!v1) {
                v12 := GenS1{F0: 1, F1: (2147483648 >= 447), F2: 887}
                switch 32768 {
                case 42:
                    switch v5 {
                    case 127:
                        v13 := GenS1{F0: 42, F1: v1, F2: 2147483647}
                        _ = v13
                    default:
                        var v14 int16 = (v3 * -1)
                        _ = v14
                    }
                case 9223372036854775807:
                    var v15 uint = v4
                    _ = v15
                case 32768:
                    if v1 {
                        var v16 int32 = 940
                        if true {
                            if true {
                                var v17 uint64 = 690
                                fmt.Println(v9, v9, v16)
                                for fi18 := 0; fi18 < 14; fi18++ {
                                    if false {
                                        for fi19 := 0; fi19 < 9; fi19++ {
                                            fmt.Println(v1, v9, v1)
                                            _ = fi19
                                        }
                                        for fi20 := 0; fi20 < 13; fi20++ {
                                            if v1 {
                                                var v21 int8 = 100
                                                _ = v21
                                            }
                                            fmt.Println(v2)
                                            v22 := GenS1{F0: -2147483648, F1: (((!(true && (v2 <= v2))) && true) || (214 <= 212)), F2: v16}
                                            _ = fi20
                                            _ = v22
                                        }
                                    } else {
                                        v4 = 255
                                        switch "foo bar" {
                                        case "hello":
                                            v5 = 247
                                        case "foo bar":
                                            v3 -= v3
                                        case "a":
                                            v9 = v9
                                        default:
                                            _ = len(v8)
                                        }
                                    }
                                    v4 = (v4 * 1)
                                    _ = fi18
                                }
                                _ = v17
                            }
                        } else {
                            var v23 uint8 = 221
                            switch (v5 - int(v3)) {
                            case -240:
                                switch (v5 / ((v5) | 1)) {
                                case -48:
                                    switch (-(v5)) {
                                    case 4294967295:
                                        var v24 int32 = 0
                                        _ = v24
                                    default:
                                        v25 := GenS1{F0: int32(v23), F1: false, F2: v16}
                                        _ = v25
                                    }
                                case -9223372036854775808:
                                    v16 += (v16 % (((v16 % ((int32(v5)) | 1))) | 1))
                                case 384:
                                    v26 := GenS1{F0: (v16 >> (((v4 - (v4 * (v4 ^ (v4 - v4))))) & 63)), F1: true, F2: (v16 / ((695) | 1))}
                                    _ = v26
                                }
                            default:
                                v23 = (v23 << ((255) & 63))
                            }
                            for fi27 := 0; fi27 < 17; fi27++ {
                                var v28 float32 = (v2 / 1.5)
                                for fi29 := 0; fi29 < 3; fi29++ {
                                    v16 = v16
                                    var v30 uint = 18446744073709551615
                                    _ = fi29
                                    _ = v30
                                }
                                _ = fi27
                                _ = v28
                            }
                            _ = v23
                        }
                        switch (v5 & 879) {
                        case 4294967295:
                            v2 = v2
                        }
                        _ = v16
                    } else {
                        v31 := GenS1{F0: 761, F1: v1, F2: -315}
                        _ = v31
                    }
                }
                for fi32 := 0; fi32 < 14; fi32++ {
                    v33 := GenS2{F0: (v3 | v3), F1: 440, F2: (!(v1 && (!v1)))}
                    _ = fi32
                    _ = v33
                }
                _ = v12
            } else {
                switch v5 {
                case 128:
                    var v34 string = "\n\t"
                    _ = v34
                case -1:
                    v9 = (1e10 + v9)
                case 0:
                    _ = len(v8)
                }
                v35 := GenS2{F0: int16(v5), F1: 1, F2: (-1 > (v5 & v5))}
                var v36 float32 = (v2 / (0.0 - 1.5))
                _ = v35
                _ = v36
            }
            var v37 int16 = v3
            _ = v37
        }
    }
    fmt.Println(v5, v1, v1)
    v1 = v1
    for fi38 := 0; fi38 < 3; fi38++ {
        if v1 {
            for fi39 := 0; fi39 < 20; fi39++ {
                v40 := GenS3{F0: v1, F1: 100, F2: 42}
                fmt.Println(v7)
                if false {
                    var v41 int = fi39
                    _ = v41
                }
                _ = fi39
                _ = v40
            }
            fmt.Println(v9, v4)
            v2 -= (1.5 - v2)
        }
        v2 = float32(-2.25)
        _ = fi38
    }
    v2 -= 0.1
    v9 = v9
    v5 = (-(v5))
    v42 := v6[42]
    if v1 {
        for fi43 := 0; fi43 < 3; fi43++ {
            delete(v6, 0)
            _ = fi43
        }
    } else {
        switch v42 {
        case 361:
            v9 -= (v9 * v9)
        case -2147483648:
            delete(v6, 624)
        default:
            v44 := GenS3{F0: (!((v1 && (1 == 42)) && ((-136 > -361) && (v42 <= v42)))), F1: 100, F2: -718}
            _ = v44
        }
        fmt.Println(v2, v1)
        fmt.Println(v1, v5, v9)
    }
    v45 := v6[32]
    var v46 uint64 = 2
    var v47 int16 = v3
    v8 = append(v8, v5)
    for fi48 := 0; fi48 < 8; fi48++ {
        for fi49 := 0; fi49 < 13; fi49++ {
            v2 -= v2
            _ = fi49
        }
        switch (fi48 >> ((380) & 63)) {
        case 4294967295:
            var v50 int32 = -1
            _ = v50
        case 1:
            var v51 uint64 = v46
            _ = v51
        case 2147483647:
            var v52 int32 = 860
            _ = v52
        }
        _ = fi48
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
