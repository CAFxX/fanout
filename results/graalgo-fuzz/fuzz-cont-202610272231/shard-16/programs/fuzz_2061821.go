package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16 }

func main() {
    var v1 uint32 = 324
    var v2 uint = 430
    v3 := make(map[uint32]string)
    v4 := []uint32{1, (v1 % (((v1 | (v1 * 2))) | 1))}
    var v5 string = "hello"
    var v6 float32 = 0.0
    fmt.Println(v1, v1, v6)
    for fi7 := 0; fi7 < 4; fi7++ {
        for fi8 := 0; fi8 < 8; fi8++ {
            v5 = v5
            v9 := GenS1{F0: (v1 << ((18446744073709551615) & 63))}
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    v1 = (v1 >> (((v2 % ((2) | 1))) & 63))
    var v10 uint16 = 1
    v11 := GenS2{F0: 985}
    switch 2147483647 {
    case 128:
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case -603:
        v2 = v2
    }
    if (v1 > 794) {
        v1 -= v1
    } else {
        v12 := GenS1{F0: v1}
        _ = v12
    }
    for fi13 := 0; fi13 < 8; fi13++ {
        fmt.Println(v5, v10, v10)
        _ = fi13
    }
    for fi14 := 0; fi14 < 6; fi14++ {
        var v15 int8 = -101
        if false {
            v16 := GenS2{F0: 42}
            if true {
                v1 = (v1 ^ 2)
                if (!false) {
                    for fi17 := 0; fi17 < 8; fi17++ {
                        v10 = (v10 + (v10 + (^(v10))))
                        _ = fi17
                    }
                } else {
                    var v18 int32 = 2147483647
                    var v19 uint32 = v1
                    _ = v18
                    _ = v19
                }
                v1 = (^(v1))
            }
            for fi20 := 0; fi20 < 13; fi20++ {
                v2 += (v2 * (v2 * (v2 << ((uint(v10)) & 63))))
                v6 = v6
                var v21 uint32 = 2147483648
                var v22 int32 = 2147483647
                var v23 string = v5
                _ = fi20
                _ = v21
                _ = v22
                _ = v23
            }
            _ = v16
        }
        if ("\n\t" < "hello") {
            var v24 string = v5
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v1 -= v1
            _ = v24
        }
        _ = fi14
        _ = v15
    }
    if (!true) {
        switch 4294967295 {
        case 32767:
            if true {
                if true {
                    v25 := v3[v1]
                    _ = v25
                }
                for fi26 := 0; fi26 < 1; fi26++ {
                    v1 = uint32(v2)
                    _ = fi26
                }
                var v27 bool = (42 <= -9223372036854775808)
                _ = v27
            }
        case -1:
            var v28 string = "hello"
            _ = v28
        case 0:
            v29 := GenS3{F0: v10}
            _ = v29
        default:
            for fi30 := 0; fi30 < 11; fi30++ {
                var v31 uint64 = uint64(v2)
                _ = fi30
                _ = v31
            }
        }
    }
    v32 := GenS1{F0: 2}
    if false {
        switch -1 {
        case 42:
            if (v5 < v5) {
                fmt.Println(v10, v1)
                v33 := GenS1{F0: v1}
                var v34 uint = 0
                _ = v33
                _ = v34
            }
        case 0:
            if len(v4) > 0 {
                v4[0] = v1
            }
        case 2147483647:
            v35 := GenS2{F0: 632}
            _ = v35
        default:
            if len(v4) > 0 {
                v4[0] = (^(v1))
            }
        }
    }
    fmt.Println(v2, v6, v6)
    var v36 float64 = 1.5
    var v37 int16 = 42
    var v38 float64 = v36
    var v39 uint8 = 255
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v32
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
