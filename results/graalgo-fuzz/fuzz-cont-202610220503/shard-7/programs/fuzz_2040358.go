package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 int32 }
type GenS2 struct { F0 uint16; F1 uint32 }
func genFn0(p0 int, p1 uint) int {
    v7 := []int{(p0 * (^(p0))), 0}
    for fi8 := 0; fi8 < 12; fi8++ {
        for fi9 := 0; fi9 < 15; fi9++ {
            p1 += p1
            var v10 int16 = -506
            _ = fi9
            _ = v10
        }
        var v11 bool = true
        _ = fi8
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v7
    return (p0 & 1)
}
func genFn1() int16 {
    var v12 uint64 = 1
    if false {
        if false {
            v13 := GenS1{F0: 100, F1: 419, F2: -2147483648}
            v12 += v12
            _ = v13
        }
        for fi14 := 0; fi14 < 19; fi14++ {
            var v15 float32 = -2.25
            _ = fi14
            _ = v15
        }
        v12 -= v12
    } else {
        switch "a" {
        case "\n\t":
            v16 := GenS1{F0: 100, F1: 1, F2: 2147483647}
            _ = v16
        case "":
            for fi17 := 0; fi17 < 10; fi17++ {
                v12 = 0
                fmt.Println(fi17, fi17)
                switch "\n\t" {
                case "hello":
                    fmt.Println(fi17, fi17)
                case "\n\t":
                    switch "a" {
                    case "\n\t":
                        v12 -= (v12 % ((2) | 1))
                    case "a":
                        v18 := GenS1{F0: -128, F1: 2147483647, F2: 0}
                        _ = v18
                    }
                case "a":
                    v19 := []int{(fi17 << ((uint(v12)) & 63)), 100}
                    _ = v19
                }
                _ = fi17
            }
        }
        var v20 uint64 = (v12 >> ((193) & 63))
        _ = v20
    }
    v12 -= v12
    _ = v12
    return 1000
}

func main() {
    v1 := make(map[uint32]string)
    var v2 uint32 = 274
    v3 := make(map[int]int)
    v4 := []string{""}
    var v5 uint8 = 128
    var v6 string = ("\n\t" + "hello")
    if true {
        if true {
            var v21 int16 = 0
            _ = v21
        }
    } else {
        var v22 float32 = 1.5
        for fi23 := 0; fi23 < 16; fi23++ {
            v3[fi23] = fi23
            var v24 uint64 = 42
            for fi25 := 0; fi25 < 15; fi25++ {
                v2 += v2
                _ = fi25
            }
            _ = fi23
            _ = v24
        }
        var v26 int16 = 0
        _ = v22
        _ = v26
    }
    for fi27 := 0; fi27 < 12; fi27++ {
        v6 = (("foo bar" + ("a" + "a")) + v6)
        v3[fi27] = genFn0((fi27 / ((32767) | 1)), 65535)
        _ = fi27
    }
    for fi28 := 0; fi28 < 16; fi28++ {
        fmt.Println(fi28, fi28)
        _ = fi28
    }
    v29 := GenS2{F0: 32767, F1: 42}
    fmt.Println(v5)
    v5 += (v5 + 127)
    switch (("foo bar" + ((v6 + v6) + "foo bar")) + v6) {
    case "":
        if len(v4) > 0 {
            v4[0] = "a"
        }
    case "hello":
        v2 = v2
    default:
        var v30 int = genFn0(445, 2)
        _ = v30
    }
    v5 -= (v5 >> ((18446744073709551615) & 63))
    if (v6 < v6) {
        v31 := GenS2{F0: 32767, F1: 65535}
        _ = v31
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
