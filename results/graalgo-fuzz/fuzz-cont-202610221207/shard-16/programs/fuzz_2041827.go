package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 int16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32 }
func genFn0(p0 int64, p1 uint32, p2 int32) int32 {
    fmt.Println(p1)
    var v5 uint16 = 32767
    if ((1 >= -1) && (!((1e10 * float32(1.5)) == 1e10))) {
        if true {
            for fi6 := 0; fi6 < 17; fi6++ {
                var v7 int16 = -32768
                p1 = (^(p1))
                _ = fi6
                _ = v7
            }
            v8 := GenS1{F0: uint8(p1), F1: "", F2: 32767}
            _ = v8
        }
        v9 := GenS1{F0: 2, F1: "\n\t", F2: 1}
        v10 := []int{42, 32768}
        _ = v10
        _ = v9
    } else {
        p2 -= (p2 * p2)
        v11 := GenS1{F0: 128, F1: "a", F2: 32767}
        switch (("" + "\n\t") + "foo bar") {
        case "hello":
            for fi12 := 0; fi12 < 20; fi12++ {
                var v13 int16 = -1
                _ = fi12
                _ = v13
            }
        }
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 2147483647
}
func genFn1(p0 int) int8 {
    var v14 uint16 = 32767
    switch "foo bar" {
    case "foo bar":
        var v15 uint64 = 18446744073709551615
        _ = v15
    case "a":
        for fi16 := 0; fi16 < 13; fi16++ {
            var v17 string = "hello"
            var v18 uint64 = 2
            _ = fi16
            _ = v17
            _ = v18
        }
    }
    _ = p0
    _ = v14
    return 1
}

func main() {
    var v1 string = ""
    v2 := []uint32{0, 374, 0}
    var v3 uint8 = 127
    var v4 int = -1
    for fi19 := 0; fi19 < 8; fi19++ {
        if true {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            v4 = fi19
            v4 = fi19
        }
        v1 = v1
        _ = fi19
    }
    if true {
        var v20 uint64 = 18446744073709551615
        _ = v20
    } else {
        v21 := GenS2{F0: 1}
        fmt.Println(v4, v3)
        _ = v21
    }
    fmt.Println(v1, v1, v4)
    if (3.14159 < ((0.1 / 1.5) / (0.0 + 0.1))) {
        if true {
            v22 := GenS2{F0: 128}
            _ = v22
        }
        switch v1 {
        case "foo bar":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case "":
            v23 := GenS2{F0: 2}
            _ = v23
        default:
            var v24 int8 = genFn1((v4 & v4))
            _ = v24
        }
    } else {
        var v25 int8 = -99
        _ = v25
    }
    var v26 int32 = genFn0(9223372036854775807, 1, -1)
    v27 := GenS2{F0: 255}
    v2 = append(v2, 65535)
    var v28 int16 = 353
    if true {
        var v29 int8 = 100
        _ = v29
    }
    for fi30 := 0; fi30 < 3; fi30++ {
        v31 := GenS1{F0: v3, F1: ((v1 + "a") + ((v1 + v1) + v1)), F2: (v28 ^ (^(v28)))}
        _ = fi30
        _ = v31
    }
    v3 = v3
    v32 := GenS3{F0: 2147483647}
    v2 = append(v2, 4294967295)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v32
    _ = v4
}
