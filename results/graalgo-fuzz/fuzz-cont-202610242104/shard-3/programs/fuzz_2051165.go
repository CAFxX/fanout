package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 uint16 }
type GenS2 struct { F0 uint8; F1 uint64; F2 uint16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint32; F1 uint64; F2 uint16 }
func genFn0(p0 int, p1 uint8) uint8 {
    var v4 int16 = -104
    v5 := []int{(p0 >> ((4294967295) & 63)), (^(p0))}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return (p1 >> ((369) & 63))
}

func main() {
    var v1 int32 = 42
    var v2 int16 = 1000
    var v3 int = 32768
    var v6 int16 = (^(v2))
    var v7 uint64 = 346
    var v8 int = 42
    var v9 int32 = (v1 | (v1 * (v1 >> ((2) & 63))))
    for fi10 := 0; fi10 < 16; fi10++ {
        if false {
            for fi11 := 0; fi11 < 3; fi11++ {
                if true {
                    v8 -= v3
                    v1 += (^(v9))
                }
                v12 := GenS1{F0: 32767, F1: int32(v6), F2: 42}
                v7 += 2
                _ = fi11
                _ = v12
            }
        }
        _ = fi10
    }
    v8 = 128
    for fi13 := 0; fi13 < 11; fi13++ {
        v14 := GenS2{F0: 1, F1: (v7 / (((v7 % ((v7) | 1))) | 1)), F2: 2}
        switch "\n\t" {
        case "\n\t":
            for fi15 := 0; fi15 < 8; fi15++ {
                if false {
                    var v16 string = ""
                    var v17 int32 = v9
                    _ = v16
                    _ = v17
                }
                _ = fi15
            }
        case "":
            var v18 uint8 = uint8(v7)
            _ = v18
        }
        v2 = v2
        _ = fi13
        _ = v14
    }
    v19 := GenS2{F0: 128, F1: v7, F2: 163}
    if true {
        for fi20 := 0; fi20 < 3; fi20++ {
            v21 := GenS1{F0: 42, F1: v1, F2: 2}
            _ = fi20
            _ = v21
        }
        v3 += v8
        v22 := GenS2{F0: 255, F1: v7, F2: 798}
        _ = v22
    } else {
        for fi23 := 0; fi23 < 4; fi23++ {
            v24 := GenS3{F0: 32768}
            var v25 int32 = (v1 + (v1 >> ((255) & 63)))
            var v26 int32 = 1
            _ = fi23
            _ = v24
            _ = v25
            _ = v26
        }
        v6 -= v2
        switch (v3 >> ((27) & 63)) {
        case 42:
            for fi27 := 0; fi27 < 18; fi27++ {
                fmt.Println(v9, v8)
                if ((^(v6)) != (^(v2))) {
                    if true {
                        v8 -= (-(v8))
                        var v28 uint32 = 2
                        var v29 float64 = 1e10
                        _ = v28
                        _ = v29
                    }
                    v2 = v2
                    var v30 int8 = 0
                    _ = v30
                } else {
                    v31 := GenS2{F0: 112, F1: v7, F2: 1}
                    var v32 uint8 = 255
                    _ = v31
                    _ = v32
                }
                var v33 int64 = -9223372036854775808
                _ = fi27
                _ = v33
            }
        default:
            v34 := []int{(v8 - (^(v8))), 42}
            _ = v34
        }
    }
    if (((0.1 * 0.0) <= 3.14159) && (v9 == v9)) {
        var v35 uint8 = 127
        v36 := []int{}
        _ = v35
        _ = v36
    }
    v37 := GenS4{F0: 2147483648, F1: (^(v7)), F2: 42}
    fmt.Println(v2)
    for fi38 := 0; fi38 < 4; fi38++ {
        v2 = (v2 % ((v6) | 1))
        _ = fi38
    }
    var v39 float64 = 0.0
    var v40 uint64 = (v7 & v7)
    fmt.Println(v9)
    if true {
        v8 = 7
        v41 := []int{}
        _ = v41
    } else {
        var v42 uint32 = 890
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v37
    _ = v39
    _ = v40
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
