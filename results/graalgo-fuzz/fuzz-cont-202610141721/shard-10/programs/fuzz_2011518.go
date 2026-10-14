package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int }
type GenS2 struct { F0 uint; F1 uint8; F2 uint64 }
type GenS3 struct { F0 int64; F1 uint }
type GenS4 struct { F0 uint16; F1 int8; F2 int8 }

func main() {
    var v1 uint16 = 42
    var v2 int32 = 2147483647
    v3 := make(map[uint32]uint32)
    var v4 float64 = (1e10 * (1e10 / -2.25))
    if (v1 >= v1) {
        for fi5 := 0; fi5 < 11; fi5++ {
            v6 := []int{9223372036854775807, fi5}
            _ = fi5
            _ = v6
        }
        fmt.Println(v2)
        v7 := GenS1{F0: 56, F1: -9223372036854775808}
        _ = v7
    } else {
        switch -897 {
        case 32768:
            var v8 int8 = 100
            _ = v8
        case 2147483647:
            v2 = v2
        default:
            switch "" {
            case "foo bar":
                v3[581] = 65535
            case "a":
                for fi9 := 0; fi9 < 7; fi9++ {
                    v10 := GenS1{F0: 93, F1: 0}
                    for fi11 := 0; fi11 < 9; fi11++ {
                        var v12 int64 = 1
                        var v13 int = fi11
                        fmt.Println(fi11, fi9, v13)
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                    _ = fi9
                    _ = v10
                }
            default:
                for fi14 := 0; fi14 < 15; fi14++ {
                    var v15 uint = 789
                    v16 := GenS2{F0: 255, F1: 2, F2: 1}
                    var v17 uint16 = v1
                    _ = fi14
                    _ = v15
                    _ = v16
                    _ = v17
                }
            }
        }
    }
    switch "foo bar" {
    case "a":
        v18 := v3[42]
        _ = v18
    }
    var v19 int32 = (v2 & (v2 << ((550) & 63)))
    for fi20 := 0; fi20 < 6; fi20++ {
        v21 := GenS1{F0: 128, F1: fi20}
        switch (("hello" + "foo bar") + "a") {
        case "\n\t":
            if false {
                fmt.Println(v1, v2)
                if true {
                    v4 += v4
                    for fi22 := 0; fi22 < 10; fi22++ {
                        v23 := []int{(fi22 << ((1) & 63))}
                        _ = fi22
                        _ = v23
                    }
                    v24 := GenS3{F0: -1, F1: 255}
                    _ = v24
                } else {
                    v2 = (v19 - v2)
                }
            } else {
                for fi25 := 0; fi25 < 3; fi25++ {
                    v26 := []int{}
                    var v27 string = "foo bar"
                    _ = fi25
                    _ = v26
                    _ = v27
                }
            }
        case "":
            if false {
                v2 = (^(v19))
            }
        default:
            v4 -= v4
        }
        v28 := GenS3{F0: -1, F1: 1}
        _ = fi20
        _ = v21
        _ = v28
    }
    v29 := GenS2{F0: 180, F1: 2, F2: 42}
    if true {
        v30 := []int{644}
        v31 := GenS3{F0: 378, F1: 18446744073709551615}
        if (-9223372036854775808 != -1) {
            if (!false) {
                v19 = v2
                v32 := GenS4{F0: (v1 >> ((18446744073709551615) & 63)), F1: 100, F2: -71}
                _ = v32
            }
            fmt.Println(v4)
            var v33 int64 = -1
            _ = v33
        } else {
            v4 = ((((1e10 - 1.5) + v4) / 0.1) - v4)
        }
        _ = v30
        _ = v31
    }
    for fi34 := 0; fi34 < 20; fi34++ {
        v35 := []int{127, -1, 32767, 128}
        for fi36 := 0; fi36 < 2; fi36++ {
            if (618 < -9223372036854775808) {
                delete(v3, 676)
                var v37 bool = false
                switch (^(fi36)) {
                case -503:
                    if v37 {
                        if len(v35) > 0 {
                            v35[0] = (fi34 << ((255) & 63))
                        }
                        v19 = v19
                        v19 = v19
                    }
                }
                _ = v37
            }
            _ = fi36
        }
        _ = fi34
        _ = v35
    }
    var v38 string = "\n\t"
    v39 := GenS2{F0: 1, F1: 128, F2: 0}
    v1 += (v1 - 32767)
    switch v38 {
    case "":
        if true {
            if (!false) {
                v40 := []int{}
                var v41 uint32 = 0
                _ = v40
                _ = v41
            }
            fmt.Println(v4, v1, v38)
        }
    case "\n\t":
        var v42 int64 = 0
        _ = v42
    default:
        fmt.Println(v19)
    }
    v43 := GenS3{F0: 42, F1: 105}
    v44 := []int{2147483647, -2147483648}
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v44:", len(v44), v44)
    _ = v1
    _ = v19
    _ = v2
    _ = v29
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v43
    _ = v44
}
