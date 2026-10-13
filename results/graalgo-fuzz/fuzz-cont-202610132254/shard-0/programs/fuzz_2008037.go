package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 bool; F1 int }
type GenS3 struct { F0 int8; F1 uint32; F2 bool }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint16 = 2
    var v3 int = 32768
    v4 := GenS1{F0: -9223372036854775808, F1: uint(v2)}
    v5 := GenS2{F0: (2 >= 2), F1: (v3 & (^(v3)))}
    v6 := []int{42}
    v7 := GenS1{F0: 1, F1: 65535}
    switch -692 {
    case 127:
        v3 = int(v2)
    default:
        v2 += v2
    }
    v8 := GenS3{F0: 100, F1: 65535, F2: false}
    v9 := v1[2147483647]
    if true {
        if (127 != 127) {
            var v10 uint16 = (v2 | 552)
            var v11 int8 = 72
            _ = v10
            _ = v11
        }
        var v12 uint8 = 1
        fmt.Println(v2, v12)
        _ = v12
    } else {
        if false {
            switch 9223372036854775807 {
            case -9223372036854775808:
                var v13 uint32 = 2
                _ = v13
            case 32768:
                var v14 uint32 = uint32(v2)
                _ = v14
            case -2147483648:
                v1[77] = 964
            }
            var v15 int = (v3 / ((v3) | 1))
            if ((127 == 42) || (!false)) {
                _ = len(v6)
                var v16 uint32 = 4294967295
                for fi17 := 0; fi17 < 5; fi17++ {
                    var v18 string = "\n\t"
                    v1[v16] = fi17
                    _ = fi17
                    _ = v18
                }
                _ = v16
            }
            _ = v15
        }
        fmt.Println(v9)
    }
    if (-1 >= 1000) {
        v9 -= v3
    } else {
        if (-32768 > 42) {
            v2 = v2
            v3 = 677
        } else {
            v3 = (v9 >> ((821) & 63))
            for fi19 := 0; fi19 < 5; fi19++ {
                var v20 string = "hello"
                fmt.Println(fi19, v2, v2)
                _ = fi19
                _ = v20
            }
            var v21 float32 = 0.1
            _ = v21
        }
        v22 := GenS1{F0: -862, F1: 18446744073709551615}
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = v22
    }
    for fi23 := 0; fi23 < 15; fi23++ {
        switch (v3 & (-(v9))) {
        case 127:
            switch 32768 {
            case -9223372036854775808:
                for fi24 := 0; fi24 < 5; fi24++ {
                    fmt.Println(v9)
                    _ = fi24
                }
            case 902:
                v2 = (v2 % (((v2 & (v2 * (v2 + (v2 / (((v2 * 32767)) | 1)))))) | 1))
            }
        case 178:
            v25 := GenS1{F0: -551, F1: 2}
            _ = v25
        }
        _ = fi23
    }
    fmt.Println(v9, v9)
    fmt.Println(v2)
    var v26 uint16 = 654
    for fi27 := 0; fi27 < 9; fi27++ {
        fmt.Println(v2)
        _ = fi27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
