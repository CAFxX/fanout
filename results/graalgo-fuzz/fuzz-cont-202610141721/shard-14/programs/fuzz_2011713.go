package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0() int32 {
    v5 := []int{}
    for fi6 := 0; fi6 < 1; fi6++ {
        switch (fi6 << ((784) & 63)) {
        case 9223372036854775807:
            fmt.Println(fi6)
        case 4294967295:
            for fi7 := 0; fi7 < 19; fi7++ {
                if len(v5) > 0 {
                    v5[0] = (fi7 % ((fi7) | 1))
                }
                var v8 int32 = -1
                var v9 uint64 = 2
                _ = fi7
                _ = v8
                _ = v9
            }
        case -133:
            var v10 float64 = 0.0
            _ = v10
        default:
            v11 := GenS1{F0: 32767}
            _ = v11
        }
        for fi12 := 0; fi12 < 19; fi12++ {
            var v13 int16 = 223
            var v14 bool = ((v13 << ((623) & 63)) > -1)
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = fi6
    }
    var v15 int = 127
    v15 += 4294967295
    var v16 float64 = 0.1
    _ = v15
    _ = v16
    _ = v5
    return 213
}
func genFn1() bool {
    var v17 bool = false
    var v18 int64 = 42
    var v19 int16 = 42
    _ = v17
    _ = v18
    _ = v19
    return v17
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[uint32]string)
    v3 := []string{("a" + "")}
    var v4 float32 = ((1.5 + (((float32(0.1) * 3.14159) - (-2.25 / -2.25)) / ((0.0 + 3.14159) + (0.0 - 1e10)))) / 0.1)
    for fi20 := 0; fi20 < 6; fi20++ {
        var v21 float64 = ((-2.25 / 3.14159) * 0.1)
        for fi22 := 0; fi22 < 3; fi22++ {
            v4 -= (v4 + v4)
            if (42 <= -618) {
                if ((fi22 << ((779) & 63)) != fi22) {
                    v4 -= v4
                    v23 := GenS1{F0: 1}
                    var v24 uint8 = 128
                    _ = v23
                    _ = v24
                } else {
                    v21 = v21
                    v21 = 1e10
                }
                fmt.Println(fi22)
                v21 -= (v21 - (((v21 * v21) * (1e10 * 0.1)) + ((3.14159 - (1e10 / ((3.14159 - (v21 + ((0.0 - (v21 - v21)) + -2.25))) / (v21 / 0.1)))) - 1.5)))
            } else {
                v25 := GenS1{F0: 32767}
                var v26 int8 = 1
                if false {
                    v21 += v21
                } else {
                    v4 = 1e10
                    var v27 bool = true
                    _ = v27
                }
                var v28 int16 = 0
                var v29 int = (fi22 & -2147483648)
                _ = v25
                _ = v26
                _ = v28
                _ = v29
            }
            switch fi20 {
            case 32767:
                v21 -= v21
            case 360:
                v30 := GenS1{F0: -1}
                _ = v30
            }
            _ = fi22
        }
        fmt.Println(fi20)
        _ = fi20
        _ = v21
    }
    var v31 string = "\n\t"
    v4 = -2.25
    _ = len(v3)
    var v32 uint16 = 42
    v32 += 2
    v32 = 32767
    fmt.Println(v32, v32, v4)
    switch "hello" {
    case "foo bar":
        v4 = (v4 - -2.25)
    }
    v31 = "a"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v4
}
