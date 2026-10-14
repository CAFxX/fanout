package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64; F2 string }
type GenS2 struct { F0 uint }
func genFn0(p0 int16, p1 int16) int {
    var v5 int64 = -9223372036854775808
    v5 = (v5 % ((-9223372036854775808) | 1))
    for fi6 := 0; fi6 < 6; fi6++ {
        var v7 int8 = -128
        _ = fi6
        _ = v7
    }
    for fi8 := 0; fi8 < 4; fi8++ {
        p1 -= -32
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = v5
    return 42
}
func genFn1(p0 uint16) bool {
    if false {
        v9 := GenS1{F0: 493, F1: 656, F2: ("a" + ("foo bar" + "a"))}
        p0 = 65535
        p0 = (p0 << ((1) & 63))
        _ = v9
    }
    p0 += 255
    var v10 int = -1
    var v11 uint32 = 2147483647
    _ = p0
    _ = v10
    _ = v11
    return false
}

func main() {
    var v1 uint16 = 2
    var v2 uint16 = v1
    var v3 uint16 = 255
    var v4 uint32 = 2147483647
    if (!false) {
        v2 = v2
        var v12 float32 = (0.1 * (1.5 - 0.0))
        _ = v12
    }
    var v13 int = 1
    if (19 <= 2) {
        var v14 int8 = -52
        _ = v14
    } else {
        var v15 bool = (18446744073709551615 <= 4294967295)
        _ = v15
    }
    switch ("a" + "hello") {
    case "hello":
        var v16 uint64 = 0
        _ = v16
    case "foo bar":
        fmt.Println(v4, v13)
    case "\n\t":
        v3 += v2
    default:
        v2 = v3
    }
    if true {
        v2 += (v1 * 255)
    } else {
        v2 -= 255
        v4 = 372
        for fi17 := 0; fi17 < 2; fi17++ {
            v18 := GenS1{F0: v4, F1: 2, F2: ((("\n\t" + ("foo bar" + ("hello" + "\n\t"))) + ("foo bar" + (("" + ("foo bar" + (("" + "a") + "hello"))) + "foo bar"))) + "")}
            for fi19 := 0; fi19 < 20; fi19++ {
                v13 -= (fi17 | fi17)
                v20 := GenS2{F0: 4294967295}
                v3 = v1
                _ = fi19
                _ = v20
            }
            switch "foo bar" {
            case "\n\t":
                v21 := []int{512, (-(fi17)), (-(fi17))}
                _ = v21
            case "":
                var v22 bool = true
                _ = v22
            case "hello":
                if true {
                    var v23 int8 = 52
                    var v24 string = "\n\t"
                    var v25 float32 = -2.25
                    _ = v23
                    _ = v24
                    _ = v25
                } else {
                    switch (fi17 & (fi17 | v13)) {
                    case 4294967295:
                        var v26 int8 = 42
                        _ = v26
                    default:
                        fmt.Println(v3)
                    }
                }
            }
            _ = fi17
            _ = v18
        }
    }
    if true {
        v27 := []int{}
        v27 = append(v27, v13)
        _ = v27
    } else {
        var v28 bool = (!((65535 == 1) || true))
        _ = v28
    }
    var v29 int16 = 678
    if false {
        switch "" {
        case "foo bar":
            v30 := []int{}
            _ = v30
        }
    } else {
        if false {
            v31 := GenS1{F0: (v4 >> ((255) & 63)), F1: 42, F2: "hello"}
            var v32 float32 = (float32(0.1) + ((3.14159 - 0.0) * 1.5))
            _ = v31
            _ = v32
        } else {
            var v33 int8 = 0
            _ = v33
        }
        v34 := []int{(v13 + v13)}
        v13 = -656
        _ = v34
    }
    v3 -= (v3 << ((508) & 63))
    v35 := []int{(v13 * v13), -9223372036854775808, (-(v13)), 905}
    for fi36 := 0; fi36 < 11; fi36++ {
        v2 = v3
        v37 := GenS1{F0: (^(v4)), F1: 4294967295, F2: ""}
        var v38 uint32 = 1
        _ = fi36
        _ = v37
        _ = v38
    }
    fmt.Println(v2, v1, v29)
    var v39 float64 = 1e10
    if len(v35) >= 2 {
        v35 = v35[1:len(v35)]
    }
    var v40 string = ""
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v13
    _ = v2
    _ = v29
    _ = v3
    _ = v35
    _ = v39
    _ = v4
    _ = v40
}
