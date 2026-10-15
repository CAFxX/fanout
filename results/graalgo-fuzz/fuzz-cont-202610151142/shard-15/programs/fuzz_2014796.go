package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16; F1 uint8; F2 uint64 }
type GenS4 struct { F0 int16; F1 int64 }
type GenS5 struct { F0 int; F1 uint8 }
type GenS6 struct { F0 uint16; F1 int32; F2 int32 }
type GenS7 struct { F0 int32 }
func genFn0() uint8 {
    var v6 uint8 = 255
    var v7 uint32 = 42
    v7 += v7
    _ = v6
    _ = v7
    return 42
}
func genFn1() uint32 {
    var v8 uint = 18446744073709551615
    if (int16(v8) == -491) {
        var v9 uint16 = 2
        for fi10 := 0; fi10 < 16; fi10++ {
            if (fi10 <= 2147483647) {
                v11 := []int{fi10, (fi10 << (((^(v8))) & 63))}
                v8 = v8
                v8 = (v8 ^ (^(v8)))
                _ = v11
            }
            _ = fi10
        }
        var v12 float64 = 3.14159
        _ = v12
        _ = v9
    } else {
        var v13 int32 = 1
        fmt.Println(v13, v13)
        v14 := GenS1{F0: 42, F1: 1}
        _ = v13
        _ = v14
    }
    switch 42 {
    case 127:
        var v15 uint16 = 42
        _ = v15
    case 305:
        if (true || false) {
            v8 = 1
        }
    default:
        v8 -= (^(v8))
    }
    _ = v8
    return 42
}

func main() {
    var v1 uint8 = 0
    v2 := []string{}
    var v3 uint32 = 2
    v4 := []string{"a", "a", ("\n\t" + "foo bar")}
    v5 := []int{}
    for fi16 := 0; fi16 < 11; fi16++ {
        switch ("\n\t" + "") {
        case "a":
            for fi17 := 0; fi17 < 9; fi17++ {
                var v18 int = (fi17 & fi17)
                _ = fi17
                _ = v18
            }
        case "hello":
            var v19 int8 = 1
            _ = v19
        }
        if (2 < 870) {
            var v20 bool = false
            if v20 {
                v21 := GenS2{F0: 0}
                _ = v21
            } else {
                v22 := GenS2{F0: 342}
                _ = v22
            }
            _ = v20
        } else {
            v2 = append(v2, ("" + ("hello" + ("" + ((("" + "") + "foo bar") + "hello")))))
            var v23 float32 = 1.5
            v24 := GenS3{F0: 423, F1: (v1 % ((2) | 1)), F2: 18446744073709551615}
            _ = v23
            _ = v24
        }
        v1 = 0
        _ = fi16
    }
    var v25 int64 = 42
    v1 = (v1 * genFn0())
    var v26 int16 = 0
    var v27 uint = 850
    for fi28 := 0; fi28 < 17; fi28++ {
        if true {
            fmt.Println(v25)
            if len(v5) > 0 {
                v5[0] = fi28
            }
            v29 := GenS1{F0: 42, F1: 523}
            _ = v29
        } else {
            v30 := GenS4{F0: int16(fi28), F1: v25}
            var v31 uint32 = v3
            _ = v30
            _ = v31
        }
        _ = fi28
    }
    _ = len(v2)
    if ((v26 & 32767) != (v26 % ((v26) | 1))) {
        v27 = (v27 + v27)
    } else {
        if true {
            for fi32 := 0; fi32 < 10; fi32++ {
                for fi33 := 0; fi33 < 13; fi33++ {
                    v34 := GenS3{F0: 2, F1: 46, F2: 4294967295}
                    switch (fi33 - 127) {
                    case 128:
                        fmt.Println(fi33, v1)
                    case 278:
                        v4 = append(v4, ((((("foo bar" + "\n\t") + "") + "foo bar") + ("\n\t" + "a")) + ("" + "\n\t")))
                    case 2147483647:
                        if true {
                            fmt.Println(v27, v1, fi33)
                            if (42 != 215) {
                                v35 := GenS5{F0: 127, F1: (v1 + v1)}
                                if (58 == 97) {
                                    v25 = int64(v27)
                                    if (18446744073709551615 <= 390) {
                                        v26 = v26
                                        fmt.Println(fi32, v25, v1)
                                    } else {
                                        if true {
                                            switch "a" {
                                            case "hello":
                                                v26 += v26
                                            case "":
                                                var v36 uint64 = 0
                                                _ = v36
                                            case "\n\t":
                                                fmt.Println(v1, v1)
                                            }
                                        } else {
                                            fmt.Println(v27, v26)
                                            v27 = (v27 & (v27 * (v27 / ((444) | 1))))
                                            _ = len(v2)
                                        }
                                        if false {
                                            v3 = uint32(v1)
                                            v5 = append(v5, -9223372036854775808)
                                            var v37 int8 = 100
                                            _ = v37
                                        } else {
                                            if false {
                                                var v38 int64 = v25
                                                v39 := GenS6{F0: uint16(v27), F1: -208, F2: -2147483648}
                                                _ = v38
                                                _ = v39
                                            } else {
                                                var v40 int64 = (^(v25))
                                                _ = len(v2)
                                                switch fi33 {
                                                case 106:
                                                    fmt.Println(v1, fi33, v1)
                                                    if true {
                                                        v25 = v40
                                                    }
                                                case 408:
                                                    v25 += (v25 ^ 0)
                                                }
                                                _ = v40
                                            }
                                            if false {
                                                var v41 float32 = float32(1e10)
                                                fmt.Println(v27, v3)
                                                _ = v41
                                            }
                                            switch ("a" + (("a" + "hello") + ("a" + ("\n\t" + ((("foo bar" + "foo bar") + ("a" + "")) + "a"))))) {
                                            case "foo bar":
                                                v42 := GenS5{F0: fi33, F1: 142}
                                                _ = v42
                                            case "hello":
                                                fmt.Println(fi32, v26, fi32)
                                            default:
                                                v43 := GenS7{F0: 1}
                                                _ = v43
                                            }
                                        }
                                    }
                                    v26 += (-(v26))
                                }
                                v27 += v27
                                _ = v35
                            } else {
                                v27 = v27
                                fmt.Println(v26)
                            }
                        } else {
                            switch "foo bar" {
                            case "foo bar":
                                var v44 uint64 = 1
                                _ = v44
                            }
                        }
                    default:
                        switch ("a" + "a") {
                        case "a":
                            if (!false) {
                                v26 += (v26 << (((^(v27))) & 63))
                                fmt.Println(v25, v26)
                                switch "\n\t" {
                                case "hello":
                                    v27 -= v27
                                case "\n\t":
                                    var v45 int32 = 1
                                    _ = v45
                                }
                            } else {
                                if (1 <= 32767) {
                                    var v46 uint = v27
                                    _ = v46
                                }
                            }
                        case "":
                            v25 += 9223372036854775807
                        case "foo bar":
                            fmt.Println(v27)
                        }
                    }
                    _ = fi33
                    _ = v34
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                _ = fi32
            }
        } else {
            var v47 int32 = 47
            _ = v47
        }
    }
    for fi48 := 0; fi48 < 1; fi48++ {
        var v49 uint8 = v1
        _ = fi48
        _ = v49
    }
    var v50 int8 = 59
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v51 := GenS6{F0: 32767, F1: 42, F2: int32(v1)}
    var v52 string = "a"
    v53 := GenS3{F0: 2, F1: 171, F2: 807}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v50:", v50)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v53
}
