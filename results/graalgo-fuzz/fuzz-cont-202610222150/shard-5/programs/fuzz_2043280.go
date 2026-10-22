package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint; F1 uint32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int; F1 uint16; F2 string }
func genFn0(p0 uint16, p1 int32) int {
    p1 += (p1 & p1)
    p0 = (^(p0))
    _ = p0
    _ = p1
    return 42
}

func main() {
    v1 := []uint32{2, 42, 958, 1}
    v2 := make(map[string]uint32)
    v3 := make(map[int]int)
    var v4 int = 128
    var v5 uint8 = 42
    v6 := []string{("a" + "\n\t"), "a", "a"}
    _ = len(v3)
    v5 -= (v5 % (((v5 << ((258) & 63))) | 1))
    for fi7 := 0; fi7 < 4; fi7++ {
        var v8 int32 = 0
        switch (-(v4)) {
        case 4294967295:
            v9 := GenS1{F0: 193}
            _ = v9
        case 1:
            var v10 uint16 = 255
            _ = v10
        }
        v11 := GenS2{F0: 242, F1: 2}
        _ = fi7
        _ = v11
        _ = v8
    }
    v12 := GenS2{F0: 1, F1: 4294967295}
    if true {
        v5 = (^(v5))
    } else {
        if (-239 != 32767) {
            var v13 float32 = (3.14159 + (-2.25 + 0.1))
            fmt.Println(v4, v5)
            _ = v13
        } else {
            var v14 uint = 0
            var v15 int8 = -128
            fmt.Println(v5)
            _ = v14
            _ = v15
        }
        var v16 float32 = (0.0 + 0.0)
        _ = v16
    }
    for fi17 := 0; fi17 < 14; fi17++ {
        switch 32768 {
        case 9223372036854775807:
            var v18 uint32 = 42
            _ = v18
        }
        v19 := GenS2{F0: 255, F1: 0}
        if true {
            var v20 int = (v4 ^ fi17)
            _ = v20
        }
        _ = fi17
        _ = v19
    }
    v4 = v4
    for fi21 := 0; fi21 < 19; fi21++ {
        for fi22 := 0; fi22 < 3; fi22++ {
            if true {
                v5 -= v5
                if ((("\n\t" + "") < "hello") || (false && ("hello" >= ("hello" + "")))) {
                    v23 := GenS3{F0: 65535}
                    _ = v23
                } else {
                    v4 += fi22
                    v4 = genFn0(1, 615)
                    v4 -= (fi21 & (fi22 / (((fi22 - (fi22 * 307))) | 1)))
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            } else {
                if true {
                    if (!false) {
                        v24 := GenS1{F0: -254}
                        v6 = append(v6, "hello")
                        var v25 uint8 = v5
                        _ = v24
                        _ = v25
                    }
                    v6 = append(v6, "hello")
                }
            }
            if false {
                switch 0 {
                case -9223372036854775808:
                    v26 := GenS3{F0: 42}
                    _ = v26
                case 1:
                    fmt.Println(v5, fi21, fi21)
                }
                var v27 int64 = 0
                _ = v27
            }
            _ = fi22
        }
        switch 1 {
        case -736:
            for fi28 := 0; fi28 < 17; fi28++ {
                v29 := GenS4{F0: v4, F1: 42, F2: ("hello" + "\n\t")}
                var v30 int = (fi28 >> ((795) & 63))
                _ = fi28
                _ = v29
                _ = v30
            }
        case 0:
            var v31 float64 = (-2.25 - float64((0.0 - (0.1 - -2.25))))
            _ = v31
        case 127:
            var v32 uint64 = 222
            _ = v32
        default:
            var v33 uint = 448
            _ = v33
        }
        var v34 string = ""
        _ = fi21
        _ = v34
    }
    switch ("\n\t" + "") {
    case "foo bar":
        var v35 int32 = -1
        _ = v35
    case "hello":
        _ = len(v2)
    }
    for fi36 := 0; fi36 < 1; fi36++ {
        if (0 == 1) {
            v37 := GenS2{F0: 958, F1: 2147483647}
            _ = v37
        }
        if (false || (458 != 411)) {
            fmt.Println(v4, fi36)
        }
        for fi38 := 0; fi38 < 17; fi38++ {
            if (3.14159 >= -2.25) {
                switch v4 {
                case 231:
                    switch v4 {
                    case 0:
                        v5 += (v5 * (v5 / ((199) | 1)))
                    }
                case 877:
                    fmt.Println(v4, fi36)
                default:
                    var v39 uint8 = v5
                    _ = v39
                }
            } else {
                var v40 string = "foo bar"
                _ = v40
            }
            _ = fi38
        }
        _ = fi36
    }
    v4 += v4
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
