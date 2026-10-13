package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 int32 }
type GenS2 struct { F0 int8; F1 uint16; F2 int8 }
func genFn0() string {
    var v5 bool = (4294967295 >= 42)
    var v6 int32 = 0
    for fi7 := 0; fi7 < 12; fi7++ {
        v8 := GenS1{F0: 585, F1: ("hello" + (("foo bar" + ("" + ("a" + "foo bar"))) + "foo bar")), F2: (v6 / ((v6) | 1))}
        _ = fi7
        _ = v8
    }
    var v9 uint32 = 65535
    var v10 int64 = int64(v9)
    _ = v10
    _ = v5
    _ = v6
    _ = v9
    return "\n\t"
}

func main() {
    v1 := []int{32768}
    v2 := []string{("\n\t" + ""), ""}
    var v3 uint64 = 0
    var v4 uint32 = 0
    var v11 float64 = -2.25
    switch "hello" {
    case "foo bar":
        var v12 float32 = 1e10
        _ = v12
    default:
        v13 := GenS1{F0: v4, F1: "", F2: 464}
        _ = v13
    }
    var v14 uint8 = 127
    switch 123 {
    case 1:
        if (!false) {
            if ((65535 >= 255) || (!false)) {
                if ((!(1.5 >= 1e10)) && ((1 != 128) && (!(2 > (v3 - v3))))) {
                    var v15 bool = true
                    _ = v15
                }
                v4 = 135
                for fi16 := 0; fi16 < 9; fi16++ {
                    switch (fi16 + (-(fi16))) {
                    case 127:
                        if (-1 == -1) {
                            v17 := GenS1{F0: uint32(v14), F1: "\n\t", F2: -2147483648}
                            _ = v17
                        } else {
                            var v18 int8 = 1
                            _ = v18
                        }
                    case 175:
                        var v19 uint = 882
                        _ = v19
                    }
                    var v20 int8 = -1
                    _ = fi16
                    _ = v20
                }
            } else {
                _ = len(v2)
                var v21 int = -427
                v3 += 0
                _ = v21
            }
            v22 := GenS1{F0: 564, F1: "a", F2: -1}
            v23 := GenS2{F0: -20, F1: 152, F2: -22}
            _ = v22
            _ = v23
        } else {
            var v24 uint = 2
            _ = v24
        }
    case 42:
        _ = len(v2)
    }
    var v25 int16 = 42
    if (true || (!true)) {
        var v26 string = "hello"
        if ((42 < 1) || true) {
            v25 = (v25 + (v25 ^ v25))
            v1 = append(v1, 2147483647)
        } else {
            var v27 int64 = 42
            fmt.Println(v26, v3, v27)
            var v28 bool = false
            _ = v27
            _ = v28
        }
        v26 = "\n\t"
        _ = v26
    }
    var v29 uint16 = 32767
    for fi30 := 0; fi30 < 18; fi30++ {
        v1 = append(v1, -2147483648)
        _ = fi30
    }
    var v31 int16 = (^(v25))
    for fi32 := 0; fi32 < 13; fi32++ {
        fmt.Println(v11)
        _ = fi32
    }
    for fi33 := 0; fi33 < 14; fi33++ {
        for fi34 := 0; fi34 < 8; fi34++ {
            var v35 int8 = -4
            _ = fi34
            _ = v35
        }
        v3 -= v3
        _ = fi33
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v14
    _ = v2
    _ = v25
    _ = v29
    _ = v3
    _ = v31
    _ = v4
}
