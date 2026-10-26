package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int32; F1 int64 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 bool; F1 int64; F2 int32 }
type GenS6 struct { F0 string; F1 int32; F2 int }
type GenS7 struct { F0 int16; F1 int; F2 uint8 }
func genFn0(p0 uint32, p1 uint8, p2 int16) uint16 {
    var v7 int = -9223372036854775808
    if (v7 >= (v7 * 32767)) {
        var v8 uint64 = 4294967295
        var v9 string = ""
        var v10 uint32 = p0
        _ = v10
        _ = v8
        _ = v9
    }
    var v11 string = ""
    v12 := []int{v7, v7, 42}
    for fi13 := 0; fi13 < 16; fi13++ {
        var v14 uint32 = p0
        _ = fi13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    _ = v7
    return 0
}
func genFn1() bool {
    v15 := GenS1{F0: 1}
    v16 := GenS2{F0: -9223372036854775808}
    switch "hello" {
    case "":
        if (!((-676 != 933) && (-81 == 0))) {
            var v17 float32 = 0.0
            for fi18 := 0; fi18 < 4; fi18++ {
                switch fi18 {
                case -9223372036854775808:
                    v19 := []int{fi18, (fi18 - fi18)}
                    _ = v19
                case -1:
                    for fi20 := 0; fi20 < 10; fi20++ {
                        switch (("" + "\n\t") + "a") {
                        case "foo bar":
                            if (42 != 127) {
                                var v21 uint = 611
                                v22 := GenS3{F0: int32(fi20), F1: -1}
                                _ = v21
                                _ = v22
                            }
                        case "":
                            var v23 string = "foo bar"
                            var v24 int16 = 1
                            _ = v23
                            _ = v24
                        default:
                            var v25 uint32 = 463
                            _ = v25
                        }
                        _ = fi20
                    }
                case 2147483647:
                    var v26 int16 = -32768
                    _ = v26
                default:
                    if false {
                        var v27 uint16 = 117
                        _ = v27
                    } else {
                        var v28 bool = ("" == (("" + ("\n\t" + "\n\t")) + "a"))
                        fmt.Println(v17, v17, v28)
                        v29 := GenS3{F0: -297, F1: 9223372036854775807}
                        _ = v28
                        _ = v29
                    }
                }
                if (!(v17 < (-2.25 - 1e10))) {
                    v17 = v17
                }
                _ = fi18
            }
            if (true && (v17 != ((v17 - 0.0) / ((1.5 / v17) - v17)))) {
                v17 = v17
            } else {
                v17 -= (v17 / (1.5 / 3.14159))
                fmt.Println(v17)
            }
            _ = v17
        } else {
            v30 := []int{-1, 194, -9223372036854775808, 128}
            for fi31 := 0; fi31 < 17; fi31++ {
                v32 := GenS4{F0: (("a" + ("\n\t" + "\n\t")) + "\n\t")}
                fmt.Println(fi31)
                _ = fi31
                _ = v32
            }
            _ = v30
        }
    case "a":
        var v33 int64 = -674
        _ = v33
    case "foo bar":
        var v34 int = -9223372036854775808
        _ = v34
    default:
        var v35 bool = (("\n\t" + "\n\t") >= ("hello" + "foo bar"))
        _ = v35
    }
    var v36 uint16 = genFn0(2147483647, 255, -1)
    _ = v15
    _ = v16
    _ = v36
    return (!true)
}

func main() {
    v1 := make(map[uint32]string)
    var v2 float64 = 1.5
    v3 := []string{("foo bar" + ("hello" + "hello"))}
    var v4 uint64 = 4294967295
    v5 := []uint32{652, 224, 2147483648}
    v6 := make(map[int]int)
    if true {
        v4 = 1
        for fi37 := 0; fi37 < 10; fi37++ {
            var v38 int16 = -1
            switch fi37 {
            case 32768:
                for fi39 := 0; fi39 < 8; fi39++ {
                    v2 = float64(1.5)
                    _ = fi39
                }
            case 42:
                var v40 int8 = -1
                _ = v40
            case 0:
                var v41 string = ("" + "a")
                _ = v41
            default:
                v38 = v38
            }
            v42 := GenS5{F0: (true && false), F1: int64(v38), F2: 0}
            _ = fi37
            _ = v38
            _ = v42
        }
        for fi43 := 0; fi43 < 10; fi43++ {
            var v44 uint16 = 32767
            v2 = float64(1e10)
            _ = fi43
            _ = v44
        }
    }
    fmt.Println(v2, v4)
    v45 := v6[534]
    var v46 uint = 42
    _ = len(v6)
    var v47 uint = 65535
    if true {
        v45 -= 9223372036854775807
        for fi48 := 0; fi48 < 12; fi48++ {
            fmt.Println(v2)
            _ = fi48
        }
    }
    v45 += v45
    switch "\n\t" {
    case "":
        v49 := GenS2{F0: (v45 * v45)}
        _ = v49
    case "a":
        _ = len(v5)
    case "foo bar":
        for fi50 := 0; fi50 < 8; fi50++ {
            if true {
                v51 := GenS1{F0: 127}
                v46 = v46
                _ = v51
            } else {
                var v52 int = (^(v45))
                var v53 uint64 = (v4 | (v4 << (((^(v47))) & 63)))
                v54 := GenS6{F0: "", F1: 2147483647, F2: (fi50 ^ v52)}
                _ = v52
                _ = v53
                _ = v54
            }
            v6[32768] = fi50
            _ = fi50
        }
    }
    fmt.Println(v46, v2, v46)
    v55 := GenS7{F0: -892, F1: (v45 >> ((v46) & 63)), F2: 0}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v47:", v47)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v47
    _ = v5
    _ = v55
    _ = v6
}
