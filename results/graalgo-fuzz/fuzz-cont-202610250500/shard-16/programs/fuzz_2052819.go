package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint64; F1 uint }
func genFn0(p0 int32, p1 uint16, p2 int8) int8 {
    fmt.Println(p0, p2)
    v7 := []int{862, 128}
    if len(v7) > 0 {
        v7[0] = 32767
    }
    var v8 uint = 14
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 40
}
func genFn1() uint {
    v9 := []int{-1, 128, 32768}
    var v10 int64 = 939
    switch -1 {
    case 32767:
        switch 127 {
        case -2147483648:
            if (255 >= 0) {
                var v11 string = ((("foo bar" + "foo bar") + "\n\t") + "foo bar")
                _ = v11
            } else {
                switch -2147483648 {
                case -2147483648:
                    if false {
                        var v12 int64 = 9223372036854775807
                        _ = v12
                    }
                default:
                    v13 := GenS1{F0: 65535}
                    _ = v13
                }
            }
        }
    default:
        v10 -= v10
    }
    _ = v10
    _ = v9
    return 65535
}

func main() {
    var v1 uint16 = 255
    v2 := make(map[int]uint32)
    v3 := make(map[uint32]string)
    v4 := make(map[int]int)
    v5 := []int{-901}
    var v6 int8 = -1
    for fi14 := 0; fi14 < 2; fi14++ {
        switch (("" + "foo bar") + "foo bar") {
        case "hello":
            if (!(42 <= -1)) {
                v5 = append(v5, (fi14 & (fi14 + fi14)))
                v15 := GenS2{F0: 578}
                _ = v15
            } else {
                v6 -= v6
                v1 = 2
            }
        case "a":
            v1 = (^(v1))
        case "":
            if (0 >= -525) {
                v16 := GenS1{F0: 42}
                fmt.Println(v1, v6, v1)
                _ = v16
            }
        default:
            v5 = append(v5, 951)
        }
        _ = fi14
    }
    var v17 int32 = -2147483648
    var v18 int16 = -1
    for fi19 := 0; fi19 < 3; fi19++ {
        if (-9223372036854775808 < 0) {
            switch "hello" {
            case "":
                var v20 uint8 = 2
                _ = v20
            }
        } else {
            for fi21 := 0; fi21 < 2; fi21++ {
                if ((v17 - (v17 - v17)) == v17) {
                    var v22 int64 = 668
                    v23 := GenS3{F0: 42, F1: 255}
                    v24 := GenS3{F0: 0, F1: 42}
                    _ = v22
                    _ = v23
                    _ = v24
                }
                v18 += v18
                switch "a" {
                case "hello":
                    if false {
                        v17 -= (v17 % ((1) | 1))
                    } else {
                        if false {
                            if false {
                                fmt.Println(fi19, v17, v6)
                            } else {
                                v6 = v6
                                v18 = (v18 / ((v18) | 1))
                            }
                            switch fi21 {
                            case 0:
                                var v25 float64 = (float64((0.1 - (3.14159 - 1e10))) + float64(float32((1e10 + 1e10))))
                                _ = v25
                            case -2147483648:
                                v26 := v2[fi19]
                                _ = v26
                            case 32767:
                                v27 := GenS1{F0: 1}
                                _ = v27
                            }
                            var v28 float64 = 0.0
                            _ = v28
                        } else {
                            v29 := GenS2{F0: 2147483647}
                            _ = v29
                        }
                    }
                }
                _ = fi21
            }
            var v30 uint16 = 32767
            _ = v30
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = fi19
    }
    v31 := GenS2{F0: 756}
    fmt.Println(v1)
    var v32 int16 = 0
    var v33 int32 = (v17 << ((3) & 63))
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    var v34 uint16 = v1
    v3[568] = ("\n\t" + "hello")
    v35 := GenS1{F0: 2}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
