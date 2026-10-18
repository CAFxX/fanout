package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8; F1 uint }
type GenS4 struct { F0 int32 }
func genFn0(p0 uint32) uint8 {
    p0 = (p0 % ((p0) | 1))
    p0 += 784
    _ = p0
    return 216
}

func main() {
    var v1 string = "\n\t"
    var v2 int8 = 42
    v3 := make(map[string]string)
    var v4 bool = false
    if true {
        var v5 uint = 42
        v2 += v2
        v6 := []int{-602, 127}
        _ = v5
        _ = v6
    } else {
        v1 = v1
        if false {
            var v7 float64 = (0.0 * 3.14159)
            _ = v7
        } else {
            v8 := []int{}
            switch v1 {
            case "hello":
                v9 := GenS1{F0: 2}
                _ = v9
            case "\n\t":
                v10 := GenS1{F0: 14}
                _ = v10
            case "a":
                var v11 uint32 = 4294967295
                _ = v11
            }
            var v12 float32 = (-2.25 + ((0.1 * (1e10 / 1e10)) * 1.5))
            _ = v12
            _ = v8
        }
        v4 = v4
    }
    if ((v2 * (v2 % ((42) | 1))) > v2) {
        for fi13 := 0; fi13 < 8; fi13++ {
            v14 := GenS1{F0: 2}
            v4 = v4
            var v15 float32 = 1.5
            _ = fi13
            _ = v14
            _ = v15
        }
        fmt.Println(v1, v4, v2)
        for fi16 := 0; fi16 < 1; fi16++ {
            var v17 string = ((v1 + v1) + v1)
            _ = fi16
            _ = v17
        }
    }
    v18 := GenS1{F0: 127}
    switch "foo bar" {
    case "a":
        var v19 int64 = -9223372036854775808
        _ = v19
    case "hello":
        if (18446744073709551615 >= 0) {
            v20 := GenS2{F0: 42}
            v21 := GenS3{F0: v2, F1: 0}
            fmt.Println(v1, v1)
            _ = v20
            _ = v21
        } else {
            if false {
                v22 := []int{127}
                if v4 {
                    if (4294967295 > 2) {
                        var v23 float32 = -2.25
                        fmt.Println(v1)
                        if len(v22) > 0 {
                            v22[0] = 32767
                        }
                        _ = v23
                    }
                    v24 := GenS1{F0: 224}
                    _ = v24
                } else {
                    var v25 int32 = -14
                    _ = v25
                }
                for fi26 := 0; fi26 < 9; fi26++ {
                    var v27 bool = v4
                    _ = fi26
                    _ = v27
                }
                _ = v22
            } else {
                v28 := GenS4{F0: 42}
                if v4 {
                    if v4 {
                        if (!v4) {
                            _ = len(v3)
                        }
                    }
                    switch (v1 + v1) {
                    case "":
                        for fi29 := 0; fi29 < 10; fi29++ {
                            var v30 uint32 = 1
                            var v31 int = -1
                            for fi32 := 0; fi32 < 15; fi32++ {
                                v4 = v4
                                var v33 uint32 = (v30 / ((v30) | 1))
                                v30 += (v33 / (((^(v30))) | 1))
                                _ = fi32
                                _ = v33
                            }
                            _ = fi29
                            _ = v30
                            _ = v31
                        }
                    }
                } else {
                    if (-374 != 0) {
                        var v34 float32 = 0.0
                        if ((v4 && (32767 >= 1)) && v4) {
                            v2 -= v2
                            switch 503 {
                            case 360:
                                v35 := GenS3{F0: v2, F1: 2}
                                _ = v35
                            }
                            fmt.Println(v34, v2)
                        } else {
                            v36 := GenS1{F0: genFn0(15)}
                            var v37 string = v1
                            var v38 uint = 814
                            _ = v36
                            _ = v37
                            _ = v38
                        }
                        _ = v34
                    }
                }
                _ = v28
            }
            v39 := []int{127}
            var v40 uint16 = 255
            _ = v39
            _ = v40
        }
    default:
        fmt.Println(v1, v4, v1)
    }
    v41 := GenS1{F0: 0}
    v4 = v4
    v4 = true
    if (-806 > -708) {
        var v42 bool = (false || ((0 > 2147483647) || v4))
        v43 := v3["a"]
        _ = v42
        _ = v43
    }
    var v44 float32 = -2.25
    var v45 int32 = 1
    for fi46 := 0; fi46 < 10; fi46++ {
        v1 = "foo bar"
        var v47 uint64 = 903
        v44 -= 1.5
        _ = fi46
        _ = v47
    }
    switch 32768 {
    case -1:
        var v48 uint8 = 5
        _ = v48
    case 32768:
        if v4 {
            var v49 bool = v4
            _ = len(v3)
            _ = v49
        }
    case 9223372036854775807:
        fmt.Println(v2)
    }
    _ = len(v3)
    v50 := []int{127, 32768}
    _ = len(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v50:", len(v50), v50)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v44
    _ = v45
    _ = v50
}
