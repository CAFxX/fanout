package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 int16 }
type GenS2 struct { F0 int16; F1 int32; F2 uint8 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int; F1 int64 }
type GenS5 struct { F0 int64; F1 int16 }
func genFn0(p0 uint16, p1 int32) string {
    v5 := []int{2147483647, 9223372036854775807, 42}
    var v6 int64 = -409
    switch 761 {
    case -449:
        fmt.Println(p0, p1)
    case 9223372036854775807:
        v7 := GenS1{F0: 245, F1: 255, F2: 42}
        _ = v7
    default:
        p0 += 894
    }
    v8 := GenS1{F0: 18446744073709551615, F1: 0, F2: -32768}
    v9 := GenS1{F0: 1, F1: 1, F2: 1000}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v8
    _ = v9
    return (("foo bar" + "\n\t") + ("a" + ("" + "")))
}

func main() {
    var v1 uint16 = 1
    var v2 uint64 = 0
    var v3 int8 = 1
    var v4 int16 = 42
    if (377 == 2) {
        v3 -= 100
    } else {
        v10 := GenS2{F0: 42, F1: 996, F2: 255}
        if (-988 == -9223372036854775808) {
            if false {
                v2 += (v2 - v2)
                fmt.Println(v1, v4)
            } else {
                var v11 uint64 = (^(v2))
                _ = v11
            }
        } else {
            if (255 == 4294967295) {
                if true {
                    fmt.Println(v4, v3, v4)
                    fmt.Println(v4, v4, v3)
                }
                var v12 uint = 18446744073709551615
                var v13 uint16 = v1
                _ = v12
                _ = v13
            }
            var v14 uint32 = 4294967295
            _ = v14
        }
        _ = v10
    }
    v15 := []int{}
    v4 += v4
    for fi16 := 0; fi16 < 20; fi16++ {
        switch "a" {
        case "a":
            var v17 uint64 = (v2 ^ v2)
            _ = v17
        case "\n\t":
            v18 := GenS2{F0: 899, F1: 972, F2: 1}
            _ = v18
        default:
            for fi19 := 0; fi19 < 4; fi19++ {
                var v20 int64 = 1
                v21 := GenS3{F0: (fi19 % ((fi19) | 1))}
                _ = fi19
                _ = v20
                _ = v21
            }
        }
        var v22 int64 = -9223372036854775808
        v3 = (v3 | (v3 / ((v3) | 1)))
        _ = fi16
        _ = v22
    }
    if len(v15) > 0 {
        v15[0] = -1
    }
    if true {
        v23 := GenS3{F0: 127}
        v3 = v3
        v4 += (v4 | v4)
        _ = v23
    }
    for fi24 := 0; fi24 < 13; fi24++ {
        _ = len(v15)
        _ = fi24
    }
    if false {
        fmt.Println(v1, v1, v3)
        if len(v15) >= 2 {
            v15 = v15[1:len(v15)]
        }
    } else {
        if len(v15) >= 2 {
            v15 = v15[1:len(v15)]
        }
        var v25 int8 = 0
        _ = v25
    }
    fmt.Println(v1)
    if false {
        v15 = append(v15, 1)
        fmt.Println(v2, v4)
    } else {
        switch -269 {
        case 4294967295:
            var v26 bool = (v4 > (v4 % ((32767) | 1)))
            _ = v26
        case 64:
            var v27 uint = 42
            _ = v27
        case -2147483648:
            var v28 int16 = 1
            _ = v28
        }
        v4 = v4
        v15 = append(v15, 4294967295)
    }
    for fi29 := 0; fi29 < 5; fi29++ {
        if (true && false) {
            v15 = append(v15, (fi29 - fi29))
        }
        _ = fi29
    }
    if false {
        var v30 int8 = v3
        v31 := GenS4{F0: -1, F1: 0}
        _ = v30
        _ = v31
    } else {
        v4 -= (v4 | (^(v4)))
        for fi32 := 0; fi32 < 17; fi32++ {
            v1 = (v1 & v1)
            v33 := GenS5{F0: 0, F1: (v4 / ((int16(v2)) | 1))}
            _ = fi32
            _ = v33
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
