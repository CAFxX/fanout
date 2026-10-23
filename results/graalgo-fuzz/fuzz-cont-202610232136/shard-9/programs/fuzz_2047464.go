package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint8; F1 bool }
func genFn0(p0 int64) uint {
    var v7 uint16 = 2
    p0 = -266
    v8 := []int{42}
    _ = p0
    _ = v7
    _ = v8
    return 2
}

func main() {
    var v1 int8 = 0
    var v2 float32 = (1e10 / (((-2.25 - 1.5) + 0.1) - (0.1 * float32(3.14159))))
    v3 := make(map[int]bool)
    v4 := []string{}
    var v5 int32 = -574
    var v6 float32 = 0.1
    v6 -= -2.25
    var v9 uint32 = 42
    v1 = v1
    v5 -= 42
    v10 := GenS1{F0: 9223372036854775807, F1: 42, F2: 32767}
    v9 += v9
    for fi11 := 0; fi11 < 20; fi11++ {
        for fi12 := 0; fi12 < 5; fi12++ {
            var v13 uint64 = 0
            v2 += (0.0 / v2)
            fmt.Println(v13, v9)
            _ = fi12
            _ = v13
        }
        var v14 uint8 = 0
        if true {
            v1 += int8(fi11)
            v15 := GenS2{F0: (("hello" + "") + "\n\t")}
            v16 := GenS1{F0: 0, F1: 1000, F2: 602}
            _ = v15
            _ = v16
        }
        _ = fi11
        _ = v14
    }
    var v17 int = -1
    v18 := GenS1{F0: 120, F1: 1000, F2: 343}
    switch v17 {
    case 42:
        switch ("\n\t" + "foo bar") {
        case "hello":
            fmt.Println(v9, v5, v6)
        default:
            var v19 uint16 = 32767
            _ = v19
        }
    case 9223372036854775807:
        if true {
            switch (v17 ^ (v17 | v17)) {
            case -623:
                var v20 uint16 = 2
                _ = v20
            case 42:
                var v21 bool = false
                _ = v21
            case -2147483648:
                for fi22 := 0; fi22 < 3; fi22++ {
                    if (false || (true || (true && true))) {
                        v1 = v1
                    }
                    for fi23 := 0; fi23 < 8; fi23++ {
                        v9 -= v9
                        v24 := GenS3{F0: 18446744073709551615}
                        v25 := GenS4{F0: 128, F1: (1 > 0)}
                        _ = fi23
                        _ = v24
                        _ = v25
                    }
                    _ = fi22
                }
            default:
                v5 -= 2147483647
            }
            v17 = (v17 | -726)
        }
    case 280:
        switch 1 {
        case 32768:
            var v26 uint32 = (v9 + (v9 + 96))
            _ = v26
        }
    }
    for fi27 := 0; fi27 < 3; fi27++ {
        switch v17 {
        case -1:
            fmt.Println(v1, v9, v6)
        default:
            v2 = 0.0
        }
        var v28 uint64 = 4294967295
        _ = fi27
        _ = v28
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
