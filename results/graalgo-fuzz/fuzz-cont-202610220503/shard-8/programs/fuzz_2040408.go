package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string }
type GenS2 struct { F0 int64; F1 uint16; F2 bool }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint16; F1 int; F2 uint32 }
type GenS5 struct { F0 int16; F1 uint32 }
type GenS6 struct { F0 int16; F1 string }

func main() {
    var v1 int64 = 0
    var v2 int64 = v1
    var v3 uint = 42
    var v4 int = 32767
    v5 := make(map[int]int)
    v4 = (v4 | (v4 + v4))
    var v6 float64 = 1e10
    if (true || (65535 >= 0)) {
        for fi7 := 0; fi7 < 7; fi7++ {
            v8 := []int{(v4 & (fi7 >> ((v3) & 63))), fi7, fi7, fi7}
            v9 := GenS1{F0: 2, F1: ("hello" + ("hello" + "a"))}
            switch 555 {
            case 1:
                v10 := GenS2{F0: (^(v1)), F1: 65535, F2: true}
                _ = v10
            case 9223372036854775807:
                for fi11 := 0; fi11 < 7; fi11++ {
                    delete(v5, int(v1))
                    var v12 float64 = v6
                    var v13 string = "foo bar"
                    _ = fi11
                    _ = v12
                    _ = v13
                }
            case 873:
                for fi14 := 0; fi14 < 10; fi14++ {
                    v15 := GenS2{F0: v2, F1: 0, F2: (!false)}
                    v16 := GenS2{F0: 714, F1: 42, F2: (1000 < -410)}
                    switch v4 {
                    case 32767:
                        v1 += (v1 << ((18446744073709551615) & 63))
                    case 898:
                        var v17 float32 = 1.5
                        _ = v17
                    default:
                        fmt.Println(v6, fi14)
                    }
                    _ = fi14
                    _ = v15
                    _ = v16
                }
            }
            _ = fi7
            _ = v8
            _ = v9
        }
        var v18 uint16 = 832
        v4 += (^(v4))
        _ = v18
    } else {
        v2 += 42
        switch "\n\t" {
        case "foo bar":
            v19 := GenS1{F0: 4294967295, F1: "hello"}
            _ = v19
        case "a":
            v20 := GenS1{F0: 65535, F1: "\n\t"}
            _ = v20
        case "\n\t":
            switch "\n\t" {
            case "foo bar":
                v21 := v5[v4]
                _ = v21
            case "":
                v2 = -517
            default:
                switch "hello" {
                case "\n\t":
                    v22 := []int{128, v4}
                    _ = v22
                }
            }
        }
        var v23 uint16 = 42
        _ = v23
    }
    if false {
        if (230 == 42) {
            if (2147483647 >= int32(v2)) {
                var v24 int64 = v1
                v25 := GenS3{F0: 18446744073709551615}
                _ = v24
                _ = v25
            }
            v1 -= (v2 + (v2 ^ 0))
            for fi26 := 0; fi26 < 8; fi26++ {
                fmt.Println(v3)
                delete(v5, 32767)
                _ = fi26
            }
        }
        if (!true) {
            v2 = (-(v2))
        } else {
            _ = len(v5)
        }
        switch ("foo bar" + "hello") {
        case "foo bar":
            for fi27 := 0; fi27 < 20; fi27++ {
                v28 := v5[(fi27 ^ (v4 | v4))]
                if true {
                    switch (v28 % ((int(v2)) | 1)) {
                    case -826:
                        v6 = v6
                    default:
                        switch "" {
                        case "\n\t":
                            var v29 int16 = 32767
                            _ = v29
                        default:
                            v6 += 0.0
                        }
                    }
                    v4 -= -84
                }
                v2 += -9223372036854775808
                _ = fi27
                _ = v28
            }
        case "hello":
            if false {
                var v30 uint = v3
                var v31 float32 = 0.0
                _ = v30
                _ = v31
            }
        default:
            var v32 float32 = float32(((v6 * (1.5 - (float64((0.0 / (3.14159 * (1e10 + -2.25)))) - ((v6 - -2.25) / v6)))) - v6))
            _ = v32
        }
    }
    if (746 != 955) {
        var v33 float32 = 0.0
        _ = v33
    } else {
        for fi34 := 0; fi34 < 10; fi34++ {
            var v35 int16 = -524
            v4 += -219
            v36 := GenS1{F0: 2, F1: "foo bar"}
            _ = fi34
            _ = v35
            _ = v36
        }
        var v37 uint16 = 1
        _ = v37
    }
    v1 = v1
    if true {
        var v38 int32 = 2147483647
        v39 := GenS4{F0: 0, F1: -1, F2: 4294967295}
        _ = v38
        _ = v39
    } else {
        v40 := GenS2{F0: (v2 / ((-1) | 1)), F1: 317, F2: (!true)}
        _ = v40
    }
    for fi41 := 0; fi41 < 17; fi41++ {
        var v42 uint16 = 554
        for fi43 := 0; fi43 < 2; fi43++ {
            v44 := GenS5{F0: 74, F1: 2147483648}
            var v45 float64 = float64(-2.25)
            _ = fi43
            _ = v44
            _ = v45
        }
        _ = fi41
        _ = v42
    }
    switch "foo bar" {
    case "\n\t":
        for fi46 := 0; fi46 < 4; fi46++ {
            v1 = -516
            _ = len(v5)
            _ = fi46
        }
    default:
        switch (v4 >> (((v3 << ((42) & 63))) & 63)) {
        case 32767:
            v47 := GenS1{F0: 2147483647, F1: ""}
            _ = v47
        default:
            v48 := GenS6{F0: 1000, F1: ""}
            _ = v48
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
