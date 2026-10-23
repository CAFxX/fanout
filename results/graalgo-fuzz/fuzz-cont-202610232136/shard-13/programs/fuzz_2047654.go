package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }

func main() {
    var v1 string = ("hello" + "")
    var v2 uint16 = 150
    v3 := make(map[uint32]string)
    v4 := make(map[int]string)
    var v5 uint8 = 127
    var v6 int = -2147483648
    for fi7 := 0; fi7 < 7; fi7++ {
        var v8 float64 = ((0.0 / ((((1e10 / -2.25) + float64(3.14159)) - 0.1) + 3.14159)) / ((1.5 * 0.1) / 1.5))
        _ = fi7
        _ = v8
    }
    var v9 bool = false
    v2 += (v2 ^ 65535)
    var v10 uint32 = 2147483648
    var v11 uint32 = v10
    if (!v9) {
        if (((!false) && true) || v9) {
            if v9 {
                v3[(v10 * v10)] = ((("foo bar" + v1) + v1) + v1)
                switch v1 {
                case "\n\t":
                    var v12 int16 = 42
                    _ = v12
                case "a":
                    fmt.Println(v6, v5)
                case "":
                    v9 = ((-234 > 532) && ((!(v5 > v5)) || true))
                default:
                    v13 := []int{612, v6}
                    _ = v13
                }
            }
        } else {
            fmt.Println(v1, v5)
            fmt.Println(v9, v2)
        }
        delete(v3, v10)
        v9 = v9
    }
    for fi14 := 0; fi14 < 14; fi14++ {
        if ((!v9) && false) {
            v15 := GenS1{F0: 0, F1: v1}
            fmt.Println(v5)
            _ = v15
        }
        _ = fi14
    }
    for fi16 := 0; fi16 < 8; fi16++ {
        switch ("\n\t" + (v1 + v1)) {
        case "":
            v6 -= (v6 / ((127) | 1))
        case "a":
            v17 := []int{v6, (v6 % ((fi16) | 1))}
            _ = v17
        case "hello":
            switch fi16 {
            case 770:
                if (-2147483648 > 0) {
                    if v9 {
                        v18 := []int{4294967295, int(v5), v6, fi16}
                        if (("\n\t" + v1) <= "a") {
                            v9 = (!false)
                            if len(v18) > 0 {
                                v18[0] = 1
                            }
                        }
                        fmt.Println(v6, v6)
                        _ = v18
                    } else {
                        v19 := []int{384, -109}
                        v6 = (v6 & (fi16 % ((v6) | 1)))
                        v6 = -9223372036854775808
                        _ = v19
                    }
                }
            case 127:
                v20 := v4[(fi16 >> ((65535) & 63))]
                _ = v20
            case 393:
                switch (v6 * (fi16 - (fi16 << ((4294967295) & 63)))) {
                case 1:
                    var v21 uint = 18446744073709551615
                    _ = v21
                }
            }
        }
        var v22 uint16 = (v2 ^ v2)
        _ = fi16
        _ = v22
    }
    var v23 uint64 = 42
    v10 = (v10 - v11)
    var v24 uint8 = v5
    var v25 int32 = 790
    for fi26 := 0; fi26 < 9; fi26++ {
        var v27 int16 = 1000
        switch "\n\t" {
        case "\n\t":
            if (!((v10 % (((v11 * 2147483648)) | 1)) <= v11)) {
                switch 689 {
                case 4294967295:
                    v27 = (-(v27))
                case -1:
                    var v28 int8 = -102
                    _ = v28
                case -9223372036854775808:
                    var v29 int8 = 127
                    _ = v29
                default:
                    v30 := GenS1{F0: (v6 + fi26), F1: ""}
                    _ = v30
                }
            }
        }
        _ = fi26
        _ = v27
    }
    for fi31 := 0; fi31 < 18; fi31++ {
        v32 := []int{}
        _ = fi31
        _ = v32
    }
    v33 := []int{(v6 ^ (v6 >> ((42) & 63))), v6, v6}
    fmt.Println(v9)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
