package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 bool; F1 bool; F2 string }
type GenS4 struct { F0 uint32; F1 bool }
func genFn0(p0 int) int16 {
    v6 := []int{127, 32767}
    var v7 float64 = 0.1
    v7 = v7
    _ = p0
    _ = v6
    _ = v7
    return -1
}
func genFn1() uint16 {
    v8 := GenS1{F0: true}
    if true {
        var v9 float64 = ((float64(0.0) * 0.0) / 1.0)
        _ = v9
    }
    _ = v8
    return 255
}

func main() {
    var v1 uint32 = 13
    v2 := make(map[string]bool)
    var v3 int = 236
    var v4 uint32 = (v1 * 0)
    var v5 int = 720
    v10 := GenS1{F0: ((!(!(false || (!(false || (!(-9223372036854775808 != 0))))))) && (813 == 0))}
    var v11 bool = true
    var v12 uint16 = 0
    v5 += (v3 + v5)
    var v13 float64 = 3.14159
    var v14 float32 = (1.5 - 0.0)
    switch (v3 % ((v3) | 1)) {
    case -493:
        var v15 float64 = (1.5 * v13)
        _ = v15
    case 9223372036854775807:
        v16 := GenS2{F0: 42}
        _ = v16
    case -822:
        v17 := GenS1{F0: true}
        _ = v17
    default:
        delete(v2, (((((("foo bar" + ("" + "")) + ("\n\t" + ((("\n\t" + "foo bar") + "a") + "a"))) + (((("" + (("hello" + "a") + "a")) + "foo bar") + "foo bar") + "foo bar")) + "\n\t") + "hello") + "\n\t"))
    }
    for fi18 := 0; fi18 < 7; fi18++ {
        switch "\n\t" {
        case "foo bar":
            switch ("hello" + "a") {
            case "foo bar":
                fmt.Println(fi18)
            case "a":
                for fi19 := 0; fi19 < 10; fi19++ {
                    var v20 int32 = -2147483648
                    var v21 uint32 = (^(v4))
                    _ = fi19
                    _ = v20
                    _ = v21
                }
            }
        case "\n\t":
            v4 -= (^(v1))
        default:
            if (("hello" + (("foo bar" + "\n\t") + (("" + "foo bar") + (("hello" + "a") + (((("a" + "a") + (("hello" + ("" + "")) + ("foo bar" + ((("\n\t" + "hello") + "foo bar") + ("hello" + "foo bar"))))) + ("hello" + "a")) + ""))))) == "a") {
                var v22 bool = (v11 || (-2147483648 != 0))
                _ = v22
            }
        }
        v12 += v12
        if (!v11) {
            v23 := GenS3{F0: ((9223372036854775807 >= 42) && ((float32((3.14159 - 0.1)) != (3.14159 - 1e10)) || false)), F1: true, F2: ""}
            _ = v23
        }
        _ = fi18
    }
    var v24 uint16 = v12
    for fi25 := 0; fi25 < 2; fi25++ {
        switch ("\n\t" + ("a" + "hello")) {
        case "hello":
            for fi26 := 0; fi26 < 15; fi26++ {
                v24 -= v24
                _ = fi26
            }
        case "\n\t":
            var v27 int8 = 0
            _ = v27
        default:
            for fi28 := 0; fi28 < 8; fi28++ {
                v24 -= (v12 ^ v12)
                switch v5 {
                case 42:
                    if v11 {
                        v14 += 3.14159
                        fmt.Println(fi25)
                        var v29 string = ""
                        v12 = v12
                        var v30 float32 = -2.25
                        v30 = (((1e10 + (v14 / ((v14 * 0.1) / (1.5 / 1.5)))) - (((v30 - v30) + -2.25) / (v30 - 1e10))) / float32((0.1 * v13)))
                        _ = v29
                        _ = v30
                    } else {
                        if v11 {
                            if v11 {
                                v24 += uint16(v1)
                                switch "" {
                                case "":
                                    if (1 < 42) {
                                        var v31 float64 = (v13 - v13)
                                        v14 -= v14
                                        _ = v31
                                    }
                                default:
                                    v2[""] = ((float32((-2.25 * (3.14159 / v13))) <= v14) || (true || v11))
                                }
                            }
                        } else {
                            v2["\n\t"] = v11
                        }
                        v32 := GenS4{F0: (v1 & v1), F1: false}
                        var v33 int64 = int64(fi28)
                        _ = v32
                        _ = v33
                    }
                case 2147483647:
                    v3 += -375
                }
                var v34 string = "\n\t"
                _ = fi28
                _ = v34
            }
        }
        _ = fi25
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
}
