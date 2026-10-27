package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint8; F1 uint8; F2 int8 }
func genFn0() uint8 {
    var v7 int8 = 51
    if true {
        fmt.Println(v7)
        v7 += (v7 / ((v7) | 1))
    } else {
        var v8 bool = (2147483648 == 4294967295)
        _ = v8
    }
    if false {
        v9 := GenS1{F0: 2, F1: -509}
        if (-19 < 0) {
            v10 := GenS1{F0: 32767, F1: 525}
            var v11 int32 = -1
            _ = v10
            _ = v11
        }
        _ = v9
    }
    v12 := []int{-2147483648}
    _ = v12
    _ = v7
    return 2
}

func main() {
    var v1 int64 = 0
    v2 := make(map[int]uint32)
    var v3 uint8 = 46
    var v4 string = "foo bar"
    v5 := make(map[uint32]int)
    var v6 uint = 458
    if false {
        v1 = v1
    }
    v4 = v4
    var v13 int16 = 1
    if ((v13 | 371) >= v13) {
        v14 := GenS2{F0: false}
        if (-1 != 42) {
            switch "a" {
            case "hello":
                v15 := []int{32768, 128, -9223372036854775808, 0}
                _ = v15
            case "\n\t":
                delete(v2, 42)
            default:
                v16 := []int{265, -474, -632}
                _ = v16
            }
            for fi17 := 0; fi17 < 19; fi17++ {
                if (-347 >= 42) {
                    if true {
                        if (2147483647 == 42) {
                            var v18 int16 = 1000
                            var v19 int16 = (v13 - (v13 << ((v6) & 63)))
                            _ = v18
                            _ = v19
                        }
                    } else {
                        v2[(fi17 * (fi17 | (fi17 >> (((v6 >> (((v6 / ((uint(v3)) | 1))) & 63))) & 63))))] = 717
                        var v20 uint8 = (v3 << (((v6 & 42)) & 63))
                        switch v4 {
                        case "hello":
                            v3 = 42
                        case "\n\t":
                            if ((v13 < (^(v13))) || ((v13 % (((v13 << ((0) & 63))) | 1)) != v13)) {
                                v3 = (v20 % (((v20 / ((genFn0()) | 1))) | 1))
                                var v21 int32 = -2147483648
                                if ((v1 + v1) >= 130) {
                                    var v22 float64 = 1e10
                                    var v23 float64 = 3.14159
                                    _ = v22
                                    _ = v23
                                }
                                _ = v21
                            } else {
                                var v24 uint16 = 65535
                                for fi25 := 0; fi25 < 1; fi25++ {
                                    v26 := GenS2{F0: (1 == 338)}
                                    var v27 uint32 = 42
                                    _ = fi25
                                    _ = v26
                                    _ = v27
                                }
                                _ = v24
                            }
                        case "":
                            if false {
                                _ = len(v2)
                                for fi28 := 0; fi28 < 6; fi28++ {
                                    v29 := []int{fi17}
                                    var v30 uint = (v6 | (v6 >> ((65535) & 63)))
                                    _ = fi28
                                    _ = v29
                                    _ = v30
                                }
                            } else {
                                v4 = "a"
                                fmt.Println(v4, v4)
                            }
                        }
                        _ = v20
                    }
                    if (true || false) {
                        v4 = "\n\t"
                        v13 = v13
                    }
                }
                _ = fi17
            }
            var v31 int32 = -1
            _ = v31
        }
        _ = v14
    } else {
        var v32 int = 814
        _ = len(v5)
        _ = v32
    }
    for fi33 := 0; fi33 < 6; fi33++ {
        v34 := GenS3{F0: (((v4 + "a") + v4) + "")}
        v1 = -9223372036854775808
        v35 := GenS1{F0: 288, F1: 2147483647}
        _ = fi33
        _ = v34
        _ = v35
    }
    v13 = (v13 / ((-1) | 1))
    fmt.Println(v4, v3)
    var v36 int16 = v13
    v1 = v1
    delete(v2, 561)
    switch 1 {
    case 306:
        var v37 uint64 = 160
        _ = v37
    case 4294967295:
        v38 := GenS4{F0: 82}
        _ = v38
    default:
        v6 += 625
    }
    if ((v13 / (((v13 ^ 32767)) | 1)) > v13) {
        v4 = v4
    } else {
        if ((((((1e10 - float64(float32(1.5))) / 1.0) + -2.25) / 1.0) > (1.5 - (1.5 + ((float64(0.1) / -2.25) - 0.0)))) && false) {
            if true {
                v39 := []int{2147483647, -490, -1, -9223372036854775808}
                var v40 float64 = (0.1 + (0.0 / -2.25))
                v2[int(v3)] = 2147483648
                _ = v39
                _ = v40
            }
        }
    }
    v13 += v36
    v36 -= 1000
    _ = len(v2)
    if (32767 < 889) {
        v41 := GenS5{F0: 255, F1: (v3 % (((v3 / (((^(v3))) | 1))) | 1)), F2: 42}
        v6 = (^(v6))
        _ = v41
    }
    if true {
        var v42 uint32 = 65535
        for fi43 := 0; fi43 < 19; fi43++ {
            fmt.Println(v3, v6)
            var v44 int = (fi43 * (fi43 << ((v6) & 63)))
            _ = fi43
            _ = v44
        }
        v3 -= v3
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
