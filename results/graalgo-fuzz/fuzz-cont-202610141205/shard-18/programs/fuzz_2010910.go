package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
func genFn0(p0 uint) uint32 {
    var v7 uint32 = 42
    p0 = (p0 + (p0 & 0))
    switch -9223372036854775808 {
    case 0:
        v8 := []int{}
        _ = v8
    case 42:
        for fi9 := 0; fi9 < 7; fi9++ {
            p0 -= (p0 | (p0 >> ((255) & 63)))
            switch "\n\t" {
            case "a":
                v10 := GenS1{F0: "foo bar"}
                _ = v10
            default:
                v7 = v7
            }
            if ((v7 / (((v7 + (v7 | (v7 | 2)))) | 1)) != v7) {
                var v11 uint16 = 255
                _ = v11
            }
            _ = fi9
        }
    default:
        fmt.Println(p0, p0)
    }
    p0 += p0
    _ = p0
    _ = v7
    return 2147483647
}
func genFn1(p0 int16) bool {
    fmt.Println(p0)
    var v12 uint8 = 83
    v12 = v12
    switch (("foo bar" + "\n\t") + "a") {
    case "":
        v13 := []int{181}
        _ = v13
    default:
        var v14 uint8 = v12
        _ = v14
    }
    _ = p0
    _ = v12
    return true
}

func main() {
    var v1 uint8 = 42
    var v2 uint64 = 1
    var v3 int64 = 9223372036854775807
    var v4 int16 = 1000
    var v5 string = "\n\t"
    var v6 int16 = (-(v4))
    v15 := []int{1, 128, -1}
    if (42 != 4294967295) {
        v15 = append(v15, 128)
        for fi16 := 0; fi16 < 5; fi16++ {
            for fi17 := 0; fi17 < 20; fi17++ {
                var v18 int = (fi16 ^ (fi16 ^ (fi17 | int(v3))))
                v1 = v1
                _ = fi17
                _ = v18
            }
            _ = fi16
        }
        v6 = 324
    }
    for fi19 := 0; fi19 < 16; fi19++ {
        if len(v15) > 0 {
            v15[0] = -594
        }
        _ = fi19
    }
    if len(v15) >= 2 {
        v15 = v15[1:len(v15)]
    }
    var v20 uint16 = 2
    switch 9223372036854775807 {
    case 378:
        if false {
            var v21 float32 = (0.1 * 1.5)
            if false {
                v3 -= -494
            }
            _ = v21
        } else {
            var v22 string = "a"
            _ = v22
        }
    default:
        v6 -= (v4 - int16(v2))
    }
    if (127 > -128) {
        if (v3 <= 0) {
            if len(v15) > 0 {
                v15[0] = 9223372036854775807
            }
        }
    } else {
        _ = len(v15)
        for fi23 := 0; fi23 < 11; fi23++ {
            if false {
                fmt.Println(v6, v5)
                for fi24 := 0; fi24 < 12; fi24++ {
                    fmt.Println(v2)
                    _ = fi24
                }
                for fi25 := 0; fi25 < 11; fi25++ {
                    fmt.Println(v20)
                    _ = fi25
                }
            }
            _ = fi23
        }
    }
    if (!true) {
        for fi26 := 0; fi26 < 1; fi26++ {
            var v27 int32 = -2147483648
            v6 += v6
            _ = fi26
            _ = v27
        }
        v15 = append(v15, 761)
        v3 += v3
    }
    v20 -= (v20 & (v20 % ((v20) | 1)))
    var v28 int64 = v3
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v20
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
