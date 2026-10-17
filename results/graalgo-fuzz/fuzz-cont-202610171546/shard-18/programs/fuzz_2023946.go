package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 string; F1 uint32; F2 bool }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint; F1 int16; F2 int64 }
type GenS5 struct { F0 string }
func genFn0() bool {
    v7 := []int{708, -606}
    var v8 int64 = 9223372036854775807
    v8 -= -137
    _ = v7
    _ = v8
    return (true || (v8 <= v8))
}
func genFn1(p0 bool, p1 int) uint16 {
    var v9 float64 = ((((3.14159 - 3.14159) - float64(1e10)) + -2.25) - 3.14159)
    v10 := []int{p1, 42, p1}
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 414
}

func main() {
    var v1 uint64 = 2
    var v2 int = 0
    var v3 int = v2
    var v4 int64 = -479
    v5 := []uint32{576, 201, 2}
    v6 := make(map[int]int)
    fmt.Println(v1, v1)
    var v11 bool = (!(1000 <= 0))
    var v12 int8 = 127
    var v13 float64 = (-2.25 * ((1e10 * 0.1) * 3.14159))
    switch (^(v2)) {
    case -467:
        _ = len(v5)
    case 127:
        v13 = v13
    case 5:
        if (genFn1(true, (^(v2))) < 276) {
            var v14 uint = 2
            v15 := GenS1{F0: 9223372036854775807, F1: 863}
            _ = v14
            _ = v15
        } else {
            switch "a" {
            case "a":
                for fi16 := 0; fi16 < 2; fi16++ {
                    fmt.Println(v1, v3)
                    var v17 int = (v2 ^ (fi16 + (fi16 >> ((42) & 63))))
                    _ = fi16
                    _ = v17
                }
            }
        }
    }
    var v18 float64 = v13
    v2 -= (v2 + v2)
    fmt.Println(v2, v13, v4)
    v19 := v6[32768]
    fmt.Println(v2, v13)
    var v20 int8 = -11
    var v21 int16 = 1
    var v22 int32 = -2147483648
    if true {
        v11 = (v20 < v20)
        for fi23 := 0; fi23 < 13; fi23++ {
            v24 := GenS1{F0: (-(v3)), F1: 0}
            var v25 int16 = 1000
            _ = fi23
            _ = v24
            _ = v25
        }
    }
    fmt.Println(v4, v1, v21)
    var v26 float32 = 1e10
    v2 -= (v3 / ((v19) | 1))
    for fi27 := 0; fi27 < 16; fi27++ {
        for fi28 := 0; fi28 < 7; fi28++ {
            var v29 uint8 = 2
            var v30 int64 = v4
            switch ("hello" + ("\n\t" + "foo bar")) {
            case "":
                if (!(!(!false))) {
                    v31 := GenS1{F0: v2, F1: 692}
                    if len(v5) >= 2 {
                        v5 = v5[1:len(v5)]
                    }
                    v32 := GenS1{F0: (^(v19)), F1: 2147483647}
                    var v33 int8 = (v12 * 42)
                    _ = v31
                    _ = v32
                    _ = v33
                } else {
                    if (v26 >= 3.14159) {
                        switch "" {
                        case "hello":
                            var v34 float64 = 3.14159
                            if (!v11) {
                                if true {
                                    switch "" {
                                    case "foo bar":
                                        if v11 {
                                            fmt.Println(v3)
                                            var v35 uint16 = 65535
                                            v36 := GenS2{F0: "hello", F1: 2147483648, F2: (int32(v30) != v22)}
                                            _ = v35
                                            _ = v36
                                        }
                                    case "hello":
                                        v21 = (-(v21))
                                    case "\n\t":
                                        var v37 int8 = int8(fi28)
                                        _ = v37
                                    default:
                                        var v38 uint64 = v1
                                        _ = v38
                                    }
                                    if v11 {
                                        v39 := GenS3{F0: 840}
                                        fmt.Println(v20, v18, v18)
                                        _ = v39
                                    }
                                } else {
                                    fmt.Println(v21, v18, v11)
                                    var v40 float64 = v18
                                    _ = v40
                                }
                                var v41 string = "a"
                                v3 = v3
                                _ = v41
                            }
                            var v42 string = "foo bar"
                            _ = v34
                            _ = v42
                        case "foo bar":
                            v43 := GenS4{F0: 65535, F1: v21, F2: v30}
                            _ = v43
                        case "a":
                            v3 = fi27
                        default:
                            v2 -= -839
                        }
                        v44 := GenS5{F0: "hello"}
                        _ = len(v6)
                        _ = v44
                    } else {
                        fmt.Println(v11)
                    }
                }
            case "a":
                switch v2 {
                case -788:
                    v30 = (v30 | v4)
                case -9223372036854775808:
                    var v45 float32 = (float32(v13) + 3.14159)
                    _ = v45
                case 32767:
                    var v46 uint = 18446744073709551615
                    _ = v46
                default:
                    var v47 uint = 69
                    _ = v47
                }
            default:
                var v48 uint8 = v29
                _ = v48
            }
            if true {
                var v49 float64 = float64(((1e10 - v26) - float32(v18)))
                v11 = (true || false)
                v50 := GenS4{F0: 1, F1: v21, F2: v30}
                _ = v49
                _ = v50
            }
            var v51 bool = ((!v11) || (!(true && false)))
            v52 := GenS4{F0: 42, F1: (v21 / (((v21 * (v21 % (((v21 * v21)) | 1)))) | 1)), F2: int64(v3)}
            v29 += v29
            if (24 != v29) {
                fmt.Println(v11, v51)
                v21 = v21
                v53 := GenS5{F0: ""}
                _ = v53
            }
            _ = fi28
            _ = v29
            _ = v30
            _ = v51
            _ = v52
        }
        var v54 string = "foo bar"
        if (v54 < (v54 + "foo bar")) {
            v55 := v6[v2]
            v12 = v12
            v56 := GenS1{F0: v3, F1: 2147483648}
            _ = v55
            _ = v56
        } else {
            if (v12 < v20) {
                for fi57 := 0; fi57 < 19; fi57++ {
                    v21 = -32768
                    _ = fi57
                }
            }
        }
        _ = fi27
        _ = v54
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
