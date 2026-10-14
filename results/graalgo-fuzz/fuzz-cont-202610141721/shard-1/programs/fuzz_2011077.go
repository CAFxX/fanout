package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 uint; F1 int8; F2 uint64 }

func main() {
    var v1 uint32 = 1
    var v2 int32 = -2147483648
    v3 := make(map[uint32]bool)
    var v4 uint16 = 42
    v5 := GenS1{F0: 4294967295, F1: 42}
    if true {
        fmt.Println(v4, v1, v1)
    } else {
        var v6 string = ""
        v6 = "hello"
        var v7 int32 = (v2 + v2)
        _ = v6
        _ = v7
    }
    switch "foo bar" {
    case "hello":
        if true {
            fmt.Println(v2)
            var v8 string = "hello"
            v2 -= 666
            _ = v8
        }
    }
    for fi9 := 0; fi9 < 9; fi9++ {
        switch fi9 {
        case 0:
            v10 := GenS2{F0: 18446744073709551615, F1: -24, F2: 195}
            _ = v10
        default:
            var v11 int64 = 9223372036854775807
            _ = v11
        }
        if (-128 == 121) {
            v12 := GenS1{F0: 0, F1: -1}
            if ("foo bar" < "a") {
                switch -2147483648 {
                case -372:
                    v13 := GenS2{F0: 65535, F1: 42, F2: 898}
                    _ = v13
                case 32767:
                    if false {
                        for fi14 := 0; fi14 < 15; fi14++ {
                            var v15 int64 = -474
                            v16 := []int{}
                            _ = fi14
                            _ = v15
                            _ = v16
                        }
                        v4 = 7
                        v4 += 2
                    } else {
                        v17 := GenS1{F0: 65535, F1: 56}
                        v4 = v4
                        for fi18 := 0; fi18 < 7; fi18++ {
                            v2 -= v2
                            v1 -= v1
                            _ = fi18
                        }
                        _ = v17
                    }
                }
            } else {
                v19 := []int{0}
                for fi20 := 0; fi20 < 6; fi20++ {
                    v1 = 2
                    _ = fi20
                }
                _ = v19
            }
            _ = v12
        }
        _ = fi9
    }
    v4 = 1
    v21 := GenS2{F0: 2, F1: 24, F2: 922}
    _ = len(v3)
    v3[v1] = (1 <= -346)
    fmt.Println(v2, v2, v4)
    for fi22 := 0; fi22 < 9; fi22++ {
        v23 := []int{-976, fi22, fi22, -115}
        _ = fi22
        _ = v23
    }
    v4 += v4
    for fi24 := 0; fi24 < 11; fi24++ {
        v25 := GenS1{F0: 552, F1: -31}
        var v26 uint8 = 2
        _ = fi24
        _ = v25
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
