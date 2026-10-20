package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 uint }
type GenS2 struct { F0 uint32; F1 string; F2 int }
type GenS3 struct { F0 uint16; F1 uint64; F2 int16 }
type GenS4 struct { F0 int16 }
func genFn0(p0 int) int64 {
    if (p0 < -1) {
        var v5 float64 = 3.14159
        for fi6 := 0; fi6 < 15; fi6++ {
            fmt.Println(p0)
            p0 = (p0 / ((p0) | 1))
            v7 := []int{-455, (p0 & 32767), -9223372036854775808}
            _ = fi6
            _ = v7
        }
        _ = v5
    } else {
        v8 := []int{p0, p0}
        switch p0 {
        case -695:
            var v9 float64 = -2.25
            _ = v9
        case 9223372036854775807:
            fmt.Println(p0)
        case 1:
            if len(v8) > 0 {
                v8[0] = (-(p0))
            }
        }
        _ = v8
    }
    fmt.Println(p0)
    p0 += -2147483648
    _ = p0
    return 516
}

func main() {
    v1 := make(map[string]string)
    var v2 uint64 = 800
    var v3 string = ""
    var v4 int16 = 1000
    v10 := GenS1{F0: 255, F1: 0, F2: 1}
    v11 := GenS1{F0: 32767, F1: 42, F2: 2}
    v12 := []int{814, 128}
    switch "foo bar" {
    case "\n\t":
        for fi13 := 0; fi13 < 20; fi13++ {
            var v14 uint32 = 1
            var v15 uint16 = 2
            switch "" {
            case "foo bar":
                v3 = (v3 + "")
            }
            _ = fi13
            _ = v14
            _ = v15
        }
    default:
        switch 32768 {
        case 42:
            if len(v12) >= 2 {
                v12 = v12[1:len(v12)]
            }
        case 500:
            for fi16 := 0; fi16 < 1; fi16++ {
                v1[""] = (("hello" + "\n\t") + (((v3 + v3) + "") + v3))
                _ = fi16
            }
        case 127:
            var v17 uint64 = v2
            _ = v17
        default:
            v18 := GenS1{F0: 125, F1: 255, F2: 647}
            _ = v18
        }
    }
    switch v3 {
    case "a":
        v2 -= v2
    }
    fmt.Println(v2)
    if true {
        fmt.Println(v3, v4)
    } else {
        if ((-2.25 * 0.0) != 0.0) {
            if false {
                var v19 int8 = 19
                if ((v2 | v2) < v2) {
                    var v20 bool = (!("a" < v3))
                    var v21 int8 = v19
                    v20 = (v4 < (-(v4)))
                    _ = v20
                    _ = v21
                }
                var v22 float64 = 3.14159
                _ = v19
                _ = v22
            }
        } else {
            if (v4 != v4) {
                var v23 bool = true
                switch -1 {
                case -2147483648:
                    if len(v12) >= 2 {
                        v12 = v12[1:len(v12)]
                    }
                default:
                    switch v3 {
                    case "":
                        var v24 int64 = genFn0(-194)
                        _ = v24
                    }
                }
                _ = v23
            }
            delete(v1, v3)
            v25 := GenS1{F0: 32767, F1: 0, F2: 184}
            _ = v25
        }
        switch (((v3 + v3) + v3) + (v3 + v3)) {
        case "hello":
            fmt.Println(v2, v4, v2)
        case "":
            switch v3 {
            case "hello":
                v3 = v3
            }
        case "\n\t":
            if true {
                for fi26 := 0; fi26 < 6; fi26++ {
                    var v27 int16 = v4
                    fmt.Println(v2, v3)
                    _ = fi26
                    _ = v27
                }
            } else {
                v28 := GenS2{F0: 236, F1: (v3 + ("a" + (v3 + ((v3 + v3) + v3)))), F2: 127}
                v29 := GenS2{F0: 4294967295, F1: "a", F2: -350}
                fmt.Println(v3, v4, v4)
                _ = v28
                _ = v29
            }
        }
        v30 := GenS3{F0: 2, F1: (v2 >> ((255) & 63)), F2: (v4 << ((631) & 63))}
        _ = v30
    }
    var v31 string = v3
    v32 := GenS2{F0: 563, F1: ((((v3 + v31) + ("foo bar" + "foo bar")) + "\n\t") + "a"), F2: 9223372036854775807}
    var v33 bool = (float64(3.14159) >= 1e10)
    fmt.Println(v2, v2)
    v3 = ("hello" + ("\n\t" + (v31 + v3)))
    var v34 uint = 2
    var v35 string = ((v31 + v3) + "")
    v36 := GenS4{F0: v4}
    var v37 uint32 = 2147483647
    switch ("\n\t" + ((v31 + "\n\t") + (v35 + ((v31 + "a") + "foo bar")))) {
    case "":
        fmt.Println(v3, v37)
    }
    fmt.Println(v34)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
}
