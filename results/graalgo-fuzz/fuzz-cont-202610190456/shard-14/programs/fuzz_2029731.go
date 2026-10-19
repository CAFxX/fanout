package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32; F2 uint32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int16; F1 bool; F2 uint8 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint32; F1 uint16 }
type GenS6 struct { F0 int64; F1 int; F2 int16 }
type GenS7 struct { F0 uint32; F1 uint8 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]string)
    var v4 uint16 = 42
    var v5 float32 = float32(0.0)
    var v6 string = ""
    v7 := GenS1{F0: 284, F1: 1, F2: 0}
    v5 = v5
    var v8 uint8 = 119
    if false {
        if (2147483647 <= 999) {
            v8 = v8
        }
        for fi9 := 0; fi9 < 11; fi9++ {
            if ((true && ((!(-1 != 42)) || (((3.14159 + -2.25) - v5) != v5))) || false) {
                switch "\n\t" {
                case "\n\t":
                    for fi10 := 0; fi10 < 14; fi10++ {
                        v11 := GenS2{F0: 19}
                        _ = fi10
                        _ = v11
                    }
                case "hello":
                    var v12 uint32 = 0
                    _ = v12
                case "foo bar":
                    for fi13 := 0; fi13 < 6; fi13++ {
                        fmt.Println(v5)
                        var v14 int = fi9
                        _ = fi13
                        _ = v14
                    }
                }
                v15 := GenS1{F0: 255, F1: -2147483648, F2: 958}
                v16 := GenS3{F0: 32767, F1: (true || (fi9 <= 0)), F2: v8}
                _ = v15
                _ = v16
            }
            if true {
                v17 := GenS4{F0: true}
                var v18 bool = false
                _ = v17
                _ = v18
            } else {
                v19 := GenS4{F0: false}
                _ = v19
            }
            _ = fi9
        }
        v6 = "hello"
    }
    for fi20 := 0; fi20 < 15; fi20++ {
        v4 -= v4
        v21 := v1[65535]
        for fi22 := 0; fi22 < 3; fi22++ {
            var v23 uint8 = 157
            _ = fi22
            _ = v23
        }
        _ = fi20
        _ = v21
    }
    var v24 int32 = -588
    for fi25 := 0; fi25 < 18; fi25++ {
        var v26 uint64 = 848
        v8 = v8
        _ = fi25
        _ = v26
    }
    if false {
        if ((^(v8)) < v8) {
            fmt.Println(v8)
            for fi27 := 0; fi27 < 15; fi27++ {
                if (1.5 <= ((1e10 / 1.0) + 0.0)) {
                    v8 -= (v8 * (v8 + (v8 << ((4294967295) & 63))))
                    var v28 uint64 = 520
                    _ = v28
                } else {
                    switch (fi27 / (((fi27 >> ((255) & 63))) | 1)) {
                    case 128:
                        v6 = v6
                    case 20:
                        fmt.Println(v24)
                    case 0:
                        for fi29 := 0; fi29 < 12; fi29++ {
                            v30 := GenS5{F0: 200, F1: v4}
                            _ = fi29
                            _ = v30
                        }
                    }
                }
                _ = fi27
            }
            fmt.Println(v24, v4)
        } else {
            var v31 int8 = 42
            v24 = v24
            _ = v31
        }
    } else {
        for fi32 := 0; fi32 < 9; fi32++ {
            _ = len(v1)
            v24 += (v24 % ((v24) | 1))
            v4 += (v4 & 42)
            _ = fi32
        }
        v6 = v6
    }
    for fi33 := 0; fi33 < 20; fi33++ {
        delete(v2, "\n\t")
        if false {
            var v34 int8 = -96
            _ = v34
        }
        _ = fi33
    }
    for fi35 := 0; fi35 < 4; fi35++ {
        switch v6 {
        case "hello":
            v5 -= v5
        case "a":
            var v36 int16 = 1
            _ = v36
        case "":
            var v37 int8 = 0
            _ = v37
        default:
            v38 := GenS6{F0: int64(v4), F1: -528, F2: -1}
            _ = v38
        }
        v4 = v4
        v39 := GenS7{F0: 2147483647, F1: (v8 / ((255) | 1))}
        _ = fi35
        _ = v39
    }
    v40 := v3[4294967295]
    v6 = v6
    v40 = ((v40 + (("\n\t" + "\n\t") + (v40 + v40))) + v6)
    for fi41 := 0; fi41 < 11; fi41++ {
        v42 := []int{fi41, (fi41 ^ (fi41 >> ((18446744073709551615) & 63)))}
        for fi43 := 0; fi43 < 13; fi43++ {
            var v44 float32 = float32(-2.25)
            var v45 float64 = (0.1 - 1.5)
            var v46 int = 127
            _ = fi43
            _ = v44
            _ = v45
            _ = v46
        }
        _ = fi41
        _ = v42
    }
    v4 += 65535
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
