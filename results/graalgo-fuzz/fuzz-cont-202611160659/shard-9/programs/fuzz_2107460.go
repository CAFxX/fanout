package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 int64 }
type GenS2 struct { F0 uint32; F1 int32 }
func genFn0(p0 bool, p1 bool, p2 uint16) string {
    var v5 bool = (!((-2.25 * -2.25) < 1e10))
    p1 = (((p2 / ((p2) | 1)) >= p2) || false)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return ("\n\t" + "\n\t")
}

func main() {
    var v1 uint32 = 2147483648
    var v2 bool = true
    var v3 uint32 = 0
    v4 := make(map[string]string)
    var v6 int = 128
    var v7 string = "a"
    fmt.Println(v6, v2)
    var v8 uint = 1
    if v2 {
        var v9 uint64 = 4294967295
        switch (v6 >> ((18446744073709551615) & 63)) {
        case 32767:
            var v10 bool = v2
            _ = v10
        case 32768:
            var v11 bool = (2 == 172)
            _ = v11
        default:
            for fi12 := 0; fi12 < 19; fi12++ {
                var v13 uint = 1
                if v2 {
                    fmt.Println(v1, v9, v8)
                    v7 = ("a" + "\n\t")
                    fmt.Println(v9, v6)
                } else {
                    if v2 {
                        v14 := GenS1{F0: -32768, F1: 32767, F2: -9223372036854775808}
                        v1 = (v3 & 333)
                        v9 = (v9 * (v9 & (v9 ^ (v9 << ((255) & 63)))))
                        _ = v14
                    } else {
                        for fi15 := 0; fi15 < 4; fi15++ {
                            var v16 float64 = 1e10
                            _ = fi15
                            _ = v16
                        }
                    }
                    switch "hello" {
                    case "hello":
                        v17 := GenS2{F0: (v3 * v3), F1: -2147483648}
                        _ = v17
                    default:
                        var v18 uint = (v13 / (((v8 * v13)) | 1))
                        _ = v18
                    }
                    v1 = (v3 ^ (v3 | 80))
                }
                if v2 {
                    for fi19 := 0; fi19 < 8; fi19++ {
                        var v20 uint16 = 32767
                        v21 := GenS1{F0: 1000, F1: 1, F2: -1}
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                    v7 = v7
                    for fi22 := 0; fi22 < 16; fi22++ {
                        v13 -= (v8 ^ (v8 & (v13 << ((793) & 63))))
                        var v23 uint32 = v3
                        _ = fi22
                        _ = v23
                    }
                }
                _ = fi12
                _ = v13
            }
        }
        var v24 int16 = 1
        _ = v24
        _ = v9
    }
    if (true && (true || v2)) {
        v6 += (v6 * v6)
        var v25 int16 = 936
        if false {
            v7 = "hello"
            if (127 < 42) {
                switch (v6 | -2147483648) {
                case 2147483647:
                    var v26 int8 = -16
                    _ = v26
                case 0:
                    v7 = v7
                }
            }
        } else {
            v27 := []int{}
            var v28 uint32 = 0
            _ = v27
            _ = v28
        }
        _ = v25
    } else {
        var v29 uint8 = 1
        fmt.Println(v2, v29)
        v2 = ((^(v29)) >= (v29 | (v29 >> ((uint(v29)) & 63))))
        _ = v29
    }
    v30 := GenS2{F0: v3, F1: 0}
    if (693 != -1) {
        _ = len(v4)
    }
    v31 := GenS2{F0: (^(v1)), F1: -803}
    fmt.Println(v6, v2)
    v32 := []int{}
    switch (v7 + v7) {
    case "\n\t":
        v7 = "foo bar"
    }
    if ((v1 | 0) > (v1 - v3)) {
        var v33 uint = 307
        _ = v33
    }
    switch v7 {
    case "hello":
        var v34 int32 = 301
        _ = v34
    case "\n\t":
        delete(v4, v7)
    default:
        var v35 uint32 = (v1 - (v3 - v3))
        _ = v35
    }
    var v36 int16 = 1000
    var v37 float64 = (float64((-2.25 - (0.0 / float32(1.5)))) + 1e10)
    if len(v32) > 0 {
        v32[0] = (^(v6))
    }
    var v38 int32 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v6
    _ = v7
    _ = v8
}
