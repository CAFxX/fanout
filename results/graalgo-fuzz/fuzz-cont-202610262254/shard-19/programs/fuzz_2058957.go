package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 uint32 }
type GenS2 struct { F0 int32; F1 uint32; F2 uint }
type GenS3 struct { F0 uint64 }
func genFn0() uint32 {
    var v4 bool = true
    v4 = false
    _ = v4
    return 2147483648
}

func main() {
    v1 := make(map[int]string)
    var v2 float64 = 1e10
    var v3 uint = 4294967295
    for fi5 := 0; fi5 < 20; fi5++ {
        var v6 float64 = v2
        fmt.Println(fi5)
        switch "" {
        case "a":
            v1[-530] = ""
        case "":
            switch (fi5 * fi5) {
            case -9223372036854775808:
                v7 := GenS1{F0: "foo bar", F1: 342, F2: uint32(v3)}
                _ = v7
            case 409:
                fmt.Println(v2, v6)
            case -2147483648:
                v8 := []int{(fi5 * fi5)}
                _ = v8
            }
        default:
            v2 = v6
        }
        _ = fi5
        _ = v6
    }
    v3 = v3
    if (0 <= 42) {
        for fi9 := 0; fi9 < 4; fi9++ {
            v10 := []int{fi9, (fi9 ^ -2147483648)}
            v10 = append(v10, fi9)
            v3 += (v3 & (v3 % ((v3) | 1)))
            _ = fi9
            _ = v10
        }
        v11 := GenS1{F0: "hello", F1: 234, F2: 1}
        v12 := GenS1{F0: "foo bar", F1: 651, F2: 65535}
        _ = v11
        _ = v12
    } else {
        var v13 int64 = 0
        v2 = (0.1 / v2)
        fmt.Println(v13, v13, v13)
        _ = v13
    }
    var v14 bool = true
    var v15 int32 = 42
    var v16 int = 32767
    v14 = v14
    for fi17 := 0; fi17 < 5; fi17++ {
        var v18 uint32 = 4294967295
        var v19 int64 = 230
        _ = fi17
        _ = v18
        _ = v19
    }
    switch ("foo bar" + "a") {
    case "a":
        var v20 uint64 = 802
        _ = v20
    }
    var v21 uint = (v3 / (((v3 | (v3 & v3))) | 1))
    v1[v16] = "foo bar"
    if ("" == "") {
        v14 = (1 <= 117)
        v14 = v14
    } else {
        if (127 >= (-(v16))) {
            var v22 int8 = -128
            _ = v22
        } else {
            var v23 int32 = -263
            _ = v23
        }
    }
    for fi24 := 0; fi24 < 12; fi24++ {
        v25 := GenS2{F0: (v15 * (v15 >> ((870) & 63))), F1: 0, F2: v21}
        switch "a" {
        case "a":
            v21 += (v21 | v21)
        case "hello":
            for fi26 := 0; fi26 < 11; fi26++ {
                if (514 >= (-(v15))) {
                    fmt.Println(fi24)
                    v2 += (0.1 / 3.14159)
                    v27 := GenS3{F0: 1}
                    _ = v27
                }
                if v14 {
                    v28 := []int{-278, (fi26 & fi24)}
                    var v29 int64 = 9223372036854775807
                    var v30 string = "foo bar"
                    _ = v28
                    _ = v29
                    _ = v30
                }
                var v31 bool = true
                _ = fi26
                _ = v31
            }
        default:
            v32 := []int{}
            _ = v32
        }
        _ = fi24
        _ = v25
    }
    v33 := GenS3{F0: 1}
    v34 := []int{v16, 940}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v34:", len(v34), v34)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v3
    _ = v33
    _ = v34
}
