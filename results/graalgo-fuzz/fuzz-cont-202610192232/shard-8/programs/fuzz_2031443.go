package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint8 }
func genFn0(p0 int, p1 int64) int16 {
    v7 := []int{p0, (p0 - 126), (p0 * (p0 + p0)), (p0 * p0)}
    var v8 float32 = (0.1 + 1.5)
    p1 += p1
    p0 = (^(p0))
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 118
}
func genFn1(p0 uint, p1 uint64) uint8 {
    var v9 uint32 = 653
    v10 := GenS1{F0: 323, F1: 2147483647}
    v11 := GenS1{F0: uint16(p0), F1: 1}
    var v12 uint8 = 2
    var v13 uint16 = 1
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v9
    return (v12 % (((v12 / ((v12) | 1))) | 1))
}

func main() {
    var v1 int16 = 42
    v2 := make(map[uint32]int)
    v3 := make(map[int]int)
    v4 := make(map[int]int)
    var v5 uint16 = 255
    v6 := []string{"hello", "", "foo bar"}
    for fi14 := 0; fi14 < 2; fi14++ {
        var v15 int16 = (^(v1))
        v5 += v5
        fmt.Println(fi14, v1)
        _ = fi14
        _ = v15
    }
    v5 -= (^(v5))
    v16 := GenS1{F0: 0, F1: -1}
    switch "foo bar" {
    case "a":
        for fi17 := 0; fi17 < 8; fi17++ {
            v18 := GenS2{F0: "\n\t"}
            _ = fi17
            _ = v18
        }
    case "hello":
        var v19 int64 = -1
        _ = v19
    }
    if (!true) {
        fmt.Println(v5, v1)
    }
    var v20 uint16 = (v5 * 1)
    for fi21 := 0; fi21 < 9; fi21++ {
        switch "" {
        case "foo bar":
            fmt.Println(v5)
        }
        if true {
            switch ((("" + ("hello" + "\n\t")) + "foo bar") + ("hello" + "hello")) {
            case "foo bar":
                v1 -= (v1 << ((1) & 63))
            case "a":
                for fi22 := 0; fi22 < 12; fi22++ {
                    v23 := GenS1{F0: 117, F1: (fi21 >> ((255) & 63))}
                    _ = fi22
                    _ = v23
                }
            case "\n\t":
                v24 := GenS3{F0: genFn1(158, 2)}
                _ = v24
            }
            v2[2147483647] = fi21
        } else {
            if false {
                if (false && (fi21 != (fi21 - fi21))) {
                    for fi25 := 0; fi25 < 6; fi25++ {
                        var v26 uint = 255
                        _ = fi25
                        _ = v26
                    }
                    var v27 int32 = 1
                    fmt.Println(v27, v5)
                    _ = v27
                }
                if (42 == 0) {
                    for fi28 := 0; fi28 < 13; fi28++ {
                        if true {
                            var v29 bool = false
                            _ = len(v6)
                            v29 = (false && v29)
                            _ = v29
                        }
                        v30 := GenS1{F0: (v5 - 255), F1: fi21}
                        v31 := v3[fi21]
                        _ = fi28
                        _ = v30
                        _ = v31
                    }
                    v32 := GenS3{F0: 142}
                    _ = v32
                }
            } else {
                var v33 int = (-(fi21))
                var v34 uint = 42
                _ = v33
                _ = v34
            }
            v35 := GenS3{F0: 2}
            _ = v35
        }
        _ = fi21
    }
    switch -1 {
    case 469:
        v3[1] = 0
    case 128:
        fmt.Println(v1, v1, v5)
    }
    v1 = (v1 * v1)
    fmt.Println(v1)
    if (!true) {
        if false {
            v5 -= 65535
            v20 = v20
        }
    }
    var v36 uint8 = genFn1(2, uint64(v20))
    if len(v6) > 0 {
        v6[0] = ("a" + "\n\t")
    }
    v2[2] = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
}
