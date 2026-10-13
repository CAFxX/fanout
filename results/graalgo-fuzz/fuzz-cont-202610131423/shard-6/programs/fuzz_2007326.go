package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool; F1 uint64 }
func genFn0(p0 uint8) int64 {
    fmt.Println(p0)
    fmt.Println(p0)
    _ = p0
    return -723
}

func main() {
    var v1 uint32 = 2147483647
    v2 := make(map[string]uint32)
    var v3 uint8 = 0
    var v4 float64 = 0.0
    var v5 uint = 510
    fmt.Println(v3, v4, v5)
    v3 += 128
    var v6 uint32 = v1
    v5 = (v5 / ((v5) | 1))
    v5 = v5
    if (false || (1 >= 853)) {
        v7 := GenS1{F0: -386}
        v4 += v4
        if true {
            fmt.Println(v4)
            delete(v2, ((("\n\t" + "a") + ("\n\t" + "")) + ("" + (("foo bar" + "a") + "a"))))
        } else {
            for fi8 := 0; fi8 < 5; fi8++ {
                var v9 uint = (v5 + uint(v3))
                var v10 uint32 = (v6 & (^(v1)))
                v11 := []int{}
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
            for fi12 := 0; fi12 < 1; fi12++ {
                for fi13 := 0; fi13 < 8; fi13++ {
                    v4 = (v4 + v4)
                    _ = fi13
                }
                switch "a" {
                case "foo bar":
                    switch "\n\t" {
                    case "a":
                        v2["hello"] = 65535
                    case "\n\t":
                        for fi14 := 0; fi14 < 20; fi14++ {
                            switch fi14 {
                            case -2147483648:
                                var v15 int = fi12
                                _ = v15
                            default:
                                v16 := []int{(fi14 << (((^(v5))) & 63)), (fi12 - (fi14 >> (((v5 / (((v5 << ((431) & 63))) | 1))) & 63))), fi12, -218}
                                _ = v16
                            }
                            _ = fi14
                        }
                    default:
                        var v17 uint32 = v6
                        _ = v17
                    }
                case "\n\t":
                    switch "foo bar" {
                    case "a":
                        v1 = (v6 | (^(v6)))
                    case "\n\t":
                        _ = len(v2)
                    case "":
                        v18 := []int{}
                        _ = v18
                    }
                default:
                    var v19 float32 = 0.1
                    _ = v19
                }
                _ = fi12
            }
        }
        _ = v7
    }
    if (false && (0.1 >= (1.5 - 3.14159))) {
        if (genFn0((v3 | uint8(v5))) >= -9223372036854775808) {
            for fi20 := 0; fi20 < 15; fi20++ {
                var v21 int64 = 42
                switch (fi20 | fi20) {
                case 0:
                    for fi22 := 0; fi22 < 19; fi22++ {
                        fmt.Println(fi20, fi22, v5)
                        v23 := GenS1{F0: 1}
                        switch 42 {
                        case 128:
                            fmt.Println(v21, v6, v1)
                        default:
                            v21 += (v21 | (v21 / ((v21) | 1)))
                        }
                        v24 := GenS2{F0: ("\n\t" <= ("\n\t" + "")), F1: 0}
                        v21 = v21
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                case 9:
                    var v25 int32 = 2147483647
                    _ = v25
                }
                for fi26 := 0; fi26 < 6; fi26++ {
                    switch "hello" {
                    case "":
                        v3 = 250
                    }
                    _ = fi26
                }
                _ = fi20
                _ = v21
            }
        } else {
            v2[("" + "a")] = 1
            v3 -= (v3 >> ((v5) & 63))
            var v27 int16 = -1
            _ = v27
        }
        var v28 uint64 = uint64(v6)
        fmt.Println(v1, v4, v5)
        _ = v28
    }
    var v29 bool = (true || (9223372036854775807 > 25))
    if v29 {
        var v30 float64 = v4
        fmt.Println(v1, v6, v4)
        _ = v30
    }
    for fi31 := 0; fi31 < 9; fi31++ {
        v32 := GenS2{F0: v29, F1: 42}
        v3 = v3
        if v29 {
            var v33 float32 = 1.5
            for fi34 := 0; fi34 < 20; fi34++ {
                v35 := v2[""]
                v35 = v35
                _ = fi34
                _ = v35
            }
            switch (("hello" + ("foo bar" + "a")) + "a") {
            case "":
                v36 := []int{fi31, 4294967295, (fi31 & (^(fi31))), -6}
                _ = v36
            case "foo bar":
                v37 := GenS2{F0: v29, F1: 42}
                _ = v37
            }
            _ = v33
        }
        _ = fi31
        _ = v32
    }
    fmt.Println(v4, v5, v6)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
