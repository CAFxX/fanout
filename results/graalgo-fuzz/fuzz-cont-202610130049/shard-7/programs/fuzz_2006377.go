package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 int8 }
type GenS3 struct { F0 uint32; F1 uint32 }
type GenS4 struct { F0 int16; F1 int16 }
type GenS5 struct { F0 uint16 }
func genFn0(p0 uint, p1 int16) string {
    var v5 int = -2147483648
    fmt.Println(v5)
    fmt.Println(v5, p1, p1)
    _ = p0
    _ = p1
    _ = v5
    return "hello"
}
func genFn1() uint32 {
    var v6 float32 = (3.14159 - (1e10 / 1.0))
    if true {
        var v7 int = 4294967295
        _ = v7
    } else {
        v6 = 1e10
        var v8 int = 1
        _ = v8
    }
    _ = v6
    return 65535
}

func main() {
    v1 := []int{}
    var v2 int64 = 42
    v3 := []uint32{4294967295, 0, 279, 1}
    var v4 int64 = v2
    v2 = v2
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v9 := GenS1{F0: 0}
    var v10 int16 = 0
    if (v10 > -930) {
        v2 -= (v2 << ((1) & 63))
        v4 += (^(v2))
        switch (("\n\t" + "") + "foo bar") {
        case "a":
            switch "hello" {
            case "\n\t":
                var v11 uint32 = 1
                _ = v11
            case "foo bar":
                if true {
                    var v12 float32 = (0.0 - 1e10)
                    _ = v12
                }
            }
        }
    }
    if ("" < "\n\t") {
        if true {
            for fi13 := 0; fi13 < 16; fi13++ {
                _ = len(v3)
                _ = len(v3)
                v10 = (v10 >> ((907) & 63))
                _ = fi13
            }
        } else {
            fmt.Println(v4)
            for fi14 := 0; fi14 < 10; fi14++ {
                if ((v10 * 1) == 42) {
                    if false {
                        fmt.Println(v10, v4, v10)
                    }
                    for fi15 := 0; fi15 < 1; fi15++ {
                        var v16 uint8 = 2
                        switch ("" + ("\n\t" + ((("a" + (("" + "hello") + ("a" + ""))) + "\n\t") + "\n\t"))) {
                        case "hello":
                            v17 := GenS1{F0: 42}
                            _ = v17
                        default:
                            v16 += (v16 - 42)
                        }
                        if true {
                            if (((genFn1() > 42) || false) || true) {
                                v4 += v2
                            }
                            v18 := GenS2{F0: 42, F1: 100}
                            _ = v18
                        }
                        _ = fi15
                        _ = v16
                    }
                } else {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    if len(v3) > 0 {
                        v3[0] = 2
                    }
                    fmt.Println(v4)
                }
                v2 -= v4
                _ = fi14
            }
        }
    } else {
        if true {
            v2 = -1
        }
        var v19 string = "foo bar"
        _ = v19
    }
    for fi20 := 0; fi20 < 17; fi20++ {
        v21 := GenS1{F0: 410}
        for fi22 := 0; fi22 < 9; fi22++ {
            switch fi20 {
            case -699:
                v23 := GenS3{F0: 2, F1: 2147483647}
                _ = v23
            }
            _ = fi22
        }
        _ = fi20
        _ = v21
    }
    v2 += 1
    fmt.Println(v10, v4)
    v4 -= v2
    v24 := GenS3{F0: 4294967295, F1: 2}
    v1 = append(v1, 720)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if true {
        v4 = (v2 | (v2 - (v2 | (v2 * (v2 % (((v2 << ((255) & 63))) | 1))))))
    } else {
        for fi25 := 0; fi25 < 10; fi25++ {
            for fi26 := 0; fi26 < 10; fi26++ {
                var v27 uint16 = 0
                var v28 uint64 = 0
                v29 := GenS4{F0: v10, F1: (v10 % ((42) | 1))}
                switch (fi25 + fi26) {
                case 42:
                    if ("hello" >= ("hello" + ("" + ("\n\t" + "hello")))) {
                        if true {
                            if (true || false) {
                                var v30 string = ""
                                fmt.Println(fi26)
                                _ = v30
                            } else {
                                var v31 int32 = 1
                                v32 := GenS3{F0: 0, F1: 42}
                                _ = v31
                                _ = v32
                            }
                            var v33 uint64 = 1
                            _ = v33
                        } else {
                            v34 := GenS4{F0: 42, F1: (v10 - -879)}
                            v27 += (^(v27))
                            fmt.Println(v2, v4)
                            var v35 uint8 = 131
                            v4 = (^(v2))
                            _ = v34
                            _ = v35
                        }
                        v10 = -1
                    }
                case 988:
                    var v36 bool = true
                    _ = v36
                default:
                    var v37 float32 = 3.14159
                    _ = v37
                }
                var v38 int = 2147483647
                if true {
                    v4 -= v2
                }
                _ = fi26
                _ = v27
                _ = v28
                _ = v29
                _ = v38
            }
            _ = fi25
        }
        v39 := GenS5{F0: 850}
        _ = v39
    }
    fmt.Println(v2)
    fmt.Println(v4)
    if len(v1) > 0 {
        v1[0] = 127
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v9
}
