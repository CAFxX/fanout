package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64; F2 int16 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint16; F1 int64 }
type GenS4 struct { F0 uint16; F1 int64; F2 int64 }

func main() {
    var v1 int = 9223372036854775807
    var v2 float32 = ((0.1 / 0.1) / (3.14159 / 1e10))
    v3 := make(map[uint32]bool)
    if (265 < -1) {
        switch v1 {
        case 58:
            var v4 uint = 65535
            _ = v4
        case 9223372036854775807:
            var v5 uint64 = 194
            _ = v5
        case 32768:
            v1 = v1
        default:
            if (65535 <= 115) {
                var v6 bool = true
                v7 := GenS1{F0: v6, F1: 4294967295, F2: 0}
                _ = v6
                _ = v7
            } else {
                v8 := GenS2{F0: 0}
                for fi9 := 0; fi9 < 9; fi9++ {
                    switch ("foo bar" + (("hello" + "a") + "\n\t")) {
                    case "\n\t":
                        for fi10 := 0; fi10 < 15; fi10++ {
                            if true {
                                var v11 uint8 = 42
                                v2 += v2
                                _ = v11
                            }
                            _ = fi10
                        }
                    default:
                        fmt.Println(v1, v1, fi9)
                    }
                    _ = fi9
                }
                _ = v8
            }
        }
    }
    if false {
        v12 := []int{}
        _ = v12
    }
    for fi13 := 0; fi13 < 9; fi13++ {
        var v14 int8 = 42
        _ = fi13
        _ = v14
    }
    v2 -= 0.0
    v1 += v1
    switch (("" + ("a" + ("hello" + ("" + ("" + ("foo bar" + ("\n\t" + "a"))))))) + "hello") {
    case "hello":
        var v15 uint = 65535
        _ = v15
    case "\n\t":
        v2 += v2
    case "a":
        if (true || true) {
            var v16 float32 = v2
            var v17 uint32 = 65535
            _ = v16
            _ = v17
        } else {
            v18 := v3[4294967295]
            _ = len(v3)
            for fi19 := 0; fi19 < 12; fi19++ {
                if (709 == 2147483647) {
                    v20 := GenS1{F0: false, F1: 795, F2: 1}
                    if (42 == 2) {
                        fmt.Println(v2)
                    }
                    var v21 int8 = 100
                    _ = v20
                    _ = v21
                }
                v22 := GenS1{F0: (!false), F1: 42, F2: -364}
                var v23 int = 328
                _ = fi19
                _ = v22
                _ = v23
            }
            _ = v18
        }
    default:
        var v24 int = v1
        _ = v24
    }
    var v25 int64 = 9223372036854775807
    var v26 uint8 = 0
    for fi27 := 0; fi27 < 3; fi27++ {
        v2 = ((v2 * 1e10) + v2)
        v1 = 4294967295
        _ = fi27
    }
    var v28 bool = ((((!true) || false) && (true || true)) || true)
    var v29 string = ("a" + "foo bar")
    v29 = "foo bar"
    v30 := GenS3{F0: 2, F1: (v25 + (v25 >> ((18446744073709551615) & 63)))}
    v31 := []int{}
    if (!v28) {
        var v32 string = v29
        v1 -= (v1 / ((1) | 1))
        _ = v32
    } else {
        if (1000 <= -564) {
            v1 -= v1
            v2 -= (v2 + 1.5)
        } else {
            v33 := GenS2{F0: 255}
            _ = v33
        }
        v3[65535] = v28
    }
    v2 += (1.5 + v2)
    v34 := GenS3{F0: 1, F1: v25}
    if true {
        fmt.Println(v1)
        if v28 {
            for fi35 := 0; fi35 < 10; fi35++ {
                var v36 string = (v29 + "foo bar")
                _ = fi35
                _ = v36
            }
            if (!(-1 == 1)) {
                var v37 int16 = 0
                if v28 {
                    var v38 int32 = 42
                    if v28 {
                        fmt.Println(v28, v38, v1)
                        var v39 uint = 65535
                        _ = v39
                    } else {
                        v29 = "a"
                    }
                    var v40 int64 = v25
                    _ = v38
                    _ = v40
                }
                _ = v37
            } else {
                v41 := GenS1{F0: (42 < 90), F1: 684, F2: 0}
                _ = v41
            }
        }
        v42 := GenS4{F0: 0, F1: (-(v25)), F2: (v25 | -9223372036854775808)}
        _ = v42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v31:", len(v31), v31)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v34
}
