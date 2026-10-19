package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string; F1 uint8; F2 bool }
type GenS3 struct { F0 int; F1 string }
func genFn0(p0 int16) uint16 {
    switch "hello" {
    case "\n\t":
        var v4 uint64 = 240
        _ = v4
    case "hello":
        p0 = p0
    case "":
        p0 = -558
    default:
        v5 := GenS1{F0: 62}
        _ = v5
    }
    var v6 uint64 = 2
    _ = p0
    _ = v6
    return 520
}
func genFn1() bool {
    var v7 string = ""
    v8 := []int{-1, 127, -9223372036854775808}
    v7 = ((v7 + ("" + "\n\t")) + v7)
    var v9 uint = 788
    _ = v7
    _ = v8
    _ = v9
    return (1 >= 42)
}

func main() {
    var v1 string = "a"
    v2 := make(map[uint32]uint32)
    var v3 bool = (!(!(!(!false))))
    v10 := GenS1{F0: 0}
    v3 = v3
    var v11 uint = 255
    v11 = (v11 - (v11 * v11))
    for fi12 := 0; fi12 < 20; fi12++ {
        if (!false) {
            v11 = (v11 >> ((v11) & 63))
        } else {
            for fi13 := 0; fi13 < 1; fi13++ {
                var v14 int16 = -32768
                if v3 {
                    var v15 float64 = 1.5
                    v16 := v2[551]
                    v14 = -32768
                    var v17 int64 = -1
                    v18 := []int{}
                    _ = v15
                    _ = v16
                    _ = v17
                    _ = v18
                }
                v19 := GenS1{F0: 0}
                _ = fi13
                _ = v14
                _ = v19
            }
        }
        v11 = 2
        _ = fi12
    }
    if (v3 || v3) {
        fmt.Println(v11, v1)
        var v20 int16 = 1
        v21 := []int{}
        _ = v20
        _ = v21
    } else {
        var v22 int = 1
        _ = v22
    }
    v1 = v1
    if (!true) {
        switch ((v1 + (v1 + v1)) + "\n\t") {
        case "a":
            switch "hello" {
            case "":
                if (!(0 <= -2147483648)) {
                    var v23 uint16 = 2
                    v3 = (42 == int64(v23))
                    var v24 uint8 = 145
                    _ = v23
                    _ = v24
                } else {
                    v1 = ((v1 + ("a" + v1)) + "foo bar")
                }
            }
        default:
            v1 = "foo bar"
        }
    } else {
        v25 := GenS2{F0: (v1 + "\n\t"), F1: 1, F2: false}
        _ = v25
    }
    v11 = (v11 & (^(v11)))
    var v26 int64 = -9223372036854775808
    var v27 uint64 = 526
    v28 := GenS3{F0: 42, F1: v1}
    fmt.Println(v1, v3)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
}
