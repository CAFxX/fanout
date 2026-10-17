package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool; F2 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool; F1 int; F2 int64 }
type GenS4 struct { F0 uint; F1 uint64; F2 int8 }

func main() {
    v1 := make(map[int]bool)
    var v2 bool = true
    v3 := make(map[int]string)
    var v4 int64 = 629
    fmt.Println(v2, v2)
    fmt.Println(v4, v2)
    v4 += (v4 + v4)
    for fi5 := 0; fi5 < 11; fi5++ {
        v4 = v4
        for fi6 := 0; fi6 < 16; fi6++ {
            if ((!v2) || (42 >= 2147483648)) {
                v7 := GenS1{F0: 111, F1: true, F2: (1e10 == -2.25)}
                var v8 float64 = ((float64(0.0) - 3.14159) + float64(float32((float64(0.1) - 0.0))))
                if ((((3.14159 / 0.1) * float32(0.0)) - (1e10 / 0.1)) >= 0.0) {
                    v9 := GenS1{F0: 0, F1: (-187 < 0), F2: ("a" != ("foo bar" + "foo bar"))}
                    fmt.Println(fi6, v8, fi5)
                    delete(v1, fi5)
                    _ = v9
                } else {
                    v10 := GenS1{F0: -1, F1: (2 > 0), F2: v2}
                    v2 = (4294967295 == 42)
                    v11 := GenS2{F0: 1}
                    _ = v10
                    _ = v11
                }
                _ = v7
                _ = v8
            }
            v3[fi5] = ""
            _ = fi6
        }
        _ = fi5
    }
    var v12 uint32 = 2147483647
    fmt.Println(v12, v2)
    if (!v2) {
        v13 := []int{-1, -915, -474}
        _ = v13
    } else {
        if v2 {
            for fi14 := 0; fi14 < 3; fi14++ {
                v15 := GenS1{F0: 29, F1: (!v2), F2: (((v4 & v4) != v4) || v2)}
                _ = fi14
                _ = v15
            }
        } else {
            if (!v2) {
                v4 = (^(v4))
                fmt.Println(v4, v4)
            }
        }
        v16 := GenS3{F0: v2, F1: 840, F2: v4}
        for fi17 := 0; fi17 < 12; fi17++ {
            for fi18 := 0; fi18 < 9; fi18++ {
                var v19 uint64 = 192
                _ = fi18
                _ = v19
            }
            v3[127] = ("foo bar" + "")
            _ = fi17
        }
        _ = v16
    }
    v2 = v2
    v20 := GenS4{F0: 18446744073709551615, F1: 2, F2: 100}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
