package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint64; F1 uint64; F2 int8 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint64; F1 int16; F2 uint8 }
func genFn0() uint8 {
    v7 := []int{128, 1, 32767}
    var v8 float64 = -2.25
    _ = v7
    _ = v8
    return 1
}

func main() {
    v1 := make(map[int]uint32)
    v2 := []string{"foo bar"}
    v3 := []int{127, -1, -2147483648}
    v4 := []uint32{}
    var v5 uint32 = 4294967295
    var v6 uint32 = 446
    v2 = append(v2, ("a" + ("" + ("" + "\n\t"))))
    var v9 int8 = 0
    for fi10 := 0; fi10 < 6; fi10++ {
        delete(v1, (fi10 % ((667) | 1)))
        _ = fi10
    }
    v5 -= v6
    var v11 bool = (false && true)
    for fi12 := 0; fi12 < 1; fi12++ {
        v11 = (4294967295 >= 1)
        v13 := GenS1{F0: 2}
        _ = fi12
        _ = v13
    }
    var v14 int64 = 888
    if (-866 >= 396) {
        for fi15 := 0; fi15 < 16; fi15++ {
            if v11 {
                for fi16 := 0; fi16 < 15; fi16++ {
                    if v11 {
                        switch "\n\t" {
                        case "":
                            fmt.Println(v5, v14, v14)
                        case "\n\t":
                            v6 = (v5 & v6)
                        case "foo bar":
                            var v17 uint64 = 42
                            _ = v17
                        }
                    } else {
                        var v18 bool = ((v14 * -150) <= v14)
                        _ = v18
                    }
                    _ = fi16
                }
                switch (fi15 >> ((2) & 63)) {
                case 881:
                    v5 = v6
                case 576:
                    switch fi15 {
                    case 0:
                        fmt.Println(v6)
                    case 32768:
                        v19 := GenS1{F0: 0}
                        _ = v19
                    }
                case -9223372036854775808:
                    _ = len(v1)
                default:
                    switch "\n\t" {
                    case "foo bar":
                        if (!(v11 && (genFn0() != 2))) {
                            var v20 int32 = -2147483648
                            v21 := GenS2{F0: 0, F1: 2, F2: (v9 << ((1) & 63))}
                            switch "foo bar" {
                            case "a":
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            case "foo bar":
                                switch ("hello" + "hello") {
                                case "\n\t":
                                    for fi22 := 0; fi22 < 9; fi22++ {
                                        v23 := GenS1{F0: 1}
                                        var v24 float32 = 0.1
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                        _ = fi22
                                        _ = v23
                                        _ = v24
                                    }
                                default:
                                    switch ("hello" + "") {
                                    case "":
                                        fmt.Println(v14)
                                    case "\n\t":
                                        v14 += (v14 * (v14 - (v14 % ((v14) | 1))))
                                    case "hello":
                                        v3 = append(v3, int(v9))
                                    }
                                }
                            default:
                                for fi25 := 0; fi25 < 10; fi25++ {
                                    v26 := GenS3{F0: 2}
                                    if len(v4) > 0 {
                                        v4[0] = (v5 / ((v6) | 1))
                                    }
                                    fmt.Println(v11, v20)
                                    var v27 bool = false
                                    _ = fi25
                                    _ = v26
                                    _ = v27
                                }
                            }
                            _ = v20
                            _ = v21
                        }
                    }
                }
                var v28 int16 = -32768
                _ = v28
            }
            for fi29 := 0; fi29 < 16; fi29++ {
                fmt.Println(v14, v11)
                _ = fi29
            }
            switch ("\n\t" + "") {
            case "a":
                v30 := GenS4{F0: 578, F1: 1, F2: 255}
                _ = v30
            default:
                for fi31 := 0; fi31 < 11; fi31++ {
                    var v32 int64 = (v14 ^ (v14 & v14))
                    _ = fi31
                    _ = v32
                }
            }
            _ = fi15
        }
    }
    var v33 float64 = -2.25
    if false {
        fmt.Println(v5, v33, v11)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
