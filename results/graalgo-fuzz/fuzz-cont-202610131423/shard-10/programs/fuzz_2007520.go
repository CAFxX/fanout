package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 int8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 bool; F1 string; F2 int64 }
type GenS4 struct { F0 uint8 }

func main() {
    var v1 int64 = 0
    var v2 bool = false
    var v3 uint = 16
    v4 := make(map[uint32]int)
    v5 := make(map[int]string)
    var v6 uint = (v3 << ((v3) & 63))
    switch 0 {
    case -352:
        switch -9223372036854775808 {
        case 32767:
            v7 := GenS1{F0: 42, F1: 1, F2: 0}
            _ = v7
        case 128:
            if true {
                var v8 uint16 = 93
                _ = v8
            } else {
                v2 = v2
                v2 = v2
            }
        case -1:
            v6 = (v6 / ((v3) | 1))
        }
    }
    if true {
        v9 := []int{}
        if v2 {
            if v2 {
                v10 := GenS1{F0: 32767, F1: 32767, F2: 127}
                _ = v10
            } else {
                var v11 float64 = 1.5
                v6 += (v3 ^ v6)
                switch 128 {
                case 32768:
                    if (0 > int32(v6)) {
                        v12 := v4[2]
                        var v13 float64 = float64(0.0)
                        v2 = ("" >= ("\n\t" + "\n\t"))
                        _ = v12
                        _ = v13
                    }
                case -274:
                    v11 -= 0.1
                case 77:
                    v14 := GenS1{F0: 837, F1: 953, F2: 1}
                    _ = v14
                }
                _ = v11
            }
            fmt.Println(v2, v3, v6)
        }
        var v15 uint = (v3 * v6)
        _ = v15
        _ = v9
    } else {
        for fi16 := 0; fi16 < 4; fi16++ {
            var v17 float64 = (1.5 + -2.25)
            v18 := GenS2{F0: 1}
            _ = fi16
            _ = v17
            _ = v18
        }
        for fi19 := 0; fi19 < 9; fi19++ {
            v3 += v3
            _ = fi19
        }
    }
    v20 := GenS1{F0: 0, F1: 65535, F2: 0}
    v6 += 255
    for fi21 := 0; fi21 < 13; fi21++ {
        fmt.Println(fi21, fi21)
        v2 = (((false && v2) && (0 >= 119)) && v2)
        _ = fi21
    }
    switch 176 {
    case 32768:
        v22 := []int{386, 32767, 9223372036854775807}
        _ = v22
    case 1:
        v23 := GenS3{F0: v2, F1: "\n\t", F2: (^(v1))}
        _ = v23
    }
    if v2 {
        for fi24 := 0; fi24 < 11; fi24++ {
            var v25 int8 = 57
            v3 = (^(v6))
            _ = fi24
            _ = v25
        }
        var v26 string = "\n\t"
        _ = v26
    }
    v2 = v2
    v1 = (v1 + (v1 << (((v3 % ((v3) | 1))) & 63)))
    fmt.Println(v2, v1, v1)
    v27 := GenS4{F0: 2}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
