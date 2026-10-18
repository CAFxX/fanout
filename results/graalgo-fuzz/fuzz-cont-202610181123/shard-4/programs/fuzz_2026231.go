package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 int32 }
type GenS2 struct { F0 int; F1 string; F2 int32 }
type GenS3 struct { F0 int32; F1 uint16; F2 int16 }
type GenS4 struct { F0 int64 }
func genFn0(p0 uint16) uint64 {
    for fi4 := 0; fi4 < 6; fi4++ {
        switch (fi4 * (fi4 + 32768)) {
        case 421:
            var v5 int8 = -128
            _ = v5
        case 1:
            var v6 int32 = -1
            _ = v6
        case 32767:
            fmt.Println(fi4, fi4)
        }
        p0 = 142
        if (!(255 >= 2)) {
            var v7 int8 = -4
            _ = v7
        } else {
            p0 -= (^(p0))
            v8 := []int{}
            _ = v8
        }
        _ = fi4
    }
    if false {
        switch "a" {
        case "\n\t":
            fmt.Println(p0)
        }
        p0 -= p0
        if (-707 != 2147483647) {
            p0 = 766
        } else {
            p0 = (p0 + 0)
            switch "" {
            case "a":
                var v9 string = "hello"
                _ = v9
            case "\n\t":
                for fi10 := 0; fi10 < 14; fi10++ {
                    v11 := []int{fi10}
                    for fi12 := 0; fi12 < 5; fi12++ {
                        var v13 uint16 = (^(p0))
                        v11 = append(v11, 32767)
                        v13 += (v13 << ((18446744073709551615) & 63))
                        _ = fi12
                        _ = v13
                    }
                    var v14 string = ""
                    _ = fi10
                    _ = v11
                    _ = v14
                }
            case "":
                for fi15 := 0; fi15 < 7; fi15++ {
                    if false {
                        for fi16 := 0; fi16 < 15; fi16++ {
                            v17 := GenS1{F0: 1, F1: 9223372036854775807, F2: -1}
                            _ = fi16
                            _ = v17
                        }
                    } else {
                        p0 += (p0 << ((127) & 63))
                    }
                    for fi18 := 0; fi18 < 16; fi18++ {
                        v19 := []int{184, 142, fi18}
                        fmt.Println(fi18, p0, fi15)
                        _ = fi18
                        _ = v19
                    }
                    for fi20 := 0; fi20 < 13; fi20++ {
                        var v21 uint64 = 0
                        if (255 > 128) {
                            var v22 int16 = 0
                            _ = v22
                        } else {
                            v23 := []int{(fi15 << ((1) & 63))}
                            var v24 int64 = 42
                            switch 1 {
                            case 4294967295:
                                v24 -= (-(v24))
                            }
                            v25 := GenS1{F0: 2147483647, F1: 0, F2: 2147483647}
                            _ = v23
                            _ = v24
                            _ = v25
                        }
                        _ = fi20
                        _ = v21
                    }
                    _ = fi15
                }
            }
            v26 := GenS2{F0: 358, F1: "hello", F2: 2147483647}
            _ = v26
        }
    }
    switch 9223372036854775807 {
    case -669:
        p0 = p0
    case 4294967295:
        for fi27 := 0; fi27 < 16; fi27++ {
            for fi28 := 0; fi28 < 17; fi28++ {
                var v29 string = "foo bar"
                _ = fi28
                _ = v29
            }
            p0 = p0
            _ = fi27
        }
    default:
        switch 537 {
        case 42:
            v30 := []int{0, 4294967295, 127}
            _ = v30
        case 9223372036854775807:
            var v31 bool = false
            _ = v31
        }
    }
    _ = p0
    return 42
}
func genFn1(p0 int8) int64 {
    v32 := GenS2{F0: -319, F1: "", F2: 42}
    v33 := []int{}
    for fi34 := 0; fi34 < 15; fi34++ {
        p0 = p0
        _ = fi34
    }
    v35 := GenS2{F0: -1, F1: "\n\t", F2: -146}
    _ = p0
    _ = v32
    _ = v33
    _ = v35
    return 686
}

func main() {
    var v1 bool = (!(113 == 1))
    var v2 string = ("foo bar" + "")
    var v3 int32 = -856
    v36 := GenS1{F0: (v3 << ((255) & 63)), F1: 9223372036854775807, F2: 2147483647}
    fmt.Println(v3, v2, v1)
    v37 := GenS2{F0: -2147483648, F1: v2, F2: (v3 ^ v3)}
    v3 = -181
    fmt.Println(v2)
    switch -9223372036854775808 {
    case 9223372036854775807:
        v2 = "\n\t"
    case 456:
        if v1 {
            v38 := GenS1{F0: v3, F1: -1, F2: (v3 | 0)}
            switch "hello" {
            case "\n\t":
                for fi39 := 0; fi39 < 9; fi39++ {
                    for fi40 := 0; fi40 < 3; fi40++ {
                        fmt.Println(v2, v1, v2)
                        _ = fi40
                    }
                    v41 := GenS3{F0: (v3 | (v3 / ((394) | 1))), F1: 584, F2: 1}
                    _ = fi39
                    _ = v41
                }
            default:
                v42 := []int{42, -166, -9223372036854775808}
                _ = v42
            }
            switch 2147483647 {
            case -9223372036854775808:
                if v1 {
                    var v43 int16 = 32767
                    _ = v43
                }
            case 32767:
                v1 = v1
            default:
                fmt.Println(v3)
            }
            _ = v38
        }
    case 937:
        v44 := GenS3{F0: v3, F1: 32767, F2: 0}
        _ = v44
    default:
        v45 := []int{127, -9223372036854775808}
        _ = v45
    }
    if v1 {
        var v46 int8 = -128
        _ = v46
    } else {
        v47 := []int{2147483647}
        switch (v2 + "a") {
        case "":
            if true {
                var v48 int32 = 42
                _ = v48
            }
        case "\n\t":
            v3 -= (v3 - 218)
        case "a":
            v2 = v2
        }
        _ = v47
    }
    switch ((v2 + "foo bar") + v2) {
    case "":
        v49 := GenS3{F0: v3, F1: 0, F2: -548}
        _ = v49
    }
    for fi50 := 0; fi50 < 1; fi50++ {
        var v51 uint32 = 105
        var v52 float32 = (-2.25 * (1e10 / (3.14159 - 0.1)))
        v53 := GenS4{F0: -935}
        _ = fi50
        _ = v51
        _ = v52
        _ = v53
    }
    for fi54 := 0; fi54 < 2; fi54++ {
        fmt.Println(v3, v2)
        _ = fi54
    }
    for fi55 := 0; fi55 < 17; fi55++ {
        var v56 int8 = -81
        _ = fi55
        _ = v56
    }
    fmt.Println(v3, v1)
    v2 = (v2 + "\n\t")
    var v57 uint = 42
    var v58 int16 = 1000
    fmt.Println(v1, v2, v57)
    var v59 float64 = (((-2.25 / -2.25) * (1e10 + 0.1)) * 0.0)
    var v60 string = "\n\t"
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v57
    _ = v58
    _ = v59
    _ = v60
}
