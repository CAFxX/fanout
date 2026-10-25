package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 int32 }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 int64 = 42
    var v2 uint8 = 255
    var v3 uint = 2
    var v4 int16 = 842
    v3 -= (^(v3))
    switch -731 {
    case 2147483647:
        for fi5 := 0; fi5 < 3; fi5++ {
            v6 := []int{fi5, fi5}
            _ = fi5
            _ = v6
        }
    case 966:
        var v7 string = ("" + "foo bar")
        _ = v7
    default:
        for fi8 := 0; fi8 < 14; fi8++ {
            v9 := []int{fi8}
            if false {
                v9 = append(v9, (-(fi8)))
            } else {
                fmt.Println(v2)
            }
            _ = len(v9)
            _ = fi8
            _ = v9
        }
    }
    v10 := GenS1{F0: v1, F1: 424, F2: -560}
    var v11 int64 = v1
    for fi12 := 0; fi12 < 3; fi12++ {
        for fi13 := 0; fi13 < 13; fi13++ {
            v11 += (v11 + v1)
            if false {
                var v14 int8 = 42
                _ = v14
            }
            _ = fi13
        }
        _ = fi12
    }
    v1 += (^(v1))
    v2 -= 128
    switch -9223372036854775808 {
    case -1:
        var v15 int64 = (v11 & v11)
        _ = v15
    case 4294967295:
        var v16 float64 = 1.5
        _ = v16
    case 0:
        v17 := []int{-2147483648}
        _ = v17
    }
    for fi18 := 0; fi18 < 13; fi18++ {
        v3 -= (^(v3))
        v1 = 409
        _ = fi18
    }
    switch 2147483647 {
    case -431:
        var v19 uint32 = 984
        _ = v19
    default:
        for fi20 := 0; fi20 < 1; fi20++ {
            for fi21 := 0; fi21 < 5; fi21++ {
                switch (("\n\t" + "\n\t") + "foo bar") {
                case "a":
                    if false {
                        var v22 float64 = ((float64((-2.25 + (1e10 - 0.1))) / -2.25) / 0.1)
                        _ = v22
                    } else {
                        v23 := []int{(-(fi21)), fi20, fi21}
                        v4 = v4
                        var v24 string = "hello"
                        _ = len(v23)
                        _ = v23
                        _ = v24
                    }
                default:
                    var v25 uint32 = 1
                    _ = v25
                }
                _ = fi21
            }
            v2 = 42
            _ = fi20
        }
    }
    for fi26 := 0; fi26 < 1; fi26++ {
        var v27 int64 = v1
        v4 = (v4 >> ((v3) & 63))
        if (false || false) {
            v27 = v27
            v28 := GenS2{F0: 731}
            v11 += -1
            _ = v28
        }
        _ = fi26
        _ = v27
    }
    var v29 float32 = (1e10 * 1.5)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v29
    _ = v3
    _ = v4
}
