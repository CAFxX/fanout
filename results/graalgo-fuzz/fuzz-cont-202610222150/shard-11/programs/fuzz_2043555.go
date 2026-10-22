package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 int16; F1 uint; F2 uint8 }
type GenS3 struct { F0 int; F1 uint32 }
type GenS4 struct { F0 int64 }
func genFn0(p0 uint, p1 uint) string {
    switch 0 {
    case 873:
        v4 := []int{-69, -288, 32768}
        _ = v4
    case -427:
        v5 := GenS1{F0: "a", F1: 1}
        _ = v5
    case 499:
        v6 := GenS1{F0: "foo bar", F1: 0}
        _ = v6
    default:
        var v7 int = 9223372036854775807
        _ = v7
    }
    v8 := []int{-9223372036854775808, 4294967295}
    for fi9 := 0; fi9 < 1; fi9++ {
        var v10 int64 = -9223372036854775808
        _ = fi9
        _ = v10
    }
    var v11 string = "hello"
    for fi12 := 0; fi12 < 4; fi12++ {
        _ = len(v8)
        switch (fi12 + (fi12 ^ (fi12 | (fi12 / (((fi12 << ((p1) & 63))) | 1))))) {
        case 42:
            _ = len(v8)
        default:
            switch fi12 {
            case 32767:
                if ("" <= (((("a" + v11) + (v11 + v11)) + "a") + v11)) {
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                }
            case -1:
                v13 := GenS2{F0: 489, F1: 39, F2: 1}
                _ = v13
            case -9223372036854775808:
                v14 := GenS3{F0: (fi12 << ((p0) & 63)), F1: 4294967295}
                _ = v14
            default:
                var v15 uint32 = 932
                _ = v15
            }
        }
        _ = fi12
    }
    _ = p0
    _ = p1
    _ = v11
    _ = v8
    return (v11 + "hello")
}

func main() {
    var v1 string = "foo bar"
    v2 := make(map[int]uint32)
    v3 := make(map[int]int)
    var v16 float64 = -2.25
    v17 := v2[42]
    if false {
        for fi18 := 0; fi18 < 10; fi18++ {
            for fi19 := 0; fi19 < 20; fi19++ {
                if false {
                    v20 := []int{fi18, 4294967295, fi18}
                    v21 := GenS2{F0: 42, F1: 39, F2: 1}
                    _ = v20
                    _ = v21
                }
                v17 = v17
                v17 = v17
                v1 = v1
                switch v1 {
                case "foo bar":
                    var v22 int = fi18
                    _ = v22
                }
                _ = fi19
            }
            v23 := GenS2{F0: -32768, F1: 2, F2: 34}
            _ = fi18
            _ = v23
        }
    } else {
        delete(v2, int(v17))
        v24 := GenS4{F0: -705}
        _ = v24
    }
    if (true || true) {
        _ = len(v3)
        fmt.Println(v16, v17, v17)
        v25 := GenS2{F0: -49, F1: 255, F2: 128}
        _ = v25
    } else {
        if (!false) {
            var v26 uint64 = 2
            v27 := GenS4{F0: -1}
            switch (((v1 + "") + v1) + "hello") {
            case "\n\t":
                v17 = 671
            case "foo bar":
                fmt.Println(v17)
            default:
                if ((v17 * v17) > v17) {
                    v16 += v16
                    v28 := []int{128, -534}
                    for fi29 := 0; fi29 < 20; fi29++ {
                        var v30 int16 = -32768
                        for fi31 := 0; fi31 < 18; fi31++ {
                            var v32 int32 = -1
                            v30 = (-(v30))
                            v30 -= int16(fi31)
                            _ = fi31
                            _ = v32
                        }
                        _ = fi29
                        _ = v30
                    }
                    _ = v28
                }
            }
            _ = v26
            _ = v27
        }
    }
    v1 = ""
    var v33 float64 = 0.1
    v17 += v17
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v33
}
