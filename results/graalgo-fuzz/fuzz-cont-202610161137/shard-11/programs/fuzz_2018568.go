package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint8; F1 int32; F2 int }
func genFn0() uint16 {
    var v4 bool = (!false)
    if v4 {
        if v4 {
            var v5 string = ("\n\t" + "foo bar")
            var v6 string = v5
            for fi7 := 0; fi7 < 19; fi7++ {
                fmt.Println(fi7)
                v8 := []int{(fi7 % (((fi7 << ((65535) & 63))) | 1)), fi7, 59}
                _ = fi7
                _ = v8
            }
            _ = v5
            _ = v6
        } else {
            v4 = (false && (((((0 > -12) || v4) || v4) && v4) || v4))
            v9 := GenS1{F0: 0, F1: 1}
            _ = v9
        }
        for fi10 := 0; fi10 < 17; fi10++ {
            var v11 int8 = -43
            v11 -= (v11 / ((v11) | 1))
            _ = fi10
            _ = v11
        }
        var v12 string = "hello"
        _ = v12
    }
    for fi13 := 0; fi13 < 2; fi13++ {
        if v4 {
            for fi14 := 0; fi14 < 6; fi14++ {
                var v15 float32 = 1.5
                _ = fi14
                _ = v15
            }
        } else {
            var v16 int8 = 19
            var v17 float32 = ((-2.25 * (((0.0 + -2.25) - float32((3.14159 - 3.14159))) + 3.14159)) * (-2.25 * 3.14159))
            _ = v16
            _ = v17
        }
        _ = fi13
    }
    _ = v4
    return 32767
}

func main() {
    var v1 uint32 = 2
    v2 := make(map[string]int)
    var v3 int16 = 1
    v1 = v1
    var v18 uint16 = 32767
    v2[""] = -320
    var v19 uint8 = 127
    for fi20 := 0; fi20 < 7; fi20++ {
        var v21 uint = 1
        for fi22 := 0; fi22 < 20; fi22++ {
            v18 = v18
            v18 = genFn0()
            _ = fi22
        }
        _ = fi20
        _ = v21
    }
    if true {
        v23 := GenS2{F0: (v1 | 1)}
        _ = v23
    } else {
        fmt.Println(v3, v19, v19)
        v1 = v1
        if false {
            for fi24 := 0; fi24 < 16; fi24++ {
                v18 = (v18 % ((v18) | 1))
                switch "" {
                case "":
                    if (!(true || true)) {
                        if false {
                            v25 := GenS2{F0: (^(v1))}
                            _ = v25
                        }
                        for fi26 := 0; fi26 < 3; fi26++ {
                            v2["\n\t"] = (-(fi26))
                            v19 += v19
                            _ = len(v2)
                            _ = fi26
                        }
                        v19 = (v19 % ((v19) | 1))
                    }
                case "\n\t":
                    for fi27 := 0; fi27 < 13; fi27++ {
                        fmt.Println(v3, v3)
                        var v28 int32 = 2147483647
                        _ = fi27
                        _ = v28
                    }
                case "foo bar":
                    var v29 int32 = 468
                    _ = v29
                }
                _ = fi24
            }
            var v30 bool = true
            v31 := GenS3{F0: v19, F1: -2147483648, F2: -2147483648}
            _ = v30
            _ = v31
        } else {
            v32 := GenS1{F0: 127, F1: 4294967295}
            var v33 int32 = -1
            var v34 float32 = (((3.14159 * float32(1e10)) - 0.1) + -2.25)
            _ = v32
            _ = v33
            _ = v34
        }
    }
    fmt.Println(v18, v19)
    var v35 float32 = 1e10
    fmt.Println(v1)
    for fi36 := 0; fi36 < 18; fi36++ {
        v18 = 940
        _ = fi36
    }
    v37 := GenS3{F0: v19, F1: 0, F2: -1}
    var v38 int8 = -1
    var v39 int8 = v38
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v35
    _ = v37
    _ = v38
    _ = v39
}
