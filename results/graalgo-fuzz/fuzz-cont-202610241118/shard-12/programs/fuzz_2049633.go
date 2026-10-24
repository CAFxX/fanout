package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 uint }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint64; F1 uint8 }

func main() {
    var v1 float64 = -2.25
    var v2 int16 = 70
    var v3 int16 = v2
    var v4 int64 = -9223372036854775808
    var v5 uint32 = 162
    var v6 int64 = int64(v5)
    fmt.Println(v2, v4)
    for fi7 := 0; fi7 < 9; fi7++ {
        fmt.Println(v1, v6, fi7)
        if false {
            var v8 int8 = 127
            v9 := GenS1{F0: 42, F1: int32(fi7), F2: 0}
            var v10 int8 = (v8 & v8)
            _ = v10
            _ = v8
            _ = v9
        } else {
            var v11 uint16 = 65535
            v12 := GenS2{F0: -9223372036854775808}
            if false {
                if false {
                    switch fi7 {
                    case 128:
                        v13 := GenS3{F0: 1}
                        _ = v13
                    case 109:
                        v11 = v11
                    }
                    v14 := GenS3{F0: v5}
                    v15 := GenS2{F0: v4}
                    _ = v14
                    _ = v15
                } else {
                    v5 = v5
                    v16 := GenS3{F0: (^(v5))}
                    v1 -= float64(3.14159)
                    _ = v16
                }
                fmt.Println(fi7, v1, v11)
            }
            _ = v11
            _ = v12
        }
        _ = fi7
    }
    v17 := []int{-9223372036854775808, 42, 32768}
    var v18 float32 = ((((-2.25 - 1e10) + 0.1) + ((float32((float64(3.14159) * v1)) / 1e10) - -2.25)) / 1.0)
    switch 32767 {
    case 9223372036854775807:
        var v19 bool = (!false)
        _ = v19
    default:
        var v20 uint8 = 1
        _ = v20
    }
    var v21 bool = false
    var v22 uint32 = 770
    var v23 bool = v21
    if ((v21 && (0 <= 1)) || false) {
        if v21 {
            if false {
                v18 -= 0.0
                v2 += v2
                fmt.Println(v2)
            }
            if ((^(v22)) >= 4294967295) {
                v24 := GenS1{F0: 977, F1: -2147483648, F2: 255}
                _ = v24
            } else {
                v25 := GenS4{F0: 641, F1: 115}
                v2 -= (v3 + 1000)
                _ = v25
            }
            var v26 int16 = 1
            _ = v26
        } else {
            _ = len(v17)
            if v21 {
                for fi27 := 0; fi27 < 20; fi27++ {
                    for fi28 := 0; fi28 < 6; fi28++ {
                        v23 = ((v1 >= v1) || v23)
                        _ = fi28
                    }
                    for fi29 := 0; fi29 < 8; fi29++ {
                        v17 = append(v17, (-(fi29)))
                        _ = fi29
                    }
                    v17 = append(v17, 302)
                    _ = fi27
                }
                for fi30 := 0; fi30 < 11; fi30++ {
                    v31 := GenS3{F0: (^(v22))}
                    _ = fi30
                    _ = v31
                }
            } else {
                var v32 string = "foo bar"
                _ = v32
            }
        }
    } else {
        for fi33 := 0; fi33 < 9; fi33++ {
            v34 := GenS2{F0: v6}
            var v35 bool = true
            switch "foo bar" {
            case "\n\t":
                for fi36 := 0; fi36 < 20; fi36++ {
                    var v37 uint16 = 65535
                    v3 -= v3
                    _ = fi36
                    _ = v37
                }
            default:
                var v38 uint8 = 177
                _ = v38
            }
            _ = fi33
            _ = v34
            _ = v35
        }
        if false {
            v6 = (v6 << ((255) & 63))
        }
        v39 := GenS3{F0: v5}
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
