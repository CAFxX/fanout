package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint32; F2 int64 }
type GenS2 struct { F0 uint8; F1 uint8; F2 uint8 }
type GenS3 struct { F0 int16; F1 int8 }

func main() {
    var v1 int64 = -9223372036854775808
    var v2 float32 = float32(-2.25)
    v3 := make(map[string]int)
    v4 := []int{}
    v5 := GenS1{F0: 0, F1: 42, F2: -1}
    var v6 int16 = -1
    for fi7 := 0; fi7 < 15; fi7++ {
        if false {
            var v8 string = ("a" + "\n\t")
            _ = v8
        }
        v2 = (1e10 / v2)
        _ = fi7
    }
    var v9 int16 = (v6 | (v6 + (v6 * 586)))
    for fi10 := 0; fi10 < 17; fi10++ {
        v9 += (-(v6))
        _ = fi10
    }
    switch 0 {
    case 0:
        _ = len(v4)
    }
    if false {
        for fi11 := 0; fi11 < 10; fi11++ {
            for fi12 := 0; fi12 < 1; fi12++ {
                if true {
                    fmt.Println(v6)
                    var v13 float32 = (v2 / -2.25)
                    if (!true) {
                        _ = len(v4)
                        var v14 bool = false
                        var v15 int16 = (v6 + (v9 >> ((87) & 63)))
                        _ = v14
                        _ = v15
                    } else {
                        var v16 int16 = -378
                        v1 = (v1 / ((0) | 1))
                        if len(v4) > 0 {
                            v4[0] = fi12
                        }
                        v16 += 1000
                        _ = v16
                    }
                    _ = v13
                }
                delete(v3, ("a" + ""))
                _ = fi12
            }
            var v17 uint = 14
            _ = fi11
            _ = v17
        }
        var v18 float64 = ((3.14159 * -2.25) - (0.0 - 1.5))
        v4 = append(v4, 0)
        _ = v18
    } else {
        switch "foo bar" {
        case "hello":
            v9 = (v9 ^ 700)
        case "foo bar":
            var v19 float32 = v2
            _ = v19
        case "a":
            v20 := GenS1{F0: 980, F1: 4294967295, F2: 0}
            _ = v20
        }
        for fi21 := 0; fi21 < 1; fi21++ {
            var v22 float64 = 1.5
            _ = fi21
            _ = v22
        }
        for fi23 := 0; fi23 < 18; fi23++ {
            switch fi23 {
            case 0:
                var v24 uint64 = 1
                _ = v24
            default:
                if false {
                    v2 = v2
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    if (true && false) {
                        var v25 int8 = 42
                        _ = v25
                    } else {
                        for fi26 := 0; fi26 < 11; fi26++ {
                            switch (("a" + "a") + "hello") {
                            case "":
                                fmt.Println(v6)
                            case "a":
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                            v27 := v3[""]
                            v28 := GenS2{F0: 42, F1: 128, F2: 236}
                            _ = fi26
                            _ = v27
                            _ = v28
                        }
                    }
                }
            }
            fmt.Println(v9, fi23)
            _ = fi23
        }
    }
    if false {
        var v29 string = "a"
        _ = len(v3)
        v2 -= (v2 + v2)
        _ = v29
    } else {
        fmt.Println(v6, v1)
        var v30 float32 = (v2 - (3.14159 / -2.25))
        _ = v30
    }
    fmt.Println(v2, v9, v6)
    fmt.Println(v6)
    var v31 int32 = 1
    var v32 uint32 = 2147483647
    v33 := GenS3{F0: v6, F1: 31}
    var v34 int = -9223372036854775808
    var v35 float32 = 1.5
    if false {
        if ((2 < 145) && false) {
            if (!(0.1 < 1e10)) {
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                v36 := GenS1{F0: 32767, F1: (v32 * (^(v32))), F2: 0}
                _ = v36
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
