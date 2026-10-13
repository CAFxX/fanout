package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 string; F1 uint; F2 int8 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint32; F1 uint8; F2 uint16 }
type GenS5 struct { F0 uint64; F1 int32 }
type GenS6 struct { F0 string; F1 uint64; F2 uint }
type GenS7 struct { F0 uint64; F1 uint32 }
func genFn0(p0 bool, p1 uint32) int {
    if (!true) {
        v7 := []int{0, 128, 127}
        v8 := GenS1{F0: 2147483647, F1: 79}
        _ = v7
        _ = v8
    }
    if (-1 != 42) {
        var v9 float64 = ((float64(3.14159) / 3.14159) / -2.25)
        var v10 uint8 = 128
        v11 := GenS2{F0: "\n\t", F1: 319, F2: -1}
        _ = v10
        _ = v11
        _ = v9
    }
    var v12 int8 = 89
    _ = p0
    _ = p1
    _ = v12
    return -181
}
func genFn1() uint16 {
    var v13 uint8 = 42
    if false {
        v14 := GenS1{F0: -630, F1: 18446744073709551615}
        var v15 uint64 = 52
        _ = v14
        _ = v15
    } else {
        var v16 uint16 = 670
        var v17 bool = (0 < 1000)
        _ = v16
        _ = v17
    }
    v18 := []int{}
    if false {
        v19 := GenS3{F0: -1}
        v13 -= 127
        v13 -= (v13 << ((65535) & 63))
        _ = v19
    } else {
        var v20 string = "foo bar"
        _ = v20
    }
    fmt.Println(v13)
    _ = v13
    _ = v18
    return 2
}

func main() {
    var v1 float64 = 0.0
    var v2 int = -540
    v3 := make(map[string]int)
    v4 := []string{("foo bar" + ""), "\n\t", "a", "\n\t"}
    var v5 uint16 = 65535
    v6 := make(map[string]uint32)
    fmt.Println(v2, v2)
    v21 := GenS4{F0: 1, F1: 2, F2: (v5 % (((v5 - (v5 - (v5 + 254)))) | 1))}
    for fi22 := 0; fi22 < 17; fi22++ {
        switch "foo bar" {
        case "":
            v23 := GenS3{F0: -1}
            _ = v23
        default:
            if (!(false || (!(!(-219 <= -1))))) {
                v24 := GenS4{F0: 2147483647, F1: 255, F2: v5}
                v2 = -502
                _ = v24
            } else {
                v5 -= (v5 % (((v5 ^ v5)) | 1))
                for fi25 := 0; fi25 < 19; fi25++ {
                    var v26 uint16 = 65535
                    _ = fi25
                    _ = v26
                }
                for fi27 := 0; fi27 < 6; fi27++ {
                    var v28 bool = (1 == 1)
                    v2 += (v2 - 32768)
                    _ = fi27
                    _ = v28
                }
            }
        }
        switch "hello" {
        case "a":
            _ = len(v4)
        case "hello":
            if false {
                v2 = fi22
                for fi29 := 0; fi29 < 13; fi29++ {
                    var v30 float32 = 0.1
                    v31 := GenS1{F0: -9223372036854775808, F1: 1}
                    _ = fi29
                    _ = v30
                    _ = v31
                }
            }
        case "foo bar":
            v5 -= genFn1()
        default:
            v32 := GenS4{F0: 586, F1: 42, F2: (v5 ^ v5)}
            _ = v32
        }
        _ = fi22
    }
    v33 := GenS4{F0: uint32(v5), F1: 2, F2: genFn1()}
    switch "" {
    case "\n\t":
        v6[""] = 554
    case "a":
        switch "\n\t" {
        case "":
            if true {
                for fi34 := 0; fi34 < 5; fi34++ {
                    v4 = append(v4, "foo bar")
                    _ = fi34
                }
                var v35 int64 = 42
                _ = v35
            } else {
                v2 = -406
                for fi36 := 0; fi36 < 1; fi36++ {
                    var v37 uint64 = 1
                    _ = fi36
                    _ = v37
                }
                if (v5 <= v5) {
                    _ = len(v4)
                    v38 := GenS3{F0: 9223372036854775807}
                    _ = v38
                } else {
                    v2 += genFn0((!(0 == 1)), 2147483648)
                }
            }
        case "hello":
            v5 -= v5
        case "\n\t":
            v39 := GenS5{F0: 18446744073709551615, F1: 1}
            _ = v39
        default:
            var v40 bool = (255 == 42)
            _ = v40
        }
    default:
        switch v2 {
        case 0:
            v1 = v1
        case -85:
            v5 -= 2
        default:
            var v41 bool = false
            _ = v41
        }
    }
    for fi42 := 0; fi42 < 7; fi42++ {
        v1 = float64((0.0 + -2.25))
        _ = len(v6)
        _ = fi42
    }
    fmt.Println(v2, v1)
    v2 -= genFn0(true, 904)
    v1 = 1e10
    if (50 >= 127) {
        switch 671 {
        case 2147483647:
            v43 := GenS6{F0: "", F1: 4294967295, F2: 4294967295}
            _ = v43
        case 32768:
            v3["hello"] = (v2 | (v2 << ((65535) & 63)))
        case 916:
            v44 := GenS5{F0: 2, F1: 125}
            _ = v44
        }
    }
    for fi45 := 0; fi45 < 16; fi45++ {
        var v46 uint64 = 2
        _ = fi45
        _ = v46
    }
    for fi47 := 0; fi47 < 12; fi47++ {
        v48 := GenS6{F0: "hello", F1: 42, F2: 687}
        _ = fi47
        _ = v48
    }
    for fi49 := 0; fi49 < 1; fi49++ {
        switch (("foo bar" + "hello") + ("a" + "a")) {
        case "hello":
            v50 := GenS7{F0: 42, F1: 65535}
            _ = v50
        case "\n\t":
            var v51 int8 = 60
            _ = v51
        default:
            var v52 float64 = float64(-2.25)
            _ = v52
        }
        _ = fi49
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
}
