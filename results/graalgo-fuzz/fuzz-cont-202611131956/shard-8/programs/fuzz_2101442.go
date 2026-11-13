package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 string; F1 int32 }
type GenS3 struct { F0 uint; F1 uint16; F2 uint16 }
type GenS4 struct { F0 uint64 }
func genFn0(p0 uint, p1 uint64, p2 uint64) uint16 {
    switch 42 {
    case 1:
        v5 := []int{2147483647, 431}
        _ = v5
    case -9223372036854775808:
        var v6 uint64 = (p2 - p2)
        _ = v6
    }
    v7 := GenS1{F0: (p2 ^ p1)}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 65535
}
func genFn1() int8 {
    var v8 float32 = 0.0
    v8 -= float32(3.14159)
    v8 = 0.1
    v8 -= ((float32(1.5) + (0.1 - v8)) + (v8 * v8))
    _ = v8
    return -128
}

func main() {
    v1 := []int{233}
    v2 := []uint32{}
    v3 := []uint32{2147483648}
    var v4 string = "foo bar"
    v9 := GenS2{F0: "foo bar", F1: 549}
    for fi10 := 0; fi10 < 20; fi10++ {
        for fi11 := 0; fi11 < 15; fi11++ {
            var v12 uint = 4294967295
            var v13 int64 = 509
            v13 = (v13 & (v13 / (((v13 * -1)) | 1)))
            _ = fi11
            _ = v12
            _ = v13
        }
        v4 = v4
        v14 := GenS1{F0: 1}
        _ = fi10
        _ = v14
    }
    for fi15 := 0; fi15 < 2; fi15++ {
        fmt.Println(v4, fi15)
        _ = fi15
    }
    if ((!false) && true) {
        switch 128 {
        case 709:
            switch ((v4 + "hello") + v4) {
            case "hello":
                var v16 float64 = -2.25
                _ = v16
            case "\n\t":
                v17 := GenS1{F0: 778}
                _ = v17
            case "a":
                if len(v3) > 0 {
                    v3[0] = 2
                }
            }
        case 74:
            var v18 bool = (v4 <= v4)
            _ = v18
        default:
            v4 = ("hello" + ((((("" + "foo bar") + v4) + (v4 + (v4 + v4))) + ("" + "")) + v4))
        }
    } else {
        if false {
            var v19 uint8 = 5
            if (255 <= 32767) {
                v19 += v19
                v20 := GenS2{F0: (((v4 + v4) + "\n\t") + v4), F1: 251}
                for fi21 := 0; fi21 < 1; fi21++ {
                    switch v4 {
                    case "\n\t":
                        fmt.Println(fi21, fi21, v19)
                    default:
                        var v22 int = -41
                        _ = v22
                    }
                    _ = fi21
                }
                _ = v20
            }
            _ = v19
        }
    }
    if (v4 <= "\n\t") {
        v4 = "\n\t"
        v23 := GenS3{F0: 2, F1: 2, F2: 1}
        fmt.Println(v4)
        _ = v23
    }
    v24 := GenS1{F0: 42}
    var v25 int64 = 42
    for fi26 := 0; fi26 < 5; fi26++ {
        if true {
            if len(v3) > 0 {
                v3[0] = 405
            }
        } else {
            fmt.Println(fi26, v25, v25)
        }
        for fi27 := 0; fi27 < 9; fi27++ {
            if (!(-2147483648 < -641)) {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            _ = fi27
        }
        switch (fi26 << ((65535) & 63)) {
        case 128:
            v4 = "hello"
        case -1:
            var v28 int32 = 1
            _ = v28
        case 0:
            if len(v2) > 0 {
                v2[0] = 858
            }
        }
        _ = fi26
    }
    if (127 <= 42) {
        var v29 uint = 255
        _ = v29
    }
    v25 += (v25 * (v25 & v25))
    switch "\n\t" {
    case "a":
        for fi30 := 0; fi30 < 10; fi30++ {
            v31 := GenS4{F0: 4294967295}
            v32 := GenS1{F0: 1}
            _ = fi30
            _ = v31
            _ = v32
        }
    case "":
        fmt.Println(v4)
    case "\n\t":
        var v33 uint32 = 2147483648
        _ = v33
    default:
        if false {
            for fi34 := 0; fi34 < 13; fi34++ {
                var v35 int32 = 0
                _ = fi34
                _ = v35
            }
        }
    }
    fmt.Println(v25)
    if len(v1) > 0 {
        v1[0] = 610
    }
    var v36 float64 = 1e10
    switch -536 {
    case -1:
        var v37 int64 = v25
        _ = v37
    case -597:
        fmt.Println(v36, v25, v25)
    case -2147483648:
        var v38 uint32 = 2147483647
        _ = v38
    default:
        v36 = (-2.25 + (3.14159 / ((0.1 - v36) + 1e10)))
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v3
    _ = v36
    _ = v4
    _ = v9
}
