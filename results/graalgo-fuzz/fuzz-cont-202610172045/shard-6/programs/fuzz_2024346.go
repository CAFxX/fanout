package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 int64; F1 int32 }

func main() {
    var v1 uint32 = 4294967295
    v2 := make(map[string]bool)
    var v3 float32 = (-2.25 - (float32(float64(0.0)) - 1e10))
    var v4 uint32 = v1
    for fi5 := 0; fi5 < 8; fi5++ {
        v6 := GenS1{F0: 989}
        if (127 == 100) {
            v7 := GenS1{F0: 0}
            v3 = v3
            var v8 string = "foo bar"
            _ = v7
            _ = v8
        }
        if true {
            v3 = v3
            v3 = v3
            fmt.Println(fi5)
        } else {
            v9 := GenS1{F0: -1}
            for fi10 := 0; fi10 < 11; fi10++ {
                switch (("foo bar" + "\n\t") + "") {
                case "foo bar":
                    if true {
                        v11 := []int{fi5, (fi5 / (((fi10 % ((42) | 1))) | 1)), fi10}
                        switch ((("foo bar" + ("" + ("a" + "\n\t"))) + "\n\t") + ("" + "")) {
                        case "a":
                            if true {
                                fmt.Println(v1, v4)
                                var v12 bool = true
                                _ = v12
                            } else {
                                _ = len(v2)
                                if len(v11) >= 2 {
                                    v11 = v11[1:len(v11)]
                                }
                                v13 := v2[""]
                                _ = v13
                            }
                        case "hello":
                            switch (fi10 * fi5) {
                            case 127:
                                var v14 int32 = 1
                                _ = v14
                            case -2147483648:
                                fmt.Println(v1)
                            }
                        }
                        _ = v11
                    }
                case "a":
                    switch -10 {
                    case 4294967295:
                        fmt.Println(fi10, v1)
                    case 32767:
                        v15 := GenS2{F0: 2}
                        _ = v15
                    case -1:
                        _ = len(v2)
                    default:
                        var v16 uint16 = 115
                        _ = v16
                    }
                case "hello":
                    v4 = (v4 ^ 969)
                }
                var v17 uint8 = 42
                _ = fi10
                _ = v17
            }
            _ = v9
        }
        _ = fi5
        _ = v6
    }
    fmt.Println(v4, v1, v1)
    for fi18 := 0; fi18 < 9; fi18++ {
        delete(v2, ("\n\t" + "\n\t"))
        v1 = 0
        _ = fi18
    }
    for fi19 := 0; fi19 < 9; fi19++ {
        fmt.Println(v3, v4)
        v20 := GenS2{F0: v4}
        var v21 uint32 = v4
        _ = fi19
        _ = v20
        _ = v21
    }
    switch "" {
    case "foo bar":
        if (false && (true || (!(1 == 127)))) {
            v4 = v1
            for fi22 := 0; fi22 < 7; fi22++ {
                v4 += v4
                _ = fi22
            }
        } else {
            var v23 bool = true
            for fi24 := 0; fi24 < 5; fi24++ {
                if v23 {
                    fmt.Println(v23, v3, v23)
                } else {
                    v25 := GenS3{F0: 1}
                    switch (^(fi24)) {
                    case 4294967295:
                        v26 := GenS4{F0: 621, F1: 2147483647}
                        _ = v26
                    case -9223372036854775808:
                        fmt.Println(fi24, v1, fi24)
                    case 32768:
                        v23 = (0 >= 2)
                    default:
                        for fi27 := 0; fi27 < 7; fi27++ {
                            v23 = ((((!(false || v23)) || v23) || v23) && (v4 > 696))
                            _ = fi27
                        }
                    }
                    _ = v25
                }
                v23 = false
                var v28 float64 = (float64((v3 * (v3 + (3.14159 + 1.5)))) - -2.25)
                _ = fi24
                _ = v28
            }
            v23 = (1 >= -1)
            _ = v23
        }
    case "a":
        for fi29 := 0; fi29 < 5; fi29++ {
            switch 128 {
            case -1:
                switch 775 {
                case 9223372036854775807:
                    if ((!false) || true) {
                        v3 += v3
                    }
                case -2147483648:
                    for fi30 := 0; fi30 < 12; fi30++ {
                        v31 := []int{(fi29 / ((fi29) | 1)), 1, fi30, fi30}
                        _ = fi30
                        _ = v31
                    }
                case 338:
                    var v32 float32 = 0.0
                    _ = v32
                default:
                    if false {
                        var v33 uint64 = 2
                        var v34 int64 = -884
                        _ = v33
                        _ = v34
                    }
                }
            case 1:
                switch "\n\t" {
                case "\n\t":
                    v1 = 790
                }
            case -215:
                v2["a"] = true
            default:
                var v35 int8 = -128
                _ = v35
            }
            v4 -= v1
            _ = len(v2)
            _ = fi29
        }
    case "hello":
        for fi36 := 0; fi36 < 12; fi36++ {
            v3 = float32(((3.14159 - 3.14159) - (1e10 * 1e10)))
            v3 = v3
            _ = fi36
        }
    default:
        fmt.Println(v3, v3, v4)
    }
    _ = len(v2)
    var v37 uint8 = 0
    if false {
        for fi38 := 0; fi38 < 5; fi38++ {
            var v39 uint16 = 2
            var v40 float64 = 0.1
            _ = fi38
            _ = v39
            _ = v40
        }
        var v41 int16 = 964
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
}
