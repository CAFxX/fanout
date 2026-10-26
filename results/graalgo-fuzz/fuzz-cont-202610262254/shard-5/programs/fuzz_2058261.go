package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 int64; F2 int }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint8 = 128
    v3 := make(map[int]string)
    v2 += (v2 >> ((2) & 63))
    v4 := []int{-1, 128, 0}
    v2 = v2
    if true {
        v2 += (^(v2))
    }
    fmt.Println(v2)
    if (true && (((("foo bar" + "a") + ("a" + "")) + "") > "hello")) {
        for fi5 := 0; fi5 < 10; fi5++ {
            for fi6 := 0; fi6 < 6; fi6++ {
                var v7 uint32 = 2147483648
                v2 = 220
                _ = fi6
                _ = v7
            }
            _ = fi5
        }
        if true {
            var v8 uint = 921
            var v9 uint16 = 342
            switch (("\n\t" + "foo bar") + "a") {
            case "\n\t":
                v8 = (^(v8))
            }
            _ = v8
            _ = v9
        } else {
            var v10 int16 = 0
            v11 := v3[2147483647]
            _ = v10
            _ = v11
        }
        v12 := GenS1{F0: 2147483647}
        _ = v12
    }
    switch "a" {
    case "a":
        v2 = (^(v2))
    }
    var v13 bool = true
    for fi14 := 0; fi14 < 19; fi14++ {
        switch "hello" {
        case "hello":
            var v15 int32 = 172
            _ = v15
        case "foo bar":
            v13 = v13
        default:
            for fi16 := 0; fi16 < 12; fi16++ {
                v2 += v2
                _ = fi16
            }
        }
        switch "hello" {
        case "a":
            for fi17 := 0; fi17 < 5; fi17++ {
                var v18 uint64 = 18446744073709551615
                v18 -= (v18 << ((429) & 63))
                _ = fi17
                _ = v18
            }
        case "foo bar":
            v2 = v2
        case "\n\t":
            switch -609 {
            case 32767:
                fmt.Println(fi14, v2)
            }
        default:
            var v19 uint8 = 128
            _ = v19
        }
        var v20 bool = (806 <= -9223372036854775808)
        _ = fi14
        _ = v20
    }
    fmt.Println(v2, v13)
    for fi21 := 0; fi21 < 15; fi21++ {
        for fi22 := 0; fi22 < 17; fi22++ {
            v13 = (42 < -1)
            var v23 float64 = (float64((float32((1e10 + (1e10 / float64(1e10)))) + 3.14159)) - float64((1.5 - (float32((0.1 / (-2.25 * float64(1.5)))) + float32((0.0 * (((float64(3.14159) + -2.25) - 0.0) + 0.0)))))))
            _ = fi22
            _ = v23
        }
        var v24 int16 = int16(v2)
        var v25 bool = v13
        _ = fi21
        _ = v24
        _ = v25
    }
    v26 := GenS2{F0: 893, F1: 9223372036854775807, F2: 4294967295}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v13
    _ = v2
    _ = v26
    _ = v3
    _ = v4
}
