package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 bool }
type GenS3 struct { F0 int32; F1 int16; F2 int }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint32; F1 bool }
type GenS6 struct { F0 int8; F1 int8; F2 int }
type GenS7 struct { F0 uint; F1 string; F2 uint8 }

func main() {
    var v1 string = "foo bar"
    var v2 float32 = 1e10
    var v3 float64 = 1e10
    var v4 int8 = 100
    var v5 uint16 = 42
    v3 = 1e10
    v4 += 127
    v6 := []int{-242, -2147483648, 0}
    var v7 int8 = (v4 >> ((140) & 63))
    switch 42 {
    case -647:
        for fi8 := 0; fi8 < 15; fi8++ {
            switch (v1 + ("a" + v1)) {
            case "":
                var v9 float64 = (float64(v2) - float64(v2))
                _ = v9
            case "\n\t":
                if true {
                    fmt.Println(v1)
                    if (-2147483648 > 42) {
                        v10 := GenS1{F0: 42}
                        for fi11 := 0; fi11 < 18; fi11++ {
                            var v12 int16 = 409
                            if false {
                                v3 = (v3 + (v3 * v3))
                                switch 0 {
                                case 978:
                                    v13 := GenS1{F0: 0}
                                    _ = v13
                                case -1:
                                    if false {
                                        var v14 uint = 741
                                        var v15 int32 = 0
                                        _ = v14
                                        _ = v15
                                    } else {
                                        var v16 uint = 42
                                        var v17 int16 = 42
                                        _ = v16
                                        _ = v17
                                    }
                                }
                            } else {
                                _ = len(v6)
                            }
                            _ = fi11
                            _ = v12
                        }
                        _ = v10
                    }
                    v2 = v2
                }
            case "hello":
                switch "" {
                case "foo bar":
                    for fi18 := 0; fi18 < 5; fi18++ {
                        switch "\n\t" {
                        case "\n\t":
                            fmt.Println(v3)
                        default:
                            v5 = (v5 / ((666) | 1))
                        }
                        if (1e10 != v2) {
                            v7 -= (v4 - (v7 % ((v4) | 1)))
                        } else {
                            if false {
                                var v19 float64 = v3
                                if (32767 > 255) {
                                    switch fi18 {
                                    case 128:
                                        var v20 int16 = 0
                                        _ = v20
                                    case 0:
                                        if len(v6) >= 2 {
                                            v6 = v6[1:len(v6)]
                                        }
                                    case 299:
                                        fmt.Println(v1, fi18, v4)
                                    default:
                                        var v21 int16 = 42
                                        _ = v21
                                    }
                                }
                                v2 = v2
                                _ = v19
                            }
                        }
                        if true {
                            v7 = v4
                        }
                        _ = fi18
                    }
                }
            default:
                v22 := GenS1{F0: -858}
                _ = v22
            }
            v5 -= v5
            _ = fi8
        }
    default:
        if false {
            for fi23 := 0; fi23 < 15; fi23++ {
                v1 = v1
                _ = fi23
            }
            var v24 int16 = -32768
            _ = v24
        } else {
            for fi25 := 0; fi25 < 3; fi25++ {
                for fi26 := 0; fi26 < 8; fi26++ {
                    var v27 bool = false
                    var v28 int = (fi25 ^ (fi25 + (fi25 | (fi25 >> ((2) & 63)))))
                    v5 += v5
                    _ = fi26
                    _ = v27
                    _ = v28
                }
                var v29 float32 = 0.0
                v30 := GenS2{F0: 957, F1: (true || ((9223372036854775807 <= 42) || ((fi25 ^ (fi25 >> ((77) & 63))) == 2147483647)))}
                _ = fi25
                _ = v29
                _ = v30
            }
            var v31 uint8 = 127
            if (!false) {
                for fi32 := 0; fi32 < 7; fi32++ {
                    var v33 int32 = 287
                    for fi34 := 0; fi34 < 2; fi34++ {
                        if true {
                            if (-1 <= 0) {
                                v4 = 42
                                if false {
                                    var v35 bool = true
                                    if len(v6) > 0 {
                                        v6[0] = (^(fi34))
                                    }
                                    _ = v35
                                }
                                v33 -= (v33 ^ (-(v33)))
                            }
                            if false {
                                switch ((v1 + (v1 + v1)) + ((("foo bar" + "hello") + v1) + v1)) {
                                case "a":
                                    var v36 float64 = (v3 - v3)
                                    _ = v36
                                case "foo bar":
                                    v31 = (v31 << ((42) & 63))
                                    var v37 uint16 = v5
                                    fmt.Println(v7, v7, v37)
                                    _ = v37
                                default:
                                    var v38 int64 = -9223372036854775808
                                    _ = v38
                                }
                                if false {
                                    v39 := GenS3{F0: (v33 << ((65535) & 63)), F1: -32768, F2: fi34}
                                    v33 = -1
                                    _ = v39
                                } else {
                                    fmt.Println(v5, v2)
                                }
                                v31 += uint8(v5)
                                v2 = (float32((v3 - 0.0)) / 1.5)
                                v40 := GenS4{F0: (fi32 >> ((2) & 63))}
                                _ = v40
                            } else {
                                if ((0 == -32768) && (true && false)) {
                                    var v41 uint = uint(v5)
                                    _ = v41
                                } else {
                                    v4 = (v7 | v4)
                                }
                                if true {
                                    if len(v6) > 0 {
                                        v6[0] = fi34
                                    }
                                }
                                v3 -= v3
                            }
                            fmt.Println(fi32, v7, v31)
                        }
                        var v42 string = (v1 + v1)
                        v6 = append(v6, 2147483647)
                        _ = fi34
                        _ = v42
                    }
                    _ = fi32
                    _ = v33
                }
            }
            _ = v31
        }
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v43 int32 = 0
    var v44 int8 = (v7 / ((1) | 1))
    _ = len(v6)
    fmt.Println(v7, v2, v44)
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    for fi45 := 0; fi45 < 12; fi45++ {
        switch (fi45 / (((^(fi45))) | 1)) {
        case 829:
            if (1 == 2147483647) {
                var v46 int32 = (v43 - v43)
                if true {
                    var v47 uint32 = 0
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    v48 := GenS1{F0: (v43 | v43)}
                    _ = v47
                    _ = v48
                } else {
                    switch (fi45 * fi45) {
                    case -2147483648:
                        fmt.Println(v46, v43)
                    case -1:
                        v49 := GenS5{F0: 2147483648, F1: ((v2 <= v2) && false)}
                        _ = v49
                    }
                    v4 = 81
                }
                _ = v46
            }
        case -9223372036854775808:
            v5 = v5
        case 2147483647:
            if (!(true || (true && true))) {
                var v50 uint32 = 103
                for fi51 := 0; fi51 < 16; fi51++ {
                    v6 = append(v6, fi45)
                    v52 := GenS6{F0: (v4 / (((-(v44))) | 1)), F1: int8(v43), F2: (fi51 * 128)}
                    _ = fi51
                    _ = v52
                }
                fmt.Println(v5, v44, v5)
                _ = v50
            }
        default:
            var v53 string = v1
            _ = v53
        }
        for fi54 := 0; fi54 < 13; fi54++ {
            _ = len(v6)
            _ = fi54
        }
        v55 := GenS7{F0: 871, F1: v1, F2: 42}
        _ = fi45
        _ = v55
    }
    v56 := GenS2{F0: (v5 ^ v5), F1: false}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v56
    _ = v6
    _ = v7
}
