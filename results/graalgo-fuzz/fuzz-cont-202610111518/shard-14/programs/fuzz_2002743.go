package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 uint8 }
type GenS2 struct { F0 int64; F1 uint }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[uint32]int)
    v3 := []uint32{65535, 4294967295}
    var v4 uint64 = 42
    var v5 uint16 = 0
    var v6 int32 = 0
    v3 = append(v3, 42)
    var v7 int8 = 0
    v6 -= v6
    switch "a" {
    case "hello":
        var v8 string = ("\n\t" + "hello")
        _ = v8
    case "a":
        var v9 string = "foo bar"
        _ = v9
    default:
        if true {
            fmt.Println(v4)
            if (2 != v5) {
                for fi10 := 0; fi10 < 7; fi10++ {
                    switch "\n\t" {
                    case "foo bar":
                        fmt.Println(v5)
                    case "a":
                        if false {
                            v11 := GenS1{F0: 40, F1: 42, F2: 127}
                            switch "a" {
                            case "":
                                v12 := GenS1{F0: 194, F1: 0, F2: 1}
                                _ = v12
                            default:
                                switch "\n\t" {
                                case "hello":
                                    v13 := GenS1{F0: 43, F1: 2, F2: 42}
                                    _ = v13
                                case "":
                                    for fi14 := 0; fi14 < 6; fi14++ {
                                        var v15 bool = false
                                        _ = fi14
                                        _ = v15
                                    }
                                case "\n\t":
                                    v16 := GenS1{F0: 127, F1: 0, F2: 1}
                                    _ = v16
                                }
                            }
                            _ = v11
                        } else {
                            v7 = (^(v7))
                            switch ("\n\t" + "hello") {
                            case "hello":
                                v4 -= v4
                            case "foo bar":
                                switch (-(fi10)) {
                                case 4294967295:
                                    switch -282 {
                                    case 4294967295:
                                        fmt.Println(v5, v5)
                                    case 2147483647:
                                        switch -1 {
                                        case 1:
                                            var v17 int8 = v7
                                            _ = v17
                                        case 128:
                                            switch fi10 {
                                            case -9223372036854775808:
                                                var v18 int32 = v6
                                                _ = v18
                                            }
                                        case 4294967295:
                                            fmt.Println(v7, v4)
                                        default:
                                            for fi19 := 0; fi19 < 16; fi19++ {
                                                v7 += v7
                                                var v20 uint16 = 1
                                                _ = fi19
                                                _ = v20
                                            }
                                        }
                                    case 128:
                                        fmt.Println(v7, v6, v4)
                                    default:
                                        var v21 uint64 = 357
                                        _ = v21
                                    }
                                case 32767:
                                    var v22 string = ""
                                    _ = v22
                                case 32768:
                                    _ = len(v1)
                                }
                            case "\n\t":
                                v23 := GenS1{F0: 2, F1: 508, F2: 128}
                                _ = v23
                            default:
                                var v24 uint64 = 739
                                _ = v24
                            }
                        }
                    case "\n\t":
                        v4 -= (v4 ^ v4)
                    }
                    _ = fi10
                }
                v25 := v2[1]
                v26 := GenS2{F0: 9223372036854775807, F1: 1}
                _ = v25
                _ = v26
            }
        }
    }
    var v27 bool = false
    var v28 uint8 = 127
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
