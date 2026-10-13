package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 uint32; F1 int }
func genFn0(p0 uint64, p1 uint16) uint {
    v5 := GenS1{F0: (("" + "hello") + ("foo bar" + "a")), F1: 2}
    var v6 int = 0
    p0 -= (p0 / (((p0 & p0)) | 1))
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 18446744073709551615
}

func main() {
    var v1 uint64 = 42
    v2 := []int{119}
    v3 := make(map[string]uint32)
    var v4 uint16 = 2
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if true {
        var v7 string = "hello"
        for fi8 := 0; fi8 < 17; fi8++ {
            var v9 int16 = -1
            fmt.Println(v1, v4)
            _ = fi8
            _ = v9
        }
        for fi10 := 0; fi10 < 11; fi10++ {
            switch "\n\t" {
            case "\n\t":
                var v11 int64 = 9223372036854775807
                _ = v11
            }
            var v12 uint32 = 798
            _ = fi10
            _ = v12
        }
        _ = v7
    } else {
        for fi13 := 0; fi13 < 5; fi13++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v14 uint = 4294967295
            var v15 uint8 = uint8(v4)
            _ = fi13
            _ = v14
            _ = v15
        }
        v16 := GenS2{F0: 548, F1: 42}
        _ = v16
    }
    var v17 int16 = 0
    v18 := GenS2{F0: 4294967295, F1: 932}
    for fi19 := 0; fi19 < 2; fi19++ {
        fmt.Println(fi19, v1, v1)
        v17 += (v17 % (((v17 % ((1000) | 1))) | 1))
        _ = fi19
    }
    var v20 uint8 = 42
    for fi21 := 0; fi21 < 7; fi21++ {
        var v22 uint8 = v20
        var v23 int8 = 27
        _ = fi21
        _ = v22
        _ = v23
    }
    v1 = uint64(v4)
    v1 -= 160
    v17 = int16(v20)
    for fi24 := 0; fi24 < 10; fi24++ {
        v25 := GenS1{F0: ("" + ("" + "")), F1: 2147483647}
        if (1 >= (^(v4))) {
            var v26 int16 = v17
            fmt.Println(v1)
            _ = v26
        }
        _ = len(v2)
        _ = fi24
        _ = v25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
