package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 int64 }
func genFn0(p0 uint, p1 int) int16 {
    var v5 string = ("\n\t" + "foo bar")
    var v6 uint32 = 1
    if true {
        switch v5 {
        case "\n\t":
            v7 := []int{1, (p1 | 1), (p1 + p1)}
            _ = v7
        case "hello":
            v8 := []int{(p1 - (p1 >> ((18446744073709551615) & 63))), 2147483647, p1}
            _ = v8
        default:
            if (18446744073709551615 == 103) {
                if true {
                    if (p0 > p0) {
                        for fi9 := 0; fi9 < 2; fi9++ {
                            p0 += (p0 ^ (p0 | (^(p0))))
                            for fi10 := 0; fi10 < 19; fi10++ {
                                var v11 uint32 = 2147483648
                                switch -546 {
                                case 2147483647:
                                    var v12 int32 = 0
                                    _ = v12
                                case -2147483648:
                                    v13 := []int{fi9, (p1 & (p1 % (((fi9 & (p1 % ((127) | 1)))) | 1)))}
                                    _ = v13
                                default:
                                    v14 := []int{p1, (fi10 % ((int(p0)) | 1)), (fi9 % ((fi10) | 1)), 4294967295}
                                    _ = v14
                                }
                                v15 := GenS1{F0: (fi10 << ((42) & 63)), F1: 9223372036854775807, F2: 9223372036854775807}
                                p1 -= (fi9 ^ (fi9 / ((fi9) | 1)))
                                v5 = "a"
                                _ = fi10
                                _ = v11
                                _ = v15
                            }
                            var v16 float32 = ((1e10 + -2.25) - 0.0)
                            _ = fi9
                            _ = v16
                        }
                        for fi17 := 0; fi17 < 9; fi17++ {
                            switch 468 {
                            case -684:
                                var v18 uint32 = (v6 * (v6 >> (((p0 >> (((p0 / ((p0) | 1))) & 63))) & 63)))
                                _ = v18
                            default:
                                v19 := []int{(fi17 % ((-1000) | 1)), (fi17 | p1), 32768}
                                _ = v19
                            }
                            _ = fi17
                        }
                    }
                    fmt.Println(v6, v6, p1)
                }
            }
        }
        if true {
            var v20 uint16 = 446
            _ = v20
        } else {
            var v21 int8 = 0
            v22 := []int{p1, (-(p1)), (^(p1))}
            if true {
                var v23 uint8 = 255
                fmt.Println(v6)
                _ = v23
            } else {
                v6 -= 42
                var v24 uint32 = (v6 << ((p0) & 63))
                var v25 bool = (1000 != 0)
                _ = v24
                _ = v25
            }
            _ = v21
            _ = v22
        }
    }
    for fi26 := 0; fi26 < 17; fi26++ {
        v27 := []int{p1, p1, (fi26 % (((fi26 & p1)) | 1))}
        p0 = (p0 - (^(p0)))
        _ = fi26
        _ = v27
    }
    fmt.Println(v5, v5)
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 1
}

func main() {
    v1 := []int{0, 5, -231}
    var v2 uint = 2
    var v3 int = 9223372036854775807
    var v4 bool = (true || (-1 == 100))
    v1 = append(v1, v3)
    v4 = (-538 <= 42)
    v2 -= v2
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi28 := 0; fi28 < 7; fi28++ {
        var v29 int64 = -1
        fmt.Println(v29, v29)
        _ = fi28
        _ = v29
    }
    switch ((("\n\t" + "a") + (("a" + "a") + (("" + ("foo bar" + (("hello" + "\n\t") + "hello"))) + "hello"))) + (("a" + "") + ("a" + "foo bar"))) {
    case "a":
        v4 = false
    default:
        v1 = append(v1, (^(v3)))
    }
    v2 -= (v2 + 42)
    v4 = false
    var v30 string = ""
    switch ((v30 + "hello") + (v30 + v30)) {
    case "a":
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
    case "\n\t":
        _ = len(v1)
    case "":
        fmt.Println(v2)
    default:
        if (701 < 158) {
            var v31 uint = v2
            _ = v31
        }
    }
    var v32 int32 = -1
    if v4 {
        v33 := GenS1{F0: v3, F1: 0, F2: 0}
        v30 = (v30 + v30)
        _ = v33
    } else {
        if len(v1) > 0 {
            v1[0] = v3
        }
    }
    var v34 int32 = v32
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v32
    _ = v34
    _ = v4
}
