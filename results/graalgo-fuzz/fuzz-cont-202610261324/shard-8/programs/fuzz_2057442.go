package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 uint; F2 int8 }
type GenS3 struct { F0 int16; F1 uint64; F2 uint16 }
func genFn0() int8 {
    if (-81 != 9223372036854775807) {
        if (472 != 344) {
            switch 0 {
            case -9223372036854775808:
                var v7 int16 = 719
                _ = v7
            case 2147483647:
                var v8 uint16 = 2
                _ = v8
            }
            var v9 uint64 = 0
            v10 := []int{486, 32768}
            _ = v10
            _ = v9
        } else {
            var v11 uint8 = 42
            var v12 uint8 = v11
            fmt.Println(v11, v11)
            _ = v11
            _ = v12
        }
        for fi13 := 0; fi13 < 19; fi13++ {
            for fi14 := 0; fi14 < 16; fi14++ {
                switch (fi14 - 4294967295) {
                case -9223372036854775808:
                    v15 := GenS1{F0: -46}
                    _ = v15
                case 910:
                    if ((("" < ("foo bar" + "\n\t")) && false) && (!true)) {
                        fmt.Println(fi14, fi14)
                    }
                case 0:
                    var v16 int16 = 1
                    _ = v16
                default:
                    v17 := GenS1{F0: 127}
                    _ = v17
                }
                var v18 uint16 = 0
                _ = fi14
                _ = v18
            }
            switch fi13 {
            case 192:
                v19 := GenS2{F0: 532, F1: 1, F2: 42}
                _ = v19
            case -874:
                v20 := GenS2{F0: 423, F1: 2, F2: -10}
                _ = v20
            case -1:
                var v21 int = fi13
                _ = v21
            default:
                for fi22 := 0; fi22 < 7; fi22++ {
                    switch fi13 {
                    case -2147483648:
                        var v23 float64 = (((1e10 + 3.14159) / ((-2.25 - float64((0.1 + (1e10 * (1e10 * 0.1))))) - (0.0 * float64(((float32((-2.25 * 1.5)) + -2.25) + 0.0))))) / 1e10)
                        _ = v23
                    default:
                        v24 := []int{(^(fi22))}
                        _ = v24
                    }
                    var v25 int64 = -742
                    _ = fi22
                    _ = v25
                }
            }
            for fi26 := 0; fi26 < 12; fi26++ {
                v27 := GenS1{F0: -1}
                var v28 int16 = 0
                switch "" {
                case "hello":
                    v29 := GenS3{F0: v28, F1: 18446744073709551615, F2: 2}
                    _ = v29
                case "a":
                    var v30 float32 = float32(1.5)
                    _ = v30
                case "foo bar":
                    v28 = v28
                default:
                    if false {
                        v28 = v28
                    } else {
                        fmt.Println(v28)
                    }
                }
                if true {
                    v28 = (v28 & v28)
                }
                _ = fi26
                _ = v27
                _ = v28
            }
            _ = fi13
        }
    }
    if (!(849 > 2147483647)) {
        var v31 int16 = 42
        v32 := GenS3{F0: 303, F1: 4294967295, F2: 875}
        _ = v31
        _ = v32
    } else {
        switch 4294967295 {
        case 193:
            var v33 uint8 = 234
            _ = v33
        default:
            var v34 int64 = 9223372036854775807
            _ = v34
        }
        var v35 int16 = 1
        v35 = (v35 << ((360) & 63))
        _ = v35
    }
    if (42 <= -930) {
        var v36 int = 9223372036854775807
        _ = v36
    }
    return -1
}

func main() {
    var v1 bool = false
    v2 := make(map[string]bool)
    var v3 string = "foo bar"
    var v4 uint = 65535
    var v5 float64 = ((float64((1.5 - 1.5)) - -2.25) * (float64(((3.14159 - 1e10) * 3.14159)) - ((1.5 + (float64(1e10) + 1.5)) / float64(3.14159))))
    var v6 bool = (v5 == v5)
    v37 := GenS1{F0: 127}
    v5 -= (0.1 + v5)
    v3 = v3
    v5 += (-2.25 + v5)
    for fi38 := 0; fi38 < 4; fi38++ {
        if (!false) {
            if (!v6) {
                fmt.Println(v1)
                v39 := []int{}
                _ = v39
            } else {
                v40 := []int{32767, -1, 127, 0}
                v4 -= 209
                var v41 int8 = genFn0()
                _ = v40
                _ = v41
            }
        } else {
            fmt.Println(v3)
            fmt.Println(v4, fi38)
        }
        if (1 >= 2) {
            var v42 float64 = v5
            fmt.Println(v6, fi38, v42)
            v43 := v2[v3]
            _ = v42
            _ = v43
        } else {
            v44 := v2[v3]
            _ = v44
        }
        v2["hello"] = (55 < 1)
        _ = fi38
    }
    fmt.Println(v1, v6)
    switch 4294967295 {
    case -1:
        for fi45 := 0; fi45 < 4; fi45++ {
            v3 = v3
            fmt.Println(v5, fi45, v6)
            _ = fi45
        }
    case 0:
        v46 := []int{-619, 1, -957}
        _ = v46
    case -895:
        var v47 uint = v4
        _ = v47
    }
    v1 = (0.1 > 0.0)
    v5 -= ((v5 - (float64(-2.25) * 3.14159)) - v5)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
