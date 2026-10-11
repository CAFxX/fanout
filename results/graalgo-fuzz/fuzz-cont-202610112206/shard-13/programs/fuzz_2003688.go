package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 string; F2 int64 }
type GenS2 struct { F0 bool; F1 int64 }
type GenS3 struct { F0 uint32; F1 int16 }
func genFn0() int {
    for fi7 := 0; fi7 < 9; fi7++ {
        fmt.Println(fi7)
        for fi8 := 0; fi8 < 18; fi8++ {
            var v9 uint32 = 566
            fmt.Println(fi7, v9, fi7)
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    var v10 uint = 255
    for fi11 := 0; fi11 < 10; fi11++ {
        var v12 float32 = 0.1
        if ((407 >= uint64(v10)) && (!false)) {
            if true {
                v13 := GenS1{F0: "", F1: "a", F2: -9223372036854775808}
                if (42 == 2) {
                    v14 := []int{(fi11 & fi11), (fi11 - fi11)}
                    if true {
                        v14 = append(v14, fi11)
                    } else {
                        if len(v14) >= 2 {
                            v14 = v14[1:len(v14)]
                        }
                    }
                    var v15 uint = (v10 | v10)
                    _ = v14
                    _ = v15
                }
                v12 -= v12
                _ = v13
            }
            var v16 uint16 = 0
            _ = v16
        } else {
            v17 := []int{9223372036854775807}
            var v18 float32 = v12
            _ = v17
            _ = v18
        }
        _ = fi11
        _ = v12
    }
    v19 := []int{127}
    _ = v10
    _ = v19
    return 127
}

func main() {
    var v1 float64 = 0.0
    v2 := []uint32{311, 4294967295, 695, 4294967295}
    var v3 int64 = 118
    v4 := make(map[int]int)
    var v5 uint8 = 255
    var v6 uint8 = 42
    v5 = v6
    var v20 float32 = 0.0
    v21 := GenS1{F0: (("\n\t" + "foo bar") + ""), F1: "hello", F2: 9223372036854775807}
    for fi22 := 0; fi22 < 11; fi22++ {
        for fi23 := 0; fi23 < 3; fi23++ {
            if true {
                switch "\n\t" {
                case "\n\t":
                    if ((^(v3)) == -267) {
                        switch fi22 {
                        case 614:
                            v1 -= ((v1 + 1.5) / v1)
                        }
                        if (!(false && (!false))) {
                            if true {
                                fmt.Println(fi22, fi23, v3)
                            }
                            delete(v4, -2147483648)
                        } else {
                            var v24 int64 = (-(v3))
                            v24 += 9223372036854775807
                            _ = v24
                        }
                    }
                case "":
                    v1 = (float64(0.0) / v1)
                case "hello":
                    v25 := GenS1{F0: "hello", F1: "", F2: v3}
                    _ = v25
                }
                fmt.Println(v20, fi23)
            } else {
                var v26 uint = 2
                if false {
                    var v27 bool = ((-1 == 42) || ((^(v26)) != v26))
                    _ = v27
                } else {
                    fmt.Println(v6, v20)
                    var v28 uint16 = 65535
                    _ = v28
                }
                v29 := GenS2{F0: true, F1: 42}
                _ = v26
                _ = v29
            }
            _ = fi23
        }
        v1 -= v1
        for fi30 := 0; fi30 < 12; fi30++ {
            if false {
                var v31 int32 = 1
                _ = v31
            } else {
                v32 := GenS2{F0: true, F1: v3}
                if false {
                    var v33 uint64 = 783
                    switch "foo bar" {
                    case "foo bar":
                        v1 += 1e10
                    }
                    _ = v33
                } else {
                    if true {
                        var v34 int64 = v3
                        _ = v34
                    }
                    var v35 float64 = (v1 / -2.25)
                    _ = v35
                }
                v36 := GenS2{F0: ((1 != -663) || (1.5 >= float32(v1))), F1: 9223372036854775807}
                _ = v32
                _ = v36
            }
            var v37 uint32 = 2147483648
            v37 = (^(v37))
            _ = fi30
            _ = v37
        }
        _ = fi22
    }
    switch 32768 {
    case -1:
        var v38 uint16 = 42
        _ = v38
    case -372:
        if ((255 > 4294967295) && (255 != 645)) {
            v2 = append(v2, 993)
        }
    case -2147483648:
        v3 -= 42
    default:
        v20 -= (1e10 + 1.5)
    }
    if (!((("hello" + ((("hello" + "") + "a") + ("hello" + "a"))) + "foo bar") > "\n\t")) {
        if true {
            v4[1] = 11
            switch -17 {
            case 9223372036854775807:
                var v39 string = (((("foo bar" + "\n\t") + "a") + "hello") + "a")
                _ = v39
            }
        }
        var v40 int32 = 411
        var v41 float32 = (v20 - 3.14159)
        _ = v40
        _ = v41
    } else {
        var v42 uint8 = v6
        _ = v42
    }
    v43 := GenS3{F0: 2147483647, F1: 695}
    v3 += (v3 / ((v3) | 1))
    if len(v2) > 0 {
        v2[0] = 1
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v43
    _ = v5
    _ = v6
}
