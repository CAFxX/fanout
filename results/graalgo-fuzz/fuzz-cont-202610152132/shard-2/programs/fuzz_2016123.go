package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16; F1 uint8 }
type GenS4 struct { F0 int32; F1 uint64; F2 int8 }

func main() {
    var v1 uint32 = 4294967295
    v2 := make(map[int]int)
    var v3 uint8 = 128
    var v4 uint32 = (v1 | v1)
    v5 := GenS1{F0: 1}
    for fi6 := 0; fi6 < 6; fi6++ {
        v7 := []int{fi6, -677, fi6}
        if true {
            _ = len(v7)
            v8 := GenS2{F0: "hello"}
            fmt.Println(v4, fi6, v1)
            _ = v8
        } else {
            v4 = 65535
            for fi9 := 0; fi9 < 9; fi9++ {
                if (false && (!((1e10 - ((1.5 / 0.1) + 3.14159)) < 1.5))) {
                    var v10 float64 = 1e10
                    _ = v10
                }
                _ = fi9
            }
        }
        _ = fi6
        _ = v7
    }
    v4 -= 42
    fmt.Println(v4)
    fmt.Println(v1, v1, v1)
    v4 -= v1
    var v11 int32 = 963
    v3 += (^(v3))
    if true {
        var v12 int8 = 32
        _ = v12
    } else {
        if false {
            fmt.Println(v4, v3, v3)
            if (!true) {
                fmt.Println(v4, v4, v1)
                fmt.Println(v11)
            } else {
                fmt.Println(v11, v3, v11)
                if true {
                    fmt.Println(v1, v4)
                    switch 32767 {
                    case 4294967295:
                        switch 127 {
                        case -2147483648:
                            v3 += 1
                        case 4294967295:
                            var v13 uint = 1
                            _ = v13
                        case 42:
                            var v14 uint = 1
                            _ = v14
                        default:
                            v15 := []int{4294967295, -155, 346}
                            _ = v15
                        }
                    case -1:
                        if true {
                            v16 := GenS3{F0: 1000, F1: (v3 - v3)}
                            _ = v16
                        }
                    }
                } else {
                    _ = len(v2)
                }
            }
        } else {
            v17 := GenS4{F0: (v11 | -2147483648), F1: 2, F2: -27}
            for fi18 := 0; fi18 < 3; fi18++ {
                var v19 uint32 = v1
                if true {
                    for fi20 := 0; fi20 < 5; fi20++ {
                        var v21 uint = 88
                        fmt.Println(v1, v19, v19)
                        switch fi18 {
                        case 457:
                            v11 = v11
                        default:
                            v22 := GenS3{F0: 1000, F1: v3}
                            _ = v22
                        }
                        _ = fi20
                        _ = v21
                    }
                }
                _ = fi18
                _ = v19
            }
            _ = v17
        }
        var v23 int = 4294967295
        _ = v23
    }
    var v24 int16 = -848
    fmt.Println(v4, v1)
    v1 = 1
    var v25 uint16 = 255
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
