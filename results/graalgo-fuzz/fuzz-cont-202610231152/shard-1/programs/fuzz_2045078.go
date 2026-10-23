package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 int }
type GenS2 struct { F0 int8; F1 uint8 }
type GenS3 struct { F0 int16; F1 uint16 }
type GenS4 struct { F0 string }

func main() {
    var v1 int = 9223372036854775807
    v2 := make(map[string]string)
    var v3 int16 = 52
    var v4 bool = true
    v5 := make(map[string]uint32)
    var v6 uint32 = 2147483648
    for fi7 := 0; fi7 < 13; fi7++ {
        v8 := GenS1{F0: 0, F1: ("a" + "foo bar"), F2: 174}
        for fi9 := 0; fi9 < 2; fi9++ {
            v10 := []int{(fi9 * fi7)}
            if (1 >= 309) {
                v1 += 841
                v11 := GenS2{F0: 1, F1: 128}
                if len(v10) > 0 {
                    v10[0] = fi7
                }
                _ = v11
            } else {
                var v12 int = (^(v1))
                var v13 uint32 = v6
                var v14 uint16 = 954
                _ = v12
                _ = v13
                _ = v14
            }
            _ = fi9
            _ = v10
        }
        if (127 > 1) {
            var v15 int32 = -1
            _ = v15
        } else {
            switch -40 {
            case 2147483647:
                v3 += (^(v3))
            case 9223372036854775807:
                var v16 uint = 0
                _ = v16
            }
            v6 = 2147483647
            v4 = true
        }
        _ = fi7
        _ = v8
    }
    switch "hello" {
    case "hello":
        if (!v4) {
            fmt.Println(v1, v4, v3)
            if v4 {
                var v17 int32 = 1
                switch "hello" {
                case "\n\t":
                    v18 := v5[("foo bar" + "a")]
                    _ = v18
                case "a":
                    if (v3 > (v3 + (-(v3)))) {
                        v6 += (v6 & (v6 - v6))
                    }
                case "":
                    for fi19 := 0; fi19 < 9; fi19++ {
                        v3 = v3
                        if (((fi19 * v1) >= fi19) && (!(v4 && true))) {
                            v17 = 1
                        } else {
                            var v20 uint = 18446744073709551615
                            for fi21 := 0; fi21 < 13; fi21++ {
                                v5[("" + "hello")] = (^(v6))
                                var v22 int8 = -1
                                switch (fi21 >> ((v20) & 63)) {
                                case 42:
                                    switch (^(fi19)) {
                                    case -1:
                                        var v23 float32 = (-2.25 + 3.14159)
                                        _ = v23
                                    case -708:
                                        v24 := []int{-907, (fi21 - 2147483647)}
                                        _ = v24
                                    }
                                }
                                _ = fi21
                                _ = v22
                            }
                            _ = v20
                        }
                        _ = fi19
                    }
                }
                v17 += (v17 * (v17 - (v17 * (-(v17)))))
                _ = v17
            } else {
                if false {
                    var v25 int8 = -1
                    _ = v25
                }
                v5[(((("\n\t" + "foo bar") + ("\n\t" + "\n\t")) + ("hello" + "\n\t")) + "foo bar")] = (v6 % ((v6) | 1))
            }
        }
    default:
        var v26 int = 1
        _ = v26
    }
    v1 -= v1
    if ((!true) && ("" > "foo bar")) {
        delete(v5, "hello")
        delete(v5, "foo bar")
        v27 := []int{v1, v1}
        _ = v27
    } else {
        v1 += 128
        if true {
            if v4 {
                for fi28 := 0; fi28 < 11; fi28++ {
                    v29 := GenS3{F0: -32768, F1: 65535}
                    _ = fi28
                    _ = v29
                }
                v4 = (18446744073709551615 < 42)
            }
            var v30 float64 = (0.0 - float64(float32((0.1 + (0.0 + float64(float32(float64(0.1))))))))
            fmt.Println(v6, v3, v1)
            _ = v30
        }
        _ = len(v2)
    }
    fmt.Println(v6)
    var v31 float64 = (float64(3.14159) + 3.14159)
    if false {
        v6 += v6
        switch v1 {
        case 1:
            for fi32 := 0; fi32 < 17; fi32++ {
                v33 := GenS4{F0: "foo bar"}
                delete(v5, "a")
                var v34 int8 = -25
                _ = fi32
                _ = v33
                _ = v34
            }
        case 2147483647:
            fmt.Println(v6)
        case 592:
            v4 = v4
        }
    }
    for fi35 := 0; fi35 < 20; fi35++ {
        for fi36 := 0; fi36 < 19; fi36++ {
            delete(v5, "\n\t")
            _ = fi36
        }
        var v37 bool = (!(18446744073709551615 < 0))
        v31 = v31
        _ = fi35
        _ = v37
    }
    if v4 {
        _ = len(v2)
        var v38 bool = (v4 || v4)
        _ = v38
    }
    var v39 int = (v1 | (v1 << ((157) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v39
    _ = v4
    _ = v5
    _ = v6
}
