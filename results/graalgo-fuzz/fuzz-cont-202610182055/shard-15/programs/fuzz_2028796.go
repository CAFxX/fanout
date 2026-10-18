package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 string; F1 uint }
func genFn0(p0 uint32, p1 int32) int {
    var v5 int64 = int64(p0)
    for fi6 := 0; fi6 < 11; fi6++ {
        if true {
            p0 += 685
        }
        if true {
            var v7 uint = 255
            v8 := GenS1{F0: 4294967295, F1: ""}
            _ = v7
            _ = v8
        }
        _ = fi6
    }
    var v9 bool = false
    for fi10 := 0; fi10 < 13; fi10++ {
        p0 -= p0
        _ = fi10
    }
    _ = p0
    _ = p1
    _ = v5
    _ = v9
    return 9223372036854775807
}
func genFn1(p0 int16, p1 int32, p2 uint16) int64 {
    p2 += p2
    v11 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 0
}

func main() {
    var v1 float32 = ((1.5 + 0.0) - 1e10)
    var v2 uint = 65535
    var v3 int64 = 9223372036854775807
    var v4 int32 = 0
    switch 128 {
    case 127:
        v12 := GenS1{F0: 0, F1: ("hello" + ("" + "a"))}
        _ = v12
    case -816:
        v13 := []int{}
        _ = v13
    }
    var v14 int8 = 127
    if (1 > 784) {
        if false {
            fmt.Println(v1, v3)
            for fi15 := 0; fi15 < 16; fi15++ {
                for fi16 := 0; fi16 < 20; fi16++ {
                    var v17 int32 = v4
                    _ = fi16
                    _ = v17
                }
                _ = fi15
            }
            var v18 uint32 = 583
            _ = v18
        }
        var v19 float32 = 1.5
        _ = v19
    } else {
        v20 := []int{genFn0(42, v4), -1, genFn0(374, v4)}
        for fi21 := 0; fi21 < 19; fi21++ {
            v1 = float32(1e10)
            var v22 int = fi21
            fmt.Println(v2)
            _ = fi21
            _ = v22
        }
        if false {
            var v23 int32 = v4
            var v24 uint8 = 80
            _ = v23
            _ = v24
        } else {
            if ((1.5 <= 1.5) || (v1 == ((1e10 - (0.1 - (v1 - (0.1 + (float32(float64(v1)) / v1))))) - v1))) {
                for fi25 := 0; fi25 < 7; fi25++ {
                    var v26 uint64 = 0
                    if false {
                        switch 128 {
                        case -9223372036854775808:
                            for fi27 := 0; fi27 < 15; fi27++ {
                                if (((!(true || (false && true))) && false) || ((v3 << ((v2) & 63)) == -793)) {
                                    var v28 bool = (0.1 > 1.5)
                                    _ = v28
                                } else {
                                    var v29 uint32 = 2147483647
                                    _ = v29
                                }
                                if (32767 > 737) {
                                    v2 = 2
                                    var v30 float32 = (v1 - 0.0)
                                    var v31 int8 = -128
                                    v3 -= (^(v3))
                                    _ = v30
                                    _ = v31
                                }
                                _ = fi27
                            }
                        default:
                            fmt.Println(v2, v1)
                        }
                    }
                    _ = fi25
                    _ = v26
                }
                v2 = v2
            }
            var v32 uint8 = 42
            _ = v32
        }
        _ = v20
    }
    v33 := []int{4294967295, 0, -112, 127}
    fmt.Println(v3, v2, v1)
    if true {
        var v34 uint32 = 2147483647
        _ = v34
    } else {
        fmt.Println(v3, v1)
        var v35 int32 = v4
        var v36 uint64 = 4294967295
        _ = v35
        _ = v36
    }
    fmt.Println(v1, v14)
    v3 = genFn1(1000, -2147483648, 0)
    if false {
        for fi37 := 0; fi37 < 5; fi37++ {
            if len(v33) > 0 {
                v33[0] = fi37
            }
            for fi38 := 0; fi38 < 2; fi38++ {
                v2 = 255
                v39 := GenS2{F0: v14}
                if (!(128 > 1)) {
                    var v40 uint32 = 2
                    v41 := GenS2{F0: (-(v14))}
                    var v42 string = "a"
                    _ = v40
                    _ = v41
                    _ = v42
                } else {
                    v1 = 3.14159
                    if len(v33) >= 2 {
                        v33 = v33[1:len(v33)]
                    }
                }
                _ = fi38
                _ = v39
            }
            v3 = (v3 << ((v2) & 63))
            _ = fi37
        }
        var v43 uint32 = 280
        _ = v43
    } else {
        v33 = append(v33, -2147483648)
        var v44 uint = v2
        _ = v44
    }
    if (-198 == -32768) {
        var v45 uint = 4294967295
        v3 -= v3
        _ = v45
    }
    v3 -= -327
    if false {
        v3 = v3
        v46 := GenS1{F0: 113, F1: (("foo bar" + "") + "\n\t")}
        _ = v46
    } else {
        v47 := GenS3{F0: "", F1: (v2 * 868)}
        for fi48 := 0; fi48 < 10; fi48++ {
            var v49 uint8 = 128
            v2 = v2
            _ = fi48
            _ = v49
        }
        _ = len(v33)
        _ = v47
    }
    if false {
        v1 += -2.25
    }
    var v50 uint64 = 4294967295
    var v51 int8 = -61
    var v52 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v4:", v4)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v50
    _ = v51
    _ = v52
}
