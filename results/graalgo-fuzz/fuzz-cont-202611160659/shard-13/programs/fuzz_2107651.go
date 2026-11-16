package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int }
type GenS2 struct { F0 int16; F1 int64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint8; F1 uint32; F2 uint64 }

func main() {
    var v1 float32 = 0.0
    v2 := make(map[int]string)
    v3 := make(map[int]string)
    switch "" {
    case "\n\t":
        v1 -= (v1 + -2.25)
    default:
        for fi4 := 0; fi4 < 14; fi4++ {
            v1 = (((v1 + -2.25) + v1) + float32((3.14159 / -2.25)))
            v5 := GenS1{F0: "", F1: 127}
            _ = fi4
            _ = v5
        }
    }
    fmt.Println(v1)
    switch -374 {
    case 561:
        for fi6 := 0; fi6 < 3; fi6++ {
            switch fi6 {
            case -789:
                var v7 float64 = 1.5
                _ = v7
            case -2147483648:
                for fi8 := 0; fi8 < 8; fi8++ {
                    var v9 uint32 = 0
                    if false {
                        _ = len(v2)
                        v10 := GenS1{F0: "foo bar", F1: -9223372036854775808}
                        _ = v10
                    } else {
                        var v11 string = ("hello" + "foo bar")
                        _ = v11
                    }
                    _ = fi8
                    _ = v9
                }
            default:
                switch "hello" {
                case "":
                    v12 := GenS2{F0: 0, F1: -1}
                    _ = v12
                case "hello":
                    v1 = (-2.25 + v1)
                default:
                    v1 -= v1
                }
            }
            for fi13 := 0; fi13 < 6; fi13++ {
                switch 837 {
                case 109:
                    switch 32767 {
                    case -252:
                        var v14 float64 = 1e10
                        _ = v14
                    case 679:
                        v15 := []int{(^(fi6)), (^(fi13))}
                        _ = v15
                    }
                case -9223372036854775808:
                    v16 := GenS3{F0: 0}
                    _ = v16
                case 128:
                    if true {
                        if true {
                            v1 = -2.25
                        }
                        v17 := GenS1{F0: ("\n\t" + "foo bar"), F1: 127}
                        _ = v17
                    } else {
                        v18 := GenS3{F0: -15}
                        _ = v18
                    }
                }
                _ = fi13
            }
            v1 -= v1
            _ = fi6
        }
    case 32767:
        v19 := GenS2{F0: -32768, F1: 0}
        _ = v19
    case -1:
        v1 += v1
    }
    v20 := GenS4{F0: 69, F1: 1, F2: 2}
    if (-9223372036854775808 > -950) {
        v21 := GenS1{F0: "a", F1: 47}
        v22 := GenS3{F0: -1}
        _ = v21
        _ = v22
    } else {
        var v23 string = ""
        for fi24 := 0; fi24 < 10; fi24++ {
            var v25 uint = 505
            v25 = v25
            _ = fi24
            _ = v25
        }
        v23 = v23
        _ = v23
    }
    if true {
        v1 += ((v1 * (float32(((0.1 * 3.14159) / (3.14159 - 1e10))) - v1)) * ((-2.25 * 1.5) - v1))
        if ((((-2.25 >= float32(-2.25)) && false) && (0 > 65535)) || (((-2.25 + 0.0) / 1e10) != -2.25)) {
            var v26 string = "\n\t"
            _ = v26
        }
        var v27 float64 = -2.25
        _ = v27
    }
    var v28 int = 42
    fmt.Println(v28, v1)
    switch (("hello" + "\n\t") + "foo bar") {
    case "":
        var v29 int = -2147483648
        _ = v29
    }
    var v30 uint64 = 132
    fmt.Println(v28, v30)
    for fi31 := 0; fi31 < 3; fi31++ {
        if false {
            for fi32 := 0; fi32 < 19; fi32++ {
                v28 = fi32
                var v33 float64 = ((float64(1e10) * (0.1 / 3.14159)) - 0.1)
                var v34 float64 = v33
                _ = fi32
                _ = v33
                _ = v34
            }
            fmt.Println(v30, fi31, v30)
        } else {
            var v35 uint64 = (^(v30))
            if true {
                v36 := GenS2{F0: 1000, F1: 42}
                fmt.Println(v30)
                var v37 uint16 = 549
                _ = v36
                _ = v37
            }
            v28 = (fi31 - fi31)
            _ = v35
        }
        _ = fi31
    }
    if false {
        if false {
            var v38 int8 = 1
            var v39 uint = 0
            _ = v38
            _ = v39
        } else {
            v40 := GenS2{F0: 702, F1: -9223372036854775808}
            v28 = 1
            v28 += v28
            _ = v40
        }
        v41 := v2[-2147483648]
        _ = v41
    }
    var v42 int = v28
    v43 := GenS2{F0: 651, F1: 1}
    v42 = (v42 % ((-309) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v2
    _ = v20
    _ = v28
    _ = v3
    _ = v30
    _ = v42
    _ = v43
}
