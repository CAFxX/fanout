package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8; F1 uint32; F2 uint8 }
type GenS3 struct { F0 int64 }
func genFn0(p0 int) uint {
    var v7 uint32 = 0
    v7 += v7
    p0 += -71
    var v8 uint8 = 1
    _ = p0
    _ = v7
    _ = v8
    return 1
}
func genFn1(p0 uint8) int {
    fmt.Println(p0)
    v9 := []int{717, 128}
    for fi10 := 0; fi10 < 11; fi10++ {
        var v11 int64 = 0
        _ = fi10
        _ = v11
    }
    _ = p0
    _ = v9
    return 9223372036854775807
}

func main() {
    var v1 uint8 = 0
    var v2 uint = 2
    v3 := make(map[uint32]int)
    var v4 uint64 = 2
    var v5 uint16 = 65535
    v6 := []string{(("a" + "\n\t") + "a"), "\n\t", "foo bar"}
    v5 = (v5 + (v5 - (v5 & (^(v5)))))
    v1 = v1
    if (!(true || true)) {
        var v12 uint32 = 382
        _ = v12
    }
    if len(v6) > 0 {
        v6[0] = ("" + "")
    }
    fmt.Println(v1, v5, v4)
    for fi13 := 0; fi13 < 8; fi13++ {
        for fi14 := 0; fi14 < 6; fi14++ {
            v15 := GenS1{F0: 903}
            _ = fi14
            _ = v15
        }
        fmt.Println(v2, v5)
        _ = fi13
    }
    for fi16 := 0; fi16 < 4; fi16++ {
        delete(v3, 42)
        v1 = (v1 - v1)
        v1 = 65
        _ = fi16
    }
    if false {
        var v17 int = 42
        fmt.Println(v17, v4, v4)
        switch (v17 + v17) {
        case 4294967295:
            var v18 uint32 = 2147483648
            _ = v18
        }
        _ = v17
    }
    switch -1 {
    case -2147483648:
        var v19 uint8 = 128
        _ = v19
    case 9223372036854775807:
        v3[2147483647] = 126
    case -415:
        if ((v4 < (v4 * 308)) || false) {
            var v20 int64 = 9223372036854775807
            _ = v20
        }
    default:
        var v21 uint = 53
        _ = v21
    }
    v22 := v3[2147483648]
    if false {
        var v23 uint64 = (v4 << ((v2) & 63))
        v24 := GenS1{F0: 32767}
        _ = v23
        _ = v24
    } else {
        if true {
            for fi25 := 0; fi25 < 15; fi25++ {
                for fi26 := 0; fi26 < 13; fi26++ {
                    switch "a" {
                    case "a":
                        v27 := GenS1{F0: int16(v2)}
                        _ = v27
                    case "\n\t":
                        v3[2147483648] = fi25
                    case "":
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                    default:
                        if ((v4 - 2) <= (v4 & (v4 >> (((v2 ^ (v2 + 4294967295))) & 63)))) {
                            var v28 int16 = 32767
                            _ = v28
                        }
                    }
                    v22 += (v22 | (fi26 * (fi25 >> ((uint(v1)) & 63))))
                    v2 = v2
                    _ = fi26
                }
                v29 := GenS1{F0: -32768}
                _ = fi25
                _ = v29
            }
            for fi30 := 0; fi30 < 4; fi30++ {
                v31 := GenS2{F0: -128, F1: 65535, F2: v1}
                _ = fi30
                _ = v31
            }
        } else {
            v32 := GenS3{F0: -9223372036854775808}
            if (((^(v2)) > v2) || false) {
                v33 := v3[1]
                var v34 int16 = 42
                for fi35 := 0; fi35 < 17; fi35++ {
                    var v36 bool = (v4 <= (v4 | (v4 & (^(v4)))))
                    _ = fi35
                    _ = v36
                }
                _ = v33
                _ = v34
            }
            _ = v32
        }
    }
    var v37 bool = (false && true)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
