package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int8; F1 uint8; F2 bool }
func genFn0(p0 uint64, p1 bool) uint64 {
    v4 := GenS1{F0: p1, F1: 57}
    p1 = (!p1)
    v5 := []int{-203, -462}
    for fi6 := 0; fi6 < 19; fi6++ {
        if (42 == -1) {
            switch "foo bar" {
            case "\n\t":
                var v7 int64 = 1
                _ = v7
            case "foo bar":
                p0 -= p0
            case "":
                switch fi6 {
                case -427:
                    if p1 {
                        for fi8 := 0; fi8 < 4; fi8++ {
                            p1 = p1
                            if (!p1) {
                                switch -2147483648 {
                                case 9223372036854775807:
                                    fmt.Println(p0, p0)
                                    p0 = (p0 + 42)
                                    var v9 int16 = 1
                                    _ = v9
                                case 264:
                                    var v10 float64 = 1.5
                                    _ = v10
                                }
                                if p1 {
                                    v5 = append(v5, fi6)
                                    p1 = (p0 != p0)
                                    fmt.Println(fi8, fi8)
                                }
                                v11 := GenS1{F0: (p1 || true), F1: 951}
                                _ = v11
                            } else {
                                var v12 uint16 = 32767
                                if false {
                                    v5 = append(v5, (fi6 << ((18446744073709551615) & 63)))
                                }
                                _ = v12
                            }
                            v13 := GenS2{F0: 313}
                            _ = fi8
                            _ = v13
                        }
                        v14 := GenS2{F0: 1}
                        for fi15 := 0; fi15 < 3; fi15++ {
                            p1 = false
                            _ = fi15
                        }
                        _ = v14
                    }
                }
            default:
                if (-19 >= 1) {
                    v16 := GenS2{F0: 2}
                    _ = v16
                }
            }
        }
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 859
}
func genFn1() uint8 {
    var v17 uint = 65535
    v17 -= (v17 & v17)
    v17 += 4294967295
    v17 = 440
    var v18 int = 817
    _ = v17
    _ = v18
    return 211
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []int{}
    var v3 int16 = 1
    if true {
        var v19 int32 = -2147483648
        v19 = v19
        v19 += v19
        _ = v19
    } else {
        var v20 uint8 = 128
        _ = v20
    }
    v3 = v3
    v21 := GenS2{F0: 127}
    for fi22 := 0; fi22 < 15; fi22++ {
        switch -2147483648 {
        case -2147483648:
            switch (fi22 % ((fi22) | 1)) {
            case 127:
                v3 += (v3 - v3)
            case -2147483648:
                v3 = v3
            }
        case 42:
            for fi23 := 0; fi23 < 14; fi23++ {
                if true {
                    switch 9223372036854775807 {
                    case 4294967295:
                        v24 := GenS2{F0: 994}
                        _ = v24
                    case 0:
                        switch ((("a" + ("" + "foo bar")) + "a") + "a") {
                        case "foo bar":
                            fmt.Println(fi23)
                        case "hello":
                            var v25 string = "foo bar"
                            _ = v25
                        }
                    }
                    v26 := GenS1{F0: true, F1: -2147483648}
                    _ = v26
                }
                v27 := v1[2147483647]
                var v28 float64 = (((0.0 - 1e10) - 0.0) * 0.1)
                _ = fi23
                _ = v27
                _ = v28
            }
        case 9223372036854775807:
            v2 = append(v2, (-(fi22)))
        }
        var v29 int32 = -265
        switch ("\n\t" + "a") {
        case "\n\t":
            var v30 int16 = -354
            _ = v30
        }
        _ = fi22
        _ = v29
    }
    v31 := GenS3{F0: -128, F1: 42, F2: true}
    _ = len(v1)
    v32 := GenS2{F0: 255}
    _ = len(v1)
    var v33 int = 42
    var v34 uint16 = 428
    var v35 int8 = 42
    _ = len(v2)
    v35 -= v35
    if (1e10 < 1.5) {
        var v36 int32 = 920
        fmt.Println(v3)
        fmt.Println(v36, v3)
        _ = v36
    }
    switch "" {
    case "":
        v3 += (v3 - v3)
    case "foo bar":
        v33 += v33
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
}
