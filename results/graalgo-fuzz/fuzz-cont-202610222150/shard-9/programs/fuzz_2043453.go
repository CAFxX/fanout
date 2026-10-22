package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 bool; F1 int8; F2 uint }

func main() {
    var v1 bool = true
    v2 := []int{-9223372036854775808, -564}
    var v3 uint64 = 4294967295
    var v4 uint = 1
    switch "hello" {
    case "":
        switch 42 {
        case 128:
            v1 = true
        case -552:
            var v5 string = "foo bar"
            _ = v5
        default:
            v3 = (v3 | 4294967295)
        }
    case "\n\t":
        if false {
            for fi6 := 0; fi6 < 18; fi6++ {
                v7 := GenS1{F0: -229}
                if (0 != 9223372036854775807) {
                    fmt.Println(fi6)
                    v3 -= v3
                    var v8 int8 = 127
                    _ = v8
                }
                v9 := GenS1{F0: 9223372036854775807}
                _ = fi6
                _ = v7
                _ = v9
            }
            var v10 uint16 = 786
            fmt.Println(v4, v4, v4)
            _ = v10
        } else {
            for fi11 := 0; fi11 < 8; fi11++ {
                var v12 int16 = 554
                v1 = (v1 && v1)
                _ = fi11
                _ = v12
            }
            fmt.Println(v3, v4, v3)
            var v13 int = -502
            _ = v13
        }
    case "foo bar":
        fmt.Println(v4, v1, v1)
    }
    v3 = 1
    v4 = v4
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi14 := 0; fi14 < 10; fi14++ {
        switch ("a" + ("foo bar" + "foo bar")) {
        case "foo bar":
            v15 := GenS1{F0: 0}
            _ = v15
        default:
            var v16 uint64 = v3
            _ = v16
        }
        _ = fi14
    }
    v17 := GenS1{F0: 0}
    v18 := GenS2{F0: v1}
    fmt.Println(v3, v4)
    if false {
        if true {
            v1 = true
            if true {
                var v19 int64 = 0
                _ = v19
            }
        }
        switch 0 {
        case -9223372036854775808:
            if true {
                v4 -= 1
            } else {
                for fi20 := 0; fi20 < 6; fi20++ {
                    v21 := GenS2{F0: (102 < -122)}
                    if false {
                        _ = len(v2)
                        var v22 uint64 = v3
                        _ = v22
                    } else {
                        _ = len(v2)
                        for fi23 := 0; fi23 < 4; fi23++ {
                            v4 = uint(v3)
                            _ = fi23
                        }
                        v3 = uint64(v4)
                    }
                    if len(v2) > 0 {
                        v2[0] = 1
                    }
                    _ = fi20
                    _ = v21
                }
                var v24 uint8 = 128
                var v25 uint8 = (v24 / ((v24) | 1))
                _ = v24
                _ = v25
            }
        }
    } else {
        switch 128 {
        case 2147483647:
            v3 = (v3 * v3)
        case -2147483648:
            fmt.Println(v3)
        default:
            switch -304 {
            case -1:
                if v1 {
                    var v26 int = 42
                    v1 = v1
                    _ = v26
                }
            case 118:
                fmt.Println(v1, v4, v4)
            case 32768:
                v27 := GenS1{F0: 632}
                _ = v27
            default:
                if ((0 >= -884) && false) {
                    var v28 uint8 = 127
                    if false {
                        fmt.Println(v3, v28)
                        v29 := GenS3{F0: v3}
                        _ = v29
                    } else {
                        if (1 >= 100) {
                            switch 127 {
                            case -934:
                                var v30 int8 = 100
                                _ = v30
                            case 1:
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                            case -209:
                                switch -988 {
                                case 981:
                                    fmt.Println(v1)
                                }
                            }
                            var v31 int16 = -302
                            _ = v31
                        }
                        v2 = append(v2, -2147483648)
                    }
                    for fi32 := 0; fi32 < 3; fi32++ {
                        var v33 uint32 = 2147483648
                        _ = fi32
                        _ = v33
                    }
                    _ = v28
                } else {
                    var v34 float64 = -2.25
                    _ = v34
                }
            }
        }
        fmt.Println(v1, v1)
    }
    var v35 uint = (v4 >> ((1) & 63))
    for fi36 := 0; fi36 < 16; fi36++ {
        for fi37 := 0; fi37 < 1; fi37++ {
            _ = len(v2)
            v38 := GenS1{F0: 1}
            _ = fi37
            _ = v38
        }
        v39 := GenS3{F0: v3}
        if (322 == v35) {
            fmt.Println(v4, v1)
            if (0 < -1) {
                for fi40 := 0; fi40 < 12; fi40++ {
                    v1 = true
                    var v41 int = fi40
                    if ((!(759 > uint16(v4))) || ((v3 >= v3) && (!(1000 == -32768)))) {
                        v42 := GenS1{F0: 687}
                        _ = v42
                    }
                    v43 := GenS1{F0: int64(fi40)}
                    var v44 uint = 18446744073709551615
                    _ = fi40
                    _ = v41
                    _ = v43
                    _ = v44
                }
                switch "foo bar" {
                case "foo bar":
                    _ = len(v2)
                case "a":
                    v3 = (v3 >> ((v4) & 63))
                }
                v4 = v4
            } else {
                switch "\n\t" {
                case "\n\t":
                    v35 = v4
                default:
                    v35 -= v35
                }
                v45 := GenS4{F0: true, F1: 65, F2: v4}
                _ = v45
            }
        }
        _ = fi36
        _ = v39
    }
    v46 := GenS4{F0: false, F1: 42, F2: (v35 % ((4294967295) | 1))}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v46
}
