package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 uint8; F1 uint16; F2 uint32 }
type GenS4 struct { F0 int16; F1 uint32 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint16 }
func genFn0(p0 uint8) int16 {
    for fi4 := 0; fi4 < 11; fi4++ {
        v5 := []int{(fi4 - (fi4 / (((^(fi4))) | 1))), fi4, (fi4 ^ fi4), (fi4 % (((-(fi4))) | 1))}
        _ = fi4
        _ = v5
    }
    var v6 uint16 = 765
    if (true && (true && ((("\n\t" + (("hello" + (((("hello" + "\n\t") + (("hello" + (("foo bar" + "a") + "foo bar")) + "")) + "") + "")) + ((("a" + "foo bar") + "") + "hello"))) <= "\n\t") && (false && (!true))))) {
        var v7 float32 = (float32(3.14159) + (0.1 / 0.1))
        _ = v7
    } else {
        v8 := GenS1{F0: v6}
        _ = v8
    }
    _ = p0
    _ = v6
    return -1
}
func genFn1(p0 string, p1 uint32) bool {
    if (255 != 1) {
        v9 := []int{}
        _ = v9
    }
    p0 = "\n\t"
    _ = p0
    _ = p1
    return (!(67 != 2))
}

func main() {
    v1 := make(map[int]bool)
    var v2 float32 = 1.5
    v3 := make(map[int]string)
    switch "a" {
    case "a":
        v10 := v1[-9223372036854775808]
        _ = v10
    }
    switch ("hello" + ("\n\t" + "hello")) {
    case "foo bar":
        switch "" {
        case "":
            v11 := GenS2{F0: ("foo bar" + ("a" + ("hello" + "foo bar"))), F1: 999}
            _ = v11
        case "\n\t":
            fmt.Println(v2)
        default:
            if true {
                v12 := GenS1{F0: 667}
                _ = v12
            } else {
                v13 := GenS2{F0: "hello", F1: 4294967295}
                _ = v13
            }
        }
    default:
        v2 -= (v2 / (v2 - v2))
    }
    for fi14 := 0; fi14 < 19; fi14++ {
        var v15 uint = 255
        if ((!(42 > 510)) || false) {
            fmt.Println(fi14, fi14)
            v15 += (v15 | 2)
            if (0 == 100) {
                if ((0.0 / 1.0) == float64(((v2 + v2) - 3.14159))) {
                    var v16 float32 = v2
                    v17 := []int{fi14, 9223372036854775807, (fi14 ^ (fi14 >> ((2) & 63)))}
                    var v18 uint64 = 1
                    _ = v16
                    _ = v17
                    _ = v18
                } else {
                    if true {
                        switch (fi14 + fi14) {
                        case 1:
                            v19 := []int{}
                            _ = v19
                        case 128:
                            switch ("foo bar" + "\n\t") {
                            case "a":
                                v20 := GenS3{F0: 87, F1: 906, F2: 223}
                                _ = v20
                            case "\n\t":
                                v21 := v3[fi14]
                                _ = v21
                            case "":
                                var v22 int32 = 0
                                _ = v22
                            default:
                                delete(v3, (fi14 % (((fi14 * -2147483648)) | 1)))
                            }
                        }
                    }
                }
                var v23 bool = true
                var v24 int8 = -123
                _ = v23
                _ = v24
            } else {
                var v25 int = (fi14 ^ -9223372036854775808)
                _ = v25
            }
        } else {
            v15 += 25
            if false {
                switch "hello" {
                case "\n\t":
                    for fi26 := 0; fi26 < 8; fi26++ {
                        var v27 uint8 = 35
                        v27 = v27
                        var v28 float64 = (1e10 - 3.14159)
                        _ = fi26
                        _ = v27
                        _ = v28
                    }
                case "hello":
                    v29 := []int{fi14}
                    _ = v29
                case "a":
                    v30 := GenS4{F0: -339, F1: 151}
                    _ = v30
                default:
                    v15 = (v15 / ((v15) | 1))
                }
            }
            var v31 int32 = 201
            _ = v31
        }
        for fi32 := 0; fi32 < 2; fi32++ {
            v2 = v2
            fmt.Println(fi32, v15, v15)
            v33 := GenS1{F0: 42}
            _ = fi32
            _ = v33
        }
        _ = fi14
        _ = v15
    }
    v34 := []int{2147483647}
    var v35 uint16 = 1
    if false {
        var v36 int32 = -1
        _ = v36
    }
    v37 := GenS2{F0: ("foo bar" + "foo bar"), F1: 571}
    if true {
        v38 := GenS5{F0: 18446744073709551615}
        _ = len(v1)
        v35 += (v35 * 37)
        _ = v38
    } else {
        delete(v1, 265)
    }
    for fi39 := 0; fi39 < 16; fi39++ {
        v40 := GenS6{F0: (v35 % ((v35) | 1))}
        v35 = v35
        for fi41 := 0; fi41 < 11; fi41++ {
            _ = len(v3)
            _ = fi41
        }
        _ = fi39
        _ = v40
    }
    fmt.Println(v35)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v37
}
