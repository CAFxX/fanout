package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 string; F1 uint8; F2 bool }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 string; F1 uint32; F2 int64 }
func genFn0() uint {
    switch 42 {
    case -163:
        var v6 bool = false
        _ = v6
    case 42:
        if true {
            v7 := []int{4294967295, -37}
            _ = v7
        }
    case -1:
        if ((false && (13 == -368)) || false) {
            var v8 uint8 = 128
            v9 := []int{-414, 42, -847}
            for fi10 := 0; fi10 < 7; fi10++ {
                var v11 bool = false
                _ = fi10
                _ = v11
            }
            _ = v8
            _ = v9
        } else {
            v12 := GenS1{F0: 78, F1: 42}
            v13 := GenS1{F0: 1, F1: 4294967295}
            _ = v12
            _ = v13
        }
    default:
        var v14 bool = true
        _ = v14
    }
    if (((false && true) || (1 != -128)) && false) {
        var v15 float32 = -2.25
        for fi16 := 0; fi16 < 7; fi16++ {
            var v17 int64 = 571
            switch "foo bar" {
            case "hello":
                v18 := []int{}
                _ = v18
            case "":
                var v19 int64 = v17
                _ = v19
            case "\n\t":
                v17 = (v17 * 42)
            default:
                v20 := GenS2{F0: "hello", F1: 1, F2: (false && true)}
                _ = v20
            }
            _ = fi16
            _ = v17
        }
        _ = v15
    } else {
        switch 930 {
        case 127:
            if (2 >= 18446744073709551615) {
                var v21 string = ("hello" + "")
                _ = v21
            }
        }
    }
    var v22 int = 9223372036854775807
    if (393 == -2147483648) {
        v22 += v22
    } else {
        v23 := []int{(^(v22)), (-(v22)), v22, v22}
        _ = v23
    }
    v24 := GenS2{F0: ("" + ("hello" + "\n\t")), F1: 0, F2: true}
    _ = v22
    _ = v24
    return 648
}

func main() {
    var v1 float32 = (-2.25 + (0.1 + 0.1))
    var v2 float64 = 0.1
    var v3 uint16 = 255
    var v4 float32 = v1
    v5 := make(map[string]bool)
    fmt.Println(v3, v2)
    var v25 float64 = 3.14159
    var v26 uint32 = 2147483648
    v26 = 644
    switch 48 {
    case -1:
        var v27 int8 = 119
        _ = v27
    case 127:
        if false {
            var v28 int64 = int64(v26)
            _ = v28
        }
    case 2147483647:
        delete(v5, ("foo bar" + ("" + ("hello" + ("a" + "foo bar")))))
    }
    v1 -= v4
    var v29 uint8 = 128
    v30 := GenS1{F0: genFn0(), F1: genFn0()}
    if false {
        var v31 uint32 = (^(v26))
        for fi32 := 0; fi32 < 6; fi32++ {
            switch "\n\t" {
            case "\n\t":
                v25 = 0.0
            case "":
                v1 -= v1
            case "foo bar":
                if false {
                    var v33 int16 = -32768
                    if false {
                        v31 = (^(v31))
                    } else {
                        v34 := GenS3{F0: 2}
                        if (true || false) {
                            var v35 string = (("hello" + ("foo bar" + "foo bar")) + ("foo bar" + (("foo bar" + "hello") + "")))
                            _ = v35
                        }
                        var v36 string = "a"
                        _ = v34
                        _ = v36
                    }
                    v26 = (v31 | 833)
                    _ = v33
                } else {
                    v37 := GenS2{F0: (("foo bar" + ("foo bar" + "\n\t")) + "foo bar"), F1: 127, F2: (!false)}
                    var v38 int16 = 1
                    switch "" {
                    case "a":
                        v39 := GenS4{F0: "foo bar", F1: (v26 >> ((uint(v29)) & 63)), F2: 42}
                        _ = v39
                    }
                    _ = v37
                    _ = v38
                }
            }
            _ = fi32
        }
        fmt.Println(v31, v3, v26)
        _ = v31
    } else {
        v29 -= (v29 | v29)
        v1 = v1
        v40 := []int{0, 32767, 9223372036854775807}
        _ = v40
    }
    for fi41 := 0; fi41 < 18; fi41++ {
        if (false && true) {
            if true {
                switch -9223372036854775808 {
                case -158:
                    v42 := []int{}
                    _ = v42
                case -2147483648:
                    v43 := GenS3{F0: (v3 & (v3 << ((1) & 63)))}
                    _ = v43
                case 2147483647:
                    v26 -= (v26 * (v26 & v26))
                }
            }
        }
        _ = fi41
    }
    switch 2147483647 {
    case 473:
        v29 += 0
    case 128:
        var v44 uint32 = (v26 ^ v26)
        _ = v44
    default:
        if (-1 <= 2147483647) {
            var v45 int = -865
            _ = v45
        } else {
            v46 := []int{32767, 32768, int(v26)}
            v2 -= float64((v1 / 1.5))
            _ = v46
        }
    }
    var v47 int64 = 741
    for fi48 := 0; fi48 < 9; fi48++ {
        v47 = v47
        _ = fi48
    }
    switch 918 {
    case -2147483648:
        v49 := []int{127}
        _ = v49
    case 2147483647:
        v1 -= 0.1
    default:
        var v50 uint64 = 18446744073709551615
        _ = v50
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v47
    _ = v5
}
