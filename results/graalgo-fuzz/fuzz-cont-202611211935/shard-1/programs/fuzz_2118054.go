package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 uint64) int16 {
    if (true || true) {
        v6 := []int{2147483647}
        if true {
            var v7 uint8 = 67
            _ = v7
        } else {
            var v8 uint = 4294967295
            p0 = 0
            _ = v8
        }
        _ = v6
    }
    v9 := []int{0}
    _ = p0
    _ = v9
    return -568
}

func main() {
    v1 := make(map[string]int)
    var v2 bool = (32767 != -1)
    var v3 uint16 = 559
    var v4 bool = v2
    var v5 uint8 = 128
    for fi10 := 0; fi10 < 14; fi10++ {
        delete(v1, "")
        switch "foo bar" {
        case "\n\t":
            switch (fi10 * fi10) {
            case 128:
                var v11 int16 = -32768
                _ = v11
            case 4294967295:
                v4 = ((fi10 / ((2147483647) | 1)) < 32767)
            case 32767:
                var v12 uint64 = 4294967295
                _ = v12
            }
        case "hello":
            v13 := GenS1{F0: 0, F1: 936}
            _ = v13
        default:
            v2 = v4
        }
        _ = fi10
    }
    v3 -= v3
    var v14 float32 = 0.0
    v15 := GenS2{F0: 0}
    for fi16 := 0; fi16 < 3; fi16++ {
        var v17 uint8 = 0
        v5 += (v17 << ((uint(v3)) & 63))
        var v18 uint16 = (v3 << ((0) & 63))
        _ = fi16
        _ = v17
        _ = v18
    }
    v19 := []int{-1, 252, 4294967295}
    var v20 int16 = -271
    v14 = 0.1
    switch "hello" {
    case "foo bar":
        var v21 uint8 = uint8(v3)
        _ = v21
    case "hello":
        v20 = v20
    }
    var v22 uint8 = 128
    var v23 int16 = genFn0(559)
    v4 = ("foo bar" <= "hello")
    var v24 float64 = 1e10
    switch "foo bar" {
    case "\n\t":
        for fi25 := 0; fi25 < 9; fi25++ {
            if (0.1 <= v24) {
                for fi26 := 0; fi26 < 20; fi26++ {
                    var v27 float64 = v24
                    fmt.Println(v4, v14)
                    var v28 uint32 = 820
                    _ = fi26
                    _ = v27
                    _ = v28
                }
                if ((v2 && (2 != 0)) && true) {
                    var v29 uint64 = 1
                    var v30 uint64 = (v29 >> ((563) & 63))
                    var v31 uint16 = (^(v3))
                    _ = v29
                    _ = v30
                    _ = v31
                }
            }
            _ = fi25
        }
    default:
        var v32 uint16 = 65535
        _ = v32
    }
    if len(v19) > 0 {
        v19[0] = 32768
    }
    _ = len(v19)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
