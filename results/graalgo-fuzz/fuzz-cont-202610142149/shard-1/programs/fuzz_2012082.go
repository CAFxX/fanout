package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 int }
type GenS2 struct { F0 int64; F1 int64 }

func main() {
    var v1 float32 = 1e10
    var v2 string = "foo bar"
    var v3 int16 = 42
    var v4 uint8 = 42
    v5 := []int{-1, 685, 32768, -299}
    var v6 bool = true
    v7 := GenS1{F0: 0, F1: 4294967295, F2: 9223372036854775807}
    for fi8 := 0; fi8 < 9; fi8++ {
        v4 = (v4 - (v4 ^ (v4 + v4)))
        for fi9 := 0; fi9 < 1; fi9++ {
            _ = len(v5)
            var v10 string = v2
            v2 = "hello"
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    var v11 string = ("" + ("\n\t" + v2))
    v5 = append(v5, 32768)
    if (v6 || (v4 != v4)) {
        var v12 uint16 = 206
        for fi13 := 0; fi13 < 19; fi13++ {
            var v14 float64 = 0.1
            _ = fi13
            _ = v14
        }
        var v15 string = v2
        _ = v12
        _ = v15
    }
    for fi16 := 0; fi16 < 17; fi16++ {
        switch fi16 {
        case 9223372036854775807:
            var v17 string = (v2 + "a")
            _ = v17
        case -1:
            if v6 {
                if (v6 && (1 > -74)) {
                    for fi18 := 0; fi18 < 5; fi18++ {
                        if len(v5) > 0 {
                            v5[0] = -2147483648
                        }
                        fmt.Println(v2, v2, v2)
                        _ = fi18
                    }
                    var v19 uint8 = (^(v4))
                    _ = v19
                }
            } else {
                v1 -= v1
                fmt.Println(v4, v6)
                var v20 int16 = -32768
                _ = v20
            }
        case 427:
            if v6 {
                var v21 string = "hello"
                v21 = "hello"
                _ = v21
            } else {
                var v22 int = (fi16 ^ fi16)
                var v23 int8 = 127
                v3 -= -1
                _ = v22
                _ = v23
            }
        }
        var v24 float64 = 1.5
        switch "\n\t" {
        case "foo bar":
            switch ("foo bar" + v11) {
            case "foo bar":
                var v25 uint32 = 523
                _ = v25
            case "hello":
                fmt.Println(v3, v6, v4)
            }
        case "a":
            for fi26 := 0; fi26 < 13; fi26++ {
                switch (fi16 >> ((664) & 63)) {
                case 32768:
                    v27 := GenS1{F0: 964, F1: 72, F2: (fi16 % (((fi26 ^ (fi26 / (((fi16 * fi26)) | 1)))) | 1))}
                    _ = v27
                }
                _ = fi26
            }
        }
        _ = fi16
        _ = v24
    }
    for fi28 := 0; fi28 < 17; fi28++ {
        fmt.Println(v2)
        v29 := GenS2{F0: -860, F1: 42}
        if (v1 < (v1 + float32(0.0))) {
            fmt.Println(v11)
            fmt.Println(v11, v1, v1)
            switch v11 {
            case "foo bar":
                v2 = v11
            case "a":
                fmt.Println(v4, v1)
            case "\n\t":
                v5 = append(v5, fi28)
            }
        } else {
            v5 = append(v5, fi28)
            for fi30 := 0; fi30 < 14; fi30++ {
                var v31 string = v2
                v32 := GenS1{F0: 42, F1: 0, F2: (fi28 ^ 1)}
                if len(v5) > 0 {
                    v5[0] = fi30
                }
                switch fi30 {
                case 42:
                    var v33 int = 127
                    _ = v33
                case 946:
                    v2 = ((((("" + "\n\t") + "hello") + "") + (v11 + v31)) + "")
                case 28:
                    _ = len(v5)
                }
                v2 = ""
                _ = fi30
                _ = v31
                _ = v32
            }
        }
        _ = fi28
        _ = v29
    }
    var v34 float64 = 3.14159
    var v35 int32 = -793
    if (127 > -1) {
        fmt.Println(v35, v3)
        v11 = ((v11 + "\n\t") + "foo bar")
    } else {
        fmt.Println(v11)
        v36 := GenS1{F0: 32767, F1: 42, F2: 1}
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
