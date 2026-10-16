package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 bool; F1 int64 }
type GenS3 struct { F0 uint8 }
func genFn0() uint {
    var v5 int8 = 1
    v6 := GenS1{F0: true}
    switch "foo bar" {
    case "a":
        v5 += v5
    case "":
        fmt.Println(v5)
    }
    _ = v5
    _ = v6
    return 4294967295
}

func main() {
    var v1 string = ""
    var v2 int8 = -38
    v3 := []string{(("a" + v1) + (v1 + "foo bar"))}
    v4 := []uint32{2147483647}
    var v7 int16 = 1000
    switch "" {
    case "foo bar":
        v3 = append(v3, v1)
    default:
        var v8 uint64 = 4294967295
        _ = v8
    }
    var v9 float32 = -2.25
    if true {
        v10 := GenS1{F0: false}
        for fi11 := 0; fi11 < 10; fi11++ {
            for fi12 := 0; fi12 < 4; fi12++ {
                switch ((v1 + v1) + ("hello" + v1)) {
                case "\n\t":
                    if false {
                        v13 := GenS1{F0: true}
                        v1 = (v1 + v1)
                        v14 := GenS1{F0: true}
                        _ = v13
                        _ = v14
                    } else {
                        if (0.0 <= (1.5 / 1.5)) {
                            v9 = v9
                            var v15 uint16 = 274
                            if false {
                                fmt.Println(v9)
                                var v16 uint = 4294967295
                                _ = v16
                            } else {
                                if true {
                                    v7 = v7
                                } else {
                                    var v17 uint8 = 1
                                    _ = v17
                                }
                                if false {
                                    if len(v3) > 0 {
                                        v3[0] = v1
                                    }
                                    v2 = -1
                                    var v18 string = v1
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    v18 = "a"
                                    _ = v18
                                }
                                v19 := GenS1{F0: true}
                                _ = v19
                            }
                            switch (fi11 + 13) {
                            case 557:
                                if ((("hello" + v1) <= (("hello" + v1) + v1)) || (!true)) {
                                    var v20 int16 = v7
                                    _ = v20
                                }
                            }
                            _ = v15
                        }
                        var v21 int8 = 82
                        _ = v21
                    }
                case "hello":
                    v22 := GenS2{F0: (2 != 651), F1: -9223372036854775808}
                    _ = v22
                default:
                    v23 := GenS1{F0: (false || true)}
                    _ = v23
                }
                fmt.Println(v2, v7)
                _ = fi12
            }
            var v24 bool = true
            var v25 uint32 = 42
            _ = fi11
            _ = v24
            _ = v25
        }
        _ = v10
    }
    v7 -= v7
    var v26 float64 = 1.5
    if true {
        fmt.Println(v26, v2, v9)
        var v27 uint = 4294967295
        v28 := GenS3{F0: 127}
        _ = v27
        _ = v28
    }
    if len(v4) > 0 {
        v4[0] = 0
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v7
    _ = v9
}
