package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 string; F1 int32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint; F1 uint64 }
type GenS6 struct { F0 uint64; F1 int; F2 int64 }

func main() {
    var v1 uint8 = 127
    v2 := make(map[int]int)
    var v3 uint = 65535
    switch -1 {
    case -2147483648:
        fmt.Println(v1, v1)
    case -9223372036854775808:
        if (!false) {
            v1 = v1
        }
    case 831:
        if false {
            fmt.Println(v1, v1)
            var v4 uint16 = 2
            delete(v2, 128)
            _ = v4
        }
    }
    var v5 int64 = 0
    for fi6 := 0; fi6 < 11; fi6++ {
        v7 := []int{(fi6 * 1), fi6, fi6}
        _ = fi6
        _ = v7
    }
    switch 9223372036854775807 {
    case 2147483647:
        var v8 float32 = (((0.1 - 1.5) + -2.25) - (0.0 / float32(-2.25)))
        _ = v8
    default:
        if (true && ((!true) && true)) {
            fmt.Println(v3)
            v5 -= v5
        } else {
            if (v1 != 42) {
                var v9 int64 = v5
                _ = v9
            } else {
                if (1 >= v5) {
                    v5 = (v5 | -258)
                    fmt.Println(v3, v3, v1)
                    for fi10 := 0; fi10 < 1; fi10++ {
                        v11 := GenS1{F0: (v1 << (((v3 & 42)) & 63))}
                        _ = fi10
                        _ = v11
                    }
                }
                switch "a" {
                case "\n\t":
                    v12 := []int{}
                    _ = v12
                case "foo bar":
                    if true {
                        if (!true) {
                            var v13 float32 = 1e10
                            v14 := GenS1{F0: v1}
                            _ = v13
                            _ = v14
                        } else {
                            var v15 int = 355
                            for fi16 := 0; fi16 < 1; fi16++ {
                                v15 -= fi16
                                var v17 uint8 = 255
                                var v18 uint64 = uint64(v17)
                                _ = fi16
                                _ = v17
                                _ = v18
                            }
                            _ = v15
                        }
                    }
                case "hello":
                    for fi19 := 0; fi19 < 3; fi19++ {
                        v5 -= v5
                        fmt.Println(v5)
                        v3 = 42
                        _ = fi19
                    }
                }
                var v20 int8 = 42
                _ = v20
            }
            for fi21 := 0; fi21 < 19; fi21++ {
                v1 = v1
                _ = fi21
            }
            for fi22 := 0; fi22 < 5; fi22++ {
                switch (fi22 & (fi22 * int(v5))) {
                case 561:
                    fmt.Println(v1)
                default:
                    var v23 uint32 = 969
                    _ = v23
                }
                _ = fi22
            }
        }
    }
    v2[32768] = -2147483648
    v3 = 4294967295
    v24 := GenS2{F0: ("hello" + "\n\t"), F1: -541}
    for fi25 := 0; fi25 < 20; fi25++ {
        if (2147483647 < 1) {
            if (!false) {
                if ("a" <= ("" + "hello")) {
                    if (-128 <= -48) {
                        fmt.Println(v5)
                        v26 := GenS2{F0: "a", F1: -1}
                        _ = v26
                    } else {
                        v5 -= (v5 * v5)
                        v27 := GenS2{F0: (("hello" + ((("hello" + (("foo bar" + (("" + "foo bar") + (("hello" + ("\n\t" + "")) + ""))) + "\n\t")) + ("\n\t" + "hello")) + "")) + ("\n\t" + "a")), F1: -2147483648}
                        v1 = (v1 >> ((65535) & 63))
                        _ = v27
                    }
                    for fi28 := 0; fi28 < 11; fi28++ {
                        var v29 int16 = -1
                        v30 := GenS3{F0: (-(v5))}
                        if false {
                            var v31 int32 = -2147483648
                            var v32 uint32 = 2
                            var v33 uint = 255
                            v34 := []int{9223372036854775807, (fi25 - (fi25 >> ((v3) & 63)))}
                            v34 = append(v34, (fi25 | (fi28 / (((fi25 | (^(fi28)))) | 1))))
                            _ = v31
                            _ = v32
                            _ = v33
                            _ = v34
                        }
                        _ = fi28
                        _ = v29
                        _ = v30
                    }
                    for fi35 := 0; fi35 < 8; fi35++ {
                        var v36 int32 = 2147483647
                        var v37 float64 = float64(-2.25)
                        var v38 uint64 = 42
                        _ = fi35
                        _ = v36
                        _ = v37
                        _ = v38
                    }
                } else {
                    switch "\n\t" {
                    case "a":
                        v39 := GenS4{F0: v1}
                        _ = v39
                    case "foo bar":
                        v40 := GenS1{F0: v1}
                        _ = v40
                    }
                    v3 -= 255
                    v5 -= v5
                }
                var v41 int32 = 42
                _ = v41
            }
            if false {
                v3 -= 257
                v3 += (v3 - v3)
                switch fi25 {
                case -1:
                    if ((v1 << (((v3 % ((4294967295) | 1))) & 63)) < v1) {
                        v42 := GenS5{F0: 2, F1: 18446744073709551615}
                        if false {
                            fmt.Println(fi25)
                        }
                        v43 := []int{(fi25 << ((v3) & 63)), fi25}
                        _ = v42
                        _ = v43
                    }
                default:
                    if (false && true) {
                        v44 := GenS1{F0: v1}
                        v5 -= -1
                        if ((0.0 / float32(((float64(1.5) * 3.14159) + 0.0))) <= 1.5) {
                            v45 := []int{fi25, fi25, (fi25 ^ -2147483648), fi25}
                            _ = v45
                        } else {
                            switch "" {
                            case "":
                                v46 := []int{}
                                _ = v46
                            }
                            v3 -= uint(v1)
                        }
                        _ = v44
                    }
                }
            }
            for fi47 := 0; fi47 < 3; fi47++ {
                v5 = -1
                _ = fi47
            }
        } else {
            fmt.Println(v3, v3)
        }
        v1 += (v1 / ((v1) | 1))
        switch ((("" + ("hello" + "a")) + "\n\t") + "foo bar") {
        case "a":
            fmt.Println(v3)
        case "\n\t":
            v48 := []int{-583, (fi25 | fi25)}
            _ = v48
        case "hello":
            v49 := GenS6{F0: 2, F1: (fi25 - (fi25 / ((fi25) | 1))), F2: -9223372036854775808}
            _ = v49
        default:
            var v50 int16 = 0
            _ = v50
        }
        _ = fi25
    }
    v51 := []int{834, 127, 32768, 306}
    delete(v2, -143)
    v5 = 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v51:", len(v51), v51)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v5
    _ = v51
}
