package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 string; F1 uint8; F2 string }
type GenS3 struct { F0 uint8; F1 int8 }
type GenS4 struct { F0 uint16 }
func genFn0(p0 bool, p1 int32) string {
    switch -800 {
    case -848:
        if (255 < 0) {
            var v7 int32 = p1
            _ = v7
        } else {
            switch "foo bar" {
            case "":
                p1 -= p1
            }
        }
    default:
        fmt.Println(p1, p1)
    }
    for fi8 := 0; fi8 < 6; fi8++ {
        p0 = p0
        p0 = true
        v9 := GenS1{F0: (-447 <= -1)}
        _ = fi8
        _ = v9
    }
    p0 = p0
    v10 := []int{}
    _ = p0
    _ = p1
    _ = v10
    return ""
}

func main() {
    var v1 int16 = 1
    v2 := []string{"\n\t"}
    var v3 string = "a"
    var v4 int16 = 42
    var v5 int64 = 1
    var v6 string = "a"
    v4 = 910
    var v11 uint16 = 255
    var v12 uint16 = (v11 | (v11 - v11))
    fmt.Println(v6, v11, v5)
    switch (v6 + "") {
    case "":
        v13 := GenS1{F0: (-1 >= 127)}
        _ = v13
    }
    v11 += v12
    var v14 int8 = 0
    v15 := GenS2{F0: "foo bar", F1: 127, F2: "foo bar"}
    var v16 int = 32767
    v16 = v16
    for fi17 := 0; fi17 < 7; fi17++ {
        switch (("hello" + "hello") + v6) {
        case "a":
            v4 += 763
        default:
            v1 = (v1 * 0)
        }
        _ = fi17
    }
    v4 = (v1 | (v1 - (v4 | 1)))
    v18 := GenS3{F0: 2, F1: (v14 / (((v14 / ((v14) | 1))) | 1))}
    if (240 >= 134) {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    } else {
        v12 += 234
        fmt.Println(v16)
        _ = len(v2)
    }
    v5 += -139
    _ = len(v2)
    v19 := GenS4{F0: (v11 + (v11 << ((1) & 63)))}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v14
    _ = v15
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
