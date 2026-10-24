package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 bool }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]bool)
    v3 := make(map[int]string)
    var v4 int64 = 1
    var v5 bool = false
    var v6 int8 = 127
    var v7 bool = ((1.5 / 1e10) == (3.14159 * 0.0))
    _ = len(v1)
    v4 -= (^(v4))
    var v8 float64 = 0.0
    var v9 int8 = 0
    var v10 int32 = 722
    switch (("a" + "hello") + "") {
    case "foo bar":
        var v11 uint8 = 128
        _ = v11
    case "hello":
        v10 = -374
    case "":
        switch ("hello" + ("\n\t" + ("" + "a"))) {
        case "hello":
            switch 128 {
            case -9223372036854775808:
                v10 += (v10 * (v10 - (v10 | (v10 * (v10 & v10)))))
            case 1:
                v12 := []int{int(v4)}
                _ = v12
            case 569:
                fmt.Println(v4, v8)
            default:
                if v5 {
                    var v13 bool = (455 != 453)
                    _ = v13
                } else {
                    if v7 {
                        v4 = (v4 & (v4 | (v4 * v4)))
                        v7 = v5
                        for fi14 := 0; fi14 < 18; fi14++ {
                            v7 = ("hello" <= ("hello" + "foo bar"))
                            v15 := []int{fi14}
                            delete(v2, fi14)
                            _ = fi14
                            _ = v15
                        }
                    }
                    v16 := []int{128, 42, 128, 32768}
                    v7 = (false && (v7 || ((v6 * v6) <= v6)))
                    _ = v16
                }
            }
        case "\n\t":
            var v17 uint64 = 0
            _ = v17
        case "":
            var v18 uint16 = 482
            _ = v18
        }
    default:
        switch "a" {
        case "a":
            if ((!v7) && (!true)) {
                for fi19 := 0; fi19 < 12; fi19++ {
                    if (1 < 0) {
                        v20 := v3[fi19]
                        fmt.Println(v8, fi19)
                        _ = v20
                    } else {
                        var v21 uint16 = 42
                        for fi22 := 0; fi22 < 9; fi22++ {
                            if (!v7) {
                                var v23 string = ("a" + "a")
                                _ = v23
                            } else {
                                v24 := []int{(fi19 % (((^(fi22))) | 1)), fi19, (fi19 >> ((0) & 63)), 42}
                                fmt.Println(v8, v8)
                                var v25 float32 = ((((0.0 - 0.1) + ((float32((0.1 - v8)) - float32(v8)) + 3.14159)) - 0.1) - float32(0.1))
                                if v7 {
                                    var v26 bool = true
                                    var v27 int32 = v10
                                    _ = v26
                                    _ = v27
                                }
                                _ = v24
                                _ = v25
                            }
                            var v28 int32 = 823
                            v29 := GenS1{F0: "", F1: 2147483647, F2: (v7 && (1.5 == (((3.14159 + 0.1) * 1.5) - 1.5)))}
                            _ = fi22
                            _ = v28
                            _ = v29
                        }
                        _ = v21
                    }
                    var v30 int = -2147483648
                    for fi31 := 0; fi31 < 5; fi31++ {
                        v10 = 1
                        _ = fi31
                    }
                    _ = fi19
                    _ = v30
                }
                fmt.Println(v9, v8)
                fmt.Println(v10, v4, v7)
            }
        }
    }
    fmt.Println(v5, v6, v10)
    v32 := GenS1{F0: "a", F1: 962, F2: v7}
    switch "" {
    case "":
        v9 -= v6
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
