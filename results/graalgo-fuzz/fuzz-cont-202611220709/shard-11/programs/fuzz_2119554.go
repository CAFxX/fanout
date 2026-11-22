package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 int32 }
type GenS2 struct { F0 uint64; F1 uint8 }
type GenS3 struct { F0 int16; F1 int32; F2 bool }
type GenS4 struct { F0 int64 }
func genFn0(p0 uint16, p1 int, p2 uint32) int {
    v7 := GenS1{F0: 127, F1: "", F2: -2147483648}
    if (!false) {
        if (-1 == 914) {
            v8 := GenS2{F0: 42, F1: 0}
            v9 := []int{}
            p0 = (p0 * p0)
            _ = v8
            _ = v9
        } else {
            if (1 < 122) {
                v10 := GenS2{F0: uint64(p2), F1: 2}
                var v11 uint64 = 18446744073709551615
                _ = v10
                _ = v11
            } else {
                var v12 int64 = 0
                var v13 uint = 65535
                var v14 int8 = 127
                _ = v12
                _ = v13
                _ = v14
            }
        }
        v15 := []int{p1}
        _ = v15
    }
    if false {
        v16 := []int{p1, -9223372036854775808}
        if len(v16) >= 2 {
            v16 = v16[1:len(v16)]
        }
        _ = v16
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return p1
}

func main() {
    var v1 bool = ((-9223372036854775808 == 9223372036854775807) || (1 <= 292))
    var v2 int64 = 42
    var v3 uint = 2
    v4 := make(map[int]string)
    var v5 float32 = (3.14159 / (((1e10 * -2.25) + 0.0) * (3.14159 / 0.1)))
    v6 := []int{42, -356, 641}
    var v17 uint = v3
    for fi18 := 0; fi18 < 8; fi18++ {
        for fi19 := 0; fi19 < 3; fi19++ {
            v20 := GenS2{F0: 2, F1: 92}
            switch "\n\t" {
            case "hello":
                var v21 bool = v1
                _ = v21
            case "a":
                v22 := GenS1{F0: -1, F1: "\n\t", F2: -1}
                _ = v22
            case "\n\t":
                var v23 bool = (-32768 >= 632)
                _ = v23
            default:
                v4[fi19] = ("" + "a")
            }
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            _ = len(v4)
            if v1 {
                fmt.Println(v3, v2, v3)
                v5 -= (1.5 / v5)
            }
            _ = fi19
            _ = v20
        }
        var v24 uint32 = 4294967295
        v25 := v4[(^(fi18))]
        _ = fi18
        _ = v24
        _ = v25
    }
    if false {
        switch genFn0(1, 42, 4294967295) {
        case 9223372036854775807:
            switch "a" {
            case "a":
                v26 := GenS1{F0: 42, F1: "\n\t", F2: -1}
                _ = v26
            default:
                if true {
                    if len(v6) > 0 {
                        v6[0] = -649
                    }
                    var v27 uint64 = 18446744073709551615
                    for fi28 := 0; fi28 < 14; fi28++ {
                        if (!v1) {
                            var v29 int16 = 143
                            _ = v29
                        }
                        _ = fi28
                    }
                    _ = v27
                }
            }
        }
        v30 := GenS1{F0: 42, F1: "hello", F2: 0}
        _ = v30
    }
    for fi31 := 0; fi31 < 20; fi31++ {
        for fi32 := 0; fi32 < 15; fi32++ {
            v33 := GenS1{F0: 0, F1: "hello", F2: -2147483648}
            var v34 uint32 = 42
            if v1 {
                v4[-9223372036854775808] = ""
            }
            var v35 string = "foo bar"
            v36 := v4[-151]
            if (128 == (fi32 / (((fi31 | fi32)) | 1))) {
                if len(v6) > 0 {
                    v6[0] = fi32
                }
                var v37 float32 = ((3.14159 - v5) + float32(float64(v5)))
                v38 := GenS1{F0: 111, F1: "", F2: -562}
                _ = v37
                _ = v38
            } else {
                v17 = 885
            }
            switch int(v3) {
            case 531:
                v39 := v4[genFn0(1, fi31, (v34 << (((^(v17))) & 63)))]
                _ = v39
            case 2147483647:
                v40 := GenS1{F0: 115, F1: (v36 + v36), F2: 42}
                _ = v40
            case 32767:
                if (227 <= 255) {
                    fmt.Println(v2, v1, v36)
                    fmt.Println(v34)
                    v41 := GenS3{F0: 1, F1: -2147483648, F2: (v35 == v35)}
                    v2 = v2
                    _ = v41
                }
            default:
                v42 := GenS3{F0: 739, F1: 42, F2: ((v36 + v35) < v36)}
                _ = v42
            }
            fmt.Println(v1)
            switch ("foo bar" + v35) {
            case "hello":
                var v43 uint16 = 671
                _ = v43
            }
            var v44 int32 = -2147483648
            var v45 float64 = 1.5
            _ = fi32
            _ = v33
            _ = v34
            _ = v35
            _ = v36
            _ = v44
            _ = v45
        }
        _ = fi31
    }
    v46 := GenS4{F0: (v2 << ((v17) & 63))}
    v47 := GenS2{F0: 0, F1: 26}
    v4[942] = ""
    v1 = true
    var v48 uint = v3
    v17 = v48
    v49 := v4[801]
    if (v5 != v5) {
        if true {
            var v50 uint8 = 0
            if (v1 || true) {
                v50 += v50
            }
            var v51 uint16 = 1
            _ = v50
            _ = v51
        }
    }
    v6 = append(v6, genFn0(0, 32768, 0))
    var v52 int32 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v52
    _ = v6
}
