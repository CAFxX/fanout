package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8; F1 int }
type GenS3 struct { F0 uint64; F1 string; F2 int64 }
type GenS4 struct { F0 uint32 }
func genFn0() uint64 {
    switch 9223372036854775807 {
    case 1:
        var v6 int32 = -2147483648
        _ = v6
    }
    v7 := []int{}
    _ = v7
    return 18446744073709551615
}
func genFn1(p0 uint) uint64 {
    if (false && ((1.5 + (1e10 / 3.14159)) != (0.0 + 0.0))) {
        var v8 float32 = 1.5
        var v9 uint32 = 0
        _ = v8
        _ = v9
    } else {
        p0 += p0
        var v10 int64 = 42
        _ = v10
    }
    switch "\n\t" {
    case "hello":
        p0 += 567
    case "a":
        if ((((true || ((!false) && true)) || (!true)) && (32767 == 42)) || false) {
            var v11 uint64 = 2
            for fi12 := 0; fi12 < 17; fi12++ {
                v13 := GenS1{F0: 1}
                var v14 int16 = -1
                _ = fi12
                _ = v13
                _ = v14
            }
            _ = v11
        } else {
            if (2147483647 <= 42) {
                p0 -= (p0 - 65535)
            }
        }
    }
    _ = p0
    return 192
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int32 = -2147483648
    var v3 bool = (true || ("foo bar" > "a"))
    var v4 float64 = 0.1
    var v5 bool = v3
    v15 := GenS2{F0: 22, F1: 1}
    if (9223372036854775807 > 1) {
        fmt.Println(v3)
        var v16 uint64 = 4294967295
        v17 := GenS2{F0: 128, F1: -434}
        _ = v16
        _ = v17
    } else {
        var v18 uint = 18446744073709551615
        _ = v18
    }
    _ = len(v1)
    switch ("foo bar" + (("\n\t" + ("\n\t" + "")) + "a")) {
    case "":
        for fi19 := 0; fi19 < 15; fi19++ {
            fmt.Println(v5, v4)
            _ = fi19
        }
    case "hello":
        if (193 > 32767) {
            fmt.Println(v3, v2)
        } else {
            var v20 bool = (!(0 == -1))
            _ = v20
        }
    case "foo bar":
        v5 = (1 == 69)
    }
    if false {
        v3 = (0 >= -589)
        var v21 float32 = 1.5
        _ = v21
    } else {
        var v22 uint16 = 674
        _ = v22
    }
    v23 := GenS1{F0: -1}
    v24 := GenS2{F0: 127, F1: -363}
    v5 = (255 > 274)
    for fi25 := 0; fi25 < 4; fi25++ {
        v26 := GenS3{F0: genFn0(), F1: ("" + "hello"), F2: -177}
        for fi27 := 0; fi27 < 11; fi27++ {
            switch 128 {
            case 1:
                var v28 string = "\n\t"
                _ = v28
            case 32768:
                var v29 int32 = (v2 * v2)
                _ = v29
            default:
                var v30 uint8 = 128
                _ = v30
            }
            fmt.Println(v4, v4, v2)
            _ = fi27
        }
        _ = fi25
        _ = v26
    }
    v31 := GenS3{F0: 419, F1: (("\n\t" + "") + ""), F2: 0}
    if false {
        var v32 int32 = v2
        v32 = v2
        _ = v32
    } else {
        v5 = false
    }
    v33 := v1[4294967295]
    fmt.Println(v4, v33)
    v34 := []int{32767, 127, 953, -436}
    v35 := GenS3{F0: 42, F1: "", F2: 446}
    v36 := GenS4{F0: (v33 - (v33 | (^(v33))))}
    v2 -= -43
    switch -1 {
    case -386:
        var v37 int64 = -810
        _ = v37
    case 32767:
        v33 = (v33 >> ((42) & 63))
    case 127:
        _ = len(v34)
    default:
        _ = len(v34)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v31
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
