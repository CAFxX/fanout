package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 uint32 }
type GenS2 struct { F0 bool; F1 uint64 }
func genFn0() string {
    if true {
        var v7 string = "\n\t"
        _ = v7
    }
    var v8 int16 = 42
    if (0.1 > 0.0) {
        if false {
            if (true && (!true)) {
                switch "" {
                case "foo bar":
                    v9 := []int{}
                    _ = v9
                }
                var v10 float64 = 0.1
                _ = v10
            } else {
                var v11 int64 = -9223372036854775808
                _ = v11
            }
            var v12 uint32 = 2147483648
            _ = v12
        } else {
            var v13 int32 = 42
            v14 := []int{765, 127}
            v13 -= (v13 ^ -2147483648)
            _ = v13
            _ = v14
        }
        var v15 uint32 = 42
        _ = v15
    }
    var v16 string = "hello"
    _ = v16
    _ = v8
    return ""
}

func main() {
    v1 := make(map[uint32]string)
    var v2 string = "hello"
    var v3 int16 = 0
    var v4 uint = 0
    var v5 uint32 = 1
    var v6 int = -32
    var v17 string = v2
    if false {
        if ((((v2 + ("foo bar" + "foo bar")) + "foo bar") == ((v2 + v17) + (v2 + v17))) || (false || (true && (!true)))) {
            var v18 int = (v6 + 492)
            fmt.Println(v6, v4, v2)
            var v19 uint8 = 0
            _ = v18
            _ = v19
        }
        var v20 uint64 = 0
        for fi21 := 0; fi21 < 14; fi21++ {
            v2 = v17
            _ = fi21
        }
        _ = v20
    }
    v6 += (v6 & 18)
    var v22 int64 = -9223372036854775808
    switch int(v5) {
    case -2147483648:
        for fi23 := 0; fi23 < 16; fi23++ {
            v24 := GenS1{F0: v6, F1: (18446744073709551615 >= 579), F2: 2147483648}
            v25 := GenS1{F0: 680, F1: (0.1 >= (-2.25 / float32((1e10 * -2.25)))), F2: (^(v5))}
            _ = fi23
            _ = v24
            _ = v25
        }
    default:
        if true {
            var v26 int8 = -128
            var v27 uint16 = 1
            var v28 uint64 = 4294967295
            _ = v26
            _ = v27
            _ = v28
        } else {
            if false {
                if false {
                    fmt.Println(v6)
                    for fi29 := 0; fi29 < 16; fi29++ {
                        switch v17 {
                        case "foo bar":
                            var v30 uint8 = 0
                            _ = v30
                        case "hello":
                            v5 = (v5 & uint32(v4))
                        }
                        var v31 int8 = -1
                        _ = fi29
                        _ = v31
                    }
                } else {
                    var v32 float64 = float64(-2.25)
                    v33 := GenS1{F0: (v6 & v6), F1: (1 <= v5), F2: v5}
                    var v34 uint8 = 227
                    _ = v32
                    _ = v33
                    _ = v34
                }
                v3 = (v3 << ((65535) & 63))
                fmt.Println(v5, v4)
            }
            if true {
                v2 = "foo bar"
            } else {
                if ((3.14159 / -2.25) < (0.0 + 0.0)) {
                    for fi35 := 0; fi35 < 3; fi35++ {
                        var v36 uint8 = 155
                        v37 := []int{(v6 % ((9223372036854775807) | 1)), (v6 & (^(fi35))), v6}
                        _ = fi35
                        _ = v36
                        _ = v37
                    }
                } else {
                    fmt.Println(v3, v6)
                    var v38 float32 = 1e10
                    v39 := GenS2{F0: ((v6 / (((^(v6))) | 1)) != (v6 & int(v4))), F1: 16}
                    _ = v38
                    _ = v39
                }
            }
            v5 = 568
        }
    }
    if ((false && true) || false) {
        v5 = (v5 + (v5 << (((^(v4))) & 63)))
        switch v6 {
        case 503:
            v40 := GenS2{F0: true, F1: 2}
            _ = v40
        default:
            v17 = "\n\t"
        }
    } else {
        v17 = "hello"
        switch "foo bar" {
        case "":
            fmt.Println(v2)
        case "\n\t":
            var v41 bool = true
            _ = v41
        default:
            for fi42 := 0; fi42 < 17; fi42++ {
                fmt.Println(v3, v4, v6)
                _ = fi42
            }
        }
    }
    v43 := []int{}
    switch v6 {
    case 42:
        fmt.Println(v4, v4, v22)
    default:
        if ((42 <= 128) || true) {
            fmt.Println(v2)
            _ = len(v1)
        } else {
            if (float64(1.5) != ((3.14159 - 1e10) / 1.5)) {
                _ = len(v43)
            } else {
                v44 := GenS2{F0: (0 <= 1), F1: 42}
                _ = v44
            }
            switch (v6 + (^(v6))) {
            case 1:
                var v45 string = v17
                _ = v45
            case -9223372036854775808:
                v46 := GenS1{F0: v6, F1: (!true), F2: v5}
                _ = v46
            }
        }
    }
    for fi47 := 0; fi47 < 12; fi47++ {
        switch -1 {
        case 0:
            v2 = v17
        default:
            var v48 uint32 = (v5 / ((1) | 1))
            _ = v48
        }
        _ = len(v43)
        _ = fi47
    }
    v4 -= (v4 << (((v4 & 4294967295)) & 63))
    v6 -= v6
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v43:", len(v43), v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v6
}
