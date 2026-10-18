package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint; F1 uint16; F2 int64 }
type GenS4 struct { F0 int16; F1 int64; F2 uint }
func genFn0(p0 uint32, p1 int64) int32 {
    fmt.Println(p1, p0)
    v6 := GenS1{F0: 65535}
    _ = p0
    _ = p1
    _ = v6
    return 834
}

func main() {
    var v1 string = "hello"
    var v2 int16 = 1000
    var v3 uint64 = 0
    var v4 int32 = 42
    var v5 uint8 = 128
    if (v3 == (v3 & (v3 ^ v3))) {
        v2 = 0
        v7 := GenS2{F0: 42}
        _ = v7
    }
    if true {
        var v8 string = v1
        v9 := []int{}
        if true {
            for fi10 := 0; fi10 < 11; fi10++ {
                v2 = (v2 + -1)
                var v11 float64 = ((1.5 + ((-2.25 - (-2.25 / 1.0)) - float64(3.14159))) * 0.0)
                _ = fi10
                _ = v11
            }
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
        }
        _ = v8
        _ = v9
    }
    v12 := []int{0, 333, 32768}
    for fi13 := 0; fi13 < 20; fi13++ {
        var v14 string = "hello"
        for fi15 := 0; fi15 < 8; fi15++ {
            if ((v4 & (v4 & (-(v4)))) > v4) {
                if false {
                    fmt.Println(v5, v3, v4)
                    if (1e10 <= (-2.25 * 3.14159)) {
                        fmt.Println(v1, fi13)
                    }
                }
                v4 -= (v4 >> ((789) & 63))
                var v16 float32 = (3.14159 / 0.1)
                var v17 uint8 = v5
                _ = v16
                _ = v17
            } else {
                var v18 int64 = -793
                if true {
                    fmt.Println(v4)
                    v18 = -196
                } else {
                    v19 := GenS1{F0: 1}
                    v20 := GenS3{F0: 2, F1: 255, F2: (v18 + (v18 >> ((65535) & 63)))}
                    var v21 uint = 1
                    _ = v19
                    _ = v20
                    _ = v21
                }
                _ = v18
            }
            v22 := GenS1{F0: 255}
            _ = fi15
            _ = v22
        }
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
        _ = fi13
        _ = v14
    }
    v12 = append(v12, -570)
    switch (("\n\t" + v1) + v1) {
    case "a":
        v23 := GenS4{F0: 42, F1: -1, F2: 255}
        _ = v23
    }
    v24 := GenS1{F0: 42}
    var v25 int32 = int32(v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
