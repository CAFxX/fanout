package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 int8; F2 uint16 }
type GenS3 struct { F0 int64; F1 uint }

func main() {
    v1 := make(map[uint32]uint32)
    v2 := make(map[int]string)
    v3 := make(map[uint32]uint32)
    v4 := make(map[uint32]uint32)
    var v5 bool = (42 < -317)
    switch 127 {
    case -2147483648:
        v5 = ("foo bar" >= "a")
    case -1:
        for fi6 := 0; fi6 < 6; fi6++ {
            switch "" {
            case "a":
                for fi7 := 0; fi7 < 15; fi7++ {
                    switch (("a" + "foo bar") + "\n\t") {
                    case "\n\t":
                        if v5 {
                            v5 = (128 <= 128)
                            var v8 string = "a"
                            _ = v8
                        }
                        v9 := GenS1{F0: 127}
                        v10 := []int{4294967295, -1, fi6}
                        var v11 int64 = 9223372036854775807
                        var v12 int8 = 3
                        v11 += (v11 * -744)
                        _ = v10
                        _ = v11
                        _ = v12
                        _ = v9
                    default:
                        v13 := GenS1{F0: (fi7 ^ (fi6 / ((127) | 1)))}
                        _ = v13
                    }
                    _ = fi7
                }
            case "foo bar":
                var v14 int64 = -1
                _ = v14
            case "":
                for fi15 := 0; fi15 < 3; fi15++ {
                    var v16 float32 = 0.1
                    v5 = (!(v5 || v5))
                    v17 := GenS1{F0: (fi6 << ((65535) & 63))}
                    _ = fi15
                    _ = v16
                    _ = v17
                }
            }
            _ = fi6
        }
    }
    v5 = (851 == 4294967295)
    switch "foo bar" {
    case "\n\t":
        v18 := []int{}
        _ = v18
    }
    var v19 int32 = -392
    var v20 int32 = 290
    fmt.Println(v20)
    for fi21 := 0; fi21 < 17; fi21++ {
        v22 := []int{(fi21 | (fi21 * (fi21 >> ((4294967295) & 63)))), 9223372036854775807, fi21}
        if (476 < -32768) {
            v23 := GenS1{F0: fi21}
            for fi24 := 0; fi24 < 15; fi24++ {
                fmt.Println(v5)
                var v25 uint64 = 218
                _ = fi24
                _ = v25
            }
            if (v5 || v5) {
                switch ("hello" + "a") {
                case "a":
                    v26 := GenS1{F0: fi21}
                    _ = v26
                }
            } else {
                for fi27 := 0; fi27 < 9; fi27++ {
                    v19 = v20
                    _ = fi27
                }
                for fi28 := 0; fi28 < 8; fi28++ {
                    v5 = v5
                    v20 = 42
                    var v29 int8 = -128
                    if (32767 < 1) {
                        var v30 int = fi28
                        if (1e10 > 0.1) {
                            v19 = 713
                            var v31 float32 = 0.0
                            _ = v31
                        } else {
                            v32 := v3[250]
                            switch "foo bar" {
                            case "\n\t":
                                var v33 bool = (false && (true || ((128 > 2) || v5)))
                                _ = v33
                            case "hello":
                                v20 = v19
                            }
                            _ = v32
                        }
                        if false {
                            switch (v30 ^ 1) {
                            case -2147483648:
                                _ = len(v4)
                            }
                        }
                        _ = v30
                    } else {
                        var v34 int64 = -9223372036854775808
                        _ = v34
                    }
                    fmt.Println(fi21, fi21)
                    var v35 uint8 = 2
                    if len(v22) > 0 {
                        v22[0] = int(v19)
                    }
                    _ = fi28
                    _ = v29
                    _ = v35
                }
                v20 = v19
            }
            _ = v23
        }
        v5 = false
        _ = fi21
        _ = v22
    }
    if true {
        v20 += (-(v20))
    } else {
        switch ("foo bar" + ("a" + "hello")) {
        case "\n\t":
            var v36 int32 = -923
            _ = v36
        case "foo bar":
            fmt.Println(v5, v19, v19)
        }
    }
    v20 += (v19 / (((^(v20))) | 1))
    var v37 int64 = 42
    if v5 {
        if v5 {
            var v38 int = 32768
            if (42 >= 42) {
                v39 := GenS2{F0: 1, F1: -128, F2: 42}
                v19 = v20
                var v40 int = v38
                _ = v39
                _ = v40
            } else {
                for fi41 := 0; fi41 < 20; fi41++ {
                    switch (fi41 + (-(fi41))) {
                    case -129:
                        if v5 {
                            for fi42 := 0; fi42 < 8; fi42++ {
                                v43 := GenS3{F0: (v37 + v37), F1: 4294967295}
                                v37 = v37
                                _ = fi42
                                _ = v43
                            }
                        }
                    case 32767:
                        fmt.Println(v38)
                    }
                    _ = fi41
                }
                fmt.Println(v37, v38, v37)
                v44 := []int{v38, (v38 | (v38 >> ((2) & 63))), (v38 >> ((4294967295) & 63)), v38}
                _ = v44
            }
            _ = v38
        }
        var v45 float64 = 1.5
        var v46 float32 = float32(v45)
        _ = v45
        _ = v46
    }
    var v47 int32 = -1
    v19 = v47
    v48 := []int{-238, 1, -143}
    v48 = append(v48, 794)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v48:", len(v48), v48)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v37
    _ = v4
    _ = v47
    _ = v48
    _ = v5
}
