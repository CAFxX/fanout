package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16 }
type GenS2 struct { F0 int; F1 int; F2 int }
func genFn0(p0 int, p1 int8) int8 {
    var v5 uint8 = 80
    v6 := GenS1{F0: 4294967295, F1: 32767}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return (p1 / (((-(p1))) | 1))
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]int)
    var v4 float32 = ((float32((1e10 + 3.14159)) / 1.0) / 1.0)
    v7 := []int{9223372036854775807, 4294967295, -244, -669}
    if (-116 <= genFn0(127, 0)) {
        v4 += v4
        v4 = 1e10
    } else {
        if len(v7) > 0 {
            v7[0] = 1
        }
    }
    if (42 != 715) {
        for fi8 := 0; fi8 < 1; fi8++ {
            switch "a" {
            case "a":
                if true {
                    var v9 bool = false
                    v10 := v3[2]
                    _ = v10
                    _ = v9
                }
            }
            switch "a" {
            case "":
                if len(v7) > 0 {
                    v7[0] = (fi8 - fi8)
                }
            }
            v4 += float32(((3.14159 + 3.14159) + 0.1))
            _ = fi8
        }
        v11 := GenS2{F0: 615, F1: -967, F2: 9223372036854775807}
        _ = v11
    }
    for fi12 := 0; fi12 < 8; fi12++ {
        v4 += v4
        v4 = v4
        v4 = 0.1
        _ = fi12
    }
    delete(v3, 41)
    var v13 int64 = 9223372036854775807
    if (837 < -2147483648) {
        v14 := v1[4294967295]
        _ = v14
    } else {
        var v15 string = "\n\t"
        fmt.Println(v4, v15)
        for fi16 := 0; fi16 < 7; fi16++ {
            if true {
                v17 := GenS2{F0: 2147483647, F1: (fi16 % (((fi16 >> ((0) & 63))) | 1)), F2: fi16}
                v13 = (v13 | -9223372036854775808)
                _ = v17
            }
            for fi18 := 0; fi18 < 6; fi18++ {
                var v19 int64 = v13
                _ = fi18
                _ = v19
            }
            var v20 int = 127
            _ = fi16
            _ = v20
        }
        _ = v15
    }
    fmt.Println(v4, v13)
    var v21 bool = (-32768 == 0)
    v13 = (v13 << ((980) & 63))
    fmt.Println(v13, v4, v21)
    fmt.Println(v13, v21)
    if ("\n\t" == "foo bar") {
        for fi22 := 0; fi22 < 19; fi22++ {
            v13 += v13
            if (!v21) {
                var v23 int8 = 100
                var v24 uint64 = 2
                _ = v23
                _ = v24
            } else {
                fmt.Println(fi22)
                switch "" {
                case "foo bar":
                    var v25 int64 = v13
                    _ = v25
                case "a":
                    delete(v2, 42)
                }
                for fi26 := 0; fi26 < 1; fi26++ {
                    v27 := GenS1{F0: 18446744073709551615, F1: 1000}
                    v28 := GenS2{F0: fi26, F1: (fi22 - fi22), F2: fi26}
                    _ = fi26
                    _ = v27
                    _ = v28
                }
            }
            _ = fi22
        }
        if len(v7) >= 2 {
            v7 = v7[1:len(v7)]
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v13
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v7
}
