package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]bool)
    v3 := make(map[int]bool)
    v4 := []int{}
    var v5 int16 = -1
    v5 = v5
    for fi6 := 0; fi6 < 3; fi6++ {
        if (-2.25 != -2.25) {
            var v7 int64 = -9223372036854775808
            v5 = v5
            if true {
                v5 += 1000
                for fi8 := 0; fi8 < 2; fi8++ {
                    var v9 int = fi8
                    _ = fi8
                    _ = v9
                }
                v5 = v5
            }
            _ = v7
        }
        var v10 float32 = -2.25
        _ = fi6
        _ = v10
    }
    if (-2.25 >= (1.5 / 0.1)) {
        v5 = 704
        if false {
            if false {
                v5 += 803
                v4 = append(v4, 1)
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            } else {
                _ = len(v1)
                fmt.Println(v5)
                if true {
                    var v11 uint = 65535
                    _ = v11
                }
            }
            switch -264 {
            case -696:
                var v12 uint = 612
                _ = v12
            case -2147483648:
                var v13 bool = (("" >= "hello") || true)
                _ = v13
            case -568:
                _ = len(v4)
            default:
                if false {
                    v5 = 1
                } else {
                    fmt.Println(v5)
                    var v14 float64 = float64((float32(3.14159) - (float32(0.1) + (0.0 / 1.0))))
                    for fi15 := 0; fi15 < 10; fi15++ {
                        v5 = 1000
                        _ = fi15
                    }
                    _ = v14
                }
            }
        } else {
            switch (("" + "hello") + ("hello" + "a")) {
            case "hello":
                var v16 uint8 = 2
                _ = v16
            case "":
                _ = len(v1)
            case "\n\t":
                if (true || true) {
                    _ = len(v3)
                    if false {
                        _ = len(v4)
                    } else {
                        v5 = (v5 - v5)
                    }
                } else {
                    for fi17 := 0; fi17 < 5; fi17++ {
                        var v18 float64 = (1e10 - 3.14159)
                        switch (("" + ("\n\t" + (("a" + "hello") + (("hello" + "foo bar") + "a")))) + "a") {
                        case "\n\t":
                            _ = len(v1)
                        }
                        for fi19 := 0; fi19 < 18; fi19++ {
                            var v20 int64 = -1
                            if (true || (385 > 838)) {
                                var v21 uint = 42
                                var v22 uint = 334
                                if true {
                                    _ = len(v1)
                                    v18 = v18
                                    var v23 float32 = (float32(float64(float32(v18))) / 1.5)
                                    _ = v23
                                }
                                _ = v21
                                _ = v22
                            }
                            _ = fi19
                            _ = v20
                        }
                        _ = fi17
                        _ = v18
                    }
                    _ = len(v2)
                }
            }
        }
        v5 -= v5
    } else {
        v24 := GenS1{F0: 2147483647, F1: 9223372036854775807}
        v5 -= (v5 * -462)
        var v25 uint64 = 2
        _ = v24
        _ = v25
    }
    v26 := GenS1{F0: 127, F1: -1}
    v27 := GenS1{F0: -1, F1: -228}
    if len(v4) > 0 {
        v4[0] = 32768
    }
    v5 += 452
    var v28 uint = 42
    fmt.Println(v5, v28)
    var v29 uint16 = 582
    v30 := GenS1{F0: 600, F1: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
