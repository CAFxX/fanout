package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 int16 }
type GenS2 struct { F0 bool; F1 uint16 }
type GenS3 struct { F0 int16; F1 int32 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 uint32; F1 uint16; F2 uint8 }

func main() {
    var v1 string = ""
    var v2 bool = ((((v1 + v1) + v1) + (v1 + "a")) != (v1 + v1))
    v3 := make(map[string]bool)
    var v4 string = (v1 + "foo bar")
    v5 := GenS1{F0: 0, F1: 18446744073709551615, F2: -32768}
    if ((!v2) && v2) {
        if ((!v2) || true) {
            v6 := GenS2{F0: (42 == -861), F1: 42}
            _ = len(v3)
            _ = v6
        } else {
            v7 := GenS1{F0: 42, F1: 42, F2: 0}
            _ = v7
        }
    } else {
        var v8 int16 = 1000
        v4 = "foo bar"
        _ = v8
    }
    v9 := GenS3{F0: 791, F1: -1}
    fmt.Println(v2)
    for fi10 := 0; fi10 < 20; fi10++ {
        var v11 int64 = 42
        if true {
            v12 := []int{fi10, 0}
            _ = v12
        } else {
            delete(v3, v4)
            if (!true) {
                if true {
                    v13 := GenS2{F0: ((42 < 868) || v2), F1: 65535}
                    v14 := GenS4{F0: (2 > 2)}
                    fmt.Println(fi10)
                    _ = v13
                    _ = v14
                } else {
                    var v15 int64 = (v11 * v11)
                    v2 = v2
                    fmt.Println(v4, v1)
                    _ = v15
                }
            } else {
                v4 = v4
                v4 = v4
                switch fi10 {
                case 0:
                    v16 := []int{(fi10 * 4294967295), 32767, fi10}
                    _ = v16
                case -9223372036854775808:
                    fmt.Println(fi10, v2, v4)
                case -662:
                    if v2 {
                        for fi17 := 0; fi17 < 18; fi17++ {
                            v1 = ("" + ("" + (((v1 + "hello") + (((v4 + v1) + v4) + "")) + "")))
                            _ = fi17
                        }
                        v4 = v1
                        if (false && v2) {
                            var v18 uint = 18446744073709551615
                            var v19 int16 = -1
                            if (!((3.14159 / 1.0) > (((3.14159 * 3.14159) / 1.5) - 0.0))) {
                                v20 := []int{fi10, (fi10 - fi10), (fi10 - (fi10 ^ (fi10 * (fi10 + fi10))))}
                                _ = v20
                            } else {
                                if v2 {
                                    if false {
                                        v21 := GenS2{F0: v2, F1: uint16(v18)}
                                        _ = v21
                                    } else {
                                        switch (fi10 >> ((v18) & 63)) {
                                        case 1:
                                            v22 := []int{}
                                            _ = v22
                                        }
                                        v23 := GenS2{F0: ((((!true) && (2147483647 <= 514)) && (65535 > 922)) || v2), F1: 65535}
                                        fmt.Println(v18, v4, v11)
                                        _ = v23
                                    }
                                    var v24 int8 = 42
                                    _ = v24
                                } else {
                                    v25 := GenS3{F0: v19, F1: -808}
                                    v11 += v11
                                    switch fi10 {
                                    case -2147483648:
                                        v18 = (v18 & (v18 * v18))
                                    case 4294967295:
                                        v26 := GenS5{F0: 874, F1: 2, F2: 0}
                                        _ = v26
                                    default:
                                        delete(v3, v4)
                                    }
                                    _ = v25
                                }
                            }
                            _ = v18
                            _ = v19
                        } else {
                            var v27 bool = v2
                            _ = v27
                        }
                    }
                default:
                    v4 = (v4 + (v1 + (v1 + "\n\t")))
                }
            }
        }
        _ = fi10
        _ = v11
    }
    var v28 int8 = 100
    v4 = v1
    if ("" > (v4 + (v4 + v1))) {
        switch (v1 + ((v4 + (((v4 + "hello") + ("foo bar" + v1)) + "hello")) + v1)) {
        case "a":
            fmt.Println(v28, v1)
        case "hello":
            v4 = ("" + "\n\t")
        case "":
            var v29 bool = v2
            _ = v29
        default:
            v28 += (v28 ^ (-(v28)))
        }
    } else {
        for fi30 := 0; fi30 < 20; fi30++ {
            if (!v2) {
                if (375 != 640) {
                    fmt.Println(v1)
                } else {
                    for fi31 := 0; fi31 < 8; fi31++ {
                        var v32 uint16 = 135
                        v4 = ("" + ((((("" + v4) + v1) + v4) + v4) + "foo bar"))
                        _ = fi31
                        _ = v32
                    }
                }
                var v33 int64 = -9223372036854775808
                _ = v33
            }
            switch -9223372036854775808 {
            case 2147483647:
                v28 = (v28 * -1)
            case -677:
                switch fi30 {
                case 4294967295:
                    fmt.Println(v28, v4)
                case -1:
                    var v34 uint = 18446744073709551615
                    _ = v34
                }
            }
            _ = fi30
        }
        for fi35 := 0; fi35 < 5; fi35++ {
            var v36 uint64 = 42
            var v37 float32 = -2.25
            _ = fi35
            _ = v36
            _ = v37
        }
    }
    v1 = "foo bar"
    v2 = v2
    delete(v3, "a")
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
