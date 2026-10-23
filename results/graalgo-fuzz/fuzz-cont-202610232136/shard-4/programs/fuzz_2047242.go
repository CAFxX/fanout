package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 int; F1 uint16; F2 uint8 }
type GenS3 struct { F0 uint64; F1 uint32 }
type GenS4 struct { F0 uint; F1 string; F2 bool }

func main() {
    var v1 uint = 294
    var v2 uint8 = 128
    var v3 int16 = -1
    v4 := make(map[uint32]int)
    var v5 int32 = 284
    var v6 uint16 = 310
    switch 32768 {
    case 32768:
        var v7 uint64 = 623
        _ = v7
    case -9223372036854775808:
        var v8 uint16 = v6
        _ = v8
    case 9223372036854775807:
        v5 += 2147483647
    }
    if false {
        delete(v4, 1)
        _ = len(v4)
    }
    v3 -= (v3 * -1)
    v9 := GenS1{F0: 32768, F1: -847}
    v5 -= (v5 | (v5 ^ v5))
    var v10 int16 = (v3 / ((v3) | 1))
    if (v1 > 730) {
        var v11 int8 = 0
        switch "" {
        case "a":
            v3 = (^(v10))
        case "\n\t":
            fmt.Println(v6, v1)
        case "":
            v2 = (v2 & v2)
        }
        _ = v11
    } else {
        v10 = 1000
        v6 = uint16(v1)
        if (true && (v6 > (^(v6)))) {
            v12 := GenS2{F0: 128, F1: uint16(v1), F2: (v2 & 214)}
            switch -1 {
            case -454:
                v13 := []int{}
                _ = v13
            case -342:
                for fi14 := 0; fi14 < 5; fi14++ {
                    var v15 uint32 = 382
                    v1 -= 70
                    v1 = 4294967295
                    _ = fi14
                    _ = v15
                }
            }
            if true {
                var v16 int32 = v5
                _ = v16
            }
            _ = v12
        }
    }
    fmt.Println(v5, v6, v10)
    var v17 uint32 = 65535
    if (true && false) {
        var v18 int32 = -2147483648
        v1 -= (v1 ^ v1)
        _ = v18
    } else {
        var v19 uint = (v1 + (^(v1)))
        _ = v19
    }
    fmt.Println(v1, v6, v3)
    switch "" {
    case "a":
        var v20 int16 = (^(v3))
        _ = v20
    case "\n\t":
        v21 := GenS1{F0: 9223372036854775807, F1: 347}
        _ = v21
    case "hello":
        for fi22 := 0; fi22 < 18; fi22++ {
            var v23 uint = (v1 & v1)
            var v24 uint16 = v6
            _ = fi22
            _ = v23
            _ = v24
        }
    }
    for fi25 := 0; fi25 < 9; fi25++ {
        v26 := GenS1{F0: 127, F1: (fi25 * 4294967295)}
        if (208 < v1) {
            var v27 int16 = (^(v3))
            _ = v27
        } else {
            fmt.Println(fi25, v17, fi25)
        }
        _ = fi25
        _ = v26
    }
    for fi28 := 0; fi28 < 14; fi28++ {
        v29 := GenS3{F0: 4294967295, F1: (v17 | v17)}
        _ = fi28
        _ = v29
    }
    v5 += (v5 ^ (v5 << ((499) & 63)))
    v30 := GenS4{F0: 0, F1: "hello", F2: (v10 >= v10)}
    if (!true) {
        var v31 int = 504
        if (!(65535 >= (v1 - 322))) {
            var v32 int = (v31 | v31)
            v4[v17] = 32767
            _ = v32
        }
        fmt.Println(v1)
        _ = v31
    } else {
        if ((!((!((51 < 42) || (v5 != 2147483647))) && (true && (-2.25 <= 1.5)))) && false) {
            v33 := GenS1{F0: 455, F1: 128}
            v34 := []int{}
            v5 -= int32(v2)
            _ = v33
            _ = v34
        } else {
            v35 := v4[2]
            v36 := []int{(v35 * 751), (v35 ^ v35)}
            var v37 int = 42
            _ = v35
            _ = v36
            _ = v37
        }
        var v38 uint = (v1 % ((v1) | 1))
        fmt.Println(v2, v1, v17)
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v17
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
