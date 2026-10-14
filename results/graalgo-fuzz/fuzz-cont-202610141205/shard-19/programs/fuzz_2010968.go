package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int32 }
type GenS2 struct { F0 uint8; F1 int16 }
type GenS3 struct { F0 uint }

func main() {
    var v1 int8 = -95
    var v2 uint = 662
    var v3 float64 = ((float64(1e10) * -2.25) - 1.5)
    var v4 uint = (v2 / ((18446744073709551615) | 1))
    v5 := []string{((("foo bar" + "") + "hello") + "a"), "hello"}
    v4 = (v4 + (v2 ^ v4))
    v2 = (^(v2))
    switch 127 {
    case 1:
        fmt.Println(v3)
    case 2147483647:
        var v6 int = -768
        _ = v6
    default:
        for fi7 := 0; fi7 < 19; fi7++ {
            v2 -= (v2 + v2)
            v8 := GenS1{F0: (v1 & 42), F1: 1}
            _ = fi7
            _ = v8
        }
    }
    _ = len(v5)
    var v9 float32 = 0.0
    var v10 bool = true
    for fi11 := 0; fi11 < 2; fi11++ {
        if ((!v10) || v10) {
            switch fi11 {
            case 1:
                fmt.Println(fi11)
            case -281:
                if (-1 >= 42) {
                    for fi12 := 0; fi12 < 15; fi12++ {
                        var v13 float32 = (3.14159 / 1.5)
                        v3 = (3.14159 / 1.0)
                        fmt.Println(fi12)
                        var v14 int = (fi12 % ((32767) | 1))
                        _ = fi12
                        _ = v13
                        _ = v14
                    }
                    var v15 string = "hello"
                    _ = v15
                }
            default:
                switch fi11 {
                case 127:
                    if false {
                        if v10 {
                            v4 += (v2 % ((65535) | 1))
                            v16 := GenS1{F0: 100, F1: 741}
                            _ = v16
                        }
                        switch fi11 {
                        case 2147483647:
                            v1 = (v1 - (v1 - v1))
                        case 823:
                            fmt.Println(v9, v10, v3)
                        case 604:
                            var v17 uint8 = 128
                            _ = v17
                        default:
                            if true {
                                v4 = v2
                            } else {
                                fmt.Println(fi11, v10)
                                switch -508 {
                                case 163:
                                    v18 := GenS2{F0: 127, F1: 42}
                                    _ = v18
                                }
                            }
                        }
                        v4 -= (v4 | 422)
                    } else {
                        var v19 int = fi11
                        var v20 int64 = -556
                        _ = v19
                        _ = v20
                    }
                case -9223372036854775808:
                    fmt.Println(v2, fi11)
                }
            }
            var v21 uint = 42
            var v22 int64 = int64(v21)
            _ = v21
            _ = v22
        } else {
            if v10 {
                var v23 uint8 = 3
                v4 = 2
                _ = v23
            } else {
                fmt.Println(v9, v4, v1)
            }
            switch (("hello" + "") + "foo bar") {
            case "":
                for fi24 := 0; fi24 < 15; fi24++ {
                    v25 := GenS1{F0: (v1 & (^(v1))), F1: 42}
                    v3 -= 1e10
                    _ = fi24
                    _ = v25
                }
            case "foo bar":
                for fi26 := 0; fi26 < 11; fi26++ {
                    v3 = v3
                    _ = fi26
                }
            default:
                for fi27 := 0; fi27 < 10; fi27++ {
                    fmt.Println(v4)
                    _ = fi27
                }
            }
        }
        var v28 int16 = 0
        _ = fi11
        _ = v28
    }
    for fi29 := 0; fi29 < 3; fi29++ {
        fmt.Println(v10, fi29, v1)
        _ = fi29
    }
    var v30 bool = ((418 > -2147483648) && (false && (!v10)))
    v31 := GenS1{F0: -6, F1: 2147483647}
    fmt.Println(v1)
    var v32 int16 = 1000
    v33 := GenS3{F0: 42}
    fmt.Println(v3)
    var v34 int32 = 0
    if false {
        v1 = (v1 % ((-128) | 1))
        v34 += v34
    } else {
        v1 = -1
    }
    for fi35 := 0; fi35 < 12; fi35++ {
        var v36 uint = 275
        _ = fi35
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v9
}
