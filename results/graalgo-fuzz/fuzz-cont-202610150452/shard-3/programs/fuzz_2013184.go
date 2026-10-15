package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 int; F1 uint }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint16; F1 uint32; F2 int64 }
func genFn0(p0 uint64, p1 int) int32 {
    p1 -= p1
    for fi6 := 0; fi6 < 4; fi6++ {
        if (!false) {
            var v7 uint64 = 42
            v8 := GenS1{F0: 925, F1: true}
            fmt.Println(p1, p0)
            _ = v7
            _ = v8
        } else {
            var v9 int64 = 42
            _ = v9
        }
        v10 := []int{}
        _ = fi6
        _ = v10
    }
    var v11 int8 = 82
    p0 += p0
    fmt.Println(p1, p0, p0)
    _ = p0
    _ = p1
    _ = v11
    return 1
}
func genFn1() uint8 {
    if true {
        if (1 != 0) {
            if (!true) {
                v12 := []int{}
                var v13 string = ("a" + ("" + ("hello" + ("\n\t" + "\n\t"))))
                fmt.Println(v13)
                _ = v12
                _ = v13
            }
        } else {
            var v14 string = "\n\t"
            _ = v14
        }
        if (!true) {
            var v15 string = ""
            _ = v15
        }
        if true {
            var v16 int8 = 18
            _ = v16
        }
    }
    var v17 int8 = 31
    if false {
        if false {
            v18 := []int{-810}
            v17 += v17
            var v19 float64 = 0.0
            _ = v18
            _ = v19
        } else {
            var v20 int16 = 252
            v17 += 100
            v17 += v17
            _ = v20
        }
    }
    _ = v17
    return 128
}

func main() {
    v1 := make(map[string]string)
    v2 := []int{127}
    var v3 float32 = float32(1e10)
    var v4 float64 = -2.25
    var v5 int32 = 293
    if (83 >= 32767) {
        var v21 uint = 679
        if false {
            switch -991 {
            case 127:
                v4 += (v4 + -2.25)
            default:
                _ = len(v1)
            }
            var v22 uint8 = 40
            _ = v22
        }
        var v23 string = "foo bar"
        _ = v21
        _ = v23
    }
    var v24 string = "\n\t"
    fmt.Println(v3, v3, v4)
    v25 := GenS1{F0: v5, F1: true}
    fmt.Println(v24, v24, v24)
    if ((!false) || (2147483647 >= v5)) {
        v24 = "\n\t"
        for fi26 := 0; fi26 < 2; fi26++ {
            v4 = 1e10
            _ = fi26
        }
        v3 = v3
    }
    if (!(128 <= 42)) {
        for fi27 := 0; fi27 < 1; fi27++ {
            if (true && true) {
                v4 += 1.5
            } else {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            var v28 float32 = v3
            for fi29 := 0; fi29 < 16; fi29++ {
                v30 := GenS1{F0: genFn0(15, 139), F1: (v24 >= "foo bar")}
                fmt.Println(v3, v4)
                _ = fi29
                _ = v30
            }
            _ = fi27
            _ = v28
        }
        v31 := GenS2{F0: 2147483647, F1: 1}
        _ = v31
    } else {
        var v32 uint16 = 42
        _ = v32
    }
    if false {
        for fi33 := 0; fi33 < 13; fi33++ {
            var v34 uint = 714
            if false {
                v35 := GenS3{F0: 255}
                for fi36 := 0; fi36 < 19; fi36++ {
                    var v37 string = ((v24 + "foo bar") + v24)
                    switch (^(fi36)) {
                    case 2147483647:
                        if (v5 == 0) {
                            if len(v2) > 0 {
                                v2[0] = -2147483648
                            }
                        }
                    case 128:
                        if true {
                            var v38 int8 = 42
                            var v39 int8 = (^(v38))
                            v38 += int8(v34)
                            _ = v38
                            _ = v39
                        }
                        if ((fi33 >> ((v34) & 63)) != 565) {
                            v40 := GenS3{F0: 42}
                            v3 = v3
                            if ((v24 <= ("hello" + ("\n\t" + v37))) && ((0 != 1) || false)) {
                                v5 = -1
                                fmt.Println(v4, v5, fi33)
                                if false {
                                    var v41 int16 = 32767
                                    var v42 uint16 = 0
                                    var v43 uint16 = (v42 & v42)
                                    _ = v41
                                    _ = v42
                                    _ = v43
                                }
                            } else {
                                if (true || (127 < 0)) {
                                    v44 := GenS2{F0: fi36, F1: 424}
                                    _ = v44
                                } else {
                                    switch 42 {
                                    case 0:
                                        _ = len(v2)
                                    }
                                    var v45 uint8 = 255
                                    fmt.Println(v37, v34, v45)
                                    v46 := GenS4{F0: 0}
                                    _ = v45
                                    _ = v46
                                }
                                v34 = v34
                                fmt.Println(v34, v3)
                            }
                            v34 = 42
                            var v47 uint = (v34 ^ (v34 + v34))
                            v48 := GenS4{F0: 4294967295}
                            v49 := GenS5{F0: 32767, F1: 2147483648, F2: 250}
                            _ = v40
                            _ = v47
                            _ = v48
                            _ = v49
                        }
                        v50 := GenS1{F0: (v5 + (v5 | v5)), F1: true}
                        _ = v50
                    }
                    _ = fi36
                    _ = v37
                }
                _ = v35
            }
            _ = fi33
            _ = v34
        }
        fmt.Println(v5, v4)
        var v51 int64 = 522
        _ = v51
    }
    v52 := GenS3{F0: 200}
    fmt.Println(v4)
    var v53 uint64 = 106
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v53
}
