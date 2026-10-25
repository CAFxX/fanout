package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
func genFn0(p0 uint32) uint8 {
    switch "hello" {
    case "foo bar":
        v6 := []int{128, -2147483648}
        _ = v6
    default:
        for fi7 := 0; fi7 < 16; fi7++ {
            fmt.Println(p0, p0)
            p0 -= (^(p0))
            var v8 string = ("a" + "hello")
            _ = fi7
            _ = v8
        }
    }
    var v9 uint32 = p0
    for fi10 := 0; fi10 < 6; fi10++ {
        p0 -= v9
        fmt.Println(v9, fi10)
        _ = fi10
    }
    _ = p0
    _ = v9
    return 56
}
func genFn1(p0 int16, p1 uint32) uint {
    fmt.Println(p1, p1)
    var v11 uint16 = 42
    switch 2147483647 {
    case 1:
        var v12 int8 = 127
        _ = v12
    default:
        v13 := []int{0, -1}
        _ = v13
    }
    v11 -= (v11 - v11)
    fmt.Println(p0, v11, p0)
    _ = p0
    _ = p1
    _ = v11
    return 363
}

func main() {
    v1 := make(map[string]int)
    var v2 int8 = -1
    v3 := []int{-1, -1, -2147483648}
    var v4 bool = false
    var v5 float64 = 3.14159
    _ = len(v3)
    switch "\n\t" {
    case "hello":
        v2 += (v2 % ((v2) | 1))
    case "":
        for fi14 := 0; fi14 < 20; fi14++ {
            fmt.Println(v2)
            _ = fi14
        }
    case "foo bar":
        if v4 {
            var v15 int8 = v2
            var v16 uint8 = 2
            _ = v15
            _ = v16
        }
    default:
        v17 := GenS1{F0: 2147483648}
        _ = v17
    }
    v5 = float64(1.5)
    switch -1 {
    case 248:
        if v4 {
            if len(v3) > 0 {
                v3[0] = 32767
            }
            fmt.Println(v4)
            v2 = -1
        }
    default:
        for fi18 := 0; fi18 < 3; fi18++ {
            v19 := GenS1{F0: 0}
            for fi20 := 0; fi20 < 19; fi20++ {
                fmt.Println(fi20, v5, v5)
                if false {
                    var v21 int8 = 100
                    v4 = (((v4 && v4) && v4) || true)
                    v22 := v1[""]
                    _ = v21
                    _ = v22
                } else {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                v23 := v1["hello"]
                fmt.Println(v2, v23)
                if v4 {
                    var v24 bool = v4
                    fmt.Println(v5, v4)
                    _ = v24
                }
                _ = fi20
                _ = v23
            }
            var v25 int = 289
            _ = fi18
            _ = v19
            _ = v25
        }
    }
    _ = len(v1)
    v5 = v5
    fmt.Println(v2, v2, v5)
    v26 := GenS1{F0: 2147483648}
    switch ("" + "") {
    case "foo bar":
        for fi27 := 0; fi27 < 20; fi27++ {
            var v28 uint = 1
            v2 = 0
            fmt.Println(v4)
            _ = fi27
            _ = v28
        }
    case "a":
        v29 := GenS1{F0: 4294967295}
        _ = v29
    case "":
        switch "\n\t" {
        case "hello":
            var v30 float32 = 0.0
            _ = v30
        case "\n\t":
            for fi31 := 0; fi31 < 15; fi31++ {
                switch (fi31 << ((genFn1(0, 1)) & 63)) {
                case 127:
                    var v32 int8 = v2
                    _ = v32
                case 0:
                    v5 = 0.1
                default:
                    var v33 uint32 = 2
                    _ = v33
                }
                _ = fi31
            }
        case "a":
            if true {
                var v34 int64 = -9223372036854775808
                _ = v34
            }
        }
    default:
        fmt.Println(v4, v5, v4)
    }
    var v35 uint64 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v35
    _ = v4
    _ = v5
}
