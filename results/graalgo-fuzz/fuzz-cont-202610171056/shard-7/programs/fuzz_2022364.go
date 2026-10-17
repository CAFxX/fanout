package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint8; F2 uint16 }
type GenS2 struct { F0 uint8; F1 uint; F2 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint32; F1 int64 }
type GenS5 struct { F0 int32 }
func genFn0(p0 int16) uint16 {
    p0 = (p0 / ((p0) | 1))
    if (true && (("hello" + "a") == "\n\t")) {
        var v5 uint32 = 42
        _ = v5
    }
    v6 := []int{854, 128, 0}
    _ = p0
    _ = v6
    return 255
}
func genFn1(p0 int64) uint8 {
    for fi7 := 0; fi7 < 5; fi7++ {
        v8 := []int{int(p0), 2147483647}
        for fi9 := 0; fi9 < 13; fi9++ {
            var v10 uint32 = 2
            if false {
                var v11 uint8 = 255
                _ = v11
            }
            _ = fi9
            _ = v10
        }
        switch "" {
        case "a":
            p0 = (p0 - -9223372036854775808)
        case "\n\t":
            p0 = (p0 * (p0 / ((p0) | 1)))
        case "foo bar":
            switch "" {
            case "a":
                p0 = p0
            case "foo bar":
                switch (("a" + "hello") + ("" + "hello")) {
                case "hello":
                    switch (("hello" + "") + (("" + "foo bar") + "a")) {
                    case "":
                        v12 := GenS1{F0: 42, F1: 205, F2: 1}
                        _ = v12
                    default:
                        var v13 string = "hello"
                        _ = v13
                    }
                }
            case "\n\t":
                v14 := GenS1{F0: 2, F1: 2, F2: 255}
                _ = v14
            }
        default:
            p0 = 9223372036854775807
        }
        _ = fi7
        _ = v8
    }
    var v15 int = -149
    v15 = 4294967295
    for fi16 := 0; fi16 < 2; fi16++ {
        var v17 bool = true
        v15 -= (v15 - (fi16 | v15))
        for fi18 := 0; fi18 < 11; fi18++ {
            if (!(937 == -1)) {
                switch fi18 {
                case 4294967295:
                    if (!v17) {
                        v19 := []int{323, (fi16 % ((fi16) | 1)), v15}
                        _ = v19
                    }
                }
            } else {
                if false {
                    v17 = ((^(p0)) == 9223372036854775807)
                    var v20 int32 = 0
                    switch ("foo bar" + "") {
                    case "hello":
                        v21 := []int{v15}
                        _ = v21
                    case "":
                        v15 -= v15
                    case "foo bar":
                        var v22 int16 = -915
                        _ = v22
                    }
                    _ = v20
                } else {
                    p0 += 42
                    if (217 > 2) {
                        var v23 uint = 65535
                        _ = v23
                    }
                }
                fmt.Println(v15, fi16)
            }
            v24 := GenS2{F0: 127, F1: 0, F2: "foo bar"}
            v25 := []int{32767}
            _ = fi18
            _ = v24
            _ = v25
        }
        _ = fi16
        _ = v17
    }
    var v26 uint64 = 2
    _ = p0
    _ = v15
    _ = v26
    return 0
}

func main() {
    var v1 int64 = 42
    v2 := make(map[string]bool)
    var v3 int32 = -1
    var v4 float32 = -2.25
    v27 := v2[("hello" + ((("a" + (("hello" + ("" + "foo bar")) + ("hello" + "\n\t"))) + "") + "hello"))]
    var v28 string = ("" + ("a" + "hello"))
    v29 := []int{-9223372036854775808, 42}
    v4 = v4
    v28 = "a"
    var v30 uint16 = 42
    fmt.Println(v27, v28)
    if ((127 <= -1) && v27) {
        v31 := GenS3{F0: 127}
        v27 = ((!v27) || (1e10 == float32(float64(-2.25))))
        _ = v31
    } else {
        if len(v29) >= 2 {
            v29 = v29[1:len(v29)]
        }
        switch 42 {
        case 9223372036854775807:
            if true {
                v32 := GenS4{F0: 65535, F1: v1}
                v1 += v1
                v33 := GenS5{F0: -1}
                _ = v32
                _ = v33
            } else {
                if (genFn0(0) < (v30 | v30)) {
                    fmt.Println(v30)
                    _ = len(v29)
                    fmt.Println(v3, v1, v28)
                }
            }
        case -71:
            fmt.Println(v27, v3, v27)
        case 345:
            if len(v29) > 0 {
                v29[0] = 0
            }
        }
    }
    if (genFn0(1) >= v30) {
        switch v28 {
        case "":
            var v34 uint32 = 2
            _ = v34
        case "foo bar":
            v35 := v2[v28]
            _ = v35
        case "a":
            fmt.Println(v4)
        }
        for fi36 := 0; fi36 < 8; fi36++ {
            var v37 int = 2147483647
            var v38 bool = false
            _ = fi36
            _ = v37
            _ = v38
        }
        v28 = ""
    }
    switch 313 {
    case -2147483648:
        v2[(((("foo bar" + "a") + v28) + v28) + "")] = v27
    case 128:
        if (!v27) {
            fmt.Println(v1, v27, v30)
        } else {
            v4 -= v4
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
}
