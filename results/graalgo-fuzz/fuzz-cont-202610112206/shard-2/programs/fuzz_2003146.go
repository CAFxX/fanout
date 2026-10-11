package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 uint }
func genFn0(p0 string) uint16 {
    var v6 uint = 736
    var v7 int16 = -32768
    var v8 uint32 = 2
    v6 = 269
    for fi9 := 0; fi9 < 12; fi9++ {
        for fi10 := 0; fi10 < 17; fi10++ {
            if (!((3.14159 <= 0.1) || (false && false))) {
                if ((!(false || (((0.0 - 0.0) / 0.1) >= 3.14159))) || false) {
                    p0 = ("foo bar" + p0)
                    v11 := GenS1{F0: 690, F1: fi10, F2: (v6 << (((v6 + 18446744073709551615)) & 63))}
                    switch ((p0 + ("foo bar" + ("foo bar" + p0))) + "") {
                    case "":
                        fmt.Println(fi10, v8)
                    case "\n\t":
                        v7 = (v7 & 0)
                    case "hello":
                        var v12 uint = (^(v6))
                        _ = v12
                    default:
                        if ((42 >= v8) && (true || ((^(v8)) > (v8 | (v8 | (v8 / (((^(v8))) | 1))))))) {
                            var v13 float64 = 1e10
                            _ = v13
                        } else {
                            var v14 float64 = 3.14159
                            fmt.Println(v6)
                            v15 := []int{fi10, 521, (fi10 | 1)}
                            _ = v14
                            _ = v15
                        }
                    }
                    _ = v11
                }
            }
            var v16 int16 = 335
            _ = fi10
            _ = v16
        }
        fmt.Println(v6, fi9)
        for fi17 := 0; fi17 < 13; fi17++ {
            v7 = v7
            if true {
                var v18 float32 = 1.5
                switch (^(fi17)) {
                case 128:
                    var v19 bool = (!((true && (fi17 <= fi9)) && (4294967295 == v6)))
                    _ = v19
                }
                p0 = ""
                _ = v18
            }
            v20 := GenS1{F0: 143, F1: int(v6), F2: 18446744073709551615}
            _ = fi17
            _ = v20
        }
        _ = fi9
    }
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 51
}

func main() {
    var v1 uint8 = 217
    v2 := []string{"", (("a" + "hello") + "foo bar"), "a"}
    v3 := []int{-9223372036854775808, 2147483647}
    var v4 int16 = -1
    v5 := make(map[string]string)
    fmt.Println(v1, v1)
    var v21 uint = 1
    var v22 uint16 = 32767
    switch (("hello" + "foo bar") + (("hello" + ("hello" + ("foo bar" + ("foo bar" + "foo bar")))) + "")) {
    case "hello":
        v21 -= (v21 - 4294967295)
    case "":
        v21 += (^(v21))
    case "\n\t":
        v23 := v5[("\n\t" + ("" + "\n\t"))]
        _ = v23
    }
    var v24 uint64 = 0
    for fi25 := 0; fi25 < 14; fi25++ {
        fmt.Println(fi25)
        v26 := GenS1{F0: 42, F1: fi25, F2: v21}
        _ = fi25
        _ = v26
    }
    var v27 bool = false
    v27 = v27
    v5[("a" + (("hello" + "\n\t") + "foo bar"))] = (("hello" + "") + "hello")
    for fi28 := 0; fi28 < 1; fi28++ {
        _ = len(v3)
        fmt.Println(v21, fi28, v4)
        _ = fi28
    }
    var v29 uint32 = 1
    var v30 int8 = 1
    fmt.Println(v24, v22)
    var v31 float32 = 0.1
    v4 = (-(v4))
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v24
    _ = v27
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
}
