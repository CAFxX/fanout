package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int64; F1 uint16 }
type GenS4 struct { F0 int32; F1 string; F2 int }
type GenS5 struct { F0 int8; F1 string; F2 int64 }
type GenS6 struct { F0 uint16; F1 string }
type GenS7 struct { F0 int32; F1 int8; F2 bool }
func genFn0() int64 {
    var v7 float32 = (3.14159 + -2.25)
    v8 := GenS1{F0: 1}
    _ = v7
    _ = v8
    return 433
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 float32 = 1e10
    var v3 uint32 = 366
    var v4 uint32 = (v3 - (v3 | 61))
    v5 := make(map[string]bool)
    v6 := []int{}
    v9 := GenS2{F0: 1}
    v2 += 3.14159
    switch (("foo bar" + "foo bar") + "\n\t") {
    case "\n\t":
        if ((true && true) && false) {
            for fi10 := 0; fi10 < 3; fi10++ {
                for fi11 := 0; fi11 < 5; fi11++ {
                    v12 := GenS1{F0: 0}
                    _ = fi11
                    _ = v12
                }
                v4 -= (v3 | v4)
                for fi13 := 0; fi13 < 17; fi13++ {
                    if true {
                        v2 += v2
                        v6 = append(v6, fi13)
                        if (!(42 < 0)) {
                            var v14 int8 = -110
                            _ = v14
                        } else {
                            _ = len(v6)
                            v2 = v2
                        }
                        fmt.Println(v4, v3, fi13)
                    }
                    _ = fi13
                }
                _ = fi10
            }
            var v15 bool = true
            _ = v15
        }
    case "hello":
        v3 += v3
    }
    for fi16 := 0; fi16 < 4; fi16++ {
        var v17 int = fi16
        if true {
            v17 = v17
            v4 = (v4 * v3)
        } else {
            if len(v6) >= 2 {
                v6 = v6[1:len(v6)]
            }
            v3 -= v3
            var v18 bool = true
            _ = v18
        }
        _ = fi16
        _ = v17
    }
    switch "a" {
    case "hello":
        if ((!((float32((1e10 * (((3.14159 * 0.1) * 3.14159) / float64(v2)))) * float32(float64(((float32(3.14159) * ((0.1 + v2) / (v2 + -2.25))) / (v2 + 1.5))))) > v2)) && ((-1 != 552) || false)) {
            var v19 float32 = v2
            _ = v19
        } else {
            if true {
                switch ("hello" + "a") {
                case "hello":
                    v20 := GenS1{F0: 839}
                    _ = v20
                case "foo bar":
                    for fi21 := 0; fi21 < 17; fi21++ {
                        var v22 float32 = v2
                        v23 := GenS3{F0: -474, F1: 2}
                        _ = fi21
                        _ = v22
                        _ = v23
                    }
                default:
                    var v24 float32 = v2
                    _ = v24
                }
            } else {
                if false {
                    for fi25 := 0; fi25 < 5; fi25++ {
                        if true {
                            var v26 uint16 = 1
                            _ = len(v6)
                            v3 = (v3 | (v3 >> ((1) & 63)))
                            _ = v26
                        } else {
                            v27 := GenS4{F0: 42, F1: "hello", F2: -458}
                            v3 -= v4
                            _ = v27
                        }
                        _ = fi25
                    }
                }
                _ = len(v6)
                for fi28 := 0; fi28 < 5; fi28++ {
                    for fi29 := 0; fi29 < 15; fi29++ {
                        if false {
                            fmt.Println(v3, fi29, v3)
                        } else {
                            v30 := GenS5{F0: 1, F1: ("a" + "foo bar"), F2: -168}
                            v31 := GenS6{F0: 255, F1: ""}
                            var v32 uint16 = 42
                            _ = v30
                            _ = v31
                            _ = v32
                        }
                        v33 := v1[(v3 - v4)]
                        _ = fi29
                        _ = v33
                    }
                    var v34 uint8 = 128
                    _ = fi28
                    _ = v34
                }
            }
            var v35 uint8 = 1
            if true {
                if false {
                    if (-985 != 32767) {
                        var v36 uint64 = 18446744073709551615
                        _ = v36
                    } else {
                        var v37 uint32 = (^(v4))
                        _ = v37
                    }
                    fmt.Println(v3, v4)
                    v38 := GenS5{F0: 0, F1: "hello", F2: 1}
                    _ = v38
                } else {
                    if (2 < 306) {
                        if (68 < -634) {
                            var v39 bool = false
                            for fi40 := 0; fi40 < 19; fi40++ {
                                v41 := GenS1{F0: 2}
                                v1[2147483647] = (v4 << ((18446744073709551615) & 63))
                                v35 -= v35
                                _ = fi40
                                _ = v41
                            }
                            v2 = (v2 * v2)
                            _ = v39
                        }
                    }
                }
                v42 := GenS4{F0: -1, F1: "", F2: 551}
                _ = v42
            } else {
                v3 = v4
            }
            _ = v35
        }
    case "\n\t":
        var v43 int64 = 1
        _ = v43
    case "":
        v44 := v5["foo bar"]
        _ = v44
    }
    v4 = 2147483648
    var v45 uint32 = v3
    var v46 int64 = 1
    if false {
        var v47 uint64 = 4294967295
        _ = v47
    }
    v46 = v46
    var v48 float32 = float32((1.5 - (0.1 * 1.5)))
    var v49 uint8 = 141
    fmt.Println(v49, v3, v2)
    var v50 int64 = (-(v46))
    v51 := GenS7{F0: 42, F1: -128, F2: ((true && ((v49 > v49) && false)) || (!(!false)))}
    var v52 string = ((("a" + "hello") + ("\n\t" + ("hello" + "hello"))) + "hello")
    var v53 uint16 = 255
    v46 = (v46 % ((genFn0()) | 1))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v48:", v48)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v48
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v52
    _ = v53
    _ = v6
    _ = v9
}
