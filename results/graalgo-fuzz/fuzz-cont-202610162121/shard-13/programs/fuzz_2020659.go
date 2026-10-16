package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint16 }
func genFn0(p0 int32) int {
    for fi6 := 0; fi6 < 19; fi6++ {
        switch (fi6 * (fi6 & fi6)) {
        case -1:
            var v7 float32 = 3.14159
            _ = v7
        case -9223372036854775808:
            for fi8 := 0; fi8 < 14; fi8++ {
                var v9 uint = 2
                _ = fi8
                _ = v9
            }
        default:
            for fi10 := 0; fi10 < 4; fi10++ {
                v11 := []int{(fi6 - 2147483647), (fi10 / ((fi10) | 1))}
                if (2 >= 255) {
                    switch (-(fi10)) {
                    case 0:
                        p0 = (p0 >> ((18446744073709551615) & 63))
                    default:
                        fmt.Println(fi10, fi10, fi10)
                    }
                    var v12 int = (fi10 * (fi6 << ((42) & 63)))
                    _ = v12
                } else {
                    v13 := GenS1{F0: 84}
                    var v14 int8 = -128
                    _ = v13
                    _ = v14
                }
                _ = fi10
                _ = v11
            }
        }
        _ = fi6
    }
    switch 128 {
    case 4294967295:
        v15 := GenS2{F0: 650}
        _ = v15
    case 2147483647:
        for fi16 := 0; fi16 < 6; fi16++ {
            var v17 int16 = -549
            fmt.Println(v17, fi16)
            fmt.Println(v17, p0, v17)
            _ = fi16
            _ = v17
        }
    case -1:
        p0 += (p0 * 1)
    }
    v18 := []int{1, -9223372036854775808, -752, 2147483647}
    var v19 int32 = (p0 | p0)
    _ = p0
    _ = v18
    _ = v19
    return 32767
}

func main() {
    var v1 float32 = 1.5
    var v2 uint64 = 42
    var v3 int32 = 2147483647
    var v4 string = "\n\t"
    var v5 uint16 = 447
    var v20 string = v4
    v21 := GenS3{F0: (v5 << ((625) & 63))}
    if false {
        v22 := []int{int(v2)}
        _ = v22
    } else {
        switch (v20 + (v20 + (v4 + v20))) {
        case "\n\t":
            v23 := []int{984}
            _ = v23
        case "":
            for fi24 := 0; fi24 < 2; fi24++ {
                for fi25 := 0; fi25 < 12; fi25++ {
                    if (genFn0(-2147483648) > fi25) {
                        var v26 float32 = v1
                        switch (fi25 / (((fi24 / ((fi24) | 1))) | 1)) {
                        case 32768:
                            var v27 float64 = -2.25
                            _ = v27
                        default:
                            var v28 string = v4
                            _ = v28
                        }
                        _ = v26
                    } else {
                        v3 = (v3 ^ v3)
                        var v29 int64 = 42
                        _ = v29
                    }
                    fmt.Println(fi24, v20)
                    _ = fi25
                }
                var v30 int16 = -1
                _ = fi24
                _ = v30
            }
        case "hello":
            var v31 float64 = 1.5
            _ = v31
        default:
            fmt.Println(v2, v5, v1)
        }
        v32 := []int{}
        _ = v32
    }
    v20 = (v4 + v20)
    for fi33 := 0; fi33 < 8; fi33++ {
        v34 := GenS1{F0: 100}
        _ = fi33
        _ = v34
    }
    v2 += v2
    v35 := []int{-9223372036854775808, genFn0(v3)}
    var v36 uint64 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
