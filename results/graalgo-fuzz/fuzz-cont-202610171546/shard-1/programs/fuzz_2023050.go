package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 int; F1 int16; F2 uint64 }
func genFn0() int {
    switch ("\n\t" + "foo bar") {
    case "a":
        switch -409 {
        case -1:
            var v7 uint8 = 175
            _ = v7
        case 2147483647:
            if false {
                v8 := []int{}
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                _ = v8
            }
        }
    default:
        if (true || (254 >= 48)) {
            v9 := []int{1}
            v10 := GenS1{F0: 42, F1: 42}
            if true {
                for fi11 := 0; fi11 < 11; fi11++ {
                    var v12 bool = ("hello" < "foo bar")
                    _ = fi11
                    _ = v12
                }
                var v13 uint32 = 4294967295
                fmt.Println(v13)
                _ = v13
            } else {
                _ = len(v9)
            }
            _ = v10
            _ = v9
        }
    }
    v14 := []int{-2147483648, 215, 395, 32768}
    _ = len(v14)
    for fi15 := 0; fi15 < 3; fi15++ {
        v16 := GenS2{F0: 42, F1: ("\n\t" + "foo bar")}
        _ = len(v14)
        _ = fi15
        _ = v16
    }
    var v17 float64 = (1e10 - 0.0)
    _ = v14
    _ = v17
    return -1
}
func genFn1(p0 uint64, p1 uint32, p2 int64) int16 {
    switch "" {
    case "hello":
        switch 887 {
        case 128:
            p1 -= (p1 ^ (p1 | (p1 << ((1) & 63))))
        case 32767:
            if false {
                p0 = (p0 % ((p0) | 1))
            } else {
                fmt.Println(p2)
            }
        default:
            var v18 int = 9223372036854775807
            _ = v18
        }
    case "":
        p0 = (^(p0))
    default:
        v19 := GenS1{F0: -587, F1: -1}
        _ = v19
    }
    var v20 float32 = float32(0.1)
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v20
    return 0
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []uint32{138, 326}
    var v3 uint32 = 761
    v4 := []uint32{(^(v3)), (v3 & v3)}
    v5 := []uint32{(v3 << ((1) & 63))}
    v6 := make(map[int]int)
    var v21 int = 32768
    var v22 uint32 = 59
    _ = len(v1)
    switch "a" {
    case "":
        v22 = v22
    case "hello":
        v22 = v22
    case "a":
        var v23 uint16 = 438
        _ = v23
    }
    switch v21 {
    case 32767:
        v24 := GenS2{F0: 1, F1: ""}
        _ = v24
    case 1:
        switch (v21 * -637) {
        case 0:
            v22 += v22
        }
    }
    fmt.Println(v22, v3, v3)
    var v25 int16 = 1000
    fmt.Println(v25, v22)
    v22 = v22
    var v26 string = "foo bar"
    var v27 int16 = v25
    _ = len(v5)
    v21 = 32767
    switch 9223372036854775807 {
    case 641:
        var v28 string = (v26 + (v26 + v26))
        _ = v28
    case 128:
        var v29 float64 = (1e10 - 1e10)
        _ = v29
    }
    v3 = (^(v22))
    fmt.Println(v3, v26, v26)
    v30 := GenS3{F0: (v21 + int(v27)), F1: 42, F2: 27}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
