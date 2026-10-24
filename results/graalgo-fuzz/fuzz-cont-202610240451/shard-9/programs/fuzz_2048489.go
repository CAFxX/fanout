package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32; F2 string }
type GenS2 struct { F0 string; F1 bool; F2 bool }
type GenS3 struct { F0 int }

func main() {
    var v1 uint16 = 65535
    var v2 uint8 = 0
    var v3 uint = 65535
    v4 := make(map[int]string)
    var v5 uint8 = (v2 % ((v2) | 1))
    var v6 bool = (0 == 992)
    fmt.Println(v3)
    v5 -= 61
    var v7 uint16 = (^(v1))
    var v8 string = "a"
    v9 := GenS1{F0: 5, F1: 1, F2: ((v8 + v8) + "\n\t")}
    switch v8 {
    case "foo bar":
        v1 -= v7
    case "hello":
        for fi10 := 0; fi10 < 13; fi10++ {
            fmt.Println(v5, v5, v1)
            _ = fi10
        }
    default:
        fmt.Println(v3, v1)
    }
    var v11 int64 = -938
    for fi12 := 0; fi12 < 6; fi12++ {
        if v6 {
            v13 := GenS1{F0: 285, F1: 65535, F2: ((v8 + v8) + v8)}
            var v14 int16 = 42
            _ = v13
            _ = v14
        } else {
            if (!(-9223372036854775808 >= (v11 + (-(v11))))) {
                fmt.Println(v1)
                var v15 float32 = 0.1
                _ = v15
            }
            v2 -= v5
        }
        for fi16 := 0; fi16 < 6; fi16++ {
            fmt.Println(v3, fi16, v6)
            _ = fi16
        }
        v17 := GenS1{F0: 688, F1: 1, F2: "hello"}
        _ = fi12
        _ = v17
    }
    switch ("hello" + "a") {
    case "foo bar":
        switch -1 {
        case -2147483648:
            var v18 uint32 = 664
            _ = v18
        case -1:
            switch 42 {
            case 9223372036854775807:
                fmt.Println(v3, v5, v3)
            default:
                v19 := GenS2{F0: v8, F1: (164 == v5), F2: v6}
                _ = v19
            }
        case 128:
            v20 := GenS3{F0: 2147483647}
            _ = v20
        default:
            if v6 {
                var v21 uint8 = (v5 * v5)
                _ = v21
            }
        }
    }
    for fi22 := 0; fi22 < 10; fi22++ {
        v8 = v8
        _ = fi22
    }
    fmt.Println(v5, v7, v6)
    if (false && false) {
        if (int32(v11) <= -410) {
            for fi23 := 0; fi23 < 19; fi23++ {
                v24 := GenS1{F0: 18446744073709551615, F1: 2147483647, F2: v8}
                var v25 int = (fi23 / (((^(fi23))) | 1))
                _ = fi23
                _ = v24
                _ = v25
            }
            v3 += (v3 >> (((v3 << (((v3 ^ 1)) & 63))) & 63))
        } else {
            v26 := GenS1{F0: 2, F1: 334, F2: v8}
            v6 = v6
            v6 = true
            _ = v26
        }
        if (!(!false)) {
            delete(v4, 755)
            v6 = true
            for fi27 := 0; fi27 < 18; fi27++ {
                for fi28 := 0; fi28 < 19; fi28++ {
                    var v29 int = (fi28 ^ fi27)
                    v30 := GenS3{F0: (v29 % (((v29 / ((v29) | 1))) | 1))}
                    v5 -= 0
                    _ = fi28
                    _ = v29
                    _ = v30
                }
                _ = fi27
            }
        } else {
            v31 := GenS1{F0: 42, F1: 4294967295, F2: ("" + v8)}
            _ = v31
        }
    } else {
        for fi32 := 0; fi32 < 4; fi32++ {
            v7 += (v7 + 191)
            _ = fi32
        }
        for fi33 := 0; fi33 < 18; fi33++ {
            v11 = v11
            _ = fi33
        }
    }
    v11 = (v11 / ((v11) | 1))
    fmt.Println(v3, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
