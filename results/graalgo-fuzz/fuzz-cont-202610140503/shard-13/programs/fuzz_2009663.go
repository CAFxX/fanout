package main

import (
    "fmt"
)

type GenS1 struct { F0 string }

func main() {
    v1 := []int{-9223372036854775808, 0, -980}
    v2 := make(map[uint32]uint32)
    var v3 float32 = 0.0
    v4 := make(map[int]uint32)
    v5 := make(map[string]string)
    var v6 float64 = 1.5
    v6 += -2.25
    switch ("foo bar" + ("\n\t" + ("" + ("\n\t" + ("" + "hello"))))) {
    case "\n\t":
        v7 := GenS1{F0: ("hello" + "\n\t")}
        _ = v7
    }
    fmt.Println(v6, v6)
    var v8 int16 = -173
    var v9 int64 = 1
    v10 := GenS1{F0: "a"}
    var v11 uint16 = 612
    v6 += 1.5
    if false {
        switch ("a" + "\n\t") {
        case "\n\t":
            if true {
                for fi12 := 0; fi12 < 4; fi12++ {
                    var v13 int8 = -1
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    _ = fi12
                    _ = v13
                }
            }
        case "hello":
            if ((v9 * 9223372036854775807) != v9) {
                _ = len(v5)
                var v14 float32 = (v3 / 3.14159)
                if (48 != 2) {
                    if (((v8 > 0) && true) || true) {
                        if (552 > 4294967295) {
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        } else {
                            var v15 uint16 = (^(v11))
                            for fi16 := 0; fi16 < 15; fi16++ {
                                fmt.Println(fi16)
                                if (1e10 != v6) {
                                    var v17 bool = false
                                    v5["a"] = ((("a" + ((("hello" + "a") + "hello") + (("hello" + ("" + "foo bar")) + "\n\t"))) + (("" + ("foo bar" + "a")) + "\n\t")) + ("\n\t" + "a"))
                                    v17 = false
                                    _ = v17
                                }
                                _ = fi16
                            }
                            _ = v15
                        }
                        v3 += -2.25
                        v3 = 0.0
                    } else {
                        var v18 uint16 = (v11 + 65535)
                        _ = v18
                    }
                    delete(v2, 0)
                } else {
                    var v19 int32 = -75
                    if (false || false) {
                        if (false && (v19 < -627)) {
                            switch 4294967295 {
                            case 128:
                                var v20 int64 = (v9 - (v9 - v9))
                                _ = v20
                            case -1:
                                var v21 float64 = (v6 - (v6 + v6))
                                _ = v21
                            case 4294967295:
                                for fi22 := 0; fi22 < 8; fi22++ {
                                    if false {
                                        var v23 int = (-(fi22))
                                        v24 := GenS1{F0: ("a" + "hello")}
                                        _ = v23
                                        _ = v24
                                    } else {
                                        var v25 uint64 = 4294967295
                                        if ((fi22 >> ((18446744073709551615) & 63)) == (fi22 & 2147483647)) {
                                            v11 = v11
                                            if true {
                                                v9 = v9
                                                v2[uint32(v25)] = 2
                                            } else {
                                                if false {
                                                    var v26 uint64 = (^(v25))
                                                    _ = v26
                                                }
                                                if false {
                                                    v4[fi22] = 2147483648
                                                } else {
                                                    v6 = float64((-2.25 - ((v3 - v14) * (((3.14159 * v14) - v14) - (v14 - 3.14159)))))
                                                    v25 = v25
                                                    var v27 uint16 = v11
                                                    _ = v27
                                                }
                                            }
                                        } else {
                                            var v28 bool = ((!(0 > -5)) && true)
                                            v6 = v6
                                            _ = v28
                                        }
                                        var v29 string = "foo bar"
                                        _ = v25
                                        _ = v29
                                    }
                                    var v30 uint32 = 0
                                    _ = fi22
                                    _ = v30
                                }
                            default:
                                v6 = float64(v3)
                            }
                            fmt.Println(v3, v11, v8)
                            v1 = append(v1, 0)
                        }
                    }
                    _ = v19
                }
                _ = v14
            }
        case "":
            var v31 float32 = 0.1
            _ = v31
        }
        v6 = (v6 / v6)
        if false {
            if true {
                v6 = v6
                v1 = append(v1, 32767)
                if false {
                    v6 -= v6
                    if true {
                        var v32 uint8 = 255
                        _ = v32
                    }
                    v6 += v6
                }
            } else {
                for fi33 := 0; fi33 < 14; fi33++ {
                    var v34 int16 = (v8 / (((-(v8))) | 1))
                    _ = fi33
                    _ = v34
                }
                delete(v4, 32767)
            }
            var v35 float64 = 0.1
            _ = v35
        }
    }
    delete(v5, "a")
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
