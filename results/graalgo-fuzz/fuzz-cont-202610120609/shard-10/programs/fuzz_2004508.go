package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 int16 }

func main() {
    v1 := make(map[string]string)
    var v2 uint64 = 0
    v3 := make(map[string]string)
    v4 := make(map[uint32]uint32)
    v2 -= (^(v2))
    var v5 int16 = 1000
    v6 := GenS1{F0: 255, F1: 917}
    v7 := GenS1{F0: 42, F1: 0}
    for fi8 := 0; fi8 < 6; fi8++ {
        if true {
            v9 := GenS2{F0: (v5 * v5)}
            _ = len(v3)
            v10 := GenS1{F0: 775, F1: 189}
            _ = v10
            _ = v9
        }
        if (1.5 >= float64(0.1)) {
            fmt.Println(v5, v5, v5)
            var v11 uint64 = 107
            v5 += (v5 & 1000)
            _ = v11
        } else {
            if (42 >= fi8) {
                for fi12 := 0; fi12 < 11; fi12++ {
                    v2 += 18446744073709551615
                    v5 = 32767
                    v5 += -530
                    _ = fi12
                }
            } else {
                fmt.Println(v5, fi8)
                v13 := []int{fi8, 32767, 42}
                for fi14 := 0; fi14 < 9; fi14++ {
                    v15 := GenS2{F0: (^(v5))}
                    switch 42 {
                    case 9223372036854775807:
                        fmt.Println(fi14, v5)
                    case -2147483648:
                        if false {
                            v2 -= v2
                            var v16 int8 = 100
                            _ = v16
                        } else {
                            switch -617 {
                            case -2147483648:
                                var v17 uint64 = (v2 - v2)
                                _ = v17
                            }
                        }
                    }
                    fmt.Println(fi8, v5, v2)
                    _ = fi14
                    _ = v15
                }
                _ = v13
            }
            var v18 int32 = 1
            _ = v18
        }
        _ = fi8
    }
    for fi19 := 0; fi19 < 1; fi19++ {
        for fi20 := 0; fi20 < 18; fi20++ {
            switch "hello" {
            case "\n\t":
                v21 := GenS1{F0: 1000, F1: -1}
                _ = v21
            case "":
                var v22 uint32 = 0
                _ = v22
            case "a":
                v5 = (v5 ^ v5)
            }
            _ = fi20
        }
        for fi23 := 0; fi23 < 14; fi23++ {
            v24 := []int{(fi23 / ((4294967295) | 1))}
            _ = fi23
            _ = v24
        }
        _ = fi19
    }
    var v25 bool = true
    if true {
        v26 := GenS1{F0: 1, F1: -384}
        var v27 uint8 = 128
        _ = v26
        _ = v27
    }
    v25 = false
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
