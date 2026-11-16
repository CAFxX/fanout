package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 int64; F1 uint64; F2 int64 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]int)
    var v3 uint64 = 18446744073709551615
    var v4 int = -1
    v5 := GenS1{F0: (!(!(!true))), F1: 18446744073709551615}
    v4 += (^(v4))
    if true {
        v4 = v4
        if false {
            if false {
                var v6 uint32 = 98
                fmt.Println(v4, v4, v6)
                _ = v6
            } else {
                fmt.Println(v4, v3)
            }
            v3 -= (v3 * 29)
        }
        v4 = v4
    } else {
        var v7 int8 = 1
        var v8 int = (^(v4))
        _ = v7
        _ = v8
    }
    for fi9 := 0; fi9 < 12; fi9++ {
        v10 := GenS1{F0: false, F1: 65535}
        for fi11 := 0; fi11 < 11; fi11++ {
            var v12 uint = 209
            if false {
                var v13 int64 = 9223372036854775807
                _ = v13
            }
            v14 := GenS1{F0: true, F1: v12}
            _ = fi11
            _ = v12
            _ = v14
        }
        v3 -= (v3 << ((2) & 63))
        _ = fi9
        _ = v10
    }
    v3 = v3
    v3 = (v3 + v3)
    var v15 uint32 = 681
    if (true || true) {
        v4 = (v4 * v4)
        v16 := []int{605, v4, -9223372036854775808}
        _ = v16
    }
    v17 := GenS1{F0: false, F1: uint(v15)}
    if (true || (-1 > -1)) {
        var v18 uint8 = 229
        switch v4 {
        case 0:
            if (("a" + ("a" + ("" + "\n\t"))) >= ("" + "hello")) {
                v19 := []int{(-(v4)), v4}
                var v20 uint32 = v15
                var v21 int64 = -1
                _ = v19
                _ = v20
                _ = v21
            }
        case 184:
            v22 := GenS2{F0: 605, F1: v3, F2: 42}
            _ = v22
        default:
            v3 = v3
        }
        _ = v18
    }
    v23 := v1[v4]
    var v24 uint64 = 2
    v25 := GenS1{F0: false, F1: 2}
    switch v4 {
    case 32767:
        switch "foo bar" {
        case "\n\t":
            switch 32767 {
            case 0:
                v26 := []int{(^(v4)), (-(v4))}
                _ = v26
            case 42:
                v3 = (v24 | v3)
            default:
                v15 = 832
            }
        case "a":
            for fi27 := 0; fi27 < 2; fi27++ {
                v28 := GenS2{F0: -9223372036854775808, F1: (^(v24)), F2: int64(v3)}
                v29 := GenS1{F0: false, F1: 132}
                _ = fi27
                _ = v28
                _ = v29
            }
        }
    }
    v30 := GenS1{F0: true, F1: 42}
    v31 := GenS1{F0: (541 <= 9223372036854775807), F1: 18446744073709551615}
    switch (v4 % (((v4 % (((v4 * v4)) | 1))) | 1)) {
    case 32767:
        v4 = -415
    case 2147483647:
        v23 = 329
    case -1:
        v23 += v15
    default:
        v15 -= 1
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v17
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
