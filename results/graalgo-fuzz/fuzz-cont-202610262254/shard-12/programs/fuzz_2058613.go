package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 uint32 }
type GenS2 struct { F0 uint64; F1 string; F2 int16 }
type GenS3 struct { F0 bool; F1 int8 }
type GenS4 struct { F0 int64; F1 int; F2 uint64 }
type GenS5 struct { F0 int32; F1 int16 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 bool; F1 int64 }
type GenS9 struct { F0 int64; F1 int16 }
type GenS10 struct { F0 int8 }
func genFn0() bool {
    if true {
        if false {
            v4 := GenS1{F0: 4294967295, F1: 1000, F2: 65535}
            var v5 int = -567
            v6 := []int{(v5 << ((910) & 63)), v5, 0, (v5 & v5)}
            _ = v4
            _ = v5
            _ = v6
        }
    }
    v7 := GenS2{F0: 670, F1: "hello", F2: 1}
    if (false || (32767 != 0)) {
        var v8 int32 = -2147483648
        v8 -= (v8 << ((65535) & 63))
        _ = v8
    } else {
        if (!("\n\t" >= "hello")) {
            switch ("a" + "\n\t") {
            case "hello":
                switch 128 {
                case 803:
                    var v9 float32 = 1e10
                    _ = v9
                case 32768:
                    var v10 bool = true
                    _ = v10
                case 127:
                    var v11 uint = 255
                    _ = v11
                }
            case "\n\t":
                var v12 uint32 = 4294967295
                _ = v12
            case "a":
                if (-1 >= 1) {
                    v13 := GenS3{F0: true, F1: -1}
                    v14 := []int{-143, 1, -9223372036854775808, 0}
                    _ = v13
                    _ = v14
                } else {
                    v15 := GenS4{F0: -227, F1: -9223372036854775808, F2: 4294967295}
                    _ = v15
                }
            default:
                v16 := []int{51, -9223372036854775808, 2147483647}
                _ = v16
            }
            switch "" {
            case "foo bar":
                switch ((("foo bar" + "a") + "\n\t") + "\n\t") {
                case "\n\t":
                    v17 := []int{0, 42}
                    _ = v17
                case "hello":
                    var v18 float64 = float64((1e10 * 3.14159))
                    _ = v18
                case "":
                    if true {
                        v19 := []int{}
                        if true {
                            v20 := GenS1{F0: 659, F1: 32767, F2: 2147483647}
                            if false {
                                v21 := GenS2{F0: 4294967295, F1: "\n\t", F2: -1}
                                _ = v21
                            } else {
                                if (128 <= 255) {
                                    var v22 float32 = (3.14159 * 1e10)
                                    switch 4294967295 {
                                    case -1:
                                        var v23 int = 9223372036854775807
                                        _ = v23
                                    case 2147483647:
                                        fmt.Println(v22)
                                    default:
                                        for fi24 := 0; fi24 < 12; fi24++ {
                                            if true {
                                                fmt.Println(fi24, fi24)
                                                var v25 float32 = v22
                                                _ = v25
                                            } else {
                                                v22 = v22
                                            }
                                            if (!false) {
                                                _ = len(v19)
                                                v22 = (((v22 - (float32((0.0 / -2.25)) + float32((1e10 + 1.5)))) * 1.5) - 3.14159)
                                            }
                                            if true {
                                                var v26 string = "a"
                                                _ = v26
                                            }
                                            _ = fi24
                                        }
                                    }
                                    for fi27 := 0; fi27 < 20; fi27++ {
                                        var v28 uint32 = 2147483647
                                        var v29 int8 = 100
                                        switch -9223372036854775808 {
                                        case 42:
                                            var v30 uint = 255
                                            _ = v30
                                        }
                                        _ = fi27
                                        _ = v28
                                        _ = v29
                                    }
                                    _ = v22
                                } else {
                                    if (1 >= 128) {
                                        var v31 uint32 = 239
                                        switch int(v31) {
                                        case 128:
                                            if len(v19) >= 2 {
                                                v19 = v19[1:len(v19)]
                                            }
                                        default:
                                            if len(v19) >= 2 {
                                                v19 = v19[1:len(v19)]
                                            }
                                        }
                                        if len(v19) >= 2 {
                                            v19 = v19[1:len(v19)]
                                        }
                                        _ = v31
                                    }
                                }
                                for fi32 := 0; fi32 < 16; fi32++ {
                                    var v33 bool = true
                                    if (!v33) {
                                        var v34 uint32 = 2147483647
                                        fmt.Println(fi32, v34, v34)
                                        _ = v34
                                    } else {
                                        if false {
                                            v33 = (42 <= -1)
                                            fmt.Println(v33)
                                            switch "foo bar" {
                                            case "hello":
                                                for fi35 := 0; fi35 < 7; fi35++ {
                                                    if v33 {
                                                        var v36 float32 = (float32(-2.25) + -2.25)
                                                        _ = v36
                                                    }
                                                    fmt.Println(fi32, fi32)
                                                    v37 := GenS1{F0: 2, F1: 0, F2: 334}
                                                    _ = fi35
                                                    _ = v37
                                                }
                                            case "foo bar":
                                                for fi38 := 0; fi38 < 1; fi38++ {
                                                    if v33 {
                                                        if v33 {
                                                            var v39 uint64 = 0
                                                            var v40 int8 = -1
                                                            _ = v39
                                                            _ = v40
                                                        }
                                                        v41 := GenS5{F0: 2147483647, F1: -450}
                                                        if len(v19) > 0 {
                                                            v19[0] = 944
                                                        }
                                                        _ = v41
                                                    } else {
                                                        var v42 uint64 = 496
                                                        var v43 uint64 = 997
                                                        _ = v42
                                                        _ = v43
                                                    }
                                                    _ = fi38
                                                }
                                            case "a":
                                                var v44 uint16 = 0
                                                _ = v44
                                            }
                                        }
                                        v33 = ((((0.1 != ((((-2.25 + 0.1) - -2.25) / 1e10) / 1.0)) || (320 > 255)) && ("\n\t" == "hello")) && v33)
                                        var v45 uint = 65535
                                        _ = v45
                                    }
                                    _ = fi32
                                    _ = v33
                                }
                            }
                            var v46 uint16 = 32767
                            _ = v20
                            _ = v46
                        }
                        v47 := GenS6{F0: 42}
                        _ = v19
                        _ = v47
                    }
                }
            case "":
                var v48 float32 = -2.25
                _ = v48
            case "a":
                v49 := GenS7{F0: (("foo bar" + ("hello" + (("a" + "") + ("hello" + "")))) + "a")}
                _ = v49
            }
        }
        for fi50 := 0; fi50 < 5; fi50++ {
            for fi51 := 0; fi51 < 19; fi51++ {
                var v52 bool = false
                v52 = v52
                var v53 int16 = 907
                _ = fi51
                _ = v52
                _ = v53
            }
            var v54 string = ("hello" + "foo bar")
            _ = fi50
            _ = v54
        }
        switch -1 {
        case -410:
            var v55 string = ""
            _ = v55
        }
    }
    _ = v7
    return (1 <= 73)
}
func genFn1() string {
    for fi56 := 0; fi56 < 19; fi56++ {
        v57 := GenS2{F0: 42, F1: "a", F2: -660}
        v58 := GenS8{F0: false, F1: 42}
        _ = fi56
        _ = v57
        _ = v58
    }
    var v59 int8 = 127
    fmt.Println(v59)
    _ = v59
    return "foo bar"
}

