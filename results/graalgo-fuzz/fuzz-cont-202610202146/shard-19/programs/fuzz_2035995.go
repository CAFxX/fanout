package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool; F2 string }
type GenS2 struct { F0 uint8 }
func genFn0() int32 {
    if ((128 <= 162) && (!(101 > -128))) {
        switch ("hello" + "") {
        case "":
            for fi4 := 0; fi4 < 2; fi4++ {
                if true {
                    if ((9223372036854775807 == 435) || (0 == 839)) {
                        v5 := []int{}
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                        _ = v5
                    }
                    if (-1 != -241) {
                        v6 := []int{}
                        _ = v6
                    }
                } else {
                    var v7 int64 = -598
                    _ = v7
                }
                v8 := []int{-2147483648, fi4, (fi4 << ((2) & 63))}
                _ = fi4
                _ = v8
            }
        default:
            for fi9 := 0; fi9 < 6; fi9++ {
                var v10 uint64 = 42
                _ = fi9
                _ = v10
            }
        }
        v11 := GenS1{F0: 127, F1: ((true && false) || ("" == "foo bar")), F2: "foo bar"}
        var v12 int16 = 1000
        _ = v11
        _ = v12
    }
    var v13 float64 = (((((float64(3.14159) + 0.1) - 1.5) - 1e10) + 1.5) - 0.1)
    var v14 int8 = -1
    _ = v13
    _ = v14
    return 879
}

func main() {
    v1 := []uint32{2147483648}
    var v2 uint16 = 1
    var v3 uint32 = 828
    v2 += v2
    v2 = 32767
    v1 = append(v1, 65535)
    v15 := GenS1{F0: 1, F1: true, F2: (((("foo bar" + "hello") + "hello") + "a") + "")}
    v2 = v2
    if true {
        var v16 uint8 = 42
        if (true && (-2.25 != 1e10)) {
            v2 = v2
            v3 += 2147483648
        } else {
            if (1 <= 612) {
                for fi17 := 0; fi17 < 8; fi17++ {
                    v3 = v3
                    _ = fi17
                }
                var v18 int = 32768
                _ = v18
            }
            v3 = (^(v3))
            if len(v1) > 0 {
                v1[0] = v3
            }
        }
        _ = v16
    }
    var v19 uint64 = 18446744073709551615
    v20 := GenS1{F0: 70, F1: (1 < -1), F2: "foo bar"}
    for fi21 := 0; fi21 < 18; fi21++ {
        v22 := GenS1{F0: 0, F1: (579 >= 2), F2: "a"}
        _ = fi21
        _ = v22
    }
    for fi23 := 0; fi23 < 10; fi23++ {
        if ((true || (false || true)) && (42 > -1)) {
            fmt.Println(v2, v2)
            if true {
                if len(v1) > 0 {
                    v1[0] = v3
                }
            } else {
                var v24 uint32 = (v3 & 326)
                v1 = append(v1, (v3 / ((49) | 1)))
                v25 := GenS1{F0: 25, F1: false, F2: "foo bar"}
                _ = v24
                _ = v25
            }
        } else {
            v3 = v3
        }
        for fi26 := 0; fi26 < 9; fi26++ {
            fmt.Println(v19, v3)
            v2 -= 2
            _ = fi26
        }
        v19 += 94
        _ = fi23
    }
    v27 := GenS2{F0: 42}
    if false {
        fmt.Println(v19, v2)
    }
    var v28 float32 = 1e10
    v19 = (v19 - (v19 / ((v19) | 1)))
    v29 := GenS1{F0: 42, F1: false, F2: ((("hello" + "\n\t") + "a") + "hello")}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v1 = append(v1, v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v27
    _ = v28
    _ = v29
    _ = v3
}
