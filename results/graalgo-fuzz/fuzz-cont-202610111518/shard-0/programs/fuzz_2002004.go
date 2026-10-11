package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 int; F1 uint32; F2 uint32 }
type GenS4 struct { F0 uint8; F1 int8; F2 int16 }

func main() {
    var v1 bool = ((!((((false || true) || false) || (0.0 < 0.0)) && true)) && ("\n\t" >= "\n\t"))
    var v2 int32 = 42
    var v3 uint64 = 190
    var v4 int8 = 0
    var v5 uint = 1
    var v6 int64 = 9223372036854775807
    for fi7 := 0; fi7 < 15; fi7++ {
        v2 = (^(v2))
        v1 = ((v4 / ((v4) | 1)) <= (^(v4)))
        switch "\n\t" {
        case "foo bar":
            v8 := GenS1{F0: (v5 ^ (v5 ^ uint(v3)))}
            _ = v8
        case "\n\t":
            v9 := []int{42, 127, 821}
            _ = v9
        }
        _ = fi7
    }
    v3 = v3
    v10 := GenS1{F0: v5}
    for fi11 := 0; fi11 < 19; fi11++ {
        for fi12 := 0; fi12 < 5; fi12++ {
            v3 = (^(v3))
            v13 := GenS1{F0: v5}
            v14 := GenS2{F0: 528, F1: v4}
            _ = fi12
            _ = v13
            _ = v14
        }
        if true {
            v2 += v2
            var v15 uint32 = 177
            v16 := GenS2{F0: 710, F1: -128}
            _ = v15
            _ = v16
        } else {
            if (v4 != v4) {
                v17 := []int{112, fi11, fi11}
                _ = v17
            } else {
                var v18 int64 = (^(v6))
                v6 += v6
                var v19 string = "a"
                _ = v18
                _ = v19
            }
        }
        fmt.Println(fi11, v1, v6)
        _ = fi11
    }
    if (false && (1.5 > (-2.25 / 1e10))) {
        if false {
            switch "a" {
            case "a":
                v20 := []int{4294967295, 42, 9223372036854775807, -555}
                _ = v20
            default:
                switch -173 {
                case -2147483648:
                    v21 := GenS2{F0: 2, F1: -39}
                    _ = v21
                case 0:
                    v22 := GenS3{F0: 1, F1: 0, F2: 1}
                    _ = v22
                default:
                    fmt.Println(v2, v4, v5)
                }
            }
            v2 = v2
        }
        var v23 int = 2147483647
        v24 := GenS3{F0: v23, F1: 474, F2: 4294967295}
        _ = v23
        _ = v24
    } else {
        for fi25 := 0; fi25 < 14; fi25++ {
            v26 := []int{(fi25 + fi25), (fi25 >> (((v5 >> ((255) & 63))) & 63))}
            _ = fi25
            _ = v26
        }
        v6 = (-(v6))
        switch ("a" + "foo bar") {
        case "a":
            var v27 uint64 = 4294967295
            _ = v27
        case "\n\t":
            for fi28 := 0; fi28 < 5; fi28++ {
                var v29 string = "hello"
                _ = fi28
                _ = v29
            }
        default:
            v3 += 2
        }
    }
    v6 -= 9223372036854775807
    v1 = false
    v30 := []int{4294967295, -592}
    var v31 uint16 = 447
    var v32 float32 = ((-2.25 - -2.25) / 1.0)
    for fi33 := 0; fi33 < 18; fi33++ {
        var v34 uint32 = 4294967295
        var v35 int8 = v4
        v36 := GenS2{F0: v31, F1: -65}
        _ = fi33
        _ = v34
        _ = v35
        _ = v36
    }
    fmt.Println(v31, v2)
    if len(v30) >= 2 {
        v30 = v30[1:len(v30)]
    }
    for fi37 := 0; fi37 < 8; fi37++ {
        var v38 uint32 = 42
        var v39 uint16 = v31
        _ = fi37
        _ = v38
        _ = v39
    }
    v4 = v4
    v40 := GenS4{F0: 1, F1: v4, F2: 42}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
