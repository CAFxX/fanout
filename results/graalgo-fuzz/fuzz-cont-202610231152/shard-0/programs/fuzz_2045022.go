package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 uint32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint32; F1 int64 }
type GenS4 struct { F0 int64; F1 uint16; F2 int64 }
type GenS5 struct { F0 uint32 }

func main() {
    var v1 int16 = 1000
    v2 := make(map[uint32]int)
    var v3 int = 42
    v4 := GenS1{F0: 42, F1: 65535, F2: 2147483648}
    if (!(127 < 127)) {
        for fi5 := 0; fi5 < 12; fi5++ {
            v3 += (fi5 << ((0) & 63))
            fmt.Println(fi5, v1, fi5)
            var v6 uint8 = 10
            _ = fi5
            _ = v6
        }
        var v7 string = "foo bar"
        _ = v7
    } else {
        fmt.Println(v3, v3)
        if (9223372036854775807 >= 1) {
            var v8 uint32 = 1
            var v9 uint16 = 32767
            _ = v8
            _ = v9
        }
        v3 += v3
    }
    v1 = (v1 + (v1 - (v1 & int16(v3))))
    for fi10 := 0; fi10 < 9; fi10++ {
        if false {
            for fi11 := 0; fi11 < 8; fi11++ {
                v12 := GenS1{F0: 42, F1: 0, F2: 2147483648}
                var v13 uint64 = 18446744073709551615
                v13 = 577
                _ = fi11
                _ = v12
                _ = v13
            }
        }
        if (-13 > -1) {
            var v14 uint64 = 0
            _ = v14
        } else {
            var v15 int8 = -115
            switch -150 {
            case 0:
                for fi16 := 0; fi16 < 18; fi16++ {
                    v2[986] = fi10
                    v17 := GenS1{F0: 2, F1: 255, F2: 570}
                    switch fi16 {
                    case -1:
                        v18 := GenS1{F0: 2147483647, F1: 1, F2: 65535}
                        _ = v18
                    case 1:
                        v19 := GenS2{F0: 539}
                        _ = v19
                    case 2147483647:
                        var v20 float32 = 0.0
                        _ = v20
                    default:
                        v21 := GenS3{F0: 2, F1: -669}
                        _ = v21
                    }
                    _ = fi16
                    _ = v17
                }
            }
            _ = v15
        }
        for fi22 := 0; fi22 < 1; fi22++ {
            v3 = (fi22 + (fi10 * (fi22 << ((18446744073709551615) & 63))))
            v23 := []int{fi22, fi10, (fi10 % (((fi10 + (v3 % (((fi10 & v3)) | 1)))) | 1))}
            _ = fi22
            _ = v23
        }
        _ = fi10
    }
    for fi24 := 0; fi24 < 9; fi24++ {
        var v25 int16 = v1
        if false {
            fmt.Println(v3, v25, v1)
            var v26 uint32 = 2
            if true {
                for fi27 := 0; fi27 < 2; fi27++ {
                    switch ("foo bar" + "a") {
                    case "a":
                        if true {
                            if false {
                                delete(v2, (v26 % ((v26) | 1)))
                                var v28 int64 = 546
                                var v29 float64 = 0.0
                                _ = v28
                                _ = v29
                            }
                        } else {
                            v2[(^(v26))] = (fi27 & (fi24 + fi27))
                            v25 = (v25 + 1)
                        }
                        v1 -= (v1 >> ((0) & 63))
                    }
                    v25 += (v25 | -32768)
                    if (57 != 1) {
                        switch 251 {
                        case 42:
                            switch "\n\t" {
                            case "":
                                v3 = (v3 / ((fi24) | 1))
                            }
                        }
                    }
                    _ = fi27
                }
            } else {
                v3 += (v3 + (v3 - 1))
                if true {
                    v25 += v1
                    for fi30 := 0; fi30 < 8; fi30++ {
                        v26 -= v26
                        _ = fi30
                    }
                    var v31 float64 = (((((0.1 - 1e10) * (float64(1e10) - 1e10)) + 0.1) + -2.25) - 0.1)
                    _ = v31
                }
                for fi32 := 0; fi32 < 14; fi32++ {
                    v33 := GenS4{F0: -636, F1: 255, F2: int64(v1)}
                    v26 += v26
                    _ = fi32
                    _ = v33
                }
            }
            _ = v26
        }
        _ = fi24
        _ = v25
    }
    var v34 uint8 = 128
    v35 := []int{(^(v3))}
    v3 += v3
    switch (-(v3)) {
    case 127:
        var v36 bool = false
        _ = v36
    case -933:
        fmt.Println(v3, v34)
    default:
        var v37 uint64 = 1
        _ = v37
    }
    fmt.Println(v1, v3, v3)
    var v38 int32 = 0
    fmt.Println(v34, v1, v1)
    var v39 uint16 = 65535
    v40 := GenS5{F0: 65535}
    var v41 int16 = v1
    _ = len(v2)
    var v42 float64 = ((1.5 + ((3.14159 - float64(float32(-2.25))) / 0.1)) + 0.1)
    v38 += v38
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v42
}
