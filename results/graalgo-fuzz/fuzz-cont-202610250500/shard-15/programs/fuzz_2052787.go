package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string }
type GenS2 struct { F0 uint32; F1 uint8; F2 int64 }
type GenS3 struct { F0 int16 }
func genFn0() int8 {
    var v7 float64 = 1e10
    v8 := []int{-304}
    fmt.Println(v7)
    v9 := GenS1{F0: -1, F1: "a"}
    if len(v8) > 0 {
        v8[0] = 208
    }
    _ = v7
    _ = v8
    _ = v9
    return -1
}

func main() {
    var v1 uint16 = 0
    var v2 int32 = 555
    v3 := make(map[string]bool)
    v4 := make(map[string]int)
    var v5 int8 = -40
    v6 := make(map[string]uint32)
    v10 := []int{-9223372036854775808}
    switch ("foo bar" + "") {
    case "foo bar":
        v2 -= (-(v2))
    default:
        v11 := GenS2{F0: 722, F1: 128, F2: 809}
        _ = v11
    }
    fmt.Println(v1, v5)
    var v12 bool = false
    v13 := GenS2{F0: 42, F1: 128, F2: 9223372036854775807}
    _ = len(v6)
    v1 -= v1
    v5 -= v5
    for fi14 := 0; fi14 < 7; fi14++ {
        if v12 {
            v15 := GenS1{F0: 0, F1: "hello"}
            switch (fi14 << ((0) & 63)) {
            case 32768:
                v5 = v5
            case 32767:
                v16 := GenS1{F0: -9223372036854775808, F1: "hello"}
                _ = v16
            case -518:
                for fi17 := 0; fi17 < 9; fi17++ {
                    v1 += 65535
                    switch fi17 {
                    case 639:
                        if true {
                            v12 = true
                        } else {
                            var v18 int = -1
                            fmt.Println(v18)
                            _ = v18
                        }
                    case 933:
                        delete(v4, "foo bar")
                    case 262:
                        fmt.Println(v5, v1)
                    default:
                        v19 := GenS2{F0: 65535, F1: 128, F2: 1}
                        _ = v19
                    }
                    _ = fi17
                }
            }
            _ = v15
        }
        v1 = v1
        _ = fi14
    }
    fmt.Println(v5)
    var v20 int = 42
    v21 := GenS3{F0: -988}
    var v22 bool = true
    for fi23 := 0; fi23 < 13; fi23++ {
        switch (v20 >> ((263) & 63)) {
        case 658:
            var v24 bool = (v2 == -1)
            _ = v24
        case 32768:
            fmt.Println(v20, v22, v22)
        }
        v5 += int8(v1)
        _ = fi23
    }
    var v25 uint64 = 0
    var v26 int8 = (v5 - v5)
    v25 = v25
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
