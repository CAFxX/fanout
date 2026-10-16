package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 int8 }
type GenS3 struct { F0 int; F1 uint8 }

func main() {
    var v1 bool = false
    v2 := []string{"\n\t", "\n\t", ("\n\t" + (("hello" + "foo bar") + "hello")), ""}
    var v3 int = 32767
    v4 := make(map[string]string)
    if v1 {
        for fi5 := 0; fi5 < 14; fi5++ {
            for fi6 := 0; fi6 < 11; fi6++ {
                v3 = (-(v3))
                v7 := GenS1{F0: "hello"}
                if (((true || (!v1)) || v1) && v1) {
                    fmt.Println(fi6, v3)
                    v8 := GenS1{F0: ("a" + ("\n\t" + "a"))}
                    _ = v8
                } else {
                    var v9 bool = v1
                    v10 := GenS1{F0: ("hello" + "a")}
                    _ = v10
                    _ = v9
                }
                _ = fi6
                _ = v7
            }
            v1 = (!(!v1))
            var v11 uint64 = 0
            _ = fi5
            _ = v11
        }
        var v12 int32 = 2147483647
        if false {
            delete(v4, "a")
            if ((v3 | 323) != v3) {
                var v13 bool = ((false && (56 == 18446744073709551615)) && false)
                _ = v13
            }
            v3 += 0
        } else {
            var v14 int16 = 42
            _ = v14
        }
        _ = v12
    }
    if v1 {
        if v1 {
            v1 = (0 > 1)
            var v15 int32 = -646
            _ = v15
        }
        var v16 float32 = (((((1.5 / 1e10) + ((-2.25 + 0.1) / 1.5)) - 0.0) * ((-2.25 + 0.0) + (1.5 + 0.0))) * (3.14159 / -2.25))
        _ = v16
    } else {
        v1 = (v1 || (242 < 763))
    }
    if v1 {
        v3 -= (v3 + v3)
    }
    v3 = 127
    v3 -= -469
    v17 := GenS2{F0: ((!(false || ((("\n\t" > (("hello" + "") + ("\n\t" + (("hello" + "hello") + ("a" + "foo bar"))))) || true) || v1))) || (2 <= 152)), F1: 0}
    switch -9223372036854775808 {
    case 32767:
        v18 := GenS2{F0: (-9223372036854775808 <= 42), F1: 100}
        _ = v18
    case 715:
        v1 = (((4294967295 < 2147483648) || (!v1)) && false)
    default:
        var v19 uint8 = 163
        _ = v19
    }
    v3 -= 32767
    v20 := GenS1{F0: ""}
    switch (v3 >> ((42) & 63)) {
    case 753:
        var v21 int = 0
        _ = v21
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v22 uint = 42
    v23 := GenS3{F0: -76, F1: 2}
    var v24 int8 = 127
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v17
    _ = v2
    _ = v20
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
