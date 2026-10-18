package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8 }
type GenS2 struct { F0 string; F1 int8 }
type GenS3 struct { F0 uint64; F1 uint }

func main() {
    var v1 uint = 58
    v2 := make(map[string]uint32)
    var v3 uint64 = 946
    var v4 uint8 = 1
    v5 := make(map[string]bool)
    v6 := make(map[int]string)
    switch 32768 {
    case 9223372036854775807:
        if (2 != 1) {
            var v7 uint32 = 721
            v3 = v3
            if false {
                for fi8 := 0; fi8 < 11; fi8++ {
                    v9 := []int{fi8}
                    _ = fi8
                    _ = v9
                }
                if false {
                    var v10 string = "\n\t"
                    _ = v10
                }
            } else {
                v11 := GenS1{F0: uint(v7), F1: (v4 ^ 1)}
                _ = v11
            }
            _ = v7
        }
    default:
        v1 -= 18446744073709551615
    }
    if (!false) {
        v1 += 40
        for fi12 := 0; fi12 < 5; fi12++ {
            if true {
                var v13 int = -539
                _ = v13
            }
            v4 += 155
            var v14 uint16 = 255
            _ = fi12
            _ = v14
        }
        fmt.Println(v3, v4)
    } else {
        v3 = v3
        var v15 int = 32768
        _ = v15
    }
    v16 := GenS1{F0: (v1 * v1), F1: v4}
    if ((false || (902 <= 0)) || ("hello" == "a")) {
        switch "\n\t" {
        case "hello":
            var v17 int16 = 32767
            _ = v17
        case "a":
            if (false && (!((false && false) && (false || (!true))))) {
                for fi18 := 0; fi18 < 1; fi18++ {
                    v19 := []int{(fi18 | (^(fi18))), fi18, (fi18 << ((v1) & 63)), fi18}
                    if false {
                        v1 = v1
                        if true {
                            for fi20 := 0; fi20 < 5; fi20++ {
                                var v21 uint64 = (v3 >> ((v1) & 63))
                                _ = fi20
                                _ = v21
                            }
                            v3 = v3
                        } else {
                            v3 += v3
                        }
                    }
                    _ = fi18
                    _ = v19
                }
            }
        case "":
            delete(v2, "")
        }
    }
    v1 = (v1 << ((v1) & 63))
    var v22 int64 = 9223372036854775807
    v23 := GenS1{F0: v1, F1: (^(v4))}
    if true {
        v4 += (^(v4))
    } else {
        var v24 bool = (((1e10 - (3.14159 + ((3.14159 * float64(((0.0 + 1e10) - 3.14159))) + float64(-2.25)))) - 1.5) > 0.1)
        v25 := GenS2{F0: "hello", F1: -110}
        v26 := []int{42, -40, 0, 42}
        _ = v24
        _ = v25
        _ = v26
    }
    v27 := GenS1{F0: (v1 / (((v1 | v1)) | 1)), F1: 108}
    v1 = (v1 - (^(v1)))
    var v28 float32 = 3.14159
    fmt.Println(v22)
    var v29 int8 = -21
    v4 = v4
    for fi30 := 0; fi30 < 19; fi30++ {
        v28 = v28
        _ = fi30
    }
    if false {
        v31 := GenS1{F0: v1, F1: (v4 << (((^(v1))) & 63))}
        v28 += v28
        _ = v31
    } else {
        if (42 >= 42) {
            fmt.Println(v28)
        }
        v32 := GenS3{F0: v3, F1: (v1 ^ (v1 ^ 65535))}
        _ = v32
    }
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v2
    _ = v22
    _ = v23
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
