package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint8 }
type GenS2 struct { F0 int16; F1 uint32; F2 uint32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint8 }
func genFn0() int32 {
    v5 := []int{-815, 128, -317, 127}
    var v6 uint64 = 39
    if ((-1 > -32768) || false) {
        fmt.Println(v6)
    } else {
        v5 = append(v5, 42)
        if len(v5) > 0 {
            v5[0] = 1
        }
    }
    var v7 uint16 = 281
    _ = v5
    _ = v6
    _ = v7
    return 42
}
func genFn1(p0 int32, p1 string, p2 int32) int32 {
    var v8 uint64 = 20
    v9 := []int{128, -626, -2147483648, 330}
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return (^(p2))
}

func main() {
    var v1 bool = false
    var v2 int = 9223372036854775807
    v3 := make(map[int]string)
    v4 := []uint32{2147483647, 4294967295}
    v10 := GenS1{F0: 127, F1: 2}
    var v11 int32 = 2147483647
    v12 := GenS2{F0: 42, F1: 0, F2: 466}
    var v13 int8 = int8(v2)
    v14 := GenS2{F0: -32768, F1: 2147483647, F2: 2147483647}
    v13 = -95
    switch (("hello" + ("" + ("a" + "a"))) + ((("hello" + "\n\t") + ("foo bar" + "")) + "a")) {
    case "":
        var v15 float32 = 1e10
        _ = v15
    case "a":
        v4 = append(v4, 22)
    case "foo bar":
        v13 -= -1
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch v2 {
    case 1:
        for fi16 := 0; fi16 < 15; fi16++ {
            var v17 int = -608
            if v1 {
                if len(v4) > 0 {
                    v4[0] = 0
                }
                v11 = 2147483647
            } else {
                v18 := GenS1{F0: (v13 & v13), F1: 95}
                _ = v18
            }
            if false {
                v19 := GenS1{F0: (v13 & (v13 << ((0) & 63))), F1: 78}
                _ = v19
            } else {
                var v20 float32 = -2.25
                var v21 int64 = 9223372036854775807
                _ = v20
                _ = v21
            }
            _ = fi16
            _ = v17
        }
    default:
        switch (("foo bar" + "\n\t") + "\n\t") {
        case "foo bar":
            var v22 bool = v1
            _ = v22
        case "a":
            switch (v2 - (v2 / (((-(v2))) | 1))) {
            case 127:
                for fi23 := 0; fi23 < 10; fi23++ {
                    switch 128 {
                    case 107:
                        var v24 uint16 = 255
                        _ = v24
                    default:
                        var v25 int16 = 32767
                        _ = v25
                    }
                    _ = fi23
                }
            case 42:
                var v26 int16 = -32768
                _ = v26
            default:
                v13 = 102
            }
        case "hello":
            for fi27 := 0; fi27 < 2; fi27++ {
                v13 = -1
                if v1 {
                    v1 = false
                    switch ("\n\t" + "a") {
                    case "\n\t":
                        if v1 {
                            if v1 {
                                if len(v4) > 0 {
                                    v4[0] = 308
                                }
                            } else {
                                var v28 uint = 42
                                fmt.Println(v13, v28)
                                var v29 uint64 = 18446744073709551615
                                _ = v28
                                _ = v29
                            }
                        }
                    }
                    if (((fi27 | v2) < (-(fi27))) || false) {
                        var v30 uint = 290
                        _ = v30
                    }
                }
                _ = fi27
            }
        default:
            for fi31 := 0; fi31 < 7; fi31++ {
                var v32 bool = true
                v4 = append(v4, 483)
                _ = fi31
                _ = v32
            }
        }
    }
    if ((!(true || (!v1))) || v1) {
        v33 := GenS3{F0: 0}
        v11 += v11
        _ = v33
    } else {
        if false {
            fmt.Println(v2, v13)
            for fi34 := 0; fi34 < 3; fi34++ {
                var v35 int64 = 1
                if ((-(v2)) >= 2147483647) {
                    var v36 int8 = (^(v13))
                    v37 := GenS4{F0: 0}
                    _ = v36
                    _ = v37
                }
                _ = fi34
                _ = v35
            }
        }
    }
    v11 += v11
    var v38 int32 = (^(v11))
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v38
    _ = v4
}
