package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 string; F1 int32 }
type GenS4 struct { F0 int64; F1 uint16; F2 int }
type GenS5 struct { F0 string; F1 int8 }
type GenS6 struct { F0 int64 }
func genFn0(p0 int16, p1 string, p2 int) int64 {
    v6 := GenS1{F0: 42, F1: 18446744073709551615, F2: (p0 & 32767)}
    if false {
        p0 = (p0 & p0)
    } else {
        v7 := []int{(p2 * (p2 % ((2147483647) | 1))), (p2 + p2), p2}
        switch p1 {
        case "hello":
            for fi8 := 0; fi8 < 15; fi8++ {
                var v9 uint8 = 42
                _ = fi8
                _ = v9
            }
        case "\n\t":
            for fi10 := 0; fi10 < 6; fi10++ {
                p1 = (p1 + p1)
                var v11 bool = true
                _ = fi10
                _ = v11
            }
        case "a":
            fmt.Println(p1, p0, p2)
        default:
            p1 = ((("\n\t" + ("a" + ("\n\t" + p1))) + p1) + "hello")
        }
        for fi12 := 0; fi12 < 11; fi12++ {
            fmt.Println(fi12)
            v7 = append(v7, (^(fi12)))
            for fi13 := 0; fi13 < 16; fi13++ {
                p2 = 32767
                if len(v7) > 0 {
                    v7[0] = 9223372036854775807
                }
                v14 := GenS1{F0: 4294967295, F1: 18446744073709551615, F2: (p0 | (p0 << ((1) & 63)))}
                _ = fi13
                _ = v14
            }
            _ = fi12
        }
        _ = v7
    }
    v15 := GenS2{F0: 100}
    if (false || (1.5 < 3.14159)) {
        p0 = 42
    } else {
        v16 := []int{32768}
        _ = v16
    }
    fmt.Println(p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    _ = v6
    return 9223372036854775807
}
func genFn1() uint {
    var v17 uint8 = 183
    if (685 <= 2147483647) {
        v18 := GenS1{F0: 2147483648, F1: 114, F2: -1}
        var v19 uint = 18446744073709551615
        _ = v18
        _ = v19
    }
    v17 = (v17 << ((0) & 63))
    fmt.Println(v17)
    for fi20 := 0; fi20 < 7; fi20++ {
        for fi21 := 0; fi21 < 20; fi21++ {
            fmt.Println(fi20)
            v17 += (v17 + 2)
            v22 := GenS2{F0: 48}
            _ = fi21
            _ = v22
        }
        v17 -= v17
        _ = fi20
    }
    _ = v17
    return 4294967295
}

func main() {
    v1 := []uint32{2147483647, 2147483647, 1}
    var v2 uint16 = 255
    var v3 uint64 = 75
    var v4 int64 = 42
    v5 := make(map[int]bool)
    for fi23 := 0; fi23 < 9; fi23++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v24 int64 = genFn0(42, "foo bar", (^(fi23)))
        v25 := GenS1{F0: 1, F1: 2, F2: -1}
        _ = fi23
        _ = v24
        _ = v25
    }
    if true {
        v26 := v5[-1]
        fmt.Println(v26, v3, v4)
        _ = v26
    }
    v27 := GenS3{F0: "a", F1: 538}
    v28 := v5[-266]
    switch 127 {
    case 207:
        v29 := GenS3{F0: "hello", F1: -2147483648}
        _ = v29
    default:
        var v30 string = ("" + "")
        _ = v30
    }
    for fi31 := 0; fi31 < 11; fi31++ {
        for fi32 := 0; fi32 < 12; fi32++ {
            if (128 > 128) {
                var v33 float32 = 0.1
                var v34 uint16 = v2
                _ = v33
                _ = v34
            } else {
                v2 = (v2 / (((^(v2))) | 1))
            }
            var v35 uint32 = uint32(v2)
            _ = fi32
            _ = v35
        }
        _ = fi31
    }
    v1 = append(v1, 1)
    var v36 int = 9223372036854775807
    switch "" {
    case "hello":
        fmt.Println(v4)
    case "foo bar":
        fmt.Println(v4)
    case "a":
        if true {
            var v37 uint16 = (v2 - (^(v2)))
            _ = len(v1)
            _ = v37
        } else {
            switch (-(v36)) {
            case 32767:
                if v28 {
                    v38 := GenS1{F0: 174, F1: 65535, F2: -15}
                    v36 = (v36 & (v36 - 32767))
                    _ = v38
                } else {
                    for fi39 := 0; fi39 < 7; fi39++ {
                        v28 = v28
                        _ = fi39
                    }
                }
            case 128:
                v40 := v5[1]
                _ = v40
            }
        }
    default:
        for fi41 := 0; fi41 < 17; fi41++ {
            var v42 float64 = (1.5 + 0.1)
            _ = fi41
            _ = v42
        }
    }
    fmt.Println(v28, v28)
    v36 = 42
    v4 -= genFn0(32767, ("" + "a"), 128)
    var v43 int64 = v4
    var v44 uint = 42
    var v45 uint32 = 775
    v46 := GenS4{F0: v4, F1: 65535, F2: (v36 + v36)}
    if (0 >= 1) {
        var v47 uint16 = (^(v2))
        v2 = 2
        _ = v47
    }
    for fi48 := 0; fi48 < 1; fi48++ {
        v49 := GenS5{F0: "foo bar", F1: 0}
        v28 = ((^(v44)) > (v44 ^ (v44 | (v44 & (v44 - (^(v44)))))))
        v50 := GenS6{F0: v4}
        _ = fi48
        _ = v49
        _ = v50
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v36
    _ = v4
    _ = v43
    _ = v44
    _ = v45
    _ = v46
    _ = v5
}
