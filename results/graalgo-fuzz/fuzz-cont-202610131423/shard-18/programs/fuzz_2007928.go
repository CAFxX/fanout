package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 uint64; F1 int32 }
type GenS4 struct { F0 uint16; F1 int8; F2 bool }
type GenS5 struct { F0 uint32; F1 uint32; F2 uint16 }
func genFn0(p0 bool, p1 string, p2 uint8) uint {
    for fi4 := 0; fi4 < 16; fi4++ {
        p2 = p2
        for fi5 := 0; fi5 < 13; fi5++ {
            var v6 uint = 0
            switch (((p1 + p1) + "") + p1) {
            case "foo bar":
                fmt.Println(fi4, p2, p1)
            case "a":
                if false {
                    v7 := []int{(fi4 & 2147483647), 127, fi4, -9223372036854775808}
                    var v8 bool = p0
                    _ = v7
                    _ = v8
                }
            default:
                p1 = (p1 + p1)
            }
            _ = fi5
            _ = v6
        }
        _ = fi4
    }
    for fi9 := 0; fi9 < 6; fi9++ {
        var v10 uint32 = 352
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 20; fi11++ {
        if ((fi11 >> ((2) & 63)) <= (fi11 + (fi11 / ((fi11) | 1)))) {
            fmt.Println(p1, fi11, p1)
        }
        p1 = "foo bar"
        _ = fi11
    }
    if (true || (-1 >= -60)) {
        var v12 float32 = 0.0
        _ = v12
    }
    v13 := GenS1{F0: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 2
}
func genFn1() bool {
    if true {
        var v14 float64 = (((0.1 - (1.5 / 1.0)) * 0.0) / 0.1)
        v15 := GenS1{F0: 42}
        v16 := GenS1{F0: 1}
        _ = v14
        _ = v15
        _ = v16
    }
    if (true && (false || (-1 == 42))) {
        if false {
            var v17 int8 = 1
            for fi18 := 0; fi18 < 9; fi18++ {
                v19 := []int{-2147483648, (fi18 - (fi18 << ((351) & 63)))}
                var v20 int16 = 623
                _ = fi18
                _ = v19
                _ = v20
            }
            v17 = -98
            _ = v17
        }
    } else {
        v21 := []int{42, 4294967295, 709}
        _ = v21
    }
    return (-128 != 127)
}

func main() {
    v1 := []uint32{0, 65535, 2147483647, 2}
    v2 := []string{"a", ("foo bar" + "foo bar"), (("hello" + "\n\t") + "hello")}
    v3 := []uint32{4294967295, 6}
    var v22 float64 = ((0.0 / 0.1) * 1.5)
    _ = len(v1)
    var v23 int16 = 0
    for fi24 := 0; fi24 < 5; fi24++ {
        v25 := GenS1{F0: 18446744073709551615}
        v26 := GenS2{F0: -70, F1: (v23 * 1000)}
        var v27 uint16 = 2
        _ = fi24
        _ = v25
        _ = v26
        _ = v27
    }
    if len(v3) > 0 {
        v3[0] = 858
    }
    for fi28 := 0; fi28 < 4; fi28++ {
        for fi29 := 0; fi29 < 4; fi29++ {
            v3 = append(v3, 579)
            v23 = (v23 / ((32767) | 1))
            _ = fi29
        }
        _ = fi28
    }
    fmt.Println(v22, v23)
    var v30 int64 = 9223372036854775807
    var v31 uint = 2
    var v32 bool = false
    for fi33 := 0; fi33 < 6; fi33++ {
        switch "a" {
        case "":
            v23 -= v23
        case "a":
            v32 = (2 > 65535)
        case "foo bar":
            var v34 int = fi33
            _ = v34
        default:
            switch fi33 {
            case 1:
                var v35 string = "\n\t"
                _ = v35
            case 2147483647:
                v36 := GenS3{F0: 2, F1: -586}
                _ = v36
            }
        }
        _ = fi33
    }
    if v32 {
        v32 = (0 != -745)
    } else {
        for fi37 := 0; fi37 < 14; fi37++ {
            if (false && v32) {
                fmt.Println(v31)
            }
            _ = fi37
        }
        v38 := GenS4{F0: 919, F1: 1, F2: (v32 && true)}
        v39 := GenS5{F0: 952, F1: 195, F2: 32767}
        _ = v38
        _ = v39
    }
    var v40 int = 9223372036854775807
    v41 := GenS1{F0: 18446744073709551615}
    switch v40 {
    case -2147483648:
        var v42 string = "\n\t"
        _ = v42
    case 1:
        v32 = (65535 != 32767)
    default:
        v3 = append(v3, 2)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v40
    _ = v41
}
