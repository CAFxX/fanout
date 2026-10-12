package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 uint8 }
type GenS2 struct { F0 uint32; F1 int8; F2 int16 }
func genFn0(p0 uint32, p1 uint) int16 {
    v7 := GenS1{F0: 707, F1: 4294967295, F2: 106}
    fmt.Println(p1, p0)
    for fi8 := 0; fi8 < 6; fi8++ {
        fmt.Println(p0, fi8)
        _ = fi8
    }
    if false {
        if true {
            v9 := []int{}
            fmt.Println(p1)
            if (("foo bar" + "a") <= ("foo bar" + "hello")) {
                var v10 uint = (p1 ^ 4294967295)
                _ = v10
            } else {
                fmt.Println(p1)
            }
            _ = v9
        }
    } else {
        var v11 float32 = (1e10 * ((0.1 + 3.14159) + (1.5 / 3.14159)))
        v12 := []int{335}
        _ = v11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = v7
    return 1000
}
func genFn1() uint64 {
    if true {
        v13 := []int{9223372036854775807, 32768, 0, 32767}
        _ = v13
    } else {
        var v14 float64 = (0.0 - float64(float32((((1e10 + 3.14159) + -2.25) / 1e10))))
        var v15 int16 = -810
        if false {
            if false {
                var v16 uint16 = 133
                _ = v16
            }
            var v17 bool = true
            var v18 uint = 42
            _ = v17
            _ = v18
        } else {
            fmt.Println(v14, v15)
            v19 := []int{683, 128, 42}
            _ = v19
        }
        _ = v14
        _ = v15
    }
    switch "\n\t" {
    case "":
        v20 := GenS1{F0: 2147483647, F1: 552, F2: 128}
        _ = v20
    case "hello":
        v21 := []int{-1}
        _ = v21
    default:
        var v22 uint64 = 4294967295
        _ = v22
    }
    switch 743 {
    case 32768:
        var v23 uint16 = 2
        _ = v23
    case 128:
        v24 := GenS2{F0: 2147483647, F1: -91, F2: 1}
        _ = v24
    case 32767:
        switch (("hello" + (("\n\t" + ("foo bar" + "hello")) + ("\n\t" + "\n\t"))) + "") {
        case "a":
            if (!true) {
                if (!(false && (3.14159 != 3.14159))) {
                    var v25 int8 = 127
                    _ = v25
                } else {
                    var v26 uint64 = 18446744073709551615
                    _ = v26
                }
                switch (("hello" + "hello") + "a") {
                case "":
                    v27 := []int{4294967295, -2147483648}
                    _ = v27
                }
            } else {
                v28 := GenS1{F0: 0, F1: 4294967295, F2: 255}
                _ = v28
            }
        case "foo bar":
            for fi29 := 0; fi29 < 2; fi29++ {
                fmt.Println(fi29)
                _ = fi29
            }
        default:
            var v30 uint16 = 2
            _ = v30
        }
    }
    return 4294967295
}

func main() {
    var v1 int16 = 225
    v2 := []int{}
    var v3 int64 = -1
    var v4 string = "\n\t"
    v5 := []int{}
    v6 := []uint32{519, 25}
    var v31 int16 = -1
    if false {
        var v32 int32 = 2147483647
        _ = v32
    }
    var v33 float32 = 0.1
    switch -9223372036854775808 {
    case 1:
        if len(v5) > 0 {
            v5[0] = 4294967295
        }
    }
    v34 := GenS1{F0: 0, F1: genFn1(), F2: 1}
    for fi35 := 0; fi35 < 5; fi35++ {
        _ = len(v6)
        _ = fi35
    }
    for fi36 := 0; fi36 < 3; fi36++ {
        if true {
            if (v4 != (("" + v4) + "")) {
                v2 = append(v2, fi36)
            } else {
                v3 = 0
                v37 := GenS1{F0: 0, F1: genFn1(), F2: 42}
                _ = v37
            }
            var v38 uint16 = 32767
            _ = v38
        }
        v31 -= (v1 << ((2) & 63))
        v39 := GenS1{F0: -2147483648, F1: 18446744073709551615, F2: 127}
        _ = fi36
        _ = v39
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    var v40 float64 = 0.0
    fmt.Println(v33, v33, v31)
    switch "foo bar" {
    case "\n\t":
        var v41 int64 = v3
        _ = v41
    case "hello":
        var v42 int64 = v3
        _ = v42
    case "a":
        v4 = v4
    }
    _ = len(v6)
    for fi43 := 0; fi43 < 5; fi43++ {
        fmt.Println(v40)
        v3 = (v3 + -1)
        if false {
            v1 -= genFn0(658, 65535)
            var v44 uint = 42
            _ = v44
        }
        _ = fi43
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v33
    _ = v34
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
