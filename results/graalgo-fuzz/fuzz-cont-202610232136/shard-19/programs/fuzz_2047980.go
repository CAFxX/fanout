package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int; F1 uint }

func main() {
    var v1 float64 = ((0.0 - (0.1 / 0.1)) - ((0.1 * 3.14159) * 1e10))
    var v2 uint32 = 1
    var v3 uint = 18446744073709551615
    v1 -= 1e10
    for fi4 := 0; fi4 < 19; fi4++ {
        if true {
            if true {
                v1 -= v1
                v5 := GenS1{F0: v2}
                _ = v5
            } else {
                v6 := []int{fi4, fi4, fi4, -9223372036854775808}
                fmt.Println(v3)
                v1 = v1
                _ = v6
            }
        } else {
            v7 := []int{}
            _ = v7
        }
        var v8 float64 = v1
        if true {
            v9 := GenS1{F0: (v2 - 2)}
            if true {
                fmt.Println(v8, v1, v2)
            } else {
                var v10 uint8 = 0
                _ = v10
            }
            if true {
                switch fi4 {
                case -708:
                    fmt.Println(v8, v1)
                }
                v11 := []int{fi4, fi4, 128}
                v12 := GenS1{F0: v2}
                _ = v11
                _ = v12
            }
            _ = v9
        } else {
            for fi13 := 0; fi13 < 9; fi13++ {
                var v14 uint64 = 0
                _ = fi13
                _ = v14
            }
            v15 := []int{}
            _ = len(v15)
            _ = v15
        }
        _ = fi4
        _ = v8
    }
    if false {
        switch "a" {
        case "a":
            v16 := GenS1{F0: v2}
            _ = v16
        case "\n\t":
            v1 += v1
        case "foo bar":
            var v17 int16 = int16(v3)
            _ = v17
        }
        var v18 float32 = 0.1
        _ = v18
    } else {
        v2 = (v2 / ((v2) | 1))
        switch 128 {
        case 0:
            var v19 uint8 = 128
            _ = v19
        }
    }
    switch "" {
    case "":
        var v20 int16 = -32768
        _ = v20
    case "foo bar":
        var v21 uint8 = 0
        _ = v21
    }
    for fi22 := 0; fi22 < 14; fi22++ {
        switch "hello" {
        case "\n\t":
            var v23 int32 = 259
            _ = v23
        case "":
            v24 := GenS1{F0: (^(v2))}
            _ = v24
        case "foo bar":
            v25 := []int{fi22, (fi22 - fi22), fi22, (fi22 | fi22)}
            _ = v25
        }
        _ = fi22
    }
    for fi26 := 0; fi26 < 18; fi26++ {
        var v27 int64 = 871
        _ = fi26
        _ = v27
    }
    if false {
        v2 += 2147483647
    } else {
        if ((-2147483648 != -1) && (true && (true && true))) {
            if (2 > 1) {
                var v28 uint = (v3 & v3)
                v1 = v1
                v3 -= (v28 ^ v28)
                _ = v28
            }
            v29 := GenS2{F0: -384, F1: (v3 << (((v3 % ((774) | 1))) & 63))}
            _ = v29
        } else {
            var v30 uint32 = v2
            v31 := []int{}
            _ = v30
            _ = v31
        }
        v32 := GenS1{F0: v2}
        for fi33 := 0; fi33 < 18; fi33++ {
            if false {
                var v34 int8 = 43
                if (v34 != (v34 & -1)) {
                    var v35 bool = (v3 < (v3 ^ v3))
                    for fi36 := 0; fi36 < 11; fi36++ {
                        if true {
                            v34 += -128
                            var v37 uint64 = 1
                            _ = v37
                        } else {
                            v38 := []int{-360, -232, fi33, fi36}
                            var v39 float64 = (v1 / v1)
                            _ = v38
                            _ = v39
                        }
                        fmt.Println(v2)
                        _ = fi36
                    }
                    _ = v35
                } else {
                    v40 := GenS1{F0: v2}
                    var v41 bool = (true || false)
                    _ = v40
                    _ = v41
                }
                _ = v34
            } else {
                v2 += 2147483648
            }
            _ = fi33
        }
        _ = v32
    }
    for fi42 := 0; fi42 < 9; fi42++ {
        v3 = (v3 * 42)
        _ = fi42
    }
    v43 := []int{-3}
    if ("hello" <= "a") {
        v2 = v2
    } else {
        var v44 uint16 = 32767
        if len(v43) >= 2 {
            v43 = v43[1:len(v43)]
        }
        _ = v44
    }
    v1 = v1
    v2 = 2147483648
    switch "hello" {
    case "\n\t":
        for fi45 := 0; fi45 < 6; fi45++ {
            v43 = append(v43, -973)
            _ = len(v43)
            _ = fi45
        }
    case "":
        var v46 float64 = 1.5
        _ = v46
    case "hello":
        fmt.Println(v1)
    default:
        if false {
            var v47 int32 = 2147483647
            _ = v47
        } else {
            fmt.Println(v1)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v43:", len(v43), v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v43
}
