package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 uint16; F1 int8; F2 uint16 }

func main() {
    var v1 uint32 = 1
    v2 := make(map[int]uint32)
    var v3 int8 = 42
    v4 := make(map[int]string)
    for fi5 := 0; fi5 < 15; fi5++ {
        var v6 int64 = 0
        _ = fi5
        _ = v6
    }
    switch 42 {
    case 2147483647:
        v7 := GenS1{F0: 255}
        _ = v7
    case -479:
        for fi8 := 0; fi8 < 7; fi8++ {
            switch -9223372036854775808 {
            case 32768:
                switch (fi8 & fi8) {
                case 32768:
                    fmt.Println(fi8, fi8)
                case 9223372036854775807:
                    switch 0 {
                    case 9223372036854775807:
                        v9 := []int{}
                        _ = v9
                    case 128:
                        v10 := GenS1{F0: 32767}
                        _ = v10
                    case 4294967295:
                        switch "hello" {
                        case "":
                            v11 := []int{}
                            _ = v11
                        case "foo bar":
                            v1 = v1
                        }
                    default:
                        v1 = 65535
                    }
                default:
                    v12 := GenS2{F0: 4294967295, F1: 165}
                    _ = v12
                }
            case -612:
                var v13 uint16 = 2
                _ = v13
            default:
                fmt.Println(v3)
            }
            _ = fi8
        }
    case 128:
        if true {
            v14 := GenS3{F0: 1, F1: (v3 + (v3 >> ((507) & 63))), F2: 32767}
            v3 -= v3
            var v15 uint32 = 519
            _ = v14
            _ = v15
        }
    }
    v3 = -128
    for fi16 := 0; fi16 < 11; fi16++ {
        var v17 bool = (!true)
        var v18 float64 = (3.14159 * (0.1 * (0.0 + 1e10)))
        _ = fi16
        _ = v17
        _ = v18
    }
    switch -9223372036854775808 {
    case 32767:
        v3 = v3
    case 127:
        var v19 uint = 18446744073709551615
        _ = v19
    default:
        var v20 int64 = -1
        _ = v20
    }
    var v21 uint32 = v1
    var v22 int = -2147483648
    v3 = v3
    switch "hello" {
    case "":
        v23 := []int{(v22 ^ v22), v22}
        _ = v23
    default:
        v21 = (v21 ^ (v21 ^ (v1 - v1)))
    }
    var v24 float32 = 1.5
    delete(v4, (v22 | (v22 * v22)))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v24
    _ = v3
    _ = v4
}
