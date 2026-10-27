package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16; F1 uint8 }
type GenS3 struct { F0 int16; F1 int64 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 bool }
func genFn0(p0 int32) int32 {
    for fi6 := 0; fi6 < 16; fi6++ {
        v7 := []int{-9223372036854775808, 1}
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 14; fi8++ {
        fmt.Println(p0, fi8)
        _ = fi8
    }
    _ = p0
    return 42
}

func main() {
    var v1 uint64 = 1
    var v2 float64 = (0.0 / ((((0.1 * 1e10) + (3.14159 * 1e10)) / float64((-2.25 / 0.1))) + 1.5))
    v3 := make(map[uint32]int)
    v4 := make(map[int]int)
    var v5 float32 = 0.0
    for fi9 := 0; fi9 < 12; fi9++ {
        var v10 uint16 = 596
        var v11 uint8 = 131
        v12 := []int{(fi9 | 127)}
        _ = fi9
        _ = v10
        _ = v11
        _ = v12
    }
    _ = len(v4)
    v13 := []int{}
    if false {
        var v14 uint16 = 322
        switch 555 {
        case 2147483647:
            switch "hello" {
            case "hello":
                switch (("foo bar" + ((("foo bar" + ("foo bar" + "")) + (("hello" + (("a" + "\n\t") + ("\n\t" + ("\n\t" + "a")))) + "foo bar")) + "hello")) + "foo bar") {
                case "a":
                    v1 += (v1 | v1)
                case "\n\t":
                    fmt.Println(v5, v1, v1)
                }
            case "\n\t":
                for fi15 := 0; fi15 < 5; fi15++ {
                    v1 += v1
                    v1 = (^(v1))
                    for fi16 := 0; fi16 < 4; fi16++ {
                        fmt.Println(fi16)
                        v3[295] = 1
                        _ = fi16
                    }
                    _ = fi15
                }
            case "a":
                if false {
                    v5 -= (v5 + (-2.25 / v5))
                } else {
                    if false {
                        var v17 string = "foo bar"
                        var v18 int16 = 0
                        if (1 == int32(v1)) {
                            v14 += 188
                            v2 -= v2
                            var v19 int16 = (v18 % (((-(v18))) | 1))
                            _ = v19
                        }
                        _ = v17
                        _ = v18
                    } else {
                        var v20 int32 = 1
                        _ = v20
                    }
                    var v21 int64 = 954
                    _ = v21
                }
            }
        case -9223372036854775808:
            if len(v13) > 0 {
                v13[0] = 4294967295
            }
        case 9223372036854775807:
            var v22 uint64 = 42
            _ = v22
        default:
            var v23 uint = 65535
            _ = v23
        }
        _ = v14
    } else {
        switch "" {
        case "\n\t":
            if true {
                for fi24 := 0; fi24 < 16; fi24++ {
                    if (!((((("a" + "\n\t") + "") + ("hello" + "hello")) > "hello") && ("\n\t" >= "a"))) {
                        v25 := GenS1{F0: 9223372036854775807}
                        _ = v25
                    } else {
                        switch 32768 {
                        case 1:
                            var v26 uint64 = v1
                            _ = v26
                        case 9223372036854775807:
                            if false {
                                v27 := GenS2{F0: -1, F1: 52}
                                if (1 <= -128) {
                                    var v28 uint8 = 55
                                    _ = v28
                                }
                                if (!true) {
                                    var v29 uint32 = 670
                                    if true {
                                        v13 = append(v13, (fi24 + fi24))
                                        var v30 int8 = 100
                                        _ = v30
                                    }
                                    if len(v13) > 0 {
                                        v13[0] = (fi24 / ((4294967295) | 1))
                                    }
                                    _ = v29
                                }
                                _ = v27
                            }
                        default:
                            if false {
                                v31 := v3[2147483648]
                                v32 := GenS3{F0: 32767, F1: 9223372036854775807}
                                var v33 int64 = -117
                                _ = v31
                                _ = v32
                                _ = v33
                            }
                        }
                        v5 = v5
                    }
                    v34 := GenS3{F0: 0, F1: 0}
                    _ = fi24
                    _ = v34
                }
                v35 := GenS4{F0: "a"}
                _ = v35
            }
        case "hello":
            for fi36 := 0; fi36 < 18; fi36++ {
                switch ("foo bar" + (("hello" + "foo bar") + "foo bar")) {
                case "\n\t":
                    v2 += (v2 + (3.14159 - (0.1 * 0.0)))
                }
                _ = fi36
            }
        case "foo bar":
            if true {
                if false {
                    var v37 uint8 = 128
                    _ = v37
                } else {
                    v38 := GenS2{F0: -32768, F1: 204}
                    var v39 uint64 = v1
                    _ = v38
                    _ = v39
                }
                var v40 int64 = 1
                var v41 int64 = 42
                _ = v40
                _ = v41
            }
        default:
            v42 := GenS4{F0: ""}
            _ = v42
        }
        v5 = (3.14159 - v5)
        var v43 float64 = v2
        _ = v43
    }
    v44 := GenS5{F0: false}
    if len(v13) > 0 {
        v13[0] = 0
    }
    v3[65535] = -83
    for fi45 := 0; fi45 < 8; fi45++ {
        var v46 bool = true
        for fi47 := 0; fi47 < 5; fi47++ {
            if false {
                delete(v3, 2147483648)
                if ((v1 / ((v1) | 1)) == v1) {
                    fmt.Println(v2, v1)
                }
            }
            _ = fi47
        }
        v13 = append(v13, fi45)
        _ = fi45
        _ = v46
    }
    v1 = 2
    v2 += 3.14159
    var v48 uint32 = 2147483648
    switch int(v1) {
    case 280:
        v2 -= v2
    case -1:
        v49 := GenS1{F0: -9223372036854775808}
        _ = v49
    default:
        fmt.Println(v1)
    }
    var v50 int32 = 958
    v5 -= v5
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v48
    _ = v5
    _ = v50
}
