package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 string }
type GenS2 struct { F0 int16; F1 uint16; F2 int64 }
type GenS3 struct { F0 uint; F1 uint64 }
type GenS4 struct { F0 int; F1 int }
type GenS5 struct { F0 uint16; F1 string; F2 uint8 }
type GenS6 struct { F0 uint8; F1 int64 }
type GenS7 struct { F0 int32; F1 uint16; F2 int64 }
type GenS8 struct { F0 uint }
func genFn0(p0 string, p1 int32, p2 uint64) string {
    v5 := GenS1{F0: p0, F1: 1, F2: (p0 + (p0 + p0))}
    p0 = ""
    fmt.Println(p0, p2)
    for fi6 := 0; fi6 < 8; fi6++ {
        fmt.Println(p0, p2, p0)
        _ = fi6
    }
    for fi7 := 0; fi7 < 9; fi7++ {
        switch -9223372036854775808 {
        case 32767:
            v8 := []int{(fi7 - fi7), fi7}
            _ = v8
        case -258:
            if true {
                for fi9 := 0; fi9 < 17; fi9++ {
                    p0 = "\n\t"
                    p1 = (p1 ^ -2147483648)
                    var v10 string = p0
                    _ = fi9
                    _ = v10
                }
                fmt.Println(p2, p1)
            } else {
                fmt.Println(p0, p1)
            }
        case -1:
            for fi11 := 0; fi11 < 2; fi11++ {
                var v12 int64 = 42
                if false {
                    var v13 float32 = 1e10
                    v13 = (v13 + (v13 / v13))
                    _ = v13
                }
                p2 += 2
                if true {
                    if false {
                        v14 := []int{}
                        var v15 uint16 = 2
                        v14 = append(v14, fi11)
                        if true {
                            if len(v14) >= 2 {
                                v14 = v14[1:len(v14)]
                            }
                        } else {
                            if (float64(((3.14159 + -2.25) - ((float32((0.0 / -2.25)) / ((3.14159 / 0.1) / 1.5)) + 0.0))) <= (0.0 + 0.0)) {
                                var v16 bool = false
                                p1 = 55
                                var v17 string = ("" + p0)
                                if len(v14) >= 2 {
                                    v14 = v14[1:len(v14)]
                                }
                                _ = v16
                                _ = v17
                            } else {
                                var v18 int8 = 100
                                switch (^(fi7)) {
                                case 128:
                                    var v19 int64 = (v12 * v12)
                                    _ = v19
                                case -1:
                                    if true {
                                        v12 += (v12 << ((18446744073709551615) & 63))
                                    }
                                }
                                v15 = (v15 << ((42) & 63))
                                fmt.Println(p0, p1, p2)
                                if (-1 == fi11) {
                                    switch 9223372036854775807 {
                                    case 317:
                                        if false {
                                            fmt.Println(fi7, p2, v15)
                                        }
                                    default:
                                        switch (fi7 | (fi7 ^ 128)) {
                                        case 554:
                                            if len(v14) >= 2 {
                                                v14 = v14[1:len(v14)]
                                            }
                                            var v20 string = "a"
                                            _ = v20
                                        default:
                                            v15 = 32767
                                        }
                                    }
                                } else {
                                    if false {
                                        var v21 uint64 = p2
                                        _ = v21
                                    }
                                }
                                _ = v18
                            }
                            var v22 uint16 = 255
                            v23 := GenS2{F0: 1000, F1: v22, F2: (^(v12))}
                            _ = v22
                            _ = v23
                        }
                        _ = v14
                        _ = v15
                    }
                    p0 = p0
                    fmt.Println(p2, v12)
                }
                var v24 uint16 = 32767
                _ = fi11
                _ = v12
                _ = v24
            }
        }
        var v25 int16 = 32767
        var v26 uint64 = (p2 << ((65535) & 63))
        _ = fi7
        _ = v25
        _ = v26
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return "a"
}
func genFn1(p0 int32, p1 uint8, p2 string) uint64 {
    p0 += p0
    var v27 int32 = p0
    v28 := GenS1{F0: "a", F1: p1, F2: p2}
    if (-840 >= -1) {
        for fi29 := 0; fi29 < 17; fi29++ {
            switch ((p2 + "\n\t") + (p2 + "")) {
            case "":
                v30 := GenS3{F0: 363, F1: 18446744073709551615}
                _ = v30
            default:
                fmt.Println(fi29)
            }
            p2 = p2
            var v31 int8 = 47
            _ = fi29
            _ = v31
        }
        p1 = p1
        var v32 uint = 18446744073709551615
        _ = v32
    }
    if true {
        p0 = v27
        fmt.Println(v27, p0, p1)
        var v33 float64 = (1e10 / 1.5)
        _ = v33
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v27
    _ = v28
    return 21
}

func main() {
    var v1 int16 = 32767
    v2 := make(map[string]int)
    v3 := make(map[int]int)
    var v4 uint32 = 211
    v1 += v1
    var v34 string = "foo bar"
    for fi35 := 0; fi35 < 12; fi35++ {
        var v36 uint16 = 42
        for fi37 := 0; fi37 < 10; fi37++ {
            v34 = ("" + v34)
            _ = fi37
        }
        _ = fi35
        _ = v36
    }
    var v38 uint32 = v4
    if false {
        v38 = v38
        var v39 int32 = 42
        v40 := GenS1{F0: (v34 + v34), F1: 128, F2: v34}
        _ = v39
        _ = v40
    } else {
        var v41 uint16 = 1
        _ = v41
    }
    fmt.Println(v38)
    v4 = v4
    fmt.Println(v34)
    if (true && true) {
        if false {
            for fi42 := 0; fi42 < 15; fi42++ {
                var v43 uint = 4294967295
                _ = fi42
                _ = v43
            }
            var v44 string = (v34 + v34)
            _ = v44
        }
        var v45 int8 = 0
        if true {
            v46 := GenS3{F0: 255, F1: 1}
            v47 := GenS4{F0: 42, F1: 9223372036854775807}
            v4 += (v4 - v4)
            _ = v46
            _ = v47
        } else {
            delete(v2, v34)
            v34 = v34
            v34 = (v34 + (v34 + "foo bar"))
        }
        _ = v45
    }
    v34 = v34
    switch v34 {
    case "":
        var v48 int32 = -997
        _ = v48
    }
    switch -198 {
    case 406:
        switch 665 {
        case 32768:
            var v49 int32 = 101
            _ = v49
        case -2147483648:
            var v50 uint8 = 21
            _ = v50
        default:
            v34 = (v34 + ("" + ((v34 + (((v34 + (v34 + v34)) + v34) + ("hello" + "foo bar"))) + ("a" + "foo bar"))))
        }
    case 127:
        if (3.14159 >= 3.14159) {
            if (!(false && true)) {
                for fi51 := 0; fi51 < 2; fi51++ {
                    if true {
                        var v52 int64 = 9223372036854775807
                        fmt.Println(v34)
                        _ = v52
                    }
                    for fi53 := 0; fi53 < 10; fi53++ {
                        var v54 int8 = 1
                        fmt.Println(v4)
                        if false {
                            v55 := GenS5{F0: 255, F1: "hello", F2: 0}
                            var v56 uint64 = 18446744073709551615
                            var v57 uint64 = 733
                            v58 := GenS6{F0: 0, F1: 9223372036854775807}
                            _ = v55
                            _ = v56
                            _ = v57
                            _ = v58
                        } else {
                            if (!(0 < 32767)) {
                                v59 := GenS3{F0: uint(v4), F1: genFn1(2147483647, 2, "hello")}
                                var v60 uint = 65535
                                v34 = v34
                                _ = v59
                                _ = v60
                            }
                        }
                        _ = fi53
                        _ = v54
                    }
                    if true {
                        var v61 int16 = 32767
                        v38 = 601
                        _ = v61
                    }
                    _ = fi51
                }
                v62 := GenS7{F0: 326, F1: 65535, F2: 42}
                _ = v62
            } else {
                fmt.Println(v34)
                v63 := []int{1, 128, 128}
                _ = v63
            }
        } else {
            var v64 float32 = 1.5
            for fi65 := 0; fi65 < 8; fi65++ {
                v38 = v4
                if true {
                    v1 = -32768
                    var v66 bool = false
                    v67 := GenS8{F0: 0}
                    _ = v66
                    _ = v67
                }
                var v68 float32 = 1e10
                _ = fi65
                _ = v68
            }
            _ = v64
        }
    case -300:
        v4 -= 42
    }
    v69 := GenS2{F0: (v1 - v1), F1: 42, F2: 0}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v38
    _ = v4
    _ = v69
}
