package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 uint32; F1 uint64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint64 }
type GenS5 struct { F0 uint16; F1 int16; F2 int }
type GenS6 struct { F0 string; F1 int32 }
type GenS7 struct { F0 uint32; F1 int16; F2 uint }
func genFn0(p0 string, p1 int) uint8 {
    v7 := []int{(p1 % ((p1) | 1)), (p1 % (((p1 / ((p1) | 1))) | 1))}
    p0 = (p0 + (p0 + p0))
    v8 := GenS1{F0: 42, F1: 2}
    var v9 float32 = 0.1
    var v10 string = ("a" + p0)
    _ = p0
    _ = p1
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 127
}
func genFn1(p0 uint64, p1 int64, p2 bool) int64 {
    var v11 int16 = -447
    for fi12 := 0; fi12 < 20; fi12++ {
        v13 := []int{(fi12 - fi12)}
        _ = len(v13)
        if (42 > 567) {
            var v14 float32 = float32(float64(0.0))
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
            v15 := GenS2{F0: 2147483648, F1: p0}
            _ = v14
            _ = v15
        } else {
            var v16 uint16 = 182
            v17 := GenS3{F0: ""}
            for fi18 := 0; fi18 < 1; fi18++ {
                if p2 {
                    v13 = append(v13, fi18)
                } else {
                    fmt.Println(fi18)
                }
                var v19 uint16 = 1
                _ = fi18
                _ = v19
            }
            _ = v16
            _ = v17
        }
        _ = fi12
        _ = v13
    }
    if p2 {
        v20 := []int{9223372036854775807}
        switch ("" + "") {
        case "foo bar":
            for fi21 := 0; fi21 < 2; fi21++ {
                var v22 uint16 = 1
                if ((-(p1)) < (p1 | p1)) {
                    p0 = (p0 & (p0 | p0))
                    var v23 float64 = ((3.14159 - 1e10) - 0.1)
                    if len(v20) >= 2 {
                        v20 = v20[1:len(v20)]
                    }
                    _ = v23
                }
                v20 = append(v20, 127)
                _ = fi21
                _ = v22
            }
        case "\n\t":
            var v24 float32 = 1.5
            _ = v24
        case "hello":
            var v25 uint8 = 212
            _ = v25
        default:
            var v26 int16 = v11
            _ = v26
        }
        var v27 int16 = (v11 + (v11 << ((255) & 63)))
        _ = v20
        _ = v27
    } else {
        for fi28 := 0; fi28 < 3; fi28++ {
            v11 -= v11
            if p2 {
                var v29 string = "foo bar"
                var v30 int64 = p1
                var v31 float32 = -2.25
                _ = v29
                _ = v30
                _ = v31
            }
            _ = fi28
        }
    }
    var v32 int64 = (p1 ^ p1)
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v32
    return p1
}

func main() {
    var v1 float64 = (1e10 * 3.14159)
    var v2 string = ""
    var v3 int = 9223372036854775807
    var v4 int8 = 112
    var v5 uint8 = 2
    v6 := make(map[uint32]uint32)
    v33 := GenS4{F0: 2}
    v3 += 888
    _ = len(v6)
    v2 = "hello"
    switch (v3 % ((v3) | 1)) {
    case 269:
        if true {
            switch int(v5) {
            case 0:
                switch v2 {
                case "foo bar":
                    v3 += (-(v3))
                }
            case -1:
                if false {
                    if true {
                        v34 := []int{(-(v3)), v3, 1}
                        if ((1e10 <= 1.5) || (2147483648 > 696)) {
                            _ = len(v6)
                        } else {
                            v35 := GenS5{F0: 42, F1: -935, F2: v3}
                            if len(v34) > 0 {
                                v34[0] = 799
                            }
                            _ = v35
                        }
                        _ = v34
                    }
                    var v36 int = (v3 & v3)
                    switch v2 {
                    case "foo bar":
                        var v37 uint64 = 1
                        _ = v37
                    }
                    _ = v36
                }
            case 9223372036854775807:
                for fi38 := 0; fi38 < 12; fi38++ {
                    v39 := []int{9223372036854775807, (^(v3)), 4294967295, -765}
                    _ = fi38
                    _ = v39
                }
            default:
                v40 := []int{}
                _ = v40
            }
            v2 = "\n\t"
            v41 := GenS6{F0: "hello", F1: 1}
            _ = v41
        } else {
            v42 := GenS7{F0: 42, F1: 32767, F2: 42}
            var v43 string = v2
            _ = v42
            _ = v43
        }
    case -693:
        var v44 int16 = 42
        _ = v44
    }
    v45 := v6[996]
    v5 = 137
    fmt.Println(v3)
    v4 = v4
    var v46 bool = (!false)
    _ = len(v6)
    fmt.Println(v4)
    var v47 int8 = -71
    var v48 int32 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v48
    _ = v5
    _ = v6
}
