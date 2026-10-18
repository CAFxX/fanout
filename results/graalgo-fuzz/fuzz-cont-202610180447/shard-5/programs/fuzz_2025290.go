package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 int; F1 string; F2 uint16 }
func genFn0(p0 string) uint64 {
    var v6 uint8 = 2
    var v7 int = 32768
    var v8 float32 = 1.5
    if (false && true) {
        fmt.Println(v6)
    }
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 448
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]uint32)
    v3 := make(map[int]int)
    var v4 int32 = 0
    var v5 uint16 = 32767
    var v9 int = 573
    v10 := GenS1{F0: 1, F1: (v5 ^ (v5 + 65535))}
    switch "foo bar" {
    case "":
        fmt.Println(v4, v5, v9)
    case "foo bar":
        v11 := []int{v9, v9}
        _ = v11
    case "a":
        for fi12 := 0; fi12 < 16; fi12++ {
            for fi13 := 0; fi13 < 5; fi13++ {
                switch "" {
                case "":
                    if false {
                        v9 = -630
                        v4 += (v4 >> ((18446744073709551615) & 63))
                        switch ((("\n\t" + "hello") + ("hello" + ("foo bar" + ("foo bar" + "foo bar")))) + "a") {
                        case "":
                            switch "foo bar" {
                            case "\n\t":
                                fmt.Println(v9, v4, v4)
                            case "a":
                                var v14 uint16 = v5
                                _ = v14
                            }
                        case "\n\t":
                            fmt.Println(v4, v4, v9)
                        case "foo bar":
                            var v15 float64 = 0.1
                            v16 := GenS1{F0: 42, F1: v5}
                            if true {
                                var v17 uint16 = v5
                                v5 = (v17 + v17)
                                var v18 bool = (false || (false || (!false)))
                                _ = v17
                                _ = v18
                            } else {
                                if false {
                                    v19 := GenS1{F0: 42, F1: (v5 % ((v5) | 1))}
                                    _ = v19
                                }
                                v20 := GenS1{F0: -705, F1: v5}
                                var v21 int16 = 32767
                                _ = v20
                                _ = v21
                            }
                            if false {
                                v22 := GenS1{F0: -448, F1: (v5 >> ((266) & 63))}
                                var v23 int64 = 1
                                v4 = (v4 * 42)
                                v24 := []int{32768, fi12}
                                v25 := GenS1{F0: (^(v23)), F1: v5}
                                _ = v22
                                _ = v23
                                _ = v24
                                _ = v25
                            }
                            var v26 string = ("\n\t" + "a")
                            _ = v15
                            _ = v16
                            _ = v26
                        }
                    }
                case "foo bar":
                    if (-2147483648 < (fi12 & v9)) {
                        v5 += (v5 - v5)
                    }
                    delete(v3, fi12)
                    var v27 uint16 = v5
                    _ = v27
                }
                _ = fi13
            }
            var v28 int64 = 0
            v28 += (v28 << ((42) & 63))
            _ = fi12
            _ = v28
        }
    default:
        v29 := GenS1{F0: 443, F1: 65535}
        _ = v29
    }
    v5 -= 1
    var v30 bool = (false && (42 != genFn0((("a" + ("hello" + (("hello" + "\n\t") + "foo bar"))) + "\n\t"))))
    v30 = v30
    v9 -= -752
    v30 = true
    v3[32767] = v9
    var v31 float64 = 3.14159
    v32 := []int{v9, v9, 4294967295}
    v33 := GenS2{F0: (v9 / ((-2147483648) | 1)), F1: "\n\t", F2: v5}
    var v34 bool = (136 < 1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v9
}
