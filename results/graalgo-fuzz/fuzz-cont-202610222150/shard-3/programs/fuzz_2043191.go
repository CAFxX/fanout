package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 uint32; F1 int16; F2 int }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint32; F1 string; F2 bool }
type GenS6 struct { F0 uint64; F1 int32; F2 bool }
func genFn0(p0 int16) uint {
    var v7 float32 = -2.25
    for fi8 := 0; fi8 < 1; fi8++ {
        switch fi8 {
        case 128:
            v7 = v7
        case 127:
            p0 = 694
        default:
            var v9 bool = false
            _ = v9
        }
        for fi10 := 0; fi10 < 11; fi10++ {
            var v11 uint64 = 18446744073709551615
            if true {
                v12 := GenS1{F0: (v11 & 316), F1: "a"}
                v13 := GenS2{F0: 212, F1: 0, F2: fi8}
                _ = v12
                _ = v13
            }
            v14 := GenS2{F0: 931, F1: (p0 | p0), F2: fi8}
            _ = fi10
            _ = v11
            _ = v14
        }
        var v15 uint8 = 154
        _ = fi8
        _ = v15
    }
    v16 := []int{-2147483648, 42}
    _ = p0
    _ = v16
    _ = v7
    return 868
}

func main() {
    var v1 string = "a"
    var v2 uint64 = 173
    var v3 int16 = 1000
    var v4 uint8 = 1
    var v5 bool = (0 >= 65535)
    v6 := make(map[int]string)
    fmt.Println(v1)
    if (false && false) {
        v17 := GenS3{F0: -9223372036854775808}
        _ = v17
    }
    if ((!true) || false) {
        v3 = 1
        v18 := GenS4{F0: (v4 ^ 255)}
        _ = v18
    }
    switch (v1 + "") {
    case "\n\t":
        v19 := GenS4{F0: (v4 * (v4 - v4))}
        _ = v19
    case "a":
        for fi20 := 0; fi20 < 11; fi20++ {
            v21 := GenS5{F0: 4294967295, F1: ((("" + v1) + v1) + v1), F2: (!false)}
            for fi22 := 0; fi22 < 7; fi22++ {
                v4 -= v4
                v4 -= (^(v4))
                v6[128] = v1
                _ = fi22
            }
            for fi23 := 0; fi23 < 9; fi23++ {
                v1 = v1
                if true {
                    v4 -= v4
                    if (true && v5) {
                        v2 = (v2 % ((838) | 1))
                    } else {
                        var v24 uint64 = (v2 | 18446744073709551615)
                        _ = v24
                    }
                    v25 := v6[-125]
                    _ = v25
                }
                var v26 uint8 = v4
                _ = fi23
                _ = v26
            }
            _ = fi20
            _ = v21
        }
    default:
        v27 := v6[128]
        _ = v27
    }
    var v28 float32 = (-2.25 * -2.25)
    fmt.Println(v28, v3)
    if (v5 || (v5 || v5)) {
        for fi29 := 0; fi29 < 5; fi29++ {
            v30 := GenS1{F0: v2, F1: "\n\t"}
            for fi31 := 0; fi31 < 15; fi31++ {
                switch fi31 {
                case 691:
                    fmt.Println(v28, v5, v2)
                case 2147483647:
                    v6[fi31] = v1
                case 9223372036854775807:
                    switch 900 {
                    case -1:
                        var v32 int8 = 29
                        _ = v32
                    case -9223372036854775808:
                        fmt.Println(v1, v3)
                    default:
                        var v33 int32 = 42
                        _ = v33
                    }
                    v34 := GenS6{F0: 42, F1: 1, F2: true}
                    _ = v34
                default:
                    var v35 uint32 = 4
                    _ = v35
                }
                _ = fi31
            }
            for fi36 := 0; fi36 < 14; fi36++ {
                fmt.Println(v3, v2, v3)
                _ = fi36
            }
            _ = fi29
            _ = v30
        }
        if v5 {
            v6[599] = ("a" + ("hello" + v1))
            for fi37 := 0; fi37 < 11; fi37++ {
                var v38 float32 = ((0.0 - 1.5) / float32(-2.25))
                for fi39 := 0; fi39 < 5; fi39++ {
                    switch "a" {
                    case "a":
                        var v40 int = -9223372036854775808
                        _ = v40
                    case "\n\t":
                        var v41 int8 = 0
                        _ = v41
                    default:
                        var v42 string = v1
                        v1 = v1
                        v28 += (float32(0.0) - -2.25)
                        _ = v42
                    }
                    var v43 uint32 = 4294967295
                    switch v1 {
                    case "a":
                        if true {
                            var v44 int = fi39
                            v45 := GenS3{F0: (fi39 / ((fi39) | 1))}
                            _ = v44
                            _ = v45
                        }
                    case "":
                        if (v1 != v1) {
                            _ = len(v6)
                        }
                    default:
                        v46 := []int{(-(fi39)), fi39}
                        _ = v46
                    }
                    _ = fi39
                    _ = v43
                }
                var v47 float64 = 1e10
                _ = fi37
                _ = v38
                _ = v47
            }
        }
    }
    var v48 int32 = -790
    var v49 uint16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v48
    _ = v49
    _ = v5
    _ = v6
}
