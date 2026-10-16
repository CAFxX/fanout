package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 int32 }
type GenS2 struct { F0 int32 }
func genFn0(p0 int) int8 {
    if false {
        for fi6 := 0; fi6 < 5; fi6++ {
            v7 := GenS1{F0: 32767, F1: 188, F2: -1}
            _ = fi6
            _ = v7
        }
    } else {
        var v8 int32 = -2147483648
        var v9 uint8 = 0
        _ = v8
        _ = v9
    }
    var v10 float32 = (((0.0 * 1.5) - 0.0) * (1.5 + (0.0 + 3.14159)))
    _ = p0
    _ = v10
    return 84
}

func main() {
    var v1 uint8 = 94
    v2 := []string{(("hello" + "") + (("foo bar" + ((("\n\t" + ("a" + "foo bar")) + "hello") + "hello")) + "foo bar")), ("\n\t" + ("\n\t" + "")), "a", "a"}
    var v3 int = 9223372036854775807
    var v4 string = ("foo bar" + "\n\t")
    var v5 string = ("a" + v4)
    var v11 bool = ((^(v1)) < 255)
    var v12 int32 = -1
    v5 = v5
    _ = len(v2)
    switch v4 {
    case "\n\t":
        var v13 string = "a"
        _ = v13
    case "hello":
        v11 = v11
    case "a":
        v11 = v11
    default:
        if len(v2) > 0 {
            v2[0] = "\n\t"
        }
    }
    if true {
        for fi14 := 0; fi14 < 19; fi14++ {
            if len(v2) > 0 {
                v2[0] = ("a" + (("\n\t" + v5) + v5))
            }
            v3 = (fi14 % ((fi14) | 1))
            v3 = (v3 ^ (v3 + (-(fi14))))
            _ = fi14
        }
        v15 := GenS1{F0: 65535, F1: v1, F2: (v12 ^ v12)}
        _ = v15
    } else {
        var v16 int8 = 127
        _ = len(v2)
        _ = v16
    }
    v4 = v4
    for fi17 := 0; fi17 < 2; fi17++ {
        fmt.Println(v5, fi17)
        var v18 int = 32767
        _ = fi17
        _ = v18
    }
    var v19 int = v3
    v1 -= 0
    var v20 int16 = 1000
    if false {
        v21 := GenS1{F0: 2, F1: 0, F2: int32(v20)}
        _ = v21
    } else {
        if (v1 > v1) {
            switch "foo bar" {
            case "\n\t":
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            var v22 uint = 0
            v23 := GenS1{F0: 0, F1: (v1 << ((v22) & 63)), F2: -2147483648}
            _ = v22
            _ = v23
        }
        var v24 float64 = ((1.5 - 1.5) * 3.14159)
        for fi25 := 0; fi25 < 4; fi25++ {
            switch -304 {
            case 497:
                switch (v3 / ((9223372036854775807) | 1)) {
                case 42:
                    v26 := GenS2{F0: (v12 | -956)}
                    _ = v26
                case 32767:
                    v27 := GenS1{F0: 0, F1: (^(v1)), F2: (v12 * 0)}
                    _ = v27
                }
            case -1:
                var v28 int8 = 127
                _ = v28
            case -9223372036854775808:
                _ = len(v2)
            default:
                if len(v2) > 0 {
                    v2[0] = "hello"
                }
            }
            var v29 bool = true
            if true {
                var v30 int32 = v12
                var v31 uint = 740
                _ = v30
                _ = v31
            } else {
                if (-2.25 < 0.0) {
                    var v32 int8 = 42
                    var v33 int8 = v32
                    v1 = (v1 | 0)
                    _ = v32
                    _ = v33
                } else {
                    v11 = (634 >= 452)
                    _ = len(v2)
                }
                if ((v20 - v20) < v20) {
                    v4 = v5
                    var v34 string = v4
                    _ = v34
                }
            }
            _ = fi25
            _ = v29
        }
        _ = v24
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
