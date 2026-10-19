package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16; F1 uint32; F2 string }
type GenS3 struct { F0 bool; F1 bool }

func main() {
    v1 := []uint32{}
    v2 := []string{""}
    var v3 string = ""
    v4 := make(map[string]uint32)
    var v5 uint8 = 128
    var v6 int = -1
    v3 = ((v3 + v3) + "")
    v7 := GenS1{F0: false}
    if (false && false) {
        if len(v2) > 0 {
            v2[0] = ((v3 + ((v3 + v3) + "a")) + "foo bar")
        }
    }
    for fi8 := 0; fi8 < 3; fi8++ {
        if ((18446744073709551615 < 565) || true) {
            var v9 uint64 = 18446744073709551615
            v3 = v3
            v10 := GenS1{F0: false}
            _ = v10
            _ = v9
        } else {
            _ = len(v4)
            for fi11 := 0; fi11 < 8; fi11++ {
                delete(v4, "hello")
                _ = fi11
            }
        }
        var v12 uint = 1
        _ = fi8
        _ = v12
    }
    v3 = (("" + v3) + v3)
    if true {
        for fi13 := 0; fi13 < 18; fi13++ {
            v14 := GenS2{F0: 1000, F1: 610, F2: v3}
            _ = fi13
            _ = v14
        }
    } else {
        for fi15 := 0; fi15 < 9; fi15++ {
            v16 := GenS2{F0: int16(v6), F1: 2, F2: ("a" + v3)}
            var v17 int64 = 9223372036854775807
            _ = fi15
            _ = v16
            _ = v17
        }
        var v18 int16 = 1
        for fi19 := 0; fi19 < 5; fi19++ {
            var v20 float64 = 1.5
            for fi21 := 0; fi21 < 4; fi21++ {
                v22 := GenS3{F0: (true || true), F1: false}
                _ = fi21
                _ = v22
            }
            _ = fi19
            _ = v20
        }
        _ = v18
    }
    if (!(!false)) {
        fmt.Println(v3, v6, v5)
        var v23 int16 = 32767
        _ = v23
    }
    fmt.Println(v3, v3, v6)
    v3 = v3
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
