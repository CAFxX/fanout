package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32; F1 int32 }

func main() {
    v1 := []uint32{2, 2147483648}
    v2 := make(map[int]int)
    var v3 uint64 = 0
    var v4 uint64 = (^(v3))
    var v5 string = "a"
    _ = len(v2)
    v6 := GenS1{F0: 82, F1: 42}
    v7 := GenS1{F0: 662, F1: 0}
    v8 := GenS2{F0: -1}
    fmt.Println(v4, v4, v4)
    var v9 uint = 2
    if ((0 == 128) || true) {
        if len(v1) > 0 {
            v1[0] = 2
        }
        if false {
            for fi10 := 0; fi10 < 4; fi10++ {
                if (false || (!true)) {
                    _ = len(v2)
                    switch fi10 {
                    case -265:
                        var v11 uint8 = 127
                        _ = v11
                    }
                }
                var v12 string = "\n\t"
                _ = fi10
                _ = v12
            }
        } else {
            for fi13 := 0; fi13 < 11; fi13++ {
                v14 := GenS2{F0: int32(v4)}
                for fi15 := 0; fi15 < 12; fi15++ {
                    v9 -= (v9 ^ 404)
                    if len(v1) > 0 {
                        v1[0] = 1
                    }
                    switch (fi15 >> (((v9 ^ 65535)) & 63)) {
                    case -991:
                        if len(v1) > 0 {
                            v1[0] = 0
                        }
                    case -1:
                        v3 -= (v4 & (v3 / (((^(v3))) | 1)))
                    default:
                        v16 := v2[656]
                        if false {
                            switch v5 {
                            case "\n\t":
                                v4 = v4
                            case "":
                                var v17 uint32 = 0
                                _ = v17
                            case "foo bar":
                                var v18 float64 = 0.1
                                _ = v18
                            default:
                                fmt.Println(v4, v16)
                            }
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        }
                        _ = v16
                    }
                    _ = fi15
                }
                _ = fi13
                _ = v14
            }
            if (false && true) {
                v1 = append(v1, 930)
                for fi19 := 0; fi19 < 17; fi19++ {
                    v9 = 18446744073709551615
                    v2[(fi19 % (((fi19 + fi19)) | 1))] = fi19
                    _ = fi19
                }
                var v20 uint8 = 128
                _ = v20
            } else {
                v4 = (^(v4))
                var v21 uint32 = uint32(v4)
                _ = v21
            }
        }
        v22 := GenS2{F0: 0}
        _ = v22
    }
    if (42 == -1) {
        if (!(56 < 459)) {
            v5 = "a"
        } else {
            var v23 int = 2147483647
            v24 := GenS3{F0: 2147483647, F1: 114}
            _ = v23
            _ = v24
        }
        v9 += (v9 ^ v9)
        fmt.Println(v9)
    } else {
        v9 = (^(v9))
        for fi25 := 0; fi25 < 10; fi25++ {
            v9 = (v9 / ((v9) | 1))
            fmt.Println(v3, fi25)
            for fi26 := 0; fi26 < 7; fi26++ {
                v27 := GenS1{F0: 440, F1: 1}
                switch (v5 + v5) {
                case "foo bar":
                    var v28 float64 = (1.5 - 0.0)
                    _ = v28
                case "\n\t":
                    delete(v2, (fi26 * (fi25 | 576)))
                case "hello":
                    v5 = ((v5 + ((v5 + v5) + v5)) + "foo bar")
                default:
                    v2[fi25] = (fi25 * int(v9))
                }
                _ = fi26
                _ = v27
            }
            _ = fi25
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
