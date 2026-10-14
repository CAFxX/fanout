package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 int8; F2 uint64 }
type GenS3 struct { F0 uint; F1 int; F2 uint32 }
type GenS4 struct { F0 uint32; F1 int; F2 uint32 }
func genFn0(p0 int16, p1 string) uint64 {
    var v5 int32 = -2147483648
    if (-1 < 96) {
        var v6 float64 = 1.5
        for fi7 := 0; fi7 < 17; fi7++ {
            fmt.Println(fi7, v6, v6)
            var v8 int64 = -142
            _ = fi7
            _ = v8
        }
        _ = v6
    }
    _ = p0
    _ = p1
    _ = v5
    return 18446744073709551615
}

func main() {
    v1 := make(map[uint32]int)
    var v2 float32 = 3.14159
    v3 := []uint32{42, 660, 65535}
    v4 := make(map[int]uint32)
    if len(v3) > 0 {
        v3[0] = 2
    }
    v9 := GenS1{F0: -978}
    v10 := GenS2{F0: 100, F1: 42, F2: 1}
    fmt.Println(v2)
    v11 := GenS2{F0: -1, F1: 1, F2: 4294967295}
    v12 := v1[2147483647]
    switch ("foo bar" + "") {
    case "hello":
        fmt.Println(v12, v12)
    case "foo bar":
        v2 += (v2 / v2)
    default:
        if ((9223372036854775807 >= 9223372036854775807) || (877 >= 32767)) {
            v13 := v1[4294967295]
            v2 = (-2.25 / ((v2 * v2) + v2))
            _ = v13
        } else {
            fmt.Println(v12)
            var v14 bool = (!(42 == 42))
            var v15 bool = v14
            _ = v14
            _ = v15
        }
    }
    var v16 uint64 = 42
    if false {
        var v17 uint64 = 4294967295
        _ = v17
    }
    v2 -= 3.14159
    for fi18 := 0; fi18 < 10; fi18++ {
        if true {
            var v19 uint = 0
            for fi20 := 0; fi20 < 17; fi20++ {
                fmt.Println(fi18)
                v19 += (v19 << ((v19) & 63))
                _ = fi20
            }
            _ = v19
        } else {
            for fi21 := 0; fi21 < 15; fi21++ {
                v3 = append(v3, 2147483647)
                _ = len(v4)
                _ = fi21
            }
            v12 = (v12 % ((42) | 1))
            fmt.Println(fi18, fi18, fi18)
        }
        for fi22 := 0; fi22 < 11; fi22++ {
            switch (-(fi22)) {
            case 42:
                v16 += genFn0(0, ("\n\t" + ""))
            case 32767:
                v12 = fi22
            default:
                _ = len(v3)
            }
            if len(v3) > 0 {
                v3[0] = 2147483647
            }
            _ = fi22
        }
        var v23 int32 = 42
        _ = fi18
        _ = v23
    }
    fmt.Println(v2, v2, v16)
    var v24 int = v12
    if (false || false) {
        v2 = v2
        if (100 == 100) {
            if ((0.0 != 3.14159) || true) {
                if (genFn0(-1, ("hello" + ("hello" + (("" + "\n\t") + "")))) != genFn0(0, "a")) {
                    if ((v16 << ((476) & 63)) <= 775) {
                        var v25 uint64 = v16
                        var v26 int = 42
                        _ = v25
                        _ = v26
                    } else {
                        var v27 bool = (float32((3.14159 + 0.1)) != -2.25)
                        v28 := GenS3{F0: 2, F1: v24, F2: 300}
                        _ = v27
                        _ = v28
                    }
                }
                var v29 uint32 = 2147483647
                v30 := GenS1{F0: (v12 ^ 2147483647)}
                _ = v29
                _ = v30
            }
            v2 = v2
        } else {
            var v31 int = v12
            _ = v31
        }
        v4[v12] = 301
    } else {
        v4[v24] = 579
    }
    if false {
        for fi32 := 0; fi32 < 8; fi32++ {
            v24 += (v24 + (^(v12)))
            v2 = v2
            _ = fi32
        }
        for fi33 := 0; fi33 < 19; fi33++ {
            var v34 uint32 = 2147483647
            v35 := GenS4{F0: v34, F1: v24, F2: v34}
            _ = fi33
            _ = v34
            _ = v35
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v9
}
