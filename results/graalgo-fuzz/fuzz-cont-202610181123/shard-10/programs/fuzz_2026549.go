package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 int16 }
type GenS2 struct { F0 int64; F1 uint16; F2 int }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int16 }
func genFn0(p0 uint8, p1 int64) int {
    switch "hello" {
    case "":
        v4 := GenS1{F0: -128, F1: "foo bar", F2: 42}
        _ = v4
    case "a":
        p1 -= p1
    case "foo bar":
        fmt.Println(p1, p0)
    default:
        for fi5 := 0; fi5 < 5; fi5++ {
            if ((42 > 42) && ((p0 + (p0 << ((2) & 63))) > 115)) {
                v6 := GenS1{F0: -126, F1: "hello", F2: 32767}
                if true {
                    var v7 uint32 = 977
                    _ = v7
                }
                var v8 uint64 = 2
                _ = v6
                _ = v8
            }
            v9 := GenS1{F0: -28, F1: "foo bar", F2: 0}
            _ = fi5
            _ = v9
        }
    }
    p0 = (^(p0))
    v10 := GenS2{F0: (p1 - -1), F1: 252, F2: 0}
    _ = p0
    _ = p1
    _ = v10
    return 9223372036854775807
}

func main() {
    var v1 float64 = 0.1
    var v2 uint = 1
    var v3 uint16 = 32767
    fmt.Println(v2)
    for fi11 := 0; fi11 < 2; fi11++ {
        for fi12 := 0; fi12 < 15; fi12++ {
            var v13 int64 = -1
            var v14 int = (fi12 >> ((v2) & 63))
            v15 := []int{}
            _ = fi12
            _ = v13
            _ = v14
            _ = v15
        }
        v16 := GenS3{F0: -521}
        _ = fi11
        _ = v16
    }
    var v17 uint16 = 1
    var v18 int64 = 42
    v19 := GenS1{F0: 0, F1: "hello", F2: 878}
    switch (("a" + "a") + "foo bar") {
    case "foo bar":
        v20 := GenS2{F0: v18, F1: (^(v17)), F2: 149}
        _ = v20
    default:
        if false {
            v21 := GenS2{F0: (v18 >> ((v2) & 63)), F1: 952, F2: -1}
            v22 := []int{2147483647, int(v2)}
            _ = v21
            _ = v22
        } else {
            if false {
                var v23 int32 = -94
                for fi24 := 0; fi24 < 7; fi24++ {
                    var v25 uint = (v2 / ((v2) | 1))
                    var v26 uint = (v2 << ((v25) & 63))
                    _ = fi24
                    _ = v25
                    _ = v26
                }
                _ = v23
            } else {
                v27 := GenS4{F0: v2}
                _ = v27
            }
            if false {
                switch 32767 {
                case 128:
                    v17 = (v17 | (v3 + v3))
                }
                var v28 float32 = (0.1 - (1e10 - (1e10 + 0.0)))
                _ = v28
            }
            for fi29 := 0; fi29 < 2; fi29++ {
                var v30 int8 = 1
                v31 := []int{fi29, -1, (fi29 * genFn0(127, (v18 * (v18 * v18)))), genFn0(128, (v18 + int64(v30)))}
                for fi32 := 0; fi32 < 4; fi32++ {
                    var v33 uint8 = 255
                    switch (fi29 ^ (fi32 >> ((2) & 63))) {
                    case 4294967295:
                        if (1 == 42) {
                            v33 = v33
                        } else {
                            fmt.Println(v33, v3)
                            fmt.Println(v17, v18)
                        }
                    case 2147483647:
                        var v34 int16 = -32768
                        _ = v34
                    default:
                        if (227 <= 834) {
                            if false {
                                var v35 bool = (!false)
                                _ = len(v31)
                                _ = v35
                            } else {
                                v36 := GenS2{F0: v18, F1: v3, F2: fi32}
                                v31 = append(v31, 127)
                                v31 = append(v31, 370)
                                _ = v36
                            }
                            var v37 int8 = (^(v30))
                            var v38 int8 = -77
                            _ = v37
                            _ = v38
                        } else {
                            v18 -= (^(v18))
                            fmt.Println(fi32, fi29, v1)
                            if (!false) {
                                v3 = 65535
                                if (42 >= 1000) {
                                    v1 = 0.0
                                } else {
                                    switch "" {
                                    case "hello":
                                        var v39 uint64 = 2
                                        switch (fi32 & fi29) {
                                        case 42:
                                            var v40 uint16 = (v3 << ((v2) & 63))
                                            _ = v40
                                        case 9223372036854775807:
                                            v30 -= v30
                                        case 32767:
                                            v41 := GenS5{F0: 1000}
                                            _ = v41
                                        default:
                                            var v42 int64 = (-(v18))
                                            _ = v42
                                        }
                                        var v43 uint = uint(v39)
                                        _ = v39
                                        _ = v43
                                    }
                                }
                                var v44 int64 = (v18 >> ((18446744073709551615) & 63))
                                _ = v44
                            } else {
                                v2 = 65535
                            }
                        }
                    }
                    var v45 int32 = -2147483648
                    _ = fi32
                    _ = v33
                    _ = v45
                }
                _ = fi29
                _ = v30
                _ = v31
            }
        }
    }
    fmt.Println(v3, v17)
    if true {
        v46 := []int{1, 9223372036854775807, -2147483648, int(v18)}
        v18 = (v18 / (((^(v18))) | 1))
        _ = v46
    } else {
        v18 += (v18 >> ((v2) & 63))
    }
    var v47 uint16 = v17
    for fi48 := 0; fi48 < 2; fi48++ {
        v2 = (v2 + (^(v2)))
        fmt.Println(v1)
        _ = fi48
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v47
}
