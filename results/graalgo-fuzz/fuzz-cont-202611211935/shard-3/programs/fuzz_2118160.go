package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string }
type GenS2 struct { F0 uint; F1 string; F2 string }
type GenS3 struct { F0 uint32; F1 int32 }
type GenS4 struct { F0 uint32; F1 string }
type GenS5 struct { F0 int32 }
func genFn0() int64 {
    var v6 int16 = 42
    v7 := GenS1{F0: 42, F1: ("foo bar" + "foo bar")}
    if (!(-9223372036854775808 >= 127)) {
        v8 := GenS2{F0: 2, F1: "", F2: (("" + "hello") + "hello")}
        v6 += 1000
        var v9 uint64 = 799
        _ = v8
        _ = v9
    } else {
        v6 = 1
        switch "" {
        case "":
            for fi10 := 0; fi10 < 13; fi10++ {
                v6 += v6
                _ = fi10
            }
        case "\n\t":
            if true {
                v11 := []int{-227, -2147483648, 42}
                _ = v11
            } else {
                for fi12 := 0; fi12 < 8; fi12++ {
                    v13 := GenS3{F0: 324, F1: -1}
                    v14 := []int{}
                    _ = fi12
                    _ = v13
                    _ = v14
                }
                var v15 uint = 2
                _ = v15
            }
        default:
            var v16 bool = (!(100 == 0))
            _ = v16
        }
    }
    v17 := GenS4{F0: 65535, F1: (("a" + "hello") + "")}
    _ = v17
    _ = v6
    _ = v7
    return 9223372036854775807
}

func main() {
    var v1 uint32 = 4294967295
    var v2 int8 = 1
    v3 := []int{127}
    v4 := make(map[string]uint32)
    var v5 uint32 = 65535
    for fi18 := 0; fi18 < 19; fi18++ {
        v4[("a" + "a")] = 2147483648
        switch -1 {
        case 32768:
            v2 += v2
        case -9223372036854775808:
            var v19 uint32 = v5
            _ = v19
        case 701:
            if ((float32(float64((0.0 * (float32((((0.0 - ((0.1 - float64(0.0)) + 3.14159)) * (float64(0.1) * float64(-2.25))) - (1e10 * 0.1))) * -2.25)))) == (-2.25 / 1e10)) || false) {
                v5 = v1
                var v20 int64 = 0
                switch ("\n\t" + "") {
                case "foo bar":
                    v2 -= (v2 << ((4294967295) & 63))
                case "hello":
                    var v21 uint32 = 42
                    _ = v21
                case "\n\t":
                    for fi22 := 0; fi22 < 12; fi22++ {
                        if len(v3) > 0 {
                            v3[0] = 127
                        }
                        _ = fi22
                    }
                }
                _ = v20
            } else {
                v1 += 468
                v1 -= (^(v5))
            }
        }
        fmt.Println(fi18, fi18, v2)
        _ = fi18
    }
    v5 -= v5
    v1 += v5
    v3 = append(v3, 32767)
    fmt.Println(v1, v5)
    var v23 float64 = 0.0
    _ = len(v4)
    if true {
        v1 += 2147483647
        if false {
            if len(v3) > 0 {
                v3[0] = 2147483647
            }
            v5 += (v5 * v5)
            var v24 string = "foo bar"
            _ = v24
        }
        switch ("a" + ("foo bar" + "")) {
        case "a":
            v25 := GenS2{F0: 0, F1: ("\n\t" + ("a" + "a")), F2: (("a" + "\n\t") + ("hello" + "\n\t"))}
            _ = v25
        case "":
            var v26 uint16 = 32767
            _ = v26
        default:
            v3 = append(v3, 726)
        }
    }
    if (0 <= -219) {
        switch -986 {
        case 128:
            v23 += v23
        }
        v27 := GenS5{F0: 0}
        v2 += (v2 + (v2 >> ((2) & 63)))
        _ = v27
    } else {
        delete(v4, "foo bar")
        v28 := GenS5{F0: -582}
        var v29 uint = 201
        _ = v28
        _ = v29
    }
    switch "hello" {
    case "":
        for fi30 := 0; fi30 < 19; fi30++ {
            v31 := GenS2{F0: 0, F1: "\n\t", F2: ""}
            if false {
                v4[("hello" + "")] = (v1 << ((255) & 63))
            }
            _ = fi30
            _ = v31
        }
    case "\n\t":
        for fi32 := 0; fi32 < 17; fi32++ {
            var v33 uint = 42
            fmt.Println(fi32, v33)
            _ = fi32
            _ = v33
        }
    }
    for fi34 := 0; fi34 < 16; fi34++ {
        v35 := GenS3{F0: v5, F1: -1}
        v5 += 506
        _ = fi34
        _ = v35
    }
    v23 -= (1e10 / -2.25)
    var v36 uint32 = v5
    for fi37 := 0; fi37 < 3; fi37++ {
        var v38 uint8 = 255
        fmt.Println(fi37)
        var v39 float64 = (v23 - 1.5)
        _ = fi37
        _ = v38
        _ = v39
    }
    if false {
        var v40 uint64 = 2
        if false {
            fmt.Println(v40, v23)
            var v41 int8 = v2
            _ = v41
        }
        _ = v40
    }
    if true {
        var v42 int16 = 1000
        var v43 bool = true
        _ = v42
        _ = v43
    } else {
        v36 = (v1 / ((v1) | 1))
        delete(v4, ("a" + "\n\t"))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v36
    _ = v4
    _ = v5
}
