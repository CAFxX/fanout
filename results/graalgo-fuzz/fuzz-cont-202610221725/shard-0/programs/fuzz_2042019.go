package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint16; F1 int8; F2 uint }
type GenS3 struct { F0 int16; F1 uint16; F2 int16 }
func genFn0(p0 bool, p1 uint64, p2 uint32) string {
    p2 -= (p2 % ((4294967295) | 1))
    var v6 float64 = -2.25
    switch -2147483648 {
    case 2147483647:
        fmt.Println(p2)
    case 0:
        var v7 string = "hello"
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return ""
}

func main() {
    var v1 int32 = 2147483647
    var v2 uint16 = 65535
    var v3 uint8 = 0
    var v4 uint8 = (v3 * 216)
    v5 := []string{""}
    fmt.Println(v2, v4, v1)
    if false {
        var v8 int32 = (v1 % (((v1 / ((2147483647) | 1))) | 1))
        v2 -= (v2 << ((255) & 63))
        _ = v8
    } else {
        for fi9 := 0; fi9 < 14; fi9++ {
            v10 := GenS1{F0: v2}
            v11 := GenS1{F0: v2}
            _ = fi9
            _ = v10
            _ = v11
        }
        v2 = (^(v2))
    }
    switch (("foo bar" + "foo bar") + "hello") {
    case "":
        v1 = 0
    case "a":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    case "foo bar":
        fmt.Println(v4, v1)
    default:
        var v12 int32 = 42
        _ = v12
    }
    if (v2 < v2) {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        switch "" {
        case "":
            var v13 int16 = 1
            _ = v13
        case "hello":
            for fi14 := 0; fi14 < 19; fi14++ {
                var v15 uint8 = v3
                var v16 int8 = -36
                v3 += 61
                _ = fi14
                _ = v15
                _ = v16
            }
        case "foo bar":
            if false {
                var v17 float32 = 0.1
                for fi18 := 0; fi18 < 6; fi18++ {
                    v19 := GenS2{F0: v2, F1: 9, F2: 956}
                    var v20 float64 = (3.14159 * ((0.0 / 1.0) / 0.1))
                    _ = fi18
                    _ = v19
                    _ = v20
                }
                if false {
                    fmt.Println(v3)
                    fmt.Println(v17)
                    v5 = append(v5, ((("foo bar" + "\n\t") + ("foo bar" + "a")) + "a"))
                }
                _ = v17
            }
        }
    } else {
        var v21 uint8 = v4
        var v22 float32 = 0.0
        _ = v21
        _ = v22
    }
    var v23 uint64 = 630
    for fi24 := 0; fi24 < 15; fi24++ {
        fmt.Println(v3, v4, v1)
        _ = fi24
    }
    v23 += v23
    var v25 int16 = 0
    v3 -= v4
    if true {
        if len(v5) > 0 {
            v5[0] = (("a" + "") + "hello")
        }
        var v26 float32 = (3.14159 * 0.1)
        _ = v26
    } else {
        for fi27 := 0; fi27 < 1; fi27++ {
            var v28 float32 = ((1.5 + -2.25) * float32(0.0))
            _ = fi27
            _ = v28
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    switch 109 {
    case -2147483648:
        _ = len(v5)
    case -150:
        var v29 uint32 = 4294967295
        _ = v29
    case -452:
        fmt.Println(v3)
    default:
        var v30 string = ""
        _ = v30
    }
    if len(v5) > 0 {
        v5[0] = (("foo bar" + "foo bar") + "hello")
    }
    v31 := GenS2{F0: (^(v2)), F1: 31, F2: 4294967295}
    for fi32 := 0; fi32 < 11; fi32++ {
        if true {
            for fi33 := 0; fi33 < 4; fi33++ {
                v34 := GenS3{F0: -905, F1: 2, F2: (v25 & v25)}
                v35 := GenS1{F0: v2}
                _ = fi33
                _ = v34
                _ = v35
            }
            if (-1 <= v25) {
                for fi36 := 0; fi36 < 20; fi36++ {
                    if (("a" + "hello") >= "\n\t") {
                        var v37 string = ("hello" + "foo bar")
                        switch v37 {
                        case "\n\t":
                            v38 := GenS2{F0: v2, F1: 100, F2: 42}
                            _ = v38
                        case "a":
                            v2 -= 42
                        }
                        _ = v37
                    } else {
                        var v39 uint8 = 42
                        _ = v39
                    }
                    _ = fi36
                }
            } else {
                var v40 int32 = v1
                v41 := GenS3{F0: 1000, F1: (v2 / (((v2 >> ((928) & 63))) | 1)), F2: (v25 | v25)}
                _ = v40
                _ = v41
            }
        } else {
            v42 := GenS3{F0: 32767, F1: 11, F2: -1}
            _ = v42
        }
        v25 = (-(v25))
        _ = fi32
    }
    var v43 uint64 = v23
    var v44 bool = true
    v43 = 2
    v5 = append(v5, ("foo bar" + ""))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v25
    _ = v3
    _ = v31
    _ = v4
    _ = v43
    _ = v44
    _ = v5
}
