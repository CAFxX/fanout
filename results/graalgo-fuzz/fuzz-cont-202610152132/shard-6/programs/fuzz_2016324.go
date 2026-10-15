package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64 }

func main() {
    var v1 uint64 = 606
    v2 := make(map[uint32]uint32)
    var v3 int = 128
    switch "hello" {
    case "a":
        var v4 int32 = 339
        _ = v4
    }
    for fi5 := 0; fi5 < 13; fi5++ {
        var v6 uint = 18446744073709551615
        _ = fi5
        _ = v6
    }
    fmt.Println(v3, v1)
    var v7 bool = (1 == -32768)
    if (v7 && true) {
        for fi8 := 0; fi8 < 13; fi8++ {
            var v9 float32 = 3.14159
            _ = fi8
            _ = v9
        }
    } else {
        delete(v2, 237)
    }
    var v10 int32 = 1
    var v11 uint16 = 723
    if false {
        v12 := v2[2147483647]
        v2[v12] = (^(v12))
        _ = v12
    } else {
        v2[42] = 4294967295
    }
    var v13 uint8 = 174
    var v14 uint32 = 883
    var v15 int32 = (v10 | (v10 >> ((858) & 63)))
    v16 := GenS1{F0: v1}
    v17 := []int{}
    v11 = 255
    fmt.Println(v1)
    v13 += 1
    if len(v17) >= 2 {
        v17 = v17[1:len(v17)]
    }
    if true {
        for fi18 := 0; fi18 < 14; fi18++ {
            switch "foo bar" {
            case "\n\t":
                var v19 uint = 1
                _ = v19
            case "hello":
                for fi20 := 0; fi20 < 18; fi20++ {
                    var v21 int16 = 209
                    _ = fi20
                    _ = v21
                }
            case "":
                v7 = (!v7)
            default:
                v14 = (v14 - (v14 + 525))
            }
            for fi22 := 0; fi22 < 4; fi22++ {
                if (!v7) {
                    var v23 int16 = -656
                    _ = len(v17)
                    fmt.Println(v14)
                    _ = v23
                } else {
                    delete(v2, v14)
                    fmt.Println(v1, v14)
                    if false {
                        var v24 int = fi18
                        var v25 uint = 1
                        _ = v24
                        _ = v25
                    } else {
                        v10 -= (v10 / (((v15 & v15)) | 1))
                        if false {
                            var v26 int16 = -32768
                            v27 := GenS2{F0: int16(v13)}
                            var v28 uint = 208
                            _ = v26
                            _ = v27
                            _ = v28
                        }
                        if true {
                            if (!true) {
                                fmt.Println(v1)
                                if true {
                                    v7 = ((^(v1)) < v1)
                                    v13 += 196
                                }
                                v2[(v14 + v14)] = uint32(v1)
                            } else {
                                v29 := GenS3{F0: 213}
                                var v30 int = fi18
                                _ = v29
                                _ = v30
                            }
                            if ((3.14159 > 3.14159) || (-1 > 1)) {
                                fmt.Println(fi22, v1)
                            }
                            v2[(^(v14))] = v14
                        } else {
                            var v31 uint32 = (v14 * uint32(v11))
                            _ = v31
                        }
                        v11 = v11
                        v3 = 127
                        v32 := GenS2{F0: 42}
                        _ = v32
                    }
                }
                var v33 int32 = (v10 & (v10 / ((0) | 1)))
                v10 = int32(v13)
                _ = fi22
                _ = v33
            }
            var v34 string = "\n\t"
            _ = fi18
            _ = v34
        }
        var v35 uint = 513
        v14 -= uint32(v35)
        _ = v35
    } else {
        for fi36 := 0; fi36 < 4; fi36++ {
            if len(v17) > 0 {
                v17[0] = (v3 * (^(fi36)))
            }
            _ = fi36
        }
        v37 := GenS3{F0: -219}
        _ = v37
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v7
}
