package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 uint32 }
type GenS2 struct { F0 int64; F1 bool; F2 uint }
type GenS3 struct { F0 int16; F1 uint32 }

func main() {
    v1 := []int{501, 422}
    var v2 float64 = (1.5 - (-2.25 * (float64(((1.5 - -2.25) / float32((((0.1 / 1.0) / 1.0) - (float64((0.1 * (3.14159 * ((0.0 * 1.5) * 0.0)))) - float64((((1e10 * 0.0) + (1.5 / 1.5)) / (1e10 + 3.14159)))))))) - float64((1.5 / 1.5)))))
    var v3 int8 = 0
    var v4 string = "\n\t"
    var v5 int16 = 1000
    var v6 int = -2147483648
    if (42 != 42) {
        _ = len(v1)
        var v7 uint32 = 0
        _ = v7
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v6 += v6
    switch "a" {
    case "hello":
        switch "\n\t" {
        case "":
            v8 := GenS1{F0: -1, F1: (v3 / ((v3) | 1)), F2: 2147483648}
            _ = v8
        case "a":
            for fi9 := 0; fi9 < 16; fi9++ {
                var v10 bool = false
                v1 = append(v1, (v6 | v6))
                for fi11 := 0; fi11 < 5; fi11++ {
                    v6 = v6
                    var v12 uint16 = 225
                    var v13 uint16 = v12
                    var v14 float64 = v2
                    switch (v4 + (v4 + (v4 + (((v4 + v4) + "a") + ((v4 + "foo bar") + v4))))) {
                    case "foo bar":
                        var v15 int = -76
                        _ = v15
                    default:
                        v16 := GenS2{F0: -529, F1: v10, F2: 42}
                        _ = v16
                    }
                    _ = fi11
                    _ = v12
                    _ = v13
                    _ = v14
                }
                _ = fi9
                _ = v10
            }
        case "hello":
            v6 = v6
        }
    case "a":
        v3 += (v3 >> ((42) & 63))
    case "foo bar":
        if false {
            fmt.Println(v4)
        } else {
            if (!false) {
                v6 = v6
                for fi17 := 0; fi17 < 11; fi17++ {
                    fmt.Println(v6)
                    _ = fi17
                }
                switch v6 {
                case 242:
                    v18 := GenS3{F0: v5, F1: 127}
                    _ = v18
                }
            } else {
                fmt.Println(v6, v6)
                for fi19 := 0; fi19 < 6; fi19++ {
                    for fi20 := 0; fi20 < 17; fi20++ {
                        v21 := GenS1{F0: 1, F1: v3, F2: 0}
                        if (false && (v4 > ((v4 + v4) + v4))) {
                            _ = len(v1)
                            v2 -= float64(1e10)
                            v5 += v5
                            var v22 uint64 = 612
                            _ = v22
                        } else {
                            v23 := GenS1{F0: -414, F1: 127, F2: 2147483647}
                            v5 += (v5 % (((v5 - -364)) | 1))
                            _ = v23
                        }
                        _ = fi20
                        _ = v21
                    }
                    _ = fi19
                }
                v3 = 89
            }
            if false {
                fmt.Println(v5, v4)
            }
        }
    default:
        var v24 uint32 = 2
        _ = v24
    }
    fmt.Println(v2)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v25 bool = ((false && true) || true)
    v6 -= 0
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
