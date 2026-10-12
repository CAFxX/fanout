package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 uint64) int {
    var v5 int8 = -1
    fmt.Println(p0, p0)
    _ = p0
    _ = v5
    return 1
}
func genFn1(p0 bool, p1 int64, p2 int64) uint32 {
    switch -9223372036854775808 {
    case 32768:
        fmt.Println(p1)
    case 0:
        switch "" {
        case "foo bar":
            for fi6 := 0; fi6 < 11; fi6++ {
                v7 := []int{fi6, genFn0(973)}
                _ = fi6
                _ = v7
            }
        case "":
            var v8 uint = 255
            _ = v8
        case "hello":
            fmt.Println(p1)
        }
    case -483:
        v9 := GenS1{F0: 0, F1: (p1 >> ((61) & 63))}
        _ = v9
    }
    p2 = p1
    for fi10 := 0; fi10 < 17; fi10++ {
        switch "hello" {
        case "hello":
            p0 = p0
        case "":
            var v11 int64 = p1
            _ = v11
        }
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = p2
    return 65535
}

func main() {
    var v1 int32 = 42
    var v2 uint64 = 18446744073709551615
    var v3 uint32 = 26
    var v4 uint8 = 50
    v12 := []int{-883, -963, 127, 128}
    for fi13 := 0; fi13 < 16; fi13++ {
        var v14 uint8 = v4
        _ = fi13
        _ = v14
    }
    v15 := GenS2{F0: genFn1(false, 1, int64(v3))}
    for fi16 := 0; fi16 < 16; fi16++ {
        v17 := GenS2{F0: genFn1(false, 396, -1)}
        switch genFn0((v2 + 4294967295)) {
        case -599:
            v2 = v2
        case 298:
            v3 -= genFn1(false, 1, 9223372036854775807)
        }
        for fi18 := 0; fi18 < 13; fi18++ {
            fmt.Println(v1)
            if true {
                fmt.Println(v2, v4)
                var v19 bool = false
                v1 -= v1
                fmt.Println(v4)
                _ = v19
            }
            var v20 float32 = (3.14159 - 1.5)
            _ = fi18
            _ = v20
        }
        _ = fi16
        _ = v17
    }
    v1 += 978
    var v21 bool = true
    v22 := GenS1{F0: 1, F1: -1}
    fmt.Println(v3)
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v15
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
