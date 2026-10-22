package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int16; F1 int16; F2 bool }
type GenS3 struct { F0 bool; F1 bool }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 string; F1 uint }
type GenS6 struct { F0 int8; F1 uint16; F2 int }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 int16; F1 int8 }
func genFn0() string {
    v4 := GenS1{F0: -9223372036854775808}
    var v5 uint16 = 42
    for fi6 := 0; fi6 < 11; fi6++ {
        v7 := GenS2{F0: -1, F1: -32768, F2: false}
        _ = fi6
        _ = v7
    }
    var v8 bool = false
    var v9 float32 = float32(float64(1e10))
    _ = v4
    _ = v5
    _ = v8
    _ = v9
    return "a"
}
func genFn1(p0 int16, p1 bool) bool {
    v10 := []int{127, 32767, 42, 32768}
    var v11 int8 = -128
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return p1
}

func main() {
    v1 := []uint32{0, 4294967295}
    v2 := []string{"foo bar"}
    v3 := make(map[int]uint32)
    var v12 uint16 = 42
    if false {
        v12 = (v12 >> ((471) & 63))
    } else {
        v13 := GenS3{F0: ((42 < 40) || (!false)), F1: true}
        fmt.Println(v12)
        _ = len(v2)
        _ = v13
    }
    for fi14 := 0; fi14 < 16; fi14++ {
        fmt.Println(fi14)
        if true {
            var v15 int = fi14
            var v16 string = "foo bar"
            _ = v15
            _ = v16
        } else {
            v12 = (^(v12))
        }
        if (499 > 255) {
            v17 := GenS4{F0: 9223372036854775807}
            _ = v17
        } else {
            fmt.Println(fi14)
            var v18 uint32 = 651
            _ = v18
        }
        _ = fi14
    }
    if false {
        if (1000 >= 1000) {
            var v19 uint = 65535
            if true {
                switch 9223372036854775807 {
                case 9223372036854775807:
                    fmt.Println(v19, v12)
                case 256:
                    _ = len(v1)
                }
                var v20 int32 = -486
                v21 := GenS5{F0: "", F1: 405}
                _ = v20
                _ = v21
            }
            if true {
                switch (("hello" + ("a" + ("\n\t" + "a"))) + ((("\n\t" + "") + "foo bar") + (("\n\t" + "foo bar") + "hello"))) {
                case "hello":
                    switch -1 {
                    case 42:
                        var v22 uint16 = (v12 - 65535)
                        _ = v22
                    default:
                        var v23 uint = 433
                        _ = v23
                    }
                case "":
                    var v24 uint64 = 0
                    _ = v24
                }
                var v25 string = "hello"
                _ = v25
            }
            _ = v19
        } else {
            v12 -= (^(v12))
            v26 := GenS5{F0: "\n\t", F1: 0}
            _ = v26
        }
        fmt.Println(v12)
        if (!false) {
            v12 -= 269
        } else {
            v12 -= v12
            switch -112 {
            case 128:
                var v27 uint8 = 2
                _ = v27
            case 9223372036854775807:
                if true {
                    if false {
                        switch "\n\t" {
                        case "foo bar":
                            v12 = v12
                        case "a":
                            var v28 int = -9223372036854775808
                            _ = v28
                        }
                        if (!false) {
                            v12 = 42
                            v12 = (v12 % ((v12) | 1))
                            v29 := GenS2{F0: -1, F1: -1, F2: false}
                            _ = v29
                        }
                        v12 = (v12 + (v12 ^ v12))
                    }
                    var v30 int = 840
                    _ = v30
                } else {
                    v12 -= 255
                }
            default:
                v12 = (^(v12))
            }
        }
    }
    v31 := GenS2{F0: -450, F1: 666, F2: ((-421 <= 640) && false)}
    var v32 int32 = 152
    for fi33 := 0; fi33 < 6; fi33++ {
        if true {
            var v34 uint64 = 1
            _ = v34
        }
        for fi35 := 0; fi35 < 17; fi35++ {
            var v36 uint16 = 288
            var v37 int32 = (v32 * (v32 & v32))
            v37 += (v37 ^ (v37 | -1))
            _ = fi35
            _ = v36
            _ = v37
        }
        if ((3.14159 - ((float32(3.14159) / 1e10) + 1e10)) == -2.25) {
            v38 := GenS6{F0: -127, F1: 1, F2: fi33}
            _ = v38
        }
        _ = fi33
    }
    switch "foo bar" {
    case "foo bar":
        var v39 int = -1
        _ = v39
    }
    _ = len(v3)
    var v40 uint8 = 1
    if true {
        if (true || (78 == 0)) {
            var v41 int32 = v32
            if true {
                if true {
                    fmt.Println(v32)
                }
                _ = len(v1)
            } else {
                var v42 float32 = -2.25
                for fi43 := 0; fi43 < 7; fi43++ {
                    switch (fi43 + 2147483647) {
                    case 91:
                        var v44 bool = (v40 <= uint8(v12))
                        _ = v44
                    case 127:
                        v40 -= 127
                    case 32767:
                        var v45 uint64 = 818
                        _ = v45
                    default:
                        for fi46 := 0; fi46 < 8; fi46++ {
                            fmt.Println(v12, v40, v12)
                            _ = fi46
                        }
                    }
                    _ = fi43
                }
                v47 := GenS7{F0: -1}
                _ = v42
                _ = v47
            }
            _ = v41
        }
        switch 42 {
        case 32768:
            v12 = v12
        case 271:
            var v48 uint8 = 198
            _ = v48
        default:
            if true {
                var v49 int32 = (v32 - v32)
                v49 -= v32
                _ = v49
            } else {
                switch "hello" {
                case "a":
                    v50 := GenS8{F0: -481}
                    _ = v50
                default:
                    v51 := v3[-2147483648]
                    _ = v51
                }
            }
        }
        for fi52 := 0; fi52 < 19; fi52++ {
            switch "\n\t" {
            case "":
                fmt.Println(fi52, v32)
            case "\n\t":
                v53 := GenS9{F0: -32768, F1: 127}
                _ = v53
            case "foo bar":
                var v54 float32 = ((0.1 / 1.5) / 1e10)
                _ = v54
            }
            for fi55 := 0; fi55 < 18; fi55++ {
                var v56 uint16 = (v12 >> ((646) & 63))
                _ = fi55
                _ = v56
            }
            _ = fi52
        }
    }
    var v57 uint16 = (v12 << ((353) & 63))
    var v58 int16 = -32768
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v40:", v40)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v40
    _ = v57
    _ = v58
}
