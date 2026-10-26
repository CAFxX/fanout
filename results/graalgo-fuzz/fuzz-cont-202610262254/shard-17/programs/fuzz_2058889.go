package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint8; F2 uint }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int; F1 uint; F2 int32 }
type GenS4 struct { F0 uint32; F1 int32; F2 uint }
func genFn0() string {
    v7 := GenS1{F0: 42, F1: 127, F2: 509}
    var v8 int16 = -404
    if false {
        switch 128 {
        case 32767:
            fmt.Println(v8)
        case 32768:
            var v9 int8 = 127
            _ = v9
        case 4294967295:
            var v10 bool = true
            _ = v10
        }
        v8 = v8
    }
    _ = v7
    _ = v8
    return "\n\t"
}
func genFn1(p0 uint64) uint {
    switch ("a" + ("\n\t" + "\n\t")) {
    case "hello":
        if (42 == 255) {
            for fi11 := 0; fi11 < 19; fi11++ {
                v12 := []int{127, (^(fi11)), (-(fi11))}
                _ = fi11
                _ = v12
            }
        }
    }
    switch (("foo bar" + (("a" + "foo bar") + "hello")) + "a") {
    case "\n\t":
        v13 := GenS1{F0: 0, F1: 80, F2: 2}
        _ = v13
    }
    _ = p0
    return 0
}

func main() {
    var v1 uint32 = 728
    v2 := make(map[string]uint32)
    var v3 uint16 = 255
    v4 := make(map[uint32]uint32)
    var v5 uint8 = 127
    v6 := make(map[string]int)
    v14 := []int{4294967295, 32768, 127, -2147483648}
    v5 -= 100
    v15 := GenS2{F0: false}
    for fi16 := 0; fi16 < 14; fi16++ {
        delete(v2, ((("\n\t" + "\n\t") + "\n\t") + "\n\t"))
        if false {
            if ((!true) && (!true)) {
                fmt.Println(v5, v3)
            } else {
                switch fi16 {
                case 0:
                    if false {
                        for fi17 := 0; fi17 < 14; fi17++ {
                            switch 42 {
                            case -1:
                                fmt.Println(fi16)
                            case 9223372036854775807:
                                fmt.Println(v5, v1)
                                v1 += 2
                                v3 += v3
                            case -661:
                                var v18 float32 = 3.14159
                                _ = v18
                            }
                            _ = fi17
                        }
                    }
                case -1:
                    switch (fi16 << ((uint(v5)) & 63)) {
                    case 42:
                        v5 = v5
                    default:
                        fmt.Println(v1, v1, v5)
                    }
                case 127:
                    for fi19 := 0; fi19 < 8; fi19++ {
                        var v20 int32 = 0
                        v3 -= (^(v3))
                        _ = len(v14)
                        _ = fi19
                        _ = v20
                    }
                }
                v3 += 562
                var v21 uint8 = 128
                _ = v21
            }
            var v22 int16 = int16(v5)
            var v23 int8 = -128
            _ = v22
            _ = v23
        } else {
            if false {
                v5 -= v5
            } else {
                var v24 int8 = -1
                var v25 uint = 65535
                var v26 uint8 = v5
                _ = v24
                _ = v25
                _ = v26
            }
            if (31 != 0) {
                fmt.Println(v5)
                var v27 int64 = 42
                _ = v27
            } else {
                for fi28 := 0; fi28 < 12; fi28++ {
                    if (2147483647 >= 0) {
                        if ((v5 % (((v5 - (v5 | (v5 - v5)))) | 1)) > v5) {
                            var v29 int = -2147483648
                            fmt.Println(v29)
                            if (((((0.0 / 3.14159) + 0.1) - float64(1e10)) / float64(3.14159)) > 1e10) {
                                fmt.Println(v5, v5, fi16)
                                var v30 string = "\n\t"
                                fmt.Println(v1, v1)
                                var v31 uint64 = 857
                                fmt.Println(v30, fi28, v3)
                                v31 -= 771
                                var v32 uint16 = 2
                                _ = v30
                                _ = v31
                                _ = v32
                            } else {
                                v1 = (^(v1))
                            }
                            var v33 int = fi16
                            var v34 float32 = ((1.5 * 1e10) - 3.14159)
                            _ = v29
                            _ = v33
                            _ = v34
                        }
                        var v35 int8 = 39
                        _ = v35
                    }
                    v2["foo bar"] = (^(v1))
                    _ = fi28
                }
                v36 := GenS2{F0: (fi16 != -1)}
                _ = v36
            }
            for fi37 := 0; fi37 < 20; fi37++ {
                if (!(false && (!(345 <= 1)))) {
                    v1 += (v1 / (((v1 | v1)) | 1))
                    v38 := GenS3{F0: fi37, F1: 665, F2: 459}
                    delete(v4, (v1 & v1))
                    _ = v38
                }
                _ = fi37
            }
        }
        for fi39 := 0; fi39 < 16; fi39++ {
            v6["\n\t"] = (fi39 - (fi39 % ((fi39) | 1)))
            _ = fi39
        }
        _ = fi16
    }
    v1 = v1
    fmt.Println(v3, v5)
    switch 42 {
    case 0:
        v5 = 0
    default:
        var v40 uint64 = 1
        _ = v40
    }
    var v41 int8 = -128
    v3 -= (v3 >> ((65535) & 63))
    v41 = (v41 - (^(v41)))
    _ = len(v4)
    for fi42 := 0; fi42 < 6; fi42++ {
        v43 := GenS4{F0: v1, F1: -1, F2: 248}
        _ = fi42
        _ = v43
    }
    v3 -= v3
    v44 := GenS4{F0: (v1 >> ((159) & 63)), F1: 247, F2: 1}
    v41 += v41
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v41
    _ = v44
    _ = v5
    _ = v6
}
