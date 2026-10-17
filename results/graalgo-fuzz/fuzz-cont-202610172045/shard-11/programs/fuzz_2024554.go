package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint }
type GenS2 struct { F0 bool; F1 int32; F2 int64 }
func genFn0(p0 string, p1 int32) int16 {
    for fi7 := 0; fi7 < 5; fi7++ {
        p1 = (p1 / (((p1 | (^(p1)))) | 1))
        _ = fi7
    }
    v8 := []int{127, -1}
    switch 2147483647 {
    case 0:
        var v9 float64 = (0.1 - 3.14159)
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v8
    return -254
}
func genFn1(p0 int32, p1 bool) int {
    p1 = p1
    var v10 int8 = -1
    v10 = 1
    var v11 bool = p1
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return 533
}

func main() {
    var v1 uint = 294
    var v2 uint32 = 165
    v3 := make(map[uint32]uint32)
    v4 := []int{32768}
    v5 := make(map[uint32]uint32)
    v6 := []string{("foo bar" + "hello"), ("foo bar" + "foo bar")}
    v3[(^(v2))] = 42
    for fi12 := 0; fi12 < 18; fi12++ {
        v2 = (v2 >> ((255) & 63))
        if (-1 >= 100) {
            if true {
                var v13 uint32 = 299
                if (-5 >= -1) {
                    v13 = v2
                    v13 = (^(v13))
                }
                delete(v3, uint32(v1))
                _ = v13
            }
            switch 756 {
            case -9223372036854775808:
                switch "" {
                case "hello":
                    switch 32768 {
                    case 32768:
                        if false {
                            var v14 uint = v1
                            var v15 int32 = 1
                            _ = v14
                            _ = v15
                        } else {
                            var v16 float64 = 3.14159
                            _ = v16
                        }
                    case 127:
                        v17 := GenS1{F0: "\n\t", F1: (v1 % ((v1) | 1))}
                        _ = v17
                    }
                case "a":
                    var v18 int8 = -128
                    _ = v18
                case "":
                    v19 := GenS1{F0: "a", F1: 2}
                    _ = v19
                }
            case 32767:
                var v20 uint16 = 32767
                _ = v20
            default:
                v21 := GenS2{F0: false, F1: 931, F2: 1}
                _ = v21
            }
            fmt.Println(v2, v1, fi12)
        }
        _ = fi12
    }
    if (!false) {
        for fi22 := 0; fi22 < 20; fi22++ {
            var v23 int8 = 19
            v24 := GenS1{F0: "", F1: v1}
            _ = fi22
            _ = v23
            _ = v24
        }
        v25 := GenS2{F0: (0.0 != float32(0.1)), F1: -2147483648, F2: -271}
        switch 562 {
        case 48:
            delete(v3, (^(v2)))
        case -2147483648:
            for fi26 := 0; fi26 < 13; fi26++ {
                var v27 bool = false
                for fi28 := 0; fi28 < 15; fi28++ {
                    if (v27 || (!v27)) {
                        v1 += v1
                        v1 = v1
                    }
                    if v27 {
                        if false {
                            v27 = (genFn0("foo bar", -2147483648) != genFn0(("foo bar" + "\n\t"), 0))
                        }
                        var v29 float64 = (0.0 - 3.14159)
                        var v30 uint8 = 76
                        _ = v29
                        _ = v30
                    } else {
                        _ = len(v5)
                    }
                    var v31 string = "foo bar"
                    _ = fi28
                    _ = v31
                }
                _ = fi26
                _ = v27
            }
        }
        _ = v25
    } else {
        if len(v6) > 0 {
            v6[0] = (("hello" + "a") + "hello")
        }
        for fi32 := 0; fi32 < 10; fi32++ {
            fmt.Println(v2, fi32, fi32)
            _ = fi32
        }
    }
    v33 := GenS1{F0: "foo bar", F1: 65535}
    switch "" {
    case "\n\t":
        v1 -= (^(v1))
    case "a":
        delete(v3, v2)
    case "foo bar":
        var v34 int64 = 9223372036854775807
        _ = v34
    }
    switch genFn1(-974, (65535 == (v2 * v2))) {
    case 2147483647:
        var v35 uint = (v1 / ((v1) | 1))
        _ = v35
    }
    var v36 uint64 = 210
    if len(v4) > 0 {
        v4[0] = 32768
    }
    var v37 int8 = 0
    fmt.Println(v1, v2, v1)
    v38 := GenS2{F0: ((1.5 + 0.0) >= 1e10), F1: 42, F2: 985}
    switch "" {
    case "a":
        var v39 uint64 = 4294967295
        _ = v39
    case "hello":
        v1 = (v1 % ((538) | 1))
    }
    v6 = append(v6, "a")
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
}
