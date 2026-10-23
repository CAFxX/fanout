package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64; F1 int16; F2 uint16 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint16; F1 int8 }
func genFn0(p0 int32, p1 uint) int16 {
    var v7 float64 = 0.1
    var v8 uint16 = 332
    var v9 int64 = -98
    v10 := []int{0, 931, 9223372036854775807}
    v11 := GenS1{F0: 1}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v7
    _ = v8
    _ = v9
    return 42
}

func main() {
    var v1 string = "hello"
    var v2 uint64 = 726
    var v3 uint32 = 65535
    v4 := make(map[uint32]bool)
    v5 := []string{(("\n\t" + v1) + ((v1 + "foo bar") + "foo bar")), ("a" + v1), "\n\t"}
    var v6 bool = true
    if v6 {
        var v12 int8 = -1
        v13 := GenS1{F0: 168}
        switch (("hello" + v1) + v1) {
        case "hello":
            v14 := GenS1{F0: 0}
            _ = v14
        default:
            switch v1 {
            case "":
                var v15 float32 = 0.1
                _ = v15
            }
        }
        _ = v12
        _ = v13
    } else {
        if v6 {
            var v16 int = 32768
            _ = v16
        } else {
            if v6 {
                var v17 bool = (1000 >= 42)
                if v6 {
                    if v6 {
                        v1 = v1
                    }
                    if (146 <= v3) {
                        if (false && v6) {
                            v18 := GenS1{F0: 255}
                            var v19 uint = 255
                            _ = v18
                            _ = v19
                        } else {
                            v3 = (v3 & (v3 / ((2147483647) | 1)))
                            switch -694 {
                            case 42:
                                var v20 uint32 = v3
                                _ = v20
                            case 0:
                                v17 = false
                            default:
                                for fi21 := 0; fi21 < 13; fi21++ {
                                    switch 445 {
                                    case 127:
                                        fmt.Println(v6, v6)
                                    case 51:
                                        v6 = (0.1 == ((0.0 + -2.25) + 0.1))
                                    }
                                    var v22 float32 = float32((-2.25 - float64(-2.25)))
                                    _ = fi21
                                    _ = v22
                                }
                            }
                            if v6 {
                                v23 := GenS1{F0: 40}
                                _ = v23
                            }
                        }
                    }
                    fmt.Println(v2, v6)
                }
                v5 = append(v5, v1)
                _ = v17
            } else {
                v24 := v4[(v3 | (v3 - (^(v3))))]
                switch (v1 + "a") {
                case "":
                    v25 := GenS2{F0: 255}
                    _ = v25
                case "foo bar":
                    fmt.Println(v1, v24)
                }
                _ = v24
            }
            v26 := GenS3{F0: -904, F1: 0, F2: 2}
            _ = v26
        }
    }
    var v27 int16 = -298
    var v28 uint = 0
    switch 127 {
    case -467:
        v29 := GenS1{F0: 1}
        _ = v29
    case 128:
        var v30 float64 = 0.0
        _ = v30
    case 683:
        if v6 {
            v31 := GenS1{F0: 255}
            var v32 float32 = 3.14159
            _ = v31
            _ = v32
        } else {
            for fi33 := 0; fi33 < 16; fi33++ {
                fmt.Println(v27, v27, v3)
                _ = fi33
            }
        }
    default:
        if true {
            var v34 int = -835
            v5 = append(v5, ("foo bar" + ("foo bar" + v1)))
            var v35 uint64 = (v2 >> (((v28 / ((v28) | 1))) & 63))
            _ = v34
            _ = v35
        }
    }
    switch 42 {
    case 32768:
        v36 := GenS3{F0: -9223372036854775808, F1: 0, F2: 0}
        _ = v36
    case 127:
        var v37 uint64 = 42
        _ = v37
    }
    delete(v4, v3)
    if len(v5) > 0 {
        v5[0] = v1
    }
    if (197 != 32767) {
        for fi38 := 0; fi38 < 6; fi38++ {
            v39 := GenS4{F0: -6}
            _ = fi38
            _ = v39
        }
        v3 -= 1
        _ = len(v4)
    }
    v40 := GenS2{F0: 53}
    for fi41 := 0; fi41 < 2; fi41++ {
        var v42 float32 = float32((1e10 / ((-2.25 - 3.14159) + 0.0)))
        for fi43 := 0; fi43 < 15; fi43++ {
            v44 := GenS5{F0: 0, F1: -128}
            _ = fi43
            _ = v44
        }
        _ = fi41
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v40
    _ = v5
    _ = v6
}
