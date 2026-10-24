package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 uint16; F1 uint32; F2 int }
type GenS3 struct { F0 uint; F1 int8 }

func main() {
    var v1 float64 = 1e10
    v2 := []string{(("foo bar" + ("hello" + "a")) + (("a" + "\n\t") + "\n\t")), "a", "\n\t"}
    var v3 int64 = -997
    v4 := []string{"a"}
    var v5 uint16 = 273
    for fi6 := 0; fi6 < 19; fi6++ {
        if false {
            var v7 int16 = 32767
            for fi8 := 0; fi8 < 8; fi8++ {
                var v9 uint = 1
                fmt.Println(v1, fi8)
                v10 := GenS1{F0: v5, F1: (^(v5))}
                _ = fi8
                _ = v10
                _ = v9
            }
            if false {
                var v11 uint32 = 4294967295
                _ = v11
            } else {
                v5 = v5
                var v12 float64 = 3.14159
                v1 = float64(0.1)
                _ = v12
            }
            _ = v7
        }
        v5 = v5
        _ = fi6
    }
    if len(v2) > 0 {
        v2[0] = (("a" + (("foo bar" + "hello") + ("foo bar" + "hello"))) + "hello")
    }
    if false {
        v1 += 1e10
        fmt.Println(v5)
        var v13 bool = (18446744073709551615 < 18446744073709551615)
        _ = v13
    }
    for fi14 := 0; fi14 < 8; fi14++ {
        for fi15 := 0; fi15 < 5; fi15++ {
            v2 = append(v2, ((((("\n\t" + "foo bar") + "hello") + ("a" + "foo bar")) + (("hello" + "\n\t") + "foo bar")) + (("\n\t" + "hello") + "")))
            var v16 string = "foo bar"
            if true {
                switch (-(fi14)) {
                case 1:
                    v17 := GenS2{F0: v5, F1: 2, F2: -649}
                    _ = v17
                case 16:
                    v3 = 42
                }
            }
            _ = fi15
            _ = v16
        }
        _ = len(v2)
        _ = fi14
    }
    _ = len(v4)
    if ((0.1 * 0.0) != float32(-2.25)) {
        v1 -= 0.1
        if false {
            var v18 uint8 = 13
            _ = v18
        } else {
            for fi19 := 0; fi19 < 11; fi19++ {
                for fi20 := 0; fi20 < 2; fi20++ {
                    v21 := GenS1{F0: 0, F1: (v5 >> ((1) & 63))}
                    _ = fi20
                    _ = v21
                }
                var v22 int64 = v3
                for fi23 := 0; fi23 < 18; fi23++ {
                    fmt.Println(v1, v5)
                    var v24 bool = (!false)
                    _ = fi23
                    _ = v24
                }
                _ = fi19
                _ = v22
            }
            if ((2 == 1) || false) {
                v1 = v1
            } else {
                if ((false || false) && true) {
                    v25 := GenS3{F0: 497, F1: 127}
                    _ = v25
                }
                var v26 int = 128
                _ = v26
            }
            v27 := GenS2{F0: (v5 - v5), F1: 2, F2: 9223372036854775807}
            _ = v27
        }
    }
    var v28 float64 = float64((1e10 / 1.0))
    fmt.Println(v5)
    v29 := GenS1{F0: v5, F1: v5}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v4
    _ = v5
}
