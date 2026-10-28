package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 uint64; F2 uint64 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint8; F1 bool }
func genFn0(p0 uint64, p1 uint8, p2 uint16) int32 {
    fmt.Println(p1, p0, p0)
    p2 = 326
    _ = p0
    _ = p1
    _ = p2
    return 42
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint8 = 255
    var v3 float32 = 0.0
    v4 := make(map[int]int)
    var v5 string = ("hello" + "a")
    var v6 float32 = v3
    switch ((v5 + v5) + ((v5 + "\n\t") + (((v5 + "a") + (v5 + ((v5 + v5) + (v5 + v5)))) + ("hello" + ("foo bar" + "hello"))))) {
    case "foo bar":
        for fi7 := 0; fi7 < 2; fi7++ {
            v8 := v4[fi7]
            _ = len(v1)
            _ = fi7
            _ = v8
        }
    case "\n\t":
        switch v5 {
        case "\n\t":
            if false {
                for fi9 := 0; fi9 < 12; fi9++ {
                    for fi10 := 0; fi10 < 1; fi10++ {
                        if true {
                            v11 := v4[fi10]
                            v12 := []int{-363, fi10, v11, (fi10 ^ fi9)}
                            v2 = (v2 << ((875) & 63))
                            _ = v11
                            _ = v12
                        } else {
                            v2 = 2
                        }
                        switch 579 {
                        case 32767:
                            v13 := GenS1{F0: 1}
                            _ = v13
                        case -518:
                            if (65535 > 711) {
                                fmt.Println(v5, v3)
                            } else {
                                if (1.5 > 1e10) {
                                    var v14 int = fi10
                                    if true {
                                        var v15 uint32 = 999
                                        _ = v15
                                    }
                                    _ = v14
                                }
                                if true {
                                    var v16 int8 = 1
                                    v2 = v2
                                    _ = v16
                                } else {
                                    v17 := GenS2{F0: -93, F1: 0, F2: 341}
                                    if (42 == -85) {
                                        fmt.Println(v3, fi10)
                                        if (!(true && false)) {
                                            fmt.Println(v6, fi9)
                                            fmt.Println(v3)
                                            v18 := GenS1{F0: 822}
                                            _ = v18
                                        }
                                        v19 := GenS2{F0: -71, F1: 42, F2: 1}
                                        _ = v19
                                    } else {
                                        var v20 int = 128
                                        switch v5 {
                                        case "":
                                            if true {
                                                fmt.Println(v20)
                                                v21 := GenS3{F0: true}
                                                switch (v5 + ((v5 + (((v5 + v5) + "a") + v5)) + "a")) {
                                                case "":
                                                    var v22 uint = 18446744073709551615
                                                    _ = v22
                                                case "hello":
                                                    switch v5 {
                                                    case "a":
                                                        var v23 float64 = -2.25
                                                        _ = v23
                                                    case "hello":
                                                        v4[v20] = (v20 * fi10)
                                                    default:
                                                        if (4294967295 > 343) {
                                                            var v24 float32 = v6
                                                            switch (fi10 + -748) {
                                                            case -325:
                                                                v25 := []int{357, -1}
                                                                _ = v25
                                                            default:
                                                                v5 = "\n\t"
                                                            }
                                                            _ = v24
                                                        } else {
                                                            v26 := GenS2{F0: -73, F1: 610, F2: 353}
                                                            if false {
                                                                delete(v1, -2147483648)
                                                            } else {
                                                                if false {
                                                                    fmt.Println(fi10, v5)
                                                                }
                                                                v6 += v3
                                                            }
                                                            var v27 string = v5
                                                            _ = v26
                                                            _ = v27
                                                        }
                                                        fmt.Println(v3, fi10)
                                                        v28 := GenS4{F0: -1}
                                                        v29 := []int{-2147483648, fi10, fi9, int(v2)}
                                                        if false {
                                                            var v30 uint8 = 0
                                                            fmt.Println(v3, v6)
                                                            _ = v30
                                                        }
                                                        var v31 int16 = 42
                                                        _ = v28
                                                        _ = v29
                                                        _ = v31
                                                    }
                                                case "\n\t":
                                                    v6 += v6
                                                default:
                                                    v5 = v5
                                                }
                                                _ = v21
                                            }
                                        default:
                                            v32 := GenS5{F0: (v2 ^ (v2 % (((v2 & (v2 / ((127) | 1)))) | 1))), F1: (1 != 112)}
                                            _ = v32
                                        }
                                        v6 = v6
                                        _ = v20
                                    }
                                    _ = v17
                                }
                            }
                        }
                        var v33 int32 = genFn0(18446744073709551615, v2, 32767)
                        _ = fi10
                        _ = v33
                    }
                    v34 := GenS4{F0: -248}
                    _ = fi9
                    _ = v34
                }
                v35 := []int{}
                _ = v35
            }
        case "":
            fmt.Println(v3, v5, v3)
        default:
            v36 := v4[4294967295]
            _ = v36
        }
    case "":
        if (v2 < 1) {
            var v37 uint = 18446744073709551615
            _ = v37
        } else {
            fmt.Println(v2, v2)
        }
    }
    switch (v5 + (v5 + v5)) {
    case "a":
        switch 32767 {
        case -2147483648:
            fmt.Println(v5, v2)
        case -441:
            if false {
                fmt.Println(v6, v6, v6)
                v3 = v3
            }
        case 32768:
            switch 32768 {
            case -9223372036854775808:
                v2 -= v2
            case -2147483648:
                delete(v1, 32768)
            case 32768:
                var v38 int32 = 2147483647
                _ = v38
            default:
                v39 := []int{}
                _ = v39
            }
        default:
            switch 127 {
            case 714:
                v5 = (v5 + (((("hello" + "\n\t") + v5) + v5) + v5))
            case 524:
                v40 := []int{int(v2), -864, -1, 4294967295}
                _ = v40
            }
        }
    case "foo bar":
        fmt.Println(v6)
    default:
        switch v5 {
        case "\n\t":
            var v41 int8 = 0
            _ = v41
        case "":
            v3 -= ((float32(0.1) / v6) + ((v3 - ((v6 / v3) - v3)) + v6))
        case "a":
            for fi42 := 0; fi42 < 12; fi42++ {
                var v43 int = (fi42 >> ((18446744073709551615) & 63))
                for fi44 := 0; fi44 < 2; fi44++ {
                    fmt.Println(v2, v2)
                    v43 -= v43
                    _ = fi44
                }
                _ = fi42
                _ = v43
            }
        }
    }
    for fi45 := 0; fi45 < 13; fi45++ {
        v46 := GenS4{F0: 0}
        _ = fi45
        _ = v46
    }
    var v47 uint32 = 412
    var v48 uint8 = 0
    if true {
        var v49 string = "foo bar"
        _ = v49
    } else {
        for fi50 := 0; fi50 < 20; fi50++ {
            v47 = (^(v47))
            _ = fi50
        }
    }
    fmt.Println(v5, v48)
    if false {
        v6 = float32((1e10 / 1.5))
        for fi51 := 0; fi51 < 12; fi51++ {
            fmt.Println(v48, fi51, v48)
            fmt.Println(v47, fi51)
            var v52 uint16 = 0
            _ = fi51
            _ = v52
        }
    }
    if (true || false) {
        v47 = v47
        fmt.Println(v2)
    }
    for fi53 := 0; fi53 < 14; fi53++ {
        if (2147483647 != -2147483648) {
            v54 := GenS5{F0: v2, F1: ((fi53 | fi53) <= fi53)}
            _ = v54
        }
        var v55 string = v5
        v56 := []int{(fi53 * fi53), 758, fi53}
        _ = fi53
        _ = v55
        _ = v56
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v48
    _ = v5
    _ = v6
}
