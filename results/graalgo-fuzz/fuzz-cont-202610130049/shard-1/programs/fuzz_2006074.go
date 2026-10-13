package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 uint }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint64; F1 int; F2 uint32 }
type GenS4 struct { F0 int64; F1 uint }
type GenS5 struct { F0 int64; F1 int }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int; F1 bool; F2 uint64 }
type GenS8 struct { F0 string; F1 uint64; F2 string }
func genFn0() uint64 {
    for fi4 := 0; fi4 < 5; fi4++ {
        if (!true) {
            switch ("" + "hello") {
            case "":
                v5 := GenS1{F0: 1, F1: 32767, F2: 670}
                _ = v5
            case "a":
                fmt.Println(fi4)
            case "\n\t":
                v6 := GenS2{F0: (("a" + (("" + ("" + (((("\n\t" + (("hello" + ("\n\t" + "foo bar")) + ("foo bar" + ""))) + "\n\t") + "hello") + "a"))) + "")) + ("a" + ((("a" + "\n\t") + "hello") + "a")))}
                _ = v6
            default:
                var v7 uint64 = 0
                _ = v7
            }
            v8 := GenS3{F0: 42, F1: 32767, F2: 375}
            _ = v8
        } else {
            var v9 float64 = 1.5
            fmt.Println(fi4, v9)
            fmt.Println(v9, v9)
            _ = v9
        }
        v10 := []int{fi4, fi4, (fi4 << ((2) & 63))}
        _ = fi4
        _ = v10
    }
    v11 := GenS4{F0: 1, F1: 923}
    v12 := []int{9223372036854775807, 127, 1}
    _ = v11
    _ = v12
    return 397
}

func main() {
    var v1 float32 = 1e10
    var v2 int16 = -32768
    var v3 int16 = (v2 * (v2 / ((v2) | 1)))
    switch -58 {
    case 840:
        if (-350 > -1) {
            for fi13 := 0; fi13 < 17; fi13++ {
                if true {
                    v1 += v1
                    var v14 string = ""
                    _ = v14
                }
                switch (("a" + ("foo bar" + "\n\t")) + "hello") {
                case "\n\t":
                    if false {
                        v3 += 32767
                    } else {
                        v15 := GenS5{F0: 0, F1: fi13}
                        var v16 int16 = (v3 << ((1) & 63))
                        _ = v15
                        _ = v16
                    }
                case "":
                    v17 := []int{}
                    _ = v17
                }
                v3 = v3
                _ = fi13
            }
            switch 127 {
            case -74:
                v18 := GenS6{F0: 981}
                _ = v18
            case -679:
                v19 := []int{-9223372036854775808, -1}
                _ = v19
            default:
                v2 -= 1
            }
            v20 := []int{9223372036854775807, -1, -9223372036854775808, -9223372036854775808}
            _ = v20
        }
    case -906:
        v21 := GenS3{F0: 18446744073709551615, F1: 127, F2: 0}
        _ = v21
    }
    v2 = (^(v3))
    v2 = v3
    v22 := GenS7{F0: 128, F1: (-282 >= -1), F2: 18446744073709551615}
    var v23 string = (("hello" + "a") + ("\n\t" + ("\n\t" + ((((("hello" + ("\n\t" + ("hello" + ("" + ("foo bar" + "foo bar"))))) + "foo bar") + (("foo bar" + "a") + "a")) + "foo bar") + ("a" + "hello")))))
    var v24 string = ((v23 + (("a" + v23) + (v23 + v23))) + v23)
    v25 := GenS8{F0: "hello", F1: 1, F2: "foo bar"}
    v26 := []int{737, -837}
    for fi27 := 0; fi27 < 16; fi27++ {
        v3 -= (v2 / ((v3) | 1))
        switch (v23 + "foo bar") {
        case "":
            if (42 < 1) {
                v2 = v3
            } else {
                if true {
                    var v28 uint16 = 144
                    v28 -= v28
                    for fi29 := 0; fi29 < 16; fi29++ {
                        fmt.Println(v1, v2, v28)
                        if true {
                            if (false || false) {
                                if len(v26) > 0 {
                                    v26[0] = 128
                                }
                                var v30 int64 = 539
                                v31 := GenS7{F0: fi29, F1: (uint64(v28) == 42), F2: genFn0()}
                                _ = v30
                                _ = v31
                            }
                            _ = len(v26)
                            var v32 int16 = v3
                            _ = v32
                        } else {
                            switch "\n\t" {
                            case "\n\t":
                                if (127 < 0) {
                                    fmt.Println(v28)
                                    var v33 bool = false
                                    var v34 string = (v23 + ((((v24 + "\n\t") + (((v23 + v24) + "foo bar") + ((("" + "a") + "a") + (v24 + v24)))) + (((v23 + v24) + v24) + "")) + "foo bar"))
                                    v26 = append(v26, (fi29 - (fi29 + 0)))
                                    v23 = ((("hello" + "foo bar") + "a") + ("a" + "hello"))
                                    _ = v33
                                    _ = v34
                                } else {
                                    v1 -= v1
                                    switch 4294967295 {
                                    case -510:
                                        v26 = append(v26, (fi27 - (^(fi27))))
                                    case 0:
                                        _ = len(v26)
                                    case -1:
                                        var v35 bool = (v24 >= (v23 + (v23 + "hello")))
                                        _ = v35
                                    default:
                                        if ((!false) || (1 <= 1)) {
                                            v26 = append(v26, (fi29 >> ((284) & 63)))
                                        } else {
                                            v23 = (v24 + ("\n\t" + v23))
                                        }
                                    }
                                }
                            case "hello":
                                switch 1 {
                                case -1:
                                    v26 = append(v26, fi29)
                                default:
                                    if ((v28 & (^(v28))) > (v28 / ((42) | 1))) {
                                        var v36 float32 = v1
                                        switch v24 {
                                        case "a":
                                            v26 = append(v26, (fi27 | fi29))
                                        default:
                                            var v37 string = "hello"
                                            _ = v37
                                        }
                                        _ = v36
                                    } else {
                                        v2 = v3
                                        if len(v26) > 0 {
                                            v26[0] = fi27
                                        }
                                        v2 -= (^(v3))
                                    }
                                }
                            case "a":
                                if true {
                                    if len(v26) > 0 {
                                        v26[0] = fi29
                                    }
                                } else {
                                    if len(v26) > 0 {
                                        v26[0] = 9223372036854775807
                                    }
                                }
                            }
                        }
                        _ = fi29
                    }
                    _ = v28
                } else {
                    v24 = "hello"
                }
            }
        case "\n\t":
            var v38 bool = false
            _ = v38
        case "hello":
            if true {
                if (42 >= 128) {
                    v1 -= v1
                    var v39 uint = 0
                    _ = v39
                }
                v3 -= (v3 & v3)
                var v40 int64 = 9223372036854775807
                _ = v40
            }
        }
        _ = fi27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
}
