package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8; F2 int16 }
type GenS2 struct { F0 int; F1 int64 }
type GenS3 struct { F0 int; F1 int16 }
type GenS4 struct { F0 int16; F1 uint64 }
func genFn0(p0 int32, p1 uint8, p2 string) uint16 {
    fmt.Println(p1, p1)
    if true {
        v7 := GenS1{F0: -1, F1: 1, F2: 0}
        _ = v7
    } else {
        p1 += 2
        p1 = (p1 ^ p1)
    }
    v8 := []int{4294967295}
    var v9 uint64 = 70
    p2 = "\n\t"
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 32767
}

func main() {
    var v1 int64 = -1
    var v2 int64 = (v1 % ((v1) | 1))
    var v3 int16 = -32768
    v4 := []int{336, -109, 9223372036854775807}
    var v5 bool = ((1e10 != -2.25) && (false || false))
    v6 := make(map[int]int)
    if true {
        if v5 {
            v5 = (128 >= 41)
            var v10 uint8 = 255
            _ = v10
        }
        if len(v4) > 0 {
            v4[0] = 960
        }
    } else {
        var v11 int8 = 127
        _ = v11
    }
    var v12 int8 = -1
    var v13 uint64 = 42
    var v14 int8 = (-(v12))
    v14 += int8(v3)
    for fi15 := 0; fi15 < 5; fi15++ {
        v3 = 1
        _ = fi15
    }
    if v5 {
        v13 = (^(v13))
        for fi16 := 0; fi16 < 1; fi16++ {
            var v17 uint = 65535
            _ = fi16
            _ = v17
        }
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    } else {
        switch "\n\t" {
        case "a":
            var v18 float32 = -2.25
            _ = v18
        }
    }
    var v19 uint = 1
    var v20 int64 = int64(v13)
    v21 := GenS1{F0: int64(v12), F1: v12, F2: 42}
    v14 += (v14 + v12)
    v22 := GenS2{F0: 71, F1: 42}
    v23 := GenS2{F0: 128, F1: 970}
    v6[-488] = -774
    switch -2147483648 {
    case -2147483648:
        switch 128 {
        case -963:
            switch "foo bar" {
            case "hello":
                fmt.Println(v2, v12, v13)
            case "":
                for fi24 := 0; fi24 < 1; fi24++ {
                    v25 := GenS3{F0: (fi24 >> (((v19 - 2)) & 63)), F1: v3}
                    _ = fi24
                    _ = v25
                }
            }
        case 1:
            v5 = (!(v19 <= (^(v19))))
        case 127:
            if v5 {
                for fi26 := 0; fi26 < 15; fi26++ {
                    v27 := GenS4{F0: (v3 & (v3 >> ((1) & 63))), F1: v13}
                    for fi28 := 0; fi28 < 4; fi28++ {
                        if len(v4) > 0 {
                            v4[0] = fi26
                        }
                        v14 = v14
                        _ = fi28
                    }
                    _ = fi26
                    _ = v27
                }
                v1 = (v20 >> ((v19) & 63))
                for fi29 := 0; fi29 < 16; fi29++ {
                    if (2147483647 < 42) {
                        var v30 int64 = -331
                        _ = v30
                    }
                    _ = fi29
                }
            }
        default:
            v14 = (v14 * -1)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
