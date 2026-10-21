package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 int16 }
type GenS2 struct { F0 uint16; F1 uint32 }
type GenS3 struct { F0 uint64; F1 uint16; F2 int8 }

func main() {
    var v1 float32 = 1.5
    var v2 bool = (!(255 == 106))
    var v3 int32 = -230
    v1 += 1e10
    var v4 int = 300
    v2 = v2
    for fi5 := 0; fi5 < 4; fi5++ {
        for fi6 := 0; fi6 < 17; fi6++ {
            var v7 uint8 = 88
            fmt.Println(v3, fi6, fi6)
            _ = fi6
            _ = v7
        }
        v4 -= (fi5 >> ((240) & 63))
        v8 := GenS1{F0: 85, F1: 42, F2: 32767}
        _ = fi5
        _ = v8
    }
    v9 := GenS1{F0: 100, F1: 1, F2: 693}
    fmt.Println(v3, v4, v2)
    v3 -= v3
    switch ("hello" + "foo bar") {
    case "\n\t":
        var v10 float32 = v1
        _ = v10
    case "a":
        v11 := GenS1{F0: 42, F1: 210, F2: 1}
        _ = v11
    default:
        var v12 int8 = -31
        _ = v12
    }
    v13 := GenS1{F0: 1, F1: 79, F2: -629}
    if true {
        for fi14 := 0; fi14 < 15; fi14++ {
            var v15 float32 = v1
            _ = fi14
            _ = v15
        }
        v16 := GenS1{F0: -109, F1: -32768, F2: 1}
        v2 = false
        _ = v16
    }
    var v17 uint8 = 128
    for fi18 := 0; fi18 < 20; fi18++ {
        v17 = (^(v17))
        _ = fi18
    }
    if v2 {
        fmt.Println(v3, v17)
    } else {
        if (v4 <= v4) {
            if (-2147483648 >= v3) {
                if v2 {
                    fmt.Println(v17, v2)
                    v19 := GenS2{F0: 482, F1: 2}
                    v20 := GenS1{F0: -128, F1: 42, F2: 42}
                    _ = v19
                    _ = v20
                }
                v21 := GenS2{F0: 840, F1: 432}
                if (2 != 32767) {
                    var v22 float32 = v1
                    if v2 {
                        var v23 uint16 = 1
                        v24 := []int{v4, -2147483648}
                        fmt.Println(v23, v23, v22)
                        _ = v23
                        _ = v24
                    }
                    _ = v22
                } else {
                    if true {
                        switch 127 {
                        case 909:
                            v17 -= (v17 | (^(v17)))
                        case -124:
                            for fi25 := 0; fi25 < 7; fi25++ {
                                v26 := []int{-9223372036854775808, v4, (-(v4))}
                                _ = fi25
                                _ = v26
                            }
                        default:
                            v27 := []int{}
                            _ = v27
                        }
                    }
                }
                _ = v21
            }
            v28 := GenS3{F0: 516, F1: 255, F2: 127}
            var v29 uint8 = (^(v17))
            _ = v28
            _ = v29
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
