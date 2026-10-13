package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int8; F1 int16; F2 int16 }
type GenS4 struct { F0 int64; F1 int16; F2 uint32 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int }

func main() {
    var v1 bool = ((true || true) || (false || false))
    var v2 int8 = -121
    var v3 int16 = 42
    fmt.Println(v1, v2)
    var v4 uint = 18446744073709551615
    var v5 int = 1
    if (25 <= -76) {
        switch (v5 ^ (v5 / ((127) | 1))) {
        case 32768:
            var v6 int64 = 498
            _ = v6
        }
        v7 := GenS1{F0: 1, F1: (^(v4))}
        _ = v7
    } else {
        v8 := []int{(v5 * v5), -9223372036854775808, 42, (-(v5))}
        _ = v8
    }
    var v9 bool = (v2 < (v2 << ((255) & 63)))
    if v9 {
        var v10 uint64 = 595
        if (!((2 == 65535) || false)) {
            var v11 float32 = 0.0
            _ = v11
        }
        v12 := []int{v5, -1}
        _ = v10
        _ = v12
    } else {
        var v13 float32 = 0.0
        fmt.Println(v5, v5)
        v14 := []int{(v5 - (v5 >> ((4294967295) & 63)))}
        _ = v13
        _ = v14
    }
    var v15 uint64 = 33
    if v1 {
        for fi16 := 0; fi16 < 14; fi16++ {
            for fi17 := 0; fi17 < 2; fi17++ {
                if (!(("hello" <= "foo bar") || v9)) {
                    v15 -= (^(v15))
                    fmt.Println(v15)
                    fmt.Println(v4)
                } else {
                    v2 -= (v2 >> (((v4 << ((v4) & 63))) & 63))
                }
                var v18 int8 = v2
                _ = fi17
                _ = v18
            }
            _ = fi16
        }
        for fi19 := 0; fi19 < 5; fi19++ {
            var v20 string = ("foo bar" + ("foo bar" + "a"))
            if v1 {
                for fi21 := 0; fi21 < 5; fi21++ {
                    var v22 int64 = -109
                    v23 := []int{(fi19 / (((fi21 << ((42) & 63))) | 1)), -2147483648, fi19, -2147483648}
                    _ = fi21
                    _ = v22
                    _ = v23
                }
                for fi24 := 0; fi24 < 7; fi24++ {
                    v25 := []int{fi19, fi19}
                    _ = fi24
                    _ = v25
                }
                v4 = v4
            }
            if (v9 || (!(1.5 == 3.14159))) {
                v26 := GenS1{F0: -1, F1: 723}
                for fi27 := 0; fi27 < 9; fi27++ {
                    fmt.Println(v2, v20)
                    var v28 float32 = 0.0
                    _ = fi27
                    _ = v28
                }
                if false {
                    v3 = (v3 & (v3 / (((v3 * int16(v15))) | 1)))
                }
                _ = v26
            } else {
                v15 = 42
                v29 := GenS2{F0: v20}
                v15 -= v15
                _ = v29
            }
            _ = fi19
            _ = v20
        }
        if (-2.25 > 0.0) {
            for fi30 := 0; fi30 < 17; fi30++ {
                if false {
                    var v31 bool = (0.0 >= -2.25)
                    _ = v31
                }
                v32 := GenS1{F0: 55, F1: (^(v4))}
                var v33 float32 = 1.5
                _ = fi30
                _ = v32
                _ = v33
            }
            var v34 float64 = (1.5 - (3.14159 + 0.1))
            _ = v34
        }
    }
    v35 := GenS2{F0: "\n\t"}
    var v36 bool = (((v3 > (v3 ^ (v3 + (^(v3))))) && (v5 == 1)) || v1)
    if false {
        for fi37 := 0; fi37 < 4; fi37++ {
            v38 := []int{(v5 ^ (v5 * v5))}
            var v39 float32 = (1.5 - 3.14159)
            v40 := GenS1{F0: 468, F1: 104}
            _ = fi37
            _ = v38
            _ = v39
            _ = v40
        }
        if true {
            v41 := GenS3{F0: (v2 ^ v2), F1: (v3 | (v3 % ((v3) | 1))), F2: v3}
            var v42 int8 = -1
            _ = v41
            _ = v42
        }
    }
    var v43 uint = 226
    var v44 int64 = -9223372036854775808
    for fi45 := 0; fi45 < 15; fi45++ {
        var v46 int64 = (v44 ^ v44)
        _ = fi45
        _ = v46
    }
    v5 += (v5 / ((v5) | 1))
    if v1 {
        v2 = -128
        if v36 {
            switch "" {
            case "":
                fmt.Println(v44, v4)
            case "a":
                if true {
                    if (!true) {
                        if (32768 != v5) {
                            v47 := GenS2{F0: "foo bar"}
                            var v48 uint32 = 2147483648
                            var v49 int32 = -1
                            _ = v47
                            _ = v48
                            _ = v49
                        } else {
                            fmt.Println(v1, v5, v43)
                            v3 = v3
                            fmt.Println(v5)
                        }
                        v50 := []int{}
                        _ = v50
                    } else {
                        var v51 int16 = v3
                        _ = v51
                    }
                    for fi52 := 0; fi52 < 5; fi52++ {
                        for fi53 := 0; fi53 < 16; fi53++ {
                            fmt.Println(v2, v15)
                            v15 += 338
                            _ = fi53
                        }
                        v43 += 970
                        v44 += (-(v44))
                        _ = fi52
                    }
                    fmt.Println(v36, v9)
                } else {
                    fmt.Println(v3)
                    fmt.Println(v9, v1, v36)
                }
            }
        } else {
            switch v5 {
            case -831:
                v54 := GenS4{F0: (^(v44)), F1: 32767, F2: 4294967295}
                _ = v54
            case -1:
                v55 := GenS1{F0: (-(v44)), F1: v4}
                _ = v55
            }
            if v1 {
                switch "hello" {
                case "hello":
                    v56 := GenS1{F0: 1, F1: 4294967295}
                    _ = v56
                case "":
                    var v57 uint8 = 2
                    _ = v57
                }
                v43 = (v43 % ((65535) | 1))
                v58 := []int{-1, v5, v5}
                _ = v58
            }
            v59 := []int{42, (v5 - (v5 >> ((156) & 63))), v5}
            _ = v59
        }
        switch "\n\t" {
        case "\n\t":
            if (!(65535 < uint32(v15))) {
                v5 = (v5 * -1)
            } else {
                var v60 int64 = (v44 ^ v44)
                var v61 int32 = 42
                var v62 string = "hello"
                _ = v60
                _ = v61
                _ = v62
            }
        case "a":
            v63 := GenS5{F0: (v5 << ((v4) & 63))}
            _ = v63
        case "foo bar":
            v64 := GenS6{F0: 273}
            _ = v64
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v9
}
