package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint }
type GenS2 struct { F0 int64 }
func genFn0(p0 string, p1 uint8, p2 uint16) bool {
    fmt.Println(p1, p0)
    v5 := []int{9223372036854775807, 2147483647, 2147483647, 857}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return false
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 int32 = 0
    var v3 int = 9223372036854775807
    var v4 string = "\n\t"
    fmt.Println(v2, v3)
    var v6 float64 = (1.5 + float64(-2.25))
    v7 := GenS1{F0: 268, F1: 65535}
    var v8 int = v3
    fmt.Println(v4, v3)
    v4 = ("\n\t" + ("hello" + ("hello" + "")))
    v2 = (v2 & v2)
    delete(v1, 2147483647)
    switch ((v4 + ((v4 + "hello") + (("foo bar" + v4) + ("\n\t" + v4)))) + (v4 + v4)) {
    case "":
        fmt.Println(v3, v4, v6)
    case "\n\t":
        var v9 uint8 = 255
        _ = v9
    default:
        var v10 bool = true
        _ = v10
    }
    switch "" {
    case "":
        var v11 string = (v4 + v4)
        _ = v11
    default:
        v12 := GenS1{F0: 32768, F1: 0}
        _ = v12
    }
    for fi13 := 0; fi13 < 1; fi13++ {
        if ((!true) && false) {
            v14 := GenS1{F0: v3, F1: 28}
            _ = v14
        } else {
            var v15 uint16 = 484
            if (!false) {
                var v16 uint16 = 919
                if false {
                    switch v8 {
                    case 700:
                        switch v4 {
                        case "\n\t":
                            v3 -= (fi13 % ((88) | 1))
                        case "a":
                            for fi17 := 0; fi17 < 1; fi17++ {
                                v18 := GenS1{F0: (^(fi13)), F1: 255}
                                var v19 string = "hello"
                                _ = fi17
                                _ = v18
                                _ = v19
                            }
                        default:
                            switch "hello" {
                            case "\n\t":
                                var v20 float64 = v6
                                _ = v20
                            case "foo bar":
                                for fi21 := 0; fi21 < 16; fi21++ {
                                    var v22 float32 = 0.0
                                    _ = fi21
                                    _ = v22
                                }
                            case "hello":
                                var v23 uint32 = 42
                                _ = v23
                            }
                        }
                    case 4294967295:
                        if false {
                            v3 -= v3
                        } else {
                            v15 = (v16 & v16)
                            if (((!true) || (2 != 65535)) || (true || (0 < 649))) {
                                var v24 int16 = 298
                                _ = v24
                            }
                        }
                    default:
                        v15 += 65535
                    }
                    switch v4 {
                    case "":
                        v1[352] = ((100 > 42) && false)
                    default:
                        var v25 uint = 2
                        _ = v25
                    }
                }
                if (false || (0 <= 127)) {
                    var v26 uint64 = 18446744073709551615
                    _ = v26
                }
                _ = v16
            }
            _ = v15
        }
        _ = fi13
    }
    if true {
        switch (v4 + v4) {
        case "":
            delete(v1, 2147483647)
        case "\n\t":
            v27 := GenS2{F0: 455}
            _ = v27
        case "foo bar":
            for fi28 := 0; fi28 < 15; fi28++ {
                var v29 uint64 = 309
                if true {
                    v30 := v1[42]
                    _ = v30
                } else {
                    v31 := []int{(^(fi28)), v3}
                    _ = v31
                }
                v4 = "hello"
                _ = fi28
                _ = v29
            }
        }
    } else {
        v3 = v3
        v32 := []int{}
        if (1e10 > float64((0.1 / 1e10))) {
            v4 = v4
            if false {
                v32 = append(v32, (-(v8)))
            } else {
                v1[2] = false
            }
            v33 := GenS1{F0: (v8 % ((1) | 1)), F1: 255}
            _ = v33
        }
        _ = v32
    }
    fmt.Println(v6)
    var v34 uint64 = 862
    fmt.Println(v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v6
    _ = v7
    _ = v8
}
