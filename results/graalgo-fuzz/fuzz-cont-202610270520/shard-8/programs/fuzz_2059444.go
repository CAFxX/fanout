package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 int16 }
type GenS2 struct { F0 uint16; F1 uint32; F2 uint32 }
type GenS3 struct { F0 bool; F1 uint64; F2 bool }
type GenS4 struct { F0 string }
type GenS5 struct { F0 int16; F1 string; F2 string }
type GenS6 struct { F0 int; F1 string }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 bool }
type GenS9 struct { F0 string; F1 uint; F2 uint }
func genFn0(p0 uint32, p1 uint32) int {
    for fi4 := 0; fi4 < 20; fi4++ {
        for fi5 := 0; fi5 < 8; fi5++ {
            v6 := GenS1{F0: (true && false), F1: 42, F2: 1000}
            p1 += p1
            p1 -= 819
            p1 -= 2147483648
            var v7 int8 = 100
            _ = fi5
            _ = v6
            _ = v7
        }
        v8 := GenS2{F0: uint16(p1), F1: p0, F2: (p1 % ((2147483647) | 1))}
        switch (-(fi4)) {
        case -1:
            if (747 == -1) {
                v9 := []int{32768}
                if false {
                    var v10 float64 = (0.1 - (0.1 - 0.0))
                    _ = v10
                }
                _ = v9
            }
        }
        _ = fi4
        _ = v8
    }
    if (!false) {
        var v11 int16 = -791
        _ = v11
    }
    p1 = (p0 ^ p1)
    p0 += 65535
    _ = p0
    _ = p1
    return 127
}
func genFn1() uint16 {
    var v12 float64 = 0.1
    switch 158 {
    case 850:
        var v13 float64 = v12
        _ = v13
    case 9223372036854775807:
        v12 = (1e10 / (1e10 * v12))
    default:
        for fi14 := 0; fi14 < 16; fi14++ {
            var v15 int16 = -695
            v16 := GenS1{F0: false, F1: 771, F2: (v15 + (-(v15)))}
            _ = fi14
            _ = v15
            _ = v16
        }
    }
    switch ((("a" + "hello") + (("a" + ("" + ("foo bar" + ("foo bar" + "\n\t")))) + "a")) + "") {
    case "a":
        v17 := GenS2{F0: 42, F1: 844, F2: 2147483647}
        _ = v17
    default:
        if (0.0 > (-2.25 / 1.0)) {
            var v18 uint = 42
            v19 := GenS3{F0: false, F1: 937, F2: (32767 == 42)}
            var v20 uint16 = 198
            _ = v18
            _ = v19
            _ = v20
        }
    }
    _ = v12
    return 869
}

func main() {
    var v1 uint16 = 283
    v2 := []uint32{34, 478, 1, 1}
    var v3 uint32 = 691
    var v21 uint16 = (v1 ^ v1)
    v3 = 595
    var v22 uint32 = (v3 & 880)
    for fi23 := 0; fi23 < 15; fi23++ {
        var v24 uint16 = 649
        _ = fi23
        _ = v24
    }
    if len(v2) > 0 {
        v2[0] = v22
    }
    if ((("" + "a") + "a") > "hello") {
        switch "\n\t" {
        case "hello":
            var v25 uint64 = 42
            _ = v25
        case "a":
            fmt.Println(v22)
        default:
            v26 := GenS4{F0: (("\n\t" + "\n\t") + "")}
            _ = v26
        }
        v27 := GenS5{F0: 42, F1: "\n\t", F2: "a"}
        _ = v27
    } else {
        for fi28 := 0; fi28 < 1; fi28++ {
            if (!(42 <= 505)) {
                v29 := GenS6{F0: fi28, F1: "hello"}
                var v30 int8 = -1
                _ = v29
                _ = v30
            }
            _ = fi28
        }
        v31 := GenS7{F0: 18446744073709551615}
        _ = v31
    }
    if (true && false) {
        if true {
            v2 = append(v2, (v22 & v3))
            for fi32 := 0; fi32 < 8; fi32++ {
                if false {
                    switch (fi32 ^ fi32) {
                    case 32768:
                        var v33 int16 = 0
                        _ = v33
                    default:
                        if true {
                            v2 = append(v2, (v3 >> ((1) & 63)))
                        }
                    }
                    if true {
                        var v34 int8 = 42
                        v35 := GenS4{F0: (("hello" + "hello") + (("hello" + ("a" + "a")) + ("a" + "\n\t")))}
                        fmt.Println(fi32, v1)
                        _ = v34
                        _ = v35
                    } else {
                        switch (("" + "hello") + "foo bar") {
                        case "foo bar":
                            v36 := GenS1{F0: false, F1: 0, F2: -1}
                            _ = v36
                        case "a":
                            for fi37 := 0; fi37 < 2; fi37++ {
                                if false {
                                    _ = len(v2)
                                } else {
                                    var v38 int64 = -797
                                    v1 -= (v21 / ((v1) | 1))
                                    _ = v38
                                }
                                _ = len(v2)
                                _ = fi37
                            }
                        }
                        switch ("" + "hello") {
                        case "hello":
                            var v39 float32 = float32(0.1)
                            _ = v39
                        case "foo bar":
                            if len(v2) > 0 {
                                v2[0] = (^(v3))
                            }
                        case "":
                            fmt.Println(v22)
                        }
                    }
                    v3 += uint32(v1)
                }
                for fi40 := 0; fi40 < 2; fi40++ {
                    if (!(true && false)) {
                        fmt.Println(v3, v1)
                    } else {
                        if len(v2) > 0 {
                            v2[0] = v22
                        }
                    }
                    var v41 uint64 = uint64(v1)
                    if (-114 == 0) {
                        var v42 float32 = (1e10 - 0.0)
                        v2 = append(v2, (v22 % (((^(v22))) | 1)))
                        _ = v42
                    } else {
                        v43 := GenS6{F0: int(v22), F1: "a"}
                        if (false && true) {
                            v1 += 255
                        } else {
                            v2 = append(v2, (v3 * 65535))
                            if ((-2.25 / -2.25) > 1.5) {
                                var v44 int64 = 1
                                _ = v44
                            }
                            v45 := GenS8{F0: true}
                            _ = v45
                        }
                        fmt.Println(fi32, fi40, v22)
                        fmt.Println(v41, v41, v21)
                        fmt.Println(fi40, fi32, v21)
                        _ = v43
                    }
                    _ = fi40
                    _ = v41
                }
                _ = fi32
            }
            v22 += 42
        }
        v46 := GenS9{F0: "", F1: 612, F2: 705}
        v47 := GenS4{F0: ""}
        _ = v46
        _ = v47
    }
    if true {
        var v48 int32 = 1
        _ = v48
    }
    v22 = (^(v22))
    v49 := GenS8{F0: (1e10 <= 3.14159)}
    for fi50 := 0; fi50 < 5; fi50++ {
        v3 -= v22
        for fi51 := 0; fi51 < 7; fi51++ {
            var v52 uint32 = (^(v22))
            _ = fi51
            _ = v52
        }
        _ = fi50
    }
    var v53 string = "a"
    fmt.Println(v1)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v54 float64 = 0.0
    v1 -= (v21 ^ (^(v1)))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v49
    _ = v53
    _ = v54
}
