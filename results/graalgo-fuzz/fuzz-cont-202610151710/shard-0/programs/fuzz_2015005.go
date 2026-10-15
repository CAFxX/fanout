package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 uint64; F1 uint16 }
type GenS3 struct { F0 bool; F1 string }
type GenS4 struct { F0 int8; F1 int8; F2 bool }
type GenS5 struct { F0 uint32; F1 int }
func genFn0(p0 uint16) uint {
    if (true || true) {
        if (true || (100 > 42)) {
            fmt.Println(p0)
            v7 := []int{}
            _ = v7
        } else {
            var v8 int32 = -777
            _ = v8
        }
    } else {
        var v9 uint64 = 949
        _ = v9
    }
    v10 := GenS1{F0: 368, F1: true}
    v11 := []int{9223372036854775807, 42, 1, -140}
    v12 := GenS2{F0: 4294967295, F1: (p0 << ((223) & 63))}
    v13 := GenS2{F0: 2, F1: p0}
    _ = p0
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    return 487
}

func main() {
    v1 := make(map[string]int)
    var v2 int = 284
    var v3 float64 = float64(3.14159)
    var v4 uint = 255
    var v5 int32 = 1
    var v6 int16 = 42
    v14 := GenS3{F0: true, F1: ("foo bar" + (("foo bar" + "") + "\n\t"))}
    for fi15 := 0; fi15 < 18; fi15++ {
        v16 := []int{(-(fi15)), (-(fi15)), (fi15 >> ((v4) & 63))}
        var v17 int64 = -9223372036854775808
        var v18 string = ("hello" + "a")
        _ = fi15
        _ = v16
        _ = v17
        _ = v18
    }
    var v19 uint32 = 0
    if (false || false) {
        var v20 int64 = -171
        for fi21 := 0; fi21 < 3; fi21++ {
            v19 -= v19
            fmt.Println(v20, v3, v3)
            _ = fi21
        }
        if true {
            if false {
                fmt.Println(v6, v4)
                v19 += v19
                v22 := v1[("hello" + ("a" + ""))]
                _ = v22
            } else {
                var v23 uint8 = 1
                v24 := v1["foo bar"]
                switch (v24 ^ (v24 / ((v2) | 1))) {
                case -2147483648:
                    var v25 bool = true
                    _ = v25
                case 9223372036854775807:
                    var v26 float32 = 1.5
                    _ = v26
                }
                _ = v23
                _ = v24
            }
            v27 := GenS3{F0: true, F1: "a"}
            for fi28 := 0; fi28 < 13; fi28++ {
                v1["a"] = 269
                var v29 bool = (!((v4 ^ v4) == v4))
                var v30 float64 = (1.5 - (3.14159 + (float64(3.14159) - v3)))
                _ = fi28
                _ = v29
                _ = v30
            }
            _ = v27
        } else {
            for fi31 := 0; fi31 < 2; fi31++ {
                var v32 uint64 = 4294967295
                switch (fi31 & fi31) {
                case -2147483648:
                    var v33 uint64 = 1
                    _ = v33
                default:
                    if false {
                        v34 := GenS3{F0: (v4 <= v4), F1: "a"}
                        v20 -= v20
                        _ = v34
                    }
                }
                _ = fi31
                _ = v32
            }
        }
        _ = v20
    } else {
        var v35 uint32 = (v19 + v19)
        v1[("a" + ("a" + "hello"))] = -392
        var v36 uint = (v4 | (v4 - 42))
        _ = v35
        _ = v36
    }
    switch "\n\t" {
    case "a":
        fmt.Println(v4, v6, v2)
    default:
        v37 := GenS4{F0: -56, F1: 18, F2: false}
        _ = v37
    }
    v19 -= 42
    for fi38 := 0; fi38 < 8; fi38++ {
        var v39 int8 = int8(v6)
        v40 := GenS3{F0: (v3 >= v3), F1: "a"}
        _ = fi38
        _ = v39
        _ = v40
    }
    v3 -= float64(3.14159)
    if true {
        if true {
            for fi41 := 0; fi41 < 17; fi41++ {
                v2 = fi41
                _ = fi41
            }
            var v42 uint32 = 720
            var v43 uint8 = 2
            _ = v42
            _ = v43
        } else {
            if false {
                for fi44 := 0; fi44 < 8; fi44++ {
                    v45 := GenS5{F0: 4294967295, F1: -595}
                    _ = fi44
                    _ = v45
                }
            } else {
                var v46 uint = 998
                v47 := []int{2147483647, (v2 * 798), (^(v2)), 42}
                v19 -= 767
                _ = v46
                _ = v47
            }
            var v48 uint32 = v19
            var v49 uint64 = 42
            _ = v48
            _ = v49
        }
    } else {
        var v50 int32 = v5
        v6 += -32768
        _ = v50
    }
    v3 += (float64((0.1 + 0.1)) + 3.14159)
    for fi51 := 0; fi51 < 10; fi51++ {
        fmt.Println(v19)
        var v52 uint = (^(v4))
        fmt.Println(v2)
        _ = fi51
        _ = v52
    }
    if false {
        fmt.Println(v2, v4, v5)
        v2 -= (v2 / (((v2 - v2)) | 1))
    }
    switch (v2 ^ 105) {
    case 32767:
        v19 += 65535
    case -399:
        v3 += v3
    default:
        fmt.Println(v6)
    }
    var v53 uint = 555
    v54 := GenS3{F0: false, F1: "hello"}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v53:", v53)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v53
    _ = v54
    _ = v6
}
