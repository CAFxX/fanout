package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int16; F1 uint64 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 uint8; F1 int; F2 int32 }
func genFn0() uint32 {
    if (278 > 0) {
        var v4 string = "a"
        _ = v4
    }
    if false {
        v5 := []int{0}
        v6 := GenS1{F0: 0, F1: 2}
        var v7 bool = (("a" + ("" + (("hello" + ("a" + (("\n\t" + ("\n\t" + ("hello" + ""))) + ""))) + ("hello" + "hello")))) != "a")
        _ = v5
        _ = v6
        _ = v7
    }
    var v8 string = "\n\t"
    var v9 uint = 1
    _ = v8
    _ = v9
    return 42
}
func genFn1(p0 uint8) uint64 {
    if (-2147483648 < 2147483647) {
        v10 := GenS2{F0: 2, F1: (p0 << ((2) & 63))}
        _ = v10
    }
    var v11 int64 = -209
    var v12 int16 = 1000
    switch 9223372036854775807 {
    case -177:
        var v13 uint64 = 1
        _ = v13
    }
    _ = p0
    _ = v11
    _ = v12
    return 42
}

func main() {
    var v1 float32 = 1e10
    v2 := []string{"hello", ((("" + "hello") + "\n\t") + ("\n\t" + "foo bar")), "\n\t"}
    var v3 string = "foo bar"
    if true {
        _ = len(v2)
    }
    for fi14 := 0; fi14 < 15; fi14++ {
        fmt.Println(v3, v3, v1)
        _ = fi14
    }
    var v15 uint32 = 118
    v16 := GenS2{F0: 706, F1: 127}
    v17 := GenS3{F0: -955}
    if false {
        fmt.Println(v15, v15)
    } else {
        if (v3 > "") {
            if false {
                v18 := GenS4{F0: -670, F1: 0}
                _ = v18
            }
            if len(v2) > 0 {
                v2[0] = "\n\t"
            }
        }
        for fi19 := 0; fi19 < 16; fi19++ {
            if (404 <= genFn1(207)) {
                for fi20 := 0; fi20 < 13; fi20++ {
                    switch v3 {
                    case "\n\t":
                        if (2147483647 == 2147483647) {
                            fmt.Println(fi20, fi20)
                            var v21 int16 = 32767
                            v3 = (v3 + "foo bar")
                            v22 := GenS5{F0: 223}
                            switch v3 {
                            case "hello":
                                var v23 bool = ((((((v1 / (v1 * ((v1 * ((((v1 / v1) * v1) * 0.0) / v1)) * 1.5))) - 0.1) - v1) * v1) - float32((0.0 - 1e10))) != (1.5 / (float32(0.0) - v1)))
                                _ = v23
                            case "\n\t":
                                if ((false || ((0 == -2147483648) && (!(!true)))) && false) {
                                    var v24 string = "\n\t"
                                    var v25 int32 = 0
                                    if true {
                                        fmt.Println(v1)
                                    } else {
                                        var v26 int16 = (v21 * (v21 & (v21 * (v21 >> ((119) & 63)))))
                                        if (744 > 65535) {
                                            v27 := GenS3{F0: (v21 % ((v26) | 1))}
                                            _ = v27
                                        } else {
                                            var v28 uint = 0
                                            switch -1 {
                                            case 32768:
                                                var v29 uint16 = 472
                                                _ = v29
                                            case -9223372036854775808:
                                                var v30 int32 = v25
                                                _ = v30
                                            case -2147483648:
                                                switch v24 {
                                                case "hello":
                                                    var v31 uint16 = 0
                                                    _ = v31
                                                case "a":
                                                    fmt.Println(fi20, v21, fi19)
                                                }
                                            default:
                                                var v32 int32 = (v25 / ((v25) | 1))
                                                _ = v32
                                            }
                                            var v33 int64 = 547
                                            _ = v28
                                            _ = v33
                                        }
                                        _ = v26
                                    }
                                    _ = v24
                                    _ = v25
                                }
                            default:
                                var v34 int32 = -414
                                _ = v34
                            }
                            _ = v21
                            _ = v22
                        } else {
                            fmt.Println(v1, fi19, v15)
                            v1 = (((1.5 * 0.1) * 1e10) + (((float32(-2.25) * (float32(((1.5 * -2.25) / 3.14159)) * v1)) / v1) / (v1 - float32(3.14159))))
                            switch (-(fi19)) {
                            case 690:
                                v2 = append(v2, v3)
                            case 2147483647:
                                v35 := GenS3{F0: 1}
                                _ = v35
                            case -416:
                                v36 := GenS1{F0: 9223372036854775807, F1: 857}
                                _ = v36
                            }
                        }
                    case "hello":
                        fmt.Println(fi20, fi20, fi19)
                    }
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    var v37 int64 = -969
                    v37 -= (v37 >> ((491) & 63))
                    _ = fi20
                    _ = v37
                }
                switch (fi19 - -1) {
                case 689:
                    v15 += (v15 & (^(v15)))
                }
            }
            _ = fi19
        }
        for fi38 := 0; fi38 < 8; fi38++ {
            v3 = (v3 + v3)
            _ = fi38
        }
    }
    v3 = "foo bar"
    v39 := GenS6{F0: 207, F1: -875, F2: -100}
    var v40 uint32 = v15
    for fi41 := 0; fi41 < 7; fi41++ {
        var v42 bool = ((!true) || true)
        v43 := GenS1{F0: 9223372036854775807, F1: genFn1(2)}
        var v44 bool = (-1 <= -1)
        _ = fi41
        _ = v42
        _ = v43
        _ = v44
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v39
    _ = v40
}
