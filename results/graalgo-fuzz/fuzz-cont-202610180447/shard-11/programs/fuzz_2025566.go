package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 int; F1 string; F2 int8 }
type GenS4 struct { F0 int64; F1 uint64 }
func genFn0() int32 {
    var v6 int32 = 42
    v6 -= v6
    _ = v6
    return -757
}
func genFn1(p0 uint32, p1 int32, p2 uint16) bool {
    for fi7 := 0; fi7 < 7; fi7++ {
        v8 := []int{fi7, fi7}
        var v9 int16 = 0
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
        _ = fi7
        _ = v8
        _ = v9
    }
    if ((18446744073709551615 > 42) || false) {
        v10 := GenS1{F0: 255, F1: 2}
        if (742 != -665) {
            v11 := []int{42}
            _ = v11
        } else {
            var v12 bool = (!(-1 != 9223372036854775807))
            p1 = (p1 >> ((1) & 63))
            _ = v12
        }
        if true {
            fmt.Println(p0)
            if true {
                v13 := GenS1{F0: 0, F1: 18446744073709551615}
                v14 := GenS2{F0: 255, F1: 42}
                var v15 int8 = 42
                _ = v13
                _ = v14
                _ = v15
            } else {
                p1 += genFn0()
            }
        }
        _ = v10
    } else {
        if false {
            switch 539 {
            case 9223372036854775807:
                var v16 uint8 = 167
                _ = v16
            case 42:
                p0 += 827
            case 256:
                fmt.Println(p2, p0, p1)
            }
            v17 := GenS1{F0: 4294967295, F1: 2}
            var v18 float32 = 1e10
            _ = v17
            _ = v18
        } else {
            var v19 bool = false
            var v20 int8 = int8(p2)
            if (2 > 0) {
                var v21 uint8 = 2
                _ = v21
            } else {
                v22 := GenS1{F0: 1, F1: uint(p0)}
                fmt.Println(v20, p2, v20)
                _ = v22
            }
            _ = v19
            _ = v20
        }
    }
    fmt.Println(p1, p0, p2)
    _ = p0
    _ = p1
    _ = p2
    return ((0 < 87) && (95 <= -128))
}

func main() {
    var v1 bool = false
    v2 := make(map[int]int)
    v3 := []int{}
    var v4 uint = 65535
    v5 := make(map[int]uint32)
    v1 = (v1 && (-2147483648 != 1))
    fmt.Println(v4, v4)
    v23 := GenS3{F0: 32768, F1: "foo bar", F2: 102}
    v4 += v4
    for fi24 := 0; fi24 < 2; fi24++ {
        var v25 uint64 = 2
        var v26 string = "hello"
        _ = fi24
        _ = v25
        _ = v26
    }
    v1 = v1
    delete(v2, 32768)
    v1 = (127 > 1)
    switch -854 {
    case 32768:
        var v27 int32 = 1
        _ = v27
    default:
        if v1 {
            var v28 bool = (v1 || v1)
            for fi29 := 0; fi29 < 16; fi29++ {
                v30 := GenS2{F0: (^(v4)), F1: -78}
                v28 = v28
                var v31 int64 = 32
                _ = fi29
                _ = v30
                _ = v31
            }
            v32 := GenS4{F0: 0, F1: 1}
            _ = v28
            _ = v32
        }
    }
    switch 127 {
    case 2147483647:
        var v33 float64 = -2.25
        _ = v33
    case -9223372036854775808:
        v1 = (v1 && v1)
    case 554:
        var v34 string = ("\n\t" + "foo bar")
        _ = v34
    default:
        delete(v5, -261)
    }
    switch -140 {
    case 32767:
        var v35 uint = (v4 ^ (v4 | (v4 | (v4 - (v4 ^ v4)))))
        _ = v35
    default:
        _ = len(v2)
    }
    switch "hello" {
    case "\n\t":
        switch "\n\t" {
        case "":
            v1 = v1
        }
    }
    _ = len(v5)
    var v36 int16 = 42
    for fi37 := 0; fi37 < 11; fi37++ {
        v1 = false
        var v38 int32 = genFn0()
        _ = fi37
        _ = v38
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v36
    _ = v4
    _ = v5
}
