package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string; F1 uint16; F2 uint8 }
type GenS3 struct { F0 uint64; F1 uint64; F2 int64 }
func genFn0(p0 uint8, p1 bool) int32 {
    p1 = p1
    fmt.Println(p1, p1)
    fmt.Println(p0, p1)
    p1 = false
    _ = p0
    _ = p1
    return int32(p0)
}

func main() {
    var v1 int32 = 42
    var v2 int64 = 0
    var v3 int = -415
    var v4 uint32 = 922
    var v5 uint64 = 168
    var v6 uint64 = 2
    if true {
        v7 := GenS1{F0: 0}
        if ((v6 <= (v5 ^ 862)) || (!true)) {
            v8 := []int{(-(v3)), (v3 - v3), (v3 ^ v3), (v3 | 4294967295)}
            _ = v8
        } else {
            switch ("\n\t" + "hello") {
            case "hello":
                v9 := GenS1{F0: 42}
                _ = v9
            case "":
                v10 := GenS2{F0: "\n\t", F1: 255, F2: 185}
                _ = v10
            }
            for fi11 := 0; fi11 < 2; fi11++ {
                fmt.Println(v1, v3, v2)
                v12 := GenS2{F0: ("" + "a"), F1: 443, F2: 147}
                _ = fi11
                _ = v12
            }
        }
        var v13 float32 = -2.25
        _ = v13
        _ = v7
    } else {
        if true {
            for fi14 := 0; fi14 < 3; fi14++ {
                v15 := []int{-1, fi14, (v3 % ((4294967295) | 1))}
                _ = fi14
                _ = v15
            }
            v6 = 438
            var v16 int8 = -95
            _ = v16
        } else {
            fmt.Println(v4)
            var v17 uint = 314
            _ = v17
        }
        v18 := []int{v3}
        _ = v18
    }
    v19 := []int{v3, 2147483647, (v3 | (v3 * 541)), (-(v3))}
    v20 := GenS1{F0: 1}
    switch 953 {
    case 127:
        for fi21 := 0; fi21 < 15; fi21++ {
            var v22 float32 = float32(float64((0.0 * (float32(float64(-2.25)) * 0.0))))
            _ = fi21
            _ = v22
        }
    case 368:
        switch ("foo bar" + "") {
        case "":
            v23 := GenS2{F0: "a", F1: 42, F2: 33}
            _ = v23
        }
    case -517:
        switch (v3 ^ -59) {
        case 35:
            if (1 < 1000) {
                if (3.14159 >= (0.0 / 1e10)) {
                    for fi24 := 0; fi24 < 18; fi24++ {
                        v1 -= genFn0(239, false)
                        if len(v19) >= 2 {
                            v19 = v19[1:len(v19)]
                        }
                        fmt.Println(v4, v2, v1)
                        _ = fi24
                    }
                } else {
                    if len(v19) > 0 {
                        v19[0] = (v3 | (v3 | (v3 ^ (v3 | 662))))
                    }
                    v19 = append(v19, v3)
                    v25 := GenS1{F0: 100}
                    _ = v25
                }
            } else {
                var v26 string = "foo bar"
                for fi27 := 0; fi27 < 9; fi27++ {
                    if len(v19) > 0 {
                        v19[0] = (-(v3))
                    }
                    for fi28 := 0; fi28 < 20; fi28++ {
                        if (0.0 >= float32(((float64(0.0) + 1e10) * 1.5))) {
                            var v29 uint8 = 1
                            _ = v29
                        } else {
                            fmt.Println(v5)
                            var v30 int32 = int32(v4)
                            _ = v30
                        }
                        var v31 uint = 0
                        if false {
                            if false {
                                var v32 uint32 = v4
                                v19 = append(v19, (fi28 | v3))
                                var v33 int8 = -128
                                _ = v32
                                _ = v33
                            } else {
                                if true {
                                    var v34 uint8 = 66
                                    if len(v19) >= 2 {
                                        v19 = v19[1:len(v19)]
                                    }
                                    _ = v34
                                }
                                if len(v19) > 0 {
                                    v19[0] = fi28
                                }
                                fmt.Println(v31, v26)
                            }
                            switch "\n\t" {
                            case "\n\t":
                                v2 += (-(v2))
                                var v35 uint = v31
                                _ = v35
                            default:
                                if true {
                                    fmt.Println(v1, v2, v1)
                                    var v36 int32 = (v1 - v1)
                                    var v37 uint16 = 255
                                    _ = v36
                                    _ = v37
                                }
                            }
                            var v38 int16 = -32768
                            _ = v38
                        }
                        _ = fi28
                        _ = v31
                    }
                    _ = fi27
                }
                _ = v26
            }
        case 32768:
            if len(v19) >= 2 {
                v19 = v19[1:len(v19)]
            }
        case 9223372036854775807:
            fmt.Println(v4, v4, v2)
        default:
            if ((v1 << ((4294967295) & 63)) < 161) {
                if (!(v4 >= (v4 % (((v4 & 686)) | 1)))) {
                    var v39 int = (^(v3))
                    v3 += v3
                    v39 = 0
                    _ = v39
                }
                fmt.Println(v4, v5, v2)
            } else {
                var v40 bool = false
                _ = v40
            }
        }
    default:
        v4 = v4
    }
    v41 := GenS1{F0: 0}
    var v42 uint16 = 0
    var v43 float32 = 1e10
    v5 = 42
    var v44 uint = 378
    v45 := GenS3{F0: 0, F1: v5, F2: (v2 | (v2 * int64(v4)))}
    v43 -= (((v43 - 0.0) * 0.0) + v43)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
