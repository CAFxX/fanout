package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint; F1 uint }
func genFn0() uint32 {
    if false {
        var v7 int16 = 974
        v7 = (v7 - v7)
        for fi8 := 0; fi8 < 12; fi8++ {
            if (!(42 >= 1)) {
                var v9 bool = true
                _ = v9
            }
            _ = fi8
        }
        _ = v7
    } else {
        if ((170 != 42) || true) {
            v10 := []int{-9223372036854775808}
            if true {
                switch 4294967295 {
                case 32768:
                    var v11 int = -2147483648
                    _ = v11
                case 63:
                    switch -2147483648 {
                    case -893:
                        var v12 int64 = 1
                        _ = v12
                    }
                default:
                    var v13 uint8 = 42
                    _ = v13
                }
                var v14 string = "hello"
                v14 = (((v14 + "foo bar") + "a") + (("a" + ("hello" + "foo bar")) + ""))
                _ = v14
            } else {
                var v15 uint64 = 521
                if false {
                    for fi16 := 0; fi16 < 20; fi16++ {
                        fmt.Println(v15)
                        _ = fi16
                    }
                } else {
                    if false {
                        var v17 int = 54
                        var v18 uint32 = 2147483648
                        if (!false) {
                            if len(v10) > 0 {
                                v10[0] = (v17 + -9223372036854775808)
                            }
                            v15 -= v15
                        }
                        _ = v17
                        _ = v18
                    } else {
                        v15 = v15
                    }
                    var v19 uint32 = 2147483647
                    _ = v19
                }
                fmt.Println(v15)
                _ = v15
            }
            _ = v10
        }
        if true {
            v20 := GenS1{F0: 2, F1: 2147483648}
            for fi21 := 0; fi21 < 16; fi21++ {
                v22 := []int{(fi21 - (-(fi21))), (fi21 % (((fi21 - (fi21 * (fi21 & -612)))) | 1)), fi21, fi21}
                _ = fi21
                _ = v22
            }
            _ = v20
        } else {
            var v23 uint = 314
            v23 = 0
            v23 -= v23
            _ = v23
        }
        var v24 float64 = (0.1 / float64(1.5))
        _ = v24
    }
    v25 := []int{}
    switch (("foo bar" + "\n\t") + "foo bar") {
    case "hello":
        var v26 float32 = ((1.5 / (1e10 / 1.0)) - 1.5)
        _ = v26
    case "":
        var v27 string = (("hello" + "") + (("a" + (("hello" + ("a" + "\n\t")) + (("" + ("a" + "foo bar")) + "hello"))) + ("hello" + "a")))
        _ = v27
    default:
        switch (("" + "a") + "hello") {
        case "":
            var v28 float64 = -2.25
            _ = v28
        case "a":
            v29 := GenS1{F0: 519, F1: 73}
            _ = v29
        case "hello":
            v30 := GenS1{F0: 975, F1: 398}
            _ = v30
        }
    }
    _ = v25
    return 42
}

func main() {
    var v1 int = 127
    var v2 int = v1
    var v3 bool = true
    var v4 uint8 = 127
    v5 := []string{""}
    v6 := make(map[uint32]uint32)
    v31 := GenS1{F0: 65535, F1: genFn0()}
    if v3 {
        v32 := GenS1{F0: 1, F1: genFn0()}
        _ = len(v6)
        _ = v32
    }
    var v33 int8 = 48
    fmt.Println(v33, v3, v2)
    var v34 uint16 = 65535
    for fi35 := 0; fi35 < 1; fi35++ {
        for fi36 := 0; fi36 < 6; fi36++ {
            if v3 {
                _ = len(v5)
            }
            v2 = 32767
            _ = len(v6)
            _ = fi36
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        for fi37 := 0; fi37 < 20; fi37++ {
            var v38 string = ("\n\t" + ("" + ""))
            switch v38 {
            case "a":
                v39 := GenS2{F0: "\n\t"}
                _ = v39
            case "foo bar":
                var v40 uint64 = 534
                _ = v40
            case "\n\t":
                var v41 float32 = 1e10
                _ = v41
            }
            _ = fi37
            _ = v38
        }
        _ = fi35
    }
    v42 := v6[439]
    if v3 {
        v34 -= v34
        var v43 uint = uint(v42)
        v6[genFn0()] = genFn0()
        _ = v43
    }
    if ((v33 << ((439) & 63)) == v33) {
        fmt.Println(v3)
        v44 := GenS3{F0: 42, F1: 255}
        _ = v44
    }
    var v45 uint32 = (v42 | v42)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v33
    _ = v34
    _ = v4
    _ = v42
    _ = v45
    _ = v5
    _ = v6
}
