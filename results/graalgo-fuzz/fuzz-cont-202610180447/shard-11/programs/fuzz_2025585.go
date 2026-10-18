package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 int }
type GenS2 struct { F0 uint32; F1 uint16 }
type GenS3 struct { F0 bool }
func genFn0(p0 int8, p1 uint16) bool {
    if true {
        var v7 uint16 = p1
        _ = v7
    } else {
        var v8 uint16 = (^(p1))
        _ = v8
    }
    fmt.Println(p0)
    fmt.Println(p1, p1)
    fmt.Println(p0)
    if true {
        v9 := []int{183, 42, 127, 370}
        _ = v9
    } else {
        p0 -= p0
    }
    _ = p0
    _ = p1
    return true
}
func genFn1(p0 int16) string {
    v10 := GenS1{F0: 842, F1: 0, F2: 32768}
    if false {
        switch ("foo bar" + ("foo bar" + "")) {
        case "foo bar":
            fmt.Println(p0)
        default:
            for fi11 := 0; fi11 < 20; fi11++ {
                fmt.Println(p0, p0)
                v12 := GenS1{F0: 18446744073709551615, F1: 42, F2: (fi11 + fi11)}
                _ = fi11
                _ = v12
            }
        }
        switch "\n\t" {
        case "":
            if false {
                switch -9223372036854775808 {
                case 4294967295:
                    var v13 uint = 2
                    _ = v13
                }
                for fi14 := 0; fi14 < 9; fi14++ {
                    var v15 bool = true
                    switch fi14 {
                    case 2147483647:
                        p0 = p0
                    case 938:
                        v16 := []int{(fi14 | (fi14 * fi14))}
                        _ = v16
                    case 127:
                        v17 := GenS1{F0: 1, F1: 73, F2: fi14}
                        _ = v17
                    default:
                        for fi18 := 0; fi18 < 16; fi18++ {
                            v15 = v15
                            fmt.Println(fi14)
                            _ = fi18
                        }
                    }
                    fmt.Println(fi14, fi14)
                    _ = fi14
                    _ = v15
                }
            }
        }
        for fi19 := 0; fi19 < 6; fi19++ {
            v20 := GenS1{F0: 897, F1: 127, F2: (fi19 / ((fi19) | 1))}
            var v21 uint = 860
            for fi22 := 0; fi22 < 15; fi22++ {
                fmt.Println(fi19)
                v21 = (v21 - 422)
                var v23 int8 = 127
                _ = fi22
                _ = v23
            }
            _ = fi19
            _ = v20
            _ = v21
        }
    }
    p0 -= p0
    if ("hello" != "hello") {
        var v24 float32 = (float32(((0.0 / 0.1) - 1e10)) - 0.0)
        p0 = (-(p0))
        _ = v24
    } else {
        var v25 string = ("\n\t" + "a")
        _ = v25
    }
    _ = p0
    _ = v10
    return (((("" + ("a" + "")) + "a") + ("hello" + "hello")) + "")
}

func main() {
    var v1 uint = 581
    v2 := []int{}
    var v3 bool = false
    var v4 uint8 = 0
    var v5 float32 = (1.5 + 1e10)
    var v6 int32 = -720
    fmt.Println(v5, v3)
    v3 = (v3 || (!(!v3)))
    v1 -= (^(v1))
    v2 = append(v2, 0)
    switch 42 {
    case -157:
        var v26 int8 = -128
        _ = v26
    case 0:
        v4 = v4
    case 204:
        for fi27 := 0; fi27 < 12; fi27++ {
            var v28 int16 = 72
            v4 = (v4 << ((v1) & 63))
            switch 2147483647 {
            case 1:
                v29 := GenS1{F0: 2, F1: (v4 % (((v4 & (v4 << ((v1) & 63)))) | 1)), F2: -329}
                _ = v29
            }
            _ = fi27
            _ = v28
        }
    }
    switch "hello" {
    case "\n\t":
        v1 -= (v1 / ((255) | 1))
    case "a":
        fmt.Println(v4, v5)
    case "foo bar":
        v2 = append(v2, 535)
    }
    switch ("hello" + ("hello" + "")) {
    case "foo bar":
        v30 := GenS2{F0: 4294967295, F1: 65535}
        _ = v30
    case "hello":
        fmt.Println(v1)
    case "\n\t":
        for fi31 := 0; fi31 < 9; fi31++ {
            fmt.Println(v4, v6)
            fmt.Println(fi31, v1)
            _ = fi31
        }
    }
    var v32 bool = v3
    if v3 {
        v4 = v4
        for fi33 := 0; fi33 < 9; fi33++ {
            if v3 {
                var v34 uint = v1
                _ = v34
            }
            switch "foo bar" {
            case "a":
                var v35 uint = 255
                _ = v35
            default:
                if (0 < -1) {
                    var v36 uint8 = v4
                    var v37 float64 = 1e10
                    v38 := GenS3{F0: true}
                    _ = v36
                    _ = v37
                    _ = v38
                }
            }
            _ = fi33
        }
        if false {
            var v39 string = "foo bar"
            _ = v39
        }
    }
    v2 = append(v2, -1)
    v40 := GenS2{F0: 2147483648, F1: 0}
    var v41 int8 = 100
    v5 = v5
    v42 := GenS1{F0: 1, F1: v4, F2: -887}
    var v43 int = -2147483648
    if len(v2) > 0 {
        v2[0] = (v43 | (v43 >> (((v1 % (((v1 % (((^(v1))) | 1))) | 1))) & 63)))
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v5
    _ = v6
}
