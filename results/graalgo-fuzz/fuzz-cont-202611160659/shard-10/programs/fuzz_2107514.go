package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint; F1 int }
type GenS3 struct { F0 uint16 }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[uint32]string)
    var v3 int8 = 127
    v4 := []uint32{175}
    var v5 int32 = -1
    v3 -= (-(v3))
    fmt.Println(v3, v3)
    var v6 uint64 = 18446744073709551615
    var v7 bool = true
    if (32768 != -2147483648) {
        if (!false) {
            if (v7 && (v5 <= v5)) {
                for fi8 := 0; fi8 < 16; fi8++ {
                    fmt.Println(v6, v5)
                    _ = fi8
                }
                if (!v7) {
                    var v9 uint8 = 0
                    _ = v9
                }
            } else {
                var v10 uint64 = v6
                var v11 int16 = -32768
                _ = v10
                _ = v11
            }
            var v12 uint32 = 0
            _ = v12
        }
    } else {
        if v7 {
            var v13 uint16 = 42
            v1[2147483647] = (0 == 503)
            _ = v13
        } else {
            v5 -= 42
        }
        var v14 uint8 = 255
        if v7 {
            v3 = v3
            v7 = v7
        } else {
            v15 := GenS1{F0: ("" + "foo bar")}
            var v16 int16 = 343
            _ = v15
            _ = v16
        }
        _ = v14
    }
    fmt.Println(v5, v3)
    for fi17 := 0; fi17 < 20; fi17++ {
        v18 := GenS2{F0: 2, F1: 32767}
        _ = fi17
        _ = v18
    }
    switch -566 {
    case -602:
        if ((v5 > v5) && false) {
            v19 := v2[782]
            v4 = append(v4, 206)
            _ = v19
        }
    default:
        fmt.Println(v5)
    }
    v20 := v2[1]
    var v21 uint32 = 4294967295
    v6 = 42
    if ((v20 + ("foo bar" + v20)) > v20) {
        for fi22 := 0; fi22 < 4; fi22++ {
            var v23 string = v20
            _ = fi22
            _ = v23
        }
        if (705 == 982) {
            v6 -= v6
            var v24 int16 = 874
            _ = v24
        }
        v25 := GenS1{F0: "foo bar"}
        _ = v25
    } else {
        v26 := GenS2{F0: uint(v21), F1: 127}
        fmt.Println(v6)
        _ = v26
    }
    v27 := GenS3{F0: 255}
    for fi28 := 0; fi28 < 15; fi28++ {
        switch (^(fi28)) {
        case 32767:
            var v29 uint = 255
            _ = v29
        case 890:
            v30 := GenS1{F0: v20}
            _ = v30
        case 96:
            var v31 bool = v7
            _ = v31
        default:
            var v32 uint16 = 32767
            _ = v32
        }
        _ = len(v4)
        fmt.Println(v3, v3, v5)
        _ = fi28
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
