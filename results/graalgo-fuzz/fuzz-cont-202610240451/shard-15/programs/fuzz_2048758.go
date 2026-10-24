package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 uint16 }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 uint16; F1 bool }
func genFn0(p0 int16) int8 {
    p0 -= p0
    v4 := []int{}
    if false {
        for fi5 := 0; fi5 < 12; fi5++ {
            var v6 uint64 = 0
            v7 := GenS1{F0: 0, F1: ("hello" + ("foo bar" + "")), F2: 42}
            _ = fi5
            _ = v6
            _ = v7
        }
    }
    switch ("\n\t" + "\n\t") {
    case "\n\t":
        if ((0.1 > (0.0 * 3.14159)) && (100 >= 95)) {
            v8 := GenS2{F0: (((("" + "\n\t") + "foo bar") + ("hello" + (((("hello" + "hello") + "") + "\n\t") + ""))) + "a"), F1: 613}
            if false {
                p0 = p0
                fmt.Println(p0)
            }
            _ = v8
        }
    case "a":
        switch "\n\t" {
        case "hello":
            for fi9 := 0; fi9 < 2; fi9++ {
                var v10 uint16 = 255
                _ = fi9
                _ = v10
            }
        case "\n\t":
            var v11 int8 = 1
            _ = v11
        }
    case "":
        p0 += (p0 % (((p0 & (-(p0)))) | 1))
    default:
        if len(v4) > 0 {
            v4[0] = 2147483647
        }
    }
    _ = p0
    _ = v4
    return 1
}

func main() {
    var v1 int64 = 9223372036854775807
    var v2 float64 = 1.5
    var v3 bool = (!true)
    if v3 {
        switch "a" {
        case "hello":
            if (2 >= 65535) {
                var v12 uint64 = 18446744073709551615
                _ = v12
            }
        case "":
            v3 = v3
        default:
            v13 := GenS2{F0: "foo bar", F1: 85}
            _ = v13
        }
    }
    if ((float32(v2) * 0.1) != 0.1) {
        fmt.Println(v1, v3, v1)
        switch ("\n\t" + "hello") {
        case "":
            switch 128 {
            case 727:
                switch "hello" {
                case "hello":
                    var v14 float32 = 1e10
                    _ = v14
                case "":
                    for fi15 := 0; fi15 < 14; fi15++ {
                        v16 := []int{fi15, 763, (^(fi15)), 42}
                        _ = fi15
                        _ = v16
                    }
                }
            default:
                fmt.Println(v1, v2, v3)
            }
        case "a":
            var v17 int = -807
            _ = v17
        case "foo bar":
            if (v3 && true) {
                var v18 uint32 = 2147483647
                _ = v18
            } else {
                switch 848 {
                case 128:
                    v19 := GenS2{F0: "\n\t", F1: 42}
                    _ = v19
                case 934:
                    for fi20 := 0; fi20 < 9; fi20++ {
                        var v21 int16 = 1000
                        _ = fi20
                        _ = v21
                    }
                default:
                    v22 := []int{-2147483648, 740}
                    _ = v22
                }
                v1 = 0
                for fi23 := 0; fi23 < 9; fi23++ {
                    var v24 uint64 = 0
                    _ = fi23
                    _ = v24
                }
            }
        }
        v1 = (v1 + (-(v1)))
    } else {
        var v25 int64 = 877
        v3 = (641 >= 0)
        _ = v25
    }
    if v3 {
        for fi26 := 0; fi26 < 18; fi26++ {
            v27 := []int{}
            var v28 uint16 = 14
            _ = fi26
            _ = v27
            _ = v28
        }
        v29 := GenS1{F0: 255, F1: "a", F2: 1}
        _ = v29
    } else {
        v30 := GenS3{F0: 42, F1: false}
        var v31 int32 = 42
        v31 = (v31 % ((v31) | 1))
        _ = v30
        _ = v31
    }
    v3 = false
    fmt.Println(v3)
    var v32 int = 128
    v33 := GenS1{F0: 0, F1: "", F2: 32767}
    var v34 string = ""
    if v3 {
        v35 := []int{}
        _ = v35
    }
    fmt.Println(v1)
    switch v34 {
    case "":
        switch "hello" {
        case "hello":
            v36 := []int{0, v32, v32, (-(v32))}
            _ = v36
        case "foo bar":
            v37 := []int{(^(v32)), v32, (v32 | (v32 >> ((42) & 63))), v32}
            _ = v37
        default:
            for fi38 := 0; fi38 < 2; fi38++ {
                v39 := []int{(fi38 % ((1) | 1)), 4294967295}
                var v40 bool = false
                _ = fi38
                _ = v39
                _ = v40
            }
        }
    default:
        for fi41 := 0; fi41 < 13; fi41++ {
            v42 := []int{}
            _ = fi41
            _ = v42
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
}
