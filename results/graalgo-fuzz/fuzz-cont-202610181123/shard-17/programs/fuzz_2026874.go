package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
func genFn0(p0 uint64, p1 uint16) int32 {
    fmt.Println(p1, p1)
    var v4 int = 3
    var v5 uint16 = 490
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 0
}
func genFn1(p0 int32) uint {
    var v6 int16 = 1000
    if false {
        v6 -= v6
        p0 = p0
        p0 += p0
    } else {
        if true {
            switch "foo bar" {
            case "foo bar":
                var v7 int = -1
                _ = v7
            }
        }
        p0 += (p0 - p0)
    }
    var v8 uint16 = 1
    fmt.Println(p0, p0, p0)
    _ = p0
    _ = v6
    _ = v8
    return 255
}

func main() {
    var v1 int64 = -1
    v2 := make(map[string]uint32)
    v3 := []int{-1, 128, 42}
    v9 := v2["hello"]
    v1 += (^(v1))
    v1 = 9223372036854775807
    var v10 uint32 = v9
    if (true && false) {
        for fi11 := 0; fi11 < 11; fi11++ {
            var v12 uint32 = 1
            switch "hello" {
            case "a":
                for fi13 := 0; fi13 < 11; fi13++ {
                    v1 = -1
                    _ = fi13
                }
            case "foo bar":
                if false {
                    v14 := GenS1{F0: 2, F1: 34}
                    _ = v14
                } else {
                    for fi15 := 0; fi15 < 17; fi15++ {
                        v16 := GenS1{F0: 1, F1: 1}
                        switch fi11 {
                        case -919:
                            if (v1 != v1) {
                                if (510 >= 65535) {
                                    fmt.Println(fi11, v12, v1)
                                    switch "hello" {
                                    case "hello":
                                        v17 := GenS1{F0: 18446744073709551615, F1: -77}
                                        _ = v17
                                    case "foo bar":
                                        fmt.Println(v9)
                                    }
                                    v9 = (^(v10))
                                } else {
                                    switch fi11 {
                                    case 42:
                                        var v18 float64 = 1.5
                                        _ = v18
                                    case -1:
                                        var v19 float32 = float32(float64(float32(((((-2.25 - (3.14159 - (1e10 - float64((1e10 * (1e10 + ((1e10 - float32((0.0 + (0.0 - ((1e10 - -2.25) / 3.14159))))) * (float32(1.5) / 0.1)))))))) * 3.14159) / 1.5) * 3.14159))))
                                        _ = v19
                                    }
                                }
                                fmt.Println(fi15, v1, fi15)
                            } else {
                                v1 -= -1
                                v9 += (v9 % (((v9 & v9)) | 1))
                                switch "hello" {
                                case "hello":
                                    var v20 uint32 = 2147483647
                                    _ = v20
                                }
                            }
                        case -2147483648:
                            v3 = append(v3, fi11)
                        }
                        var v21 int16 = 0
                        _ = fi15
                        _ = v16
                        _ = v21
                    }
                }
            }
            var v22 bool = true
            _ = fi11
            _ = v12
            _ = v22
        }
        fmt.Println(v9)
    } else {
        _ = len(v2)
        v10 = (v9 * 463)
        var v23 int = -9223372036854775808
        _ = v23
    }
    switch "hello" {
    case "\n\t":
        for fi24 := 0; fi24 < 19; fi24++ {
            for fi25 := 0; fi25 < 19; fi25++ {
                _ = len(v2)
                v26 := GenS1{F0: 125, F1: 1}
                _ = fi25
                _ = v26
            }
            _ = fi24
        }
    case "a":
        var v27 int16 = 1000
        _ = v27
    default:
        fmt.Println(v1, v10)
    }
    v28 := v2[("a" + ("a" + (("hello" + "") + "foo bar")))]
    if false {
        v10 = (v28 & (v28 - (v9 ^ v9)))
        switch ((((("a" + (("\n\t" + "a") + "\n\t")) + "\n\t") + "a") + "") + "hello") {
        case "a":
            v29 := GenS1{F0: 4294967295, F1: 0}
            _ = v29
        case "foo bar":
            v9 = 4294967295
        case "":
            var v30 int32 = genFn0(549, 0)
            _ = v30
        }
        fmt.Println(v28, v9, v28)
    } else {
        v31 := GenS1{F0: 4, F1: -1}
        _ = v31
    }
    v32 := GenS1{F0: 686, F1: 1}
    fmt.Println(v10)
    fmt.Println(v28, v1, v28)
    if false {
        fmt.Println(v28, v28)
    }
    var v33 uint16 = 32767
    var v34 uint = 1
    v35 := GenS1{F0: 0, F1: 100}
    var v36 bool = ((false && (true || false)) || (int16(v33) >= 42))
    v33 = 255
    fmt.Println(v36, v34)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v28
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v9
}
