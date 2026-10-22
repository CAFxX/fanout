package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 int }

func main() {
    v1 := make(map[string]int)
    v2 := make(map[int]uint32)
    var v3 int16 = 1
    fmt.Println(v3)
    if false {
        for fi4 := 0; fi4 < 20; fi4++ {
            for fi5 := 0; fi5 < 18; fi5++ {
                v3 = (v3 >> ((490) & 63))
                var v6 string = "\n\t"
                if ((!true) || (!false)) {
                    fmt.Println(fi5, v3)
                } else {
                    fmt.Println(fi5, v6, fi5)
                    var v7 int16 = v3
                    var v8 uint16 = 32767
                    _ = v7
                    _ = v8
                }
                _ = fi5
                _ = v6
            }
            fmt.Println(v3, fi4)
            _ = fi4
        }
        for fi9 := 0; fi9 < 2; fi9++ {
            var v10 float32 = float32(-2.25)
            _ = fi9
            _ = v10
        }
    }
    switch -921 {
    case 32768:
        v3 -= (v3 - 0)
    default:
        var v11 uint32 = 42
        _ = v11
    }
    if (2 >= 1) {
        var v12 float32 = 1e10
        _ = v12
    } else {
        var v13 bool = (731 >= 4294967295)
        for fi14 := 0; fi14 < 10; fi14++ {
            var v15 uint32 = 65535
            v16 := []int{fi14, (fi14 + fi14), 32768}
            _ = fi14
            _ = v15
            _ = v16
        }
        for fi17 := 0; fi17 < 1; fi17++ {
            v18 := GenS1{F0: (!(65535 >= 255)), F1: -9223372036854775808}
            var v19 uint64 = 167
            switch fi17 {
            case 2147483647:
                v19 += v19
            case 32768:
                v20 := GenS2{F0: fi17}
                _ = v20
            case -9223372036854775808:
                v21 := GenS2{F0: fi17}
                _ = v21
            }
            _ = fi17
            _ = v18
            _ = v19
        }
        _ = v13
    }
    v3 = (v3 << ((42) & 63))
    _ = len(v1)
    switch 127 {
    case -2147483648:
        var v22 int64 = 9223372036854775807
        _ = v22
    case 9223372036854775807:
        v23 := GenS1{F0: ((9223372036854775807 > 911) || false), F1: 9223372036854775807}
        _ = v23
    }
    v2[4294967295] = 270
    var v24 uint = 18446744073709551615
    v25 := GenS2{F0: 5}
    if true {
        for fi26 := 0; fi26 < 3; fi26++ {
            v24 = 4294967295
            var v27 uint = v24
            v3 -= (^(v3))
            _ = fi26
            _ = v27
        }
    }
    if (255 <= 1) {
        switch -9223372036854775808 {
        case 42:
            var v28 int8 = -114
            _ = v28
        case 0:
            var v29 int8 = 81
            _ = v29
        case -595:
            v30 := []int{9223372036854775807, -199, 32768}
            _ = v30
        default:
            v31 := GenS1{F0: true, F1: 9223372036854775807}
            _ = v31
        }
        if ((128 == 4294967295) && true) {
            v3 -= (v3 | (v3 * v3))
            var v32 uint = v24
            _ = v32
        }
    }
    v33 := []int{477, 127, -2147483648, 32768}
    var v34 uint32 = 42
    v24 = (v24 + v24)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v33:", len(v33), v33)
    fmt.Println("v34:", v34)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v33
    _ = v34
}
