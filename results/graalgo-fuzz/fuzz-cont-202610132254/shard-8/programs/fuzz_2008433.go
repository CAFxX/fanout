package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 uint64 }

func main() {
    v1 := make(map[string]int)
    var v2 uint64 = 18446744073709551615
    v3 := []int{}
    var v4 int8 = 64
    v5 := make(map[string]string)
    var v6 uint8 = 128
    v7 := v1["foo bar"]
    for fi8 := 0; fi8 < 20; fi8++ {
        var v9 int = fi8
        v3 = append(v3, (fi8 ^ fi8))
        for fi10 := 0; fi10 < 9; fi10++ {
            v9 -= v9
            if true {
                var v11 float32 = 0.1
                _ = v11
            } else {
                if (!(!(false && (int(v2) > (-(fi8)))))) {
                    v4 -= (v4 - (v4 & (-(v4))))
                }
                if ((-1 < -766) && (4294967295 == uint32(v2))) {
                    v4 = v4
                    var v12 int = 2147483647
                    fmt.Println(v4)
                    fmt.Println(v2, v12, v12)
                    _ = v12
                } else {
                    v13 := GenS1{F0: true}
                    var v14 int64 = 1
                    var v15 float32 = 1e10
                    _ = v13
                    _ = v14
                    _ = v15
                }
                _ = len(v1)
            }
            _ = fi10
        }
        _ = fi8
        _ = v9
    }
    v16 := GenS1{F0: true}
    fmt.Println(v4)
    v17 := v1[("" + "")]
    var v18 uint8 = (v6 & (v6 >> ((252) & 63)))
    v17 += v17
    var v19 uint = 42
    v3 = append(v3, 128)
    switch ("" + "\n\t") {
    case "a":
        var v20 uint32 = 1
        _ = v20
    case "hello":
        fmt.Println(v4, v2)
    case "foo bar":
        if true {
            switch "foo bar" {
            case "hello":
                v2 = v2
            case "\n\t":
                v21 := GenS2{F0: 42}
                _ = v21
            case "foo bar":
                var v22 int32 = -2147483648
                _ = v22
            }
        } else {
            _ = len(v3)
            v18 -= v18
        }
    default:
        if true {
            if (("hello" + "") < ("hello" + "foo bar")) {
                v4 -= v4
                v23 := GenS3{F0: 740}
                _ = v23
            } else {
                var v24 uint64 = (v2 + (v2 * v2))
                for fi25 := 0; fi25 < 8; fi25++ {
                    v6 -= (v6 % (((v6 | 255)) | 1))
                    if len(v3) > 0 {
                        v3[0] = v17
                    }
                    _ = fi25
                }
                _ = v24
            }
            v26 := GenS4{F0: (^(v2))}
            _ = v26
        }
    }
    for fi27 := 0; fi27 < 3; fi27++ {
        fmt.Println(v7)
        _ = fi27
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
