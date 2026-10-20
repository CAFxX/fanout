package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int64; F2 uint16 }
type GenS2 struct { F0 uint8; F1 uint64; F2 bool }
type GenS3 struct { F0 uint16 }
func genFn0(p0 bool, p1 int16) uint {
    var v7 uint = 968
    if (!((("hello" + ("hello" + ("\n\t" + ((("a" + "a") + ("" + (("" + (("\n\t" + (("foo bar" + "") + "foo bar")) + "foo bar")) + ("" + ("\n\t" + "hello"))))) + "\n\t")))) + "foo bar") <= "a")) {
        v7 += (v7 * (v7 / ((4294967295) | 1)))
    }
    var v8 int16 = p1
    var v9 uint8 = 232
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    _ = v9
    return v7
}
func genFn1(p0 string, p1 uint32) uint {
    fmt.Println(p1, p0)
    v10 := []int{-1, -2147483648, 4294967295, 42}
    var v11 uint64 = 42
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    return 23
}

func main() {
    var v1 bool = (128 <= 347)
    var v2 uint64 = 2
    var v3 uint16 = 2
    v4 := make(map[uint32]string)
    v5 := make(map[string]string)
    v6 := make(map[string]bool)
    v12 := GenS1{F0: -312, F1: 273, F2: (v3 << ((uint(v3)) & 63))}
    if v1 {
        for fi13 := 0; fi13 < 16; fi13++ {
            v14 := GenS1{F0: -507, F1: 804, F2: (v3 << ((1) & 63))}
            v3 -= (v3 * 255)
            if true {
                v15 := []int{fi13, int(v2), -1}
                _ = v15
            } else {
                v16 := GenS1{F0: 0, F1: 42, F2: v3}
                if v1 {
                    var v17 bool = (-2147483648 >= 42)
                    _ = v17
                } else {
                    var v18 int64 = 242
                    _ = v18
                }
                _ = v16
            }
            _ = fi13
            _ = v14
        }
    } else {
        var v19 uint16 = v3
        _ = v19
    }
    if false {
        var v20 bool = v1
        v20 = (!(!v1))
        _ = v20
    }
    v1 = (false || (v3 == (v3 >> ((4294967295) & 63))))
    v21 := GenS2{F0: 42, F1: v2, F2: v1}
    if (((("\n\t" + (("foo bar" + "foo bar") + "")) >= "hello") && ((v2 + v2) < (^(v2)))) && v1) {
        fmt.Println(v2, v1)
        if (1000 > -448) {
            if v1 {
                var v22 uint32 = 1
                v23 := GenS3{F0: v3}
                _ = v22
                _ = v23
            }
            for fi24 := 0; fi24 < 9; fi24++ {
                switch ("foo bar" + "") {
                case "":
                    v25 := GenS2{F0: uint8(v2), F1: (v2 & 4294967295), F2: v1}
                    _ = v25
                default:
                    switch 128 {
                    case -9223372036854775808:
                        for fi26 := 0; fi26 < 4; fi26++ {
                            fmt.Println(fi26, v3, v3)
                            if (v1 || false) {
                                var v27 bool = true
                                v28 := GenS3{F0: v3}
                                v29 := GenS2{F0: uint8(v2), F1: (v2 % (((v2 % (((v2 / ((v2) | 1))) | 1))) | 1)), F2: false}
                                v30 := v6[("\n\t" + "foo bar")]
                                _ = v27
                                _ = v28
                                _ = v29
                                _ = v30
                            }
                            var v31 uint32 = 945
                            _ = fi26
                            _ = v31
                        }
                    }
                }
                if v1 {
                    if (!(!v1)) {
                        var v32 float64 = 0.1
                        _ = v32
                    }
                } else {
                    v1 = v1
                    var v33 uint16 = v3
                    delete(v5, "foo bar")
                    _ = v33
                }
                var v34 string = "hello"
                _ = fi24
                _ = v34
            }
            if ((!true) || (("" + ("foo bar" + ("a" + "\n\t"))) == "")) {
                fmt.Println(v2, v2)
            } else {
                var v35 string = ("\n\t" + "")
                if (1000 <= 32767) {
                    v36 := []int{}
                    for fi37 := 0; fi37 < 12; fi37++ {
                        _ = len(v36)
                        if v1 {
                            if true {
                                if v1 {
                                    v36 = append(v36, (fi37 + (fi37 / (((fi37 & (fi37 * fi37))) | 1))))
                                    for fi38 := 0; fi38 < 20; fi38++ {
                                        var v39 uint32 = 2147483647
                                        _ = fi38
                                        _ = v39
                                    }
                                    fmt.Println(v3)
                                }
                            } else {
                                var v40 string = v35
                                var v41 float32 = 0.0
                                _ = v40
                                _ = v41
                            }
                        }
                        _ = fi37
                    }
                    var v42 int8 = -30
                    _ = v36
                    _ = v42
                }
                var v43 float32 = 0.0
                _ = v35
                _ = v43
            }
        } else {
            switch "a" {
            case "hello":
                var v44 int32 = 2147483647
                _ = v44
            case "":
                var v45 bool = v1
                _ = v45
            }
            switch -975 {
            case -2147483648:
                if false {
                    v46 := GenS3{F0: 42}
                    _ = v46
                }
            case 4294967295:
                v47 := GenS3{F0: (v3 ^ (v3 >> ((414) & 63)))}
                _ = v47
            default:
                v1 = (!v1)
            }
        }
    } else {
        fmt.Println(v1, v1)
        v3 = 88
    }
    v2 = v2
    v3 = (v3 | v3)
    var v48 bool = v1
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v48:", v48)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v12
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v48
    _ = v5
    _ = v6
}
