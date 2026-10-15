package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 int8 }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 int; F1 int }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint; F1 uint32; F2 uint64 }
func genFn0(p0 uint32) int8 {
    var v5 int32 = 798
    var v6 bool = (-32768 <= -1)
    if true {
        fmt.Println(v5)
        switch 988 {
        case 0:
            v7 := GenS1{F0: -628, F1: v5, F2: 42}
            _ = v7
        case 42:
            var v8 int16 = -911
            _ = v8
        default:
            var v9 uint32 = 1
            _ = v9
        }
        for fi10 := 0; fi10 < 8; fi10++ {
            for fi11 := 0; fi11 < 15; fi11++ {
                v12 := GenS1{F0: v5, F1: v5, F2: 121}
                fmt.Println(fi10)
                _ = fi11
                _ = v12
            }
            v13 := []int{(fi10 << ((18446744073709551615) & 63)), fi10, -9223372036854775808}
            fmt.Println(p0)
            _ = fi10
            _ = v13
        }
    }
    for fi14 := 0; fi14 < 15; fi14++ {
        v6 = v6
        _ = fi14
    }
    if v6 {
        for fi15 := 0; fi15 < 16; fi15++ {
            fmt.Println(v5, p0, fi15)
            v16 := []int{(fi15 | (fi15 - (fi15 % ((fi15) | 1)))), (^(fi15)), (fi15 - fi15), (fi15 - -549)}
            p0 = p0
            _ = fi15
            _ = v16
        }
        if v6 {
            v17 := GenS1{F0: -457, F1: (v5 + (v5 & 42)), F2: 42}
            if (!false) {
                if v6 {
                    var v18 uint64 = 42
                    _ = v18
                }
                p0 = p0
            }
            _ = v17
        }
    }
    _ = p0
    _ = v5
    _ = v6
    return -128
}
func genFn1(p0 uint32, p1 int64) int8 {
    if ((true || (!true)) && true) {
        p0 = 2147483647
        var v19 int8 = 42
        _ = v19
    }
    var v20 int32 = 42
    v21 := []int{-1, 1}
    v20 -= (v20 ^ 0)
    var v22 float32 = 1.5
    _ = p0
    _ = p1
    _ = v20
    _ = v21
    _ = v22
    return 1
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]uint32)
    var v3 float32 = -2.25
    v4 := make(map[string]string)
    var v23 uint8 = 166
    var v24 uint8 = v23
    v3 = v3
    if false {
        switch "foo bar" {
        case "":
            var v25 uint = 813
            _ = v25
        }
        v26 := GenS2{F0: "", F1: 347}
        _ = v26
    } else {
        var v27 uint32 = 949
        _ = v27
    }
    var v28 int = 1
    v29 := []int{(v28 % ((42) | 1))}
    v23 -= (v24 - (v24 % ((2) | 1)))
    switch ("" + ("\n\t" + "\n\t")) {
    case "\n\t":
        switch 9223372036854775807 {
        case 203:
            if true {
                v28 = v28
                var v30 uint32 = 671
                _ = v30
            } else {
                if false {
                    v28 += (v28 | (^(v28)))
                    delete(v1, 1)
                    if len(v29) >= 2 {
                        v29 = v29[1:len(v29)]
                    }
                }
                for fi31 := 0; fi31 < 14; fi31++ {
                    v23 += (v23 << ((230) & 63))
                    v32 := GenS3{F0: v28, F1: (-(fi31))}
                    for fi33 := 0; fi33 < 10; fi33++ {
                        fmt.Println(fi33, fi31)
                        v34 := GenS4{F0: 0}
                        v3 = -2.25
                        _ = fi33
                        _ = v34
                    }
                    _ = fi31
                    _ = v32
                }
                var v35 uint16 = 32767
                _ = v35
            }
        default:
            switch "" {
            case "hello":
                if len(v29) >= 2 {
                    v29 = v29[1:len(v29)]
                }
            case "\n\t":
                v29 = append(v29, 1)
            }
        }
    case "hello":
        for fi36 := 0; fi36 < 11; fi36++ {
            fmt.Println(v3, v23)
            v23 = 1
            _ = fi36
        }
    }
    v37 := GenS5{F0: 712, F1: 4294967295, F2: 404}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v28
    _ = v29
    _ = v3
    _ = v37
    _ = v4
}
