package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 int32; F1 int8; F2 uint }
type GenS3 struct { F0 uint8 }

func main() {
    var v1 uint8 = 42
    var v2 float32 = ((1e10 + 3.14159) / 1.0)
    v3 := []uint32{}
    var v4 uint32 = 613
    v5 := GenS1{F0: 677, F1: false}
    v2 = -2.25
    var v6 uint = 926
    var v7 uint16 = 32767
    var v8 uint64 = 42
    var v9 bool = ((v8 - (v8 + v8)) < (v8 / (((^(v8))) | 1)))
    v10 := GenS2{F0: -2147483648, F1: -121, F2: (v6 >> ((65535) & 63))}
    if (!v9) {
        if (!true) {
            if v9 {
                if (1000 <= 0) {
                    if len(v3) > 0 {
                        v3[0] = (^(v4))
                    }
                }
            } else {
                v6 = v6
            }
            for fi11 := 0; fi11 < 18; fi11++ {
                var v12 string = ("hello" + "")
                if v9 {
                    if v9 {
                        v13 := GenS2{F0: 0, F1: 0, F2: 18446744073709551615}
                        var v14 int16 = 0
                        v4 += uint32(v6)
                        _ = v13
                        _ = v14
                    } else {
                        v3 = append(v3, v4)
                        if (!((v6 & v6) < (v6 << ((990) & 63)))) {
                            var v15 float64 = (1e10 + -2.25)
                            if v9 {
                                var v16 int16 = 568
                                _ = len(v3)
                                _ = v16
                            } else {
                                fmt.Println(v1, v6, v2)
                            }
                            _ = v15
                        } else {
                            var v17 int8 = -128
                            v17 = v17
                            v18 := GenS2{F0: 42, F1: -75, F2: (v6 - 658)}
                            _ = v17
                            _ = v18
                        }
                        v3 = append(v3, 2147483647)
                    }
                    v3 = append(v3, (^(v4)))
                    switch fi11 {
                    case 32768:
                        v19 := GenS3{F0: 0}
                        _ = v19
                    }
                } else {
                    for fi20 := 0; fi20 < 8; fi20++ {
                        fmt.Println(v6, v12, v1)
                        _ = fi20
                    }
                    for fi21 := 0; fi21 < 8; fi21++ {
                        v22 := GenS2{F0: -2147483648, F1: -1, F2: 226}
                        _ = fi21
                        _ = v22
                    }
                }
                for fi23 := 0; fi23 < 5; fi23++ {
                    v1 += (v1 - (v1 & (v1 << (((v6 % (((v6 << ((v6) & 63))) | 1))) & 63))))
                    _ = fi23
                }
                _ = fi11
                _ = v12
            }
        }
    } else {
        fmt.Println(v9)
    }
    var v24 int32 = 42
    fmt.Println(v24)
    var v25 bool = false
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v26 := GenS2{F0: -1, F1: 51, F2: v6}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
