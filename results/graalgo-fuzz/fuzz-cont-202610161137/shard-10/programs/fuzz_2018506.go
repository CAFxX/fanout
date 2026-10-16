package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 int16; F1 int8 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 bool; F1 int8; F2 int8 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 uint }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[uint32]string)
    var v3 int16 = 885
    var v4 int = 4294967295
    var v5 float64 = (((1e10 / 3.14159) / 3.14159) - 3.14159)
    for fi6 := 0; fi6 < 12; fi6++ {
        if (!((("\n\t" + "a") + ("\n\t" + "\n\t")) == "foo bar")) {
            for fi7 := 0; fi7 < 18; fi7++ {
                v1[(-(fi7))] = 571
                _ = fi7
            }
        }
        var v8 uint16 = 227
        _ = fi6
        _ = v8
    }
    for fi9 := 0; fi9 < 1; fi9++ {
        v10 := GenS1{F0: 2, F1: 4294967295}
        fmt.Println(v4, v3)
        _ = fi9
        _ = v10
    }
    for fi11 := 0; fi11 < 6; fi11++ {
        for fi12 := 0; fi12 < 5; fi12++ {
            fmt.Println(fi12, v5)
            v5 += v5
            v13 := GenS2{F0: (v3 | 1), F1: -9}
            var v14 int = fi12
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = fi11
    }
    fmt.Println(v4)
    for fi15 := 0; fi15 < 19; fi15++ {
        fmt.Println(fi15, fi15)
        v16 := []int{0, v4, (fi15 & (fi15 ^ v4)), 42}
        _ = fi15
        _ = v16
    }
    v17 := []int{(v4 << ((4294967295) & 63)), (v4 | 32768), (v4 >> ((65535) & 63))}
    if true {
        for fi18 := 0; fi18 < 18; fi18++ {
            _ = len(v2)
            v19 := GenS1{F0: 65535, F1: 666}
            _ = fi18
            _ = v19
        }
    }
    if false {
        var v20 float32 = 1e10
        v5 = v5
        if (false || (1 == 9223372036854775807)) {
            v5 += 0.0
        }
        _ = v20
    }
    if ("" >= ("\n\t" + "a")) {
        var v21 uint = 25
        _ = v21
    } else {
        v3 += -32768
        if len(v17) > 0 {
            v17[0] = (v4 * (v4 % ((v4) | 1)))
        }
    }
    v2[2] = ("hello" + ("hello" + "\n\t"))
    if len(v17) > 0 {
        v17[0] = v4
    }
    if (366 <= 965) {
        for fi22 := 0; fi22 < 20; fi22++ {
            switch (-(fi22)) {
            case 127:
                v3 -= (v3 & (v3 & 32767))
            default:
                for fi23 := 0; fi23 < 9; fi23++ {
                    v24 := GenS3{F0: (65535 != 2)}
                    var v25 string = "foo bar"
                    _ = fi23
                    _ = v24
                    _ = v25
                }
            }
            if (!false) {
                for fi26 := 0; fi26 < 11; fi26++ {
                    v17 = append(v17, 32768)
                    if false {
                        _ = len(v17)
                        v5 = ((float64(0.1) * v5) - ((float64(((-2.25 * 1e10) / 1.0)) / (-2.25 / float64(1.5))) - v5))
                    }
                    _ = fi26
                }
                var v27 bool = true
                if len(v17) >= 2 {
                    v17 = v17[1:len(v17)]
                }
                _ = v27
            } else {
                switch "hello" {
                case "":
                    v28 := GenS4{F0: (false || ((1e10 * 0.1) > -2.25)), F1: -16, F2: 0}
                    _ = v28
                }
                v2[2147483648] = ("foo bar" + "a")
                v29 := GenS5{F0: 854}
                _ = v29
            }
            _ = fi22
        }
        v4 += 32767
        v4 = int(v3)
    } else {
        var v30 int16 = (v3 << ((858) & 63))
        if (0 != 4294967295) {
            if (!false) {
                var v31 uint = 42
                for fi32 := 0; fi32 < 14; fi32++ {
                    v4 = v4
                    if len(v17) > 0 {
                        v17[0] = v4
                    }
                    var v33 bool = (!false)
                    _ = fi32
                    _ = v33
                }
                _ = v31
            } else {
                var v34 uint = 4294967295
                v30 += -581
                _ = v34
            }
            if ((v3 % (((v3 * v3)) | 1)) >= (-(v30))) {
                v35 := GenS1{F0: 42, F1: 1}
                _ = len(v2)
                _ = v35
            }
            var v36 uint8 = 2
            _ = v36
        }
        for fi37 := 0; fi37 < 9; fi37++ {
            v38 := GenS6{F0: 255}
            v30 = v30
            _ = fi37
            _ = v38
        }
        _ = v30
    }
    var v39 uint64 = 18446744073709551615
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v17:", len(v17), v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v5
}
