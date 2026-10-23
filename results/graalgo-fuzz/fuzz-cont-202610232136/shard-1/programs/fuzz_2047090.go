package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 int32; F1 uint; F2 bool }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint16; F1 uint8 }
func genFn0(p0 uint16, p1 bool) uint {
    var v6 int64 = -9223372036854775808
    v6 = (v6 / ((v6) | 1))
    var v7 int = -336
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 790
}
func genFn1(p0 int64) bool {
    for fi8 := 0; fi8 < 19; fi8++ {
        for fi9 := 0; fi9 < 5; fi9++ {
            fmt.Println(fi8, fi9)
            _ = fi9
        }
        p0 += p0
        _ = fi8
    }
    v10 := []int{-1, -654, 4294967295, 128}
    v10 = append(v10, 843)
    switch "\n\t" {
    case "foo bar":
        v11 := GenS1{F0: 42, F1: -438}
        _ = v11
    case "\n\t":
        fmt.Println(p0)
    case "a":
        if true {
            v12 := GenS1{F0: -1, F1: 690}
            var v13 int64 = p0
            _ = v12
            _ = v13
        } else {
            fmt.Println(p0)
        }
    default:
        switch (("hello" + "\n\t") + "hello") {
        case "foo bar":
            p0 = (p0 - (p0 + p0))
        default:
            v14 := GenS1{F0: 42, F1: -280}
            _ = v14
        }
    }
    _ = p0
    _ = v10
    return false
}

func main() {
    var v1 float64 = (3.14159 * 0.1)
    var v2 int16 = 0
    v3 := make(map[int]string)
    var v4 uint8 = 127
    v5 := make(map[string]string)
    if ((v2 << ((367) & 63)) >= (v2 ^ (v2 & (v2 % ((v2) | 1))))) {
        if (!((!true) || (v1 < v1))) {
            v2 = (v2 << ((259) & 63))
            for fi15 := 0; fi15 < 4; fi15++ {
                for fi16 := 0; fi16 < 7; fi16++ {
                    var v17 float64 = (v1 - float64((3.14159 * 1e10)))
                    var v18 int = 4294967295
                    var v19 int32 = -1
                    _ = fi16
                    _ = v17
                    _ = v18
                    _ = v19
                }
                v20 := []int{799, (^(fi15)), (fi15 % (((-(fi15))) | 1)), 9223372036854775807}
                v2 -= (^(v2))
                _ = fi15
                _ = v20
            }
            switch "" {
            case "\n\t":
                var v21 int16 = int16(v4)
                _ = v21
            case "hello":
                switch "hello" {
                case "":
                    v22 := GenS1{F0: 32767, F1: v2}
                    _ = v22
                case "a":
                    v1 -= (((float64((1e10 - 3.14159)) * float64(3.14159)) + ((3.14159 / v1) / (3.14159 + -2.25))) - ((v1 / v1) / 1.5))
                case "hello":
                    fmt.Println(v1, v4)
                default:
                    for fi23 := 0; fi23 < 10; fi23++ {
                        for fi24 := 0; fi24 < 7; fi24++ {
                            switch "\n\t" {
                            case "foo bar":
                                v25 := GenS1{F0: v2, F1: v2}
                                _ = v25
                            case "":
                                v2 -= v2
                            case "a":
                                if (!true) {
                                    if true {
                                        var v26 int = (fi24 + (fi24 % (((fi24 << ((1) & 63))) | 1)))
                                        v27 := []int{}
                                        v28 := GenS2{F0: -935, F1: 255, F2: (2147483647 == 2147483648)}
                                        _ = v26
                                        _ = v27
                                        _ = v28
                                    }
                                    var v29 uint64 = 630
                                    v1 = v1
                                    switch ("a" + "hello") {
                                    case "a":
                                        var v30 int16 = v2
                                        _ = v30
                                    case "hello":
                                        v4 = v4
                                    case "":
                                        if (!(false || true)) {
                                            v31 := []int{128, fi24, fi23}
                                            var v32 int = fi24
                                            v33 := GenS1{F0: (v2 - 439), F1: -1}
                                            _ = v31
                                            _ = v32
                                            _ = v33
                                        } else {
                                            if false {
                                                var v34 float64 = ((v1 + v1) + float64(-2.25))
                                                _ = v34
                                            }
                                            if false {
                                                var v35 uint = genFn0(32767, false)
                                                v36 := GenS1{F0: (v2 | 42), F1: v2}
                                                if false {
                                                    v29 -= (v29 | (v29 - (v29 + (^(v29)))))
                                                    var v37 bool = false
                                                    v2 = (v2 - v2)
                                                    _ = v37
                                                }
                                                _ = v35
                                                _ = v36
                                            }
                                            v38 := v5["hello"]
                                            _ = v38
                                        }
                                    default:
                                        v4 -= (v4 / (((v4 - (v4 << ((4294967295) & 63)))) | 1))
                                    }
                                    v39 := GenS3{F0: 42}
                                    _ = v29
                                    _ = v39
                                }
                            }
                            if true {
                                if true {
                                    delete(v3, fi23)
                                    var v40 uint = genFn0(42, (0 >= 1))
                                    _ = v40
                                }
                                var v41 uint = genFn0(967, (42 != 1))
                                var v42 bool = (0 <= 81)
                                _ = v41
                                _ = v42
                            }
                            var v43 uint64 = 4294967295
                            _ = fi24
                            _ = v43
                        }
                        v44 := []int{}
                        _ = fi23
                        _ = v44
                    }
                }
            default:
                v2 = (-(v2))
            }
        }
        v45 := GenS4{F0: 42, F1: 2}
        _ = v45
    } else {
        v2 = (v2 >> ((626) & 63))
    }
    if (-2.25 == 1e10) {
        var v46 int8 = -19
        v47 := []int{42, 9223372036854775807}
        _ = v46
        _ = v47
    } else {
        _ = len(v3)
        var v48 float32 = 3.14159
        _ = v48
    }
    v49 := GenS2{F0: -2147483648, F1: 18446744073709551615, F2: true}
    var v50 int8 = -1
    v1 = 1e10
    if (!false) {
        fmt.Println(v2)
    }
    v50 = (^(v50))
    v2 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
}
