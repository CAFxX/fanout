package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 uint32 }
type GenS2 struct { F0 string; F1 uint64 }

func main() {
    v1 := []string{("hello" + "\n\t")}
    var v2 uint32 = 4294967295
    v3 := []uint32{2147483648, v2, v2, v2}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if (("foo bar" < "") && false) {
        v4 := GenS1{F0: "", F1: "\n\t", F2: (v2 % (((^(v2))) | 1))}
        fmt.Println(v2)
        _ = v4
    } else {
        fmt.Println(v2)
        if true {
            fmt.Println(v2)
            if len(v3) > 0 {
                v3[0] = 1
            }
        }
    }
    if true {
        for fi5 := 0; fi5 < 1; fi5++ {
            var v6 int = (^(fi5))
            _ = fi5
            _ = v6
        }
        if false {
            if (true && ((-9223372036854775808 < 839) || true)) {
                var v7 int64 = -425
                _ = v7
            } else {
                v8 := GenS1{F0: "\n\t", F1: "\n\t", F2: v2}
                fmt.Println(v2)
                var v9 int32 = 2147483647
                _ = v8
                _ = v9
            }
        } else {
            fmt.Println(v2)
            var v10 bool = true
            _ = v10
        }
    }
    v11 := GenS1{F0: "", F1: "foo bar", F2: (^(v2))}
    if false {
        v2 = 1
        if false {
            v2 -= (v2 / ((v2) | 1))
        }
        v2 += 4294967295
    }
    var v12 float32 = 1.5
    for fi13 := 0; fi13 < 10; fi13++ {
        for fi14 := 0; fi14 < 3; fi14++ {
            v1 = append(v1, "\n\t")
            _ = fi14
        }
        _ = len(v3)
        _ = fi13
    }
    switch ("a" + "foo bar") {
    case "hello":
        v12 = v12
    case "foo bar":
        var v15 float64 = (1.5 - (float64((((float32(3.14159) / v12) * float32(0.1)) * 1e10)) * 0.0))
        _ = v15
    default:
        var v16 int16 = -948
        _ = v16
    }
    fmt.Println(v12, v12)
    var v17 int64 = 194
    v17 -= v17
    if len(v1) > 0 {
        v1[0] = "\n\t"
    }
    var v18 int16 = -32768
    fmt.Println(v2, v17, v18)
    if true {
        var v19 int64 = v17
        for fi20 := 0; fi20 < 9; fi20++ {
            v21 := GenS1{F0: (("\n\t" + "a") + (("a" + ("hello" + ((("\n\t" + ("" + "foo bar")) + "\n\t") + ("" + ("foo bar" + ""))))) + "")), F1: ("\n\t" + ("a" + ("hello" + "\n\t"))), F2: v2}
            for fi22 := 0; fi22 < 17; fi22++ {
                var v23 uint64 = 0
                fmt.Println(v23, v23, fi20)
                if (!(!false)) {
                    v24 := GenS2{F0: "", F1: 18446744073709551615}
                    _ = v24
                }
                _ = fi22
                _ = v23
            }
            _ = fi20
            _ = v21
        }
        _ = v19
    }
    v17 += (^(v17))
    fmt.Println(v17, v12)
    v17 = v17
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
