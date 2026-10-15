package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 uint16 }
type GenS2 struct { F0 int64; F1 string }
type GenS3 struct { F0 int; F1 uint16; F2 uint16 }
type GenS4 struct { F0 int8; F1 uint32; F2 int8 }
type GenS5 struct { F0 int8; F1 int8 }
type GenS6 struct { F0 uint32; F1 string }
func genFn0() int8 {
    var v4 float64 = -2.25
    for fi5 := 0; fi5 < 10; fi5++ {
        v6 := []int{}
        _ = fi5
        _ = v6
    }
    v7 := GenS1{F0: 42, F1: -1, F2: 255}
    fmt.Println(v4)
    _ = v4
    _ = v7
    return 0
}

func main() {
    v1 := []string{"hello", "a"}
    var v2 int8 = -128
    var v3 int16 = 90
    for fi8 := 0; fi8 < 10; fi8++ {
        if ((0.1 <= 1e10) || (294 < 192)) {
            v9 := GenS1{F0: 4294967295, F1: 32767, F2: 42}
            _ = v9
        }
        _ = fi8
    }
    var v10 int16 = 42
    var v11 uint8 = 42
    if (!false) {
        v3 = 32767
    }
    if (56 != 0) {
        if ((1 != (^(v11))) && (true && (330 <= -9223372036854775808))) {
            var v12 uint32 = 2147483648
            if false {
                fmt.Println(v12, v10)
                if (!true) {
                    fmt.Println(v10, v12, v10)
                    if true {
                        for fi13 := 0; fi13 < 9; fi13++ {
                            v14 := GenS1{F0: 868, F1: (v3 % (((-(v10))) | 1)), F2: 255}
                            _ = fi13
                            _ = v14
                        }
                    } else {
                        v11 += v11
                    }
                }
            }
            _ = v12
        } else {
            v10 = 539
            v3 = 1000
        }
        for fi15 := 0; fi15 < 10; fi15++ {
            switch "" {
            case "hello":
                switch fi15 {
                case -295:
                    var v16 string = ("\n\t" + (((("hello" + "") + ("foo bar" + "\n\t")) + "hello") + ""))
                    _ = v16
                case 128:
                    var v17 uint = 132
                    _ = v17
                }
            case "\n\t":
                v18 := GenS1{F0: 42, F1: v10, F2: 32767}
                _ = v18
            case "":
                if false {
                    var v19 float64 = (1.5 * (0.1 + 0.1))
                    _ = v19
                }
            default:
                v20 := GenS2{F0: 696, F1: (("foo bar" + "") + ("" + "\n\t"))}
                _ = v20
            }
            for fi21 := 0; fi21 < 7; fi21++ {
                v3 = v3
                var v22 int16 = (^(v10))
                _ = fi21
                _ = v22
            }
            if (((434 == 1) || (!(v2 >= (v2 * (v2 % ((int8(v3)) | 1)))))) || (v11 == v11)) {
                fmt.Println(v10, v2, v2)
                if (fi15 == (fi15 * fi15)) {
                    var v23 int16 = (^(v3))
                    _ = v23
                }
            }
            _ = fi15
        }
        if (42 < v10) {
            switch -2147483648 {
            case 127:
                v24 := GenS3{F0: -685, F1: 2, F2: 313}
                _ = v24
            }
            v1 = append(v1, ("" + "hello"))
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        } else {
            if (!true) {
                for fi25 := 0; fi25 < 11; fi25++ {
                    if true {
                        for fi26 := 0; fi26 < 6; fi26++ {
                            var v27 uint = 123
                            if len(v1) > 0 {
                                v1[0] = "\n\t"
                            }
                            v27 = v27
                            _ = fi26
                            _ = v27
                        }
                    } else {
                        v3 += (v3 % (((v10 - v10)) | 1))
                    }
                    _ = fi25
                }
                if true {
                    _ = len(v1)
                } else {
                    if len(v1) > 0 {
                        v1[0] = ""
                    }
                    var v28 float64 = 3.14159
                    fmt.Println(v28, v10, v10)
                    _ = v28
                }
            }
        }
    }
    if false {
        var v29 float64 = (((1.5 * 1e10) / (0.1 - -2.25)) - -2.25)
        _ = v29
    }
    var v30 int32 = -2147483648
    v31 := GenS4{F0: (v2 ^ (v2 % ((genFn0()) | 1))), F1: 289, F2: v2}
    v11 = (v11 << ((255) & 63))
    var v32 int64 = -9223372036854775808
    v11 = v11
    v2 += genFn0()
    v33 := GenS5{F0: int8(v11), F1: genFn0()}
    switch int(v2) {
    case 955:
        var v34 string = "a"
        _ = v34
    case 4294967295:
        v35 := GenS6{F0: 0, F1: "hello"}
        _ = v35
    case 128:
        _ = len(v1)
    }
    var v36 int = -9223372036854775808
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v32:", v32)
    fmt.Println("v36:", v36)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v33
    _ = v36
}
