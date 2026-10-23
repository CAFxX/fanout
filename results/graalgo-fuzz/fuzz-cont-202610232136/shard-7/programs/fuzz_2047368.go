package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 uint64 }
type GenS2 struct { F0 uint32; F1 uint; F2 uint8 }
func genFn0(p0 int8, p1 int64) string {
    fmt.Println(p0, p1)
    v6 := GenS1{F0: 2, F1: 127, F2: 2}
    v7 := []int{-1, -874}
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return ""
}

func main() {
    var v1 string = "hello"
    v2 := make(map[string]bool)
    var v3 uint = 255
    var v4 int64 = 1
    v5 := []int{}
    var v8 uint64 = 298
    fmt.Println(v4, v8, v1)
    fmt.Println(v4, v8, v1)
    switch 2147483647 {
    case -1:
        if true {
            v5 = append(v5, 128)
            var v9 int64 = (v4 ^ -844)
            v2[(v1 + "\n\t")] = false
            _ = v9
        } else {
            if (false && (!true)) {
                var v10 float32 = -2.25
                v4 = (v4 << ((v3) & 63))
                fmt.Println(v3)
                _ = v10
            }
        }
    case 128:
        _ = len(v2)
    }
    switch ("" + v1) {
    case "":
        v11 := GenS2{F0: 0, F1: (v3 - v3), F2: 75}
        _ = v11
    }
    switch v1 {
    case "foo bar":
        v12 := GenS1{F0: 65535, F1: 128, F2: v8}
        _ = v12
    }
    switch ((v1 + (v1 + "")) + v1) {
    case "\n\t":
        fmt.Println(v4, v4, v1)
    case "foo bar":
        fmt.Println(v4)
    case "a":
        var v13 bool = true
        _ = v13
    default:
        v8 = (v8 & 0)
    }
    for fi14 := 0; fi14 < 18; fi14++ {
        for fi15 := 0; fi15 < 7; fi15++ {
            if (100 < -1) {
                var v16 uint16 = 255
                fmt.Println(fi15, fi14)
                _ = v16
            } else {
                var v17 float32 = 3.14159
                _ = v17
            }
            switch ((v1 + v1) + v1) {
            case "\n\t":
                fmt.Println(v8, fi15, v4)
            case "foo bar":
                v18 := GenS2{F0: uint32(v8), F1: (v3 / (((v3 * v3)) | 1)), F2: 255}
                fmt.Println(v4)
                _ = v18
            case "a":
                if false {
                    _ = len(v5)
                    var v19 int8 = -108
                    _ = len(v5)
                    _ = v19
                } else {
                    if len(v5) > 0 {
                        v5[0] = (-(fi15))
                    }
                    var v20 int32 = 0
                    v20 = (v20 | v20)
                    var v21 uint64 = (^(v8))
                    _ = v20
                    _ = v21
                }
            default:
                fmt.Println(v4, fi15, fi14)
            }
            var v22 uint32 = 2
            _ = fi15
            _ = v22
        }
        _ = fi14
    }
    v4 += v4
    switch v1 {
    case "":
        var v23 uint32 = 65535
        _ = v23
    case "\n\t":
        _ = len(v5)
    case "a":
        if (1 <= -97) {
            for fi24 := 0; fi24 < 15; fi24++ {
                v4 = v4
                v8 -= (v8 | (v8 & v8))
                if false {
                    fmt.Println(fi24)
                } else {
                    switch (^(fi24)) {
                    case 9223372036854775807:
                        for fi25 := 0; fi25 < 5; fi25++ {
                            switch v1 {
                            case "":
                                var v26 string = "\n\t"
                                _ = v26
                            case "a":
                                v3 += 245
                            default:
                                fmt.Println(v1, fi24, v4)
                            }
                            var v27 uint = v3
                            _ = fi25
                            _ = v27
                        }
                    case 128:
                        v1 = "\n\t"
                    case 234:
                        v4 -= v4
                    }
                    var v28 int = int(v3)
                    _ = v28
                }
                _ = fi24
            }
            for fi29 := 0; fi29 < 7; fi29++ {
                fmt.Println(v4, v1, v8)
                _ = fi29
            }
        } else {
            fmt.Println(v3, v1)
            fmt.Println(v4, v1)
            switch 2147483647 {
            case -1:
                var v30 uint = 641
                _ = v30
            case 774:
                v5 = append(v5, 1)
            case -224:
                v3 = v3
            }
        }
    }
    if true {
        v4 = (-(v4))
        _ = len(v5)
        var v31 uint32 = 65535
        _ = v31
    } else {
        v2[(v1 + "a")] = (false && false)
    }
    var v32 float32 = 0.1
    if true {
        v33 := GenS2{F0: 42, F1: 0, F2: 137}
        delete(v2, v1)
        _ = v33
    }
    if true {
        var v34 int64 = (v4 - 42)
        switch (((((("\n\t" + v1) + "") + v1) + "foo bar") + (v1 + "hello")) + "\n\t") {
        case "a":
            v3 = v3
        case "foo bar":
            var v35 string = (v1 + ((v1 + "foo bar") + ""))
            _ = v35
        case "":
            var v36 uint = 1
            _ = v36
        }
        _ = v34
    }
    var v37 string = v1
    v32 -= -2.25
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v37
    _ = v4
    _ = v5
    _ = v8
}
