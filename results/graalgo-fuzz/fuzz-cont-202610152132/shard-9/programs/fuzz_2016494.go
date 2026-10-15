package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8; F1 int64; F2 int8 }
type GenS4 struct { F0 int16 }
func genFn0(p0 string) bool {
    for fi7 := 0; fi7 < 6; fi7++ {
        if (true || false) {
            var v8 uint64 = 950
            var v9 bool = false
            p0 = (p0 + ("\n\t" + ""))
            _ = v8
            _ = v9
        }
        v10 := GenS1{F0: 128, F1: 65535}
        _ = fi7
        _ = v10
    }
    var v11 float64 = 0.1
    for fi12 := 0; fi12 < 18; fi12++ {
        p0 = p0
        v13 := GenS2{F0: 18446744073709551615}
        _ = fi12
        _ = v13
    }
    if (!true) {
        var v14 float64 = 0.1
        p0 = "a"
        _ = v14
    }
    v11 += v11
    _ = p0
    _ = v11
    return true
}
func genFn1(p0 uint64, p1 int) int32 {
    v15 := GenS3{F0: 1, F1: 42, F2: 0}
    v16 := GenS2{F0: 4294967295}
    _ = p0
    _ = p1
    _ = v15
    _ = v16
    return int32(p0)
}

func main() {
    var v1 string = (("hello" + "foo bar") + ("\n\t" + "hello"))
    var v2 uint64 = 1
    var v3 string = v1
    var v4 uint64 = (v2 | 619)
    v5 := []string{((v3 + ("a" + "\n\t")) + "a"), v3, (((v3 + ((v1 + v3) + v1)) + "foo bar") + (v3 + v3)), ("foo bar" + "a")}
    var v6 int64 = -1
    _ = len(v5)
    var v17 uint8 = 255
    v2 -= v2
    switch (v1 + v3) {
    case "":
        v18 := GenS1{F0: (^(v17)), F1: 0}
        _ = v18
    default:
        v19 := GenS4{F0: -154}
        _ = v19
    }
    switch 42 {
    case -2147483648:
        switch v3 {
        case "a":
            switch ((v3 + v3) + v1) {
            case "":
                var v20 uint32 = 2147483648
                _ = v20
            case "\n\t":
                fmt.Println(v1, v6, v2)
            case "a":
                v21 := GenS3{F0: 127, F1: v6, F2: -128}
                _ = v21
            default:
                for fi22 := 0; fi22 < 9; fi22++ {
                    if false {
                        var v23 uint32 = 2147483647
                        v17 += (v17 ^ (v17 * v17))
                        var v24 uint = 42
                        _ = v23
                        _ = v24
                    }
                    var v25 int32 = 2147483647
                    _ = fi22
                    _ = v25
                }
            }
        case "hello":
            for fi26 := 0; fi26 < 6; fi26++ {
                fmt.Println(v6, v17, fi26)
                _ = fi26
            }
        }
    }
    for fi27 := 0; fi27 < 7; fi27++ {
        var v28 uint32 = 0
        _ = fi27
        _ = v28
    }
    if false {
        v4 = (^(v4))
        for fi29 := 0; fi29 < 13; fi29++ {
            v4 += (v2 * 592)
            _ = fi29
        }
        if len(v5) > 0 {
            v5[0] = "foo bar"
        }
    } else {
        if true {
            if false {
                if true {
                    var v30 uint64 = (v4 | uint64(v17))
                    _ = v30
                }
            }
            for fi31 := 0; fi31 < 16; fi31++ {
                switch (fi31 % (((fi31 & 42)) | 1)) {
                case -2147483648:
                    var v32 int16 = 42
                    _ = v32
                default:
                    var v33 int8 = -62
                    _ = v33
                }
                if false {
                    v2 -= (v4 / ((18446744073709551615) | 1))
                } else {
                    v3 = v1
                    v17 = v17
                    v5 = append(v5, ((v1 + ((("" + "a") + v1) + (v3 + "\n\t"))) + "a"))
                }
                v3 = (((((v3 + "foo bar") + v3) + (((v3 + (("\n\t" + v1) + ("a" + v3))) + v3) + ("hello" + v3))) + v3) + "foo bar")
                _ = fi31
            }
        } else {
            var v34 float64 = 1e10
            _ = v34
        }
        switch -1 {
        case 32767:
            v35 := GenS4{F0: 890}
            _ = v35
        case 9223372036854775807:
            switch "hello" {
            case "hello":
                v2 = 1
            default:
                v3 = "\n\t"
            }
        case 0:
            v2 = (^(v4))
        }
        for fi36 := 0; fi36 < 2; fi36++ {
            var v37 int32 = 2147483647
            switch v3 {
            case "foo bar":
                var v38 int8 = 0
                _ = v38
            default:
                fmt.Println(v1, v37, v17)
            }
            if (-9223372036854775808 > v6) {
                var v39 uint = 490
                var v40 int = (fi36 << ((65535) & 63))
                v3 = (v3 + v3)
                _ = v39
                _ = v40
            }
            _ = fi36
            _ = v37
        }
    }
    switch v1 {
    case "hello":
        for fi41 := 0; fi41 < 18; fi41++ {
            var v42 string = "\n\t"
            fmt.Println(v2, fi41)
            fmt.Println(v4, fi41)
            _ = fi41
            _ = v42
        }
    default:
        v1 = ("a" + "foo bar")
    }
    v3 = ""
    fmt.Println(v2, v17, v1)
    var v43 float32 = float32(((0.1 + ((0.0 * 1.5) - ((1.5 * (0.1 * -2.25)) - 0.0))) * 1.5))
    if len(v5) > 0 {
        v5[0] = "a"
    }
    var v44 int64 = (^(v6))
    if false {
        v45 := GenS4{F0: 1000}
        var v46 uint8 = uint8(v4)
        var v47 uint = 153
        _ = v45
        _ = v46
        _ = v47
    }
    if (false || (!(9223372036854775807 > -1))) {
        if len(v5) > 0 {
            v5[0] = v3
        }
        var v48 int16 = -264
        _ = v48
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
