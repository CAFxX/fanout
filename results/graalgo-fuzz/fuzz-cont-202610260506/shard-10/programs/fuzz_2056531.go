package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int8; F1 uint; F2 uint64 }
type GenS4 struct { F0 int64 }
func genFn0(p0 uint64) uint64 {
    p0 = 0
    var v6 uint32 = 449
    v6 = v6
    switch 820 {
    case 4294967295:
        var v7 uint = 263
        _ = v7
    }
    _ = p0
    _ = v6
    return 1
}

func main() {
    v1 := make(map[int]bool)
    var v2 int64 = 1
    v3 := make(map[uint32]int)
    var v4 int16 = 321
    v5 := []string{"", "a", "\n\t"}
    v4 += (v4 & v4)
    var v8 int8 = 127
    if true {
        switch "" {
        case "":
            v2 = v2
        }
    } else {
        for fi9 := 0; fi9 < 9; fi9++ {
            v2 = v2
            for fi10 := 0; fi10 < 3; fi10++ {
                v11 := v3[1]
                v12 := GenS1{F0: 32767, F1: -813}
                _ = fi10
                _ = v11
                _ = v12
            }
            _ = fi9
        }
    }
    if (v4 > 32767) {
        v2 = v2
        fmt.Println(v2)
    }
    v13 := GenS1{F0: 258, F1: -2147483648}
    var v14 uint64 = 0
    fmt.Println(v8, v14)
    v4 = 0
    if (((v4 & v4) == v4) || true) {
        var v15 uint32 = 65535
        var v16 int16 = v4
        for fi17 := 0; fi17 < 16; fi17++ {
            for fi18 := 0; fi18 < 6; fi18++ {
                v19 := GenS1{F0: 32767, F1: 1}
                v20 := GenS1{F0: 32767, F1: 2147483647}
                var v21 bool = true
                _ = fi18
                _ = v19
                _ = v20
                _ = v21
            }
            for fi22 := 0; fi22 < 5; fi22++ {
                var v23 uint16 = 0
                v24 := GenS1{F0: (v23 % ((v23) | 1)), F1: -2147483648}
                _ = fi22
                _ = v23
                _ = v24
            }
            _ = fi17
        }
        _ = v15
        _ = v16
    } else {
        for fi25 := 0; fi25 < 15; fi25++ {
            if len(v5) > 0 {
                v5[0] = "hello"
            }
            fmt.Println(v8)
            _ = fi25
        }
        var v26 int32 = 1
        switch -531 {
        case 127:
            var v27 int16 = v4
            _ = v27
        case -9223372036854775808:
            v28 := GenS2{F0: 502}
            _ = v28
        default:
            _ = len(v1)
        }
        _ = v26
    }
    if (true && true) {
        if true {
            var v29 uint = 2
            v30 := GenS2{F0: v29}
            var v31 int64 = v2
            _ = v29
            _ = v30
            _ = v31
        } else {
            fmt.Println(v4, v4, v4)
            for fi32 := 0; fi32 < 17; fi32++ {
                switch fi32 {
                case -364:
                    switch fi32 {
                    case 128:
                        for fi33 := 0; fi33 < 11; fi33++ {
                            if len(v5) > 0 {
                                v5[0] = (("hello" + ("" + "")) + (("a" + "a") + ("" + ("hello" + ""))))
                            }
                            v34 := GenS3{F0: (^(v8)), F1: 255, F2: 0}
                            _ = fi33
                            _ = v34
                        }
                    case 32767:
                        switch "foo bar" {
                        case "a":
                            if false {
                                fmt.Println(fi32)
                            }
                        }
                    case 1:
                        v35 := GenS2{F0: 957}
                        _ = v35
                    }
                case -281:
                    for fi36 := 0; fi36 < 12; fi36++ {
                        if true {
                            v8 = 42
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                        }
                        var v37 float32 = (0.0 + ((float32((float64(1e10) + ((float64(1.5) * 1.5) - 1.5))) - (3.14159 - 1.5)) * 1e10))
                        v8 = (v8 % ((v8) | 1))
                        _ = fi36
                        _ = v37
                    }
                case -568:
                    fmt.Println(v8, v2, v14)
                default:
                    v14 -= v14
                }
                _ = fi32
            }
            if true {
                _ = len(v1)
                v38 := GenS2{F0: 18446744073709551615}
                switch "" {
                case "a":
                    var v39 uint64 = (v14 - 2)
                    _ = v39
                }
                _ = v38
            } else {
                v8 = (v8 & (v8 + 127))
                var v40 float32 = -2.25
                _ = v40
            }
        }
        var v41 float32 = 3.14159
        _ = v41
    }
    v42 := GenS4{F0: v2}
    _ = len(v5)
    var v43 string = (("foo bar" + "a") + "hello")
    fmt.Println(v2, v14)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v8
}
