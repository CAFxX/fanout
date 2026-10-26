package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 string; F2 int16 }
type GenS3 struct { F0 uint16; F1 uint }
func genFn0(p0 int64) int8 {
    p0 = 1
    p0 += (p0 - p0)
    var v7 float32 = ((((0.1 + ((3.14159 - (((((3.14159 - 3.14159) / (3.14159 + (1e10 - 1e10))) + ((0.1 * 0.0) - 3.14159)) * 1.5) + ((3.14159 * 1e10) / float32((3.14159 * 1.5))))) - 0.1)) * 3.14159) / (1.5 - float32(3.14159))) / 1.5)
    var v8 uint = 23
    _ = p0
    _ = v7
    _ = v8
    return 71
}

func main() {
    var v1 string = "\n\t"
    v2 := make(map[uint32]int)
    var v3 uint32 = 137
    var v4 string = v1
    v5 := make(map[int]bool)
    v6 := make(map[uint32]uint32)
    v1 = v4
    switch (v1 + "hello") {
    case "\n\t":
        v9 := v5[2147483647]
        _ = v9
    case "foo bar":
        for fi10 := 0; fi10 < 20; fi10++ {
            for fi11 := 0; fi11 < 4; fi11++ {
                _ = len(v2)
                _ = fi11
            }
            switch fi10 {
            case 128:
                switch fi10 {
                case -33:
                    if false {
                        if true {
                            if true {
                                switch 32768 {
                                case 32767:
                                    var v12 uint8 = 2
                                    _ = v12
                                case 50:
                                    for fi13 := 0; fi13 < 7; fi13++ {
                                        fmt.Println(v1, fi10)
                                        _ = fi13
                                    }
                                default:
                                    v4 = "a"
                                }
                                switch (fi10 >> ((4294967295) & 63)) {
                                case 1:
                                    v4 = ("a" + "\n\t")
                                }
                                for fi14 := 0; fi14 < 11; fi14++ {
                                    var v15 string = v1
                                    _ = fi14
                                    _ = v15
                                }
                            } else {
                                v16 := []int{}
                                _ = v16
                            }
                        } else {
                            if (0 >= -1) {
                                var v17 float64 = (float64(((1.5 / 1e10) / 1.0)) + (1.5 * 0.0))
                                _ = v17
                            } else {
                                v3 = v3
                                v1 = v1
                                if (-837 > -1) {
                                    v3 = v3
                                } else {
                                    v3 = (v3 + 677)
                                    fmt.Println(v4, v4, v4)
                                    v3 += (v3 % ((4294967295) | 1))
                                }
                            }
                        }
                    }
                }
            case 42:
                v18 := GenS1{F0: v1}
                _ = v18
            case 2147483647:
                v19 := []int{(fi10 | (^(fi10))), fi10, fi10}
                _ = v19
            }
            _ = fi10
        }
    }
    if true {
        delete(v6, (v3 * v3))
    } else {
        var v20 int8 = -128
        if (true || (-32768 > 0)) {
            delete(v2, (v3 % (((v3 << ((255) & 63))) | 1)))
            v20 = (v20 << ((255) & 63))
            if false {
                v21 := GenS1{F0: v4}
                _ = v21
            } else {
                for fi22 := 0; fi22 < 5; fi22++ {
                    var v23 uint32 = (v3 * (v3 >> ((0) & 63)))
                    switch 9223372036854775807 {
                    case 1:
                        var v24 string = ""
                        _ = v24
                    default:
                        for fi25 := 0; fi25 < 13; fi25++ {
                            v26 := v5[32767]
                            v1 = v4
                            v27 := []int{int(v3)}
                            _ = fi25
                            _ = v26
                            _ = v27
                        }
                    }
                    v28 := []int{fi22, (fi22 * (fi22 / ((2147483647) | 1))), fi22}
                    _ = fi22
                    _ = v23
                    _ = v28
                }
                fmt.Println(v20, v4, v20)
                v4 = (v4 + v4)
            }
        } else {
            fmt.Println(v1, v3)
            switch 32768 {
            case -1:
                var v29 float32 = 3.14159
                _ = v29
            case 598:
                var v30 uint32 = v3
                _ = v30
            }
            switch (v4 + v1) {
            case "a":
                switch 32767 {
                case 1:
                    switch 1 {
                    case 2147483647:
                        v31 := GenS1{F0: ("a" + (v4 + "foo bar"))}
                        _ = v31
                    case 128:
                        v4 = "hello"
                    default:
                        fmt.Println(v20)
                    }
                }
            default:
                fmt.Println(v4, v20)
            }
        }
        _ = v20
    }
    switch 0 {
    case -688:
        v1 = v1
    }
    var v32 uint = 255
    var v33 int = 1
    v3 -= v3
    v32 += v32
    v34 := v6[v3]
    var v35 int32 = -2147483648
    switch "a" {
    case "hello":
        var v36 string = (v1 + v4)
        _ = v36
    case "a":
        v34 -= v34
    case "":
        for fi37 := 0; fi37 < 4; fi37++ {
            v38 := GenS2{F0: (!((!true) && false)), F1: (("a" + (v4 + v1)) + v1), F2: 362}
            v4 = "a"
            _ = fi37
            _ = v38
        }
    }
    for fi39 := 0; fi39 < 10; fi39++ {
        v4 = v4
        var v40 int = (^(fi39))
        v41 := GenS3{F0: 2, F1: v32}
        _ = fi39
        _ = v40
        _ = v41
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
