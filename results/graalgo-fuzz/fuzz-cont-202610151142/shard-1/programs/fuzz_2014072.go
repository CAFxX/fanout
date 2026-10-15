package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 uint8 }
type GenS2 struct { F0 string; F1 uint; F2 bool }
type GenS3 struct { F0 string; F1 string }
type GenS4 struct { F0 string; F1 uint32; F2 string }
func genFn0() int32 {
    var v5 float32 = (-2.25 * 1e10)
    v5 -= float32(0.1)
    switch "foo bar" {
    case "a":
        if false {
            v5 = v5
            if true {
                v6 := []int{-1, 510}
                _ = v6
            }
        } else {
            v7 := GenS1{F0: 2, F1: 42, F2: 2}
            var v8 uint8 = 252
            _ = v7
            _ = v8
        }
    case "":
        if false {
            var v9 uint = 1
            var v10 int64 = -1
            _ = v10
            _ = v9
        }
    }
    switch -792 {
    case 2147483647:
        var v11 uint = 228
        _ = v11
    default:
        v5 = v5
    }
    if (1 == 1) {
        v5 = v5
        if (((false || ((!true) && true)) && false) || (42 == 0)) {
            if ((577 >= 486) && true) {
                if true {
                    v5 -= ((3.14159 - (v5 * float32((1e10 / (1.5 / 1e10))))) + 0.1)
                    v12 := []int{128}
                    v13 := GenS1{F0: 1, F1: -1, F2: 2}
                    _ = v12
                    _ = v13
                }
                v14 := GenS2{F0: (("a" + ("hello" + ("" + "\n\t"))) + ""), F1: 4294967295, F2: true}
                _ = v14
            }
            switch "foo bar" {
            case "hello":
                var v15 float64 = 1e10
                _ = v15
            default:
                var v16 string = ("\n\t" + (("foo bar" + "\n\t") + ""))
                _ = v16
            }
        } else {
            if ((((("hello" + "hello") + ("\n\t" + (("\n\t" + ("hello" + ("foo bar" + "hello"))) + ("" + ("a" + ("" + ("foo bar" + "\n\t"))))))) + ((("hello" + (("a" + ("" + "\n\t")) + ("foo bar" + "\n\t"))) + "hello") + "a")) + "\n\t") == ("\n\t" + "")) {
                fmt.Println(v5)
            } else {
                v5 += ((v5 * 1.5) + 3.14159)
            }
            v17 := GenS3{F0: "", F1: "foo bar"}
            _ = v17
        }
        v5 -= (3.14159 * ((3.14159 + 3.14159) / ((0.1 * (((1e10 * v5) + v5) - float32(((float64(((-2.25 / 0.1) - 0.0)) * -2.25) - 0.1)))) + (v5 - 3.14159))))
    }
    _ = v5
    return 2147483647
}

func main() {
    var v1 int8 = 22
    var v2 uint8 = 128
    var v3 bool = false
    var v4 bool = (1 <= 0)
    v18 := GenS3{F0: (("foo bar" + (("hello" + "") + ("hello" + ""))) + ("a" + "\n\t")), F1: "a"}
    var v19 int8 = (v1 - (v1 / ((v1) | 1)))
    var v20 uint32 = 2
    if v4 {
        v21 := GenS1{F0: v2, F1: -1, F2: (v2 + 231)}
        _ = v21
    }
    v22 := []int{-2147483648, -1}
    switch "" {
    case "\n\t":
        _ = len(v22)
    }
    var v23 bool = (!v3)
    for fi24 := 0; fi24 < 19; fi24++ {
        var v25 bool = true
        for fi26 := 0; fi26 < 11; fi26++ {
            v27 := GenS4{F0: (("a" + ("" + "foo bar")) + ("hello" + "a")), F1: v20, F2: "foo bar"}
            v28 := GenS1{F0: (v2 | 2), F1: 42, F2: v2}
            _ = fi26
            _ = v27
            _ = v28
        }
        var v29 uint64 = 42
        _ = fi24
        _ = v25
        _ = v29
    }
    var v30 uint64 = uint64(v2)
    if (-961 == 9223372036854775807) {
        v31 := GenS4{F0: ("a" + ("\n\t" + (("foo bar" + "\n\t") + (("\n\t" + ("\n\t" + ("" + "hello"))) + "a")))), F1: 1, F2: ""}
        _ = v31
    } else {
        for fi32 := 0; fi32 < 1; fi32++ {
            if len(v22) > 0 {
                v22[0] = (fi32 - fi32)
            }
            _ = fi32
        }
        var v33 int32 = 122
        v2 = (v2 ^ v2)
        _ = v33
    }
    v34 := GenS3{F0: "\n\t", F1: ("hello" + "")}
    fmt.Println(v20)
    var v35 string = "a"
    fmt.Println(v35, v1)
    var v36 int32 = 457
    var v37 int64 = -1
    var v38 uint = 578
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v3
    _ = v30
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
}
