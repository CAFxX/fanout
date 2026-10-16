package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 int32 }
type GenS2 struct { F0 uint8; F1 int64; F2 int8 }
type GenS3 struct { F0 uint8; F1 int64 }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 string; F1 uint16 }
type GenS6 struct { F0 int; F1 bool; F2 int8 }

func main() {
    v1 := make(map[string]int)
    var v2 uint8 = 1
    v3 := make(map[string]int)
    v4 := make(map[int]bool)
    v5 := []string{"a", "a"}
    var v6 uint64 = 997
    if len(v5) > 0 {
        v5[0] = "foo bar"
    }
    v2 = (v2 | (v2 % ((v2) | 1)))
    _ = len(v4)
    v6 += 253
    fmt.Println(v2, v6)
    if false {
        if false {
            v2 = (v2 / ((v2) | 1))
        }
        fmt.Println(v2, v2)
    } else {
        var v7 float64 = (-2.25 / float64((3.14159 * (3.14159 * 1.5))))
        switch (("a" + "a") + "") {
        case "":
            switch ("foo bar" + ("\n\t" + ("" + ("hello" + ("" + "\n\t"))))) {
            case "foo bar":
                v8 := GenS1{F0: -528, F1: 148, F2: -2147483648}
                _ = v8
            case "hello":
                fmt.Println(v6)
            }
        case "\n\t":
            var v9 uint8 = 134
            _ = v9
        case "hello":
            switch "\n\t" {
            case "foo bar":
                var v10 uint64 = v6
                _ = v10
            case "a":
                v6 = v6
            default:
                v7 = ((v7 + -2.25) * 0.1)
            }
        }
        _ = len(v5)
        _ = v7
    }
    v11 := GenS1{F0: 0, F1: -745, F2: 42}
    for fi12 := 0; fi12 < 15; fi12++ {
        for fi13 := 0; fi13 < 9; fi13++ {
            switch "foo bar" {
            case "":
                _ = len(v1)
                v6 -= 1
            case "a":
                v2 = (^(v2))
            }
            var v14 uint64 = 18446744073709551615
            var v15 int = int(v2)
            _ = fi13
            _ = v14
            _ = v15
        }
        _ = fi12
    }
    var v16 uint8 = (v2 / ((127) | 1))
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    v6 = v6
    if ("hello" >= "") {
        v16 -= 195
        for fi17 := 0; fi17 < 16; fi17++ {
            for fi18 := 0; fi18 < 12; fi18++ {
                var v19 int64 = -655
                var v20 float32 = 0.1
                _ = fi18
                _ = v19
                _ = v20
            }
            if (v2 == (v16 * v16)) {
                if true {
                    v16 = (^(v16))
                    if (!((uint32(v2) < 65535) && (18446744073709551615 != 4294967295))) {
                        v6 -= v6
                        v21 := GenS2{F0: v16, F1: -512, F2: 1}
                        _ = v21
                    }
                    var v22 uint8 = (^(v16))
                    _ = v22
                } else {
                    if (0.1 < 1e10) {
                        for fi23 := 0; fi23 < 4; fi23++ {
                            if false {
                                var v24 int16 = -612
                                if len(v5) > 0 {
                                    v5[0] = ("foo bar" + "a")
                                }
                                _ = v24
                            } else {
                                fmt.Println(fi23, fi23)
                            }
                            _ = fi23
                        }
                    } else {
                        for fi25 := 0; fi25 < 14; fi25++ {
                            fmt.Println(v16, fi17)
                            _ = fi25
                        }
                    }
                    v26 := GenS2{F0: v16, F1: 1, F2: -4}
                    _ = v26
                }
                if (!(!true)) {
                    for fi27 := 0; fi27 < 10; fi27++ {
                        v6 = 4294967295
                        var v28 uint = 42
                        v2 -= (v2 << ((v28) & 63))
                        _ = fi27
                        _ = v28
                    }
                    v29 := GenS3{F0: (v16 >> ((44) & 63)), F1: 1}
                    var v30 float64 = 3.14159
                    _ = v29
                    _ = v30
                } else {
                    v31 := GenS4{F0: 127}
                    _ = v31
                }
                switch -1 {
                case -703:
                    v32 := GenS5{F0: ("a" + "hello"), F1: 896}
                    _ = v32
                case 776:
                    fmt.Println(fi17)
                }
            }
            var v33 uint16 = 65535
            _ = fi17
            _ = v33
        }
    } else {
        fmt.Println(v16)
        fmt.Println(v2)
        if true {
            for fi34 := 0; fi34 < 3; fi34++ {
                fmt.Println(v2, fi34, fi34)
                _ = fi34
            }
            v2 += (v2 * 1)
        }
    }
    if (v6 >= 201) {
        var v35 uint64 = (v6 ^ 2)
        v16 = v16
        v16 += (^(v2))
        _ = v35
    }
    v6 += (v6 | 0)
    v5 = append(v5, "hello")
    v36 := GenS4{F0: 127}
    v37 := GenS6{F0: 4294967295, F1: true, F2: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v16
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v4
    _ = v5
    _ = v6
}
