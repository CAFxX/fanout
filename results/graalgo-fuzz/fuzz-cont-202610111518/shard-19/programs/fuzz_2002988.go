package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint; F1 uint8; F2 uint }
type GenS3 struct { F0 uint64; F1 uint32; F2 uint64 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int8 }
func genFn0(p0 int64) uint32 {
    var v5 uint8 = 17
    fmt.Println(v5, p0)
    switch 674 {
    case 769:
        var v6 float64 = float64(-2.25)
        _ = v6
    case 4294967295:
        for fi7 := 0; fi7 < 2; fi7++ {
            switch (fi7 ^ fi7) {
            case -100:
                var v8 string = ((("a" + "hello") + ("hello" + "\n\t")) + "hello")
                _ = v8
            default:
                var v9 uint16 = 255
                _ = v9
            }
            if (!(!(((127 < fi7) || false) && (((v5 | (v5 + v5)) < v5) || (103 == 140))))) {
                var v10 uint32 = 65535
                v11 := GenS1{F0: v10}
                _ = v10
                _ = v11
            } else {
                var v12 uint32 = 812
                _ = v12
            }
            for fi13 := 0; fi13 < 17; fi13++ {
                var v14 uint32 = 2
                v15 := []int{}
                _ = fi13
                _ = v14
                _ = v15
            }
            _ = fi7
        }
    }
    fmt.Println(v5, p0)
    _ = p0
    _ = v5
    return 29
}

func main() {
    var v1 float64 = 1.5
    v2 := make(map[string]int)
    v3 := make(map[int]string)
    var v4 uint64 = 2
    var v16 int64 = 9223372036854775807
    if (2 > 1) {
        for fi17 := 0; fi17 < 14; fi17++ {
            if (2 == 128) {
                for fi18 := 0; fi18 < 8; fi18++ {
                    if (!true) {
                        v19 := GenS2{F0: 4294967295, F1: 186, F2: 65535}
                        fmt.Println(v4, v1, v1)
                        switch ("" + "hello") {
                        case "a":
                            if (!(!(((3.14159 / (v1 / v1)) * v1) <= v1))) {
                                var v20 float64 = ((((3.14159 / ((v1 / -2.25) * v1)) + v1) - ((v1 - v1) / (1e10 - -2.25))) * (((v1 * (float64(-2.25) - 0.1)) - float64((1e10 + 0.1))) - v1))
                                _ = v20
                            }
                        }
                        _ = v19
                    }
                    var v21 uint = 42
                    if (v4 <= (v4 ^ (v4 << ((v21) & 63)))) {
                        v22 := GenS2{F0: (^(v21)), F1: 213, F2: 734}
                        _ = v22
                    }
                    _ = fi18
                    _ = v21
                }
                var v23 int16 = 32767
                fmt.Println(v4)
                _ = v23
            }
            if true {
                var v24 uint8 = 128
                _ = v24
            }
            var v25 int32 = 42
            _ = fi17
            _ = v25
        }
        var v26 float64 = v1
        v27 := []int{-445}
        _ = v26
        _ = v27
    }
    v1 -= 3.14159
    if true {
        var v28 int16 = 1
        switch -34 {
        case -338:
            if false {
                v16 -= v16
                v29 := GenS3{F0: v4, F1: 2, F2: (v4 % (((v4 / (((v4 ^ 137)) | 1))) | 1))}
                _ = v29
            }
        case 1:
            for fi30 := 0; fi30 < 12; fi30++ {
                v31 := v3[278]
                var v32 int64 = v16
                _ = fi30
                _ = v31
                _ = v32
            }
        case 4294967295:
            v33 := []int{9223372036854775807, 127, -9223372036854775808}
            _ = v33
        }
        switch ("a" + "\n\t") {
        case "foo bar":
            var v34 int8 = int8(v4)
            _ = v34
        default:
            v35 := GenS4{F0: (v16 >> ((1) & 63))}
            _ = v35
        }
        _ = v28
    } else {
        var v36 float64 = ((v1 - -2.25) - v1)
        _ = v36
    }
    v37 := []int{-9223372036854775808, -2147483648, 9223372036854775807}
    switch (("hello" + ((("foo bar" + "a") + "foo bar") + ("" + "\n\t"))) + "a") {
    case "\n\t":
        v4 -= 2
    }
    fmt.Println(v4, v1)
    _ = len(v2)
    if true {
        var v38 uint16 = 255
        for fi39 := 0; fi39 < 4; fi39++ {
            v40 := GenS5{F0: 0}
            fmt.Println(v4)
            var v41 string = ("a" + ("" + (("a" + "a") + "hello")))
            _ = fi39
            _ = v40
            _ = v41
        }
        _ = v38
    } else {
        v37 = append(v37, 9223372036854775807)
    }
    fmt.Println(v1, v4, v1)
    if len(v37) > 0 {
        v37[0] = 32767
    }
    v42 := GenS6{F0: 100}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v37:", len(v37), v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v42
}
