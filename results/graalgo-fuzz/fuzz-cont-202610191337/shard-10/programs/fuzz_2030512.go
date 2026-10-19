package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint8 = 0
    var v3 int64 = -213
    v4 := make(map[uint32]uint32)
    switch "\n\t" {
    case "":
        v5 := []int{}
        _ = v5
    case "\n\t":
        fmt.Println(v2, v2)
    case "a":
        if (!(!(0 == 255))) {
            switch -619 {
            case 1:
                for fi6 := 0; fi6 < 2; fi6++ {
                    fmt.Println(v3)
                    _ = fi6
                }
            default:
                for fi7 := 0; fi7 < 7; fi7++ {
                    v8 := []int{867, (fi7 | (fi7 - 127)), (fi7 / ((fi7) | 1))}
                    _ = fi7
                    _ = v8
                }
            }
            v9 := GenS1{F0: -1, F1: 1}
            switch "a" {
            case "a":
                var v10 uint16 = 632
                _ = v10
            case "hello":
                switch "hello" {
                case "foo bar":
                    var v11 int8 = 14
                    _ = v11
                default:
                    switch 42 {
                    case -483:
                        for fi12 := 0; fi12 < 20; fi12++ {
                            switch fi12 {
                            case -205:
                                v13 := []int{fi12, 42}
                                _ = v13
                            case 0:
                                v1[2147483647] = (850 > 4294967295)
                            case 2147483647:
                                if (false && true) {
                                    v2 = (v2 & (^(v2)))
                                }
                            }
                            v3 += v3
                            _ = fi12
                        }
                    case -2147483648:
                        if true {
                            switch -1 {
                            case 634:
                                var v14 int32 = -1
                                _ = v14
                            case -1:
                                var v15 uint8 = 0
                                _ = v15
                            case -885:
                                v3 = -1
                            }
                        } else {
                            fmt.Println(v3, v2)
                            for fi16 := 0; fi16 < 3; fi16++ {
                                v2 = 127
                                v17 := GenS1{F0: -723, F1: 65535}
                                var v18 float64 = (3.14159 / 3.14159)
                                _ = fi16
                                _ = v17
                                _ = v18
                            }
                        }
                    }
                }
            case "":
                var v19 int8 = 0
                _ = v19
            }
            _ = v9
        }
    default:
        v20 := GenS1{F0: 42, F1: 65535}
        _ = v20
    }
    v2 = (v2 % ((v2) | 1))
    var v21 uint = 45
    v2 = v2
    fmt.Println(v3)
    fmt.Println(v2, v3)
    fmt.Println(v3, v21)
    v21 = 882
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
