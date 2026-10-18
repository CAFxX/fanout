package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 int16; F1 int32; F2 int8 }
type GenS3 struct { F0 uint }
func genFn0() int8 {
    var v6 int64 = 933
    v6 = 42
    var v7 uint16 = 255
    v8 := []int{1, -868, 32767, -9223372036854775808}
    _ = v6
    _ = v7
    _ = v8
    return -1
}
func genFn1(p0 bool) uint {
    switch 382 {
    case 127:
        var v9 uint32 = 42
        _ = v9
    case 0:
        switch "foo bar" {
        case "a":
            switch 67 {
            case 1:
                v10 := GenS1{F0: 2147483648, F1: 1}
                _ = v10
            case 32767:
                p0 = (128 > 127)
            case 431:
                for fi11 := 0; fi11 < 16; fi11++ {
                    for fi12 := 0; fi12 < 9; fi12++ {
                        if (!(42 < 42)) {
                            v13 := []int{-193, fi11, fi12}
                            var v14 int32 = 0
                            _ = v13
                            _ = v14
                        }
                        _ = fi12
                    }
                    switch (fi11 & 42) {
                    case -520:
                        p0 = false
                    case 128:
                        if p0 {
                            v15 := GenS1{F0: 29, F1: 2}
                            _ = v15
                        }
                    case -9223372036854775808:
                        fmt.Println(p0, p0)
                    default:
                        var v16 int32 = 948
                        _ = v16
                    }
                    _ = fi11
                }
            }
        default:
            fmt.Println(p0)
        }
    case -577:
        v17 := GenS1{F0: 2, F1: 986}
        _ = v17
    default:
        if (((0 == 619) && true) || (-1 >= -78)) {
            var v18 uint = 65535
            v18 -= v18
            for fi19 := 0; fi19 < 16; fi19++ {
                if (!(26 >= -32768)) {
                    for fi20 := 0; fi20 < 6; fi20++ {
                        v21 := []int{-1, fi19, (fi19 / ((fi19) | 1))}
                        fmt.Println(fi20, fi20, fi19)
                        p0 = false
                        _ = fi20
                        _ = v21
                    }
                    v22 := GenS1{F0: 42, F1: v18}
                    switch (("a" + "\n\t") + "hello") {
                    case "\n\t":
                        for fi23 := 0; fi23 < 4; fi23++ {
                            v24 := []int{}
                            if p0 {
                                var v25 float32 = -2.25
                                if true {
                                    if len(v24) > 0 {
                                        v24[0] = fi23
                                    }
                                    v24 = append(v24, fi19)
                                } else {
                                    v26 := GenS1{F0: 2147483647, F1: (v18 & v18)}
                                    if false {
                                        if len(v24) > 0 {
                                            v24[0] = 32767
                                        }
                                        var v27 uint8 = 2
                                        _ = v27
                                    } else {
                                        p0 = true
                                        if (p0 && true) {
                                            v25 -= v25
                                        }
                                    }
                                    _ = v26
                                }
                                _ = v25
                            }
                            _ = fi23
                            _ = v24
                        }
                    case "hello":
                        var v28 float64 = -2.25
                        _ = v28
                    case "foo bar":
                        p0 = (!(!p0))
                    default:
                        if true {
                            if p0 {
                                v18 = (^(v18))
                                v29 := []int{9223372036854775807}
                                _ = v29
                            }
                        }
                    }
                    _ = v22
                } else {
                    p0 = p0
                }
                if ((p0 && p0) || (false && p0)) {
                    v30 := []int{(fi19 - (fi19 >> (((v18 / ((v18) | 1))) & 63)))}
                    v31 := GenS1{F0: 2147483647, F1: (^(v18))}
                    _ = v30
                    _ = v31
                }
                _ = fi19
            }
            _ = v18
        }
    }
    v32 := []int{}
    _ = p0
    _ = v32
    return 173
}

func main() {
    v1 := []int{}
    var v2 int64 = -1
    var v3 int = 9223372036854775807
    var v4 uint16 = 65535
    var v5 uint8 = 1
    var v33 int32 = -198
    fmt.Println(v5, v5, v5)
    v5 -= 0
    if false {
        v1 = append(v1, int(v2))
    } else {
        v34 := GenS1{F0: 4, F1: 0}
        _ = v34
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if ((1e10 + -2.25) >= 3.14159) {
        var v35 string = "\n\t"
        _ = v35
    }
    v36 := GenS2{F0: 1, F1: v33, F2: -1}
    _ = len(v1)
    v37 := GenS3{F0: 2}
    if true {
        if (-207 > 284) {
            var v38 uint32 = 65535
            _ = v38
        }
        v39 := GenS1{F0: 2147483648, F1: genFn1(true)}
        v5 -= v5
        _ = v39
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
