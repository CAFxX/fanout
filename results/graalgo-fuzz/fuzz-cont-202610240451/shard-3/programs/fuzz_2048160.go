package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 int8 }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 uint8; F1 uint }
func genFn0(p0 uint8, p1 uint, p2 uint) bool {
    switch "\n\t" {
    case "hello":
        switch 42 {
        case 32767:
            var v4 int8 = 42
            _ = v4
        }
    case "foo bar":
        if true {
            p1 -= 42
        }
    }
    if (-2.25 <= 0.1) {
        v5 := GenS1{F0: 9223372036854775807, F1: 4294967295, F2: 0}
        fmt.Println(p1)
        p1 = 1
        _ = v5
    } else {
        v6 := []int{}
        var v7 int8 = -128
        _ = v6
        _ = v7
    }
    if (18446744073709551615 != 42) {
        switch 127 {
        case 2147483647:
            v8 := []int{55, 972, -2147483648, 32767}
            _ = v8
        case 1:
            v9 := []int{0, 9223372036854775807, 594, 1}
            _ = v9
        }
        for fi10 := 0; fi10 < 3; fi10++ {
            v11 := []int{(fi10 - fi10), (fi10 << ((255) & 63))}
            _ = fi10
            _ = v11
        }
        fmt.Println(p1)
    }
    _ = p0
    _ = p1
    _ = p2
    return true
}
func genFn1(p0 uint32) uint64 {
    p0 -= (p0 << ((uint(p0)) & 63))
    var v12 string = "foo bar"
    v13 := []int{-2147483648, 289}
    var v14 float64 = 0.1
    _ = p0
    _ = v12
    _ = v13
    _ = v14
    return 949
}

func main() {
    var v1 string = ""
    v2 := []string{}
    var v3 int8 = 1
    var v15 uint32 = 47
    for fi16 := 0; fi16 < 10; fi16++ {
        var v17 int16 = 0
        for fi18 := 0; fi18 < 18; fi18++ {
            var v19 uint = uint(v15)
            v20 := GenS1{F0: int64(v19), F1: 42, F2: 0}
            if (1 >= -1) {
                if len(v2) > 0 {
                    v2[0] = v1
                }
                v17 += v17
                v1 = "hello"
            } else {
                switch "" {
                case "":
                    switch (v1 + v1) {
                    case "":
                        v1 = v1
                    default:
                        if true {
                            var v21 int = (fi16 & 128)
                            v22 := GenS1{F0: -9223372036854775808, F1: 731, F2: (v3 | (v3 & 100))}
                            _ = v21
                            _ = v22
                        }
                    }
                case "a":
                    v3 = v3
                case "foo bar":
                    v17 = v17
                default:
                    switch ((v1 + "\n\t") + (v1 + v1)) {
                    case "":
                        v1 = v1
                    case "\n\t":
                        v15 = v15
                    case "foo bar":
                        v23 := GenS2{F0: 255, F1: v1}
                        _ = v23
                    default:
                        fmt.Println(fi16, v3)
                    }
                }
                v1 = "foo bar"
                if (!false) {
                    v2 = append(v2, v1)
                    _ = len(v2)
                } else {
                    switch (("hello" + "foo bar") + ((v1 + ((v1 + v1) + "\n\t")) + "a")) {
                    case "foo bar":
                        v24 := GenS2{F0: 254, F1: v1}
                        _ = v24
                    case "\n\t":
                        if len(v2) > 0 {
                            v2[0] = v1
                        }
                    case "hello":
                        var v25 float32 = 3.14159
                        _ = v25
                    default:
                        v15 = v15
                    }
                    var v26 uint32 = 2147483648
                    _ = v26
                }
                v1 = v1
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            var v27 uint8 = 42
            v28 := GenS2{F0: v19, F1: ("foo bar" + ((v1 + (v1 + "foo bar")) + "a"))}
            _ = fi18
            _ = v19
            _ = v20
            _ = v27
            _ = v28
        }
        fmt.Println(fi16, v1)
        _ = fi16
        _ = v17
    }
    switch ("" + (v1 + v1)) {
    case "hello":
        v15 = (v15 - v15)
    case "foo bar":
        switch 636 {
        case 682:
            var v29 uint64 = genFn1((v15 - 1))
            _ = v29
        case 127:
            fmt.Println(v3)
        case 1:
            v15 += v15
        }
    default:
        v3 = v3
    }
    for fi30 := 0; fi30 < 20; fi30++ {
        for fi31 := 0; fi31 < 4; fi31++ {
            fmt.Println(v3)
            _ = fi31
        }
        _ = fi30
    }
    for fi32 := 0; fi32 < 8; fi32++ {
        var v33 float64 = (3.14159 / 1e10)
        if (("hello" + v1) >= v1) {
            if ((v33 <= v33) || true) {
                v34 := GenS2{F0: 18446744073709551615, F1: v1}
                if true {
                    if (!(255 >= uint8(v15))) {
                        var v35 bool = false
                        v15 -= 2147483648
                        fmt.Println(v33)
                        _ = v35
                    } else {
                        v1 = v1
                        if (979 <= 1) {
                            var v36 uint8 = uint8(v15)
                            v2 = append(v2, "\n\t")
                            for fi37 := 0; fi37 < 20; fi37++ {
                                fmt.Println(fi32, v1)
                                fmt.Println(v33)
                                _ = len(v2)
                                _ = fi37
                            }
                            _ = v36
                        }
                        v2 = append(v2, (v1 + v1))
                    }
                }
                _ = v34
            } else {
                for fi38 := 0; fi38 < 13; fi38++ {
                    if (1000 != 42) {
                        var v39 uint64 = 2
                        v15 = 1
                        if false {
                            var v40 int64 = -240
                            var v41 int16 = 32767
                            _ = v40
                            _ = v41
                        }
                        _ = v39
                    }
                    _ = fi38
                }
            }
            v42 := GenS2{F0: 18446744073709551615, F1: ""}
            _ = v42
        }
        if (1 < -474) {
            if (!false) {
                if false {
                    var v43 bool = true
                    _ = v43
                } else {
                    var v44 int16 = 1
                    _ = v44
                }
                v45 := GenS2{F0: 4294967295, F1: ("foo bar" + v1)}
                _ = v45
            }
            var v46 string = ""
            var v47 uint16 = 1
            _ = v46
            _ = v47
        } else {
            _ = len(v2)
            if len(v2) > 0 {
                v2[0] = "\n\t"
            }
            v1 = "foo bar"
        }
        _ = fi32
        _ = v33
    }
    for fi48 := 0; fi48 < 14; fi48++ {
        for fi49 := 0; fi49 < 4; fi49++ {
            if (466 > 0) {
                v15 = v15
                fmt.Println(v15)
            }
            v3 -= (v3 + v3)
            v50 := GenS1{F0: -9223372036854775808, F1: 322, F2: 0}
            _ = fi49
            _ = v50
        }
        _ = fi48
    }
    v15 = 1
    v15 -= 662
    v51 := GenS3{F0: 210, F1: 255}
    var v52 int32 = -2147483648
    var v53 uint16 = 255
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v51
    _ = v52
    _ = v53
}
