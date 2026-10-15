package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16; F1 int16; F2 int32 }
type GenS3 struct { F0 int; F1 int8; F2 uint64 }
func genFn0() uint32 {
    for fi6 := 0; fi6 < 4; fi6++ {
        var v7 uint = 2
        _ = fi6
        _ = v7
    }
    var v8 int8 = -128
    v9 := []int{}
    for fi10 := 0; fi10 < 17; fi10++ {
        v8 -= v8
        _ = fi10
    }
    _ = v8
    _ = v9
    return 1
}

func main() {
    var v1 int32 = 0
    var v2 string = "foo bar"
    var v3 bool = true
    v4 := make(map[int]string)
    v5 := make(map[string]bool)
    v11 := v5[("hello" + (v2 + v2))]
    for fi12 := 0; fi12 < 4; fi12++ {
        if (!true) {
            if (65535 <= 42) {
                v13 := GenS1{F0: 1}
                _ = v13
            } else {
                v14 := v5[v2]
                var v15 uint8 = 0
                _ = v14
                _ = v15
            }
        }
        switch "foo bar" {
        case "\n\t":
            v4[(fi12 / (((fi12 + fi12)) | 1))] = (v2 + v2)
        case "":
            var v16 uint = 42
            _ = v16
        default:
            for fi17 := 0; fi17 < 14; fi17++ {
                v1 = v1
                var v18 int8 = -118
                _ = fi17
                _ = v18
            }
        }
        _ = len(v5)
        _ = fi12
    }
    fmt.Println(v2, v3)
    v11 = v3
    v1 += 740
    var v19 string = ("foo bar" + v2)
    switch (v2 + ((v19 + v19) + ((v2 + "hello") + "hello"))) {
    case "\n\t":
        switch "foo bar" {
        case "hello":
            var v20 int = -812
            _ = v20
        }
    case "foo bar":
        v21 := v5[(v2 + "foo bar")]
        _ = v21
    default:
        if (!(3.14159 >= 0.0)) {
            var v22 string = ""
            _ = v22
        }
    }
    var v23 string = v2
    var v24 float32 = 1.5
    if v11 {
        switch 966 {
        case 32767:
            for fi25 := 0; fi25 < 13; fi25++ {
                var v26 string = (("foo bar" + v23) + ((v23 + ("" + v19)) + (v23 + "a")))
                v27 := GenS1{F0: 730}
                _ = fi25
                _ = v26
                _ = v27
            }
        case -1:
            if (!false) {
                v28 := []int{-23, -511, -2147483648, -496}
                v11 = ((!v3) && (v1 <= v1))
                _ = v28
            }
        case -2147483648:
            if ("foo bar" != (v19 + (v2 + "hello"))) {
                var v29 int64 = 634
                fmt.Println(v29, v2)
                _ = v29
            } else {
                if true {
                    if (!(!false)) {
                        for fi30 := 0; fi30 < 2; fi30++ {
                            v3 = false
                            v11 = true
                            var v31 int64 = -9223372036854775808
                            _ = fi30
                            _ = v31
                        }
                        v32 := GenS2{F0: 42, F1: -146, F2: (v1 * v1)}
                        _ = v32
                    } else {
                        fmt.Println(v24, v1)
                        var v33 float64 = ((-2.25 * -2.25) * -2.25)
                        var v34 int32 = -367
                        _ = v33
                        _ = v34
                    }
                    fmt.Println(v3, v11)
                } else {
                    var v35 float64 = 1e10
                    var v36 int = 42
                    for fi37 := 0; fi37 < 20; fi37++ {
                        for fi38 := 0; fi38 < 7; fi38++ {
                            v3 = v3
                            switch (v2 + ((v2 + v23) + v19)) {
                            case "a":
                                var v39 bool = v11
                                _ = v39
                            }
                            v23 = v2
                            _ = fi38
                        }
                        _ = fi37
                    }
                    _ = v35
                    _ = v36
                }
            }
        }
    } else {
        v40 := GenS1{F0: 266}
        v41 := GenS3{F0: 42, F1: -84, F2: 121}
        if false {
            v1 += v1
            if v3 {
                v42 := []int{589, -829, 9223372036854775807}
                _ = v42
            } else {
                if false {
                    v5["foo bar"] = v3
                }
                v43 := GenS2{F0: 967, F1: 32767, F2: v1}
                v2 = (v23 + v19)
                _ = v43
            }
            var v44 uint8 = 42
            _ = v44
        } else {
            fmt.Println(v11)
        }
        _ = v40
        _ = v41
    }
    v11 = (!((0.0 + (v24 / (v24 / 1.0))) == 1.5))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v19
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
