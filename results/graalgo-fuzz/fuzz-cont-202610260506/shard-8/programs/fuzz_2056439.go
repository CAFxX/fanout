package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string; F1 int8; F2 int32 }
func genFn0(p0 uint) int64 {
    p0 = 2
    var v6 int = 42
    v6 += v6
    v7 := GenS1{F0: 207}
    if false {
        for fi8 := 0; fi8 < 13; fi8++ {
            switch (v6 + fi8) {
            case -1:
                v6 += (^(v6))
            default:
                v9 := []int{2147483647, -2147483648, fi8}
                _ = v9
            }
            if true {
                if ((-2.25 >= (3.14159 + -2.25)) || (0 <= 598)) {
                    v10 := []int{(fi8 | (fi8 - 32768)), (v6 << ((873) & 63)), 4294967295, 1}
                    _ = v10
                } else {
                    v11 := []int{v6}
                    fmt.Println(v6, fi8)
                    _ = len(v11)
                    _ = v11
                }
                fmt.Println(p0, p0, p0)
            }
            v12 := GenS2{F0: "hello", F1: -1, F2: -2147483648}
            _ = fi8
            _ = v12
        }
    } else {
        switch (v6 >> ((p0) & 63)) {
        case 9223372036854775807:
            var v13 bool = false
            _ = v13
        case 1:
            v14 := []int{v6, 127, 410, (v6 ^ (-(v6)))}
            _ = v14
        case 797:
            var v15 float32 = (1.5 - 1e10)
            _ = v15
        }
    }
    _ = p0
    _ = v6
    _ = v7
    return 0
}
func genFn1() bool {
    var v16 uint = 42
    var v17 int = 32768
    var v18 uint16 = 2
    v18 = (^(v18))
    v18 = v18
    _ = v16
    _ = v17
    _ = v18
    return true
}

func main() {
    var v1 float32 = float32(1.5)
    v2 := make(map[uint32]int)
    var v3 int8 = -1
    var v4 bool = (true || (-1 <= -382))
    var v5 float64 = (3.14159 * 0.0)
    delete(v2, 2)
    for fi19 := 0; fi19 < 12; fi19++ {
        var v20 uint16 = 0
        _ = fi19
        _ = v20
    }
    v4 = (170 <= 42)
    _ = len(v2)
    for fi21 := 0; fi21 < 4; fi21++ {
        v22 := []int{550, fi21, (fi21 | fi21)}
        _ = fi21
        _ = v22
    }
    for fi23 := 0; fi23 < 16; fi23++ {
        for fi24 := 0; fi24 < 1; fi24++ {
            v1 += float32((v5 * v5))
            v25 := GenS1{F0: 278}
            _ = fi24
            _ = v25
        }
        if false {
            if ((42 < 2147483647) || (!(-349 == 0))) {
                v26 := []int{}
                v27 := GenS2{F0: ("hello" + "hello"), F1: (-(v3)), F2: 2147483647}
                _ = v26
                _ = v27
            } else {
                v28 := []int{}
                _ = v28
            }
            if (4294967295 <= 18446744073709551615) {
                var v29 uint8 = 255
                var v30 float64 = -2.25
                v31 := []int{(fi23 >> ((4294967295) & 63))}
                _ = v29
                _ = v30
                _ = v31
            } else {
                v1 -= v1
                for fi32 := 0; fi32 < 17; fi32++ {
                    var v33 int64 = genFn0(644)
                    v34 := GenS1{F0: 42}
                    _ = fi32
                    _ = v33
                    _ = v34
                }
            }
        }
        _ = fi23
    }
    v35 := []int{-9223372036854775808, 127, 42}
    switch "\n\t" {
    case "a":
        v3 = v3
    case "":
        v36 := GenS1{F0: 2}
        _ = v36
    }
    switch (("\n\t" + "a") + "") {
    case "foo bar":
        var v37 int8 = v3
        _ = v37
    case "\n\t":
        v1 -= (v1 + (v1 / 1.5))
    default:
        if (1 < v3) {
            v38 := GenS2{F0: "\n\t", F1: v3, F2: 1}
            _ = v38
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v4
    _ = v5
}
