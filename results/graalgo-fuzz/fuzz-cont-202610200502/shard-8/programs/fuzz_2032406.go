package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int; F2 uint }
type GenS2 struct { F0 string; F1 int; F2 uint8 }
type GenS3 struct { F0 uint64 }

func main() {
    var v1 int32 = 1
    v2 := make(map[int]int)
    v3 := []int{-9223372036854775808, 0, 128}
    fmt.Println(v1)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v2[-9223372036854775808] = -1
    if true {
        if true {
            for fi4 := 0; fi4 < 1; fi4++ {
                v1 -= (v1 ^ v1)
                _ = fi4
            }
        } else {
            fmt.Println(v1)
            var v5 int = 128
            v3 = append(v3, 127)
            _ = v5
        }
    } else {
        v1 = (^(v1))
        fmt.Println(v1)
        var v6 int16 = 32767
        _ = v6
    }
    if (!(true && true)) {
        v7 := GenS1{F0: (v1 >> ((42) & 63)), F1: -184, F2: 4294967295}
        switch "hello" {
        case "\n\t":
            if false {
                var v8 int = -120
                if (-2.25 <= (1.5 / (1.5 / 1e10))) {
                    for fi9 := 0; fi9 < 17; fi9++ {
                        if (!false) {
                            if (42 != -128) {
                                _ = len(v3)
                            }
                        }
                        var v10 bool = true
                        _ = fi9
                        _ = v10
                    }
                }
                var v11 uint64 = 543
                _ = v11
                _ = v8
            } else {
                v1 = v1
                var v12 float64 = 3.14159
                v3 = append(v3, 172)
                _ = v12
            }
        }
        var v13 string = "a"
        _ = v13
        _ = v7
    } else {
        var v14 int = 32768
        _ = v14
    }
    v15 := GenS1{F0: -1, F1: 32768, F2: 255}
    var v16 uint16 = 255
    for fi17 := 0; fi17 < 9; fi17++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi17
    }
    var v18 uint8 = 173
    var v19 int32 = 2147483647
    v19 = -1
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if false {
        fmt.Println(v19)
        v20 := v2[-2147483648]
        for fi21 := 0; fi21 < 14; fi21++ {
            for fi22 := 0; fi22 < 10; fi22++ {
                var v23 int64 = -9223372036854775808
                if true {
                    var v24 string = (("\n\t" + "foo bar") + ("a" + "foo bar"))
                    _ = v24
                } else {
                    var v25 int8 = -128
                    switch (("\n\t" + "a") + "\n\t") {
                    case "a":
                        if true {
                            fmt.Println(v16, v19)
                            v25 = v25
                            switch "" {
                            case "\n\t":
                                var v26 uint64 = 18446744073709551615
                                _ = v26
                            }
                        }
                    case "\n\t":
                        v27 := GenS2{F0: ("" + "a"), F1: fi22, F2: (v18 + 127)}
                        _ = v27
                    default:
                        if true {
                            v28 := GenS3{F0: 4294967295}
                            _ = v28
                        } else {
                            var v29 int = -2147483648
                            if ((true && true) && true) {
                                var v30 int16 = -1
                                _ = v30
                            } else {
                                v31 := GenS2{F0: ("" + ""), F1: 32767, F2: v18}
                                _ = v31
                            }
                            if true {
                                v19 = v1
                            }
                            _ = v29
                        }
                    }
                    var v32 int16 = -32768
                    _ = v25
                    _ = v32
                }
                var v33 int = 199
                var v34 uint32 = 42
                v2[128] = v33
                _ = fi22
                _ = v23
                _ = v33
                _ = v34
            }
            var v35 uint8 = (v18 - (v18 + (v18 * 255)))
            _ = fi21
            _ = v35
        }
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v3
}
