package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint32; F1 int64 }
type GenS4 struct { F0 uint32; F1 int8 }
type GenS5 struct { F0 int; F1 int8; F2 uint16 }
func genFn0(p0 uint) int32 {
    v6 := GenS1{F0: -9223372036854775808}
    p0 = 42
    fmt.Println(p0)
    _ = p0
    _ = v6
    return -2147483648
}

func main() {
    v1 := []string{("hello" + ("\n\t" + "foo bar")), "\n\t"}
    v2 := make(map[string]string)
    var v3 uint32 = 2147483647
    var v4 float32 = 0.1
    v5 := make(map[uint32]bool)
    var v7 int16 = -32768
    switch ("foo bar" + (("hello" + ("hello" + "a")) + (("a" + (("" + ("foo bar" + "a")) + ("" + "hello"))) + "a"))) {
    case "":
        if true {
            var v8 int32 = -1
            _ = v8
        } else {
            var v9 uint = 581
            _ = v9
        }
    default:
        _ = len(v1)
    }
    if ((-9223372036854775808 <= -1) && true) {
        for fi10 := 0; fi10 < 19; fi10++ {
            for fi11 := 0; fi11 < 10; fi11++ {
                fmt.Println(fi10)
                _ = fi11
            }
            v12 := GenS2{F0: ("a" + "a")}
            switch (fi10 | (fi10 | (fi10 + (fi10 | fi10)))) {
            case 2147483647:
                v7 += v7
            }
            _ = fi10
            _ = v12
        }
        for fi13 := 0; fi13 < 15; fi13++ {
            v7 -= v7
            for fi14 := 0; fi14 < 4; fi14++ {
                v7 = (v7 - -32768)
                _ = fi14
            }
            _ = fi13
        }
    } else {
        fmt.Println(v7, v4)
    }
    if true {
        if (false && false) {
            v15 := GenS3{F0: (v3 >> ((2) & 63)), F1: -9223372036854775808}
            var v16 uint = 0
            _ = v15
            _ = v16
        } else {
            if (v4 <= (float32(1.5) / ((1e10 / v4) - (3.14159 + 3.14159)))) {
                if true {
                    if ((!true) && (!false)) {
                        var v17 uint = 855
                        var v18 int32 = -2147483648
                        _ = v17
                        _ = v18
                    }
                } else {
                    v19 := GenS2{F0: "\n\t"}
                    _ = v19
                }
                v1 = append(v1, "foo bar")
                for fi20 := 0; fi20 < 7; fi20++ {
                    if (!false) {
                        if ((!false) || true) {
                            if true {
                                var v21 uint16 = 32767
                                fmt.Println(v21)
                                var v22 uint = 65535
                                _ = v21
                                _ = v22
                            }
                            for fi23 := 0; fi23 < 9; fi23++ {
                                switch ("a" + ("a" + "")) {
                                case "a":
                                    if (!false) {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        var v24 int16 = 42
                                        _ = v24
                                    } else {
                                        var v25 float32 = v4
                                        delete(v2, ("foo bar" + "hello"))
                                        _ = v25
                                    }
                                case "foo bar":
                                    if ((42 != 79) && (fi20 < 42)) {
                                        var v26 uint8 = 127
                                        if (true || false) {
                                            if true {
                                                v2["foo bar"] = (("foo bar" + "") + "hello")
                                                if false {
                                                    fmt.Println(v7, v4, fi20)
                                                } else {
                                                    if false {
                                                        if false {
                                                            v4 = ((1e10 + (v4 / (v4 + 1.5))) + v4)
                                                            fmt.Println(fi20)
                                                        }
                                                    }
                                                    v26 += v26
                                                }
                                                var v27 int32 = -1
                                                _ = v27
                                            } else {
                                                v26 = v26
                                                var v28 bool = false
                                                v29 := GenS2{F0: "hello"}
                                                _ = v28
                                                _ = v29
                                            }
                                        } else {
                                            v5[(v3 << ((1) & 63))] = (true && true)
                                            v26 = v26
                                        }
                                        switch (fi23 % (((fi20 | (^(fi23)))) | 1)) {
                                        case 9223372036854775807:
                                            fmt.Println(v3, v4, fi23)
                                        case 0:
                                            v30 := GenS3{F0: (v3 ^ 381), F1: 9223372036854775807}
                                            _ = v30
                                        case 42:
                                            fmt.Println(v4, v4)
                                        }
                                        _ = v26
                                    }
                                    v31 := GenS4{F0: v3, F1: -1}
                                    fmt.Println(v3, fi23)
                                    _ = v31
                                case "hello":
                                    var v32 uint = 18446744073709551615
                                    _ = v32
                                }
                                _ = fi23
                            }
                            switch (("" + ("" + "\n\t")) + "a") {
                            case "\n\t":
                                v4 = v4
                            case "hello":
                                switch fi20 {
                                case 42:
                                    v1 = append(v1, ("\n\t" + "foo bar"))
                                default:
                                    v33 := v2["hello"]
                                    _ = v33
                                }
                            }
                        }
                        var v34 bool = false
                        _ = v34
                    } else {
                        for fi35 := 0; fi35 < 1; fi35++ {
                            v3 += (^(v3))
                            _ = fi35
                        }
                    }
                    var v36 int16 = 0
                    v5[v3] = (((657 <= 2) || false) || (-129 >= genFn0(489)))
                    _ = fi20
                    _ = v36
                }
            } else {
                fmt.Println(v7)
                for fi37 := 0; fi37 < 20; fi37++ {
                    var v38 int = -2147483648
                    fmt.Println(v3, v7)
                    var v39 int8 = 100
                    _ = fi37
                    _ = v38
                    _ = v39
                }
            }
            var v40 float64 = 3.14159
            _ = v40
        }
    } else {
        v7 -= (^(v7))
    }
    fmt.Println(v3, v7)
    switch 128 {
    case -582:
        for fi41 := 0; fi41 < 18; fi41++ {
            v42 := GenS5{F0: int(v3), F1: 0, F2: 377}
            _ = fi41
            _ = v42
        }
    case -376:
        var v43 int64 = 565
        _ = v43
    default:
        fmt.Println(v4, v4, v3)
    }
    v3 -= (v3 - v3)
    v4 = v4
    if false {
        var v44 float64 = float64(1e10)
        _ = v44
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
