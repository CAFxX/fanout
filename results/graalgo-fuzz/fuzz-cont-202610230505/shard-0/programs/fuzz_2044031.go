package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 int8; F2 uint32 }
func genFn0(p0 uint64, p1 uint64) uint64 {
    var v5 int32 = 272
    v6 := []int{-568, 4294967295}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 1
}

func main() {
    var v1 uint16 = 856
    v2 := []int{42}
    v3 := make(map[int]uint32)
    var v4 bool = false
    v7 := GenS1{F0: 42}
    v4 = (1 <= -232)
    v4 = (!(float64(0.1) == 3.14159))
    fmt.Println(v1, v1)
    if v4 {
        v4 = (!false)
    } else {
        switch ("hello" + ("hello" + "hello")) {
        case "":
            var v8 int = 265
            _ = v8
        case "foo bar":
            v1 = 255
        default:
            switch "foo bar" {
            case "a":
                fmt.Println(v1, v1)
            case "hello":
                v9 := GenS2{F0: 0, F1: 1, F2: 42}
                _ = v9
            case "foo bar":
                for fi10 := 0; fi10 < 14; fi10++ {
                    for fi11 := 0; fi11 < 15; fi11++ {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        v12 := GenS1{F0: 32767}
                        var v13 uint32 = 4294967295
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                    var v14 bool = v4
                    _ = fi10
                    _ = v14
                }
            default:
                var v15 int16 = 1
                _ = v15
            }
        }
        var v16 bool = false
        _ = v16
    }
    delete(v3, -1)
    if false {
        v4 = v4
    }
    if v4 {
        v17 := GenS2{F0: 0, F1: 1, F2: 0}
        fmt.Println(v1, v4)
        for fi18 := 0; fi18 < 12; fi18++ {
            switch "hello" {
            case "foo bar":
                if v4 {
                    _ = len(v2)
                    v4 = ((^(v1)) != (v1 % ((255) | 1)))
                }
            case "hello":
                v19 := GenS2{F0: -128, F1: 34, F2: 2147483648}
                _ = v19
            default:
                v3[(-(fi18))] = 391
            }
            if v4 {
                var v20 int16 = -1
                var v21 int64 = -9223372036854775808
                v21 += (v21 / (((v21 & v21)) | 1))
                _ = v20
                _ = v21
            }
            _ = fi18
        }
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v7
}
