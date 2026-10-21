package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 uint32; F1 bool; F2 uint32 }
type GenS3 struct { F0 int64; F1 uint32 }
func genFn0(p0 int16, p1 int32, p2 uint) uint8 {
    switch "foo bar" {
    case "foo bar":
        v6 := GenS1{F0: -550, F1: 42}
        _ = v6
    case "\n\t":
        v7 := GenS1{F0: (p0 << ((18446744073709551615) & 63)), F1: 243}
        _ = v7
    }
    for fi8 := 0; fi8 < 19; fi8++ {
        fmt.Println(p1)
        v9 := []int{(fi8 / ((-1) | 1)), fi8, fi8}
        _ = fi8
        _ = v9
    }
    v10 := []int{}
    v11 := GenS2{F0: 2147483648, F1: false, F2: 715}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return 108
}
func genFn1(p0 uint32, p1 int, p2 uint64) int16 {
    if true {
        var v12 uint64 = (p2 >> ((255) & 63))
        v13 := []int{(p1 % ((p1) | 1)), (-(p1)), (p1 | p1)}
        v14 := GenS1{F0: -1, F1: -9223372036854775808}
        _ = v12
        _ = v13
        _ = v14
    } else {
        var v15 uint32 = p0
        v16 := []int{(p1 - (p1 >> ((0) & 63))), p1}
        _ = v15
        _ = v16
    }
    p2 = p2
    p2 += (p2 - 2)
    p1 = (p1 << ((uint(p2)) & 63))
    switch (p1 ^ (p1 % ((852) | 1))) {
    case 128:
        v17 := []int{32767, -873}
        _ = v17
    case -99:
        var v18 string = ""
        _ = v18
    case 4294967295:
        v19 := GenS2{F0: (p0 - (p0 ^ 2147483648)), F1: (1 != 1), F2: p0}
        _ = v19
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}

func main() {
    v1 := make(map[string]int)
    var v2 uint16 = 2
    var v3 uint = 10
    v4 := make(map[string]int)
    var v5 uint16 = (^(v2))
    v20 := GenS3{F0: -1, F1: 65535}
    v21 := v4[""]
    if true {
        var v22 int16 = 42
        v23 := GenS2{F0: 2147483648, F1: (!(0.0 >= float64(1.5))), F2: 4294967295}
        _ = v22
        _ = v23
    }
    v5 -= (v5 >> (((^(v3))) & 63))
    v21 += (^(v21))
    switch "\n\t" {
    case "":
        for fi24 := 0; fi24 < 1; fi24++ {
            for fi25 := 0; fi25 < 5; fi25++ {
                var v26 int16 = 0
                v27 := GenS3{F0: -647, F1: 431}
                v4["a"] = -1
                v1[((("" + ("hello" + "a")) + ("foo bar" + ("" + "hello"))) + "\n\t")] = (v21 | -154)
                _ = fi25
                _ = v26
                _ = v27
            }
            v28 := GenS3{F0: 9223372036854775807, F1: 1}
            for fi29 := 0; fi29 < 5; fi29++ {
                v30 := GenS2{F0: 106, F1: false, F2: 2147483647}
                var v31 uint = v3
                fmt.Println(v31, v31, v5)
                v2 -= 32767
                _ = fi29
                _ = v30
                _ = v31
            }
            _ = fi24
            _ = v28
        }
    case "a":
        v3 -= 2
    case "foo bar":
        v21 -= (v21 | v21)
    }
    var v32 uint32 = 2147483648
    var v33 float32 = 1e10
    var v34 float64 = 1.5
    if (!true) {
        v5 += (v5 ^ (v5 * v5))
        if (1e10 != v34) {
            v5 -= 32767
        }
    } else {
        fmt.Println(v21)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
