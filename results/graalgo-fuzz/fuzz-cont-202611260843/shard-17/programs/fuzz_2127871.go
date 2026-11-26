package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32 }
func genFn0(p0 int16, p1 string, p2 uint) int64 {
    v5 := []int{}
    var v6 uint8 = 1
    if true {
        if false {
            v6 += 128
            for fi7 := 0; fi7 < 15; fi7++ {
                v8 := GenS1{F0: p1}
                if (!true) {
                    if len(v5) > 0 {
                        v5[0] = (fi7 / ((fi7) | 1))
                    }
                }
                _ = fi7
                _ = v8
            }
        }
    }
    for fi9 := 0; fi9 < 13; fi9++ {
        switch int(p2) {
        case 32768:
            var v10 uint32 = 42
            _ = v10
        case 9223372036854775807:
            var v11 int16 = (p0 / ((p0) | 1))
            _ = v11
        }
        fmt.Println(fi9, fi9, v6)
        _ = fi9
    }
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return int64(v6)
}

func main() {
    v1 := make(map[string]string)
    var v2 float64 = (((((float64(0.0) + float64(((1e10 - (-2.25 * (-2.25 / 1e10))) - ((0.0 + 3.14159) * 1e10)))) - -2.25) / ((3.14159 - float64(1e10)) * 1.5)) / 1.0) - 0.1)
    var v3 uint = 42
    var v4 int = -9223372036854775808
    fmt.Println(v2)
    var v12 int64 = genFn0(1000, "hello", (^(v3)))
    v13 := GenS2{F0: 0}
    v12 += genFn0(0, "hello", v3)
    v4 += (v4 >> (((v3 % (((v3 | v3)) | 1))) & 63))
    fmt.Println(v3, v4, v4)
    v2 = (v2 - float64(-2.25))
    var v14 uint8 = 128
    for fi15 := 0; fi15 < 2; fi15++ {
        var v16 int = (v4 << (((v3 - 2)) & 63))
        _ = fi15
        _ = v16
    }
    v14 -= v14
    var v17 uint8 = 42
    fmt.Println(v17)
    if (-8 < -128) {
        v4 = v4
        _ = len(v1)
    } else {
        v18 := GenS2{F0: 36}
        for fi19 := 0; fi19 < 4; fi19++ {
            delete(v1, ("\n\t" + "hello"))
            v17 -= 127
            v2 += v2
            _ = fi19
        }
        _ = v18
    }
    v20 := GenS1{F0: "\n\t"}
    v21 := GenS1{F0: (("a" + "a") + "foo bar")}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
