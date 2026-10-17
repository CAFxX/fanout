package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 bool }
func genFn0(p0 uint16, p1 string, p2 uint16) int64 {
    fmt.Println(p2)
    if false {
        var v6 uint8 = 149
        if false {
            if true {
                for fi7 := 0; fi7 < 2; fi7++ {
                    v8 := []int{fi7, (fi7 + (^(fi7))), fi7, fi7}
                    v9 := GenS1{F0: (^(v6)), F1: 18446744073709551615, F2: (952 != 65535)}
                    p2 += p0
                    _ = fi7
                    _ = v8
                    _ = v9
                }
                v10 := GenS1{F0: v6, F1: 448, F2: false}
                v11 := []int{-1, 9223372036854775807, 146}
                _ = v10
                _ = v11
            } else {
                for fi12 := 0; fi12 < 14; fi12++ {
                    v13 := GenS1{F0: (v6 - (v6 - v6)), F1: 65535, F2: false}
                    _ = fi12
                    _ = v13
                }
                var v14 uint64 = uint64(v6)
                var v15 bool = true
                _ = v14
                _ = v15
            }
        } else {
            p0 += p2
        }
        var v16 string = ""
        _ = v16
        _ = v6
    }
    var v17 uint32 = 2147483647
    for fi18 := 0; fi18 < 17; fi18++ {
        var v19 float64 = float64(((1.5 - (0.0 * 3.14159)) / -2.25))
        if false {
            switch ((p1 + (p1 + "foo bar")) + "hello") {
            case "a":
                switch fi18 {
                case 559:
                    v20 := []int{-377}
                    _ = v20
                case 1:
                    fmt.Println(p0, v17)
                default:
                    var v21 uint = 18446744073709551615
                    _ = v21
                }
            case "\n\t":
                if false {
                    var v22 string = "a"
                    fmt.Println(fi18, v19, fi18)
                    var v23 bool = ((1 >= 1) || (true && ((-787 == -9223372036854775808) || (!true))))
                    _ = v22
                    _ = v23
                } else {
                    v24 := GenS1{F0: 0, F1: 170, F2: true}
                    v25 := []int{fi18, (fi18 & (-(fi18))), fi18}
                    _ = v24
                    _ = v25
                }
            }
            v26 := GenS1{F0: 2, F1: 676, F2: (!(-1 <= 100))}
            v27 := GenS1{F0: 13, F1: 2, F2: true}
            _ = v26
            _ = v27
        }
        p2 -= uint16(v17)
        _ = fi18
        _ = v19
    }
    v17 += (v17 - v17)
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    return -927
}

func main() {
    var v1 int8 = 1
    var v2 int16 = 42
    var v3 int = 127
    var v4 uint16 = 804
    var v5 string = "\n\t"
    v2 = (-(v2))
    v28 := []int{}
    v3 = 846
    v1 = v1
    if true {
        var v29 int64 = -27
        _ = v29
    } else {
        for fi30 := 0; fi30 < 4; fi30++ {
            if (!false) {
                var v31 int16 = v2
                _ = v31
            }
            if len(v28) > 0 {
                v28[0] = v3
            }
            _ = fi30
        }
    }
    v2 = v2
    for fi32 := 0; fi32 < 14; fi32++ {
        v33 := GenS1{F0: 65, F1: 351, F2: (false || true)}
        switch v3 {
        case 620:
            v1 += (v1 + v1)
        case 0:
            v34 := GenS1{F0: 137, F1: 0, F2: (127 >= 127)}
            _ = v34
        case 32767:
            v3 -= (v3 - v3)
        }
        switch "" {
        case "a":
            v4 = 0
        case "":
            if len(v28) > 0 {
                v28[0] = int(v4)
            }
        case "hello":
            _ = len(v28)
        default:
            for fi35 := 0; fi35 < 12; fi35++ {
                fmt.Println(fi32, v2)
                var v36 uint32 = 2147483648
                v3 += fi32
                _ = fi35
                _ = v36
            }
        }
        _ = fi32
        _ = v33
    }
    var v37 uint64 = 2
    for fi38 := 0; fi38 < 6; fi38++ {
        var v39 int8 = 0
        var v40 uint64 = 42
        _ = fi38
        _ = v39
        _ = v40
    }
    v5 = v5
    v41 := GenS1{F0: 128, F1: 255, F2: false}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v37
    _ = v4
    _ = v41
    _ = v5
}
