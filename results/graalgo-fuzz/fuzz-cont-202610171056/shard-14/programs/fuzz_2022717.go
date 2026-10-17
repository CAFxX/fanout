package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int; F1 uint32 }
func genFn0(p0 int8, p1 int32, p2 uint8) int {
    switch 9223372036854775807 {
    case -9223372036854775808:
        for fi6 := 0; fi6 < 10; fi6++ {
            for fi7 := 0; fi7 < 9; fi7++ {
                if true {
                    v8 := []int{(-(fi6)), 32767, (fi6 / ((46) | 1))}
                    _ = len(v8)
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                    _ = v8
                }
                var v9 uint64 = 1
                var v10 float64 = ((-2.25 - 1e10) / -2.25)
                _ = fi7
                _ = v10
                _ = v9
            }
            _ = fi6
        }
    case 127:
        var v11 int64 = 185
        _ = v11
    }
    fmt.Println(p2)
    v12 := []int{}
    var v13 uint8 = (p2 | p2)
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v13
    return 2147483647
}
func genFn1() string {
    var v14 uint16 = 255
    var v15 uint32 = 2147483648
    _ = v14
    _ = v15
    return ("foo bar" + (("hello" + "") + "\n\t"))
}

func main() {
    var v1 float64 = 0.0
    v2 := make(map[int]uint32)
    var v3 int64 = 42
    var v4 uint64 = 42
    var v5 int32 = 2147483647
    var v16 uint8 = 113
    v17 := GenS1{F0: ((("a" + "hello") + "a") + ("hello" + (("a" + ("a" + "a")) + "hello"))), F1: v16}
    v3 = v3
    if (353 <= 1) {
        if (!(true || (!true))) {
            var v18 int64 = v3
            _ = v18
        }
        v16 = v16
    } else {
        var v19 int16 = -1
        for fi20 := 0; fi20 < 20; fi20++ {
            var v21 float32 = 1.5
            _ = fi20
            _ = v21
        }
        switch (("\n\t" + "\n\t") + ("hello" + "foo bar")) {
        case "":
            var v22 uint8 = 45
            _ = v22
        case "\n\t":
            v16 = 42
        case "foo bar":
            v23 := GenS1{F0: "a", F1: v16}
            _ = v23
        }
        _ = v19
    }
    switch -561 {
    case 883:
        fmt.Println(v16, v16, v16)
    }
    switch "hello" {
    case "a":
        v24 := GenS2{F0: -764}
        _ = v24
    }
    fmt.Println(v3, v3)
    for fi25 := 0; fi25 < 19; fi25++ {
        v26 := GenS3{F0: (v5 * (v5 >> ((0) & 63)))}
        v5 -= 42
        _ = fi25
        _ = v26
    }
    if (!false) {
        v2[genFn0(42, (v5 & (v5 + v5)), v16)] = 2147483648
        switch 2147483647 {
        case -877:
            var v27 int8 = -128
            _ = v27
        default:
            for fi28 := 0; fi28 < 18; fi28++ {
                for fi29 := 0; fi29 < 19; fi29++ {
                    delete(v2, (fi29 ^ (fi28 & (fi28 - (fi28 ^ (fi28 << ((42) & 63)))))))
                    _ = fi29
                }
                v30 := []int{42, fi28}
                fmt.Println(v5, v4, v1)
                _ = fi28
                _ = v30
            }
        }
    } else {
        var v31 int8 = -54
        _ = v31
    }
    v32 := GenS4{F0: genFn0(1, (-(v5)), (v16 & 1)), F1: 42}
    var v33 int = -9223372036854775808
    if false {
        var v34 uint64 = 18446744073709551615
        var v35 uint32 = 1
        _ = v34
        _ = v35
    } else {
        switch (-(v33)) {
        case 9223372036854775807:
            var v36 float32 = -2.25
            _ = v36
        case -9223372036854775808:
            var v37 string = "\n\t"
            _ = v37
        }
        fmt.Println(v1, v33, v33)
    }
    v38 := []int{(v33 / ((v33) | 1)), 128, (v33 - v33), v33}
    if false {
        v39 := GenS1{F0: ("hello" + ""), F1: 128}
        for fi40 := 0; fi40 < 4; fi40++ {
            var v41 int8 = -119
            var v42 int16 = -348
            if false {
                v43 := GenS4{F0: -1, F1: uint32(v16)}
                v5 -= v5
                _ = v43
            } else {
                fmt.Println(v16, v4, v5)
            }
            _ = fi40
            _ = v41
            _ = v42
        }
        fmt.Println(v3, v5, v3)
        _ = v39
    }
    v3 += (v3 % ((0) | 1))
    v44 := GenS3{F0: (v5 / ((v5) | 1))}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v38
    _ = v4
    _ = v44
    _ = v5
}
