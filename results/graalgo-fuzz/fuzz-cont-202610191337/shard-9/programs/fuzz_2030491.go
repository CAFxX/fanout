package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16; F1 int16 }
type GenS3 struct { F0 uint8; F1 bool }
func genFn0() uint64 {
    v6 := []int{536}
    var v7 int16 = -32768
    _ = v6
    _ = v7
    return 1
}
func genFn1() int64 {
    for fi8 := 0; fi8 < 14; fi8++ {
        var v9 int = fi8
        v10 := []int{(-(fi8)), (-(v9)), -593}
        _ = fi8
        _ = v10
        _ = v9
    }
    var v11 uint = 2
    v12 := []int{4294967295}
    fmt.Println(v11)
    _ = v11
    _ = v12
    return 9223372036854775807
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int16 = -662
    var v3 int64 = 9223372036854775807
    var v4 int32 = 1
    var v5 int16 = (v2 >> ((4294967295) & 63))
    if false {
        var v13 int16 = v2
        _ = v13
    } else {
        var v14 string = ""
        if ((v4 - v4) == (v4 % ((v4) | 1))) {
            v15 := []int{127, 4294967295, 42, 367}
            var v16 uint16 = 65535
            _ = v15
            _ = v16
        }
        _ = v14
    }
    v17 := GenS1{F0: 18446744073709551615}
    fmt.Println(v5)
    var v18 uint16 = 600
    switch ("\n\t" + ("" + "foo bar")) {
    case "a":
        v5 = -986
    case "hello":
        if (("hello" != "foo bar") && (true || false)) {
            var v19 int32 = v4
            _ = v19
        } else {
            if (2147483647 == (v4 % (((-(v4))) | 1))) {
                if true {
                    v5 = v2
                    if (true && (100 >= -128)) {
                        if true {
                            v5 = v2
                        } else {
                            v20 := GenS1{F0: 2}
                            var v21 uint64 = 18446744073709551615
                            _ = v20
                            _ = v21
                        }
                    }
                    v2 = v2
                }
            }
        }
    default:
        var v22 int32 = (v4 & 418)
        _ = v22
    }
    v23 := v1[32767]
    if false {
        var v24 bool = true
        v25 := GenS1{F0: 4294967295}
        for fi26 := 0; fi26 < 9; fi26++ {
            fmt.Println(v2, v18, v24)
            for fi27 := 0; fi27 < 13; fi27++ {
                switch "a" {
                case "a":
                    var v28 uint = 18446744073709551615
                    _ = v28
                }
                _ = fi27
            }
            var v29 float32 = ((3.14159 / 1.0) + 3.14159)
            _ = fi26
            _ = v29
        }
        _ = v24
        _ = v25
    } else {
        _ = len(v1)
        for fi30 := 0; fi30 < 17; fi30++ {
            fmt.Println(v5)
            _ = fi30
        }
        var v31 uint8 = 128
        _ = v31
    }
    var v32 bool = false
    v5 = (v2 + v5)
    switch 42 {
    case 4294967295:
        for fi33 := 0; fi33 < 6; fi33++ {
            for fi34 := 0; fi34 < 15; fi34++ {
                v1[(fi34 % (((fi34 * fi34)) | 1))] = (v23 * v23)
                if (65535 < 550) {
                    v18 = v18
                    v35 := []int{(fi34 + int(v18))}
                    _ = v35
                }
                if (255 > 255) {
                    var v36 int64 = v3
                    _ = v36
                } else {
                    var v37 uint8 = 127
                    _ = v37
                }
                var v38 int16 = v5
                var v39 bool = (0.0 > 1e10)
                v40 := GenS2{F0: (v5 - v5), F1: 1}
                _ = fi34
                _ = v38
                _ = v39
                _ = v40
            }
            fmt.Println(v3, v2, v5)
            _ = fi33
        }
    case 127:
        var v41 uint32 = (^(v23))
        _ = v41
    default:
        v42 := GenS3{F0: 71, F1: true}
        _ = v42
    }
    var v43 uint8 = 22
    fmt.Println(v3, v5)
    switch ((((("foo bar" + "hello") + ("foo bar" + "")) + "") + "a") + "") {
    case "\n\t":
        v44 := GenS2{F0: v2, F1: v5}
        _ = v44
    case "foo bar":
        fmt.Println(v4, v43, v4)
    case "a":
        var v45 uint32 = (^(v23))
        _ = v45
    }
    if (true || (2147483647 < 752)) {
        v46 := v1[476]
        var v47 int8 = 100
        if (18446744073709551615 < 4294967295) {
            var v48 float32 = (-2.25 * 0.0)
            _ = v48
        } else {
            v49 := []int{419, 9223372036854775807}
            _ = v49
        }
        _ = v46
        _ = v47
    } else {
        if false {
            var v50 float32 = 1e10
            v32 = v32
            _ = v50
        }
        var v51 uint = 42
        v52 := GenS2{F0: (v5 ^ (v5 - (v2 + v5))), F1: (v5 >> ((1) & 63))}
        _ = v51
        _ = v52
    }
    var v53 float32 = 1e10
    for fi54 := 0; fi54 < 12; fi54++ {
        v32 = false
        for fi55 := 0; fi55 < 12; fi55++ {
            var v56 uint64 = 733
            var v57 uint = 4294967295
            v56 = genFn0()
            _ = fi55
            _ = v56
            _ = v57
        }
        _ = fi54
    }
    v23 = 782
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v23
    _ = v3
    _ = v32
    _ = v4
    _ = v43
    _ = v5
    _ = v53
}
