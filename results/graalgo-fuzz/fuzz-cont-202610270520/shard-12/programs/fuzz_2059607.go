package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 uint }
type GenS2 struct { F0 int16; F1 bool }
func genFn0(p0 uint, p1 int8, p2 bool) uint64 {
    if ((!(-867 > 42)) && (0.1 > 1.5)) {
        switch -2147483648 {
        case -1:
            switch (("a" + ("foo bar" + "\n\t")) + "") {
            case "\n\t":
                var v6 bool = ((p2 && p2) || p2)
                _ = v6
            }
        case 1:
            fmt.Println(p0)
        case -794:
            v7 := []int{4294967295, 42, 1}
            _ = v7
        }
        v8 := GenS1{F0: 1, F1: 1, F2: (^(p0))}
        _ = v8
    } else {
        p2 = p2
    }
    switch 0 {
    case 128:
        p0 -= 322
    case 2147483647:
        p1 = 1
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}
func genFn1(p0 int16) bool {
    if (610 > 42) {
        p0 = p0
    } else {
        if (1e10 >= (1.5 * 1.5)) {
            for fi9 := 0; fi9 < 18; fi9++ {
                var v10 int8 = -78
                _ = fi9
                _ = v10
            }
        }
    }
    var v11 int32 = 0
    p0 = p0
    _ = p0
    _ = v11
    return (625 == -2147483648)
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []int{0, 374}
    var v3 uint8 = 0
    v4 := make(map[uint32]bool)
    var v5 int16 = 1000
    fmt.Println(v5, v3)
    if true {
        v12 := v1[32768]
        v12 -= (v12 * v12)
        _ = v12
    }
    if (2147483647 >= 1) {
        if len(v2) > 0 {
            v2[0] = 249
        }
        v3 = (v3 % ((v3) | 1))
        if true {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        } else {
            var v13 int32 = -1
            _ = v13
        }
    }
    fmt.Println(v3)
    var v14 string = "foo bar"
    _ = len(v2)
    switch ((v14 + v14) + v14) {
    case "":
        if true {
            var v15 string = "foo bar"
            v15 = v15
            _ = v15
        } else {
            v16 := GenS2{F0: v5, F1: false}
            for fi17 := 0; fi17 < 11; fi17++ {
                var v18 int64 = 1
                if true {
                    fmt.Println(v5)
                }
                v4[4294967295] = false
                _ = fi17
                _ = v18
            }
            _ = v16
        }
    case "a":
        var v19 int32 = 440
        _ = v19
    }
    if (!true) {
        fmt.Println(v5, v5, v5)
        v20 := GenS1{F0: 1, F1: 807, F2: 18446744073709551615}
        for fi21 := 0; fi21 < 15; fi21++ {
            for fi22 := 0; fi22 < 17; fi22++ {
                if (87 == (v3 ^ v3)) {
                    v5 = (v5 ^ -32768)
                }
                fmt.Println(fi22)
                v23 := GenS1{F0: -2147483648, F1: 327, F2: 18446744073709551615}
                var v24 int64 = -9223372036854775808
                v25 := v1[(fi22 - (fi21 - (fi22 ^ fi21)))]
                _ = fi22
                _ = v23
                _ = v24
                _ = v25
            }
            v26 := GenS1{F0: int(v3), F1: 989, F2: 508}
            var v27 uint16 = 255
            _ = fi21
            _ = v26
            _ = v27
        }
        _ = v20
    }
    v28 := GenS2{F0: v5, F1: ((42 < 450) || false)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
