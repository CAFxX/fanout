package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint16; F1 int }
type GenS4 struct { F0 int8 }

func main() {
    var v1 int64 = 0
    v2 := make(map[int]int)
    v3 := []int{-568, 0, 566}
    var v4 uint16 = 716
    var v5 int8 = 127
    v6 := []string{}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v7 uint64 = 18446744073709551615
    var v8 int16 = 0
    _ = len(v2)
    var v9 int64 = -575
    if false {
        v9 = 0
    }
    var v10 int = -2147483648
    if (true && true) {
        switch (("" + ((((("\n\t" + "foo bar") + "\n\t") + ("\n\t" + "foo bar")) + "hello") + "\n\t")) + "") {
        case "":
            var v11 int16 = v8
            _ = v11
        case "\n\t":
            v12 := GenS1{F0: v10, F1: v7}
            _ = v12
        default:
            var v13 uint8 = 42
            _ = v13
        }
        switch (v10 + (v10 >> ((65535) & 63))) {
        case 2147483647:
            if (0.1 != (0.0 * 0.0)) {
                v14 := GenS2{F0: 42}
                if (v10 == -634) {
                    v3 = append(v3, v10)
                    var v15 uint16 = v4
                    v5 += v5
                    _ = v15
                } else {
                    fmt.Println(v10, v4)
                    v16 := GenS3{F0: v4, F1: v10}
                    v5 -= 0
                    _ = v16
                }
                var v17 int = v10
                _ = v14
                _ = v17
            } else {
                v18 := GenS2{F0: -2147483648}
                delete(v2, 533)
                if (v10 != v10) {
                    fmt.Println(v5, v1, v1)
                    v1 += (v9 + 9223372036854775807)
                }
                _ = v18
            }
        case 127:
            if ((1 < 2) || false) {
                fmt.Println(v1)
            } else {
                fmt.Println(v7, v4)
            }
        }
        v2[32767] = (^(v10))
    } else {
        v19 := v2[(v10 ^ 32767)]
        _ = v19
    }
    v20 := GenS3{F0: v4, F1: int(v8)}
    fmt.Println(v5)
    var v21 uint64 = (v7 % ((v7) | 1))
    fmt.Println(v7, v5)
    v9 -= 816
    fmt.Println(v9, v21, v21)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch "hello" {
    case "":
        for fi22 := 0; fi22 < 3; fi22++ {
            v23 := GenS4{F0: (^(v5))}
            v24 := GenS4{F0: 0}
            _ = fi22
            _ = v23
            _ = v24
        }
    case "foo bar":
        v25 := GenS2{F0: 2147483647}
        _ = v25
    }
    v26 := v2[(v10 ^ v10)]
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v21
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
