package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 string; F2 uint16 }
type GenS3 struct { F0 uint }
func genFn0(p0 int32) int32 {
    var v6 float64 = 0.1
    v6 -= float64((3.14159 * 0.1))
    var v7 uint32 = 0
    v6 -= (v6 * v6)
    _ = p0
    _ = v6
    _ = v7
    return (^(p0))
}

func main() {
    var v1 float32 = (0.0 - 3.14159)
    var v2 float64 = 0.0
    v3 := make(map[uint32]uint32)
    var v4 float64 = float64(float32(1e10))
    var v5 int64 = -642
    v1 = 3.14159
    var v8 int64 = (v5 | v5)
    v9 := GenS1{F0: 797}
    var v10 int16 = 42
    for fi11 := 0; fi11 < 19; fi11++ {
        v2 = 1e10
        v4 += (v2 + (v2 / (v2 / 1.0)))
        switch "hello" {
        case "":
            for fi12 := 0; fi12 < 20; fi12++ {
                if (!true) {
                    var v13 int32 = -516
                    switch (fi12 & int(v13)) {
                    case 803:
                        v8 = 0
                    case -242:
                        v14 := GenS1{F0: 32767}
                        _ = v14
                    default:
                        fmt.Println(v2, v8)
                    }
                    _ = v13
                }
                v5 += v8
                var v15 string = "\n\t"
                _ = fi12
                _ = v15
            }
        default:
            v16 := GenS2{F0: -1, F1: ((("hello" + ("foo bar" + "")) + "\n\t") + "hello"), F2: 255}
            _ = v16
        }
        _ = fi11
    }
    for fi17 := 0; fi17 < 20; fi17++ {
        v18 := []int{(fi17 ^ 128)}
        _ = fi17
        _ = v18
    }
    v19 := []int{0}
    if false {
        var v20 uint16 = 1
        for fi21 := 0; fi21 < 1; fi21++ {
            var v22 uint = 42
            _ = fi21
            _ = v22
        }
        _ = v20
    }
    _ = len(v3)
    fmt.Println(v2)
    v4 = 0.0
    _ = len(v3)
    _ = len(v3)
    for fi23 := 0; fi23 < 16; fi23++ {
        if true {
            fmt.Println(v2, v1)
            v24 := GenS1{F0: fi23}
            _ = v24
        } else {
            v25 := GenS3{F0: 2}
            for fi26 := 0; fi26 < 1; fi26++ {
                switch (^(fi26)) {
                case 2147483647:
                    if true {
                        v2 -= float64((3.14159 / 3.14159))
                        var v27 bool = ((42 <= 42) || ("" > ("foo bar" + "a")))
                        _ = v27
                    } else {
                        if true {
                            switch fi23 {
                            case 32767:
                                var v28 uint64 = 4294967295
                                _ = v28
                            default:
                                var v29 uint8 = 0
                                _ = v29
                            }
                            v30 := GenS2{F0: 100, F1: ("\n\t" + (("foo bar" + "foo bar") + "\n\t")), F2: 788}
                            _ = v30
                        } else {
                            v1 -= v1
                            v19 = append(v19, fi23)
                        }
                        v3[2147483647] = 505
                        fmt.Println(v5)
                    }
                }
                var v31 uint16 = 625
                if len(v19) > 0 {
                    v19[0] = -2147483648
                }
                _ = fi26
                _ = v31
            }
            var v32 string = ((("a" + "hello") + "foo bar") + "a")
            _ = v25
            _ = v32
        }
        for fi33 := 0; fi33 < 2; fi33++ {
            switch "hello" {
            case "\n\t":
                v1 -= float32(((v2 + float64(-2.25)) / v2))
            case "":
                v2 = 1.5
            case "foo bar":
                v19 = append(v19, fi23)
            }
            _ = fi33
        }
        switch (fi23 % ((fi23) | 1)) {
        case 422:
            v2 -= v2
        }
        _ = fi23
    }
    var v34 uint16 = 42
    for fi35 := 0; fi35 < 16; fi35++ {
        var v36 int32 = 1
        fmt.Println(v10)
        _ = fi35
        _ = v36
    }
    switch 1 {
    case -2147483648:
        if false {
            v37 := GenS3{F0: 4294967295}
            switch "" {
            case "foo bar":
                v4 = v4
            default:
                switch int(v10) {
                case 4294967295:
                    var v38 float64 = 0.0
                    _ = v38
                case 2147483647:
                    _ = len(v19)
                default:
                    var v39 float32 = v1
                    _ = v39
                }
            }
            v19 = append(v19, 32768)
            _ = v37
        } else {
            var v40 bool = ((843 <= 909) && true)
            var v41 uint8 = 42
            v40 = v40
            _ = v40
            _ = v41
        }
    case 127:
        var v42 float64 = 1e10
        _ = v42
    }
    v43 := GenS2{F0: 42, F1: "\n\t", F2: v34}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v43
    _ = v5
    _ = v8
    _ = v9
}
