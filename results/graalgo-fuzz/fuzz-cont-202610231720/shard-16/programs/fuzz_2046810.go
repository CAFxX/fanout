package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 int16; F1 int }
type GenS3 struct { F0 bool; F1 string; F2 int32 }

func main() {
    var v1 int8 = 100
    v2 := []uint32{2147483648, 900, 65535}
    var v3 int32 = -1
    var v4 int16 = -750
    v5 := make(map[uint32]string)
    if ((-(v4)) <= (^(v4))) {
        var v6 uint8 = 136
        _ = v6
    } else {
        var v7 uint16 = 65535
        v8 := GenS1{F0: int16(v7), F1: 42}
        if (false && false) {
            switch 0 {
            case 128:
                v4 -= (v4 * 42)
            case -9223372036854775808:
                v9 := GenS1{F0: v4, F1: uint(v7)}
                _ = v9
            }
        }
        _ = v7
        _ = v8
    }
    v10 := GenS2{F0: v4, F1: -371}
    var v11 int = 0
    for fi12 := 0; fi12 < 18; fi12++ {
        v3 = v3
        _ = fi12
    }
    for fi13 := 0; fi13 < 10; fi13++ {
        for fi14 := 0; fi14 < 18; fi14++ {
            if (true || (true || (3.14159 <= 3.14159))) {
                var v15 uint32 = 859
                v11 = (v11 >> ((934) & 63))
                v4 = (v4 & v4)
                _ = v15
            } else {
                v3 = v3
                if (65535 >= 0) {
                    v5[4294967295] = "hello"
                    v1 += (v1 << ((798) & 63))
                } else {
                    fmt.Println(fi14, v11)
                    var v16 uint16 = 91
                    fmt.Println(v11, v3)
                    _ = v16
                }
            }
            v11 = fi14
            _ = fi14
        }
        fmt.Println(v11, fi13, fi13)
        _ = fi13
    }
    v17 := GenS1{F0: (v4 + (v4 & (v4 >> ((0) & 63)))), F1: 65535}
    fmt.Println(v4, v1, v1)
    v11 = -95
    v18 := v5[0]
    v3 = v3
    delete(v5, 42)
    var v19 uint32 = 2147483648
    var v20 uint = 2
    fmt.Println(v4)
    switch v11 {
    case 392:
        if (42 > 601) {
            var v21 int = (-(v11))
            fmt.Println(v1)
            _ = v21
        } else {
            v4 -= v4
            var v22 string = ""
            _ = v22
        }
    case -1:
        for fi23 := 0; fi23 < 4; fi23++ {
            if (v20 >= 255) {
                if len(v2) > 0 {
                    v2[0] = v19
                }
                if false {
                    if true {
                        delete(v5, v19)
                    } else {
                        v4 += 1
                        var v24 uint16 = 32767
                        _ = v24
                    }
                    v25 := GenS2{F0: v4, F1: 42}
                    _ = v25
                }
                var v26 float32 = 0.0
                _ = v26
            }
            _ = fi23
        }
    }
    switch v11 {
    case -9223372036854775808:
        switch (v11 - (v11 ^ 128)) {
        case 32767:
            if false {
                var v27 uint64 = 4294967295
                var v28 int = (v11 & (v11 & (v11 << ((v20) & 63))))
                _ = v27
                _ = v28
            } else {
                v29 := GenS1{F0: v4, F1: (^(v20))}
                _ = v29
            }
        case -9223372036854775808:
            fmt.Println(v1)
        case 128:
            var v30 int8 = (-(v1))
            _ = v30
        }
    case -622:
        for fi31 := 0; fi31 < 20; fi31++ {
            fmt.Println(v18)
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi31
        }
    case -1:
        for fi32 := 0; fi32 < 2; fi32++ {
            for fi33 := 0; fi33 < 13; fi33++ {
                if (255 < 42) {
                    v34 := GenS1{F0: v4, F1: 36}
                    _ = v34
                }
                if ((3.14159 + 0.0) != (0.0 / 0.1)) {
                    var v35 string = (v18 + "hello")
                    _ = v35
                }
                v3 = (v3 + v3)
                _ = fi33
            }
            switch v11 {
            case -1:
                if ((!(!false)) && true) {
                    v20 -= (^(v20))
                } else {
                    v36 := GenS2{F0: v4, F1: -1}
                    var v37 uint64 = 42
                    _ = v36
                    _ = v37
                }
            case 4294967295:
                for fi38 := 0; fi38 < 5; fi38++ {
                    var v39 string = "foo bar"
                    _ = fi38
                    _ = v39
                }
            case -301:
                v40 := GenS2{F0: (v4 & 1), F1: (v11 | (v11 << ((v20) & 63)))}
                _ = v40
            default:
                v41 := GenS3{F0: true, F1: ((v18 + "hello") + ("" + (v18 + ("" + "foo bar")))), F2: (v3 >> ((v20) & 63))}
                _ = v41
            }
            _ = fi32
        }
    default:
        var v42 int16 = v4
        _ = v42
    }
    var v43 string = (("foo bar" + "\n\t") + v18)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v11
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v43
    _ = v5
}
