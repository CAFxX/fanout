package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 int8; F2 uint32 }
type GenS3 struct { F0 int64; F1 int }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int8; F1 int64 }
func genFn0() uint8 {
    var v6 int8 = 42
    for fi7 := 0; fi7 < 3; fi7++ {
        v8 := []int{}
        v8 = append(v8, fi7)
        _ = fi7
        _ = v8
    }
    _ = v6
    return 42
}

func main() {
    var v1 int = 42
    v2 := make(map[uint32]uint32)
    var v3 int8 = 0
    v4 := make(map[string]uint32)
    var v5 int32 = -955
    v9 := []int{(v1 * (-(v1)))}
    v1 = (v1 * v1)
    if true {
        fmt.Println(v5, v5)
        v1 += v1
        if (false || ((!(v3 <= v3)) && true)) {
            if (4294967295 == 125) {
                if false {
                    fmt.Println(v5)
                    v1 = v1
                    v10 := GenS1{F0: 65535}
                    _ = v10
                } else {
                    v9 = append(v9, 128)
                    if false {
                        v5 -= (v5 % ((1) | 1))
                        fmt.Println(v3, v1, v3)
                    }
                    switch "hello" {
                    case "hello":
                        for fi11 := 0; fi11 < 12; fi11++ {
                            v12 := GenS1{F0: 255}
                            _ = fi11
                            _ = v12
                        }
                    case "":
                        var v13 int = v1
                        _ = v13
                    case "foo bar":
                        for fi14 := 0; fi14 < 20; fi14++ {
                            v5 += 905
                            var v15 uint32 = 832
                            _ = fi14
                            _ = v15
                        }
                    default:
                        switch ("a" + ((("foo bar" + "") + ("\n\t" + ("" + "hello"))) + "foo bar")) {
                        case "a":
                            if false {
                                var v16 string = "hello"
                                _ = v16
                            }
                        case "hello":
                            v17 := GenS2{F0: ("\n\t" + ""), F1: (v3 << ((1) & 63)), F2: 4294967295}
                            _ = v17
                        case "":
                            v5 = (v5 ^ v5)
                        default:
                            if len(v9) > 0 {
                                v9[0] = v1
                            }
                        }
                    }
                }
            }
            v18 := GenS3{F0: -437, F1: v1}
            _ = v18
        } else {
            for fi19 := 0; fi19 < 8; fi19++ {
                v20 := GenS4{F0: v3}
                _ = fi19
                _ = v20
            }
        }
    }
    var v21 string = "hello"
    if true {
        var v22 int8 = 1
        _ = v22
    } else {
        fmt.Println(v3, v21)
        switch (-(v1)) {
        case 128:
            v1 = 130
        case 424:
            var v23 string = (v21 + "hello")
            _ = v23
        case 42:
            v21 = ((v21 + "hello") + "a")
        default:
            fmt.Println(v5, v3)
        }
        fmt.Println(v3)
    }
    v5 -= 639
    for fi24 := 0; fi24 < 1; fi24++ {
        fmt.Println(v5, v21, v3)
        _ = fi24
    }
    switch v1 {
    case -9223372036854775808:
        var v25 int16 = -32768
        _ = v25
    case -80:
        delete(v4, "a")
    default:
        v26 := GenS5{F0: 127, F1: 9223372036854775807}
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
