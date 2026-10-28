package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
type GenS2 struct { F0 uint16; F1 uint64; F2 bool }
type GenS3 struct { F0 int; F1 bool }
func genFn0(p0 uint64) uint64 {
    fmt.Println(p0)
    v6 := GenS1{F0: 926, F1: 42}
    switch ("\n\t" + "\n\t") {
    case "\n\t":
        var v7 uint = 0
        _ = v7
    case "":
        p0 = 1
    }
    var v8 uint32 = 42
    v8 -= 538
    _ = p0
    _ = v6
    _ = v8
    return p0
}

func main() {
    var v1 int8 = 127
    var v2 int8 = v1
    var v3 string = ""
    v4 := make(map[int]uint32)
    var v5 uint32 = 670
    if (float32((0.1 * ((3.14159 / 1e10) + 3.14159))) < 1e10) {
        fmt.Println(v2, v1)
    } else {
        v1 += -27
    }
    v9 := GenS2{F0: 1, F1: 18446744073709551615, F2: (0.0 != 3.14159)}
    if false {
        v5 -= (v5 << ((2) & 63))
        v5 -= v5
    }
    fmt.Println(v1, v3, v2)
    fmt.Println(v2)
    for fi10 := 0; fi10 < 10; fi10++ {
        v11 := v4[fi10]
        _ = fi10
        _ = v11
    }
    v12 := GenS1{F0: 0, F1: 1}
    for fi13 := 0; fi13 < 6; fi13++ {
        if false {
            switch "a" {
            case "hello":
                switch (v3 + v3) {
                case "foo bar":
                    v14 := GenS1{F0: 255, F1: 93}
                    _ = v14
                }
            case "\n\t":
                for fi15 := 0; fi15 < 1; fi15++ {
                    delete(v4, (fi15 ^ -780))
                    if true {
                        v16 := GenS2{F0: 1, F1: 18446744073709551615, F2: false}
                        _ = v16
                    }
                    var v17 int32 = 0
                    _ = fi15
                    _ = v17
                }
            }
        }
        v5 = v5
        _ = fi13
    }
    if (18446744073709551615 >= genFn0(771)) {
        if (63 >= 128) {
            fmt.Println(v2, v2, v2)
            v18 := GenS2{F0: 32767, F1: 18446744073709551615, F2: (false && true)}
            _ = v18
        }
    }
    v19 := []int{188, -830}
    switch -1 {
    case 461:
        v20 := GenS2{F0: 255, F1: 1, F2: false}
        _ = v20
    case -284:
        var v21 float32 = 0.0
        _ = v21
    case 2147483647:
        v2 += v2
    default:
        var v22 float64 = 1.5
        _ = v22
    }
    v5 = (v5 * 532)
    switch "foo bar" {
    case "\n\t":
        v23 := GenS2{F0: 682, F1: 638, F2: (((true && false) || (127 != 127)) || (1e10 >= 0.1))}
        _ = v23
    case "foo bar":
        fmt.Println(v3)
    }
    delete(v4, 127)
    v24 := GenS3{F0: 115, F1: (!(false || (((v1 & (v1 / (((v2 | v2)) | 1))) < (v1 >> ((65535) & 63))) && (9223372036854775807 <= -40))))}
    var v25 uint8 = 61
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v19
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
