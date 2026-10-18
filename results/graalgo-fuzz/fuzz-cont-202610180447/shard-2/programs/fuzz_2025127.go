package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint32; F1 int8; F2 uint }
type GenS3 struct { F0 string; F1 uint32; F2 uint }
type GenS4 struct { F0 int32; F1 uint64; F2 int }
func genFn0(p0 int8, p1 bool, p2 int32) int32 {
    v6 := []int{}
    if len(v6) > 0 {
        v6[0] = 460
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return p2
}
func genFn1(p0 int, p1 string, p2 int8) int16 {
    p2 = p2
    for fi7 := 0; fi7 < 9; fi7++ {
        v8 := GenS1{F0: p2}
        v9 := []int{(^(p0)), 189}
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 int16 = -385
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return (-(v10))
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]int)
    v3 := []string{(("foo bar" + "a") + ""), (("\n\t" + ("foo bar" + "a")) + "hello"), "hello"}
    var v4 int8 = 127
    var v5 string = ((("\n\t" + "") + "") + "foo bar")
    for fi11 := 0; fi11 < 9; fi11++ {
        v12 := GenS1{F0: v4}
        v13 := GenS1{F0: v4}
        _ = fi11
        _ = v12
        _ = v13
    }
    v5 = "a"
    var v14 uint8 = 128
    v15 := GenS2{F0: 0, F1: v4, F2: 18446744073709551615}
    for fi16 := 0; fi16 < 2; fi16++ {
        fmt.Println(v5, v5, fi16)
        v14 += v14
        _ = fi16
    }
    v2[("\n\t" + v5)] = -9223372036854775808
    var v17 uint32 = 594
    v18 := v2[v5]
    if (v5 <= "a") {
        v19 := GenS2{F0: 2147483647, F1: -1, F2: 0}
        _ = v19
    } else {
        fmt.Println(v14, v5, v17)
        var v20 int32 = -1
        var v21 uint64 = 4294967295
        _ = v20
        _ = v21
    }
    for fi22 := 0; fi22 < 14; fi22++ {
        fmt.Println(v14, v17, fi22)
        switch ((((v5 + "a") + "hello") + v5) + (v5 + (("a" + v5) + ("foo bar" + v5)))) {
        case "a":
            for fi23 := 0; fi23 < 20; fi23++ {
                v24 := GenS1{F0: v4}
                v17 = (v17 * (v17 / (((v17 & v17)) | 1)))
                _ = fi23
                _ = v24
            }
        default:
            v17 = (v17 & (v17 | (v17 / ((v17) | 1))))
        }
        _ = len(v2)
        _ = fi22
    }
    if true {
        fmt.Println(v5, v18)
        for fi25 := 0; fi25 < 4; fi25++ {
            switch ("a" + v5) {
            case "a":
                v17 -= (^(v17))
            case "":
                v3 = append(v3, "hello")
            default:
                for fi26 := 0; fi26 < 15; fi26++ {
                    fmt.Println(v17, v18)
                    switch (fi26 - fi26) {
                    case -672:
                        delete(v2, v5)
                    case -2147483648:
                        var v27 int16 = -32768
                        _ = v27
                    }
                    _ = fi26
                }
            }
            var v28 int8 = v4
            v3 = append(v3, v5)
            _ = fi25
            _ = v28
        }
    } else {
        v29 := GenS3{F0: v5, F1: (v17 | (v17 << ((717) & 63))), F2: 2}
        var v30 uint32 = v17
        _ = v29
        _ = v30
    }
    v14 -= (^(v14))
    fmt.Println(v4, v14, v14)
    switch v18 {
    case 1:
        v31 := GenS4{F0: 2147483647, F1: 0, F2: -1}
        _ = v31
    case -981:
        switch (("" + (v5 + "hello")) + (v5 + (v5 + (v5 + v5)))) {
        case "\n\t":
            if true {
                fmt.Println(v5)
                fmt.Println(v4, v14)
                v4 = (v4 & 1)
            }
        case "a":
            for fi32 := 0; fi32 < 9; fi32++ {
                var v33 uint8 = v14
                if (18446744073709551615 != 18446744073709551615) {
                    v34 := GenS1{F0: v4}
                    _ = v34
                }
                v35 := GenS2{F0: v17, F1: v4, F2: 65535}
                _ = fi32
                _ = v33
                _ = v35
            }
        default:
            var v36 uint8 = (^(v14))
            _ = v36
        }
    case -1:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    default:
        for fi37 := 0; fi37 < 16; fi37++ {
            v4 = -2
            fmt.Println(v5)
            _ = fi37
        }
    }
    v3 = append(v3, "\n\t")
    v4 = (v4 / ((-128) | 1))
    v17 -= uint32(v14)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
