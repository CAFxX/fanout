package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int64; F1 int }
func genFn0(p0 int16) int64 {
    var v7 string = ""
    var v8 int32 = 2147483647
    v7 = (v7 + v7)
    _ = p0
    _ = v7
    _ = v8
    return 9223372036854775807
}

func main() {
    var v1 uint32 = 654
    v2 := make(map[int]string)
    var v3 int16 = 1000
    v4 := []int{}
    var v5 int16 = (v3 | (-(v3)))
    v6 := make(map[uint32]int)
    fmt.Println(v1)
    for fi9 := 0; fi9 < 14; fi9++ {
        v10 := GenS1{F0: 95}
        fmt.Println(v5, v5)
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 5; fi11++ {
        v3 += (v3 << ((1) & 63))
        v5 -= (v3 >> ((1) & 63))
        for fi12 := 0; fi12 < 17; fi12++ {
            var v13 string = "foo bar"
            _ = fi12
            _ = v13
        }
        _ = fi11
    }
    var v14 uint16 = 2
    for fi15 := 0; fi15 < 19; fi15++ {
        fmt.Println(v3)
        _ = fi15
    }
    var v16 uint16 = 206
    v1 = v1
    var v17 uint = 2
    for fi18 := 0; fi18 < 15; fi18++ {
        var v19 uint64 = 987
        if false {
            var v20 int16 = (v5 + v3)
            _ = v20
        } else {
            for fi21 := 0; fi21 < 9; fi21++ {
                if (!false) {
                    _ = len(v6)
                    _ = len(v4)
                }
                _ = fi21
            }
            if true {
                v17 = v17
                for fi22 := 0; fi22 < 4; fi22++ {
                    v4 = append(v4, (fi18 / ((fi22) | 1)))
                    v17 -= uint(v1)
                    fmt.Println(v5, fi22)
                    _ = fi22
                }
                switch ("hello" + "foo bar") {
                case "hello":
                    _ = len(v6)
                case "a":
                    v23 := GenS2{F0: 9223372036854775807, F1: fi18}
                    _ = v23
                }
            }
        }
        v24 := v2[(fi18 & fi18)]
        _ = fi18
        _ = v19
        _ = v24
    }
    switch 9223372036854775807 {
    case -779:
        for fi25 := 0; fi25 < 12; fi25++ {
            v16 -= (v16 % ((65535) | 1))
            var v26 uint32 = v1
            _ = fi25
            _ = v26
        }
    }
    var v27 uint64 = 806
    if (101 >= (v17 % (((v17 & (v17 >> ((361) & 63)))) | 1))) {
        v16 += v16
        for fi28 := 0; fi28 < 6; fi28++ {
            var v29 string = ("a" + "hello")
            if ((v3 ^ (v3 ^ (v5 - v5))) <= 32767) {
                v16 = (v16 * (v16 & (v16 ^ v14)))
            }
            v1 = (v1 * 4294967295)
            _ = fi28
            _ = v29
        }
    } else {
        v30 := GenS1{F0: uint8(v17)}
        var v31 uint = (^(v17))
        v32 := GenS2{F0: 0, F1: 916}
        _ = v30
        _ = v31
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v16
    _ = v17
    _ = v2
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
