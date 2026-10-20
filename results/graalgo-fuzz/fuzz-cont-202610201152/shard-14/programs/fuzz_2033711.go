package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 uint64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int; F1 bool; F2 uint }
type GenS4 struct { F0 int64; F1 int16 }

func main() {
    var v1 int32 = 1
    var v2 uint32 = 4294967295
    v3 := []int{}
    switch 127 {
    case 4294967295:
        var v4 uint32 = 1
        _ = v4
    case -1:
        v1 -= 0
    }
    v1 -= (v1 / (((-(v1))) | 1))
    var v5 float32 = 1e10
    if (-2.25 < 3.14159) {
        fmt.Println(v5, v2, v1)
        switch "\n\t" {
        case "hello":
            v5 = 0.0
        case "":
            v1 += v1
        case "a":
            v1 -= (v1 | 0)
        }
        for fi6 := 0; fi6 < 16; fi6++ {
            v2 -= v2
            v7 := GenS1{F0: (v1 + v1), F1: 932, F2: 1}
            if (v5 < 1.5) {
                switch (("a" + ("hello" + "")) + ("hello" + "a")) {
                case "hello":
                    if false {
                        fmt.Println(v5)
                    }
                }
            } else {
                var v8 int32 = v1
                var v9 uint32 = (^(v2))
                _ = v8
                _ = v9
            }
            _ = fi6
            _ = v7
        }
    }
    var v10 uint16 = 255
    v11 := GenS2{F0: v10}
    if true {
        for fi12 := 0; fi12 < 1; fi12++ {
            if len(v3) > 0 {
                v3[0] = fi12
            }
            _ = fi12
        }
    } else {
        if false {
            switch "" {
            case "foo bar":
                v5 = (1.5 + ((v5 / (float32(((1e10 + 3.14159) * 3.14159)) / v5)) + 1e10))
            }
            var v13 float32 = v5
            fmt.Println(v5)
            _ = v13
        }
        if true {
            for fi14 := 0; fi14 < 12; fi14++ {
                var v15 uint16 = (v10 ^ (v10 >> ((255) & 63)))
                _ = fi14
                _ = v15
            }
        } else {
            fmt.Println(v1, v1)
            var v16 uint8 = 42
            _ = v16
        }
        for fi17 := 0; fi17 < 19; fi17++ {
            var v18 float32 = float32(((1.5 + (1e10 - 3.14159)) / (-2.25 * (float64((v5 + v5)) / float64(v5)))))
            _ = fi17
            _ = v18
        }
    }
    v2 = 4294967295
    v19 := GenS3{F0: 477, F1: false, F2: 4294967295}
    var v20 float64 = 0.0
    if (int8(v2) == 100) {
        v21 := GenS4{F0: 9223372036854775807, F1: 1000}
        _ = v21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v5
}
