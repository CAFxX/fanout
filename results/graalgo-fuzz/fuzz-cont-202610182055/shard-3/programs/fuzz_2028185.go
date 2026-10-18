package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 bool; F1 uint8 }

func main() {
    var v1 float64 = (1.5 - 1.5)
    v2 := []int{0, -34}
    var v3 int8 = -128
    var v4 int8 = v3
    v5 := []string{"\n\t", "a", "hello", ("" + ("foo bar" + "\n\t"))}
    switch 138 {
    case -9223372036854775808:
        var v6 uint = 42
        _ = v6
    case -1:
        if ((!false) || true) {
            if false {
                if false {
                    v1 = (1.5 * (v1 + (3.14159 * 0.1)))
                    v3 = v4
                }
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                switch -98 {
                case 2147483647:
                    v4 = v3
                }
            } else {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            _ = len(v5)
        } else {
            v7 := GenS1{F0: 255, F1: 1000}
            for fi8 := 0; fi8 < 4; fi8++ {
                v4 = v4
                _ = fi8
            }
            var v9 uint8 = 2
            _ = v7
            _ = v9
        }
    case 627:
        for fi10 := 0; fi10 < 15; fi10++ {
            _ = len(v5)
            v3 += (v3 + v3)
            _ = fi10
        }
    }
    var v11 string = "a"
    v12 := GenS1{F0: 2, F1: 4294967295}
    var v13 int8 = 100
    switch ((v11 + ((v11 + v11) + "\n\t")) + (v11 + "a")) {
    case "a":
        var v14 uint8 = 2
        _ = v14
    case "\n\t":
        v13 = 0
    case "foo bar":
        v13 -= (v3 | v3)
    }
    for fi15 := 0; fi15 < 11; fi15++ {
        if (!("\n\t" == "hello")) {
            var v16 int32 = -2147483648
            _ = v16
        } else {
            v11 = "\n\t"
            v17 := GenS1{F0: 0, F1: 0}
            _ = v17
        }
        _ = fi15
    }
    switch 32767 {
    case 213:
        var v18 int32 = 1
        _ = v18
    case -871:
        switch v11 {
        case "hello":
            _ = len(v5)
        case "":
            for fi19 := 0; fi19 < 6; fi19++ {
                for fi20 := 0; fi20 < 4; fi20++ {
                    v3 -= (v13 & (v13 * -126))
                    _ = fi20
                }
                _ = fi19
            }
        case "\n\t":
            for fi21 := 0; fi21 < 15; fi21++ {
                v13 += (-(v4))
                var v22 int8 = v3
                _ = fi21
                _ = v22
            }
        default:
            for fi23 := 0; fi23 < 7; fi23++ {
                var v24 bool = (4294967295 != 67)
                for fi25 := 0; fi25 < 8; fi25++ {
                    v3 = (v3 | 127)
                    _ = fi25
                }
                _ = fi23
                _ = v24
            }
        }
    case 2147483647:
        for fi26 := 0; fi26 < 3; fi26++ {
            v27 := GenS1{F0: 1, F1: 42}
            if ((0.0 >= 0.0) || (-32768 > -509)) {
                if len(v2) > 0 {
                    v2[0] = (-(fi26))
                }
            }
            if (("foo bar" + v11) <= v11) {
                if len(v5) > 0 {
                    v5[0] = v11
                }
                for fi28 := 0; fi28 < 7; fi28++ {
                    v13 -= (v3 + v3)
                    v29 := GenS2{F0: true}
                    v3 = (v13 + v4)
                    _ = fi28
                    _ = v29
                }
                var v30 uint = 65535
                _ = v30
            }
            _ = fi26
            _ = v27
        }
    }
    v31 := GenS1{F0: 2, F1: 704}
    if (1 != 2) {
        var v32 float32 = -2.25
        _ = v32
    }
    switch 2147483647 {
    case 42:
        v4 += v3
    case -2147483648:
        var v33 int = 128
        _ = v33
    default:
        var v34 string = (v11 + v11)
        _ = v34
    }
    var v35 uint8 = 42
    if (-262 == 1) {
        v11 = (v11 + v11)
        var v36 bool = false
        _ = v36
    } else {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        for fi37 := 0; fi37 < 15; fi37++ {
            var v38 uint = 0
            if true {
                var v39 uint = 609
                if false {
                    var v40 bool = (!(!(1.5 < 3.14159)))
                    for fi41 := 0; fi41 < 2; fi41++ {
                        v38 -= v39
                        _ = fi41
                    }
                    _ = v40
                } else {
                    if len(v2) > 0 {
                        v2[0] = 9223372036854775807
                    }
                    var v42 uint64 = 700
                    for fi43 := 0; fi43 < 14; fi43++ {
                        if (1 >= 2147483647) {
                            var v44 bool = (!(1 != 9223372036854775807))
                            _ = v44
                        } else {
                            if false {
                                if len(v5) > 0 {
                                    v5[0] = v11
                                }
                            } else {
                                var v45 int = (-(fi37))
                                _ = v45
                            }
                            if true {
                                v46 := GenS3{F0: -457}
                                switch v11 {
                                case "a":
                                    var v47 uint32 = uint32(v35)
                                    _ = v47
                                }
                                _ = v46
                            } else {
                                if ((v1 + (((v1 / v1) * float64((1.5 + 1e10))) - v1)) == v1) {
                                    var v48 string = (v11 + v11)
                                    v11 = ""
                                    v39 += (v39 ^ 4294967295)
                                    _ = v48
                                } else {
                                    var v49 bool = true
                                    _ = v49
                                }
                                v13 = v4
                            }
                            v50 := GenS4{F0: ((((fi37 == fi43) && (684 <= 354)) && (true || (0.0 >= float32(v1)))) || false), F1: 95}
                            _ = v50
                        }
                        fmt.Println(v11)
                        v38 -= 65535
                        _ = fi43
                    }
                    _ = v42
                }
                _ = v39
            }
            if false {
                if len(v5) > 0 {
                    v5[0] = (v11 + v11)
                }
            } else {
                if len(v2) > 0 {
                    v2[0] = fi37
                }
                v1 -= v1
                switch v11 {
                case "hello":
                    v51 := GenS4{F0: true, F1: v35}
                    _ = v51
                case "":
                    var v52 uint8 = (v35 % ((v35) | 1))
                    _ = v52
                }
            }
            _ = fi37
            _ = v38
        }
    }
    var v53 int8 = (v3 << ((687) & 63))
    if (65535 != 32767) {
        _ = len(v5)
        fmt.Println(v1, v35, v4)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v31
    _ = v35
    _ = v4
    _ = v5
    _ = v53
}
