package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 uint64; F1 uint16 }
type GenS3 struct { F0 int64; F1 bool; F2 int8 }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]bool)
    var v3 uint32 = 261
    var v4 uint8 = 18
    v5 := make(map[uint32]string)
    var v6 uint16 = 2
    v7 := GenS1{F0: -1, F1: -1}
    v8 := []int{-183, -526}
    if len(v8) > 0 {
        v8[0] = 2147483647
    }
    v6 = uint16(v3)
    if true {
        var v9 float64 = float64(1e10)
        if (false && false) {
            for fi10 := 0; fi10 < 9; fi10++ {
                v1[-443] = ((("foo bar" + ("hello" + "")) + (("hello" + "foo bar") + "\n\t")) > "")
                var v11 float64 = v9
                var v12 float64 = (v11 / float64(((1.5 - 1e10) + (3.14159 - 0.1))))
                _ = fi10
                _ = v11
                _ = v12
            }
        }
        v4 -= 2
        _ = v9
    }
    v13 := GenS2{F0: 1, F1: (^(v6))}
    v3 = v3
    fmt.Println(v4)
    var v14 uint = 433
    v15 := GenS3{F0: 1, F1: true, F2: 42}
    if false {
        for fi16 := 0; fi16 < 16; fi16++ {
            delete(v5, v3)
            var v17 uint16 = (^(v6))
            v18 := GenS1{F0: 73, F1: 535}
            _ = fi16
            _ = v17
            _ = v18
        }
        if true {
            var v19 string = "a"
            _ = v19
        }
    } else {
        for fi20 := 0; fi20 < 7; fi20++ {
            for fi21 := 0; fi21 < 12; fi21++ {
                if (uint64(v14) < 620) {
                    v6 -= (v6 | v6)
                }
                var v22 bool = false
                _ = len(v1)
                _ = fi21
                _ = v22
            }
            _ = fi20
        }
        var v23 float64 = 1.5
        var v24 int = 32768
        _ = v23
        _ = v24
    }
    var v25 int = 732
    if true {
        v14 -= (v14 << ((2) & 63))
        var v26 float64 = (((((1.5 * (0.1 * -2.25)) - float64(1e10)) / (1.5 / 1.0)) + (0.1 * float64(1e10))) / 3.14159)
        _ = v26
    } else {
        switch (v25 * v25) {
        case -1:
            fmt.Println(v6)
        case 9223372036854775807:
            v4 = uint8(v3)
        }
    }
    var v27 bool = ((("a" + "hello") + "") != "\n\t")
    var v28 int16 = -835
    var v29 float64 = 3.14159
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v25
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
