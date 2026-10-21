package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 uint16; F1 int32; F2 int32 }
type GenS3 struct { F0 int; F1 int8; F2 int32 }
func genFn0(p0 uint) uint {
    p0 = 0
    var v5 int32 = -201
    _ = p0
    _ = v5
    return 2
}
func genFn1(p0 uint16, p1 int) uint16 {
    switch p1 {
    case 98:
        var v6 int64 = -463
        _ = v6
    case 2147483647:
        v7 := []int{-2147483648, 1, p1}
        _ = v7
    case -2147483648:
        var v8 float32 = float32((3.14159 * 0.0))
        _ = v8
    default:
        p1 = p1
    }
    v9 := []int{(p1 + (p1 << ((65535) & 63))), (p1 % ((127) | 1)), p1, 67}
    for fi10 := 0; fi10 < 14; fi10++ {
        if false {
            var v11 uint16 = p0
            _ = v11
        } else {
            p0 = (p0 * 32767)
            p0 += (p0 ^ p0)
        }
        if len(v9) > 0 {
            v9[0] = 599
        }
        _ = fi10
    }
    v12 := GenS1{F0: 42, F1: 9223372036854775807}
    p0 = p0
    _ = p0
    _ = p1
    _ = v12
    _ = v9
    return (p0 + p0)
}

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[string]bool)
    v3 := make(map[int]bool)
    v4 := []uint32{428, 2147483647, 2}
    var v13 uint = 0
    fmt.Println(v13)
    var v14 uint8 = 255
    var v15 uint64 = 42
    v4 = append(v4, 411)
    switch -721 {
    case 32767:
        for fi16 := 0; fi16 < 15; fi16++ {
            v17 := GenS2{F0: 2, F1: -1, F2: 728}
            for fi18 := 0; fi18 < 1; fi18++ {
                var v19 uint16 = 959
                var v20 float32 = (((0.0 + 1e10) - 0.0) + (1.5 / 0.1))
                _ = fi18
                _ = v19
                _ = v20
            }
            v14 = v14
            _ = fi16
            _ = v17
        }
    default:
        v21 := GenS2{F0: 32767, F1: 0, F2: -1}
        _ = v21
    }
    v13 -= (^(v13))
    v4 = append(v4, 65535)
    v14 = (^(v14))
    var v22 string = (((("" + ("a" + "a")) + "") + "hello") + ("foo bar" + "\n\t"))
    var v23 uint64 = (v15 >> (((v13 - 2)) & 63))
    v24 := GenS3{F0: 32768, F1: 100, F2: -2147483648}
    switch -2147483648 {
    case 32767:
        for fi25 := 0; fi25 < 15; fi25++ {
            if true {
                var v26 uint8 = v14
                _ = len(v4)
                _ = v26
            } else {
                switch (v22 + "a") {
                case "a":
                    v15 -= 1
                case "\n\t":
                    for fi27 := 0; fi27 < 19; fi27++ {
                        fmt.Println(v23, v23)
                        switch 42 {
                        case 32767:
                            v1[2] = v22
                        }
                        _ = fi27
                    }
                }
                var v28 string = v22
                _ = v28
            }
            fmt.Println(fi25, v22)
            _ = len(v4)
            _ = fi25
        }
    default:
        if false {
            fmt.Println(v22, v22, v15)
            var v29 bool = true
            _ = v29
        } else {
            var v30 int8 = 100
            _ = v30
        }
    }
    if true {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        v13 = v13
    }
    v31 := GenS1{F0: v13, F1: -9223372036854775808}
    var v32 int16 = 797
    var v33 float32 = ((-2.25 / 1.0) * (1.5 - (1.5 * (-2.25 - -2.25))))
    for fi34 := 0; fi34 < 5; fi34++ {
        var v35 uint16 = 65535
        _ = fi34
        _ = v35
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
}
