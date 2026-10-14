package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32; F1 int64 }
type GenS3 struct { F0 int16 }

func main() {
    var v1 uint8 = 183
    var v2 uint64 = 0
    v3 := make(map[uint32]string)
    var v4 uint16 = 255
    v5 := []string{"foo bar", "foo bar", "a", "a"}
    v6 := []uint32{380, 2147483647, 0, 65535}
    v1 -= v1
    v4 += v4
    if true {
        v4 = v4
    } else {
        var v7 uint64 = (v2 >> ((65535) & 63))
        _ = v7
    }
    switch -1 {
    case 0:
        for fi8 := 0; fi8 < 3; fi8++ {
            for fi9 := 0; fi9 < 2; fi9++ {
                v2 += v2
                _ = fi9
            }
            for fi10 := 0; fi10 < 16; fi10++ {
                if true {
                    v1 = v1
                    switch fi10 {
                    case -669:
                        if (0.1 > float64(1e10)) {
                            v11 := GenS1{F0: int(v4)}
                            _ = v11
                        }
                    }
                    var v12 uint8 = 42
                    var v13 uint8 = v12
                    _ = v12
                    _ = v13
                } else {
                    v2 = 18446744073709551615
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                }
                var v14 int16 = 32767
                _ = fi10
                _ = v14
            }
            if true {
                fmt.Println(fi8)
                v4 += v4
            }
            _ = fi8
        }
    case -2147483648:
        switch ("foo bar" + ("a" + "a")) {
        case "foo bar":
            var v15 int8 = -1
            _ = v15
        }
    default:
        v16 := GenS2{F0: 177, F1: -1}
        _ = v16
    }
    v4 = 65535
    var v17 uint64 = 4294967295
    delete(v3, 710)
    v18 := GenS3{F0: -54}
    var v19 uint = 85
    if true {
        fmt.Println(v4, v17)
    }
    var v20 uint16 = uint16(v2)
    v2 -= v17
    delete(v3, 65535)
    v17 = (^(v17))
    v2 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
