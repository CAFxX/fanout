package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string; F2 int8 }
type GenS2 struct { F0 uint }
func genFn0() uint16 {
    switch -2147483648 {
    case 32768:
        var v7 uint64 = 42
        _ = v7
    case -100:
        var v8 uint32 = 42
        _ = v8
    }
    v9 := GenS1{F0: 975, F1: "hello", F2: -52}
    _ = v9
    return 255
}
func genFn1(p0 int8, p1 bool, p2 int) int32 {
    for fi10 := 0; fi10 < 10; fi10++ {
        fmt.Println(p0)
        p0 -= p0
        var v11 int = (^(p2))
        _ = fi10
        _ = v11
    }
    if p1 {
        v12 := GenS1{F0: 358, F1: "\n\t", F2: 100}
        for fi13 := 0; fi13 < 19; fi13++ {
            fmt.Println(p1, p1, fi13)
            if p1 {
                if p1 {
                    fmt.Println(p1)
                    p2 -= p2
                    for fi14 := 0; fi14 < 11; fi14++ {
                        var v15 uint16 = 0
                        _ = fi14
                        _ = v15
                    }
                }
                switch "foo bar" {
                case "foo bar":
                    for fi16 := 0; fi16 < 17; fi16++ {
                        fmt.Println(fi16, p2, p0)
                        fmt.Println(fi16, p0, p2)
                        if (477 <= 2147483647) {
                            var v17 uint8 = 255
                            p2 -= (fi13 << ((42) & 63))
                            v18 := []int{128, fi16, 0}
                            _ = v17
                            _ = v18
                        }
                        _ = fi16
                    }
                }
            }
            _ = fi13
        }
        _ = v12
    } else {
        v19 := []int{2147483647, p2, (^(p2))}
        _ = v19
    }
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    _ = p2
    return -873
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]int)
    v3 := make(map[int]bool)
    var v4 int = 2147483647
    v5 := make(map[string]uint32)
    v6 := make(map[uint32]string)
    var v20 int32 = 0
    for fi21 := 0; fi21 < 5; fi21++ {
        v20 = (^(v20))
        if (true && false) {
            for fi22 := 0; fi22 < 20; fi22++ {
                v4 = (v4 >> ((18446744073709551615) & 63))
                v23 := GenS2{F0: 893}
                v24 := GenS2{F0: 796}
                _ = fi22
                _ = v23
                _ = v24
            }
        }
        for fi25 := 0; fi25 < 14; fi25++ {
            fmt.Println(v20)
            v4 = -325
            _ = fi25
        }
        _ = fi21
    }
    var v26 float64 = ((1.5 / (1e10 / (float64(((1e10 + 0.1) / 1e10)) + (3.14159 / (1e10 * 0.1))))) - ((float64(((0.0 - (1.5 - 3.14159)) - 1e10)) + (1.5 - float64((((1.5 / 0.1) / float32(float64((3.14159 + float32(0.1))))) * -2.25)))) * -2.25))
    var v27 int64 = 0
    v4 -= v4
    v4 += (^(v4))
    v26 -= 0.0
    var v28 int8 = 127
    v27 -= 9223372036854775807
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v20
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
