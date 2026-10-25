package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8; F1 int; F2 int64 }
func genFn0() int16 {
    var v7 uint64 = 66
    for fi8 := 0; fi8 < 7; fi8++ {
        for fi9 := 0; fi9 < 7; fi9++ {
            v7 -= 1
            v10 := []int{(fi9 - (fi9 % ((fi8) | 1))), (fi8 ^ (fi9 & fi8)), fi8}
            var v11 uint = 279
            if len(v10) > 0 {
                v10[0] = fi8
            }
            if false {
                v11 = (^(v11))
                fmt.Println(v11, fi8)
            }
            _ = fi9
            _ = v10
            _ = v11
        }
        v12 := GenS1{F0: "hello"}
        _ = fi8
        _ = v12
    }
    _ = v7
    return -32768
}
func genFn1() uint16 {
    var v13 int8 = 127
    fmt.Println(v13)
    var v14 int8 = 42
    _ = v13
    _ = v14
    return 255
}

func main() {
    v1 := make(map[string]string)
    var v2 bool = false
    var v3 int64 = 9223372036854775807
    var v4 string = "hello"
    var v5 int16 = -440
    var v6 uint16 = 0
    v15 := GenS1{F0: ("\n\t" + v4)}
    v3 += (v3 * (v3 * -235))
    v16 := GenS1{F0: ""}
    v17 := GenS1{F0: (v4 + v4)}
    var v18 int32 = 1
    v2 = ((v18 <= (-(v18))) || (!(32768 >= -1)))
    switch v4 {
    case "hello":
        for fi19 := 0; fi19 < 4; fi19++ {
            var v20 int8 = -54
            if (v2 && false) {
                v18 += 1
                var v21 int16 = (-(v5))
                _ = v21
            } else {
                v22 := []int{(fi19 << ((255) & 63)), -989}
                _ = v22
            }
            _ = fi19
            _ = v20
        }
    case "a":
        v23 := []int{128, int(v18), 127}
        _ = v23
    case "foo bar":
        var v24 int = 9223372036854775807
        _ = v24
    }
    v2 = (v2 || (true && (32768 <= -9223372036854775808)))
    switch -9223372036854775808 {
    case 128:
        if (630 >= 18446744073709551615) {
            v25 := []int{9223372036854775807, -1, int(v6), -2147483648}
            var v26 int = 806
            v6 -= v6
            _ = v25
            _ = v26
        }
    case 416:
        switch v4 {
        case "":
            if ((v6 >> ((42) & 63)) != v6) {
                if false {
                    var v27 int32 = (v18 - (v18 * v18))
                    v28 := v1[v4]
                    _ = v27
                    _ = v28
                } else {
                    switch 659 {
                    case -9223372036854775808:
                        if (v6 != (^(v6))) {
                            v29 := GenS1{F0: v4}
                            v30 := GenS1{F0: v4}
                            _ = v29
                            _ = v30
                        }
                    case 0:
                        v31 := v1[v4]
                        _ = v31
                    case -489:
                        v6 -= 0
                    default:
                        switch 625 {
                        case -618:
                            for fi32 := 0; fi32 < 13; fi32++ {
                                var v33 int32 = v18
                                v34 := []int{2147483647}
                                if (42 > 128) {
                                    v34 = append(v34, 538)
                                    fmt.Println(v2)
                                    var v35 uint = 675
                                    _ = v35
                                }
                                _ = fi32
                                _ = v33
                                _ = v34
                            }
                        }
                    }
                }
                if (-775 <= 32767) {
                    fmt.Println(v6, v18)
                    v4 = ((v4 + v4) + v4)
                    var v36 float64 = 1.5
                    _ = v36
                }
                var v37 float64 = float64(-2.25)
                _ = v37
            }
        case "hello":
            v38 := GenS1{F0: (v4 + (v4 + ((v4 + v4) + "foo bar")))}
            _ = v38
        default:
            v39 := GenS2{F0: v4}
            _ = v39
        }
    case 0:
        v40 := []int{-424}
        _ = v40
    default:
        var v41 string = v4
        _ = v41
    }
    if v2 {
        v42 := []int{}
        if len(v42) > 0 {
            v42[0] = 2147483647
        }
        if (v2 || v2) {
            if (!((v3 & v3) >= v3)) {
                fmt.Println(v5, v5)
                v1[((("hello" + v4) + "hello") + ((v4 + v4) + ("foo bar" + (("a" + v4) + v4))))] = v4
            } else {
                switch 1 {
                case 128:
                    v2 = ((-618 > 42) || v2)
                }
                var v43 int64 = 0
                fmt.Println(v43, v2)
                _ = v43
            }
        } else {
            v5 = (v5 + genFn0())
            v2 = (float64(1.5) > 1.5)
            v44 := GenS3{F0: 0, F1: 42, F2: -9223372036854775808}
            _ = v44
        }
        _ = v42
    }
    var v45 float64 = 0.1
    v4 = v4
    v5 = genFn0()
    v46 := []int{-9223372036854775808, 953, 128}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v46:", len(v46), v46)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v5
    _ = v6
}
