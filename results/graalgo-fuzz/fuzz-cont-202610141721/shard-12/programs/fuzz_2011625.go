package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint; F1 int64; F2 int32 }
type GenS3 struct { F0 bool; F1 uint8 }
func genFn0() int32 {
    for fi7 := 0; fi7 < 3; fi7++ {
        var v8 uint16 = 817
        _ = fi7
        _ = v8
    }
    v9 := GenS1{F0: 2}
    v10 := GenS1{F0: 128}
    for fi11 := 0; fi11 < 3; fi11++ {
        if ((fi11 ^ (^(fi11))) >= fi11) {
            var v12 int64 = 1
            var v13 uint16 = 1
            v13 -= v13
            _ = v12
            _ = v13
        } else {
            v14 := GenS2{F0: 757, F1: 128, F2: 2147483647}
            _ = v14
        }
        var v15 float64 = (((3.14159 - ((float64(float32((-2.25 - ((-2.25 * 1.5) / (float64(1.5) / 3.14159))))) + 1.5) + float64((((0.0 / (1.5 * 3.14159)) * -2.25) * 1.5)))) - 0.0) - float64(float32(1e10)))
        _ = fi11
        _ = v15
    }
    var v16 uint16 = 32767
    _ = v10
    _ = v16
    _ = v9
    return -2147483648
}

func main() {
    v1 := []string{"", "\n\t", "foo bar"}
    v2 := make(map[string]int)
    v3 := make(map[uint32]int)
    var v4 int64 = 36
    v5 := make(map[uint32]int)
    v6 := []string{"", ("\n\t" + ("foo bar" + "a")), "hello"}
    v4 = v4
    var v17 uint16 = 32767
    v3[132] = -815
    fmt.Println(v17)
    if ((((1.5 - 1.5) == 3.14159) || (1 != 861)) || false) {
        v18 := GenS3{F0: false, F1: 42}
        v19 := GenS3{F0: false, F1: 0}
        _ = v18
        _ = v19
    }
    for fi20 := 0; fi20 < 10; fi20++ {
        if (!(1000 < -1)) {
            fmt.Println(v17, v17)
        } else {
            for fi21 := 0; fi21 < 9; fi21++ {
                v17 = 475
                _ = fi21
            }
            v17 -= v17
            if ((1 <= (v17 / ((65535) | 1))) || true) {
                fmt.Println(v17, v4)
            }
        }
        fmt.Println(fi20)
        _ = fi20
    }
    v17 += v17
    var v22 int32 = -45
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
