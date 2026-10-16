package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8; F1 uint32; F2 int32 }
type GenS3 struct { F0 uint64; F1 int64; F2 int8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int64; F1 int; F2 uint32 }
func genFn0(p0 uint64, p1 int32, p2 uint8) string {
    for fi4 := 0; fi4 < 18; fi4++ {
        switch "foo bar" {
        case "\n\t":
            fmt.Println(p0, p1, p1)
        case "hello":
            v5 := GenS1{F0: 255}
            _ = v5
        default:
            p1 += -10
        }
        var v6 uint16 = 0
        v7 := []int{42, -1, 0, fi4}
        _ = fi4
        _ = v6
        _ = v7
    }
    p1 = (-(p1))
    v8 := []int{}
    v9 := GenS2{F0: -1, F1: 2, F2: 0}
    p1 += (-(p1))
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return ("hello" + "a")
}
func genFn1() int16 {
    for fi10 := 0; fi10 < 5; fi10++ {
        v11 := []int{(^(fi10)), (fi10 * 9223372036854775807)}
        var v12 int64 = -562
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 float64 = -2.25
    _ = v13
    return 371
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[uint32]string)
    v3 := make(map[int]bool)
    if false {
        for fi14 := 0; fi14 < 20; fi14++ {
            var v15 uint = 1
            v16 := v3[fi14]
            _ = fi14
            _ = v15
            _ = v16
        }
        var v17 uint = 1
        _ = v17
    } else {
        if (-223 != 0) {
            v18 := GenS3{F0: 2, F1: -951, F2: -116}
            var v19 uint16 = 42
            _ = v18
            _ = v19
        } else {
            var v20 int16 = 32767
            for fi21 := 0; fi21 < 6; fi21++ {
                v20 += (v20 >> ((42) & 63))
                _ = fi21
            }
            if true {
                if (true && (genFn1() != v20)) {
                    for fi22 := 0; fi22 < 17; fi22++ {
                        v23 := []int{}
                        _ = fi22
                        _ = v23
                    }
                    v20 += (-(v20))
                    v24 := GenS4{F0: -2147483648}
                    _ = v24
                } else {
                    fmt.Println(v20)
                    v25 := GenS5{F0: 29}
                    v20 -= v20
                    _ = v25
                }
                if (4294967295 == 123) {
                    var v26 int8 = 0
                    v27 := []int{245, 42, -471}
                    _ = v26
                    _ = v27
                }
            }
            _ = v20
        }
    }
    var v28 int = 128
    if (float32(-2.25) > 0.1) {
        switch ("hello" + "foo bar") {
        case "foo bar":
            if true {
                for fi29 := 0; fi29 < 8; fi29++ {
                    v30 := GenS2{F0: 1, F1: 938, F2: 0}
                    _ = fi29
                    _ = v30
                }
                for fi31 := 0; fi31 < 11; fi31++ {
                    v32 := GenS1{F0: 102}
                    if false {
                        v28 = 804
                    }
                    v33 := GenS3{F0: 4294967295, F1: 42, F2: -38}
                    _ = fi31
                    _ = v32
                    _ = v33
                }
                _ = len(v1)
            }
        case "hello":
            var v34 uint = 786
            _ = v34
        }
    }
    v35 := []int{v28, 42}
    for fi36 := 0; fi36 < 7; fi36++ {
        v37 := GenS6{F0: -9223372036854775808}
        _ = fi36
        _ = v37
    }
    fmt.Println(v28)
    v38 := v2[1]
    var v39 float64 = 3.14159
    var v40 bool = (false || true)
    v41 := GenS7{F0: 1, F1: v28, F2: 2}
    v42 := GenS6{F0: -812}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v35
    _ = v38
    _ = v39
    _ = v40
    _ = v41
    _ = v42
}
