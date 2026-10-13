package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint32; F1 uint8 }
type GenS3 struct { F0 string; F1 uint64 }

func main() {
    v1 := make(map[int]string)
    var v2 uint8 = 2
    v3 := make(map[string]bool)
    if (v2 != v2) {
        var v4 bool = true
        _ = v4
    } else {
        v5 := GenS1{F0: (float32(3.14159) < (-2.25 - -2.25))}
        _ = v5
    }
    for fi6 := 0; fi6 < 13; fi6++ {
        var v7 int8 = 1
        _ = fi6
        _ = v7
    }
    v2 = v2
    var v8 string = "foo bar"
    v8 = v8
    var v9 int8 = 100
    var v10 int = 168
    var v11 string = "\n\t"
    fmt.Println(v9, v2)
    for fi12 := 0; fi12 < 14; fi12++ {
        switch v8 {
        case "foo bar":
            if (false && true) {
                v13 := GenS2{F0: 967, F1: v2}
                v14 := GenS3{F0: v8, F1: 0}
                fmt.Println(v10)
                _ = v13
                _ = v14
            }
        case "":
            v2 -= 1
        default:
            v15 := []int{fi12, (^(v10)), 4294967295}
            _ = v15
        }
        _ = fi12
    }
    v16 := GenS1{F0: true}
    var v17 uint16 = 255
    var v18 uint16 = (^(v17))
    var v19 int8 = (-(v9))
    switch 32767 {
    case 2147483647:
        v11 = ((("a" + "\n\t") + "a") + v11)
    case 1:
        for fi20 := 0; fi20 < 19; fi20++ {
            fmt.Println(v17)
            _ = fi20
        }
    default:
        v21 := []int{(v10 & v10), (v10 + (v10 ^ (v10 & v10))), 0, (v10 / ((v10) | 1))}
        _ = v21
    }
    var v22 uint8 = 0
    if false {
        if false {
            switch (v10 >> ((0) & 63)) {
            case 128:
                var v23 int16 = 1000
                _ = v23
            case 2147483647:
                v18 -= v17
            }
            if (float32(1e10) == float32(3.14159)) {
                fmt.Println(v10, v11)
                v18 += v17
            } else {
                switch v10 {
                case 128:
                    switch 1 {
                    case 1:
                        for fi24 := 0; fi24 < 6; fi24++ {
                            v11 = "foo bar"
                            _ = fi24
                        }
                    case 0:
                        if ((v9 - 98) < -104) {
                            v2 -= 2
                            v22 -= v2
                            v25 := v3[(v8 + "hello")]
                            _ = v25
                        }
                    case 32768:
                        if false {
                            v17 += (v18 | (v18 / ((v17) | 1)))
                            v26 := GenS2{F0: uint32(v17), F1: (v22 & uint8(v18))}
                            _ = v26
                        }
                    }
                case 32767:
                    v19 = (-(v19))
                case 127:
                    var v27 int64 = -1
                    _ = v27
                default:
                    v17 += (v17 << ((670) & 63))
                }
                _ = len(v1)
                var v28 uint = 1
                _ = v28
            }
        }
    }
    if true {
        fmt.Println(v17)
        v29 := []int{9223372036854775807, int(v19), v10, v10}
        _ = v29
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v8
    _ = v9
}
