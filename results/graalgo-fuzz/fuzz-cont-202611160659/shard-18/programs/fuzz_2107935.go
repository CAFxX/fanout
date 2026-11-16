package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    var v1 int64 = 42
    v2 := []string{("" + ("a" + "\n\t")), (("hello" + "") + "\n\t"), "\n\t"}
    var v3 string = ("\n\t" + (("" + "foo bar") + "\n\t"))
    var v4 int = 32768
    v5 := make(map[uint32]uint32)
    switch v3 {
    case "\n\t":
        v5[257] = 42
    case "hello":
        for fi6 := 0; fi6 < 13; fi6++ {
            fmt.Println(v1, fi6, fi6)
            switch v3 {
            case "a":
                v1 = 1
            default:
                for fi7 := 0; fi7 < 8; fi7++ {
                    v3 = v3
                    v1 = (v1 & v1)
                    if true {
                        if (18446744073709551615 > 2) {
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            var v8 uint = 0
                            switch "a" {
                            case "foo bar":
                                var v9 int8 = 42
                                _ = v9
                            default:
                                v1 += (v1 >> ((1) & 63))
                            }
                            var v10 float64 = 1.5
                            switch "" {
                            case "\n\t":
                                var v11 string = "\n\t"
                                _ = v11
                            }
                            _ = v10
                            _ = v8
                        }
                    } else {
                        if ((0 >= -622) && true) {
                            v1 = (^(v1))
                        }
                    }
                    _ = fi7
                }
            }
            _ = fi6
        }
    case "a":
        var v12 int64 = v1
        _ = v12
    default:
        v1 = (v1 / (((v1 << ((289) & 63))) | 1))
    }
    fmt.Println(v3, v1, v3)
    v5[24] = 24
    v3 = v3
    fmt.Println(v1, v4)
    v13 := GenS1{F0: (v4 | v4)}
    var v14 int16 = 32767
    v15 := GenS1{F0: (v4 ^ v4)}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
