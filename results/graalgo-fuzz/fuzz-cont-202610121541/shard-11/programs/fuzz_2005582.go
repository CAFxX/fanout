package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 int; F1 int16 }
type GenS3 struct { F0 string; F1 int32 }

func main() {
    var v1 bool = true
    var v2 uint32 = 208
    v3 := []int{1}
    v4 := []uint32{v2, 457}
    var v5 int = 1
    v5 = 155
    var v6 int16 = 0
    v7 := GenS1{F0: 4294967295, F1: 0}
    v1 = (false && v1)
    switch "a" {
    case "a":
        if v1 {
            for fi8 := 0; fi8 < 17; fi8++ {
                v6 -= (^(v6))
                _ = fi8
            }
            var v9 uint16 = 660
            _ = v9
        }
    case "foo bar":
        if v1 {
            _ = len(v4)
        }
    case "hello":
        if len(v4) > 0 {
            v4[0] = v2
        }
    default:
        switch 2147483647 {
        case 32767:
            fmt.Println(v5)
        default:
            v10 := GenS2{F0: v5, F1: 0}
            _ = v10
        }
    }
    var v11 int8 = -127
    v1 = false
    if v1 {
        var v12 uint16 = 42
        _ = v12
    } else {
        v2 += 1
        switch ("foo bar" + ("foo bar" + "")) {
        case "foo bar":
            var v13 int8 = v11
            _ = v13
        case "hello":
            if v1 {
                v6 = v6
            } else {
                var v14 int32 = 0
                fmt.Println(v14, v2, v5)
                _ = v14
            }
        case "":
            if true {
                if v1 {
                    var v15 int32 = -1
                    v16 := GenS2{F0: v5, F1: (^(v6))}
                    for fi17 := 0; fi17 < 20; fi17++ {
                        if v1 {
                            var v18 int = v5
                            var v19 float64 = 1.5
                            _ = v18
                            _ = v19
                        }
                        v20 := GenS1{F0: 0, F1: 255}
                        _ = fi17
                        _ = v20
                    }
                    _ = v15
                    _ = v16
                }
                if len(v3) > 0 {
                    v3[0] = (v5 << ((0) & 63))
                }
            }
        default:
            v5 = (v5 << ((255) & 63))
        }
    }
    v21 := GenS3{F0: "\n\t", F1: -677}
    if false {
        v11 = (v11 & 127)
        var v22 uint = 255
        _ = v22
    } else {
        fmt.Println(v2)
    }
    for fi23 := 0; fi23 < 7; fi23++ {
        fmt.Println(v2, v2)
        v1 = ((v2 & v2) == v2)
        _ = fi23
    }
    if (v1 || (((((float32((0.0 + 3.14159)) / float32(1e10)) - 0.1) + 1.5) * float32((-2.25 * 1.5))) != 0.0)) {
        var v24 uint32 = v2
        fmt.Println(v2, v1)
        v25 := GenS1{F0: 2, F1: 2}
        _ = v24
        _ = v25
    } else {
        var v26 uint64 = 1
        _ = len(v4)
        _ = v26
    }
    for fi27 := 0; fi27 < 5; fi27++ {
        for fi28 := 0; fi28 < 17; fi28++ {
            v1 = v1
            fmt.Println(v11)
            fmt.Println(v5, v11, fi27)
            _ = fi28
        }
        var v29 uint32 = 2
        _ = fi27
        _ = v29
    }
    v4 = append(v4, 0)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
