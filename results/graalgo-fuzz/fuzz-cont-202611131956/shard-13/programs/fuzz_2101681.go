package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 uint8 }
type GenS2 struct { F0 uint8; F1 int8; F2 bool }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 string; F1 uint8 }
func genFn0(p0 int, p1 int16) bool {
    for fi4 := 0; fi4 < 8; fi4++ {
        fmt.Println(fi4)
        _ = fi4
    }
    for fi5 := 0; fi5 < 19; fi5++ {
        for fi6 := 0; fi6 < 17; fi6++ {
            var v7 int16 = (p1 * p1)
            var v8 int16 = v7
            _ = fi6
            _ = v7
            _ = v8
        }
        p0 = 310
        _ = fi5
    }
    var v9 int16 = 705
    _ = p0
    _ = p1
    _ = v9
    return true
}
func genFn1() uint64 {
    v10 := []int{-2147483648}
    for fi11 := 0; fi11 < 6; fi11++ {
        var v12 int16 = 32767
        _ = fi11
        _ = v12
    }
    _ = v10
    return 982
}

func main() {
    var v1 int8 = 1
    var v2 uint = 18446744073709551615
    v3 := make(map[string]bool)
    v13 := []int{32768, 0, 2147483647, -1}
    v14 := GenS1{F0: v1, F1: 1, F2: 0}
    var v15 uint64 = 2
    v13 = append(v13, 32768)
    var v16 uint8 = 255
    v17 := GenS2{F0: (v16 % ((v16) | 1)), F1: (-(v1)), F2: (v1 <= -1)}
    var v18 uint = v2
    if ((42 >= 32767) || true) {
        var v19 uint = v2
        v20 := GenS3{F0: 42}
        if (true || false) {
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
            for fi21 := 0; fi21 < 13; fi21++ {
                v2 = v18
                v22 := GenS3{F0: 255}
                for fi23 := 0; fi23 < 19; fi23++ {
                    switch "a" {
                    case "hello":
                        var v24 string = (("hello" + "\n\t") + ((("\n\t" + "") + "\n\t") + ((("a" + "a") + "foo bar") + "hello")))
                        _ = v24
                    }
                    _ = fi23
                }
                _ = fi21
                _ = v22
            }
            var v25 float64 = 0.1
            _ = v25
        } else {
            v26 := GenS2{F0: 1, F1: (v1 * (v1 % ((100) | 1))), F2: (int16(v1) <= 1)}
            var v27 int16 = -32768
            _ = v26
            _ = v27
        }
        _ = v19
        _ = v20
    }
    v28 := v3[("a" + "a")]
    fmt.Println(v1, v15, v1)
    var v29 int16 = 122
    fmt.Println(v28)
    if true {
        v30 := GenS3{F0: 42}
        if v28 {
            for fi31 := 0; fi31 < 8; fi31++ {
                var v32 int64 = 0
                switch "a" {
                case "":
                    fmt.Println(v15, fi31, v32)
                case "hello":
                    v32 -= 0
                case "foo bar":
                    var v33 float32 = 3.14159
                    _ = v33
                }
                _ = fi31
                _ = v32
            }
            v15 += (v15 | v15)
        } else {
            v34 := GenS3{F0: 724}
            v35 := GenS4{F0: ("" + (("\n\t" + ("" + "\n\t")) + "")), F1: (^(v16))}
            _ = v34
            _ = v35
        }
        var v36 string = "foo bar"
        _ = v30
        _ = v36
    }
    if (("" + "hello") <= "foo bar") {
        _ = len(v3)
        for fi37 := 0; fi37 < 7; fi37++ {
            v38 := GenS4{F0: ("foo bar" + ""), F1: v16}
            _ = fi37
            _ = v38
        }
        for fi39 := 0; fi39 < 19; fi39++ {
            v1 = (-(v1))
            var v40 uint16 = 42
            fmt.Println(v18)
            _ = fi39
            _ = v40
        }
    }
    fmt.Println(v2, v2)
    var v41 float32 = 1e10
    v42 := GenS4{F0: "\n\t", F1: 255}
    var v43 float32 = ((v41 * float32(1.5)) - 3.14159)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v41
    _ = v42
    _ = v43
}
