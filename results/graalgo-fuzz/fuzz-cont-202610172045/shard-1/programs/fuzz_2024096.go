package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 uint8 }
func genFn0() int16 {
    var v4 int64 = 0
    for fi5 := 0; fi5 < 19; fi5++ {
        v6 := []int{128, (-(fi5)), fi5, fi5}
        var v7 uint16 = 946
        _ = fi5
        _ = v6
        _ = v7
    }
    var v8 int64 = (v4 % (((^(v4))) | 1))
    v9 := GenS1{F0: 2147483647, F1: "", F2: 42}
    v4 = v8
    _ = v4
    _ = v8
    _ = v9
    return -256
}

func main() {
    var v1 uint64 = 28
    var v2 int8 = 42
    v3 := []string{}
    var v10 float32 = 1.5
    if true {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        for fi11 := 0; fi11 < 2; fi11++ {
            for fi12 := 0; fi12 < 12; fi12++ {
                if len(v3) > 0 {
                    v3[0] = ""
                }
                if len(v3) > 0 {
                    v3[0] = "a"
                }
                _ = fi12
            }
            var v13 float32 = float32((0.1 * 3.14159))
            for fi14 := 0; fi14 < 3; fi14++ {
                var v15 int = (fi11 + (fi11 & fi14))
                var v16 uint8 = 107
                _ = fi14
                _ = v15
                _ = v16
            }
            _ = fi11
            _ = v13
        }
    }
    for fi17 := 0; fi17 < 7; fi17++ {
        var v18 int64 = -9223372036854775808
        v3 = append(v3, ("hello" + "\n\t"))
        var v19 int32 = 42
        _ = fi17
        _ = v18
        _ = v19
    }
    if len(v3) > 0 {
        v3[0] = ((("\n\t" + ("foo bar" + (("a" + "hello") + "hello"))) + "a") + ("hello" + ("hello" + "\n\t")))
    }
    if (false && false) {
        var v20 float64 = 0.0
        _ = v20
    } else {
        v21 := GenS1{F0: 0, F1: "\n\t", F2: 2}
        v22 := GenS1{F0: -1, F1: "", F2: 2}
        var v23 uint8 = 128
        _ = v21
        _ = v22
        _ = v23
    }
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v2 -= v2
    v24 := GenS1{F0: -1, F1: "", F2: 2}
    _ = len(v3)
    if true {
        switch -488 {
        case 32768:
            v3 = append(v3, (("a" + ("" + "\n\t")) + ""))
        default:
            var v25 uint32 = 42
            _ = v25
        }
        v10 = (-2.25 + v10)
        for fi26 := 0; fi26 < 20; fi26++ {
            for fi27 := 0; fi27 < 10; fi27++ {
                var v28 bool = (false && false)
                v29 := GenS1{F0: 265, F1: "", F2: 2}
                var v30 float64 = 1e10
                _ = fi27
                _ = v28
                _ = v29
                _ = v30
            }
            for fi31 := 0; fi31 < 4; fi31++ {
                v32 := GenS1{F0: 0, F1: "a", F2: 117}
                if (42 >= 127) {
                    switch fi31 {
                    case -1:
                        var v33 string = "foo bar"
                        _ = v33
                    }
                    v1 -= (v1 % ((v1) | 1))
                }
                v10 = v10
                _ = len(v3)
                var v34 bool = (894 != 837)
                _ = fi31
                _ = v32
                _ = v34
            }
            _ = fi26
        }
    } else {
        v2 += 1
    }
    if true {
        switch ("hello" + "") {
        case "hello":
            v2 = v2
        case "a":
            if ((false || true) || (490 > -312)) {
                fmt.Println(v10)
                if len(v3) > 0 {
                    v3[0] = "\n\t"
                }
                var v35 uint16 = 255
                _ = v35
            }
        case "foo bar":
            var v36 int16 = genFn0()
            _ = v36
        default:
            v10 = (v10 + 0.0)
        }
        if (18446744073709551615 <= 302) {
            v10 -= (v10 - ((v10 - v10) * float32(1.5)))
        } else {
            v10 = ((0.0 / float32((-2.25 + ((-2.25 + 1e10) + 3.14159)))) + float32(0.0))
            v1 = v1
            if (float64(v10) != (1.5 / -2.25)) {
                v2 = -128
            } else {
                _ = len(v3)
                var v37 uint32 = 2147483648
                if len(v3) > 0 {
                    v3[0] = "a"
                }
                _ = v37
            }
        }
        v10 -= ((v10 - v10) - float32((0.1 - ((3.14159 + 1e10) / ((1.5 / 1.5) / 1.5)))))
    }
    v1 = (v1 % ((v1) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v10
    _ = v2
    _ = v24
    _ = v3
}
