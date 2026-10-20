package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16 }
type GenS2 struct { F0 uint16; F1 int32; F2 uint32 }
type GenS3 struct { F0 int; F1 int8; F2 uint }

func main() {
    var v1 int32 = -1
    var v2 uint16 = 64
    var v3 uint16 = v2
    var v4 string = ""
    v5 := make(map[int]bool)
    var v6 uint = 4294967295
    if true {
        _ = len(v5)
    } else {
        var v7 uint64 = 18446744073709551615
        _ = v7
    }
    var v8 int32 = v1
    switch "foo bar" {
    case "hello":
        var v9 string = (v4 + (v4 + ((v4 + (v4 + v4)) + v4)))
        _ = v9
    case "a":
        if (uint8(v6) < 1) {
            var v10 uint8 = 22
            v11 := v5[-694]
            v12 := GenS1{F0: (^(v6)), F1: -184}
            _ = v10
            _ = v11
            _ = v12
        }
    default:
        for fi13 := 0; fi13 < 10; fi13++ {
            v14 := []int{fi13, (fi13 - fi13)}
            var v15 int = (fi13 | 9223372036854775807)
            _ = fi13
            _ = v14
            _ = v15
        }
    }
    for fi16 := 0; fi16 < 4; fi16++ {
        v17 := GenS1{F0: v6, F1: 308}
        v18 := []int{fi16, fi16, (fi16 << (((v6 ^ v6)) & 63)), 1}
        _ = fi16
        _ = v17
        _ = v18
    }
    fmt.Println(v3)
    var v19 int16 = 42
    switch 42 {
    case -939:
        switch v4 {
        case "a":
            if (!(1 >= 0)) {
                v20 := GenS2{F0: v2, F1: v8, F2: 366}
                _ = v20
            } else {
                var v21 int = 32768
                fmt.Println(v21)
                _ = v21
            }
        case "foo bar":
            v1 += v1
        case "\n\t":
            var v22 int8 = -128
            _ = v22
        }
    case 2147483647:
        for fi23 := 0; fi23 < 1; fi23++ {
            if (904 > 1) {
                v3 -= 42
            } else {
                v24 := []int{(fi23 * -1), fi23, fi23, int(v6)}
                _ = v24
            }
            _ = fi23
        }
    default:
        v6 = v6
    }
    var v25 float64 = -2.25
    fmt.Println(v25, v8)
    var v26 float32 = 1e10
    delete(v5, -652)
    switch -9223372036854775808 {
    case 4294967295:
        switch -839 {
        case 9223372036854775807:
            v26 -= float32((v25 - 1e10))
        case -598:
            for fi27 := 0; fi27 < 16; fi27++ {
                var v28 int = 127
                var v29 float32 = (v26 / v26)
                for fi30 := 0; fi30 < 17; fi30++ {
                    switch ("\n\t" + v4) {
                    case "\n\t":
                        var v31 float32 = v29
                        _ = v31
                    case "":
                        fmt.Println(v3, v19, v25)
                    default:
                        v32 := GenS2{F0: (v2 - 42), F1: v1, F2: 2147483648}
                        _ = v32
                    }
                    fmt.Println(v29, v19, fi30)
                    v33 := GenS3{F0: fi27, F1: -20, F2: v6}
                    var v34 float32 = v29
                    _ = fi30
                    _ = v33
                    _ = v34
                }
                _ = fi27
                _ = v28
                _ = v29
            }
        default:
            v35 := []int{}
            _ = v35
        }
    default:
        v36 := GenS1{F0: v6, F1: (v19 + v19)}
        _ = v36
    }
    v37 := GenS1{F0: 286, F1: int16(v1)}
    v38 := v5[292]
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v19
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
