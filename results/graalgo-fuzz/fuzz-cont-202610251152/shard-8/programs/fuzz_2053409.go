package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 int8; F1 uint }
type GenS3 struct { F0 uint8; F1 string; F2 int16 }
type GenS4 struct { F0 uint8; F1 int8; F2 uint8 }
type GenS5 struct { F0 uint64; F1 uint8; F2 int16 }
type GenS6 struct { F0 string; F1 int64 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 string; F1 int32 }
type GenS9 struct { F0 string }
func genFn0(p0 uint8, p1 string, p2 uint16) int32 {
    var v4 uint8 = (p0 >> ((0) & 63))
    switch p1 {
    case "hello":
        var v5 int = 4294967295
        _ = v5
    case "\n\t":
        var v6 int8 = -95
        _ = v6
    case "":
        v7 := GenS1{F0: -677, F1: 18446744073709551615}
        _ = v7
    default:
        p1 = p1
    }
    if true {
        v8 := []int{2147483647, 32768}
        var v9 int32 = 42
        var v10 uint = 309
        _ = v10
        _ = v8
        _ = v9
    }
    var v11 int8 = -128
    v12 := GenS2{F0: (v11 >> ((18446744073709551615) & 63)), F1: 4294967295}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    _ = v4
    return 42
}
func genFn1(p0 int, p1 int) bool {
    for fi13 := 0; fi13 < 9; fi13++ {
        v14 := []int{p0}
        var v15 int8 = -128
        _ = fi13
        _ = v14
        _ = v15
    }
    var v16 int32 = -1
    v17 := GenS3{F0: 2, F1: ("" + "\n\t"), F2: 0}
    switch ("" + ("\n\t" + "hello")) {
    case "foo bar":
        v18 := []int{p0, (p0 | (p1 % ((9223372036854775807) | 1))), p0}
        _ = v18
    default:
        switch (p1 >> ((65535) & 63)) {
        case -2147483648:
            var v19 uint8 = 116
            _ = v19
        case 128:
            for fi20 := 0; fi20 < 15; fi20++ {
                switch ("" + ("\n\t" + "foo bar")) {
                case "":
                    v21 := []int{(p0 ^ fi20), p1}
                    _ = v21
                default:
                    for fi22 := 0; fi22 < 11; fi22++ {
                        p0 = -1
                        var v23 int64 = 1
                        fmt.Println(fi20)
                        v24 := []int{(p0 - (-(fi20))), fi22}
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                }
                var v25 int = fi20
                _ = fi20
                _ = v25
            }
        default:
            if (!("" == "a")) {
                fmt.Println(v16)
                var v26 int64 = -9223372036854775808
                _ = v26
            }
        }
    }
    _ = p0
    _ = p1
    _ = v16
    _ = v17
    return ((v16 - genFn0(128, "hello", 913)) <= (v16 << ((4294967295) & 63)))
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]uint32)
    var v3 string = "\n\t"
    if false {
        var v27 int64 = 1
        _ = v27
    } else {
        v28 := GenS4{F0: 128, F1: 100, F2: 128}
        if true {
            if ((float32(float64((1.5 / 1.5))) <= 0.1) && (255 >= 255)) {
                var v29 uint = 0
                _ = v29
            }
        } else {
            var v30 string = (v3 + "\n\t")
            _ = v30
        }
        _ = v28
    }
    v3 = ((v3 + "a") + v3)
    fmt.Println(v3)
    switch v3 {
    case "a":
        for fi31 := 0; fi31 < 5; fi31++ {
            if (827 <= 0) {
                v32 := []int{127, 127, (fi31 / (((-(fi31))) | 1))}
                for fi33 := 0; fi33 < 7; fi33++ {
                    v34 := GenS5{F0: 18446744073709551615, F1: 42, F2: 0}
                    v3 = v3
                    v32 = append(v32, 1)
                    _ = fi33
                    _ = v34
                }
                _ = v32
            } else {
                v35 := v1[335]
                _ = v35
            }
            for fi36 := 0; fi36 < 9; fi36++ {
                var v37 string = "foo bar"
                _ = fi36
                _ = v37
            }
            _ = fi31
        }
    case "":
        v38 := GenS6{F0: "a", F1: -120}
        _ = v38
    case "hello":
        v39 := GenS7{F0: 1000}
        _ = v39
    default:
        switch ((((((v3 + v3) + v3) + (v3 + v3)) + "\n\t") + v3) + v3) {
        case "\n\t":
            v40 := GenS6{F0: ((v3 + v3) + v3), F1: 9223372036854775807}
            _ = v40
        case "foo bar":
            var v41 uint32 = 0
            _ = v41
        case "a":
            if (0 <= 0) {
                v3 = v3
                v42 := GenS8{F0: v3, F1: 2147483647}
                if true {
                    if (-451 >= 2147483647) {
                        fmt.Println(v3)
                    }
                    var v43 uint = 2
                    _ = v43
                }
                _ = v42
            } else {
                v3 = v3
            }
        }
    }
    switch -83 {
    case 444:
        fmt.Println(v3)
    case 9223372036854775807:
        fmt.Println(v3)
    case -9223372036854775808:
        if true {
            var v44 int8 = 127
            _ = v44
        } else {
            v3 = ((v3 + v3) + v3)
            if false {
                v3 = "foo bar"
                var v45 float32 = 1e10
                _ = v45
            }
            v46 := GenS9{F0: v3}
            _ = v46
        }
    default:
        var v47 float64 = 0.0
        _ = v47
    }
    var v48 uint16 = 65535
    for fi49 := 0; fi49 < 11; fi49++ {
        var v50 uint8 = 128
        if false {
            v3 = v3
            if (!false) {
                var v51 uint32 = 828
                _ = v51
            }
            _ = len(v2)
        } else {
            v50 = v50
            delete(v2, 1)
            var v52 string = "a"
            _ = v52
        }
        _ = fi49
        _ = v50
    }
    for fi53 := 0; fi53 < 17; fi53++ {
        var v54 int16 = 32767
        _ = fi53
        _ = v54
    }
    if true {
        var v55 int8 = -65
        _ = v55
    }
    if ("foo bar" < v3) {
        v3 = ((("\n\t" + v3) + ("a" + v3)) + ("" + ""))
        var v56 uint16 = v48
        _ = v56
    }
    if false {
        if (1 != 42) {
            var v57 uint8 = 198
            v48 -= 612
            for fi58 := 0; fi58 < 13; fi58++ {
                for fi59 := 0; fi59 < 2; fi59++ {
                    v2[487] = 2
                    var v60 int16 = 42
                    v57 += 128
                    _ = fi59
                    _ = v60
                }
                _ = fi58
            }
            _ = v57
        }
    }
    v3 = ""
    var v61 int = -2147483648
    var v62 uint8 = 255
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v48:", v48)
    fmt.Println("v61:", v61)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v2
    _ = v3
    _ = v48
    _ = v61
    _ = v62
}
