package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 int }
type GenS2 struct { F0 int; F1 int64; F2 int32 }
type GenS3 struct { F0 uint64; F1 int16; F2 uint32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 string; F1 uint; F2 int32 }
type GenS6 struct { F0 string; F1 uint8; F2 int32 }
type GenS7 struct { F0 int32; F1 uint16 }
func genFn0(p0 uint8, p1 uint8) int8 {
    var v7 int = 32768
    fmt.Println(v7, p0)
    fmt.Println(p1)
    var v8 uint8 = (p0 | (p0 | 128))
    switch (v7 % ((v7) | 1)) {
    case 127:
        for fi9 := 0; fi9 < 6; fi9++ {
            p1 = (^(p1))
            for fi10 := 0; fi10 < 9; fi10++ {
                if ((fi9 >> ((42) & 63)) >= 368) {
                    var v11 int64 = -612
                    _ = v11
                } else {
                    switch (fi9 & fi10) {
                    case 0:
                        p0 -= v8
                    case 32767:
                        fmt.Println(fi9, v7)
                    }
                    var v12 bool = false
                    if v12 {
                        var v13 int64 = 1
                        var v14 string = "foo bar"
                        _ = v13
                        _ = v14
                    } else {
                        v15 := GenS1{F0: (!v12), F1: 1, F2: v7}
                        v16 := GenS1{F0: v12, F1: int16(fi9), F2: -474}
                        _ = v15
                        _ = v16
                    }
                    if (false || (0 != 0)) {
                        var v17 float64 = 1e10
                        if v12 {
                            v18 := GenS2{F0: (^(fi10)), F1: -1, F2: 0}
                            if (!(4294967295 == 1)) {
                                switch fi9 {
                                case 42:
                                    v17 += v17
                                }
                                var v19 int16 = -1
                                p1 += v8
                                _ = v19
                            } else {
                                v20 := []int{(fi10 - fi10), (-(fi9)), 2147483647, (fi9 & 32768)}
                                p0 += p1
                                var v21 uint32 = 890
                                v20 = append(v20, 2147483647)
                                _ = v20
                                _ = v21
                            }
                            _ = v18
                        } else {
                            var v22 string = ("a" + "\n\t")
                            _ = v22
                        }
                        _ = v17
                    } else {
                        var v23 string = (("a" + "hello") + ("\n\t" + "foo bar"))
                        _ = v23
                    }
                    _ = v12
                }
                v24 := GenS2{F0: -1, F1: 34, F2: -2147483648}
                v25 := GenS2{F0: (v7 / (((fi9 << ((328) & 63))) | 1)), F1: 9223372036854775807, F2: 0}
                _ = fi10
                _ = v24
                _ = v25
            }
            if true {
                p1 = (p1 - p1)
                switch "\n\t" {
                case "foo bar":
                    if (!(-1 > -1)) {
                        v26 := GenS1{F0: (9223372036854775807 != 881), F1: 0, F2: v7}
                        _ = v26
                    }
                case "a":
                    fmt.Println(fi9)
                default:
                    v27 := GenS3{F0: 395, F1: 1, F2: 42}
                    _ = v27
                }
                fmt.Println(v7, v8, v7)
            } else {
                p1 -= (v8 % (((p1 | p1)) | 1))
            }
            _ = fi9
        }
    default:
        if false {
            if (!(!false)) {
                fmt.Println(v8)
            } else {
                if false {
                    for fi28 := 0; fi28 < 8; fi28++ {
                        for fi29 := 0; fi29 < 7; fi29++ {
                            var v30 uint8 = 42
                            switch "\n\t" {
                            case "hello":
                                v31 := []int{-2147483648, 0}
                                _ = v31
                            case "":
                                fmt.Println(p1, p1)
                            }
                            _ = fi29
                            _ = v30
                        }
                        for fi32 := 0; fi32 < 4; fi32++ {
                            v33 := []int{(fi28 << ((2) & 63)), -9223372036854775808}
                            var v34 float64 = (-2.25 / 1.0)
                            p1 += 0
                            p1 = (^(v8))
                            _ = len(v33)
                            _ = fi32
                            _ = v33
                            _ = v34
                        }
                        _ = fi28
                    }
                }
                v7 = 42
            }
            p1 += 127
            p1 += (p0 | (v8 + (p0 ^ (v8 & (p1 % (((p1 ^ p0)) | 1))))))
        }
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return -128
}

func main() {
    v1 := make(map[string]string)
    var v2 int8 = 105
    v3 := make(map[int]int)
    var v4 int = int(v2)
    v5 := make(map[uint32]uint32)
    v6 := make(map[int]int)
    v35 := GenS4{F0: ("a" + "a")}
    if (!true) {
        var v36 uint8 = 128
        for fi37 := 0; fi37 < 5; fi37++ {
            if (344 > 442) {
                if (true && false) {
                    var v38 float64 = ((0.0 + (1e10 + -2.25)) + ((((float64(3.14159) + float64(3.14159)) / float64((1.5 * 1e10))) + 1e10) * (-2.25 * 3.14159)))
                    v39 := GenS4{F0: ""}
                    v40 := GenS4{F0: "hello"}
                    _ = v38
                    _ = v39
                    _ = v40
                } else {
                    var v41 string = "hello"
                    _ = v41
                }
                for fi42 := 0; fi42 < 10; fi42++ {
                    var v43 uint = 2
                    _ = fi42
                    _ = v43
                }
            }
            if true {
                if true {
                    fmt.Println(v2, v4)
                }
                v44 := GenS5{F0: "", F1: 2, F2: 1}
                fmt.Println(v2, v4)
                _ = v44
            }
            var v45 uint64 = 777
            _ = fi37
            _ = v45
        }
        _ = v36
    } else {
        if (!false) {
            var v46 int16 = 0
            _ = v46
        }
        v47 := GenS6{F0: "", F1: 0, F2: 2147483647}
        if (1 >= 42) {
            var v48 uint16 = 255
            _ = v48
        } else {
            v4 = (v4 & (v4 << ((761) & 63)))
            var v49 bool = true
            fmt.Println(v2, v2)
            _ = v49
        }
        _ = v47
    }
    var v50 uint8 = 188
    v50 += (^(v50))
    if ((18446744073709551615 <= 4294967295) || (("hello" > ("foo bar" + (((("\n\t" + "a") + ("a" + ("hello" + "\n\t"))) + "") + (("a" + "") + "hello")))) && true)) {
        v50 -= 127
        v2 -= v2
        v51 := GenS7{F0: 42, F1: 534}
        _ = v51
    }
    v2 = v2
    v52 := GenS5{F0: "hello", F1: 4294967295, F2: 2147483647}
    var v53 string = "a"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("v53:", v53)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v50
    _ = v52
    _ = v53
    _ = v6
}
