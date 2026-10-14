package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 int8 }
type GenS3 struct { F0 uint8; F1 uint64; F2 int8 }

func main() {
    v1 := make(map[int]string)
    var v2 string = "\n\t"
    v3 := make(map[int]string)
    v4 := make(map[string]bool)
    var v5 uint64 = 496
    for fi6 := 0; fi6 < 12; fi6++ {
        v7 := GenS1{F0: 32767}
        for fi8 := 0; fi8 < 6; fi8++ {
            var v9 int = 42
            _ = fi8
            _ = v9
        }
        var v10 int32 = 1
        _ = fi6
        _ = v10
        _ = v7
    }
    fmt.Println(v2)
    for fi11 := 0; fi11 < 14; fi11++ {
        switch fi11 {
        case 4294967295:
            v12 := GenS1{F0: 65535}
            _ = v12
        case 2147483647:
            switch fi11 {
            case 549:
                var v13 uint32 = 427
                _ = v13
            case 42:
                v5 += 634
            case 0:
                switch ((v2 + ((v2 + (v2 + (v2 + v2))) + v2)) + "foo bar") {
                case "\n\t":
                    for fi14 := 0; fi14 < 14; fi14++ {
                        var v15 float64 = float64(1.5)
                        _ = fi14
                        _ = v15
                    }
                case "hello":
                    v5 -= (^(v5))
                case "a":
                    fmt.Println(v5, v2)
                default:
                    v16 := GenS1{F0: 32767}
                    _ = v16
                }
            default:
                v17 := GenS1{F0: uint16(v5)}
                _ = v17
            }
        case 0:
            if false {
                var v18 bool = false
                var v19 string = v2
                _ = v18
                _ = v19
            } else {
                fmt.Println(v2, v2, fi11)
                if true {
                    for fi20 := 0; fi20 < 15; fi20++ {
                        var v21 uint64 = (^(v5))
                        var v22 int32 = -1
                        _ = fi20
                        _ = v21
                        _ = v22
                    }
                    if true {
                        v2 = v2
                        v5 = (v5 / ((v5) | 1))
                    } else {
                        var v23 int16 = -256
                        v24 := GenS2{F0: 100, F1: -116}
                        _ = v23
                        _ = v24
                    }
                    fmt.Println(v5, v5, v5)
                }
            }
        }
        v25 := GenS3{F0: 99, F1: v5, F2: 1}
        v26 := []int{fi11, fi11}
        _ = fi11
        _ = v25
        _ = v26
    }
    v3[-1] = (((("hello" + v2) + ((v2 + v2) + "a")) + "foo bar") + v2)
    v27 := GenS3{F0: 241, F1: v5, F2: 39}
    for fi28 := 0; fi28 < 14; fi28++ {
        var v29 int32 = 1
        _ = fi28
        _ = v29
    }
    v5 = v5
    var v30 float32 = (0.0 / -2.25)
    switch v2 {
    case "hello":
        v5 -= v5
    case "":
        for fi31 := 0; fi31 < 2; fi31++ {
            switch "a" {
            case "a":
                v1[fi31] = ("foo bar" + "a")
            }
            _ = fi31
        }
    }
    v2 = v2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
