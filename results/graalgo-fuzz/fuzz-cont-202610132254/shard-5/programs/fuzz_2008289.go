package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 uint }
type GenS2 struct { F0 uint; F1 int8 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint64 = 800
    v3 := make(map[int]int)
    var v4 bool = ((0.0 / ((-2.25 / 0.1) - 1.5)) < float32((0.1 / -2.25)))
    var v5 uint8 = 0
    fmt.Println(v2)
    if v4 {
        switch 127 {
        case 128:
            switch ("\n\t" + (("a" + ("foo bar" + "hello")) + "")) {
            case "hello":
                v6 := []int{0}
                _ = v6
            case "":
                v7 := GenS1{F0: 173, F1: 543, F2: 246}
                _ = v7
            case "a":
                if (65535 != 4294967295) {
                    var v8 float64 = ((float64((float32(0.1) * -2.25)) * -2.25) * 0.0)
                    fmt.Println(v5)
                    switch "hello" {
                    case "\n\t":
                        switch 0 {
                        case -895:
                            var v9 string = ("\n\t" + "")
                            _ = v9
                        case -492:
                            for fi10 := 0; fi10 < 14; fi10++ {
                                v5 = v5
                                var v11 uint32 = 766
                                v12 := GenS2{F0: 42, F1: 100}
                                _ = fi10
                                _ = v11
                                _ = v12
                            }
                        case 886:
                            var v13 int = 390
                            _ = v13
                        }
                    default:
                        v8 -= (v8 + -2.25)
                    }
                    _ = v8
                } else {
                    v14 := []int{-266, 1}
                    fmt.Println(v2, v4, v4)
                    fmt.Println(v2, v2, v5)
                    _ = v14
                }
            }
        }
        v4 = (v4 || (("a" < (("a" + ("" + "\n\t")) + "")) && v4))
    } else {
        var v15 float32 = (1.5 * float32(float64(-2.25)))
        for fi16 := 0; fi16 < 17; fi16++ {
            v17 := []int{fi16, fi16, fi16}
            fmt.Println(v4, v5, v5)
            _ = fi16
            _ = v17
        }
        _ = v15
    }
    fmt.Println(v4, v5, v5)
    for fi18 := 0; fi18 < 8; fi18++ {
        fmt.Println(v5, v5)
        var v19 int8 = 0
        for fi20 := 0; fi20 < 5; fi20++ {
            v21 := v3[1]
            v19 -= 15
            _ = fi20
            _ = v21
        }
        _ = fi18
        _ = v19
    }
    fmt.Println(v4)
    var v22 uint = 65535
    var v23 int32 = 2147483647
    if v4 {
        fmt.Println(v4, v23)
    } else {
        v23 -= (-(v23))
        v23 = v23
    }
    for fi24 := 0; fi24 < 9; fi24++ {
        v22 += (v22 & 255)
        _ = fi24
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
