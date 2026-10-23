package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 uint16; F1 uint8 }
type GenS5 struct { F0 int64; F1 uint16 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 string; F1 string }
type GenS8 struct { F0 uint32; F1 string; F2 int8 }
type GenS9 struct { F0 uint32; F1 uint64 }
func genFn0(p0 int, p1 int8, p2 bool) uint16 {
    var v7 int16 = -1
    for fi8 := 0; fi8 < 11; fi8++ {
        v9 := GenS1{F0: 4294967295}
        _ = fi8
        _ = v9
    }
    var v10 float64 = 1e10
    for fi11 := 0; fi11 < 7; fi11++ {
        for fi12 := 0; fi12 < 8; fi12++ {
            if p2 {
                if false {
                    var v13 uint16 = 99
                    v14 := []int{fi11, 127, (-(fi11))}
                    fmt.Println(v13, p1)
                    _ = v13
                    _ = v14
                } else {
                    v15 := GenS1{F0: (^(fi12))}
                    switch (p0 + (fi11 + (p0 | (fi12 >> ((1) & 63))))) {
                    case 9223372036854775807:
                        if false {
                            v16 := GenS2{F0: 968}
                            if p2 {
                                if p2 {
                                    if (p2 || p2) {
                                        v17 := GenS1{F0: p0}
                                        p1 = p1
                                        _ = v17
                                    } else {
                                        p1 -= 0
                                    }
                                }
                            } else {
                                var v18 uint64 = 755
                                v19 := GenS3{F0: 0}
                                _ = v18
                                _ = v19
                            }
                            _ = v16
                        }
                    case -1:
                        var v20 bool = p2
                        _ = v20
                    case 1:
                        fmt.Println(p0)
                    }
                    _ = v15
                }
            }
            var v21 uint = 65535
            _ = fi12
            _ = v21
        }
        v22 := []int{(p0 & p0), p0, 4294967295, p0}
        _ = fi11
        _ = v22
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v7
    return 255
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 int = 127
    var v3 int32 = 0
    var v4 uint32 = 338
    v5 := []int{(v2 << ((2) & 63))}
    var v6 uint8 = 1
    if true {
        if ((float32(0.0) / 1.0) > -2.25) {
            if (1000 > -1) {
                switch ("a" + (("foo bar" + "\n\t") + "")) {
                case "\n\t":
                    fmt.Println(v4, v2)
                case "foo bar":
                    v4 -= (v4 | 65535)
                }
                for fi23 := 0; fi23 < 20; fi23++ {
                    fmt.Println(fi23, v2, fi23)
                    _ = fi23
                }
            }
            for fi24 := 0; fi24 < 8; fi24++ {
                for fi25 := 0; fi25 < 5; fi25++ {
                    if true {
                        var v26 uint64 = 1
                        _ = v26
                    } else {
                        _ = len(v1)
                        var v27 uint32 = v4
                        if (!(!false)) {
                            fmt.Println(v3, v27)
                            var v28 uint8 = (v6 + 0)
                            v27 = (v27 << ((2) & 63))
                            _ = v28
                        }
                        _ = v27
                    }
                    delete(v1, (v4 & (v4 - 2)))
                    v29 := GenS1{F0: (fi25 - 42)}
                    if true {
                        if (!(((0.0 + 1.5) / (1.5 + (0.1 * 0.1))) <= (3.14159 - 3.14159))) {
                            v5 = append(v5, (v2 ^ fi24))
                            _ = len(v1)
                        }
                    } else {
                        v6 -= v6
                    }
                    _ = fi25
                    _ = v29
                }
                v2 += fi24
                _ = fi24
            }
        } else {
            for fi30 := 0; fi30 < 17; fi30++ {
                v2 = v2
                var v31 int64 = 1
                _ = fi30
                _ = v31
            }
            fmt.Println(v4, v6, v4)
            v32 := GenS2{F0: 239}
            _ = v32
        }
        switch "" {
        case "\n\t":
            v6 = v6
        case "a":
            v4 = 42
        default:
            for fi33 := 0; fi33 < 19; fi33++ {
                if (true || (!true)) {
                    v6 = v6
                    fmt.Println(fi33)
                } else {
                    if false {
                        _ = len(v5)
                        v34 := v1[v4]
                        _ = v34
                    }
                    v35 := GenS4{F0: 42, F1: v6}
                    var v36 int32 = v3
                    _ = v35
                    _ = v36
                }
                _ = fi33
            }
        }
        if ("a" == ((("hello" + "hello") + "foo bar") + "\n\t")) {
            switch v2 {
            case -737:
                if (989 < (v4 % ((1) | 1))) {
                    var v37 uint = 1
                    _ = v37
                } else {
                    _ = len(v1)
                    v38 := GenS5{F0: 0, F1: 0}
                    _ = v38
                }
            }
        } else {
            delete(v1, v4)
            switch "hello" {
            case "foo bar":
                var v39 int = 0
                _ = v39
            case "\n\t":
                fmt.Println(v6, v6, v2)
            }
            v4 -= v4
        }
    }
    var v40 float32 = (0.0 / -2.25)
    if true {
        if (!(42 < 972)) {
            v3 += v3
            if (42 == 42) {
                v41 := GenS6{F0: 4294967295}
                var v42 int32 = (v3 ^ (v3 / ((0) | 1)))
                _ = v41
                _ = v42
            }
        }
    }
    var v43 uint32 = 7
    v44 := GenS7{F0: ("foo bar" + ("hello" + (("hello" + "") + "hello"))), F1: "hello"}
    v2 -= -2147483648
    if true {
        var v45 int8 = -5
        var v46 float64 = (-2.25 / (-2.25 + ((1.5 * 1.5) + (((((3.14159 - 1e10) / -2.25) + 0.0) + 1e10) + (1.5 * (0.0 + (1e10 + (((float64((-2.25 + v40)) + 0.0) * -2.25) * float64(float32(((0.0 + (float64(0.0) + (float64(v40) - -2.25))) - 1e10)))))))))))
        _ = v45
        _ = v46
    } else {
        var v47 uint32 = v43
        fmt.Println(v4, v4)
        var v48 uint8 = v6
        _ = v47
        _ = v48
    }
    fmt.Println(v2)
    if len(v5) > 0 {
        v5[0] = v2
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    for fi49 := 0; fi49 < 14; fi49++ {
        var v50 int16 = 1
        _ = fi49
        _ = v50
    }
    if (false && false) {
        v51 := GenS8{F0: v43, F1: "a", F2: -1}
        v52 := GenS9{F0: (v4 * v4), F1: 42}
        _ = v51
        _ = v52
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v43
    _ = v44
    _ = v5
    _ = v6
}
