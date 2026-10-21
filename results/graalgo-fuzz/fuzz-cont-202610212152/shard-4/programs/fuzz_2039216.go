package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32 }
type GenS2 struct { F0 string; F1 uint64 }
type GenS3 struct { F0 bool; F1 int8; F2 uint32 }

func main() {
    var v1 int16 = -491
    var v2 int = -596
    var v3 int8 = 0
    var v4 uint16 = 942
    fmt.Println(v4)
    v5 := GenS1{F0: -35, F1: 0}
    switch "" {
    case "hello":
        switch "hello" {
        case "\n\t":
            var v6 int64 = -9223372036854775808
            _ = v6
        case "hello":
            var v7 float64 = (1.5 + 1e10)
            _ = v7
        case "foo bar":
            v2 = -453
        }
    case "foo bar":
        v3 -= (-(v3))
    }
    if false {
        for fi8 := 0; fi8 < 9; fi8++ {
            v3 = (-(v3))
            v9 := []int{(fi8 ^ fi8)}
            _ = fi8
            _ = v9
        }
        for fi10 := 0; fi10 < 17; fi10++ {
            var v11 float64 = 0.0
            var v12 uint64 = 42
            v13 := []int{v2, fi10, (fi10 % (((v2 >> ((316) & 63))) | 1))}
            _ = fi10
            _ = v11
            _ = v12
            _ = v13
        }
        fmt.Println(v3)
    } else {
        v14 := []int{128, v2, 1, (v2 >> ((734) & 63))}
        v15 := GenS1{F0: v3, F1: 0}
        switch (v2 / (((-(v2))) | 1)) {
        case 0:
            v16 := GenS2{F0: ("\n\t" + "foo bar"), F1: 1}
            _ = v16
        case 32768:
            var v17 bool = true
            _ = v17
        }
        _ = v14
        _ = v15
    }
    for fi18 := 0; fi18 < 1; fi18++ {
        switch (-(v2)) {
        case 173:
            switch fi18 {
            case 0:
                var v19 int8 = (v3 % ((v3) | 1))
                _ = v19
            case -9223372036854775808:
                if false {
                    v3 -= (v3 % (((v3 - v3)) | 1))
                    for fi20 := 0; fi20 < 7; fi20++ {
                        v21 := GenS2{F0: ((("a" + ("" + "\n\t")) + "foo bar") + "hello"), F1: 418}
                        v1 = 42
                        _ = fi20
                        _ = v21
                    }
                } else {
                    v2 = (v2 << ((627) & 63))
                }
            case 587:
                var v22 float64 = 0.0
                _ = v22
            }
        case 2147483647:
            for fi23 := 0; fi23 < 13; fi23++ {
                v4 = v4
                v24 := GenS2{F0: ("" + "foo bar"), F1: 612}
                _ = fi23
                _ = v24
            }
        }
        switch "hello" {
        case "foo bar":
            v25 := GenS3{F0: (-9223372036854775808 >= -42), F1: (v3 | -58), F2: 4294967295}
            _ = v25
        case "a":
            var v26 string = ""
            _ = v26
        default:
            v3 = (^(v3))
        }
        if false {
            var v27 float64 = 0.0
            _ = v27
        }
        _ = fi18
    }
    var v28 string = "a"
    var v29 float32 = 3.14159
    v30 := []int{v2, v2, v2}
    v4 = (^(v4))
    switch v28 {
    case "":
        v30 = append(v30, v2)
    case "hello":
        if true {
            var v31 uint32 = 0
            if len(v30) >= 2 {
                v30 = v30[1:len(v30)]
            }
            _ = v31
        }
    case "a":
        v32 := GenS2{F0: "a", F1: 18446744073709551615}
        _ = v32
    }
    switch ("" + v28) {
    case "foo bar":
        v28 = v28
    case "\n\t":
        for fi33 := 0; fi33 < 4; fi33++ {
            v4 = v4
            switch 1 {
            case -2147483648:
                v30 = append(v30, (fi33 & 778))
            default:
                fmt.Println(v29, v1)
            }
            _ = fi33
        }
    case "hello":
        v3 += v3
    default:
        v28 = (v28 + ("\n\t" + (v28 + v28)))
    }
    var v34 uint32 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v34
    _ = v4
    _ = v5
}
