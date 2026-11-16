package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 uint }
type GenS2 struct { F0 uint; F1 string; F2 uint8 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int8; F1 int16; F2 uint8 }
type GenS5 struct { F0 uint8; F1 int64; F2 bool }
type GenS6 struct { F0 int32; F1 uint16 }
type GenS7 struct { F0 uint; F1 string; F2 int64 }
func genFn0(p0 int64, p1 int64) int32 {
    var v7 float32 = 0.1
    fmt.Println(v7)
    var v8 uint16 = 395
    v9 := GenS1{F0: 255, F1: 1, F2: 18446744073709551615}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return 2147483647
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[int]string)
    v3 := make(map[int]int)
    var v4 int8 = -121
    v5 := make(map[string]int)
    var v6 int = -74
    switch (((("hello" + ("" + ("\n\t" + "foo bar"))) + "foo bar") + "foo bar") + "hello") {
    case "foo bar":
        v4 = 127
    case "hello":
        switch ("hello" + "foo bar") {
        case "\n\t":
            switch ("" + "") {
            case "a":
                for fi10 := 0; fi10 < 10; fi10++ {
                    if true {
                        v4 = v4
                        switch fi10 {
                        case 404:
                            var v11 uint8 = 18
                            _ = v11
                        case -9223372036854775808:
                            if (750 > 4294967295) {
                                v6 -= int(v4)
                            } else {
                                switch v6 {
                                case 32768:
                                    _ = len(v3)
                                }
                                v1[42] = "foo bar"
                                v12 := []int{}
                                _ = v12
                            }
                        }
                    }
                    _ = fi10
                }
            case "hello":
                for fi13 := 0; fi13 < 18; fi13++ {
                    v14 := GenS2{F0: 255, F1: (("\n\t" + (("a" + ((("\n\t" + "") + "a") + "a")) + "foo bar")) + (("\n\t" + ((((("\n\t" + "") + (("foo bar" + (("hello" + "a") + "hello")) + "foo bar")) + ((("foo bar" + "a") + "foo bar") + "foo bar")) + ("foo bar" + "a")) + "")) + "a")), F2: 0}
                    fmt.Println(fi13)
                    _ = fi13
                    _ = v14
                }
            case "":
                v15 := []int{-484, (v6 / (((v6 >> ((4294967295) & 63))) | 1)), 0, v6}
                _ = v15
            }
        case "foo bar":
            var v16 int64 = 9223372036854775807
            _ = v16
        case "hello":
            switch "a" {
            case "hello":
                v6 -= v6
            case "foo bar":
                v1[2147483647] = "foo bar"
            }
        default:
            v6 = (v6 << ((18446744073709551615) & 63))
        }
    case "a":
        if false {
            var v17 uint8 = 255
            v18 := []int{877, v6, v6, (v6 & v6)}
            for fi19 := 0; fi19 < 9; fi19++ {
                v20 := GenS3{F0: 65535}
                delete(v1, uint32(v17))
                v21 := GenS1{F0: 65535, F1: v4, F2: 18446744073709551615}
                _ = fi19
                _ = v20
                _ = v21
            }
            _ = v17
            _ = v18
        } else {
            var v22 int32 = -973
            _ = v22
        }
    default:
        v23 := GenS4{F0: v4, F1: 1, F2: 75}
        _ = v23
    }
    if true {
        var v24 string = "a"
        if (42 == 2) {
            var v25 int16 = 32767
            if (1e10 >= -2.25) {
                var v26 int8 = -128
                v27 := GenS5{F0: 21, F1: 42, F2: false}
                _ = v26
                _ = v27
            } else {
                v25 -= (v25 * (v25 | v25))
                var v28 uint8 = 72
                _ = v28
            }
            _ = v25
        } else {
            var v29 uint8 = 128
            _ = len(v2)
            _ = v29
        }
        v6 = v6
        _ = v24
    }
    if true {
        if (133 > 0) {
            for fi30 := 0; fi30 < 3; fi30++ {
                v31 := GenS6{F0: 415, F1: 0}
                fmt.Println(fi30)
                for fi32 := 0; fi32 < 17; fi32++ {
                    var v33 int = v6
                    _ = fi32
                    _ = v33
                }
                _ = fi30
                _ = v31
            }
            v4 = 0
        } else {
            var v34 string = ("a" + "foo bar")
            var v35 string = v34
            var v36 bool = ((!((float32(3.14159) + (-2.25 * (0.0 * 1.5))) > 0.1)) && false)
            _ = v34
            _ = v35
            _ = v36
        }
        var v37 uint32 = 275
        _ = v37
    } else {
        for fi38 := 0; fi38 < 15; fi38++ {
            v4 = (-(v4))
            _ = fi38
        }
    }
    v39 := GenS7{F0: 4294967295, F1: "\n\t", F2: 186}
    for fi40 := 0; fi40 < 16; fi40++ {
        switch int(v4) {
        case 32767:
            var v41 int16 = 608
            _ = v41
        case -9223372036854775808:
            if true {
                var v42 string = (("foo bar" + "a") + (("" + ("foo bar" + "\n\t")) + "foo bar"))
                v43 := v2[fi40]
                v6 = fi40
                _ = v42
                _ = v43
            } else {
                if ((!true) && false) {
                    if ((640 <= 49) || false) {
                        for fi44 := 0; fi44 < 7; fi44++ {
                            v6 = v6
                            v45 := []int{-979, (-(fi44)), v6, (v6 | (fi44 / ((2147483647) | 1)))}
                            if false {
                                v46 := v3[v6]
                                var v47 int8 = 42
                                _ = v46
                                _ = v47
                            }
                            _ = fi44
                            _ = v45
                        }
                        fmt.Println(v6, fi40)
                    } else {
                        v48 := []int{(fi40 + v6)}
                        for fi49 := 0; fi49 < 11; fi49++ {
                            v50 := GenS3{F0: 386}
                            v6 -= 0
                            _ = fi49
                            _ = v50
                        }
                        var v51 int16 = 42
                        _ = v48
                        _ = v51
                    }
                    fmt.Println(v6)
                }
            }
        default:
            v52 := []int{}
            _ = v52
        }
        var v53 bool = false
        _ = fi40
        _ = v53
    }
    for fi54 := 0; fi54 < 3; fi54++ {
        var v55 float64 = 3.14159
        _ = fi54
        _ = v55
    }
    if (128 >= 255) {
        for fi56 := 0; fi56 < 13; fi56++ {
            switch (v6 % ((32767) | 1)) {
            case 0:
                var v57 uint = 4294967295
                _ = v57
            case 2147483647:
                v6 = (fi56 << ((0) & 63))
            }
            v58 := GenS5{F0: 130, F1: -734, F2: true}
            v6 = (fi56 | 263)
            _ = fi56
            _ = v58
        }
    }
    fmt.Println(v4, v6)
    var v59 bool = true
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v59:", v59)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v59
    _ = v6
}
