package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 uint }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 int16; F1 int }
type GenS4 struct { F0 string; F1 uint }
func genFn0() int {
    var v5 uint = 18446744073709551615
    v5 -= 265
    for fi6 := 0; fi6 < 1; fi6++ {
        var v7 bool = true
        switch (^(fi6)) {
        case 362:
            for fi8 := 0; fi8 < 14; fi8++ {
                var v9 int32 = 2147483647
                v10 := []int{}
                if (!v7) {
                    v11 := GenS1{F0: (v5 << ((v5) & 63)), F1: 200, F2: (v5 >> (((v5 / (((v5 * (v5 - v5))) | 1))) & 63))}
                    v9 = (v9 * (v9 >> ((18446744073709551615) & 63)))
                    var v12 int32 = v9
                    _ = v11
                    _ = v12
                } else {
                    if v7 {
                        v9 = v9
                        v9 = -1
                        v7 = ((!(v5 <= v5)) || v7)
                    } else {
                        v13 := GenS1{F0: v5, F1: 18446744073709551615, F2: v5}
                        if (!(true && v7)) {
                            switch (fi8 / ((fi6) | 1)) {
                            case -263:
                                _ = len(v10)
                            case -249:
                                v7 = (false && ((fi8 + (fi8 - fi8)) < (fi8 * fi6)))
                            }
                            v7 = true
                            if len(v10) > 0 {
                                v10[0] = (fi6 | (fi8 << (((v5 - (^(v5)))) & 63)))
                            }
                        } else {
                            if v7 {
                                if v7 {
                                    v9 = -1
                                }
                            }
                            fmt.Println(v5)
                        }
                        v14 := GenS2{F0: 842, F1: 32767}
                        v10 = append(v10, 981)
                        _ = v13
                        _ = v14
                    }
                }
                if len(v10) >= 2 {
                    v10 = v10[1:len(v10)]
                }
                _ = fi8
                _ = v10
                _ = v9
            }
        case -2147483648:
            fmt.Println(v7)
        case 32768:
            var v15 string = "a"
            _ = v15
        default:
            v16 := GenS1{F0: (v5 | (v5 << ((8) & 63))), F1: 0, F2: 0}
            _ = v16
        }
        _ = fi6
        _ = v7
    }
    fmt.Println(v5)
    var v17 uint16 = 32767
    _ = v17
    _ = v5
    return 0
}
func genFn1() int16 {
    var v18 int8 = -115
    fmt.Println(v18)
    v19 := GenS3{F0: 32767, F1: -366}
    fmt.Println(v18)
    var v20 uint32 = 276
    _ = v18
    _ = v19
    _ = v20
    return -1
}

func main() {
    v1 := make(map[uint32]int)
    v2 := []string{"\n\t", (("\n\t" + "foo bar") + ("" + "a")), ("hello" + "a"), "foo bar"}
    v3 := []int{}
    var v4 int32 = 0
    switch "hello" {
    case "hello":
        var v21 string = "\n\t"
        _ = v21
    case "":
        fmt.Println(v4)
    default:
        v3 = append(v3, 1)
    }
    fmt.Println(v4)
    v4 -= (v4 >> ((4294967295) & 63))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v22 int8 = int8(v4)
    var v23 float32 = ((0.1 + 1.5) * 0.1)
    var v24 int8 = (v22 | (v22 << ((911) & 63)))
    var v25 bool = true
    switch 2147483647 {
    case -805:
        var v26 uint8 = 0
        _ = v26
    case 0:
        if v25 {
            var v27 int8 = (v24 / ((-98) | 1))
            for fi28 := 0; fi28 < 7; fi28++ {
                v29 := GenS2{F0: 42, F1: 139}
                var v30 int = 2147483647
                fmt.Println(v27, v24)
                _ = fi28
                _ = v29
                _ = v30
            }
            _ = v27
        }
    case -312:
        v31 := GenS3{F0: 32767, F1: 32767}
        _ = v31
    default:
        switch ("hello" + ("" + "hello")) {
        case "foo bar":
            switch "" {
            case "foo bar":
                v32 := v1[4294967295]
                _ = v32
            case "a":
                for fi33 := 0; fi33 < 4; fi33++ {
                    v34 := GenS4{F0: "hello", F1: 255}
                    switch ("hello" + "hello") {
                    case "":
                        var v35 uint8 = 0
                        _ = v35
                    case "\n\t":
                        var v36 uint64 = 2
                        _ = v36
                    }
                    for fi37 := 0; fi37 < 11; fi37++ {
                        v25 = ((((0.0 / 3.14159) + 0.1) / (0.0 / v23)) == (v23 + v23))
                        v4 = 1
                        fmt.Println(v25, v24, fi33)
                        _ = fi37
                    }
                    _ = fi33
                    _ = v34
                }
            default:
                v25 = true
            }
        case "\n\t":
            v25 = v25
        }
    }
    v23 = v23
    v4 += (v4 + 1)
    for fi38 := 0; fi38 < 1; fi38++ {
        v4 = v4
        _ = fi38
    }
    fmt.Println(v25, v4, v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v4
}
