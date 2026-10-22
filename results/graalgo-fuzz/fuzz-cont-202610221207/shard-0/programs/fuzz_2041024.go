package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 string; F1 uint16 }
func genFn0(p0 bool) bool {
    p0 = (42 < 4294967295)
    v7 := GenS1{F0: 4294967295, F1: 2}
    var v8 int64 = 0
    v8 = v8
    var v9 uint8 = 1
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return p0
}

func main() {
    v1 := make(map[string]bool)
    var v2 uint32 = 666
    var v3 bool = (float32((0.0 - -2.25)) >= (0.1 * 1.5))
    var v4 int16 = -1
    v5 := make(map[int]int)
    v6 := make(map[uint32]bool)
    var v10 int = -620
    if true {
        v5[0] = v10
    }
    fmt.Println(v10, v2)
    v4 = 1000
    if (("" > "\n\t") || ((0.1 + 3.14159) >= -2.25)) {
        fmt.Println(v4, v3)
    } else {
        var v11 bool = true
        if v3 {
            v12 := GenS1{F0: 457, F1: 0}
            v4 = (v4 ^ (v4 % (((^(v4))) | 1)))
            v11 = v3
            _ = v12
        }
        _ = v11
    }
    var v13 uint8 = 255
    switch ("foo bar" + "") {
    case "":
        if v3 {
            fmt.Println(v3, v10)
            v13 -= 187
            for fi14 := 0; fi14 < 10; fi14++ {
                var v15 uint32 = v2
                v3 = false
                _ = fi14
                _ = v15
            }
        }
    default:
        switch v10 {
        case -2147483648:
            var v16 int16 = (^(v4))
            _ = v16
        default:
            _ = len(v5)
        }
    }
    switch (-(v10)) {
    case -9223372036854775808:
        v10 = (^(v10))
    case 42:
        var v17 int = 32768
        _ = v17
    case 9223372036854775807:
        switch "hello" {
        case "\n\t":
            v3 = ("" != ("hello" + (("hello" + "hello") + "\n\t")))
        case "":
            for fi18 := 0; fi18 < 16; fi18++ {
                v3 = (-2.25 != 3.14159)
                var v19 bool = false
                for fi20 := 0; fi20 < 4; fi20++ {
                    var v21 uint = 65535
                    var v22 int32 = 42
                    _ = fi20
                    _ = v21
                    _ = v22
                }
                _ = fi18
                _ = v19
            }
        default:
            for fi23 := 0; fi23 < 1; fi23++ {
                v24 := v5[4294967295]
                _ = fi23
                _ = v24
            }
        }
    }
    var v25 float64 = 1e10
    switch (^(v10)) {
    case 4294967295:
        if ((v13 < 134) || ((v2 * v2) > 65535)) {
            v13 = 1
        }
    default:
        for fi26 := 0; fi26 < 13; fi26++ {
            v27 := GenS2{F0: "foo bar", F1: 32767}
            if ("foo bar" < (("a" + (("foo bar" + "hello") + ("hello" + (("a" + "") + "")))) + "foo bar")) {
                for fi28 := 0; fi28 < 10; fi28++ {
                    if v3 {
                        v29 := GenS2{F0: "hello", F1: 42}
                        _ = v29
                    }
                    fmt.Println(fi28)
                    _ = fi28
                }
            }
            _ = fi26
            _ = v27
        }
    }
    for fi30 := 0; fi30 < 18; fi30++ {
        var v31 uint16 = 2
        if (0 > 65535) {
            v32 := GenS1{F0: 4294967295, F1: v31}
            _ = v32
        }
        _ = fi30
        _ = v31
    }
    if (true || (18446744073709551615 != 65535)) {
        var v33 int64 = 66
        if (!(v13 == 2)) {
            if (v13 < 1) {
                if v3 {
                    v34 := GenS1{F0: 243, F1: 65535}
                    var v35 int = v10
                    delete(v5, v10)
                    _ = v34
                    _ = v35
                }
            }
            switch ("hello" + "\n\t") {
            case "\n\t":
                v10 = (-(v10))
            case "":
                v33 -= (-(v33))
            case "hello":
                fmt.Println(v2)
            }
        }
        _ = v33
    }
    v3 = (!v3)
    switch "" {
    case "\n\t":
        var v36 float64 = ((1.5 + v25) / 3.14159)
        _ = v36
    case "hello":
        var v37 float32 = ((0.0 + 3.14159) * (-2.25 / (3.14159 - 0.0)))
        _ = v37
    default:
        for fi38 := 0; fi38 < 11; fi38++ {
            v39 := []int{127, 2147483647}
            fmt.Println(v2, v13)
            _ = fi38
            _ = v39
        }
    }
    fmt.Println(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
