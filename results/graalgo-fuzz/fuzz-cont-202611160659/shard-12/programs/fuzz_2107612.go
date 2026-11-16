package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 string; F2 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 bool; F1 int }
type GenS4 struct { F0 uint; F1 uint16 }
func genFn0() int64 {
    v7 := []int{378, 9223372036854775807}
    var v8 int16 = -276
    switch 9223372036854775807 {
    case 4294967295:
        fmt.Println(v8)
    }
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 float32 = float32(0.1)
    var v2 float64 = 3.14159
    v3 := make(map[string]int)
    v4 := []int{32767}
    v5 := make(map[int]string)
    var v6 string = ""
    v1 -= 1.5
    v2 -= v2
    if len(v4) > 0 {
        v4[0] = 32767
    }
    var v9 uint32 = 2147483647
    fmt.Println(v9, v1)
    fmt.Println(v9, v9, v1)
    for fi10 := 0; fi10 < 5; fi10++ {
        switch 32768 {
        case 399:
            if (true || true) {
                _ = len(v5)
                switch "hello" {
                case "foo bar":
                    for fi11 := 0; fi11 < 2; fi11++ {
                        v12 := GenS1{F0: 42, F1: (((v6 + v6) + v6) + v6), F2: 1000}
                        fmt.Println(fi10)
                        var v13 int8 = 1
                        v14 := GenS2{F0: (0 == 2147483647)}
                        if (-125 != v13) {
                            var v15 float64 = ((3.14159 / v2) * (0.0 + (v2 + ((v2 / (v2 + (float64(float32((1.5 * v2))) * (3.14159 / -2.25)))) * (v2 / v2)))))
                            _ = v15
                        }
                        v2 = -2.25
                        _ = fi11
                        _ = v12
                        _ = v13
                        _ = v14
                    }
                case "hello":
                    for fi16 := 0; fi16 < 2; fi16++ {
                        fmt.Println(fi10, v9)
                        var v17 int64 = 42
                        _ = fi16
                        _ = v17
                    }
                case "":
                    v2 = v2
                }
            } else {
                var v18 int32 = 2147483647
                var v19 uint8 = 47
                fmt.Println(v19, v9)
                _ = v18
                _ = v19
            }
        case -1:
            if (0 > -438) {
                for fi20 := 0; fi20 < 14; fi20++ {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = fi20
                }
                v21 := GenS3{F0: true, F1: (fi10 >> ((42) & 63))}
                if (-1 < 42) {
                    var v22 float64 = 0.1
                    _ = v22
                } else {
                    v4 = append(v4, 32768)
                    v9 += (v9 / ((v9) | 1))
                    for fi23 := 0; fi23 < 3; fi23++ {
                        v4 = append(v4, fi10)
                        v24 := v3[v6]
                        _ = fi23
                        _ = v24
                    }
                }
                _ = v21
            } else {
                var v25 int32 = -139
                var v26 int16 = -32768
                _ = v25
                _ = v26
            }
        case 32768:
            v27 := GenS3{F0: false, F1: fi10}
            _ = v27
        }
        _ = fi10
    }
    v9 -= v9
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v28 := GenS4{F0: 945, F1: 33}
    v29 := GenS1{F0: 1, F1: v6, F2: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
