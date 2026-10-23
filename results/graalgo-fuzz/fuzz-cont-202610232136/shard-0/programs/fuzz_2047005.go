package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint64; F1 int64; F2 uint8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 string; F1 int32 }
type GenS5 struct { F0 int16; F1 string; F2 uint64 }
func genFn0(p0 uint16, p1 string, p2 uint16) int {
    if false {
        p1 = "foo bar"
        var v4 bool = true
        _ = v4
    }
    var v5 float32 = float32(((((0.0 - float64(1e10)) - 1.5) - -2.25) - 0.1))
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return -2147483648
}

func main() {
    var v1 int8 = 67
    v2 := make(map[string]int)
    var v3 uint32 = 679
    v1 = -1
    if true {
        v6 := GenS1{F0: -685}
        _ = v6
    } else {
        v7 := GenS2{F0: 4294967295, F1: 9223372036854775807, F2: 160}
        var v8 string = ((("hello" + ("foo bar" + ((("foo bar" + "\n\t") + "hello") + ""))) + "foo bar") + "foo bar")
        _ = v7
        _ = v8
    }
    var v9 uint64 = 4294967295
    v10 := GenS1{F0: -307}
    if false {
        for fi11 := 0; fi11 < 13; fi11++ {
            var v12 string = (("\n\t" + "a") + "a")
            _ = fi11
            _ = v12
        }
        if (588 != 42) {
            switch 9223372036854775807 {
            case 4294967295:
                fmt.Println(v1, v3, v1)
            }
        }
        var v13 bool = (!(-1 == -9223372036854775808))
        _ = v13
    } else {
        v14 := []int{1, int(v1), genFn0(42, ((("a" + ("foo bar" + "a")) + (((("a" + "foo bar") + (((("" + "foo bar") + ("" + (("a" + ("foo bar" + ((("foo bar" + "hello") + "hello") + (((("foo bar" + ((("foo bar" + "hello") + (("a" + "a") + ("a" + ("foo bar" + "")))) + "")) + "foo bar") + "hello") + (("a" + "hello") + "foo bar"))))) + "a"))) + "") + "hello")) + "") + "a")) + ""), 626), genFn0(2, ("" + "a"), 1)}
        v15 := GenS3{F0: 42}
        switch "foo bar" {
        case "a":
            if false {
                var v16 uint = 123
                v17 := GenS4{F0: "foo bar", F1: -2147483648}
                _ = v16
                _ = v17
            }
        default:
            v3 = (v3 ^ (v3 | 2))
        }
        _ = v14
        _ = v15
    }
    var v18 string = "hello"
    var v19 uint32 = (v3 * v3)
    for fi20 := 0; fi20 < 1; fi20++ {
        if true {
            var v21 int64 = 0
            for fi22 := 0; fi22 < 9; fi22++ {
                if false {
                    v18 = "\n\t"
                } else {
                    var v23 uint32 = v3
                    if (false || (783 < 592)) {
                        v19 += (v3 / ((v3) | 1))
                        v24 := []int{(fi20 % ((fi20) | 1)), 42}
                        _ = v24
                    }
                    v25 := v2[(v18 + (v18 + v18))]
                    _ = v23
                    _ = v25
                }
                v26 := GenS5{F0: 0, F1: v18, F2: (v9 % (((v9 << ((255) & 63))) | 1))}
                v3 = v19
                v1 -= (v1 - (v1 >> ((65535) & 63)))
                _ = fi22
                _ = v26
            }
            _ = v21
        }
        _ = fi20
    }
    v27 := GenS2{F0: v9, F1: -9223372036854775808, F2: 128}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v18
    _ = v19
    _ = v2
    _ = v27
    _ = v3
    _ = v9
}
