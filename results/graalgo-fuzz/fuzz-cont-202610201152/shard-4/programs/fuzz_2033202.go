package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 int32 }
type GenS2 struct { F0 int32; F1 bool; F2 uint32 }
type GenS3 struct { F0 uint8; F1 uint64; F2 uint64 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint16 }

func main() {
    var v1 int8 = 111
    var v2 int16 = 1
    var v3 uint64 = 999
    v4 := make(map[string]uint32)
    var v5 bool = (true && false)
    var v6 bool = v5
    v5 = false
    switch "\n\t" {
    case "a":
        delete(v4, "hello")
    }
    v1 = v1
    switch "a" {
    case "hello":
        var v7 int32 = 2147483647
        _ = v7
    }
    v8 := GenS1{F0: 91, F1: 231, F2: 390}
    v1 += v1
    var v9 uint8 = 255
    for fi10 := 0; fi10 < 15; fi10++ {
        switch 32768 {
        case 32767:
            switch "\n\t" {
            case "\n\t":
                if ((-1 > 0) || v5) {
                    fmt.Println(v9, v5, v1)
                    switch "hello" {
                    case "\n\t":
                        fmt.Println(fi10, v9, fi10)
                    case "hello":
                        v11 := GenS1{F0: (v9 & v9), F1: (v9 - v9), F2: -1}
                        _ = v11
                    case "a":
                        if true {
                            if (v1 <= (v1 / ((v1) | 1))) {
                                if false {
                                    var v12 uint16 = 42
                                    _ = v12
                                } else {
                                    if (-9223372036854775808 <= 42) {
                                        v13 := GenS2{F0: -1, F1: (!(true || true)), F2: 94}
                                        var v14 int64 = 42
                                        _ = v13
                                        _ = v14
                                    } else {
                                        for fi15 := 0; fi15 < 2; fi15++ {
                                            fmt.Println(v9, v3)
                                            fmt.Println(fi10, v2)
                                            switch "a" {
                                            case "hello":
                                                fmt.Println(v1, v6)
                                            case "foo bar":
                                                var v16 bool = v5
                                                var v17 uint64 = 759
                                                v18 := GenS1{F0: 205, F1: v9, F2: 0}
                                                _ = v16
                                                _ = v17
                                                _ = v18
                                            default:
                                                fmt.Println(fi15, v2)
                                            }
                                            _ = fi15
                                        }
                                        if true {
                                            var v19 uint32 = 1
                                            _ = v19
                                        } else {
                                            var v20 int8 = (-(v1))
                                            v21 := GenS3{F0: 83, F1: v3, F2: v3}
                                            var v22 float32 = 3.14159
                                            _ = v20
                                            _ = v21
                                            _ = v22
                                        }
                                    }
                                }
                                delete(v4, (("hello" + "hello") + "hello"))
                            } else {
                                switch (fi10 >> ((76) & 63)) {
                                case 121:
                                    v23 := []int{(fi10 >> ((1) & 63)), (fi10 + fi10)}
                                    _ = v23
                                }
                                for fi24 := 0; fi24 < 16; fi24++ {
                                    var v25 float64 = 0.1
                                    v26 := []int{-2147483648}
                                    var v27 uint64 = v3
                                    switch (("foo bar" + ("\n\t" + "\n\t")) + "hello") {
                                    case "a":
                                        fmt.Println(v2, v1)
                                        if ((("\n\t" + "") + ("a" + "a")) > (("a" + "\n\t") + "hello")) {
                                            switch (fi24 << ((65535) & 63)) {
                                            case -420:
                                                if v6 {
                                                    v25 += (v25 + 3.14159)
                                                    v25 += v25
                                                    v3 -= (v3 | (v27 | v27))
                                                } else {
                                                    fmt.Println(fi24, v5)
                                                }
                                            default:
                                                if v6 {
                                                    fmt.Println(fi24, v27)
                                                    v28 := GenS1{F0: (v9 << ((1) & 63)), F1: (v9 % (((v9 & (v9 % ((v9) | 1)))) | 1)), F2: 0}
                                                    v2 -= -32768
                                                    _ = v28
                                                } else {
                                                    v29 := GenS1{F0: (^(v9)), F1: 137, F2: 42}
                                                    var v30 float64 = (v25 + v25)
                                                    if v6 {
                                                        var v31 int32 = -2147483648
                                                        switch (fi24 * 0) {
                                                        case 0:
                                                            v4[(("hello" + "foo bar") + "hello")] = 42
                                                        case 4294967295:
                                                            if ("\n\t" != (("a" + "hello") + "")) {
                                                                v1 = (v1 + (-(v1)))
                                                                _ = len(v4)
                                                                if ((42 >= 65535) && v6) {
                                                                    v1 = -66
                                                                }
                                                            }
                                                            var v32 uint8 = (^(v9))
                                                            _ = v32
                                                        default:
                                                            v30 += 0.1
                                                        }
                                                        v33 := GenS4{F0: int8(v2)}
                                                        _ = v31
                                                        _ = v33
                                                    } else {
                                                        if v6 {
                                                            v34 := GenS4{F0: (v1 * (v1 >> ((18446744073709551615) & 63)))}
                                                            var v35 bool = v5
                                                            var v36 int = fi10
                                                            _ = v34
                                                            _ = v35
                                                            _ = v36
                                                        }
                                                    }
                                                    _ = v29
                                                    _ = v30
                                                }
                                            }
                                            var v37 uint16 = 0
                                            _ = len(v4)
                                            _ = v37
                                        } else {
                                            switch (^(fi10)) {
                                            case 32767:
                                                if v5 {
                                                    v2 = 223
                                                }
                                            default:
                                                if v6 {
                                                    v38 := GenS5{F0: ("\n\t" + "foo bar")}
                                                    var v39 string = "foo bar"
                                                    if len(v26) > 0 {
                                                        v26[0] = fi10
                                                    }
                                                    v3 += v3
                                                    _ = v38
                                                    _ = v39
                                                } else {
                                                    var v40 uint = 609
                                                    v41 := GenS6{F0: 845}
                                                    _ = v40
                                                    _ = v41
                                                }
                                            }
                                        }
                                        var v42 int64 = 1
                                        if (!(v9 >= (v9 * (v9 % ((v9) | 1))))) {
                                            v6 = (!false)
                                            v6 = false
                                        }
                                        var v43 int32 = -2147483648
                                        fmt.Println(v25)
                                        var v44 uint64 = (v3 * v27)
                                        _ = v42
                                        _ = v43
                                        _ = v44
                                    }
                                    _ = fi24
                                    _ = v25
                                    _ = v26
                                    _ = v27
                                }
                            }
                            fmt.Println(v2, v9)
                            v6 = ((v6 || ((^(v2)) >= v2)) || (!((fi10 + 562) <= fi10)))
                        }
                    default:
                        v6 = (287 == 515)
                    }
                }
            case "":
                for fi45 := 0; fi45 < 20; fi45++ {
                    v3 = (v3 * v3)
                    v6 = true
                    _ = fi45
                }
            default:
                v5 = false
            }
        case 9223372036854775807:
            if true {
                v9 += 127
                var v46 bool = ((-625 > 9223372036854775807) && ((2147483647 == 910) || (4294967295 > 301)))
                var v47 int64 = 0
                _ = v46
                _ = v47
            }
        default:
            v48 := []int{}
            _ = v48
        }
        v5 = true
        fmt.Println(v1, fi10)
        _ = fi10
    }
    for fi49 := 0; fi49 < 2; fi49++ {
        if v6 {
            var v50 bool = ((v3 / ((v3) | 1)) > v3)
            for fi51 := 0; fi51 < 18; fi51++ {
                v5 = (true && (v2 > -32768))
                if false {
                    v3 -= v3
                    fmt.Println(fi51, v2)
                    v52 := []int{-972, fi49}
                    _ = v52
                }
                _ = fi51
            }
            _ = v50
        }
        _ = fi49
    }
    fmt.Println(v1, v1)
    v5 = true
    v6 = ("foo bar" == "\n\t")
    for fi53 := 0; fi53 < 7; fi53++ {
        for fi54 := 0; fi54 < 1; fi54++ {
            v6 = (0 < 4294967295)
            var v55 uint32 = 143
            if v5 {
                v9 -= (^(v9))
                fmt.Println(v5, fi54)
            } else {
                fmt.Println(v5)
                v56 := []int{(fi54 ^ (fi53 >> ((2) & 63))), (fi53 % ((fi53) | 1)), (fi54 / ((32767) | 1))}
                _ = v56
            }
            _ = fi54
            _ = v55
        }
        v57 := GenS2{F0: 317, F1: (false || true), F2: 271}
        _ = fi53
        _ = v57
    }
    var v58 float32 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v58:", v58)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v6
    _ = v8
    _ = v9
}