func main() {
    var v1 uint32 = 991
    v2 := []uint32{v1}
    var v3 int = 1
    var v60 bool = true
    v61 := GenS9{F0: 438, F1: 33}
    var v62 uint = 1
    v63 := GenS2{F0: 42, F1: (("a" + "\n\t") + ""), F2: 32767}
    v64 := GenS10{F0: 0}
    fmt.Println(v62, v62)
    for fi65 := 0; fi65 < 10; fi65++ {
        v60 = true
        fmt.Println(fi65, v62)
        for fi66 := 0; fi66 < 13; fi66++ {
            var v67 uint32 = v1
            _ = fi66
            _ = v67
        }
        _ = fi65
    }
    if len(v2) > 0 {
        v2[0] = (v1 >> ((18446744073709551615) & 63))
    }
    if len(v2) > 0 {
        v2[0] = (^(v1))
    }
    v68 := GenS2{F0: 160, F1: "", F2: 42}
    fmt.Println(v1)
    if len(v2) > 0 {
        v2[0] = v1
    }
    var v69 uint = 1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v60:", v60)
    fmt.Println("v62:", v62)
    fmt.Println("v69:", v69)
    _ = v1
    _ = v2
    _ = v3
    _ = v60
    _ = v61
    _ = v62
    _ = v63
    _ = v64
    _ = v68
    _ = v69
}
