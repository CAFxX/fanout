package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64 }
type GenS2 struct { F0 int32 }
func genFn0(p0 int8, p1 uint8, p2 uint16) int {
    for fi6 := 0; fi6 < 5; fi6++ {
        p1 += p1
        _ = fi6
    }
    p2 += (p2 & p2)
    _ = p0
    _ = p1
    _ = p2
    return -109
}

func main() {
    var v1 uint64 = 552
    var v2 uint64 = (v1 % (((v1 * v1)) | 1))
    var v3 float32 = (-2.25 + 0.1)
    v4 := []int{-2147483648, 1, -1, -2147483648}
    var v5 uint32 = 1
    switch -408 {
    case 572:
        v3 = v3
    case 2147483647:
        v3 += (((v3 / (0.1 / v3)) - float32(((1.5 + -2.25) - (3.14159 * (0.0 / 1.0))))) / v3)
    }
    _ = len(v4)
    v1 += v1
    fmt.Println(v3, v2, v2)
    var v7 string = "hello"
    var v8 int32 = int32(v5)
    if false {
        v3 -= v3
    } else {
        for fi9 := 0; fi9 < 10; fi9++ {
            if true {
                if (-1 <= -910) {
                    v10 := GenS1{F0: 127, F1: v1}
                    var v11 int16 = 815
                    _ = v10
                    _ = v11
                }
                v7 = (v7 + v7)
                var v12 uint32 = v5
                _ = v12
            } else {
                v1 += v2
                v13 := GenS2{F0: (v8 & -2147483648)}
                var v14 float32 = ((v3 * v3) + v3)
                _ = v13
                _ = v14
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            var v15 uint8 = 127
            _ = fi9
            _ = v15
        }
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if true {
        if true {
            for fi16 := 0; fi16 < 16; fi16++ {
                for fi17 := 0; fi17 < 7; fi17++ {
                    v3 = (1e10 / v3)
                    var v18 uint = 1
                    _ = fi17
                    _ = v18
                }
                _ = fi16
            }
            var v19 uint32 = 65535
            v5 = (v19 - (v19 * v19))
            _ = v19
        } else {
            var v20 uint16 = 255
            _ = v20
        }
        v5 = v5
        v1 = 4294967295
    }
    var v21 uint64 = 0
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v22 float32 = v3
    v1 += v2
    if true {
        v23 := GenS2{F0: v8}
        if false {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            fmt.Println(v3, v7, v3)
            var v24 uint64 = v2
            _ = v24
        }
        _ = v23
    }
    if len(v4) > 0 {
        v4[0] = genFn0(58, 156, 864)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
