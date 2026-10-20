package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int16; F1 int32 }
type GenS4 struct { F0 int64; F1 int; F2 uint32 }
func genFn0(p0 int32, p1 uint16) uint32 {
    fmt.Println(p0, p0)
    var v7 int16 = 32767
    var v8 uint = 912
    for fi9 := 0; fi9 < 4; fi9++ {
        fmt.Println(v8, v7)
        if true {
            for fi10 := 0; fi10 < 15; fi10++ {
                v11 := GenS1{F0: 1, F1: -128}
                _ = fi10
                _ = v11
            }
        } else {
            v12 := GenS1{F0: 42, F1: int8(p1)}
            for fi13 := 0; fi13 < 12; fi13++ {
                v14 := GenS2{F0: 90}
                var v15 int8 = 42
                _ = fi13
                _ = v14
                _ = v15
            }
            _ = v12
        }
        p0 += 42
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 42
}

func main() {
    v1 := []int{1, 32767}
    v2 := make(map[uint32]uint32)
    var v3 uint16 = 65535
    v4 := []string{"a"}
    v5 := []string{("foo bar" + "a")}
    v6 := []int{}
    fmt.Println(v3)
    _ = len(v1)
    if true {
        var v16 int8 = -43
        var v17 uint32 = 4294967295
        v18 := GenS3{F0: 0, F1: 931}
        _ = v16
        _ = v17
        _ = v18
    } else {
        if ((0.1 + 3.14159) >= 0.1) {
            var v19 string = "a"
            v2[1] = 747
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = v19
        }
        v20 := GenS4{F0: -600, F1: 2147483647, F2: 42}
        _ = v20
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v21 float64 = 1.5
    var v22 int8 = -1
    v4 = append(v4, "")
    fmt.Println(v22, v21, v3)
    var v23 string = ""
    if (0 != 595) {
        if true {
            delete(v2, 65535)
            var v24 int = 128
            _ = v24
        }
    }
    var v25 uint16 = 255
    for fi26 := 0; fi26 < 12; fi26++ {
        if true {
            var v27 int32 = 2147483647
            _ = v27
        }
        v28 := GenS2{F0: 127}
        _ = fi26
        _ = v28
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
