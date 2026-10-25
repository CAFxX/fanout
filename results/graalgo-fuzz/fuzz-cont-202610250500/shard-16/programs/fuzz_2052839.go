package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 uint64; F1 int64; F2 int }
type GenS3 struct { F0 bool; F1 int; F2 uint64 }
type GenS4 struct { F0 uint16 }
func genFn0(p0 int8) bool {
    for fi7 := 0; fi7 < 2; fi7++ {
        switch fi7 {
        case 32767:
            v8 := GenS1{F0: "hello", F1: 243}
            _ = v8
        case -1:
            v9 := GenS2{F0: 0, F1: -1, F2: -2147483648}
            _ = v9
        default:
            if (!true) {
                p0 -= p0
                for fi10 := 0; fi10 < 14; fi10++ {
                    p0 = -128
                    p0 += (p0 % (((p0 << ((854) & 63))) | 1))
                    var v11 int = (-(fi7))
                    _ = fi10
                    _ = v11
                }
                v12 := GenS1{F0: ("foo bar" + (((((("a" + "") + "foo bar") + "a") + "foo bar") + "hello") + ((("hello" + "a") + "foo bar") + ""))), F1: 145}
                _ = v12
            } else {
                for fi13 := 0; fi13 < 9; fi13++ {
                    v14 := GenS3{F0: ((((("a" + "hello") + "foo bar") + "hello") == "hello") || (0.1 > float64((1e10 / (float32(1e10) / ((-2.25 * -2.25) * ((1.5 / 1.0) / 1.0))))))), F1: fi13, F2: 4294967295}
                    p0 -= (p0 - (p0 ^ (^(p0))))
                    switch (-(fi7)) {
                    case 4294967295:
                        if (((0.0 / 0.1) / (1e10 * (-2.25 - (0.1 - (float32(3.14159) * ((0.1 + 1.5) + (1e10 / 1.5))))))) > 1.5) {
                            p0 = (^(p0))
                            v15 := GenS1{F0: "\n\t", F1: 1}
                            p0 += p0
                            v16 := []int{}
                            _ = v15
                            _ = v16
                        }
                    case 32767:
                        p0 -= (p0 - (p0 ^ (p0 >> ((753) & 63))))
                        v17 := GenS4{F0: 32767}
                        _ = v17
                    default:
                        v18 := []int{666}
                        _ = v18
                    }
                    _ = fi13
                    _ = v14
                }
                var v19 uint8 = 42
                for fi20 := 0; fi20 < 8; fi20++ {
                    var v21 string = "foo bar"
                    _ = fi20
                    _ = v21
                }
                _ = v19
            }
        }
        v22 := []int{338, 4294967295, -929}
        v23 := GenS4{F0: 175}
        _ = fi7
        _ = v22
        _ = v23
    }
    if ((((0.1 / 1.5) * -2.25) * float32(((0.0 - 0.0) * -2.25))) >= (1.5 * 1.5)) {
        v24 := GenS3{F0: true, F1: -9223372036854775808, F2: 290}
        switch ("a" + "\n\t") {
        case "a":
            switch 418 {
            case -257:
                fmt.Println(p0)
            case 32768:
                v25 := GenS3{F0: (!(!("foo bar" < "hello"))), F1: -2147483648, F2: 0}
                _ = v25
            case 4294967295:
                p0 = p0
            }
        default:
            switch -2147483648 {
            case 4294967295:
                v26 := GenS4{F0: 2}
                _ = v26
            case 603:
                var v27 string = (("a" + "foo bar") + "\n\t")
                _ = v27
            default:
                v28 := []int{32768, -9223372036854775808}
                _ = v28
            }
        }
        _ = v24
    }
    p0 += -15
    _ = p0
    return false
}

func main() {
    v1 := []int{}
    var v2 string = ""
    var v3 int = 0
    var v4 bool = true
    v5 := []string{v2, v2, "hello", "foo bar"}
    v6 := []int{(-(v3))}
    var v29 int64 = 36
    if len(v5) > 0 {
        v5[0] = (((v2 + (((v2 + v2) + (("a" + ("a" + (("a" + (v2 + v2)) + v2))) + v2)) + v2)) + ("" + ((("a" + (v2 + (v2 + ""))) + v2) + ("foo bar" + v2)))) + v2)
    }
    var v30 bool = v4
    v31 := GenS3{F0: (v2 <= (v2 + (v2 + (("\n\t" + "foo bar") + "foo bar")))), F1: (-(v3)), F2: 2}
    var v32 uint16 = 2
    var v33 uint8 = 122
    var v34 string = (v2 + "a")
    var v35 float64 = float64(((1.5 * ((float32(-2.25) / 1.0) + ((0.1 - 3.14159) + (((float32((0.0 + (float64(1e10) + 1.5))) / float32(3.14159)) * 3.14159) - ((-2.25 * 0.1) - 1e10))))) / -2.25))
    var v36 bool = v30
    var v37 string = v34
    if true {
        v38 := GenS2{F0: 42, F1: v29, F2: -9223372036854775808}
        _ = v38
    }
    _ = len(v5)
    v34 = v34
    var v39 float32 = 1.5
    if (false || ((2147483648 != 302) || v36)) {
        var v40 uint64 = 4294967295
        var v41 float32 = 3.14159
        v6 = append(v6, int(v32))
        _ = v40
        _ = v41
    }
    var v42 int64 = v29
    _ = len(v1)
    fmt.Println(v29)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v39
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
