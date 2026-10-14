package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 bool; F1 int8 }
type GenS4 struct { F0 int64; F1 uint8; F2 uint }
type GenS5 struct { F0 bool; F1 uint16 }
type GenS6 struct { F0 uint64 }
func genFn0(p0 uint16, p1 int8, p2 uint) uint32 {
    switch -9223372036854775808 {
    case 32767:
        v5 := []int{}
        _ = v5
    case 2147483647:
        var v6 uint64 = 4294967295
        _ = v6
    case 128:
        fmt.Println(p1, p1)
    }
    var v7 bool = false
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 0
}
func genFn1(p0 uint16) uint {
    var v8 float64 = (1.5 * ((((((1e10 / 1.0) / 0.1) + (0.0 + (1.5 * 1e10))) / (float64(-2.25) + (1.5 / 1.5))) - (3.14159 + 0.0)) + (3.14159 * 1.5)))
    for fi9 := 0; fi9 < 19; fi9++ {
        var v10 uint = 496
        v8 = 1e10
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 2; fi11++ {
        v8 = 0.0
        v12 := []int{1, fi11, fi11, 927}
        _ = fi11
        _ = v12
    }
    _ = p0
    _ = v8
    return 0
}

func main() {
    var v1 int8 = 42
    var v2 int16 = 1000
    var v3 bool = (!true)
    v4 := make(map[uint32]bool)
    for fi13 := 0; fi13 < 9; fi13++ {
        var v14 uint32 = 906
        _ = fi13
        _ = v14
    }
    var v15 int32 = 1
    var v16 uint64 = 4294967295
    v17 := GenS1{F0: (v1 ^ v1), F1: 0}
    if true {
        var v18 int16 = (^(v2))
        v19 := []int{1, 32767, 127}
        _ = v18
        _ = v19
    }
    if (42 > 2) {
        var v20 uint8 = 58
        v2 -= 32767
        _ = v20
    } else {
        if ((2147483647 <= 4294967295) || (int64(v1) < 0)) {
            switch 42 {
            case -713:
                if (v1 < (v1 >> ((2) & 63))) {
                    v16 = v16
                    var v21 string = "hello"
                    _ = v21
                } else {
                    if ((!(((^(v1)) > v1) || ("hello" >= "foo bar"))) && v3) {
                        var v22 int16 = (v2 % ((-303) | 1))
                        switch 512 {
                        case 4294967295:
                            switch ("hello" + "foo bar") {
                            case "foo bar":
                                v22 = v22
                            }
                        }
                        _ = v22
                    }
                    fmt.Println(v16)
                }
            case 9223372036854775807:
                fmt.Println(v1, v3)
            case -9223372036854775808:
                v1 -= v1
            }
            var v23 int16 = (v2 & (v2 + v2))
            _ = v23
        }
        v15 = 0
        v2 = (v2 + -32768)
    }
    if v3 {
        var v24 float32 = -2.25
        _ = v24
    } else {
        v25 := GenS1{F0: 100, F1: 1}
        v1 = v1
        if v3 {
            switch "foo bar" {
            case "":
                if (v3 || (4294967295 > 1)) {
                    if v3 {
                        var v26 uint = genFn1(255)
                        v27 := []int{-916, 32768, 2147483647}
                        _ = v26
                        _ = v27
                    } else {
                        var v28 string = "hello"
                        _ = v28
                    }
                    var v29 int16 = v2
                    _ = v29
                }
            case "foo bar":
                if true {
                    v30 := GenS2{F0: ""}
                    _ = v30
                }
            case "hello":
                v31 := GenS1{F0: (^(v1)), F1: 2}
                _ = v31
            }
        }
        _ = v25
    }
    v32 := GenS3{F0: true, F1: 127}
    var v33 float64 = 1.5
    if v3 {
        delete(v4, 4294967295)
        v34 := GenS1{F0: (v1 << ((uint(v16)) & 63)), F1: 127}
        var v35 int8 = 0
        _ = v34
        _ = v35
    } else {
        v36 := GenS4{F0: 0, F1: 1, F2: 45}
        for fi37 := 0; fi37 < 16; fi37++ {
            v38 := GenS5{F0: true, F1: 65535}
            _ = fi37
            _ = v38
        }
        v3 = (!v3)
        _ = v36
    }
    for fi39 := 0; fi39 < 20; fi39++ {
        for fi40 := 0; fi40 < 9; fi40++ {
            v41 := GenS2{F0: (("\n\t" + ("foo bar" + "\n\t")) + "foo bar")}
            v42 := []int{int(v15)}
            v33 += (v33 + 0.0)
            _ = fi40
            _ = v41
            _ = v42
        }
        for fi43 := 0; fi43 < 12; fi43++ {
            v44 := GenS6{F0: (v16 ^ (v16 ^ v16))}
            _ = fi43
            _ = v44
        }
        _ = fi39
    }
    v45 := GenS4{F0: 0, F1: 42, F2: 0}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v4
    _ = v45
}
