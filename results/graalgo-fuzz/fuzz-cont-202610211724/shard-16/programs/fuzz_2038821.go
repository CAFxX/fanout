package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32; F1 int32 }
type GenS4 struct { F0 int8; F1 int16; F2 int8 }
type GenS5 struct { F0 int64 }
func genFn0(p0 int16, p1 uint32) uint8 {
    if true {
        v5 := []int{}
        p1 = 42
        _ = v5
    } else {
        var v6 int = -99
        _ = v6
    }
    v7 := []int{}
    v7 = append(v7, 54)
    _ = p0
    _ = p1
    _ = v7
    return 127
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int64 = 440
    v3 := []uint32{42, 65535, 42}
    var v4 int = 1
    v8 := GenS1{F0: v2, F1: 2}
    if false {
        var v9 int8 = -18
        var v10 uint64 = 0
        var v11 uint32 = 379
        _ = v10
        _ = v11
        _ = v9
    }
    switch ("hello" + "\n\t") {
    case "foo bar":
        if false {
            if true {
                for fi12 := 0; fi12 < 18; fi12++ {
                    for fi13 := 0; fi13 < 13; fi13++ {
                        v14 := GenS2{F0: 18446744073709551615}
                        if true {
                            v2 = v2
                            var v15 string = ("\n\t" + "")
                            _ = v15
                        }
                        v2 += v2
                        _ = fi13
                        _ = v14
                    }
                    v2 = (^(v2))
                    if len(v3) > 0 {
                        v3[0] = 4294967295
                    }
                    _ = fi12
                }
                var v16 string = "\n\t"
                v4 = 409
                _ = v16
            } else {
                if false {
                    var v17 bool = true
                    if (4294967295 > 4294967295) {
                        v18 := GenS3{F0: 2147483647, F1: 0}
                        _ = v18
                    } else {
                        if (-32768 == 1000) {
                            v19 := GenS4{F0: 100, F1: -32768, F2: 100}
                            var v20 uint16 = 287
                            _ = v19
                            _ = v20
                        } else {
                            fmt.Println(v2)
                        }
                        v21 := GenS2{F0: 17}
                        _ = v21
                    }
                    v1[1] = ((("\n\t" + "hello") + "\n\t") + ("hello" + "\n\t"))
                    _ = v17
                } else {
                    v22 := GenS5{F0: (-(v2))}
                    v2 = (v2 % ((v2) | 1))
                    var v23 float64 = -2.25
                    _ = v22
                    _ = v23
                }
                if (!true) {
                    v2 = 1
                }
            }
            var v24 uint16 = 662
            if (!false) {
                if false {
                    for fi25 := 0; fi25 < 20; fi25++ {
                        var v26 uint = 42
                        v27 := GenS3{F0: 42, F1: 42}
                        v26 = uint(v24)
                        _ = fi25
                        _ = v26
                        _ = v27
                    }
                }
                var v28 int64 = (v2 + 42)
                _ = v28
            } else {
                if false {
                    switch ("\n\t" + "\n\t") {
                    case "hello":
                        var v29 uint = 318
                        _ = v29
                    }
                } else {
                    var v30 int64 = v2
                    _ = v30
                }
                v31 := GenS2{F0: uint(v24)}
                for fi32 := 0; fi32 < 11; fi32++ {
                    fmt.Println(v4, fi32)
                    _ = fi32
                }
                _ = v31
            }
            _ = v24
        } else {
            v2 += (^(v2))
            var v33 int8 = 51
            fmt.Println(v2, v33, v2)
            _ = v33
        }
    case "\n\t":
        _ = len(v1)
    case "a":
        var v34 int16 = 1
        _ = v34
    }
    var v35 float64 = (3.14159 + 3.14159)
    v4 += int(v2)
    var v36 float64 = v35
    fmt.Println(v35, v35)
    var v37 int64 = (v2 * v2)
    for fi38 := 0; fi38 < 19; fi38++ {
        v4 -= v4
        _ = fi38
    }
    var v39 int = v4
    v35 += v36
    for fi40 := 0; fi40 < 17; fi40++ {
        if (0 >= 127) {
            var v41 int64 = -9223372036854775808
            v4 = 4294967295
            fmt.Println(v41, v35, v36)
            _ = v41
        }
        _ = fi40
    }
    for fi42 := 0; fi42 < 2; fi42++ {
        var v43 float32 = 0.1
        _ = fi42
        _ = v43
    }
    v37 = v2
    if (!(65535 < 0)) {
        var v44 float64 = 0.0
        v35 = 0.0
        _ = v44
    } else {
        if (100 >= 46) {
            var v45 int8 = -16
            _ = v45
        } else {
            if false {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
            } else {
                v39 = v4
            }
        }
    }
    var v46 float32 = float32(v35)
    var v47 int16 = 1000
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v39
    _ = v4
    _ = v46
    _ = v47
    _ = v8
}
