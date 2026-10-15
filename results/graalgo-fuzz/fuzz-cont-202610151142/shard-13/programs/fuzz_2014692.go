package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 uint8 }
type GenS2 struct { F0 uint16; F1 int32; F2 uint }

func main() {
    var v1 uint64 = 401
    var v2 int64 = 42
    var v3 float32 = 0.1
    var v4 int = 445
    var v5 uint32 = 788
    var v6 uint8 = 2
    if false {
        if true {
            for fi7 := 0; fi7 < 16; fi7++ {
                v8 := GenS1{F0: v6, F1: (("a" + "") + "a"), F2: (v6 << ((2) & 63))}
                v9 := []int{32767}
                var v10 uint64 = v1
                _ = fi7
                _ = v10
                _ = v8
                _ = v9
            }
        }
        v11 := GenS1{F0: v6, F1: "a", F2: 123}
        _ = v11
    }
    if false {
        for fi12 := 0; fi12 < 8; fi12++ {
            if false {
                v13 := GenS2{F0: 42, F1: -23, F2: uint(v5)}
                v14 := GenS2{F0: 65535, F1: -1, F2: 65535}
                _ = v13
                _ = v14
            }
            _ = fi12
        }
        v4 = 9223372036854775807
        for fi15 := 0; fi15 < 9; fi15++ {
            v2 = -961
            switch "foo bar" {
            case "\n\t":
                switch (("foo bar" + "foo bar") + "\n\t") {
                case "foo bar":
                    v1 += v1
                default:
                    var v16 uint32 = v5
                    _ = v16
                }
            case "foo bar":
                var v17 string = "hello"
                _ = v17
            }
            _ = fi15
        }
    }
    if (((1 <= -473) || (false || ((v5 << ((uint(v6)) & 63)) > (v5 + (v5 >> ((0) & 63)))))) || true) {
        v1 = 1
    } else {
        for fi18 := 0; fi18 < 2; fi18++ {
            if true {
                v19 := []int{fi18, (fi18 << ((42) & 63))}
                _ = v19
            }
            fmt.Println(v5, v5, v2)
            for fi20 := 0; fi20 < 17; fi20++ {
                var v21 int64 = v2
                _ = fi20
                _ = v21
            }
            _ = fi18
        }
    }
    switch (-(v4)) {
    case 717:
        var v22 uint = 2
        _ = v22
    default:
        if true {
            var v23 uint8 = (v6 << ((584) & 63))
            if (603 > 598) {
                if (0 <= 2) {
                    var v24 uint8 = v23
                    v4 = (v4 | 1)
                    _ = v24
                }
                var v25 int8 = 1
                _ = v25
            } else {
                v23 -= v6
                fmt.Println(v4)
            }
            _ = v23
        } else {
            for fi26 := 0; fi26 < 14; fi26++ {
                if true {
                    v2 -= v2
                }
                _ = fi26
            }
            if (v2 <= 224) {
                if true {
                    for fi27 := 0; fi27 < 16; fi27++ {
                        v28 := GenS1{F0: 128, F1: "", F2: (v6 | 42)}
                        v1 = (^(v1))
                        _ = fi27
                        _ = v28
                    }
                    for fi29 := 0; fi29 < 6; fi29++ {
                        var v30 uint32 = (^(v5))
                        _ = fi29
                        _ = v30
                    }
                }
                v6 += (v6 << ((436) & 63))
                var v31 int32 = 129
                _ = v31
            }
        }
    }
    v2 = v2
    v32 := GenS2{F0: 446, F1: 42, F2: 0}
    v6 += v6
    for fi33 := 0; fi33 < 9; fi33++ {
        v34 := GenS1{F0: v6, F1: (("foo bar" + "foo bar") + "a"), F2: v6}
        v5 -= 0
        _ = fi33
        _ = v34
    }
    var v35 uint8 = (^(v6))
    v1 = v1
    if ((-128 < 0) && (65535 <= 0)) {
        v6 = 183
    } else {
        if (42 <= (v2 & v2)) {
            switch "hello" {
            case "":
                var v36 int8 = -1
                _ = v36
            case "\n\t":
                switch ("a" + "") {
                case "hello":
                    v4 = v4
                default:
                    fmt.Println(v1)
                }
            }
            switch (("a" + (("a" + ("foo bar" + "\n\t")) + "")) + "hello") {
            case "a":
                var v37 uint = 65535
                _ = v37
            case "\n\t":
                v35 -= (v35 / ((v35) | 1))
            case "hello":
                switch v4 {
                case -2147483648:
                    var v38 int16 = -438
                    _ = v38
                case -1:
                    v2 = 0
                default:
                    v39 := GenS1{F0: (^(v6)), F1: ("foo bar" + "foo bar"), F2: v35}
                    _ = v39
                }
            default:
                var v40 uint32 = (^(v5))
                _ = v40
            }
            v4 = v4
        }
    }
    var v41 uint32 = (v5 + 0)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v35
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
