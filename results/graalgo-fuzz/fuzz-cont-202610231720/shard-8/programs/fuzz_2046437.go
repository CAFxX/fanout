package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 int }
func genFn0(p0 int) int {
    p0 = p0
    var v5 float32 = -2.25
    _ = p0
    _ = v5
    return p0
}

func main() {
    v1 := []int{258}
    v2 := []string{}
    var v3 uint32 = 42
    v4 := make(map[uint32]uint32)
    switch -2147483648 {
    case 127:
        v3 += v3
    case 32768:
        if true {
            v3 += v3
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
    case 2147483647:
        var v6 uint64 = 18446744073709551615
        _ = v6
    default:
        for fi7 := 0; fi7 < 9; fi7++ {
            v3 += (^(v3))
            var v8 bool = (784 >= 32767)
            _ = fi7
            _ = v8
        }
    }
    switch 32768 {
    case 730:
        v9 := GenS1{F0: false, F1: -956}
        _ = v9
    }
    for fi10 := 0; fi10 < 1; fi10++ {
        var v11 uint32 = v3
        v3 = v3
        var v12 uint = 1
        _ = fi10
        _ = v11
        _ = v12
    }
    _ = len(v1)
    if true {
        fmt.Println(v3)
        fmt.Println(v3)
    } else {
        var v13 uint32 = v3
        if true {
            var v14 int64 = 911
            if ((!true) || (false || (-1 >= 0))) {
                switch genFn0(42) {
                case 42:
                    for fi15 := 0; fi15 < 12; fi15++ {
                        if true {
                            for fi16 := 0; fi16 < 6; fi16++ {
                                var v17 uint16 = 1
                                v3 = (v13 & v13)
                                v18 := GenS1{F0: false, F1: 1}
                                _ = fi16
                                _ = v17
                                _ = v18
                            }
                        }
                        var v19 int16 = 429
                        fmt.Println(fi15, v19, fi15)
                        _ = fi15
                        _ = v19
                    }
                case -9223372036854775808:
                    var v20 int = genFn0(9223372036854775807)
                    _ = v20
                default:
                    var v21 uint32 = v13
                    _ = v21
                }
            }
            for fi22 := 0; fi22 < 5; fi22++ {
                var v23 string = ("a" + "foo bar")
                v14 = 9223372036854775807
                _ = fi22
                _ = v23
            }
            _ = v14
        }
        for fi24 := 0; fi24 < 9; fi24++ {
            for fi25 := 0; fi25 < 13; fi25++ {
                var v26 uint32 = v13
                v26 = v13
                _ = fi25
                _ = v26
            }
            v27 := GenS1{F0: (float32(float64(1.5)) <= float32((-2.25 + 1.5))), F1: -1}
            switch -2147483648 {
            case 128:
                var v28 float32 = 1e10
                _ = v28
            case -9223372036854775808:
                for fi29 := 0; fi29 < 1; fi29++ {
                    if len(v2) > 0 {
                        v2[0] = "a"
                    }
                    switch (^(fi29)) {
                    case 0:
                        if (true || false) {
                            delete(v4, v3)
                        }
                    case 2147483647:
                        if true {
                            v13 = 42
                            v30 := GenS1{F0: false, F1: 1}
                            delete(v4, 337)
                            _ = v30
                        } else {
                            v13 = 267
                            var v31 float64 = float64(1.5)
                            if false {
                                if len(v2) > 0 {
                                    v2[0] = "a"
                                }
                                v31 = v31
                                var v32 uint8 = 128
                                _ = v32
                            }
                            _ = v31
                        }
                    }
                    v33 := GenS2{F0: (fi24 << ((1) & 63))}
                    _ = fi29
                    _ = v33
                }
            case 2147483647:
                delete(v4, v13)
            }
            _ = fi24
            _ = v27
        }
        _ = v13
    }
    for fi34 := 0; fi34 < 17; fi34++ {
        for fi35 := 0; fi35 < 13; fi35++ {
            switch (-(fi35)) {
            case 740:
                if true {
                    switch (fi35 | 4294967295) {
                    case 204:
                        v3 += v3
                    case -2147483648:
                        v3 = v3
                    default:
                        var v36 string = "a"
                        _ = v36
                    }
                    v3 += 276
                }
            case 127:
                var v37 string = ""
                _ = v37
            }
            _ = fi35
        }
        v38 := GenS2{F0: fi34}
        if (254 == 65535) {
            switch genFn0(fi34) {
            case -184:
                if (0 == 2) {
                    var v39 int8 = 37
                    v39 = 1
                    if true {
                        v3 = (^(v3))
                        _ = len(v4)
                    }
                    _ = v39
                }
            default:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
        }
        _ = fi34
        _ = v38
    }
    var v40 uint16 = 255
    fmt.Println(v3, v40)
    v41 := GenS1{F0: (0 >= 169), F1: 0}
    if false {
        v3 += 750
    }
    v42 := GenS1{F0: ((1e10 / (3.14159 * 3.14159)) == 1.5), F1: 649}
    var v43 int16 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
}
