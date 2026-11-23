package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint32; F1 int32; F2 uint16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint64; F1 uint32 }

func main() {
    v1 := []int{32768}
    var v2 float32 = 3.14159
    v3 := make(map[uint32]string)
    v4 := []uint32{}
    var v5 int = 4294967295
    if (-1 == int8(v5)) {
        var v6 uint32 = 2147483647
        _ = v6
    }
    _ = len(v4)
    v7 := GenS1{F0: 1}
    if true {
        v8 := GenS2{F0: 420, F1: 0, F2: 32767}
        _ = v8
    }
    var v9 int32 = 42
    var v10 float64 = -2.25
    var v11 int64 = 0
    v11 = 0
    var v12 uint32 = 731
    for fi13 := 0; fi13 < 20; fi13++ {
        var v14 uint64 = 885
        fmt.Println(v10, v5, v12)
        for fi15 := 0; fi15 < 1; fi15++ {
            var v16 uint16 = 1
            var v17 int8 = -73
            v5 = fi15
            _ = fi15
            _ = v16
            _ = v17
        }
        _ = fi13
        _ = v14
    }
    var v18 uint16 = 32767
    switch "hello" {
    case "foo bar":
        if (("" <= "\n\t") && true) {
            fmt.Println(v5)
            if len(v1) > 0 {
                v1[0] = v5
            }
        } else {
            v18 = v18
        }
    case "hello":
        for fi19 := 0; fi19 < 8; fi19++ {
            v3[521] = "foo bar"
            v20 := GenS1{F0: v18}
            _ = fi19
            _ = v20
        }
    }
    v3[(v12 * v12)] = ("a" + "a")
    v21 := GenS3{F0: false}
    var v22 uint64 = 523
    v12 = v12
    if true {
        var v23 int16 = -32768
        _ = v23
    } else {
        if false {
            var v24 int8 = 0
            for fi25 := 0; fi25 < 10; fi25++ {
                if ("hello" > ("a" + "foo bar")) {
                    var v26 uint = 4294967295
                    var v27 int = v5
                    _ = v26
                    _ = v27
                } else {
                    v28 := GenS3{F0: ((v12 * v12) >= v12)}
                    v18 += (v18 ^ v18)
                    for fi29 := 0; fi29 < 4; fi29++ {
                        v30 := GenS4{F0: (v22 - (v22 << ((4294967295) & 63))), F1: (^(v12))}
                        var v31 uint16 = (^(v18))
                        _ = fi29
                        _ = v30
                        _ = v31
                    }
                    _ = v28
                }
                _ = fi25
            }
            _ = v24
        } else {
            var v32 uint64 = v22
            switch "hello" {
            case "":
                v12 = (v12 << ((2) & 63))
            }
            _ = v32
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v9
}
