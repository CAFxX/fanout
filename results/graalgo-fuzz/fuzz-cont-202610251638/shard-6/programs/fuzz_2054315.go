package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 uint }
func genFn0(p0 uint8, p1 int64) uint32 {
    switch 1 {
    case 408:
        for fi7 := 0; fi7 < 3; fi7++ {
            v8 := GenS1{F0: fi7}
            p1 = 1
            var v9 int8 = -1
            _ = fi7
            _ = v8
            _ = v9
        }
    }
    var v10 uint = 1
    if (!(!false)) {
        fmt.Println(p0, p1)
        for fi11 := 0; fi11 < 4; fi11++ {
            switch ("" + "a") {
            case "foo bar":
                v12 := GenS1{F0: 128}
                _ = v12
            case "hello":
                v13 := GenS1{F0: (^(fi11))}
                _ = v13
            case "a":
                p0 = p0
            default:
                fmt.Println(v10, p0, p1)
            }
            v14 := []int{fi11, (fi11 % ((fi11) | 1)), (fi11 / ((fi11) | 1)), fi11}
            _ = fi11
            _ = v14
        }
    }
    v15 := GenS2{F0: p0, F1: (v10 ^ (^(v10)))}
    _ = p0
    _ = p1
    _ = v10
    _ = v15
    return 2147483647
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []int{128, -2147483648, 0, 127}
    var v3 string = "\n\t"
    var v4 uint8 = 42
    v5 := make(map[string]string)
    var v6 float32 = 3.14159
    var v16 float32 = (v6 - float32(3.14159))
    v6 = ((v6 / float32(3.14159)) - (-2.25 + float32(1e10)))
    v4 = v4
    var v17 int16 = 42
    switch "\n\t" {
    case "\n\t":
        if (-2147483648 > 4294967295) {
            if len(v2) > 0 {
                v2[0] = 32767
            }
            switch 42 {
            case 127:
                switch 4294967295 {
                case -9223372036854775808:
                    v17 = (^(v17))
                }
            case 32768:
                _ = len(v1)
            }
            if (819 == 42) {
                v16 = (v6 + 1e10)
                v18 := v5[(v3 + ((v3 + v3) + v3))]
                _ = v18
            } else {
                var v19 bool = (1 <= 42)
                fmt.Println(v19, v16, v3)
                _ = v19
            }
        }
    case "hello":
        v4 -= 26
    }
    v20 := GenS1{F0: 128}
    if (1e10 <= 0.0) {
        if true {
            var v21 uint8 = v4
            var v22 uint = 197
            _ = v21
            _ = v22
        }
        for fi23 := 0; fi23 < 19; fi23++ {
            var v24 uint32 = 42
            v25 := GenS1{F0: (fi23 + (fi23 ^ fi23))}
            _ = fi23
            _ = v24
            _ = v25
        }
    }
    fmt.Println(v17, v3)
    v26 := GenS1{F0: 242}
    v16 += ((v16 / v16) - -2.25)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
