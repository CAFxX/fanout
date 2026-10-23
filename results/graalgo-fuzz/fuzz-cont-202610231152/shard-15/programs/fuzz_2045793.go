package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int16; F1 uint; F2 uint16 }

func main() {
    v1 := make(map[int]bool)
    v2 := []int{-2147483648, -1}
    var v3 float64 = 0.0
    if false {
        var v4 bool = (4294967295 <= 850)
        fmt.Println(v4, v3)
        _ = v4
    }
    v3 = (v3 - 0.1)
    switch ("hello" + "hello") {
    case "foo bar":
        if len(v2) > 0 {
            v2[0] = 1
        }
    default:
        for fi5 := 0; fi5 < 5; fi5++ {
            for fi6 := 0; fi6 < 2; fi6++ {
                switch fi5 {
                case 1:
                    if len(v2) > 0 {
                        v2[0] = fi5
                    }
                case -13:
                    var v7 int = (fi6 & (fi6 ^ (-(fi6))))
                    _ = v7
                }
                var v8 bool = false
                if len(v2) > 0 {
                    v2[0] = 533
                }
                _ = fi6
                _ = v8
            }
            _ = fi5
        }
    }
    fmt.Println(v3)
    if true {
        var v9 int8 = 127
        var v10 int64 = -1
        var v11 uint64 = 1
        _ = v10
        _ = v11
        _ = v9
    }
    v12 := GenS1{F0: 2147483647}
    v3 = float64(((1.5 / ((-2.25 / 1.5) / 1.0)) - (1.5 * float32(1.5))))
    switch -2147483648 {
    case -590:
        v13 := GenS1{F0: 473}
        _ = v13
    case 32768:
        fmt.Println(v3)
    case 430:
        switch "a" {
        case "a":
            var v14 int16 = -692
            _ = v14
        case "hello":
            switch 127 {
            case 43:
                v3 = 1.5
            }
        }
    default:
        if ("\n\t" < ("a" + "hello")) {
            var v15 uint64 = 42
            var v16 float32 = (0.0 * (1.5 / 0.1))
            _ = v15
            _ = v16
        }
    }
    v3 += (1.5 - float64(0.1))
    switch 1 {
    case 2147483647:
        v3 -= v3
    case 32768:
        for fi17 := 0; fi17 < 4; fi17++ {
            if true {
                for fi18 := 0; fi18 < 4; fi18++ {
                    v3 = v3
                    fmt.Println(fi18, fi18)
                    v3 -= v3
                    _ = fi18
                }
                v1[(fi17 ^ -1)] = true
                v3 -= v3
            }
            _ = fi17
        }
    default:
        var v19 int32 = 92
        _ = v19
    }
    if (42 <= -294) {
        v20 := GenS2{F0: 559}
        v21 := GenS3{F0: 1000, F1: 526, F2: 255}
        var v22 int = -267
        _ = v20
        _ = v21
        _ = v22
    }
    v3 = ((float64(float32((v3 * 1.5))) / (float64(1.5) * float64((3.14159 - float32(1.5))))) * (v3 - 0.0))
    v3 += (v3 + v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
}
