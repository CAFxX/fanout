package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint32; F1 int8; F2 uint64 }

func main() {
    v1 := make(map[string]bool)
    v2 := []string{"foo bar", "a"}
    var v3 uint16 = 42
    var v4 uint16 = v3
    var v5 int64 = 42
    v6 := make(map[int]bool)
    v4 = v3
    if true {
        var v7 uint8 = 127
        fmt.Println(v4)
        _ = v7
    } else {
        fmt.Println(v3, v5)
        if true {
            v2 = append(v2, "a")
        }
        var v8 uint32 = 2147483647
        _ = v8
    }
    _ = len(v2)
    for fi9 := 0; fi9 < 7; fi9++ {
        v10 := GenS1{F0: 356}
        _ = fi9
        _ = v10
    }
    var v11 uint32 = 289
    var v12 uint64 = 809
    for fi13 := 0; fi13 < 12; fi13++ {
        for fi14 := 0; fi14 < 1; fi14++ {
            if (-643 == -1) {
                v1[(("" + "foo bar") + "\n\t")] = ((127 <= 255) && (v5 > (-(v5))))
                switch (fi13 | (fi14 & (fi14 | (fi14 - fi14)))) {
                case 32768:
                    if true {
                        var v15 int32 = -86
                        v12 = (v12 / ((440) | 1))
                        v15 -= v15
                        _ = v15
                    } else {
                        v16 := GenS1{F0: 514}
                        var v17 uint32 = v11
                        v17 = 4294967295
                        v1["a"] = ((124 == 127) && true)
                        _ = v16
                        _ = v17
                    }
                default:
                    v12 -= 0
                    if true {
                        _ = len(v2)
                        v1[("hello" + "foo bar")] = true
                        v3 = 42
                        v3 -= 442
                    }
                }
                v12 -= v12
            }
            _ = fi14
        }
        _ = fi13
    }
    if ((3.14159 == (1.5 + (1e10 * 0.1))) && false) {
        v18 := GenS1{F0: 42}
        v3 = v4
        _ = len(v2)
        _ = v18
    } else {
        switch "a" {
        case "hello":
            if false {
                if len(v2) > 0 {
                    v2[0] = "\n\t"
                }
            } else {
                for fi19 := 0; fi19 < 14; fi19++ {
                    v2 = append(v2, ("foo bar" + ("foo bar" + ("hello" + ("a" + "hello")))))
                    v20 := GenS1{F0: 18446744073709551615}
                    switch fi19 {
                    case 1:
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                    _ = fi19
                    _ = v20
                }
                if false {
                    v5 -= v5
                }
            }
        case "\n\t":
            for fi21 := 0; fi21 < 17; fi21++ {
                v4 = (v4 - v3)
                v22 := GenS2{F0: "\n\t"}
                v23 := GenS3{F0: v11, F1: -78, F2: (v12 ^ v12)}
                _ = fi21
                _ = v22
                _ = v23
            }
        case "":
            fmt.Println(v12)
        }
        var v24 uint32 = v11
        for fi25 := 0; fi25 < 17; fi25++ {
            if (!(true || false)) {
                if len(v2) > 0 {
                    v2[0] = "foo bar"
                }
                switch ("hello" + "hello") {
                case "\n\t":
                    if (v5 > (v5 % (((v5 * (-(v5)))) | 1))) {
                        var v26 float32 = 3.14159
                        if (!(!(v12 < v12))) {
                            var v27 int64 = v5
                            if false {
                                var v28 int64 = -347
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                v27 += int64(v24)
                                _ = v28
                            }
                            _ = v27
                        }
                        var v29 int8 = 25
                        _ = v26
                        _ = v29
                    }
                case "":
                    v11 -= v11
                case "foo bar":
                    v24 = v11
                }
            }
            for fi30 := 0; fi30 < 12; fi30++ {
                var v31 int8 = 0
                var v32 bool = false
                v12 = (v12 | 4294967295)
                _ = fi30
                _ = v31
                _ = v32
            }
            _ = fi25
        }
        _ = v24
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
