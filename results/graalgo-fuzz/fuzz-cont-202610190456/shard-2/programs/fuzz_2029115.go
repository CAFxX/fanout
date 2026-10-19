package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int; F2 string }
type GenS2 struct { F0 int8 }

func main() {
    var v1 float32 = 1e10
    var v2 int32 = 1
    var v3 bool = false
    var v4 int = 42
    var v5 int8 = -128
    v6 := []string{("" + "foo bar"), ""}
    for fi7 := 0; fi7 < 4; fi7++ {
        var v8 float32 = 1.5
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        v6 = append(v6, (("" + "foo bar") + (("" + ("foo bar" + "\n\t")) + "hello")))
        _ = fi7
        _ = v8
    }
    var v9 int = v4
    v1 -= float32(-2.25)
    fmt.Println(v9)
    v10 := GenS1{F0: (v9 / (((-(v4))) | 1)), F1: (v4 * (v4 << ((18446744073709551615) & 63))), F2: ""}
    _ = len(v6)
    switch -838 {
    case 128:
        v11 := GenS1{F0: v4, F1: v4, F2: "\n\t"}
        _ = v11
    case 0:
        fmt.Println(v9)
    }
    var v12 uint64 = 369
    if false {
        if ((42 <= 755) && (false && v3)) {
            if (2 != 42) {
                v13 := GenS1{F0: v4, F1: 2147483647, F2: ("" + "\n\t")}
                if v3 {
                    if v3 {
                        var v14 uint = 65535
                        _ = v14
                    } else {
                        var v15 float64 = 0.0
                        var v16 int16 = 32767
                        _ = v15
                        _ = v16
                    }
                    var v17 uint16 = 814
                    _ = v17
                }
                _ = v13
            }
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
        }
        if v3 {
            var v18 string = ("a" + "a")
            var v19 int32 = (^(v2))
            _ = v18
            _ = v19
        } else {
            v20 := GenS1{F0: (v4 / (((v9 + (v9 + v9))) | 1)), F1: v4, F2: "foo bar"}
            var v21 int8 = (-(v5))
            fmt.Println(v9, v2)
            _ = v20
            _ = v21
        }
        for fi22 := 0; fi22 < 4; fi22++ {
            var v23 uint16 = 255
            _ = fi22
            _ = v23
        }
    }
    switch v9 {
    case -416:
        v6 = append(v6, "foo bar")
    case 548:
        switch ("\n\t" + (("" + "a") + ("foo bar" + "a"))) {
        case "\n\t":
            var v24 uint32 = 120
            _ = v24
        default:
            switch (v9 % (((-(v9))) | 1)) {
            case -9223372036854775808:
                var v25 int16 = 0
                _ = v25
            default:
                switch "hello" {
                case "foo bar":
                    fmt.Println(v4)
                default:
                    for fi26 := 0; fi26 < 2; fi26++ {
                        if (false && (-620 == 42)) {
                            v3 = v3
                            var v27 bool = (-2.25 <= 3.14159)
                            _ = v27
                        }
                        v6 = append(v6, "")
                        _ = fi26
                    }
                }
            }
        }
    case 964:
        for fi28 := 0; fi28 < 17; fi28++ {
            v29 := GenS1{F0: (v9 << ((255) & 63)), F1: (^(v4)), F2: "foo bar"}
            var v30 int = -105
            var v31 uint = 42
            _ = fi28
            _ = v29
            _ = v30
            _ = v31
        }
    }
    for fi32 := 0; fi32 < 10; fi32++ {
        switch ("foo bar" + "hello") {
        case "hello":
            fmt.Println(v9, v3, v2)
        case "":
            v1 += float32(0.1)
        case "a":
            switch "\n\t" {
            case "hello":
                if (float32(0.0) >= v1) {
                    var v33 float64 = 1e10
                    _ = v33
                } else {
                    switch ("hello" + "hello") {
                    case "":
                        if true {
                            v34 := GenS1{F0: (v4 & (v9 ^ -780)), F1: -421, F2: "hello"}
                            for fi35 := 0; fi35 < 18; fi35++ {
                                fmt.Println(v3, v9, v2)
                                _ = fi35
                            }
                            _ = v34
                        }
                    }
                    if v3 {
                        v3 = ((!v3) || (false && ((!v3) && ((v5 % ((v5) | 1)) != -61))))
                        var v36 string = "hello"
                        _ = v36
                    } else {
                        v4 = 32767
                        for fi37 := 0; fi37 < 2; fi37++ {
                            switch "a" {
                            case "hello":
                                var v38 uint64 = (v12 + (^(v12)))
                                _ = v38
                            default:
                                v2 += 0
                            }
                            v3 = (4 > 42)
                            v1 = (float32((-2.25 + 3.14159)) - v1)
                            _ = fi37
                        }
                    }
                    v39 := GenS1{F0: (^(v9)), F1: v9, F2: "a"}
                    _ = v39
                }
            case "foo bar":
                var v40 int16 = -32768
                _ = v40
            }
        }
        if (v3 || v3) {
            v41 := GenS2{F0: (v5 / ((v5) | 1))}
            var v42 int64 = int64(v5)
            v42 = 9223372036854775807
            _ = v41
            _ = v42
        }
        _ = fi32
    }
    v5 -= 127
    if len(v6) > 0 {
        v6[0] = "\n\t"
    }
    v43 := GenS1{F0: 0, F1: v4, F2: ""}
    var v44 uint = 0
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v6
    _ = v9
}
