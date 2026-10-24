package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 int; F1 int64 }
type GenS4 struct { F0 int; F1 uint32; F2 uint16 }

func main() {
    var v1 int16 = 819
    var v2 int64 = -783
    var v3 float32 = (1e10 / 1.0)
    v4 := []uint32{1, 2147483648}
    switch "\n\t" {
    case "\n\t":
        v3 -= (float32(-2.25) - (3.14159 + ((v3 - v3) + (v3 - 1.5))))
    default:
        var v5 int32 = 715
        _ = v5
    }
    var v6 string = "hello"
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi7 := 0; fi7 < 10; fi7++ {
        for fi8 := 0; fi8 < 2; fi8++ {
            fmt.Println(v3)
            if false {
                if false {
                    v2 += (v2 << ((255) & 63))
                    if false {
                        fmt.Println(v2)
                    }
                } else {
                    v6 = v6
                }
                var v9 uint32 = 622
                var v10 int16 = int16(v9)
                _ = v10
                _ = v9
            }
            var v11 uint8 = 3
            _ = fi8
            _ = v11
        }
        _ = fi7
    }
    var v12 int = -9223372036854775808
    if (1 >= -1) {
        v13 := GenS1{F0: -435}
        fmt.Println(v12, v1)
        _ = v13
    }
    var v14 int8 = int8(v2)
    v15 := GenS1{F0: (^(v12))}
    v2 = v2
    v16 := GenS1{F0: int(v2)}
    v17 := GenS1{F0: (v12 >> ((18446744073709551615) & 63))}
    for fi18 := 0; fi18 < 19; fi18++ {
        fmt.Println(v2)
        switch "foo bar" {
        case "a":
            v6 = (v6 + "\n\t")
        case "\n\t":
            if (("\n\t" + (v6 + (v6 + ""))) >= ("a" + "\n\t")) {
                v19 := GenS1{F0: fi18}
                _ = v19
            } else {
                v2 = v2
            }
        default:
            v20 := GenS2{F0: v6, F1: ((fi18 | 32768) < (fi18 % ((int(v1)) | 1)))}
            _ = v20
        }
        var v21 int64 = v2
        _ = fi18
        _ = v21
    }
    if false {
        switch v12 {
        case 964:
            v3 = 3.14159
        }
        switch 4294967295 {
        case 4294967295:
            if len(v4) > 0 {
                v4[0] = 4294967295
            }
        default:
            if true {
                if true {
                    v12 = v12
                }
                v14 += (v14 ^ (v14 | (v14 ^ 127)))
            }
        }
    }
    var v22 int32 = 31
    switch v6 {
    case "a":
        if false {
            v23 := GenS3{F0: 32767, F1: (v2 >> ((4294967295) & 63))}
            _ = v23
        }
    default:
        switch v12 {
        case 271:
            var v24 int16 = 1
            _ = v24
        case 127:
            v25 := GenS4{F0: (v12 % (((v12 & -91)) | 1)), F1: 4294967295, F2: 0}
            _ = v25
        default:
            var v26 uint32 = 0
            _ = v26
        }
    }
    v3 -= v3
    for fi27 := 0; fi27 < 15; fi27++ {
        fmt.Println(fi27, v3, v3)
        _ = fi27
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v6
}
