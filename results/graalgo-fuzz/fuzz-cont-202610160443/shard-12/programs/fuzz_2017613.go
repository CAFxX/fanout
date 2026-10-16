package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 bool }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32; F1 int8; F2 int }
type GenS4 struct { F0 uint16; F1 bool }
type GenS5 struct { F0 int; F1 bool }
func genFn0() bool {
    var v6 uint16 = 362
    fmt.Println(v6)
    for fi7 := 0; fi7 < 3; fi7++ {
        if false {
            if false {
                fmt.Println(fi7)
            }
        } else {
            for fi8 := 0; fi8 < 4; fi8++ {
                switch (fi7 - (^(fi7))) {
                case 772:
                    switch "\n\t" {
                    case "":
                        v6 = 32767
                        switch (fi8 ^ 787) {
                        case -173:
                            var v9 int8 = 1
                            _ = v9
                        case -2147483648:
                            v6 = 545
                        case 42:
                            if ((v6 ^ (^(v6))) <= 255) {
                                v6 = 940
                                var v10 uint64 = 2
                                _ = v10
                            } else {
                                if (!true) {
                                    var v11 uint8 = 2
                                    _ = v11
                                } else {
                                    v12 := GenS1{F0: true, F1: 1, F2: ((v6 << ((uint(v6)) & 63)) >= (v6 + v6))}
                                    _ = v12
                                }
                            }
                        default:
                            v6 = (^(v6))
                        }
                    }
                case -2147483648:
                    var v13 int8 = 100
                    _ = v13
                case 354:
                    v14 := []int{42, fi8, fi7, fi8}
                    fmt.Println(fi7)
                    v15 := GenS1{F0: true, F1: -202, F2: (v6 >= (v6 & (v6 + v6)))}
                    _ = v14
                    _ = v15
                default:
                    var v16 int8 = 42
                    _ = v16
                }
                v6 = (v6 ^ 255)
                _ = fi8
            }
            switch (-(fi7)) {
            case 0:
                v17 := GenS1{F0: (fi7 == fi7), F1: -9223372036854775808, F2: ((0.0 - 3.14159) >= -2.25)}
                _ = v17
            }
            v18 := GenS2{F0: 46}
            _ = v18
        }
        _ = fi7
    }
    var v19 int = 32767
    _ = v19
    _ = v6
    return (!false)
}
func genFn1(p0 uint32) uint16 {
    if false {
        switch ("a" + (((("\n\t" + "hello") + "a") + ((("hello" + "foo bar") + "\n\t") + "a")) + "hello")) {
        case "a":
            for fi20 := 0; fi20 < 2; fi20++ {
                p0 = 65535
                _ = fi20
            }
        case "":
            fmt.Println(p0)
        default:
            var v21 uint16 = 2
            _ = v21
        }
        v22 := GenS3{F0: 42, F1: 72, F2: -2147483648}
        for fi23 := 0; fi23 < 1; fi23++ {
            v24 := GenS2{F0: 2}
            _ = fi23
            _ = v24
        }
        _ = v22
    }
    for fi25 := 0; fi25 < 6; fi25++ {
        v26 := GenS1{F0: (!((!true) && true)), F1: -367, F2: (!(99 > 146))}
        _ = fi25
        _ = v26
    }
    v27 := GenS4{F0: 255, F1: false}
    _ = p0
    _ = v27
    return 65535
}

func main() {
    v1 := []int{190, 2147483647, 998}
    v2 := make(map[uint32]uint32)
    var v3 int = 9223372036854775807
    var v4 int = (v3 >> ((2) & 63))
    var v5 uint32 = 952
    var v28 string = ((("\n\t" + "hello") + "\n\t") + "a")
    v29 := GenS5{F0: (^(v4)), F1: true}
    switch ("hello" + v28) {
    case "hello":
        var v30 float32 = 1.5
        _ = v30
    case "\n\t":
        v31 := GenS2{F0: 4294967295}
        _ = v31
    default:
        for fi32 := 0; fi32 < 4; fi32++ {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            switch v28 {
            case "foo bar":
                var v33 uint32 = v5
                _ = v33
            case "hello":
                switch "hello" {
                case "a":
                    for fi34 := 0; fi34 < 3; fi34++ {
                        v4 = (v3 << ((2) & 63))
                        switch v4 {
                        case -2147483648:
                            v35 := GenS4{F0: 2, F1: false}
                            _ = v35
                        default:
                            v28 = (v28 + "foo bar")
                        }
                        _ = fi34
                    }
                case "\n\t":
                    v36 := GenS5{F0: v4, F1: false}
                    _ = v36
                }
            case "a":
                var v37 int16 = -1
                _ = v37
            default:
                var v38 int64 = 42
                _ = v38
            }
            _ = fi32
        }
    }
    v4 = v4
    for fi39 := 0; fi39 < 6; fi39++ {
        switch (v3 | -649) {
        case -2147483648:
            var v40 uint64 = 671
            _ = v40
        case -9223372036854775808:
            v5 -= 472
        }
        _ = fi39
    }
    v3 += v3
    var v41 uint16 = 65535
    var v42 float64 = ((0.1 + 1e10) / 3.14159)
    var v43 int8 = -78
    v4 += v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v5
}
