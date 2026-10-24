package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 uint16; F1 uint8; F2 bool }

func main() {
    var v1 uint8 = 0
    var v2 int32 = 42
    var v3 uint = 915
    var v4 int16 = -1
    switch 32767 {
    case -293:
        if true {
            for fi5 := 0; fi5 < 5; fi5++ {
                v4 = v4
                for fi6 := 0; fi6 < 4; fi6++ {
                    v3 += v3
                    _ = fi6
                }
                _ = fi5
            }
            v3 = (v3 >> (((v3 * (v3 >> ((4294967295) & 63)))) & 63))
            fmt.Println(v3, v1, v2)
        }
    case 1:
        fmt.Println(v1)
    }
    var v7 float64 = 3.14159
    var v8 uint32 = 2147483647
    var v9 int32 = (v2 - (v2 & v2))
    for fi10 := 0; fi10 < 12; fi10++ {
        fmt.Println(v2, v8)
        v11 := GenS1{F0: 1, F1: 2}
        _ = fi10
        _ = v11
    }
    v7 += v7
    if (0 == 42) {
        if (!(319 > 2)) {
            var v12 float64 = v7
            _ = v12
        } else {
            var v13 int32 = v9
            _ = v13
        }
        for fi14 := 0; fi14 < 10; fi14++ {
            var v15 uint64 = 42
            for fi16 := 0; fi16 < 16; fi16++ {
                if false {
                    v17 := []int{(-(fi16))}
                    if len(v17) > 0 {
                        v17[0] = fi16
                    }
                    v15 = v15
                    _ = v17
                } else {
                    v2 += (-(v9))
                    if false {
                        fmt.Println(v8, v15)
                    } else {
                        v18 := GenS1{F0: 18446744073709551615, F1: 65535}
                        if (!(!true)) {
                            v3 = (v3 | 4294967295)
                        }
                        _ = v18
                    }
                    switch (-(fi14)) {
                    case 2147483647:
                        var v19 uint32 = (^(v8))
                        _ = v19
                    case 32767:
                        switch "\n\t" {
                        case "a":
                            var v20 float64 = float64(-2.25)
                            _ = v20
                        }
                    case 42:
                        var v21 int16 = v4
                        _ = v21
                    default:
                        v22 := []int{(fi14 * (fi14 - 9223372036854775807))}
                        _ = v22
                    }
                }
                var v23 uint16 = uint16(v15)
                v9 = v9
                _ = fi16
                _ = v23
            }
            var v24 uint = 888
            _ = fi14
            _ = v15
            _ = v24
        }
    } else {
        var v25 bool = false
        _ = v25
    }
    for fi26 := 0; fi26 < 17; fi26++ {
        for fi27 := 0; fi27 < 4; fi27++ {
            v28 := GenS1{F0: 2, F1: 65535}
            var v29 uint16 = 0
            v1 = v1
            _ = fi27
            _ = v28
            _ = v29
        }
        v30 := GenS1{F0: 0, F1: 2}
        _ = fi26
        _ = v30
    }
    var v31 uint = v3
    for fi32 := 0; fi32 < 18; fi32++ {
        v1 += v1
        if true {
            v33 := GenS1{F0: 42, F1: 366}
            var v34 uint64 = 145
            _ = v33
            _ = v34
        }
        var v35 uint32 = v8
        _ = fi32
        _ = v35
    }
    var v36 uint32 = (v8 % ((v8) | 1))
    var v37 int16 = v4
    v38 := GenS2{F0: 595, F1: (^(v1)), F2: true}
    v39 := []int{-463, -537, 4294967295}
    if len(v39) > 0 {
        v39[0] = -877
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
