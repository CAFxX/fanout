package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 string; F2 uint8 }
type GenS3 struct { F0 uint32; F1 string }
type GenS4 struct { F0 uint32; F1 int8; F2 int32 }

func main() {
    var v1 float64 = ((1e10 + -2.25) * -2.25)
    v2 := []string{"foo bar"}
    var v3 uint8 = 1
    var v4 int16 = 32767
    var v5 uint8 = 42
    var v6 int64 = 870
    if false {
        var v7 uint32 = 42
        _ = v7
    }
    switch ("hello" + "a") {
    case "foo bar":
        for fi8 := 0; fi8 < 15; fi8++ {
            for fi9 := 0; fi9 < 18; fi9++ {
                v3 = (^(v3))
                var v10 float32 = 0.0
                v11 := GenS1{F0: ("hello" + ("\n\t" + "\n\t"))}
                if true {
                    switch "" {
                    case "hello":
                        v3 = 204
                    case "":
                        var v12 float64 = float64(v10)
                        v13 := GenS1{F0: (("foo bar" + (("foo bar" + (("" + (("\n\t" + "\n\t") + "\n\t")) + (("foo bar" + ("\n\t" + "a")) + "hello"))) + ("a" + "foo bar"))) + "")}
                        v14 := GenS1{F0: ""}
                        _ = v12
                        _ = v13
                        _ = v14
                    }
                } else {
                    var v15 uint64 = 100
                    _ = v15
                }
                fmt.Println(v10)
                _ = fi9
                _ = v10
                _ = v11
            }
            var v16 uint32 = 507
            v17 := GenS1{F0: "\n\t"}
            _ = fi8
            _ = v16
            _ = v17
        }
    case "\n\t":
        switch -401 {
        case 802:
            switch "" {
            case "":
                if len(v2) > 0 {
                    v2[0] = "foo bar"
                }
            }
        case 42:
            var v18 int = 4294967295
            _ = v18
        default:
            v6 = v6
        }
    case "":
        for fi19 := 0; fi19 < 12; fi19++ {
            v20 := GenS1{F0: "\n\t"}
            v21 := GenS2{F0: (v4 & (v4 + v4)), F1: "foo bar", F2: v3}
            var v22 uint = 18446744073709551615
            _ = fi19
            _ = v20
            _ = v21
            _ = v22
        }
    }
    fmt.Println(v5, v1, v3)
    var v23 float64 = v1
    if false {
        if (127 > -128) {
            var v24 float32 = -2.25
            v3 = (^(v3))
            _ = v24
        }
    } else {
        var v25 int = 128
        switch v25 {
        case 42:
            var v26 bool = ((!(0.1 > v23)) || (!(-1 < -2147483648)))
            _ = v26
        case -799:
            v4 -= v4
        }
        _ = v25
    }
    v27 := GenS1{F0: ("foo bar" + ("\n\t" + (("" + "foo bar") + "a")))}
    v5 = 1
    v2 = append(v2, "\n\t")
    v28 := GenS1{F0: ""}
    if (!((128 != 9223372036854775807) && (((-2147483648 >= -34) && (true || (1000 <= v4))) && false))) {
        var v29 int16 = 1
        for fi30 := 0; fi30 < 5; fi30++ {
            var v31 bool = (1 < 0)
            _ = fi30
            _ = v31
        }
        _ = v29
    } else {
        if len(v2) > 0 {
            v2[0] = "\n\t"
        }
        v23 -= v1
        switch -1 {
        case 35:
            for fi32 := 0; fi32 < 2; fi32++ {
                v33 := GenS3{F0: 4294967295, F1: ""}
                v3 = 42
                switch "\n\t" {
                case "a":
                    var v34 uint32 = 2147483648
                    _ = v34
                }
                _ = fi32
                _ = v33
            }
        case 937:
            v35 := GenS2{F0: v4, F1: ("hello" + ""), F2: v3}
            _ = v35
        }
    }
    v6 = (v6 - (-(v6)))
    switch "\n\t" {
    case "hello":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    default:
        v6 -= -100
    }
    for fi36 := 0; fi36 < 18; fi36++ {
        v3 = (v5 >> ((4294967295) & 63))
        _ = fi36
    }
    var v37 float64 = -2.25
    v38 := GenS3{F0: uint32(v5), F1: ""}
    for fi39 := 0; fi39 < 12; fi39++ {
        var v40 uint32 = 245
        v41 := GenS4{F0: (v40 >> ((uint(v40)) & 63)), F1: -116, F2: 811}
        for fi42 := 0; fi42 < 2; fi42++ {
            var v43 int32 = 42
            _ = fi42
            _ = v43
        }
        _ = fi39
        _ = v40
        _ = v41
    }
    var v44 uint16 = 0
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v27
    _ = v28
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v44
    _ = v5
    _ = v6
}
