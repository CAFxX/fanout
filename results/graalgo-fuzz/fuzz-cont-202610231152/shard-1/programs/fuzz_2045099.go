package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 int16; F1 int64 }

func main() {
    var v1 int16 = 32767
    var v2 uint = 659
    var v3 uint32 = 0
    var v4 uint16 = 65535
    for fi5 := 0; fi5 < 2; fi5++ {
        fmt.Println(v1, v4)
        for fi6 := 0; fi6 < 18; fi6++ {
            v1 = v1
            if false {
                v4 = 847
                v7 := []int{(-(fi5))}
                var v8 float32 = 1e10
                _ = v7
                _ = v8
            } else {
                var v9 uint64 = 4294967295
                _ = v9
            }
            var v10 bool = false
            _ = fi6
            _ = v10
        }
        _ = fi5
    }
    switch ("foo bar" + ("hello" + "a")) {
    case "":
        var v11 int32 = 823
        _ = v11
    case "hello":
        v12 := []int{-653, 253}
        _ = v12
    case "a":
        v13 := []int{4294967295, -270}
        _ = v13
    default:
        v14 := GenS1{F0: 1, F1: uint32(v4)}
        _ = v14
    }
    var v15 bool = true
    var v16 uint64 = 18446744073709551615
    var v17 int = 517
    for fi18 := 0; fi18 < 1; fi18++ {
        var v19 uint64 = 275
        for fi20 := 0; fi20 < 9; fi20++ {
            v4 = v4
            v3 += v3
            v4 -= v4
            _ = fi20
        }
        _ = fi18
        _ = v19
    }
    var v21 int = (v17 | v17)
    fmt.Println(v15)
    v17 = 1
    v22 := []int{(v21 * v21)}
    fmt.Println(v4, v4)
    v16 = (v16 ^ v16)
    for fi23 := 0; fi23 < 16; fi23++ {
        var v24 uint64 = (v16 ^ v16)
        switch ("a" + "a") {
        case "":
            var v25 bool = false
            _ = v25
        default:
            var v26 uint = v2
            _ = v26
        }
        _ = fi23
        _ = v24
    }
    var v27 uint = 0
    for fi28 := 0; fi28 < 7; fi28++ {
        for fi29 := 0; fi29 < 10; fi29++ {
            v2 = v27
            fmt.Println(v4)
            if false {
                var v30 int16 = (v1 % ((int16(fi28)) | 1))
                switch "foo bar" {
                case "a":
                    v1 += v1
                default:
                    fmt.Println(v1)
                }
                _ = v30
            }
            _ = fi29
        }
        v31 := GenS2{F0: (v1 ^ (v1 / (((v1 | int16(v2))) | 1))), F1: 0}
        var v32 string = ""
        _ = fi28
        _ = v31
        _ = v32
    }
    v17 += v21
    v4 = v4
    v33 := GenS1{F0: uint8(v4), F1: v3}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v27
    _ = v3
    _ = v33
    _ = v4
}
