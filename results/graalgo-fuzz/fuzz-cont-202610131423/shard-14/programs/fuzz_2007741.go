package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32; F1 uint16; F2 uint8 }
type GenS3 struct { F0 int8; F1 uint64 }
type GenS4 struct { F0 bool }
func genFn0() int64 {
    v7 := GenS1{F0: (!false)}
    for fi8 := 0; fi8 < 6; fi8++ {
        if false {
            if false {
                var v9 uint32 = 65535
                _ = v9
            }
            v10 := []int{}
            _ = v10
        } else {
            for fi11 := 0; fi11 < 1; fi11++ {
                if true {
                    fmt.Println(fi8)
                    var v12 string = "foo bar"
                    _ = v12
                }
                var v13 int16 = 32767
                v13 = (v13 - -852)
                _ = fi11
                _ = v13
            }
            if false {
                var v14 uint8 = 95
                _ = v14
            }
        }
        _ = fi8
    }
    var v15 uint64 = 353
    v16 := GenS1{F0: (!true)}
    _ = v15
    _ = v16
    _ = v7
    return 0
}
func genFn1(p0 int, p1 int) int8 {
    v17 := []int{p1, p0, -1}
    if (181 <= 1) {
        if false {
            switch ("a" + "") {
            case "":
                var v18 uint8 = 42
                _ = v18
            case "foo bar":
                for fi19 := 0; fi19 < 11; fi19++ {
                    for fi20 := 0; fi20 < 12; fi20++ {
                        fmt.Println(fi19, fi20)
                        _ = fi20
                    }
                    p1 = (fi19 << ((4294967295) & 63))
                    switch 683 {
                    case 9223372036854775807:
                        switch (p0 | (^(p1))) {
                        case 0:
                            p0 += 128
                        case -2147483648:
                            p1 -= (fi19 * p0)
                        case 950:
                            switch "" {
                            case "":
                                fmt.Println(p0)
                            default:
                                if false {
                                    switch "" {
                                    case "hello":
                                        var v21 float64 = -2.25
                                        _ = v21
                                    }
                                    _ = len(v17)
                                }
                            }
                        default:
                            switch "foo bar" {
                            case "a":
                                switch ((("\n\t" + ("foo bar" + "foo bar")) + "\n\t") + "\n\t") {
                                case "":
                                    if (false && false) {
                                        if (!((false && true) || true)) {
                                            var v22 bool = false
                                            fmt.Println(fi19, v22)
                                            _ = v22
                                        } else {
                                            v17 = append(v17, (-(p1)))
                                            var v23 uint16 = 0
                                            fmt.Println(fi19)
                                            _ = v23
                                        }
                                    }
                                case "hello":
                                    fmt.Println(fi19, p0)
                                }
                            case "\n\t":
                                if true {
                                    fmt.Println(p0)
                                    switch (p0 + p1) {
                                    case -846:
                                        var v24 int8 = 100
                                        _ = v24
                                    }
                                }
                            default:
                                p1 = p1
                            }
                        }
                    case 128:
                        var v25 uint16 = 632
                        _ = v25
                    default:
                        v26 := GenS2{F0: 2147483648, F1: 65535, F2: 128}
                        _ = v26
                    }
                    _ = fi19
                }
            default:
                p0 = (p1 << ((475) & 63))
            }
            v17 = append(v17, (p1 & (-(p0))))
            if len(v17) > 0 {
                v17[0] = p0
            }
        }
        v27 := GenS3{F0: -108, F1: 2}
        for fi28 := 0; fi28 < 9; fi28++ {
            switch 9223372036854775807 {
            case 1:
                fmt.Println(p1)
            case -588:
                var v29 bool = (0 > 977)
                _ = v29
            case -9223372036854775808:
                p1 = p0
            default:
                p1 -= (p0 * p0)
            }
            p0 += (fi28 << ((4294967295) & 63))
            _ = fi28
        }
        _ = v27
    } else {
        var v30 bool = (-2.25 <= 3.14159)
        if (false || (!(42 == 1))) {
            for fi31 := 0; fi31 < 15; fi31++ {
                for fi32 := 0; fi32 < 5; fi32++ {
                    v33 := GenS4{F0: false}
                    var v34 uint8 = 128
                    if v30 {
                        switch "hello" {
                        case "a":
                            var v35 int16 = 0
                            _ = v35
                        default:
                            v30 = v30
                        }
                        var v36 uint8 = 68
                        _ = v36
                    }
                    if (!(788 >= 1)) {
                        var v37 uint8 = v34
                        fmt.Println(v34)
                        v30 = v30
                        _ = v37
                    } else {
                        v17 = append(v17, p0)
                        if (!v30) {
                            if len(v17) >= 2 {
                                v17 = v17[1:len(v17)]
                            }
                            p0 = p1
                        }
                    }
                    _ = fi32
                    _ = v33
                    _ = v34
                }
                if len(v17) >= 2 {
                    v17 = v17[1:len(v17)]
                }
                fmt.Println(v30, v30)
                _ = fi31
            }
            if len(v17) > 0 {
                v17[0] = (p1 << ((124) & 63))
            }
        }
        for fi38 := 0; fi38 < 20; fi38++ {
            v30 = true
            for fi39 := 0; fi39 < 6; fi39++ {
                v40 := GenS3{F0: 0, F1: 14}
                _ = fi39
                _ = v40
            }
            _ = fi38
        }
        _ = v30
    }
    if len(v17) > 0 {
        v17[0] = (p0 - (p0 >> ((569) & 63)))
    }
    var v41 uint16 = 260
    _ = p0
    _ = p1
    _ = v17
    _ = v41
    return 33
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 uint64 = 1
    v3 := make(map[uint32]uint32)
    var v4 uint32 = 4294967295
    var v5 int = 674
    v6 := make(map[uint32]int)
    switch (v5 >> ((18446744073709551615) & 63)) {
    case 127:
        var v42 int16 = 42
        _ = v42
    }
    if (!((v4 == (v4 & v4)) && false)) {
        var v43 uint16 = 2
        fmt.Println(v4)
        _ = v43
    } else {
        for fi44 := 0; fi44 < 12; fi44++ {
            v5 = fi44
            var v45 bool = (!false)
            _ = fi44
            _ = v45
        }
        var v46 bool = false
        for fi47 := 0; fi47 < 13; fi47++ {
            v5 = fi47
            v2 = (v1 << ((65535) & 63))
            _ = fi47
        }
        _ = v46
    }
    v2 += 0
    switch (v5 + v5) {
    case -1:
        v48 := v6[(v4 | 95)]
        _ = v48
    }
    var v49 bool = false
    var v50 int = v5
    if true {
        fmt.Println(v50, v5)
        var v51 int16 = 32767
        if (!v49) {
            var v52 int = v5
            _ = v52
        }
        _ = v51
    }
    v4 = (v4 << ((519) & 63))
    v4 = v4
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v6
}
