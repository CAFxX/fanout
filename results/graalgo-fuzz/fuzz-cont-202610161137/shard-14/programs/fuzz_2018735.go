package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 string; F1 uint; F2 int }
type GenS3 struct { F0 int }
func genFn0() int64 {
    switch -383 {
    case 0:
        v6 := []int{4294967295, 32768, 358}
        _ = v6
    case -608:
        var v7 int32 = 252
        _ = v7
    }
    var v8 uint32 = 155
    v8 = 2147483648
    _ = v8
    return -1
}

func main() {
    v1 := []uint32{1, 640, 42}
    v2 := []string{}
    var v3 int64 = 82
    var v4 uint32 = 2147483647
    var v5 uint32 = (v4 % (((v4 << ((37) & 63))) | 1))
    v5 += (v5 << ((4294967295) & 63))
    if true {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v9 int16 = -32768
        _ = v9
    } else {
        var v10 int = -1
        if (!true) {
            v10 = 127
        } else {
            if len(v2) > 0 {
                v2[0] = ""
            }
            if (((((0.1 / 3.14159) * (3.14159 * 1e10)) <= 1.5) && false) && (2 != 827)) {
                v10 = -390
                v5 = v4
            }
        }
        _ = v10
    }
    for fi11 := 0; fi11 < 17; fi11++ {
        if (42 != 18446744073709551615) {
            fmt.Println(fi11)
            fmt.Println(v3, fi11)
        }
        v4 += v4
        if false {
            var v12 int8 = 100
            v5 -= 1
            _ = v12
        }
        _ = fi11
    }
    if len(v2) > 0 {
        v2[0] = (("a" + "") + "foo bar")
    }
    var v13 float32 = float32(0.0)
    var v14 uint32 = 0
    switch (((((((("a" + "hello") + "foo bar") + ("hello" + "a")) + (((("foo bar" + (("foo bar" + "\n\t") + ("hello" + ("foo bar" + "foo bar")))) + (("foo bar" + "\n\t") + ((("foo bar" + ("\n\t" + ("" + "foo bar"))) + "hello") + ("foo bar" + "hello")))) + "a") + "foo bar")) + ("foo bar" + "hello")) + "") + "foo bar") + "a") {
    case "":
        v2 = append(v2, "\n\t")
    }
    switch 9223372036854775807 {
    case 32768:
        v15 := GenS1{F0: 88, F1: ""}
        _ = v15
    default:
        v13 = v13
    }
    fmt.Println(v5, v3, v4)
    if (1 != 65) {
        for fi16 := 0; fi16 < 1; fi16++ {
            v1 = append(v1, (^(v14)))
            v1 = append(v1, 2147483648)
            _ = fi16
        }
        for fi17 := 0; fi17 < 14; fi17++ {
            v2 = append(v2, "\n\t")
            v14 -= (v5 - v4)
            _ = len(v1)
            _ = fi17
        }
    } else {
        switch (("hello" + "a") + ("a" + ((("a" + ((("hello" + ("\n\t" + (("hello" + ("\n\t" + ((("foo bar" + "\n\t") + ("hello" + "foo bar")) + "hello"))) + "a"))) + "hello") + "\n\t")) + ("" + "a")) + "foo bar"))) {
        case "hello":
            if ((v3 != v3) && false) {
                v4 += (v4 >> ((18446744073709551615) & 63))
            } else {
                v13 = v13
                v18 := GenS1{F0: 2, F1: "a"}
                _ = v18
            }
        case "":
            v19 := GenS2{F0: "", F1: 95, F2: 42}
            _ = v19
        case "a":
            v3 += v3
        default:
            var v20 int32 = -1
            _ = v20
        }
        fmt.Println(v14)
    }
    v21 := GenS3{F0: 971}
    v13 = 0.1
    fmt.Println(v13)
    _ = len(v2)
    fmt.Println(v5)
    for fi22 := 0; fi22 < 7; fi22++ {
        v3 += v3
        if (false && (100 != 117)) {
            v13 = v13
        } else {
            var v23 uint16 = 42
            v1 = append(v1, 163)
            v14 = v4
            _ = v23
        }
        if true {
            if len(v1) > 0 {
                v1[0] = (^(v14))
            }
            _ = len(v1)
        }
        _ = fi22
    }
    if true {
        var v24 uint8 = 42
        _ = v24
    } else {
        v13 -= v13
        v2 = append(v2, "a")
    }
    if (v4 <= v4) {
        var v25 uint = 42
        v3 -= (v3 / ((42) | 1))
        _ = v25
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
