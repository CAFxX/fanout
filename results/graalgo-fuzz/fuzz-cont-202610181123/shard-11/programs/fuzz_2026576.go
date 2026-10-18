package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 int }
type GenS2 struct { F0 bool; F1 int }
type GenS3 struct { F0 int32; F1 uint64; F2 int8 }
type GenS4 struct { F0 bool; F1 int32; F2 int8 }
type GenS5 struct { F0 uint64; F1 int }
func genFn0() int16 {
    v7 := GenS1{F0: 0, F1: 1, F2: 487}
    for fi8 := 0; fi8 < 19; fi8++ {
        var v9 int16 = 727
        _ = fi8
        _ = v9
    }
    var v10 int16 = -32768
    if (110 >= 1) {
        v11 := []int{127, 742, 4294967295}
        var v12 uint32 = 1
        _ = v11
        _ = v12
    } else {
        v13 := []int{-1, 2147483647, 35}
        _ = v13
    }
    if ((1.5 * 0.1) > 0.1) {
        var v14 float64 = 3.14159
        v10 = 1
        v14 += float64(float32(v14))
        _ = v14
    }
    _ = v10
    _ = v7
    return -1
}
func genFn1() int32 {
    var v15 string = ("foo bar" + "")
    var v16 bool = (!(0.0 >= 0.1))
    v16 = v16
    v15 = "foo bar"
    v17 := []int{}
    _ = v15
    _ = v16
    _ = v17
    return 1
}

func main() {
    v1 := make(map[int]int)
    v2 := []string{"a", "a"}
    var v3 int32 = -2147483648
    v4 := make(map[string]string)
    var v5 bool = false
    v6 := make(map[int]string)
    var v18 int16 = -48
    if (0 != 65535) {
        var v19 uint = 211
        _ = v19
    } else {
        for fi20 := 0; fi20 < 5; fi20++ {
            if ((fi20 + fi20) == (fi20 - fi20)) {
                for fi21 := 0; fi21 < 20; fi21++ {
                    v5 = false
                    v22 := GenS2{F0: (v5 && false), F1: fi20}
                    _ = fi21
                    _ = v22
                }
                v5 = v5
                var v23 uint = 0
                _ = v23
            } else {
                v18 = genFn0()
            }
            var v24 uint64 = 0
            var v25 uint32 = 1
            _ = fi20
            _ = v24
            _ = v25
        }
    }
    fmt.Println(v18)
    v26 := GenS3{F0: v3, F1: 501, F2: 100}
    switch 2147483647 {
    case -151:
        var v27 int16 = v18
        _ = v27
    case 530:
        _ = len(v2)
    default:
        v28 := GenS3{F0: genFn1(), F1: 4294967295, F2: -1}
        _ = v28
    }
    fmt.Println(v5)
    if true {
        v29 := GenS4{F0: ((!(42 != 177)) && (-1 != 525)), F1: v3, F2: 1}
        if (-2.25 != 0.1) {
            switch "a" {
            case "":
                var v30 bool = v5
                _ = v30
            case "foo bar":
                var v31 int64 = -9223372036854775808
                _ = v31
            case "\n\t":
                v32 := GenS5{F0: 42, F1: 127}
                _ = v32
            default:
                v33 := GenS2{F0: v5, F1: -498}
                _ = v33
            }
        }
        _ = v29
    }
    v3 += (v3 << ((610) & 63))
    switch "hello" {
    case "":
        for fi34 := 0; fi34 < 3; fi34++ {
            if (v5 || (!v5)) {
                _ = len(v2)
            }
            _ = fi34
        }
    case "a":
        for fi35 := 0; fi35 < 8; fi35++ {
            if true {
                if v5 {
                    var v36 uint32 = 777
                    v37 := GenS2{F0: (true && (2147483647 <= (fi35 | fi35))), F1: -783}
                    var v38 uint16 = 2
                    _ = v36
                    _ = v37
                    _ = v38
                } else {
                    switch (fi35 & (fi35 | fi35)) {
                    case 42:
                        var v39 bool = (v5 || (42 <= 1))
                        _ = v39
                    }
                }
                if (!true) {
                    var v40 uint64 = 632
                    fmt.Println(v40, fi35)
                    _ = v40
                } else {
                    _ = len(v2)
                }
                var v41 int16 = v18
                _ = v41
            }
            fmt.Println(v5)
            _ = fi35
        }
    }
    v18 -= v18
    var v42 uint32 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v18
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v6
}
