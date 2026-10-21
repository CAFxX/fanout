package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 int8 }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 int64 }

func main() {
    var v1 bool = true
    var v2 int = 32767
    v3 := make(map[uint32]uint32)
    v4 := make(map[int]uint32)
    v5 := v4[(^(v2))]
    v6 := []int{}
    switch (v2 >> ((599) & 63)) {
    case -9223372036854775808:
        fmt.Println(v5)
    case -971:
        var v7 int64 = 0
        _ = v7
    default:
        v5 = v5
    }
    v5 -= v5
    var v8 int = v2
    fmt.Println(v1, v5)
    switch ("a" + "") {
    case "a":
        v9 := v4[v8]
        _ = v9
    case "foo bar":
        for fi10 := 0; fi10 < 9; fi10++ {
            fmt.Println(v8)
            switch ("\n\t" + (((((("hello" + ("foo bar" + (("hello" + "\n\t") + "hello"))) + "foo bar") + "\n\t") + "") + ("hello" + "")) + "hello")) {
            case "\n\t":
                if true {
                    var v11 int = 681
                    v4[v11] = (v5 * (v5 - v5))
                    _ = v11
                } else {
                    v12 := v3[v5]
                    _ = v12
                }
            case "a":
                for fi13 := 0; fi13 < 20; fi13++ {
                    var v14 uint8 = 199
                    _ = fi13
                    _ = v14
                }
            case "foo bar":
                _ = len(v6)
            }
            _ = fi10
        }
    case "\n\t":
        v15 := GenS1{F0: (v5 ^ 553), F1: 32767, F2: 20}
        _ = v15
    }
    var v16 uint8 = 127
    for fi17 := 0; fi17 < 3; fi17++ {
        fmt.Println(v2, fi17, v5)
        v16 = (v16 % ((42) | 1))
        _ = fi17
    }
    v18 := GenS1{F0: 2147483648, F1: 1, F2: -109}
    var v19 uint = uint(v16)
    var v20 float32 = -2.25
    if true {
        switch (v8 - (v2 * int(v16))) {
        case 4294967295:
            v2 = v2
        case -449:
            for fi21 := 0; fi21 < 17; fi21++ {
                switch 0 {
                case -2147483648:
                    var v22 int16 = 1
                    _ = v22
                case 846:
                    if v1 {
                        switch v8 {
                        case 616:
                            v5 += (v5 + (v5 * (v5 * 2)))
                        case 1:
                            v23 := GenS1{F0: (v5 / ((42) | 1)), F1: 761, F2: 1}
                            _ = v23
                        }
                        if true {
                            v20 = -2.25
                            switch -434 {
                            case 32768:
                                var v24 uint64 = 1
                                _ = v24
                            case 32767:
                                v5 = 271
                            }
                        } else {
                            var v25 float32 = (v20 - 3.14159)
                            v26 := v4[v8]
                            _ = v25
                            _ = v26
                        }
                    } else {
                        if len(v6) >= 2 {
                            v6 = v6[1:len(v6)]
                        }
                        switch ("foo bar" + "\n\t") {
                        case "\n\t":
                            if v1 {
                                v19 += (^(v19))
                            }
                        case "foo bar":
                            v2 = 1
                        case "hello":
                            v27 := v3[(^(v5))]
                            _ = v27
                        }
                    }
                case 369:
                    var v28 int = -9223372036854775808
                    _ = v28
                default:
                    v2 = v8
                }
                _ = len(v4)
                if (v1 || (!(65535 == 42))) {
                    var v29 int8 = -128
                    switch "\n\t" {
                    case "":
                        v5 = (v5 & 2)
                    case "foo bar":
                        v30 := GenS1{F0: v5, F1: 32767, F2: (v29 * (v29 >> (((v19 ^ 1)) & 63)))}
                        _ = v30
                    case "\n\t":
                        switch "" {
                        case "hello":
                            for fi31 := 0; fi31 < 17; fi31++ {
                                fmt.Println(v5, v8, v1)
                                switch (v2 >> (((v19 & v19)) & 63)) {
                                case 42:
                                    v8 += v8
                                case -751:
                                    var v32 uint8 = v16
                                    _ = v32
                                case -546:
                                    v33 := GenS2{F0: 167, F1: "a"}
                                    _ = v33
                                }
                                var v34 float32 = (-2.25 - v20)
                                v2 = 121
                                v29 -= 100
                                switch int(v19) {
                                case -9223372036854775808:
                                    v35 := GenS3{F0: -671}
                                    _ = v35
                                }
                                _ = fi31
                                _ = v34
                            }
                        case "\n\t":
                            v6 = append(v6, (v2 ^ v2))
                        case "foo bar":
                            var v36 string = ""
                            _ = v36
                        }
                    }
                    if (v1 && (1 <= 1)) {
                        var v37 int = fi21
                        v5 -= v5
                        v29 -= v29
                        _ = v37
                    } else {
                        var v38 int = v8
                        switch "hello" {
                        case "":
                            v2 = (fi21 >> (((v19 ^ v19)) & 63))
                        }
                        _ = v38
                    }
                    _ = v29
                }
                _ = fi21
            }
        default:
            delete(v4, v8)
        }
    }
    v2 -= v8
    var v39 float64 = 1e10
    v16 -= (v16 * 30)
    v20 = (1.5 + v20)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
