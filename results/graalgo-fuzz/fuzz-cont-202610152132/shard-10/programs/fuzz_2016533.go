package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 int32 }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 int16; F1 int64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint16; F1 uint64; F2 uint8 }
func genFn0() bool {
    for fi4 := 0; fi4 < 7; fi4++ {
        var v5 int64 = 442
        _ = fi4
        _ = v5
    }
    var v6 uint32 = 42
    for fi7 := 0; fi7 < 14; fi7++ {
        switch fi7 {
        case -206:
            v8 := []int{fi7, fi7, (-(fi7)), (fi7 & fi7)}
            _ = v8
        }
        for fi9 := 0; fi9 < 11; fi9++ {
            v10 := []int{}
            _ = fi9
            _ = v10
        }
        v11 := []int{4294967295, 4294967295, fi7, fi7}
        _ = fi7
        _ = v11
    }
    switch "" {
    case "\n\t":
        v12 := GenS1{F0: 128, F1: 116, F2: -1}
        _ = v12
    case "foo bar":
        var v13 int16 = 1000
        _ = v13
    }
    _ = v6
    return true
}

func main() {
    var v1 int = 9223372036854775807
    var v2 int16 = -1
    var v3 string = (("foo bar" + "foo bar") + "foo bar")
    var v14 bool = true
    for fi15 := 0; fi15 < 2; fi15++ {
        v2 += v2
        fmt.Println(v1, v14)
        _ = fi15
    }
    for fi16 := 0; fi16 < 4; fi16++ {
        fmt.Println(fi16, v14)
        var v17 string = (v3 + v3)
        var v18 int8 = -128
        _ = fi16
        _ = v17
        _ = v18
    }
    for fi19 := 0; fi19 < 6; fi19++ {
        for fi20 := 0; fi20 < 9; fi20++ {
            switch (v1 ^ (v1 << ((921) & 63))) {
            case 400:
                if ((v2 - (v2 + (v2 ^ 1))) >= 973) {
                    v21 := GenS2{F0: v2, F1: false}
                    var v22 float32 = 3.14159
                    v23 := GenS3{F0: v2, F1: 933}
                    var v24 int64 = 0
                    _ = v21
                    _ = v22
                    _ = v23
                    _ = v24
                } else {
                    v25 := GenS4{F0: v1}
                    v26 := []int{v1, -791}
                    v1 = fi19
                    _ = v25
                    _ = v26
                }
            }
            fmt.Println(v2, v3, fi20)
            if v14 {
                v27 := GenS1{F0: (^(fi20)), F1: 0, F2: 2147483647}
                v28 := GenS3{F0: v2, F1: 42}
                _ = v27
                _ = v28
            }
            _ = fi20
        }
        v29 := []int{(fi19 ^ fi19)}
        _ = fi19
        _ = v29
    }
    var v30 bool = (0.1 >= ((-2.25 * 0.1) / -2.25))
    for fi31 := 0; fi31 < 15; fi31++ {
        var v32 uint = 787
        var v33 uint64 = 130
        if false {
            v32 = (^(v32))
            for fi34 := 0; fi34 < 7; fi34++ {
                v33 -= v33
                _ = fi34
            }
            switch v3 {
            case "\n\t":
                v35 := GenS5{F0: 65535, F1: (v33 % ((v33) | 1)), F2: 42}
                _ = v35
            case "hello":
                v30 = (((!(v14 && v30)) || (-1 >= 42)) && (42 >= 128))
            case "foo bar":
                var v36 float64 = 0.0
                _ = v36
            default:
                var v37 int32 = 42
                _ = v37
            }
        } else {
            v38 := []int{(fi31 - -482), (-(fi31))}
            _ = v38
        }
        _ = fi31
        _ = v32
        _ = v33
    }
    v2 = (v2 ^ -1)
    switch v3 {
    case "":
        v39 := []int{}
        _ = v39
    case "\n\t":
        var v40 bool = (42 < 859)
        _ = v40
    case "foo bar":
        switch v1 {
        case 128:
            var v41 bool = v30
            _ = v41
        }
    default:
        v2 += v2
    }
    fmt.Println(v1)
    for fi42 := 0; fi42 < 8; fi42++ {
        v43 := []int{v1, fi42, v1, fi42}
        var v44 int64 = -9223372036854775808
        switch -844 {
        case 1:
            v2 += v2
        case 127:
            _ = len(v43)
        }
        _ = fi42
        _ = v43
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v30
}
