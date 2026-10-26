package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint8; F1 uint32; F2 int8 }
type GenS5 struct { F0 uint64; F1 int; F2 uint }
func genFn0() uint8 {
    for fi4 := 0; fi4 < 20; fi4++ {
        switch (^(fi4)) {
        case 32767:
            if (!(246 >= 255)) {
                if false {
                    var v5 uint16 = 278
                    var v6 float32 = (0.1 * float32((1.5 - float64(0.0))))
                    var v7 string = "foo bar"
                    _ = v5
                    _ = v6
                    _ = v7
                }
                v8 := GenS1{F0: 0}
                var v9 int = fi4
                _ = v8
                _ = v9
            }
        case -627:
            if true {
                fmt.Println(fi4)
                var v10 int16 = 305
                _ = v10
            }
        }
        v11 := []int{(fi4 - fi4), 209, fi4}
        _ = fi4
        _ = v11
    }
    for fi12 := 0; fi12 < 4; fi12++ {
        for fi13 := 0; fi13 < 19; fi13++ {
            fmt.Println(fi13, fi13)
            _ = fi13
        }
        _ = fi12
    }
    return 7
}
func genFn1() uint16 {
    switch ((("hello" + "foo bar") + "\n\t") + "hello") {
    case "hello":
        var v14 uint16 = 65535
        _ = v14
    default:
        v15 := GenS2{F0: 2147483648}
        _ = v15
    }
    switch "hello" {
    case "":
        if (0 == 729) {
            var v16 int8 = -112
            v16 = v16
            v16 = (^(v16))
            _ = v16
        }
    }
    return 76
}

func main() {
    var v1 int32 = -1
    v2 := make(map[string]string)
    var v3 string = "a"
    for fi17 := 0; fi17 < 5; fi17++ {
        fmt.Println(fi17, fi17)
        var v18 uint32 = 2
        _ = fi17
        _ = v18
    }
    v19 := []int{}
    switch -9223372036854775808 {
    case 42:
        v20 := GenS3{F0: "hello"}
        _ = v20
    default:
        var v21 int64 = 0
        _ = v21
    }
    var v22 float32 = ((0.1 * float32(0.0)) + (0.0 * (0.0 / 1.0)))
    v23 := GenS2{F0: 4294967295}
    var v24 float64 = float64((v22 + ((((3.14159 - float32(0.0)) * 0.0) / 1.5) + v22)))
    delete(v2, "hello")
    if (true || true) {
        for fi25 := 0; fi25 < 8; fi25++ {
            for fi26 := 0; fi26 < 15; fi26++ {
                v2[("foo bar" + "foo bar")] = (("a" + "\n\t") + v3)
                if true {
                    _ = len(v19)
                    fmt.Println(v1, v1, v24)
                    v24 += 3.14159
                    v1 = v1
                    if (!false) {
                        if (-1 > -128) {
                            v2[("foo bar" + v3)] = v3
                        }
                    } else {
                        if true {
                            v27 := v2[(((((("a" + ((v3 + (("foo bar" + v3) + ("hello" + ((v3 + v3) + ("hello" + "foo bar"))))) + "\n\t")) + (("a" + v3) + v3)) + (v3 + (((((v3 + v3) + v3) + v3) + v3) + ((((v3 + ("" + v3)) + "\n\t") + v3) + v3)))) + "a") + (v3 + "foo bar")) + "hello")]
                            var v28 int32 = (v1 / ((-968) | 1))
                            if true {
                                v29 := GenS4{F0: 36, F1: 65535, F2: 111}
                                if (2 >= 55) {
                                    if len(v19) > 0 {
                                        v19[0] = 2147483647
                                    }
                                } else {
                                    var v30 uint64 = 418
                                    v30 = (v30 + (v30 | 1))
                                    if len(v19) > 0 {
                                        v19[0] = -2147483648
                                    }
                                    _ = v30
                                }
                                var v31 uint64 = 18446744073709551615
                                _ = v29
                                _ = v31
                            }
                            _ = v27
                            _ = v28
                        }
                        v32 := GenS1{F0: 4294967295}
                        switch fi26 {
                        case -529:
                            var v33 uint8 = 255
                            _ = v33
                        case 1:
                            v2["foo bar"] = v3
                        case -1:
                            v19 = append(v19, fi25)
                        default:
                            fmt.Println(fi25, fi26, v1)
                        }
                        _ = v32
                    }
                } else {
                    var v34 float32 = (3.14159 - v22)
                    var v35 int16 = 42
                    fmt.Println(fi26)
                    _ = v34
                    _ = v35
                }
                _ = fi26
            }
            for fi36 := 0; fi36 < 3; fi36++ {
                var v37 float64 = float64((v22 * 0.0))
                _ = fi36
                _ = v37
            }
            v38 := v2[""]
            _ = fi25
            _ = v38
        }
        var v39 uint16 = 754
        _ = len(v2)
        _ = v39
    }
    fmt.Println(v3, v1, v24)
    var v40 int64 = -794
    for fi41 := 0; fi41 < 14; fi41++ {
        var v42 int = fi41
        _ = fi41
        _ = v42
    }
    if false {
        v3 = (v3 + ("foo bar" + ("foo bar" + (((v3 + "foo bar") + ((((v3 + v3) + "hello") + "") + v3)) + "hello"))))
        fmt.Println(v22)
    } else {
        for fi43 := 0; fi43 < 19; fi43++ {
            var v44 float64 = v24
            if (false || false) {
                var v45 int16 = 160
                if ((!false) && false) {
                    switch 127 {
                    case -789:
                        if len(v19) >= 2 {
                            v19 = v19[1:len(v19)]
                        }
                    case -856:
                        if len(v19) > 0 {
                            v19[0] = fi43
                        }
                    case 127:
                        v46 := GenS5{F0: 42, F1: 128, F2: 1}
                        _ = v46
                    default:
                        var v47 int32 = v1
                        _ = v47
                    }
                }
                var v48 int = 1
                _ = v45
                _ = v48
            } else {
                v3 = "hello"
                var v49 int32 = (v1 - v1)
                _ = v49
            }
            _ = fi43
            _ = v44
        }
        for fi50 := 0; fi50 < 18; fi50++ {
            fmt.Println(fi50, v1)
            _ = fi50
        }
        v40 -= -1
    }
    v51 := GenS5{F0: 482, F1: -1, F2: 65535}
    v24 = v24
    v3 = v3
    _ = len(v2)
    var v52 uint32 = 559
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("v40:", v40)
    fmt.Println("v52:", v52)
    _ = v1
    _ = v19
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v40
    _ = v51
    _ = v52
}
