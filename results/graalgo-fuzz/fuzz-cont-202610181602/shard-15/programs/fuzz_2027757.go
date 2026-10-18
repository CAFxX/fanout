package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string; F1 uint32; F2 uint64 }
type GenS3 struct { F0 uint64; F1 int64 }
type GenS4 struct { F0 int8; F1 int16; F2 int16 }
type GenS5 struct { F0 uint; F1 int8 }
func genFn0(p0 int8, p1 int16) int32 {
    if (false && (32767 >= 134)) {
        if (false && false) {
            if (1 < 18446744073709551615) {
                p1 = -1
            }
        } else {
            for fi4 := 0; fi4 < 19; fi4++ {
                if (255 >= 97) {
                    v5 := []int{(fi4 >> ((65535) & 63)), fi4}
                    switch "" {
                    case "\n\t":
                        p1 = (p1 << ((65535) & 63))
                    }
                    switch ("foo bar" + ((("foo bar" + "foo bar") + "a") + (("\n\t" + ("foo bar" + "")) + "foo bar"))) {
                    case "foo bar":
                        v6 := GenS1{F0: (p0 * -74)}
                        _ = v6
                    default:
                        var v7 bool = (0 == 37)
                        _ = v7
                    }
                    _ = v5
                } else {
                    p0 += -97
                    var v8 int = (fi4 & fi4)
                    for fi9 := 0; fi9 < 6; fi9++ {
                        v10 := GenS2{F0: "foo bar", F1: 0, F2: 0}
                        _ = fi9
                        _ = v10
                    }
                    _ = v8
                }
                v11 := GenS3{F0: 152, F1: -882}
                _ = fi4
                _ = v11
            }
        }
        v12 := []int{9223372036854775807, 2147483647, -573, 127}
        p0 = p0
        _ = v12
    }
    var v13 uint8 = 42
    _ = p0
    _ = p1
    _ = v13
    return -650
}

func main() {
    var v1 bool = true
    var v2 uint = 800
    var v3 uint8 = 42
    v14 := GenS4{F0: 42, F1: 42, F2: 1}
    if true {
        switch "\n\t" {
        case "\n\t":
            v1 = (!false)
        case "":
            v15 := []int{}
            _ = v15
        case "foo bar":
            for fi16 := 0; fi16 < 19; fi16++ {
                v17 := GenS5{F0: 42, F1: -128}
                v18 := []int{fi16, -1}
                _ = fi16
                _ = v17
                _ = v18
            }
        }
        v19 := GenS5{F0: 2, F1: 100}
        fmt.Println(v3, v3)
        _ = v19
    }
    var v20 int32 = int32(v2)
    v2 = v2
    v21 := []int{2147483647, 1, 820, -2147483648}
    if len(v21) > 0 {
        v21[0] = -2147483648
    }
    var v22 uint32 = 881
    for fi23 := 0; fi23 < 14; fi23++ {
        v24 := GenS3{F0: 4294967295, F1: 1}
        v21 = append(v21, fi23)
        var v25 string = ((("a" + "") + "a") + "a")
        _ = fi23
        _ = v24
        _ = v25
    }
    v22 -= (^(v22))
    v20 = (v20 >> ((v2) & 63))
    v21 = append(v21, 1)
    if (v22 >= (^(v22))) {
        v21 = append(v21, 386)
        v21 = append(v21, 2147483647)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
}
