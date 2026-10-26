package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8; F1 uint16 }
type GenS3 struct { F0 int32; F1 string }
func genFn0() uint64 {
    if true {
        var v7 uint32 = 216
        _ = v7
    }
    for fi8 := 0; fi8 < 8; fi8++ {
        v9 := []int{fi8}
        _ = fi8
        _ = v9
    }
    return 48
}

func main() {
    v1 := make(map[string]string)
    v2 := []uint32{}
    v3 := []string{}
    v4 := make(map[string]bool)
    v5 := []string{"", (("" + (("a" + "") + "")) + "foo bar"), "a"}
    v6 := make(map[uint32]string)
    if (652 >= 505) {
        var v10 float32 = 1e10
        _ = v10
    } else {
        switch "hello" {
        case "\n\t":
            v11 := GenS1{F0: (("" + ("a" + "foo bar")) + ("foo bar" + "foo bar"))}
            _ = v11
        default:
            if (-2147483648 >= 980) {
                var v12 float32 = 0.0
                for fi13 := 0; fi13 < 3; fi13++ {
                    for fi14 := 0; fi14 < 18; fi14++ {
                        switch fi14 {
                        case 0:
                            var v15 uint8 = 0
                            _ = v15
                        case 494:
                            var v16 uint8 = 87
                            _ = v16
                        }
                        var v17 float64 = -2.25
                        if (-680 != -786) {
                            fmt.Println(fi14, v12, fi13)
                            switch "a" {
                            case "":
                                v18 := GenS2{F0: -128, F1: 808}
                                _ = v18
                            case "hello":
                                if len(v3) > 0 {
                                    v3[0] = ("foo bar" + ("\n\t" + "foo bar"))
                                }
                            case "foo bar":
                                v19 := GenS1{F0: "hello"}
                                _ = v19
                            }
                        }
                        _ = fi14
                        _ = v17
                    }
                    v2 = append(v2, 0)
                    for fi20 := 0; fi20 < 12; fi20++ {
                        var v21 int64 = -9223372036854775808
                        _ = fi20
                        _ = v21
                    }
                    _ = fi13
                }
                _ = v12
            } else {
                var v22 bool = (65535 != 211)
                _ = v22
            }
        }
        v23 := GenS2{F0: 1, F1: 42}
        for fi24 := 0; fi24 < 4; fi24++ {
            var v25 int32 = 686
            var v26 float64 = -2.25
            var v27 int8 = -78
            _ = fi24
            _ = v25
            _ = v26
            _ = v27
        }
        _ = v23
    }
    switch -314 {
    case 496:
        var v28 uint64 = 2
        _ = v28
    case 1:
        var v29 int8 = -128
        _ = v29
    case 0:
        var v30 uint = 18446744073709551615
        _ = v30
    default:
        for fi31 := 0; fi31 < 11; fi31++ {
            switch ("hello" + "") {
            case "hello":
                v32 := GenS3{F0: -1, F1: "a"}
                _ = v32
            default:
                var v33 uint16 = 2
                _ = v33
            }
            _ = fi31
        }
    }
    var v34 uint8 = 42
    v34 = 0
    fmt.Println(v34)
    var v35 float64 = (1.5 * 0.1)
    fmt.Println(v35, v34)
    var v36 uint = 0
    v35 += ((v35 + 0.1) - (3.14159 / 1.5))
    fmt.Println(v35)
    var v37 int16 = 42
    v36 = v36
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
