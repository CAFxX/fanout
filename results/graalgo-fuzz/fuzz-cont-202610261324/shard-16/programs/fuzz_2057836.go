package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 int64; F1 int8 }
type GenS3 struct { F0 uint8; F1 uint8; F2 uint32 }
type GenS4 struct { F0 int16; F1 int64; F2 int16 }
type GenS5 struct { F0 uint16; F1 string; F2 uint16 }
type GenS6 struct { F0 int32; F1 int64; F2 string }
type GenS7 struct { F0 int8; F1 int32; F2 string }
func genFn0() int {
    var v7 uint = 4294967295
    fmt.Println(v7)
    v7 = 1
    v8 := []int{}
    if (2 >= 126) {
        fmt.Println(v7)
    } else {
        switch 1 {
        case -2147483648:
            fmt.Println(v7)
        default:
            var v9 float64 = ((((0.1 * 3.14159) - -2.25) * 3.14159) - float64(1.5))
            _ = v9
        }
        v10 := GenS1{F0: -961, F1: 21}
        _ = v10
    }
    _ = v7
    _ = v8
    return -1
}

func main() {
    v1 := []string{"hello"}
    var v2 int64 = -474
    var v3 float64 = 3.14159
    v4 := make(map[string]string)
    v5 := make(map[string]bool)
    v6 := make(map[uint32]uint32)
    if true {
        for fi11 := 0; fi11 < 19; fi11++ {
            var v12 bool = (false || (true && true))
            var v13 float32 = 0.1
            v12 = true
            _ = fi11
            _ = v12
            _ = v13
        }
        v14 := v6[2]
        fmt.Println(v2, v3)
        _ = v14
    } else {
        v15 := GenS2{F0: v2, F1: 125}
        _ = v15
    }
    v3 -= float64(1.5)
    var v16 int32 = -1
    v16 = v16
    for fi17 := 0; fi17 < 3; fi17++ {
        if (true || (2 > 42)) {
            switch ("\n\t" + "a") {
            case "\n\t":
                var v18 uint = 683
                _ = v18
            }
            v19 := GenS3{F0: 0, F1: 200, F2: 2147483647}
            _ = v19
        } else {
            if false {
                if (((float32(v3) * (-2.25 - (1.5 - 1e10))) * 3.14159) <= 0.1) {
                    switch "" {
                    case "foo bar":
                        if (!(false && ((v3 + v3) >= ((0.1 * float64((0.1 / ((float32((float64(3.14159) + v3)) + 1.5) * (0.1 - ((1.5 + ((1.5 + -2.25) - 0.0)) * 0.0)))))) * float64(((-2.25 - (0.0 - (1.5 * (1.5 + float32(0.1))))) - ((float32(-2.25) / 1.0) / float32(v3)))))))) {
                            for fi20 := 0; fi20 < 20; fi20++ {
                                delete(v5, "a")
                                if (0 >= 529) {
                                    v21 := GenS4{F0: 0, F1: 8, F2: 32767}
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    var v22 string = ("hello" + "a")
                                    _ = v21
                                    _ = v22
                                }
                                if true {
                                    var v23 int8 = 83
                                    _ = v23
                                } else {
                                    fmt.Println(v16, v2, v2)
                                }
                                _ = fi20
                            }
                            for fi24 := 0; fi24 < 14; fi24++ {
                                fmt.Println(v2, v3, v2)
                                v3 -= 0.1
                                v2 += v2
                                _ = fi24
                            }
                        }
                    default:
                        v16 += v16
                    }
                    var v25 float32 = 0.1
                    _ = v25
                }
            }
            v2 = v2
            v26 := GenS5{F0: 0, F1: "foo bar", F2: 622}
            _ = v26
        }
        var v27 int = fi17
        for fi28 := 0; fi28 < 6; fi28++ {
            if (1 == (^(v2))) {
                var v29 bool = true
                var v30 string = "a"
                _ = v29
                _ = v30
            }
            v31 := GenS4{F0: 1000, F1: v2, F2: 0}
            var v32 bool = true
            _ = fi28
            _ = v31
            _ = v32
        }
        _ = fi17
        _ = v27
    }
    v33 := GenS6{F0: v16, F1: (v2 >> ((37) & 63)), F2: ("a" + (("a" + "hello") + "a"))}
    var v34 int64 = (v2 | 0)
    v16 = (-(v16))
    v5["\n\t"] = true
    for fi35 := 0; fi35 < 15; fi35++ {
        if (234 > 127) {
            fmt.Println(v2)
        } else {
            v34 = v2
        }
        if len(v1) > 0 {
            v1[0] = "foo bar"
        }
        _ = fi35
    }
    var v36 uint32 = 283
    for fi37 := 0; fi37 < 8; fi37++ {
        fmt.Println(v3, v16, v36)
        _ = fi37
    }
    v38 := GenS7{F0: 0, F1: (^(v16)), F2: "\n\t"}
    if true {
        v34 = v2
    }
    v16 = (v16 << ((42) & 63))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v34:", v34)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v36
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
