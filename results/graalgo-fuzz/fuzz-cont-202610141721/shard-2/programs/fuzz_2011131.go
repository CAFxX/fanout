package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64; F1 int }
func genFn0(p0 uint16) uint64 {
    var v7 uint8 = 128
    fmt.Println(v7)
    _ = p0
    _ = v7
    return 4294967295
}
func genFn1(p0 uint64, p1 int64, p2 uint8) uint32 {
    v8 := []int{0, 767}
    var v9 uint8 = (p2 + 2)
    _ = len(v8)
    v9 = (p2 * (v9 / (((^(p2))) | 1)))
    if false {
        var v10 string = "\n\t"
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 65535
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]bool)
    v3 := make(map[string]int)
    v4 := make(map[uint32]int)
    var v5 uint64 = 4294967295
    var v6 uint32 = 2147483647
    v11 := []int{-2147483648, -624, 128, -883}
    for fi12 := 0; fi12 < 9; fi12++ {
        var v13 bool = false
        v5 -= (v5 ^ v5)
        _ = fi12
        _ = v13
    }
    v6 -= 65535
    if (128 != -828) {
        for fi14 := 0; fi14 < 6; fi14++ {
            v15 := GenS1{F0: 1, F1: 42}
            v5 += v5
            _ = fi14
            _ = v15
        }
        fmt.Println(v5)
    } else {
        var v16 int8 = -128
        _ = v16
    }
    fmt.Println(v6, v5)
    fmt.Println(v6)
    v5 += v5
    var v17 string = "a"
    v11 = append(v11, 32768)
    v18 := GenS2{F0: (v6 & v6)}
    for fi19 := 0; fi19 < 10; fi19++ {
        switch ((v17 + (((v17 + ("foo bar" + ((v17 + ("\n\t" + (v17 + ("foo bar" + v17)))) + "\n\t"))) + ((("foo bar" + v17) + ("a" + (("foo bar" + "foo bar") + v17))) + v17)) + ("foo bar" + ("hello" + "\n\t")))) + "\n\t") {
        case "":
            var v20 int64 = -9223372036854775808
            _ = v20
        case "foo bar":
            v21 := GenS2{F0: 74}
            _ = v21
        }
        var v22 float32 = 0.1
        _ = fi19
        _ = v22
    }
    v6 = (v6 / (((v6 << ((854) & 63))) | 1))
    v5 -= 477
    var v23 float64 = 0.0
    for fi24 := 0; fi24 < 15; fi24++ {
        var v25 uint32 = 687
        v26 := GenS3{F0: 1, F1: (fi24 ^ fi24)}
        if (!(true && (!false))) {
            var v27 int16 = -32768
            _ = v27
        }
        _ = fi24
        _ = v25
        _ = v26
    }
    var v28 uint8 = 255
    var v29 float32 = (0.1 + float32((float64(1e10) + ((float64((0.0 / (1.5 * -2.25))) - float64(3.14159)) * 3.14159))))
    fmt.Println(v6)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v11:", len(v11), v11)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
