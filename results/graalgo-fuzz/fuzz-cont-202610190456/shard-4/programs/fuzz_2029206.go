package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 uint64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int; F1 int16; F2 bool }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint8; F1 int64 }
type GenS6 struct { F0 int64; F1 int32 }
type GenS7 struct { F0 bool }
func genFn0(p0 int64, p1 int8) bool {
    v4 := GenS1{F0: "hello", F1: 2147483647, F2: 2}
    v5 := []int{-336, -689, 2147483647}
    var v6 float64 = 0.1
    v7 := GenS2{F0: ((p1 | p1) > p1)}
    var v8 int32 = -748
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return false
}

func main() {
    v1 := []uint32{}
    var v2 uint8 = 128
    var v3 uint16 = 802
    for fi9 := 0; fi9 < 17; fi9++ {
        v3 = 65535
        if (float64(1e10) != float64((1e10 + (0.0 * (((0.0 + 3.14159) * (-2.25 / 3.14159)) + 1e10))))) {
            if (true && (false || false)) {
                var v10 bool = true
                v10 = (-39 > fi9)
                _ = v10
            }
            v2 += v2
        } else {
            switch (("\n\t" + "hello") + "\n\t") {
            case "foo bar":
                v2 = v2
            }
            fmt.Println(v3, fi9)
            v11 := GenS1{F0: "", F1: 42, F2: 42}
            _ = v11
        }
        _ = fi9
    }
    var v12 uint64 = 4294967295
    fmt.Println(v3)
    if false {
        var v13 uint = 359
        v14 := GenS3{F0: 9223372036854775807, F1: -1, F2: ((v2 * 127) < v2)}
        var v15 uint16 = v3
        _ = v13
        _ = v14
        _ = v15
    } else {
        var v16 bool = (false || ((!(127 < -128)) && false))
        _ = len(v1)
        v17 := GenS3{F0: -1, F1: -979, F2: (v16 || (!v16))}
        _ = v16
        _ = v17
    }
    v3 -= 32767
    v1 = append(v1, 858)
    v12 = 2
    switch (("" + (("\n\t" + "a") + "\n\t")) + "\n\t") {
    case "foo bar":
        for fi18 := 0; fi18 < 7; fi18++ {
            var v19 uint8 = uint8(v3)
            fmt.Println(v19, v2, v12)
            var v20 uint = 65535
            _ = fi18
            _ = v19
            _ = v20
        }
    default:
        for fi21 := 0; fi21 < 17; fi21++ {
            v3 = (^(v3))
            _ = fi21
        }
    }
    if (((((1.5 / 0.1) > (1.5 * 0.1)) && ((v12 % ((2) | 1)) < (v12 % (((v12 * v12)) | 1)))) && (!false)) && false) {
        for fi22 := 0; fi22 < 4; fi22++ {
            if len(v1) > 0 {
                v1[0] = 65535
            }
            _ = fi22
        }
    }
    if true {
        if false {
            v1 = append(v1, 65535)
            var v23 bool = false
            _ = v23
        } else {
            if (("" + ("foo bar" + (("a" + "foo bar") + "hello"))) <= ("hello" + ("\n\t" + "foo bar"))) {
                v2 = (v2 >> ((2) & 63))
                v12 -= (v12 / ((2) | 1))
                v24 := GenS4{F0: 0}
                _ = v24
            } else {
                var v25 bool = (!(false || (!(1 != 45))))
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                var v26 bool = false
                _ = v25
                _ = v26
            }
            if (42 < 692) {
                for fi27 := 0; fi27 < 10; fi27++ {
                    for fi28 := 0; fi28 < 2; fi28++ {
                        var v29 bool = (("" + "") < (("foo bar" + (("a" + "\n\t") + "\n\t")) + "foo bar"))
                        if false {
                            v30 := GenS5{F0: (v2 / ((v2) | 1)), F1: 933}
                            v3 = (v3 + 65535)
                            _ = v30
                        } else {
                            if len(v1) > 0 {
                                v1[0] = 4294967295
                            }
                            var v31 int16 = -955
                            _ = v31
                        }
                        var v32 int = 9223372036854775807
                        _ = fi28
                        _ = v29
                        _ = v32
                    }
                    _ = fi27
                }
                v33 := GenS6{F0: -362, F1: 440}
                for fi34 := 0; fi34 < 4; fi34++ {
                    v12 -= (v12 >> ((1) & 63))
                    var v35 int32 = 956
                    fmt.Println(v35, v3)
                    _ = fi34
                    _ = v35
                }
                _ = v33
            } else {
                v36 := GenS2{F0: false}
                v37 := GenS5{F0: (v2 | 128), F1: -977}
                _ = v36
                _ = v37
            }
            v1 = append(v1, 4294967295)
        }
        for fi38 := 0; fi38 < 9; fi38++ {
            fmt.Println(v3)
            fmt.Println(v3, v12)
            if true {
                v3 = v3
            }
            _ = fi38
        }
        v3 = 2
    } else {
        for fi39 := 0; fi39 < 5; fi39++ {
            fmt.Println(v3, v12, v12)
            var v40 float64 = 3.14159
            v41 := GenS5{F0: (v2 * v2), F1: -63}
            _ = fi39
            _ = v40
            _ = v41
        }
        for fi42 := 0; fi42 < 5; fi42++ {
            v43 := GenS6{F0: -384, F1: -2147483648}
            _ = fi42
            _ = v43
        }
        var v44 int8 = -128
        _ = v44
    }
    var v45 float64 = 1.5
    switch -2147483648 {
    case 841:
        if (2 > 65535) {
            v46 := GenS7{F0: false}
            _ = len(v1)
            v3 = 19
            _ = v46
        } else {
            if true {
                var v47 int16 = -32768
                switch -71 {
                case -9223372036854775808:
                    var v48 uint8 = v2
                    _ = v48
                case 32768:
                    v1 = append(v1, 4294967295)
                }
                var v49 float32 = 1.5
                _ = v47
                _ = v49
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    }
    if len(v1) > 0 {
        v1[0] = 65535
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v45
}
