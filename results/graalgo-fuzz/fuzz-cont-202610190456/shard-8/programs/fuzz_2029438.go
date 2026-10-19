package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 int64) uint {
    var v4 bool = (1 <= 0)
    fmt.Println(v4, v4)
    _ = p0
    _ = v4
    return 65535
}
func genFn1(p0 uint8, p1 int, p2 uint32) bool {
    v5 := GenS1{F0: 100}
    var v6 int8 = 1
    var v7 uint = 18446744073709551615
    for fi8 := 0; fi8 < 7; fi8++ {
        switch (p1 % (((fi8 / ((int(p2)) | 1))) | 1)) {
        case -518:
            fmt.Println(fi8, p0, p2)
        default:
            v9 := GenS1{F0: v6}
            _ = v9
        }
        v10 := []int{(^(p1))}
        if (-2147483648 == 1) {
            fmt.Println(fi8, p1)
            v11 := GenS1{F0: (v6 & int8(v7))}
            v12 := GenS1{F0: (v6 >> (((v7 | genFn0(1))) & 63))}
            _ = v11
            _ = v12
        }
        _ = fi8
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return (true && false)
}

func main() {
    v1 := make(map[string]string)
    v2 := make(map[uint32]int)
    v3 := []string{"\n\t"}
    var v13 int32 = -2147483648
    v13 = v13
    _ = len(v3)
    v14 := GenS2{F0: 42}
    v13 -= (v13 | 2147483647)
    for fi15 := 0; fi15 < 16; fi15++ {
        if true {
            v16 := GenS1{F0: -28}
            for fi17 := 0; fi17 < 6; fi17++ {
                v18 := v2[2147483647]
                _ = len(v3)
                if true {
                    switch fi15 {
                    case -2147483648:
                        v18 -= v18
                    case -9223372036854775808:
                        v13 = v13
                    }
                    switch (fi15 - fi17) {
                    case -639:
                        v18 -= (^(v18))
                    case -26:
                        v19 := GenS2{F0: 625}
                        _ = v19
                    }
                }
                _ = fi17
                _ = v18
            }
            switch fi15 {
            case -580:
                v13 -= (v13 & v13)
            }
            _ = v16
        } else {
            fmt.Println(v13, fi15)
            v13 -= -1
        }
        var v20 float64 = 1e10
        fmt.Println(fi15)
        _ = fi15
        _ = v20
    }
    var v21 float64 = ((-2.25 - 1e10) * float64(1.5))
    delete(v2, 4294967295)
    switch -1 {
    case -416:
        if (false || false) {
            var v22 uint16 = 1
            if (65535 <= 799) {
                var v23 int32 = -2147483648
                switch 2147483647 {
                case 473:
                    fmt.Println(v13, v22)
                case 42:
                    switch 127 {
                    case 9223372036854775807:
                        var v24 int32 = v13
                        _ = v24
                    default:
                        v25 := GenS1{F0: -90}
                        _ = v25
                    }
                case -2147483648:
                    if true {
                        var v26 int16 = 32767
                        _ = v26
                    }
                default:
                    switch 42 {
                    case 226:
                        var v27 int = 0
                        _ = v27
                    case 815:
                        var v28 int = 252
                        _ = v28
                    case 127:
                        if true {
                            var v29 int64 = 9223372036854775807
                            _ = v29
                        } else {
                            v30 := v1["foo bar"]
                            _ = v30
                        }
                    }
                }
                _ = v23
            }
            _ = v22
        } else {
            var v31 uint16 = 42
            _ = v31
        }
    case 128:
        if (false && (!(("" != ("a" + ((("" + "a") + "foo bar") + ("hello" + ("foo bar" + "hello"))))) && (false || true)))) {
            var v32 uint8 = 255
            delete(v2, 325)
            if (true || (v13 < 0)) {
                var v33 int16 = 1000
                v33 -= v33
                _ = v33
            } else {
                if (true && false) {
                    for fi34 := 0; fi34 < 2; fi34++ {
                        v32 = (^(v32))
                        _ = fi34
                    }
                    var v35 bool = (0.0 == 3.14159)
                    _ = v35
                }
            }
            _ = v32
        }
    }
    for fi36 := 0; fi36 < 12; fi36++ {
        switch (fi36 + (fi36 * 32768)) {
        case 127:
            v13 = 0
        default:
            var v37 float32 = (float32(((v21 / (v21 / -2.25)) / -2.25)) - 1e10)
            _ = v37
        }
        v21 += (v21 / (v21 / ((3.14159 + 3.14159) + 1.5)))
        _ = fi36
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v21
    _ = v3
}
