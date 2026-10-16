package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 string }
type GenS2 struct { F0 uint32; F1 string; F2 uint8 }
func genFn0(p0 uint64, p1 bool) uint32 {
    if p1 {
        p0 -= p0
        if ((476 != 1) || p1) {
            v4 := GenS1{F0: "", F1: 42, F2: "hello"}
            v5 := GenS1{F0: "foo bar", F1: 993, F2: "foo bar"}
            _ = v4
            _ = v5
        }
    } else {
        v6 := []int{128, -393, 658, -273}
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = v6
    }
    if false {
        for fi7 := 0; fi7 < 20; fi7++ {
            p1 = p1
            _ = fi7
        }
        switch ("hello" + ((("" + "\n\t") + (("" + (("\n\t" + "a") + "\n\t")) + "")) + "\n\t")) {
        case "":
            p1 = false
        case "a":
            v8 := GenS2{F0: 4294967295, F1: "\n\t", F2: 127}
            _ = v8
        case "\n\t":
            v9 := []int{0}
            _ = v9
        default:
            p1 = p1
        }
        var v10 string = (("" + "hello") + "")
        _ = v10
    } else {
        v11 := []int{-1, 978}
        _ = v11
    }
    v12 := []int{1, -484}
    switch "foo bar" {
    case "a":
        v12 = append(v12, 9223372036854775807)
    default:
        var v13 float32 = -2.25
        _ = v13
    }
    _ = p0
    _ = p1
    _ = v12
    return 445
}

func main() {
    var v1 int = 2147483647
    var v2 int = v1
    var v3 uint = 314
    v14 := []int{v2, 2147483647}
    for fi15 := 0; fi15 < 13; fi15++ {
        v2 += v1
        switch "foo bar" {
        case "":
            var v16 int32 = 18
            _ = v16
        case "\n\t":
            var v17 int16 = -32768
            _ = v17
        case "hello":
            var v18 int = (fi15 ^ v1)
            _ = v18
        }
        if true {
            var v19 uint64 = 2
            _ = v19
        } else {
            var v20 string = "foo bar"
            v1 += v1
            _ = v20
        }
        _ = fi15
    }
    if (!((1.5 < 3.14159) || true)) {
        var v21 int8 = -128
        for fi22 := 0; fi22 < 14; fi22++ {
            for fi23 := 0; fi23 < 19; fi23++ {
                v2 = fi22
                if (1e10 > 1.5) {
                    var v24 uint = (v3 ^ v3)
                    _ = v24
                }
                var v25 int = (fi23 % ((v1) | 1))
                _ = fi23
                _ = v25
            }
            _ = fi22
        }
        _ = v21
    }
    for fi26 := 0; fi26 < 19; fi26++ {
        var v27 float64 = 3.14159
        var v28 uint = (v3 ^ v3)
        fmt.Println(v28, v2, v27)
        _ = fi26
        _ = v27
        _ = v28
    }
    v14 = append(v14, 193)
    fmt.Println(v3, v3, v2)
    var v29 uint16 = 65535
    if false {
        if (false || (!true)) {
            v1 -= v1
        }
    }
    var v30 float32 = (0.0 - (float32(1.5) * 3.14159))
    v29 -= v29
    for fi31 := 0; fi31 < 7; fi31++ {
        v32 := GenS2{F0: 2147483648, F1: (("foo bar" + "hello") + (("" + (("" + ("a" + "")) + "hello")) + "hello")), F2: 89}
        v3 -= (^(v3))
        _ = fi31
        _ = v32
    }
    v3 -= 1
    if len(v14) > 0 {
        v14[0] = v1
    }
    v3 = v3
    if false {
        for fi33 := 0; fi33 < 19; fi33++ {
            if (true && (("\n\t" >= "\n\t") || (0 < 1))) {
                var v34 string = ((("" + ("a" + "a")) + "foo bar") + "\n\t")
                _ = len(v14)
                _ = v34
            } else {
                v2 -= (fi33 ^ -259)
                v3 = 65535
            }
            _ = fi33
        }
        if (false || true) {
            var v35 int8 = 57
            if len(v14) > 0 {
                v14[0] = v1
            }
            v29 += v29
            _ = v35
        }
        var v36 uint32 = 2
        _ = v36
    } else {
        for fi37 := 0; fi37 < 3; fi37++ {
            v1 += -9223372036854775808
            if (!false) {
                v29 += (^(v29))
                fmt.Println(v3, v3, v2)
                var v38 bool = true
                _ = v38
            }
            _ = fi37
        }
        if false {
            var v39 int8 = 100
            v30 -= v30
            _ = v39
        }
    }
    fmt.Println(v30, v2, v29)
    var v40 uint64 = 0
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v14
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v40
}
