package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 bool }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint8; F1 int32; F2 uint }
type GenS4 struct { F0 int8; F1 string; F2 uint16 }
func genFn0(p0 int64, p1 uint64) int {
    if false {
        var v5 int16 = 1
        _ = v5
    }
    var v6 int8 = -1
    for fi7 := 0; fi7 < 4; fi7++ {
        switch "a" {
        case "a":
            switch (fi7 >> ((0) & 63)) {
            case 128:
                fmt.Println(p0, p0)
            case 32768:
                fmt.Println(p1, p1)
            case 4294967295:
                v8 := GenS1{F0: -964, F1: 1, F2: false}
                _ = v8
            }
        case "":
            var v9 int16 = 92
            _ = v9
        case "foo bar":
            switch 2147483647 {
            case -2147483648:
                for fi10 := 0; fi10 < 12; fi10++ {
                    p0 = p0
                    p0 += (p0 * (p0 * p0))
                    _ = fi10
                }
            case 755:
                var v11 string = "foo bar"
                _ = v11
            case 315:
                for fi12 := 0; fi12 < 6; fi12++ {
                    v13 := []int{901, -896}
                    if len(v13) > 0 {
                        v13[0] = fi7
                    }
                    _ = fi12
                    _ = v13
                }
            }
        }
        p1 = (p1 >> ((289) & 63))
        p1 = (p1 & p1)
        _ = fi7
    }
    switch "a" {
    case "a":
        v14 := GenS2{F0: -1}
        _ = v14
    case "hello":
        var v15 uint64 = (p1 | 4294967295)
        _ = v15
    }
    _ = p0
    _ = p1
    _ = v6
    return -1
}
func genFn1() int64 {
    var v16 uint64 = 18446744073709551615
    for fi17 := 0; fi17 < 19; fi17++ {
        if ((fi17 + (fi17 >> ((42) & 63))) == (fi17 | (-(fi17)))) {
            v18 := []int{fi17, (fi17 ^ (fi17 % (((fi17 - (fi17 / ((-266) | 1)))) | 1))), (^(fi17)), (fi17 % ((32767) | 1))}
            _ = v18
        }
        v19 := GenS2{F0: 42}
        _ = fi17
        _ = v19
    }
    _ = v16
    return 0
}

func main() {
    var v1 uint = 936
    v2 := []string{"foo bar"}
    var v3 string = ("\n\t" + "\n\t")
    v4 := make(map[int]int)
    var v20 int32 = -509
    var v21 float64 = ((0.1 - 3.14159) + 1.5)
    v3 = v3
    v22 := GenS3{F0: 128, F1: v20, F2: v1}
    v23 := GenS4{F0: -128, F1: v3, F2: 32767}
    v24 := GenS2{F0: 1}
    switch -9223372036854775808 {
    case -2147483648:
        if len(v2) > 0 {
            v2[0] = ("a" + "hello")
        }
    case 193:
        switch "" {
        case "hello":
            fmt.Println(v3, v20, v1)
        case "a":
            for fi25 := 0; fi25 < 9; fi25++ {
                v26 := GenS1{F0: 0, F1: 1, F2: false}
                _ = fi25
                _ = v26
            }
        case "\n\t":
            var v27 uint8 = 42
            _ = v27
        }
    }
    var v28 int32 = -1
    v3 = v3
    switch 42 {
    case 9223372036854775807:
        fmt.Println(v21, v21, v3)
    default:
        if false {
            delete(v4, 4294967295)
            if (792 <= 0) {
                v4[32767] = 0
            }
            var v29 int = genFn0(-9223372036854775808, 4294967295)
            _ = v29
        } else {
            var v30 int64 = 1
            _ = v30
        }
    }
    var v31 uint8 = 42
    var v32 uint64 = 18446744073709551615
    v1 += 487
    var v33 string = (v3 + "")
    v28 = v20
    for fi34 := 0; fi34 < 7; fi34++ {
        fmt.Println(v31)
        _ = fi34
    }
    switch (v3 + v3) {
    case "\n\t":
        var v35 uint32 = 4294967295
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v28
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
}
