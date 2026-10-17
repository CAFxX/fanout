package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 uint32; F1 bool; F2 uint16 }
func genFn0(p0 string, p1 uint32) bool {
    for fi7 := 0; fi7 < 10; fi7++ {
        switch p0 {
        case "\n\t":
            if false {
                var v8 int64 = 0
                if true {
                    p1 += p1
                } else {
                    v9 := []int{(fi7 / ((int(v8)) | 1)), (fi7 & 127), fi7}
                    if len(v9) > 0 {
                        v9[0] = fi7
                    }
                    var v10 uint16 = 1
                    _ = v10
                    _ = v9
                }
                _ = v8
            } else {
                v11 := GenS1{F0: "\n\t", F1: -320}
                _ = v11
            }
        }
        _ = fi7
    }
    p1 -= (p1 >> ((1) & 63))
    _ = p0
    _ = p1
    return (false && false)
}
func genFn1() int {
    var v12 int64 = 1
    v13 := GenS1{F0: "", F1: v12}
    for fi14 := 0; fi14 < 3; fi14++ {
        v12 += (v12 + v12)
        var v15 uint = 18446744073709551615
        _ = fi14
        _ = v15
    }
    var v16 float32 = 1.5
    _ = v12
    _ = v13
    _ = v16
    return 77
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint8 = 1
    var v3 uint32 = 24
    var v4 int8 = -128
    var v5 float64 = 0.0
    var v6 float64 = ((0.1 - float64(float32(-2.25))) / (v5 - (v5 - v5)))
    if ((v2 != 2) && (!true)) {
        for fi17 := 0; fi17 < 13; fi17++ {
            var v18 uint = 255
            _ = fi17
            _ = v18
        }
    }
    v4 = v4
    fmt.Println(v3, v3, v2)
    v2 -= (v2 + 238)
    for fi19 := 0; fi19 < 7; fi19++ {
        var v20 uint64 = 18446744073709551615
        for fi21 := 0; fi21 < 3; fi21++ {
            var v22 float32 = 1e10
            v23 := v1["hello"]
            v24 := []int{(fi21 + fi21), fi19, (fi21 >> ((4294967295) & 63)), (fi21 * (fi19 % ((127) | 1)))}
            _ = fi21
            _ = v22
            _ = v23
            _ = v24
        }
        v5 += (v5 + v6)
        _ = fi19
        _ = v20
    }
    v3 = 65535
    v2 = (v2 >> ((317) & 63))
    var v25 uint16 = 684
    switch "\n\t" {
    case "hello":
        v3 = (v3 * (^(v3)))
    }
    switch 817 {
    case 2147483647:
        var v26 uint64 = 4294967295
        _ = v26
    case -866:
        switch "" {
        case "\n\t":
            v6 = (v6 + float64((float32((-2.25 - 3.14159)) / 3.14159)))
        case "a":
            v6 = v6
        default:
            var v27 int8 = 90
            _ = v27
        }
    default:
        for fi28 := 0; fi28 < 4; fi28++ {
            switch 32767 {
            case 9223372036854775807:
                for fi29 := 0; fi29 < 7; fi29++ {
                    delete(v1, "foo bar")
                    v25 -= 65535
                    switch (("a" + (("\n\t" + ("foo bar" + "a")) + "")) + "foo bar") {
                    case "\n\t":
                        var v30 uint64 = 778
                        _ = v30
                    case "":
                        v1["a"] = false
                    default:
                        v4 += v4
                    }
                    if (32767 >= (fi29 ^ (fi28 | 9223372036854775807))) {
                        v31 := GenS1{F0: "foo bar", F1: 9223372036854775807}
                        _ = v31
                    } else {
                        if (true && (!true)) {
                            var v32 float64 = (float64(1e10) * ((v5 + v5) / v6))
                            var v33 uint64 = 275
                            _ = v32
                            _ = v33
                        } else {
                            v3 = (v3 - 845)
                            if (genFn1() != genFn1()) {
                                if false {
                                    v3 -= (v3 << ((255) & 63))
                                    var v34 uint32 = v3
                                    if (false && false) {
                                        v2 = (v2 * (v2 ^ (v2 | v2)))
                                    } else {
                                        v6 += float64(0.0)
                                        v35 := []int{-299, fi29}
                                        _ = v35
                                    }
                                    _ = v34
                                }
                                var v36 int16 = 1
                                var v37 float64 = 0.1
                                _ = v36
                                _ = v37
                            }
                        }
                    }
                    _ = fi29
                }
            default:
                v2 += (v2 ^ v2)
            }
            _ = fi28
        }
    }
    v38 := GenS2{F0: v3, F1: ((!(!true)) || true), F2: 255}
    var v39 int = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
