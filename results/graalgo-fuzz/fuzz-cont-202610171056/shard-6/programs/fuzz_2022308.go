package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint64; F1 int32; F2 int64 }
func genFn0(p0 uint64) uint32 {
    if (false && (354 < 0)) {
        v5 := []int{-894, 2147483647}
        for fi6 := 0; fi6 < 2; fi6++ {
            var v7 int8 = 1
            p0 += p0
            _ = fi6
            _ = v7
        }
        _ = v5
    } else {
        for fi8 := 0; fi8 < 15; fi8++ {
            p0 += p0
            p0 = 42
            fmt.Println(p0, p0)
            _ = fi8
        }
        var v9 bool = (255 >= 128)
        fmt.Println(p0, p0)
        _ = v9
    }
    v10 := GenS1{F0: 2}
    _ = p0
    _ = v10
    return 2147483647
}
func genFn1() uint8 {
    v11 := []int{42, 42}
    for fi12 := 0; fi12 < 13; fi12++ {
        v13 := GenS2{F0: -2147483648}
        _ = fi12
        _ = v13
    }
    var v14 uint8 = 14
    v14 -= v14
    v14 += v14
    _ = v11
    _ = v14
    return (v14 / (((v14 ^ (v14 / ((v14) | 1)))) | 1))
}

func main() {
    var v1 uint64 = 1
    var v2 uint32 = 303
    v3 := make(map[int]int)
    var v4 float32 = 0.1
    v15 := GenS2{F0: -144}
    for fi16 := 0; fi16 < 4; fi16++ {
        switch "" {
        case "a":
            delete(v3, (fi16 & fi16))
        case "foo bar":
            if (-665 > -392) {
                var v17 uint8 = 1
                v18 := GenS2{F0: 0}
                var v19 uint64 = uint64(v17)
                _ = v17
                _ = v18
                _ = v19
            }
        case "hello":
            var v20 int16 = 296
            _ = v20
        }
        if (15 == 255) {
            var v21 float64 = 1.5
            if ((!false) && true) {
                if true {
                    var v22 bool = true
                    _ = v22
                } else {
                    v4 = v4
                    v23 := GenS2{F0: 42}
                    var v24 int = (fi16 | (fi16 & (-(fi16))))
                    _ = v23
                    _ = v24
                }
                _ = len(v3)
            } else {
                v3[(fi16 % (((fi16 % ((fi16) | 1))) | 1))] = fi16
                var v25 uint16 = 0
                switch fi16 {
                case -503:
                    v1 = v1
                default:
                    var v26 int32 = 0
                    _ = v26
                }
                _ = v25
            }
            switch (fi16 + (fi16 / ((fi16) | 1))) {
            case 127:
                v4 -= v4
            case -921:
                if false {
                    var v27 int = 2147483647
                    switch "\n\t" {
                    case "":
                        v1 = v1
                    case "\n\t":
                        for fi28 := 0; fi28 < 15; fi28++ {
                            switch v27 {
                            case -1:
                                if (false && false) {
                                    var v29 int8 = 7
                                    var v30 uint8 = 255
                                    _ = v29
                                    _ = v30
                                } else {
                                    if false {
                                        v1 = (v1 << ((uint(v1)) & 63))
                                    }
                                    v31 := GenS2{F0: -1}
                                    var v32 uint64 = 142
                                    v33 := GenS1{F0: v1}
                                    _ = v31
                                    _ = v32
                                    _ = v33
                                }
                            case 0:
                                var v34 uint32 = (v2 << ((571) & 63))
                                _ = v34
                            }
                            v35 := GenS3{F0: v1, F1: 2147483647, F2: 42}
                            if true {
                                v4 += 1e10
                                fmt.Println(v4)
                            }
                            _ = fi28
                            _ = v35
                        }
                    }
                    v36 := GenS3{F0: v1, F1: 1, F2: -733}
                    _ = v27
                    _ = v36
                }
            default:
                if true {
                    for fi37 := 0; fi37 < 1; fi37++ {
                        var v38 float64 = (v21 / (3.14159 * v21))
                        fmt.Println(v2, v4)
                        v38 += 3.14159
                        _ = fi37
                        _ = v38
                    }
                }
            }
            _ = v21
        }
        var v39 string = "\n\t"
        _ = fi16
        _ = v39
    }
    v4 -= float32((0.1 + (3.14159 - float64(1.5))))
    var v40 int16 = 866
    if true {
        var v41 int32 = 1
        v40 = (v40 / ((v40) | 1))
        _ = v41
    }
    var v42 uint16 = 1
    delete(v3, -1)
    v43 := GenS3{F0: 18446744073709551615, F1: 281, F2: 933}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v42
    _ = v43
}
