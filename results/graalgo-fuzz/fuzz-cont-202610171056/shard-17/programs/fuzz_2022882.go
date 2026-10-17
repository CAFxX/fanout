package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 uint }
type GenS2 struct { F0 string; F1 int8; F2 bool }
type GenS3 struct { F0 bool; F1 uint16; F2 string }
func genFn0() int {
    v6 := []int{-9223372036854775808, 565}
    var v7 float64 = 0.1
    _ = v6
    _ = v7
    return -239
}

func main() {
    var v1 int64 = -693
    v2 := []int{500, 2147483647, 128, -2147483648}
    v3 := make(map[uint32]string)
    v4 := make(map[uint32]string)
    var v5 bool = true
    if (true || (4294967295 > 2)) {
        var v8 int8 = -23
        for fi9 := 0; fi9 < 14; fi9++ {
            v1 -= 164
            switch "a" {
            case "\n\t":
                for fi10 := 0; fi10 < 7; fi10++ {
                    var v11 uint64 = 42
                    var v12 string = (("\n\t" + ((("hello" + "\n\t") + "a") + "a")) + "foo bar")
                    _ = fi10
                    _ = v11
                    _ = v12
                }
            case "a":
                if true {
                    for fi13 := 0; fi13 < 4; fi13++ {
                        v8 -= 42
                        _ = fi13
                    }
                }
            }
            v5 = v5
            _ = fi9
        }
        _ = v8
    }
    var v14 float32 = ((float32(0.0) + 0.0) / -2.25)
    fmt.Println(v5, v14, v14)
    var v15 uint32 = 812
    for fi16 := 0; fi16 < 2; fi16++ {
        var v17 uint = 255
        v15 += (v15 - v15)
        v5 = (!v5)
        _ = fi16
        _ = v17
    }
    v5 = (!true)
    if (v1 <= (v1 >> ((2) & 63))) {
        for fi18 := 0; fi18 < 12; fi18++ {
            if (false || v5) {
                for fi19 := 0; fi19 < 19; fi19++ {
                    if v5 {
                        v20 := GenS1{F0: fi18, F1: (^(v15)), F2: 2}
                        _ = v20
                    }
                    v2 = append(v2, 32767)
                    _ = fi19
                }
                v1 -= 0
            }
            var v21 uint16 = 32767
            v1 = v1
            _ = fi18
            _ = v21
        }
    } else {
        v2 = append(v2, 854)
        switch 127 {
        case 9223372036854775807:
            for fi22 := 0; fi22 < 10; fi22++ {
                fmt.Println(fi22, v1)
                _ = fi22
            }
        case -1:
            switch -461 {
            case 629:
                v14 = (0.1 - -2.25)
            default:
                _ = len(v3)
            }
        }
        switch genFn0() {
        case 2147483647:
            var v23 string = ((("\n\t" + "\n\t") + "") + "a")
            _ = v23
        case 4294967295:
            _ = len(v3)
        case 32768:
            var v24 uint32 = v15
            _ = v24
        default:
            var v25 uint = 42
            _ = v25
        }
    }
    v5 = false
    _ = len(v3)
    v26 := GenS2{F0: "hello", F1: -1, F2: (1 >= int16(v1))}
    v27 := GenS2{F0: "a", F1: -1, F2: (2147483647 > 2147483647)}
    for fi28 := 0; fi28 < 16; fi28++ {
        v4[v15] = ""
        if (76 >= 127) {
            var v29 int8 = -1
            var v30 string = "\n\t"
            _ = v29
            _ = v30
        } else {
            for fi31 := 0; fi31 < 7; fi31++ {
                if v5 {
                    v32 := GenS3{F0: false, F1: 65535, F2: "\n\t"}
                    v1 -= (v1 & int64(fi28))
                    v15 -= (v15 * 2)
                    fmt.Println(fi28, v14, v5)
                    _ = v32
                }
                _ = fi31
            }
            if (fi28 >= -157) {
                for fi33 := 0; fi33 < 15; fi33++ {
                    if true {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        if false {
                            v14 = 0.1
                            var v34 bool = false
                            var v35 int16 = 882
                            var v36 string = "a"
                            v5 = (v15 <= 42)
                            _ = v34
                            _ = v35
                            _ = v36
                        } else {
                            if v5 {
                                var v37 uint8 = 2
                                var v38 int16 = -318
                                _ = v37
                                _ = v38
                            }
                            v14 += v14
                            v39 := GenS1{F0: int(v15), F1: v15, F2: 4294967295}
                            _ = v39
                        }
                        v15 = 4294967295
                    } else {
                        v14 = v14
                        v40 := GenS2{F0: "a", F1: -128, F2: false}
                        var v41 bool = (255 < 242)
                        _ = v40
                        _ = v41
                    }
                    _ = len(v3)
                    v1 -= -1
                    _ = fi33
                }
                switch "\n\t" {
                case "foo bar":
                    var v42 int16 = -243
                    _ = v42
                case "a":
                    if len(v2) > 0 {
                        v2[0] = (fi28 >> ((255) & 63))
                    }
                case "":
                    v15 = v15
                default:
                    v5 = (!v5)
                }
            }
        }
        var v43 float32 = v14
        _ = fi28
        _ = v43
    }
    fmt.Println(v15, v5)
    fmt.Println(v14, v5)
    var v44 int32 = -218
    v1 -= (v1 * (v1 | (v1 * v1)))
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v44
    _ = v5
}
