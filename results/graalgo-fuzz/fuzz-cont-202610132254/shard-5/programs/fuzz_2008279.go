package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64 }
type GenS2 struct { F0 int8; F1 uint64 }
type GenS3 struct { F0 int8; F1 bool; F2 string }
type GenS4 struct { F0 int; F1 bool; F2 int32 }
type GenS5 struct { F0 int32; F1 uint64; F2 string }
func genFn0(p0 uint8) int {
    switch "\n\t" {
    case "hello":
        if false {
            p0 -= (p0 % (((^(p0))) | 1))
            p0 = 255
            if (false || (857 >= 42)) {
                v5 := []int{32768}
                _ = v5
            } else {
                p0 = p0
                for fi6 := 0; fi6 < 19; fi6++ {
                    v7 := []int{fi6, (fi6 & 9223372036854775807), 127}
                    _ = fi6
                    _ = v7
                }
            }
        }
    case "foo bar":
        for fi8 := 0; fi8 < 9; fi8++ {
            p0 += p0
            var v9 uint32 = 781
            _ = fi8
            _ = v9
        }
    default:
        v10 := []int{}
        _ = v10
    }
    fmt.Println(p0)
    if (32767 >= 309) {
        fmt.Println(p0)
        var v11 int64 = -9223372036854775808
        _ = v11
    } else {
        v12 := []int{4294967295, -9223372036854775808, 2147483647, -1}
        _ = v12
    }
    _ = p0
    return -592
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]uint32)
    var v3 int = -2147483648
    v4 := []int{(^(v3)), (v3 << ((554) & 63)), 32768}
    v13 := GenS1{F0: 2, F1: -9223372036854775808}
    v3 += 127
    if false {
        var v14 string = "hello"
        var v15 int32 = 1
        _ = len(v2)
        _ = v14
        _ = v15
    }
    var v16 uint64 = 2
    var v17 float32 = (3.14159 * (float32((-2.25 - 1e10)) / float32(-2.25)))
    v17 += ((-2.25 - (v17 / 0.1)) / float32(0.1))
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    _ = len(v4)
    var v18 float64 = 1.5
    for fi19 := 0; fi19 < 16; fi19++ {
        var v20 float32 = 3.14159
        v21 := GenS2{F0: 127, F1: v16}
        _ = fi19
        _ = v20
        _ = v21
    }
    v22 := GenS3{F0: 127, F1: false, F2: ("hello" + "")}
    v23 := GenS4{F0: (v3 << ((255) & 63)), F1: true, F2: 2147483647}
    switch v3 {
    case 976:
        v24 := GenS1{F0: 1, F1: 648}
        _ = v24
    case -468:
        v4 = append(v4, v3)
    case 4294967295:
        v3 -= 2147483647
    }
    switch ("foo bar" + "a") {
    case "a":
        switch (v3 % ((v3) | 1)) {
        case -951:
            v25 := GenS5{F0: 2147483647, F1: v16, F2: "foo bar"}
            _ = v25
        case -1:
            if len(v4) > 0 {
                v4[0] = (^(v3))
            }
        case -894:
            fmt.Println(v16, v16)
        default:
            v18 -= 1.5
        }
    default:
        switch (v3 >> ((777) & 63)) {
        case -1:
            if (-72 <= 100) {
                v26 := v1[(v3 - (v3 & (v3 << ((2) & 63))))]
                v27 := GenS3{F0: -10, F1: (1 <= 100), F2: "\n\t"}
                _ = v26
                _ = v27
            } else {
                fmt.Println(v3, v18)
                v3 = 640
            }
        case 686:
            if ((!(!true)) || false) {
                var v28 bool = false
                fmt.Println(v18)
                var v29 uint32 = 4294967295
                _ = v28
                _ = v29
            }
        default:
            switch "a" {
            case "hello":
                v30 := GenS5{F0: 42, F1: (v16 << ((214) & 63)), F2: (("" + ("foo bar" + "\n\t")) + "\n\t")}
                _ = v30
            case "foo bar":
                var v31 uint32 = 42
                _ = v31
            default:
                var v32 uint16 = 2
                _ = v32
            }
        }
    }
    fmt.Println(v3, v17, v3)
    var v33 bool = true
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v13
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v33
    _ = v4
}
