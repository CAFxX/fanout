package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string }
type GenS2 struct { F0 uint8; F1 uint; F2 int64 }
type GenS3 struct { F0 uint64; F1 uint32 }
func genFn0() string {
    v6 := []int{128, 42, -303, 128}
    var v7 string = "hello"
    var v8 int8 = 0
    v7 = "\n\t"
    v8 = (v8 % ((v8) | 1))
    _ = v6
    _ = v7
    _ = v8
    return ("\n\t" + ("foo bar" + ((v7 + (v7 + v7)) + ("" + "\n\t"))))
}

func main() {
    var v1 bool = false
    var v2 uint32 = 2147483648
    var v3 uint32 = v2
    v4 := make(map[string]uint32)
    v5 := []string{}
    if false {
        v2 = v2
        v9 := GenS1{F0: 42, F1: ("foo bar" + "foo bar")}
        var v10 int = 4294967295
        _ = v10
        _ = v9
    }
    v11 := GenS1{F0: 1, F1: ("a" + "")}
    switch "a" {
    case "a":
        v3 += v2
    case "foo bar":
        v2 += v3
    }
    v1 = v1
    var v12 int8 = 100
    fmt.Println(v1, v2)
    v1 = v1
    for fi13 := 0; fi13 < 2; fi13++ {
        v14 := GenS1{F0: 251, F1: "\n\t"}
        _ = fi13
        _ = v14
    }
    switch -2147483648 {
    case 1:
        var v15 int32 = 1
        _ = v15
    case -1:
        switch "a" {
        case "foo bar":
            v2 = v2
        }
    case -398:
        v16 := GenS2{F0: 2, F1: 4294967295, F2: -1}
        _ = v16
    default:
        var v17 int8 = v12
        _ = v17
    }
    switch "\n\t" {
    case "foo bar":
        v18 := GenS1{F0: 1, F1: "foo bar"}
        _ = v18
    case "hello":
        v12 -= v12
    default:
        v19 := GenS3{F0: 42, F1: v3}
        _ = v19
    }
    fmt.Println(v12)
    v20 := GenS2{F0: 128, F1: 42, F2: -9223372036854775808}
    v21 := GenS3{F0: uint64(v3), F1: (v2 << ((1) & 63))}
    v3 -= (v3 << ((65535) & 63))
    var v22 bool = false
    for fi23 := 0; fi23 < 2; fi23++ {
        if len(v5) > 0 {
            v5[0] = ("" + ("foo bar" + (("hello" + ((((("\n\t" + "") + "") + "") + "") + "hello")) + "hello")))
        }
        _ = fi23
    }
    switch -1 {
    case 127:
        for fi24 := 0; fi24 < 4; fi24++ {
            if ((v1 && (2 == 0)) && (591 <= -364)) {
                v2 -= (v3 << ((900) & 63))
                if (!(!(!(false || v1)))) {
                    fmt.Println(v12)
                } else {
                    if v22 {
                        v22 = (!v22)
                        if ((v12 / ((v12) | 1)) != -128) {
                            var v25 float64 = 0.0
                            _ = v25
                        }
                        var v26 uint32 = (v2 / ((318) | 1))
                        _ = v26
                    }
                    var v27 uint = 42
                    delete(v4, "a")
                    _ = v27
                }
            }
            v5 = append(v5, "\n\t")
            _ = fi24
        }
    case -977:
        for fi28 := 0; fi28 < 8; fi28++ {
            for fi29 := 0; fi29 < 10; fi29++ {
                switch "foo bar" {
                case "":
                    v2 = (v3 % ((v3) | 1))
                }
                if (0.0 <= float64((3.14159 + 0.1))) {
                    v1 = (v1 || false)
                    v30 := GenS2{F0: 42, F1: 2, F2: 0}
                    _ = v30
                } else {
                    if v22 {
                        v12 = (-(v12))
                        if len(v5) > 0 {
                            v5[0] = "hello"
                        }
                        if v22 {
                            if ((3.14159 < 1e10) && false) {
                                v2 += (v3 | (^(v3)))
                                var v31 float32 = -2.25
                                _ = v31
                            } else {
                                var v32 uint16 = 990
                                var v33 int32 = 1
                                if v22 {
                                    var v34 int8 = (v12 / ((int8(v2)) | 1))
                                    var v35 float32 = 0.1
                                    fmt.Println(v35, v35)
                                    _ = len(v5)
                                    _ = v34
                                    _ = v35
                                }
                                _ = v32
                                _ = v33
                            }
                            var v36 int = (^(fi28))
                            _ = v36
                        }
                    }
                    var v37 uint64 = 42
                    _ = v37
                }
                _ = fi29
            }
            _ = fi28
        }
    case 2147483647:
        v38 := GenS2{F0: 229, F1: 65535, F2: -96}
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
