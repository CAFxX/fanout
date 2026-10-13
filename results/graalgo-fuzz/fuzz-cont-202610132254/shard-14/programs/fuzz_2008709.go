package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64; F1 int8; F2 string }
func genFn0() uint32 {
    var v5 float32 = -2.25
    for fi6 := 0; fi6 < 14; fi6++ {
        for fi7 := 0; fi7 < 2; fi7++ {
            fmt.Println(fi7, v5)
            if false {
                fmt.Println(fi7)
                v5 = ((v5 - v5) / -2.25)
                var v8 uint32 = 606
                _ = v8
            }
            var v9 int8 = 42
            _ = fi7
            _ = v9
        }
        _ = fi6
    }
    v10 := GenS1{F0: (18446744073709551615 > 65535)}
    v11 := []int{}
    v11 = append(v11, 9223372036854775807)
    _ = v10
    _ = v11
    _ = v5
    return 814
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[string]string)
    var v3 uint8 = 2
    var v4 uint32 = 52
    v12 := v1[1]
    v13 := GenS1{F0: true}
    var v14 float64 = 0.0
    for fi15 := 0; fi15 < 11; fi15++ {
        var v16 uint32 = (^(v4))
        _ = fi15
        _ = v16
    }
    v17 := v2[v12]
    v14 = 0.0
    v18 := GenS1{F0: false}
    var v19 int64 = 730
    v20 := GenS1{F0: (false || false)}
    for fi21 := 0; fi21 < 6; fi21++ {
        var v22 int = (fi21 | fi21)
        _ = fi21
        _ = v22
    }
    if false {
        v23 := []int{127, 9223372036854775807}
        _ = v23
    }
    for fi24 := 0; fi24 < 4; fi24++ {
        v14 = ((((float64(3.14159) - v14) - 3.14159) / ((0.0 + v14) - v14)) / v14)
        v17 = v17
        switch fi24 {
        case 2147483647:
            v25 := GenS1{F0: false}
            _ = v25
        case 210:
            var v26 uint32 = genFn0()
            _ = v26
        case 32767:
            v27 := []int{fi24, fi24, fi24}
            _ = v27
        default:
            for fi28 := 0; fi28 < 11; fi28++ {
                if (true || (-374 == (fi24 >> ((42) & 63)))) {
                    if (((!(v14 > v14)) && true) || (18446744073709551615 < 18446744073709551615)) {
                        v4 += (v4 >> ((397) & 63))
                        v19 = (v19 & v19)
                    }
                } else {
                    switch (fi24 + fi28) {
                    case 32767:
                        v29 := []int{}
                        _ = v29
                    case 127:
                        var v30 uint = 281
                        _ = v30
                    case -481:
                        var v31 uint16 = 2
                        _ = v31
                    default:
                        if (!((2 > 1) && (!("foo bar" == (v17 + v12))))) {
                            if true {
                                var v32 int16 = 1
                                _ = v32
                            } else {
                                fmt.Println(v19, fi24, v19)
                                if (true || (!false)) {
                                    v4 = 65535
                                }
                                v33 := []int{fi24}
                                _ = v33
                            }
                        } else {
                            v34 := GenS2{F0: 2, F1: 1, F2: ""}
                            _ = v34
                        }
                    }
                    var v35 uint8 = 42
                    _ = v35
                }
                var v36 uint32 = v4
                fmt.Println(v12, v36, v14)
                _ = fi28
                _ = v36
            }
        }
        _ = fi24
    }
    v4 = 1
    for fi37 := 0; fi37 < 9; fi37++ {
        switch (("a" + v17) + v17) {
        case "foo bar":
            var v38 string = v17
            _ = v38
        }
        _ = fi37
    }
    v3 = (v3 & (^(v3)))
    _ = len(v1)
    v17 = v17
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
