package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint16 }
func genFn0() int64 {
    var v5 int64 = -9223372036854775808
    v6 := []int{-943, 1, -97, 657}
    if false {
        fmt.Println(v5)
    } else {
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        if (0 <= 42) {
            if false {
                v7 := GenS1{F0: -1, F1: "hello"}
                var v8 int64 = (^(v5))
                fmt.Println(v5, v5)
                _ = v7
                _ = v8
            }
        } else {
            if (!(930 > 17)) {
                var v9 uint16 = 32767
                _ = v9
            } else {
                var v10 int8 = -128
                var v11 string = ("hello" + ("\n\t" + (((("hello" + ("foo bar" + "")) + "a") + "a") + "hello")))
                _ = v10
                _ = v11
            }
            for fi12 := 0; fi12 < 11; fi12++ {
                for fi13 := 0; fi13 < 1; fi13++ {
                    var v14 float64 = (0.0 * 0.0)
                    _ = fi13
                    _ = v14
                }
                var v15 uint64 = 2
                _ = fi12
                _ = v15
            }
        }
    }
    _ = v5
    _ = v6
    return (v5 % ((v5) | 1))
}

func main() {
    var v1 uint32 = 276
    v2 := make(map[int]bool)
    v3 := []uint32{v1, (v1 * v1)}
    v4 := make(map[string]bool)
    var v16 int8 = -85
    for fi17 := 0; fi17 < 20; fi17++ {
        var v18 float64 = 0.1
        _ = fi17
        _ = v18
    }
    var v19 uint8 = 128
    if (!(!false)) {
        v16 += 100
    }
    _ = len(v4)
    switch 9223372036854775807 {
    case 4294967295:
        v4["foo bar"] = (false && ("hello" == "foo bar"))
    default:
        switch ("foo bar" + "a") {
        case "foo bar":
            if (true && false) {
                v16 = -31
                v20 := GenS1{F0: 767, F1: ("foo bar" + "a")}
                v21 := GenS1{F0: -532, F1: "hello"}
                _ = v20
                _ = v21
            } else {
                for fi22 := 0; fi22 < 9; fi22++ {
                    if ((true && false) || (fi22 < fi22)) {
                        fmt.Println(v16)
                        var v23 uint8 = (v19 >> ((533) & 63))
                        _ = v23
                    } else {
                        var v24 float64 = (1.5 - 1.5)
                        _ = len(v3)
                        for fi25 := 0; fi25 < 12; fi25++ {
                            fmt.Println(v24, v19, v16)
                            _ = len(v3)
                            v1 += (v1 / ((v1) | 1))
                            _ = fi25
                        }
                        _ = v24
                    }
                    var v26 float64 = (0.1 - float64((-2.25 / (((1.5 + 1.5) / 3.14159) - (((0.1 * -2.25) + 0.1) - (((((0.1 * ((1.5 / 1.5) - (1.5 + (0.0 * 3.14159)))) + (-2.25 * 1e10)) - 1.5) * 0.1) * 1.5))))))
                    _ = fi22
                    _ = v26
                }
            }
        case "a":
            switch "foo bar" {
            case "hello":
                if true {
                    var v27 bool = true
                    _ = v27
                }
            case "\n\t":
                fmt.Println(v19)
            }
        }
    }
    for fi28 := 0; fi28 < 7; fi28++ {
        _ = len(v4)
        v1 = v1
        for fi29 := 0; fi29 < 10; fi29++ {
            if true {
                if (-2147483648 == 42) {
                    switch "" {
                    case "foo bar":
                        v30 := GenS2{F0: 32767, F1: v16}
                        _ = v30
                    case "a":
                        if true {
                            v3 = append(v3, v1)
                            fmt.Println(fi28, fi28, v16)
                        } else {
                            var v31 int32 = -440
                            _ = v31
                        }
                    case "hello":
                        if (434 <= 4294967295) {
                            v32 := GenS3{F0: "\n\t"}
                            _ = v32
                        } else {
                            v33 := GenS4{F0: 32767}
                            _ = v33
                        }
                    }
                    if (42 > 460) {
                        switch "" {
                        case "hello":
                            fmt.Println(v1, fi28)
                        }
                        if (!false) {
                            if false {
                                v1 -= (v1 * 869)
                                if (!(int16(v16) >= -1)) {
                                    v34 := GenS2{F0: 532, F1: v16}
                                    v2[(fi28 & (^(fi29)))] = true
                                    var v35 int64 = 21
                                    _ = v34
                                    _ = v35
                                }
                                v36 := GenS2{F0: 65535, F1: v16}
                                _ = v36
                            }
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                        } else {
                            v19 -= (^(v19))
                        }
                        var v37 uint16 = 255
                        _ = v37
                    }
                }
            }
            v19 -= (^(v19))
            _ = fi29
        }
        _ = fi28
    }
    v3 = append(v3, (v1 | (v1 * (v1 * (v1 & (v1 % ((v1) | 1)))))))
    for fi38 := 0; fi38 < 3; fi38++ {
        switch (fi38 + fi38) {
        case 771:
            var v39 uint32 = v1
            _ = v39
        case -2147483648:
            var v40 uint16 = 100
            _ = v40
        case -550:
            fmt.Println(v1)
        default:
            _ = len(v2)
        }
        var v41 int64 = 1
        fmt.Println(v1)
        _ = fi38
        _ = v41
    }
    if ((false || false) && (255 >= 877)) {
        v42 := GenS1{F0: -983, F1: "a"}
        v3 = append(v3, v1)
        v43 := GenS4{F0: 32767}
        _ = v42
        _ = v43
    }
    fmt.Println(v19, v1, v1)
    v16 = v16
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v16
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
