package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 int; F1 uint }
type GenS3 struct { F0 string; F1 uint8 }
func genFn0(p0 uint32, p1 bool) uint32 {
    var v5 int16 = -1
    v6 := GenS1{F0: 42, F1: (127 > 128)}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 246
}

func main() {
    var v1 bool = true
    var v2 int8 = 1
    v3 := make(map[uint32]string)
    var v4 int64 = -924
    if v1 {
        var v7 int8 = 0
        for fi8 := 0; fi8 < 3; fi8++ {
            v9 := GenS2{F0: (fi8 * fi8), F1: 172}
            v7 = (v7 << ((0) & 63))
            _ = fi8
            _ = v9
        }
        v4 += (v4 * v4)
        _ = v7
    }
    v10 := []int{-23, 850, 154}
    switch -784 {
    case -545:
        var v11 uint64 = 4294967295
        _ = v11
    case 32767:
        for fi12 := 0; fi12 < 3; fi12++ {
            v10 = append(v10, (fi12 ^ 0))
            if (v4 <= v4) {
                v13 := GenS1{F0: 1, F1: (v1 && (!false))}
                if true {
                    for fi14 := 0; fi14 < 19; fi14++ {
                        _ = len(v10)
                        var v15 uint64 = 0
                        _ = fi14
                        _ = v15
                    }
                    v2 = v2
                    v1 = ((fi12 * -9223372036854775808) < (fi12 + fi12))
                }
                _ = len(v10)
                _ = v13
            }
            var v16 float64 = 0.1
            _ = fi12
            _ = v16
        }
    }
    delete(v3, 820)
    switch 93 {
    case 32767:
        var v17 int16 = 1
        _ = v17
    case 42:
        fmt.Println(v4, v4)
    }
    _ = len(v3)
    v2 += v2
    for fi18 := 0; fi18 < 6; fi18++ {
        for fi19 := 0; fi19 < 3; fi19++ {
            if ((3.14159 + -2.25) >= 1e10) {
                var v20 bool = false
                fmt.Println(v1)
                _ = v20
            }
            _ = fi19
        }
        v21 := v3[0]
        var v22 int8 = v2
        _ = fi18
        _ = v21
        _ = v22
    }
    for fi23 := 0; fi23 < 1; fi23++ {
        delete(v3, 2)
        var v24 uint64 = 0
        _ = fi23
        _ = v24
    }
    if true {
        v25 := GenS3{F0: ("\n\t" + ("a" + ("hello" + "foo bar"))), F1: 128}
        for fi26 := 0; fi26 < 3; fi26++ {
            if false {
                delete(v3, 1)
            }
            _ = fi26
        }
        for fi27 := 0; fi27 < 13; fi27++ {
            switch ("" + "hello") {
            case "hello":
                v4 = int64(v2)
            default:
                fmt.Println(fi27)
            }
            _ = fi27
        }
        _ = v25
    }
    var v28 int8 = -45
    var v29 int16 = 42
    fmt.Println(v4, v4)
    for fi30 := 0; fi30 < 1; fi30++ {
        fmt.Println(v1, fi30)
        _ = fi30
    }
    v31 := GenS2{F0: 42, F1: 65535}
    for fi32 := 0; fi32 < 6; fi32++ {
        switch "" {
        case "":
            fmt.Println(v29, v2)
        case "\n\t":
            var v33 int64 = 42
            _ = v33
        }
        for fi34 := 0; fi34 < 8; fi34++ {
            _ = len(v10)
            _ = fi34
        }
        _ = fi32
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v31
    _ = v4
}
