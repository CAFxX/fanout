package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32; F1 int64 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint32 }
func genFn0() int8 {
    v5 := []int{39, -9223372036854775808, 1, 196}
    var v6 float32 = 0.1
    var v7 int32 = 42
    _ = v5
    _ = v6
    _ = v7
    return -1
}

func main() {
    var v1 int64 = -181
    var v2 bool = false
    v3 := make(map[uint32]string)
    v4 := make(map[string]bool)
    fmt.Println(v1, v2)
    var v8 int16 = 42
    if (65535 >= 0) {
        v2 = v2
    } else {
        if (4294967295 < 42) {
            var v9 int32 = -2147483648
            v10 := []int{1, -413, 4294967295}
            _ = v10
            _ = v9
        }
    }
    fmt.Println(v1, v1, v1)
    switch 32767 {
    case 0:
        if (!v2) {
            v11 := v4["foo bar"]
            v12 := []int{32767, 4294967295, 1, 494}
            for fi13 := 0; fi13 < 17; fi13++ {
                var v14 uint16 = 1
                _ = fi13
                _ = v14
            }
            _ = v11
            _ = v12
        }
    }
    var v15 float32 = ((float32((((-2.25 / float64(float32(3.14159))) + (-2.25 * 0.0)) / (1e10 - 3.14159))) * -2.25) * 0.1)
    switch 42 {
    case -452:
        for fi16 := 0; fi16 < 15; fi16++ {
            v1 = v1
            _ = fi16
        }
    default:
        v17 := GenS1{F0: 1}
        _ = v17
    }
    for fi18 := 0; fi18 < 11; fi18++ {
        if false {
            v19 := GenS2{F0: 2147483648, F1: (v1 / (((v1 % ((v1) | 1))) | 1))}
            var v20 uint32 = 2147483647
            if ((v1 + (v1 ^ (v1 - 9223372036854775807))) <= v1) {
                v21 := GenS1{F0: (fi18 % ((fi18) | 1))}
                fmt.Println(v15)
                if ("" == (("foo bar" + ("a" + "foo bar")) + "")) {
                    v22 := GenS2{F0: (v20 + (v20 & v20)), F1: -1}
                    v1 -= v1
                    _ = v22
                } else {
                    v23 := []int{fi18, 2147483647}
                    _ = v23
                }
                _ = v21
            }
            _ = v19
            _ = v20
        } else {
            if (921 > 757) {
                if (v15 > ((1.5 - float32(0.0)) / v15)) {
                    v8 -= -32768
                } else {
                    v1 = 9223372036854775807
                    v1 = -791
                    var v24 uint8 = 255
                    _ = v24
                }
            }
            switch "" {
            case "a":
                v8 += (v8 << ((4294967295) & 63))
            case "\n\t":
                v25 := v3[448]
                _ = v25
            case "foo bar":
                fmt.Println(v1, v8)
            }
            switch "" {
            case "a":
                if v2 {
                    _ = len(v3)
                    v8 = v8
                    var v26 uint8 = 0
                    _ = v26
                }
            case "\n\t":
                for fi27 := 0; fi27 < 20; fi27++ {
                    v2 = true
                    switch "a" {
                    case "":
                        var v28 string = "a"
                        _ = v28
                    case "foo bar":
                        if v2 {
                            v1 = v1
                        } else {
                            var v29 uint16 = 181
                            var v30 uint16 = 1
                            v3[2] = "a"
                            if true {
                                var v31 string = (("" + ("" + "foo bar")) + ((("" + ("hello" + (("a" + "hello") + "foo bar"))) + "") + (("hello" + "") + "\n\t")))
                                delete(v4, "hello")
                                v4[v31] = v2
                                _ = v31
                            } else {
                                if ("\n\t" == "") {
                                    fmt.Println(v29, v29, v30)
                                    v32 := []int{}
                                    _ = v32
                                } else {
                                    var v33 int64 = (v1 >> ((4294967295) & 63))
                                    var v34 int8 = 86
                                    v35 := GenS3{F0: 127}
                                    _ = v33
                                    _ = v34
                                    _ = v35
                                }
                            }
                            _ = v29
                            _ = v30
                        }
                    default:
                        var v36 int32 = -2147483648
                        _ = v36
                    }
                    _ = fi27
                }
            case "hello":
                v37 := []int{}
                _ = v37
            default:
                for fi38 := 0; fi38 < 6; fi38++ {
                    fmt.Println(fi38, v1)
                    _ = fi38
                }
            }
        }
        _ = fi18
    }
    v39 := []int{}
    switch "foo bar" {
    case "hello":
        var v40 float32 = v15
        _ = v40
    case "\n\t":
        var v41 int64 = v1
        _ = v41
    }
    var v42 int8 = 0
    for fi43 := 0; fi43 < 5; fi43++ {
        for fi44 := 0; fi44 < 12; fi44++ {
            v45 := GenS4{F0: 2}
            _ = fi44
            _ = v45
        }
        v8 += v8
        _ = fi43
    }
    var v46 float32 = ((v15 - (v15 - v15)) + float32(3.14159))
    v3[4294967295] = ("foo bar" + "hello")
    if (2147483648 <= 2147483648) {
        var v47 uint32 = 0
        _ = v47
    }
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v39:", len(v39), v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v46:", v46)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v42
    _ = v46
    _ = v8
}
