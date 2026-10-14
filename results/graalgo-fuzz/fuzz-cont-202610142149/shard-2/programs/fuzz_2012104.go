package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int64; F1 int }
func genFn0(p0 bool) int32 {
    fmt.Println(p0)
    var v4 float64 = 1.5
    _ = p0
    _ = v4
    return 58
}

func main() {
    v1 := []string{"\n\t", ("hello" + "hello")}
    v2 := make(map[int]bool)
    v3 := make(map[int]bool)
    var v5 int = 32767
    var v6 int64 = 581
    var v7 int32 = 0
    delete(v3, (v5 | (v5 & v5)))
    _ = len(v1)
    v7 += v7
    if true {
        var v8 uint64 = 0
        var v9 int16 = 1000
        v6 += 0
        _ = v8
        _ = v9
    }
    if true {
        if (!(genFn0(true) <= (-(v7)))) {
            switch (v5 % (((v5 & int(v6))) | 1)) {
            case 32768:
                v5 += (v5 ^ -719)
            }
            _ = len(v1)
            v5 = (v5 ^ v5)
        }
        switch (v5 | v5) {
        case -707:
            if len(v1) > 0 {
                v1[0] = "a"
            }
        case 128:
            for fi10 := 0; fi10 < 8; fi10++ {
                fmt.Println(v5, v5)
                var v11 string = "foo bar"
                v11 = v11
                _ = fi10
                _ = v11
            }
        }
    }
    for fi12 := 0; fi12 < 9; fi12++ {
        _ = len(v1)
        if (true && (!true)) {
            v6 = (v6 & (v6 >> ((42) & 63)))
            v13 := v2[(fi12 >> ((835) & 63))]
            var v14 uint64 = 1
            _ = v13
            _ = v14
        } else {
            fmt.Println(v7, v6, v7)
            v15 := v2[-2147483648]
            v16 := GenS1{F0: v7}
            _ = v15
            _ = v16
        }
        _ = fi12
    }
    v17 := GenS2{F0: v6, F1: v5}
    var v18 float32 = ((float32(0.0) + 3.14159) + (((float32(0.0) + ((0.0 * 3.14159) * 3.14159)) - ((((-2.25 * (3.14159 + (((1.5 / 1.0) / ((3.14159 + 1.5) / (float32((-2.25 - float64(1.5))) + ((1e10 * 0.0) * 3.14159)))) - (0.0 * 0.0)))) - (((1e10 - (float32((3.14159 * 1e10)) - 3.14159)) + 0.0) / -2.25)) * -2.25) / float32((1e10 - 1.5)))) / 3.14159))
    fmt.Println(v6, v6, v18)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
}
