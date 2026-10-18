package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 int }
type GenS2 struct { F0 uint32; F1 string; F2 int8 }
type GenS3 struct { F0 int; F1 string }
type GenS4 struct { F0 int; F1 uint16 }

func main() {
    var v1 float64 = ((1.5 / (3.14159 - (1.5 * float64(1.5)))) - 1.5)
    v2 := []int{2147483647, 42, 127, -540}
    var v3 bool = (128 > 2)
    v4 := make(map[int]int)
    v5 := []string{"a", "a", ("foo bar" + ("" + "\n\t")), (("" + "") + "a")}
    var v6 int8 = -65
    switch 560 {
    case 42:
        for fi7 := 0; fi7 < 16; fi7++ {
            for fi8 := 0; fi8 < 1; fi8++ {
                var v9 float32 = (0.1 * 0.1)
                var v10 bool = (!v3)
                fmt.Println(v3, v10, v10)
                if v3 {
                    fmt.Println(fi7)
                    v9 = ((v9 / 3.14159) + v9)
                }
                var v11 uint32 = 399
                _ = fi8
                _ = v10
                _ = v11
                _ = v9
            }
            fmt.Println(v3)
            var v12 int8 = 53
            _ = fi7
            _ = v12
        }
    case -1:
        fmt.Println(v3, v1, v6)
    case 2147483647:
        if len(v5) > 0 {
            v5[0] = ("hello" + "\n\t")
        }
    default:
        var v13 int64 = 42
        _ = v13
    }
    if ((!v3) && ((!v3) && v3)) {
        v3 = (v3 && (127 != 0))
        v14 := GenS1{F0: 0, F1: -128, F2: 230}
        _ = v14
    } else {
        switch ("a" + "") {
        case "a":
            if (false || (!(v3 && v3))) {
                var v15 int = 710
                _ = v15
            }
        case "hello":
            v16 := GenS1{F0: 0, F1: (v6 << ((18446744073709551615) & 63)), F2: 42}
            _ = v16
        }
        delete(v4, 127)
        for fi17 := 0; fi17 < 9; fi17++ {
            var v18 int64 = 1
            v4[-668] = fi17
            _ = fi17
            _ = v18
        }
    }
    for fi19 := 0; fi19 < 10; fi19++ {
        v1 = (1.5 + 0.0)
        v20 := v4[9223372036854775807]
        _ = fi19
        _ = v20
    }
    v21 := GenS2{F0: 2147483647, F1: "\n\t", F2: -1}
    var v22 int8 = v6
    if v3 {
        fmt.Println(v6, v6)
    }
    var v23 int64 = -9223372036854775808
    if (!(false || (42 < -9223372036854775808))) {
        if v3 {
            var v24 float64 = 1.5
            if false {
                v25 := GenS3{F0: 2147483647, F1: "foo bar"}
                v3 = (true || (v3 || (!(255 != 0))))
                v26 := GenS4{F0: 489, F1: 32767}
                _ = v25
                _ = v26
            } else {
                v27 := GenS3{F0: -9223372036854775808, F1: ""}
                v22 = 1
                var v28 string = ""
                _ = v27
                _ = v28
            }
            _ = v24
        } else {
            switch 9223372036854775807 {
            case -961:
                _ = len(v4)
            default:
                if true {
                    if ((v3 || v3) || ("hello" >= (("" + ("foo bar" + ("hello" + ""))) + "foo bar"))) {
                        fmt.Println(v3, v6)
                        if (42 >= 397) {
                            if (v3 || (0 > 42)) {
                                for fi29 := 0; fi29 < 8; fi29++ {
                                    var v30 uint16 = 65535
                                    _ = fi29
                                    _ = v30
                                }
                                var v31 int16 = -1
                                _ = v31
                            } else {
                                var v32 float64 = (v1 - v1)
                                _ = v32
                            }
                        }
                    }
                    var v33 bool = (255 > 0)
                    _ = v33
                }
            }
        }
        if len(v5) > 0 {
            v5[0] = "hello"
        }
    }
    var v34 uint16 = 255
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
