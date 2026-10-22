package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64; F2 uint16 }
type GenS2 struct { F0 int8 }
func genFn0(p0 uint64) uint32 {
    if ((((((((-2.25 / float32(float64((1.5 - (0.0 + (1e10 + float32(1.5))))))) * -2.25) + 3.14159) - 1e10) - 1.5) + 0.0) == -2.25) || false) {
        if ((1 != 2147483647) && false) {
            if (!false) {
                var v7 uint32 = 42
                fmt.Println(v7, p0)
                var v8 int = 32767
                _ = v7
                _ = v8
            } else {
                var v9 int8 = 120
                v10 := []int{-1}
                v9 -= 120
                _ = v10
                _ = v9
            }
        }
        v11 := GenS1{F0: 255, F1: -895, F2: 2}
        _ = v11
    } else {
        var v12 int8 = 127
        switch 9223372036854775807 {
        case 2147483647:
            var v13 uint = 623
            _ = v13
        case 32768:
            v14 := []int{32767}
            _ = v14
        }
        if true {
            var v15 uint = 845
            v16 := []int{}
            _ = v15
            _ = v16
        } else {
            v17 := GenS1{F0: 206, F1: -9223372036854775808, F2: 32767}
            v12 = (-(v12))
            v18 := []int{32768, 32768, -369}
            _ = v17
            _ = v18
        }
        _ = v12
    }
    var v19 uint = 0
    p0 = p0
    var v20 uint16 = 741
    _ = p0
    _ = v19
    _ = v20
    return 321
}
func genFn1(p0 int16) uint8 {
    var v21 uint8 = 1
    fmt.Println(p0)
    v22 := GenS1{F0: 0, F1: -9223372036854775808, F2: uint16(v21)}
    _ = p0
    _ = v21
    _ = v22
    return 128
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint16 = 346
    v3 := make(map[uint32]string)
    var v4 int = -2147483648
    var v5 int16 = 42
    var v6 int8 = 127
    for fi23 := 0; fi23 < 1; fi23++ {
        if true {
            fmt.Println(v6, v6)
        } else {
            var v24 uint16 = (v2 / ((v2) | 1))
            if true {
                fmt.Println(v4)
            } else {
                var v25 string = ("a" + "")
                if true {
                    v26 := []int{(fi23 << ((656) & 63)), (v4 << ((4294967295) & 63)), (v4 ^ fi23), int(v24)}
                    _ = len(v3)
                    _ = v26
                } else {
                    var v27 float64 = 0.0
                    _ = v27
                }
                if (!false) {
                    if (127 != 128) {
                        v6 = (v6 * 127)
                    }
                    v24 = v2
                    fmt.Println(v2, fi23, v2)
                } else {
                    v5 -= (v5 | v5)
                    v28 := GenS2{F0: (v6 & v6)}
                    var v29 int = v4
                    _ = v28
                    _ = v29
                }
                _ = v25
            }
            var v30 uint8 = 42
            _ = v24
            _ = v30
        }
        var v31 uint = 755
        _ = fi23
        _ = v31
    }
    switch ("hello" + "foo bar") {
    case "a":
        v2 = v2
    case "":
        fmt.Println(v5)
    case "hello":
        var v32 uint64 = 794
        _ = v32
    }
    v33 := []int{(v4 << ((4294967295) & 63)), (^(v4)), v4, v4}
    for fi34 := 0; fi34 < 5; fi34++ {
        if false {
            for fi35 := 0; fi35 < 18; fi35++ {
                switch "" {
                case "foo bar":
                    v6 = (v6 / ((v6) | 1))
                }
                v2 = (v2 / ((v2) | 1))
                var v36 uint64 = 997
                _ = fi35
                _ = v36
            }
            v1[1] = (!(v6 <= (v6 / ((122) | 1))))
        }
        _ = len(v3)
        _ = fi34
    }
    var v37 int = (v4 / ((v4) | 1))
    var v38 uint = uint(v2)
    fmt.Println(v6, v37, v4)
    switch "hello" {
    case "hello":
        v37 = v37
    }
    v4 -= 32767
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
