package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint; F1 int64; F2 uint16 }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]string)
    var v3 float32 = (((((float32(0.0) * (1e10 / 1.5)) / 1.5) + 3.14159) / (1e10 * 1.5)) - ((-2.25 * (1.5 - 3.14159)) * -2.25))
    var v4 int64 = 1
    v5 := make(map[uint32]int)
    var v6 int32 = -2147483648
    v4 = (v4 & (-(v4)))
    for fi7 := 0; fi7 < 8; fi7++ {
        v2[(^(fi7))] = ("hello" + "a")
        if (false || false) {
            var v8 float64 = 3.14159
            for fi9 := 0; fi9 < 5; fi9++ {
                fmt.Println(fi7, fi9)
                var v10 string = "\n\t"
                _ = fi9
                _ = v10
            }
            _ = v8
        }
        _ = fi7
    }
    v11 := GenS1{F0: 42}
    v3 = 0.1
    if true {
        var v12 string = "hello"
        var v13 uint8 = 42
        _ = v12
        _ = v13
    }
    switch 42 {
    case 32767:
        var v14 uint32 = 181
        _ = v14
    case 127:
        v15 := GenS1{F0: 32767}
        _ = v15
    }
    if (42 != 2) {
        for fi16 := 0; fi16 < 8; fi16++ {
            if (!(0 > 271)) {
                var v17 int32 = (v6 | (^(v6)))
                v17 = v6
                var v18 int = (fi16 + 906)
                _ = v17
                _ = v18
            } else {
                for fi19 := 0; fi19 < 9; fi19++ {
                    v3 = 0.0
                    if false {
                        if (1 <= 216) {
                            v20 := GenS1{F0: 255}
                            switch fi19 {
                            case 128:
                                if (v3 < v3) {
                                    v3 = float32(float64((3.14159 - v3)))
                                    v3 = v3
                                } else {
                                    fmt.Println(v4)
                                    v21 := GenS1{F0: 42}
                                    v6 = 678
                                    _ = v21
                                }
                            default:
                                delete(v2, (fi16 ^ fi19))
                            }
                            if true {
                                switch (fi19 ^ 0) {
                                case 127:
                                    if true {
                                        v22 := []int{(fi16 << ((42) & 63)), (fi19 / ((fi19) | 1))}
                                        var v23 uint = 928
                                        _ = v22
                                        _ = v23
                                    }
                                case -875:
                                    if (false || (4294967295 > 4294967295)) {
                                        var v24 uint32 = 0
                                        _ = v24
                                    }
                                case -1:
                                    v6 = (^(v6))
                                }
                                _ = len(v1)
                            }
                            _ = v20
                        } else {
                            switch (fi16 >> ((65535) & 63)) {
                            case 32768:
                                var v25 uint = 2
                                _ = v25
                            case 128:
                                v4 = 1
                            default:
                                v6 = v6
                            }
                            fmt.Println(v6)
                        }
                        v26 := GenS2{F0: 4294967295, F1: -9223372036854775808, F2: 1}
                        _ = v26
                    } else {
                        v6 += v6
                    }
                    _ = fi19
                }
            }
            switch fi16 {
            case 32767:
                for fi27 := 0; fi27 < 6; fi27++ {
                    v1["hello"] = (0 >= 65535)
                    _ = fi27
                }
            }
            _ = fi16
        }
    }
    fmt.Println(v6)
    for fi28 := 0; fi28 < 15; fi28++ {
        switch ("" + "a") {
        case "\n\t":
            delete(v1, ("foo bar" + ("foo bar" + "hello")))
        case "hello":
            v6 -= v6
        }
        _ = fi28
    }
    var v29 int32 = v6
    v3 = v3
    if false {
        var v30 uint8 = 42
        v6 += (v29 ^ (v29 ^ -2147483648))
        _ = v30
    }
    v6 -= v29
    v6 = (v6 - -1)
    fmt.Println(v29)
    for fi31 := 0; fi31 < 2; fi31++ {
        fmt.Println(v4, v6)
        _ = fi31
    }
    switch 128 {
    case 42:
        var v32 uint64 = 693
        _ = v32
    case -465:
        var v33 uint64 = 18446744073709551615
        _ = v33
    default:
        for fi34 := 0; fi34 < 6; fi34++ {
            v35 := []int{}
            var v36 uint = 530
            v29 = v29
            _ = fi34
            _ = v35
            _ = v36
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
