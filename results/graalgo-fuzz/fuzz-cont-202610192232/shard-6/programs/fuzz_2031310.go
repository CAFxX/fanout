package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int16; F1 int8; F2 string }
type GenS3 struct { F0 uint32; F1 string }
type GenS4 struct { F0 uint32; F1 uint8 }

func main() {
    var v1 int = -972
    var v2 string = "\n\t"
    var v3 uint16 = 42
    v3 += v3
    var v4 uint32 = 765
    var v5 int8 = 42
    v6 := GenS1{F0: v5}
    v7 := []int{(v1 | v1)}
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    for fi8 := 0; fi8 < 5; fi8++ {
        v1 += fi8
        _ = fi8
    }
    v9 := GenS2{F0: -32768, F1: (v5 + v5), F2: v2}
    switch (v1 | -2147483648) {
    case -833:
        for fi10 := 0; fi10 < 16; fi10++ {
            v5 -= v5
            v11 := GenS3{F0: (v4 + v4), F1: v2}
            _ = fi10
            _ = v11
        }
    }
    var v12 uint32 = (v4 % ((324) | 1))
    if (v1 >= (v1 << ((1) & 63))) {
        v13 := GenS3{F0: 782, F1: (v2 + (((v2 + (v2 + "\n\t")) + (v2 + "\n\t")) + v2))}
        var v14 int16 = 1000
        _ = len(v7)
        _ = v13
        _ = v14
    }
    v2 = v2
    if len(v7) > 0 {
        v7[0] = (v1 - v1)
    }
    if false {
        if (false && (int8(v12) >= v5)) {
            var v15 int = v1
            if len(v7) > 0 {
                v7[0] = v15
            }
            var v16 uint16 = v3
            _ = v15
            _ = v16
        }
        for fi17 := 0; fi17 < 17; fi17++ {
            v18 := GenS3{F0: v4, F1: "a"}
            if true {
                v1 -= 1
                var v19 uint16 = 959
                for fi20 := 0; fi20 < 2; fi20++ {
                    var v21 int8 = (v5 & (v5 | v5))
                    var v22 float32 = 0.0
                    _ = fi20
                    _ = v21
                    _ = v22
                }
                _ = v19
            } else {
                if true {
                    v7 = append(v7, 42)
                    var v23 int32 = 98
                    _ = v23
                } else {
                    for fi24 := 0; fi24 < 20; fi24++ {
                        switch (v1 / (((fi24 & fi17)) | 1)) {
                        case -1:
                            var v25 float64 = 1e10
                            _ = v25
                        default:
                            v1 -= (fi17 / ((v1) | 1))
                        }
                        _ = fi24
                    }
                    v26 := GenS4{F0: v12, F1: uint8(v12)}
                    _ = v26
                }
                _ = len(v7)
            }
            var v27 uint16 = v3
            _ = fi17
            _ = v18
            _ = v27
        }
    }
    var v28 float32 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v12
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
