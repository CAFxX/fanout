package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 bool; F1 bool; F2 int8 }
type GenS3 struct { F0 uint32; F1 int8 }
type GenS4 struct { F0 int64; F1 uint }
type GenS5 struct { F0 int; F1 bool; F2 uint }
type GenS6 struct { F0 uint16; F1 int64 }
func genFn0(p0 uint, p1 int16) uint64 {
    p1 = (p1 - 287)
    p0 += (^(p0))
    if (42 > 2) {
        if true {
            v7 := []int{189, -9223372036854775808, 32767}
            _ = v7
        } else {
            if true {
                if true {
                    switch ("a" + ("hello" + "")) {
                    case "a":
                        p0 -= 156
                    }
                } else {
                    v8 := []int{127, 127, -2147483648}
                    _ = v8
                }
            } else {
                v9 := []int{32768, 935}
                _ = v9
            }
            var v10 bool = false
            _ = v10
        }
        p0 -= (p0 + (p0 + p0))
    } else {
        switch "hello" {
        case "\n\t":
            var v11 uint32 = uint32(p0)
            _ = v11
        case "a":
            var v12 int64 = -1
            _ = v12
        case "":
            v13 := []int{9223372036854775807}
            _ = v13
        default:
            switch 4294967295 {
            case 2147483647:
                var v14 bool = ((128 <= 42) || (true || (32768 < 2147483647)))
                _ = v14
            case 4294967295:
                v15 := GenS1{F0: ("hello" + (("" + "") + ("hello" + "\n\t")))}
                _ = v15
            case 9223372036854775807:
                if false {
                    v16 := GenS1{F0: "a"}
                    p0 -= p0
                    var v17 int16 = 42
                    _ = v16
                    _ = v17
                }
            }
        }
    }
    if false {
        v18 := GenS2{F0: true, F1: (128 <= 32767), F2: -13}
        _ = v18
    } else {
        v19 := []int{}
        _ = v19
    }
    _ = p0
    _ = p1
    return 573
}
func genFn1(p0 uint64, p1 int64) int8 {
    var v20 uint16 = 32767
    var v21 uint32 = 864
    for fi22 := 0; fi22 < 19; fi22++ {
        v21 = v21
        _ = fi22
    }
    _ = p0
    _ = p1
    _ = v20
    _ = v21
    return -84
}

func main() {
    var v1 int8 = 42
    v2 := make(map[uint32]int)
    v3 := make(map[int]uint32)
    v4 := make(map[int]int)
    v5 := []uint32{}
    var v6 uint8 = 255
    if false {
        v1 -= 100
    }
    v23 := GenS3{F0: 632, F1: 100}
    var v24 int8 = int8(v6)
    switch (("hello" + "hello") + "foo bar") {
    case "a":
        if len(v5) > 0 {
            v5[0] = 4294967295
        }
    case "":
        for fi25 := 0; fi25 < 3; fi25++ {
            var v26 int64 = -9223372036854775808
            _ = len(v5)
            _ = fi25
            _ = v26
        }
    case "\n\t":
        var v27 uint32 = 171
        _ = v27
    }
    if true {
        if (!(false && true)) {
            v28 := GenS4{F0: int64(v6), F1: 218}
            _ = v28
        } else {
            var v29 int64 = -1
            v1 = (v24 * 63)
            _ = v29
        }
        for fi30 := 0; fi30 < 18; fi30++ {
            if false {
                for fi31 := 0; fi31 < 14; fi31++ {
                    var v32 uint16 = 255
                    _ = fi31
                    _ = v32
                }
                for fi33 := 0; fi33 < 20; fi33++ {
                    v24 = v24
                    _ = fi33
                }
                v1 = v24
            } else {
                for fi34 := 0; fi34 < 13; fi34++ {
                    if (866 != 642) {
                        v1 = (^(v24))
                    }
                    _ = fi34
                }
                for fi35 := 0; fi35 < 14; fi35++ {
                    var v36 float32 = float32(1.5)
                    var v37 uint = 811
                    _ = fi35
                    _ = v36
                    _ = v37
                }
                v1 = v1
            }
            _ = fi30
        }
        for fi38 := 0; fi38 < 14; fi38++ {
            v39 := GenS5{F0: fi38, F1: false, F2: 42}
            v40 := GenS5{F0: (fi38 >> ((805) & 63)), F1: (true || (9223372036854775807 == -9223372036854775808)), F2: 815}
            _ = fi38
            _ = v39
            _ = v40
        }
    } else {
        fmt.Println(v24)
    }
    if (53 >= 42) {
        var v41 uint64 = 230
        var v42 int8 = (v24 & v1)
        _ = v41
        _ = v42
    } else {
        v24 -= v1
    }
    var v43 float32 = 1.5
    v44 := GenS6{F0: 351, F1: 1}
    fmt.Println(v1, v6)
    for fi45 := 0; fi45 < 10; fi45++ {
        v46 := v2[2147483647]
        var v47 uint8 = (v6 * v6)
        var v48 uint16 = 271
        _ = fi45
        _ = v46
        _ = v47
        _ = v48
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
