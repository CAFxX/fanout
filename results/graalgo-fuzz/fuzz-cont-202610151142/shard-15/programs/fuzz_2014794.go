package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 int32; F1 int64; F2 int16 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]string)
    var v4 uint64 = 605
    v5 := []int{32768, 128, -2147483648}
    var v6 int16 = 0
    var v7 bool = (uint8(v4) > 88)
    v7 = ((v7 || false) && true)
    var v8 uint = 255
    v6 += (v6 << (((v8 + (v8 ^ v8))) & 63))
    fmt.Println(v6, v8, v6)
    v8 -= v8
    if ((v6 << ((1) & 63)) >= (v6 * v6)) {
        v9 := GenS1{F0: 65535}
        _ = v9
    }
    v7 = v7
    if ((!v7) || false) {
        var v10 int16 = (^(v6))
        switch -496 {
        case 2147483647:
            switch 2147483647 {
            case 2147483647:
                v11 := GenS2{F0: uint32(v8), F1: v4}
                _ = v11
            }
        case 32768:
            for fi12 := 0; fi12 < 6; fi12++ {
                switch fi12 {
                case 32768:
                    if ((fi12 ^ -500) == (fi12 * (fi12 | -2147483648))) {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    }
                }
                fmt.Println(v8, v4)
                _ = fi12
            }
        case 32767:
            if v7 {
                v7 = true
                v5 = append(v5, 85)
                var v13 int8 = 42
                _ = v13
            } else {
                v14 := GenS1{F0: 42}
                _ = v14
            }
        }
        _ = v10
    } else {
        v4 = (v4 / ((0) | 1))
    }
    v15 := v1[-405]
    var v16 uint64 = 4294967295
    for fi17 := 0; fi17 < 20; fi17++ {
        var v18 int32 = -1
        var v19 uint = (v8 / ((0) | 1))
        var v20 uint32 = 2
        _ = fi17
        _ = v18
        _ = v19
        _ = v20
    }
    _ = len(v5)
    switch ("hello" + "foo bar") {
    case "hello":
        v4 -= 2
    case "a":
        v21 := GenS3{F0: 2147483647, F1: -1, F2: (v6 + 50)}
        _ = v21
    case "\n\t":
        switch (("\n\t" + "foo bar") + ("hello" + "a")) {
        case "":
            for fi22 := 0; fi22 < 19; fi22++ {
                fmt.Println(v8, v4)
                _ = fi22
            }
        case "foo bar":
            fmt.Println(v8)
        default:
            for fi23 := 0; fi23 < 9; fi23++ {
                var v24 int16 = v6
                _ = fi23
                _ = v24
            }
        }
    default:
        var v25 uint64 = (v16 % ((v16) | 1))
        _ = v25
    }
    v7 = (!v7)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
