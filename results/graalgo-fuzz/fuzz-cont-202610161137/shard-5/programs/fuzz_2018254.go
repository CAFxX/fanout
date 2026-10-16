package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 int }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int; F1 uint8 }
type GenS5 struct { F0 int16; F1 int32; F2 bool }
type GenS6 struct { F0 uint16; F1 bool; F2 bool }
type GenS7 struct { F0 uint64; F1 bool; F2 uint16 }
type GenS8 struct { F0 string; F1 int16 }
func genFn0(p0 int64, p1 uint8) uint16 {
    p1 = 255
    p1 = p1
    if (!false) {
        if (-634 == 1) {
            for fi5 := 0; fi5 < 10; fi5++ {
                var v6 uint16 = uint16(p1)
                _ = fi5
                _ = v6
            }
            var v7 uint = 255
            _ = v7
        }
        p1 -= (p1 << ((255) & 63))
        v8 := []int{-634}
        _ = v8
    }
    v9 := []int{-2147483648, 2147483647}
    p1 -= p1
    _ = p0
    _ = p1
    _ = v9
    return 65535
}

func main() {
    var v1 int32 = -420
    v2 := []int{}
    v3 := []string{"a", ""}
    var v4 float64 = 1.5
    v10 := GenS1{F0: 0, F1: 874, F2: 764}
    for fi11 := 0; fi11 < 15; fi11++ {
        if len(v2) > 0 {
            v2[0] = 0
        }
        switch 73 {
        case 9223372036854775807:
            v4 += 3.14159
        case -708:
            v4 = v4
        case 32768:
            if true {
                v1 = v1
                var v12 bool = (!true)
                switch "hello" {
                case "\n\t":
                    for fi13 := 0; fi13 < 5; fi13++ {
                        v12 = v12
                        v14 := GenS1{F0: 1, F1: (fi13 / ((127) | 1)), F2: fi13}
                        v15 := GenS1{F0: 278, F1: fi11, F2: 42}
                        if true {
                            if (true && ((v4 / v4) == (v4 - float64(((1.5 * 0.1) / 1.0))))) {
                                v4 = (1e10 + ((v4 + 1e10) - v4))
                            }
                            fmt.Println(v12, v1)
                        } else {
                            v16 := GenS2{F0: fi13}
                            if true {
                                fmt.Println(v1, v1, fi13)
                            } else {
                                v12 = v12
                            }
                            _ = v16
                        }
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                default:
                    if v12 {
                        for fi17 := 0; fi17 < 11; fi17++ {
                            v18 := GenS1{F0: 18446744073709551615, F1: fi11, F2: 127}
                            switch (fi11 << ((42) & 63)) {
                            case 0:
                                switch fi17 {
                                case -239:
                                    _ = len(v2)
                                default:
                                    if len(v3) > 0 {
                                        v3[0] = ("\n\t" + "")
                                    }
                                }
                            default:
                                var v19 bool = (genFn0(-9223372036854775808, 51) > 160)
                                _ = v19
                            }
                            v2 = append(v2, (fi17 ^ fi11))
                            _ = fi17
                            _ = v18
                        }
                        var v20 int16 = 760
                        v21 := GenS3{F0: 522}
                        _ = v20
                        _ = v21
                    }
                }
                _ = v12
            }
        default:
            switch 127 {
            case 2147483647:
                fmt.Println(v4)
            case 42:
                var v22 uint64 = 18446744073709551615
                _ = v22
            case 4294967295:
                var v23 uint64 = 50
                _ = v23
            default:
                v24 := GenS3{F0: 9223372036854775807}
                _ = v24
            }
        }
        if len(v3) > 0 {
            v3[0] = "foo bar"
        }
        _ = fi11
    }
    if false {
        if false {
            v25 := GenS1{F0: 798, F1: -9223372036854775808, F2: -1}
            v1 = -590
            _ = v25
        }
    } else {
        for fi26 := 0; fi26 < 2; fi26++ {
            var v27 uint16 = 227
            _ = fi26
            _ = v27
        }
    }
    for fi28 := 0; fi28 < 1; fi28++ {
        var v29 int = -1
        switch v29 {
        case 42:
            v4 = 1.5
        case -2147483648:
            var v30 int32 = v1
            _ = v30
        }
        if len(v3) > 0 {
            v3[0] = "hello"
        }
        _ = fi28
        _ = v29
    }
    for fi31 := 0; fi31 < 5; fi31++ {
        var v32 uint32 = 663
        switch "foo bar" {
        case "":
            v33 := GenS4{F0: (fi31 | fi31), F1: 128}
            _ = v33
        case "hello":
            if ("foo bar" == "\n\t") {
                for fi34 := 0; fi34 < 20; fi34++ {
                    switch (^(fi31)) {
                    case 13:
                        if true {
                            v35 := GenS2{F0: 128}
                            _ = v35
                        }
                    default:
                        var v36 bool = false
                        _ = v36
                    }
                    var v37 int32 = (v1 % ((v1) | 1))
                    _ = len(v2)
                    _ = fi34
                    _ = v37
                }
                if false {
                    var v38 uint32 = 729
                    v4 -= (1.5 - 1e10)
                    _ = v38
                }
                v1 = (v1 + (v1 >> ((18446744073709551615) & 63)))
            } else {
                fmt.Println(v1, v4, v4)
            }
        }
        v3 = append(v3, "foo bar")
        _ = fi31
        _ = v32
    }
    v4 += 3.14159
    v1 = (v1 << ((4294967295) & 63))
    for fi39 := 0; fi39 < 16; fi39++ {
        v40 := GenS5{F0: 430, F1: 0, F2: true}
        _ = fi39
        _ = v40
    }
    fmt.Println(v1, v4)
    if (true || false) {
        var v41 bool = true
        if true {
            for fi42 := 0; fi42 < 3; fi42++ {
                v43 := GenS6{F0: 42, F1: (!false), F2: (v41 || (65535 > 42))}
                _ = len(v2)
                var v44 int8 = -35
                _ = fi42
                _ = v43
                _ = v44
            }
        }
        _ = v41
    } else {
        v45 := GenS7{F0: 472, F1: false, F2: 32767}
        v46 := GenS6{F0: 32767, F1: false, F2: false}
        if (!true) {
            v47 := GenS7{F0: 141, F1: (v1 != -2147483648), F2: 1}
            _ = v47
        } else {
            v4 -= v4
            switch ("a" + "foo bar") {
            case "":
                var v48 int8 = 42
                _ = v48
            case "\n\t":
                v4 += v4
            default:
                for fi49 := 0; fi49 < 10; fi49++ {
                    v50 := GenS1{F0: 413, F1: 693, F2: (fi49 ^ (^(fi49)))}
                    _ = fi49
                    _ = v50
                }
            }
            switch 1 {
            case 822:
                v1 = (v1 / (((^(v1))) | 1))
            case 32767:
                v4 = (v4 / v4)
            case 32768:
                var v51 int = 1
                _ = v51
            default:
                if (float64(0.1) == ((v4 * v4) - -2.25)) {
                    v2 = append(v2, 868)
                }
            }
        }
        _ = v45
        _ = v46
    }
    fmt.Println(v4, v1)
    v52 := GenS8{F0: "a", F1: -998}
    for fi53 := 0; fi53 < 3; fi53++ {
        fmt.Println(fi53, v4)
        var v54 uint64 = 1
        _ = fi53
        _ = v54
    }
    fmt.Println(v1, v4)
    v1 = (v1 << ((4294967295) & 63))
    v3 = append(v3, ("hello" + "a"))
    var v55 int8 = -113
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v55:", v55)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v52
    _ = v55
}
