package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int; F1 int8; F2 bool }
type GenS4 struct { F0 int32 }

func main() {
    v1 := make(map[string]uint32)
    var v2 uint = 42
    var v3 uint32 = 1
    var v4 string = ("" + "a")
    var v5 uint = v2
    for fi6 := 0; fi6 < 11; fi6++ {
        fmt.Println(v5)
        for fi7 := 0; fi7 < 13; fi7++ {
            var v8 int16 = -32768
            if ((182 != 42) || (true && (((0.1 - float64(((float32(0.0) / -2.25) + 0.0))) / (((0.0 / 1.0) - 1.5) * 3.14159)) > -2.25))) {
                v4 = v4
                v9 := []int{fi7, fi6, (fi6 & (fi6 % ((745) | 1)))}
                v4 = "\n\t"
                _ = v9
            } else {
                if ((877 > 970) && true) {
                    switch ("a" + "a") {
                    case "a":
                        v1[((v4 + "hello") + v4)] = (^(v3))
                    case "hello":
                        if true {
                            var v10 int64 = 886
                            _ = v10
                        }
                    case "foo bar":
                        fmt.Println(fi6, v4, v8)
                    default:
                        fmt.Println(fi7)
                    }
                    fmt.Println(v3)
                }
                switch (fi7 | (^(fi6))) {
                case -9223372036854775808:
                    if (false || (!(1 == 2))) {
                        v11 := GenS1{F0: -1, F1: v3}
                        var v12 float32 = (0.1 - 1.5)
                        fmt.Println(v12, v2, fi6)
                        v13 := GenS1{F0: 144, F1: 65535}
                        v4 = ""
                        _ = v11
                        _ = v12
                        _ = v13
                    }
                case 4294967295:
                    v14 := []int{127, fi6, fi6, (^(fi6))}
                    _ = v14
                }
            }
            v5 -= v2
            _ = fi7
            _ = v8
        }
        v5 += v5
        _ = fi6
    }
    v3 += (v3 >> ((v5) & 63))
    fmt.Println(v2, v3)
    if false {
        if true {
            for fi15 := 0; fi15 < 18; fi15++ {
                v16 := GenS2{F0: 522}
                v4 = "hello"
                var v17 bool = false
                _ = fi15
                _ = v16
                _ = v17
            }
            var v18 uint8 = 14
            _ = v18
        } else {
            fmt.Println(v4, v5, v3)
            v3 -= (v3 << ((v5) & 63))
        }
        v2 -= v2
        switch 9223372036854775807 {
        case 2147483647:
            var v19 uint = 42
            _ = v19
        case 32768:
            switch -2147483648 {
            case 2147483647:
                v20 := GenS1{F0: 2147483647, F1: v3}
                _ = v20
            case 32767:
                var v21 int32 = -2147483648
                _ = v21
            case 1:
                v22 := GenS3{F0: -892, F1: -1, F2: (0 >= v5)}
                _ = v22
            default:
                v23 := []int{32767, 649, -133}
                _ = v23
            }
        case -554:
            var v24 int8 = -41
            _ = v24
        default:
            v25 := GenS4{F0: 42}
            _ = v25
        }
    } else {
        var v26 uint64 = 18446744073709551615
        _ = v26
    }
    var v27 uint64 = 0
    v28 := []int{2147483647, 42, -22, 42}
    for fi29 := 0; fi29 < 16; fi29++ {
        _ = len(v28)
        _ = fi29
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
