package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32 }
type GenS2 struct { F0 int32; F1 int16; F2 int32 }
type GenS3 struct { F0 uint8; F1 string }
type GenS4 struct { F0 uint32; F1 uint8; F2 int64 }
func genFn0(p0 uint8, p1 int32) int8 {
    for fi7 := 0; fi7 < 12; fi7++ {
        for fi8 := 0; fi8 < 19; fi8++ {
            p0 = p0
            _ = fi8
        }
        var v9 uint = 42
        _ = fi7
        _ = v9
    }
    var v10 uint8 = (^(p0))
    v11 := []int{-2147483648, 32768, -9223372036854775808, -165}
    var v12 uint8 = (v10 << ((698) & 63))
    if true {
        fmt.Println(p1, v12, v10)
        v10 = p0
    } else {
        fmt.Println(v10, p0, v12)
        for fi13 := 0; fi13 < 1; fi13++ {
            var v14 string = ""
            v15 := GenS1{F0: v12, F1: p1}
            _ = fi13
            _ = v14
            _ = v15
        }
        var v16 bool = (true && false)
        _ = v16
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v12
    return 100
}
func genFn1(p0 uint64, p1 uint8) int32 {
    p0 += p0
    switch "\n\t" {
    case "foo bar":
        fmt.Println(p1, p1)
    }
    fmt.Println(p1, p1)
    v17 := []int{2147483647, -2147483648}
    _ = p0
    _ = p1
    _ = v17
    return -2147483648
}

func main() {
    var v1 int16 = 891
    var v2 uint8 = 248
    var v3 int32 = 513
    v4 := []uint32{4294967295, 0, 4294967295}
    v5 := []uint32{}
    var v6 int16 = v1
    v1 -= (v1 & v1)
    v3 += v3
    if (false || (0 > -967)) {
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    } else {
        switch 42 {
        case -335:
            switch 32767 {
            case 9223372036854775807:
                switch "hello" {
                case "":
                    v18 := GenS1{F0: (v2 ^ v2), F1: int32(v2)}
                    _ = v18
                }
            }
        default:
            fmt.Println(v1, v6, v3)
        }
        v19 := GenS2{F0: (^(v3)), F1: (v1 | v6), F2: genFn1(0, v2)}
        var v20 int64 = 699
        _ = v19
        _ = v20
    }
    fmt.Println(v1)
    for fi21 := 0; fi21 < 5; fi21++ {
        if (684 == 555) {
            fmt.Println(fi21, v3)
            var v22 uint64 = 1
            var v23 float32 = 1e10
            _ = v22
            _ = v23
        }
        switch (("foo bar" + "hello") + "hello") {
        case "foo bar":
            v6 = v6
        case "a":
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
        }
        switch 32768 {
        case 9223372036854775807:
            v24 := GenS3{F0: 2, F1: ""}
            _ = v24
        }
        _ = fi21
    }
    var v25 uint = 65535
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v26 bool = true
    for fi27 := 0; fi27 < 8; fi27++ {
        v28 := GenS4{F0: 754, F1: 1, F2: 42}
        var v29 int8 = 100
        fmt.Println(v3, v2)
        _ = fi27
        _ = v28
        _ = v29
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
