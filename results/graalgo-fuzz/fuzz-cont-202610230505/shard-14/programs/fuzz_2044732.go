package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 bool }

func main() {
    var v1 int = 32768
    var v2 uint = 4294967295
    var v3 int = (v1 | (v1 ^ v1))
    var v4 uint32 = 564
    var v5 int64 = 59
    v4 -= uint32(v2)
    for fi6 := 0; fi6 < 6; fi6++ {
        for fi7 := 0; fi7 < 1; fi7++ {
            var v8 int64 = 9223372036854775807
            _ = fi7
            _ = v8
        }
        var v9 uint8 = 201
        _ = fi6
        _ = v9
    }
    if false {
        for fi10 := 0; fi10 < 15; fi10++ {
            for fi11 := 0; fi11 < 17; fi11++ {
                var v12 int16 = 32767
                v2 = 65535
                var v13 int64 = v5
                _ = fi11
                _ = v12
                _ = v13
            }
            if ((v2 << (((v2 % ((v2) | 1))) & 63)) >= (v2 / (((v2 % ((uint(v4)) | 1))) | 1))) {
                v14 := []int{(v3 - v3)}
                for fi15 := 0; fi15 < 4; fi15++ {
                    var v16 int = (fi15 / ((fi10) | 1))
                    _ = fi15
                    _ = v16
                }
                var v17 uint = (v2 ^ (v2 * (v2 % ((v2) | 1))))
                _ = v14
                _ = v17
            } else {
                v18 := []int{(^(fi10))}
                fmt.Println(v1, v1)
                var v19 int64 = (^(v5))
                _ = v18
                _ = v19
            }
            _ = fi10
        }
        for fi20 := 0; fi20 < 14; fi20++ {
            var v21 int = (v1 | (-(v3)))
            v2 += uint(v4)
            for fi22 := 0; fi22 < 13; fi22++ {
                var v23 uint16 = 42
                var v24 int16 = -715
                v5 += int64(v1)
                v25 := GenS1{F0: v5}
                _ = fi22
                _ = v23
                _ = v24
                _ = v25
            }
            _ = fi20
            _ = v21
        }
        for fi26 := 0; fi26 < 10; fi26++ {
            for fi27 := 0; fi27 < 11; fi27++ {
                fmt.Println(v2)
                if true {
                    var v28 int32 = -807
                    _ = v28
                } else {
                    v29 := []int{v3}
                    _ = v29
                }
                _ = fi27
            }
            v30 := GenS2{F0: 220}
            _ = fi26
            _ = v30
        }
    } else {
        v31 := []int{(-(v1)), v3, (v1 << ((v2) & 63)), (v3 >> (((v2 ^ v2)) & 63))}
        v32 := GenS1{F0: (v5 << ((18446744073709551615) & 63))}
        v33 := GenS1{F0: (v5 - v5)}
        _ = v31
        _ = v32
        _ = v33
    }
    v34 := []int{0}
    switch (v3 / ((v3) | 1)) {
    case -2147483648:
        for fi35 := 0; fi35 < 17; fi35++ {
            var v36 uint8 = 0
            if true {
                v37 := GenS3{F0: 42}
                var v38 int64 = -430
                v1 += fi35
                _ = v37
                _ = v38
            } else {
                _ = len(v34)
            }
            _ = fi35
            _ = v36
        }
    case 1:
        if true {
            switch "" {
            case "hello":
                v1 = int(v4)
            case "\n\t":
                fmt.Println(v2, v3, v3)
            default:
                var v39 bool = false
                _ = v39
            }
        }
    default:
        if false {
            _ = len(v34)
            v4 = (v4 * 65535)
        }
    }
    if len(v34) > 0 {
        v34[0] = v3
    }
    v34 = append(v34, v1)
    if true {
        var v40 uint16 = 476
        var v41 int64 = 1
        _ = v40
        _ = v41
    } else {
        v5 -= (v5 * 42)
    }
    v42 := GenS4{F0: false}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v42
    _ = v5
}
