package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 string }
type GenS2 struct { F0 uint32; F1 int64 }
type GenS3 struct { F0 uint8; F1 int16; F2 string }
func genFn0(p0 int32, p1 uint) int8 {
    for fi7 := 0; fi7 < 9; fi7++ {
        if (-415 == 0) {
            for fi8 := 0; fi8 < 20; fi8++ {
                v9 := []int{fi8}
                _ = fi8
                _ = v9
            }
        }
        if true {
            fmt.Println(p0, fi7)
            v10 := []int{4294967295, (fi7 << ((p1) & 63))}
            v11 := GenS1{F0: 4294967295, F1: "", F2: "hello"}
            _ = v10
            _ = v11
        } else {
            switch 32767 {
            case 0:
                p0 += (p0 * -1)
            case 4294967295:
                v12 := []int{199}
                _ = v12
            case -2147483648:
                var v13 uint = (^(p1))
                _ = v13
            }
            p1 += 0
        }
        p0 += (p0 % (((p0 << ((p1) & 63))) | 1))
        _ = fi7
    }
    v14 := []int{32767, 128, 257}
    p1 = p1
    _ = p0
    _ = p1
    _ = v14
    return 127
}

func main() {
    var v1 uint = 4294967295
    v2 := []int{32767, 9223372036854775807, 1, 11}
    v3 := []uint32{65535, 1, 42}
    var v4 int64 = -9223372036854775808
    v5 := []string{("\n\t" + "\n\t"), ""}
    v6 := make(map[string]string)
    if false {
        if len(v5) > 0 {
            v5[0] = "hello"
        }
    }
    v15 := GenS2{F0: 42, F1: (v4 & -9223372036854775808)}
    switch 2147483647 {
    case 127:
        fmt.Println(v1, v4)
    default:
        v6[""] = (("hello" + "\n\t") + "")
    }
    v16 := GenS3{F0: 2, F1: 42, F2: (("\n\t" + ("\n\t" + ((("a" + ("a" + "")) + "a") + "\n\t"))) + "foo bar")}
    fmt.Println(v4)
    v17 := GenS2{F0: 4294967295, F1: (-(v4))}
    if len(v3) > 0 {
        v3[0] = 791
    }
    if (42 >= 42) {
        switch ("\n\t" + (("" + "") + "hello")) {
        case "foo bar":
            delete(v6, "foo bar")
        case "\n\t":
            fmt.Println(v4, v4)
        default:
            fmt.Println(v1, v1)
        }
    } else {
        v1 = (^(v1))
        var v18 uint = v1
        var v19 int64 = v4
        _ = v18
        _ = v19
    }
    v4 = (-(v4))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
