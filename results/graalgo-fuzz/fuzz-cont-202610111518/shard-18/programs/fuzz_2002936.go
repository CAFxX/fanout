package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }
type GenS2 struct { F0 int64 }
func genFn0(p0 int8, p1 uint32, p2 uint16) int64 {
    v7 := GenS1{F0: 42, F1: 32768}
    var v8 uint32 = (p1 | p1)
    v9 := []int{-2147483648, -2147483648, -690}
    if len(v9) > 0 {
        v9[0] = 42
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return -614
}
func genFn1(p0 uint16) string {
    switch 647 {
    case 9223372036854775807:
        v10 := []int{}
        _ = v10
    case -9223372036854775808:
        for fi11 := 0; fi11 < 3; fi11++ {
            var v12 string = "\n\t"
            _ = fi11
            _ = v12
        }
    case 2147483647:
        fmt.Println(p0)
    default:
        var v13 float64 = (1e10 + 1e10)
        _ = v13
    }
    for fi14 := 0; fi14 < 13; fi14++ {
        p0 += p0
        v15 := GenS1{F0: fi14, F1: fi14}
        var v16 float64 = (1.5 - ((float64(1.5) + 0.1) / 3.14159))
        _ = fi14
        _ = v15
        _ = v16
    }
    v17 := []int{}
    var v18 int = 32768
    var v19 float64 = 0.0
    _ = p0
    _ = v17
    _ = v18
    _ = v19
    return ((("a" + "foo bar") + "foo bar") + "\n\t")
}

func main() {
    var v1 uint = 65535
    var v2 int16 = -90
    var v3 int32 = -1
    var v4 uint8 = 180
    v5 := []string{"foo bar", "\n\t", "foo bar", ((((("foo bar" + "hello") + ("\n\t" + "")) + "hello") + "foo bar") + ("" + "hello"))}
    var v6 int32 = v3
    if false {
        fmt.Println(v1)
        var v20 int8 = 56
        _ = v20
    } else {
        var v21 uint32 = 2147483648
        v2 = v2
        _ = v21
    }
    if true {
        var v22 int64 = -1
        _ = v22
    } else {
        if len(v5) > 0 {
            v5[0] = "hello"
        }
        if true {
            v6 -= -95
            _ = len(v5)
        } else {
            var v23 bool = false
            v24 := GenS1{F0: 1, F1: -1}
            fmt.Println(v4)
            _ = v23
            _ = v24
        }
        v25 := GenS1{F0: -9223372036854775808, F1: -797}
        _ = v25
    }
    var v26 int8 = 0
    v2 = v2
    var v27 uint8 = 2
    if ("" > (("a" + "\n\t") + "")) {
        v5 = append(v5, "")
        if ((v1 >= 4294967295) || true) {
            switch 128 {
            case 0:
                v6 = v6
            case -787:
                v6 = v3
            }
            fmt.Println(v1, v2, v6)
        }
        switch 9223372036854775807 {
        case 910:
            var v28 float32 = 1e10
            _ = v28
        default:
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
    }
    if true {
        switch (("" + "\n\t") + "foo bar") {
        case "hello":
            v29 := GenS1{F0: 128, F1: 128}
            _ = v29
        case "a":
            var v30 uint8 = 2
            _ = v30
        default:
            var v31 float32 = (1.5 / -2.25)
            _ = v31
        }
        for fi32 := 0; fi32 < 16; fi32++ {
            v33 := GenS1{F0: (fi32 / ((fi32) | 1)), F1: fi32}
            if len(v5) > 0 {
                v5[0] = ""
            }
            for fi34 := 0; fi34 < 18; fi34++ {
                v3 = v6
                v35 := GenS1{F0: (fi32 / ((fi34) | 1)), F1: int(v4)}
                _ = fi34
                _ = v35
            }
            _ = fi32
            _ = v33
        }
        var v36 float64 = 0.0
        _ = v36
    } else {
        for fi37 := 0; fi37 < 18; fi37++ {
            for fi38 := 0; fi38 < 5; fi38++ {
                var v39 int = 2147483647
                switch "\n\t" {
                case "\n\t":
                    _ = len(v5)
                    var v40 int32 = v3
                    var v41 int = 42
                    fmt.Println(v3, v26)
                    v42 := GenS1{F0: (fi37 >> (((v1 % ((v1) | 1))) & 63)), F1: fi37}
                    _ = len(v5)
                    var v43 uint8 = v4
                    _ = v40
                    _ = v41
                    _ = v42
                    _ = v43
                }
                var v44 uint64 = 4294967295
                _ = fi38
                _ = v39
                _ = v44
            }
            switch ((("a" + "hello") + "\n\t") + "hello") {
            case "":
                var v45 int64 = 42
                _ = v45
            default:
                var v46 int64 = 42
                _ = v46
            }
            v2 = (^(v2))
            _ = fi37
        }
        v2 -= v2
    }
    v5 = append(v5, ((("" + "") + ("" + "")) + ""))
    _ = len(v5)
    v47 := GenS2{F0: genFn0(v26, 2147483647, 714)}
    v6 = (v3 - v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v6
}
