package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16 }
type GenS2 struct { F0 bool; F1 int8 }
type GenS3 struct { F0 uint32; F1 int8; F2 int64 }
type GenS4 struct { F0 int64; F1 uint; F2 int64 }
func genFn0() uint16 {
    var v5 int64 = -1
    v5 -= (v5 << ((42) & 63))
    var v6 int = 0
    if (345 <= 1) {
        v5 -= (v5 - (v5 & (^(v5))))
        switch (v6 / ((-284) | 1)) {
        case 32767:
            fmt.Println(v6, v5)
        case -9223372036854775808:
            if true {
                v7 := GenS1{F0: 36, F1: 1000}
                var v8 uint64 = 4294967295
                v9 := []int{42}
                _ = v7
                _ = v8
                _ = v9
            }
        }
        var v10 int = (v6 - (v6 << ((336) & 63)))
        _ = v10
    }
    _ = v5
    _ = v6
    return 1
}
func genFn1(p0 uint, p1 uint16) uint8 {
    switch "foo bar" {
    case "":
        var v11 bool = (1 != 1)
        _ = v11
    case "foo bar":
        var v12 int32 = -1
        _ = v12
    }
    var v13 uint8 = 72
    if false {
        var v14 int = 156
        p1 = (p1 | genFn0())
        _ = v14
    } else {
        p1 -= (^(p1))
    }
    for fi15 := 0; fi15 < 13; fi15++ {
        if (4294967295 > 4294967295) {
            var v16 int64 = -7
            _ = v16
        }
        p0 -= (p0 >> ((p0) & 63))
        v17 := []int{(fi15 + fi15)}
        _ = fi15
        _ = v17
    }
    if true {
        var v18 int8 = 127
        p1 += p1
        v19 := GenS2{F0: false, F1: v18}
        _ = v18
        _ = v19
    } else {
        v20 := GenS3{F0: 54, F1: 1, F2: 1}
        _ = v20
    }
    _ = p0
    _ = p1
    _ = v13
    return v13
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 string = ("a" + "foo bar")
    v3 := []uint32{4294967295}
    var v4 float32 = 3.14159
    v2 = "foo bar"
    v2 = v2
    v4 = (v4 - v4)
    if len(v3) > 0 {
        v3[0] = 65535
    }
    switch 2147483647 {
    case 32768:
        fmt.Println(v4, v4)
    case 439:
        switch v2 {
        case "hello":
            var v21 int32 = 0
            _ = v21
        default:
            var v22 int = 750
            _ = v22
        }
    }
    fmt.Println(v2, v4)
    var v23 bool = false
    if (-281 < -9223372036854775808) {
        if (0 == 647) {
            v24 := GenS2{F0: false, F1: -1}
            _ = v24
        }
    } else {
        switch "a" {
        case "":
            for fi25 := 0; fi25 < 20; fi25++ {
                v26 := v1[2147483648]
                if (!(!(v26 && v23))) {
                    var v27 string = v2
                    _ = v27
                }
                _ = fi25
                _ = v26
            }
        default:
            for fi28 := 0; fi28 < 8; fi28++ {
                if false {
                    var v29 int8 = -1
                    _ = v29
                } else {
                    var v30 int16 = -667
                    var v31 string = v2
                    var v32 int32 = -450
                    _ = v30
                    _ = v31
                    _ = v32
                }
                _ = fi28
            }
        }
    }
    switch v2 {
    case "foo bar":
        v33 := GenS4{F0: 1, F1: 4294967295, F2: 42}
        _ = v33
    }
    _ = len(v3)
    v2 = (("\n\t" + ("hello" + v2)) + "a")
    for fi34 := 0; fi34 < 7; fi34++ {
        v35 := GenS4{F0: 357, F1: 18446744073709551615, F2: 212}
        var v36 int16 = 0
        _ = fi34
        _ = v35
        _ = v36
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v4
}
