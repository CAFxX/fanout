package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int32; F1 uint; F2 uint }

func main() {
    var v1 uint32 = 42
    var v2 uint8 = 42
    var v3 int = -1
    var v4 uint64 = uint64(v1)
    v5 := []int{(v3 | v3)}
    switch ("hello" + "\n\t") {
    case "foo bar":
        v1 -= v1
    case "hello":
        var v6 uint32 = v1
        _ = v6
    default:
        for fi7 := 0; fi7 < 16; fi7++ {
            v3 = 32767
            _ = fi7
        }
    }
    v8 := GenS1{F0: 100, F1: 0}
    v4 -= v4
    v2 = (^(v2))
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v9 := GenS2{F0: 42}
    for fi10 := 0; fi10 < 17; fi10++ {
        if false {
            _ = len(v5)
        }
        _ = fi10
    }
    switch "\n\t" {
    case "foo bar":
        v4 = (v4 * v4)
    case "a":
        var v11 int = v3
        _ = v11
    default:
        var v12 bool = true
        _ = v12
    }
    v4 = (v4 / ((287) | 1))
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    _ = len(v5)
    for fi13 := 0; fi13 < 5; fi13++ {
        if false {
            fmt.Println(v1, v1)
            fmt.Println(v1)
            for fi14 := 0; fi14 < 8; fi14++ {
                var v15 int = fi14
                if (100 >= -62) {
                    v4 -= v4
                }
                v2 = v2
                _ = fi14
                _ = v15
            }
        } else {
            var v16 int16 = -1
            _ = v16
        }
        v17 := GenS3{F0: 42}
        _ = fi13
        _ = v17
    }
    for fi18 := 0; fi18 < 8; fi18++ {
        v19 := GenS3{F0: -1}
        if ((v2 / (((v2 | v2)) | 1)) >= 127) {
            fmt.Println(v4, v3, v4)
            v3 += fi18
            v3 = (fi18 + (v3 + (fi18 & (fi18 | 9223372036854775807))))
        }
        v4 = (^(v4))
        _ = fi18
        _ = v19
    }
    switch v3 {
    case 2147483647:
        v4 = (^(v4))
    default:
        if (-2147483648 == -775) {
            if false {
                if false {
                    v4 = (^(v4))
                } else {
                    v5 = append(v5, 816)
                }
            } else {
                v5 = append(v5, v3)
                v20 := GenS2{F0: -32768}
                v3 -= (v3 * v3)
                _ = v20
            }
            v4 -= v4
        } else {
            if false {
                var v21 uint = 18446744073709551615
                for fi22 := 0; fi22 < 8; fi22++ {
                    for fi23 := 0; fi23 < 4; fi23++ {
                        switch "\n\t" {
                        case "hello":
                            switch v3 {
                            case 1:
                                v24 := GenS4{F0: v2}
                                _ = v24
                            default:
                                var v25 float64 = (0.1 / 1.5)
                                _ = v25
                            }
                            var v26 int16 = -32768
                            if len(v5) > 0 {
                                v5[0] = -978
                            }
                            _ = v26
                        case "":
                            _ = len(v5)
                        }
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        v1 = v1
                        _ = fi23
                    }
                    if false {
                        v27 := GenS5{F0: 835, F1: (v21 + (v21 + 0)), F2: uint(v4)}
                        _ = v27
                    }
                    _ = len(v5)
                    _ = fi22
                }
                _ = v21
            }
            var v28 bool = true
            for fi29 := 0; fi29 < 12; fi29++ {
                v5 = append(v5, fi29)
                _ = fi29
            }
            _ = v28
        }
    }
    v1 -= v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
