package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64; F1 bool; F2 int16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 string }

func main() {
    v1 := make(map[string]bool)
    var v2 int16 = 1000
    v3 := []int{32767, 0}
    var v4 uint32 = 1
    v5 := []int{0, -9223372036854775808, 4294967295, 2147483647}
    var v6 uint32 = 2147483647
    v7 := GenS1{F0: -444}
    v2 -= v2
    v8 := GenS1{F0: -2147483648}
    v9 := GenS2{F0: true}
    if (1000 >= v2) {
        v2 -= -1
        if len(v5) > 0 {
            v5[0] = -9223372036854775808
        }
        fmt.Println(v2)
    } else {
        for fi10 := 0; fi10 < 15; fi10++ {
            if len(v3) > 0 {
                v3[0] = 9223372036854775807
            }
            if true {
                fmt.Println(fi10, v4, fi10)
            } else {
                var v11 string = "\n\t"
                var v12 uint64 = 2
                if true {
                    for fi13 := 0; fi13 < 20; fi13++ {
                        v14 := GenS3{F0: 563, F1: (false && ((-2.25 / 1.5) <= 1.5)), F2: (v2 << ((18446744073709551615) & 63))}
                        v15 := GenS1{F0: -2147483648}
                        _ = fi13
                        _ = v14
                        _ = v15
                    }
                    v16 := GenS4{F0: (v11 + "a")}
                    _ = v16
                } else {
                    v12 = 42
                }
                _ = v11
                _ = v12
            }
            fmt.Println(v4, fi10, fi10)
            _ = fi10
        }
        if (0.0 != -2.25) {
            fmt.Println(v2, v4)
        } else {
            for fi17 := 0; fi17 < 15; fi17++ {
                var v18 int8 = -128
                fmt.Println(fi17, v6)
                v19 := GenS5{F0: (("hello" + "hello") + "\n\t")}
                _ = fi17
                _ = v18
                _ = v19
            }
        }
    }
    if len(v3) > 0 {
        v3[0] = 2147483647
    }
    v6 = (v4 >> ((65535) & 63))
    var v20 uint32 = 2147483647
    for fi21 := 0; fi21 < 15; fi21++ {
        for fi22 := 0; fi22 < 1; fi22++ {
            fmt.Println(v2)
            var v23 uint64 = 1
            _ = fi22
            _ = v23
        }
        _ = fi21
    }
    var v24 int = -1
    v24 = 358
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
