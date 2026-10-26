package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 uint32 }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 string; F1 string }
func genFn0(p0 uint64, p1 string, p2 bool) int {
    p2 = p2
    v6 := GenS1{F0: p1, F1: -69, F2: 135}
    var v7 string = "\n\t"
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return -890
}
func genFn1(p0 int, p1 int16) bool {
    switch p0 {
    case 4294967295:
        for fi8 := 0; fi8 < 4; fi8++ {
            p0 -= fi8
            var v9 int64 = -9223372036854775808
            var v10 uint16 = 42
            _ = fi8
            _ = v10
            _ = v9
        }
    case 128:
        for fi11 := 0; fi11 < 14; fi11++ {
            var v12 int = genFn0(2, "hello", (!(false || false)))
            switch fi11 {
            case -613:
                p1 -= p1
            case -582:
                var v13 uint8 = 42
                _ = v13
            case 306:
                fmt.Println(v12)
            }
            _ = fi11
            _ = v12
        }
    case -9223372036854775808:
        p1 += p1
    }
    for fi14 := 0; fi14 < 15; fi14++ {
        v15 := GenS1{F0: "a", F1: 133, F2: 65535}
        _ = fi14
        _ = v15
    }
    p0 = p0
    switch (p0 + p0) {
    case -1:
        v16 := []int{p0, (-(p0)), p0}
        _ = v16
    case 42:
        p1 -= p1
    }
    _ = p0
    _ = p1
    return true
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]uint32)
    var v3 float64 = 1.5
    v4 := make(map[int]string)
    v5 := make(map[int]string)
    for fi17 := 0; fi17 < 20; fi17++ {
        switch (^(fi17)) {
        case 9223372036854775807:
            v18 := GenS1{F0: "hello", F1: 878, F2: 706}
            _ = v18
        default:
            delete(v5, (fi17 + (fi17 >> ((25) & 63))))
        }
        if true {
            v19 := []int{-1, (-(fi17))}
            v20 := GenS1{F0: "a", F1: -706, F2: 6}
            _ = v19
            _ = v20
        } else {
            var v21 uint64 = 1
            if true {
                fmt.Println(fi17, v21, v3)
                if true {
                    v3 -= v3
                    v22 := GenS2{F0: 1, F1: 2147483648}
                    _ = v22
                } else {
                    v23 := []int{1}
                    _ = v23
                }
                for fi24 := 0; fi24 < 9; fi24++ {
                    v21 = (v21 << ((0) & 63))
                    _ = fi24
                }
            } else {
                for fi25 := 0; fi25 < 3; fi25++ {
                    var v26 bool = false
                    switch ("" + "foo bar") {
                    case "foo bar":
                        v27 := GenS2{F0: 487, F1: 65535}
                        _ = v27
                    case "hello":
                        var v28 uint64 = 170
                        _ = v28
                    }
                    fmt.Println(fi17, fi17)
                    _ = fi25
                    _ = v26
                }
                v29 := GenS2{F0: 42, F1: 1}
                _ = v29
            }
            var v30 uint64 = 4294967295
            _ = v21
            _ = v30
        }
        v3 = v3
        _ = fi17
    }
    v31 := GenS3{F0: ("" + "\n\t"), F1: "foo bar"}
    var v32 int8 = -1
    var v33 int32 = 799
    var v34 float32 = ((0.0 + (float32((v3 / v3)) - float32(((float64((0.0 + 1e10)) - float64((0.1 * (0.1 / float32(1.5))))) + (v3 * (v3 - float64(-2.25))))))) + float32(0.0))
    var v35 int64 = 9223372036854775807
    v36 := []int{-423, 9223372036854775807, -9223372036854775808}
    v3 = 0.1
    if true {
        fmt.Println(v34, v32)
        for fi37 := 0; fi37 < 5; fi37++ {
            var v38 string = (("a" + ("" + "foo bar")) + ("foo bar" + ("a" + ("a" + ""))))
            fmt.Println(v3, v3)
            _ = fi37
            _ = v38
        }
    } else {
        var v39 int64 = -198
        v35 = (v39 >> ((18446744073709551615) & 63))
        _ = v39
    }
    switch "a" {
    case "hello":
        if true {
            for fi40 := 0; fi40 < 17; fi40++ {
                if true {
                    v41 := GenS3{F0: "\n\t", F1: "\n\t"}
                    _ = v41
                } else {
                    fmt.Println(fi40, v32, v34)
                    var v42 float64 = float64(float32((((v3 + (((1.5 / 1.0) / 1.0) - 1.5)) + v3) * v3)))
                    if false {
                        fmt.Println(v42)
                        var v43 int8 = v32
                        _ = v43
                    }
                    _ = v42
                }
                v32 += (v32 / (((^(v32))) | 1))
                _ = fi40
            }
            var v44 uint = 814
            _ = v44
        }
    case "":
        if (262 < 2) {
            var v45 int = 152
            _ = v45
        } else {
            v46 := GenS1{F0: "\n\t", F1: v35, F2: 2147483647}
            _ = v46
        }
    case "\n\t":
        if len(v36) > 0 {
            v36[0] = genFn0(42, "hello", true)
        }
    default:
        fmt.Println(v32)
    }
    v3 -= float64(((((v34 - v34) + v34) + float32((v3 / 1.5))) - 3.14159))
    fmt.Println(v33, v35)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
