package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int16; F1 uint64 }
type GenS4 struct { F0 uint }
func genFn0(p0 int64) int {
    v5 := []int{-9223372036854775808, 1, 2147483647, 2147483647}
    switch ("a" + "") {
    case "hello":
        var v6 int8 = -128
        _ = v6
    case "":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    case "foo bar":
        for fi7 := 0; fi7 < 11; fi7++ {
            for fi8 := 0; fi8 < 6; fi8++ {
                if (true && (!true)) {
                    v9 := GenS1{F0: fi8}
                    var v10 bool = true
                    _ = v10
                    _ = v9
                } else {
                    v11 := GenS1{F0: (fi8 % ((2147483647) | 1))}
                    p0 = -1
                    _ = v11
                }
                _ = fi8
            }
            v12 := GenS1{F0: 9223372036854775807}
            var v13 uint8 = 31
            _ = fi7
            _ = v12
            _ = v13
        }
    }
    v14 := GenS2{F0: 79}
    _ = p0
    _ = v14
    _ = v5
    return -1
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[int]bool)
    var v3 int8 = 1
    v4 := []uint32{2147483647}
    var v15 uint = 564
    v4 = append(v4, 894)
    for fi16 := 0; fi16 < 11; fi16++ {
        var v17 float64 = float64(1e10)
        v18 := GenS3{F0: -132, F1: 0}
        _ = fi16
        _ = v17
        _ = v18
    }
    for fi19 := 0; fi19 < 8; fi19++ {
        var v20 string = "\n\t"
        v21 := v2[32768]
        _ = fi19
        _ = v20
        _ = v21
    }
    for fi22 := 0; fi22 < 15; fi22++ {
        var v23 int16 = 1
        if false {
            v15 += v15
            v4 = append(v4, 2)
            var v24 uint = 0
            _ = v24
        } else {
            var v25 uint64 = 2
            if len(v4) > 0 {
                v4[0] = 1
            }
            v23 = -766
            _ = v25
        }
        var v26 float32 = float32(0.0)
        _ = fi22
        _ = v23
        _ = v26
    }
    var v27 uint = (v15 / ((v15) | 1))
    for fi28 := 0; fi28 < 16; fi28++ {
        for fi29 := 0; fi29 < 3; fi29++ {
            v15 += v27
            switch genFn0(279) {
            case 576:
                var v30 int16 = -32768
                _ = v30
            case 2147483647:
                v31 := GenS4{F0: v15}
                _ = v31
            case -468:
                v27 -= (v15 << ((v15) & 63))
            }
            _ = fi29
        }
        var v32 string = ""
        v27 = (v27 % ((44) | 1))
        _ = fi28
        _ = v32
    }
    var v33 uint16 = 68
    var v34 int = -83
    if false {
        if (v33 == 42) {
            switch (v34 >> ((698) & 63)) {
            case -9223372036854775808:
                var v35 uint32 = 0
                _ = v35
            case 42:
                v34 += v34
            }
            fmt.Println(v15)
            var v36 uint16 = (^(v33))
            _ = v36
        } else {
            v3 -= v3
            v3 = (v3 ^ v3)
        }
        if (true || (-1 == -9223372036854775808)) {
            var v37 int = (^(v34))
            _ = v37
        } else {
            for fi38 := 0; fi38 < 18; fi38++ {
                v39 := GenS4{F0: (v27 ^ (v27 + (^(v27))))}
                v15 = (^(v15))
                _ = fi38
                _ = v39
            }
            var v40 int64 = 329
            _ = v40
        }
        var v41 int16 = 1000
        _ = v41
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v27
    _ = v3
    _ = v33
    _ = v34
    _ = v4
}
