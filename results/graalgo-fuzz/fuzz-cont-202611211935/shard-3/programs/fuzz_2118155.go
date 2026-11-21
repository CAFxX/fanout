package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 bool; F1 int }
type GenS4 struct { F0 int8; F1 int }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 string }
func genFn0(p0 uint) string {
    fmt.Println(p0)
    var v5 uint = 314
    _ = p0
    _ = v5
    return "foo bar"
}
func genFn1() uint64 {
    var v6 int32 = 399
    for fi7 := 0; fi7 < 18; fi7++ {
        v8 := []int{}
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        var v9 int8 = 127
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 int64 = 0
    if (false && (32768 == 0)) {
        v6 -= (v6 * 0)
        for fi11 := 0; fi11 < 10; fi11++ {
            var v12 uint8 = 42
            _ = fi11
            _ = v12
        }
        switch 132 {
        case 370:
            var v13 bool = true
            _ = v13
        case 395:
            if true {
                var v14 int16 = 1000
                _ = v14
            } else {
                v6 = -2147483648
                v15 := []int{-2147483648}
                fmt.Println(v10, v6)
                _ = v15
            }
        }
    } else {
        v16 := GenS1{F0: 0}
        switch 343 {
        case 42:
            v17 := []int{450, 128}
            _ = v17
        case 0:
            fmt.Println(v10)
        }
        _ = v16
    }
    for fi18 := 0; fi18 < 11; fi18++ {
        v19 := []int{(fi18 + (fi18 ^ fi18)), (fi18 % ((fi18) | 1)), (^(fi18)), (fi18 ^ (fi18 << ((18446744073709551615) & 63)))}
        var v20 int16 = 1
        _ = fi18
        _ = v19
        _ = v20
    }
    _ = v10
    _ = v6
    return 772
}

func main() {
    var v1 int16 = 801
    v2 := make(map[int]string)
    var v3 int32 = 2147483647
    v4 := make(map[string]string)
    switch 128 {
    case -649:
        if false {
            v3 += -2147483648
            v3 = v3
        } else {
            if false {
                for fi21 := 0; fi21 < 14; fi21++ {
                    var v22 int8 = 12
                    _ = fi21
                    _ = v22
                }
                v1 += v1
            } else {
                v1 = v1
            }
            if false {
                for fi23 := 0; fi23 < 14; fi23++ {
                    v24 := GenS1{F0: 1}
                    _ = fi23
                    _ = v24
                }
            } else {
                if ((v1 & v1) > v1) {
                    switch 4294967295 {
                    case 42:
                        var v25 int8 = 1
                        _ = v25
                    case 4294967295:
                        fmt.Println(v3)
                    default:
                        v26 := GenS1{F0: 2}
                        _ = v26
                    }
                    v27 := GenS2{F0: 1}
                    _ = v27
                } else {
                    if true {
                        switch "" {
                        case "foo bar":
                            v3 = 42
                        case "hello":
                            fmt.Println(v3)
                        default:
                            var v28 float64 = -2.25
                            _ = v28
                        }
                        v29 := GenS3{F0: true, F1: 371}
                        _ = v29
                    } else {
                        v3 -= v3
                        v30 := []int{-9223372036854775808}
                        _ = v30
                    }
                    var v31 uint32 = 65535
                    _ = v31
                }
                var v32 uint8 = 155
                var v33 float32 = 3.14159
                _ = v32
                _ = v33
            }
        }
    default:
        var v34 bool = (1 > 1)
        _ = v34
    }
    v35 := []int{192, 384, 0, -2147483648}
    fmt.Println(v1, v3)
    if true {
        v3 = (v3 >> ((0) & 63))
        var v36 uint = 352
        _ = v36
    }
    v37 := GenS4{F0: 42, F1: -9223372036854775808}
    fmt.Println(v3, v3)
    if false {
        v38 := GenS4{F0: -117, F1: 0}
        _ = v38
    } else {
        v39 := GenS5{F0: (!(1e10 != 0.0))}
        if len(v35) > 0 {
            v35[0] = 440
        }
        _ = v39
    }
    v40 := v4[("" + "\n\t")]
    if (!false) {
        v41 := GenS1{F0: 65535}
        var v42 uint32 = 649
        _ = v41
        _ = v42
    } else {
        v1 -= v1
        for fi43 := 0; fi43 < 4; fi43++ {
            var v44 string = "a"
            var v45 int64 = 9223372036854775807
            v45 = (v45 & v45)
            _ = fi43
            _ = v44
            _ = v45
        }
        v40 = v40
    }
    for fi46 := 0; fi46 < 2; fi46++ {
        if true {
            v47 := GenS1{F0: 2}
            _ = v47
        } else {
            fmt.Println(v1)
            v4[("\n\t" + (v40 + ""))] = v40
        }
        if false {
            v40 = v40
        } else {
            for fi48 := 0; fi48 < 8; fi48++ {
                var v49 int32 = (v3 << ((42) & 63))
                v49 = v3
                v3 = (v3 | v3)
                v1 = -1
                _ = fi48
                _ = v49
            }
        }
        _ = fi46
    }
    if (!(v1 > (v1 | v1))) {
        v50 := GenS6{F0: v40}
        _ = v50
    } else {
        var v51 uint16 = 28
        v52 := GenS6{F0: ((v40 + ("a" + v40)) + (v40 + "a"))}
        if (!(2147483647 > -9223372036854775808)) {
            fmt.Println(v51)
            v1 -= (v1 + -32768)
            v1 = -974
        }
        _ = v51
        _ = v52
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v37
    _ = v4
    _ = v40
}
