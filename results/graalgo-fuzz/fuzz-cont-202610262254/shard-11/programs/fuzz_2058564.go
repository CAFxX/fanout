package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool }
type GenS2 struct { F0 uint8; F1 uint; F2 uint8 }
func genFn0() uint {
    for fi7 := 0; fi7 < 9; fi7++ {
        var v8 int64 = 145
        var v9 int16 = 96
        var v10 int16 = (v9 | v9)
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    for fi11 := 0; fi11 < 20; fi11++ {
        switch "hello" {
        case "foo bar":
            switch -364 {
            case 42:
                var v12 int64 = -1
                _ = v12
            case 907:
                var v13 float32 = (0.0 - (float32(1.5) + (((float32(float64(0.1)) - (0.1 - (3.14159 - -2.25))) - float32((0.1 * (1.5 + -2.25)))) - 0.1)))
                _ = v13
            case 2147483647:
                var v14 uint8 = 2
                _ = v14
            default:
                for fi15 := 0; fi15 < 20; fi15++ {
                    var v16 int16 = 1000
                    if true {
                        var v17 int64 = 994
                        _ = v17
                    }
                    v18 := []int{int(v16), (fi11 - fi11)}
                    if true {
                        v19 := GenS1{F0: (-(fi15)), F1: (0.0 < -2.25)}
                        v20 := GenS1{F0: (fi11 / ((2147483647) | 1)), F1: ("" >= "")}
                        var v21 uint8 = 90
                        _ = v19
                        _ = v20
                        _ = v21
                    } else {
                        var v22 uint64 = 18446744073709551615
                        _ = v22
                    }
                    if len(v18) > 0 {
                        v18[0] = 128
                    }
                    _ = fi15
                    _ = v16
                    _ = v18
                }
            }
        }
        var v23 string = "\n\t"
        if (-120 > 1) {
            var v24 float64 = 3.14159
            v25 := GenS1{F0: fi11, F1: false}
            _ = v24
            _ = v25
        } else {
            for fi26 := 0; fi26 < 12; fi26++ {
                v27 := []int{(-(fi11)), fi11, fi26, (fi26 - (fi11 % ((fi26) | 1)))}
                _ = fi26
                _ = v27
            }
        }
        _ = fi11
        _ = v23
    }
    return 678
}

func main() {
    var v1 int16 = -244
    var v2 int16 = (v1 / ((-32768) | 1))
    v3 := []string{"a", ("" + ("foo bar" + "\n\t"))}
    var v4 int16 = 42
    var v5 uint64 = 1
    var v6 float64 = (3.14159 * 0.1)
    v28 := GenS2{F0: 255, F1: 18446744073709551615, F2: 2}
    v2 = (^(v4))
    if (v4 > 0) {
        for fi29 := 0; fi29 < 1; fi29++ {
            v30 := GenS2{F0: 2, F1: 255, F2: 255}
            _ = fi29
            _ = v30
        }
        v1 = 0
        v31 := GenS1{F0: 1, F1: true}
        _ = v31
    }
    v2 = -1
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    var v32 int32 = -2147483648
    for fi33 := 0; fi33 < 10; fi33++ {
        fmt.Println(fi33, fi33)
        _ = fi33
    }
    for fi34 := 0; fi34 < 11; fi34++ {
        for fi35 := 0; fi35 < 9; fi35++ {
            switch "a" {
            case "foo bar":
                v32 = 87
            default:
                if len(v3) > 0 {
                    v3[0] = (("foo bar" + "hello") + (("" + "foo bar") + "foo bar"))
                }
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                if ((((42 == 1) || (true && false)) && (true || (-597 != -215))) && false) {
                    v32 += (v32 | v32)
                }
            }
            _ = fi35
        }
        var v36 int32 = (v32 / ((v32) | 1))
        var v37 int8 = 0
        _ = fi34
        _ = v36
        _ = v37
    }
    v4 = (v4 * (v2 * v1))
    v3 = append(v3, ("hello" + ("a" + "a")))
    if true {
        fmt.Println(v5)
        if (1.5 != 3.14159) {
            for fi38 := 0; fi38 < 1; fi38++ {
                v4 = (v1 & (v2 + (^(v1))))
                _ = fi38
            }
        } else {
            v2 -= v4
        }
        for fi39 := 0; fi39 < 20; fi39++ {
            v32 = v32
            v5 = v5
            _ = fi39
        }
    }
    for fi40 := 0; fi40 < 17; fi40++ {
        if (true && (!(!(971 >= 496)))) {
            for fi41 := 0; fi41 < 9; fi41++ {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi41
            }
        }
        var v42 uint = 2
        _ = fi40
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
