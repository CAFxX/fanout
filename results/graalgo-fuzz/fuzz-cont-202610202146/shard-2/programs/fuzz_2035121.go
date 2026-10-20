package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 int64 }
func genFn0(p0 int8, p1 int16, p2 uint64) int {
    var v5 int16 = 1
    if true {
        var v6 uint = 935
        v7 := []int{32768}
        p1 = 299
        _ = v6
        _ = v7
    }
    for fi8 := 0; fi8 < 10; fi8++ {
        switch fi8 {
        case 127:
            for fi9 := 0; fi9 < 9; fi9++ {
                var v10 int16 = -1
                _ = fi9
                _ = v10
            }
        case 2147483647:
            switch fi8 {
            case 128:
                var v11 int64 = 34
                _ = v11
            case -35:
                var v12 int16 = (v5 % (((p1 ^ 0)) | 1))
                _ = v12
            }
        case -56:
            if true {
                v5 -= p1
                switch (fi8 % ((fi8) | 1)) {
                case 1:
                    var v13 int8 = p0
                    _ = v13
                case 42:
                    switch -549 {
                    case 164:
                        v5 = v5
                    case 942:
                        var v14 uint8 = 1
                        _ = v14
                    default:
                        var v15 int8 = (^(p0))
                        _ = v15
                    }
                default:
                    v16 := GenS1{F0: "", F1: ""}
                    _ = v16
                }
                var v17 float32 = 0.0
                _ = v17
            }
        default:
            v18 := GenS1{F0: (("foo bar" + "") + (("hello" + ("hello" + "hello")) + (("hello" + ("\n\t" + "a")) + "a"))), F1: "hello"}
            _ = v18
        }
        _ = fi8
    }
    v19 := []int{int(v5), 32767, 44, 128}
    _ = p0
    _ = p1
    _ = p2
    _ = v19
    _ = v5
    return 199
}

func main() {
    var v1 int64 = -9223372036854775808
    var v2 string = ("\n\t" + ("" + ""))
    var v3 uint8 = 14
    var v4 bool = false
    if ((v2 + v2) != v2) {
        v20 := GenS2{F0: 0, F1: 535}
        _ = v20
    }
    var v21 uint64 = 18446744073709551615
    for fi22 := 0; fi22 < 2; fi22++ {
        v23 := []int{(fi22 % (((-(fi22))) | 1)), fi22}
        var v24 uint8 = (^(v3))
        _ = fi22
        _ = v23
        _ = v24
    }
    v21 += v21
    var v25 int32 = 2147483647
    for fi26 := 0; fi26 < 15; fi26++ {
        var v27 int8 = -128
        _ = fi26
        _ = v27
    }
    if (v1 >= (v1 - v1)) {
        switch 32767 {
        case 42:
            for fi28 := 0; fi28 < 12; fi28++ {
                for fi29 := 0; fi29 < 5; fi29++ {
                    v30 := GenS3{F0: v1}
                    v3 += (v3 >> ((644) & 63))
                    var v31 uint64 = v21
                    _ = fi29
                    _ = v30
                    _ = v31
                }
                _ = fi28
            }
        case 128:
            for fi32 := 0; fi32 < 11; fi32++ {
                fmt.Println(v2, v4)
                _ = fi32
            }
        }
        v1 -= (-(v1))
    }
    var v33 int8 = 42
    v4 = (v3 < (v3 + (v3 / ((v3) | 1))))
    if (32767 == -903) {
        v34 := []int{167}
        _ = v34
    } else {
        v35 := []int{42, 4294967295}
        _ = v35
    }
    var v36 uint64 = (v21 + (v21 >> ((65535) & 63)))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v25
    _ = v3
    _ = v33
    _ = v36
    _ = v4
}
