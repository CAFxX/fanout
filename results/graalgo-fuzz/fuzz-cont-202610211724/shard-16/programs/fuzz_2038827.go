package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 uint64 }
type GenS2 struct { F0 int16; F1 string }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint16; F1 uint16; F2 int }

func main() {
    v1 := []int{-603}
    var v2 uint64 = 4294967295
    v3 := []int{-2147483648, 4294967295}
    v4 := []string{}
    v5 := make(map[int]uint32)
    if len(v1) > 0 {
        v1[0] = 0
    }
    v2 = (v2 << ((18446744073709551615) & 63))
    v2 -= v2
    switch "hello" {
    case "hello":
        v1 = append(v1, -534)
    case "":
        v2 -= 2
    }
    _ = len(v1)
    v6 := GenS1{F0: -47, F1: 127, F2: (v2 + v2)}
    if (32767 > -32768) {
        var v7 uint16 = 32767
        _ = v7
    }
    fmt.Println(v2)
    fmt.Println(v2)
    switch -571 {
    case 4294967295:
        v8 := GenS2{F0: -485, F1: ("hello" + ("" + ("" + "\n\t")))}
        _ = v8
    case 897:
        v2 -= v2
    case -147:
        var v9 uint = 1
        _ = v9
    }
    if (-1 == 733) {
        v5[-276] = 4294967295
        v2 = (v2 + (^(v2)))
        if true {
            if (127 != 62) {
                if (!(2 == 1)) {
                    if (1 < 127) {
                        v5[32768] = 4294967295
                    } else {
                        for fi10 := 0; fi10 < 8; fi10++ {
                            fmt.Println(fi10, fi10)
                            var v11 uint8 = 127
                            v2 = v2
                            _ = fi10
                            _ = v11
                        }
                        v2 -= v2
                        var v12 uint16 = 2
                        _ = v12
                    }
                } else {
                    v2 = (^(v2))
                    v2 += 1
                }
                if true {
                    var v13 int64 = 42
                    switch "\n\t" {
                    case "":
                        if ((-2147483648 >= 2147483647) && true) {
                            v2 -= (^(v2))
                            delete(v5, 42)
                        }
                    case "\n\t":
                        var v14 int8 = 89
                        _ = v14
                    default:
                        var v15 float32 = 3.14159
                        _ = v15
                    }
                    _ = v13
                }
                v1 = append(v1, 42)
            }
        } else {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            v2 = (^(v2))
            v2 -= v2
        }
    }
    var v16 bool = false
    fmt.Println(v2, v16)
    fmt.Println(v2, v16)
    for fi17 := 0; fi17 < 18; fi17++ {
        var v18 int8 = 126
        if true {
            v19 := GenS3{F0: (^(v2))}
            _ = v19
        } else {
            v20 := GenS4{F0: 32767, F1: 42, F2: fi17}
            v16 = (!(993 > 933))
            _ = v20
        }
        _ = fi17
        _ = v18
    }
    for fi21 := 0; fi21 < 12; fi21++ {
        v22 := GenS2{F0: 284, F1: "a"}
        _ = fi21
        _ = v22
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
