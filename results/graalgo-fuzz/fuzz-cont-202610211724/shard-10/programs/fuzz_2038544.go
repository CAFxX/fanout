package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64; F1 int; F2 int }
type GenS3 struct { F0 uint64; F1 int16 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int8 = 0
    var v3 uint32 = 0
    var v4 bool = false
    var v5 int16 = 1
    var v6 int8 = v2
    v7 := GenS1{F0: true}
    for fi8 := 0; fi8 < 1; fi8++ {
        v9 := GenS1{F0: false}
        _ = fi8
        _ = v9
    }
    _ = len(v1)
    for fi10 := 0; fi10 < 4; fi10++ {
        v11 := GenS1{F0: v4}
        var v12 float32 = 1.5
        _ = fi10
        _ = v11
        _ = v12
    }
    if v4 {
        v5 += 0
        var v13 uint16 = 534
        switch "a" {
        case "\n\t":
            v2 -= int8(v3)
        }
        _ = v13
    }
    for fi14 := 0; fi14 < 5; fi14++ {
        v15 := []int{(^(fi14)), (fi14 + 617), (fi14 << ((556) & 63)), fi14}
        fmt.Println(v6, v3, v6)
        v15 = append(v15, 32767)
        _ = fi14
        _ = v15
    }
    var v16 float64 = float64(0.1)
    v3 = v3
    if (-51 > 1) {
        v4 = v4
    } else {
        for fi17 := 0; fi17 < 16; fi17++ {
            var v18 uint64 = 1
            if ("" <= "") {
                for fi19 := 0; fi19 < 5; fi19++ {
                    var v20 uint8 = 40
                    if (v20 < 1) {
                        v21 := []int{(fi19 * fi19), fi17, 127, fi19}
                        _ = v21
                    }
                    v6 = v2
                    _ = fi19
                    _ = v20
                }
            }
            var v22 uint16 = 0
            _ = fi17
            _ = v18
            _ = v22
        }
        v23 := GenS2{F0: 42, F1: 163, F2: -566}
        _ = v23
    }
    if (18446744073709551615 < 0) {
        var v24 int64 = 9223372036854775807
        _ = v24
    }
    var v25 int16 = v5
    var v26 int8 = (^(v2))
    v5 = (v5 << ((2) & 63))
    if v4 {
        v1[(v3 + (v3 & (v3 % ((v3) | 1))))] = v3
        for fi27 := 0; fi27 < 3; fi27++ {
            var v28 bool = v4
            _ = fi27
            _ = v28
        }
        if v4 {
            if (v4 || ((255 == 32767) || (32767 <= 0))) {
                v29 := []int{128, -471, 8, 1}
                v30 := GenS1{F0: (!(v16 < 0.1))}
                _ = v29
                _ = v30
            }
        } else {
            for fi31 := 0; fi31 < 17; fi31++ {
                var v32 uint16 = 42
                v33 := []int{(fi31 * fi31)}
                var v34 uint8 = 42
                _ = fi31
                _ = v32
                _ = v33
                _ = v34
            }
            v5 += 42
        }
    } else {
        if v4 {
            v35 := GenS3{F0: 18446744073709551615, F1: (v25 / ((32767) | 1))}
            _ = v35
        }
        v36 := v1[526]
        _ = v36
    }
    var v37 uint16 = 366
    for fi38 := 0; fi38 < 6; fi38++ {
        fmt.Println(v26)
        var v39 uint8 = 201
        var v40 int8 = -74
        _ = fi38
        _ = v39
        _ = v40
    }
    fmt.Println(v37, v26, v3)
    v26 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
