package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint32; F1 uint8; F2 string }
type GenS3 struct { F0 bool; F1 uint8 }
type GenS4 struct { F0 int; F1 uint8; F2 string }
func genFn0(p0 uint8) int64 {
    var v7 string = ("foo bar" + "hello")
    var v8 int8 = 100
    var v9 int8 = v8
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return -1
}

func main() {
    var v1 int8 = 0
    var v2 int8 = v1
    v3 := []string{("\n\t" + "\n\t"), ("hello" + ("foo bar" + ("hello" + "hello")))}
    var v4 int32 = -1
    var v5 uint = 18446744073709551615
    var v6 int8 = (v2 & v2)
    if false {
        v3 = append(v3, (("a" + "hello") + (("" + ("" + "foo bar")) + "\n\t")))
    } else {
        if false {
            v10 := GenS1{F0: 0}
            _ = v10
        } else {
            var v11 uint16 = 433
            fmt.Println(v1, v4, v2)
            _ = v11
        }
        var v12 int8 = v6
        var v13 int8 = (v12 % ((-1) | 1))
        _ = v12
        _ = v13
    }
    v3 = append(v3, "a")
    if true {
        v5 = 4294967295
        v5 += 1
    } else {
        for fi14 := 0; fi14 < 10; fi14++ {
            if ((false && (("" + ("" + "foo bar")) < (((("a" + (("a" + "\n\t") + (("" + ("hello" + "\n\t")) + "foo bar"))) + "hello") + "") + "\n\t"))) && false) {
                var v15 float32 = (0.0 / 1.0)
                var v16 int = 0
                v17 := GenS2{F0: 2147483648, F1: 118, F2: "\n\t"}
                _ = v15
                _ = v16
                _ = v17
            }
            _ = fi14
        }
        v18 := GenS3{F0: (!(1.5 != float32(0.0))), F1: 127}
        _ = v18
    }
    if false {
        v1 += -1
        switch "\n\t" {
        case "a":
            v2 += v2
        default:
            v19 := GenS4{F0: -2147483648, F1: 108, F2: "\n\t"}
            _ = v19
        }
        var v20 int8 = v2
        _ = v20
    }
    v3 = append(v3, ("\n\t" + ("\n\t" + "")))
    v6 = (v6 - (v1 / (((v6 * (v1 - (v2 + v6)))) | 1)))
    if (!(v5 <= (v5 >> ((v5) & 63)))) {
        if true {
            v2 = -7
            for fi21 := 0; fi21 < 14; fi21++ {
                for fi22 := 0; fi22 < 15; fi22++ {
                    switch fi21 {
                    case 4294967295:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        switch "foo bar" {
                        case "":
                            fmt.Println(v2, fi22, v5)
                        case "a":
                            var v23 int64 = 0
                            _ = v23
                        default:
                            if (0 != -1) {
                                if ((-2.25 / 1.0) > 1.5) {
                                    fmt.Println(v1)
                                }
                                v2 = v2
                            }
                        }
                        var v24 int32 = v4
                        if len(v3) > 0 {
                            v3[0] = "foo bar"
                        }
                        var v25 int32 = (^(v24))
                        _ = v24
                        _ = v25
                    default:
                        if len(v3) > 0 {
                            v3[0] = ("a" + ("hello" + "a"))
                        }
                    }
                    v26 := GenS2{F0: 4294967295, F1: 2, F2: ""}
                    var v27 bool = (!(-32768 >= 32767))
                    var v28 bool = v27
                    _ = fi22
                    _ = v26
                    _ = v27
                    _ = v28
                }
                for fi29 := 0; fi29 < 3; fi29++ {
                    v5 = v5
                    v5 = v5
                    _ = fi29
                }
                _ = fi21
            }
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        for fi30 := 0; fi30 < 6; fi30++ {
            var v31 uint64 = 1000
            var v32 int32 = v4
            v4 += -2147483648
            _ = fi30
            _ = v31
            _ = v32
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
    } else {
        var v33 uint64 = 42
        _ = v33
    }
    var v34 float32 = 1e10
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v6
}
