package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64 }
type GenS2 struct { F0 bool; F1 string }
type GenS3 struct { F0 int16 }
func genFn0() uint32 {
    v5 := []int{}
    var v6 uint16 = 2
    if (false && true) {
        for fi7 := 0; fi7 < 5; fi7++ {
            v8 := GenS1{F0: (v6 << ((511) & 63)), F1: 697}
            _ = fi7
            _ = v8
        }
        v5 = append(v5, 1)
        v5 = append(v5, -958)
    } else {
        v5 = append(v5, 768)
    }
    v6 = 32767
    _ = v5
    _ = v6
    return 1
}

func main() {
    v1 := []uint32{65535}
    v2 := make(map[int]uint32)
    v3 := make(map[int]bool)
    var v4 float32 = 3.14159
    v4 += 0.0
    fmt.Println(v4)
    for fi9 := 0; fi9 < 7; fi9++ {
        v10 := GenS2{F0: false, F1: (("" + "hello") + "\n\t")}
        _ = fi9
        _ = v10
    }
    var v11 uint = 42
    if true {
        var v12 uint8 = 1
        v13 := GenS1{F0: 2, F1: 963}
        _ = v12
        _ = v13
    }
    v14 := v3[32767]
    v11 -= v11
    fmt.Println(v4)
    if true {
        if (0 < 125) {
            var v15 uint8 = 127
            fmt.Println(v14, v11, v11)
            var v16 float64 = (3.14159 - 0.0)
            _ = v15
            _ = v16
        }
        for fi17 := 0; fi17 < 4; fi17++ {
            for fi18 := 0; fi18 < 4; fi18++ {
                if (-1 >= 2147483647) {
                    v19 := GenS1{F0: 2, F1: 18446744073709551615}
                    v4 += (float32((3.14159 / (((float64(v4) * -2.25) * -2.25) + 1.5))) - (float32(float64(v4)) - v4))
                    _ = v19
                }
                v20 := GenS3{F0: -1}
                _ = fi18
                _ = v20
            }
            _ = fi17
        }
    }
    var v21 float32 = (((v4 - v4) / 1.0) / ((v4 * (0.1 / v4)) / v4))
    for fi22 := 0; fi22 < 20; fi22++ {
        v23 := GenS1{F0: 2, F1: 165}
        var v24 float32 = 0.1
        _ = fi22
        _ = v23
        _ = v24
    }
    var v25 float64 = ((-2.25 * ((-2.25 / 1e10) * 1.5)) + -2.25)
    v26 := GenS3{F0: 0}
    var v27 uint16 = 350
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v14
    _ = v2
    _ = v21
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
