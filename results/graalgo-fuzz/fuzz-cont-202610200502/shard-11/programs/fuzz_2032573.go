package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8 }
type GenS2 struct { F0 uint32; F1 bool }
type GenS3 struct { F0 int8; F1 bool; F2 int16 }
type GenS4 struct { F0 uint16; F1 string }
type GenS5 struct { F0 string; F1 uint; F2 int32 }
type GenS6 struct { F0 int8; F1 uint }
type GenS7 struct { F0 uint16; F1 uint64 }
type GenS8 struct { F0 uint8 }
func genFn0(p0 int64) int32 {
    if true {
        v5 := GenS1{F0: (!true), F1: 42}
        if true {
            v6 := GenS2{F0: 4294967295, F1: (-700 != 42)}
            v7 := []int{0, -453}
            _ = v6
            _ = v7
        }
        _ = v5
    }
    if true {
        var v8 int32 = 0
        v8 = 42
        p0 = 42
        _ = v8
    } else {
        v9 := []int{32768, 127, 816, 9223372036854775807}
        p0 -= 42
        fmt.Println(p0)
        _ = v9
    }
    _ = p0
    return -1
}
func genFn1(p0 uint16, p1 int, p2 bool) string {
    if p2 {
        v10 := GenS3{F0: 1, F1: (1 != -2147483648), F2: 32767}
        _ = v10
    } else {
        v11 := GenS4{F0: p0, F1: "foo bar"}
        _ = v11
    }
    if (-1 > 1000) {
        p2 = p2
    }
    _ = p0
    _ = p1
    _ = p2
    return ("hello" + ("" + ""))
}

func main() {
    var v1 uint16 = 2
    v2 := make(map[int]bool)
    v3 := make(map[string]string)
    var v4 uint16 = 0
    if (42 < 1) {
        v4 = (v4 & (^(v1)))
        if true {
            var v12 uint8 = 53
            v12 = 2
            _ = v12
        }
        var v13 uint8 = 42
        _ = v13
    }
    var v14 uint16 = v4
    var v15 int64 = -1
    switch "hello" {
    case "foo bar":
        v16 := GenS5{F0: "", F1: 2, F2: -2147483648}
        _ = v16
    case "\n\t":
        for fi17 := 0; fi17 < 18; fi17++ {
            if ("\n\t" <= "a") {
                v18 := GenS6{F0: int8(v4), F1: 2}
                _ = v18
            }
            switch fi17 {
            case 32767:
                v4 += (v14 * v14)
            case -1:
                _ = len(v2)
            }
            _ = fi17
        }
    }
    v19 := []int{32768}
    if true {
        switch -952 {
        case 0:
            v1 -= 42
        case 128:
            fmt.Println(v4)
        case -703:
            if false {
                v20 := GenS7{F0: v4, F1: 997}
                var v21 int32 = -1
                _ = v20
                _ = v21
            }
        }
        if (422 == (^(v15))) {
            v22 := GenS8{F0: 2}
            _ = v22
        }
    }
    for fi23 := 0; fi23 < 5; fi23++ {
        var v24 uint16 = 1
        switch (-(fi23)) {
        case 343:
            for fi25 := 0; fi25 < 2; fi25++ {
                if false {
                    var v26 float32 = -2.25
                    fmt.Println(v1, v15, v4)
                    _ = v26
                }
                var v27 int32 = genFn0((v15 >> ((2) & 63)))
                if len(v19) > 0 {
                    v19[0] = (fi23 ^ (fi25 >> ((18446744073709551615) & 63)))
                }
                var v28 uint16 = (v14 >> ((255) & 63))
                _ = fi25
                _ = v27
                _ = v28
            }
        case 359:
            var v29 int = (fi23 | 4294967295)
            _ = v29
        case -9223372036854775808:
            var v30 uint16 = (v24 * (v24 * v14))
            _ = v30
        }
        v14 -= 665
        _ = fi23
        _ = v24
    }
    for fi31 := 0; fi31 < 14; fi31++ {
        fmt.Println(v14, v14, v15)
        var v32 float32 = 1e10
        _ = fi31
        _ = v32
    }
    v33 := GenS5{F0: "foo bar", F1: 42, F2: -2147483648}
    var v34 int16 = -32768
    for fi35 := 0; fi35 < 13; fi35++ {
        v36 := GenS3{F0: -128, F1: true, F2: v34}
        _ = fi35
        _ = v36
    }
    var v37 int16 = v34
    var v38 int64 = 9223372036854775807
    if (0 >= 127) {
        fmt.Println(v38, v15, v37)
        fmt.Println(v38)
    } else {
        v39 := v2[-9223372036854775808]
        if len(v19) > 0 {
            v19[0] = -726
        }
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v37
    _ = v38
    _ = v4
}
