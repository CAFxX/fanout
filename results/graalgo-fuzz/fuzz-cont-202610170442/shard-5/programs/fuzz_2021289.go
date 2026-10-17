package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 uint64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint8; F1 string; F2 string }
type GenS5 struct { F0 int32; F1 uint16 }
func genFn0() uint16 {
    var v4 uint16 = 255
    v5 := GenS1{F0: v4, F1: 18446744073709551615, F2: 44}
    v6 := []int{2147483647}
    _ = v4
    _ = v5
    _ = v6
    return v4
}

func main() {
    var v1 int = 42
    var v2 int8 = -1
    var v3 uint = 65535
    for fi7 := 0; fi7 < 4; fi7++ {
        if false {
            fmt.Println(fi7, v3, fi7)
            for fi8 := 0; fi8 < 20; fi8++ {
                v1 += v1
                v2 -= (v2 ^ 100)
                v1 = v1
                _ = fi8
            }
        }
        var v9 float32 = 1.5
        var v10 int32 = 2147483647
        _ = fi7
        _ = v10
        _ = v9
    }
    fmt.Println(v3)
    v1 -= v1
    v11 := GenS1{F0: 75, F1: 2, F2: 866}
    switch (("hello" + "") + "\n\t") {
    case "a":
        var v12 uint32 = 2
        _ = v12
    case "":
        if true {
            v13 := []int{0, v1}
            _ = v13
        }
    case "foo bar":
        for fi14 := 0; fi14 < 9; fi14++ {
            var v15 uint16 = genFn0()
            if ((0.1 - 1e10) <= 1e10) {
                for fi16 := 0; fi16 < 4; fi16++ {
                    v17 := GenS1{F0: 2, F1: 336, F2: 2}
                    _ = fi16
                    _ = v17
                }
                var v18 int64 = 42
                v19 := []int{(fi14 | -875), (v1 >> (((v3 >> ((4294967295) & 63))) & 63))}
                _ = v18
                _ = v19
            } else {
                v20 := GenS2{F0: 4294967295}
                v15 += v15
                v21 := GenS3{F0: 950}
                _ = v20
                _ = v21
            }
            _ = fi14
            _ = v15
        }
    }
    v22 := GenS1{F0: genFn0(), F1: 18446744073709551615, F2: 2}
    for fi23 := 0; fi23 < 7; fi23++ {
        fmt.Println(v2)
        switch ("foo bar" + ("foo bar" + "\n\t")) {
        case "hello":
            if (!true) {
                v24 := GenS4{F0: 42, F1: "\n\t", F2: "hello"}
                _ = v24
            }
        default:
            v3 = v3
        }
        _ = fi23
    }
    var v25 float64 = 1e10
    if false {
        if true {
            if true {
                v26 := []int{(v1 & v1), -2147483648, v1, (v1 / ((-1) | 1))}
                _ = v26
            } else {
                var v27 uint8 = 127
                for fi28 := 0; fi28 < 14; fi28++ {
                    switch "foo bar" {
                    case "hello":
                        var v29 string = ""
                        _ = v29
                    }
                    v30 := []int{v1, 32768, v1}
                    v2 = (v2 % (((v2 >> (((v3 * v3)) & 63))) | 1))
                    _ = fi28
                    _ = v30
                }
                v25 = -2.25
                _ = v27
            }
            var v31 uint64 = 18446744073709551615
            v25 = float64(0.0)
            _ = v31
        } else {
            for fi32 := 0; fi32 < 4; fi32++ {
                var v33 uint16 = 42
                v34 := GenS5{F0: 0, F1: (v33 << ((v3) & 63))}
                var v35 uint32 = 2
                _ = fi32
                _ = v33
                _ = v34
                _ = v35
            }
            fmt.Println(v25, v2)
            fmt.Println(v25, v2, v25)
        }
    } else {
        v1 = (v1 & (v1 % (((v1 - (v1 * v1))) | 1)))
        v1 += v1
    }
    var v36 string = "\n\t"
    var v37 uint8 = 42
    switch "\n\t" {
    case "hello":
        fmt.Println(v3, v36)
    case "a":
        v2 = v2
    }
    var v38 int = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    _ = v1
    _ = v11
    _ = v2
    _ = v22
    _ = v25
    _ = v3
    _ = v36
    _ = v37
    _ = v38
}
