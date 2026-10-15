package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64; F2 int }
type GenS2 struct { F0 string; F1 uint32 }
type GenS3 struct { F0 uint64; F1 int16 }
type GenS4 struct { F0 int8; F1 int64 }
type GenS5 struct { F0 bool; F1 int32 }
func genFn0() int8 {
    for fi5 := 0; fi5 < 16; fi5++ {
        v6 := GenS1{F0: 4294967295, F1: 0, F2: fi5}
        _ = fi5
        _ = v6
    }
    var v7 float32 = -2.25
    if true {
        fmt.Println(v7)
        v8 := GenS2{F0: "hello", F1: 2147483647}
        for fi9 := 0; fi9 < 20; fi9++ {
            fmt.Println(v7, v7)
            _ = fi9
        }
        _ = v8
    } else {
        var v10 uint64 = 18446744073709551615
        _ = v10
    }
    v11 := GenS3{F0: 4294967295, F1: 1}
    for fi12 := 0; fi12 < 13; fi12++ {
        fmt.Println(v7, fi12)
        _ = fi12
    }
    _ = v11
    _ = v7
    return 72
}
func genFn1(p0 uint8, p1 uint16, p2 int64) uint8 {
    if (p1 <= (p1 | (p1 % (((p1 / ((p1) | 1))) | 1)))) {
        v13 := []int{}
        switch 32767 {
        case 2147483647:
            p2 = (p2 % (((p2 - p2)) | 1))
        case 32768:
            var v14 uint32 = 1
            _ = v14
        default:
            p1 = 42
        }
        _ = v13
    } else {
        var v15 int16 = 679
        if true {
            var v16 int16 = 42
            v17 := []int{-134}
            v18 := GenS1{F0: -9223372036854775808, F1: 623, F2: 4294967295}
            _ = v16
            _ = v17
            _ = v18
        }
        _ = v15
    }
    switch ("a" + (("a" + "") + "")) {
    case "foo bar":
        if true {
            v19 := GenS4{F0: 100, F1: -851}
            fmt.Println(p0)
            v20 := []int{788, -632}
            _ = v19
            _ = v20
        } else {
            v21 := []int{}
            switch 32767 {
            case 411:
                var v22 int64 = p2
                _ = v22
            }
            _ = v21
        }
    case "":
        fmt.Println(p1, p1, p0)
    default:
        if false {
            if true {
                var v23 int32 = -1
                _ = v23
            } else {
                var v24 string = "foo bar"
                v25 := GenS1{F0: 127, F1: 662, F2: 352}
                _ = v24
                _ = v25
            }
        }
    }
    switch -2147483648 {
    case -1:
        switch 42 {
        case 32767:
            var v26 uint = 0
            _ = v26
        }
    case 558:
        p1 = (p1 + 32767)
    }
    fmt.Println(p2, p1, p0)
    var v27 int64 = p2
    _ = p0
    _ = p1
    _ = p2
    _ = v27
    return uint8(p1)
}

func main() {
    var v1 int8 = -1
    var v2 int16 = 0
    var v3 int16 = v2
    var v4 int8 = (v1 << ((4294967295) & 63))
    v3 += (v3 << ((225) & 63))
    fmt.Println(v3)
    for fi28 := 0; fi28 < 11; fi28++ {
        fmt.Println(v2, v4)
        v2 += 0
        for fi29 := 0; fi29 < 5; fi29++ {
            var v30 int = 32767
            v4 = 127
            fmt.Println(v30, v30)
            _ = fi29
            _ = v30
        }
        _ = fi28
    }
    v3 -= 287
    v1 = v4
    var v31 uint16 = 32767
    if true {
        fmt.Println(v31)
    } else {
        var v32 bool = true
        _ = v32
    }
    v33 := GenS5{F0: (0.0 == 1.5), F1: -1}
    var v34 float32 = 1e10
    var v35 uint16 = 820
    v35 = (^(v31))
    v36 := []int{-2147483648, 32767}
    v37 := GenS5{F0: true, F1: 0}
    fmt.Println(v31, v31, v1)
    v38 := GenS5{F0: (false && false), F1: 0}
    v1 += (v4 % ((-128) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("v34:", v34)
    fmt.Println("v35:", v35)
    fmt.Println("slicelen v36:", len(v36), v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v4
}
