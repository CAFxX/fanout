package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int; F1 uint16; F2 bool }
type GenS3 struct { F0 int16; F1 uint; F2 uint16 }
func genFn0(p0 uint) uint32 {
    v6 := GenS1{F0: true}
    if (!(1 == -112)) {
        if true {
            p0 = 655
        } else {
            p0 -= (p0 << ((117) & 63))
        }
        p0 -= (p0 * p0)
        for fi7 := 0; fi7 < 6; fi7++ {
            if true {
                p0 = p0
                switch (fi7 - -366) {
                case 128:
                    var v8 uint32 = 65535
                    _ = v8
                }
            }
            p0 = 4294967295
            _ = fi7
        }
    } else {
        v9 := []int{-9223372036854775808, 0, 32768, 9223372036854775807}
        p0 = p0
        v10 := GenS1{F0: true}
        _ = v10
        _ = v9
    }
    v11 := []int{-345, 128}
    p0 = (p0 & p0)
    _ = p0
    _ = v11
    _ = v6
    return 114
}
func genFn1(p0 int64, p1 uint) uint16 {
    fmt.Println(p1, p0)
    for fi12 := 0; fi12 < 9; fi12++ {
        v13 := GenS1{F0: (-2.25 != ((((3.14159 * 0.1) + 1e10) + 0.1) - (float64(0.1) * 1.5)))}
        _ = fi12
        _ = v13
    }
    if ("hello" > ("hello" + ("a" + "a"))) {
        p0 = p0
        fmt.Println(p0, p0)
        v14 := []int{524, 2147483647}
        _ = v14
    }
    p0 = (p0 ^ p0)
    var v15 uint = p1
    _ = p0
    _ = p1
    _ = v15
    return 2
}

func main() {
    var v1 uint = 1
    var v2 int64 = 765
    var v3 float32 = (-2.25 / (-2.25 / 1e10))
    var v4 uint8 = 127
    var v5 int32 = -2147483648
    v3 = (-2.25 * (float32(0.0) + 1e10))
    v16 := GenS2{F0: -9223372036854775808, F1: 957, F2: false}
    v4 += (v4 / ((v4) | 1))
    var v17 uint8 = v4
    v18 := []int{161}
    for fi19 := 0; fi19 < 20; fi19++ {
        _ = len(v18)
        v20 := GenS3{F0: -32768, F1: v1, F2: 2}
        _ = fi19
        _ = v20
    }
    v18 = append(v18, 42)
    fmt.Println(v1, v5, v1)
    v21 := GenS2{F0: 2147483647, F1: 1, F2: (32767 < 42)}
    if len(v18) >= 2 {
        v18 = v18[1:len(v18)]
    }
    fmt.Println(v2, v5)
    fmt.Println(v5, v17, v1)
    v5 = 2147483647
    v1 = v1
    for fi22 := 0; fi22 < 2; fi22++ {
        for fi23 := 0; fi23 < 15; fi23++ {
            if len(v18) > 0 {
                v18[0] = fi23
            }
            _ = fi23
        }
        if (true || false) {
            fmt.Println(v4, v3, v1)
            fmt.Println(v2, v4)
        }
        var v24 int = 9223372036854775807
        _ = fi22
        _ = v24
    }
    var v25 uint32 = genFn0(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
