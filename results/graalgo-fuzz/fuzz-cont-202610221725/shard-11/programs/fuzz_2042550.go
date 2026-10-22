package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8; F2 uint }
type GenS2 struct { F0 int16; F1 uint8; F2 uint16 }
type GenS3 struct { F0 int; F1 int64; F2 bool }
type GenS4 struct { F0 uint; F1 int8 }
type GenS5 struct { F0 uint64; F1 int }
type GenS6 struct { F0 uint32; F1 uint16; F2 int }
type GenS7 struct { F0 int16 }
func genFn0(p0 string, p1 string, p2 uint8) int64 {
    for fi4 := 0; fi4 < 6; fi4++ {
        if true {
            var v5 int = 127
            _ = v5
        } else {
            p2 = p2
        }
        _ = fi4
    }
    for fi6 := 0; fi6 < 8; fi6++ {
        if false {
            if false {
                var v7 float64 = 1e10
                v8 := []int{fi6, fi6, (fi6 >> ((192) & 63)), fi6}
                if len(v8) > 0 {
                    v8[0] = (-(fi6))
                }
                _ = v7
                _ = v8
            } else {
                if false {
                    p1 = (((p0 + (p1 + p0)) + p1) + (("" + ("foo bar" + "\n\t")) + ("\n\t" + ((p0 + p0) + p1))))
                } else {
                    v9 := []int{(fi6 - (-(fi6)))}
                    fmt.Println(fi6, p0)
                    _ = v9
                }
                v10 := GenS1{F0: -128, F1: p2, F2: 17}
                v11 := GenS2{F0: 368, F1: 248, F2: 1}
                _ = v10
                _ = v11
            }
            p2 += (p2 | p2)
        }
        _ = fi6
    }
    for fi12 := 0; fi12 < 9; fi12++ {
        switch (p1 + p1) {
        case "foo bar":
            p0 = p1
        }
        _ = fi12
    }
    _ = p0
    _ = p1
    _ = p2
    return 0
}
func genFn1() int {
    v13 := GenS1{F0: 43, F1: 128, F2: 0}
    if true {
        if (false && (114 <= 131)) {
            for fi14 := 0; fi14 < 5; fi14++ {
                var v15 uint16 = 392
                _ = fi14
                _ = v15
            }
            for fi16 := 0; fi16 < 1; fi16++ {
                fmt.Println(fi16)
                var v17 uint8 = 2
                var v18 uint8 = v17
                _ = fi16
                _ = v17
                _ = v18
            }
        }
    }
    var v19 int16 = -630
    _ = v13
    _ = v19
    return 42
}

func main() {
    var v1 int16 = -1
    var v2 float64 = 1e10
    var v3 uint = 65535
    v3 = 712
    if true {
        v20 := GenS3{F0: -648, F1: 9223372036854775807, F2: false}
        for fi21 := 0; fi21 < 13; fi21++ {
            v22 := []int{902, (fi21 * fi21)}
            _ = fi21
            _ = v22
        }
        _ = v20
    } else {
        v23 := []int{368, genFn1(), 128}
        var v24 int64 = 42
        _ = v23
        _ = v24
    }
    v1 += 0
    if (0 >= -1) {
        v25 := []int{}
        _ = v25
    } else {
        fmt.Println(v2, v3, v3)
        if (!true) {
            if (18446744073709551615 >= (v3 + 0)) {
                var v26 int = 4294967295
                switch ("a" + "\n\t") {
                case "a":
                    v27 := []int{v26, genFn1(), 32768}
                    _ = v27
                case "\n\t":
                    var v28 string = "a"
                    _ = v28
                }
                _ = v26
            } else {
                v1 += 1000
            }
        }
        v29 := []int{2147483647, 0, 32768}
        _ = v29
    }
    var v30 float32 = -2.25
    v1 = v1
    fmt.Println(v2)
    fmt.Println(v3, v3)
    v31 := GenS4{F0: (v3 ^ 0), F1: -1}
    for fi32 := 0; fi32 < 11; fi32++ {
        if (false && (42 == 9223372036854775807)) {
            switch (("" + "foo bar") + ("a" + "\n\t")) {
            case "\n\t":
                fmt.Println(v30, v30, v1)
            }
        } else {
            if true {
                if (false || (true || (0 > 1))) {
                    switch (fi32 << ((v3) & 63)) {
                    case 42:
                        var v33 uint64 = 2
                        _ = v33
                    case 611:
                        v30 += v30
                    }
                    if (false || (1 != 42)) {
                        v34 := []int{fi32}
                        switch "" {
                        case "hello":
                            var v35 int = (fi32 * genFn1())
                            _ = v35
                        case "\n\t":
                            fmt.Println(v30, v1)
                        default:
                            switch "\n\t" {
                            case "\n\t":
                                _ = len(v34)
                            case "":
                                v3 = 65535
                            }
                        }
                        for fi36 := 0; fi36 < 16; fi36++ {
                            if len(v34) > 0 {
                                v34[0] = -740
                            }
                            var v37 bool = true
                            fmt.Println(v37, v3, v37)
                            _ = fi36
                            _ = v37
                        }
                        _ = v34
                    } else {
                        v38 := []int{fi32}
                        v39 := GenS4{F0: v3, F1: -128}
                        v40 := GenS5{F0: 4294967295, F1: fi32}
                        _ = v38
                        _ = v39
                        _ = v40
                    }
                    var v41 int8 = 42
                    _ = v41
                } else {
                    var v42 uint32 = 0
                    _ = v42
                }
            } else {
                var v43 uint32 = 2147483648
                v44 := GenS6{F0: v43, F1: 420, F2: genFn1()}
                v45 := GenS1{F0: -128, F1: 42, F2: v3}
                _ = v43
                _ = v44
                _ = v45
            }
            v46 := GenS7{F0: (v1 / (((v1 + (v1 - (v1 | (v1 ^ (^(v1))))))) | 1))}
            v47 := []int{(fi32 << ((v3) & 63))}
            _ = v46
            _ = v47
        }
        _ = fi32
    }
    var v48 uint8 = 159
    var v49 uint64 = 0
    v2 -= 0.1
    var v50 int64 = -5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v48
    _ = v49
    _ = v50
}
