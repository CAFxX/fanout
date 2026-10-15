package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 uint32; F1 uint64; F2 string }
type GenS3 struct { F0 uint; F1 uint64 }
type GenS4 struct { F0 bool; F1 int8 }
type GenS5 struct { F0 int; F1 int }
func genFn0() int16 {
    var v6 uint64 = 4294967295
    var v7 uint = 18446744073709551615
    v8 := []int{-2147483648, 32767, -9223372036854775808, -2147483648}
    var v9 int16 = -935
    if ((-837 > -1) || true) {
        v10 := GenS1{F0: 88, F1: (v9 - (-(v9)))}
        var v11 uint8 = 2
        _ = v10
        _ = v11
    }
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return (v9 << (((v7 ^ (v7 % ((255) | 1)))) & 63))
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []string{"hello", (("a" + "a") + "")}
    var v3 uint16 = 32767
    var v4 float64 = 0.0
    v5 := []string{"", ("a" + "")}
    for fi12 := 0; fi12 < 5; fi12++ {
        if (((^(v3)) == v3) && (fi12 >= 126)) {
            if (42 != 65535) {
                if (-128 != 100) {
                    v3 += v3
                    v3 -= 951
                } else {
                    if len(v2) > 0 {
                        v2[0] = "\n\t"
                    }
                    v13 := GenS1{F0: 155, F1: -875}
                    var v14 uint16 = v3
                    _ = v13
                    _ = v14
                }
                for fi15 := 0; fi15 < 16; fi15++ {
                    fmt.Println(fi15, fi12, fi15)
                    _ = len(v1)
                    _ = fi15
                }
            }
        } else {
            if len(v2) > 0 {
                v2[0] = ""
            }
            var v16 bool = ((v3 | v3) != (v3 + v3))
            for fi17 := 0; fi17 < 14; fi17++ {
                v4 = 0.1
                v18 := GenS2{F0: 2147483648, F1: 452, F2: "a"}
                _ = fi17
                _ = v18
            }
            _ = v16
        }
        _ = fi12
    }
    for fi19 := 0; fi19 < 11; fi19++ {
        v20 := GenS1{F0: 0, F1: 42}
        _ = fi19
        _ = v20
    }
    var v21 string = "a"
    for fi22 := 0; fi22 < 20; fi22++ {
        switch ((v21 + ("" + (v21 + "a"))) + "foo bar") {
        case "":
            var v23 uint64 = 1
            _ = v23
        default:
            for fi24 := 0; fi24 < 2; fi24++ {
                if (75 > 49) {
                    v25 := v1[42]
                    var v26 float64 = ((0.1 * v4) * (v4 - v4))
                    _ = v25
                    _ = v26
                }
                v27 := GenS3{F0: 255, F1: 18446744073709551615}
                v4 = ((v4 * v4) * 1e10)
                _ = fi24
                _ = v27
            }
        }
        var v28 uint16 = v3
        _ = fi22
        _ = v28
    }
    switch v21 {
    case "a":
        switch ("foo bar" + "hello") {
        case "a":
            for fi29 := 0; fi29 < 19; fi29++ {
                v1[226] = (fi29 - fi29)
                switch fi29 {
                case -2147483648:
                    v30 := GenS1{F0: 0, F1: 327}
                    _ = v30
                }
                _ = fi29
            }
        case "\n\t":
            for fi31 := 0; fi31 < 19; fi31++ {
                _ = len(v5)
                v21 = v21
                _ = fi31
            }
        }
    }
    var v32 bool = (1 < 1)
    switch 1 {
    case 0:
        fmt.Println(v32, v4, v4)
    }
    if v32 {
        if (1 != 0) {
            switch (("foo bar" + v21) + v21) {
            case "hello":
                var v33 uint8 = 93
                _ = v33
            default:
                v3 -= (v3 >> ((1) & 63))
            }
            v32 = (0 <= -1)
            for fi34 := 0; fi34 < 13; fi34++ {
                if true {
                    if (1.5 >= v4) {
                        if false {
                            if (211 == 239) {
                                v32 = (9223372036854775807 <= -1)
                            }
                        } else {
                            v32 = true
                            var v35 bool = (1 == 1)
                            for fi36 := 0; fi36 < 16; fi36++ {
                                if v32 {
                                    v32 = (296 == 434)
                                    var v37 uint16 = 65535
                                    var v38 float32 = (float32(v4) / -2.25)
                                    if (883 < 65535) {
                                        v39 := GenS3{F0: 255, F1: 1}
                                        _ = v39
                                    }
                                    _ = v37
                                    _ = v38
                                }
                                v35 = true
                                _ = fi36
                            }
                            _ = v35
                        }
                        if (18446744073709551615 < 0) {
                            if true {
                                if v32 {
                                    switch fi34 {
                                    case 4294967295:
                                        if (((v21 + ("\n\t" + "")) <= v21) || v32) {
                                            var v40 uint64 = 61
                                            _ = v40
                                        } else {
                                            var v41 float64 = float64(1.5)
                                            v41 -= v41
                                            fmt.Println(v4)
                                            _ = v41
                                        }
                                    case -1:
                                        v21 = (v21 + v21)
                                    case 9223372036854775807:
                                        _ = len(v1)
                                    default:
                                        if false {
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                        }
                                    }
                                } else {
                                    v42 := GenS1{F0: 127, F1: -32768}
                                    v43 := v1[0]
                                    _ = v42
                                    _ = v43
                                }
                                var v44 int64 = -345
                                _ = v44
                            }
                        }
                    }
                    fmt.Println(v32, v21, v21)
                } else {
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    var v45 string = v21
                    _ = v45
                }
                _ = fi34
            }
        } else {
            if (!false) {
                fmt.Println(v21, v32, v32)
                if (("a" > v21) || v32) {
                    v46 := GenS3{F0: 1, F1: 173}
                    v47 := GenS1{F0: 127, F1: 1000}
                    _ = v46
                    _ = v47
                } else {
                    for fi48 := 0; fi48 < 5; fi48++ {
                        fmt.Println(v4, v3)
                        var v49 int32 = 973
                        var v50 int32 = (v49 & 2147483647)
                        _ = fi48
                        _ = v49
                        _ = v50
                    }
                }
            } else {
                v51 := GenS4{F0: (v32 || (uint(v3) == 255)), F1: 100}
                v52 := GenS4{F0: v32, F1: 51}
                _ = v51
                _ = v52
            }
        }
    } else {
        delete(v1, 65535)
    }
    for fi53 := 0; fi53 < 16; fi53++ {
        var v54 int64 = 0
        _ = fi53
        _ = v54
    }
    v55 := GenS5{F0: -282, F1: 9223372036854775807}
    v3 = 32767
    if false {
        if (!v32) {
            var v56 int16 = 1000
            _ = v56
        } else {
            var v57 int64 = -459
            _ = v57
        }
        v58 := GenS3{F0: 65535, F1: 952}
        _ = v58
    }
    fmt.Println(v21)
    if false {
        var v59 int8 = 127
        v60 := GenS5{F0: 0, F1: 127}
        v1[2147483648] = 2147483647
        _ = v59
        _ = v60
    }
    switch 0 {
    case -9223372036854775808:
        switch 42 {
        case 0:
            _ = len(v1)
        case -2147483648:
            _ = len(v1)
        case 4294967295:
            fmt.Println(v3)
        }
    default:
        delete(v1, 2147483647)
    }
    v3 = (v3 >> ((0) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v55
}
