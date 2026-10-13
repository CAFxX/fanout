package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 string }
type GenS2 struct { F0 int; F1 bool; F2 string }
func genFn0(p0 uint8, p1 string, p2 int32) uint64 {
    if false {
        if (42 >= 1) {
            if false {
                fmt.Println(p0)
                p2 = (p2 % (((p2 & p2)) | 1))
                v6 := []int{}
                _ = v6
            } else {
                var v7 int8 = 42
                _ = v7
            }
        }
    } else {
        if false {
            v8 := GenS1{F0: 1, F1: 42, F2: "foo bar"}
            var v9 int16 = -32768
            for fi10 := 0; fi10 < 13; fi10++ {
                if (1 <= -605) {
                    switch (fi10 + (fi10 ^ (fi10 >> ((1) & 63)))) {
                    case 0:
                        v9 -= v9
                    case -9223372036854775808:
                        fmt.Println(p0)
                    case 633:
                        p0 = p0
                    default:
                        for fi11 := 0; fi11 < 19; fi11++ {
                            fmt.Println(fi11)
                            if (p0 > (p0 * p0)) {
                                v12 := GenS1{F0: 1, F1: -1, F2: ("a" + "foo bar")}
                                if true {
                                    v13 := []int{fi10, 0, fi10, (fi10 / ((int(p0)) | 1))}
                                    if len(v13) >= 2 {
                                        v13 = v13[1:len(v13)]
                                    }
                                    p2 = -1
                                    _ = v13
                                }
                                _ = v12
                            }
                            if (289 >= 65535) {
                                var v14 float64 = -2.25
                                v9 = (v9 - v9)
                                _ = v14
                            }
                            _ = fi11
                        }
                    }
                    fmt.Println(p1)
                    if (!(4294967295 >= 2147483647)) {
                        v15 := []int{fi10}
                        v16 := GenS2{F0: fi10, F1: (0.1 != 1e10), F2: "\n\t"}
                        _ = v15
                        _ = v16
                    }
                } else {
                    v17 := GenS1{F0: 762, F1: 1, F2: (p1 + "hello")}
                    _ = v17
                }
                var v18 float32 = 1e10
                switch p1 {
                case "\n\t":
                    var v19 int16 = (v9 << ((577) & 63))
                    _ = v19
                }
                _ = fi10
                _ = v18
            }
            _ = v8
            _ = v9
        } else {
            var v20 bool = false
            var v21 int64 = 1
            _ = v20
            _ = v21
        }
        p0 = p0
        var v22 uint64 = 42
        _ = v22
    }
    v23 := GenS1{F0: -605, F1: 577, F2: (p1 + (p1 + "\n\t"))}
    v24 := GenS1{F0: 1, F1: 0, F2: (p1 + "foo bar")}
    var v25 uint32 = 270
    _ = p0
    _ = p1
    _ = p2
    _ = v23
    _ = v24
    _ = v25
    return 1
}
func genFn1() int32 {
    var v26 int32 = -460
    var v27 uint = 466
    _ = v26
    _ = v27
    return v26
}

func main() {
    var v1 uint16 = 255
    var v2 float64 = 1.5
    var v3 int64 = -9223372036854775808
    var v4 uint16 = v1
    v5 := []string{}
    switch int(v1) {
    case -1:
        switch "foo bar" {
        case "a":
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        case "":
            v4 += v1
        default:
            var v28 int64 = v3
            _ = v28
        }
    case 0:
        var v29 uint64 = 2
        _ = v29
    case -543:
        fmt.Println(v4)
    default:
        v5 = append(v5, "")
    }
    var v30 int32 = genFn1()
    if true {
        v30 += v30
    } else {
        v2 = (v2 * v2)
    }
    v2 = v2
    v31 := GenS2{F0: 1, F1: false, F2: ""}
    v4 = (^(v1))
    v2 -= 1.5
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if true {
        var v32 float64 = (((v2 * (v2 * v2)) / 1.0) - -2.25)
        _ = v32
    } else {
        if (0 > 979) {
            if len(v5) > 0 {
                v5[0] = "\n\t"
            }
            _ = len(v5)
        } else {
            var v33 uint64 = 1
            var v34 string = "foo bar"
            v1 = v4
            _ = v33
            _ = v34
        }
        var v35 float32 = 1e10
        v30 = v30
        _ = v35
    }
    for fi36 := 0; fi36 < 1; fi36++ {
        for fi37 := 0; fi37 < 3; fi37++ {
            if false {
                var v38 float32 = 0.0
                _ = v38
            }
            _ = fi37
        }
        _ = fi36
    }
    var v39 int32 = v30
    fmt.Println(v2)
    _ = len(v5)
    var v40 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v39
    _ = v4
    _ = v40
    _ = v5
}
