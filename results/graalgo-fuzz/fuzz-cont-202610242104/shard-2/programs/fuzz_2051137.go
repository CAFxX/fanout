package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 uint32 }
type GenS2 struct { F0 bool; F1 int; F2 uint }
type GenS3 struct { F0 uint16 }
func genFn0(p0 int32, p1 uint16) int64 {
    v7 := []int{-139, 1}
    if false {
        p1 += p1
        switch 295 {
        case 546:
            fmt.Println(p0, p0)
        case 32768:
            var v8 string = "\n\t"
            _ = v8
        default:
            fmt.Println(p0, p0)
        }
        p0 = p0
    } else {
        fmt.Println(p1, p1)
        var v9 float64 = (0.0 / (1.5 + 0.1))
        v9 = (-2.25 * 3.14159)
        _ = v9
    }
    p1 -= 255
    _ = p0
    _ = p1
    _ = v7
    return -1
}
func genFn1(p0 int8, p1 int, p2 int8) int64 {
    var v10 uint8 = 42
    switch "" {
    case "":
        v11 := []int{int(v10), (p1 & p1), p1, (^(p1))}
        _ = v11
    case "\n\t":
        var v12 uint64 = 805
        _ = v12
    case "a":
        var v13 int = p1
        _ = v13
    }
    p1 += (p1 << ((651) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return 1
}

func main() {
    var v1 int8 = 0
    var v2 uint16 = 197
    v3 := make(map[string]string)
    var v4 uint32 = 4294967295
    var v5 int16 = 1
    v6 := []string{""}
    if len(v6) > 0 {
        v6[0] = ""
    }
    var v14 string = ""
    if len(v6) > 0 {
        v6[0] = (v14 + "\n\t")
    }
    delete(v3, (v14 + v14))
    if (!false) {
        for fi15 := 0; fi15 < 15; fi15++ {
            if false {
                v14 = v14
            } else {
                var v16 uint32 = (v4 / ((2) | 1))
                if ((v5 % ((v5) | 1)) != (v5 ^ v5)) {
                    v4 -= 65535
                    v5 = 1
                } else {
                    var v17 uint32 = (v16 << ((4294967295) & 63))
                    var v18 int64 = -9223372036854775808
                    _ = v17
                    _ = v18
                }
                _ = v16
            }
            _ = fi15
        }
    } else {
        v14 = v14
        for fi19 := 0; fi19 < 11; fi19++ {
            v20 := GenS1{F0: fi19, F1: 9223372036854775807, F2: v4}
            var v21 int64 = 0
            _ = fi19
            _ = v20
            _ = v21
        }
    }
    v1 = v1
    _ = len(v6)
    fmt.Println(v5)
    for fi22 := 0; fi22 < 8; fi22++ {
        v1 -= (v1 ^ (^(v1)))
        var v23 uint32 = v4
        _ = fi22
        _ = v23
    }
    for fi24 := 0; fi24 < 15; fi24++ {
        v2 = v2
        _ = fi24
    }
    var v25 int32 = int32(v2)
    switch "a" {
    case "":
        fmt.Println(v4, v5, v5)
    case "\n\t":
        switch 128 {
        case 0:
            v26 := GenS1{F0: -2147483648, F1: -9223372036854775808, F2: v4}
            _ = v26
        case 128:
            if (0 == 42) {
                if false {
                    fmt.Println(v2, v1, v1)
                } else {
                    v27 := v3[v14]
                    _ = v27
                }
                var v28 bool = false
                _ = v28
            } else {
                v29 := GenS1{F0: -255, F1: -1, F2: (^(v4))}
                _ = v29
            }
        default:
            if (v5 == 1000) {
                var v30 uint16 = v2
                _ = v30
            } else {
                v4 -= 4294967295
            }
        }
    case "hello":
        var v31 bool = false
        _ = v31
    }
    for fi32 := 0; fi32 < 7; fi32++ {
        var v33 uint8 = 52
        v1 += -31
        _ = fi32
        _ = v33
    }
    switch ((v14 + (v14 + "\n\t")) + v14) {
    case "a":
        v34 := GenS1{F0: -114, F1: genFn0((v25 | (v25 | (v25 | 2147483647))), (^(v2))), F2: (v4 | v4)}
        _ = v34
    case "\n\t":
        if true {
            v5 = v5
            var v35 int = -299
            _ = v35
        }
    case "foo bar":
        v36 := GenS1{F0: 127, F1: -591, F2: v4}
        _ = v36
    }
    if (genFn1(v1, -719, -97) <= int64(v4)) {
        var v37 int64 = 0
        fmt.Println(v1)
        if (0 > v2) {
            var v38 int16 = (-(v5))
            _ = v38
        } else {
            v39 := GenS2{F0: (941 <= 4294967295), F1: 4294967295, F2: 65535}
            v3[""] = v14
            v40 := GenS3{F0: v2}
            _ = v39
            _ = v40
        }
        _ = v37
    }
    v14 = v14
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
