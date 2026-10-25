package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 int64) uint64 {
    fmt.Println(p0)
    if false {
        if (0 <= 1000) {
            p0 += 9223372036854775807
            p0 -= p0
            p0 = (p0 / ((-9223372036854775808) | 1))
        }
        for fi7 := 0; fi7 < 18; fi7++ {
            for fi8 := 0; fi8 < 1; fi8++ {
                v9 := GenS1{F0: "\n\t"}
                p0 += (^(p0))
                _ = fi8
                _ = v9
            }
            p0 += -9223372036854775808
            v10 := []int{-371}
            _ = fi7
            _ = v10
        }
    } else {
        for fi11 := 0; fi11 < 6; fi11++ {
            var v12 int64 = (p0 & 595)
            var v13 float32 = 0.0
            _ = fi11
            _ = v12
            _ = v13
        }
        if false {
            fmt.Println(p0)
        } else {
            v14 := GenS1{F0: "foo bar"}
            var v15 uint32 = 65535
            v16 := []int{127, 127}
            _ = v14
            _ = v15
            _ = v16
        }
    }
    _ = p0
    return 42
}
func genFn1(p0 bool) uint {
    var v17 int32 = -796
    switch ("" + "hello") {
    case "a":
        var v18 uint32 = 0
        _ = v18
    case "":
        var v19 int8 = -128
        _ = v19
    case "foo bar":
        v17 -= v17
    }
    var v20 string = "\n\t"
    _ = p0
    _ = v17
    _ = v20
    return 65535
}

func main() {
    var v1 float64 = 0.1
    var v2 bool = true
    var v3 int32 = 0
    var v4 float64 = ((float64(float32((v1 + 0.1))) - 1.5) + v1)
    v5 := []uint32{2147483648}
    var v6 uint8 = 120
    v21 := GenS1{F0: "foo bar"}
    if (42 >= 2147483648) {
        if len(v5) > 0 {
            v5[0] = 65535
        }
        fmt.Println(v4)
        for fi22 := 0; fi22 < 1; fi22++ {
            v23 := GenS1{F0: "hello"}
            _ = fi22
            _ = v23
        }
    } else {
        fmt.Println(v1, v3, v1)
        if (18446744073709551615 >= 0) {
            var v24 uint32 = 0
            fmt.Println(v4, v24, v1)
            _ = v24
        } else {
            var v25 string = "hello"
            _ = v25
        }
        for fi26 := 0; fi26 < 5; fi26++ {
            for fi27 := 0; fi27 < 14; fi27++ {
                v2 = v2
                fmt.Println(v3, v2)
                v1 = (1e10 / v1)
                _ = fi27
            }
            _ = fi26
        }
    }
    v5 = append(v5, 628)
    switch ("a" + ("foo bar" + "\n\t")) {
    case "":
        if false {
            for fi28 := 0; fi28 < 16; fi28++ {
                v29 := GenS1{F0: ("hello" + "")}
                _ = fi28
                _ = v29
            }
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            v2 = (((32767 != -32768) || v2) || v2)
        }
    }
    var v30 uint64 = 42
    v3 += (-(v3))
    for fi31 := 0; fi31 < 6; fi31++ {
        v3 = 549
        v6 = v6
        _ = fi31
    }
    switch "" {
    case "hello":
        if v2 {
            if len(v5) > 0 {
                v5[0] = 300
            }
        }
    case "":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    case "\n\t":
        v5 = append(v5, 237)
    default:
        if len(v5) > 0 {
            v5[0] = uint32(v6)
        }
    }
    for fi32 := 0; fi32 < 4; fi32++ {
        var v33 uint32 = 817
        v2 = ((("" + ("" + "foo bar")) + "hello") != "")
        if (!(!true)) {
            for fi34 := 0; fi34 < 1; fi34++ {
                v6 += (v6 % ((150) | 1))
                _ = fi34
            }
            var v35 bool = (302 == genFn0(9223372036854775807))
            _ = v35
        }
        _ = fi32
        _ = v33
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
