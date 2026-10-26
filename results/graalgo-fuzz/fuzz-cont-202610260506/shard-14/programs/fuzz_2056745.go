package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 uint }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint64; F1 int32; F2 uint64 }
type GenS4 struct { F0 uint16; F1 int }
type GenS5 struct { F0 int64; F1 uint8 }
type GenS6 struct { F0 int32; F1 uint16; F2 int }
type GenS7 struct { F0 int8; F1 uint64; F2 uint16 }
type GenS8 struct { F0 bool }
func genFn0(p0 bool, p1 uint64, p2 uint) int16 {
    if p0 {
        v7 := []int{}
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
        switch -9223372036854775808 {
        case 2147483647:
            if p0 {
                var v8 int32 = 0
                _ = v8
            } else {
                fmt.Println(p1, p2)
                if (65535 == uint16(p1)) {
                    var v9 uint16 = 32767
                    var v10 bool = p0
                    _ = v10
                    _ = v9
                } else {
                    p1 -= 755
                    if len(v7) > 0 {
                        v7[0] = -391
                    }
                    var v11 uint = p2
                    _ = v11
                }
                p0 = (p1 >= p1)
            }
        case 32768:
            for fi12 := 0; fi12 < 20; fi12++ {
                var v13 uint = p2
                v14 := GenS1{F0: p0, F1: 283, F2: 58}
                _ = fi12
                _ = v13
                _ = v14
            }
        case -9223372036854775808:
            for fi15 := 0; fi15 < 13; fi15++ {
                for fi16 := 0; fi16 < 4; fi16++ {
                    v17 := GenS1{F0: (!(fi16 >= (-(fi15)))), F1: -32768, F2: 499}
                    _ = fi16
                    _ = v17
                }
                if len(v7) > 0 {
                    v7[0] = (fi15 >> ((p2) & 63))
                }
                _ = fi15
            }
        default:
            var v18 uint = 42
            _ = v18
        }
        _ = v7
    }
    for fi19 := 0; fi19 < 2; fi19++ {
        v20 := []int{fi19}
        v21 := GenS2{F0: (p1 + p1)}
        p1 += (p1 << ((p2) & 63))
        _ = fi19
        _ = v20
        _ = v21
    }
    v22 := []int{}
    var v23 string = ""
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    _ = v23
    return 280
}

