package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 uint; F2 uint64 }
func genFn0(p0 uint, p1 int64, p2 uint8) int8 {
    for fi7 := 0; fi7 < 20; fi7++ {
        switch fi7 {
        case 1:
            switch (^(fi7)) {
            case 128:
                for fi8 := 0; fi8 < 20; fi8++ {
                    v9 := []int{(^(fi7)), (fi8 * (-(fi7)))}
                    var v10 int64 = 953
                    _ = fi8
                    _ = v10
                    _ = v9
                }
            case 4294967295:
                fmt.Println(p0, p1, p1)
            case -2147483648:
                v11 := GenS1{F0: ((("" + "") + "\n\t") + ("foo bar" + "\n\t"))}
                _ = v11
            default:
                p1 -= (p1 << ((p0) & 63))
            }
        }
        _ = fi7
    }
    p2 = 1
    v12 := []int{-9223372036854775808, 34}
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return -1
}
func genFn1(p0 uint32, p1 int64) uint32 {
    for fi13 := 0; fi13 < 9; fi13++ {
        v14 := GenS1{F0: ("foo bar" + ((("a" + ("" + "\n\t")) + "") + "a"))}
        var v15 uint16 = 508
        p0 += p0
        _ = fi13
        _ = v14
        _ = v15
    }
    var v16 float64 = (1e10 + float64(0.1))
    _ = p0
    _ = p1
    _ = v16
    return p0
}

func main() {
    v1 := []int{1, 32768}
    var v2 uint = 0
    v3 := make(map[string]uint32)
    v4 := make(map[string]int)
    v5 := make(map[string]uint32)
    v6 := []int{}
    var v17 int8 = 14
    switch 127 {
    case -428:
        if (799 < 0) {
            _ = len(v4)
            switch -1 {
            case 4294967295:
                if (-957 > 1) {
                    v18 := GenS1{F0: ("a" + "hello")}
                    _ = v18
                } else {
                    if len(v1) > 0 {
                        v1[0] = 4294967295
                    }
                    fmt.Println(v2, v17)
                }
            case -9223372036854775808:
                if true {
                    for fi19 := 0; fi19 < 16; fi19++ {
                        v20 := GenS1{F0: ("foo bar" + "a")}
                        for fi21 := 0; fi21 < 9; fi21++ {
                            var v22 uint16 = 690
                            if false {
                                var v23 uint8 = 42
                                _ = v23
                            }
                            fmt.Println(fi19, v2)
                            _ = fi21
                            _ = v22
                        }
                        _ = fi19
                        _ = v20
                    }
                }
            case 32767:
                fmt.Println(v2, v2)
            }
            v24 := GenS1{F0: "hello"}
            _ = v24
        }
    }
    v25 := GenS2{F0: true, F1: v2, F2: 1}
    _ = len(v3)
    v26 := v3[("hello" + "hello")]
    v2 -= (v2 * 18446744073709551615)
    var v27 float64 = float64(1.5)
    var v28 uint32 = genFn1(v26, 9223372036854775807)
    v28 = (v28 | (v26 * (v28 ^ v26)))
    v29 := GenS2{F0: false, F1: (v2 - (v2 % ((v2) | 1))), F2: 840}
    if len(v1) > 0 {
        v1[0] = 32767
    }
    _ = len(v4)
    v30 := GenS2{F0: true, F1: v2, F2: 2}
    v26 = genFn1((v28 - (v26 | (v28 & v28))), 1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
