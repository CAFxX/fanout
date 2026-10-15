package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int16; F1 uint16 }
type GenS7 struct { F0 int; F1 int32 }
type GenS8 struct { F0 uint16; F1 uint; F2 uint64 }
func genFn0() int {
    var v7 int = 523
    switch "\n\t" {
    case "hello":
        switch "a" {
        case "a":
            fmt.Println(v7)
        }
    case "\n\t":
        v8 := GenS1{F0: v7}
        _ = v8
    case "a":
        if (42 < 26) {
            var v9 int = 1
            v9 = -585
            v10 := GenS1{F0: (v7 << ((42) & 63))}
            _ = v10
            _ = v9
        }
    }
    _ = v7
    return (v7 % ((1) | 1))
}
func genFn1() int32 {
    var v11 int = 32767
    var v12 int32 = 1
    v13 := []int{}
    fmt.Println(v12)
    v14 := GenS2{F0: 42}
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    return v12
}

func main() {
    var v1 string = "a"
    var v2 int = 2147483647
    v3 := []string{(("" + v1) + "\n\t"), (((v1 + v1) + v1) + (v1 + v1))}
    v4 := make(map[string]string)
    v5 := make(map[int]bool)
    var v6 uint32 = 289
    if len(v3) > 0 {
        v3[0] = "\n\t"
    }
    var v15 float32 = ((0.0 - 1e10) * 1e10)
    switch v2 {
    case 745:
        if (v15 >= (v15 / 1.0)) {
            v16 := GenS3{F0: (v6 ^ (v6 ^ (^(v6))))}
            v17 := GenS4{F0: 42}
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = v16
            _ = v17
        }
    case 32767:
        for fi18 := 0; fi18 < 16; fi18++ {
            v5[fi18] = true
            _ = fi18
        }
    case 127:
        if true {
            if true {
                delete(v4, "\n\t")
            } else {
                var v19 bool = (!(true && false))
                switch v2 {
                case 190:
                    var v20 int64 = 0
                    _ = v20
                case 2147483647:
                    v21 := GenS5{F0: -90}
                    _ = v21
                }
                _ = v19
            }
            v22 := GenS3{F0: v6}
            v23 := GenS4{F0: 42}
            _ = v22
            _ = v23
        }
    default:
        v15 += (((1.5 - (v15 + v15)) + 1.5) * float32(0.1))
    }
    fmt.Println(v6, v1, v2)
    var v24 float32 = 1e10
    var v25 string = "foo bar"
    var v26 float32 = (v15 / v24)
    if (42 < -128) {
        v27 := GenS6{F0: 0, F1: 255}
        for fi28 := 0; fi28 < 14; fi28++ {
            var v29 uint8 = 42
            v29 -= v29
            for fi30 := 0; fi30 < 19; fi30++ {
                var v31 uint64 = 2
                fmt.Println(v6, v2, v31)
                if ((1e10 > ((((((3.14159 - (((0.0 * -2.25) / float64(v26)) / 0.1)) * -2.25) / 1.5) + (0.1 + 0.1)) + 1e10) - 0.1)) && (!(0 >= 124))) {
                    var v32 int16 = 1000
                    fmt.Println(v24)
                    _ = v32
                } else {
                    var v33 uint32 = (v6 >> ((42) & 63))
                    if false {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                    _ = v33
                }
                _ = fi30
                _ = v31
            }
            _ = fi28
            _ = v29
        }
        _ = v27
    }
    v34 := GenS1{F0: v2}
    v2 += 32767
    v15 -= (v15 / v26)
    var v35 float32 = ((-2.25 / (0.1 * v24)) + (v26 - 3.14159))
    var v36 float64 = 0.1
    v1 = "\n\t"
    var v37 uint = 183
    for fi38 := 0; fi38 < 2; fi38++ {
        v39 := GenS7{F0: -1, F1: genFn1()}
        v35 = (v24 - (0.0 * (v15 * (float32(1.5) + v24))))
        if true {
            v35 -= v26
            for fi40 := 0; fi40 < 13; fi40++ {
                v4[("foo bar" + (v25 + (v25 + v25)))] = ""
                v3 = append(v3, "foo bar")
                if true {
                    var v41 int32 = -396
                    var v42 int64 = 121
                    v26 = 0.0
                    v41 -= 0
                    switch "hello" {
                    case "foo bar":
                        if true {
                            var v43 float64 = 0.1
                            if false {
                                v25 = ""
                                v4[v25] = v25
                            }
                            var v44 int16 = -1
                            _ = v43
                            _ = v44
                        }
                    case "hello":
                        v41 -= (^(v41))
                    default:
                        v42 = v42
                    }
                    _ = v41
                    _ = v42
                } else {
                    switch genFn0() {
                    case -692:
                        v45 := v4[""]
                        _ = v45
                    case 0:
                        var v46 int64 = -459
                        _ = v46
                    case -2147483648:
                        v26 = v26
                    default:
                        if false {
                            var v47 uint64 = 371
                            _ = v47
                        } else {
                            v48 := GenS7{F0: fi38, F1: 42}
                            var v49 uint8 = 0
                            var v50 float32 = (((v24 - 1e10) + 1.5) / (v26 / (v26 / v35)))
                            _ = v48
                            _ = v49
                            _ = v50
                        }
                    }
                }
                _ = fi40
            }
            var v51 int16 = 401
            _ = v51
        } else {
            var v52 uint = v37
            _ = v52
        }
        _ = fi38
        _ = v39
    }
    switch ("\n\t" + (v1 + (v25 + "a"))) {
    case "hello":
        v53 := GenS5{F0: v2}
        _ = v53
    case "a":
        v54 := GenS8{F0: 849, F1: v37, F2: 2}
        _ = v54
    }
    var v55 uint32 = 2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v55:", v55)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v55
    _ = v6
}
