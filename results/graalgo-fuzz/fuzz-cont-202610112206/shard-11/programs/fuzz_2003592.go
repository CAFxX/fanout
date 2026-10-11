package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint; F2 uint32 }
type GenS2 struct { F0 uint32; F1 int64; F2 int16 }
type GenS3 struct { F0 int32; F1 uint32; F2 uint }
type GenS4 struct { F0 string; F1 uint; F2 int32 }
func genFn0(p0 int8) uint64 {
    if (42 < 1) {
        if true {
            p0 -= -128
            fmt.Println(p0)
            v7 := []int{4294967295, 402, 4294967295}
            _ = v7
        } else {
            if true {
                switch ("" + ("hello" + "\n\t")) {
                case "hello":
                    p0 = -1
                case "\n\t":
                    v8 := GenS1{F0: 73, F1: 1, F2: 65535}
                    _ = v8
                default:
                    switch ("" + "hello") {
                    case "a":
                        var v9 int = 32767
                        _ = v9
                    case "foo bar":
                        if (!((false && (-950 > -9223372036854775808)) || ((65535 <= 126) && false))) {
                            var v10 int32 = -2147483648
                            v11 := []int{-331}
                            if true {
                                v12 := GenS1{F0: 2, F1: 229, F2: 2}
                                _ = v12
                            }
                            _ = v10
                            _ = v11
                        } else {
                            v13 := []int{32768, 9223372036854775807}
                            p0 = (^(p0))
                            _ = v13
                        }
                    case "":
                        for fi14 := 0; fi14 < 1; fi14++ {
                            switch "\n\t" {
                            case "\n\t":
                                if (false && true) {
                                    var v15 int16 = -1
                                    v16 := []int{(fi14 | (fi14 - (fi14 * -2147483648)))}
                                    _ = v15
                                    _ = v16
                                }
                            case "hello":
                                v17 := []int{fi14}
                                _ = v17
                            case "foo bar":
                                var v18 int32 = 42
                                _ = v18
                            }
                            _ = fi14
                        }
                    }
                }
                var v19 string = "hello"
                _ = v19
            } else {
                var v20 uint32 = 1
                _ = v20
            }
            p0 += p0
        }
    }
    switch -152 {
    case 32767:
        p0 -= (p0 + p0)
    case 2147483647:
        p0 = (-(p0))
    case 486:
        switch "" {
        case "":
            p0 -= (-(p0))
        case "foo bar":
            if (4294967295 >= 255) {
                if ((0.0 / 0.1) > (1.5 / 1e10)) {
                    v21 := []int{-2147483648}
                    _ = v21
                }
            }
        case "hello":
            if (42 >= 42) {
                v22 := []int{4294967295, 0, -1, 42}
                if true {
                    var v23 int8 = (p0 - (p0 + (p0 + 42)))
                    _ = v23
                } else {
                    var v24 uint64 = 4294967295
                    var v25 uint8 = 242
                    _ = v24
                    _ = v25
                }
                _ = v22
            } else {
                if false {
                    v26 := GenS1{F0: 65535, F1: 181, F2: 42}
                    var v27 string = "hello"
                    _ = v26
                    _ = v27
                }
                var v28 float64 = 0.1
                _ = v28
            }
        default:
            p0 = 29
        }
    }
    fmt.Println(p0)
    v29 := GenS1{F0: 65535, F1: 923, F2: 868}
    for fi30 := 0; fi30 < 1; fi30++ {
        v31 := GenS1{F0: 221, F1: 781, F2: 65535}
        v32 := GenS2{F0: 72, F1: -1, F2: 925}
        var v33 float64 = 1e10
        _ = fi30
        _ = v31
        _ = v32
        _ = v33
    }
    _ = p0
    _ = v29
    return 0
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint8 = 42
    v3 := []int{-238, 72, -9223372036854775808}
    v4 := []string{"foo bar", ""}
    v5 := []int{42, -1, 2147483647, 9223372036854775807}
    var v6 string = "hello"
    var v34 int8 = 100
    v35 := GenS3{F0: 2147483647, F1: 557, F2: 977}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v34 = (v34 | v34)
    delete(v1, 281)
    v34 = v34
    switch ("a" + v6) {
    case "a":
        for fi36 := 0; fi36 < 11; fi36++ {
            if (255 > 0) {
                v2 = 0
            }
            v37 := GenS4{F0: (v6 + v6), F1: 42, F2: -868}
            _ = fi36
            _ = v37
        }
    case "hello":
        v34 = (v34 >> ((42) & 63))
    }
    v6 = ("hello" + ((v6 + v6) + v6))
    fmt.Println(v6, v6, v34)
    var v38 float64 = 1e10
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
