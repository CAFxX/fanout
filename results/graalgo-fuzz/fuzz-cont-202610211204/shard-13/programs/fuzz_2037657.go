package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint; F1 uint32; F2 uint64 }
type GenS3 struct { F0 string; F1 int64 }

func main() {
    var v1 string = "hello"
    var v2 float32 = (3.14159 * 0.1)
    var v3 int16 = -943
    if true {
        v3 = v3
        if true {
            var v4 float64 = ((1.5 / ((1.5 / float64((0.0 + (1e10 - (v2 - v2))))) * float64(((1e10 + ((0.1 / 1e10) * -2.25)) * 3.14159)))) / 1.5)
            switch v1 {
            case "hello":
                v3 = (v3 >> ((437) & 63))
            }
            _ = v4
        } else {
            v5 := []int{-733, 42, -9223372036854775808}
            v1 = v1
            switch 32768 {
            case 42:
                fmt.Println(v1)
            case 2147483647:
                v1 = v1
            default:
                var v6 int64 = -1
                _ = v6
            }
            _ = v5
        }
    } else {
        v2 -= v2
    }
    v2 = v2
    for fi7 := 0; fi7 < 8; fi7++ {
        switch (v1 + v1) {
        case "a":
            fmt.Println(v2, v1, v1)
        case "hello":
            var v8 uint8 = 255
            _ = v8
        case "":
            v2 += float32(((3.14159 + float64((float32(0.0) * (float32(float64(((0.0 - (v2 * v2)) * (v2 + v2)))) - (v2 + v2))))) / 3.14159))
        }
        v1 = "foo bar"
        fmt.Println(v1)
        _ = fi7
    }
    var v9 bool = (true || false)
    v9 = v9
    for fi10 := 0; fi10 < 12; fi10++ {
        for fi11 := 0; fi11 < 17; fi11++ {
            if (1 < 1) {
                switch (fi11 << ((0) & 63)) {
                case 913:
                    switch v1 {
                    case "":
                        var v12 float32 = 1e10
                        _ = v12
                    case "hello":
                        switch (v1 + "\n\t") {
                        case "":
                            v13 := []int{}
                            _ = v13
                        case "hello":
                            var v14 int8 = 100
                            _ = v14
                        case "foo bar":
                            v15 := GenS1{F0: (int32(v3) <= -1)}
                            _ = v15
                        default:
                            switch "a" {
                            case "":
                                var v16 string = v1
                                _ = v16
                            }
                        }
                    case "a":
                        v2 -= 1.5
                    default:
                        var v17 int16 = v3
                        var v18 int16 = 1000
                        _ = v17
                        _ = v18
                    }
                case 238:
                    fmt.Println(v1)
                default:
                    v2 -= v2
                }
                if false {
                    fmt.Println(v9, fi11, v2)
                }
                if ((v3 % ((v3) | 1)) != int16(fi11)) {
                    fmt.Println(fi10, v3)
                    fmt.Println(v2)
                }
            }
            v1 = ""
            switch v1 {
            case "":
                v19 := []int{(fi10 + 42), (fi11 / (((fi11 / ((fi11) | 1))) | 1))}
                _ = v19
            case "a":
                v20 := []int{}
                _ = v20
            }
            switch "\n\t" {
            case "\n\t":
                switch v1 {
                case "a":
                    if false {
                        v1 = v1
                        var v21 uint = 851
                        _ = v21
                    }
                default:
                    switch "foo bar" {
                    case "foo bar":
                        v22 := GenS2{F0: 42, F1: 1, F2: 42}
                        _ = v22
                    case "a":
                        v23 := GenS1{F0: false}
                        _ = v23
                    }
                }
            }
            var v24 uint = 960
            _ = fi11
            _ = v24
        }
        fmt.Println(v2, v9)
        _ = fi10
    }
    v25 := []int{128, -344}
    v26 := GenS3{F0: v1, F1: 474}
    v27 := GenS3{F0: (v1 + ((v1 + v1) + "hello")), F1: -1}
    var v28 int = 9223372036854775807
    for fi29 := 0; fi29 < 11; fi29++ {
        for fi30 := 0; fi30 < 16; fi30++ {
            v9 = true
            switch v1 {
            case "foo bar":
                var v31 uint16 = 0
                _ = v31
            case "hello":
                var v32 float32 = 1.5
                _ = v32
            case "\n\t":
                switch -232 {
                case -2147483648:
                    v33 := GenS3{F0: ("a" + (v1 + ("a" + "hello"))), F1: 0}
                    _ = v33
                case 679:
                    v9 = v9
                }
            }
            var v34 uint = 12
            _ = fi30
            _ = v34
        }
        _ = fi29
    }
    v25 = append(v25, (v28 - -2147483648))
    fmt.Println(v2, v9)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v9
}
