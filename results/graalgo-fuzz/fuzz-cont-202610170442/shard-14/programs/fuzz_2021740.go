package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int8; F1 int64 }
type GenS4 struct { F0 uint16; F1 int32; F2 int64 }
type GenS5 struct { F0 uint16; F1 int32; F2 int8 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint16; F1 int16; F2 int16 }

func main() {
    v1 := make(map[uint32]string)
    var v2 float64 = (1e10 * (3.14159 / -2.25))
    v3 := []string{("" + "a"), "foo bar"}
    var v4 uint16 = 65535
    for fi5 := 0; fi5 < 20; fi5++ {
        var v6 int = 4294967295
        if false {
            if false {
                v6 += v6
            } else {
                for fi7 := 0; fi7 < 20; fi7++ {
                    switch v6 {
                    case 32768:
                        var v8 uint32 = 0
                        v9 := GenS1{F0: 0}
                        var v10 int64 = 0
                        _ = v10
                        _ = v8
                        _ = v9
                    case 123:
                        v11 := GenS2{F0: v4}
                        _ = v11
                    }
                    var v12 uint64 = 42
                    _ = fi7
                    _ = v12
                }
                v13 := GenS2{F0: v4}
                _ = v13
            }
            if (2 < 4294967295) {
                var v14 int8 = 100
                var v15 int64 = 0
                _ = v14
                _ = v15
            }
        }
        var v16 float64 = (((3.14159 / (v2 - (v2 * v2))) / v2) - v2)
        _ = fi5
        _ = v16
        _ = v6
    }
    if true {
        var v17 uint = 42
        v18 := GenS3{F0: -128, F1: 0}
        v19 := GenS4{F0: (v4 - 1), F1: -806, F2: 42}
        _ = v17
        _ = v18
        _ = v19
    } else {
        fmt.Println(v4)
        var v20 int64 = -9223372036854775808
        for fi21 := 0; fi21 < 9; fi21++ {
            if (-125 >= -1) {
                if false {
                    v2 += (((1e10 * v2) * ((((float64(3.14159) * (v2 - v2)) - v2) - v2) / 1.5)) / v2)
                    switch ("hello" + ("hello" + "")) {
                    case "foo bar":
                        var v22 int8 = -5
                        _ = v22
                    case "hello":
                        v23 := GenS3{F0: 1, F1: (-(v20))}
                        _ = v23
                    }
                } else {
                    if false {
                        var v24 uint32 = 2147483647
                        v1[2147483647] = "\n\t"
                        _ = v24
                    } else {
                        delete(v1, 864)
                        var v25 int64 = -1
                        var v26 int = (fi21 / ((fi21) | 1))
                        _ = v25
                        _ = v26
                    }
                    for fi27 := 0; fi27 < 18; fi27++ {
                        if len(v3) > 0 {
                            v3[0] = "foo bar"
                        }
                        _ = fi27
                    }
                }
            } else {
                switch ("hello" + (((("hello" + "") + "hello") + ("hello" + "hello")) + "")) {
                case "":
                    v28 := GenS4{F0: (v4 % ((v4) | 1)), F1: 622, F2: (v20 >> ((884) & 63))}
                    _ = v28
                case "a":
                    v1[0] = ("" + "")
                case "\n\t":
                    var v29 int32 = 1
                    _ = v29
                default:
                    var v30 uint64 = 2
                    _ = v30
                }
                switch ((("hello" + (((("a" + "\n\t") + ("hello" + "hello")) + ("a" + "foo bar")) + "")) + "\n\t") + "") {
                case "a":
                    if true {
                        v4 = (^(v4))
                    }
                case "foo bar":
                    var v31 float32 = 3.14159
                    _ = v31
                case "":
                    if true {
                        v4 += v4
                        switch (fi21 % ((fi21) | 1)) {
                        case 2147483647:
                            var v32 bool = true
                            _ = v32
                        default:
                            v2 += float64((1.5 - 1.5))
                        }
                        if (187 == 0) {
                            var v33 bool = (!true)
                            v20 += v20
                            var v34 int = 127
                            _ = v33
                            _ = v34
                        }
                    }
                default:
                    v20 = (v20 >> ((255) & 63))
                }
                var v35 bool = true
                _ = v35
            }
            _ = fi21
        }
        _ = v20
    }
    fmt.Println(v2, v4)
    v4 += (v4 ^ 255)
    v36 := GenS5{F0: v4, F1: 178, F2: 0}
    v4 -= (v4 << ((988) & 63))
    if (-2147483648 >= 2147483647) {
        if false {
            var v37 string = "a"
            _ = v37
        } else {
            var v38 uint32 = 65535
            v39 := GenS6{F0: true}
            _ = len(v1)
            _ = v38
            _ = v39
        }
        v4 = 0
    } else {
        v4 = v4
        v40 := GenS7{F0: 4294967295}
        _ = v40
    }
    fmt.Println(v4)
    v41 := GenS8{F0: 65535, F1: 42, F2: 1000}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v41
}
