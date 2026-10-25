package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 int8 }

func main() {
    v1 := []uint32{110, 220, 421, 19}
    var v2 bool = (-1 <= 32767)
    v3 := []uint32{2147483647, 4294967295, 2147483648, 2147483647}
    var v4 float32 = float32(1.5)
    fmt.Println(v4, v2)
    v4 = (v4 / (v4 - v4))
    var v5 uint64 = 1
    fmt.Println(v2, v2, v4)
    var v6 uint16 = 247
    var v7 uint = 487
    if len(v3) > 0 {
        v3[0] = 2147483647
    }
    var v8 uint16 = (v6 ^ (v6 % ((v6) | 1)))
    for fi9 := 0; fi9 < 7; fi9++ {
        var v10 uint16 = (v6 & 0)
        var v11 uint = uint(v5)
        _ = fi9
        _ = v10
        _ = v11
    }
    var v12 int16 = 1
    switch "\n\t" {
    case "a":
        _ = len(v1)
    case "foo bar":
        if (!v2) {
            if v2 {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v13 string = "\n\t"
                switch (((v13 + ((v13 + "a") + "")) + (v13 + ("hello" + v13))) + (v13 + "a")) {
                case "a":
                    var v14 int = int(v7)
                    _ = v14
                default:
                    if (!v2) {
                        for fi15 := 0; fi15 < 5; fi15++ {
                            fmt.Println(v7, v5, v8)
                            switch 4294967295 {
                            case -2147483648:
                                fmt.Println(fi15, v12)
                            case 128:
                                if true {
                                    if (((v13 + "a") + v13) >= "foo bar") {
                                        var v16 int64 = 42
                                        _ = v16
                                    } else {
                                        switch fi15 {
                                        case 42:
                                            for fi17 := 0; fi17 < 13; fi17++ {
                                                v18 := GenS1{F0: 4294967295, F1: (^(v8)), F2: -62}
                                                v7 = v7
                                                _ = fi17
                                                _ = v18
                                            }
                                        case -420:
                                            for fi19 := 0; fi19 < 13; fi19++ {
                                                var v20 string = (("foo bar" + v13) + (v13 + v13))
                                                var v21 uint16 = (^(v8))
                                                _ = fi19
                                                _ = v20
                                                _ = v21
                                            }
                                        }
                                    }
                                }
                            case 4294967295:
                                for fi22 := 0; fi22 < 6; fi22++ {
                                    v1 = append(v1, 681)
                                    _ = fi22
                                }
                            }
                            var v23 bool = v2
                            _ = fi15
                            _ = v23
                        }
                    }
                }
                _ = v13
            } else {
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
        }
    }
    v24 := GenS1{F0: 1, F1: (v6 % ((v8) | 1)), F2: 0}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
