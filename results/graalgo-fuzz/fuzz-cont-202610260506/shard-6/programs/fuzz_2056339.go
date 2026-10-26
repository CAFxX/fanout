package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint64; F1 uint }
type GenS4 struct { F0 int64; F1 bool; F2 uint }
type GenS5 struct { F0 int; F1 bool }

func main() {
    v1 := make(map[string]string)
    v2 := []int{-1, -211, 32767, 4294967295}
    v3 := make(map[int]int)
    v4 := []int{4294967295, 1, 4294967295}
    var v5 uint32 = 0
    if true {
        switch ("" + ("hello" + "")) {
        case "\n\t":
            var v6 bool = true
            _ = v6
        case "":
            for fi7 := 0; fi7 < 17; fi7++ {
                var v8 int = fi7
                var v9 uint64 = 2
                switch ("foo bar" + "foo bar") {
                case "a":
                    var v10 uint = 42
                    _ = v10
                }
                _ = fi7
                _ = v8
                _ = v9
            }
        case "a":
            var v11 int8 = 42
            _ = v11
        default:
            for fi12 := 0; fi12 < 17; fi12++ {
                for fi13 := 0; fi13 < 15; fi13++ {
                    switch fi12 {
                    case 1:
                        var v14 uint = 244
                        _ = v14
                    case 2147483647:
                        v5 += v5
                    }
                    v5 = v5
                    _ = fi13
                }
                _ = fi12
            }
        }
        v3[42] = 9223372036854775807
        v15 := GenS1{F0: "\n\t", F1: 9223372036854775807}
        _ = v15
    }
    switch "hello" {
    case "\n\t":
        fmt.Println(v5)
    case "hello":
        v16 := GenS2{F0: 207}
        _ = v16
    }
    var v17 uint8 = 150
    if false {
        fmt.Println(v5)
        switch "hello" {
        case "\n\t":
            if true {
                v18 := GenS1{F0: "hello", F1: -26}
                _ = v18
            } else {
                v19 := GenS1{F0: "foo bar", F1: -419}
                v5 = v5
                _ = v19
            }
        case "foo bar":
            fmt.Println(v17, v17)
        case "hello":
            v20 := GenS1{F0: "", F1: 461}
            _ = v20
        default:
            delete(v1, ("\n\t" + "\n\t"))
        }
    } else {
        v5 -= v5
    }
    v21 := GenS1{F0: "hello", F1: 127}
    v17 -= (v17 * v17)
    var v22 float32 = float32((1e10 * 0.0))
    var v23 int8 = 127
    var v24 bool = (!true)
    v25 := GenS1{F0: ("" + "foo bar"), F1: 4294967295}
    switch (("\n\t" + ("hello" + "\n\t")) + ("\n\t" + "hello")) {
    case "hello":
        switch 9223372036854775807 {
        case 32767:
            var v26 uint64 = 2
            _ = v26
        case -832:
            var v27 int64 = 9223372036854775807
            _ = v27
        case 343:
            fmt.Println(v5, v23)
        default:
            switch -2147483648 {
            case 34:
                v24 = v24
            default:
                v24 = false
            }
        }
    case "\n\t":
        if v24 {
            for fi28 := 0; fi28 < 10; fi28++ {
                for fi29 := 0; fi29 < 16; fi29++ {
                    if (1e10 <= float64(1e10)) {
                        v30 := GenS3{F0: 0, F1: 2}
                        _ = v30
                    } else {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        var v31 float32 = 1.5
                        _ = v31
                    }
                    v17 = (^(v17))
                    if len(v2) > 0 {
                        v2[0] = fi28
                    }
                    _ = fi29
                }
                _ = fi28
            }
            var v32 string = "\n\t"
            _ = v32
        }
    case "":
        var v33 int = -1
        _ = v33
    }
    v34 := GenS4{F0: 0, F1: (!true), F2: 2}
    switch "" {
    case "foo bar":
        for fi35 := 0; fi35 < 11; fi35++ {
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = fi35
        }
    case "":
        v17 += v17
    default:
        if ((!(v24 && (269 == 2))) && v24) {
            var v36 uint8 = 126
            v37 := GenS5{F0: int(v5), F1: v24}
            if false {
                var v38 string = (((("\n\t" + "") + "") + ("\n\t" + "hello")) + "")
                var v39 int16 = 42
                _ = v38
                _ = v39
            }
            _ = v36
            _ = v37
        } else {
            v24 = v24
            delete(v1, "hello")
            for fi40 := 0; fi40 < 5; fi40++ {
                v2 = append(v2, (fi40 | -505))
                _ = fi40
            }
        }
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi41 := 0; fi41 < 7; fi41++ {
        var v42 int8 = v23
        _ = fi41
        _ = v42
    }
    switch "foo bar" {
    case "\n\t":
        var v43 uint8 = v17
        _ = v43
    default:
        for fi44 := 0; fi44 < 10; fi44++ {
            v5 += v5
            var v45 float64 = 0.1
            var v46 int64 = 239
            _ = fi44
            _ = v45
            _ = v46
        }
    }
    if v24 {
        v47 := GenS2{F0: v17}
        switch "a" {
        case "hello":
            if false {
                var v48 uint16 = 42
                _ = v48
            }
        case "foo bar":
            fmt.Println(v5, v24)
        }
        _ = v47
    } else {
        fmt.Println(v24, v23)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v34
    _ = v4
    _ = v5
}
