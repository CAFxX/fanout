package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64 }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 int32; F1 string; F2 uint }

func main() {
    var v1 uint = 18446744073709551615
    var v2 string = ("\n\t" + "a")
    v3 := make(map[int]int)
    v1 = v1
    delete(v3, 2147483647)
    v4 := []int{-288}
    var v5 int32 = 1
    v6 := v3[32767]
    fmt.Println(v1, v6, v6)
    var v7 int8 = -1
    for fi8 := 0; fi8 < 5; fi8++ {
        v2 = "foo bar"
        v9 := GenS1{F0: (v6 >> ((255) & 63)), F1: -337}
        _ = fi8
        _ = v9
    }
    if (42 != 534) {
        if true {
            var v10 uint8 = 135
            _ = v10
        }
        v3[(v6 - (v6 << (((v1 + (^(v1)))) & 63)))] = -1
        fmt.Println(v6, v5)
    } else {
        for fi11 := 0; fi11 < 1; fi11++ {
            _ = len(v4)
            _ = fi11
        }
    }
    if true {
        var v12 int = (^(v6))
        v6 = v6
        for fi13 := 0; fi13 < 16; fi13++ {
            fmt.Println(v7, v7)
            _ = fi13
        }
        _ = v12
    }
    if false {
        for fi14 := 0; fi14 < 13; fi14++ {
            v6 = fi14
            _ = fi14
        }
        fmt.Println(v1, v5, v6)
        fmt.Println(v5, v5)
    } else {
        if true {
            v15 := GenS1{F0: (v6 % ((v6) | 1)), F1: 0}
            v16 := GenS2{F0: 42, F1: (v6 & v6)}
            _ = v15
            _ = v16
        } else {
            for fi17 := 0; fi17 < 8; fi17++ {
                if true {
                    for fi18 := 0; fi18 < 5; fi18++ {
                        v19 := GenS1{F0: (fi18 - (-(v6))), F1: -1}
                        _ = fi18
                        _ = v19
                    }
                    v20 := GenS3{F0: -1, F1: "foo bar", F2: 903}
                    v4 = append(v4, (v6 - (v6 * 2147483647)))
                    _ = v20
                }
                _ = fi17
            }
            _ = len(v4)
            switch ("hello" + (v2 + "")) {
            case "a":
                fmt.Println(v1, v1, v5)
            default:
                var v21 string = "hello"
                _ = v21
            }
        }
        if ((32767 >= -1) && true) {
            var v22 uint32 = 2
            var v23 bool = (!false)
            var v24 int8 = (v7 ^ v7)
            _ = v22
            _ = v23
            _ = v24
        }
        v25 := GenS1{F0: v6, F1: 160}
        _ = v25
    }
    v3[(v6 + (v6 % ((v6) | 1)))] = 32768
    var v26 int16 = 32767
    v26 -= v26
    var v27 int = (v6 | v6)
    _ = len(v3)
    v28 := v3[v6]
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
