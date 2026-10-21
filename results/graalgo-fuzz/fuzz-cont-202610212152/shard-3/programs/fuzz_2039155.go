package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 bool; F1 bool; F2 string }
func genFn0() uint8 {
    v7 := GenS1{F0: -9223372036854775808, F1: 42}
    var v8 string = "\n\t"
    for fi9 := 0; fi9 < 15; fi9++ {
        v10 := []int{9223372036854775807, (fi9 / ((fi9) | 1)), (fi9 % ((4294967295) | 1))}
        v10 = append(v10, (fi9 % ((42) | 1)))
        _ = fi9
        _ = v10
    }
    var v11 uint64 = 1
    if (42 > 42) {
        switch 4294967295 {
        case 32767:
            fmt.Println(v11, v11)
        }
    } else {
        if (2 > 2) {
            if true {
                var v12 int64 = 42
                var v13 uint64 = 384
                _ = v12
                _ = v13
            } else {
                v14 := GenS1{F0: 874, F1: v11}
                _ = v14
            }
            for fi15 := 0; fi15 < 4; fi15++ {
                var v16 uint = 0
                v17 := []int{fi15}
                _ = fi15
                _ = v16
                _ = v17
            }
            v18 := []int{9223372036854775807}
            _ = v18
        }
        var v19 int8 = 49
        _ = v19
    }
    _ = v11
    _ = v7
    _ = v8
    return 127
}
func genFn1() int64 {
    var v20 uint16 = 32767
    v20 += v20
    fmt.Println(v20)
    v20 -= (v20 / (((v20 % ((v20) | 1))) | 1))
    _ = v20
    return 1
}

func main() {
    var v1 float64 = 3.14159
    var v2 bool = (262 >= -753)
    var v3 uint8 = 1
    var v4 bool = false
    var v5 uint16 = 2
    var v6 uint32 = 2147483647
    for fi21 := 0; fi21 < 20; fi21++ {
        switch fi21 {
        case -291:
            v22 := GenS1{F0: -1, F1: 2}
            _ = v22
        case 1:
            var v23 int16 = 698
            _ = v23
        case 42:
            switch "" {
            case "a":
                v1 += v1
            case "foo bar":
                var v24 uint16 = (v5 % (((^(v5))) | 1))
                _ = v24
            case "hello":
                switch fi21 {
                case 9223372036854775807:
                    fmt.Println(v4, v3)
                case -9223372036854775808:
                    if false {
                        v5 = (v5 << ((241) & 63))
                        var v25 uint64 = 42
                        v26 := GenS2{F0: v2, F1: v2, F2: "hello"}
                        _ = v25
                        _ = v26
                    } else {
                        v27 := GenS1{F0: -298, F1: 42}
                        v28 := GenS1{F0: 0, F1: uint64(v3)}
                        if v2 {
                            fmt.Println(v5, v4)
                        }
                        _ = v27
                        _ = v28
                    }
                case -1:
                    v29 := []int{(fi21 & fi21), 1, 127}
                    _ = v29
                }
            default:
                switch -675 {
                case 32768:
                    v4 = false
                case 9223372036854775807:
                    for fi30 := 0; fi30 < 13; fi30++ {
                        var v31 int64 = 703
                        v32 := []int{(fi21 + fi21), (fi21 & (fi30 << ((0) & 63))), (fi21 >> ((18446744073709551615) & 63)), (fi21 * (-(fi30)))}
                        _ = fi30
                        _ = v31
                        _ = v32
                    }
                }
            }
        }
        _ = fi21
    }
    for fi33 := 0; fi33 < 10; fi33++ {
        fmt.Println(v4)
        v34 := []int{127, (fi33 - fi33), (fi33 & (fi33 & (fi33 / ((fi33) | 1))))}
        _ = fi33
        _ = v34
    }
    var v35 uint32 = v6
    var v36 int64 = 955
    for fi37 := 0; fi37 < 12; fi37++ {
        for fi38 := 0; fi38 < 5; fi38++ {
            v39 := []int{fi38, fi37}
            _ = fi38
            _ = v39
        }
        var v40 bool = v4
        v41 := []int{fi37, 2147483647}
        _ = fi37
        _ = v40
        _ = v41
    }
    var v42 int = 0
    v36 = v36
    var v43 float32 = ((((-2.25 - ((1e10 + 0.1) * 3.14159)) / 1.0) / 3.14159) - 0.1)
    v44 := []int{42, (v42 >> ((368) & 63))}
    var v45 int32 = -229
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v44:", len(v44), v44)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
