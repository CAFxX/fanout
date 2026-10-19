package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int; F1 int16; F2 bool }
type GenS3 struct { F0 int16; F1 uint8; F2 uint }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint16; F1 int32 }
type GenS6 struct { F0 string; F1 int16; F2 int8 }
type GenS7 struct { F0 int32; F1 int }
type GenS8 struct { F0 int16; F1 int32 }
func genFn0() int16 {
    var v7 int = -639
    for fi8 := 0; fi8 < 16; fi8++ {
        v7 = 128
        if true {
            if false {
                fmt.Println(fi8, fi8)
                if false {
                    switch "a" {
                    case "foo bar":
                        v9 := GenS1{F0: true}
                        _ = v9
                    }
                    var v10 uint64 = 18446744073709551615
                    _ = v10
                } else {
                    fmt.Println(fi8, v7)
                    v7 = -2147483648
                }
                var v11 uint8 = 2
                _ = v11
            } else {
                for fi12 := 0; fi12 < 6; fi12++ {
                    var v13 float32 = -2.25
                    _ = fi12
                    _ = v13
                }
                v14 := GenS2{F0: -662, F1: 0, F2: false}
                v15 := GenS2{F0: (fi8 >> ((168) & 63)), F1: 1, F2: ((("hello" <= "") || (true || (!false))) && true)}
                _ = v14
                _ = v15
            }
            v7 = v7
        }
        v7 = 128
        _ = fi8
    }
    _ = v7
    return 1000
}
func genFn1(p0 int, p1 int64) uint {
    var v16 string = "a"
    fmt.Println(p0, p1, v16)
    p0 = (p0 << ((18446744073709551615) & 63))
    var v17 int64 = p1
    p0 = 317
    _ = p0
    _ = p1
    _ = v16
    _ = v17
    return 646
}

func main() {
    var v1 int = -1
    v2 := make(map[uint32]int)
    var v3 uint16 = 32767
    v4 := []uint32{296, 65535, 228, 2147483647}
    var v5 uint16 = (v3 ^ (v3 + v3))
    var v6 int = (v1 & v1)
    if false {
        v18 := GenS3{F0: 94, F1: 127, F2: 255}
        v19 := GenS1{F0: (128 < (v6 % ((v6) | 1)))}
        _ = v18
        _ = v19
    } else {
        switch (v6 * (v6 + v6)) {
        case 397:
            v20 := GenS4{F0: 709}
            _ = v20
        case -2147483648:
            switch "foo bar" {
            case "\n\t":
                v6 -= (v1 >> ((genFn1((^(v1)), 1)) & 63))
            case "foo bar":
                v5 = 0
            default:
                v2[405] = 97
            }
        }
        for fi21 := 0; fi21 < 11; fi21++ {
            v5 = 2
            for fi22 := 0; fi22 < 17; fi22++ {
                switch fi22 {
                case -9223372036854775808:
                    v5 = v3
                case 42:
                    var v23 int = (fi21 | 4294967295)
                    _ = v23
                case -197:
                    v6 -= (v6 | v1)
                }
                delete(v2, 42)
                v24 := GenS2{F0: (v6 ^ 32768), F1: 991, F2: false}
                _ = fi22
                _ = v24
            }
            var v25 int64 = -9223372036854775808
            _ = fi21
            _ = v25
        }
        v1 = v6
    }
    if (false && (true || (v3 == (^(v5))))) {
        fmt.Println(v1)
        fmt.Println(v5)
        v1 = 12
    }
    fmt.Println(v5, v1, v1)
    v26 := GenS5{F0: v5, F1: 1}
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch "foo bar" {
    case "":
        for fi27 := 0; fi27 < 18; fi27++ {
            var v28 string = "hello"
            var v29 uint64 = 4294967295
            _ = fi27
            _ = v28
            _ = v29
        }
    case "\n\t":
        v30 := GenS2{F0: v6, F1: 42, F2: (0 >= 972)}
        _ = v30
    case "hello":
        v1 += v1
    default:
        fmt.Println(v1, v1)
    }
    _ = len(v2)
    for fi31 := 0; fi31 < 1; fi31++ {
        for fi32 := 0; fi32 < 16; fi32++ {
            switch v1 {
            case -1:
                v1 += fi31
            case 42:
                if false {
                    v3 = v5
                } else {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
            case 32767:
                v2[65535] = v1
            }
            v33 := GenS6{F0: "a", F1: genFn0(), F2: -1}
            _ = fi32
            _ = v33
        }
        v5 = v3
        v5 = v5
        _ = fi31
    }
    var v34 uint32 = 150
    if true {
        v1 = v6
        v35 := GenS7{F0: 2147483647, F1: (v1 >> ((65535) & 63))}
        _ = v35
    } else {
        v5 -= (v3 & (^(v5)))
        var v36 int32 = 42
        _ = v36
    }
    for fi37 := 0; fi37 < 4; fi37++ {
        if (177 > 127) {
            v5 = v3
            if true {
                v38 := GenS4{F0: v34}
                _ = v38
            } else {
                v3 -= v3
            }
        }
        if (1000 > 32767) {
            fmt.Println(v1, v6)
        }
        v39 := GenS5{F0: (v5 % ((v5) | 1)), F1: -821}
        _ = fi37
        _ = v39
    }
    for fi40 := 0; fi40 < 8; fi40++ {
        delete(v2, (v34 & v34))
        var v41 bool = true
        _ = fi40
        _ = v41
    }
    var v42 uint16 = (v5 | (v3 ^ uint16(v34)))
    if (206 < 691) {
        v43 := v2[v34]
        if true {
            var v44 uint = uint(v34)
            var v45 string = "a"
            _ = v44
            _ = v45
        }
        v43 += (v6 ^ -9223372036854775808)
        _ = v43
    } else {
        v46 := GenS8{F0: 1, F1: -2147483648}
        _ = v46
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v34
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
