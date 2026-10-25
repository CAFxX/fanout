package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int; F1 uint8 }

func main() {
    v1 := []int{279, -347, 418}
    v2 := []string{"foo bar", "\n\t", "foo bar"}
    v3 := []uint32{2147483647, 0, 1, 1}
    var v4 uint32 = 92
    var v5 uint8 = 124
    v6 := GenS1{F0: 1}
    var v7 int32 = 0
    _ = len(v1)
    var v8 uint64 = 0
    if false {
        var v9 float32 = 0.1
        for fi10 := 0; fi10 < 7; fi10++ {
            var v11 int64 = -889
            for fi12 := 0; fi12 < 20; fi12++ {
                if (!(42 >= uint16(v8))) {
                    v3 = append(v3, 42)
                    if (-49 < v7) {
                        var v13 uint8 = 1
                        _ = v13
                    }
                }
                v7 -= v7
                v14 := GenS1{F0: -9223372036854775808}
                fmt.Println(v4, fi10)
                var v15 uint16 = 972
                _ = fi12
                _ = v14
                _ = v15
            }
            _ = fi10
            _ = v11
        }
        if len(v1) > 0 {
            v1[0] = -9223372036854775808
        }
        _ = v9
    } else {
        for fi16 := 0; fi16 < 5; fi16++ {
            for fi17 := 0; fi17 < 6; fi17++ {
                v4 += 2147483648
                var v18 int32 = (v7 ^ -592)
                _ = fi17
                _ = v18
            }
            v4 += v4
            _ = fi16
        }
        switch -9223372036854775808 {
        case -552:
            if (v5 < v5) {
                var v19 uint32 = uint32(v5)
                var v20 int32 = v7
                _ = v19
                _ = v20
            }
        case 128:
            if false {
                v8 -= 73
            } else {
                if true {
                    _ = len(v1)
                    var v21 int = 0
                    _ = v21
                } else {
                    v5 = 79
                }
                var v22 bool = true
                if (v8 <= (v8 ^ (v8 | 2))) {
                    v7 = v7
                }
                _ = v22
            }
        case 581:
            v23 := GenS2{F0: 992, F1: v5}
            _ = v23
        }
    }
    v24 := GenS2{F0: int(v5), F1: 0}
    var v25 uint32 = (v4 << ((2) & 63))
    var v26 int = 832
    v8 -= (v8 % (((v8 * 0)) | 1))
    var v27 float64 = (((float64(3.14159) / 1e10) - (-2.25 / 1e10)) / 1.5)
    for fi28 := 0; fi28 < 15; fi28++ {
        v29 := GenS2{F0: (fi28 / (((fi28 - (fi28 << ((uint(v8)) & 63)))) | 1)), F1: v5}
        switch ("" + "") {
        case "foo bar":
            v30 := GenS2{F0: (fi28 % ((32768) | 1)), F1: (v5 * v5)}
            _ = v30
        case "":
            v31 := GenS1{F0: int(v5)}
            _ = v31
        case "a":
            var v32 uint = uint(v4)
            _ = v32
        }
        _ = fi28
        _ = v29
    }
    v8 = 337
    var v33 int64 = -1
    v27 = ((3.14159 / (v27 * v27)) + ((1.5 / ((1e10 + ((v27 * v27) / 1.0)) * 1.5)) / (1e10 / ((v27 * 0.0) - v27))))
    v7 = v7
    var v34 uint16 = 255
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
