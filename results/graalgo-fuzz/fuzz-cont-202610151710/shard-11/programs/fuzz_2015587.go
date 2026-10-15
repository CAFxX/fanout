package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32; F1 int64 }
type GenS4 struct { F0 bool }
func genFn0(p0 bool, p1 string, p2 uint16) string {
    if (p2 > p2) {
        var v4 uint32 = 446
        for fi5 := 0; fi5 < 14; fi5++ {
            v6 := GenS1{F0: 91}
            var v7 float64 = (((3.14159 * ((3.14159 - 3.14159) * float64(((1.5 - (float32(0.1) * -2.25)) + 3.14159)))) + 0.1) - (3.14159 * 0.0))
            _ = fi5
            _ = v6
            _ = v7
        }
        switch -76 {
        case -751:
            var v8 string = ("\n\t" + "foo bar")
            _ = v8
        case -2147483648:
            for fi9 := 0; fi9 < 6; fi9++ {
                var v10 float64 = 3.14159
                if p0 {
                    v11 := GenS2{F0: 2}
                    switch (fi9 >> ((255) & 63)) {
                    case 32767:
                        p2 -= p2
                    case 2147483647:
                        p1 = p1
                    }
                    p1 = ((p1 + "a") + "\n\t")
                    _ = v11
                }
                _ = fi9
                _ = v10
            }
        default:
            v4 += (v4 | 2147483648)
        }
        _ = v4
    } else {
        p2 += p2
        v12 := []int{42, 1, 9223372036854775807}
        _ = v12
    }
    for fi13 := 0; fi13 < 2; fi13++ {
        var v14 float32 = float32(0.0)
        _ = fi13
        _ = v14
    }
    var v15 float64 = -2.25
    p0 = true
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return p1
}

func main() {
    v1 := make(map[int]int)
    var v2 int16 = 1000
    var v3 int = 32767
    v16 := GenS2{F0: 2}
    v3 -= v3
    v2 -= v2
    if ((255 >= 383) && (("foo bar" + "") < "\n\t")) {
        switch v3 {
        case 1:
            fmt.Println(v2, v3)
        }
        v17 := []int{v3, (^(v3)), (v3 & 128), (v3 + v3)}
        _ = v17
    } else {
        if true {
            v3 = v3
            if (128 != v3) {
                v18 := []int{}
                _ = len(v18)
                if (65 < 128) {
                    v19 := GenS1{F0: 0}
                    if true {
                        v20 := GenS3{F0: -1, F1: -243}
                        _ = v20
                    } else {
                        v3 = (v3 - -1)
                    }
                    delete(v1, (v3 | 670))
                    _ = v19
                }
                _ = v18
            } else {
                if true {
                    v21 := []int{v3, 32767}
                    var v22 uint16 = 796
                    _ = v21
                    _ = v22
                } else {
                    for fi23 := 0; fi23 < 20; fi23++ {
                        fmt.Println(fi23, v2)
                        v24 := GenS3{F0: 2147483647, F1: -509}
                        v3 = 1
                        _ = fi23
                        _ = v24
                    }
                }
                for fi25 := 0; fi25 < 8; fi25++ {
                    if (!false) {
                        var v26 uint32 = 788
                        v27 := []int{(fi25 >> ((18446744073709551615) & 63)), v3, 127, fi25}
                        v3 = (v3 - (fi25 & (v3 - 98)))
                        _ = v26
                        _ = v27
                    } else {
                        v28 := []int{}
                        fmt.Println(fi25, v2)
                        v29 := GenS3{F0: 585, F1: -9223372036854775808}
                        _ = v28
                        _ = v29
                    }
                    switch fi25 {
                    case 32767:
                        v3 = (fi25 & -9223372036854775808)
                    }
                    fmt.Println(v2, v2)
                    _ = fi25
                }
                var v30 string = "hello"
                _ = v30
            }
            v31 := GenS4{F0: ("foo bar" <= "a")}
            _ = v31
        } else {
            var v32 bool = (((127 <= -1) || (677 >= 505)) || (!false))
            _ = v32
        }
        var v33 float32 = 1e10
        for fi34 := 0; fi34 < 19; fi34++ {
            var v35 bool = false
            var v36 string = ""
            _ = fi34
            _ = v35
            _ = v36
        }
        _ = v33
    }
    v2 += v2
    v3 += v3
    fmt.Println(v3, v3)
    v37 := []int{}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v37:", len(v37), v37)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v37
}
