package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int; F2 uint32 }
type GenS2 struct { F0 string; F1 int32; F2 int64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint; F1 uint64; F2 bool }
type GenS5 struct { F0 uint8; F1 int32; F2 int8 }
func genFn0(p0 int64, p1 int32) uint8 {
    p0 = (p0 * (-(p0)))
    v7 := GenS1{F0: 42, F1: -2147483648, F2: 279}
    switch "a" {
    case "a":
        var v8 uint8 = 1
        _ = v8
    case "":
        v9 := GenS1{F0: 2, F1: 0, F2: 262}
        _ = v9
    case "hello":
        var v10 int16 = 1000
        _ = v10
    }
    switch ("\n\t" + "\n\t") {
    case "foo bar":
        v11 := GenS2{F0: (("hello" + "foo bar") + "hello"), F1: p1, F2: 0}
        _ = v11
    case "\n\t":
        switch -566 {
        case 1:
            v12 := []int{197, -9223372036854775808, 32767, -514}
            _ = v12
        default:
            for fi13 := 0; fi13 < 17; fi13++ {
                fmt.Println(p0, p0, p1)
                if true {
                    var v14 uint8 = 1
                    _ = v14
                }
                switch fi13 {
                case -763:
                    fmt.Println(p0, p1, p0)
                case -2147483648:
                    for fi15 := 0; fi15 < 14; fi15++ {
                        if true {
                            switch "hello" {
                            case "":
                                var v16 uint32 = 2147483647
                                _ = v16
                            case "foo bar":
                                switch 4294967295 {
                                case 128:
                                    var v17 int8 = 1
                                    _ = v17
                                case 32768:
                                    if ((0.0 * (0.0 / 0.1)) != (float64((3.14159 - (((1e10 - -2.25) * ((-2.25 - 0.0) + (3.14159 + 1.5))) + -2.25))) * 0.1)) {
                                        if ((2 == 1) && ("\n\t" > "a")) {
                                            if false {
                                                fmt.Println(fi15, p0)
                                            } else {
                                                var v18 float64 = -2.25
                                                var v19 bool = ((-(p0)) >= 0)
                                                if (float32(1.5) == float32(v18)) {
                                                    v20 := GenS3{F0: (p1 - (p1 >> ((65535) & 63)))}
                                                    v21 := GenS3{F0: (p1 | p1)}
                                                    _ = v20
                                                    _ = v21
                                                } else {
                                                    switch fi15 {
                                                    case 42:
                                                        v22 := GenS2{F0: "a", F1: p1, F2: 0}
                                                        _ = v22
                                                    }
                                                }
                                                _ = v18
                                                _ = v19
                                            }
                                            switch (fi15 << ((255) & 63)) {
                                            case -2147483648:
                                                p1 = (p1 % ((-2147483648) | 1))
                                            case 42:
                                                if true {
                                                    var v23 uint8 = 128
                                                    v24 := []int{}
                                                    _ = v23
                                                    _ = v24
                                                }
                                            case -1:
                                                if true {
                                                    p0 += p0
                                                    v25 := []int{}
                                                    var v26 int64 = 9223372036854775807
                                                    _ = v25
                                                    _ = v26
                                                }
                                            }
                                        }
                                    } else {
                                        v27 := []int{}
                                        p0 += (^(p0))
                                        _ = v27
                                    }
                                }
                            }
                            p0 -= p0
                            switch (^(fi13)) {
                            case 937:
                                v28 := []int{-9223372036854775808, (fi15 / ((fi15) | 1)), (fi13 - fi13)}
                                _ = v28
                            case 209:
                                var v29 uint32 = 324
                                _ = v29
                            default:
                                fmt.Println(fi15, p0, fi15)
                            }
                        } else {
                            v30 := GenS4{F0: 65535, F1: 2, F2: ("\n\t" > "foo bar")}
                            if (true && ((1 == 788) && (-128 != 42))) {
                                switch fi13 {
                                case -9223372036854775808:
                                    var v31 uint16 = 0
                                    _ = v31
                                case 4294967295:
                                    switch ("a" + ("hello" + "")) {
                                    case "\n\t":
                                        var v32 int8 = 100
                                        _ = v32
                                    case "a":
                                        switch (fi15 & (fi13 << ((18446744073709551615) & 63))) {
                                        case -454:
                                            switch fi13 {
                                            case 9223372036854775807:
                                                var v33 uint8 = 127
                                                _ = v33
                                            }
                                        }
                                    default:
                                        v34 := []int{(fi15 << ((255) & 63)), (fi13 | fi15)}
                                        _ = v34
                                    }
                                case 42:
                                    p1 = -618
                                }
                                fmt.Println(fi13, fi13)
                                var v35 string = ("" + "a")
                                _ = v35
                            }
                            v36 := GenS1{F0: 65535, F1: (fi15 ^ 763), F2: 4294967295}
                            _ = v30
                            _ = v36
                        }
                        var v37 int16 = 1
                        v38 := []int{}
                        _ = fi15
                        _ = v37
                        _ = v38
                    }
                case 794:
                    v39 := GenS1{F0: 2, F1: fi13, F2: 0}
                    _ = v39
                }
                _ = fi13
            }
        }
    default:
        v40 := []int{32768, 4294967295, -9223372036854775808, -42}
        _ = v40
    }
    _ = p0
    _ = p1
    _ = v7
    return 1
}

func main() {
    var v1 uint = 18446744073709551615
    v2 := make(map[uint32]uint32)
    v3 := make(map[int]bool)
    v4 := make(map[int]uint32)
    v5 := make(map[uint32]bool)
    var v6 int16 = 0
    switch 9223372036854775807 {
    case -9223372036854775808:
        var v41 string = ("" + ("" + "a"))
        _ = v41
    }
    for fi42 := 0; fi42 < 7; fi42++ {
        v1 = v1
        _ = fi42
    }
    v6 += (v6 % ((v6) | 1))
    var v43 float32 = 1e10
    delete(v2, 2147483647)
    for fi44 := 0; fi44 < 20; fi44++ {
        v6 -= 764
        var v45 float32 = (0.1 + (1.5 * v43))
        switch fi44 {
        case -1:
            v3[fi44] = ((false && (193 == 4294967295)) || false)
        case 127:
            var v46 uint32 = 2
            _ = v46
        case 2147483647:
            var v47 int = fi44
            _ = v47
        }
        _ = fi44
        _ = v45
    }
    v6 = v6
    v48 := GenS5{F0: 128, F1: 2147483647, F2: 127}
    v6 = (v6 ^ 42)
    var v49 uint = (v1 % ((v1) | 1))
    _ = len(v3)
    switch "" {
    case "hello":
        v50 := v2[2147483647]
        _ = v50
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v48
    _ = v49
    _ = v5
    _ = v6
}
