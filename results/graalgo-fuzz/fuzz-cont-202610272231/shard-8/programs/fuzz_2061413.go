package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64; F2 uint }
type GenS2 struct { F0 uint; F1 int8; F2 int16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 string; F1 int8 }

func main() {
    var v1 float64 = -2.25
    var v2 int8 = 1
    v3 := make(map[uint32]bool)
    var v4 int8 = (-(v2))
    var v5 string = ("" + ("hello" + ("" + ("hello" + "a"))))
    var v6 string = v5
    for fi7 := 0; fi7 < 3; fi7++ {
        v1 -= (v1 / (float64(0.0) - v1))
        _ = fi7
    }
    var v8 int = 117
    var v9 float32 = 0.0
    v6 = v6
    if (!(255 == 343)) {
        v10 := GenS1{F0: 0, F1: 1, F2: 255}
        _ = v10
    } else {
        v11 := GenS1{F0: 1, F1: 663, F2: 575}
        var v12 uint64 = 0
        var v13 float64 = (((0.0 + v1) * (float64(v9) + v1)) / v1)
        _ = v11
        _ = v12
        _ = v13
    }
    v4 -= v4
    v5 = v6
    v14 := []int{(v8 ^ 128), int(v4)}
    for fi15 := 0; fi15 < 15; fi15++ {
        v16 := GenS2{F0: 255, F1: (-(v2)), F2: 1}
        _ = fi15
        _ = v16
    }
    if true {
        v8 -= v8
        v9 -= 1.5
    }
    v17 := GenS2{F0: 2, F1: v4, F2: -32768}
    for fi18 := 0; fi18 < 8; fi18++ {
        for fi19 := 0; fi19 < 3; fi19++ {
            var v20 string = "foo bar"
            if (0 <= 1) {
                v21 := GenS3{F0: ((v2 % ((v4) | 1)) != 1)}
                _ = v21
            } else {
                var v22 int32 = -664
                switch v6 {
                case "hello":
                    v8 = v8
                case "\n\t":
                    v23 := GenS1{F0: 2, F1: 42, F2: 2}
                    _ = v23
                default:
                    var v24 int16 = -504
                    _ = v24
                }
                var v25 int = -9223372036854775808
                _ = v22
                _ = v25
            }
            switch 32767 {
            case 2147483647:
                if len(v14) >= 2 {
                    v14 = v14[1:len(v14)]
                }
            case -9223372036854775808:
                switch int(v4) {
                case 573:
                    var v26 string = ""
                    _ = v26
                default:
                    var v27 uint16 = 1
                    _ = v27
                }
            case 1:
                v28 := GenS4{F0: v5, F1: v2}
                _ = v28
            }
            _ = fi19
            _ = v20
        }
        _ = fi18
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
