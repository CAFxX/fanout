package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 uint }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 int16; F1 bool }
type GenS4 struct { F0 int; F1 int }

func main() {
    v1 := make(map[string]string)
    var v2 float32 = ((0.0 - 1e10) - (-2.25 / (1.5 * (((((float32(1.5) / (1.5 + float32(1e10))) - 3.14159) + (1.5 + 1e10)) - -2.25) / float32(3.14159)))))
    var v3 uint32 = 2147483647
    var v4 float64 = (3.14159 / ((1.5 + ((0.1 / 1.5) * float64((float32(((float64(v2) - ((3.14159 + (3.14159 * 1e10)) * 1.5)) * (0.1 - 0.1))) + v2)))) / ((1e10 - float64(((((v2 / -2.25) - (v2 / v2)) * float32(float64(3.14159))) - v2))) / -2.25)))
    if false {
        var v5 int16 = -32768
        _ = v5
    } else {
        fmt.Println(v4)
        if (-2147483648 >= 32767) {
            switch 32768 {
            case 32767:
                v6 := GenS1{F0: 4294967295, F1: 0, F2: 1}
                _ = v6
            case -919:
                var v7 uint32 = v3
                _ = v7
            case 0:
                v4 -= ((v4 * (((v4 / v4) - v4) - v4)) / 1.0)
            }
            v8 := []int{982, 9223372036854775807, 32768, -496}
            _ = v8
        } else {
            v2 += 1e10
            var v9 float64 = 1.5
            var v10 int8 = 0
            _ = v10
            _ = v9
        }
        v11 := []int{1, 1}
        _ = v11
    }
    v12 := GenS1{F0: 0, F1: 871, F2: 1}
    v2 = (1e10 + v2)
    v3 -= 1
    v13 := GenS1{F0: 881, F1: 975, F2: 2}
    v14 := v1[("foo bar" + "foo bar")]
    switch (v14 + (v14 + v14)) {
    case "":
        fmt.Println(v14)
    }
    if (!(111 == 65535)) {
        fmt.Println(v4, v3)
        v15 := GenS1{F0: 647, F1: 153, F2: 65535}
        for fi16 := 0; fi16 < 14; fi16++ {
            for fi17 := 0; fi17 < 20; fi17++ {
                var v18 float32 = (1e10 * 3.14159)
                _ = fi17
                _ = v18
            }
            _ = fi16
        }
        _ = v15
    } else {
        v19 := GenS2{F0: "\n\t", F1: (42 > 65535)}
        _ = v19
    }
    v20 := GenS3{F0: -1, F1: true}
    for fi21 := 0; fi21 < 17; fi21++ {
        var v22 uint64 = 4294967295
        for fi23 := 0; fi23 < 20; fi23++ {
            switch fi21 {
            case 2147483647:
                _ = len(v1)
            case 1:
                v3 = v3
            case 4294967295:
                v14 = (v14 + v14)
            }
            switch v14 {
            case "\n\t":
                if false {
                    v22 += 753
                } else {
                    v24 := GenS4{F0: (fi23 % ((2147483647) | 1)), F1: fi21}
                    _ = v24
                }
            }
            _ = fi23
        }
        _ = len(v1)
        _ = fi21
        _ = v22
    }
    v4 += v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
