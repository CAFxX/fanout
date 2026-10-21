package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 int32 }
type GenS3 struct { F0 uint8 }
func genFn0() int32 {
    var v4 float64 = 1e10
    v4 += v4
    var v5 uint = 0
    if false {
        v6 := GenS1{F0: 208}
        _ = v6
    }
    _ = v4
    _ = v5
    return -200
}

func main() {
    var v1 string = ""
    var v2 uint16 = 42
    var v3 uint64 = 18446744073709551615
    var v7 bool = false
    fmt.Println(v7, v1)
    var v8 uint64 = v3
    var v9 uint16 = (^(v2))
    switch 32768 {
    case -907:
        var v10 int32 = -815
        _ = v10
    case 1:
        for fi11 := 0; fi11 < 20; fi11++ {
            v9 -= v2
            var v12 uint16 = v2
            if (1 < genFn0()) {
                v13 := GenS2{F0: v2, F1: -2147483648}
                v14 := GenS3{F0: uint8(v3)}
                _ = v13
                _ = v14
            } else {
                for fi15 := 0; fi15 < 12; fi15++ {
                    var v16 uint64 = (v3 >> ((2) & 63))
                    v2 -= (v2 + v2)
                    switch 32768 {
                    case -10:
                        var v17 int8 = 42
                        _ = v17
                    }
                    _ = fi15
                    _ = v16
                }
                var v18 int8 = 42
                _ = v18
            }
            _ = fi11
            _ = v12
        }
    case -281:
        switch "a" {
        case "foo bar":
            for fi19 := 0; fi19 < 20; fi19++ {
                for fi20 := 0; fi20 < 2; fi20++ {
                    v8 = v3
                    if true {
                        if v7 {
                            v9 = v2
                        }
                    }
                    v7 = (2147483647 > 42)
                    var v21 float64 = (1e10 - 1e10)
                    _ = fi20
                    _ = v21
                }
                switch ((("a" + v1) + v1) + v1) {
                case "a":
                    v22 := []int{(fi19 % ((1) | 1)), fi19, -1, fi19}
                    _ = v22
                }
                _ = fi19
            }
        }
    default:
        for fi23 := 0; fi23 < 6; fi23++ {
            var v24 string = ((v1 + v1) + v1)
            _ = fi23
            _ = v24
        }
    }
    v25 := GenS1{F0: 0}
    for fi26 := 0; fi26 < 18; fi26++ {
        var v27 uint8 = 255
        v28 := []int{fi26, fi26, fi26, (-(fi26))}
        _ = fi26
        _ = v27
        _ = v28
    }
    v29 := GenS2{F0: 512, F1: 2147483647}
    var v30 int32 = 2147483647
    v2 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v30
    _ = v7
    _ = v8
    _ = v9
}
