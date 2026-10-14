package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 uint8; F1 int32 }
type GenS3 struct { F0 uint64; F1 uint8; F2 int32 }

func main() {
    var v1 uint32 = 0
    v2 := make(map[string]string)
    var v3 uint8 = 0
    v4 := GenS1{F0: v3, F1: (!true)}
    switch 4294967295 {
    case -168:
        v5 := GenS1{F0: (v3 + 127), F1: true}
        _ = v5
    case -9223372036854775808:
        if true {
            switch -1 {
            case 4294967295:
                v6 := GenS2{F0: 1, F1: 227}
                _ = v6
            case 32768:
                var v7 uint64 = 0
                _ = v7
            }
            var v8 int64 = 9223372036854775807
            v9 := []int{-5}
            _ = v8
            _ = v9
        } else {
            var v10 int8 = 127
            fmt.Println(v3, v10)
            _ = v10
        }
    default:
        v11 := GenS3{F0: 2, F1: v3, F2: 42}
        _ = v11
    }
    for fi12 := 0; fi12 < 17; fi12++ {
        if ((("hello" + "\n\t") + "a") <= ("a" + (("foo bar" + "foo bar") + "hello"))) {
            if ((0 == uint16(v1)) || (v3 == v3)) {
                for fi13 := 0; fi13 < 5; fi13++ {
                    v14 := []int{(^(fi13))}
                    switch "" {
                    case "":
                        v15 := GenS2{F0: v3, F1: 2147483647}
                        _ = v15
                    case "foo bar":
                        fmt.Println(v3, fi13)
                    }
                    if (!((v3 | (^(v3))) >= v3)) {
                        var v16 bool = true
                        var v17 float64 = ((float64(float32(1e10)) / (0.0 - (3.14159 / float64(1e10)))) - (float64(((0.1 * 3.14159) / float32((0.1 + (((3.14159 - 3.14159) - 0.1) + 1e10))))) / ((1e10 - (((1e10 * float64((1e10 + (1.5 / 1e10)))) + -2.25) / (0.0 + (1.5 + (0.1 + 1e10))))) + 1e10)))
                        var v18 int16 = 167
                        _ = v16
                        _ = v17
                        _ = v18
                    } else {
                        if len(v14) >= 2 {
                            v14 = v14[1:len(v14)]
                        }
                    }
                    _ = fi13
                    _ = v14
                }
                var v19 int64 = -294
                v3 -= v3
                _ = v19
            }
        } else {
            v3 += (v3 % (((^(v3))) | 1))
            v20 := GenS2{F0: v3, F1: -906}
            fmt.Println(v1, v3, v3)
            _ = v20
        }
        v1 = 451
        _ = fi12
    }
    var v21 int64 = 42
    for fi22 := 0; fi22 < 8; fi22++ {
        v3 += 1
        _ = fi22
    }
    for fi23 := 0; fi23 < 6; fi23++ {
        var v24 bool = (v3 >= v3)
        _ = fi23
        _ = v24
    }
    for fi25 := 0; fi25 < 15; fi25++ {
        v26 := []int{fi25, (^(fi25))}
        _ = fi25
        _ = v26
    }
    var v27 uint16 = 589
    var v28 int = 1
    v3 = 0
    fmt.Println(v28)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v21
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
