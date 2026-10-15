package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 int32 }
type GenS2 struct { F0 int8; F1 bool }
func genFn0(p0 uint16) int8 {
    fmt.Println(p0)
    p0 = (p0 + (p0 / ((32767) | 1)))
    v4 := GenS1{F0: p0, F1: 1, F2: 0}
    p0 -= p0
    p0 = p0
    _ = p0
    _ = v4
    return -1
}
func genFn1() int {
    var v5 uint8 = 78
    v5 += (v5 & v5)
    v5 -= v5
    _ = v5
    return 42
}

func main() {
    v1 := []uint32{42}
    var v2 float64 = (float64(0.1) - 3.14159)
    v3 := []int{-2147483648, 2147483647, 0}
    fmt.Println(v2)
    if true {
        v2 = v2
        var v6 int8 = 127
        v2 = v2
        _ = v6
    }
    fmt.Println(v2)
    var v7 int8 = -128
    for fi8 := 0; fi8 < 7; fi8++ {
        v9 := GenS1{F0: 0, F1: -2147483648, F2: -1}
        switch (("foo bar" + "\n\t") + "") {
        case "\n\t":
            if false {
                switch "" {
                case "foo bar":
                    var v10 string = (((("a" + ("" + ("foo bar" + "\n\t"))) + (("a" + "") + (("hello" + "foo bar") + ("foo bar" + "hello")))) + "a") + "a")
                    _ = v10
                case "":
                    v7 = v7
                }
            } else {
                for fi11 := 0; fi11 < 3; fi11++ {
                    fmt.Println(v7, fi11)
                    _ = fi11
                }
                if (65535 <= 80) {
                    v7 += v7
                    for fi12 := 0; fi12 < 17; fi12++ {
                        var v13 float64 = (3.14159 * (float64(-2.25) * v2))
                        var v14 uint = 255
                        v13 -= float64(0.1)
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                    v15 := GenS1{F0: 2, F1: 1, F2: 1}
                    _ = v15
                }
            }
        case "a":
            if (924 < -838) {
                v2 += v2
                v2 -= 0.0
            } else {
                var v16 uint16 = 65535
                if len(v1) > 0 {
                    v1[0] = 440
                }
                _ = v16
            }
        case "foo bar":
            v2 += ((v2 + v2) - (v2 / (v2 / v2)))
        }
        fmt.Println(v7)
        _ = fi8
        _ = v9
    }
    var v17 float64 = 1.5
    if len(v3) > 0 {
        v3[0] = 2147483647
    }
    fmt.Println(v17, v2, v2)
    fmt.Println(v2, v2, v17)
    for fi18 := 0; fi18 < 6; fi18++ {
        fmt.Println(v2)
        if len(v3) > 0 {
            v3[0] = (fi18 >> ((0) & 63))
        }
        _ = fi18
    }
    var v19 int8 = 0
    for fi20 := 0; fi20 < 2; fi20++ {
        v19 -= (v7 * v7)
        _ = fi20
    }
    v21 := GenS1{F0: 862, F1: 42, F2: -671}
    if (-71 != genFn0(65535)) {
        switch "a" {
        case "a":
            for fi22 := 0; fi22 < 19; fi22++ {
                if len(v1) > 0 {
                    v1[0] = 874
                }
                var v23 float64 = v2
                _ = fi22
                _ = v23
            }
        case "hello":
            v3 = append(v3, 42)
        case "\n\t":
            v24 := GenS1{F0: 65535, F1: -1, F2: 1}
            _ = v24
        }
    } else {
        if false {
            v25 := GenS2{F0: 100, F1: false}
            _ = v25
        }
    }
    v7 = v7
    if len(v1) > 0 {
        v1[0] = 2147483647
    }
    v2 -= ((1.5 / float64(-2.25)) / float64(3.14159))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v19
    _ = v2
    _ = v21
    _ = v3
    _ = v7
}
