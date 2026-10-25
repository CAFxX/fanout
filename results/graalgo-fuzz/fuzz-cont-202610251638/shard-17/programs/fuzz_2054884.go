package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int32; F1 uint; F2 bool }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint32; F1 bool; F2 uint8 }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int64; F1 uint16; F2 uint16 }
func genFn0(p0 uint64) string {
    switch "foo bar" {
    case "\n\t":
        if false {
            v7 := GenS1{F0: (p0 + (p0 % (((p0 / ((p0) | 1))) | 1))), F1: 127}
            _ = v7
        } else {
            var v8 int64 = 625
            fmt.Println(v8, v8)
            _ = v8
        }
    case "hello":
        v9 := GenS2{F0: 9223372036854775807}
        _ = v9
    case "a":
        if (false && (1 < -2147483648)) {
            v10 := []int{9223372036854775807, -366, 809, 128}
            _ = v10
        } else {
            for fi11 := 0; fi11 < 19; fi11++ {
                v12 := []int{(fi11 - (fi11 - fi11)), 4294967295}
                _ = len(v12)
                var v13 int = fi11
                _ = fi11
                _ = v12
                _ = v13
            }
            switch 32768 {
            case 4294967295:
                if false {
                    p0 -= (p0 ^ (p0 - p0))
                    v14 := GenS2{F0: 1}
                    v15 := []int{-932, -9223372036854775808}
                    _ = v14
                    _ = v15
                } else {
                    var v16 int16 = 42
                    _ = v16
                }
            default:
                var v17 float64 = 0.0
                _ = v17
            }
            for fi18 := 0; fi18 < 19; fi18++ {
                p0 -= 2
                _ = fi18
            }
        }
    }
    var v19 int32 = 2147483647
    _ = p0
    _ = v19
    return "\n\t"
}

func main() {
    var v1 int16 = 103
    var v2 uint64 = 18446744073709551615
    v3 := []string{"\n\t", ("\n\t" + "foo bar"), "hello", ("a" + "foo bar")}
    var v4 bool = (-58 < -504)
    var v5 int16 = (v1 - v1)
    var v6 bool = v4
    v20 := GenS3{F0: 0, F1: 4294967295, F2: (329 == 181)}
    var v21 uint = 687
    v4 = ((!(!((("\n\t" + "hello") + "hello") == "a"))) && (-191 <= 0))
    var v22 string = (((("a" + "") + ("\n\t" + "hello")) + ("\n\t" + "a")) + ("foo bar" + (("" + "foo bar") + ("\n\t" + (("foo bar" + (((((("\n\t" + (("" + "foo bar") + (("\n\t" + "foo bar") + ("foo bar" + ((("hello" + "") + "hello") + ""))))) + "hello") + "\n\t") + (("hello" + ("foo bar" + "foo bar")) + (("\n\t" + "a") + ("foo bar" + "foo bar")))) + "foo bar") + "\n\t")) + "")))))
    if (3.14159 < 3.14159) {
        if v4 {
            switch 4294967295 {
            case -9223372036854775808:
                v4 = false
            case -2147483648:
                v23 := GenS4{F0: 255}
                _ = v23
            case 32768:
                var v24 string = "foo bar"
                _ = v24
            }
        } else {
            if (false && (!v6)) {
                v25 := GenS2{F0: -1}
                var v26 float64 = 0.0
                var v27 uint16 = 900
                _ = v25
                _ = v26
                _ = v27
            }
            v3 = append(v3, "hello")
            if len(v3) > 0 {
                v3[0] = ""
            }
        }
        for fi28 := 0; fi28 < 5; fi28++ {
            _ = len(v3)
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = fi28
        }
        var v29 int16 = (-(v1))
        _ = v29
    }
    var v30 int32 = 1
    var v31 uint64 = 4294967295
    if true {
        for fi32 := 0; fi32 < 6; fi32++ {
            var v33 int = int(v31)
            v33 = (^(fi32))
            _ = fi32
            _ = v33
        }
    } else {
        var v34 uint = uint(v2)
        fmt.Println(v1, v5)
        _ = v34
    }
    v6 = (!(393 >= 2))
    if (false && (v4 && (2 > 0))) {
        switch v22 {
        case "foo bar":
            v6 = v4
        case "hello":
            for fi35 := 0; fi35 < 6; fi35++ {
                var v36 float64 = ((0.1 * (0.0 / -2.25)) / 3.14159)
                if len(v3) > 0 {
                    v3[0] = (v22 + "hello")
                }
                _ = fi35
                _ = v36
            }
        }
        for fi37 := 0; fi37 < 9; fi37++ {
            if v4 {
                fmt.Println(v30)
            }
            _ = fi37
        }
    } else {
        var v38 uint64 = v31
        if (-39 == 1) {
            if len(v3) > 0 {
                v3[0] = v22
            }
            v4 = v6
        }
        _ = v38
    }
    v6 = (724 != -1)
    v39 := GenS5{F0: 2147483648, F1: (1 >= 100), F2: 69}
    if (v22 == (v22 + v22)) {
        for fi40 := 0; fi40 < 8; fi40++ {
            v41 := GenS1{F0: v31, F1: 128}
            v42 := GenS6{F0: -2147483648}
            var v43 uint32 = 451
            _ = fi40
            _ = v41
            _ = v42
            _ = v43
        }
    } else {
        var v44 int8 = 51
        v30 -= (v30 % ((v30) | 1))
        _ = len(v3)
        _ = v44
    }
    for fi45 := 0; fi45 < 5; fi45++ {
        for fi46 := 0; fi46 < 18; fi46++ {
            v47 := GenS7{F0: -9223372036854775808, F1: 691, F2: 0}
            _ = fi46
            _ = v47
        }
        v48 := GenS2{F0: -1}
        _ = fi45
        _ = v48
    }
    if v6 {
        for fi49 := 0; fi49 < 13; fi49++ {
            switch v22 {
            case "hello":
                var v50 float64 = 3.14159
                _ = v50
            }
            v30 -= v30
            if v6 {
                for fi51 := 0; fi51 < 8; fi51++ {
                    fmt.Println(v6, v4)
                    if false {
                        fmt.Println(v22)
                        if false {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            v5 -= v5
                            fmt.Println(v5)
                        } else {
                            v3 = append(v3, (v22 + (v22 + "foo bar")))
                        }
                    }
                    _ = fi51
                }
                var v52 float64 = 1e10
                v22 = v22
                _ = v52
            }
            _ = fi49
        }
        var v53 uint = v21
        for fi54 := 0; fi54 < 11; fi54++ {
            var v55 int32 = (v30 - 2147483647)
            fmt.Println(v31, v5, fi54)
            var v56 float64 = 0.0
            _ = fi54
            _ = v55
            _ = v56
        }
        _ = v53
    }
    var v57 int32 = v30
    fmt.Println(v1, v5)
    v4 = ((v57 % ((0) | 1)) == v57)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v30
    _ = v31
    _ = v39
    _ = v4
    _ = v5
    _ = v57
    _ = v6
}
