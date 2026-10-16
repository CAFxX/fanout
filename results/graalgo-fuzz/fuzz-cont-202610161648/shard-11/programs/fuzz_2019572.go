package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int8; F2 int8 }
type GenS2 struct { F0 int16; F1 int16; F2 bool }
type GenS3 struct { F0 bool; F1 uint16 }
func genFn0(p0 int8) uint32 {
    var v7 uint8 = 128
    p0 = 0
    var v8 int8 = p0
    for fi9 := 0; fi9 < 15; fi9++ {
        v10 := []int{-1, (-(fi9)), (^(fi9))}
        if false {
            fmt.Println(v8, v7)
            _ = len(v10)
        } else {
            v10 = append(v10, (fi9 + 2147483647))
        }
        _ = fi9
        _ = v10
    }
    if (255 == 255) {
        switch 9223372036854775807 {
        case 127:
            var v11 uint8 = v7
            _ = v11
        case -5:
            v12 := GenS1{F0: 2, F1: (v8 / ((v8) | 1)), F2: 100}
            _ = v12
        }
        for fi13 := 0; fi13 < 18; fi13++ {
            for fi14 := 0; fi14 < 4; fi14++ {
                fmt.Println(p0, p0)
                switch fi14 {
                case 85:
                    if (v8 == p0) {
                        switch "foo bar" {
                        case "a":
                            p0 = -1
                        }
                        var v15 float64 = 0.0
                        _ = v15
                    }
                default:
                    v16 := []int{fi13, fi13}
                    _ = v16
                }
                _ = fi14
            }
            v17 := []int{}
            for fi18 := 0; fi18 < 12; fi18++ {
                if len(v17) > 0 {
                    v17[0] = int(v8)
                }
                v7 -= v7
                if len(v17) > 0 {
                    v17[0] = (fi13 % ((fi18) | 1))
                }
                _ = fi18
            }
            _ = fi13
            _ = v17
        }
    }
    _ = p0
    _ = v7
    _ = v8
    return 4294967295
}

func main() {
    var v1 float64 = 1e10
    v2 := []uint32{0}
    v3 := make(map[string]int)
    var v4 string = (("" + "hello") + "hello")
    var v5 int64 = 9223372036854775807
    var v6 int64 = (v5 << ((1) & 63))
    switch "" {
    case "\n\t":
        var v19 uint64 = 900
        _ = v19
    case "":
        if true {
            var v20 string = "foo bar"
            delete(v3, "hello")
            _ = v20
        }
    }
    if false {
        for fi21 := 0; fi21 < 10; fi21++ {
            switch fi21 {
            case 0:
                for fi22 := 0; fi22 < 5; fi22++ {
                    v6 += (^(v6))
                    v23 := GenS2{F0: 0, F1: -405, F2: (!true)}
                    _ = fi22
                    _ = v23
                }
            case -9223372036854775808:
                fmt.Println(v1, fi21)
            }
            v5 = v6
            var v24 int16 = -497
            _ = fi21
            _ = v24
        }
    } else {
        switch ((("a" + v4) + "\n\t") + "hello") {
        case "":
            var v25 int64 = (v5 / ((v6) | 1))
            _ = v25
        case "a":
            if false {
                fmt.Println(v4, v1, v1)
                if true {
                    var v26 uint64 = 2
                    v27 := GenS1{F0: 36, F1: 0, F2: -128}
                    _ = v26
                    _ = v27
                } else {
                    delete(v3, "foo bar")
                    v28 := GenS3{F0: (2147483647 <= 411), F1: 65535}
                    v4 = v4
                    _ = v28
                }
            } else {
                v5 = v5
            }
        case "hello":
            v29 := GenS3{F0: (true || false), F1: 2}
            _ = v29
        default:
            switch (v4 + "\n\t") {
            case "a":
                var v30 bool = (!(!true))
                _ = v30
            case "\n\t":
                fmt.Println(v6, v6, v4)
            case "foo bar":
                v31 := GenS3{F0: false, F1: 866}
                _ = v31
            }
        }
    }
    var v32 uint16 = 65535
    for fi33 := 0; fi33 < 7; fi33++ {
        v1 -= v1
        var v34 string = v4
        v35 := GenS3{F0: true, F1: (^(v32))}
        _ = fi33
        _ = v34
        _ = v35
    }
    var v36 uint16 = 482
    var v37 int64 = (v6 / (((v5 | (v5 / (((v6 / ((v5) | 1))) | 1)))) | 1))
    switch ("foo bar" + (((("a" + (v4 + v4)) + v4) + "") + v4)) {
    case "":
        if true {
            var v38 int8 = 127
            delete(v3, "hello")
            v32 += 944
            _ = v38
        } else {
            fmt.Println(v36, v1)
            v3["hello"] = -2147483648
        }
    case "\n\t":
        for fi39 := 0; fi39 < 5; fi39++ {
            var v40 int8 = -128
            fmt.Println(v32, v36, v4)
            v41 := GenS2{F0: 42, F1: -864, F2: (false && (2 < 1))}
            _ = fi39
            _ = v40
            _ = v41
        }
    case "foo bar":
        v32 += (v36 + 334)
    default:
        var v42 uint8 = 216
        _ = v42
    }
    v1 = (float64(1.5) - ((((-2.25 - ((0.0 / (float64(0.1) - v1)) / (v1 - ((v1 / v1) * v1)))) * v1) / v1) / (-2.25 / float64((float32((float64(0.1) - (v1 * float64(0.1)))) / 1e10)))))
    switch v4 {
    case "a":
        v36 += v32
    default:
        for fi43 := 0; fi43 < 14; fi43++ {
            v1 = (v1 * (v1 - v1))
            fmt.Println(v4, v6, v36)
            _ = fi43
        }
    }
    v32 += (v32 << ((1) & 63))
    _ = len(v2)
    fmt.Println(v32, v4)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
