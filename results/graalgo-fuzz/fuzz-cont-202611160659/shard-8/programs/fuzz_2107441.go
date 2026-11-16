package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16; F1 int; F2 bool }
type GenS3 struct { F0 uint16; F1 uint32; F2 int16 }
func genFn0(p0 bool, p1 string, p2 int32) int {
    v4 := []int{-9223372036854775808, 253}
    var v5 float32 = 1e10
    p0 = p0
    p2 = p2
    fmt.Println(v5)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return -51
}
func genFn1(p0 uint64, p1 uint16, p2 string) uint32 {
    fmt.Println(p1)
    if (!true) {
        for fi6 := 0; fi6 < 1; fi6++ {
            if true {
                for fi7 := 0; fi7 < 10; fi7++ {
                    var v8 int8 = -1
                    var v9 uint = 42
                    v9 += v9
                    _ = fi7
                    _ = v8
                    _ = v9
                }
                p1 -= 671
            }
            _ = fi6
        }
        v10 := []int{-76, 32768, 1, -990}
        var v11 uint = 65535
        _ = v10
        _ = v11
    } else {
        v12 := []int{}
        _ = v12
    }
    switch p2 {
    case "a":
        var v13 string = p2
        _ = v13
    }
    fmt.Println(p2, p2)
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    return 4294967295
}

func main() {
    var v1 uint8 = 1
    var v2 int8 = -16
    v3 := make(map[int]int)
    for fi14 := 0; fi14 < 7; fi14++ {
        for fi15 := 0; fi15 < 14; fi15++ {
            var v16 string = ("a" + ("foo bar" + "\n\t"))
            _ = fi15
            _ = v16
        }
        v17 := GenS1{F0: fi14}
        v2 += 114
        _ = fi14
        _ = v17
    }
    v18 := GenS2{F0: 695, F1: genFn0(true, "hello", -2147483648), F2: (967 < 20)}
    fmt.Println(v2, v2)
    for fi19 := 0; fi19 < 11; fi19++ {
        fmt.Println(v1, v2, v2)
        v20 := GenS1{F0: genFn0((793 != 42), "hello", 2147483647)}
        _ = fi19
        _ = v20
    }
    v2 = v2
    for fi21 := 0; fi21 < 12; fi21++ {
        switch (fi21 + fi21) {
        case 127:
            v22 := GenS1{F0: 32767}
            _ = v22
        case 128:
            v23 := []int{fi21, (-(fi21)), (fi21 | (fi21 - -90)), (fi21 ^ (fi21 * -136))}
            _ = v23
        default:
            for fi24 := 0; fi24 < 5; fi24++ {
                v25 := GenS2{F0: 255, F1: (fi24 & (fi21 >> ((4294967295) & 63))), F2: true}
                var v26 bool = false
                _ = fi24
                _ = v25
                _ = v26
            }
        }
        v27 := GenS3{F0: 255, F1: uint32(v1), F2: 860}
        if true {
            switch fi21 {
            case 32768:
                if (false || false) {
                    v2 -= -1
                } else {
                    if (-868 > -1) {
                        var v28 int32 = -1
                        v28 -= v28
                        _ = v28
                    } else {
                        v29 := GenS1{F0: (fi21 >> ((202) & 63))}
                        _ = v29
                    }
                    var v30 uint = 0
                    v2 += v2
                    _ = v30
                }
            case 1:
                var v31 uint16 = 477
                _ = v31
            default:
                v2 = v2
            }
        }
        _ = fi21
        _ = v27
    }
    var v32 uint16 = 2
    if (1 == -229) {
        v2 += (v2 * v2)
    }
    var v33 int32 = 1
    var v34 int64 = 1
    var v35 uint16 = uint16(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
}
