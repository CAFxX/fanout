package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 uint8 }
type GenS2 struct { F0 bool; F1 int16; F2 string }
func genFn0() uint64 {
    if (false || false) {
        var v7 int = 0
        _ = v7
    } else {
        if (true || false) {
            for fi8 := 0; fi8 < 4; fi8++ {
                fmt.Println(fi8)
                if (!false) {
                    var v9 uint16 = 114
                    fmt.Println(v9)
                    v9 -= v9
                    _ = v9
                }
                if false {
                    if (32767 >= 65535) {
                        fmt.Println(fi8)
                    } else {
                        switch (("hello" + "") + "hello") {
                        case "\n\t":
                            switch -776 {
                            case -733:
                                for fi10 := 0; fi10 < 8; fi10++ {
                                    switch (fi8 << ((42) & 63)) {
                                    case 0:
                                        v11 := GenS1{F0: 18446744073709551615, F1: 198, F2: 255}
                                        v12 := []int{fi8, fi10, (fi10 * 128), fi10}
                                        var v13 uint = 255
                                        _ = v11
                                        _ = v12
                                        _ = v13
                                    case -874:
                                        var v14 float64 = 0.1
                                        _ = v14
                                    }
                                    var v15 bool = (!true)
                                    _ = fi10
                                    _ = v15
                                }
                            }
                        case "a":
                            var v16 float32 = (((0.0 / 0.1) - 0.0) + (1.5 + 1.5))
                            _ = v16
                        case "":
                            if true {
                                var v17 uint64 = 42
                                _ = v17
                            } else {
                                for fi18 := 0; fi18 < 9; fi18++ {
                                    v19 := GenS1{F0: 4294967295, F1: 127, F2: 128}
                                    var v20 string = ("a" + "hello")
                                    _ = fi18
                                    _ = v19
                                    _ = v20
                                }
                            }
                        default:
                            for fi21 := 0; fi21 < 6; fi21++ {
                                if true {
                                    var v22 int16 = 364
                                    fmt.Println(fi8, v22)
                                    if true {
                                        v23 := GenS1{F0: 652, F1: 33, F2: 127}
                                        var v24 uint64 = 1
                                        _ = v23
                                        _ = v24
                                    }
                                    if (2147483647 != -2147483648) {
                                        v25 := []int{(fi21 ^ fi8), (fi8 << ((4294967295) & 63)), (-(fi8)), fi21}
                                        fmt.Println(v22, v22)
                                        _ = v25
                                    }
                                    _ = v22
                                } else {
                                    var v26 uint32 = 42
                                    v27 := []int{fi8}
                                    _ = v26
                                    _ = v27
                                }
                                var v28 float64 = float64(((0.1 / 3.14159) + (3.14159 * float32(0.1))))
                                fmt.Println(v28, fi8, fi8)
                                _ = fi21
                                _ = v28
                            }
                        }
                    }
                }
                _ = fi8
            }
            var v29 uint16 = 101
            var v30 uint64 = 42
            _ = v29
            _ = v30
        }
        var v31 uint64 = 18446744073709551615
        _ = v31
    }
    var v32 uint64 = 4294967295
    for fi33 := 0; fi33 < 9; fi33++ {
        var v34 uint = 42
        var v35 uint64 = 1
        _ = fi33
        _ = v34
        _ = v35
    }
    if (-1 > -463) {
        v32 -= v32
        for fi36 := 0; fi36 < 16; fi36++ {
            v32 = v32
            _ = fi36
        }
    } else {
        switch 128 {
        case 9223372036854775807:
            var v37 int = 0
            _ = v37
        case -2147483648:
            v38 := GenS2{F0: false, F1: 42, F2: "a"}
            _ = v38
        default:
            v39 := GenS1{F0: v32, F1: 2, F2: 2}
            _ = v39
        }
        var v40 uint8 = 128
        _ = v40
    }
    _ = v32
    return (^(v32))
}

func main() {
    var v1 int8 = -28
    v2 := []string{"foo bar", "", ""}
    var v3 float64 = 0.0
    var v4 string = ("" + ("\n\t" + (("foo bar" + ("foo bar" + "foo bar")) + "hello")))
    v5 := []string{v4, "foo bar"}
    v6 := make(map[int]string)
    v1 += v1
    v1 = v1
    v1 = v1
    var v41 string = v4
    switch "\n\t" {
    case "hello":
        _ = len(v6)
    case "foo bar":
        switch 568 {
        case 127:
            var v42 uint = 437
            _ = v42
        default:
            fmt.Println(v41)
        }
    }
    switch -179 {
    case 42:
        v3 = (v3 + v3)
    }
    var v43 uint16 = 32767
    if (!(9223372036854775807 >= 9223372036854775807)) {
        var v44 int = 327
        v45 := v6[(v44 >> ((uint(v43)) & 63))]
        _ = v44
        _ = v45
    } else {
        var v46 int32 = 0
        _ = v46
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v43
    _ = v5
    _ = v6
}
