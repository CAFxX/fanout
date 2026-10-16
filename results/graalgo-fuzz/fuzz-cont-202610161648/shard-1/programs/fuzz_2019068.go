package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16; F1 int8 }
type GenS4 struct { F0 int }

func main() {
    var v1 uint16 = 42
    var v2 string = ("\n\t" + ((((("hello" + "foo bar") + "a") + ("a" + ("hello" + ""))) + "foo bar") + ""))
    var v3 int8 = 127
    var v4 uint16 = 60
    var v5 int = 128
    v6 := make(map[int]string)
    if (!true) {
        fmt.Println(v1, v3, v5)
        var v7 uint64 = 4294967295
        _ = v7
    } else {
        v6[(v5 % (((-(v5))) | 1))] = "foo bar"
        v1 -= (v4 / (((^(v4))) | 1))
    }
    var v8 uint = 906
    fmt.Println(v4, v2)
    fmt.Println(v5, v1)
    var v9 int64 = 9223372036854775807
    var v10 uint32 = 2147483647
    v11 := GenS1{F0: -124}
    switch int(v10) {
    case 493:
        v10 = v10
    case -9223372036854775808:
        v4 += 1
    default:
        v1 += (v1 | 42)
    }
    v12 := GenS1{F0: (v5 * 42)}
    for fi13 := 0; fi13 < 8; fi13++ {
        var v14 uint16 = 65535
        for fi15 := 0; fi15 < 8; fi15++ {
            delete(v6, 32768)
            if false {
                v16 := []int{}
                var v17 uint32 = 2147483648
                _ = v16
                _ = v17
            }
            _ = fi15
        }
        fmt.Println(v1, v10)
        _ = fi13
        _ = v14
    }
    var v18 int = -1
    switch (v2 + ("foo bar" + v2)) {
    case "hello":
        _ = len(v6)
    }
    for fi19 := 0; fi19 < 2; fi19++ {
        for fi20 := 0; fi20 < 6; fi20++ {
            fmt.Println(v2, fi20, fi20)
            v21 := []int{(fi20 ^ (v18 | v5))}
            _ = fi20
            _ = v21
        }
        switch (fi19 << (((v8 ^ v8)) & 63)) {
        case 32768:
            v22 := []int{v18, 4294967295, (fi19 | 2147483647)}
            _ = v22
        case 2147483647:
            v9 = v9
        case -4:
            v23 := GenS2{F0: -827}
            _ = v23
        default:
            if false {
                for fi24 := 0; fi24 < 10; fi24++ {
                    if ((-54 > 386) && ((v10 % (((v10 >> ((v8) & 63))) | 1)) < 739)) {
                        switch v2 {
                        case "\n\t":
                            v25 := GenS1{F0: fi19}
                            _ = v25
                        case "a":
                            var v26 uint8 = 42
                            _ = v26
                        case "":
                            if false {
                                if true {
                                    v8 -= 972
                                    fmt.Println(fi24, v1, v3)
                                    v1 = (v4 - (^(v4)))
                                }
                                v4 += uint16(v10)
                                switch fi24 {
                                case -25:
                                    if (-14 == -1) {
                                        var v27 string = v2
                                        if false {
                                            v27 = ((v27 + "hello") + ((v2 + "hello") + v27))
                                            v28 := []int{fi24, (v5 % (((v5 % ((fi19) | 1))) | 1)), (fi24 % ((v18) | 1))}
                                            _ = v28
                                        }
                                        _ = v27
                                    }
                                case 32767:
                                    switch (v18 ^ 32768) {
                                    case 32767:
                                        switch v2 {
                                        case "":
                                            var v29 int16 = -782
                                            _ = v29
                                        }
                                    case 22:
                                        switch fi19 {
                                        case 462:
                                            var v30 int64 = (v9 - (v9 - int64(v5)))
                                            _ = v30
                                        }
                                        if (1e10 != float32(1e10)) {
                                            var v31 uint = v8
                                            _ = v31
                                        }
                                    case 1:
                                        v32 := GenS3{F0: -44, F1: v3}
                                        _ = v32
                                    }
                                case -713:
                                    fmt.Println(v9)
                                }
                            }
                        default:
                            v8 = v8
                        }
                        if false {
                            v33 := []int{-9223372036854775808}
                            var v34 int8 = (v3 / ((-60) | 1))
                            var v35 int16 = 42
                            _ = v33
                            _ = v34
                            _ = v35
                        }
                    }
                    switch v5 {
                    case 2147483647:
                        delete(v6, 0)
                    case -2147483648:
                        var v36 string = v2
                        _ = v36
                    case -1:
                        switch "a" {
                        case "a":
                            v10 -= v10
                        case "":
                            v2 = "\n\t"
                        default:
                            var v37 string = v2
                            _ = v37
                        }
                    default:
                        v38 := GenS3{F0: 494, F1: (v3 & (v3 | 100))}
                        _ = v38
                    }
                    _ = fi24
                }
                var v39 uint = v8
                v39 += (v8 - v39)
                _ = v39
            }
        }
        v40 := v6[v5]
        _ = fi19
        _ = v40
    }
    v41 := GenS4{F0: -2147483648}
    fmt.Println(v9)
    var v42 float64 = 3.14159
    var v43 float64 = float64(3.14159)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
