package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0() int {
    if (1 < 0) {
        var v4 uint64 = 18446744073709551615
        if (!(false || (9223372036854775807 != 1))) {
            v4 += (v4 + v4)
            var v5 uint32 = 926
            _ = v5
        }
        _ = v4
    } else {
        var v6 uint = 18446744073709551615
        _ = v6
    }
    switch ("foo bar" + ("" + (("a" + "") + "\n\t"))) {
    case "hello":
        var v7 int16 = 1
        _ = v7
    case "foo bar":
        var v8 int8 = -128
        _ = v8
    case "":
        if true {
            var v9 uint = 42
            _ = v9
        } else {
            var v10 uint = 255
            v11 := []int{}
            switch 128 {
            case -719:
                if len(v11) >= 2 {
                    v11 = v11[1:len(v11)]
                }
            case -854:
                if false {
                    fmt.Println(v10)
                    if false {
                        var v12 float32 = (float32(3.14159) * (0.1 * 0.1))
                        _ = v12
                    }
                    for fi13 := 0; fi13 < 13; fi13++ {
                        for fi14 := 0; fi14 < 7; fi14++ {
                            v15 := GenS1{F0: -1}
                            var v16 uint = v10
                            _ = fi14
                            _ = v15
                            _ = v16
                        }
                        var v17 uint32 = 42
                        fmt.Println(v17, v17)
                        _ = fi13
                        _ = v17
                    }
                }
            case -9223372036854775808:
                switch 32768 {
                case -442:
                    if len(v11) >= 2 {
                        v11 = v11[1:len(v11)]
                    }
                case 32768:
                    if false {
                        v10 = (v10 + (^(v10)))
                    }
                default:
                    if false {
                        var v18 int64 = -1
                        _ = v18
                    } else {
                        var v19 uint8 = 42
                        v10 = v10
                        var v20 int64 = 9223372036854775807
                        _ = v19
                        _ = v20
                    }
                }
            }
            _ = v10
            _ = v11
        }
    default:
        var v21 int8 = 1
        _ = v21
    }
    return 1
}
func genFn1() int8 {
    var v22 int8 = -1
    for fi23 := 0; fi23 < 1; fi23++ {
        switch genFn0() {
        case 32767:
            var v24 string = ("hello" + ("" + "\n\t"))
            _ = v24
        case 127:
            for fi25 := 0; fi25 < 4; fi25++ {
                v22 -= (-(v22))
                v22 += v22
                _ = fi25
            }
        default:
            var v26 int = (fi23 + (fi23 * 32768))
            _ = v26
        }
        v22 -= -32
        v27 := []int{fi23}
        _ = fi23
        _ = v27
    }
    _ = v22
    return v22
}

func main() {
    v1 := []uint32{84}
    var v2 uint = 282
    var v3 float32 = 1e10
    v2 += 1
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v28 := GenS1{F0: 42}
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    for fi29 := 0; fi29 < 6; fi29++ {
        v1 = append(v1, 615)
        fmt.Println(v2, v2, fi29)
        _ = fi29
    }
    var v30 int8 = -55
    var v31 uint64 = 0
    var v32 int = genFn0()
    switch ("" + "a") {
    case "\n\t":
        v33 := GenS1{F0: 42}
        _ = v33
    case "":
        fmt.Println(v30, v30, v30)
    case "foo bar":
        v32 = (v32 % ((genFn0()) | 1))
    }
    var v34 bool = false
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v34
}
