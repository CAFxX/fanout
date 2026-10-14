package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool }
type GenS2 struct { F0 int; F1 int; F2 int32 }
type GenS3 struct { F0 int8 }

func main() {
    var v1 string = "\n\t"
    v2 := []int{}
    v3 := make(map[uint32]string)
    v4 := make(map[uint32]string)
    var v5 int8 = -9
    var v6 float32 = 1.5
    fmt.Println(v6, v6)
    var v7 uint64 = 4294967295
    var v8 bool = false
    var v9 uint16 = 2
    if (42 != v7) {
        var v10 int64 = 663
        delete(v3, 42)
        _ = v10
    } else {
        switch "" {
        case "foo bar":
            v8 = false
        default:
            if ((!true) && (v8 && ((v6 - (v6 * ((-2.25 - v6) - float32((0.0 / ((0.0 * float64(1.5)) - (3.14159 * 1e10))))))) > v6))) {
                v1 = v1
                v11 := GenS1{F0: uint32(v9), F1: (v8 && (!true))}
                _ = v11
            }
        }
        fmt.Println(v1, v5, v5)
    }
    for fi12 := 0; fi12 < 3; fi12++ {
        if ("hello" <= "hello") {
            fmt.Println(v9, v9, fi12)
            var v13 float32 = (v6 / v6)
            if ((!v8) || v8) {
                var v14 int = fi12
                for fi15 := 0; fi15 < 10; fi15++ {
                    fmt.Println(v7, v5)
                    var v16 string = v1
                    _ = fi15
                    _ = v16
                }
                _ = v14
            }
            _ = v13
        } else {
            switch (v1 + v1) {
            case "hello":
                switch (fi12 * -381) {
                case 2147483647:
                    v17 := GenS1{F0: 0, F1: true}
                    _ = v17
                case 1:
                    switch "hello" {
                    case "a":
                        v9 = (v9 << ((18446744073709551615) & 63))
                    case "\n\t":
                        var v18 float64 = 3.14159
                        _ = v18
                    }
                case 127:
                    v19 := GenS1{F0: 2, F1: (89 > 65535)}
                    _ = v19
                }
            }
            fmt.Println(v6, fi12)
        }
        v8 = v8
        for fi20 := 0; fi20 < 12; fi20++ {
            v7 += (v7 << ((20) & 63))
            v21 := GenS2{F0: -421, F1: (fi20 << ((651) & 63)), F2: -2147483648}
            v22 := GenS3{F0: v5}
            _ = fi20
            _ = v21
            _ = v22
        }
        _ = fi12
    }
    if (0 <= 42) {
        if len(v2) > 0 {
            v2[0] = 128
        }
        var v23 string = (v1 + v1)
        v6 = float32(1.5)
        _ = v23
    }
    _ = len(v2)
    var v24 bool = v8
    var v25 uint = 42
    v3[42] = v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
