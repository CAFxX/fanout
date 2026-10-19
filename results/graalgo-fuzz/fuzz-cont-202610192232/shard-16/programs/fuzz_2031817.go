package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 bool }
type GenS2 struct { F0 int; F1 int16; F2 int8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int; F1 uint }
type GenS5 struct { F0 string }
func genFn0(p0 uint16, p1 bool, p2 uint16) int {
    switch -9223372036854775808 {
    case -1:
        var v7 uint64 = 1
        _ = v7
    }
    var v8 int64 = -1
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 0
}
func genFn1(p0 int) string {
    v9 := []int{(p0 >> ((65535) & 63))}
    _ = len(v9)
    _ = p0
    _ = v9
    return "\n\t"
}

func main() {
    var v1 bool = (134 > 128)
    var v2 bool = v1
    v3 := make(map[uint32]uint32)
    v4 := make(map[int]string)
    var v5 uint64 = 952
    v6 := []string{(("foo bar" + "a") + (("hello" + "\n\t") + "\n\t"))}
    var v10 int16 = 32767
    for fi11 := 0; fi11 < 11; fi11++ {
        v12 := GenS1{F0: fi11, F1: 1, F2: true}
        _ = fi11
        _ = v12
    }
    if (2 >= 4294967295) {
        var v13 uint64 = v5
        v5 -= 581
        v2 = v2
        _ = v13
    } else {
        switch 128 {
        case -861:
            if (1e10 == 1e10) {
                var v14 uint = 59
                if true {
                    switch ("" + "") {
                    case "\n\t":
                        if (!(true && (!true))) {
                            v1 = (v2 && v1)
                            v2 = true
                        }
                    }
                    for fi15 := 0; fi15 < 13; fi15++ {
                        var v16 float64 = -2.25
                        _ = fi15
                        _ = v16
                    }
                }
                _ = v14
            } else {
                for fi17 := 0; fi17 < 2; fi17++ {
                    v18 := v3[2]
                    var v19 uint64 = uint64(v18)
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                v10 = v10
            }
        case 4294967295:
            var v20 uint = 175
            _ = v20
        case 756:
            switch "hello" {
            case "a":
                if ((304 >= 0) && (997 < 0)) {
                    var v21 bool = ((v1 && (genFn0(32767, (18446744073709551615 < 4294967295), 0) <= -811)) || (-100 <= -1))
                    var v22 int = -9223372036854775808
                    switch "foo bar" {
                    case "":
                        switch int(v5) {
                        case 42:
                            var v23 uint8 = 116
                            _ = v23
                        case 32768:
                            var v24 int = (v22 << ((697) & 63))
                            _ = v24
                        default:
                            v3[2] = 4294967295
                        }
                    case "a":
                        v2 = (((((42 >= -9223372036854775808) || v21) && false) && true) && (681 > -9223372036854775808))
                    default:
                        v25 := GenS2{F0: (^(v22)), F1: v10, F2: -44}
                        _ = v25
                    }
                    _ = v21
                    _ = v22
                } else {
                    for fi26 := 0; fi26 < 18; fi26++ {
                        v27 := GenS1{F0: genFn0(1, ((v10 & 32767) < v10), 0), F1: 2, F2: (v2 || (!(v10 > v10)))}
                        _ = fi26
                        _ = v27
                    }
                    fmt.Println(v1)
                }
            }
        }
        for fi28 := 0; fi28 < 13; fi28++ {
            var v29 uint16 = 42
            v30 := GenS3{F0: 1}
            v10 = (^(v10))
            _ = fi28
            _ = v29
            _ = v30
        }
    }
    fmt.Println(v10, v1)
    for fi31 := 0; fi31 < 5; fi31++ {
        var v32 string = ("a" + ("foo bar" + "hello"))
        fmt.Println(v2, fi31)
        v33 := v4[fi31]
        _ = fi31
        _ = v32
        _ = v33
    }
    for fi34 := 0; fi34 < 7; fi34++ {
        var v35 uint16 = 1
        _ = fi34
        _ = v35
    }
    if v2 {
        v5 -= (v5 >> ((0) & 63))
        if (v2 && true) {
            fmt.Println(v10, v2, v5)
            v10 += (v10 % ((0) | 1))
            fmt.Println(v5, v5)
        }
        if false {
            for fi36 := 0; fi36 < 12; fi36++ {
                for fi37 := 0; fi37 < 17; fi37++ {
                    if true {
                        v10 += v10
                    }
                    switch fi37 {
                    case 32767:
                        v38 := GenS1{F0: fi37, F1: 1, F2: (v1 || v1)}
                        v39 := GenS2{F0: 2147483647, F1: v10, F2: 127}
                        if len(v6) > 0 {
                            v6[0] = ("\n\t" + "\n\t")
                        }
                        _ = v38
                        _ = v39
                    case 2147483647:
                        v40 := GenS4{F0: (fi36 / (((fi37 << ((237) & 63))) | 1)), F1: 2}
                        _ = v40
                    case 44:
                        if (733 < 255) {
                            if v2 {
                                switch "\n\t" {
                                case "hello":
                                    v41 := GenS5{F0: (("\n\t" + "foo bar") + "")}
                                    _ = v41
                                case "":
                                    v5 -= (v5 + v5)
                                default:
                                    v42 := GenS5{F0: (("a" + ("hello" + "hello")) + "a")}
                                    _ = v42
                                }
                            }
                        } else {
                            v4[fi37] = ("hello" + "a")
                        }
                    }
                    var v43 int8 = 1
                    switch "a" {
                    case "a":
                        v43 += v43
                        _ = len(v6)
                    case "\n\t":
                        v43 = -78
                    case "hello":
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                    }
                    _ = fi37
                    _ = v43
                }
                var v44 uint8 = 223
                _ = fi36
                _ = v44
            }
        } else {
            for fi45 := 0; fi45 < 19; fi45++ {
                var v46 uint16 = 2
                _ = fi45
                _ = v46
            }
            v10 -= (v10 / (((v10 * v10)) | 1))
            for fi47 := 0; fi47 < 18; fi47++ {
                fmt.Println(v1)
                _ = fi47
            }
        }
    } else {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        v48 := GenS3{F0: 175}
        v6 = append(v6, "hello")
        _ = v48
    }
    fmt.Println(v2, v1)
    var v49 uint64 = v5
    delete(v3, 416)
    switch genFn0(806, (false && v1), 32767) {
    case 32768:
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    case 4294967295:
        for fi50 := 0; fi50 < 8; fi50++ {
            for fi51 := 0; fi51 < 15; fi51++ {
                var v52 int16 = (v10 >> ((1) & 63))
                _ = fi51
                _ = v52
            }
            _ = fi50
        }
    case -691:
        if (!(v10 >= (v10 | (v10 & v10)))) {
            v2 = v2
            v10 = 0
            if len(v6) > 0 {
                v6[0] = "a"
            }
        }
    }
    v53 := GenS3{F0: 2}
    var v54 uint8 = 65
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v53
    _ = v54
    _ = v6
}
