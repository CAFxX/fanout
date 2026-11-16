package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint8 }
type GenS2 struct { F0 int32; F1 uint16 }
type GenS3 struct { F0 int; F1 int; F2 int16 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int64; F1 uint32; F2 uint16 }

func main() {
    v1 := make(map[uint32]int)
    var v2 int8 = 100
    var v3 uint16 = 819
    v4 := GenS1{F0: 1, F1: 128}
    if false {
        v5 := GenS1{F0: -591, F1: 2}
        var v6 bool = false
        var v7 uint8 = 255
        _ = v5
        _ = v6
        _ = v7
    } else {
        fmt.Println(v3, v3)
        var v8 uint64 = 18446744073709551615
        _ = v8
    }
    switch 128 {
    case -2147483648:
        v9 := []int{2147483647, 9223372036854775807, 32767}
        _ = v9
    }
    v10 := v1[42]
    v3 = v3
    if false {
        for fi11 := 0; fi11 < 5; fi11++ {
            v10 = (v10 >> ((0) & 63))
            for fi12 := 0; fi12 < 13; fi12++ {
                v13 := GenS1{F0: -1, F1: 0}
                var v14 uint64 = 2
                v15 := GenS1{F0: 42, F1: 255}
                _ = fi12
                _ = v13
                _ = v14
                _ = v15
            }
            var v16 uint16 = (v3 % ((842) | 1))
            _ = fi11
            _ = v16
        }
    } else {
        v17 := []int{(v10 & v10)}
        fmt.Println(v2)
        _ = v17
    }
    v2 -= 42
    v18 := GenS1{F0: 0, F1: 128}
    var v19 bool = true
    if true {
        var v20 uint16 = (v3 % ((65535) | 1))
        var v21 uint16 = (v3 >> ((42) & 63))
        _ = v20
        _ = v21
    } else {
        fmt.Println(v3, v3)
        var v22 int8 = v2
        for fi23 := 0; fi23 < 19; fi23++ {
            for fi24 := 0; fi24 < 10; fi24++ {
                v25 := GenS2{F0: -2147483648, F1: v3}
                v26 := []int{(fi24 >> ((uint(v3)) & 63)), fi23, 32767}
                if (1 < int32(v2)) {
                    v3 += (v3 ^ (v3 + (v3 - (v3 * (v3 | v3)))))
                    switch fi23 {
                    case -763:
                        v27 := GenS3{F0: 420, F1: fi24, F2: -1}
                        _ = v27
                    case -472:
                        var v28 bool = false
                        v2 = v2
                        _ = v28
                    case -58:
                        if v19 {
                            var v29 uint32 = 65535
                            _ = v29
                        } else {
                            if false {
                                var v30 float64 = (0.1 * (-2.25 + 1e10))
                                v31 := GenS2{F0: 2147483647, F1: (v3 << ((800) & 63))}
                                _ = v30
                                _ = v31
                            } else {
                                var v32 int16 = 42
                                var v33 uint32 = 2
                                _ = v32
                                _ = v33
                            }
                            var v34 string = "\n\t"
                            v19 = ((42 <= 1) || v19)
                            if len(v26) >= 2 {
                                v26 = v26[1:len(v26)]
                            }
                            var v35 int32 = 0
                            _ = v34
                            _ = v35
                        }
                    default:
                        v36 := GenS2{F0: -2147483648, F1: (v3 - (v3 % (((v3 - 0)) | 1)))}
                        _ = v36
                    }
                    fmt.Println(v3, v10)
                }
                _ = fi24
                _ = v25
                _ = v26
            }
            for fi37 := 0; fi37 < 1; fi37++ {
                var v38 bool = ("a" <= "foo bar")
                v19 = v38
                _ = fi37
                _ = v38
            }
            _ = fi23
        }
        _ = v22
    }
    var v39 int32 = -2147483648
    var v40 int8 = 80
    v41 := GenS2{F0: 42, F1: v3}
    v42 := GenS4{F0: -55}
    for fi43 := 0; fi43 < 19; fi43++ {
        var v44 int16 = -1
        v45 := GenS5{F0: 42, F1: 2147483648, F2: (v3 ^ (v3 << ((uint(v3)) & 63)))}
        _ = fi43
        _ = v44
        _ = v45
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v10
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
}
