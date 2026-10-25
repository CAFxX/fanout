package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32; F1 uint32; F2 uint8 }

func main() {
    var v1 uint = 18446744073709551615
    v2 := []int{32768, 32768}
    v3 := []string{("a" + "foo bar"), "foo bar", ""}
    for fi4 := 0; fi4 < 10; fi4++ {
        if false {
            v3 = append(v3, ("a" + "hello"))
        }
        _ = fi4
    }
    v1 -= v1
    var v5 string = ("\n\t" + "\n\t")
    for fi6 := 0; fi6 < 11; fi6++ {
        var v7 int64 = 1
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 18; fi8++ {
        var v9 int16 = -1
        _ = fi8
        _ = v9
    }
    if true {
        v2 = append(v2, 808)
        v1 = 2
    } else {
        var v10 int = 32767
        if false {
            v11 := GenS1{F0: 42, F1: 42}
            fmt.Println(v5, v1, v5)
            var v12 float32 = 0.1
            _ = v11
            _ = v12
        }
        switch "\n\t" {
        case "hello":
            v13 := GenS2{F0: -1}
            _ = v13
        case "":
            if (0 >= 0) {
                switch v10 {
                case -2147483648:
                    if false {
                        var v14 int32 = -7
                        switch "hello" {
                        case "\n\t":
                            switch 42 {
                            case -76:
                                v15 := GenS1{F0: -114, F1: 65535}
                                _ = v15
                            case 9223372036854775807:
                                var v16 string = v5
                                _ = v16
                            default:
                                v14 = 340
                            }
                        case "hello":
                            switch (v10 * 2147483647) {
                            case 42:
                                fmt.Println(v14, v14, v5)
                            default:
                                v14 = int32(v10)
                            }
                        default:
                            v17 := GenS2{F0: (-(v14))}
                            _ = v17
                        }
                        fmt.Println(v10, v5)
                        _ = v14
                    }
                default:
                    var v18 int64 = 9223372036854775807
                    _ = v18
                }
                var v19 uint = (v1 * (^(v1)))
                _ = v19
            } else {
                fmt.Println(v10)
                v20 := GenS3{F0: 832, F1: 605, F2: 0}
                _ = v20
            }
        case "\n\t":
            var v21 int32 = 475
            _ = v21
        }
        _ = v10
    }
    fmt.Println(v5, v5)
    fmt.Println(v1, v5)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v5
}
