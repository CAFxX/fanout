package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
func genFn0(p0 string) bool {
    for fi5 := 0; fi5 < 1; fi5++ {
        for fi6 := 0; fi6 < 2; fi6++ {
            var v7 uint64 = 52
            var v8 int64 = -1
            _ = fi6
            _ = v7
            _ = v8
        }
        _ = fi5
    }
    for fi9 := 0; fi9 < 6; fi9++ {
        fmt.Println(p0, fi9)
        _ = fi9
    }
    _ = p0
    return true
}

func main() {
    v1 := make(map[uint32]string)
    var v2 int64 = -1
    var v3 uint32 = 0
    v4 := make(map[string]int)
    if false {
        if (65535 <= 65535) {
            if true {
                v2 = (v2 >> ((0) & 63))
            }
            fmt.Println(v2)
        } else {
            if (-84 > v2) {
                if true {
                    v10 := GenS1{F0: 2, F1: 0}
                    for fi11 := 0; fi11 < 1; fi11++ {
                        for fi12 := 0; fi12 < 12; fi12++ {
                            switch "" {
                            case "a":
                                if true {
                                    var v13 int = -1
                                    v14 := []int{-604, (fi11 & fi12)}
                                    _ = v13
                                    _ = v14
                                } else {
                                    switch fi12 {
                                    case 32768:
                                        fmt.Println(v3)
                                    default:
                                        v15 := []int{(^(fi11)), -2147483648, (-(fi11))}
                                        _ = v15
                                    }
                                }
                            case "\n\t":
                                delete(v1, (v3 << ((42) & 63)))
                            case "":
                                var v16 uint16 = 32767
                                _ = v16
                            default:
                                switch (("hello" + "\n\t") + ((("foo bar" + (((("a" + ((("a" + "hello") + "a") + "hello")) + "\n\t") + "hello") + "a")) + ("a" + "a")) + ("hello" + ("hello" + "a")))) {
                                case "\n\t":
                                    fmt.Println(fi12, v3)
                                default:
                                    var v17 uint16 = 239
                                    _ = v17
                                }
                            }
                            v4["a"] = fi12
                            _ = fi12
                        }
                        _ = fi11
                    }
                    _ = v10
                } else {
                    var v18 int32 = 129
                    var v19 float32 = -2.25
                    _ = v18
                    _ = v19
                }
                if (-470 != -1) {
                    v3 = 65535
                }
            }
            v20 := GenS1{F0: 412, F1: v2}
            v2 = (v2 + 9223372036854775807)
            _ = v20
        }
        switch "\n\t" {
        case "\n\t":
            var v21 int64 = 42
            _ = v21
        }
    } else {
        var v22 float32 = float32(0.0)
        v3 += 2147483648
        _ = v22
    }
    delete(v1, (^(v3)))
    v3 += (v3 * v3)
    v3 += (v3 | (v3 / ((42) | 1)))
    var v23 int64 = v2
    fmt.Println(v23, v23)
    v3 = (v3 + v3)
    if false {
        v24 := []int{42, 0, -965}
        _ = v24
    } else {
        v23 -= v23
        var v25 uint16 = 2
        _ = v25
    }
    switch "foo bar" {
    case "a":
        switch 2147483647 {
        case 128:
            delete(v4, "\n\t")
        case -213:
            v3 = 65535
        case -43:
            v2 -= v2
        }
    case "foo bar":
        if (false || false) {
            v26 := []int{32767, -1, 32768}
            _ = v26
        }
    case "hello":
        v23 = (v23 >> ((42) & 63))
    }
    for fi27 := 0; fi27 < 1; fi27++ {
        if false {
            var v28 uint64 = 734
            if (!(4294967295 == 42)) {
                v23 += int64(fi27)
                for fi29 := 0; fi29 < 4; fi29++ {
                    v2 -= v2
                    var v30 int32 = 173
                    fmt.Println(v2, v23, fi27)
                    _ = fi29
                    _ = v30
                }
            }
            _ = v28
        } else {
            for fi31 := 0; fi31 < 20; fi31++ {
                var v32 float64 = (0.0 - (-2.25 - ((0.0 + ((-2.25 - 0.0) + 1e10)) - (float64(0.1) * 0.1))))
                fmt.Println(v32, v23, fi31)
                _ = fi31
                _ = v32
            }
            v33 := []int{fi27, fi27}
            var v34 float32 = 3.14159
            _ = v33
            _ = v34
        }
        _ = fi27
    }
    v3 -= 2
    v35 := []int{}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v35
    _ = v4
}
