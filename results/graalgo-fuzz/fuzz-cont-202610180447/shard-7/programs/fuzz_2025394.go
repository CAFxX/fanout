package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint8; F1 bool }
type GenS3 struct { F0 int; F1 int32; F2 int64 }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 bool; F1 uint64 }
type GenS6 struct { F0 int16; F1 int8 }
func genFn0() string {
    var v7 uint16 = 1
    v8 := GenS1{F0: 2}
    _ = v7
    _ = v8
    return ("\n\t" + "foo bar")
}
func genFn1(p0 uint16) int8 {
    v9 := []int{1, -1, 127, 127}
    for fi10 := 0; fi10 < 3; fi10++ {
        if len(v9) > 0 {
            v9[0] = (fi10 + 42)
        }
        if (!false) {
            for fi11 := 0; fi11 < 13; fi11++ {
                p0 -= p0
                var v12 uint32 = 764
                var v13 int16 = -636
                if ((18446744073709551615 < 888) || (-93 > 63)) {
                    fmt.Println(fi10)
                }
                _ = fi11
                _ = v12
                _ = v13
            }
            if false {
                if (("foo bar" + "") == "") {
                    var v14 float64 = 1e10
                    fmt.Println(p0, p0)
                    _ = v14
                }
                p0 += 32767
            }
        }
        for fi15 := 0; fi15 < 15; fi15++ {
            p0 += (p0 * (^(p0)))
            v16 := GenS2{F0: 127, F1: (((float64((-2.25 / 1.0)) - 0.0) * ((-2.25 * (1e10 - 3.14159)) - 3.14159)) >= (0.0 + (3.14159 * (0.1 * ((((-2.25 - float64(1e10)) + 3.14159) + ((0.0 * (1.5 / 1.0)) + 1e10)) - ((1.5 * 1e10) * float64(0.0)))))))}
            if false {
                if true {
                    var v17 uint32 = 82
                    _ = v17
                }
                if true {
                    if false {
                        p0 += p0
                    } else {
                        v18 := GenS2{F0: 127, F1: false}
                        p0 -= p0
                        var v19 int64 = 9223372036854775807
                        if true {
                            v20 := GenS1{F0: 505}
                            v9 = append(v9, fi15)
                            _ = v20
                        }
                        var v21 float64 = (3.14159 - 1e10)
                        _ = v18
                        _ = v19
                        _ = v21
                    }
                    fmt.Println(fi15, fi15, fi15)
                }
            } else {
                p0 -= (p0 >> ((255) & 63))
            }
            _ = fi15
            _ = v16
        }
        _ = fi10
    }
    p0 -= (^(p0))
    p0 = (p0 | (^(p0)))
    if false {
        v22 := GenS3{F0: 128, F1: -2147483648, F2: 42}
        _ = v22
    } else {
        p0 += 0
        for fi23 := 0; fi23 < 7; fi23++ {
            p0 = 2
            p0 = 32767
            _ = fi23
        }
    }
    _ = p0
    _ = v9
    return -123
}

func main() {
    var v1 int8 = -17
    var v2 float64 = -2.25
    var v3 int32 = 891
    v4 := make(map[string]int)
    v5 := make(map[string]string)
    v6 := make(map[int]string)
    for fi24 := 0; fi24 < 20; fi24++ {
        switch (((("\n\t" + "\n\t") + "foo bar") + "foo bar") + "foo bar") {
        case "hello":
            var v25 uint8 = 197
            _ = v25
        case "foo bar":
            for fi26 := 0; fi26 < 3; fi26++ {
                var v27 uint8 = 16
                _ = fi26
                _ = v27
            }
        }
        _ = fi24
    }
    for fi28 := 0; fi28 < 16; fi28++ {
        v29 := GenS4{F0: -32768}
        var v30 float32 = 1e10
        _ = fi28
        _ = v29
        _ = v30
    }
    var v31 uint = 18446744073709551615
    v32 := GenS2{F0: 255, F1: (42 == 255)}
    v33 := []int{}
    if (42 == 0) {
        v3 = -427
        v2 += (0.1 / v2)
    }
    v34 := GenS5{F0: false, F1: 2}
    for fi35 := 0; fi35 < 1; fi35++ {
        switch 32768 {
        case 4294967295:
            for fi36 := 0; fi36 < 16; fi36++ {
                v1 += genFn1(255)
                if (((fi36 % ((-654) | 1)) < 32768) || ((v3 >> (((v31 & 542)) & 63)) != v3)) {
                    var v37 uint64 = 42
                    _ = v37
                }
                _ = fi36
            }
        case -1:
            fmt.Println(v3, fi35)
        }
        v2 = 1.5
        _ = fi35
    }
    var v38 float64 = 3.14159
    for fi39 := 0; fi39 < 18; fi39++ {
        fmt.Println(v1, v31, v3)
        _ = fi39
    }
    v40 := GenS6{F0: 32767, F1: 100}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v38:", v38)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v38
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
