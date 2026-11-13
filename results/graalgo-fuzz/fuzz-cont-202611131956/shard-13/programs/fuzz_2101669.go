package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint; F1 int; F2 int64 }

func main() {
    var v1 int = 0
    var v2 int32 = 1
    var v3 int8 = 42
    var v4 float64 = float64(-2.25)
    var v5 int32 = (v2 & v2)
    v6 := make(map[int]bool)
    var v7 uint32 = 4294967295
    var v8 uint16 = 432
    if true {
        v8 = (v8 % ((0) | 1))
    } else {
        switch ("" + "a") {
        case "":
            v9 := GenS1{F0: v7}
            _ = v9
        case "a":
            if false {
                v5 = (v2 * v2)
                v10 := GenS2{F0: 255, F1: (v1 + v1), F2: 518}
                var v11 float32 = 1e10
                _ = v10
                _ = v11
            } else {
                fmt.Println(v5, v7, v4)
                switch "" {
                case "":
                    delete(v6, 0)
                }
                for fi12 := 0; fi12 < 6; fi12++ {
                    fmt.Println(v7, v1, v1)
                    var v13 int8 = 1
                    _ = fi12
                    _ = v13
                }
            }
        }
        switch 393 {
        case 127:
            if false {
                var v14 uint = 1
                v2 = -1
                _ = v14
            }
        case 42:
            fmt.Println(v2, v3)
        }
    }
    fmt.Println(v3, v7)
    fmt.Println(v3)
    v15 := GenS2{F0: 904, F1: (^(v1)), F2: 0}
    v16 := GenS2{F0: 42, F1: 4294967295, F2: 42}
    v8 += v8
    for fi17 := 0; fi17 < 1; fi17++ {
        v6[(fi17 + (fi17 | v1))] = true
        _ = fi17
    }
    if false {
        var v18 bool = (!(v4 != v4))
        v19 := GenS1{F0: v7}
        _ = v18
        _ = v19
    } else {
        for fi20 := 0; fi20 < 11; fi20++ {
            v3 -= -67
            if false {
                v1 = 42
                var v21 uint8 = 2
                _ = v21
            }
            _ = fi20
        }
        var v22 bool = false
        v23 := GenS2{F0: 65535, F1: 0, F2: -9223372036854775808}
        _ = v22
        _ = v23
    }
    v24 := []int{(v1 * v1), v1}
    v8 += (^(v8))
    for fi25 := 0; fi25 < 1; fi25++ {
        var v26 float64 = v4
        _ = len(v6)
        v3 = v3
        _ = fi25
        _ = v26
    }
    var v27 int8 = v3
    var v28 int16 = 0
    if true {
        v8 -= (v8 & v8)
    }
    var v29 int = int(v3)
    var v30 int8 = (v3 + 100)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v24
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
