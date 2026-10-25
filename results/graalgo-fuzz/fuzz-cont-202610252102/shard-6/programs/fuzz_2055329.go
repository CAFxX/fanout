package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int16; F2 uint16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint32; F1 int16; F2 uint32 }
type GenS5 struct { F0 string; F1 int; F2 int32 }
func genFn0(p0 string) uint32 {
    var v7 float32 = float32((-2.25 + -2.25))
    v8 := []int{}
    p0 = p0
    var v9 int32 = -653
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return 42
}
func genFn1(p0 int8, p1 uint32) int8 {
    fmt.Println(p0, p1)
    var v10 int = 4294967295
    v11 := []int{}
    v12 := GenS1{F0: 18446744073709551615, F1: -32768, F2: 65535}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v12
    return p0
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint32 = 0
    var v3 uint = 255
    var v4 float64 = (0.0 * 1.5)
    v5 := make(map[string]bool)
    v6 := make(map[int]uint32)
    if (1 > 0) {
        v13 := GenS2{F0: 0}
        if true {
            for fi14 := 0; fi14 < 11; fi14++ {
                v15 := v1[(v2 + 94)]
                var v16 float64 = v4
                v4 += float64((0.0 / (1e10 * 0.1)))
                _ = fi14
                _ = v15
                _ = v16
            }
            switch (("foo bar" + "hello") + "") {
            case "":
                v3 = v3
            case "\n\t":
                for fi17 := 0; fi17 < 9; fi17++ {
                    var v18 float64 = v4
                    delete(v6, (-(fi17)))
                    _ = fi17
                    _ = v18
                }
            case "a":
                v2 += v2
            }
        }
        for fi19 := 0; fi19 < 1; fi19++ {
            v20 := GenS3{F0: "hello"}
            _ = fi19
            _ = v20
        }
        _ = v13
    }
    v2 += (v2 + (v2 * v2))
    var v21 uint64 = 121
    v3 = 57
    if true {
        if (v3 >= 255) {
            switch 32767 {
            case 4294967295:
                var v22 uint64 = (v21 ^ v21)
                _ = v22
            }
        } else {
            var v23 float32 = (3.14159 / ((1e10 - (float32((-2.25 + (((1e10 * 0.1) * (v4 + float64(1.5))) * 0.1))) - 1e10)) * (0.0 * (1e10 / float32(1e10)))))
            for fi24 := 0; fi24 < 6; fi24++ {
                var v25 bool = (uint64(v2) == 753)
                _ = fi24
                _ = v25
            }
            var v26 uint8 = 128
            _ = v23
            _ = v26
        }
        for fi27 := 0; fi27 < 17; fi27++ {
            if true {
                for fi28 := 0; fi28 < 3; fi28++ {
                    if ((1 > -1) && true) {
                        var v29 int = (fi28 + (fi27 & (-(fi27))))
                        switch fi27 {
                        case -79:
                            if (!(1000 == 466)) {
                                switch ("a" + ("foo bar" + "foo bar")) {
                                case "hello":
                                    v3 -= (v3 ^ (v3 % ((v3) | 1)))
                                }
                                var v30 uint16 = 1
                                switch -491 {
                                case 1:
                                    v30 = (v30 ^ 2)
                                default:
                                    var v31 float32 = 0.1
                                    _ = v31
                                }
                                _ = v30
                            }
                        case -459:
                            v32 := []int{127}
                            var v33 int16 = -32768
                            if len(v32) > 0 {
                                v32[0] = fi28
                            }
                            var v34 int = -386
                            var v35 bool = (859 <= 9223372036854775807)
                            _ = v32
                            _ = v33
                            _ = v34
                            _ = v35
                        }
                        v36 := GenS4{F0: v2, F1: -1, F2: (v2 ^ v2)}
                        _ = v29
                        _ = v36
                    }
                    _ = fi28
                }
                var v37 string = "foo bar"
                if ("hello" <= ("hello" + ("a" + (v37 + "hello")))) {
                    fmt.Println(v3, v2)
                    v38 := GenS3{F0: ("hello" + "")}
                    for fi39 := 0; fi39 < 5; fi39++ {
                        v4 = (v4 * (1e10 / 0.1))
                        _ = fi39
                    }
                    _ = v38
                }
                _ = v37
            } else {
                v3 -= v3
                var v40 int64 = -1
                _ = v40
            }
            v2 = v2
            v41 := GenS5{F0: "", F1: (fi27 >> (((v3 / ((2) | 1))) & 63)), F2: 73}
            _ = fi27
            _ = v41
        }
        v4 -= -2.25
    }
    fmt.Println(v3, v21, v4)
    switch "" {
    case "foo bar":
        var v42 int16 = 0
        _ = v42
    default:
        var v43 int64 = -9223372036854775808
        _ = v43
    }
    _ = len(v5)
    if true {
        for fi44 := 0; fi44 < 8; fi44++ {
            var v45 int8 = genFn1(116, 2147483647)
            var v46 int = (fi44 << ((1) & 63))
            _ = fi44
            _ = v45
            _ = v46
        }
    } else {
        v5[("a" + "")] = false
    }
    var v47 uint = v3
    if ((0.1 < ((1e10 / (float32((v4 / v4)) + ((3.14159 + 0.0) / (1e10 + -2.25)))) / 1e10)) || (false || (65535 >= 345))) {
        v3 += (^(v47))
        v48 := GenS1{F0: (v21 - 155), F1: 351, F2: 32767}
        var v49 uint64 = 0
        _ = v48
        _ = v49
    }
    v21 += uint64(v2)
    v2 += v2
    for fi50 := 0; fi50 < 1; fi50++ {
        fmt.Println(v4, v21, v21)
        _ = fi50
    }
    var v51 uint32 = (v2 * v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v47:", v47)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v51:", v51)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v51
    _ = v6
}
