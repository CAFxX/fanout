package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32; F1 bool }
func genFn0() string {
    if true {
        var v6 bool = (("hello" + "foo bar") != "\n\t")
        for fi7 := 0; fi7 < 11; fi7++ {
            switch (fi7 * fi7) {
            case 243:
                if v6 {
                    fmt.Println(v6)
                    fmt.Println(fi7, v6)
                    if (42 < 2) {
                        for fi8 := 0; fi8 < 6; fi8++ {
                            var v9 bool = true
                            if true {
                                v10 := []int{(fi7 ^ (fi8 % (((fi8 & (fi7 % (((fi8 ^ fi7)) | 1)))) | 1))), (^(fi7))}
                                v11 := GenS1{F0: -9223372036854775808}
                                v6 = v6
                                _ = v10
                                _ = v11
                            } else {
                                v9 = (!false)
                            }
                            var v12 string = "hello"
                            _ = fi8
                            _ = v12
                            _ = v9
                        }
                    } else {
                        v13 := []int{fi7, (fi7 % (((fi7 - fi7)) | 1)), fi7}
                        _ = v13
                    }
                }
            }
            _ = fi7
        }
        _ = v6
    } else {
        var v14 uint16 = 648
        v15 := []int{-9223372036854775808, 32768, 9223372036854775807, 32767}
        var v16 string = (("\n\t" + (("foo bar" + "hello") + ("" + (("hello" + "") + "foo bar")))) + "")
        _ = v14
        _ = v15
        _ = v16
    }
    v17 := GenS1{F0: 127}
    var v18 int8 = 0
    fmt.Println(v18)
    _ = v17
    _ = v18
    return "foo bar"
}
func genFn1(p0 int64) int8 {
    v19 := []int{9223372036854775807, -9223372036854775808, 42, -2147483648}
    var v20 float64 = (((1.5 + 1.5) - 0.0) / 1e10)
    _ = p0
    _ = v19
    _ = v20
    return -128
}

func main() {
    var v1 uint64 = 537
    v2 := []int{1, -2147483648}
    v3 := make(map[int]uint32)
    var v4 int = 0
    var v5 int32 = 0
    fmt.Println(v4, v4)
    v4 -= 127
    switch v4 {
    case 2147483647:
        var v21 int = (v4 & (v4 & -1))
        _ = v21
    case 1:
        for fi22 := 0; fi22 < 4; fi22++ {
            var v23 bool = true
            var v24 uint = 18446744073709551615
            v4 -= 2147483647
            _ = fi22
            _ = v23
            _ = v24
        }
    }
    v5 -= (^(v5))
    var v25 int16 = -32768
    var v26 uint64 = (v1 ^ (v1 << ((232) & 63)))
    if true {
        _ = len(v3)
        if true {
            for fi27 := 0; fi27 < 11; fi27++ {
                var v28 int32 = (-(v5))
                var v29 uint16 = 2
                v5 += 42
                _ = fi27
                _ = v28
                _ = v29
            }
        }
    } else {
        var v30 int32 = (v5 | v5)
        _ = v30
    }
    fmt.Println(v4)
    v4 = v4
    v5 -= -477
    v5 += v5
    switch "a" {
    case "a":
        var v31 uint64 = 0
        _ = v31
    case "foo bar":
        v32 := GenS1{F0: v4}
        _ = v32
    case "\n\t":
        v33 := GenS1{F0: v4}
        _ = v33
    default:
        for fi34 := 0; fi34 < 16; fi34++ {
            var v35 bool = (1 >= 1)
            _ = fi34
            _ = v35
        }
    }
    if (!(false && (!true))) {
        v25 = v25
        var v36 uint16 = 255
        v4 -= -1
        _ = v36
    } else {
        for fi37 := 0; fi37 < 15; fi37++ {
            delete(v3, (fi37 % ((698) | 1)))
            _ = fi37
        }
        switch (-(v4)) {
        case 42:
            v25 = (v25 / ((v25) | 1))
        case 9223372036854775807:
            var v38 uint16 = 32767
            _ = v38
        case -747:
            switch v4 {
            case 42:
                v39 := v3[(v4 << ((253) & 63))]
                _ = v39
            case 834:
                if (true || false) {
                    v40 := GenS1{F0: (-(v4))}
                    v41 := GenS2{F0: v5, F1: true}
                    _ = v40
                    _ = v41
                }
            case 128:
                v42 := GenS1{F0: 1}
                _ = v42
            }
        default:
            fmt.Println(v25, v5, v26)
        }
        v5 = v5
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
