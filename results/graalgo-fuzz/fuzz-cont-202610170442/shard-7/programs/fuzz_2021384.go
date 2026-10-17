package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 uint32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 uint; F1 uint32 }
func genFn0(p0 int32) uint {
    p0 += p0
    for fi4 := 0; fi4 < 19; fi4++ {
        p0 += p0
        _ = fi4
    }
    _ = p0
    return 1
}
func genFn1(p0 int, p1 int64) uint64 {
    if ((0.0 + 0.0) != 0.0) {
        for fi5 := 0; fi5 < 10; fi5++ {
            v6 := GenS1{F0: fi5, F1: fi5, F2: 42}
            _ = fi5
            _ = v6
        }
        for fi7 := 0; fi7 < 9; fi7++ {
            p0 += (-(fi7))
            v8 := []int{(p0 / ((4294967295) | 1)), (p0 % (((p0 - p0)) | 1)), 32768, (^(p0))}
            for fi9 := 0; fi9 < 13; fi9++ {
                if (!true) {
                    if len(v8) > 0 {
                        v8[0] = (^(fi7))
                    }
                    if (((fi9 ^ fi9) <= fi7) || (0 == -810)) {
                        var v10 bool = (!(-2.25 < 0.1))
                        _ = v10
                    }
                    if (86 <= 151) {
                        switch 0 {
                        case 474:
                            p1 += (p1 >> ((4294967295) & 63))
                        }
                        var v11 int = fi7
                        _ = v11
                    } else {
                        v12 := GenS1{F0: (fi9 ^ fi9), F1: -9223372036854775808, F2: 780}
                        _ = v12
                    }
                    if (2147483647 <= 42) {
                        p1 -= (^(p1))
                    }
                    fmt.Println(fi7)
                }
                p0 = fi9
                v8 = append(v8, fi9)
                v13 := GenS2{F0: 0}
                p0 = (fi7 >> ((0) & 63))
                _ = fi9
                _ = v13
            }
            _ = fi7
            _ = v8
        }
        fmt.Println(p0, p0)
    } else {
        for fi14 := 0; fi14 < 3; fi14++ {
            fmt.Println(p1, p1)
            switch "foo bar" {
            case "foo bar":
                p1 = p1
            case "a":
                fmt.Println(fi14)
            }
            _ = fi14
        }
        v15 := []int{}
        _ = v15
    }
    var v16 float64 = 3.14159
    _ = p0
    _ = p1
    _ = v16
    return 4294967295
}

func main() {
    v1 := make(map[int]int)
    var v2 uint8 = 15
    var v3 float32 = -2.25
    v3 -= v3
    switch 481 {
    case 4294967295:
        v2 += 178
    case 590:
        var v17 int8 = 100
        _ = v17
    }
    v18 := GenS1{F0: 1, F1: -1, F2: 4294967295}
    var v19 int32 = 42
    fmt.Println(v2, v19, v19)
    v20 := []int{}
    v2 = (v2 % ((16) | 1))
    if ((-53 == 1) && true) {
        if false {
            var v21 int8 = 42
            v21 = -5
            for fi22 := 0; fi22 < 11; fi22++ {
                var v23 float32 = -2.25
                _ = fi22
                _ = v23
            }
            _ = v21
        }
        v19 -= -432
        for fi24 := 0; fi24 < 13; fi24++ {
            switch "" {
            case "a":
                v19 -= 587
            case "hello":
                for fi25 := 0; fi25 < 14; fi25++ {
                    if true {
                        fmt.Println(v3, v3)
                        fmt.Println(fi25)
                        v19 = (v19 + v19)
                    } else {
                        v26 := GenS1{F0: fi25, F1: 32768, F2: 1}
                        _ = v26
                    }
                    var v27 int32 = (v19 ^ -1)
                    fmt.Println(v27, v27, v2)
                    _ = fi25
                    _ = v27
                }
            default:
                v3 += (1e10 - 3.14159)
            }
            if (983 <= 814) {
                _ = len(v20)
                switch 884 {
                case 32768:
                    for fi28 := 0; fi28 < 5; fi28++ {
                        var v29 int32 = (-(v19))
                        _ = fi28
                        _ = v29
                    }
                case -9223372036854775808:
                    var v30 float32 = (float32((float64(((v3 * v3) * float32(-2.25))) * float64((v3 / (0.1 + v3))))) - 1e10)
                    _ = v30
                }
            } else {
                v31 := GenS3{F0: 2, F1: 2147483648}
                _ = v31
            }
            v19 += v19
            _ = fi24
        }
    } else {
        v32 := GenS2{F0: 18446744073709551615}
        _ = v32
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
