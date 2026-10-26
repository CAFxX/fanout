package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 int64; F1 bool; F2 uint32 }
type GenS3 struct { F0 int16; F1 uint16 }
type GenS4 struct { F0 int64; F1 int16 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int32; F1 int }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 uint32; F1 string }
func genFn0() string {
    var v7 float64 = 1.5
    v8 := []int{127, -2147483648, -2147483648, 32768}
    v7 -= v7
    switch -345 {
    case 128:
        switch ((("foo bar" + ("" + "")) + ("hello" + "")) + "\n\t") {
        case "a":
            switch (("hello" + ("hello" + "foo bar")) + "") {
            case "hello":
                switch 42 {
                case 4294967295:
                    var v9 uint16 = 1
                    _ = v9
                case 9223372036854775807:
                    v8 = append(v8, 127)
                default:
                    switch ((("foo bar" + ("foo bar" + "hello")) + ("" + "a")) + "a") {
                    case "\n\t":
                        if (false || false) {
                            var v10 int8 = 84
                            fmt.Println(v10, v7)
                            _ = v10
                        } else {
                            if len(v8) >= 2 {
                                v8 = v8[1:len(v8)]
                            }
                            v7 += v7
                        }
                    case "foo bar":
                        v7 += float64(0.0)
                    default:
                        for fi11 := 0; fi11 < 4; fi11++ {
                            v12 := GenS1{F0: false, F1: 857}
                            v13 := GenS2{F0: 0, F1: (!(547 < 1)), F2: 2147483648}
                            var v14 uint8 = 128
                            _ = fi11
                            _ = v12
                            _ = v13
                            _ = v14
                        }
                    }
                }
            case "\n\t":
                for fi15 := 0; fi15 < 8; fi15++ {
                    for fi16 := 0; fi16 < 14; fi16++ {
                        if (51 == 26) {
                            var v17 uint32 = 42
                            _ = v17
                        } else {
                            v7 = v7
                            v18 := GenS2{F0: 0, F1: true, F2: 399}
                            _ = v18
                        }
                        _ = fi16
                    }
                    _ = fi15
                }
            }
        case "foo bar":
            for fi19 := 0; fi19 < 8; fi19++ {
                v20 := GenS3{F0: -1, F1: 2}
                var v21 int64 = 0
                v22 := GenS4{F0: (v21 + (v21 / ((v21) | 1))), F1: 1}
                _ = fi19
                _ = v20
                _ = v21
                _ = v22
            }
        case "\n\t":
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
        }
    }
    _ = v7
    _ = v8
    return "foo bar"
}

func main() {
    v1 := make(map[int]string)
    var v2 bool = (0.1 < 0.0)
    var v3 float32 = 3.14159
    var v4 string = ""
    var v5 float64 = -2.25
    var v6 uint = 741
    v6 -= v6
    for fi23 := 0; fi23 < 6; fi23++ {
        v24 := GenS5{F0: 32768}
        _ = fi23
        _ = v24
    }
    switch -349 {
    case 128:
        v25 := v1[-107]
        _ = v25
    }
    fmt.Println(v2)
    v3 = (v3 * float32((1.5 / v5)))
    for fi26 := 0; fi26 < 20; fi26++ {
        if false {
            var v27 float64 = (-2.25 + v5)
            v2 = false
            v4 = ""
            _ = v27
        }
        _ = fi26
    }
    v28 := []int{-486, 2147483647, 2147483647, 358}
    for fi29 := 0; fi29 < 16; fi29++ {
        delete(v1, 128)
        _ = fi29
    }
    for fi30 := 0; fi30 < 10; fi30++ {
        var v31 int64 = 561
        for fi32 := 0; fi32 < 17; fi32++ {
            var v33 uint = (^(v6))
            v34 := GenS6{F0: -2147483648, F1: fi32}
            v35 := GenS7{F0: -988}
            _ = fi32
            _ = v33
            _ = v34
            _ = v35
        }
        _ = fi30
        _ = v31
    }
    if (4294967295 != 325) {
        for fi36 := 0; fi36 < 6; fi36++ {
            for fi37 := 0; fi37 < 16; fi37++ {
                var v38 uint16 = 2
                v2 = ((4294967295 >= 1) || (v2 || ((!v2) || v2)))
                _ = fi37
                _ = v38
            }
            v3 = v3
            _ = fi36
        }
    }
    v39 := GenS5{F0: -534}
    delete(v1, 4294967295)
    for fi40 := 0; fi40 < 15; fi40++ {
        delete(v1, fi40)
        v41 := GenS8{F0: 42}
        _ = fi40
        _ = v41
    }
    v28 = append(v28, 1)
    v6 = (v6 << ((v6) & 63))
    if (false && v2) {
        var v42 uint64 = 2
        if v2 {
            v43 := GenS9{F0: 0, F1: "foo bar"}
            _ = len(v28)
            _ = v43
        } else {
            fmt.Println(v5)
            switch 127 {
            case -732:
                v3 -= (float32(0.0) + v3)
            default:
                v44 := GenS6{F0: 42, F1: 760}
                _ = v44
            }
            v2 = ((-459 < 197) && false)
        }
        for fi45 := 0; fi45 < 11; fi45++ {
            v42 = v42
            _ = fi45
        }
        _ = v42
    }
    fmt.Println(v2)
    switch v4 {
    case "foo bar":
        v5 = -2.25
    case "hello":
        v4 = "\n\t"
    default:
        v46 := GenS7{F0: 950}
        _ = v46
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
