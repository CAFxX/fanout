package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 bool; F1 uint32 }
func genFn0(p0 int16, p1 int64) string {
    v7 := []int{433, -1, 4294967295, 2147483647}
    p0 -= 0
    switch -899 {
    case -2147483648:
        var v8 int8 = -117
        _ = v8
    case 4294967295:
        for fi9 := 0; fi9 < 10; fi9++ {
            p0 = p0
            _ = fi9
        }
    default:
        fmt.Println(p0, p1)
    }
    p0 -= (-(p0))
    _ = p0
    _ = p1
    _ = v7
    return ""
}

func main() {
    var v1 float32 = (3.14159 * 1.5)
    v2 := make(map[int]uint32)
    v3 := make(map[string]string)
    var v4 bool = (255 == 128)
    v5 := make(map[uint32]int)
    var v6 string = "hello"
    v10 := GenS1{F0: v4, F1: 65535}
    if (0 != 75) {
        switch 127 {
        case 866:
            if (2147483647 < 1) {
                var v11 int8 = 52
                _ = v11
            } else {
                v12 := v5[65535]
                _ = v12
            }
        default:
            var v13 int = -2147483648
            _ = v13
        }
        fmt.Println(v6, v6, v4)
    } else {
        for fi14 := 0; fi14 < 10; fi14++ {
            v15 := []int{fi14}
            _ = fi14
            _ = v15
        }
        for fi16 := 0; fi16 < 10; fi16++ {
            var v17 uint = 18446744073709551615
            _ = fi16
            _ = v17
        }
        v1 += v1
    }
    v1 = v1
    var v18 string = ""
    v19 := GenS2{F0: (1 <= 2), F1: 1}
    if ((v18 + v18) >= "a") {
        var v20 string = "a"
        for fi21 := 0; fi21 < 13; fi21++ {
            switch (v6 + v20) {
            case "a":
                v20 = v18
            }
            for fi22 := 0; fi22 < 13; fi22++ {
                v20 = "\n\t"
                v23 := GenS1{F0: v4, F1: 131}
                if v4 {
                    v24 := GenS1{F0: v4, F1: 1}
                    var v25 uint64 = 0
                    switch v18 {
                    case "\n\t":
                        fmt.Println(v1, v20)
                    default:
                        v26 := []int{(fi21 % (((fi22 >> ((65535) & 63))) | 1)), (fi22 ^ fi21), fi21}
                        _ = v26
                    }
                    fmt.Println(v1, fi22, v18)
                    _ = v24
                    _ = v25
                }
                _ = fi22
                _ = v23
            }
            _ = fi21
        }
        var v27 int8 = -119
        _ = v20
        _ = v27
    } else {
        var v28 uint32 = 2
        _ = v28
    }
    var v29 uint32 = 2147483647
    var v30 float64 = 3.14159
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v18
    _ = v19
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
