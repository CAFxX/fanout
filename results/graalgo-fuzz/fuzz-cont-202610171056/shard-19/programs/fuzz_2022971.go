package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 uint64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32; F1 string; F2 uint }
func genFn0(p0 uint64, p1 bool) int32 {
    v4 := GenS1{F0: 42, F1: 32767, F2: (p0 + 4294967295)}
    if true {
        p0 += (p0 / (((p0 << ((255) & 63))) | 1))
    }
    _ = p0
    _ = p1
    _ = v4
    return 14
}
func genFn1(p0 int8) int {
    for fi5 := 0; fi5 < 11; fi5++ {
        v6 := []int{(fi5 * 128), -87, (fi5 + -317), -1}
        v7 := GenS1{F0: 0, F1: 141, F2: 647}
        var v8 int = -9223372036854775808
        _ = fi5
        _ = v6
        _ = v7
        _ = v8
    }
    v9 := []int{412, -625}
    for fi10 := 0; fi10 < 11; fi10++ {
        var v11 uint = 478
        v12 := GenS2{F0: 630}
        p0 += p0
        _ = fi10
        _ = v11
        _ = v12
    }
    switch ((("a" + "\n\t") + "a") + "") {
    case "":
        fmt.Println(p0)
    case "\n\t":
        for fi13 := 0; fi13 < 18; fi13++ {
            if true {
                if (fi13 > fi13) {
                    var v14 uint = 18446744073709551615
                    v14 = (v14 - (v14 | 1))
                    switch fi13 {
                    case 128:
                        var v15 string = ("hello" + "")
                        _ = v15
                    case 0:
                        for fi16 := 0; fi16 < 16; fi16++ {
                            fmt.Println(v14, p0, fi13)
                            _ = fi16
                        }
                    case -113:
                        p0 += (p0 ^ p0)
                    default:
                        var v17 uint8 = 145
                        _ = v17
                    }
                    _ = v14
                } else {
                    var v18 float64 = 0.1
                    v19 := GenS2{F0: 42}
                    p0 = p0
                    _ = v18
                    _ = v19
                }
                v20 := GenS2{F0: 2}
                _ = v20
            }
            var v21 uint16 = 743
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            _ = fi13
            _ = v21
        }
    }
    _ = p0
    _ = v9
    return 127
}

func main() {
    var v1 uint64 = 392
    v2 := make(map[string]int)
    var v3 uint64 = 42
    v2[""] = 128
    var v22 float64 = 1e10
    var v23 uint = 18446744073709551615
    fmt.Println(v22, v3)
    if false {
        v24 := GenS2{F0: (v3 << ((v23) & 63))}
        v1 += (^(v1))
        _ = v24
    } else {
        v25 := GenS3{F0: 0, F1: "a", F2: v23}
        if true {
            v1 = (v3 - v3)
            fmt.Println(v22, v3)
        }
        _ = v25
    }
    if (!false) {
        fmt.Println(v3)
        v2["\n\t"] = -1
        var v26 uint = v23
        _ = v26
    }
    v23 += v23
    switch "foo bar" {
    case "foo bar":
        switch "" {
        case "a":
            v22 = (3.14159 + ((3.14159 * (v22 + (v22 / 1.0))) / float64(1.5)))
        default:
            var v27 int16 = -1
            _ = v27
        }
    case "":
        for fi28 := 0; fi28 < 14; fi28++ {
            v29 := []int{fi28, 4294967295, -1, (fi28 << ((v23) & 63))}
            _ = fi28
            _ = v29
        }
    default:
        for fi30 := 0; fi30 < 17; fi30++ {
            switch ("a" + "foo bar") {
            case "foo bar":
                v2["foo bar"] = (fi30 % ((-757) | 1))
            }
            v1 = (v1 & 18446744073709551615)
            _ = fi30
        }
    }
    v1 -= (v3 % ((4294967295) | 1))
    fmt.Println(v23, v22)
    fmt.Println(v3)
    if (("a" + ("\n\t" + "foo bar")) > "a") {
        v3 = (v1 / ((0) | 1))
        var v31 float64 = 3.14159
        _ = v31
    }
    var v32 int16 = 442
    fmt.Println(v23, v22, v3)
    v33 := []int{858}
    v33 = append(v33, -541)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v33:", len(v33), v33)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v32
    _ = v33
}
