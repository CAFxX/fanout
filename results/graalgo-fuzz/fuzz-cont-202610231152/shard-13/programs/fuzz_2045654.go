package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16; F2 bool }
type GenS2 struct { F0 uint; F1 int32 }
type GenS3 struct { F0 uint8; F1 uint16; F2 uint16 }
type GenS4 struct { F0 int8; F1 uint16 }
type GenS5 struct { F0 uint8; F1 uint8; F2 int }

func main() {
    var v1 float64 = 3.14159
    var v2 string = ("a" + ("hello" + "foo bar"))
    var v3 uint16 = 65535
    for fi4 := 0; fi4 < 20; fi4++ {
        v5 := GenS1{F0: -1, F1: -1, F2: false}
        var v6 string = (v2 + "hello")
        _ = fi4
        _ = v5
        _ = v6
    }
    v7 := GenS2{F0: 1, F1: 837}
    switch -2147483648 {
    case -9223372036854775808:
        v2 = v2
    }
    var v8 string = ("hello" + v2)
    v9 := []int{42}
    fmt.Println(v1)
    if (2147483647 >= 42) {
        var v10 uint8 = 170
        v1 = (1e10 * (-2.25 - v1))
        _ = v10
    } else {
        for fi11 := 0; fi11 < 7; fi11++ {
            switch fi11 {
            case -898:
                for fi12 := 0; fi12 < 16; fi12++ {
                    var v13 int8 = -128
                    var v14 uint8 = 2
                    var v15 uint = 18446744073709551615
                    _ = fi12
                    _ = v13
                    _ = v14
                    _ = v15
                }
            case -780:
                v2 = v8
            case 4294967295:
                if false {
                    fmt.Println(fi11)
                    for fi16 := 0; fi16 < 2; fi16++ {
                        fmt.Println(v3, fi11, v1)
                        if len(v9) >= 2 {
                            v9 = v9[1:len(v9)]
                        }
                        var v17 int8 = 0
                        var v18 uint8 = 2
                        v1 = 0.1
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                }
            }
            v19 := GenS1{F0: -229, F1: 0, F2: ((false && (!(false || (false && false)))) && true)}
            fmt.Println(v2, v2)
            _ = fi11
            _ = v19
        }
        v20 := GenS3{F0: 1, F1: (v3 / ((v3) | 1)), F2: (v3 % ((v3) | 1))}
        _ = v20
    }
    var v21 string = v8
    v1 += (float64(0.1) / ((0.0 * (1e10 - (((1.5 + ((1.5 * v1) * v1)) - (v1 * 3.14159)) * v1))) / 1.5))
    v8 = v8
    fmt.Println(v1)
    v21 = v2
    if len(v9) > 0 {
        v9[0] = 42
    }
    v2 = v2
    v22 := GenS3{F0: 255, F1: v3, F2: v3}
    v8 = "hello"
    var v23 bool = (127 <= 44)
    if (false && v23) {
        switch 128 {
        case 1:
            var v24 int16 = -773
            _ = v24
        case 42:
            if v23 {
                v3 = v3
                var v25 int8 = -43
                switch v8 {
                case "foo bar":
                    var v26 int64 = 42
                    _ = v26
                case "a":
                    if v23 {
                        v27 := GenS3{F0: 2, F1: 32767, F2: v3}
                        v23 = (!((!(v25 > (v25 + (v25 % ((v25) | 1))))) || v23))
                        var v28 float32 = (3.14159 + 0.1)
                        _ = v27
                        _ = v28
                    }
                case "hello":
                    fmt.Println(v25, v1)
                default:
                    v3 = v3
                }
                _ = v25
            } else {
                if (v8 > ("hello" + "hello")) {
                    v29 := GenS1{F0: -1, F1: 539, F2: v23}
                    for fi30 := 0; fi30 < 12; fi30++ {
                        v1 -= float64((1e10 * float32(v1)))
                        if true {
                            if ((fi30 % (((fi30 << ((18446744073709551615) & 63))) | 1)) >= (fi30 / (((^(fi30))) | 1))) {
                                fmt.Println(v8, v8)
                            }
                            if v23 {
                                switch v8 {
                                case "hello":
                                    var v31 float64 = (v1 / v1)
                                    _ = v31
                                case "a":
                                    var v32 int16 = -32768
                                    _ = v32
                                case "":
                                    fmt.Println(v21)
                                default:
                                    var v33 int8 = -14
                                    _ = v33
                                }
                                if false {
                                    fmt.Println(v23)
                                }
                            } else {
                                v34 := GenS2{F0: 458, F1: -1}
                                _ = v34
                            }
                        } else {
                            var v35 int8 = 1
                            if true {
                                var v36 int32 = -2147483648
                                var v37 int32 = v36
                                _ = v36
                                _ = v37
                            }
                            switch fi30 {
                            case 42:
                                v38 := GenS4{F0: (v35 & v35), F1: 2}
                                _ = v38
                            }
                            _ = v35
                        }
                        _ = fi30
                    }
                    _ = v29
                }
            }
        }
        v23 = (2 > 438)
        for fi39 := 0; fi39 < 16; fi39++ {
            switch (v21 + "a") {
            case "hello":
                switch 916 {
                case 32768:
                    v2 = v21
                }
            case "foo bar":
                v1 = (v1 * 0.0)
            case "\n\t":
                if (!(4294967295 <= 42)) {
                    v40 := GenS5{F0: 1, F1: 42, F2: -375}
                    switch (fi39 | (fi39 >> ((925) & 63))) {
                    case 127:
                        var v41 float32 = 1e10
                        _ = v41
                    default:
                        var v42 int32 = -2147483648
                        _ = v42
                    }
                    _ = v40
                } else {
                    v21 = (((v8 + ("\n\t" + ("hello" + v21))) + v21) + v21)
                    fmt.Println(v2, v1, v21)
                }
            }
            _ = fi39
        }
    } else {
        v43 := GenS1{F0: -1, F1: -890, F2: (v3 > (^(v3)))}
        _ = v43
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
