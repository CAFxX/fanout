package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint8; F2 string }
type GenS2 struct { F0 uint8; F1 uint16; F2 uint16 }
type GenS3 struct { F0 int; F1 uint32 }
type GenS4 struct { F0 uint; F1 int8; F2 int }
type GenS5 struct { F0 int; F1 uint16; F2 bool }
func genFn0(p0 int16, p1 bool, p2 int) int8 {
    p1 = (152 == 0)
    if (145 <= 1) {
        v6 := GenS1{F0: 0, F1: 42, F2: ""}
        var v7 uint64 = 691
        fmt.Println(p1)
        _ = v6
        _ = v7
    }
    var v8 int64 = 479
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    return 42
}

func main() {
    var v1 int32 = 0
    var v2 float32 = (0.0 * 3.14159)
    var v3 int16 = 101
    var v4 string = ""
    var v5 int64 = -1
    v3 += (v3 << ((2) & 63))
    v4 = "hello"
    v5 -= int64(v3)
    v3 = 514
    switch -1 {
    case 32767:
        var v9 uint = 4294967295
        _ = v9
    default:
        v10 := GenS2{F0: 90, F1: 55, F2: 0}
        _ = v10
    }
    var v11 uint16 = 255
    v12 := []int{}
    for fi13 := 0; fi13 < 1; fi13++ {
        if false {
            for fi14 := 0; fi14 < 16; fi14++ {
                _ = len(v12)
                var v15 int64 = (v5 - v5)
                v15 = (v15 | (v5 / (((v15 - v15)) | 1)))
                switch fi14 {
                case 2147483647:
                    var v16 float64 = float64(v2)
                    _ = v16
                case 128:
                    var v17 float32 = -2.25
                    _ = v17
                case 32768:
                    var v18 uint64 = 18446744073709551615
                    _ = v18
                }
                v3 = (v3 ^ v3)
                _ = fi14
                _ = v15
            }
            if len(v12) > 0 {
                v12[0] = fi13
            }
        } else {
            var v19 uint64 = 42
            _ = v19
        }
        if (!(fi13 > 2147483647)) {
            var v20 int8 = 1
            for fi21 := 0; fi21 < 2; fi21++ {
                v3 -= (v3 + (v3 >> ((0) & 63)))
                v3 += -32768
                v22 := GenS2{F0: 128, F1: (^(v11)), F2: 2}
                _ = fi21
                _ = v22
            }
            _ = len(v12)
            _ = v20
        } else {
            switch ((v4 + v4) + v4) {
            case "\n\t":
                if false {
                    switch (fi13 >> ((94) & 63)) {
                    case 128:
                        v1 += -2147483648
                    case 32767:
                        if (255 > 2) {
                            v23 := GenS3{F0: fi13, F1: 851}
                            var v24 string = "hello"
                            _ = v23
                            _ = v24
                        }
                    case 487:
                        for fi25 := 0; fi25 < 3; fi25++ {
                            v26 := GenS2{F0: 255, F1: v11, F2: (v11 >> ((675) & 63))}
                            _ = fi25
                            _ = v26
                        }
                    default:
                        var v27 uint8 = 245
                        _ = v27
                    }
                    switch -548 {
                    case 650:
                        if (true || ((int(v11) < fi13) || true)) {
                            if len(v12) >= 2 {
                                v12 = v12[1:len(v12)]
                            }
                            var v28 int32 = (v1 | (v1 >> ((42) & 63)))
                            switch (((v4 + (("a" + ("hello" + "hello")) + v4)) + (v4 + (v4 + ("hello" + v4)))) + "a") {
                            case "foo bar":
                                if len(v12) > 0 {
                                    v12[0] = fi13
                                }
                            case "hello":
                                var v29 uint64 = 957
                                _ = v29
                            }
                            _ = v28
                        }
                    }
                    var v30 int32 = v1
                    _ = v30
                }
            case "":
                for fi31 := 0; fi31 < 9; fi31++ {
                    switch ("" + "\n\t") {
                    case "\n\t":
                        var v32 uint8 = 2
                        _ = v32
                    case "hello":
                        v33 := GenS4{F0: 18, F1: 127, F2: -2147483648}
                        _ = v33
                    }
                    var v34 int8 = -42
                    var v35 uint16 = 65535
                    _ = fi31
                    _ = v34
                    _ = v35
                }
            }
            var v36 int64 = (-(v5))
            _ = v36
        }
        for fi37 := 0; fi37 < 9; fi37++ {
            v38 := GenS5{F0: (-(fi37)), F1: (v11 >> ((42) & 63)), F2: (false || true)}
            _ = len(v12)
            v11 = v11
            _ = fi37
            _ = v38
        }
        _ = fi13
    }
    switch 128 {
    case -53:
        v39 := GenS4{F0: 2, F1: genFn0(1000, true, -110), F2: int(v1)}
        _ = v39
    default:
        v5 -= (v5 ^ (v5 * -9223372036854775808))
    }
    v2 = v2
    fmt.Println(v3, v5)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
