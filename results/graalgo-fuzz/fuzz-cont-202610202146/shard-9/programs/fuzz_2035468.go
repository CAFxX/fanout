package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32; F1 int8; F2 int64 }
func genFn0() uint {
    for fi7 := 0; fi7 < 11; fi7++ {
        v8 := GenS1{F0: (^(fi7))}
        v9 := []int{(^(fi7)), (-(fi7)), fi7}
        var v10 uint64 = 42
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    v11 := GenS1{F0: -2147483648}
    var v12 float32 = 0.1
    _ = v11
    _ = v12
    return 65535
}
func genFn1(p0 uint, p1 bool, p2 int32) uint32 {
    v13 := []int{-1, 0, 42}
    var v14 bool = (!(-2.25 != 3.14159))
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v14
    return 0
}

func main() {
    v1 := []string{"foo bar"}
    v2 := make(map[string]bool)
    var v3 int32 = -2147483648
    v4 := []string{("hello" + "a")}
    var v5 bool = (-128 >= 73)
    v6 := []uint32{}
    v5 = (true || (1 != 4294967295))
    switch 366 {
    case -380:
        v15 := GenS1{F0: -2147483648}
        _ = v15
    case -734:
        v5 = v5
    case -141:
        var v16 uint8 = 14
        _ = v16
    }
    v5 = (float64((3.14159 / 1.0)) <= 0.1)
    v5 = (!v5)
    switch -16 {
    case 4294967295:
        _ = len(v2)
    case 887:
        var v17 int64 = 1
        _ = v17
    }
    v5 = (1 == 65535)
    v5 = false
    var v18 uint = 100
    for fi19 := 0; fi19 < 9; fi19++ {
        v20 := GenS1{F0: fi19}
        for fi21 := 0; fi21 < 17; fi21++ {
            if (42 < -9223372036854775808) {
                if len(v1) > 0 {
                    v1[0] = "a"
                }
                var v22 string = ""
                var v23 int16 = 42
                _ = v22
                _ = v23
            }
            switch 164 {
            case 783:
                if false {
                    v5 = false
                    v3 += v3
                } else {
                    var v24 int64 = 1
                    var v25 int32 = v3
                    _ = v24
                    _ = v25
                }
                var v26 int64 = -9223372036854775808
                _ = v26
            default:
                v27 := GenS2{F0: 2, F1: 42, F2: 42}
                _ = v27
            }
            v3 = v3
            _ = fi21
        }
        _ = fi19
        _ = v20
    }
    if false {
        var v28 int32 = 248
        v28 += v3
        _ = v28
    } else {
        var v29 string = ("\n\t" + "foo bar")
        if true {
            var v30 float32 = 1e10
            _ = v30
        }
        fmt.Println(v18, v3)
        _ = v29
    }
    var v31 uint = (^(v18))
    v18 = v18
    v31 += (v18 % ((genFn0()) | 1))
    if (127 >= 100) {
        for fi32 := 0; fi32 < 13; fi32++ {
            for fi33 := 0; fi33 < 9; fi33++ {
                switch "\n\t" {
                case "foo bar":
                    _ = len(v2)
                default:
                    fmt.Println(v3, v31, v18)
                }
                var v34 uint8 = 2
                _ = fi33
                _ = v34
            }
            v35 := v2[("a" + "foo bar")]
            switch (fi32 - (fi32 ^ fi32)) {
            case 837:
                v36 := GenS1{F0: fi32}
                _ = v36
            default:
                fmt.Println(v35, v18)
            }
            _ = fi32
            _ = v35
        }
    } else {
        for fi37 := 0; fi37 < 17; fi37++ {
            v38 := GenS1{F0: fi37}
            v3 -= v3
            delete(v2, "hello")
            _ = fi37
            _ = v38
        }
    }
    for fi39 := 0; fi39 < 4; fi39++ {
        v2["hello"] = (((((("\n\t" + ("foo bar" + "foo bar")) + ("foo bar" + "hello")) + ("a" + "\n\t")) > "") || true) && (v5 || true))
        _ = fi39
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
