package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8 }
type GenS2 struct { F0 int32; F1 uint8 }
type GenS3 struct { F0 bool; F1 int16; F2 uint16 }
type GenS4 struct { F0 int16; F1 int64; F2 uint }
func genFn0() int32 {
    if (false || (390 == -189)) {
        if ((42 < 983) || true) {
            if false {
                var v5 int32 = 167
                _ = v5
            } else {
                for fi6 := 0; fi6 < 13; fi6++ {
                    var v7 uint = 614
                    var v8 int16 = 42
                    _ = fi6
                    _ = v7
                    _ = v8
                }
                if true {
                    switch "foo bar" {
                    case "":
                        switch ("hello" + "hello") {
                        case "a":
                            v9 := []int{32768}
                            _ = v9
                        default:
                            v10 := GenS1{F0: -693, F1: 2}
                            _ = v10
                        }
                    case "\n\t":
                        var v11 uint64 = 18446744073709551615
                        _ = v11
                    }
                    v12 := GenS1{F0: 9223372036854775807, F1: 128}
                    v13 := []int{}
                    _ = v12
                    _ = v13
                }
                switch 128 {
                case -83:
                    for fi14 := 0; fi14 < 1; fi14++ {
                        var v15 uint8 = 42
                        _ = fi14
                        _ = v15
                    }
                default:
                    var v16 string = ((("a" + "a") + "hello") + "a")
                    _ = v16
                }
            }
            var v17 uint16 = 104
            var v18 int32 = 0
            _ = v17
            _ = v18
        } else {
            var v19 string = ((("hello" + ("\n\t" + "a")) + "foo bar") + "foo bar")
            if true {
                var v20 int8 = 127
                for fi21 := 0; fi21 < 14; fi21++ {
                    v22 := []int{}
                    var v23 int16 = 1
                    var v24 float64 = 0.1
                    _ = fi21
                    _ = v22
                    _ = v23
                    _ = v24
                }
                v19 = ("a" + v19)
                _ = v20
            }
            _ = v19
        }
    }
    v25 := []int{}
    _ = v25
    return 0
}
func genFn1(p0 uint16) uint16 {
    for fi26 := 0; fi26 < 1; fi26++ {
        v27 := GenS1{F0: (fi26 * (fi26 % (((^(fi26))) | 1))), F1: 128}
        _ = fi26
        _ = v27
    }
    switch ("" + "") {
    case "hello":
        fmt.Println(p0)
    case "foo bar":
        fmt.Println(p0)
    case "":
        if true {
            for fi28 := 0; fi28 < 17; fi28++ {
                v29 := GenS1{F0: fi28, F1: 1}
                _ = fi28
                _ = v29
            }
            fmt.Println(p0)
        } else {
            switch 2147483647 {
            case 127:
                v30 := []int{}
                _ = v30
            case 32768:
                var v31 int16 = -1
                _ = v31
            case 4294967295:
                v32 := []int{-682, 690, 602, 1}
                _ = v32
            default:
                for fi33 := 0; fi33 < 18; fi33++ {
                    var v34 uint = 0
                    _ = fi33
                    _ = v34
                }
            }
            var v35 float64 = 0.1
            _ = v35
        }
    default:
        var v36 uint16 = 32767
        _ = v36
    }
    if (((-180 >= -32768) && true) && (1 != 1)) {
        v37 := GenS2{F0: -2147483648, F1: 127}
        v38 := GenS1{F0: 643, F1: 2}
        _ = v37
        _ = v38
    } else {
        for fi39 := 0; fi39 < 20; fi39++ {
            v40 := []int{fi39, fi39, (fi39 * fi39), (fi39 ^ fi39)}
            fmt.Println(fi39)
            v41 := GenS3{F0: false, F1: 1000, F2: 526}
            _ = fi39
            _ = v40
            _ = v41
        }
    }
    var v42 uint = 299
    v43 := []int{2147483647, -9223372036854775808, -9223372036854775808, -9223372036854775808}
    _ = p0
    _ = v42
    _ = v43
    return (p0 | (p0 / ((p0) | 1)))
}

func main() {
    var v1 uint16 = 525
    var v2 string = (((("a" + ("hello" + "hello")) + "foo bar") + "hello") + ("foo bar" + "hello"))
    var v3 int8 = 69
    var v4 bool = false
    v44 := []int{4294967295, 4294967295, 9223372036854775807}
    var v45 string = v2
    v46 := GenS3{F0: true, F1: -742, F2: v1}
    for fi47 := 0; fi47 < 9; fi47++ {
        switch "hello" {
        case "":
            var v48 bool = (!(2147483647 <= 2147483647))
            _ = v48
        }
        var v49 int16 = 42
        v3 = (-(v3))
        _ = fi47
        _ = v49
    }
    for fi50 := 0; fi50 < 10; fi50++ {
        for fi51 := 0; fi51 < 6; fi51++ {
            var v52 int64 = -9223372036854775808
            v3 = 1
            var v53 int64 = v52
            _ = fi51
            _ = v52
            _ = v53
        }
        if v4 {
            v3 -= v3
            fmt.Println(v45)
        }
        _ = fi50
    }
    var v54 uint32 = 0
    for fi55 := 0; fi55 < 15; fi55++ {
        var v56 uint16 = (v1 & (v1 / ((v1) | 1)))
        switch v45 {
        case "":
            v1 = (v56 / (((v1 >> ((65535) & 63))) | 1))
        case "\n\t":
            v56 -= genFn1((^(v56)))
        case "hello":
            if ((18446744073709551615 == 0) && true) {
                v57 := GenS4{F0: 69, F1: 1, F2: 0}
                _ = v57
            }
        }
        var v58 uint16 = (v56 & (^(v56)))
        _ = fi55
        _ = v56
        _ = v58
    }
    var v59 uint16 = (v1 / ((0) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v44:", len(v44), v44)
    fmt.Println("v45:", v45)
    fmt.Println("v54:", v54)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
    _ = v54
    _ = v59
}
