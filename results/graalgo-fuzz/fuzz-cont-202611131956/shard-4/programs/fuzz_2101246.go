package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64; F1 int32 }
func genFn0(p0 int32, p1 int32, p2 uint16) int16 {
    for fi4 := 0; fi4 < 10; fi4++ {
        if false {
            p1 = 0
            fmt.Println(fi4, fi4, p1)
            v5 := GenS1{F0: 1}
            _ = v5
        }
        if (0.1 <= 0.0) {
            for fi6 := 0; fi6 < 6; fi6++ {
                v7 := GenS2{F0: uint64(p2), F1: p0}
                switch fi4 {
                case 9223372036854775807:
                    switch ("a" + "") {
                    case "":
                        if (!false) {
                            v8 := []int{(fi6 / ((-9223372036854775808) | 1)), fi4, fi4, fi6}
                            p0 += (p1 ^ (p1 >> ((4294967295) & 63)))
                            _ = v8
                        } else {
                            p1 = p0
                            var v9 uint64 = 18446744073709551615
                            _ = v9
                        }
                    default:
                        switch fi4 {
                        case 2147483647:
                            p2 = p2
                        case 9223372036854775807:
                            v10 := GenS2{F0: 0, F1: p1}
                            _ = v10
                        case 464:
                            p2 = (p2 << ((979) & 63))
                        default:
                            v11 := []int{(fi6 | fi4), fi6, (fi4 >> ((230) & 63))}
                            fmt.Println(p1)
                            _ = v11
                        }
                    }
                case 32768:
                    switch fi6 {
                    case -1:
                        v12 := GenS2{F0: 322, F1: p0}
                        _ = v12
                    default:
                        p2 = p2
                    }
                default:
                    p2 += (p2 ^ 0)
                }
                _ = fi6
                _ = v7
            }
            var v13 uint16 = p2
            switch (-(fi4)) {
            case 32768:
                var v14 bool = false
                _ = v14
            case 2147483647:
                for fi15 := 0; fi15 < 13; fi15++ {
                    fmt.Println(v13, p1, fi15)
                    _ = fi15
                }
            case -1:
                var v16 uint = 65535
                _ = v16
            default:
                var v17 int8 = 0
                _ = v17
            }
            _ = v13
        } else {
            switch 468 {
            case 0:
                p1 += p1
            case 4294967295:
                var v18 int16 = 0
                _ = v18
            case 128:
                var v19 int8 = -1
                _ = v19
            }
        }
        fmt.Println(fi4)
        _ = fi4
    }
    switch ("hello" + "a") {
    case "foo bar":
        v20 := GenS1{F0: 1000}
        _ = v20
    case "a":
        fmt.Println(p2)
    case "hello":
        var v21 float64 = -2.25
        _ = v21
    }
    _ = p0
    _ = p1
    _ = p2
    return -32768
}

func main() {
    var v1 uint16 = 2
    v2 := make(map[string]string)
    var v3 uint16 = v1
    v3 = (^(v3))
    var v22 uint = 2
    v3 += (^(v1))
    var v23 float32 = -2.25
    v24 := GenS2{F0: 18446744073709551615, F1: -1}
    var v25 uint64 = 1
    for fi26 := 0; fi26 < 15; fi26++ {
        var v27 uint64 = 1
        _ = fi26
        _ = v27
    }
    var v28 string = "hello"
    v29 := GenS2{F0: uint64(v3), F1: -482}
    var v30 float32 = float32(3.14159)
    v23 = 1.5
    fmt.Println(v30, v22)
    v28 = v28
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v28
    _ = v29
    _ = v3
    _ = v30
}
