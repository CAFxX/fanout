package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool }
type GenS2 struct { F0 string; F1 int64; F2 int64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 bool; F1 uint; F2 string }
func genFn0(p0 int8, p1 bool, p2 uint) string {
    for fi7 := 0; fi7 < 17; fi7++ {
        v8 := GenS1{F0: p1, F1: (p2 > (p2 & (p2 >> ((p2) & 63))))}
        if p1 {
            if true {
                var v9 int64 = -592
                var v10 uint = p2
                var v11 int64 = 1
                _ = v10
                _ = v11
                _ = v9
            }
            if (!p1) {
                for fi12 := 0; fi12 < 6; fi12++ {
                    fmt.Println(p2, p0, fi12)
                    p2 = (p2 + 42)
                    v13 := []int{}
                    _ = fi12
                    _ = v13
                }
                fmt.Println(p1)
                p0 = (p0 | -1)
            }
            v14 := []int{(fi7 * (fi7 % (((fi7 * (fi7 / ((fi7) | 1)))) | 1)))}
            _ = v14
        }
        var v15 int = (fi7 | fi7)
        _ = fi7
        _ = v15
        _ = v8
    }
    for fi16 := 0; fi16 < 19; fi16++ {
        var v17 uint16 = 776
        _ = fi16
        _ = v17
    }
    if false {
        v18 := []int{-196}
        v19 := GenS2{F0: "a", F1: 975, F2: 717}
        _ = v18
        _ = v19
    } else {
        var v20 uint16 = 383
        v21 := []int{}
        if len(v21) > 0 {
            v21[0] = 128
        }
        _ = v20
        _ = v21
    }
    _ = p0
    _ = p1
    _ = p2
    return "foo bar"
}

func main() {
    var v1 int64 = 639
    var v2 int64 = v1
    v3 := make(map[string]uint32)
    var v4 uint32 = 65535
    v5 := []string{"\n\t", ("foo bar" + "a")}
    v6 := make(map[int]int)
    switch 4294967295 {
    case 2147483647:
        if (false || false) {
            v2 = (v1 + v1)
            var v22 uint = 18446744073709551615
            _ = v22
        } else {
            if false {
                var v23 int32 = -309
                fmt.Println(v4, v23)
                _ = v23
            }
            v4 = (^(v4))
        }
    default:
        v2 += v1
    }
    var v24 uint64 = 0
    fmt.Println(v4)
    v25 := GenS1{F0: (439 <= 42), F1: (true || ((3.14159 == 1e10) && false))}
    v26 := GenS3{F0: 0}
    var v27 int64 = (v1 + v1)
    var v28 uint64 = v24
    var v29 uint8 = 127
    var v30 int32 = 0
    for fi31 := 0; fi31 < 15; fi31++ {
        switch "" {
        case "":
            if false {
                var v32 bool = (false || (-436 != 32767))
                if (!v32) {
                    var v33 float64 = 0.0
                    v3[("" + "a")] = 4294967295
                    if (32767 < -1) {
                        v2 += (v1 * (v27 % ((int64(v4)) | 1)))
                    } else {
                        var v34 uint32 = (v4 % ((v4) | 1))
                        if ((3.14159 == 3.14159) || true) {
                            v34 = 2
                        }
                        var v35 uint64 = 42
                        _ = v34
                        _ = v35
                    }
                    _ = v33
                } else {
                    v36 := GenS1{F0: (!false), F1: (v1 < (v27 & v1))}
                    _ = v36
                }
                _ = len(v6)
                _ = v32
            }
        case "hello":
            v37 := GenS4{F0: false, F1: 920, F2: ""}
            _ = v37
        }
        _ = fi31
    }
    var v38 uint16 = 32767
    var v39 string = "hello"
    v40 := GenS3{F0: 945}
    v2 -= int64(v24)
    if len(v5) > 0 {
        v5[0] = ((("\n\t" + "hello") + "foo bar") + v39)
    }
    v28 = (v24 - 42)
    v41 := GenS2{F0: (v39 + v39), F1: v27, F2: v27}
    switch "a" {
    case "foo bar":
        var v42 uint64 = v24
        _ = v42
    case "":
        var v43 bool = true
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
