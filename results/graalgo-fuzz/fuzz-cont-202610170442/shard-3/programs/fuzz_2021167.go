package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64; F2 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int; F1 uint64 }
type GenS4 struct { F0 int8; F1 uint32 }

func main() {
    var v1 int8 = -128
    var v2 int = 437
    var v3 string = "a"
    var v4 uint8 = 220
    var v5 int8 = -1
    switch v2 {
    case -1:
        v1 += 1
    case 19:
        switch v2 {
        case 42:
            v6 := []int{(v2 % ((v2) | 1)), v2, (v2 >> ((42) & 63)), v2}
            _ = v6
        default:
            var v7 int8 = v1
            _ = v7
        }
    case 9223372036854775807:
        if (563 > 42) {
            v8 := GenS1{F0: 32767, F1: int64(v1), F2: v4}
            v9 := GenS1{F0: 984, F1: 42, F2: (^(v4))}
            if (((199 < 2147483647) || true) && (4294967295 >= 1)) {
                for fi10 := 0; fi10 < 9; fi10++ {
                    fmt.Println(fi10)
                    _ = fi10
                }
            }
            _ = v8
            _ = v9
        } else {
            v11 := []int{488}
            v12 := GenS2{F0: true}
            v13 := GenS1{F0: 0, F1: 0, F2: v4}
            _ = v11
            _ = v12
            _ = v13
        }
    }
    v4 = 0
    var v14 int = v2
    var v15 float32 = -2.25
    var v16 uint64 = 18446744073709551615
    fmt.Println(v5, v15, v4)
    v14 += -9223372036854775808
    switch 32767 {
    case 757:
        switch (v2 ^ v2) {
        case 32767:
            if false {
                for fi17 := 0; fi17 < 14; fi17++ {
                    v18 := []int{fi17, fi17}
                    _ = len(v18)
                    var v19 string = ""
                    _ = fi17
                    _ = v18
                    _ = v19
                }
                v20 := []int{-852}
                for fi21 := 0; fi21 < 15; fi21++ {
                    v16 = v16
                    _ = fi21
                }
                _ = v20
            }
        case 128:
            v15 = 3.14159
        case 9223372036854775807:
            if (!false) {
                v2 -= v14
                switch (v14 >> ((2) & 63)) {
                case 32768:
                    var v22 uint = 255
                    _ = v22
                case -517:
                    for fi23 := 0; fi23 < 16; fi23++ {
                        var v24 uint32 = 2147483648
                        v1 += 61
                        if true {
                            v25 := GenS3{F0: v14, F1: v16}
                            var v26 float32 = 1e10
                            _ = v25
                            _ = v26
                        } else {
                            v14 = v14
                            var v27 bool = true
                            _ = v27
                        }
                        _ = fi23
                        _ = v24
                    }
                }
                v14 = v2
            }
        default:
            var v28 float64 = 3.14159
            _ = v28
        }
    case 9223372036854775807:
        fmt.Println(v2, v2)
    case -1:
        var v29 uint8 = 42
        _ = v29
    default:
        var v30 string = ""
        _ = v30
    }
    if true {
        var v31 int = v14
        _ = v31
    } else {
        switch v3 {
        case "":
            fmt.Println(v4, v3, v4)
        case "foo bar":
            for fi32 := 0; fi32 < 5; fi32++ {
                v15 = v15
                switch "a" {
                case "hello":
                    switch v3 {
                    case "hello":
                        v4 = v4
                    case "":
                        for fi33 := 0; fi33 < 10; fi33++ {
                            v14 -= v14
                            v16 = v16
                            if true {
                                v34 := []int{(-(v14)), 88}
                                _ = v34
                            } else {
                                v35 := GenS2{F0: (9223372036854775807 <= 1)}
                                _ = v35
                            }
                            _ = fi33
                        }
                    case "a":
                        v36 := GenS4{F0: (-(v1)), F1: 0}
                        _ = v36
                    }
                case "foo bar":
                    for fi37 := 0; fi37 < 1; fi37++ {
                        v4 = (v4 * 191)
                        _ = fi37
                    }
                case "":
                    v14 += fi32
                }
                var v38 int16 = -1
                _ = fi32
                _ = v38
            }
        default:
            fmt.Println(v14, v1, v3)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
