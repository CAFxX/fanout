package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 uint16; F1 int64 }
type GenS3 struct { F0 bool; F1 int8; F2 uint8 }
func genFn0() uint64 {
    v6 := []int{-1, 1, -657, -139}
    var v7 int8 = 100
    v7 -= (v7 / (((-(v7))) | 1))
    for fi8 := 0; fi8 < 9; fi8++ {
        if len(v6) > 0 {
            v6[0] = fi8
        }
        var v9 string = "hello"
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 10; fi10++ {
        var v11 uint64 = 18446744073709551615
        _ = fi10
        _ = v11
    }
    _ = v6
    _ = v7
    return 0
}
func genFn1(p0 uint) uint8 {
    v12 := []int{32767, 251, 9223372036854775807}
    v13 := GenS1{F0: 2, F1: 1}
    p0 += p0
    var v14 bool = (1 <= -1)
    for fi15 := 0; fi15 < 3; fi15++ {
        switch "\n\t" {
        case "foo bar":
            for fi16 := 0; fi16 < 17; fi16++ {
                var v17 uint = (p0 | (^(p0)))
                v17 = (v17 | v17)
                fmt.Println(fi16, fi16, p0)
                _ = fi16
                _ = v17
            }
        case "a":
            if (882 < 255) {
                var v18 bool = v14
                if (1 == 0) {
                    var v19 uint64 = 1
                    _ = v19
                }
                v14 = (-260 == -9223372036854775808)
                _ = v18
            } else {
                v20 := GenS2{F0: 0, F1: -634}
                _ = v20
            }
        default:
            v14 = (false && (32767 >= -416))
        }
        var v21 uint64 = 4294967295
        v22 := GenS3{F0: false, F1: 100, F2: 127}
        _ = fi15
        _ = v21
        _ = v22
    }
    _ = p0
    _ = v12
    _ = v13
    _ = v14
    return 42
}

func main() {
    var v1 float32 = 0.1
    v2 := make(map[uint32]int)
    var v3 float32 = 0.1
    var v4 int = 42
    v5 := make(map[uint32]string)
    v3 += ((v1 / v1) + (((0.0 + v3) - (v1 * 1.5)) / (1e10 / (float32(0.0) * v3))))
    var v23 int64 = 1
    var v24 int16 = -32768
    for fi25 := 0; fi25 < 3; fi25++ {
        switch ("a" + (("" + "hello") + "")) {
        case "foo bar":
            fmt.Println(v23)
        default:
            var v26 uint64 = 1
            _ = v26
        }
        v3 = (v3 + v3)
        delete(v5, 4294967295)
        _ = fi25
    }
    v4 += v4
    for fi27 := 0; fi27 < 19; fi27++ {
        v28 := v2[2]
        var v29 int32 = 484
        var v30 int32 = (-(v29))
        _ = fi27
        _ = v28
        _ = v29
        _ = v30
    }
    fmt.Println(v24)
    if false {
        v23 = v23
        v1 = -2.25
        switch ("a" + ("a" + "a")) {
        case "foo bar":
            v31 := []int{4294967295, (v4 % ((v4) | 1)), -895, v4}
            _ = v31
        case "\n\t":
            for fi32 := 0; fi32 < 7; fi32++ {
                v3 += (v1 / v3)
                if (false || false) {
                    if ("\n\t" > "hello") {
                        v4 = (v4 % (((fi32 % ((v4) | 1))) | 1))
                        var v33 float64 = 1e10
                        _ = v33
                    }
                }
                v3 = float32(0.1)
                _ = fi32
            }
        }
    } else {
        v2[2147483647] = (v4 ^ (v4 * v4))
    }
    v34 := []int{(v4 & -133)}
    for fi35 := 0; fi35 < 20; fi35++ {
        v24 = (v24 & (^(v24)))
        for fi36 := 0; fi36 < 4; fi36++ {
            var v37 uint64 = 4294967295
            _ = fi36
            _ = v37
        }
        switch ((("hello" + ("" + "a")) + ("hello" + (("hello" + "a") + "a"))) + "") {
        case "foo bar":
            v1 += ((((v1 - (0.1 - (1.5 / v3))) * v1) / v3) / (3.14159 + (3.14159 / float32(1e10))))
        }
        _ = fi35
    }
    delete(v5, 959)
    var v38 string = (("\n\t" + "a") + "hello")
    v4 += 4294967295
    var v39 uint64 = 600
    v38 = ("" + "\n\t")
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v34
    _ = v38
    _ = v39
    _ = v4
    _ = v5
}
