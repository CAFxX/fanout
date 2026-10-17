package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 int64; F1 uint; F2 int64 }
type GenS3 struct { F0 uint32; F1 string }
type GenS4 struct { F0 int64; F1 uint64 }
func genFn0(p0 int32) int {
    p0 = (p0 | p0)
    p0 += (p0 << ((2) & 63))
    for fi4 := 0; fi4 < 11; fi4++ {
        v5 := []int{fi4, fi4, (^(fi4)), 32767}
        _ = fi4
        _ = v5
    }
    _ = p0
    return 128
}

func main() {
    var v1 uint8 = 17
    var v2 uint = 18446744073709551615
    var v3 int32 = 1
    var v6 bool = (80 == 73)
    var v7 uint64 = uint64(v2)
    if (!(v6 || v6)) {
        var v8 int32 = (-(v3))
        _ = v8
    } else {
        v9 := []int{1, 485, 32768}
        _ = v9
    }
    fmt.Println(v1)
    v10 := []int{9223372036854775807, 32767, -1, 42}
    var v11 uint8 = v1
    for fi12 := 0; fi12 < 6; fi12++ {
        switch fi12 {
        case 42:
            v6 = false
        case 32768:
            fmt.Println(v11, fi12, v6)
        }
        var v13 int8 = 105
        v7 -= v7
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 5; fi14++ {
        if true {
            for fi15 := 0; fi15 < 17; fi15++ {
                v7 = (v7 - v7)
                var v16 bool = (!(!true))
                if v6 {
                    fmt.Println(v16, v1, v2)
                    var v17 string = ("foo bar" + "a")
                    _ = v17
                } else {
                    v18 := GenS1{F0: (^(v2)), F1: (!((v7 >= (^(v7))) || false))}
                    switch "\n\t" {
                    case "\n\t":
                        v6 = v6
                    case "":
                        if v16 {
                            if (!true) {
                                var v19 int64 = -745
                                var v20 uint32 = 2
                                v6 = true
                                _ = v19
                                _ = v20
                            }
                        } else {
                            var v21 int = fi14
                            _ = v21
                        }
                        v22 := GenS2{F0: -1, F1: v2, F2: -9223372036854775808}
                        var v23 uint64 = 1
                        _ = v22
                        _ = v23
                    default:
                        var v24 int64 = 999
                        _ = v24
                    }
                    if (255 != 2) {
                        v2 -= (v2 ^ v2)
                    } else {
                        _ = len(v10)
                        var v25 float64 = ((1e10 + 0.0) / 0.1)
                        _ = v25
                    }
                    _ = v18
                }
                fmt.Println(fi15, v11, v1)
                v3 += (^(v3))
                _ = fi15
                _ = v16
            }
            v10 = append(v10, fi14)
            v11 -= 200
        }
        var v26 uint16 = 2
        _ = fi14
        _ = v26
    }
    switch "hello" {
    case "":
        for fi27 := 0; fi27 < 20; fi27++ {
            fmt.Println(v7)
            _ = fi27
        }
    case "foo bar":
        switch 128 {
        case 127:
            var v28 bool = (v1 != (v1 | v11))
            _ = v28
        case 569:
            switch genFn0(2147483647) {
            case 32767:
                v1 += uint8(v2)
            }
        default:
            v29 := GenS2{F0: 9223372036854775807, F1: v2, F2: -9223372036854775808}
            _ = v29
        }
    case "hello":
        fmt.Println(v3, v6, v2)
    }
    for fi30 := 0; fi30 < 1; fi30++ {
        v31 := GenS2{F0: int64(v2), F1: 255, F2: -1}
        if (!((!(!("\n\t" >= (("\n\t" + "\n\t") + ("hello" + ("foo bar" + ((("" + "\n\t") + "a") + ""))))))) || (0 != 0))) {
            if v6 {
                _ = len(v10)
            } else {
                v6 = false
                for fi32 := 0; fi32 < 5; fi32++ {
                    switch "hello" {
                    case "":
                        var v33 string = "a"
                        _ = v33
                    case "foo bar":
                        if v6 {
                            var v34 uint16 = 42
                            _ = v34
                        } else {
                            v35 := GenS3{F0: 577, F1: ("foo bar" + "hello")}
                            _ = v35
                        }
                    case "hello":
                        v36 := GenS4{F0: -1, F1: (v7 * v7)}
                        _ = v36
                    }
                    var v37 uint64 = (v7 - v7)
                    _ = len(v10)
                    _ = fi32
                    _ = v37
                }
                var v38 int64 = 42
                _ = v38
            }
            if true {
                fmt.Println(v6, v3, fi30)
                if (("\n\t" + "hello") <= "foo bar") {
                    switch "hello" {
                    case "foo bar":
                        var v39 int64 = -1
                        _ = v39
                    case "":
                        v40 := GenS4{F0: -543, F1: v7}
                        _ = v40
                    }
                }
                v6 = v6
            }
        }
        var v41 uint = 465
        _ = fi30
        _ = v31
        _ = v41
    }
    for fi42 := 0; fi42 < 18; fi42++ {
        fmt.Println(v11, v2)
        var v43 string = "\n\t"
        v3 = (v3 * (v3 | v3))
        _ = fi42
        _ = v43
    }
    var v44 int16 = 1000
    if len(v10) > 0 {
        v10[0] = 1
    }
    v6 = (v44 < v44)
    for fi45 := 0; fi45 < 1; fi45++ {
        var v46 int32 = (v3 / (((^(v3))) | 1))
        var v47 uint64 = v7
        var v48 int64 = 1
        _ = fi45
        _ = v46
        _ = v47
        _ = v48
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v44:", v44)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v44
    _ = v6
    _ = v7
}
