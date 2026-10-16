package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 bool; F1 int64; F2 uint64 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int8; F1 uint16; F2 int32 }
func genFn0() int {
    v7 := GenS1{F0: 494}
    var v8 uint16 = 255
    switch 691 {
    case 32767:
        fmt.Println(v8)
    case 0:
        if false {
            fmt.Println(v8)
            for fi9 := 0; fi9 < 20; fi9++ {
                fmt.Println(v8, v8)
                if (1 != 1) {
                    v8 = v8
                    v8 -= 693
                    fmt.Println(v8, fi9)
                }
                if (9223372036854775807 < (^(fi9))) {
                    for fi10 := 0; fi10 < 7; fi10++ {
                        var v11 int32 = 5
                        fmt.Println(v11)
                        v12 := GenS1{F0: v11}
                        _ = fi10
                        _ = v11
                        _ = v12
                    }
                    var v13 bool = (!true)
                    v8 -= (v8 & v8)
                    _ = v13
                }
                _ = fi9
            }
        }
    default:
        switch ("hello" + ("a" + (("a" + ("" + "\n\t")) + ("a" + "\n\t")))) {
        case "":
            var v14 int32 = 2147483647
            _ = v14
        case "foo bar":
            v8 = v8
        }
    }
    _ = v7
    _ = v8
    return 32768
}

func main() {
    var v1 string = "\n\t"
    var v2 uint32 = 4294967295
    var v3 uint32 = 42
    var v4 int32 = 42
    v5 := make(map[int]uint32)
    var v6 uint32 = v3
    v5[-806] = v3
    var v15 string = "foo bar"
    v15 = v15
    var v16 int = 1
    if (1.5 >= 0.0) {
        v1 = v1
        var v17 uint32 = (^(v2))
        _ = v17
    }
    delete(v5, genFn0())
    var v18 int = (v16 + 607)
    var v19 uint = 578
    var v20 int = v18
    v21 := GenS2{F0: false, F1: -1, F2: 18446744073709551615}
    for fi22 := 0; fi22 < 16; fi22++ {
        v23 := GenS1{F0: v4}
        if false {
            if false {
                v24 := []int{(v18 * 4), fi22, (v20 ^ 2147483647), (v18 & fi22)}
                v25 := GenS2{F0: false, F1: 366, F2: 0}
                for fi26 := 0; fi26 < 7; fi26++ {
                    v4 = (^(v4))
                    v18 += fi22
                    v15 = ("hello" + v1)
                    _ = fi26
                }
                _ = v24
                _ = v25
            } else {
                for fi27 := 0; fi27 < 7; fi27++ {
                    var v28 float64 = 3.14159
                    if (!false) {
                        v29 := v5[genFn0()]
                        _ = v29
                    }
                    var v30 int = v20
                    _ = fi27
                    _ = v28
                    _ = v30
                }
                if true {
                    for fi31 := 0; fi31 < 17; fi31++ {
                        fmt.Println(v1, v20, fi22)
                        _ = fi31
                    }
                    var v32 uint32 = (v6 ^ v3)
                    _ = v32
                } else {
                    var v33 int = (fi22 + 128)
                    v3 -= v6
                    v34 := GenS3{F0: 0}
                    _ = v33
                    _ = v34
                }
            }
            var v35 uint16 = 2
            if false {
                v16 += (fi22 - v20)
            }
            _ = v35
        } else {
            var v36 string = (v1 + "foo bar")
            if true {
                switch v16 {
                case 128:
                    var v37 float32 = 1e10
                    _ = v37
                case 4294967295:
                    for fi38 := 0; fi38 < 1; fi38++ {
                        fmt.Println(v18, v18, v6)
                        switch -493 {
                        case 111:
                            v39 := GenS4{F0: 41, F1: 399, F2: v4}
                            _ = v39
                        case 1:
                            var v40 int = v20
                            _ = v40
                        case 64:
                            if (v19 >= (v19 | (^(v19)))) {
                                v16 = 1
                                var v41 int = (^(fi22))
                                _ = v41
                            }
                        }
                        switch v36 {
                        case "a":
                            var v42 int = 0
                            _ = v42
                        case "foo bar":
                            var v43 uint64 = 440
                            _ = v43
                        default:
                            var v44 float32 = 1.5
                            _ = v44
                        }
                        _ = fi38
                    }
                default:
                    var v45 uint64 = 784
                    _ = v45
                }
                v18 = (-(v18))
                var v46 int16 = -1
                _ = v46
            }
            _ = v36
        }
        for fi47 := 0; fi47 < 5; fi47++ {
            var v48 int8 = int8(v6)
            _ = fi47
            _ = v48
        }
        _ = fi22
        _ = v23
    }
    v49 := []int{(v16 + v20)}
    var v50 int16 = 1000
    var v51 int64 = 0
    _ = len(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v49:", len(v49), v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v16
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v6
}
