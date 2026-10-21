package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8; F1 int8 }
func genFn0(p0 bool) int8 {
    v5 := []int{536, 0, 1, -2147483648}
    p0 = (true || p0)
    for fi6 := 0; fi6 < 7; fi6++ {
        for fi7 := 0; fi7 < 4; fi7++ {
            p0 = true
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            v8 := GenS1{F0: 263}
            _ = fi7
            _ = v8
        }
        var v9 string = (("a" + "") + ("hello" + "hello"))
        for fi10 := 0; fi10 < 17; fi10++ {
            _ = len(v5)
            var v11 int32 = -2147483648
            v12 := GenS2{F0: 42, F1: 100}
            _ = len(v5)
            if (p0 || (!(0 < 283))) {
                var v13 uint64 = 18446744073709551615
                _ = v13
            }
            if (-432 == (v11 ^ 617)) {
                p0 = ((0.1 / -2.25) == 1e10)
                switch "foo bar" {
                case "":
                    var v14 uint64 = 2
                    _ = v14
                case "\n\t":
                    switch "" {
                    case "foo bar":
                        v5 = append(v5, fi10)
                    case "a":
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    case "":
                        v9 = v9
                    }
                }
            } else {
                switch (v9 + v9) {
                case "a":
                    var v15 int8 = 100
                    _ = v15
                default:
                    v11 -= v11
                }
                _ = len(v5)
                v11 = (-(v11))
            }
            fmt.Println(v11, fi10, v11)
            _ = fi10
            _ = v11
            _ = v12
        }
        _ = fi6
        _ = v9
    }
    for fi16 := 0; fi16 < 15; fi16++ {
        for fi17 := 0; fi17 < 13; fi17++ {
            p0 = (!p0)
            _ = fi17
        }
        var v18 uint64 = 2
        _ = fi16
        _ = v18
    }
    _ = p0
    _ = v5
    return -128
}

func main() {
    var v1 uint16 = 1
    var v2 bool = true
    v3 := make(map[string]bool)
    var v4 string = (("\n\t" + (("\n\t" + (("\n\t" + "\n\t") + "hello")) + ("a" + ("\n\t" + "foo bar")))) + ((("a" + "a") + "hello") + "a"))
    var v19 int32 = 42
    if (v2 || v2) {
        v20 := []int{933, -1}
        for fi21 := 0; fi21 < 9; fi21++ {
            var v22 uint32 = 42
            fmt.Println(v2, v1)
            _ = fi21
            _ = v22
        }
        if v2 {
            v2 = v2
            for fi23 := 0; fi23 < 12; fi23++ {
                v19 += 0
                _ = fi23
            }
        }
        _ = v20
    } else {
        for fi24 := 0; fi24 < 13; fi24++ {
            if false {
                v2 = (!true)
                var v25 uint8 = 196
                var v26 uint16 = v1
                _ = v25
                _ = v26
            } else {
                fmt.Println(v19, v4, v2)
                if v2 {
                    v1 = v1
                    switch v4 {
                    case "\n\t":
                        var v27 uint32 = 65535
                        _ = v27
                    case "hello":
                        var v28 int8 = -128
                        _ = v28
                    default:
                        v4 = "foo bar"
                    }
                    fmt.Println(fi24)
                }
                var v29 int8 = -82
                _ = v29
            }
            _ = fi24
        }
        for fi30 := 0; fi30 < 3; fi30++ {
            fmt.Println(fi30, v2, fi30)
            _ = fi30
        }
        fmt.Println(v19, v4, v2)
    }
    v31 := v3[(v4 + (v4 + "a"))]
    _ = len(v3)
    v32 := []int{275, 32768, 768}
    v33 := GenS1{F0: 18446744073709551615}
    var v34 int64 = 42
    if false {
        v35 := GenS1{F0: 2}
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
}
