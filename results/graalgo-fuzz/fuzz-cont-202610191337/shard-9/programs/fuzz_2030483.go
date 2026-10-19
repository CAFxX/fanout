package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 uint32; F1 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 bool; F1 int8 }
func genFn0(p0 int64, p1 string, p2 bool) int8 {
    p2 = p2
    var v6 float64 = 0.1
    switch p1 {
    case "":
        var v7 uint32 = 42
        _ = v7
    case "a":
        var v8 uint8 = 127
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 100
}

func main() {
    var v1 int64 = 0
    var v2 int8 = 100
    v3 := []string{(("" + ("\n\t" + ((("hello" + "hello") + "") + ("foo bar" + "")))) + "a"), ("foo bar" + "a"), ("\n\t" + "a")}
    var v4 int16 = -130
    var v5 int32 = 42
    v9 := GenS1{F0: v5, F1: 4294967295}
    v4 = v4
    if (!((("a" > "") && true) && (!(42 == 1)))) {
        v3 = append(v3, "foo bar")
    } else {
        v10 := GenS2{F0: 2, F1: "hello"}
        v2 -= v2
        _ = v10
    }
    v1 = v1
    v2 += 93
    var v11 uint64 = 2
    var v12 uint = 65535
    v12 = (v12 & 2)
    v4 = 1000
    if (false && (((("hello" + "foo bar") + "a") < (((("" + "") + "\n\t") + "hello") + ("hello" + ("\n\t" + "\n\t")))) || true)) {
        switch 9223372036854775807 {
        case 2147483647:
            _ = len(v3)
        case 127:
            fmt.Println(v12)
        case -2147483648:
            var v13 float64 = (0.0 * 0.1)
            _ = v13
        default:
            if len(v3) > 0 {
                v3[0] = (("\n\t" + "") + "")
            }
        }
        var v14 string = ("hello" + ("a" + ((("a" + "a") + "hello") + ((("hello" + "foo bar") + "hello") + (("a" + (("\n\t" + "a") + "\n\t")) + ("\n\t" + "foo bar"))))))
        _ = v14
    }
    v12 = 0
    v15 := GenS3{F0: 2}
    if false {
        for fi16 := 0; fi16 < 8; fi16++ {
            for fi17 := 0; fi17 < 18; fi17++ {
                fmt.Println(v12, v1, fi17)
                _ = fi17
            }
            for fi18 := 0; fi18 < 10; fi18++ {
                fmt.Println(fi18, fi18)
                _ = fi18
            }
            if (!false) {
                v19 := GenS4{F0: genFn0(v1, ("hello" + "a"), (!((float64(0.1) != 1.5) || false)))}
                v20 := GenS1{F0: (v5 << ((v12) & 63)), F1: (v12 - v12)}
                _ = v19
                _ = v20
            }
            _ = fi16
        }
    } else {
        v1 -= (v1 - 9223372036854775807)
        for fi21 := 0; fi21 < 5; fi21++ {
            v3 = append(v3, "foo bar")
            _ = fi21
        }
    }
    v22 := GenS5{F0: false, F1: v2}
    var v23 uint16 = 311
    fmt.Println(v11, v23, v5)
    fmt.Println(v12, v1, v2)
    v23 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v15
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
