package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 uint8 }
type GenS2 struct { F0 uint; F1 uint }
func genFn0(p0 int8, p1 uint32) uint32 {
    p1 = (^(p1))
    v7 := []int{-2147483648, 0, 128}
    if false {
        if ((42 >= -1) || true) {
            v7 = append(v7, 1)
            if (false && ((p0 <= 0) || true)) {
                var v8 int8 = p0
                p1 += (^(p1))
                _ = v8
            }
            v9 := GenS1{F0: 135, F1: int16(p0), F2: 128}
            _ = v9
        } else {
            var v10 int16 = -246
            fmt.Println(p0)
            v10 -= v10
            _ = v10
        }
        fmt.Println(p1, p0)
        p1 = 353
    } else {
        p1 += p1
        fmt.Println(p1, p0)
    }
    _ = p0
    _ = p1
    _ = v7
    return (p1 >> ((255) & 63))
}

func main() {
    var v1 uint32 = 225
    v2 := make(map[int]int)
    var v3 uint64 = 18446744073709551615
    v4 := make(map[int]bool)
    var v5 string = "foo bar"
    v6 := make(map[uint32]int)
    var v11 int64 = 1
    var v12 int16 = 34
    v13 := GenS1{F0: v11, F1: (v12 ^ int16(v1)), F2: 2}
    v11 = -1
    v14 := GenS2{F0: 42, F1: 977}
    if false {
        if true {
            if false {
                var v15 int64 = (^(v11))
                for fi16 := 0; fi16 < 17; fi16++ {
                    v17 := GenS2{F0: 2, F1: 18446744073709551615}
                    v5 = (v5 + "a")
                    _ = fi16
                    _ = v17
                }
                _ = v15
            }
            var v18 uint32 = genFn0(-1, 472)
            _ = v18
        }
    } else {
        for fi19 := 0; fi19 < 6; fi19++ {
            switch (v5 + (((v5 + "") + v5) + v5)) {
            case "a":
                switch (v5 + v5) {
                case "":
                    var v20 string = v5
                    _ = v20
                }
            case "":
                v5 = (v5 + "foo bar")
            default:
                v21 := []int{}
                _ = v21
            }
            _ = fi19
        }
        fmt.Println(v12, v12, v11)
    }
    v22 := []int{32767, -309, 0, 32768}
    var v23 bool = ((int16(v11) > v12) && (0 < 2147483647))
    fmt.Println(v12, v5)
    if len(v22) >= 2 {
        v22 = v22[1:len(v22)]
    }
    var v24 uint16 = 42
    _ = len(v4)
    var v25 int8 = -128
    var v26 int = -2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
