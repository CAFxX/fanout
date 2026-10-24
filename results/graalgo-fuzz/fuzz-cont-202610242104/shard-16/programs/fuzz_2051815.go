package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 uint16; F1 uint64; F2 uint16 }

func main() {
    v1 := make(map[string]string)
    var v2 float64 = float64(float32(1.5))
    var v3 uint = 53
    v4 := []string{""}
    var v5 uint = 2
    fmt.Println(v5)
    v2 = (1.5 + 3.14159)
    v6 := GenS1{F0: 520, F1: -1}
    if true {
        fmt.Println(v5, v5, v2)
    }
    v7 := v1["hello"]
    v3 = (^(v3))
    fmt.Println(v3, v3)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if true {
        v8 := GenS2{F0: 32767, F1: 339, F2: 65535}
        for fi9 := 0; fi9 < 5; fi9++ {
            switch (fi9 + 128) {
            case 421:
                for fi10 := 0; fi10 < 8; fi10++ {
                    v11 := GenS1{F0: 4294967295, F1: 427}
                    v5 -= v3
                    _ = fi10
                    _ = v11
                }
            case 824:
                var v12 float32 = (((-2.25 * (-2.25 - (-2.25 - (((0.0 / 1e10) - (1.5 - 0.1)) / 3.14159)))) - -2.25) / ((3.14159 + float32((v2 / 1.0))) / 1.5))
                _ = v12
            case 127:
                var v13 int16 = -61
                _ = v13
            default:
                var v14 int8 = -33
                _ = v14
            }
            var v15 int64 = 1
            _ = fi9
            _ = v15
        }
        var v16 int16 = -1
        _ = v16
        _ = v8
    }
    fmt.Println(v7)
    if (false || (false && (1000 > 32767))) {
        var v17 uint8 = 2
        var v18 int16 = -880
        if (v18 <= v18) {
            v5 += (^(v3))
            v7 = v7
        } else {
            fmt.Println(v7, v7)
            for fi19 := 0; fi19 < 11; fi19++ {
                switch (v7 + "foo bar") {
                case "\n\t":
                    _ = len(v4)
                case "hello":
                    delete(v1, v7)
                }
                _ = fi19
            }
            if (true && false) {
                v20 := GenS1{F0: 18446744073709551615, F1: -337}
                if (v2 > v2) {
                    v21 := GenS1{F0: 4294967295, F1: 0}
                    var v22 int16 = -1
                    v23 := GenS1{F0: 297, F1: -2147483648}
                    _ = v21
                    _ = v22
                    _ = v23
                }
                _ = v20
            }
        }
        _ = v17
        _ = v18
    }
    v2 = v2
    if ((-2.25 * v2) > float64((1e10 * (1.5 + 3.14159)))) {
        for fi24 := 0; fi24 < 11; fi24++ {
            v1[v7] = (v7 + v7)
            v2 += v2
            _ = fi24
        }
        v3 = v5
    } else {
        delete(v1, ("hello" + (v7 + v7)))
        v4 = append(v4, v7)
    }
    var v25 uint32 = 4294967295
    for fi26 := 0; fi26 < 20; fi26++ {
        v27 := v1["\n\t"]
        _ = fi26
        _ = v27
    }
    v1[v7] = v7
    v25 = 2147483647
    var v28 float64 = float64(0.1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v25
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
