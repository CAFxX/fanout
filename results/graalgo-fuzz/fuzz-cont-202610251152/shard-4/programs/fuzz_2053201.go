package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 int8 }
type GenS2 struct { F0 uint64; F1 uint64; F2 uint64 }
type GenS3 struct { F0 uint8; F1 uint8; F2 uint16 }

func main() {
    var v1 float64 = float64(-2.25)
    var v2 uint = 255
    var v3 uint = 255
    v4 := make(map[int]string)
    v5 := make(map[int]bool)
    v6 := make(map[int]string)
    var v7 uint = v2
    var v8 uint64 = 18446744073709551615
    if true {
        fmt.Println(v3, v8)
    } else {
        var v9 string = ""
        if (!false) {
            fmt.Println(v7)
            v1 = v1
            fmt.Println(v1, v8, v9)
        }
        for fi10 := 0; fi10 < 14; fi10++ {
            switch fi10 {
            case 92:
                switch 93 {
                case 1:
                    v11 := GenS1{F0: fi10, F1: 0, F2: -1}
                    _ = v11
                case 42:
                    if true {
                        if true {
                            delete(v5, (fi10 << ((v3) & 63)))
                            v12 := GenS1{F0: fi10, F1: 100, F2: -1}
                            _ = v12
                        }
                        v2 += (v2 | v2)
                    } else {
                        fmt.Println(fi10)
                        for fi13 := 0; fi13 < 6; fi13++ {
                            if (-9223372036854775808 == 1) {
                                var v14 uint32 = 4294967295
                                v15 := []int{fi10, -1, fi13, -964}
                                _ = v14
                                _ = v15
                            }
                            _ = fi13
                        }
                        switch v9 {
                        case "":
                            v16 := GenS1{F0: 4294967295, F1: 100, F2: 127}
                            _ = v16
                        }
                    }
                }
            case 32767:
                var v17 uint = v7
                _ = v17
            }
            v18 := GenS2{F0: (v8 * 2), F1: 0, F2: (v8 * (v8 & 42))}
            _ = fi10
            _ = v18
        }
        _ = v9
    }
    for fi19 := 0; fi19 < 9; fi19++ {
        delete(v4, fi19)
        _ = fi19
    }
    v20 := GenS3{F0: uint8(v3), F1: 218, F2: 2}
    v3 -= 2
    var v21 int64 = -9223372036854775808
    var v22 int8 = 0
    fmt.Println(v2, v21)
    var v23 uint64 = v8
    switch 128 {
    case -758:
        switch -1 {
        case 1:
            switch -492 {
            case -2147483648:
                v22 += v22
            case 2147483647:
                var v24 uint = v7
                _ = v24
            case 128:
                v7 -= (^(v2))
            default:
                if false {
                    if (v21 <= v21) {
                        if true {
                            fmt.Println(v23)
                        }
                    } else {
                        fmt.Println(v1, v8, v21)
                        v1 = (0.0 * ((v1 + v1) / (float64(3.14159) * v1)))
                    }
                    var v25 int32 = 42
                    _ = v25
                }
            }
        case -9223372036854775808:
            fmt.Println(v22, v7, v22)
        }
    case 127:
        v26 := []int{479, -590, -2147483648, 1}
        _ = v26
    }
    v8 += v23
    var v27 int16 = 662
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
