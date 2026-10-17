package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int8 }

func main() {
    var v1 int64 = 42
    var v2 int = 317
    v3 := make(map[int]uint32)
    v4 := GenS1{F0: 158}
    var v5 uint64 = 129
    var v6 float64 = (((1e10 * 3.14159) * -2.25) - 0.1)
    var v7 float32 = -2.25
    v8 := []int{}
    switch v2 {
    case 1:
        v9 := GenS1{F0: 678}
        _ = v9
    case 127:
        var v10 int8 = -85
        _ = v10
    default:
        switch "a" {
        case "a":
            v6 = v6
        case "\n\t":
            if false {
                v2 = -944
            }
        case "foo bar":
            var v11 uint64 = (v5 >> ((65535) & 63))
            _ = v11
        }
    }
    var v12 uint16 = 2
    switch 806 {
    case -1:
        if false {
            switch "a" {
            case "":
                v1 -= v1
            default:
                var v13 int8 = 0
                _ = v13
            }
            for fi14 := 0; fi14 < 6; fi14++ {
                switch v2 {
                case -217:
                    v8 = append(v8, int(v1))
                case 128:
                    for fi15 := 0; fi15 < 15; fi15++ {
                        if false {
                            fmt.Println(v1)
                            fmt.Println(v5)
                        }
                        if (false && (false || (v1 == 1))) {
                            if len(v8) >= 2 {
                                v8 = v8[1:len(v8)]
                            }
                        } else {
                            var v16 uint16 = (^(v12))
                            _ = v16
                        }
                        _ = fi15
                    }
                }
                v17 := GenS2{F0: (v12 >> ((1) & 63))}
                _ = fi14
                _ = v17
            }
        } else {
            var v18 uint16 = (v12 / ((v12) | 1))
            switch ("" + ("hello" + ("foo bar" + "hello"))) {
            case "a":
                v18 = 1
            case "":
                v19 := GenS1{F0: 18446744073709551615}
                _ = v19
            case "hello":
                fmt.Println(v7, v7, v6)
            default:
                var v20 int8 = 75
                _ = v20
            }
            if false {
                if ((^(v2)) < v2) {
                    v1 -= v1
                    var v21 uint64 = v5
                    var v22 uint64 = 2
                    _ = v21
                    _ = v22
                }
            } else {
                for fi23 := 0; fi23 < 17; fi23++ {
                    for fi24 := 0; fi24 < 13; fi24++ {
                        if (uint(v5) != 1) {
                            delete(v3, fi24)
                            if (int32(v18) <= 922) {
                                if false {
                                    var v25 uint8 = 184
                                    _ = v25
                                } else {
                                    v26 := GenS2{F0: (v12 - (v12 * 2))}
                                    v27 := GenS3{F0: -11}
                                    _ = v26
                                    _ = v27
                                }
                                fmt.Println(v18)
                            } else {
                                var v28 bool = ((((^(v12)) == v12) || (1 == 47)) && true)
                                fmt.Println(v18, fi24)
                                _ = v28
                            }
                            v12 += (v18 / ((uint16(v5)) | 1))
                        }
                        v8 = append(v8, int(v1))
                        v29 := GenS3{F0: 1}
                        _ = fi24
                        _ = v29
                    }
                    _ = fi23
                }
            }
            _ = v18
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
