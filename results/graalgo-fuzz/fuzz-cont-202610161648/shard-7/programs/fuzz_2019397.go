package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 int16; F1 uint }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 bool; F1 int64 }
type GenS7 struct { F0 uint64; F1 uint }

func main() {
    var v1 float64 = float64((3.14159 - ((1e10 * float32((0.1 + 0.1))) * 1.5)))
    var v2 int = -9223372036854775808
    var v3 uint8 = 128
    var v4 uint64 = 18446744073709551615
    var v5 uint64 = (v4 - v4)
    if false {
        v3 -= (v3 - (v3 + v3))
        v2 += (v2 >> ((255) & 63))
        var v6 uint16 = 893
        _ = v6
    } else {
        if false {
            v7 := []int{9223372036854775807, (v2 % ((v2) | 1)), (^(v2)), v2}
            v8 := GenS1{F0: true}
            var v9 bool = ((v5 ^ 42) > 811)
            _ = v7
            _ = v8
            _ = v9
        } else {
            var v10 uint32 = 2
            _ = v10
        }
        v4 -= 647
    }
    v11 := []int{1, v2}
    switch v2 {
    case 661:
        v1 -= 1.5
    }
    for fi12 := 0; fi12 < 8; fi12++ {
        if false {
            var v13 int8 = 127
            _ = v13
        }
        _ = fi12
    }
    var v14 int8 = 0
    if true {
        switch "foo bar" {
        case "a":
            if false {
                v3 = v3
                switch v2 {
                case -121:
                    switch ("" + "") {
                    case "\n\t":
                        if len(v11) >= 2 {
                            v11 = v11[1:len(v11)]
                        }
                    case "foo bar":
                        if (0 > 1) {
                            if len(v11) >= 2 {
                                v11 = v11[1:len(v11)]
                            }
                            for fi15 := 0; fi15 < 17; fi15++ {
                                v2 -= (^(fi15))
                                v1 -= v1
                                _ = fi15
                            }
                            v16 := GenS1{F0: true}
                            _ = v16
                        } else {
                            if len(v11) > 0 {
                                v11[0] = 998
                            }
                        }
                    case "":
                        switch ("foo bar" + (("\n\t" + ("\n\t" + "hello")) + "")) {
                        case "a":
                            fmt.Println(v2, v3, v3)
                        case "foo bar":
                            v1 = (v1 / float64(float32(0.1)))
                        case "":
                            switch v2 {
                            case 128:
                                var v17 int64 = 732
                                _ = v17
                            case -250:
                                var v18 int = 2147483647
                                _ = v18
                            }
                        }
                    default:
                        for fi19 := 0; fi19 < 11; fi19++ {
                            v20 := GenS2{F0: 672, F1: -1}
                            _ = fi19
                            _ = v20
                        }
                    }
                default:
                    for fi21 := 0; fi21 < 9; fi21++ {
                        switch (v2 % ((-1) | 1)) {
                        case -221:
                            v2 = (fi21 << ((0) & 63))
                        }
                        fmt.Println(v5)
                        v14 += v14
                        _ = fi21
                    }
                }
                _ = len(v11)
            } else {
                v22 := GenS2{F0: -1, F1: -1}
                _ = v22
            }
        case "":
            _ = len(v11)
        default:
            v3 = (v3 & (v3 * (v3 + 128)))
        }
    } else {
        if false {
            var v23 uint8 = v3
            v23 -= v23
            _ = v23
        }
        if (false && false) {
            if (0.1 < (3.14159 * 1.5)) {
                v5 -= (v5 | (v5 / (((v4 / (((^(v4))) | 1))) | 1)))
            } else {
                v1 = 1e10
            }
            v14 += v14
            if ((88 == 989) || false) {
                for fi24 := 0; fi24 < 20; fi24++ {
                    v25 := GenS3{F0: 113}
                    _ = fi24
                    _ = v25
                }
            }
        } else {
            for fi26 := 0; fi26 < 19; fi26++ {
                var v27 uint = 1
                for fi28 := 0; fi28 < 2; fi28++ {
                    if true {
                        var v29 uint64 = v4
                        v30 := GenS1{F0: true}
                        v31 := GenS4{F0: -592, F1: v27}
                        _ = v29
                        _ = v30
                        _ = v31
                    }
                    var v32 string = ("\n\t" + "")
                    switch "hello" {
                    case "foo bar":
                        var v33 bool = false
                        _ = v33
                    case "\n\t":
                        var v34 int32 = 2147483647
                        _ = v34
                    case "hello":
                        v5 = v4
                        var v35 uint16 = 732
                        _ = v35
                    }
                    _ = fi28
                    _ = v32
                }
                v5 = 31
                _ = fi26
                _ = v27
            }
        }
        var v36 int8 = -128
        _ = v36
    }
    fmt.Println(v4, v1)
    v37 := GenS5{F0: 42}
    for fi38 := 0; fi38 < 10; fi38++ {
        var v39 int = 0
        v40 := GenS5{F0: 42}
        for fi41 := 0; fi41 < 10; fi41++ {
            var v42 uint = 65535
            v3 = v3
            fmt.Println(v2, fi41, v5)
            _ = fi41
            _ = v42
        }
        _ = fi38
        _ = v39
        _ = v40
    }
    v3 -= (v3 % ((2) | 1))
    if (!true) {
        switch "hello" {
        case "hello":
            if len(v11) > 0 {
                v11[0] = (v2 | (v2 - (-(v2))))
            }
        default:
            if true {
                if (-32768 <= 755) {
                    var v43 uint32 = 65535
                    var v44 uint16 = 32767
                    var v45 string = ("\n\t" + "hello")
                    _ = v43
                    _ = v44
                    _ = v45
                } else {
                    for fi46 := 0; fi46 < 7; fi46++ {
                        v2 += fi46
                        v3 = v3
                        _ = fi46
                    }
                }
            } else {
                if len(v11) > 0 {
                    v11[0] = (v2 | (v2 * (^(v2))))
                }
                v47 := GenS4{F0: 170, F1: 519}
                for fi48 := 0; fi48 < 15; fi48++ {
                    v49 := GenS6{F0: (true && (!false)), F1: 0}
                    var v50 int16 = 941
                    var v51 string = ""
                    _ = fi48
                    _ = v49
                    _ = v50
                    _ = v51
                }
                _ = v47
            }
        }
    } else {
        if false {
            fmt.Println(v3)
            switch (-(v2)) {
            case -2147483648:
                for fi52 := 0; fi52 < 2; fi52++ {
                    switch v2 {
                    case -728:
                        v5 = 42
                    default:
                        fmt.Println(v1)
                    }
                    _ = fi52
                }
            case 103:
                v3 -= (v3 % ((v3) | 1))
            }
            var v53 int32 = 0
            _ = v53
        } else {
            switch (v2 - 4294967295) {
            case -340:
                switch -2147483648 {
                case 90:
                    v54 := GenS1{F0: true}
                    _ = v54
                case -1:
                    v14 = v14
                }
            case 816:
                var v55 string = ""
                _ = v55
            case -982:
                fmt.Println(v2)
            }
            if false {
                var v56 float32 = 1.5
                _ = v56
            } else {
                v11 = append(v11, v2)
                v57 := GenS7{F0: v4, F1: 562}
                _ = v57
            }
        }
        var v58 int16 = -32768
        fmt.Println(v2, v5)
        _ = v58
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
}
