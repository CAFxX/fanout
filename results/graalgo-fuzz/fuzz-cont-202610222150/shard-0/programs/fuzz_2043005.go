package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 int8 }
func genFn0(p0 bool, p1 int16) uint16 {
    var v6 uint8 = 128
    v7 := []int{42, 32768, 9223372036854775807}
    if len(v7) > 0 {
        v7[0] = 2147483647
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 1
}
func genFn1(p0 string, p1 uint8, p2 uint8) int32 {
    fmt.Println(p0)
    v8 := GenS1{F0: (((true && (false && true)) || (-32768 <= 1000)) || true), F1: 208}
    var v9 int16 = 42
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 int8 = 0
    v2 := make(map[string]string)
    var v3 bool = (!(((!false) || (false || (1 != 145))) && (42 != 830)))
    var v4 int64 = -9223372036854775808
    v5 := []string{("" + "a"), ("a" + "foo bar")}
    for fi10 := 0; fi10 < 8; fi10++ {
        v3 = false
        var v11 string = "a"
        _ = fi10
        _ = v11
    }
    var v12 uint16 = genFn0(v3, 32767)
    var v13 float64 = (((1.5 - ((((((0.0 + (0.0 - (((3.14159 * 0.0) * 1e10) - (1.5 / 1e10)))) + float64(3.14159)) + (1e10 + 0.1)) + 0.0) + 1e10) - float64((float32((float64(0.0) / 3.14159)) - 1e10)))) + 0.1) - float64(0.1))
    v2["a"] = ("\n\t" + "hello")
    for fi14 := 0; fi14 < 5; fi14++ {
        v1 = (v1 >> ((42) & 63))
        _ = fi14
    }
    var v15 float64 = ((float64(0.1) / (((((v13 / (v13 + v13)) / (float64(0.1) * (v13 - (v13 + 0.1)))) / v13) * (v13 / v13)) / (float64(3.14159) * v13))) / (v13 - 0.0))
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi16 := 0; fi16 < 2; fi16++ {
        if (false || (1 <= 2)) {
            for fi17 := 0; fi17 < 18; fi17++ {
                var v18 uint32 = 2
                v12 -= v12
                if true {
                    v18 = 0
                }
                _ = fi17
                _ = v18
            }
        }
        var v19 uint32 = 2147483648
        v13 = (v15 - 0.0)
        _ = fi16
        _ = v19
    }
    v3 = ("foo bar" >= ("\n\t" + "foo bar"))
    var v20 int32 = 162
    var v21 int16 = 792
    switch 127 {
    case 128:
        fmt.Println(v13, v15)
    case 1:
        var v22 int8 = -1
        _ = v22
    case -396:
        v21 = (v21 - (v21 | (v21 % ((v21) | 1))))
    }
    switch 9223372036854775807 {
    case 32767:
        v23 := GenS1{F0: v3, F1: -558}
        _ = v23
    }
    var v24 uint32 = 1
    for fi25 := 0; fi25 < 12; fi25++ {
        _ = len(v2)
        v26 := GenS1{F0: false, F1: (fi25 >> ((1) & 63))}
        var v27 string = "\n\t"
        _ = fi25
        _ = v26
        _ = v27
    }
    for fi28 := 0; fi28 < 5; fi28++ {
        for fi29 := 0; fi29 < 8; fi29++ {
            v30 := GenS2{F0: v1}
            _ = fi29
            _ = v30
        }
        _ = fi28
    }
    switch (("foo bar" + "\n\t") + "\n\t") {
    case "":
        for fi31 := 0; fi31 < 6; fi31++ {
            v15 = float64(1.5)
            v12 = (^(v12))
            _ = fi31
        }
    case "hello":
        if ((v3 || v3) || v3) {
            delete(v2, ("foo bar" + ("" + ((("foo bar" + "a") + ("a" + ("a" + "hello"))) + ("foo bar" + "\n\t")))))
        } else {
            v32 := GenS2{F0: (v1 + (^(v1)))}
            var v33 float32 = (1e10 - 0.1)
            var v34 int16 = v21
            _ = v32
            _ = v33
            _ = v34
        }
    case "\n\t":
        v12 = (v12 | (v12 >> ((18446744073709551615) & 63)))
    }
    var v35 int64 = (v4 ^ 0)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v24
    _ = v3
    _ = v35
    _ = v4
    _ = v5
}
