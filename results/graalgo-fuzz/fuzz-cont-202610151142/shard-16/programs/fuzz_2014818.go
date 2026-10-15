package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint8; F1 uint8 }
func genFn0(p0 uint16, p1 int8) uint32 {
    for fi6 := 0; fi6 < 2; fi6++ {
        var v7 int8 = (p1 << ((335) & 63))
        fmt.Println(p1, fi6)
        _ = fi6
        _ = v7
    }
    v8 := []int{}
    var v9 bool = (-128 >= 100)
    for fi10 := 0; fi10 < 18; fi10++ {
        var v11 bool = ((!v9) && v9)
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 0
}
func genFn1(p0 uint16) uint8 {
    var v12 uint16 = 807
    v12 = (p0 >> ((1) & 63))
    for fi13 := 0; fi13 < 4; fi13++ {
        switch 128 {
        case 9223372036854775807:
            v14 := GenS1{F0: false}
            _ = v14
        default:
            var v15 float32 = (3.14159 / (1.5 + float32(0.1)))
            _ = v15
        }
        v16 := GenS2{F0: 9223372036854775807}
        for fi17 := 0; fi17 < 18; fi17++ {
            p0 = (p0 % ((v12) | 1))
            if ((!false) || false) {
                if (true && ("" >= "hello")) {
                    var v18 float64 = (3.14159 - -2.25)
                    var v19 int16 = 42
                    v20 := []int{fi17}
                    switch "foo bar" {
                    case "hello":
                        if len(v20) > 0 {
                            v20[0] = 9223372036854775807
                        }
                    default:
                        if (0 >= 18446744073709551615) {
                            var v21 uint = 255
                            var v22 int16 = v19
                            var v23 float32 = float32(0.0)
                            _ = v21
                            _ = v22
                            _ = v23
                        }
                    }
                    _ = v18
                    _ = v19
                    _ = v20
                } else {
                    switch (fi17 & (fi17 - fi17)) {
                    case 32767:
                        v24 := GenS3{F0: 2, F1: 47}
                        _ = v24
                    case -362:
                        fmt.Println(fi13, fi13, fi17)
                    }
                    fmt.Println(p0, fi13)
                    v12 += 65535
                }
                v25 := []int{-1, (-(fi17)), fi17, -9223372036854775808}
                _ = v25
            } else {
                p0 = (p0 << ((339) & 63))
                v12 = (v12 / ((42) | 1))
                v12 = (p0 & 65535)
            }
            _ = fi17
        }
        _ = fi13
        _ = v16
    }
    p0 = v12
    p0 = 1
    _ = p0
    _ = v12
    return 0
}

func main() {
    var v1 float64 = 1e10
    var v2 int = -179
    var v3 uint64 = 2
    v4 := []int{}
    v5 := []int{}
    if true {
        for fi26 := 0; fi26 < 4; fi26++ {
            fmt.Println(v1, v2, v2)
            _ = fi26
        }
    } else {
        var v27 int16 = 1
        var v28 int32 = 42
        _ = v27
        _ = v28
    }
    var v29 string = "\n\t"
    var v30 int16 = 0
    v30 = v30
    fmt.Println(v3)
    var v31 int = 2147483647
    if false {
        for fi32 := 0; fi32 < 13; fi32++ {
            var v33 int8 = 100
            _ = fi32
            _ = v33
        }
    } else {
        v29 = "foo bar"
    }
    v1 = v1
    v29 = v29
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