func main() {
    var v1 uint64 = 1
    var v2 float64 = (3.14159 / 1e10)
    var v3 float32 = 0.0
    var v4 string = "foo bar"
    var v5 int32 = -2147483648
    var v6 string = ("a" + "")
    var v24 int = 32767
    v25 := GenS1{F0: false, F1: 1, F2: 0}
    v26 := GenS1{F0: true, F1: 32767, F2: 806}
    if false {
        for fi27 := 0; fi27 < 8; fi27++ {
            v28 := []int{v24, (^(fi27)), v24, (^(v24))}
            var v29 int32 = 2147483647
            _ = fi27
            _ = v28
            _ = v29
        }
        switch v24 {
        case 0:
            var v30 uint16 = 770
            _ = v30
        default:
            var v31 uint32 = 65535
            _ = v31
        }
        if ((-2.25 + v3) == (float32(0.0) + 1e10)) {
            for fi32 := 0; fi32 < 3; fi32++ {
                switch (v6 + "") {
                case "foo bar":
                    var v33 int16 = genFn0(true, v1, 4294967295)
                    _ = v33
                case "hello":
                    for fi34 := 0; fi34 < 7; fi34++ {
                        fmt.Println(v3, v24)
                        if true {
                            v35 := GenS2{F0: v1}
                            if (true || (42 >= 2)) {
                                var v36 string = (((("foo bar" + (v6 + "a")) + v6) + v6) + ((v4 + ("foo bar" + v6)) + v4))
                                var v37 bool = false
                                _ = v36
                                _ = v37
                            }
                            _ = v35
                        } else {
                            if (42 <= 2147483647) {
                                var v38 uint32 = 0
                                fmt.Println(v6, v38)
                                _ = v38
                            }
                            v39 := GenS1{F0: true, F1: 47, F2: 367}
                            if true {
                                v3 = float32(v2)
                                var v40 string = v6
                                _ = v40
                            }
                            _ = v39
                        }
                        if ((255 > 128) && false) {
                            v5 += -2147483648
                            if ((880 > 42) && (!true)) {
                                fmt.Println(fi32, v6)
                            }
                        } else {
                            if (394 == -172) {
                                var v41 uint64 = (v1 << ((4294967295) & 63))
                                _ = v41
                            } else {
                                switch fi34 {
                                case 9223372036854775807:
                                    fmt.Println(v4, v6)
                                }
                                v5 = -2147483648
                            }
                            var v42 uint = 42
                            _ = v42
                        }
                        _ = fi34
                    }
                default:
                    v43 := []int{fi32, fi32}
                    _ = v43
                }
                var v44 uint16 = 2
                _ = fi32
                _ = v44
            }
            if true {
                fmt.Println(v5, v2)
            } else {
                v45 := GenS3{F0: v1, F1: (v5 * -463), F2: 0}
                _ = v45
            }
        }
    } else {
        v46 := []int{v24}
        _ = v46
    }
    v47 := GenS3{F0: v1, F1: (v5 * (-(v5))), F2: v1}
    fmt.Println(v2)
    if ((true && (585 < 4294967295)) && (!true)) {
        if true {
            v48 := []int{(v24 + v24), (v24 << ((60) & 63))}
            _ = v48
        }
        v2 += float64((v3 * float32((3.14159 + 1.5))))
    } else {
        var v49 int = (v24 >> ((65535) & 63))
        _ = v49
    }
    var v50 bool = true
    switch 42 {
    case 32767:
        v4 = v6
    case -819:
        switch "a" {
        case "foo bar":
            var v51 bool = v50
            _ = v51
        case "\n\t":
            v52 := []int{}
            _ = v52
        default:
            for fi53 := 0; fi53 < 11; fi53++ {
                v54 := GenS2{F0: (v1 | (v1 & v1))}
                _ = fi53
                _ = v54
            }
        }
    }
    var v55 string = ((v4 + v6) + v4)
    switch (((v4 + (v55 + ("\n\t" + "hello"))) + "hello") + ("foo bar" + v4)) {
    case "":
        var v56 int8 = 1
        _ = v56
    case "a":
        var v57 float64 = (0.1 * v2)
        _ = v57
    default:
        v58 := GenS4{F0: 65535, F1: (v24 | -9223372036854775808)}
        _ = v58
    }
    for fi59 := 0; fi59 < 15; fi59++ {
        v60 := GenS5{F0: int64(fi59), F1: 127}
        _ = fi59
        _ = v60
    }
    switch v55 {
    case "a":
        var v61 float64 = 0.1
        _ = v61
    case "foo bar":
        v62 := GenS6{F0: v5, F1: 42, F2: 127}
        _ = v62
    case "":
        var v63 bool = v50
        _ = v63
    }
    v64 := []int{v24, v24}
    if (!(108 > 921)) {
        if (-1 > 0) {
            v65 := GenS7{F0: 76, F1: v1, F2: 941}
            v1 -= (v1 >> ((2) & 63))
            v55 = v55
            _ = v65
        }
        for fi66 := 0; fi66 < 2; fi66++ {
            switch -189 {
            case -9223372036854775808:
                fmt.Println(v4, fi66)
            case 26:
                var v67 int32 = (^(v5))
                _ = v67
            case 1:
                if len(v64) >= 2 {
                    v64 = v64[1:len(v64)]
                }
            default:
                fmt.Println(v55, v6)
            }
            if (1.5 > (0.0 - 1e10)) {
                fmt.Println(v2, v6, v6)
            }
            var v68 uint64 = v1
            _ = fi66
            _ = v68
        }
    }
    v1 = (v1 % ((18446744073709551615) | 1))
    v69 := GenS8{F0: (true && (v50 || (!v50)))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v55:", v55)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v64:", len(v64), v64)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v50
    _ = v55
    _ = v6
    _ = v64
    _ = v69
}
