package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8; F1 uint16; F2 string }
type GenS3 struct { F0 uint32; F1 int32 }
func genFn0(p0 bool, p1 bool) uint16 {
    var v5 float64 = 1e10
    fmt.Println(p1, v5, p1)
    _ = p0
    _ = p1
    _ = v5
    return 32767
}
func genFn1(p0 int, p1 uint8) int {
    if false {
        var v6 int64 = 177
        _ = v6
    } else {
        switch (p0 + (p0 ^ (-(p0)))) {
        case 4294967295:
            switch 1 {
            case 127:
                switch "foo bar" {
                case "\n\t":
                    for fi7 := 0; fi7 < 16; fi7++ {
                        for fi8 := 0; fi8 < 14; fi8++ {
                            var v9 int = -759
                            fmt.Println(fi7, fi8)
                            _ = fi8
                            _ = v9
                        }
                        p0 -= (fi7 - 1)
                        _ = fi7
                    }
                case "foo bar":
                    for fi10 := 0; fi10 < 5; fi10++ {
                        p1 = (p1 << ((4294967295) & 63))
                        v11 := GenS1{F0: -9223372036854775808}
                        var v12 uint8 = 42
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                case "a":
                    v13 := []int{p0, p0, p0, p0}
                    _ = v13
                default:
                    v14 := GenS1{F0: 0}
                    _ = v14
                }
            case 128:
                var v15 float64 = 0.0
                _ = v15
            case 169:
                var v16 int8 = 100
                _ = v16
            }
        case 9223372036854775807:
            v17 := []int{(p0 ^ (p0 + -422)), p0, (p0 ^ p0)}
            _ = v17
        case 0:
            var v18 float32 = 1e10
            _ = v18
        default:
            v19 := []int{p0}
            _ = v19
        }
        p0 -= p0
        v20 := GenS1{F0: -611}
        _ = v20
    }
    v21 := GenS2{F0: 42, F1: 401, F2: "a"}
    var v22 uint = 2
    _ = p0
    _ = p1
    _ = v21
    _ = v22
    return (^(p0))
}

func main() {
    var v1 uint64 = 341
    var v2 uint16 = 32767
    var v3 uint8 = 0
    var v4 float64 = 3.14159
    v2 += 260
    for fi23 := 0; fi23 < 8; fi23++ {
        v24 := []int{fi23}
        fmt.Println(v1)
        switch "\n\t" {
        case "foo bar":
            var v25 int16 = 32767
            _ = v25
        case "a":
            _ = len(v24)
        case "hello":
            v4 += v4
        default:
            var v26 float32 = ((1.5 * 0.0) - (-2.25 + 0.0))
            _ = v26
        }
        _ = fi23
        _ = v24
    }
    if (0 == -2147483648) {
        for fi27 := 0; fi27 < 9; fi27++ {
            var v28 int64 = -9223372036854775808
            for fi29 := 0; fi29 < 16; fi29++ {
                v30 := GenS3{F0: 505, F1: -1}
                v31 := []int{fi27}
                v4 += v4
                _ = fi29
                _ = v30
                _ = v31
            }
            _ = fi27
            _ = v28
        }
        switch genFn1(127, (^(v3))) {
        case 128:
            if (false || (!(2147483647 == 0))) {
                switch "" {
                case "foo bar":
                    v4 = 0.1
                case "hello":
                    var v32 uint8 = v3
                    _ = v32
                }
            }
        case 412:
            v33 := GenS3{F0: 778, F1: 42}
            _ = v33
        default:
            if (false && true) {
                v4 = 0.1
                var v34 int = genFn1(684, v3)
                switch "" {
                case "\n\t":
                    var v35 uint = uint(v3)
                    _ = v35
                case "hello":
                    var v36 int16 = 32767
                    _ = v36
                case "":
                    var v37 int16 = 1000
                    _ = v37
                }
                _ = v34
            }
        }
    }
    fmt.Println(v4)
    v4 = v4
    var v38 uint32 = 2
    fmt.Println(v4, v38, v38)
    fmt.Println(v2)
    if true {
        v4 -= float64(1e10)
        if true {
            fmt.Println(v38, v4)
            fmt.Println(v1, v1)
        } else {
            switch ("hello" + "") {
            case "hello":
                fmt.Println(v1, v2)
            }
            v39 := []int{42}
            _ = v39
        }
        var v40 uint64 = (v1 / (((v1 * (v1 - (v1 ^ (v1 / ((v1) | 1)))))) | 1))
        _ = v40
    }
    var v41 int16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v41
}
