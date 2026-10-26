package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 uint16 }
type GenS2 struct { F0 uint8; F1 uint }
type GenS3 struct { F0 int32 }

func main() {
    var v1 uint64 = 4294967295
    var v2 uint8 = 42
    v3 := make(map[int]int)
    var v4 string = "\n\t"
    v1 -= (^(v1))
    for fi5 := 0; fi5 < 17; fi5++ {
        v6 := GenS1{F0: -9223372036854775808, F1: 42, F2: 248}
        var v7 int8 = -1
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 int64 = 9223372036854775807
    v8 -= 0
    v9 := GenS2{F0: v2, F1: 2}
    v1 = (^(v1))
    var v10 bool = true
    for fi11 := 0; fi11 < 9; fi11++ {
        if v10 {
            if true {
                v12 := v3[(fi11 << ((0) & 63))]
                _ = v12
            } else {
                v3[(-(fi11))] = fi11
                fmt.Println(v2)
            }
            switch int(v2) {
            case 4294967295:
                var v13 bool = (v10 || v10)
                _ = v13
            case 0:
                v14 := []int{}
                _ = v14
            case 1:
                for fi15 := 0; fi15 < 18; fi15++ {
                    v16 := GenS2{F0: (v2 >> ((4294967295) & 63)), F1: 320}
                    _ = fi15
                    _ = v16
                }
            }
            switch v4 {
            case "a":
                v1 = v1
            case "":
                v4 = (v4 + v4)
            default:
                v17 := GenS3{F0: 42}
                _ = v17
            }
        } else {
            v18 := []int{fi11}
            if (65535 >= 32767) {
                delete(v3, (fi11 >> ((952) & 63)))
            } else {
                fmt.Println(v1, v10)
                v19 := v3[(fi11 % ((fi11) | 1))]
                _ = v19
            }
            var v20 uint = 42
            _ = v18
            _ = v20
        }
        _ = fi11
    }
    var v21 uint8 = v2
    var v22 float32 = 0.1
    v22 -= v22
    for fi23 := 0; fi23 < 7; fi23++ {
        var v24 int = -1
        _ = fi23
        _ = v24
    }
    switch -670 {
    case 127:
        v1 = v1
    }
    var v25 int64 = (-(v8))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v22
    _ = v25
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
