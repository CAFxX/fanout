package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8 }

func main() {
    var v1 int = 127
    var v2 uint = 716
    var v3 int = v1
    var v4 int8 = 100
    v5 := []uint32{2}
    var v6 string = "a"
    v5 = append(v5, 1)
    var v7 uint8 = 127
    switch v3 {
    case 42:
        v8 := GenS1{F0: (v7 ^ v7)}
        _ = v8
    }
    var v9 uint = 0
    fmt.Println(v7, v4, v6)
    var v10 int64 = int64(v9)
    for fi11 := 0; fi11 < 4; fi11++ {
        _ = len(v5)
        v9 += (v2 % ((v2) | 1))
        for fi12 := 0; fi12 < 17; fi12++ {
            v4 -= v4
            _ = fi12
        }
        _ = fi11
    }
    switch (v3 * 42) {
    case 128:
        for fi13 := 0; fi13 < 18; fi13++ {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            _ = fi13
        }
    case 1:
        v14 := GenS1{F0: uint8(v9)}
        _ = v14
    }
    for fi15 := 0; fi15 < 20; fi15++ {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        for fi16 := 0; fi16 < 4; fi16++ {
            v7 += v7
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            v3 += (^(v3))
            var v17 uint8 = (v7 / (((v7 % ((73) | 1))) | 1))
            v2 += 18446744073709551615
            v10 += (v10 % (((-(v10))) | 1))
            v18 := GenS2{F0: 100}
            _ = fi16
            _ = v17
            _ = v18
        }
        v19 := GenS2{F0: v4}
        _ = fi15
        _ = v19
    }
    v20 := GenS3{F0: (v4 << (((v2 % (((v2 % (((v2 + v2)) | 1))) | 1))) & 63))}
    var v21 uint = v9
    var v22 float32 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
