package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 int16 }
type GenS2 struct { F0 uint64; F1 int32; F2 int8 }
type GenS3 struct { F0 int64 }

func main() {
    var v1 uint16 = 0
    var v2 uint8 = 42
    var v3 string = ""
    v4 := []uint32{0, 2147483648, 42}
    v5 := make(map[int]uint32)
    var v6 uint32 = 2147483648
    fmt.Println(v1)
    for fi7 := 0; fi7 < 17; fi7++ {
        for fi8 := 0; fi8 < 15; fi8++ {
            v2 += v2
            v5[(fi8 - -515)] = (v6 - v6)
            var v9 int = 32768
            _ = fi8
            _ = v9
        }
        v2 += v2
        _ = fi7
    }
    var v10 float64 = (1e10 - (1.5 / (float64(3.14159) * -2.25)))
    v11 := GenS1{F0: -9223372036854775808, F1: ((v3 + v3) + v3), F2: -499}
    if true {
        var v12 int32 = 0
        fmt.Println(v12)
        _ = v12
    }
    for fi13 := 0; fi13 < 4; fi13++ {
        var v14 uint8 = v2
        if true {
            if true {
                fmt.Println(v10, v1, fi13)
                v5[(-(fi13))] = (v6 << ((42) & 63))
            } else {
                switch fi13 {
                case 68:
                    delete(v5, (-(fi13)))
                case 1:
                    if (true && (699 > 1)) {
                        v10 -= (3.14159 * (v10 + v10))
                    }
                }
                for fi15 := 0; fi15 < 14; fi15++ {
                    v16 := GenS2{F0: 0, F1: -14, F2: 20}
                    switch (fi15 & fi13) {
                    case 169:
                        var v17 int = 134
                        _ = v17
                    case 528:
                        if true {
                            v3 = "hello"
                        }
                        var v18 string = ((("foo bar" + v3) + v3) + "hello")
                        fmt.Println(fi13, v18)
                        var v19 uint64 = 2
                        v18 = (v18 + v18)
                        _ = v18
                        _ = v19
                    default:
                        v2 -= (v2 - 128)
                    }
                    _ = fi15
                    _ = v16
                }
                for fi20 := 0; fi20 < 8; fi20++ {
                    var v21 string = ("\n\t" + v3)
                    _ = fi20
                    _ = v21
                }
            }
            if (false || (42 >= 832)) {
                fmt.Println(v14, v1, fi13)
            }
        } else {
            if ((^(v6)) == 635) {
                for fi22 := 0; fi22 < 3; fi22++ {
                    fmt.Println(v6, v1, v10)
                    var v23 float32 = 0.1
                    _ = fi22
                    _ = v23
                }
                v3 = "\n\t"
                var v24 uint = 4294967295
                _ = v24
            } else {
                v2 -= v14
                v25 := GenS1{F0: -1, F1: ("foo bar" + v3), F2: 1}
                _ = v25
            }
            var v26 int8 = -88
            v6 -= (v6 & v6)
            _ = v26
        }
        if false {
            if (!(!((!true) && (-128 != -16)))) {
                var v27 int32 = 42
                _ = v27
            }
            var v28 int = fi13
            _ = v28
        }
        _ = fi13
        _ = v14
    }
    for fi29 := 0; fi29 < 17; fi29++ {
        var v30 uint32 = (v6 * (v6 >> ((637) & 63)))
        v31 := GenS1{F0: -473, F1: "\n\t", F2: 1000}
        for fi32 := 0; fi32 < 6; fi32++ {
            var v33 uint8 = v2
            var v34 int64 = -9223372036854775808
            v35 := GenS1{F0: (v34 / ((-101) | 1)), F1: ((((v3 + "") + ("hello" + (v3 + "hello"))) + ("hello" + ((v3 + (v3 + v3)) + "hello"))) + ((v3 + v3) + "foo bar")), F2: 0}
            _ = fi32
            _ = v33
            _ = v34
            _ = v35
        }
        _ = fi29
        _ = v30
        _ = v31
    }
    var v36 int16 = 32767
    fmt.Println(v3, v6, v3)
    v3 = v3
    v37 := GenS3{F0: 889}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
