package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 bool; F2 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int64; F1 uint16 }

func main() {
    var v1 bool = false
    v2 := []int{0}
    var v3 int32 = 0
    for fi4 := 0; fi4 < 3; fi4++ {
        v5 := GenS1{F0: 4294967295, F1: v1, F2: 65535}
        v3 = (v3 & v3)
        if (!false) {
            fmt.Println(v1)
        } else {
            v3 = 0
        }
        _ = fi4
        _ = v5
    }
    var v6 uint64 = 2
    v2 = append(v2, 359)
    var v7 int8 = 0
    fmt.Println(v6)
    var v8 int8 = 127
    switch 737 {
    case 32768:
        v6 = 4294967295
    }
    if v1 {
        v9 := GenS1{F0: 288, F1: v1, F2: 495}
        v7 = 90
        _ = v9
    } else {
        if (-1 > -32768) {
            for fi10 := 0; fi10 < 2; fi10++ {
                switch -9223372036854775808 {
                case 1:
                    v11 := GenS2{F0: 100}
                    _ = v11
                case -1:
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                }
                v12 := GenS1{F0: 42, F1: (float64(0.1) <= 0.0), F2: 436}
                _ = fi10
                _ = v12
            }
            v13 := GenS3{F0: -1, F1: 926}
            _ = v13
        }
    }
    var v14 uint = 404
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
}
