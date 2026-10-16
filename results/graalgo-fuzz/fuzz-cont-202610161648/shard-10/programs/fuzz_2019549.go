package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool }
func genFn0() int {
    if (-452 > -2147483648) {
        var v5 int32 = -446
        var v6 uint = 18446744073709551615
        v6 -= v6
        _ = v5
        _ = v6
    }
    for fi7 := 0; fi7 < 6; fi7++ {
        if (!false) {
            v8 := []int{(fi7 * (fi7 * fi7)), fi7, (fi7 + fi7)}
            _ = v8
        } else {
            for fi9 := 0; fi9 < 12; fi9++ {
                v10 := GenS1{F0: 874, F1: 18446744073709551615}
                v11 := GenS2{F0: "hello"}
                v12 := GenS2{F0: ""}
                _ = fi9
                _ = v10
                _ = v11
                _ = v12
            }
        }
        v13 := GenS2{F0: "\n\t"}
        var v14 int64 = 0
        _ = fi7
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 20; fi15++ {
        v16 := GenS3{F0: true}
        _ = fi15
        _ = v16
    }
    return -2147483648
}
func genFn1(p0 string, p1 uint, p2 uint8) uint {
    var v17 string = p0
    if (1 <= 2) {
        var v18 uint64 = 1
        v19 := GenS1{F0: (v18 | (v18 ^ 4294967295)), F1: (p1 & 42)}
        v20 := GenS2{F0: v17}
        _ = v18
        _ = v19
        _ = v20
    }
    v21 := GenS2{F0: ("\n\t" + "")}
    var v22 int = -9223372036854775808
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v21
    _ = v22
    return 255
}

func main() {
    v1 := make(map[int]int)
    var v2 bool = false
    v3 := make(map[int]string)
    var v4 bool = v2
    fmt.Println(v2, v4)
    v2 = v4
    var v23 uint16 = 0
    v2 = false
    v24 := GenS2{F0: "\n\t"}
    fmt.Println(v4, v2)
    if (!((0.0 - ((-2.25 / ((1.5 * (1e10 - 3.14159)) + (-2.25 - 0.1))) * 3.14159)) != 0.1)) {
        for fi25 := 0; fi25 < 4; fi25++ {
            var v26 uint = 2
            for fi27 := 0; fi27 < 17; fi27++ {
                v23 -= 42
                v28 := []int{-2147483648}
                v23 = 65535
                _ = fi27
                _ = v28
            }
            _ = fi25
            _ = v26
        }
    }
    v4 = v4
    v23 = 908
    switch 128 {
    case 4294967295:
        v29 := []int{int(v23), -34, -1}
        _ = v29
    default:
        var v30 int = -421
        _ = v30
    }
    v4 = v4
    switch "hello" {
    case "":
        if v2 {
            delete(v1, genFn0())
            v31 := GenS3{F0: v4}
            if (v2 || v4) {
                fmt.Println(v2, v4, v23)
                if v4 {
                    var v32 bool = (!v4)
                    var v33 int8 = 127
                    var v34 string = ("foo bar" + "\n\t")
                    _ = v32
                    _ = v33
                    _ = v34
                } else {
                    v23 += v23
                    fmt.Println(v23, v23, v23)
                    var v35 bool = v2
                    _ = v35
                }
            } else {
                for fi36 := 0; fi36 < 20; fi36++ {
                    v23 = (v23 % (((v23 << ((633) & 63))) | 1))
                    var v37 string = "foo bar"
                    var v38 int16 = 1000
                    _ = fi36
                    _ = v37
                    _ = v38
                }
                var v39 uint64 = 418
                _ = v39
            }
            _ = v31
        } else {
            fmt.Println(v23)
            switch 2147483647 {
            case 0:
                for fi40 := 0; fi40 < 9; fi40++ {
                    var v41 uint16 = 0
                    for fi42 := 0; fi42 < 9; fi42++ {
                        var v43 int32 = 0
                        if (42 >= 935) {
                            if ("\n\t" <= "") {
                                var v44 int8 = -1
                                v45 := []int{}
                                _ = v44
                                _ = v45
                            } else {
                                v46 := []int{(fi40 * -2147483648), (fi40 - (fi40 * fi40)), genFn0()}
                                fmt.Println(v23, v41)
                                _ = v46
                            }
                        } else {
                            var v47 int64 = 1
                            _ = v47
                        }
                        _ = fi42
                        _ = v43
                    }
                    _ = fi40
                    _ = v41
                }
            }
            v2 = (42 > 487)
        }
    case "hello":
        var v48 uint8 = 179
        _ = v48
    }
    v23 -= 65535
    v49 := []int{}
    fmt.Println(v2, v2, v23)
    var v50 int = 700
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v49:", len(v49), v49)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v49
    _ = v50
}
