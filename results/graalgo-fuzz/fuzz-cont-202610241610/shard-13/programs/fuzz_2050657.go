package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 uint64; F1 int32; F2 int }
type GenS3 struct { F0 uint16; F1 int16 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint8; F1 int }
type GenS6 struct { F0 uint }
func genFn0() uint16 {
    if true {
        v7 := GenS1{F0: 2147483647, F1: 872}
        if true {
            var v8 uint16 = 255
            v9 := GenS1{F0: 1, F1: v8}
            _ = v8
            _ = v9
        } else {
            if false {
                for fi10 := 0; fi10 < 8; fi10++ {
                    v11 := []int{(fi10 + fi10)}
                    v12 := GenS1{F0: -586, F1: 42}
                    if true {
                        v13 := GenS1{F0: 2147483647, F1: 42}
                        for fi14 := 0; fi14 < 2; fi14++ {
                            fmt.Println(fi10)
                            var v15 float64 = float64(1e10)
                            v15 = v15
                            _ = fi14
                            _ = v15
                        }
                        var v16 float32 = (((1.5 - 0.1) - 1.5) * 0.1)
                        _ = v13
                        _ = v16
                    } else {
                        var v17 int16 = 1
                        _ = v17
                    }
                    _ = fi10
                    _ = v11
                    _ = v12
                }
            }
        }
        _ = v7
    } else {
        var v18 float64 = 1e10
        v19 := GenS1{F0: 0, F1: 255}
        _ = v18
        _ = v19
    }
    var v20 int16 = 936
    for fi21 := 0; fi21 < 17; fi21++ {
        fmt.Println(v20)
        if false {
            v22 := []int{fi21, 861, fi21, (fi21 & -469)}
            _ = v22
        }
        _ = fi21
    }
    for fi23 := 0; fi23 < 5; fi23++ {
        if (true || false) {
            v24 := GenS1{F0: -2147483648, F1: 468}
            switch fi23 {
            case 42:
                if (false && false) {
                    fmt.Println(v20, v20)
                    v20 -= (v20 % ((v20) | 1))
                    var v25 float32 = 0.0
                    _ = v25
                }
            case -2147483648:
                for fi26 := 0; fi26 < 3; fi26++ {
                    v27 := []int{fi26}
                    if false {
                        var v28 uint32 = 0
                        if len(v27) > 0 {
                            v27[0] = (fi23 | fi26)
                        }
                        v29 := GenS2{F0: 774, F1: -2147483648, F2: fi23}
                        _ = v28
                        _ = v29
                    }
                    v20 -= (v20 >> ((28) & 63))
                    v30 := GenS3{F0: 1, F1: (v20 / ((v20) | 1))}
                    _ = fi26
                    _ = v27
                    _ = v30
                }
            case 0:
                if true {
                    var v31 uint = 1
                    _ = v31
                } else {
                    if false {
                        v20 = v20
                    }
                    switch (("" + "a") + ("hello" + ("\n\t" + "foo bar"))) {
                    case "\n\t":
                        switch fi23 {
                        case -426:
                            v32 := GenS4{F0: 255}
                            _ = v32
                        }
                    default:
                        v33 := []int{fi23, fi23, (^(fi23)), fi23}
                        _ = v33
                    }
                }
            }
            for fi34 := 0; fi34 < 5; fi34++ {
                switch "" {
                case "\n\t":
                    v20 += v20
                case "hello":
                    var v35 int8 = -74
                    _ = v35
                case "foo bar":
                    if false {
                        v20 = int16(fi23)
                        v36 := GenS2{F0: 387, F1: 42, F2: 4294967295}
                        if (18446744073709551615 > 65535) {
                            v20 = (v20 + -74)
                            v37 := []int{}
                            if len(v37) >= 2 {
                                v37 = v37[1:len(v37)]
                            }
                            _ = v37
                        }
                        _ = v36
                    }
                default:
                    v38 := []int{fi34, -807}
                    _ = v38
                }
                var v39 int8 = 127
                var v40 uint16 = 699
                v41 := []int{}
                v42 := GenS3{F0: v40, F1: v20}
                _ = fi34
                _ = v39
                _ = v40
                _ = v41
                _ = v42
            }
            _ = v24
        } else {
            var v43 int32 = 618
            _ = v43
        }
        v20 += v20
        _ = fi23
    }
    _ = v20
    return 300
}
func genFn1(p0 uint32, p1 int16, p2 int) int32 {
    p2 = p2
    p1 = -32768
    switch p2 {
    case 621:
        p2 = (p2 % (((p2 | (p2 >> ((2) & 63)))) | 1))
    case 2147483647:
        var v44 float64 = float64(0.1)
        _ = v44
    case -548:
        fmt.Println(p0, p1, p0)
    }
    p1 += p1
    _ = p0
    _ = p1
    _ = p2
    return -2147483648
}

func main() {
    var v1 bool = true
    v2 := make(map[int]string)
    v3 := make(map[string]uint32)
    v4 := []uint32{}
    var v5 int64 = 42
    var v6 string = "hello"
    for fi45 := 0; fi45 < 2; fi45++ {
        if (!v1) {
            switch (("foo bar" + (v6 + v6)) + v6) {
            case "":
                if (0.1 <= 1.5) {
                    v46 := GenS5{F0: 127, F1: fi45}
                    v6 = ""
                    _ = len(v3)
                    _ = v46
                }
            case "a":
                delete(v3, v6)
            case "\n\t":
                switch v6 {
                case "foo bar":
                    v5 += v5
                }
            }
            switch (fi45 >> ((4294967295) & 63)) {
            case 127:
                v47 := GenS2{F0: 42, F1: 1, F2: fi45}
                _ = v47
            case 32767:
                var v48 uint16 = 2
                _ = v48
            case 32768:
                v49 := GenS6{F0: 1}
                _ = v49
            }
        } else {
            v6 = v6
        }
        if (!v1) {
            v5 += v5
            fmt.Println(v6, fi45)
            v50 := v2[819]
            _ = v50
        }
        switch "a" {
        case "foo bar":
            var v51 uint = 84
            _ = v51
        }
        _ = fi45
    }
    fmt.Println(v6)
    _ = len(v4)
    v5 = (v5 ^ v5)
    var v52 float64 = float64((0.0 - 0.1))
    var v53 bool = v1
    var v54 uint8 = 14
    delete(v2, 9223372036854775807)
    var v55 bool = v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("v54:", v54)
    fmt.Println("v55:", v55)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v53
    _ = v54
    _ = v55
    _ = v6
}
