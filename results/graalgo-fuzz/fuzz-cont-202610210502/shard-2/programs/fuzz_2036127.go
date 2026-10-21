package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 int8) uint32 {
    p0 = (p0 / ((74) | 1))
    var v7 float32 = float32(3.14159)
    switch 32767 {
    case 1:
        v7 += v7
    case 2147483647:
        var v8 int8 = -73
        _ = v8
    default:
        var v9 int32 = 42
        _ = v9
    }
    v10 := GenS1{F0: 1}
    _ = p0
    _ = v10
    _ = v7
    return 183
}
func genFn1(p0 uint16) bool {
    switch "\n\t" {
    case "\n\t":
        var v11 int = 1
        _ = v11
    default:
        switch 826 {
        case 4294967295:
            switch 129 {
            case -979:
                p0 = (p0 / (((p0 & 907)) | 1))
            case -9223372036854775808:
                var v12 int = 127
                _ = v12
            case -2147483648:
                p0 -= (p0 & (p0 ^ (^(p0))))
            }
        default:
            for fi13 := 0; fi13 < 20; fi13++ {
                v14 := []int{(-(fi13))}
                _ = fi13
                _ = v14
            }
        }
    }
    if (!true) {
        var v15 int8 = -128
        _ = v15
    } else {
        var v16 uint32 = 2
        v17 := []int{32767, 127}
        var v18 float32 = (1e10 - -2.25)
        _ = v16
        _ = v17
        _ = v18
    }
    p0 = 0
    var v19 uint8 = 255
    _ = p0
    _ = v19
    return false
}

func main() {
    v1 := []string{("foo bar" + (("hello" + "\n\t") + ("" + "foo bar"))), "foo bar"}
    var v2 int32 = -2147483648
    var v3 int8 = -128
    var v4 bool = (1 != 2)
    var v5 uint16 = 35
    v6 := []int{9223372036854775807}
    var v20 bool = (2 < 562)
    switch (("foo bar" + "hello") + ("hello" + (("foo bar" + "foo bar") + "a"))) {
    case "a":
        v21 := GenS2{F0: 42}
        _ = v21
    }
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v22 bool = v4
    v20 = false
    if v4 {
        switch -2147483648 {
        case 165:
            var v23 bool = true
            _ = v23
        case 4294967295:
            if (((^(v5)) < (v5 << ((475) & 63))) && false) {
                v24 := GenS1{F0: v2}
                if (true || (uint32(v5) != genFn0((v3 * v3)))) {
                    if v22 {
                        var v25 float64 = 0.0
                        _ = v25
                    } else {
                        switch 32768 {
                        case 944:
                            v5 -= 1
                        case 1:
                            v22 = v20
                        }
                        fmt.Println(v5)
                    }
                    var v26 uint64 = 18446744073709551615
                    _ = v26
                } else {
                    switch 42 {
                    case 2147483647:
                        v1 = append(v1, "foo bar")
                    case 552:
                        if (18446744073709551615 > 45) {
                            var v27 uint16 = v5
                            for fi28 := 0; fi28 < 12; fi28++ {
                                var v29 int32 = (v2 - -2147483648)
                                _ = fi28
                                _ = v29
                            }
                            var v30 float32 = 1.5
                            _ = v27
                            _ = v30
                        }
                    case 588:
                        v22 = v20
                    }
                }
                v6 = append(v6, -813)
                _ = v24
            }
        default:
            _ = len(v1)
        }
    } else {
        var v31 uint8 = 72
        v32 := GenS1{F0: 42}
        _ = v31
        _ = v32
    }
    v20 = v4
    for fi33 := 0; fi33 < 12; fi33++ {
        switch "\n\t" {
        case "a":
            if v22 {
                if true {
                    var v34 float64 = 3.14159
                    if (1 >= 0) {
                        var v35 float32 = 3.14159
                        _ = len(v1)
                        _ = v35
                    } else {
                        _ = len(v6)
                    }
                    _ = v34
                } else {
                    v36 := GenS1{F0: v2}
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    _ = v36
                }
            }
        case "\n\t":
            if len(v1) > 0 {
                v1[0] = "a"
            }
        case "foo bar":
            var v37 uint32 = 329
            _ = v37
        default:
            var v38 float32 = 1.5
            _ = v38
        }
        for fi39 := 0; fi39 < 1; fi39++ {
            var v40 int = fi39
            _ = fi39
            _ = v40
        }
        v3 = v3
        _ = fi33
    }
    var v41 bool = (!true)
    var v42 int64 = 1
    var v43 int16 = 0
    var v44 uint32 = genFn0(0)
    v6 = append(v6, -572)
    for fi45 := 0; fi45 < 15; fi45++ {
        v5 = (^(v5))
        var v46 uint32 = (v44 + 65535)
        _ = fi45
        _ = v46
    }
    if len(v6) > 0 {
        v6[0] = 9223372036854775807
    }
    var v47 uint8 = 91
    if len(v6) > 0 {
        v6[0] = -9223372036854775808
    }
    var v48 string = ("hello" + (("a" + "") + ("a" + "")))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v22
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v47
    _ = v48
    _ = v5
    _ = v6
}
