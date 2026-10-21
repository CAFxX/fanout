package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 uint8; F2 string }
func genFn0(p0 uint, p1 uint32) int32 {
    v6 := []int{}
    var v7 int8 = -128
    p1 = p1
    var v8 int64 = 0
    v6 = append(v6, 4294967295)
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    _ = v8
    return -865
}
func genFn1(p0 uint16, p1 int16) int8 {
    p0 -= 2
    v9 := GenS1{F0: 1}
    switch -9223372036854775808 {
    case 32767:
        fmt.Println(p1, p0)
    }
    fmt.Println(p0, p0)
    _ = p0
    _ = p1
    _ = v9
    return 127
}

func main() {
    var v1 uint8 = 2
    var v2 int64 = -802
    v3 := []int{1, 32767, 0}
    var v4 int8 = 42
    v5 := []uint32{}
    var v10 float32 = (float32((float64(3.14159) - 0.0)) / (-2.25 + 3.14159))
    for fi11 := 0; fi11 < 19; fi11++ {
        v2 = v2
        _ = fi11
    }
    var v12 int8 = genFn1(2, -1)
    v5 = append(v5, 42)
    var v13 uint32 = 1
    v2 -= v2
    if len(v5) > 0 {
        v5[0] = (v13 % ((v13) | 1))
    }
    v13 = v13
    fmt.Println(v2)
    for fi14 := 0; fi14 < 4; fi14++ {
        v12 = (v4 - -107)
        if true {
            var v15 uint64 = 370
            _ = v15
        }
        v1 = (v1 / ((v1) | 1))
        _ = fi14
    }
    var v16 uint8 = 2
    switch 32768 {
    case 127:
        var v17 uint = 4294967295
        _ = v17
    }
    switch ("" + "") {
    case "foo bar":
        for fi18 := 0; fi18 < 13; fi18++ {
            var v19 uint16 = 42
            v20 := GenS1{F0: (^(v13))}
            _ = fi18
            _ = v19
            _ = v20
        }
    }
    v21 := GenS2{F0: 372, F1: (v1 << ((0) & 63)), F2: ("hello" + "a")}
    var v22 int = 1
    if false {
        switch ("foo bar" + ("hello" + "\n\t")) {
        case "a":
            v23 := GenS1{F0: (^(v13))}
            _ = v23
        case "\n\t":
            var v24 int = -869
            _ = v24
        }
    } else {
        v25 := GenS2{F0: -344, F1: 2, F2: "foo bar"}
        switch (("foo bar" + "foo bar") + "hello") {
        case "":
            if (2 == uint16(v16)) {
                if true {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
            }
        }
        _ = v25
    }
    var v26 int = v22
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
