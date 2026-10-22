package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int8; F1 int8; F2 int16 }
type GenS3 struct { F0 uint16; F1 int }
func genFn0() uint32 {
    var v7 int8 = 1
    if true {
        var v8 int32 = 1
        _ = v8
    } else {
        fmt.Println(v7)
    }
    _ = v7
    return 1
}
func genFn1() uint8 {
    for fi9 := 0; fi9 < 18; fi9++ {
        v10 := []int{(fi9 | fi9), 32767, 128, (fi9 >> ((255) & 63))}
        var v11 int32 = -2147483648
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 int = 9223372036854775807
    if false {
        if false {
            var v13 int64 = 0
            v12 -= (v12 >> ((255) & 63))
            _ = v13
        }
    } else {
        v14 := GenS1{F0: 487}
        _ = v14
    }
    if true {
        fmt.Println(v12)
        v12 += 42
        var v15 int16 = 42
        _ = v15
    } else {
        v12 += (^(v12))
    }
    _ = v12
    return 255
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 bool = true
    var v3 int32 = -2147483648
    var v4 uint32 = 42
    var v5 bool = v2
    var v6 uint64 = 4294967295
    v16 := GenS2{F0: -79, F1: -1, F2: -1}
    v4 += genFn0()
    var v17 uint32 = v4
    switch "" {
    case "hello":
        var v18 bool = v2
        _ = v18
    case "":
        if ((!true) && true) {
            if (v5 || ((!((46 == 187) || v2)) && v2)) {
                v4 -= 2147483647
                var v19 uint = 1
                _ = v19
            } else {
                v2 = (v2 && true)
            }
            for fi20 := 0; fi20 < 18; fi20++ {
                for fi21 := 0; fi21 < 15; fi21++ {
                    var v22 int8 = 27
                    v23 := GenS3{F0: 42, F1: fi20}
                    _ = fi21
                    _ = v22
                    _ = v23
                }
                switch "" {
                case "hello":
                    v24 := GenS1{F0: (v4 % ((662) | 1))}
                    _ = v24
                }
                _ = fi20
            }
            v5 = (-9223372036854775808 < 42)
        } else {
            v4 = (v17 & (v17 % (((v17 / ((0) | 1))) | 1)))
        }
    case "foo bar":
        v25 := GenS3{F0: 2, F1: int(v6)}
        _ = v25
    }
    v26 := []int{127, 127, 9223372036854775807}
    var v27 int32 = v3
    var v28 string = "hello"
    var v29 uint8 = 82
    if v2 {
        v29 = (v29 & v29)
    }
    v4 = genFn0()
    var v30 int16 = int16(v3)
    var v31 int8 = 1
    v32 := GenS3{F0: 2, F1: -2147483648}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
