package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32 }
type GenS2 struct { F0 uint32; F1 int32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint16; F1 string }
type GenS5 struct { F0 string; F1 uint; F2 uint8 }
func genFn0() uint16 {
    v4 := GenS1{F0: "a", F1: 207}
    v5 := []int{32768}
    v6 := GenS1{F0: ("foo bar" + ("hello" + ("a" + ""))), F1: 2}
    for fi7 := 0; fi7 < 16; fi7++ {
        var v8 uint16 = 1
        fmt.Println(fi7)
        switch ("a" + "a") {
        case "":
            var v9 uint16 = (^(v8))
            _ = v9
        case "a":
            if (1 < -115) {
                v8 += v8
                var v10 uint16 = 1
                _ = v10
            }
        case "hello":
            var v11 int = (-(fi7))
            _ = v11
        }
        _ = fi7
        _ = v8
    }
    _ = v4
    _ = v5
    _ = v6
    return 42
}

func main() {
    v1 := make(map[uint32]int)
    var v2 float64 = 0.0
    v3 := []string{(("hello" + ((("" + "hello") + "a") + "\n\t")) + ("a" + "foo bar")), ((("hello" + ("hello" + "foo bar")) + ("" + "a")) + "\n\t"), "foo bar"}
    fmt.Println(v2)
    var v12 int = -241
    v2 -= (float64(((1.5 + ((-2.25 * 0.0) / 3.14159)) - 0.0)) * v2)
    for fi13 := 0; fi13 < 12; fi13++ {
        for fi14 := 0; fi14 < 12; fi14++ {
            fmt.Println(v12, fi14)
            v15 := GenS2{F0: 510, F1: 0}
            _ = fi14
            _ = v15
        }
        var v16 int16 = 1000
        for fi17 := 0; fi17 < 10; fi17++ {
            switch ("a" + "\n\t") {
            case "hello":
                var v18 uint32 = 406
                _ = v18
            case "":
                switch (("" + "a") + ("\n\t" + "")) {
                case "\n\t":
                    if (("a" + "\n\t") > "") {
                        var v19 int8 = 1
                        if (!true) {
                            v12 -= 32767
                            var v20 int16 = (v16 * int16(v19))
                            _ = v20
                        } else {
                            v21 := GenS2{F0: 65535, F1: -1}
                            _ = v21
                        }
                        v22 := GenS1{F0: "\n\t", F1: 181}
                        _ = v19
                        _ = v22
                    } else {
                        var v23 float64 = 3.14159
                        _ = v23
                    }
                default:
                    var v24 float64 = float64(((-2.25 - 1.5) + 3.14159))
                    _ = v24
                }
                var v25 int32 = -222
                _ = v25
            case "foo bar":
                var v26 bool = (int8(v16) > -1)
                _ = v26
            }
            v27 := v1[0]
            fmt.Println(v16, v27, fi17)
            _ = fi17
            _ = v27
        }
        _ = fi13
        _ = v16
    }
    for fi28 := 0; fi28 < 7; fi28++ {
        v29 := GenS3{F0: 184}
        v30 := GenS4{F0: 545, F1: ("" + "a")}
        _ = fi28
        _ = v29
        _ = v30
    }
    v12 += v12
    fmt.Println(v12)
    switch (v12 - -9223372036854775808) {
    case 128:
        if false {
            var v31 uint = 259
            _ = v31
        } else {
            if (2147483648 >= 4294967295) {
                v2 = (0.0 * (v2 / v2))
                switch (("\n\t" + "") + "hello") {
                case "foo bar":
                    if false {
                        _ = len(v1)
                        v2 = float64(float32(0.0))
                    } else {
                        if true {
                            v2 += v2
                        } else {
                            v2 -= float64(1.5)
                        }
                        var v32 float32 = 0.0
                        fmt.Println(v2)
                        _ = v32
                    }
                default:
                    var v33 int16 = 0
                    _ = v33
                }
            } else {
                v34 := GenS1{F0: "\n\t", F1: 2}
                _ = v34
            }
            if (0 <= 4294967295) {
                v35 := GenS5{F0: ("foo bar" + ("foo bar" + ("\n\t" + ("" + "foo bar")))), F1: 4294967295, F2: 0}
                fmt.Println(v2, v12)
                var v36 float32 = 3.14159
                _ = v35
                _ = v36
            }
        }
    case 4294967295:
        var v37 uint32 = 2147483648
        _ = v37
    default:
        var v38 int = 0
        _ = v38
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
}
