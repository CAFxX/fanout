package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 uint; F1 bool; F2 uint64 }
type GenS3 struct { F0 string; F1 int16 }
func genFn0(p0 int16) int64 {
    for fi7 := 0; fi7 < 8; fi7++ {
        p0 -= 593
        _ = fi7
    }
    var v8 float32 = 0.0
    _ = p0
    _ = v8
    return 42
}

func main() {
    v1 := []uint32{2147483647, 965, 65535}
    var v2 bool = false
    var v3 string = ""
    var v4 uint = 601
    var v5 bool = v2
    var v6 int32 = 42
    var v9 float64 = (0.1 / 1.5)
    v10 := GenS1{F0: (v6 - 2147483647), F1: 9223372036854775807}
    var v11 bool = false
    var v12 int16 = 42
    switch 4294967295 {
    case 127:
        v13 := GenS2{F0: 4294967295, F1: v11, F2: 905}
        _ = v13
    case 235:
        fmt.Println(v5)
    default:
        v11 = false
    }
    v12 += v12
    for fi14 := 0; fi14 < 4; fi14++ {
        if v2 {
            var v15 int = fi14
            _ = v15
        }
        if len(v1) > 0 {
            v1[0] = 158
        }
        _ = fi14
    }
    v11 = ((!true) || (-922 >= 2147483647))
    v1 = append(v1, 65535)
    switch v3 {
    case "foo bar":
        switch -642 {
        case 9223372036854775807:
            var v16 int = -996
            _ = v16
        case 42:
            fmt.Println(v2, v5, v11)
        case -9223372036854775808:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        default:
            var v17 float32 = (0.1 + (1.5 - 1.5))
            _ = v17
        }
    }
    if true {
        fmt.Println(v4, v2)
        for fi18 := 0; fi18 < 4; fi18++ {
            _ = len(v1)
            var v19 uint = (v4 & 49)
            _ = fi18
            _ = v19
        }
        var v20 string = v3
        _ = v20
    } else {
        v2 = (255 >= 51)
        var v21 float64 = v9
        for fi22 := 0; fi22 < 5; fi22++ {
            fmt.Println(v2, v21)
            if (42 >= 255) {
                if len(v1) > 0 {
                    v1[0] = 749
                }
            }
            v23 := GenS2{F0: v4, F1: v2, F2: 4294967295}
            _ = fi22
            _ = v23
        }
        _ = v21
    }
    var v24 float32 = 1.5
    v12 += 1
    if (((976 > 0) && (42 == 18446744073709551615)) && (0 != v6)) {
        var v25 int32 = (v6 | (v6 % (((v6 >> ((v4) & 63))) | 1)))
        if (-128 != 0) {
            var v26 uint8 = 251
            switch v3 {
            case "\n\t":
                if (v5 || (!true)) {
                    v26 += 127
                    var v27 uint64 = 347
                    for fi28 := 0; fi28 < 13; fi28++ {
                        switch -678 {
                        case 566:
                            v29 := GenS3{F0: ("foo bar" + "foo bar"), F1: (v12 + -1)}
                            _ = v29
                        case 32767:
                            if len(v1) > 0 {
                                v1[0] = 1
                            }
                        case 2147483647:
                            v30 := GenS3{F0: ((v3 + v3) + ((("" + v3) + v3) + v3)), F1: v12}
                            _ = v30
                        default:
                            fmt.Println(v6, v12)
                        }
                        v3 = ((v3 + ("a" + "foo bar")) + v3)
                        _ = fi28
                    }
                    _ = v27
                } else {
                    _ = len(v1)
                    v6 -= 42
                    switch (v3 + ("a" + v3)) {
                    case "foo bar":
                        v31 := GenS1{F0: v25, F1: -2147483648}
                        _ = v31
                    }
                }
            }
            _ = v26
        } else {
            v5 = false
            v32 := GenS2{F0: (v4 >> (((v4 - 446)) & 63)), F1: (!("foo bar" <= (((v3 + (("\n\t" + "") + v3)) + v3) + v3))), F2: 0}
            _ = v32
        }
        _ = v25
    } else {
        fmt.Println(v5, v2, v11)
        if v11 {
            switch v3 {
            case "a":
                switch 261 {
                case -783:
                    fmt.Println(v3, v2)
                case 863:
                    var v33 bool = (374 == 2)
                    _ = v33
                case 42:
                    _ = len(v1)
                }
            case "hello":
                v34 := GenS3{F0: "foo bar", F1: -1}
                _ = v34
            case "foo bar":
                v11 = true
            default:
                v5 = (v12 != v12)
            }
        } else {
            v3 = (("\n\t" + (v3 + (("foo bar" + (v3 + v3)) + v3))) + "\n\t")
            for fi35 := 0; fi35 < 14; fi35++ {
                v6 = (v6 + (v6 & (-(v6))))
                v2 = (((v24 > v24) && (genFn0(1) <= 9223372036854775807)) || (v12 >= (v12 & (^(v12)))))
                var v36 string = (((((v3 + (("foo bar" + ("" + (v3 + v3))) + "a")) + (((v3 + "a") + v3) + v3)) + "\n\t") + (v3 + "a")) + v3)
                _ = fi35
                _ = v36
            }
            fmt.Println(v4, v5)
        }
    }
    for fi37 := 0; fi37 < 7; fi37++ {
        var v38 bool = v2
        _ = fi37
        _ = v38
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
