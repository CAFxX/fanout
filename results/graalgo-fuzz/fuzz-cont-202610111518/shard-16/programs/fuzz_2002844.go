package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int64; F1 int32 }
type GenS3 struct { F0 uint32; F1 int8 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int16; F1 uint64; F2 uint64 }
func genFn0() int32 {
    for fi6 := 0; fi6 < 11; fi6++ {
        var v7 uint16 = 0
        _ = fi6
        _ = v7
    }
    var v8 float64 = ((3.14159 + 0.0) * 1e10)
    var v9 uint16 = 42
    switch "a" {
    case "":
        v10 := []int{42, 9223372036854775807, -9223372036854775808, 128}
        _ = v10
    case "\n\t":
        v11 := []int{}
        _ = v11
    case "foo bar":
        v12 := []int{}
        _ = v12
    }
    _ = v8
    _ = v9
    return -608
}

func main() {
    var v1 float64 = 1e10
    var v2 uint64 = 1
    var v3 uint64 = v2
    var v4 string = ""
    v5 := make(map[uint32]string)
    var v13 bool = true
    v14 := GenS1{F0: 42}
    fmt.Println(v2, v3)
    var v15 int16 = 0
    fmt.Println(v2, v15)
    var v16 int32 = 42
    v3 = v3
    var v17 int64 = 997
    v15 = 42
    v16 += genFn0()
    switch ((v4 + v4) + v4) {
    case "hello":
        for fi18 := 0; fi18 < 4; fi18++ {
            v19 := GenS2{F0: v17, F1: (v16 & (v16 >> ((0) & 63)))}
            fmt.Println(v1, v1, v15)
            _ = fi18
            _ = v19
        }
    case "\n\t":
        if (true && (166 < 270)) {
            for fi20 := 0; fi20 < 13; fi20++ {
                if v13 {
                    switch 660 {
                    case -2147483648:
                        for fi21 := 0; fi21 < 2; fi21++ {
                            v22 := []int{fi21, (fi20 >> ((4294967295) & 63)), (fi21 % ((int(v16)) | 1)), (fi20 / (((fi20 & (fi21 << ((868) & 63)))) | 1))}
                            var v23 int64 = v17
                            _ = fi21
                            _ = v22
                            _ = v23
                        }
                    case 179:
                        for fi24 := 0; fi24 < 15; fi24++ {
                            v16 += genFn0()
                            _ = fi24
                        }
                    case 4294967295:
                        for fi25 := 0; fi25 < 17; fi25++ {
                            v16 = genFn0()
                            switch v4 {
                            case "hello":
                                if v13 {
                                    fmt.Println(fi20)
                                }
                            case "a":
                                v2 -= (v3 | (^(v2)))
                            case "\n\t":
                                fmt.Println(v4)
                            }
                            if (4294967295 == 146) {
                                v26 := []int{fi25, 506}
                                _ = v26
                            }
                            if false {
                                v27 := GenS1{F0: 65535}
                                _ = v27
                            }
                            v28 := GenS3{F0: 2147483648, F1: -128}
                            var v29 uint = 2
                            _ = fi25
                            _ = v28
                            _ = v29
                        }
                    default:
                        var v30 int8 = 71
                        _ = v30
                    }
                    if (1e10 != (v1 * v1)) {
                        v3 = 18446744073709551615
                        if ((fi20 + -2147483648) < int(v16)) {
                            fmt.Println(v13)
                            _ = len(v5)
                            var v31 float64 = v1
                            _ = v31
                        } else {
                            v16 += (v16 & int32(fi20))
                        }
                    }
                    switch (v4 + ("a" + v4)) {
                    case "":
                        var v32 uint8 = 0
                        _ = v32
                    }
                }
                _ = fi20
            }
            v33 := GenS4{F0: 255}
            if (0.1 == 0.1) {
                v2 = (v2 * (v2 ^ (v2 ^ (v3 >> ((42) & 63)))))
                var v34 string = ""
                _ = v34
            }
            _ = v33
        } else {
            if (127 < -9223372036854775808) {
                if v13 {
                    if (255 < 127) {
                        v35 := GenS3{F0: 729, F1: 127}
                        for fi36 := 0; fi36 < 15; fi36++ {
                            var v37 int = (fi36 + fi36)
                            var v38 uint16 = 201
                            v39 := []int{(fi36 * v37)}
                            _ = fi36
                            _ = v37
                            _ = v38
                            _ = v39
                        }
                        _ = v35
                    } else {
                        var v40 bool = (v13 && v13)
                        fmt.Println(v40, v16)
                        _ = v40
                    }
                    v41 := GenS5{F0: v15, F1: (v3 + (^(v3))), F2: (v2 & 745)}
                    _ = v41
                }
            }
            v42 := []int{}
            _ = v42
        }
    case "":
        v17 = 213
    default:
        v5[5] = (v4 + v4)
    }
    v43 := GenS4{F0: 668}
    var v44 int32 = 2147483647
    switch 762 {
    case 2147483647:
        v44 = 2147483647
    case 128:
        for fi45 := 0; fi45 < 12; fi45++ {
            var v46 float32 = (0.0 - -2.25)
            var v47 uint = 52
            _ = fi45
            _ = v46
            _ = v47
        }
    case -490:
        v48 := GenS2{F0: -9223372036854775808, F1: (v44 & genFn0())}
        _ = v48
    }
    if v13 {
        v1 = float64(1e10)
        v2 += (v3 & 2)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v5
}
