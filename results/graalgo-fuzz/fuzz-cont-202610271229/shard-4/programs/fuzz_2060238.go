package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint; F1 bool; F2 bool }
func genFn0(p0 int8, p1 string) uint32 {
    fmt.Println(p0, p1)
    var v4 uint32 = 42
    fmt.Println(p0, p1, p1)
    v5 := GenS1{F0: 151, F1: int(p0)}
    v6 := GenS1{F0: 0, F1: 42}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    return v4
}

func main() {
    v1 := []int{}
    var v2 float64 = -2.25
    var v3 float32 = 1.5
    for fi7 := 0; fi7 < 17; fi7++ {
        if false {
            for fi8 := 0; fi8 < 11; fi8++ {
                v3 += 3.14159
                switch (^(fi8)) {
                case 32767:
                    if true {
                        fmt.Println(fi7, fi7)
                    } else {
                        v3 += ((v3 / v3) * (v3 * v3))
                        if (true && (!(float32(float64(3.14159)) > float32((v2 * v2))))) {
                            var v9 int8 = 1
                            fmt.Println(fi8, v9, v3)
                            if true {
                                if len(v1) > 0 {
                                    v1[0] = fi7
                                }
                                v10 := GenS1{F0: 331, F1: int(v9)}
                                _ = v10
                            } else {
                                var v11 float32 = v3
                                var v12 uint8 = 151
                                _ = v11
                                _ = v12
                            }
                            v3 = 3.14159
                            fmt.Println(fi8)
                            var v13 bool = (!("hello" == "foo bar"))
                            switch "a" {
                            case "\n\t":
                                v14 := GenS1{F0: 369, F1: (fi7 / (((fi7 * fi8)) | 1))}
                                _ = v14
                            case "hello":
                                _ = len(v1)
                            case "a":
                                fmt.Println(fi7, fi7)
                            default:
                                fmt.Println(fi7, fi7, v9)
                            }
                            _ = v13
                            _ = v9
                        }
                    }
                case 309:
                    var v15 uint = 919
                    _ = v15
                default:
                    v1 = append(v1, fi8)
                }
                _ = fi8
            }
        }
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        for fi16 := 0; fi16 < 16; fi16++ {
            v17 := GenS1{F0: -184, F1: fi7}
            if len(v1) > 0 {
                v1[0] = (fi7 | fi7)
            }
            var v18 uint16 = 192
            var v19 uint = 18446744073709551615
            _ = fi16
            _ = v17
            _ = v18
            _ = v19
        }
        _ = fi7
    }
    fmt.Println(v3)
    v20 := GenS2{F0: 661}
    v2 += v2
    v1 = append(v1, 2147483647)
    switch -480 {
    case -2147483648:
        switch 9223372036854775807 {
        case -1:
            _ = len(v1)
        }
    case -799:
        for fi21 := 0; fi21 < 17; fi21++ {
            var v22 int32 = int32(fi21)
            _ = fi21
            _ = v22
        }
    case 1:
        v2 -= ((float64(-2.25) / -2.25) - v2)
    }
    var v23 float32 = ((v3 + v3) / v3)
    for fi24 := 0; fi24 < 9; fi24++ {
        _ = len(v1)
        if true {
            for fi25 := 0; fi25 < 9; fi25++ {
                v26 := GenS3{F0: 134, F1: true, F2: (1 > 255)}
                _ = fi25
                _ = v26
            }
            if (!true) {
                for fi27 := 0; fi27 < 7; fi27++ {
                    var v28 int32 = -702
                    var v29 bool = true
                    _ = fi27
                    _ = v28
                    _ = v29
                }
                v3 = v3
            }
        }
        _ = fi24
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v3
}
