package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint }
type GenS2 struct { F0 uint16; F1 uint; F2 int16 }
type GenS3 struct { F0 int64; F1 int8 }
type GenS4 struct { F0 uint32; F1 uint8; F2 int8 }
type GenS5 struct { F0 uint64; F1 int64 }
type GenS6 struct { F0 int32; F1 int64; F2 uint32 }
func genFn0() uint32 {
    var v5 int32 = 61
    switch -1 {
    case -1:
        var v6 uint8 = 45
        _ = v6
    }
    v7 := GenS1{F0: 65535, F1: 255}
    for fi8 := 0; fi8 < 6; fi8++ {
        switch (("hello" + "a") + "hello") {
        case "a":
            var v9 int = (^(fi8))
            _ = v9
        case "\n\t":
            fmt.Println(v5, v5)
        case "hello":
            v10 := GenS1{F0: 709, F1: 42}
            _ = v10
        }
        var v11 int8 = 100
        for fi12 := 0; fi12 < 3; fi12++ {
            var v13 int = -611
            if true {
                v5 = v5
            } else {
                v13 = (v13 / ((-588) | 1))
            }
            _ = fi12
            _ = v13
        }
        _ = fi8
        _ = v11
    }
    _ = v5
    _ = v7
    return 973
}

func main() {
    var v1 float32 = 1e10
    v2 := make(map[string]uint32)
    var v3 float64 = -2.25
    var v4 int = 485
    v14 := GenS2{F0: 689, F1: 65535, F2: 1000}
    v2[("a" + "foo bar")] = 0
    v15 := GenS3{F0: 739, F1: -96}
    v16 := GenS2{F0: 1, F1: 65535, F2: 42}
    v17 := []int{v4, 65}
    fmt.Println(v1)
    if (true || (!(655 == -359))) {
        var v18 int32 = -2147483648
        fmt.Println(v4)
        v1 = 3.14159
        _ = v18
    }
    for fi19 := 0; fi19 < 16; fi19++ {
        fmt.Println(v3)
        var v20 uint = 1
        _ = fi19
        _ = v20
    }
    var v21 float32 = 0.1
    for fi22 := 0; fi22 < 15; fi22++ {
        fmt.Println(v4, v3)
        if false {
            fmt.Println(v3, v3, v4)
            var v23 uint64 = 84
            for fi24 := 0; fi24 < 3; fi24++ {
                v4 += (fi22 ^ (v4 | fi22))
                _ = fi24
            }
            _ = v23
        } else {
            if ((fi22 == (-(v4))) || (v21 >= (1.5 / (v21 * 3.14159)))) {
                v25 := GenS1{F0: 2147483647, F1: 42}
                fmt.Println(fi22)
                _ = v25
            } else {
                for fi26 := 0; fi26 < 18; fi26++ {
                    var v27 int32 = 2147483647
                    var v28 uint8 = 2
                    fmt.Println(v4, v27, v4)
                    _ = fi26
                    _ = v27
                    _ = v28
                }
                if (false && true) {
                    fmt.Println(v3, fi22, v1)
                    v29 := v2["hello"]
                    _ = v29
                }
                switch "\n\t" {
                case "\n\t":
                    v30 := GenS3{F0: 0, F1: 100}
                    _ = v30
                }
            }
            fmt.Println(v1, v21, v21)
        }
        _ = fi22
    }
    switch (("a" + (("" + "hello") + ("" + "foo bar"))) + ("foo bar" + "foo bar")) {
    case "hello":
        v31 := GenS1{F0: 0, F1: 2}
        _ = v31
    case "":
        v32 := GenS4{F0: genFn0(), F1: 244, F2: -85}
        _ = v32
    case "a":
        delete(v2, "a")
    }
    for fi33 := 0; fi33 < 9; fi33++ {
        delete(v2, "hello")
        _ = fi33
    }
    if false {
        v34 := GenS5{F0: 0, F1: 42}
        var v35 int16 = 42
        _ = v34
        _ = v35
    }
    v17 = append(v17, -354)
    for fi36 := 0; fi36 < 16; fi36++ {
        fmt.Println(fi36, fi36, v1)
        var v37 uint32 = 4294967295
        _ = fi36
        _ = v37
    }
    for fi38 := 0; fi38 < 15; fi38++ {
        if false {
            v39 := GenS6{F0: -1, F1: -743, F2: 65535}
            for fi40 := 0; fi40 < 13; fi40++ {
                if false {
                    var v41 int16 = 954
                    var v42 uint64 = 4294967295
                    v42 += v42
                    fmt.Println(v3, v41)
                    _ = v41
                    _ = v42
                }
                _ = fi40
            }
            _ = v39
        } else {
            for fi43 := 0; fi43 < 17; fi43++ {
                switch 32767 {
                case 128:
                    fmt.Println(v1, fi43)
                    var v44 bool = true
                    _ = v44
                case -9223372036854775808:
                    var v45 uint64 = 42
                    _ = v45
                case 32768:
                    var v46 int64 = 9223372036854775807
                    var v47 uint64 = 1
                    _ = v46
                    _ = v47
                }
                _ = fi43
            }
            for fi48 := 0; fi48 < 4; fi48++ {
                if len(v17) >= 2 {
                    v17 = v17[1:len(v17)]
                }
                if true {
                    var v49 uint16 = 960
                    _ = v49
                } else {
                    switch 619 {
                    case -9223372036854775808:
                        v3 -= v3
                    case -54:
                        var v50 uint32 = 4294967295
                        _ = v50
                    default:
                        fmt.Println(v21, v3, fi48)
                    }
                    fmt.Println(fi38, v21)
                }
                var v51 uint16 = 434
                _ = fi48
                _ = v51
            }
        }
        if (!(true && true)) {
            v3 = float64(float32(-2.25))
        } else {
            v1 -= v21
            v17 = append(v17, v4)
        }
        for fi52 := 0; fi52 < 10; fi52++ {
            v4 += v4
            v4 -= v4
            v2["a"] = 128
            _ = fi52
        }
        _ = fi38
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
