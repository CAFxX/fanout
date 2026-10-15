package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 int16; F1 bool }
func genFn0(p0 uint8, p1 int32) int8 {
    v5 := []int{-298}
    for fi6 := 0; fi6 < 17; fi6++ {
        if true {
            if len(v5) > 0 {
                v5[0] = fi6
            }
            var v7 uint8 = 0
            _ = v7
        }
        for fi8 := 0; fi8 < 5; fi8++ {
            var v9 bool = ("a" >= "\n\t")
            v5 = append(v5, 625)
            _ = fi8
            _ = v9
        }
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = v5
    return 100
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int32 = -2147483648
    var v3 int64 = -1
    var v4 uint64 = 717
    for fi10 := 0; fi10 < 1; fi10++ {
        for fi11 := 0; fi11 < 6; fi11++ {
            switch (fi11 | (fi10 / ((fi10) | 1))) {
            case -9223372036854775808:
                switch (fi10 >> ((uint(v4)) & 63)) {
                case 384:
                    var v12 bool = (877 <= 339)
                    _ = v12
                case 2147483647:
                    if (!false) {
                        var v13 uint16 = 2
                        _ = v13
                    } else {
                        var v14 uint64 = 249
                        _ = v14
                    }
                }
            case 128:
                fmt.Println(fi10)
            case 125:
                var v15 uint = 42
                _ = v15
            }
            v16 := []int{(fi11 << ((18446744073709551615) & 63))}
            _ = fi11
            _ = v16
        }
        var v17 bool = (int32(v3) != int32(v4))
        _ = fi10
        _ = v17
    }
    for fi18 := 0; fi18 < 11; fi18++ {
        var v19 int64 = 1
        _ = fi18
        _ = v19
    }
    v2 += v2
    v3 += (v3 % ((v3) | 1))
    v20 := GenS1{F0: (v4 + (v4 << ((42) & 63))), F1: 555}
    var v21 uint32 = 65535
    v3 = v3
    var v22 int32 = 2147483647
    fmt.Println(v2)
    _ = len(v1)
    v23 := GenS1{F0: v4, F1: 792}
    for fi24 := 0; fi24 < 1; fi24++ {
        v2 = v22
        v1[(fi24 % ((fi24) | 1))] = v21
        _ = fi24
    }
    var v25 uint32 = (v21 >> ((4294967295) & 63))
    var v26 int = -475
    v27 := GenS1{F0: v4, F1: 65535}
    v28 := GenS1{F0: 1, F1: 750}
    for fi29 := 0; fi29 < 6; fi29++ {
        for fi30 := 0; fi30 < 14; fi30++ {
            if ("" == ("\n\t" + "hello")) {
                var v31 uint32 = (^(v25))
                fmt.Println(v26)
                _ = v31
            }
            _ = fi30
        }
        switch 552 {
        case 0:
            var v32 int8 = int8(fi29)
            _ = v32
        case 32767:
            v33 := []int{}
            _ = v33
        case 127:
            if (!false) {
                fmt.Println(v21, v26)
            } else {
                _ = len(v1)
                switch "hello" {
                case "hello":
                    v2 += (v2 << ((65535) & 63))
                case "":
                    var v34 int8 = genFn0(160, v2)
                    _ = v34
                default:
                    switch v26 {
                    case -2147483648:
                        var v35 int32 = v2
                        _ = v35
                    default:
                        for fi36 := 0; fi36 < 3; fi36++ {
                            _ = len(v1)
                            _ = fi36
                        }
                    }
                }
                if true {
                    switch "\n\t" {
                    case "":
                        if true {
                            delete(v1, fi29)
                            var v37 int32 = -422
                            _ = v37
                        } else {
                            v38 := GenS2{F0: -718, F1: true}
                            v4 = (v4 % (((v4 - v4)) | 1))
                            v22 = -270
                            _ = v38
                        }
                    case "hello":
                        switch "a" {
                        case "":
                            var v39 int32 = (v22 * v22)
                            _ = v39
                        case "\n\t":
                            v40 := []int{int(v3), (v26 ^ fi29), 32767}
                            _ = v40
                        default:
                            fmt.Println(v22)
                        }
                    default:
                        v41 := []int{}
                        _ = v41
                    }
                    v2 -= -1
                } else {
                    v42 := GenS2{F0: -32768, F1: false}
                    _ = v42
                }
            }
        default:
            fmt.Println(v26, v21)
        }
        v2 = v22
        _ = fi29
    }
    v4 += (^(v4))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
