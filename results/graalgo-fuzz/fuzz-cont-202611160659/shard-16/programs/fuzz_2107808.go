package main

import (
    "fmt"
)

func genFn0() bool {
    switch "hello" {
    case "":
        if (!false) {
            v7 := []int{}
            _ = v7
        }
    default:
        var v8 uint64 = 1
        _ = v8
    }
    if (true || true) {
        var v9 uint = 768
        for fi10 := 0; fi10 < 7; fi10++ {
            for fi11 := 0; fi11 < 4; fi11++ {
                if false {
                    v9 = 2
                    fmt.Println(fi10, fi10, v9)
                }
                _ = fi11
            }
            for fi12 := 0; fi12 < 16; fi12++ {
                if (819 >= -142) {
                    fmt.Println(fi12, fi10, v9)
                    var v13 uint32 = 259
                    var v14 int16 = -1
                    _ = v13
                    _ = v14
                } else {
                    v9 = (v9 / ((65535) | 1))
                }
                fmt.Println(v9, fi10)
                _ = fi12
            }
            _ = fi10
        }
        _ = v9
    }
    var v15 uint32 = 2147483647
    var v16 uint32 = 2147483648
    _ = v15
    _ = v16
    return false
}

func main() {
    var v1 float64 = 1.5
    v2 := make(map[string]bool)
    var v3 uint = 255
    v4 := make(map[int]uint32)
    v5 := make(map[uint32]int)
    var v6 float64 = 0.0
    v17 := []int{9223372036854775807, -1}
    var v18 int16 = 1
    switch "hello" {
    case "\n\t":
        var v19 uint8 = 2
        _ = v19
    default:
        v3 += (v3 << (((v3 >> (((v3 - (v3 * (^(v3))))) & 63))) & 63))
    }
    for fi20 := 0; fi20 < 20; fi20++ {
        var v21 int8 = -128
        _ = fi20
        _ = v21
    }
    var v22 uint16 = 65535
    var v23 int = 1
    var v24 uint = (^(v3))
    var v25 int = (v23 + 32768)
    for fi26 := 0; fi26 < 17; fi26++ {
        for fi27 := 0; fi27 < 3; fi27++ {
            fmt.Println(fi27)
            _ = fi27
        }
        var v28 bool = false
        var v29 int64 = -979
        _ = fi26
        _ = v28
        _ = v29
    }
    if (true && true) {
        var v30 int16 = 751
        if true {
            switch "a" {
            case "\n\t":
                v22 = 0
            case "a":
                delete(v2, "hello")
            case "foo bar":
                v3 -= v3
            default:
                switch "" {
                case "a":
                    fmt.Println(v6)
                }
            }
        }
        v31 := v4[v25]
        _ = v30
        _ = v31
    } else {
        for fi32 := 0; fi32 < 6; fi32++ {
            var v33 int = 1
            var v34 int = (-(fi32))
            _ = fi32
            _ = v33
            _ = v34
        }
        var v35 uint64 = 196
        var v36 bool = false
        _ = v35
        _ = v36
    }
    delete(v5, 2)
    for fi37 := 0; fi37 < 3; fi37++ {
        _ = len(v17)
        _ = fi37
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
