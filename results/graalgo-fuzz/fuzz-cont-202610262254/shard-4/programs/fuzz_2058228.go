package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64; F1 int32; F2 uint16 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int64; F1 int8 }
func genFn0() uint32 {
    var v7 uint64 = 0
    v7 += v7
    switch 32767 {
    case 4294967295:
        v8 := GenS1{F0: 1}
        _ = v8
    }
    v9 := GenS1{F0: -1}
    _ = v7
    _ = v9
    return 65535
}
func genFn1(p0 uint8) int16 {
    p0 = p0
    for fi10 := 0; fi10 < 20; fi10++ {
        switch (fi10 << ((65535) & 63)) {
        case 9223372036854775807:
            for fi11 := 0; fi11 < 15; fi11++ {
                fmt.Println(p0)
                v12 := GenS1{F0: 0}
                p0 += (p0 >> ((42) & 63))
                _ = fi11
                _ = v12
            }
        case -9223372036854775808:
            p0 -= 128
        case -2147483648:
            if true {
                v13 := []int{-594}
                _ = v13
            } else {
                var v14 uint16 = 42
                var v15 float64 = -2.25
                if false {
                    v16 := GenS2{F0: 18446744073709551615, F1: 1, F2: (v14 | v14)}
                    _ = v16
                }
                _ = v14
                _ = v15
            }
        default:
            v17 := GenS3{F0: genFn0()}
            _ = v17
        }
        _ = fi10
    }
    p0 = (p0 * p0)
    for fi18 := 0; fi18 < 2; fi18++ {
        var v19 int32 = 2147483647
        fmt.Println(p0)
        _ = fi18
        _ = v19
    }
    _ = p0
    return 1
}

func main() {
    var v1 int8 = 42
    var v2 string = "hello"
    v3 := make(map[uint32]uint32)
    v4 := make(map[int]bool)
    var v5 uint16 = 0
    var v6 int8 = 0
    var v20 int16 = 32767
    for fi21 := 0; fi21 < 19; fi21++ {
        for fi22 := 0; fi22 < 16; fi22++ {
            var v23 int16 = v20
            _ = fi22
            _ = v23
        }
        for fi24 := 0; fi24 < 1; fi24++ {
            if false {
                fmt.Println(v1, v5)
            }
            _ = fi24
        }
        _ = fi21
    }
    var v25 int16 = 795
    v26 := []int{32768}
    v5 = (v5 % ((2) | 1))
    fmt.Println(v5, v5)
    fmt.Println(v1)
    v27 := GenS4{F0: 2147483647}
    v2 = v2
    v3[1] = 2147483648
    v28 := GenS5{F0: 0, F1: (v1 | v6)}
    var v29 bool = true
    var v30 uint64 = 1
    for fi31 := 0; fi31 < 5; fi31++ {
        v20 = v20
        if true {
            switch v2 {
            case "foo bar":
                var v32 float32 = (0.0 * float32(1.5))
                _ = v32
            }
        } else {
            fmt.Println(v30, fi31, v29)
        }
        _ = fi31
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
