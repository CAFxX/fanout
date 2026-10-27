package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint }
func genFn0() int16 {
    var v6 uint32 = 2147483647
    if true {
        for fi7 := 0; fi7 < 4; fi7++ {
            if false {
                v8 := []int{(fi7 ^ -513), (-(fi7))}
                if false {
                    v6 = 65535
                } else {
                    var v9 string = "hello"
                    v6 -= (v6 | 2147483648)
                    v9 = ((v9 + ("hello" + v9)) + "a")
                    _ = v9
                }
                _ = v8
            }
            var v10 float32 = 1e10
            _ = fi7
            _ = v10
        }
        if true {
            v6 += 883
            v6 = (v6 / ((v6) | 1))
            v11 := GenS1{F0: 716}
            _ = v11
        }
        var v12 int = 2147483647
        _ = v12
    } else {
        var v13 int16 = -452
        var v14 string = (((("hello" + (("hello" + "a") + ("a" + "\n\t"))) + "a") + (("a" + "\n\t") + "a")) + "hello")
        if (1 <= 0) {
            switch 4294967295 {
            case 675:
                var v15 uint8 = 128
                _ = v15
            case 153:
                v16 := GenS2{F0: 4294967295}
                _ = v16
            case 616:
                v14 = v14
            }
        } else {
            v17 := []int{4294967295, -9223372036854775808, 128}
            var v18 float32 = 1.5
            if true {
                var v19 uint64 = 946
                _ = v19
            } else {
                var v20 int64 = 515
                if false {
                    var v21 uint16 = 1
                    v17 = append(v17, -1)
                    _ = v21
                }
                _ = v20
            }
            _ = v17
            _ = v18
        }
        _ = v13
        _ = v14
    }
    var v22 int = 32767
    v6 += 929
    if false {
        var v23 float32 = 0.0
        v24 := GenS2{F0: 4294967295}
        v25 := GenS1{F0: 42}
        _ = v23
        _ = v24
        _ = v25
    }
    _ = v22
    _ = v6
    return 1000
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int8 = 100
    var v3 float64 = (1.5 / ((1e10 * ((float64((0.1 * (0.1 / float32(1.5)))) / 1.0) - (0.0 + float64((3.14159 - 0.0))))) + float64((-2.25 * 0.0))))
    var v4 int32 = 1
    v5 := []int{-121, -662, -2147483648, 9223372036854775807}
    switch "" {
    case "\n\t":
        if (1.5 == (1.5 + v3)) {
            for fi26 := 0; fi26 < 5; fi26++ {
                fmt.Println(fi26, v3)
                var v27 uint16 = 0
                _ = fi26
                _ = v27
            }
        }
    case "a":
        for fi28 := 0; fi28 < 16; fi28++ {
            v2 -= (v2 ^ v2)
            _ = fi28
        }
    case "foo bar":
        switch 32768 {
        case 0:
            for fi29 := 0; fi29 < 15; fi29++ {
                fmt.Println(v2, fi29, v2)
                var v30 int32 = (^(v4))
                _ = fi29
                _ = v30
            }
        case 687:
            switch ("a" + "hello") {
            case "foo bar":
                var v31 float32 = float32(v3)
                _ = v31
            case "a":
                v5 = append(v5, 1)
            case "\n\t":
                if true {
                    if (false || ((!false) || true)) {
                        if true {
                            delete(v1, 648)
                            var v32 int64 = -161
                            var v33 int16 = genFn0()
                            _ = v32
                            _ = v33
                        }
                        fmt.Println(v4, v2)
                        var v34 int64 = -1
                        _ = v34
                    } else {
                        var v35 int = 938
                        _ = v35
                    }
                }
            default:
                var v36 uint = 2
                _ = v36
            }
        case 127:
            for fi37 := 0; fi37 < 10; fi37++ {
                v4 = -1
                _ = fi37
            }
        default:
            v2 = v2
        }
    }
    v4 = (v4 / ((1) | 1))
    var v38 float64 = v3
    var v39 int32 = (^(v4))
    v3 += 3.14159
    v40 := GenS2{F0: 1}
    var v41 uint8 = 128
    v3 += float64((1.5 / float32(1.5)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v41:", v41)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
}
