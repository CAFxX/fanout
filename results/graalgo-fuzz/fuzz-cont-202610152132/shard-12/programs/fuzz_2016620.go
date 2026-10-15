package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint; F1 int32 }
type GenS3 struct { F0 int8; F1 uint32 }
type GenS4 struct { F0 int8 }

func main() {
    var v1 uint = 42
    var v2 int16 = 0
    var v3 float64 = (0.0 / 3.14159)
    if false {
        for fi4 := 0; fi4 < 11; fi4++ {
            var v5 uint = v1
            if (1 >= 1) {
                v2 = 0
                v6 := []int{fi4}
                v7 := GenS1{F0: (v1 * v1)}
                _ = v6
                _ = v7
            }
            v8 := GenS1{F0: (v5 % ((42) | 1))}
            _ = fi4
            _ = v5
            _ = v8
        }
        switch 32768 {
        case -111:
            if false {
                if false {
                    var v9 int = 4294967295
                    v10 := []int{v9, 9223372036854775807}
                    _ = v10
                    _ = v9
                } else {
                    for fi11 := 0; fi11 < 12; fi11++ {
                        fmt.Println(v2)
                        _ = fi11
                    }
                    var v12 uint16 = 65535
                    if ((v1 ^ v1) > v1) {
                        v13 := GenS1{F0: v1}
                        if ((false && true) && true) {
                            v2 = 0
                        }
                        v1 -= 346
                        _ = v13
                    }
                    _ = v12
                }
                var v14 int = 9223372036854775807
                switch "" {
                case "\n\t":
                    var v15 float32 = 1e10
                    _ = v15
                case "foo bar":
                    v16 := GenS1{F0: 255}
                    _ = v16
                case "":
                    v17 := []int{v14, v14, (v14 | v14), v14}
                    _ = v17
                default:
                    v18 := GenS1{F0: v1}
                    _ = v18
                }
                _ = v14
            } else {
                v19 := GenS1{F0: v1}
                v20 := []int{-2147483648, 2147483647}
                _ = v19
                _ = v20
            }
        case 127:
            switch 9223372036854775807 {
            case 9223372036854775807:
                v3 += v3
            case 167:
                var v21 bool = true
                _ = v21
            }
        default:
            v22 := []int{0, 42}
            _ = v22
        }
        for fi23 := 0; fi23 < 6; fi23++ {
            v24 := []int{(-(fi23))}
            _ = fi23
            _ = v24
        }
    }
    v3 += float64(1e10)
    switch -1 {
    case 0:
        for fi25 := 0; fi25 < 2; fi25++ {
            v26 := []int{fi25, (fi25 << (((v1 | (v1 * 341))) & 63))}
            for fi27 := 0; fi27 < 18; fi27++ {
                if (4294967295 < 65535) {
                    if false {
                        v28 := GenS1{F0: (v1 >> ((42) & 63))}
                        _ = v28
                    } else {
                        fmt.Println(v3, fi27, fi25)
                        var v29 uint8 = 127
                        _ = v29
                    }
                }
                _ = fi27
            }
            _ = fi25
            _ = v26
        }
    case -1:
        var v30 int16 = -77
        _ = v30
    case -9223372036854775808:
        for fi31 := 0; fi31 < 8; fi31++ {
            v32 := GenS2{F0: v1, F1: 2147483647}
            v3 = 1.5
            switch ("" + (("" + "a") + "hello")) {
            case "foo bar":
                v1 -= v1
            }
            _ = fi31
            _ = v32
        }
    default:
        var v33 int = -2147483648
        _ = v33
    }
    var v34 float32 = (-2.25 * 0.0)
    for fi35 := 0; fi35 < 5; fi35++ {
        switch "foo bar" {
        case "foo bar":
            v36 := GenS2{F0: v1, F1: 1}
            _ = v36
        }
        v37 := []int{fi35, -1, (fi35 * fi35), fi35}
        _ = fi35
        _ = v37
    }
    var v38 uint = (v1 ^ v1)
    for fi39 := 0; fi39 < 13; fi39++ {
        if (100 >= -1) {
            var v40 int64 = -490
            v34 = -2.25
            v38 = 1
            _ = v40
        }
        v41 := GenS3{F0: 100, F1: 306}
        v3 = 1.5
        _ = fi39
        _ = v41
    }
    if (0.1 <= (v3 - (v3 * v3))) {
        v42 := []int{342, 42, 42, 396}
        _ = v42
    }
    v1 += 4294967295
    v43 := GenS4{F0: 42}
    v44 := []int{}
    for fi45 := 0; fi45 < 3; fi45++ {
        if true {
            var v46 uint8 = 2
            _ = v46
        } else {
            if len(v44) >= 2 {
                v44 = v44[1:len(v44)]
            }
            switch (fi45 % (((fi45 * (fi45 % ((fi45) | 1)))) | 1)) {
            case -1:
                v1 = 42
            }
            if false {
                v44 = append(v44, (fi45 | (-(fi45))))
                var v47 uint8 = 255
                _ = v47
            }
        }
        _ = fi45
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v44:", len(v44), v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v38
    _ = v43
    _ = v44
}
