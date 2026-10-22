package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 uint8; F1 uint64 }
type GenS3 struct { F0 int16 }
func genFn0() uint8 {
    v7 := []int{-2147483648, -977, -441}
    switch -2147483648 {
    case -766:
        var v8 string = "foo bar"
        _ = v8
    case -1:
        if (!false) {
            if true {
                var v9 uint16 = 1
                v9 += (v9 ^ v9)
                _ = v9
            }
            if len(v7) > 0 {
                v7[0] = 32767
            }
        }
    }
    var v10 float32 = (3.14159 * -2.25)
    switch 32768 {
    case -216:
        fmt.Println(v10)
    }
    _ = v10
    _ = v7
    return 1
}

func main() {
    v1 := make(map[string]int)
    var v2 bool = (1 == 2)
    var v3 float64 = 0.0
    var v4 uint16 = 32767
    var v5 uint16 = (v4 - 1)
    var v6 bool = v2
    if ((0.0 * float32(0.1)) <= (0.1 + 0.0)) {
        for fi11 := 0; fi11 < 18; fi11++ {
            v12 := []int{fi11, -2147483648, 201}
            fmt.Println(v5, v6)
            for fi13 := 0; fi13 < 19; fi13++ {
                fmt.Println(fi11)
                if ((false || v2) && (!false)) {
                    var v14 uint16 = (v5 * v5)
                    _ = v14
                } else {
                    v2 = ((v3 * v3) <= v3)
                }
                if (!v6) {
                    switch ("a" + "a") {
                    case "":
                        var v15 int16 = 42
                        _ = v15
                    case "foo bar":
                        fmt.Println(v5, v5)
                    }
                }
                _ = fi13
            }
            _ = fi11
            _ = v12
        }
    } else {
        v3 = float64(1e10)
    }
    v16 := GenS1{F0: 2147483647, F1: v6}
    v17 := GenS2{F0: 128, F1: 18446744073709551615}
    v5 += v4
    v18 := GenS2{F0: 0, F1: 42}
    for fi19 := 0; fi19 < 3; fi19++ {
        v20 := GenS2{F0: 255, F1: 168}
        v5 += v5
        for fi21 := 0; fi21 < 9; fi21++ {
            v4 = (v4 << ((1) & 63))
            _ = fi21
        }
        _ = fi19
        _ = v20
    }
    var v22 bool = false
    fmt.Println(v5)
    for fi23 := 0; fi23 < 15; fi23++ {
        v24 := GenS2{F0: genFn0(), F1: 1}
        _ = fi23
        _ = v24
    }
    switch (((("" + "hello") + "foo bar") + "") + "a") {
    case "\n\t":
        v25 := GenS3{F0: -760}
        _ = v25
    }
    var v26 uint8 = 255
    v26 = 42
    fmt.Println(v6)
    switch 9223372036854775807 {
    case 889:
        var v27 float64 = 0.1
        _ = v27
    }
    switch "a" {
    case "foo bar":
        var v28 int32 = 1
        _ = v28
    case "":
        for fi29 := 0; fi29 < 8; fi29++ {
            v2 = (-934 < -710)
            for fi30 := 0; fi30 < 1; fi30++ {
                switch "hello" {
                case "a":
                    v26 -= genFn0()
                case "foo bar":
                    var v31 int8 = 112
                    _ = v31
                }
                v32 := []int{(fi29 ^ (fi30 | (^(fi30)))), -44, 2147483647, -460}
                _ = fi30
                _ = v32
            }
            _ = fi29
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
