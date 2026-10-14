package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int; F2 uint32 }
func genFn0(p0 uint64, p1 uint, p2 int8) int8 {
    var v6 bool = ((p2 | (p2 + 1)) != p2)
    v7 := []int{4294967295, 32767}
    v6 = true
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return p2
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{"hello", "hello", "hello", "foo bar"}
    v3 := []string{"a", ("hello" + (("a" + "\n\t") + "hello")), ((("" + "hello") + "foo bar") + (((("\n\t" + "foo bar") + ("foo bar" + "a")) + "\n\t") + (("\n\t" + "") + "")))}
    var v4 int32 = -2147483648
    var v5 uint32 = 65535
    var v8 int64 = 42
    v5 = v5
    if false {
        fmt.Println(v4, v4)
        switch 323 {
        case 888:
            fmt.Println(v4)
        case -1:
            if (!false) {
                v9 := GenS1{F0: "a", F1: 32767, F2: v5}
                fmt.Println(v4, v8, v5)
                var v10 int32 = (v4 - (v4 + (v4 + (v4 + (v4 & v4)))))
                _ = v10
                _ = v9
            }
        case 128:
            v11 := v1[2147483647]
            _ = v11
        default:
            v5 -= 4294967295
        }
        var v12 float64 = (-2.25 + -2.25)
        _ = v12
    }
    var v13 bool = (615 <= 2)
    var v14 float32 = 0.1
    var v15 int16 = 1
    switch 4294967295 {
    case -2147483648:
        v16 := GenS1{F0: "", F1: -1, F2: 4294967295}
        _ = v16
    case -9223372036854775808:
        for fi17 := 0; fi17 < 20; fi17++ {
            fmt.Println(v13, fi17, v5)
            if (fi17 == fi17) {
                for fi18 := 0; fi18 < 17; fi18++ {
                    var v19 uint64 = 374
                    var v20 bool = false
                    if ((18446744073709551615 < 65535) && false) {
                        if len(v3) > 0 {
                            v3[0] = ("foo bar" + "")
                        }
                    }
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                if (v13 && (!v13)) {
                    for fi21 := 0; fi21 < 19; fi21++ {
                        v5 = (v5 / ((v5) | 1))
                        v1[(v5 << ((0) & 63))] = ((42 <= 42) || true)
                        _ = fi21
                    }
                }
            }
            var v22 uint32 = v5
            _ = fi17
            _ = v22
        }
    }
    fmt.Println(v8, v15)
    var v23 int8 = genFn0(608, 545, -128)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
