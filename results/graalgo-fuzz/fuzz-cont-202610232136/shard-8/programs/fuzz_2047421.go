package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint32; F1 int16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int32 }

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint64 = 4294967295
    var v3 uint32 = 695
    var v4 uint8 = 49
    var v5 int32 = 652
    v6 := []int{118, -2147483648}
    delete(v1, v3)
    v3 = uint32(v2)
    if (true && (0 <= 711)) {
        v3 -= v3
        v7 := GenS1{F0: (v3 << ((2) & 63))}
        _ = len(v6)
        _ = v7
    }
    var v8 float64 = (-2.25 - 3.14159)
    fmt.Println(v3)
    if true {
        if (true && ((58 < uint(v2)) || (true || (true && true)))) {
            v9 := GenS2{F0: v3, F1: 1000}
            _ = v9
        }
        for fi10 := 0; fi10 < 18; fi10++ {
            v2 -= 0
            for fi11 := 0; fi11 < 11; fi11++ {
                v8 += 1.5
                var v12 int64 = 0
                v3 -= (v3 / ((42) | 1))
                var v13 int64 = v12
                var v14 uint32 = (^(v3))
                _ = fi11
                _ = v12
                _ = v13
                _ = v14
            }
            _ = fi10
        }
        delete(v1, (v3 + v3))
    } else {
        if false {
            v5 = v5
            if (v8 <= (0.0 + (0.0 * float64(0.0)))) {
                if ((("a" == "foo bar") || (2 != 65535)) && ("hello" != "a")) {
                    fmt.Println(v3)
                    v3 += uint32(v2)
                    if (2 <= 1) {
                        for fi15 := 0; fi15 < 7; fi15++ {
                            v5 -= 2147483647
                            _ = fi15
                        }
                        var v16 uint = 4294967295
                        v8 -= (v8 * v8)
                        _ = v16
                    } else {
                        for fi17 := 0; fi17 < 4; fi17++ {
                            if (false && (true && (!((0.1 * float32((float64(0.0) - ((v8 - v8) - v8)))) != (3.14159 * 0.1))))) {
                                if true {
                                    fmt.Println(v8, v5)
                                } else {
                                    switch "foo bar" {
                                    case "foo bar":
                                        switch ("foo bar" + "") {
                                        case "":
                                            var v18 uint8 = 136
                                            _ = v18
                                        case "a":
                                            v19 := GenS2{F0: 560, F1: -458}
                                            _ = v19
                                        }
                                    case "":
                                        for fi20 := 0; fi20 < 17; fi20++ {
                                            v21 := GenS1{F0: (v3 - (v3 * uint32(v4)))}
                                            _ = fi20
                                            _ = v21
                                        }
                                    }
                                }
                                var v22 bool = false
                                _ = v22
                            }
                            v4 = (v4 | v4)
                            _ = fi17
                        }
                        v23 := GenS3{F0: (v5 << ((4294967295) & 63))}
                        _ = v23
                    }
                } else {
                    for fi24 := 0; fi24 < 4; fi24++ {
                        var v25 uint32 = v3
                        var v26 int32 = (v5 * v5)
                        v26 = v5
                        _ = fi24
                        _ = v25
                        _ = v26
                    }
                }
            }
        }
        var v27 int8 = 0
        var v28 float32 = (-2.25 / 0.1)
        _ = v27
        _ = v28
    }
    switch 4294967295 {
    case -969:
        if (!((!false) || true)) {
            if false {
                delete(v1, 0)
                v29 := GenS2{F0: 243, F1: 1}
                for fi30 := 0; fi30 < 20; fi30++ {
                    v2 += v2
                    var v31 string = ("a" + "")
                    var v32 bool = (-128 != 42)
                    _ = fi30
                    _ = v31
                    _ = v32
                }
                _ = v29
            } else {
                fmt.Println(v5, v5)
                var v33 bool = true
                var v34 uint64 = 1
                _ = v33
                _ = v34
            }
            v35 := GenS3{F0: -1}
            if (42 < 42) {
                v4 -= v4
                var v36 uint32 = (v3 / ((0) | 1))
                _ = v36
            } else {
                v37 := GenS4{F0: (v5 ^ v5)}
                _ = v37
            }
            _ = v35
        }
    case -1:
        v4 -= 86
    }
    _ = len(v6)
    v5 = -1
    var v38 uint16 = 0
    switch -886 {
    case 4294967295:
        if true {
            var v39 uint = 2
            _ = v39
        }
    case 9223372036854775807:
        var v40 int = 127
        _ = v40
    default:
        var v41 int32 = v5
        _ = v41
    }
    var v42 int64 = 9223372036854775807
    v43 := GenS1{F0: (^(v3))}
    var v44 string = "a"
    var v45 int16 = -1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v6
    _ = v8
}
