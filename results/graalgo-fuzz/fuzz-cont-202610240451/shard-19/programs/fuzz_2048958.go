package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int; F1 int64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int; F1 bool }
func genFn0(p0 int32, p1 uint8, p2 int32) uint8 {
    if false {
        for fi6 := 0; fi6 < 8; fi6++ {
            p2 = 1
            p1 += (p1 - p1)
            p0 += p2
            _ = fi6
        }
    }
    switch "foo bar" {
    case "":
        var v7 int = 127
        _ = v7
    default:
        p2 = -359
    }
    v8 := []int{2147483647, -639, 1, -2147483648}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 1
}

func main() {
    var v1 int16 = -266
    v2 := make(map[int]uint32)
    var v3 uint32 = 0
    var v4 uint = 2
    var v5 uint = 0
    var v9 uint8 = 1
    v10 := GenS1{F0: (v3 % (((v3 / ((146) | 1))) | 1)), F1: true}
    if true {
        if (!((((3.14159 + -2.25) - -2.25) > 3.14159) && true)) {
            v5 = v5
            var v11 int = -2147483648
            _ = v11
        }
        for fi12 := 0; fi12 < 18; fi12++ {
            if ((!true) || (604 == 18446744073709551615)) {
                var v13 int32 = -2147483648
                switch "foo bar" {
                case "foo bar":
                    fmt.Println(v1)
                case "hello":
                    switch "hello" {
                    case "a":
                        if false {
                            switch fi12 {
                            case -185:
                                v1 -= -392
                            }
                            v13 += v13
                            var v14 uint16 = 42
                            _ = v14
                        } else {
                            v15 := GenS2{F0: 65535}
                            _ = v15
                        }
                    }
                case "a":
                    var v16 uint32 = v3
                    _ = v16
                }
                switch (fi12 / (((fi12 << ((18446744073709551615) & 63))) | 1)) {
                case -9223372036854775808:
                    v4 = (v5 & 592)
                case 645:
                    v5 -= 42
                case 543:
                    v17 := GenS1{F0: v3, F1: false}
                    _ = v17
                default:
                    var v18 int64 = 0
                    _ = v18
                }
                _ = v13
            } else {
                for fi19 := 0; fi19 < 10; fi19++ {
                    var v20 int64 = 0
                    v21 := GenS3{F0: 127, F1: v20}
                    v22 := GenS2{F0: 65535}
                    _ = fi19
                    _ = v20
                    _ = v21
                    _ = v22
                }
                fmt.Println(v1)
                var v23 uint = (v5 ^ v5)
                _ = v23
            }
            v24 := v2[fi12]
            if true {
                if false {
                    v4 -= (v4 >> (((v4 & 42)) & 63))
                }
            }
            _ = fi12
            _ = v24
        }
        var v25 string = "\n\t"
        _ = v25
    }
    v26 := []int{-149, 2147483647, 42}
    v27 := GenS2{F0: 255}
    var v28 float64 = 0.1
    v29 := v2[0]
    _ = len(v2)
    v30 := GenS1{F0: v3, F1: (-9223372036854775808 == 402)}
    for fi31 := 0; fi31 < 6; fi31++ {
        v32 := GenS2{F0: uint16(v4)}
        _ = fi31
        _ = v32
    }
    for fi33 := 0; fi33 < 10; fi33++ {
        v29 = (v3 + v3)
        switch "" {
        case "foo bar":
            if true {
                fmt.Println(fi33, fi33, v3)
                v34 := GenS4{F0: -9223372036854775808}
                v35 := GenS1{F0: v3, F1: (false && false)}
                _ = v34
                _ = v35
            }
        case "":
            for fi36 := 0; fi36 < 1; fi36++ {
                var v37 bool = true
                if v37 {
                    if len(v26) > 0 {
                        v26[0] = 1
                    }
                }
                _ = fi36
                _ = v37
            }
        case "\n\t":
            if true {
                var v38 string = ("hello" + (("a" + "foo bar") + ("hello" + "hello")))
                _ = v38
            } else {
                for fi39 := 0; fi39 < 4; fi39++ {
                    if true {
                        switch ((("hello" + "\n\t") + (("" + "a") + "a")) + ("foo bar" + "")) {
                        case "hello":
                            if (2147483647 <= (fi39 % ((-9223372036854775808) | 1))) {
                                if true {
                                    fmt.Println(v28, v4, fi33)
                                    var v40 int32 = 2147483647
                                    switch "\n\t" {
                                    case "\n\t":
                                        fmt.Println(v28, v28)
                                    case "a":
                                        var v41 float64 = ((0.0 - float64(float32(v28))) * 1e10)
                                        _ = v41
                                    case "hello":
                                        v42 := GenS1{F0: v3, F1: false}
                                        _ = v42
                                    }
                                    _ = v40
                                } else {
                                    v2[2147483647] = (v3 - (v29 / ((2) | 1)))
                                    var v43 uint8 = (v9 / (((v9 + 135)) | 1))
                                    v26 = append(v26, (fi33 | fi39))
                                    _ = v43
                                }
                            }
                        }
                        v9 = (v9 + 232)
                        v5 += v5
                    } else {
                        if len(v26) >= 2 {
                            v26 = v26[1:len(v26)]
                        }
                    }
                    _ = fi39
                }
            }
        }
        v3 -= (v3 + (v29 >> ((v4) & 63)))
        _ = fi33
    }
    if (2 < 18446744073709551615) {
        v26 = append(v26, 0)
    } else {
        for fi44 := 0; fi44 < 3; fi44++ {
            v45 := GenS5{F0: 1}
            var v46 float32 = -2.25
            switch "\n\t" {
            case "":
                var v47 uint16 = 0
                _ = v47
            case "a":
                switch "" {
                case "a":
                    v1 -= v1
                case "":
                    var v48 int64 = 742
                    _ = v48
                }
            default:
                switch "a" {
                case "a":
                    v49 := GenS6{F0: 0, F1: (9223372036854775807 > 0)}
                    _ = v49
                case "\n\t":
                    for fi50 := 0; fi50 < 3; fi50++ {
                        var v51 bool = (18446744073709551615 < (v4 << (((v4 * (v4 % (((v5 >> ((v5) & 63))) | 1)))) & 63)))
                        v52 := GenS6{F0: fi44, F1: (true && v51)}
                        _ = fi50
                        _ = v51
                        _ = v52
                    }
                case "":
                    v5 += (v5 | 255)
                default:
                    v28 += v28
                }
            }
            _ = fi44
            _ = v45
            _ = v46
        }
        v4 -= v4
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v9
}
