package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 int32 }

func main() {
    var v1 float64 = 1.5
    var v2 bool = false
    var v3 int = 265
    var v4 uint64 = 254
    fmt.Println(v1, v1)
    for fi5 := 0; fi5 < 7; fi5++ {
        for fi6 := 0; fi6 < 19; fi6++ {
            var v7 uint8 = 201
            _ = fi6
            _ = v7
        }
        _ = fi5
    }
    var v8 int = 970
    if v2 {
        for fi9 := 0; fi9 < 4; fi9++ {
            var v10 uint64 = (^(v4))
            _ = fi9
            _ = v10
        }
        var v11 uint64 = (v4 << ((472) & 63))
        if (true && (!(v2 && (!(32767 != 2))))) {
            var v12 uint16 = 32767
            fmt.Println(v1)
            if v2 {
                v13 := GenS1{F0: (v12 ^ v12), F1: "a"}
                if (1000 < int16(v8)) {
                    v11 -= 1
                    v14 := []int{v3}
                    var v15 int32 = 2147483647
                    _ = v14
                    _ = v15
                }
                _ = v13
            }
            _ = v12
        }
        _ = v11
    } else {
        var v16 string = "hello"
        v17 := GenS2{F0: 42, F1: (v16 + v16)}
        v18 := GenS2{F0: 292, F1: v16}
        _ = v16
        _ = v17
        _ = v18
    }
    if true {
        var v19 float32 = 0.1
        var v20 uint = 255
        _ = v19
        _ = v20
    }
    var v21 int32 = 1
    switch "hello" {
    case "":
        v8 += v8
    case "hello":
        v22 := GenS3{F0: 906}
        _ = v22
    }
    if true {
        if v2 {
            for fi23 := 0; fi23 < 5; fi23++ {
                v8 = (v8 - 4294967295)
                _ = fi23
            }
            v24 := GenS1{F0: 65535, F1: (("\n\t" + ("\n\t" + ("hello" + "a"))) + ("\n\t" + "hello"))}
            var v25 uint64 = (v4 << ((1) & 63))
            _ = v24
            _ = v25
        }
    }
    var v26 uint = 2
    v27 := []int{(-(v8)), int(v4), -2147483648, (v3 - v3)}
    var v28 uint = v26
    var v29 float32 = 1e10
    for fi30 := 0; fi30 < 4; fi30++ {
        v2 = (!false)
        switch "hello" {
        case "a":
            if len(v27) > 0 {
                v27[0] = 781
            }
        }
        _ = fi30
    }
    v31 := GenS3{F0: (-(v21))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v27:", len(v27), v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v21
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v31
    _ = v4
    _ = v8
}
