package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 int64; F1 uint }
type GenS3 struct { F0 bool; F1 uint32; F2 string }
type GenS4 struct { F0 bool; F1 uint64; F2 uint16 }
type GenS5 struct { F0 int32; F1 uint8 }
func genFn0(p0 uint16, p1 int32) uint64 {
    p1 += p1
    for fi6 := 0; fi6 < 17; fi6++ {
        if true {
            v7 := GenS1{F0: (fi6 >> ((42) & 63)), F1: 42}
            if true {
                for fi8 := 0; fi8 < 6; fi8++ {
                    fmt.Println(fi8, fi6, p0)
                    v9 := []int{fi8, (fi8 ^ (fi8 << ((412) & 63)))}
                    v10 := GenS2{F0: 1, F1: 390}
                    if true {
                        fmt.Println(fi6)
                        var v11 uint16 = p0
                        _ = v11
                    }
                    switch "hello" {
                    case "":
                        if (357 != -461) {
                            _ = len(v9)
                            v12 := GenS1{F0: fi6, F1: 18446744073709551615}
                            _ = v12
                        } else {
                            var v13 int32 = (p1 ^ p1)
                            _ = len(v9)
                            p1 -= p1
                            _ = v13
                        }
                    case "foo bar":
                        _ = len(v9)
                    case "\n\t":
                        switch "" {
                        case "a":
                            p0 = p0
                        }
                    default:
                        var v14 bool = false
                        _ = v14
                    }
                    _ = fi8
                    _ = v10
                    _ = v9
                }
                p1 = p1
            }
            v15 := []int{}
            _ = v15
            _ = v7
        }
        _ = fi6
    }
    fmt.Println(p0)
    if false {
        for fi16 := 0; fi16 < 4; fi16++ {
            if (!(456 > 42)) {
                var v17 uint8 = 127
                for fi18 := 0; fi18 < 4; fi18++ {
                    var v19 uint16 = (p0 * 65535)
                    var v20 string = "a"
                    v21 := GenS3{F0: false, F1: 586, F2: "foo bar"}
                    _ = fi18
                    _ = v19
                    _ = v20
                    _ = v21
                }
                _ = v17
            }
            v22 := []int{fi16, (fi16 / (((fi16 / (((^(fi16))) | 1))) | 1))}
            var v23 float32 = 0.0
            _ = fi16
            _ = v22
            _ = v23
        }
        var v24 string = "\n\t"
        var v25 int16 = 0
        _ = v24
        _ = v25
    }
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    return 2
}

func main() {
    var v1 float32 = -2.25
    var v2 int64 = 9223372036854775807
    v3 := make(map[int]string)
    var v4 int32 = 2147483647
    v5 := []int{-615, 32768, 4294967295}
    for fi26 := 0; fi26 < 13; fi26++ {
        v1 = float32((-2.25 * 3.14159))
        v1 = (((v1 * float32((float64((v1 - v1)) - -2.25))) + (v1 + (((1e10 * (float32(float64((v1 - v1))) * (0.0 + 1.5))) - v1) / 3.14159))) + ((1.5 / float32(3.14159)) * 3.14159))
        _ = fi26
    }
    v4 -= 94
    var v27 int = -2147483648
    for fi28 := 0; fi28 < 19; fi28++ {
        switch (-(fi28)) {
        case 9223372036854775807:
            var v29 uint8 = 0
            _ = v29
        case 127:
            var v30 uint16 = 65535
            _ = v30
        }
        _ = fi28
    }
    v31 := GenS3{F0: true, F1: 139, F2: ("a" + ("foo bar" + "a"))}
    var v32 uint64 = 1
    if (!false) {
        if (!true) {
            var v33 uint8 = 42
            v33 += (v33 | 255)
            v4 = v4
            _ = v33
        } else {
            v2 = (v2 * v2)
            v4 = 2147483647
            var v34 bool = ((v27 != v27) && false)
            _ = v34
        }
        v1 += v1
        v1 += float32(0.0)
    }
    v35 := GenS1{F0: v27, F1: 2}
    fmt.Println(v4, v2, v2)
    v36 := GenS4{F0: true, F1: (v32 & genFn0(2, v4)), F2: 65535}
    var v37 uint32 = 179
    for fi38 := 0; fi38 < 4; fi38++ {
        v39 := GenS5{F0: -118, F1: 42}
        _ = fi38
        _ = v39
    }
    v40 := GenS3{F0: (((!(float32((float64(v1) * 1.5)) != ((v1 / 1e10) - (1e10 / 1.5)))) || (!true)) || false), F1: (^(v37)), F2: "foo bar"}
    var v41 int8 = 7
    var v42 uint64 = (v32 % ((v32) | 1))
    var v43 int64 = (^(v2))
    v44 := GenS5{F0: v4, F1: 84}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v3
    _ = v31
    _ = v32
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v5
}
