package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 int8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int32; F1 uint32; F2 bool }
func genFn0(p0 int32, p1 uint, p2 bool) int32 {
    var v6 int = 414
    if p2 {
        for fi7 := 0; fi7 < 1; fi7++ {
            var v8 int32 = (p0 - p0)
            v9 := []int{fi7, v6, int(p1)}
            v10 := GenS1{F0: 81, F1: 974, F2: 100}
            _ = fi7
            _ = v10
            _ = v8
            _ = v9
        }
        if (false && (p0 < (p0 >> ((4294967295) & 63)))) {
            v11 := []int{128, v6, (-(v6))}
            for fi12 := 0; fi12 < 2; fi12++ {
                var v13 uint8 = 255
                _ = fi12
                _ = v13
            }
            if p2 {
                var v14 int8 = 100
                _ = v14
            }
            _ = v11
        } else {
            fmt.Println(p0, p2, p1)
        }
    }
    fmt.Println(p1)
    if (-128 < 124) {
        for fi15 := 0; fi15 < 4; fi15++ {
            p1 = p1
            _ = fi15
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return -1
}
func genFn1(p0 int64, p1 uint64) uint32 {
    switch ("foo bar" + "") {
    case "":
        for fi16 := 0; fi16 < 13; fi16++ {
            for fi17 := 0; fi17 < 19; fi17++ {
                p1 = (p1 / (((p1 / ((p1) | 1))) | 1))
                _ = fi17
            }
            var v18 float32 = (0.1 / (1e10 - 1.5))
            _ = fi16
            _ = v18
        }
    case "\n\t":
        p0 = p0
    case "foo bar":
        var v19 float64 = -2.25
        _ = v19
    }
    v20 := []int{0}
    _ = p0
    _ = p1
    _ = v20
    return 2147483648
}

func main() {
    v1 := []int{-1, 799}
    var v2 uint32 = 324
    v3 := []int{}
    v4 := make(map[int]int)
    v5 := make(map[int]string)
    var v21 float32 = ((1.5 + 0.0) / 0.1)
    if (293 > 1) {
        for fi22 := 0; fi22 < 8; fi22++ {
            var v23 int32 = -1
            v2 -= 2147483647
            _ = fi22
            _ = v23
        }
        var v24 uint8 = 111
        if len(v3) > 0 {
            v3[0] = 32767
        }
        _ = v24
    }
    v2 = (v2 % ((978) | 1))
    var v25 uint64 = 388
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch "a" {
    case "":
        v26 := GenS1{F0: 255, F1: (v25 + v25), F2: -72}
        _ = v26
    case "foo bar":
        var v27 uint32 = (v2 + 1)
        _ = v27
    case "a":
        fmt.Println(v21, v25)
    default:
        if false {
            v21 = (0.1 - v21)
            delete(v4, 289)
            v21 += ((3.14159 * 0.1) + (v21 - v21))
        } else {
            switch 0 {
            case -2147483648:
                switch "a" {
                case "a":
                    v28 := GenS1{F0: 165, F1: v25, F2: -1}
                    _ = v28
                case "hello":
                    for fi29 := 0; fi29 < 15; fi29++ {
                        var v30 int16 = 0
                        _ = fi29
                        _ = v30
                    }
                case "":
                    fmt.Println(v21, v2)
                default:
                    var v31 uint = 688
                    _ = v31
                }
            }
            v25 = v25
        }
    }
    var v32 int64 = 1
    if false {
        if (true && false) {
            var v33 int8 = 100
            _ = v33
        } else {
            v34 := GenS2{F0: (false || true)}
            switch "foo bar" {
            case "hello":
                if false {
                    fmt.Println(v25, v32, v25)
                } else {
                    var v35 int64 = v32
                    var v36 int8 = 0
                    _ = v35
                    _ = v36
                }
            case "foo bar":
                for fi37 := 0; fi37 < 20; fi37++ {
                    var v38 bool = (!false)
                    _ = fi37
                    _ = v38
                }
            case "\n\t":
                var v39 int32 = genFn0(int32(v2), 4294967295, (v32 == v32))
                _ = v39
            default:
                v2 = 42
            }
            _ = v34
        }
        for fi40 := 0; fi40 < 3; fi40++ {
            switch (("a" + ("\n\t" + ("foo bar" + "a"))) + "\n\t") {
            case "a":
                var v41 float32 = (v21 + v21)
                _ = v41
            case "":
                if (false || false) {
                    v42 := GenS3{F0: 0}
                    delete(v4, fi40)
                    v21 = v21
                    _ = v42
                }
            }
            _ = fi40
        }
    } else {
        _ = len(v4)
    }
    for fi43 := 0; fi43 < 15; fi43++ {
        _ = len(v1)
        var v44 int64 = v32
        v32 = (v44 & v44)
        _ = fi43
        _ = v44
    }
    switch 128 {
    case 9223372036854775807:
        var v45 uint8 = 0
        _ = v45
    case -337:
        switch "" {
        case "foo bar":
            var v46 float32 = (1.5 / v21)
            _ = v46
        case "\n\t":
            v2 = genFn1(v32, (v25 * 83))
        }
    default:
        var v47 uint64 = 18446744073709551615
        _ = v47
    }
    if false {
        var v48 int32 = 42
        v21 += ((v21 - float32((-2.25 + (-2.25 + 0.1)))) * (-2.25 / 1.0))
        _ = v48
    }
    v1 = append(v1, 42)
    v49 := GenS4{F0: -2147483648, F1: v2, F2: ((!true) || (true || (!false)))}
    v50 := GenS4{F0: -602, F1: v2, F2: (v21 != ((v21 / (0.0 - (1e10 * v21))) * float32(float64(v21))))}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v21
    _ = v25
    _ = v3
    _ = v32
    _ = v4
    _ = v49
    _ = v5
    _ = v50
}
