package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint64; F1 bool }
type GenS3 struct { F0 uint8; F1 int64 }
type GenS4 struct { F0 int32; F1 int32; F2 int8 }
func genFn0(p0 bool, p1 uint32, p2 uint8) int {
    if p0 {
        var v7 uint = 255
        p1 -= 1
        if (!p0) {
            if (1 <= (p2 * 0)) {
                p0 = false
            }
        } else {
            var v8 uint64 = 18446744073709551615
            _ = v8
        }
        _ = v7
    }
    p0 = ("a" == "foo bar")
    _ = p0
    _ = p1
    _ = p2
    return -2147483648
}

func main() {
    v1 := make(map[uint32]int)
    var v2 uint32 = 2147483648
    var v3 float64 = 1e10
    var v4 uint16 = 42
    v5 := make(map[uint32]bool)
    var v6 float64 = v3
    if (2147483647 != -308) {
        var v9 int = 1
        if false {
            v3 -= float64(-2.25)
            var v10 uint64 = 42
            var v11 uint32 = v2
            _ = v10
            _ = v11
        }
        for fi12 := 0; fi12 < 3; fi12++ {
            switch (fi12 / ((genFn0((!false), (^(v2)), 42)) | 1)) {
            case -2147483648:
                _ = len(v5)
            case 42:
                v13 := []int{}
                _ = v13
            case -468:
                for fi14 := 0; fi14 < 11; fi14++ {
                    v15 := GenS1{F0: v9}
                    fmt.Println(fi12, fi14)
                    fmt.Println(v6, v9)
                    if (2 != v4) {
                        v4 -= v4
                        v6 += 0.0
                    }
                    v16 := GenS2{F0: uint64(v4), F1: false}
                    _ = fi14
                    _ = v15
                    _ = v16
                }
            }
            if false {
                if false {
                    v17 := []int{genFn0(false, v2, 59), v9, (v9 + (fi12 & fi12)), int(v4)}
                    for fi18 := 0; fi18 < 18; fi18++ {
                        if (true || (0 <= -761)) {
                            var v19 int64 = 42
                            _ = v19
                        }
                        _ = fi18
                    }
                    _ = v17
                }
                switch "" {
                case "":
                    v20 := []int{genFn0(true, (v2 ^ v2), 2), 0}
                    _ = v20
                case "\n\t":
                    v21 := GenS3{F0: 2, F1: 9223372036854775807}
                    _ = v21
                default:
                    var v22 int = v9
                    _ = v22
                }
            } else {
                v6 = v3
                var v23 bool = (!false)
                v5[(v2 | v2)] = true
                _ = v23
            }
            _ = fi12
        }
        _ = v9
    }
    var v24 int8 = 127
    if (!false) {
        delete(v1, v2)
        _ = len(v1)
    } else {
        v4 = (v4 + 32767)
        var v25 string = ""
        if false {
            fmt.Println(v24, v4, v4)
            for fi26 := 0; fi26 < 10; fi26++ {
                var v27 float32 = -2.25
                _ = fi26
                _ = v27
            }
            v2 = (v2 << ((586) & 63))
        } else {
            v28 := GenS4{F0: 0, F1: 201, F2: v24}
            v24 -= 100
            v29 := []int{0, 0}
            _ = v28
            _ = v29
        }
        _ = v25
    }
    v3 = v6
    var v30 uint = 4294967295
    v31 := GenS4{F0: 1, F1: 0, F2: -128}
    fmt.Println(v24, v6)
    var v32 int32 = 2147483647
    if false {
        var v33 int = 128
        if (v32 <= v32) {
            switch ("hello" + ("hello" + (((("\n\t" + "hello") + "hello") + "a") + ("" + "hello")))) {
            case "\n\t":
                for fi34 := 0; fi34 < 4; fi34++ {
                    v35 := v5[v2]
                    fmt.Println(v33)
                    _ = fi34
                    _ = v35
                }
            case "hello":
                var v36 float64 = v6
                _ = v36
            case "a":
                v37 := GenS4{F0: v32, F1: (-(v32)), F2: 65}
                _ = v37
            }
            v30 -= (v30 * 773)
            switch "a" {
            case "\n\t":
                v3 = (((1e10 + ((((0.0 - v3) + v6) + 1e10) * 3.14159)) / 0.1) * 3.14159)
            }
        } else {
            v33 = genFn0((int32(v2) <= 42), (^(v2)), 128)
            var v38 int8 = 1
            _ = v38
        }
        v39 := []int{(-(v33)), (v33 & genFn0(false, (^(v2)), 2)), int(v2), int(v4)}
        _ = v33
        _ = v39
    }
    for fi40 := 0; fi40 < 1; fi40++ {
        v24 -= v24
        _ = fi40
    }
    fmt.Println(v6, v30, v4)
    v4 = v4
    v41 := []int{32767, -2147483648, 127, 32767}
    for fi42 := 0; fi42 < 2; fi42++ {
        switch "hello" {
        case "foo bar":
            var v43 int = (^(fi42))
            _ = v43
        }
        for fi44 := 0; fi44 < 16; fi44++ {
            var v45 bool = ((v32 + v32) != v32)
            fmt.Println(v24, v32, v30)
            _ = fi44
            _ = v45
        }
        v30 = uint(v2)
        _ = fi42
    }
    if (v30 > (v30 - v30)) {
        if false {
            var v46 string = ("" + "foo bar")
            _ = v46
        } else {
            v47 := GenS3{F0: 2, F1: 0}
            _ = v47
        }
        fmt.Println(v30, v6)
        v48 := GenS4{F0: v32, F1: (^(v32)), F2: (v24 << ((uint(v4)) & 63))}
        _ = v48
    } else {
        var v49 string = (("hello" + ("foo bar" + ("\n\t" + ""))) + ("" + "a"))
        _ = v49
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v41:", len(v41), v41)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v41
    _ = v5
    _ = v6
}
