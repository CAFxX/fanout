package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 int32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int; F1 uint64; F2 uint16 }

func main() {
    var v1 uint16 = 42
    var v2 uint16 = v1
    v3 := make(map[uint32]string)
    var v4 string = (("foo bar" + "hello") + "a")
    var v5 uint64 = 31
    switch v4 {
    case "a":
        fmt.Println(v2)
    case "hello":
        if (false || false) {
            v4 = v4
            if (!(1 >= 2147483647)) {
                switch 32767 {
                case 816:
                    v4 = ((v4 + v4) + (v4 + v4))
                case 1:
                    for fi6 := 0; fi6 < 11; fi6++ {
                        for fi7 := 0; fi7 < 16; fi7++ {
                            var v8 bool = false
                            fmt.Println(v1)
                            _ = fi7
                            _ = v8
                        }
                        v3[578] = v4
                        var v9 float32 = (((((1e10 / 1.0) / float32(3.14159)) / float32(1e10)) + 0.0) / ((((-2.25 / 0.1) - (0.1 * 1e10)) / (-2.25 + 1.5)) * 3.14159))
                        _ = fi6
                        _ = v9
                    }
                case 128:
                    v4 = v4
                default:
                    v2 = (v1 / (((v2 * (v2 / (((v1 >> ((42) & 63))) | 1)))) | 1))
                }
                fmt.Println(v1, v4, v5)
            } else {
                for fi10 := 0; fi10 < 19; fi10++ {
                    v11 := []int{}
                    _ = fi10
                    _ = v11
                }
            }
            for fi12 := 0; fi12 < 4; fi12++ {
                v13 := []int{(fi12 * (fi12 - fi12)), 856, fi12, 32767}
                var v14 int = 1
                if (478 >= 18446744073709551615) {
                    fmt.Println(v4)
                    v1 = 0
                }
                _ = fi12
                _ = v13
                _ = v14
            }
        } else {
            v2 -= (v2 | v1)
        }
    case "foo bar":
        fmt.Println(v5, v5)
    default:
        var v15 int32 = 0
        _ = v15
    }
    v2 = 255
    for fi16 := 0; fi16 < 11; fi16++ {
        v17 := GenS1{F0: v4, F1: (fi16 << ((65535) & 63)), F2: 1}
        _ = fi16
        _ = v17
    }
    v18 := GenS2{F0: -937}
    switch 2147483647 {
    case 42:
        switch -57 {
        case -1:
            v5 -= v5
        default:
            v19 := GenS1{F0: ("foo bar" + "foo bar"), F1: 4294967295, F2: 122}
            _ = v19
        }
    }
    v2 = (v1 << ((4294967295) & 63))
    fmt.Println(v4, v1, v4)
    var v20 int8 = -1
    switch 2147483647 {
    case -906:
        v21 := []int{}
        _ = v21
    case 42:
        v22 := []int{42, -9223372036854775808, 2147483647}
        _ = v22
    case -2147483648:
        switch "" {
        case "hello":
            if false {
                if true {
                    switch -9223372036854775808 {
                    case 32767:
                        var v23 uint = 756
                        _ = v23
                    case 1:
                        for fi24 := 0; fi24 < 4; fi24++ {
                            v5 = (v5 | 18446744073709551615)
                            for fi25 := 0; fi25 < 16; fi25++ {
                                if ((!(529 == 522)) && false) {
                                    var v26 int16 = 725
                                    v4 = v4
                                    _ = v26
                                }
                                _ = fi25
                            }
                            v20 = 0
                            _ = fi24
                        }
                    case 32768:
                        var v27 int32 = 1
                        _ = v27
                    }
                    var v28 uint32 = 1
                    _ = v28
                }
                for fi29 := 0; fi29 < 8; fi29++ {
                    v1 = 65535
                    _ = fi29
                }
                var v30 float64 = (float64(1.5) + 3.14159)
                _ = v30
            } else {
                if (1 >= 0) {
                    v31 := GenS1{F0: "", F1: 790, F2: -2147483648}
                    _ = v31
                }
                v32 := v3[1]
                _ = v32
            }
        case "a":
            v33 := GenS1{F0: ("hello" + "foo bar"), F1: 128, F2: 710}
            _ = v33
        case "foo bar":
            for fi34 := 0; fi34 < 5; fi34++ {
                var v35 int16 = 32767
                v36 := GenS2{F0: (^(v35))}
                _ = fi34
                _ = v35
                _ = v36
            }
        }
    }
    var v37 string = ((v4 + (("foo bar" + v4) + "\n\t")) + v4)
    if true {
        var v38 int16 = -867
        fmt.Println(v20)
        var v39 uint64 = 301
        _ = v38
        _ = v39
    } else {
        v40 := GenS3{F0: 4294967295, F1: (v5 % (((^(v5))) | 1)), F2: v1}
        if ((^(v1)) <= (v2 ^ 42)) {
            v3[4294967295] = (v4 + v4)
            switch (v4 + "hello") {
            case "\n\t":
                delete(v3, 65535)
            case "hello":
                var v41 int16 = 604
                _ = v41
            default:
                for fi42 := 0; fi42 < 13; fi42++ {
                    fmt.Println(fi42)
                    var v43 uint8 = 42
                    _ = fi42
                    _ = v43
                }
            }
        }
        fmt.Println(v5, v5)
        _ = v40
    }
    _ = len(v3)
    var v44 float32 = (3.14159 * (1.5 - ((1.5 * 1e10) / 1e10)))
    for fi45 := 0; fi45 < 20; fi45++ {
        delete(v3, 471)
        v5 -= v5
        v46 := GenS3{F0: (fi45 << ((366) & 63)), F1: (^(v5)), F2: (v2 << ((0) & 63))}
        _ = fi45
        _ = v46
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v2
    _ = v20
    _ = v3
    _ = v37
    _ = v4
    _ = v44
    _ = v5
}
