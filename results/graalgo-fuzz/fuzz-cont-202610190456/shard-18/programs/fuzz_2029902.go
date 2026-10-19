package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16; F1 uint }
type GenS3 struct { F0 int16 }

func main() {
    v1 := make(map[string]string)
    var v2 int32 = 537
    v3 := make(map[uint32]int)
    v2 = 42
    v4 := GenS1{F0: 183}
    v2 = (v2 ^ (v2 >> ((1) & 63)))
    if (9223372036854775807 < 9223372036854775807) {
        switch "a" {
        case "a":
            delete(v3, 0)
        case "\n\t":
            var v5 int32 = v2
            _ = v5
        case "":
            if (!true) {
                if ((87 <= 255) || (false || (0 == 577))) {
                    if true {
                        switch ("hello" + "") {
                        case "a":
                            var v6 uint16 = 42
                            _ = v6
                        case "foo bar":
                            v2 = v2
                        case "\n\t":
                            v2 = (v2 - 2147483647)
                        default:
                            fmt.Println(v2)
                        }
                        var v7 uint = 65535
                        _ = v7
                    }
                } else {
                    for fi8 := 0; fi8 < 17; fi8++ {
                        fmt.Println(v2, fi8)
                        for fi9 := 0; fi9 < 7; fi9++ {
                            switch (("\n\t" + "\n\t") + "") {
                            case "a":
                                switch fi9 {
                                case 32767:
                                    var v10 uint = 0
                                    _ = v10
                                case 604:
                                    switch "hello" {
                                    case "a":
                                        var v11 int32 = (v2 - v2)
                                        _ = v11
                                    case "\n\t":
                                        if (1 > 4294967295) {
                                            v12 := []int{(fi8 >> ((595) & 63))}
                                            v13 := GenS1{F0: 42}
                                            _ = v12
                                            _ = v13
                                        }
                                    }
                                default:
                                    var v14 uint32 = 65535
                                    _ = v14
                                }
                            case "\n\t":
                                if false {
                                    fmt.Println(fi8, fi8, fi8)
                                }
                            default:
                                var v15 int32 = (-(v2))
                                _ = v15
                            }
                            v16 := GenS2{F0: -1, F1: 1}
                            fmt.Println(fi9)
                            v17 := []int{-63}
                            var v18 int = (fi9 << ((57) & 63))
                            _ = fi9
                            _ = v16
                            _ = v17
                            _ = v18
                        }
                        for fi19 := 0; fi19 < 15; fi19++ {
                            v2 = v2
                            fmt.Println(fi19)
                            _ = fi19
                        }
                        _ = fi8
                    }
                }
            } else {
                var v20 uint64 = 2
                _ = v20
            }
        }
        var v21 int8 = -86
        _ = v21
    }
    v22 := v3[0]
    v23 := GenS3{F0: -924}
    fmt.Println(v22)
    for fi24 := 0; fi24 < 1; fi24++ {
        var v25 uint64 = 2
        _ = fi24
        _ = v25
    }
    v2 = 673
    var v26 float32 = 0.1
    v22 = -9223372036854775808
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v26
    _ = v3
    _ = v4
}
