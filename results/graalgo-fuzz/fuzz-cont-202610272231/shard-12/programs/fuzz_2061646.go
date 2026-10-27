package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint64; F1 int64 }
type GenS6 struct { F0 int; F1 uint16 }
func genFn0(p0 bool, p1 uint32) uint16 {
    for fi7 := 0; fi7 < 1; fi7++ {
        p0 = true
        fmt.Println(p0)
        _ = fi7
    }
    v8 := GenS1{F0: 866}
    var v9 float32 = 0.0
    _ = p0
    _ = p1
    _ = v8
    _ = v9
    return 0
}
func genFn1() int16 {
    if false {
        if false {
            var v10 uint8 = 0
            _ = v10
        }
        var v11 float64 = (float64((-2.25 / float32((1.5 + (0.0 / 1e10))))) / (3.14159 + -2.25))
        if true {
            if true {
                if false {
                    var v12 int64 = 9223372036854775807
                    _ = v12
                }
            } else {
                v13 := []int{}
                switch 724 {
                case 0:
                    var v14 uint64 = 42
                    _ = v14
                case 266:
                    if false {
                        v11 = float64(0.0)
                        if (68 >= 468) {
                            v11 += 1.5
                            var v15 bool = (!(!(42 == 0)))
                            _ = v15
                        } else {
                            for fi16 := 0; fi16 < 9; fi16++ {
                                var v17 uint64 = 18446744073709551615
                                v13 = append(v13, -9223372036854775808)
                                _ = fi16
                                _ = v17
                            }
                        }
                    } else {
                        fmt.Println(v11)
                    }
                case -440:
                    if (1 >= 959) {
                        fmt.Println(v11)
                        if true {
                            v18 := GenS1{F0: 2147483648}
                            _ = v18
                        } else {
                            var v19 int8 = 60
                            _ = v19
                        }
                    } else {
                        v13 = append(v13, -827)
                        for fi20 := 0; fi20 < 18; fi20++ {
                            for fi21 := 0; fi21 < 8; fi21++ {
                                switch ("foo bar" + "hello") {
                                case "a":
                                    v22 := GenS1{F0: 743}
                                    _ = v22
                                }
                                v23 := GenS2{F0: -32768}
                                var v24 string = (("\n\t" + "hello") + "\n\t")
                                _ = fi21
                                _ = v23
                                _ = v24
                            }
                            _ = fi20
                        }
                    }
                default:
                    if (4294967295 < 2) {
                        v11 += -2.25
                        v25 := GenS1{F0: 2147483648}
                        var v26 float32 = ((1e10 / 1.5) * (0.0 * 1e10))
                        _ = v25
                        _ = v26
                    }
                }
                var v27 uint = 0
                _ = v13
                _ = v27
            }
        }
        _ = v11
    }
    for fi28 := 0; fi28 < 19; fi28++ {
        fmt.Println(fi28)
        var v29 string = ("\n\t" + "a")
        _ = fi28
        _ = v29
    }
    if (65535 > 18446744073709551615) {
        v30 := []int{-2147483648, 924, -769, 0}
        var v31 int8 = 0
        _ = v30
        _ = v31
    } else {
        if ((true || (42 < -914)) || (3.14159 > (0.0 + (0.0 + 0.1)))) {
            if false {
                var v32 bool = false
                _ = v32
            } else {
                var v33 uint32 = 884
                v34 := []int{int(v33), 9223372036854775807}
                _ = v33
                _ = v34
            }
        }
        var v35 string = "foo bar"
        _ = v35
    }
    v36 := GenS3{F0: 42}
    var v37 bool = false
    _ = v36
    _ = v37
    return -32768
}

func main() {
    var v1 uint32 = 344
    var v2 float64 = (3.14159 + 1.5)
    var v3 int = 1
    var v4 float32 = 0.1
    v5 := make(map[string]int)
    var v6 int16 = -494
    var v38 int8 = -128
    switch v3 {
    case 42:
        v39 := []int{0, (v3 << ((285) & 63))}
        _ = v39
    case 1:
        for fi40 := 0; fi40 < 5; fi40++ {
            var v41 int64 = -303
            v5[("\n\t" + "foo bar")] = 2147483647
            _ = fi40
            _ = v41
        }
    default:
        var v42 uint32 = (v1 | 48)
        _ = v42
    }
    var v43 uint16 = genFn0((!(!false)), (v1 - 4294967295))
    var v44 uint32 = 193
    if ((^(v44)) == (v1 << ((50) & 63))) {
        for fi45 := 0; fi45 < 18; fi45++ {
            var v46 uint32 = v1
            v47 := GenS1{F0: (^(v1))}
            delete(v5, "foo bar")
            _ = fi45
            _ = v46
            _ = v47
        }
        v48 := GenS3{F0: 1}
        _ = v48
    } else {
        var v49 string = "foo bar"
        if (!true) {
            fmt.Println(v2, v6)
            v43 -= v43
        }
        _ = v49
    }
    fmt.Println(v3)
    _ = len(v5)
    var v50 int8 = 1
    for fi51 := 0; fi51 < 1; fi51++ {
        v52 := GenS4{F0: -223}
        _ = fi51
        _ = v52
    }
    v53 := GenS3{F0: 1}
    v1 = (v44 | v44)
    var v54 int64 = -1
    v44 = 0
    v55 := GenS5{F0: 4294967295, F1: 351}
    var v56 string = ("hello" + ("foo bar" + "foo bar"))
    v57 := GenS6{F0: (v3 << ((829) & 63)), F1: v43}
    var v58 uint32 = (v44 << ((4294967295) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("v54:", v54)
    fmt.Println("v56:", v56)
    fmt.Println("v58:", v58)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v50
    _ = v53
    _ = v54
    _ = v55
    _ = v56
    _ = v57
    _ = v58
    _ = v6
}
