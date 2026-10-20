package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 int8; F1 uint64; F2 string }

func main() {
    var v1 uint = 18446744073709551615
    v2 := make(map[uint32]bool)
    var v3 bool = true
    var v4 int16 = 1000
    var v5 int8 = 52
    var v6 int16 = v4
    if v3 {
        var v7 float64 = float64(1e10)
        _ = v7
    } else {
        v8 := []int{32768}
        _ = v8
    }
    if v3 {
        if (-2.25 <= 0.1) {
            var v9 bool = (4294967295 == -9223372036854775808)
            _ = v9
        }
        if (!(2 >= 42)) {
            v6 -= (v4 & (v4 ^ -35))
            switch 32768 {
            case 1:
                v10 := GenS1{F0: 18446744073709551615}
                _ = v10
            }
            if (-2147483648 > -807) {
                v4 = (^(v6))
            }
        }
    } else {
        if false {
            delete(v2, 793)
            v4 = v4
            v3 = (127 != 134)
        }
        var v11 uint64 = 487
        _ = v11
    }
    switch "" {
    case "\n\t":
        switch 1 {
        case 219:
            v5 += (v5 % ((v5) | 1))
        }
    }
    var v12 int16 = -995
    if v3 {
        switch -2147483648 {
        case -1:
            fmt.Println(v5, v12, v3)
        case 1:
            if v3 {
                v4 -= v6
                var v13 uint16 = 1
                _ = v13
            } else {
                if v3 {
                    if (v3 || v3) {
                        delete(v2, 2147483647)
                    }
                    var v14 float64 = float64(3.14159)
                    _ = v14
                }
                if ((v6 | v4) >= -404) {
                    v3 = (789 == -9223372036854775808)
                }
                var v15 int8 = -1
                _ = v15
            }
        case 127:
            v16 := GenS2{F0: -9223372036854775808, F1: v4}
            _ = v16
        default:
            v17 := v2[515]
            _ = v17
        }
        for fi18 := 0; fi18 < 16; fi18++ {
            var v19 int8 = (v5 / ((v5) | 1))
            v20 := v2[990]
            v21 := GenS1{F0: 805}
            _ = fi18
            _ = v19
            _ = v20
            _ = v21
        }
    }
    if true {
        var v22 bool = v3
        for fi23 := 0; fi23 < 1; fi23++ {
            v3 = v3
            _ = fi23
        }
        v24 := GenS1{F0: 226}
        _ = v22
        _ = v24
    }
    fmt.Println(v6, v6)
    v1 += (v1 / ((597) | 1))
    var v25 string = ""
    for fi26 := 0; fi26 < 8; fi26++ {
        var v27 uint32 = 2
        _ = fi26
        _ = v27
    }
    fmt.Println(v4, v3)
    for fi28 := 0; fi28 < 8; fi28++ {
        v29 := GenS3{F0: -119, F1: 2, F2: ((v25 + v25) + "foo bar")}
        v25 = v25
        _ = fi28
        _ = v29
    }
    var v30 string = "hello"
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v25
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
