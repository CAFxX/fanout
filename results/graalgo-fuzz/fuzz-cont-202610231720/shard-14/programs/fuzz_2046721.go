package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint; F1 int; F2 int }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint8 = 163
    var v3 uint64 = 18446744073709551615
    v4 := []int{-692, 127}
    var v5 int16 = 42
    fmt.Println(v5, v5)
    v4 = append(v4, 42)
    if false {
        var v6 float32 = 1e10
        var v7 string = "foo bar"
        fmt.Println(v3, v3)
        _ = v6
        _ = v7
    } else {
        for fi8 := 0; fi8 < 4; fi8++ {
            v2 -= (v2 * v2)
            _ = fi8
        }
        v2 += (v2 + 1)
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    }
    var v9 string = ""
    if true {
        var v10 int32 = 1
        _ = v10
    }
    if false {
        v3 -= (v3 & v3)
        var v11 int = 32767
        v12 := GenS1{F0: 654}
        _ = v11
        _ = v12
    } else {
        var v13 float64 = (0.1 - (float64(1e10) - ((1.5 * -2.25) * 3.14159)))
        var v14 int8 = 0
        fmt.Println(v2, v5)
        _ = v13
        _ = v14
    }
    if true {
        for fi15 := 0; fi15 < 19; fi15++ {
            if false {
                for fi16 := 0; fi16 < 5; fi16++ {
                    if len(v4) > 0 {
                        v4[0] = fi15
                    }
                    _ = fi16
                }
                if false {
                    var v17 string = v9
                    for fi18 := 0; fi18 < 20; fi18++ {
                        var v19 int32 = 42
                        _ = len(v1)
                        var v20 uint64 = (v3 >> ((42) & 63))
                        _ = fi18
                        _ = v19
                        _ = v20
                    }
                    if false {
                        v21 := GenS1{F0: 0}
                        v1[-9223372036854775808] = 42
                        v9 = (v17 + v9)
                        _ = v21
                    } else {
                        var v22 int8 = 41
                        _ = v22
                    }
                    _ = v17
                } else {
                    var v23 float32 = 3.14159
                    v2 -= v2
                    _ = v23
                }
            } else {
                if len(v4) > 0 {
                    v4[0] = fi15
                }
                if len(v4) > 0 {
                    v4[0] = fi15
                }
            }
            fmt.Println(v3, v9)
            var v24 uint32 = 2147483648
            _ = fi15
            _ = v24
        }
    }
    switch 2147483647 {
    case 32767:
        v2 = v2
    }
    for fi25 := 0; fi25 < 8; fi25++ {
        var v26 uint = 405
        for fi27 := 0; fi27 < 8; fi27++ {
            fmt.Println(v2, v3, v9)
            v28 := GenS2{F0: v26, F1: fi27, F2: fi25}
            _ = fi27
            _ = v28
        }
        var v29 uint32 = 2147483647
        _ = fi25
        _ = v26
        _ = v29
    }
    v9 = (v9 + (v9 + ("\n\t" + "hello")))
    var v30 uint16 = 32767
    for fi31 := 0; fi31 < 12; fi31++ {
        var v32 int = (fi31 / ((fi31) | 1))
        v9 = "hello"
        v3 -= 1
        _ = fi31
        _ = v32
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v9
}
