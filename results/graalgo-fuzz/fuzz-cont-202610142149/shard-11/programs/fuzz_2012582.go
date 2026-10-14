package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 uint; F1 uint16; F2 uint32 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]int)
    v3 := make(map[string]string)
    for fi4 := 0; fi4 < 16; fi4++ {
        delete(v3, "hello")
        _ = fi4
    }
    if false {
        var v5 int = -753
        switch (v5 >> ((4294967295) & 63)) {
        case 32767:
            switch (v5 ^ -654) {
            case -512:
                var v6 string = "\n\t"
                _ = v6
            default:
                delete(v1, 28)
            }
        case 2147483647:
            switch 153 {
            case 624:
                v7 := []int{v5, v5}
                _ = v7
            case 352:
                for fi8 := 0; fi8 < 6; fi8++ {
                    fmt.Println(v5, v5)
                    _ = fi8
                }
            case 346:
                if (!true) {
                    fmt.Println(v5)
                }
            default:
                if true {
                    var v9 float64 = -2.25
                    _ = v9
                } else {
                    delete(v2, 2147483647)
                    v5 -= v5
                    var v10 float32 = 3.14159
                    _ = v10
                }
            }
        default:
            v11 := GenS1{F0: (v5 / (((v5 % (((-(v5))) | 1))) | 1)), F1: -32768}
            _ = v11
        }
        _ = v5
    } else {
        switch 337 {
        case -1:
            var v12 uint8 = 2
            _ = v12
        case 32768:
            var v13 uint32 = 2
            _ = v13
        case -9223372036854775808:
            switch 9223372036854775807 {
            case 454:
                var v14 bool = (true && true)
                _ = v14
            }
        }
        if true {
            var v15 bool = (!(32767 > 722))
            v16 := []int{128, -294, -422}
            _ = v15
            _ = v16
        }
    }
    v17 := []int{0}
    v17 = append(v17, 552)
    var v18 int64 = 381
    var v19 int32 = 2147483647
    v18 = v18
    v19 -= (v19 << ((42) & 63))
    v20 := GenS2{F0: 120, F1: 2, F2: 807}
    _ = len(v2)
    if true {
        var v21 int8 = -64
        _ = v21
    }
    if len(v17) > 0 {
        v17[0] = 1
    }
    switch 42 {
    case 449:
        switch int(v18) {
        case 421:
            var v22 uint16 = 32767
            _ = v22
        default:
            for fi23 := 0; fi23 < 16; fi23++ {
                v24 := GenS2{F0: 18446744073709551615, F1: 216, F2: 2147483647}
                fmt.Println(v18)
                if (true || false) {
                    var v25 uint16 = 1
                    v26 := GenS2{F0: 2, F1: v25, F2: 567}
                    v27 := GenS1{F0: 1, F1: 0}
                    _ = v25
                    _ = v26
                    _ = v27
                } else {
                    var v28 int8 = 100
                    _ = v28
                }
                _ = fi23
                _ = v24
            }
        }
    default:
        if true {
            v19 -= 2147483647
            if true {
                var v29 int64 = 0
                _ = v29
            } else {
                _ = len(v17)
                fmt.Println(v18)
            }
            fmt.Println(v18)
        }
    }
    v19 = (^(v19))
    if false {
        delete(v3, "hello")
        v17 = append(v17, 1)
    } else {
        if true {
            fmt.Println(v19, v18)
            switch -9223372036854775808 {
            case 128:
                v30 := GenS1{F0: -2147483648, F1: 1}
                _ = v30
            }
        } else {
            var v31 uint64 = 0
            v18 = v18
            _ = v31
        }
    }
    v18 += v18
    switch -684 {
    case 2147483647:
        var v32 int16 = 32767
        _ = v32
    case 296:
        var v33 int = int(v19)
        _ = v33
    }
    var v34 uint8 = 227
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v34
}
