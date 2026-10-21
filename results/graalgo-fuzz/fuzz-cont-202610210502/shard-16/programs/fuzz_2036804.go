package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 int32 }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 float64 = ((3.14159 + float64(1e10)) / 0.1)
    var v2 int = -716
    v3 := make(map[string]bool)
    var v4 int = v2
    v5 := []int{-2147483648}
    if false {
        var v6 int8 = 29
        _ = v6
    } else {
        var v7 int16 = 1000
        _ = v7
    }
    switch (v4 >> ((42) & 63)) {
    case 0:
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    case 9223372036854775807:
        var v8 int = (v2 - v2)
        _ = v8
    }
    var v9 string = (((("foo bar" + ((("foo bar" + (("hello" + ("\n\t" + (((("hello" + "") + "\n\t") + "\n\t") + ("\n\t" + "")))) + (("foo bar" + "foo bar") + ("" + "foo bar")))) + "\n\t") + "hello")) + "\n\t") + (("foo bar" + "a") + "\n\t")) + "foo bar")
    v3[v9] = false
    v5 = append(v5, v4)
    switch ((v9 + v9) + v9) {
    case "":
        if false {
            fmt.Println(v4)
            v2 = (v4 ^ -217)
        }
    }
    for fi10 := 0; fi10 < 2; fi10++ {
        v9 = v9
        for fi11 := 0; fi11 < 4; fi11++ {
            if (true && ((false && false) && (0.0 == (3.14159 + ((1.5 + 1.5) - 3.14159))))) {
                v2 += 128
                if true {
                    var v12 int = (v2 + (v2 * v4))
                    v9 = v9
                    v2 += v4
                    v5 = append(v5, 2147483647)
                    _ = v12
                } else {
                    v4 += (v4 + (v4 + (fi10 - (-(fi10)))))
                }
            } else {
                v5 = append(v5, v2)
                v3[v9] = (!true)
            }
            var v13 int32 = 916
            fmt.Println(v4, v1, v9)
            _ = fi11
            _ = v13
        }
        _ = fi10
    }
    if true {
        var v14 int8 = -123
        var v15 bool = (false || true)
        for fi16 := 0; fi16 < 5; fi16++ {
            for fi17 := 0; fi17 < 12; fi17++ {
                v2 = fi17
                var v18 uint32 = 58
                var v19 uint32 = v18
                _ = fi17
                _ = v18
                _ = v19
            }
            _ = fi16
        }
        _ = v14
        _ = v15
    } else {
        switch v4 {
        case 127:
            var v20 uint8 = 255
            _ = v20
        case 1:
            if (true || false) {
                var v21 uint64 = 274
                if ((v4 * (v4 * (v2 / ((v2) | 1)))) >= (v4 % (((^(v4))) | 1))) {
                    var v22 int8 = 0
                    v23 := GenS1{F0: (v22 << ((18446744073709551615) & 63)), F1: v21, F2: -68}
                    fmt.Println(v1, v9, v2)
                    _ = v22
                    _ = v23
                }
                _ = v21
            } else {
                var v24 bool = (0.0 < (float32(0.0) + (float32(((-2.25 - 0.1) + 1e10)) * 3.14159)))
                if false {
                    for fi25 := 0; fi25 < 15; fi25++ {
                        v1 = v1
                        _ = fi25
                    }
                    v26 := GenS1{F0: -128, F1: 4294967295, F2: 42}
                    var v27 float32 = -2.25
                    _ = v26
                    _ = v27
                }
                v9 = "\n\t"
                _ = v24
            }
        }
        var v28 int16 = -1
        v1 = v1
        _ = v28
    }
    switch v9 {
    case "hello":
        _ = len(v5)
    case "a":
        v4 += (v4 / ((v2) | 1))
    case "\n\t":
        v2 = v2
    default:
        for fi29 := 0; fi29 < 9; fi29++ {
            var v30 float32 = 0.1
            _ = fi29
            _ = v30
        }
    }
    v31 := GenS1{F0: 127, F1: 4294967295, F2: 300}
    v1 = (v1 + float64((1.5 / 1.0)))
    if false {
        var v32 uint64 = 4294967295
        _ = v32
    } else {
        v33 := v3[(v9 + ("\n\t" + ((("a" + v9) + v9) + v9)))]
        _ = v33
    }
    switch v9 {
    case "\n\t":
        v9 = v9
    }
    var v34 int32 = -1
    for fi35 := 0; fi35 < 6; fi35++ {
        v36 := GenS2{F0: 0}
        switch v9 {
        case "foo bar":
            v4 += (^(fi35))
        case "hello":
            var v37 int16 = 608
            _ = v37
        case "":
            if (92 == 0) {
                var v38 uint8 = 2
                _ = v38
            }
        }
        _ = fi35
        _ = v36
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v34
    _ = v4
    _ = v5
    _ = v9
}
