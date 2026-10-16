package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64; F1 int8; F2 uint8 }
type GenS4 struct { F0 uint8; F1 int16 }
type GenS5 struct { F0 int; F1 int64 }
type GenS6 struct { F0 string; F1 int16 }
func genFn0(p0 uint64) uint64 {
    p0 -= (p0 >> ((710) & 63))
    if (true || ((42 < 42) && true)) {
        p0 = p0
    }
    if ((228 > 255) || (("a" > ("" + "hello")) || ((1 < 42) || false))) {
        var v4 bool = (2 > 0)
        if false {
            if v4 {
                var v5 uint16 = 267
                _ = v5
            }
        } else {
            var v6 uint = 566
            _ = v6
        }
        _ = v4
    }
    _ = p0
    return 2
}

func main() {
    var v1 float32 = 0.0
    var v2 int = 13
    var v3 uint = 18446744073709551615
    for fi7 := 0; fi7 < 4; fi7++ {
        v8 := GenS1{F0: 0}
        _ = fi7
        _ = v8
    }
    v9 := GenS2{F0: -535}
    if false {
        if false {
            if true {
                if false {
                    v10 := GenS1{F0: 563}
                    _ = v10
                } else {
                    v11 := GenS3{F0: 982, F1: 100, F2: 0}
                    _ = v11
                }
                for fi12 := 0; fi12 < 3; fi12++ {
                    if false {
                        v1 = ((((1e10 + -2.25) - 0.1) + -2.25) - 3.14159)
                    } else {
                        var v13 int = 1
                        var v14 int64 = 1
                        switch 32768 {
                        case 127:
                            fmt.Println(fi12)
                        default:
                            v15 := GenS4{F0: 30, F1: 32767}
                            _ = v15
                        }
                        _ = v13
                        _ = v14
                    }
                    v16 := []int{int(v3), v2}
                    _ = fi12
                    _ = v16
                }
                var v17 uint8 = 120
                _ = v17
            }
        }
        var v18 int64 = int64(v3)
        _ = v18
    }
    fmt.Println(v3, v2, v3)
    v19 := GenS5{F0: v2, F1: 0}
    var v20 uint64 = 255
    if false {
        v21 := []int{}
        v22 := GenS2{F0: 32767}
        v23 := GenS2{F0: -1}
        _ = v21
        _ = v22
        _ = v23
    } else {
        v24 := []int{v2, v2, v2, (-(v2))}
        if true {
            switch (v2 - (v2 + 32767)) {
            case -809:
                v2 = (^(v2))
            case -2147483648:
                var v25 int = v2
                _ = v25
            case 2147483647:
                v26 := GenS6{F0: "\n\t", F1: 1}
                _ = v26
            }
        }
        v27 := GenS1{F0: 42}
        _ = v24
        _ = v27
    }
    v1 = v1
    for fi28 := 0; fi28 < 4; fi28++ {
        fmt.Println(v1)
        fmt.Println(v3)
        v2 = (v2 ^ v2)
        _ = fi28
    }
    v3 -= v3
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v9
}
