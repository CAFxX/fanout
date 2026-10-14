package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 bool; F1 int; F2 uint32 }
func genFn0(p0 int8, p1 bool) uint {
    if (p1 && (4294967295 < 1)) {
        switch 4294967295 {
        case 433:
            switch ("hello" + "foo bar") {
            case "a":
                p0 += p0
            }
        }
    }
    var v7 uint64 = 4294967295
    v8 := []int{}
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 1
}

func main() {
    var v1 float32 = 1e10
    var v2 int32 = 919
    var v3 bool = (!true)
    var v4 uint32 = 348
    var v5 int64 = 0
    v6 := make(map[uint32]int)
    var v9 uint16 = 2
    var v10 uint16 = (v9 ^ (v9 % ((v9) | 1)))
    fmt.Println(v3)
    switch "\n\t" {
    case "":
        for fi11 := 0; fi11 < 17; fi11++ {
            v12 := GenS1{F0: 32767}
            v13 := v6[v4]
            v9 -= 65535
            _ = fi11
            _ = v12
            _ = v13
        }
    case "a":
        v10 = v9
    case "hello":
        var v14 int16 = -811
        _ = v14
    default:
        for fi15 := 0; fi15 < 19; fi15++ {
            var v16 uint = genFn0(-1, (1.5 <= 1.5))
            for fi17 := 0; fi17 < 20; fi17++ {
                var v18 int64 = v5
                var v19 uint16 = v10
                v20 := GenS1{F0: 1}
                _ = fi17
                _ = v18
                _ = v19
                _ = v20
            }
            _ = fi15
            _ = v16
        }
    }
    v10 = (v10 - v9)
    var v21 string = "hello"
    for fi22 := 0; fi22 < 6; fi22++ {
        switch v21 {
        case "":
            var v23 string = v21
            _ = v23
        case "hello":
            v24 := []int{}
            _ = v24
        case "foo bar":
            if true {
                if true {
                    var v25 uint16 = uint16(v4)
                    var v26 string = v21
                    _ = v25
                    _ = v26
                }
                var v27 int8 = -1
                _ = v27
            } else {
                if v3 {
                    v9 = (v10 - v10)
                    switch "foo bar" {
                    case "":
                        v28 := GenS1{F0: 1000}
                        _ = v28
                    case "foo bar":
                        var v29 int = fi22
                        _ = v29
                    default:
                        if (-1 == -9223372036854775808) {
                            v4 += (v4 / (((v4 - (v4 * v4))) | 1))
                        }
                    }
                } else {
                    for fi30 := 0; fi30 < 20; fi30++ {
                        v21 = "foo bar"
                        _ = fi30
                    }
                    if (!(v3 && (4294967295 >= 18446744073709551615))) {
                        var v31 string = "\n\t"
                        if true {
                            v9 -= 761
                            switch (fi22 & 456) {
                            case -828:
                                if false {
                                    fmt.Println(v31, v2, v2)
                                    var v32 float32 = 3.14159
                                    _ = v32
                                }
                            }
                        } else {
                            for fi33 := 0; fi33 < 2; fi33++ {
                                if false {
                                    v34 := GenS2{F0: ((v4 ^ v4) <= 42), F1: (fi33 ^ 2147483647), F2: (v4 / (((^(v4))) | 1))}
                                    _ = v34
                                }
                                _ = fi33
                            }
                        }
                        var v35 int8 = 100
                        _ = v31
                        _ = v35
                    } else {
                        switch fi22 {
                        case 1:
                            var v36 int = (fi22 << ((18446744073709551615) & 63))
                            _ = v36
                        case 4294967295:
                            v37 := []int{(fi22 * (fi22 ^ (fi22 - fi22)))}
                            _ = v37
                        case -864:
                            var v38 int16 = 1000
                            _ = v38
                        default:
                            var v39 int64 = v5
                            _ = v39
                        }
                        v3 = ((v5 | (v5 << ((1) & 63))) < (v5 + (-(v5))))
                        for fi40 := 0; fi40 < 18; fi40++ {
                            v41 := GenS1{F0: 42}
                            _ = fi40
                            _ = v41
                        }
                    }
                }
                delete(v6, (v4 << ((4294967295) & 63)))
                var v42 float64 = 1.5
                _ = v42
            }
        default:
            v4 -= (v4 & 617)
        }
        var v43 int8 = -128
        _ = fi22
        _ = v43
    }
    v44 := []int{-797, 108}
    v9 += (^(v10))
    var v45 uint16 = v9
    var v46 uint = uint(v10)
    v46 = v46
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v44:", len(v44), v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v5
    _ = v6
    _ = v9
}
