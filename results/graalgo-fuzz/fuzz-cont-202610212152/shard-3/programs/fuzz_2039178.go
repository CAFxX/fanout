package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 uint16 }
type GenS2 struct { F0 int8; F1 uint64; F2 uint16 }
type GenS3 struct { F0 int64; F1 int32 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint8; F1 int16 }

func main() {
    var v1 int = 32767
    v2 := []string{("" + "")}
    v3 := []uint32{317}
    var v4 int = (v1 + v1)
    var v5 int8 = 118
    var v6 int16 = 48
    var v7 uint32 = 65535
    switch "\n\t" {
    case "a":
        v7 += v7
    case "hello":
        var v8 uint8 = 164
        _ = v8
    case "\n\t":
        if (false && (true && true)) {
            var v9 float64 = float64((1e10 * float32(0.0)))
            v4 = 862
            _ = v9
        } else {
            var v10 int16 = (v6 + v6)
            if false {
                if (-2.25 < -2.25) {
                    fmt.Println(v4, v4)
                    var v11 int16 = (v10 | (-(v10)))
                    if false {
                        v7 = (v7 << ((1) & 63))
                        v11 -= -886
                        var v12 float32 = 1e10
                        _ = v12
                    }
                    _ = v11
                }
                if (false || true) {
                    v13 := GenS1{F0: 42, F1: false, F2: 32767}
                    if len(v3) > 0 {
                        v3[0] = v7
                    }
                    _ = v13
                }
            }
            _ = v10
        }
    }
    if (false && true) {
        v14 := GenS1{F0: 32767, F1: (true && false), F2: 777}
        v6 = v6
        _ = v14
    }
    var v15 uint = 0
    var v16 uint64 = 18446744073709551615
    v17 := GenS1{F0: v6, F1: true, F2: 2}
    fmt.Println(v6, v6)
    var v18 bool = true
    if (1.5 != float64(((0.0 - -2.25) * 1.5))) {
        if (1 < -1) {
            if v18 {
                v19 := GenS2{F0: int8(v15), F1: v16, F2: 576}
                switch "foo bar" {
                case "hello":
                    v20 := GenS3{F0: 1, F1: -30}
                    _ = v20
                case "":
                    v7 += (v7 / (((v7 << (((v15 | v15)) & 63))) | 1))
                }
                v7 -= v7
                _ = v19
            } else {
                fmt.Println(v5, v1)
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if v18 {
            v21 := GenS4{F0: -32768}
            for fi22 := 0; fi22 < 14; fi22++ {
                v23 := GenS5{F0: v5}
                _ = fi22
                _ = v23
            }
            var v24 string = "hello"
            _ = v21
            _ = v24
        }
    } else {
        if v18 {
            v25 := GenS6{F0: 42}
            if true {
                v7 = v7
            }
            _ = v25
        } else {
            if v18 {
                v26 := GenS7{F0: 5, F1: (v6 * v6)}
                fmt.Println(v7, v4, v7)
                var v27 int64 = 0
                _ = v26
                _ = v27
            }
            var v28 int8 = -1
            _ = v28
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
