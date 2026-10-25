package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 int64 }
type GenS2 struct { F0 uint64; F1 uint64; F2 uint }
type GenS3 struct { F0 bool; F1 int32; F2 int16 }

func main() {
    var v1 int16 = -32768
    var v2 uint = 1
    var v3 uint64 = 553
    v4 := make(map[string]string)
    v5 := []string{"hello", "foo bar"}
    fmt.Println(v2)
    v2 = (v2 & 0)
    switch "\n\t" {
    case "\n\t":
        if true {
            var v6 uint32 = 2147483647
            if false {
                v2 = v2
            }
            _ = v6
        } else {
            v2 += 576
        }
    default:
        if false {
            fmt.Println(v3, v3)
            for fi7 := 0; fi7 < 11; fi7++ {
                v5 = append(v5, "a")
                for fi8 := 0; fi8 < 10; fi8++ {
                    var v9 uint = (v2 + v2)
                    _ = fi8
                    _ = v9
                }
                _ = fi7
            }
            fmt.Println(v2, v3)
        }
    }
    if (42 != -1) {
        for fi10 := 0; fi10 < 14; fi10++ {
            var v11 int = fi10
            v4[""] = ("hello" + "a")
            _ = fi10
            _ = v11
        }
    } else {
        for fi12 := 0; fi12 < 18; fi12++ {
            if true {
                if (false || true) {
                    var v13 float64 = 3.14159
                    _ = v13
                } else {
                    v14 := GenS1{F0: 2147483648, F1: 0, F2: 42}
                    _ = v14
                }
            } else {
                var v15 float32 = float32(((float64(1e10) + -2.25) - ((float64(float32(3.14159)) - 1e10) / 1e10)))
                _ = v15
            }
            fmt.Println(v3, fi12, v3)
            if ("foo bar" != "") {
                if ("foo bar" >= "") {
                    v16 := v4[("a" + "foo bar")]
                    _ = v16
                }
            } else {
                var v17 int64 = 0
                var v18 float32 = (1.5 - -2.25)
                _ = v17
                _ = v18
            }
            _ = fi12
        }
        v19 := GenS1{F0: 2, F1: 0, F2: 9223372036854775807}
        _ = v19
    }
    fmt.Println(v3, v3, v2)
    v20 := GenS1{F0: 363, F1: 0, F2: 1}
    v2 -= v2
    if true {
        v1 -= v1
    } else {
        var v21 uint64 = 637
        var v22 int32 = 0
        _ = v21
        _ = v22
    }
    fmt.Println(v2)
    if false {
        v23 := GenS2{F0: (v3 | v3), F1: v3, F2: 42}
        fmt.Println(v1, v3)
        if (!(!(false && (42 > -2147483648)))) {
            var v24 int64 = 1
            var v25 bool = (!(200 != (v2 * v2)))
            _ = v24
            _ = v25
        }
        _ = v23
    } else {
        if false {
            switch 4294967295 {
            case -109:
                var v26 uint8 = 255
                _ = v26
            case 127:
                v2 = (v2 << ((v2) & 63))
            }
        }
        delete(v4, (("\n\t" + ("foo bar" + "\n\t")) + "a"))
        for fi27 := 0; fi27 < 7; fi27++ {
            v28 := GenS3{F0: (!(!((("foo bar" + ("foo bar" + (((("hello" + "hello") + ("" + "\n\t")) + "") + "a"))) >= ("\n\t" + "a")) || (!(42 < -1))))), F1: 42, F2: (^(v1))}
            var v29 uint64 = v3
            _ = fi27
            _ = v28
            _ = v29
        }
    }
    var v30 uint64 = (^(v3))
    for fi31 := 0; fi31 < 13; fi31++ {
        v32 := GenS1{F0: 2, F1: 795, F2: 9223372036854775807}
        switch "\n\t" {
        case "a":
            var v33 uint = (v2 % (((^(v2))) | 1))
            _ = v33
        case "":
            if len(v5) > 0 {
                v5[0] = "a"
            }
        default:
            delete(v4, ("foo bar" + "hello"))
        }
        var v34 int8 = 42
        _ = fi31
        _ = v32
        _ = v34
    }
    var v35 int8 = 65
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v35:", v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v20
    _ = v3
    _ = v30
    _ = v35
    _ = v4
    _ = v5
}
