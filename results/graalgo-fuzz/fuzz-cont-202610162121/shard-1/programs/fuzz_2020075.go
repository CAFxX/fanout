package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int; F1 int32 }
type GenS3 struct { F0 uint8 }
func genFn0(p0 int8) int8 {
    for fi6 := 0; fi6 < 20; fi6++ {
        fmt.Println(p0, fi6)
        v7 := []int{}
        _ = fi6
        _ = v7
    }
    switch -251 {
    case -9223372036854775808:
        fmt.Println(p0)
    case 127:
        v8 := []int{32768, -9223372036854775808, -1, 1}
        _ = v8
    case 260:
        var v9 uint8 = 2
        _ = v9
    }
    _ = p0
    return 100
}
func genFn1(p0 int16, p1 string, p2 int8) uint64 {
    p1 = ""
    switch 127 {
    case 9223372036854775807:
        p2 = p2
    case -1:
        p1 = p1
    default:
        var v10 int16 = p0
        _ = v10
    }
    v11 := []int{-694, -9223372036854775808, -2147483648}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 18446744073709551615
}

func main() {
    v1 := []string{(("foo bar" + ("" + ("\n\t" + (("" + "\n\t") + "hello")))) + "a"), "a", ("" + "\n\t"), ("hello" + "\n\t")}
    var v2 uint32 = 881
    var v3 bool = false
    var v4 bool = (v2 >= (v2 ^ v2))
    var v5 bool = v3
    var v12 bool = v5
    v13 := GenS1{F0: uint8(v2)}
    v5 = v5
    var v14 int32 = 1
    var v15 string = "hello"
    for fi16 := 0; fi16 < 8; fi16++ {
        fmt.Println(v3)
        _ = fi16
    }
    for fi17 := 0; fi17 < 15; fi17++ {
        for fi18 := 0; fi18 < 20; fi18++ {
            v19 := GenS2{F0: 32768, F1: (v14 ^ v14)}
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi18
            _ = v19
        }
        var v20 int64 = 503
        _ = fi17
        _ = v20
    }
    fmt.Println(v3, v4)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi21 := 0; fi21 < 18; fi21++ {
        if (!(!true)) {
            if v5 {
                for fi22 := 0; fi22 < 19; fi22++ {
                    var v23 uint8 = 128
                    _ = fi22
                    _ = v23
                }
            }
            fmt.Println(v12, v4, v14)
            if (v5 || ((3.14159 * (1e10 + 1e10)) <= (3.14159 / ((0.0 + float32((-2.25 - (0.1 - 3.14159)))) + -2.25)))) {
                fmt.Println(v5, v14)
            } else {
                _ = len(v1)
                for fi24 := 0; fi24 < 13; fi24++ {
                    var v25 int32 = 2147483647
                    var v26 uint16 = 2
                    switch "\n\t" {
                    case "foo bar":
                        if ((-(fi24)) != (fi21 & (fi24 & fi24))) {
                            v1 = append(v1, ("a" + ""))
                            v4 = false
                            v5 = v3
                        } else {
                            v27 := GenS2{F0: 314, F1: v25}
                            _ = v27
                        }
                    case "\n\t":
                        v25 = v25
                    case "a":
                        var v28 int8 = genFn0(genFn0(127))
                        v29 := GenS3{F0: 2}
                        fmt.Println(v28, v5, v15)
                        _ = v28
                        _ = v29
                    default:
                        _ = len(v1)
                    }
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = fi24
                    _ = v25
                    _ = v26
                }
                var v30 int32 = (v14 - -1)
                _ = v30
            }
        }
        for fi31 := 0; fi31 < 8; fi31++ {
            v5 = (v14 >= v14)
            v5 = (0.1 < (1.5 + (0.0 / (-2.25 * (((3.14159 + 1.5) * -2.25) + 0.1)))))
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi31
        }
        v12 = v12
        _ = fi21
    }
    var v32 bool = v3
    _ = len(v1)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v2, v4, v2)
    var v33 uint8 = 128
    v34 := GenS3{F0: v33}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
}
