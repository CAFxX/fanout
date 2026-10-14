package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 int32; F1 uint64; F2 uint16 }
type GenS4 struct { F0 int16; F1 string }
type GenS5 struct { F0 int; F1 int64 }
type GenS6 struct { F0 bool; F1 bool; F2 uint }
func genFn0(p0 bool) bool {
    var v5 int64 = 1
    v6 := []int{32767, 1}
    _ = p0
    _ = v5
    _ = v6
    return (true && (1.5 <= float64(0.0)))
}
func genFn1(p0 bool) int {
    switch 1 {
    case 42:
        for fi7 := 0; fi7 < 20; fi7++ {
            if (((("hello" + (("a" + "foo bar") + "foo bar")) + ("\n\t" + "a")) + "foo bar") == "a") {
                var v8 uint = 639
                _ = v8
            } else {
                var v9 uint8 = 0
                _ = v9
            }
            if ((-128 >= 1) || (true || (33 < 2147483648))) {
                v10 := GenS1{F0: -9223372036854775808}
                if p0 {
                    for fi11 := 0; fi11 < 15; fi11++ {
                        var v12 int32 = 1
                        v13 := []int{(fi7 - (fi7 * (fi7 * (fi11 >> ((745) & 63))))), (fi7 * (fi7 * fi7)), fi11, (fi11 >> ((102) & 63))}
                        fmt.Println(fi11, p0)
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                } else {
                    var v14 int = (fi7 - (fi7 % (((-(fi7))) | 1)))
                    v14 -= fi7
                    fmt.Println(fi7)
                    _ = v14
                }
                if false {
                    fmt.Println(fi7, p0)
                    var v15 int8 = 1
                    p0 = (629 != 42)
                    _ = v15
                } else {
                    var v16 int64 = 0
                    v17 := GenS1{F0: 1}
                    _ = v16
                    _ = v17
                }
                _ = v10
            } else {
                p0 = (255 == 255)
                if (255 < 2) {
                    v18 := []int{(-(fi7)), fi7}
                    for fi19 := 0; fi19 < 15; fi19++ {
                        p0 = true
                        var v20 float64 = -2.25
                        var v21 int = fi19
                        _ = fi19
                        _ = v20
                        _ = v21
                    }
                    _ = v18
                } else {
                    p0 = (0 == 0)
                    p0 = (1 > 190)
                    v22 := []int{fi7}
                    _ = v22
                }
            }
            _ = fi7
        }
    default:
        p0 = (!p0)
    }
    switch (("foo bar" + "a") + "foo bar") {
    case "hello":
        fmt.Println(p0)
    default:
        if false {
            v23 := GenS2{F0: "foo bar", F1: -32768}
            p0 = p0
            var v24 int16 = 32767
            _ = v23
            _ = v24
        } else {
            v25 := []int{42, -2147483648}
            for fi26 := 0; fi26 < 20; fi26++ {
                var v27 bool = (1000 < 1)
                _ = len(v25)
                _ = fi26
                _ = v27
            }
            p0 = p0
            _ = v25
        }
    }
    _ = p0
    return 1
}

func main() {
    var v1 int = 4294967295
    var v2 uint = 65535
    var v3 uint = v2
    var v4 uint = v2
    for fi28 := 0; fi28 < 14; fi28++ {
        fmt.Println(v3)
        v29 := GenS3{F0: 2147483647, F1: 4294967295, F2: 32767}
        var v30 uint = v4
        _ = fi28
        _ = v29
        _ = v30
    }
    v3 = (v2 % (((v4 * (v3 & v4))) | 1))
    v1 -= -85
    v3 = v2
    if true {
        v31 := GenS3{F0: -2147483648, F1: 4294967295, F2: uint16(v4)}
        fmt.Println(v3)
        switch v1 {
        case 127:
            fmt.Println(v4)
        case 0:
            for fi32 := 0; fi32 < 10; fi32++ {
                if false {
                    fmt.Println(v4)
                    v33 := []int{v1}
                    v34 := GenS4{F0: 42, F1: ""}
                    _ = v33
                    _ = v34
                } else {
                    var v35 float64 = 1e10
                    if true {
                        var v36 int64 = -9223372036854775808
                        for fi37 := 0; fi37 < 6; fi37++ {
                            fmt.Println(v2, v1)
                            switch "" {
                            case "":
                                v38 := []int{fi32, (fi32 << ((255) & 63)), v1, 1}
                                _ = v38
                            case "\n\t":
                                v3 = v3
                            }
                            _ = fi37
                        }
                        for fi39 := 0; fi39 < 4; fi39++ {
                            fmt.Println(v1, v3, v36)
                            var v40 uint16 = 0
                            v41 := []int{}
                            fmt.Println(v3, v40, v36)
                            _ = fi39
                            _ = v40
                            _ = v41
                        }
                        _ = v36
                    }
                    v42 := GenS5{F0: v1, F1: 9223372036854775807}
                    _ = v35
                    _ = v42
                }
                v43 := []int{fi32, genFn1(true), (fi32 >> (((v2 % (((^(v4))) | 1))) & 63)), v1}
                var v44 uint8 = 29
                _ = fi32
                _ = v43
                _ = v44
            }
        }
        _ = v31
    }
    v2 = (v2 * (v4 + (v2 & v4)))
    var v45 float64 = 3.14159
    v46 := GenS5{F0: v1, F1: -9223372036854775808}
    switch ("" + ("" + ("a" + ("a" + "hello")))) {
    case "foo bar":
        fmt.Println(v3, v4)
    case "\n\t":
        v47 := GenS6{F0: false, F1: (false && (996 < 65535)), F2: v3}
        _ = v47
    default:
        switch "foo bar" {
        case "\n\t":
            v4 -= v2
        default:
            if false {
                var v48 float32 = 1.5
                _ = v48
            } else {
                v1 -= v1
                for fi49 := 0; fi49 < 8; fi49++ {
                    v1 = (fi49 + fi49)
                    _ = fi49
                }
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
}
