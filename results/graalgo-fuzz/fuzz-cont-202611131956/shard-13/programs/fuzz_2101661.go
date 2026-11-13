package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 uint8; F1 int16 }
func genFn0() uint {
    switch "a" {
    case "foo bar":
        var v6 uint16 = 42
        _ = v6
    case "hello":
        var v7 uint = 939
        _ = v7
    default:
        v8 := GenS1{F0: "hello", F1: ("\n\t" + ("hello" + "foo bar"))}
        _ = v8
    }
    var v9 int64 = 829
    var v10 int = 42
    _ = v10
    _ = v9
    return 42
}
func genFn1() int8 {
    var v11 uint = 1
    var v12 uint = (v11 >> ((v11) & 63))
    _ = v11
    _ = v12
    return 1
}

func main() {
    var v1 uint = 18446744073709551615
    v2 := make(map[int]uint32)
    var v3 int16 = -571
    v4 := []int{}
    var v5 int64 = 1
    v3 = -1
    _ = len(v4)
    var v13 uint8 = uint8(v1)
    if len(v4) > 0 {
        v4[0] = 127
    }
    var v14 int64 = v5
    v5 = v5
    var v15 uint8 = v13
    v13 = v15
    v15 = v15
    for fi16 := 0; fi16 < 5; fi16++ {
        v17 := GenS2{F0: 42, F1: int16(v5)}
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi16
        _ = v17
    }
    for fi18 := 0; fi18 < 8; fi18++ {
        _ = len(v2)
        _ = len(v2)
        _ = fi18
    }
    if false {
        if false {
            var v19 int32 = 1
            _ = v19
        } else {
            v20 := GenS1{F0: "a", F1: "hello"}
            _ = v20
        }
    } else {
        if true {
            switch 0 {
            case -1:
                for fi21 := 0; fi21 < 20; fi21++ {
                    var v22 uint64 = 1
                    if (2 <= 65535) {
                        var v23 uint32 = 0
                        _ = v23
                    }
                    _ = fi21
                    _ = v22
                }
            case -77:
                if len(v4) > 0 {
                    v4[0] = -2147483648
                }
            }
        } else {
            for fi24 := 0; fi24 < 16; fi24++ {
                fmt.Println(v13, v13, v14)
                _ = fi24
            }
            v3 += (v3 % ((1) | 1))
        }
    }
    var v25 float32 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
}
