package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 uint64 }
type GenS2 struct { F0 int16; F1 uint; F2 int8 }
type GenS3 struct { F0 uint16 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]string)
    var v3 uint = 1
    v4 := make(map[string]string)
    fmt.Println(v3)
    v5 := []int{128, 4294967295}
    v3 = v3
    fmt.Println(v3)
    v1[9223372036854775807] = 0
    if false {
        if true {
            switch 128 {
            case 0:
                var v6 uint32 = 557
                _ = v6
            case 292:
                switch ("hello" + "foo bar") {
                case "foo bar":
                    v3 = (v3 - (v3 >> ((v3) & 63)))
                case "a":
                    v3 = (^(v3))
                case "":
                    var v7 float32 = 0.0
                    _ = v7
                }
            }
        }
    }
    if true {
        if (1 <= 1) {
            _ = len(v4)
            v3 = (^(v3))
            v8 := GenS1{F0: "foo bar", F1: 1, F2: 42}
            _ = v8
        }
    } else {
        v9 := GenS1{F0: ("\n\t" + ("hello" + ("" + ((("hello" + "") + ("hello" + "foo bar")) + "\n\t")))), F1: 9223372036854775807, F2: 4294967295}
        var v10 float64 = 0.0
        for fi11 := 0; fi11 < 20; fi11++ {
            switch (fi11 + 0) {
            case 78:
                if false {
                    if (true && false) {
                        for fi12 := 0; fi12 < 16; fi12++ {
                            var v13 int = fi12
                            _ = fi12
                            _ = v13
                        }
                        v3 -= v3
                    } else {
                        if true {
                            switch (fi11 & fi11) {
                            case 128:
                                v10 = 1.5
                            case 2147483647:
                                delete(v1, fi11)
                            default:
                                switch "hello" {
                                case "":
                                    fmt.Println(v3, v3)
                                case "a":
                                    fmt.Println(v10, fi11, v10)
                                case "hello":
                                    v3 = v3
                                default:
                                    var v14 int16 = 42
                                    _ = v14
                                }
                            }
                            var v15 int64 = 322
                            var v16 float64 = (((3.14159 + 1e10) * 3.14159) / 1e10)
                            _ = v15
                            _ = v16
                        } else {
                            var v17 uint64 = 4294967295
                            var v18 int8 = 1
                            var v19 uint32 = 71
                            _ = v17
                            _ = v18
                            _ = v19
                        }
                        v5 = append(v5, -2147483648)
                    }
                } else {
                    if (!(0.1 <= ((1.5 - 3.14159) * ((0.1 + 0.1) / 1.0)))) {
                        if ((v10 >= 0.1) || true) {
                            fmt.Println(v3, v10)
                        }
                        var v20 string = "\n\t"
                        fmt.Println(v3)
                        _ = v20
                    } else {
                        var v21 uint64 = 1
                        v22 := GenS2{F0: int16(fi11), F1: (^(v3)), F2: 42}
                        _ = v21
                        _ = v22
                    }
                }
            case 736:
                v23 := GenS3{F0: 275}
                _ = v23
            }
            _ = fi11
        }
        _ = v10
        _ = v9
    }
    var v24 string = "\n\t"
    v24 = (v24 + "foo bar")
    v24 = ((v24 + "hello") + "foo bar")
    v24 = ((v24 + (v24 + v24)) + v24)
    var v25 int = int(v3)
    if len(v5) > 0 {
        v5[0] = 4294967295
    }
    _ = len(v1)
    v2[472] = v24
    for fi26 := 0; fi26 < 4; fi26++ {
        v24 = v24
        _ = fi26
    }
    v27 := GenS3{F0: 759}
    var v28 uint16 = 32767
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
