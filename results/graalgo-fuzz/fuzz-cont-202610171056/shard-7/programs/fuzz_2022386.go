package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16; F1 string; F2 uint16 }
type GenS4 struct { F0 uint8; F1 int16; F2 int64 }
type GenS5 struct { F0 int8 }

func main() {
    v1 := make(map[int]bool)
    v2 := []string{"a", ""}
    var v3 uint = 1
    var v4 string = ("hello" + ("" + "\n\t"))
    v5 := make(map[int]bool)
    var v6 uint16 = 507
    _ = len(v1)
    for fi7 := 0; fi7 < 1; fi7++ {
        switch (fi7 >> ((0) & 63)) {
        case -9223372036854775808:
            var v8 uint32 = 2147483647
            _ = v8
        case 32768:
            fmt.Println(v3, v3, v6)
        case 197:
            if false {
                fmt.Println(v4, v3, v6)
            }
        }
        var v9 uint16 = v6
        _ = fi7
        _ = v9
    }
    var v10 uint64 = 978
    fmt.Println(v3)
    delete(v1, 0)
    v6 = 0
    var v11 int = -1
    for fi12 := 0; fi12 < 1; fi12++ {
        if true {
            switch v4 {
            case "a":
                v3 = (v3 * (v3 % (((v3 >> ((v3) & 63))) | 1)))
            case "\n\t":
                if false {
                    fmt.Println(v6, v10)
                } else {
                    switch (fi12 % ((v11) | 1)) {
                    case -645:
                        v13 := GenS1{F0: -1, F1: true, F2: (false && (-2.25 <= 1e10))}
                        _ = v13
                    case 0:
                        fmt.Println(v10, v10, fi12)
                    default:
                        if false {
                            if (2147483647 <= 55) {
                                v11 -= v11
                                var v14 uint64 = v10
                                if (true && (-32768 == 1000)) {
                                    var v15 uint8 = 2
                                    v16 := GenS2{F0: v15}
                                    v17 := GenS3{F0: 65535, F1: v4, F2: v6}
                                    _ = v15
                                    _ = v16
                                    _ = v17
                                } else {
                                    var v18 uint32 = 2
                                    _ = v18
                                }
                                _ = v14
                            } else {
                                var v19 string = (((v4 + "") + "") + ("foo bar" + v4))
                                for fi20 := 0; fi20 < 20; fi20++ {
                                    v21 := GenS3{F0: 961, F1: (v19 + v4), F2: 2}
                                    v22 := GenS1{F0: 1, F1: false, F2: false}
                                    v11 = v11
                                    _ = fi20
                                    _ = v21
                                    _ = v22
                                }
                                _ = v19
                            }
                        }
                    }
                }
            case "":
                v3 -= 697
            }
            v23 := GenS4{F0: 32, F1: -1, F2: 196}
            _ = v23
        } else {
            if true {
                if len(v2) > 0 {
                    v2[0] = (v4 + v4)
                }
            } else {
                v3 = v3
                var v24 uint8 = 113
                for fi25 := 0; fi25 < 5; fi25++ {
                    v10 = v10
                    delete(v1, int(v24))
                    _ = len(v2)
                    _ = fi25
                }
                _ = v24
            }
            fmt.Println(v11, v11)
            switch 9223372036854775807 {
            case -908:
                if true {
                    var v26 float32 = ((((3.14159 + ((3.14159 - 0.1) * float32((-2.25 / 1.5)))) + ((3.14159 - -2.25) + (float32(1e10) * (0.0 + -2.25)))) * 0.0) + 1.5)
                    v27 := GenS4{F0: 127, F1: int16(v10), F2: -1}
                    _ = v26
                    _ = v27
                }
            case 1:
                var v28 uint16 = 365
                _ = v28
            }
        }
        _ = len(v2)
        v29 := GenS5{F0: -128}
        _ = fi12
        _ = v29
    }
    if false {
        var v30 int64 = -1
        v11 = v11
        _ = v30
    }
    var v31 int8 = 32
    v4 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
