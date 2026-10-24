package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 int8; F1 int64; F2 uint64 }
type GenS3 struct { F0 int16 }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[uint32]bool)
    var v3 uint32 = 2147483648
    var v4 float32 = 0.1
    var v5 string = "hello"
    v6 := make(map[uint32]string)
    v3 += (v3 | v3)
    switch 32768 {
    case -388:
        if true {
            var v7 uint8 = 255
            _ = v7
        } else {
            v3 = (v3 * 890)
        }
    }
    v3 = (^(v3))
    v4 += float32((((-2.25 * 0.1) * (0.0 / 3.14159)) + 0.1))
    if (!(674 > 1)) {
        v8 := GenS1{F0: true, F1: 18446744073709551615}
        var v9 uint16 = 0
        _ = v8
        _ = v9
    } else {
        v10 := []int{214}
        v10 = append(v10, 1)
        _ = v10
    }
    fmt.Println(v3)
    for fi11 := 0; fi11 < 7; fi11++ {
        switch (v5 + v5) {
        case "hello":
            v12 := GenS1{F0: (3.14159 > 3.14159), F1: 65535}
            _ = v12
        }
        _ = fi11
    }
    for fi13 := 0; fi13 < 1; fi13++ {
        v14 := GenS1{F0: (!true), F1: 228}
        _ = fi13
        _ = v14
    }
    fmt.Println(v5, v4, v4)
    v3 = (v3 | v3)
    var v15 uint16 = 725
    for fi16 := 0; fi16 < 7; fi16++ {
        v17 := GenS2{F0: 127, F1: -1, F2: 216}
        v18 := []int{}
        fmt.Println(v15)
        _ = fi16
        _ = v17
        _ = v18
    }
    switch -622 {
    case 2147483647:
        v3 = (v3 >> ((255) & 63))
    case 9223372036854775807:
        var v19 bool = (true && false)
        _ = v19
    case 127:
        if true {
            v20 := []int{-1, -9223372036854775808, -723, 4294967295}
            _ = v20
        } else {
            _ = len(v1)
            if (false || true) {
                var v21 uint32 = v3
                _ = v21
            } else {
                v22 := []int{-775, 4294967295, -1, 42}
                for fi23 := 0; fi23 < 10; fi23++ {
                    v24 := GenS3{F0: 0}
                    _ = fi23
                    _ = v24
                }
                _ = v22
            }
        }
    default:
        v4 -= (3.14159 + ((float32(float64((v4 - (v4 * float32(1e10))))) - v4) * ((v4 * v4) / v4)))
    }
    if true {
        if false {
            var v25 uint32 = v3
            _ = v25
        } else {
            switch "\n\t" {
            case "\n\t":
                v26 := GenS2{F0: 20, F1: 1, F2: 18446744073709551615}
                _ = v26
            case "foo bar":
                for fi27 := 0; fi27 < 9; fi27++ {
                    v28 := []int{(fi27 | (fi27 << ((127) & 63)))}
                    for fi29 := 0; fi29 < 19; fi29++ {
                        var v30 string = ""
                        var v31 uint16 = 649
                        if false {
                            if (0.1 == 1e10) {
                                if true {
                                    v28 = append(v28, (fi29 + fi29))
                                    v28 = append(v28, fi27)
                                } else {
                                    v15 += 588
                                    if (18446744073709551615 > 0) {
                                        var v32 uint = 248
                                        v4 += 1.5
                                        v1[v3] = (v3 - (v3 ^ v3))
                                        v28 = append(v28, (^(fi29)))
                                        v3 = (v3 - 2147483647)
                                        _ = v32
                                    }
                                }
                            } else {
                                if true {
                                    var v33 float64 = 0.0
                                    _ = v33
                                } else {
                                    v5 = ((v30 + ((v30 + ((v5 + (("" + v5) + "")) + (((("foo bar" + ((((v30 + v30) + (((v5 + (v5 + ("a" + v30))) + "hello") + "\n\t")) + "\n\t") + "")) + v5) + ("hello" + v5)) + "foo bar"))) + "foo bar")) + "hello")
                                    var v34 uint32 = 927
                                    var v35 float32 = v4
                                    _ = v34
                                    _ = v35
                                }
                            }
                        }
                        _ = fi29
                        _ = v30
                        _ = v31
                    }
                    _ = fi27
                    _ = v28
                }
            case "a":
                var v36 int64 = -1
                _ = v36
            default:
                switch "foo bar" {
                case "a":
                    var v37 uint = 42
                    _ = v37
                case "foo bar":
                    v3 = v3
                case "hello":
                    fmt.Println(v4)
                }
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
