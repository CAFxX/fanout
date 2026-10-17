package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 uint64; F2 uint8 }
type GenS3 struct { F0 uint16; F1 int64 }
type GenS4 struct { F0 uint16 }
func genFn0(p0 string, p1 uint32) int32 {
    var v4 uint16 = 42
    v5 := GenS1{F0: -128}
    v4 = v4
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 1
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []string{}
    v3 := []int{-795, -546, 904}
    var v6 bool = (1e10 > 1.5)
    switch "foo bar" {
    case "hello":
        v6 = v6
    case "":
        v7 := GenS2{F0: 65535, F1: 378, F2: 27}
        _ = v7
    default:
        if true {
            var v8 uint8 = 250
            _ = v8
        } else {
            var v9 uint32 = 2
            if (2 >= 2) {
                var v10 uint64 = uint64(v9)
                for fi11 := 0; fi11 < 10; fi11++ {
                    fmt.Println(v10)
                    var v12 uint = uint(v9)
                    v1[42] = "\n\t"
                    _ = fi11
                    _ = v12
                }
                _ = v10
            } else {
                if v6 {
                    v6 = false
                } else {
                    v13 := GenS3{F0: 255, F1: 9223372036854775807}
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = v13
                }
                v6 = v6
            }
            _ = len(v1)
            _ = v9
        }
    }
    if v6 {
        fmt.Println(v6)
        var v14 int32 = 1
        _ = v14
    } else {
        v2 = append(v2, "foo bar")
        fmt.Println(v6)
        var v15 int16 = 984
        _ = v15
    }
    var v16 uint64 = 1
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v17 := GenS1{F0: -115}
    if (0 <= -41) {
        var v18 uint64 = v16
        var v19 uint64 = 399
        _ = v18
        _ = v19
    } else {
        var v20 string = ("" + "a")
        _ = v20
    }
    v21 := GenS3{F0: 568, F1: -1}
    if (1 <= 2147483647) {
        fmt.Println(v16, v6)
        if (v6 || v6) {
            v22 := GenS4{F0: 65535}
            var v23 bool = v6
            _ = v22
            _ = v23
        } else {
            var v24 uint8 = 1
            _ = v24
        }
        if v6 {
            for fi25 := 0; fi25 < 8; fi25++ {
                fmt.Println(v16)
                v16 = (v16 << ((1) & 63))
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi25
            }
            fmt.Println(v6)
            fmt.Println(v16, v6)
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v6
}
