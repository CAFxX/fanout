package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 int8 }
type GenS2 struct { F0 string; F1 int64; F2 uint16 }
type GenS3 struct { F0 bool; F1 uint8; F2 int8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint64; F1 string }
func genFn0(p0 uint, p1 uint32, p2 uint32) int8 {
    fmt.Println(p1)
    v6 := []int{}
    switch "\n\t" {
    case "hello":
        fmt.Println(p2, p1, p0)
    case "foo bar":
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    }
    var v7 uint64 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 42
}
func genFn1() string {
    for fi8 := 0; fi8 < 2; fi8++ {
        if true {
            var v9 bool = true
            _ = v9
        }
        _ = fi8
    }
    if true {
        var v10 int64 = -1
        v11 := []int{9223372036854775807}
        var v12 bool = (1.5 == ((((float64((1.5 - 1.5)) / 1.0) + 0.1) / float64((1.5 - float32((0.0 + 1e10))))) / -2.25))
        _ = v10
        _ = v11
        _ = v12
    } else {
        var v13 float64 = 0.1
        _ = v13
    }
    var v14 string = "hello"
    v14 = "hello"
    _ = v14
    return (v14 + (("\n\t" + "foo bar") + "\n\t"))
}

func main() {
    var v1 int8 = 100
    var v2 float32 = ((float32(1e10) * (-2.25 / -2.25)) / (0.0 - 3.14159))
    v3 := []string{("\n\t" + ("foo bar" + ((("foo bar" + (("hello" + "a") + "foo bar")) + "") + "foo bar"))), "foo bar"}
    v4 := make(map[uint32]uint32)
    v5 := make(map[string]uint32)
    fmt.Println(v1, v2)
    v15 := GenS1{F0: 0, F1: 980, F2: genFn0(523, 343, 42)}
    fmt.Println(v2, v1)
    v16 := GenS2{F0: "\n\t", F1: 505, F2: 2}
    var v17 uint8 = 79
    v18 := GenS2{F0: "", F1: 9223372036854775807, F2: 2}
    if false {
        var v19 uint16 = 0
        var v20 int16 = 32767
        var v21 string = ("" + "foo bar")
        _ = v19
        _ = v20
        _ = v21
    }
    v17 -= (v17 * 255)
    for fi22 := 0; fi22 < 5; fi22++ {
        for fi23 := 0; fi23 < 6; fi23++ {
            fmt.Println(v2, v2, fi23)
            if (210 == uint16(v17)) {
                var v24 uint32 = 2147483648
                switch (("" + "hello") + "\n\t") {
                case "\n\t":
                    if ((v1 % (((v1 >> ((1) & 63))) | 1)) == (v1 * (v1 | v1))) {
                        v2 = v2
                        var v25 int = fi23
                        var v26 int8 = (v1 * int8(fi23))
                        _ = len(v4)
                        _ = v25
                        _ = v26
                    } else {
                        var v27 float32 = -2.25
                        v1 -= genFn0(677, v24, 4294967295)
                        var v28 int16 = -32768
                        _ = v27
                        _ = v28
                    }
                case "foo bar":
                    if false {
                        var v29 uint32 = v24
                        if true {
                            delete(v4, (v29 + v29))
                            _ = len(v5)
                        }
                        var v30 float32 = v2
                        _ = v29
                        _ = v30
                    }
                    var v31 int = fi23
                    if false {
                        _ = len(v3)
                        var v32 uint8 = 127
                        delete(v5, "\n\t")
                        v4[0] = v24
                        var v33 int32 = 42
                        _ = v32
                        _ = v33
                    }
                    _ = v31
                }
                fmt.Println(v17, v17, fi23)
                v17 += (v17 << ((18446744073709551615) & 63))
                v34 := GenS3{F0: false, F1: (v17 + 128), F2: genFn0(4294967295, (v24 + (v24 >> ((226) & 63))), 42)}
                _ = v24
                _ = v34
            } else {
                fmt.Println(v1)
                if (0.0 >= 1.5) {
                    delete(v4, 86)
                } else {
                    if ((549 <= 1) || (true && true)) {
                        var v35 uint = 4
                        var v36 uint = 4294967295
                        _ = v35
                        _ = v36
                    }
                }
            }
            _ = fi23
        }
        var v37 int16 = 42
        if true {
            var v38 int = fi22
            v39 := GenS4{F0: 2147483648}
            _ = v38
            _ = v39
        }
        _ = fi22
        _ = v37
    }
    v17 += v17
    switch 42 {
    case -9223372036854775808:
        var v40 int8 = v1
        _ = v40
    default:
        var v41 uint16 = 0
        _ = v41
    }
    if (-1 >= 9223372036854775807) {
        if false {
            if (!true) {
                v1 += (v1 % ((1) | 1))
                var v42 uint16 = 215
                _ = v42
            } else {
                for fi43 := 0; fi43 < 11; fi43++ {
                    var v44 int64 = 0
                    var v45 int8 = (v1 * (v1 % ((-1) | 1)))
                    _ = fi43
                    _ = v44
                    _ = v45
                }
                for fi46 := 0; fi46 < 8; fi46++ {
                    fmt.Println(v1)
                    var v47 float64 = 1.5
                    _ = fi46
                    _ = v47
                }
                v3 = append(v3, ("" + "\n\t"))
            }
            var v48 int = -350
            _ = v48
        }
    }
    v49 := GenS5{F0: 1, F1: ("\n\t" + "foo bar")}
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
}
