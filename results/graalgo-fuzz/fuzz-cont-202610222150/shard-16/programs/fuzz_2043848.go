package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int32 }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 uint16; F1 int; F2 int32 }
type GenS4 struct { F0 string }
func genFn0() int32 {
    v6 := []int{}
    var v7 uint8 = 127
    v7 += v7
    for fi8 := 0; fi8 < 18; fi8++ {
        var v9 float32 = (0.1 * 1.5)
        _ = fi8
        _ = v9
    }
    _ = v6
    _ = v7
    return -935
}
func genFn1(p0 int, p1 uint) uint {
    p0 += -2147483648
    switch ((("hello" + "a") + "a") + "\n\t") {
    case "a":
        if true {
            v10 := []int{}
            var v11 uint16 = 1
            if len(v10) >= 2 {
                v10 = v10[1:len(v10)]
            }
            _ = v10
            _ = v11
        }
    default:
        for fi12 := 0; fi12 < 20; fi12++ {
            var v13 int32 = 1
            v14 := GenS1{F0: p1, F1: 2147483647}
            _ = fi12
            _ = v13
            _ = v14
        }
    }
    for fi15 := 0; fi15 < 19; fi15++ {
        fmt.Println(p1)
        var v16 int = fi15
        _ = fi15
        _ = v16
    }
    _ = p0
    _ = p1
    return 996
}

func main() {
    v1 := []int{1, -78}
    var v2 int64 = 42
    v3 := []int{879, -248, -327, -9223372036854775808}
    var v4 uint16 = 2
    v5 := []int{}
    switch "foo bar" {
    case "foo bar":
        var v17 int32 = 1
        _ = v17
    default:
        var v18 uint64 = 42
        _ = v18
    }
    var v19 uint64 = 18446744073709551615
    v19 = v19
    _ = len(v5)
    if false {
        var v20 bool = ((!(!true)) && true)
        v2 += (v2 & (-(v2)))
        if (42 < 1000) {
            var v21 float64 = 0.1
            if ((2 < 4294967295) || (1000 <= -1)) {
                v19 = (v19 << ((255) & 63))
            } else {
                v22 := GenS1{F0: 0, F1: -1}
                v19 += 737
                v2 += (v2 << ((65535) & 63))
                _ = v22
            }
            _ = v21
        } else {
            for fi23 := 0; fi23 < 16; fi23++ {
                v24 := GenS2{F0: "hello", F1: v20}
                var v25 float64 = 0.1
                v26 := GenS1{F0: 65535, F1: 0}
                _ = fi23
                _ = v24
                _ = v25
                _ = v26
            }
            var v27 float32 = 0.1
            if v20 {
                v28 := GenS3{F0: (v4 >> ((2) & 63)), F1: 334, F2: -1}
                if len(v5) > 0 {
                    v5[0] = 258
                }
                _ = v28
            }
            _ = v27
        }
        _ = v20
    } else {
        var v29 uint64 = v19
        _ = v29
    }
    switch 358 {
    case 128:
        var v30 uint16 = 399
        _ = v30
    case -2147483648:
        var v31 int = 32768
        _ = v31
    }
    fmt.Println(v2, v19)
    if (539 != 128) {
        if true {
            var v32 uint64 = (v19 + v19)
            v33 := GenS2{F0: "hello", F1: (343 <= -2147483648)}
            _ = v32
            _ = v33
        } else {
            v2 = v2
            if ("" == "\n\t") {
                v34 := GenS2{F0: "a", F1: ((!(0 < 69)) || false)}
                for fi35 := 0; fi35 < 15; fi35++ {
                    v19 = v19
                    if (!true) {
                        v36 := GenS3{F0: (v4 & (v4 & 0)), F1: 128, F2: -974}
                        v2 += 9223372036854775807
                        _ = v36
                    }
                    fmt.Println(v19)
                    _ = fi35
                }
                var v37 string = ("a" + "hello")
                _ = v34
                _ = v37
            } else {
                var v38 int8 = -128
                var v39 float32 = 3.14159
                _ = v38
                _ = v39
            }
        }
        switch "hello" {
        case "":
            var v40 uint64 = 42
            _ = v40
        default:
            v5 = append(v5, 32767)
        }
        v41 := GenS1{F0: 533, F1: 694}
        _ = v41
    } else {
        for fi42 := 0; fi42 < 4; fi42++ {
            if (int32(v4) <= genFn0()) {
                if false {
                    fmt.Println(v19, v2)
                } else {
                    v2 -= v2
                    _ = len(v3)
                }
                var v43 int8 = 1
                _ = v43
            }
            _ = fi42
        }
        v44 := GenS2{F0: "hello", F1: (v4 >= 72)}
        _ = v44
    }
    v19 -= v19
    if (-18 <= -9223372036854775808) {
        var v45 uint8 = 2
        _ = v45
    }
    var v46 int32 = 1
    var v47 uint = 1
    var v48 uint64 = uint64(v4)
    fmt.Println(v2)
    v49 := GenS4{F0: ""}
    var v50 int16 = -41
    if true {
        v48 = (v19 >> ((genFn1(32768, v47)) & 63))
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("v48:", v48)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v47
    _ = v48
    _ = v49
    _ = v5
    _ = v50
}
