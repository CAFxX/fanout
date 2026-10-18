package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16; F1 uint16 }
type GenS4 struct { F0 uint16; F1 uint }
type GenS5 struct { F0 string; F1 int8; F2 string }
func genFn0(p0 int) uint32 {
    for fi7 := 0; fi7 < 17; fi7++ {
        fmt.Println(p0)
        _ = fi7
    }
    v8 := GenS1{F0: 42}
    _ = p0
    _ = v8
    return 0
}

func main() {
    var v1 bool = (127 != 0)
    var v2 uint8 = 127
    v3 := make(map[int]string)
    var v4 string = "\n\t"
    v5 := make(map[uint32]bool)
    var v6 uint64 = 0
    if true {
        var v9 int16 = 0
        v9 = v9
        for fi10 := 0; fi10 < 11; fi10++ {
            delete(v3, 4294967295)
            _ = len(v3)
            _ = fi10
        }
        _ = v9
    }
    var v11 uint = 577
    for fi12 := 0; fi12 < 8; fi12++ {
        for fi13 := 0; fi13 < 3; fi13++ {
            v14 := GenS2{F0: (("" + v4) + "foo bar")}
            v15 := GenS2{F0: v4}
            v5[404] = v1
            _ = fi13
            _ = v14
            _ = v15
        }
        switch v4 {
        case "":
            switch "hello" {
            case "\n\t":
                for fi16 := 0; fi16 < 1; fi16++ {
                    v17 := []int{fi16}
                    v1 = v1
                    if (1.5 != (3.14159 + 1.5)) {
                        var v18 uint16 = 662
                        _ = v18
                    } else {
                        v19 := GenS3{F0: 1, F1: 65535}
                        fmt.Println(v6, v1, v11)
                        _ = len(v17)
                        _ = len(v17)
                        var v20 float32 = float32(1e10)
                        v11 += (v11 & (v11 & v11))
                        if v1 {
                            v6 += v6
                        }
                        var v21 uint8 = 125
                        v22 := GenS2{F0: "\n\t"}
                        _ = v19
                        _ = v20
                        _ = v21
                        _ = v22
                    }
                    _ = fi16
                    _ = v17
                }
            case "foo bar":
                v6 += (v6 ^ v6)
            case "a":
                if (!((!false) && (!(1 < -989)))) {
                    switch 9223372036854775807 {
                    case 145:
                        var v23 uint8 = (v2 - 178)
                        _ = v23
                    }
                }
            default:
                for fi24 := 0; fi24 < 15; fi24++ {
                    v1 = true
                    var v25 int64 = 9223372036854775807
                    var v26 string = ((v4 + "a") + ("hello" + (v4 + v4)))
                    _ = fi24
                    _ = v25
                    _ = v26
                }
            }
        case "hello":
            v11 = 1
        default:
            v27 := []int{-1, -9223372036854775808}
            _ = v27
        }
        _ = fi12
    }
    _ = len(v5)
    if (!v1) {
        delete(v3, 127)
        if (1 >= 935) {
            fmt.Println(v2, v1, v1)
        } else {
            var v28 float64 = (3.14159 / 1.0)
            _ = v28
        }
    }
    fmt.Println(v2, v2)
    for fi29 := 0; fi29 < 19; fi29++ {
        if (1e10 != (1e10 * 3.14159)) {
            var v30 string = ""
            v31 := GenS4{F0: 32767, F1: 525}
            _ = v30
            _ = v31
        } else {
            if false {
                v5[genFn0((-(fi29)))] = ((true && v1) && false)
                v32 := GenS2{F0: v4}
                _ = v32
            } else {
                var v33 uint8 = 255
                var v34 int = (fi29 & (fi29 / ((fi29) | 1)))
                switch (-(fi29)) {
                case 580:
                    if ((false && (v2 < v33)) || true) {
                        var v35 string = v4
                        _ = v35
                    } else {
                        fmt.Println(v2)
                        v36 := []int{32767, (-(v34))}
                        var v37 bool = true
                        _ = v36
                        _ = v37
                    }
                case 1:
                    for fi38 := 0; fi38 < 1; fi38++ {
                        v39 := v3[fi38]
                        if true {
                            switch v4 {
                            case "":
                                var v40 uint64 = 929
                                _ = v40
                            default:
                                v33 += (v33 + (v33 & (v33 % (((v33 >> ((0) & 63))) | 1))))
                            }
                            var v41 string = v39
                            _ = v41
                        } else {
                            var v42 bool = true
                            v4 = "foo bar"
                            fmt.Println(v6, fi29)
                            v43 := GenS1{F0: v11}
                            _ = v42
                            _ = v43
                        }
                        var v44 int16 = 0
                        fmt.Println(v33, v4)
                        v45 := v3[(^(fi29))]
                        _ = fi38
                        _ = v39
                        _ = v44
                        _ = v45
                    }
                default:
                    var v46 uint8 = (v2 / (((v2 + v33)) | 1))
                    _ = v46
                }
                _ = v33
                _ = v34
            }
            v47 := []int{fi29, (fi29 / (((fi29 - (-(fi29)))) | 1))}
            _ = v47
        }
        v2 = v2
        v48 := GenS5{F0: (v4 + ("\n\t" + v4)), F1: -128, F2: ("hello" + (v4 + "foo bar"))}
        _ = fi29
        _ = v48
    }
    v49 := []int{464, -6, 2147483647}
    for fi50 := 0; fi50 < 7; fi50++ {
        if (true && v1) {
            fmt.Println(fi50, fi50, v11)
        }
        _ = fi50
    }
    var v51 uint32 = 1
    v2 += (v2 >> (((v11 ^ (^(v11)))) & 63))
    _ = len(v49)
    v1 = false
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v49:", len(v49), v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v51:", v51)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v51
    _ = v6
}
