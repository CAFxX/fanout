package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
func genFn0(p0 int16, p1 uint16) uint32 {
    var v6 uint64 = 4294967295
    fmt.Println(p1)
    v7 := []int{4294967295, 1, 2147483647}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 725
}

func main() {
    var v1 int = 4294967295
    v2 := []string{"foo bar"}
    var v3 int8 = -59
    v4 := make(map[uint32]int)
    var v5 bool = false
    if len(v2) > 0 {
        v2[0] = "a"
    }
    if len(v2) > 0 {
        v2[0] = ((("\n\t" + ("foo bar" + "")) + "\n\t") + ("" + "foo bar"))
    }
    var v8 uint16 = 2
    var v9 uint = 29
    v1 = 859
    if (v8 != (^(v8))) {
        var v10 string = "hello"
        _ = v10
    }
    var v11 bool = true
    v11 = (-9223372036854775808 > (v1 | 9223372036854775807))
    switch (v1 % ((v1) | 1)) {
    case 2147483647:
        v12 := GenS1{F0: v9, F1: ""}
        _ = v12
    default:
        v9 = (v9 / ((v9) | 1))
    }
    if false {
        v8 -= v8
        v4[760] = 128
        delete(v4, 42)
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if ((v8 & v8) == v8) {
            v8 = v8
        } else {
            fmt.Println(v11, v5, v9)
            var v13 int = v1
            if (!(genFn0(1, (v8 | (v8 * v8))) != 65535)) {
                var v14 int32 = -977
                if v11 {
                    fmt.Println(v5, v3)
                    var v15 string = ("foo bar" + "a")
                    _ = v15
                }
                _ = v14
            }
            _ = v13
        }
        switch "hello" {
        case "hello":
            if (!v11) {
                if v11 {
                    for fi16 := 0; fi16 < 5; fi16++ {
                        var v17 bool = ((v3 & -1) <= 0)
                        if (v9 >= v9) {
                            var v18 bool = true
                            _ = v18
                        }
                        _ = fi16
                        _ = v17
                    }
                    v3 = (v3 | (^(v3)))
                }
            } else {
                var v19 uint8 = 0
                _ = v19
            }
        case "\n\t":
            switch v1 {
            case 560:
                fmt.Println(v1, v11)
            case 42:
                for fi20 := 0; fi20 < 5; fi20++ {
                    var v21 int8 = (v3 / ((-1) | 1))
                    var v22 int = (^(fi20))
                    _ = fi20
                    _ = v21
                    _ = v22
                }
            default:
                if (v9 < 4294967295) {
                    v23 := v4[2]
                    v1 += 1
                    _ = v23
                } else {
                    var v24 float32 = (((3.14159 - (3.14159 / 1e10)) - float32(((1.5 / 0.1) / -2.25))) + 1e10)
                    _ = v24
                }
            }
        default:
            var v25 int32 = -2147483648
            _ = v25
        }
    }
    if v5 {
        v26 := GenS1{F0: (v9 - v9), F1: ("a" + "foo bar")}
        _ = v26
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
