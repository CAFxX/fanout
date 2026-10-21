package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8; F1 int8 }

func main() {
    v1 := []uint32{}
    var v2 float64 = 1.5
    v3 := []uint32{}
    v4 := make(map[int]string)
    var v5 uint8 = 255
    switch 32767 {
    case 42:
        v2 = v2
    case 963:
        var v6 int64 = -9223372036854775808
        _ = v6
    case 2147483647:
        fmt.Println(v5)
    default:
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    }
    var v7 bool = false
    switch 128 {
    case 0:
        v8 := GenS1{F0: 2}
        _ = v8
    case -770:
        v5 = (^(v5))
    default:
        v7 = (!(107 <= 832))
    }
    var v9 float32 = ((1e10 + float32(float64(((1e10 - 1e10) / ((0.1 / -2.25) - 1.5))))) * (((1.5 / (1e10 - (0.1 * -2.25))) + 3.14159) / 0.1))
    switch 890 {
    case 2147483647:
        fmt.Println(v2)
    case -1:
        switch ("a" + "") {
        case "":
            var v10 string = "hello"
            _ = v10
        case "hello":
            for fi11 := 0; fi11 < 8; fi11++ {
                v7 = (true || v7)
                var v12 uint8 = v5
                if true {
                    if len(v3) > 0 {
                        v3[0] = 0
                    }
                }
                _ = fi11
                _ = v12
            }
        }
    case 128:
        v13 := GenS1{F0: 42}
        _ = v13
    default:
        v14 := GenS1{F0: 869}
        _ = v14
    }
    var v15 int64 = -898
    var v16 uint64 = 1
    var v17 bool = (!(!(v15 >= v15)))
    v15 -= (-(v15))
    v4[32767] = "\n\t"
    v15 = (v15 << ((4294967295) & 63))
    fmt.Println(v9)
    v7 = (-813 <= -172)
    if ((^(v15)) == (v15 - v15)) {
        for fi18 := 0; fi18 < 17; fi18++ {
            var v19 uint8 = 241
            _ = fi18
            _ = v19
        }
        switch -1 {
        case 9223372036854775807:
            v7 = (-561 == -9223372036854775808)
        case 1:
            var v20 float64 = v2
            _ = v20
        case 32767:
            for fi21 := 0; fi21 < 2; fi21++ {
                switch -829 {
                case -1:
                    fmt.Println(fi21, v17, v15)
                case 988:
                    v22 := GenS2{F0: 0, F1: 0}
                    _ = v22
                case 128:
                    v5 = v5
                }
                if (false && (65535 == 2147483648)) {
                    for fi23 := 0; fi23 < 16; fi23++ {
                        if (!(-32768 < 1000)) {
                            var v24 bool = (2 < 2147483647)
                            v25 := v4[fi21]
                            if false {
                                v26 := v4[(fi21 << ((585) & 63))]
                                if v17 {
                                    _ = len(v1)
                                } else {
                                    _ = len(v1)
                                }
                                _ = v26
                            }
                            _ = v24
                            _ = v25
                        } else {
                            v3 = append(v3, 34)
                        }
                        var v27 float64 = float64(v9)
                        var v28 uint64 = v16
                        _ = fi23
                        _ = v27
                        _ = v28
                    }
                } else {
                    var v29 int = (fi21 ^ (fi21 - 32768))
                    switch "a" {
                    case "a":
                        switch v29 {
                        case 0:
                            v3 = append(v3, 318)
                        }
                    default:
                        v9 -= 1e10
                    }
                    _ = v29
                }
                v7 = ((v2 + v2) >= v2)
                _ = fi21
            }
        }
        v7 = true
    }
    var v30 uint32 = 2147483647
    switch 128 {
    case 32767:
        v17 = v7
    default:
        delete(v4, int(v30))
    }
    switch "\n\t" {
    case "hello":
        v31 := v4[1]
        _ = v31
    default:
        var v32 float32 = float32(v2)
        _ = v32
    }
    var v33 int8 = -128
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v30
    _ = v33
    _ = v4
    _ = v5
    _ = v7
    _ = v9
}
