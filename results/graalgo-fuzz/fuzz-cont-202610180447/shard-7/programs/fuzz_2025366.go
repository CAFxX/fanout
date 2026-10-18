package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64; F1 int8; F2 uint32 }
func genFn0(p0 uint16, p1 uint8) uint64 {
    for fi5 := 0; fi5 < 7; fi5++ {
        fmt.Println(fi5, p0, fi5)
        var v6 uint8 = (p1 % ((p1) | 1))
        v7 := []int{fi5, fi5, fi5}
        _ = fi5
        _ = v6
        _ = v7
    }
    if true {
        p1 += (p1 / (((p1 * 2)) | 1))
        var v8 int8 = -27
        if (!false) {
            v9 := GenS1{F0: 42, F1: 1}
            v10 := []int{}
            var v11 uint64 = uint64(p1)
            _ = v10
            _ = v11
            _ = v9
        } else {
            v12 := GenS2{F0: 42}
            _ = v12
        }
        _ = v8
    }
    _ = p0
    _ = p1
    return 18446744073709551615
}
func genFn1(p0 int32) int16 {
    v13 := []int{2147483647}
    fmt.Println(p0)
    for fi14 := 0; fi14 < 10; fi14++ {
        v15 := GenS1{F0: 759, F1: 964}
        p0 = (p0 & p0)
        var v16 uint16 = 486
        _ = fi14
        _ = v15
        _ = v16
    }
    fmt.Println(p0)
    v17 := GenS3{F0: 18446744073709551615, F1: 42, F2: 2147483648}
    _ = p0
    _ = v13
    _ = v17
    return -358
}

func main() {
    var v1 string = ""
    var v2 uint16 = 255
    v3 := []uint32{65535, 224, 42, 208}
    var v4 int64 = -9223372036854775808
    if true {
        v2 -= v2
    }
    v18 := GenS1{F0: -9223372036854775808, F1: 710}
    switch 32768 {
    case 323:
        var v19 uint8 = 121
        _ = v19
    case 32767:
        v4 -= (v4 << ((42) & 63))
    }
    for fi20 := 0; fi20 < 12; fi20++ {
        for fi21 := 0; fi21 < 11; fi21++ {
            switch (v1 + "") {
            case "hello":
                var v22 uint = 18446744073709551615
                _ = v22
            }
            fmt.Println(fi20, fi21)
            _ = fi21
        }
        _ = fi20
    }
    var v23 string = (v1 + ("" + (v1 + ((v1 + v1) + v1))))
    if true {
        if (40 != 375) {
            v24 := GenS2{F0: 1}
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
            _ = v24
        }
        if true {
            _ = len(v3)
        } else {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
    } else {
        switch v1 {
        case "":
            var v25 bool = (42 <= -124)
            _ = v25
        case "\n\t":
            fmt.Println(v1, v2)
        }
    }
    var v26 int = 32767
    fmt.Println(v1, v4)
    var v27 int16 = 42
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v23
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
