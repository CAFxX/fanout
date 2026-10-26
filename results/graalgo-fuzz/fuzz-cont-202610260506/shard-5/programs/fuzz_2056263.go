package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 string, p1 uint16) int64 {
    v5 := []int{2147483647, 32768}
    if (42 <= -2147483648) {
        if len(v5) > 0 {
            v5[0] = 4294967295
        }
        var v6 uint16 = (p1 - (^(p1)))
        var v7 uint8 = 255
        _ = v6
        _ = v7
    } else {
        var v8 uint64 = 42
        _ = v8
    }
    v5 = append(v5, 128)
    _ = p0
    _ = p1
    _ = v5
    return 455
}
func genFn1() int {
    var v9 uint64 = 773
    v9 -= v9
    v10 := GenS1{F0: 42, F1: -1}
    v11 := GenS1{F0: 0, F1: -1}
    v9 -= v9
    _ = v10
    _ = v11
    _ = v9
    return 9223372036854775807
}

func main() {
    var v1 uint8 = 255
    var v2 uint8 = v1
    v3 := []uint32{42, 2147483648, 2147483647}
    var v4 uint8 = (v2 / ((v2) | 1))
    v2 -= v1
    if true {
        var v12 int8 = -128
        if (false || (false || (512 <= 1))) {
            var v13 uint16 = 830
            _ = v13
        } else {
            var v14 string = "foo bar"
            var v15 bool = false
            _ = v14
            _ = v15
        }
        switch -824 {
        case 128:
            v16 := GenS2{F0: 2}
            _ = v16
        case -25:
            if true {
                v4 = 2
                fmt.Println(v1)
            }
        case 4294967295:
            v3 = append(v3, 4294967295)
        }
        _ = v12
    } else {
        v17 := GenS2{F0: 65535}
        for fi18 := 0; fi18 < 9; fi18++ {
            v3 = append(v3, 2147483647)
            switch (fi18 - (fi18 >> ((2) & 63))) {
            case 725:
                if true {
                    if ((!(0.0 > 3.14159)) || ((v4 * (v4 - 255)) > 0)) {
                        if len(v3) > 0 {
                            v3[0] = 588
                        }
                        switch "\n\t" {
                        case "hello":
                            var v19 uint32 = 2
                            _ = v19
                        default:
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        }
                    } else {
                        v4 = v2
                        v3 = append(v3, 42)
                    }
                    for fi20 := 0; fi20 < 12; fi20++ {
                        var v21 float64 = 0.1
                        v22 := GenS1{F0: -1, F1: 756}
                        fmt.Println(fi18)
                        _ = fi20
                        _ = v21
                        _ = v22
                    }
                }
            case -2147483648:
                v4 = v4
            }
            for fi23 := 0; fi23 < 6; fi23++ {
                var v24 int8 = -128
                _ = fi23
                _ = v24
            }
            _ = fi18
        }
        var v25 uint8 = v2
        _ = v17
        _ = v25
    }
    switch "foo bar" {
    case "foo bar":
        if true {
            fmt.Println(v4, v4, v1)
        } else {
            if ((-1 > 0) || false) {
                fmt.Println(v4, v4, v1)
                v2 -= v4
                var v26 float64 = 0.0
                _ = v26
            }
            for fi27 := 0; fi27 < 5; fi27++ {
                var v28 uint8 = 0
                for fi29 := 0; fi29 < 3; fi29++ {
                    if false {
                        var v30 bool = (1 == 630)
                        v2 = (v28 / ((v28) | 1))
                        _ = v30
                    }
                    _ = fi29
                }
                v31 := GenS1{F0: -495, F1: -2147483648}
                _ = fi27
                _ = v28
                _ = v31
            }
        }
    default:
        fmt.Println(v1)
    }
    if false {
        v1 = 172
        fmt.Println(v4)
    } else {
        for fi32 := 0; fi32 < 19; fi32++ {
            var v33 bool = false
            var v34 uint32 = 826
            fmt.Println(v4, v1, v2)
            _ = fi32
            _ = v33
            _ = v34
        }
    }
    var v35 bool = (!false)
    v1 = 42
    var v36 uint16 = 65535
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v37 := GenS1{F0: 1, F1: 1}
    v2 -= v1
    if len(v3) > 0 {
        v3[0] = 32
    }
    var v38 int16 = 236
    for fi39 := 0; fi39 < 13; fi39++ {
        _ = len(v3)
        _ = fi39
    }
    var v40 uint = 255
    for fi41 := 0; fi41 < 6; fi41++ {
        v3 = append(v3, 4294967295)
        var v42 uint64 = 42
        if false {
            var v43 uint64 = (^(v42))
            var v44 int = fi41
            _ = v43
            _ = v44
        } else {
            var v45 uint8 = v4
            for fi46 := 0; fi46 < 5; fi46++ {
                var v47 int16 = -680
                v35 = (-2.25 != 0.0)
                _ = fi46
                _ = v47
            }
            _ = v45
        }
        _ = fi41
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v40
}
