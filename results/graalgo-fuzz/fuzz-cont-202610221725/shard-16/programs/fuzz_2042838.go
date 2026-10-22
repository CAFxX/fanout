package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 int16 }
type GenS2 struct { F0 int16; F1 int64 }
type GenS3 struct { F0 int }
func genFn0(p0 string, p1 int32, p2 uint8) uint64 {
    var v7 string = "a"
    v8 := GenS1{F0: 42, F1: 18446744073709551615, F2: 0}
    v9 := []int{128}
    fmt.Println(p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return 465
}

func main() {
    var v1 bool = true
    var v2 int = 914
    var v3 uint8 = 1
    v4 := []int{(v2 & v2), (v2 ^ v2), (v2 / ((v2) | 1))}
    v5 := make(map[int]uint32)
    var v6 int = -1
    switch ("\n\t" + "\n\t") {
    case "\n\t":
        v2 -= v2
    case "hello":
        switch v6 {
        case -9223372036854775808:
            v10 := GenS1{F0: 30, F1: 1, F2: -1}
            _ = v10
        default:
            v4 = append(v4, v6)
        }
    case "foo bar":
        v11 := GenS2{F0: 404, F1: 9223372036854775807}
        _ = v11
    }
    for fi12 := 0; fi12 < 9; fi12++ {
        var v13 uint = 0
        if true {
            v3 = (v3 << ((v13) & 63))
        } else {
            var v14 float32 = 1.5
            _ = v14
        }
        v6 = v6
        _ = fi12
        _ = v13
    }
    for fi15 := 0; fi15 < 18; fi15++ {
        for fi16 := 0; fi16 < 7; fi16++ {
            var v17 bool = ((599 < -1) && v1)
            var v18 int8 = 127
            fmt.Println(v3, v18)
            var v19 uint = 2
            _ = fi16
            _ = v17
            _ = v18
            _ = v19
        }
        if (-1 > 127) {
            v20 := v5[(fi15 & (fi15 << ((255) & 63)))]
            for fi21 := 0; fi21 < 1; fi21++ {
                if true {
                    var v22 bool = (42 != v3)
                    v2 += (fi15 - v6)
                    if (255 > 2) {
                        var v23 string = "a"
                        v24 := GenS3{F0: (v2 >> ((56) & 63))}
                        _ = v23
                        _ = v24
                    }
                    _ = v22
                } else {
                    var v25 uint64 = 42
                    _ = v25
                }
                if v1 {
                    var v26 int8 = 0
                    fmt.Println(v26, fi21)
                    _ = v26
                }
                v6 = (-(fi21))
                _ = fi21
            }
            _ = v20
        }
        _ = fi15
    }
    fmt.Println(v3, v1)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if v1 {
        fmt.Println(v6, v3, v6)
        if true {
            v3 += v3
            fmt.Println(v3, v1)
            fmt.Println(v1, v1, v6)
        }
        fmt.Println(v1)
    }
    for fi27 := 0; fi27 < 13; fi27++ {
        for fi28 := 0; fi28 < 8; fi28++ {
            v2 = (v2 >> ((65535) & 63))
            _ = fi28
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi27
    }
    var v29 int = v2
    var v30 bool = false
    v3 = v3
    var v31 int16 = 1000
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
