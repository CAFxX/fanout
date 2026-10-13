package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 string }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 uint; F1 int; F2 int64 }
func genFn0(p0 uint16, p1 uint8) int8 {
    v6 := []int{138, 4294967295}
    for fi7 := 0; fi7 < 2; fi7++ {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        var v8 int = (fi7 | 32767)
        _ = fi7
        _ = v8
    }
    if false {
        fmt.Println(p1, p1)
    }
    switch 32767 {
    case 127:
        for fi9 := 0; fi9 < 18; fi9++ {
            var v10 int64 = -812
            p1 += (p1 * (^(p1)))
            _ = fi9
            _ = v10
        }
    default:
        if len(v6) > 0 {
            v6[0] = 0
        }
    }
    _ = p0
    _ = p1
    _ = v6
    return 100
}

func main() {
    v1 := make(map[string]int)
    var v2 uint32 = 909
    var v3 float32 = (1e10 + -2.25)
    var v4 bool = false
    v5 := []string{"a", "hello", ("\n\t" + "foo bar")}
    v11 := GenS1{F0: 0, F1: 43, F2: "hello"}
    switch -187 {
    case 1:
        if v4 {
            var v12 int32 = 1
            _ = v12
        } else {
            v1["hello"] = 785
            var v13 int16 = 0
            _ = v13
        }
    case -776:
        v1["\n\t"] = 127
    case 127:
        delete(v1, "")
    }
    if ("foo bar" == ("\n\t" + ("a" + (("foo bar" + "foo bar") + ((((("foo bar" + ("" + ("" + "\n\t"))) + "foo bar") + "") + "hello") + "\n\t"))))) {
        for fi14 := 0; fi14 < 19; fi14++ {
            v15 := GenS1{F0: 0, F1: 242, F2: "hello"}
            _ = fi14
            _ = v15
        }
        switch 32768 {
        case -454:
            var v16 float32 = 1.5
            _ = v16
        }
        fmt.Println(v2)
    }
    v2 += 42
    fmt.Println(v2)
    _ = len(v5)
    v4 = v4
    v2 = (v2 & (v2 >> ((668) & 63)))
    v4 = (!(42 == 192))
    _ = len(v5)
    if (!v4) {
        for fi17 := 0; fi17 < 10; fi17++ {
            if (!false) {
                fmt.Println(v4, v2, v4)
                fmt.Println(v2)
            }
            if false {
                v18 := GenS1{F0: 18446744073709551615, F1: 205, F2: ("hello" + "foo bar")}
                v4 = v4
                var v19 float32 = v3
                _ = v18
                _ = v19
            }
            if v4 {
                v20 := GenS1{F0: 1, F1: 127, F2: ("\n\t" + "\n\t")}
                var v21 bool = false
                if ((v2 << ((1) & 63)) <= 2) {
                    switch "\n\t" {
                    case "":
                        v22 := GenS2{F0: 0, F1: ""}
                        _ = v22
                    }
                    v2 -= 42
                    var v23 int8 = 127
                    _ = v23
                }
                _ = v20
                _ = v21
            } else {
                v24 := v1[("hello" + "\n\t")]
                _ = v24
            }
            _ = fi17
        }
        if v4 {
            fmt.Println(v3)
            v25 := GenS1{F0: 1, F1: 128, F2: "\n\t"}
            v26 := GenS2{F0: 1, F1: "a"}
            _ = v25
            _ = v26
        } else {
            v27 := GenS3{F0: 42, F1: 42, F2: 771}
            _ = v27
        }
        fmt.Println(v2)
    } else {
        var v28 int32 = 42
        _ = v28
    }
    v3 = 1e10
    v2 += 0
    fmt.Println(v4, v2, v4)
    for fi29 := 0; fi29 < 7; fi29++ {
        var v30 uint32 = 65535
        v31 := GenS2{F0: 2, F1: ("foo bar" + "a")}
        var v32 int32 = 1
        _ = fi29
        _ = v30
        _ = v31
        _ = v32
    }
    fmt.Println(v4, v3)
    v3 = -2.25
    var v33 string = ("foo bar" + ("" + "\n\t"))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
}
