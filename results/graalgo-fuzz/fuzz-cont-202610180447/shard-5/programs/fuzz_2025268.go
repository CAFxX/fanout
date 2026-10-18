package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32 }
type GenS2 struct { F0 int64; F1 int8 }
func genFn0(p0 uint64, p1 int32, p2 uint64) int64 {
    v7 := []int{128}
    p2 += (p2 % ((p2) | 1))
    for fi8 := 0; fi8 < 16; fi8++ {
        v7 = append(v7, (fi8 % ((-1) | 1)))
        v9 := GenS1{F0: 18446744073709551615, F1: (p1 - (p1 + p1))}
        var v10 string = "a"
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 336
}

func main() {
    var v1 float64 = 0.0
    v2 := make(map[string]uint32)
    var v3 int16 = 0
    var v4 int8 = -99
    v5 := make(map[string]uint32)
    v6 := make(map[string]uint32)
    switch "\n\t" {
    case "foo bar":
        v1 = (3.14159 + float64(-2.25))
    }
    v1 += (3.14159 + 0.1)
    switch "hello" {
    case "":
        for fi11 := 0; fi11 < 1; fi11++ {
            if true {
                fmt.Println(v4, v1)
            }
            _ = fi11
        }
    case "hello":
        if true {
            if (-2147483648 > 2147483647) {
                fmt.Println(v3, v4)
            }
            switch 0 {
            case 128:
                delete(v5, "hello")
            }
            if (true && (128 != 0)) {
                var v12 uint64 = 579
                if true {
                    var v13 uint64 = (v12 + v12)
                    switch 42 {
                    case 127:
                        var v14 int64 = -1
                        _ = v14
                    case 32767:
                        if true {
                            v13 = (v12 - 1)
                            if ((0.1 * float64(float32((v1 / v1)))) > v1) {
                                v15 := GenS2{F0: -9223372036854775808, F1: v4}
                                delete(v2, "hello")
                                var v16 bool = (-48 == 2147483647)
                                _ = v15
                                _ = v16
                            }
                        }
                    }
                    var v17 int = -166
                    _ = v13
                    _ = v17
                } else {
                    var v18 int = -104
                    switch ("\n\t" + "hello") {
                    case "a":
                        fmt.Println(v18, v4)
                    case "foo bar":
                        var v19 float32 = (float32((v1 - float64(float32(v1)))) + 1e10)
                        _ = v19
                    case "hello":
                        var v20 uint16 = 2
                        _ = v20
                    default:
                        fmt.Println(v3, v1)
                    }
                    _ = v18
                }
                _ = v12
            }
        } else {
            for fi21 := 0; fi21 < 17; fi21++ {
                v22 := []int{int(v3), fi21, fi21}
                v1 += v1
                _ = fi21
                _ = v22
            }
            _ = len(v6)
        }
    case "a":
        var v23 bool = false
        _ = v23
    default:
        v24 := []int{2147483647}
        _ = v24
    }
    v25 := GenS1{F0: 42, F1: -1}
    if (18446744073709551615 < 724) {
        v26 := GenS2{F0: -868, F1: v4}
        _ = v26
    } else {
        for fi27 := 0; fi27 < 5; fi27++ {
            if true {
                v3 -= 629
                var v28 int8 = (v4 - -128)
                _ = v28
            }
            v3 -= v3
            var v29 float64 = (v1 + v1)
            _ = fi27
            _ = v29
        }
        switch ("a" + ("hello" + ("hello" + "foo bar"))) {
        case "foo bar":
            fmt.Println(v1, v4)
        case "a":
            if true {
                var v30 bool = false
                v1 = 3.14159
                _ = v30
            } else {
                v2[(("a" + "a") + "\n\t")] = 1
                var v31 uint = 198
                _ = v31
            }
        default:
            v4 = (v4 >> ((382) & 63))
        }
    }
    if true {
        var v32 uint16 = 407
        _ = v32
    } else {
        for fi33 := 0; fi33 < 13; fi33++ {
            _ = len(v2)
            _ = fi33
        }
        v4 = -128
        for fi34 := 0; fi34 < 7; fi34++ {
            for fi35 := 0; fi35 < 3; fi35++ {
                var v36 bool = (int(v3) > fi35)
                v3 += -382
                _ = fi35
                _ = v36
            }
            _ = fi34
        }
    }
    if false {
        v4 -= v4
    } else {
        v1 += 1.5
    }
    v37 := v5["hello"]
    v3 -= 924
    var v38 int32 = -2147483648
    switch 1 {
    case -2147483648:
        v4 += 127
    default:
        v37 = 81
    }
    v1 -= float64((1.5 + 0.1))
    switch -9223372036854775808 {
    case 4294967295:
        var v39 uint = 299
        _ = v39
    case 0:
        delete(v5, ("\n\t" + "\n\t"))
    case 62:
        fmt.Println(v37, v38, v4)
    }
    v2["hello"] = 1
    var v40 int32 = 0
    var v41 int16 = (v3 % (((v3 & (v3 - v3))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v37
    _ = v38
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
