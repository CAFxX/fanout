package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int8; F1 int8; F2 int }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 bool; F1 uint; F2 uint }
type GenS5 struct { F0 uint16; F1 uint16 }
type GenS6 struct { F0 uint8 }
func genFn0() uint64 {
    for fi4 := 0; fi4 < 8; fi4++ {
        if (!true) {
            var v5 int16 = 1
            for fi6 := 0; fi6 < 17; fi6++ {
                v7 := []int{}
                fmt.Println(v5, fi4, fi6)
                if len(v7) > 0 {
                    v7[0] = (fi4 >> ((958) & 63))
                }
                _ = fi6
                _ = v7
            }
            _ = v5
        } else {
            v8 := GenS1{F0: -128}
            _ = v8
        }
        _ = fi4
    }
    switch "\n\t" {
    case "":
        var v9 float64 = (-2.25 / (-2.25 + 3.14159))
        _ = v9
    case "a":
        if false {
            var v10 float32 = 1e10
            var v11 string = ("foo bar" + ("\n\t" + "\n\t"))
            _ = v10
            _ = v11
        } else {
            v12 := []int{}
            v12 = append(v12, 9223372036854775807)
            var v13 uint32 = 65535
            _ = v12
            _ = v13
        }
    case "foo bar":
        v14 := []int{128, 566, 127}
        _ = v14
    }
    var v15 bool = (!(42 <= 42))
    for fi16 := 0; fi16 < 7; fi16++ {
        v17 := GenS1{F0: -125}
        _ = fi16
        _ = v17
    }
    var v18 uint32 = 933
    _ = v15
    _ = v18
    return 18446744073709551615
}

func main() {
    v1 := []uint32{4294967295, 2147483648}
    v2 := []string{"a", "\n\t"}
    v3 := make(map[uint32]int)
    var v19 int16 = 1
    if (0 > 595) {
        for fi20 := 0; fi20 < 2; fi20++ {
            fmt.Println(fi20, fi20)
            for fi21 := 0; fi21 < 4; fi21++ {
                var v22 int = 2147483647
                _ = fi21
                _ = v22
            }
            _ = fi20
        }
    }
    fmt.Println(v19)
    fmt.Println(v19)
    if (468 >= 1) {
        if (true || (!(!true))) {
            switch "a" {
            case "hello":
                fmt.Println(v19)
            case "":
                fmt.Println(v19)
            default:
                v23 := GenS2{F0: 126, F1: 42, F2: -1}
                _ = v23
            }
            if len(v1) > 0 {
                v1[0] = 227
            }
            if false {
                v19 -= (v19 << ((845) & 63))
                var v24 int16 = v19
                _ = v24
            } else {
                v25 := GenS2{F0: -4, F1: 55, F2: -2147483648}
                v26 := GenS3{F0: 1}
                _ = v25
                _ = v26
            }
        }
    } else {
        if ((!false) || false) {
            v19 = (v19 ^ v19)
        } else {
            v27 := GenS4{F0: (((-2.25 * 1.5) == (-2.25 - 0.0)) || false), F1: 4294967295, F2: 65535}
            var v28 uint64 = 673
            _ = v27
            _ = v28
        }
        var v29 uint8 = 255
        var v30 uint32 = 4294967295
        _ = v29
        _ = v30
    }
    fmt.Println(v19)
    var v31 int = 2147483647
    for fi32 := 0; fi32 < 4; fi32++ {
        switch "hello" {
        case "hello":
            for fi33 := 0; fi33 < 4; fi33++ {
                v19 += v19
                _ = fi33
            }
        case "foo bar":
            v34 := GenS5{F0: 42, F1: 42}
            _ = v34
        case "a":
            if false {
                var v35 float32 = -2.25
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                v36 := GenS6{F0: 247}
                _ = v35
                _ = v36
            } else {
                v37 := v3[422]
                _ = v37
            }
        }
        _ = fi32
    }
    if (0 >= 2147483647) {
        v19 = -293
        var v38 uint32 = 4294967295
        if len(v2) > 0 {
            v2[0] = "\n\t"
        }
        _ = v38
    }
    v2 = append(v2, "hello")
    if ((!true) && (7 <= 1)) {
        fmt.Println(v19, v19)
        if false {
            var v39 int = (v31 * (-(v31)))
            var v40 uint8 = 142
            switch ("hello" + "a") {
            case "":
                var v41 float64 = (float64(1.5) / -2.25)
                _ = v41
            case "foo bar":
                v3[2] = v39
            case "a":
                var v42 uint32 = 686
                _ = v42
            default:
                v19 += v19
            }
            _ = v39
            _ = v40
        }
    }
    if (692 <= 1) {
        if len(v2) > 0 {
            v2[0] = ((("hello" + "") + ("a" + "")) + (("a" + "hello") + "\n\t"))
        }
        if (4294967295 <= 167) {
            switch v31 {
            case 0:
                _ = len(v1)
            default:
                v43 := GenS4{F0: (!(v19 != (v19 * (v19 % (((v19 ^ 737)) | 1))))), F1: 1, F2: 65535}
                _ = v43
            }
            v31 = 32767
        }
    }
    v44 := GenS3{F0: 42}
    v45 := GenS2{F0: 127, F1: -15, F2: v31}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v31
    _ = v44
    _ = v45
}
