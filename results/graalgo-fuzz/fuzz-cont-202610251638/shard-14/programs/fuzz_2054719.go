package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int; F2 uint64 }
type GenS2 struct { F0 int8; F1 bool; F2 bool }
type GenS3 struct { F0 uint32 }
func genFn0() uint16 {
    var v7 uint64 = 42
    v8 := []int{9223372036854775807}
    _ = len(v8)
    _ = v7
    _ = v8
    return 0
}
func genFn1() uint16 {
    for fi9 := 0; fi9 < 17; fi9++ {
        var v10 string = "\n\t"
        v10 = "foo bar"
        v10 = (("hello" + "") + v10)
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 8; fi11++ {
        var v12 int64 = 948
        _ = fi11
        _ = v12
    }
    return 2
}

func main() {
    var v1 int8 = 1
    v2 := make(map[string]int)
    var v3 string = "a"
    v4 := make(map[uint32]string)
    var v5 float32 = (1e10 - 0.1)
    v6 := make(map[uint32]bool)
    if (!(2147483647 == 42)) {
        for fi13 := 0; fi13 < 17; fi13++ {
            var v14 uint16 = genFn0()
            if true {
                for fi15 := 0; fi15 < 1; fi15++ {
                    v3 = (v3 + (v3 + v3))
                    v16 := []int{}
                    v1 = v1
                    _ = fi15
                    _ = v16
                }
            }
            var v17 int64 = 42
            _ = fi13
            _ = v14
            _ = v17
        }
        var v18 int8 = 1
        var v19 uint32 = 450
        _ = v18
        _ = v19
    }
    switch v3 {
    case "hello":
        switch v3 {
        case "foo bar":
            v20 := GenS1{F0: 385, F1: -681, F2: 42}
            _ = v20
        case "a":
            v3 = ((v3 + ((v3 + "\n\t") + v3)) + v3)
        }
    }
    if (true || false) {
        if (!false) {
            for fi21 := 0; fi21 < 14; fi21++ {
                v5 = v5
                var v22 float64 = 0.0
                for fi23 := 0; fi23 < 6; fi23++ {
                    var v24 int32 = -244
                    _ = fi23
                    _ = v24
                }
                _ = fi21
                _ = v22
            }
            v3 = (v3 + v3)
        }
    }
    v5 += (1.5 - float32(-2.25))
    var v25 uint16 = 221
    v1 = v1
    var v26 int32 = 42
    v27 := GenS2{F0: v1, F1: true, F2: false}
    v4[2147483647] = ((((v3 + (v3 + v3)) + "a") + ("a" + (("hello" + (v3 + "\n\t")) + (v3 + v3)))) + v3)
    if ((v25 << ((1) & 63)) == (v25 / (((v25 & 2)) | 1))) {
        for fi28 := 0; fi28 < 14; fi28++ {
            fmt.Println(v5, v25, v3)
            _ = fi28
        }
        for fi29 := 0; fi29 < 19; fi29++ {
            v30 := GenS3{F0: 893}
            _ = fi29
            _ = v30
        }
    }
    v26 -= (v26 >> ((0) & 63))
    var v31 uint32 = 0
    v32 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
