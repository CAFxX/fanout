package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint; F2 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint32; F1 uint64 }
type GenS4 struct { F0 int8; F1 uint32; F2 string }
func genFn0(p0 int32, p1 string) uint8 {
    if (!true) {
        p0 += -252
        var v6 int64 = 42
        _ = v6
    } else {
        for fi7 := 0; fi7 < 3; fi7++ {
            var v8 int = 2147483647
            _ = fi7
            _ = v8
        }
    }
    if true {
        v9 := []int{-2147483648, 9223372036854775807, 4294967295}
        v9 = append(v9, 9223372036854775807)
        _ = v9
    } else {
        if false {
            for fi10 := 0; fi10 < 9; fi10++ {
                for fi11 := 0; fi11 < 5; fi11++ {
                    v12 := GenS1{F0: 996, F1: 704, F2: (fi11 * fi10)}
                    _ = fi11
                    _ = v12
                }
                v13 := []int{(-(fi10)), (fi10 + fi10)}
                p0 += p0
                _ = fi10
                _ = v13
            }
        } else {
            v14 := GenS1{F0: 850, F1: 42, F2: 937}
            var v15 int8 = -128
            _ = v14
            _ = v15
        }
        for fi16 := 0; fi16 < 4; fi16++ {
            var v17 int = 127
            _ = fi16
            _ = v17
        }
    }
    fmt.Println(p1)
    switch p1 {
    case "hello":
        v18 := GenS2{F0: 1000}
        _ = v18
    case "\n\t":
        switch 4294967295 {
        case 397:
            var v19 bool = (32767 != 851)
            _ = v19
        default:
            p0 += (p0 % (((p0 % (((p0 ^ p0)) | 1))) | 1))
        }
    case "foo bar":
        p1 = (p1 + "")
    }
    var v20 string = p1
    _ = p0
    _ = p1
    _ = v20
    return 2
}
func genFn1(p0 int16, p1 int16) uint8 {
    var v21 string = ("foo bar" + ("a" + ""))
    if true {
        var v22 int32 = -1
        for fi23 := 0; fi23 < 13; fi23++ {
            p1 -= p1
            fmt.Println(p0, v22)
            _ = fi23
        }
        _ = v22
    }
    var v24 float64 = 0.1
    fmt.Println(p0, p0, v21)
    _ = p0
    _ = p1
    _ = v21
    _ = v24
    return 42
}

func main() {
    var v1 bool = true
    var v2 int16 = 32767
    var v3 int = 128
    var v4 uint32 = 2
    var v5 float64 = (3.14159 / 1.5)
    for fi25 := 0; fi25 < 11; fi25++ {
        switch v3 {
        case 9223372036854775807:
            v26 := []int{(fi25 * (fi25 * (fi25 - 127)))}
            _ = v26
        }
        _ = fi25
    }
    v4 -= (v4 | v4)
    var v27 int8 = 0
    if v1 {
        v4 += 65535
        switch v3 {
        case 0:
            v28 := GenS3{F0: 42, F1: 62}
            _ = v28
        }
        v4 += v4
    } else {
        v29 := []int{v3}
        _ = v29
    }
    v30 := GenS3{F0: (v4 * v4), F1: 18446744073709551615}
    if ((v1 || (!v1)) && true) {
        fmt.Println(v5, v27, v3)
        v31 := []int{}
        _ = v31
    }
    v32 := GenS4{F0: v27, F1: 2147483647, F2: ""}
    v2 += (v2 - v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v30
    _ = v32
    _ = v4
    _ = v5
}
