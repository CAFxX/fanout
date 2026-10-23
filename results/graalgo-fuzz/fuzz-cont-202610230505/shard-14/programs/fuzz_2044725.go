package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool; F1 uint }
type GenS4 struct { F0 int64; F1 int32 }

func main() {
    var v1 int8 = -128
    v2 := []uint32{}
    v3 := make(map[int]int)
    v4 := []int{}
    if false {
        v5 := GenS1{F0: (-(v1))}
        for fi6 := 0; fi6 < 4; fi6++ {
            var v7 float32 = (float32(1e10) * 0.0)
            _ = fi6
            _ = v7
        }
        _ = v5
    } else {
        if true {
            switch ("" + ("" + "a")) {
            case "\n\t":
                fmt.Println(v1)
            case "hello":
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            default:
                var v8 uint32 = 65535
                _ = v8
            }
        } else {
            var v9 int16 = 1000
            for fi10 := 0; fi10 < 17; fi10++ {
                if true {
                    v3[(fi10 % ((fi10) | 1))] = (fi10 % ((fi10) | 1))
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                } else {
                    v11 := GenS1{F0: (^(v1))}
                    switch 2147483647 {
                    case 9223372036854775807:
                        _ = len(v4)
                    case -2147483648:
                        fmt.Println(fi10)
                    case 904:
                        for fi12 := 0; fi12 < 7; fi12++ {
                            var v13 uint16 = 255
                            _ = fi12
                            _ = v13
                        }
                    }
                    switch (^(fi10)) {
                    case 155:
                        fmt.Println(v9, fi10, v1)
                    }
                    _ = v11
                }
                _ = fi10
            }
            v3[2147483647] = 42
            _ = v9
        }
        var v14 uint16 = 173
        _ = v14
    }
    v1 = v1
    for fi15 := 0; fi15 < 19; fi15++ {
        if true {
            for fi16 := 0; fi16 < 6; fi16++ {
                v1 -= (v1 & (v1 * v1))
                _ = fi16
            }
            v1 = v1
            if true {
                var v17 int16 = 741
                for fi18 := 0; fi18 < 1; fi18++ {
                    var v19 string = "\n\t"
                    _ = fi18
                    _ = v19
                }
                _ = v17
            } else {
                if (fi15 <= fi15) {
                    var v20 float32 = -2.25
                    v21 := GenS1{F0: (v1 | v1)}
                    _ = v20
                    _ = v21
                }
                var v22 uint16 = 32767
                var v23 int = fi15
                _ = v22
                _ = v23
            }
        }
        v1 = (^(v1))
        _ = fi15
    }
    var v24 int32 = 1
    switch -314 {
    case 4294967295:
        var v25 uint64 = 794
        _ = v25
    }
    v24 += (v24 - -483)
    for fi26 := 0; fi26 < 14; fi26++ {
        v27 := GenS2{F0: 0}
        v24 -= -1
        _ = fi26
        _ = v27
    }
    var v28 int = -44
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v29 := GenS1{F0: (v1 >> ((475) & 63))}
    v30 := GenS3{F0: false, F1: 0}
    v3[(v28 | 32768)] = (v28 % (((v28 * -9223372036854775808)) | 1))
    var v31 int16 = int16(v24)
    if false {
        v32 := GenS4{F0: -512, F1: (v24 / ((2147483647) | 1))}
        _ = v32
    }
    for fi33 := 0; fi33 < 7; fi33++ {
        v34 := GenS3{F0: (873 < 0), F1: 4294967295}
        var v35 uint16 = 2
        _ = fi33
        _ = v34
        _ = v35
    }
    v28 -= -910
    v31 -= (v31 << ((1) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
}
