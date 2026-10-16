package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32; F1 int32; F2 int }
type GenS4 struct { F0 uint; F1 string; F2 int32 }
func genFn0(p0 uint8, p1 uint) int64 {
    var v4 float64 = ((3.14159 + (1.5 * 3.14159)) * 0.0)
    v5 := []int{540, 32768}
    switch int(p1) {
    case 127:
        var v6 uint32 = 42
        _ = v6
    case 2147483647:
        var v7 int8 = 100
        _ = v7
    }
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 9223372036854775807
}

func main() {
    v1 := []uint32{4294967295, 42}
    v2 := make(map[int]uint32)
    var v3 uint16 = 0
    v3 = (v3 >> ((65535) & 63))
    if (true && (-52 == -10)) {
        var v8 uint32 = 0
        switch -9223372036854775808 {
        case 32768:
            var v9 uint8 = 127
            _ = v9
        case 2147483647:
            var v10 uint32 = 444
            _ = v10
        }
        _ = v8
    }
    if true {
        v11 := GenS1{F0: 168, F1: 1}
        for fi12 := 0; fi12 < 9; fi12++ {
            if (!true) {
                for fi13 := 0; fi13 < 7; fi13++ {
                    v3 -= (v3 - (v3 - v3))
                    v14 := GenS2{F0: 42}
                    _ = fi13
                    _ = v14
                }
                v2[fi12] = 42
                var v15 uint8 = 128
                _ = v15
            }
            _ = fi12
        }
        _ = v11
    }
    if (877 <= 1) {
        fmt.Println(v3)
        v3 -= (v3 & (v3 + (v3 + v3)))
    }
    if (1 == 1) {
        switch 9223372036854775807 {
        case 9223372036854775807:
            for fi16 := 0; fi16 < 17; fi16++ {
                v17 := GenS3{F0: -2147483648, F1: 1, F2: fi16}
                var v18 uint = 18446744073709551615
                _ = fi16
                _ = v17
                _ = v18
            }
        }
        var v19 string = ""
        _ = len(v1)
        _ = v19
    }
    v3 -= 0
    v3 = v3
    for fi20 := 0; fi20 < 11; fi20++ {
        _ = len(v1)
        var v21 uint64 = 1
        _ = fi20
        _ = v21
    }
    var v22 uint32 = 725
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v23 bool = (!false)
    v3 += (v3 & v3)
    if true {
        v22 = 1
        fmt.Println(v3, v23, v3)
        delete(v2, -9223372036854775808)
    } else {
        var v24 uint8 = 255
        if (v23 && (v23 && (!(v3 > (^(v3)))))) {
            for fi25 := 0; fi25 < 12; fi25++ {
                var v26 string = ("a" + "")
                fmt.Println(fi25, v22)
                _ = fi25
                _ = v26
            }
            var v27 uint32 = v22
            for fi28 := 0; fi28 < 5; fi28++ {
                v23 = (((^(v24)) != (v24 | v24)) && v23)
                _ = fi28
            }
            _ = v27
        } else {
            var v29 uint16 = (^(v3))
            fmt.Println(v24, v3)
            for fi30 := 0; fi30 < 5; fi30++ {
                fmt.Println(v24)
                fmt.Println(v24, v29, v3)
                _ = fi30
            }
            _ = v29
        }
        _ = v24
    }
    var v31 uint16 = (^(v3))
    fmt.Println(v22, v31, v3)
    v32 := GenS4{F0: 18446744073709551615, F1: "hello", F2: 42}
    var v33 uint16 = (v3 / (((v31 ^ v31)) | 1))
    v22 = v22
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v31
    _ = v32
    _ = v33
}
