package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32; F2 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint64 }

func main() {
    var v1 uint = 4294967295
    var v2 uint16 = 32767
    v3 := make(map[int]int)
    var v4 bool = false
    v5 := []int{3, 879, 0, 32768}
    var v6 uint64 = 4294967295
    v1 += 2
    switch -245 {
    case 2147483647:
        v7 := GenS1{F0: 65535, F1: 1, F2: 240}
        _ = v7
    case -2147483648:
        var v8 uint8 = 0
        _ = v8
    case 666:
        v1 -= (v1 / ((42) | 1))
    default:
        switch int(v1) {
        case -9223372036854775808:
            fmt.Println(v6, v2)
        case 68:
            delete(v3, 32767)
        }
    }
    if false {
        if v4 {
            var v9 float64 = (0.0 / 1e10)
            _ = v9
        } else {
            if len(v5) > 0 {
                v5[0] = -153
            }
        }
        for fi10 := 0; fi10 < 1; fi10++ {
            switch "hello" {
            case "a":
                switch (fi10 % ((128) | 1)) {
                case 9223372036854775807:
                    v1 = (^(v1))
                }
            default:
                v11 := GenS1{F0: 2, F1: 2147483648, F2: int16(v6)}
                _ = v11
            }
            switch (fi10 / (((fi10 ^ (fi10 & fi10))) | 1)) {
            case 42:
                v2 = 42
            case 127:
                if len(v5) > 0 {
                    v5[0] = fi10
                }
            case 128:
                var v12 int16 = 810
                _ = v12
            default:
                switch (("foo bar" + "a") + ("\n\t" + "hello")) {
                case "foo bar":
                    if (v4 || v4) {
                        for fi13 := 0; fi13 < 3; fi13++ {
                            v4 = true
                            var v14 uint32 = 1
                            _ = fi13
                            _ = v14
                        }
                        v3[fi10] = fi10
                    }
                }
            }
            fmt.Println(v4, v4, v4)
            _ = fi10
        }
        var v15 uint8 = 255
        _ = v15
    }
    if v4 {
        for fi16 := 0; fi16 < 9; fi16++ {
            var v17 int8 = 0
            v18 := v3[fi16]
            _ = fi16
            _ = v17
            _ = v18
        }
        if len(v5) > 0 {
            v5[0] = 32768
        }
    } else {
        for fi19 := 0; fi19 < 9; fi19++ {
            var v20 uint64 = 18446744073709551615
            var v21 int32 = -1
            if v4 {
                v22 := v3[fi19]
                v23 := GenS2{F0: -313}
                v2 = (v2 & 42)
                _ = v22
                _ = v23
            }
            _ = fi19
            _ = v20
            _ = v21
        }
    }
    var v24 int8 = 100
    v6 = v6
    var v25 int16 = -32768
    if v4 {
        switch (("a" + "\n\t") + "foo bar") {
        case "hello":
            v24 = v24
        case "foo bar":
            fmt.Println(v2)
        case "a":
            switch "hello" {
            case "hello":
                v6 = (^(v6))
            }
        default:
            v24 -= -58
        }
        v6 -= 42
        var v26 bool = v4
        _ = v26
    }
    v2 = v2
    v27 := GenS2{F0: -1}
    v28 := GenS3{F0: (v1 - 109)}
    v29 := GenS4{F0: (^(v6))}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
