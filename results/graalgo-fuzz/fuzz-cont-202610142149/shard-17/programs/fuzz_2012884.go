package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 int8 }
type GenS2 struct { F0 uint; F1 int64 }
func genFn0(p0 string, p1 uint8, p2 uint8) int16 {
    v4 := GenS1{F0: 2, F1: 170, F2: -128}
    p2 -= p2
    p2 = 108
    for fi5 := 0; fi5 < 20; fi5++ {
        v6 := []int{fi5}
        if ((false || false) && false) {
            fmt.Println(p1)
            var v7 bool = (255 != 741)
            _ = v7
        }
        v8 := GenS2{F0: 255, F1: 407}
        _ = fi5
        _ = v6
        _ = v8
    }
    if true {
        fmt.Println(p0, p1, p0)
        var v9 bool = true
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return 1
}

func main() {
    var v1 bool = true
    v2 := make(map[uint32]uint32)
    var v3 int = 113
    fmt.Println(v3, v3)
    v10 := GenS2{F0: 955, F1: 1}
    switch "foo bar" {
    case "":
        var v11 string = ""
        _ = v11
    case "a":
        if ((genFn0(("" + ""), 27, 160) != 32767) || (((0 != int16(v3)) || (v1 || v1)) || (v1 || (-1 == 0)))) {
            v3 = (v3 & 9223372036854775807)
        }
    }
    var v12 uint64 = 18446744073709551615
    v12 += v12
    fmt.Println(v1, v1, v12)
    var v13 uint = 255
    v14 := v2[4294967295]
    if ((^(v13)) == (^(v13))) {
        for fi15 := 0; fi15 < 13; fi15++ {
            v12 -= (v12 & v12)
            for fi16 := 0; fi16 < 1; fi16++ {
                if (fi15 <= 806) {
                    fmt.Println(v12, v3)
                } else {
                    if v1 {
                        v13 = v13
                    } else {
                        var v17 string = ("hello" + "hello")
                        var v18 float32 = 1e10
                        v1 = (-1 > -1)
                        _ = v17
                        _ = v18
                    }
                    fmt.Println(v13)
                }
                _ = fi16
            }
            _ = fi15
        }
    }
    v12 += v12
    var v19 string = (("\n\t" + "") + "foo bar")
    v20 := GenS1{F0: 102, F1: 0, F2: 100}
    if false {
        v14 += uint32(v13)
        v12 += (v12 | (v12 | (v12 * uint64(v13))))
        for fi21 := 0; fi21 < 11; fi21++ {
            var v22 int = v3
            _ = fi21
            _ = v22
        }
    }
    v1 = false
    v23 := GenS1{F0: 1, F1: -931, F2: -1}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v23
    _ = v3
}
