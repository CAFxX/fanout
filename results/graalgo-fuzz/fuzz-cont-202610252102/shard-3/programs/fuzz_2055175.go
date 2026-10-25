package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 uint16 }
type GenS2 struct { F0 int16 }
func genFn0(p0 int16, p1 bool, p2 uint32) int16 {
    v6 := []int{1, -1, 654, -2147483648}
    switch -737 {
    case 0:
        switch 367 {
        case 192:
            for fi7 := 0; fi7 < 18; fi7++ {
                var v8 uint64 = 0
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                _ = fi7
                _ = v8
            }
        case -1:
            if false {
                p1 = false
                if p1 {
                    p0 = 1000
                    _ = len(v6)
                } else {
                    var v9 float64 = ((((1e10 * 1e10) + (1.5 - 0.0)) - 0.1) - -2.25)
                    p0 = 0
                    _ = v9
                }
            }
        }
    }
    if (("\n\t" >= "foo bar") && (true || (p0 <= 0))) {
        switch ("hello" + ("" + "foo bar")) {
        case "a":
            if len(v6) > 0 {
                v6[0] = 9223372036854775807
            }
        case "":
            var v10 bool = p1
            _ = v10
        case "hello":
            var v11 int16 = -251
            _ = v11
        }
        for fi12 := 0; fi12 < 2; fi12++ {
            v13 := GenS1{F0: 283, F1: 42, F2: 263}
            var v14 uint16 = 2
            p2 = 327
            _ = fi12
            _ = v13
            _ = v14
        }
        var v15 uint = 992
        _ = v15
    } else {
        var v16 uint64 = 42
        p0 -= -134
        v17 := GenS1{F0: 2, F1: (v16 & (v16 - v16)), F2: 943}
        _ = v16
        _ = v17
    }
    for fi18 := 0; fi18 < 6; fi18++ {
        p1 = p1
        v19 := GenS1{F0: 884, F1: 18446744073709551615, F2: 65535}
        var v20 float64 = float64(1e10)
        _ = fi18
        _ = v19
        _ = v20
    }
    if (p1 || false) {
        var v21 uint16 = 1
        _ = v21
    } else {
        v22 := GenS2{F0: (^(p0))}
        _ = v22
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return p0
}
func genFn1() int8 {
    for fi23 := 0; fi23 < 1; fi23++ {
        var v24 string = "a"
        var v25 float64 = 3.14159
        var v26 string = v24
        _ = fi23
        _ = v24
        _ = v25
        _ = v26
    }
    v27 := []int{}
    _ = v27
    return 42
}

func main() {
    var v1 float64 = 1.5
    var v2 int = 32768
    var v3 int32 = 2147483647
    var v4 int64 = 39
    v5 := []uint32{633, 244, 2147483647, 42}
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi28 := 0; fi28 < 3; fi28++ {
        v1 = 0.1
        _ = fi28
    }
    var v29 float32 = ((0.1 / ((float32(((0.1 - float64(1.5)) - v1)) * 0.1) / 3.14159)) + 1.5)
    v3 -= v3
    switch ("\n\t" + ((("hello" + ("foo bar" + "\n\t")) + "foo bar") + (("foo bar" + "\n\t") + "foo bar"))) {
    case "foo bar":
        fmt.Println(v4, v2)
    case "a":
        if false {
            var v30 bool = (v4 != 42)
            _ = v30
        }
    case "\n\t":
        var v31 uint64 = 1
        _ = v31
    default:
        var v32 bool = false
        _ = v32
    }
    for fi33 := 0; fi33 < 12; fi33++ {
        if (!true) {
            v1 = (v1 * ((v1 / v1) - v1))
        }
        _ = fi33
    }
    fmt.Println(v4, v29, v29)
    v34 := GenS1{F0: 0, F1: 2, F2: 65535}
    v2 = v2
    for fi35 := 0; fi35 < 1; fi35++ {
        for fi36 := 0; fi36 < 9; fi36++ {
            v3 = (v3 & -1)
            v29 += 1.5
            _ = fi36
        }
        _ = fi35
    }
    v37 := GenS2{F0: 0}
    v38 := GenS1{F0: 255, F1: 4294967295, F2: 0}
    v1 = v1
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if (v4 > (v4 / ((953) | 1))) {
        var v39 bool = (18446744073709551615 < 18446744073709551615)
        _ = len(v5)
        _ = v39
    } else {
        if true {
            v3 += (v3 ^ 42)
            if len(v5) > 0 {
                v5[0] = 384
            }
        }
    }
    for fi40 := 0; fi40 < 7; fi40++ {
        v1 = (float64(v29) + ((0.0 + v1) - v1))
        _ = fi40
    }
    var v41 float32 = 1.5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v34
    _ = v37
    _ = v38
    _ = v4
    _ = v41
    _ = v5
}
