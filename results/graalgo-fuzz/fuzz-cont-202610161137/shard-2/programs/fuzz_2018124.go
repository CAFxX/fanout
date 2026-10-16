package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
func genFn0(p0 bool, p1 uint8) int16 {
    p1 = 1
    fmt.Println(p1)
    p1 = p1
    for fi6 := 0; fi6 < 12; fi6++ {
        v7 := []int{fi6, (fi6 * fi6)}
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    return -1
}

func main() {
    var v1 float64 = 0.0
    v2 := []string{"foo bar", "", ("foo bar" + "hello")}
    v3 := make(map[uint32]string)
    var v4 int16 = 42
    v5 := make(map[string]uint32)
    for fi8 := 0; fi8 < 12; fi8++ {
        v1 = float64((-2.25 - 0.0))
        for fi9 := 0; fi9 < 7; fi9++ {
            v10 := GenS1{F0: "\n\t", F1: 2}
            fmt.Println(fi8, fi9, fi9)
            fmt.Println(v4, v1)
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    v11 := v5["foo bar"]
    v4 = (v4 ^ v4)
    _ = len(v2)
    fmt.Println(v4, v4)
    var v12 float64 = (0.0 - 3.14159)
    for fi13 := 0; fi13 < 3; fi13++ {
        _ = len(v5)
        if true {
            var v14 uint = 42
            _ = len(v2)
            v14 = (v14 ^ (v14 / ((v14) | 1)))
            _ = v14
        }
        _ = fi13
    }
    v1 -= (v12 + v1)
    delete(v5, "foo bar")
    switch 42 {
    case 127:
        _ = len(v5)
    case 9223372036854775807:
        for fi15 := 0; fi15 < 12; fi15++ {
            v1 = v1
            for fi16 := 0; fi16 < 12; fi16++ {
                switch (fi16 / (((fi15 >> ((699) & 63))) | 1)) {
                case 454:
                    v17 := GenS1{F0: ("\n\t" + "a"), F1: 255}
                    _ = v17
                case 4294967295:
                    fmt.Println(v11)
                case 32767:
                    var v18 int16 = (v4 << ((2) & 63))
                    _ = v18
                }
                _ = fi16
            }
            _ = fi15
        }
    default:
        v12 = 0.0
    }
    v4 += (v4 * (v4 ^ (v4 & genFn0((true && true), 99))))
    for fi19 := 0; fi19 < 11; fi19++ {
        if false {
            switch ((("a" + "\n\t") + "") + "a") {
            case "\n\t":
                for fi20 := 0; fi20 < 18; fi20++ {
                    switch fi19 {
                    case 746:
                        var v21 uint32 = 1
                        _ = v21
                    case 614:
                        fmt.Println(v1)
                    case 1:
                        v3[v11] = "hello"
                    }
                    _ = fi20
                }
            case "foo bar":
                for fi22 := 0; fi22 < 12; fi22++ {
                    var v23 uint8 = 128
                    fmt.Println(v11, v11, v4)
                    v11 = (v11 >> ((0) & 63))
                    _ = fi22
                    _ = v23
                }
            default:
                var v24 uint = 42
                _ = v24
            }
            var v25 uint8 = 255
            delete(v3, (v11 * v11))
            _ = v25
        }
        fmt.Println(v12, v11, v12)
        _ = fi19
    }
    fmt.Println(v1, v12, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
