package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64; F1 int16 }
func genFn0(p0 string, p1 int32) int64 {
    var v4 int32 = -2147483648
    for fi5 := 0; fi5 < 4; fi5++ {
        v6 := []int{fi5, 180, (fi5 ^ fi5), (fi5 - fi5)}
        _ = fi5
        _ = v6
    }
    fmt.Println(p0, p0, p0)
    for fi7 := 0; fi7 < 14; fi7++ {
        v4 = v4
        v8 := []int{(fi7 | fi7), (fi7 & 0)}
        v8 = append(v8, fi7)
        _ = fi7
        _ = v8
    }
    _ = p0
    _ = p1
    _ = v4
    return 42
}

func main() {
    var v1 uint64 = 4294967295
    v2 := make(map[int]uint32)
    var v3 uint32 = 42
    if true {
        v9 := []int{32767, 1, -576, 1}
        _ = v9
    } else {
        v3 += (^(v3))
    }
    if true {
        if ((v1 % ((1) | 1)) != 0) {
            var v10 uint16 = 255
            _ = v10
        } else {
            if (255 != 255) {
                switch -511 {
                case 42:
                    for fi11 := 0; fi11 < 17; fi11++ {
                        switch "hello" {
                        case "a":
                            v12 := GenS1{F0: 42}
                            _ = v12
                        default:
                            switch (fi11 / ((9223372036854775807) | 1)) {
                            case -9223372036854775808:
                                for fi13 := 0; fi13 < 16; fi13++ {
                                    v3 = v3
                                    v14 := GenS2{F0: 255}
                                    var v15 uint8 = 249
                                    _ = fi13
                                    _ = v14
                                    _ = v15
                                }
                            default:
                                v1 = (^(v1))
                            }
                        }
                        for fi16 := 0; fi16 < 19; fi16++ {
                            if false {
                                v1 -= (v1 & v1)
                                v1 = (v1 & v1)
                            }
                            switch (^(fi11)) {
                            case 1:
                                v17 := []int{}
                                _ = v17
                            default:
                                var v18 uint32 = (v3 % (((v3 + v3)) | 1))
                                _ = v18
                            }
                            _ = fi16
                        }
                        v19 := GenS3{F0: 9223372036854775807, F1: 0}
                        _ = fi11
                        _ = v19
                    }
                case 0:
                    fmt.Println(v3, v1)
                case -2147483648:
                    v20 := []int{514}
                    _ = v20
                default:
                    var v21 uint8 = 127
                    _ = v21
                }
                v22 := []int{}
                fmt.Println(v1, v3)
                _ = v22
            }
        }
        var v23 float32 = (0.1 - 1e10)
        var v24 int16 = 42
        _ = v23
        _ = v24
    }
    if (0.1 >= 3.14159) {
        v25 := GenS3{F0: -9223372036854775808, F1: 0}
        var v26 uint32 = (v3 / ((617) | 1))
        v27 := []int{894, 4294967295, 9223372036854775807, 954}
        _ = v25
        _ = v26
        _ = v27
    } else {
        switch ("\n\t" + "a") {
        case "foo bar":
            v1 -= (v1 | (v1 * v1))
        default:
            delete(v2, 2147483647)
        }
    }
    if ((v3 | 2) >= (v3 % (((v3 | v3)) | 1))) {
        for fi28 := 0; fi28 < 3; fi28++ {
            v1 = 18446744073709551615
            var v29 uint16 = 65535
            _ = fi28
            _ = v29
        }
    }
    var v30 uint = 0
    var v31 bool = ((v1 % ((42) | 1)) <= v1)
    v32 := GenS2{F0: 32767}
    var v33 bool = (-1 != 0)
    var v34 int16 = 0
    fmt.Println(v3, v33, v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
}
