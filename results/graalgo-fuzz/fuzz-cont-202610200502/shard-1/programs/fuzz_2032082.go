package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 int16 }

func main() {
    var v1 uint = 18446744073709551615
    var v2 float32 = 0.1
    var v3 int64 = 9223372036854775807
    v4 := []int{}
    if false {
        if len(v4) > 0 {
            v4[0] = 1
        }
    }
    var v5 uint64 = 42
    var v6 float64 = (1e10 + (float64(v2) - ((1e10 - float64(v2)) * (1e10 / -2.25))))
    if true {
        v7 := GenS1{F0: (("hello" + "a") + "\n\t"), F1: 186, F2: -32768}
        switch -215 {
        case -2147483648:
            if (!false) {
                var v8 float32 = v2
                _ = v8
            }
        case 1:
            for fi9 := 0; fi9 < 1; fi9++ {
                for fi10 := 0; fi10 < 3; fi10++ {
                    fmt.Println(v3, v2)
                    var v11 int8 = 42
                    fmt.Println(fi10)
                    _ = fi10
                    _ = v11
                }
                var v12 float64 = v6
                _ = len(v4)
                _ = fi9
                _ = v12
            }
        case 32767:
            if true {
                fmt.Println(v3, v6, v6)
                if ((((832 <= 65535) && (-126 < 1000)) || true) || (-800 >= 1)) {
                    for fi13 := 0; fi13 < 10; fi13++ {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        if true {
                            for fi14 := 0; fi14 < 4; fi14++ {
                                switch "hello" {
                                case "":
                                    var v15 uint32 = 2147483647
                                    _ = v15
                                case "\n\t":
                                    var v16 float64 = float64((-2.25 + 3.14159))
                                    _ = v16
                                case "foo bar":
                                    v1 -= (^(v1))
                                default:
                                    var v17 float64 = -2.25
                                    _ = v17
                                }
                                _ = fi14
                            }
                            var v18 int64 = -9223372036854775808
                            _ = v18
                        } else {
                            switch fi13 {
                            case -9223372036854775808:
                                v19 := GenS1{F0: "\n\t", F1: 88, F2: -1}
                                _ = v19
                            }
                            for fi20 := 0; fi20 < 7; fi20++ {
                                var v21 bool = false
                                v3 = (v3 ^ (v3 >> ((v1) & 63)))
                                fmt.Println(fi20)
                                _ = fi20
                                _ = v21
                            }
                        }
                        _ = fi13
                    }
                } else {
                    v2 -= ((v2 * v2) - (v2 * -2.25))
                    v1 = v1
                    for fi22 := 0; fi22 < 19; fi22++ {
                        v23 := GenS1{F0: "\n\t", F1: 82, F2: -838}
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        _ = fi22
                        _ = v23
                    }
                }
            } else {
                v24 := GenS1{F0: "", F1: 186, F2: 1}
                fmt.Println(v5, v6, v3)
                var v25 uint8 = 42
                _ = v24
                _ = v25
            }
        }
        _ = v7
    }
    v6 = (v6 + v6)
    v6 = float64(3.14159)
    v26 := GenS1{F0: (("\n\t" + ((("" + "foo bar") + ("\n\t" + "\n\t")) + "a")) + "\n\t"), F1: 173, F2: -277}
    for fi27 := 0; fi27 < 6; fi27++ {
        v1 -= v1
        _ = fi27
    }
    if true {
        switch ("a" + "a") {
        case "":
            if (1 < -19) {
                var v28 string = "foo bar"
                var v29 int32 = 1
                var v30 int8 = -128
                _ = v28
                _ = v29
                _ = v30
            }
        case "\n\t":
            var v31 uint = v1
            _ = v31
        case "hello":
            if false {
                var v32 uint16 = 848
                _ = v32
            } else {
                switch (("" + "a") + "\n\t") {
                case "hello":
                    var v33 bool = (-32768 <= 0)
                    _ = v33
                }
                v4 = append(v4, -2147483648)
                v2 -= v2
            }
        default:
            if (true || true) {
                v3 += (v3 | v3)
                if len(v4) > 0 {
                    v4[0] = 0
                }
            } else {
                fmt.Println(v5, v6)
                var v34 float64 = ((v6 / (v6 + 1.5)) - v6)
                var v35 uint16 = 42
                _ = v34
                _ = v35
            }
        }
        var v36 int8 = -77
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
