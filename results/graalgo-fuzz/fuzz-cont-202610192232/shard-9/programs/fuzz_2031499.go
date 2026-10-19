package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int; F1 bool }
func genFn0() int16 {
    var v6 float64 = float64(float32((((3.14159 / 3.14159) + 3.14159) + float64(0.0))))
    if (2 <= 52) {
        fmt.Println(v6)
        for fi7 := 0; fi7 < 12; fi7++ {
            var v8 float32 = 3.14159
            switch (("foo bar" + ("foo bar" + "\n\t")) + "a") {
            case "foo bar":
                var v9 bool = true
                _ = v9
            case "":
                fmt.Println(fi7)
            }
            _ = fi7
            _ = v8
        }
        v6 = (v6 * ((((((v6 * v6) * -2.25) + (1.5 - v6)) * (((v6 / (((0.1 - ((v6 / float64(((3.14159 * 3.14159) + (3.14159 / -2.25)))) / 1.5)) - (v6 + v6)) + (v6 / 3.14159))) / v6) - (v6 * 0.1))) - (v6 - (float64(1e10) / 3.14159))) - -2.25))
    }
    _ = v6
    return 32767
}

func main() {
    var v1 string = "foo bar"
    var v2 int = 4294967295
    var v3 int32 = 1
    v4 := make(map[uint32]uint32)
    var v5 float32 = 0.1
    for fi10 := 0; fi10 < 6; fi10++ {
        v2 = v2
        var v11 uint16 = 32767
        for fi12 := 0; fi12 < 1; fi12++ {
            switch v1 {
            case "":
                var v13 float32 = (v5 - v5)
                _ = v13
            default:
                if (v5 < 1.5) {
                    switch (fi10 / ((fi12) | 1)) {
                    case -568:
                        fmt.Println(v2)
                    case 2147483647:
                        delete(v4, 2147483648)
                    }
                    if ((-(v3)) <= v3) {
                        var v14 int64 = 42
                        _ = v14
                    } else {
                        v1 = ((v1 + "hello") + v1)
                    }
                    var v15 uint = 360
                    _ = v15
                } else {
                    v1 = ((v1 + (("" + "\n\t") + v1)) + v1)
                    var v16 string = v1
                    _ = v16
                }
            }
            v17 := GenS1{F0: 32767}
            _ = fi12
            _ = v17
        }
        _ = fi10
        _ = v11
    }
    if (42 < 9223372036854775807) {
        var v18 uint64 = 18446744073709551615
        if (v18 >= 0) {
            v18 += 18446744073709551615
            var v19 int32 = v3
            v20 := v4[835]
            _ = v19
            _ = v20
        }
        _ = v18
    }
    var v21 int32 = v3
    if (!((v1 + v1) >= v1)) {
        var v22 uint = 18446744073709551615
        for fi23 := 0; fi23 < 7; fi23++ {
            var v24 int8 = 42
            for fi25 := 0; fi25 < 9; fi25++ {
                v1 = "foo bar"
                _ = fi25
            }
            _ = fi23
            _ = v24
        }
        if (-32768 != 1) {
            _ = len(v4)
            if true {
                for fi26 := 0; fi26 < 13; fi26++ {
                    var v27 int32 = (v3 & (v3 | 2147483647))
                    _ = fi26
                    _ = v27
                }
                v28 := GenS2{F0: v2, F1: false}
                _ = v28
            } else {
                v3 = (v3 % (((v3 % ((int32(v2)) | 1))) | 1))
            }
        }
        _ = v22
    }
    v3 = int32(v2)
    fmt.Println(v5, v1)
    var v29 uint = 42
    v30 := []int{v2, (v2 * (v2 + (v2 | (v2 | (v2 ^ -910))))), (^(v2))}
    fmt.Println(v3, v29, v2)
    v21 -= 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
