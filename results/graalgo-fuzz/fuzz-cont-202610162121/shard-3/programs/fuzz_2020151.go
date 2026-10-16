package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 int8 }
type GenS2 struct { F0 uint32; F1 uint64; F2 int16 }
type GenS3 struct { F0 uint64; F1 uint64 }
type GenS4 struct { F0 uint16; F1 uint16; F2 uint }
type GenS5 struct { F0 uint32 }

func main() {
    var v1 uint = 1
    var v2 bool = (0 > 163)
    var v3 uint = 0
    var v4 string = ""
    var v5 uint32 = 0
    switch int(v1) {
    case 9223372036854775807:
        for fi6 := 0; fi6 < 5; fi6++ {
            v7 := GenS1{F0: -9223372036854775808, F1: -2147483648, F2: 0}
            var v8 uint32 = (v5 | v5)
            fmt.Println(v2)
            _ = fi6
            _ = v7
            _ = v8
        }
    }
    if false {
        fmt.Println(v3, v2)
        for fi9 := 0; fi9 < 13; fi9++ {
            v10 := []int{fi9}
            _ = fi9
            _ = v10
        }
    }
    for fi11 := 0; fi11 < 15; fi11++ {
        switch fi11 {
        case -411:
            fmt.Println(v3, v3, v4)
        case 0:
            switch v4 {
            case "\n\t":
                v12 := GenS2{F0: (v5 / (((v5 - v5)) | 1)), F1: 38, F2: 0}
                _ = v12
            case "hello":
                for fi13 := 0; fi13 < 15; fi13++ {
                    var v14 int16 = 203
                    v15 := GenS1{F0: 0, F1: -1, F2: -128}
                    v16 := []int{fi13}
                    _ = fi13
                    _ = v14
                    _ = v15
                    _ = v16
                }
            }
        case 128:
            var v17 uint32 = 1
            _ = v17
        }
        switch -1 {
        case -986:
            var v18 uint64 = 142
            _ = v18
        default:
            v19 := GenS1{F0: fi11, F1: -1, F2: -1}
            _ = v19
        }
        _ = fi11
    }
    v1 -= 448
    switch (v4 + v4) {
    case "a":
        var v20 uint16 = 861
        _ = v20
    case "":
        for fi21 := 0; fi21 < 4; fi21++ {
            fmt.Println(fi21, v1, v5)
            var v22 float64 = -2.25
            _ = fi21
            _ = v22
        }
    }
    var v23 int64 = -705
    for fi24 := 0; fi24 < 1; fi24++ {
        var v25 bool = v2
        var v26 uint = (v1 * 937)
        _ = fi24
        _ = v25
        _ = v26
    }
    switch v4 {
    case "foo bar":
        switch "a" {
        case "foo bar":
            var v27 uint64 = 894
            _ = v27
        case "a":
            for fi28 := 0; fi28 < 5; fi28++ {
                switch v4 {
                case "foo bar":
                    for fi29 := 0; fi29 < 4; fi29++ {
                        var v30 uint32 = (v5 / ((42) | 1))
                        fmt.Println(fi28, v4, v3)
                        fmt.Println(fi28, v2, v5)
                        v30 -= v5
                        _ = fi29
                        _ = v30
                    }
                case "hello":
                    v31 := GenS3{F0: 4294967295, F1: 2}
                    _ = v31
                }
                _ = fi28
            }
        case "\n\t":
            for fi32 := 0; fi32 < 14; fi32++ {
                v33 := []int{(fi32 | (fi32 | fi32)), (fi32 % ((-2147483648) | 1))}
                _ = fi32
                _ = v33
            }
        }
    case "hello":
        v2 = false
    default:
        v34 := GenS4{F0: 488, F1: 2, F2: (v1 / (((^(v1))) | 1))}
        _ = v34
    }
    v23 = (v23 - v23)
    v35 := GenS5{F0: v5}
    v36 := GenS3{F0: 245, F1: 42}
    fmt.Println(v1)
    v2 = true
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
