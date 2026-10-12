package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int64; F1 int64; F2 int }

func main() {
    var v1 int64 = -94
    v2 := make(map[int]bool)
    var v3 int = -419
    var v4 float32 = ((0.0 / 1.0) + 0.0)
    switch ("hello" + "a") {
    case "":
        v5 := []int{v3, (v3 + v3)}
        _ = v5
    case "foo bar":
        v1 = (v1 & (v1 * (v1 / ((-48) | 1))))
    default:
        for fi6 := 0; fi6 < 13; fi6++ {
            fmt.Println(fi6, v3)
            for fi7 := 0; fi7 < 3; fi7++ {
                if (42 >= 32767) {
                    if false {
                        v8 := v2[915]
                        _ = v8
                    }
                    fmt.Println(fi7)
                } else {
                    fmt.Println(fi6, fi6, fi6)
                    var v9 int64 = (v1 & v1)
                    _ = v9
                }
                v10 := GenS1{F0: 280, F1: -343}
                var v11 float64 = float64((-2.25 + v4))
                _ = fi7
                _ = v10
                _ = v11
            }
            if (128 > 255) {
                v12 := []int{(-(v3)), (v3 ^ -1)}
                _ = v12
            } else {
                v13 := GenS2{F0: 65535}
                var v14 bool = (42 == 18446744073709551615)
                delete(v2, fi6)
                _ = v13
                _ = v14
            }
            _ = fi6
        }
    }
    var v15 int32 = -2147483648
    var v16 uint32 = 223
    v17 := []int{}
    _ = len(v17)
    switch "\n\t" {
    case "":
        var v18 bool = (v16 == (v16 << ((4294967295) & 63)))
        _ = v18
    case "a":
        v19 := v2[v3]
        _ = v19
    }
    fmt.Println(v3)
    v1 += v1
    switch "" {
    case "":
        var v20 int = (v3 ^ (^(v3)))
        _ = v20
    }
    fmt.Println(v4, v1, v1)
    v21 := GenS1{F0: 65535, F1: 0}
    for fi22 := 0; fi22 < 10; fi22++ {
        var v23 int64 = v1
        var v24 bool = true
        fmt.Println(v23)
        _ = fi22
        _ = v23
        _ = v24
    }
    var v25 float32 = v4
    for fi26 := 0; fi26 < 9; fi26++ {
        _ = len(v17)
        var v27 uint16 = 868
        for fi28 := 0; fi28 < 7; fi28++ {
            var v29 int16 = int16(v27)
            var v30 uint32 = v16
            _ = fi28
            _ = v29
            _ = v30
        }
        _ = fi26
        _ = v27
    }
    if len(v17) >= 2 {
        v17 = v17[1:len(v17)]
    }
    if false {
        switch "a" {
        case "hello":
            fmt.Println(v4)
        case "a":
            v31 := GenS1{F0: 0, F1: 42}
            _ = v31
        case "foo bar":
            for fi32 := 0; fi32 < 16; fi32++ {
                var v33 int8 = 127
                _ = fi32
                _ = v33
            }
        default:
            v34 := GenS3{F0: (v1 / (((v1 | (v1 * (^(v1))))) | 1)), F1: -1, F2: 739}
            _ = v34
        }
        var v35 int32 = (v15 << ((255) & 63))
        _ = v35
    }
    if true {
        v36 := GenS1{F0: 255, F1: 267}
        for fi37 := 0; fi37 < 12; fi37++ {
            if (false || (fi37 > (v3 | v3))) {
                v1 = 1
                v2[v3] = ((v1 >> ((567) & 63)) > -1)
            } else {
                for fi38 := 0; fi38 < 20; fi38++ {
                    v39 := GenS1{F0: 65535, F1: -480}
                    var v40 int = fi38
                    var v41 uint8 = 2
                    _ = fi38
                    _ = v39
                    _ = v40
                    _ = v41
                }
                var v42 bool = true
                v2[(v3 * fi37)] = ((!((("hello" + (("" + "hello") + "\n\t")) < ("foo bar" + "\n\t")) && false)) || (32767 != 32767))
                _ = v42
            }
            v4 += v25
            _ = len(v17)
            _ = fi37
        }
        var v43 int16 = 1000
        _ = v36
        _ = v43
    } else {
        _ = len(v2)
    }
    switch (-(v3)) {
    case -1:
        var v44 int64 = (v1 + v1)
        _ = v44
    default:
        _ = len(v2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v25
    _ = v3
    _ = v4
}
