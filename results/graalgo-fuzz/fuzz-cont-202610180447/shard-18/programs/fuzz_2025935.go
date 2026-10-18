package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint; F1 uint8; F2 int16 }
func genFn0(p0 int64, p1 int8) uint {
    if false {
        for fi4 := 0; fi4 < 5; fi4++ {
            fmt.Println(fi4)
            _ = fi4
        }
        if (1 >= 1) {
            switch "foo bar" {
            case "":
                if (1 >= 60) {
                    p1 = (^(p1))
                }
            case "foo bar":
                for fi5 := 0; fi5 < 18; fi5++ {
                    fmt.Println(fi5, fi5)
                    for fi6 := 0; fi6 < 17; fi6++ {
                        switch (fi5 / ((fi6) | 1)) {
                        case -1:
                            v7 := GenS1{F0: ((false || false) || ("\n\t" >= "")), F1: fi6}
                            _ = v7
                        default:
                            p0 = p0
                        }
                        v8 := []int{42, fi5}
                        _ = fi6
                        _ = v8
                    }
                    if true {
                        var v9 uint = 670
                        _ = v9
                    }
                    _ = fi5
                }
            case "a":
                v10 := []int{-9223372036854775808, -2147483648, 329, -142}
                _ = v10
            }
        } else {
            var v11 int64 = 0
            _ = v11
        }
        var v12 float32 = 1.5
        _ = v12
    }
    if false {
        v13 := []int{9223372036854775807, -20, 0}
        _ = v13
    } else {
        switch "foo bar" {
        case "a":
            p0 -= (p0 * p0)
        }
    }
    _ = p0
    _ = p1
    return 65535
}
func genFn1(p0 int, p1 int8) int8 {
    if ((-(p1)) < int8(p0)) {
        v14 := GenS1{F0: ((0.0 / -2.25) >= 3.14159), F1: p0}
        v15 := []int{}
        p0 = 32768
        _ = v14
        _ = v15
    } else {
        var v16 uint32 = 2147483647
        _ = v16
    }
    var v17 uint = 65535
    v18 := GenS2{F0: (p1 | (p1 % ((p1) | 1)))}
    v19 := []int{p0}
    v19 = append(v19, (p0 ^ (-(p0))))
    _ = p0
    _ = p1
    _ = v17
    _ = v18
    _ = v19
    return p1
}

func main() {
    v1 := []int{127, 42}
    var v2 int = -9223372036854775808
    v3 := make(map[string]uint32)
    var v20 bool = false
    v1 = append(v1, v2)
    v20 = (0 <= 0)
    var v21 uint32 = 65535
    v21 += 813
    var v22 float64 = (((3.14159 - (0.0 + 0.1)) * 3.14159) - 1e10)
    fmt.Println(v2, v20, v22)
    var v23 uint = 2
    if (2147483647 >= (v2 & v2)) {
        fmt.Println(v2, v21)
    }
    for fi24 := 0; fi24 < 19; fi24++ {
        v22 = float64(0.0)
        if (v20 || (false || v20)) {
            if v20 {
                v25 := GenS2{F0: 42}
                v3[""] = 843
                var v26 uint8 = 42
                _ = v25
                _ = v26
            } else {
                var v27 int8 = 1
                v28 := GenS3{F0: 0, F1: 42, F2: 1000}
                _ = v27
                _ = v28
            }
            v29 := GenS1{F0: v20, F1: (fi24 ^ 32767)}
            _ = v29
        }
        v3[("hello" + ("foo bar" + "hello"))] = v21
        _ = fi24
    }
    var v30 uint = 4294967295
    v2 -= 127
    for fi31 := 0; fi31 < 10; fi31++ {
        var v32 uint = v23
        fmt.Println(v30, v21, v23)
        _ = fi31
        _ = v32
    }
    fmt.Println(v23, v20, v20)
    v33 := GenS3{F0: 2, F1: 255, F2: 0}
    v3[("hello" + ("\n\t" + "hello"))] = uint32(v30)
    v30 += (v30 * 1)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v30
    _ = v33
}
