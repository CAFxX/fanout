package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint64; F2 uint8 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint32; F1 uint8 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int64 }
func genFn0(p0 int) uint32 {
    switch 0 {
    case 127:
        var v5 bool = true
        _ = v5
    case -571:
        for fi6 := 0; fi6 < 3; fi6++ {
            var v7 int16 = 1
            v8 := GenS1{F0: 446, F1: 848, F2: 2}
            _ = fi6
            _ = v7
            _ = v8
        }
    }
    v9 := []int{p0}
    _ = len(v9)
    _ = p0
    _ = v9
    return 2147483648
}
func genFn1(p0 uint8, p1 int32, p2 int) int8 {
    fmt.Println(p2, p2)
    v10 := GenS1{F0: p1, F1: 2, F2: 58}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    return -67
}

func main() {
    var v1 int16 = -1
    v2 := make(map[uint32]string)
    v3 := make(map[int]bool)
    v4 := []string{"", "a", ("foo bar" + "a"), ""}
    v1 += -1
    if (true || false) {
        if true {
            if true {
                v4 = append(v4, ("a" + "hello"))
            }
            v11 := v2[863]
            v4 = append(v4, ((("hello" + ("" + (v11 + "foo bar"))) + (v11 + "")) + v11))
            _ = v11
        } else {
            if false {
                fmt.Println(v1)
                var v12 int64 = 165
                _ = v12
            }
            if false {
                v4 = append(v4, "a")
                var v13 uint32 = 2
                switch ("\n\t" + ("" + "a")) {
                case "":
                    v13 = (v13 / ((v13) | 1))
                case "foo bar":
                    for fi14 := 0; fi14 < 4; fi14++ {
                        v13 -= (v13 >> ((284) & 63))
                        _ = fi14
                    }
                case "hello":
                    v1 -= v1
                }
                _ = v13
            } else {
                v1 = (v1 & -747)
            }
        }
        v1 = (v1 >> ((234) & 63))
        var v15 int32 = -2147483648
        _ = v15
    }
    switch -2147483648 {
    case 32767:
        if false {
            var v16 int = 4294967295
            _ = v16
        }
    }
    var v17 float32 = (3.14159 * 0.0)
    var v18 int64 = 9223372036854775807
    v1 = -472
    var v19 float64 = 0.0
    if (true || false) {
        var v20 float32 = 0.1
        _ = v20
    } else {
        v1 = 1000
        v21 := v3[2147483647]
        var v22 int = -1
        _ = v21
        _ = v22
    }
    if (!false) {
        for fi23 := 0; fi23 < 17; fi23++ {
            v24 := GenS2{F0: 340}
            _ = fi23
            _ = v24
        }
        fmt.Println(v19, v19)
        var v25 uint64 = 4294967295
        _ = v25
    }
    if (false || false) {
        for fi26 := 0; fi26 < 14; fi26++ {
            switch fi26 {
            case 192:
                if (0 == 363) {
                    var v27 int64 = v18
                    v4 = append(v4, "foo bar")
                    _ = v27
                }
            case -757:
                fmt.Println(v17, v19)
            }
            _ = fi26
        }
        v28 := GenS3{F0: true}
        fmt.Println(v18, v19)
        _ = v28
    }
    v1 -= (-(v1))
    if false {
        fmt.Println(v1, v1)
        var v29 int16 = v1
        _ = v29
    } else {
        v30 := GenS4{F0: 193, F1: 160}
        v31 := GenS5{F0: "foo bar"}
        _ = v30
        _ = v31
    }
    for fi32 := 0; fi32 < 1; fi32++ {
        v33 := GenS4{F0: 79, F1: 70}
        var v34 float32 = ((float32(((v19 * v19) / 1.5)) / v17) - 1e10)
        _ = fi32
        _ = v33
        _ = v34
    }
    v1 = (v1 * (v1 * -1))
    for fi35 := 0; fi35 < 2; fi35++ {
        if true {
            fmt.Println(v17, v1)
            v18 = v18
            v36 := GenS6{F0: (v18 - 1)}
            _ = v36
        } else {
            switch "a" {
            case "a":
                v19 += v19
            case "":
                switch (-(fi35)) {
                case -1:
                    if true {
                        var v37 int = (fi35 >> ((42) & 63))
                        _ = v37
                    }
                case 32768:
                    var v38 bool = true
                    _ = v38
                }
            }
            var v39 int8 = 77
            _ = v39
        }
        _ = fi35
    }
    var v40 uint32 = genFn0(845)
    v18 -= (v18 - v18)
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v40
}
