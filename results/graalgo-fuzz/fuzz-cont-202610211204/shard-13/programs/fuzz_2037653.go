package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 bool; F1 int32; F2 int64 }
type GenS7 struct { F0 uint64 }
func genFn0(p0 uint, p1 bool, p2 bool) int16 {
    p0 = 1
    p1 = ((-1 < 1000) || p2)
    fmt.Println(p1, p1, p2)
    p2 = (42 >= -128)
    var v7 int8 = 118
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 0
}

func main() {
    var v1 bool = ("\n\t" == "a")
    var v2 int16 = 42
    v3 := []string{("foo bar" + "\n\t")}
    var v4 int = 32768
    var v5 uint16 = 2
    v6 := []int{9223372036854775807, 9223372036854775807, (v4 ^ v4), (-(v4))}
    v8 := GenS1{F0: 42}
    _ = len(v6)
    var v9 int16 = 1000
    var v10 uint16 = (^(v5))
    if ((v5 << ((255) & 63)) > v10) {
        if len(v6) > 0 {
            v6[0] = v4
        }
        if v1 {
            if (int8(v10) != 0) {
                for fi11 := 0; fi11 < 11; fi11++ {
                    var v12 int16 = v9
                    var v13 uint16 = (v5 / ((v10) | 1))
                    for fi14 := 0; fi14 < 14; fi14++ {
                        v12 -= v9
                        switch ("hello" + "\n\t") {
                        case "hello":
                            var v15 int32 = -1
                            _ = v15
                        case "foo bar":
                            var v16 string = "a"
                            var v17 uint64 = 2
                            if v1 {
                                var v18 int64 = 314
                                v19 := GenS1{F0: 2147483647}
                                switch v16 {
                                case "hello":
                                    var v20 int32 = -2147483648
                                    _ = v20
                                case "foo bar":
                                    var v21 float64 = float64(0.0)
                                    _ = v21
                                case "\n\t":
                                    v22 := GenS2{F0: int(v18)}
                                    _ = v22
                                }
                                _ = v18
                                _ = v19
                            }
                            _ = v16
                            _ = v17
                        default:
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                        }
                        if (0 > -595) {
                            switch "\n\t" {
                            case "":
                                if len(v3) > 0 {
                                    v3[0] = ("a" + "a")
                                }
                            case "\n\t":
                                if ((v1 || v1) && (false && v1)) {
                                    var v23 uint = 631
                                    _ = v23
                                }
                            }
                            v24 := GenS3{F0: 77}
                            switch -9223372036854775808 {
                            case 0:
                                if (555 == 2147483648) {
                                    v25 := GenS4{F0: (v5 << ((65535) & 63))}
                                    _ = v25
                                }
                                v26 := GenS4{F0: v13}
                                v27 := GenS4{F0: (v5 | 574)}
                                _ = v26
                                _ = v27
                            default:
                                if len(v6) > 0 {
                                    v6[0] = (fi14 >> ((1) & 63))
                                }
                            }
                            _ = v24
                        }
                        _ = fi14
                    }
                    _ = fi11
                    _ = v12
                    _ = v13
                }
                v5 = v10
                var v28 uint8 = 0
                _ = v28
            } else {
                var v29 string = "\n\t"
                _ = v29
            }
            switch -9223372036854775808 {
            case -9223372036854775808:
                v30 := GenS3{F0: 124}
                _ = v30
            case -2147483648:
                var v31 int8 = 1
                _ = v31
            }
        }
    } else {
        v5 += (v10 & v5)
        for fi32 := 0; fi32 < 10; fi32++ {
            for fi33 := 0; fi33 < 9; fi33++ {
                v5 = v10
                v6 = append(v6, 9223372036854775807)
                _ = fi33
            }
            if len(v6) > 0 {
                v6[0] = v4
            }
            _ = fi32
        }
        v34 := GenS5{F0: genFn0(65535, v1, (!v1))}
        _ = v34
    }
    for fi35 := 0; fi35 < 16; fi35++ {
        v2 = v2
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi35
    }
    v2 = genFn0(4294967295, v1, v1)
    for fi36 := 0; fi36 < 10; fi36++ {
        fmt.Println(v4, v2, v9)
        var v37 float32 = 1e10
        _ = fi36
        _ = v37
    }
    fmt.Println(v1)
    switch (v4 * v4) {
    case 42:
        var v38 int32 = -1
        _ = v38
    }
    for fi39 := 0; fi39 < 13; fi39++ {
        var v40 int32 = 582
        v41 := GenS6{F0: v1, F1: (v40 % ((v40) | 1)), F2: -146}
        _ = fi39
        _ = v40
        _ = v41
    }
    var v42 int16 = 1000
    var v43 uint64 = 18446744073709551615
    switch "" {
    case "a":
        if len(v3) > 0 {
            v3[0] = "hello"
        }
    default:
        v6 = append(v6, 42)
    }
    for fi44 := 0; fi44 < 2; fi44++ {
        for fi45 := 0; fi45 < 20; fi45++ {
            var v46 int32 = -1
            v47 := GenS2{F0: (^(v4))}
            v48 := GenS7{F0: (v43 * 42)}
            _ = fi45
            _ = v46
            _ = v47
            _ = v48
        }
        v1 = v1
        if len(v6) > 0 {
            v6[0] = 4294967295
        }
        _ = fi44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
