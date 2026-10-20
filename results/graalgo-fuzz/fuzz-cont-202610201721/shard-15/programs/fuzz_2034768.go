package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16 }
type GenS2 struct { F0 int16; F1 int32; F2 bool }
type GenS3 struct { F0 uint8; F1 uint16; F2 uint64 }
type GenS4 struct { F0 uint32; F1 int; F2 int32 }
type GenS5 struct { F0 uint8; F1 int32 }
func genFn0(p0 int64, p1 uint64, p2 string) int8 {
    v5 := GenS1{F0: int(p1), F1: -1}
    if (42 <= (p1 & (p1 >> ((8) & 63)))) {
        v6 := GenS1{F0: 0, F1: -164}
        var v7 float32 = 3.14159
        p0 -= p0
        _ = v6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 127
}
func genFn1() int {
    var v8 float32 = (float32((0.1 - (0.1 - 1.5))) + (-2.25 * ((float32(0.1) - 1.5) * (3.14159 * ((0.1 / 1.0) + (-2.25 / 1.0))))))
    if true {
        fmt.Println(v8)
        v9 := []int{}
        _ = v9
    } else {
        if ((100 >= 1) || true) {
            var v10 bool = (!(945 > -1))
            _ = v10
        }
        for fi11 := 0; fi11 < 20; fi11++ {
            var v12 float64 = 0.0
            _ = fi11
            _ = v12
        }
        if true {
            if ("a" == "") {
                switch -1 {
                case -710:
                    v13 := []int{-2147483648, -2147483648, -14, 9223372036854775807}
                    _ = v13
                }
                var v14 uint16 = 0
                _ = v14
            }
        }
    }
    fmt.Println(v8)
    v15 := GenS2{F0: 0, F1: -390, F2: true}
    _ = v15
    _ = v8
    return -2147483648
}

func main() {
    var v1 int32 = 42
    var v2 int16 = -421
    v3 := []string{(("" + "hello") + ""), ("foo bar" + ("" + ("" + "a"))), "foo bar"}
    v4 := []string{"foo bar"}
    v1 -= v1
    v16 := GenS2{F0: -32768, F1: -609, F2: false}
    if len(v4) > 0 {
        v4[0] = (("a" + "a") + ("\n\t" + ""))
    }
    v17 := GenS3{F0: 0, F1: 0, F2: 4294967295}
    v1 = (v1 & (v1 + (v1 ^ v1)))
    var v18 uint8 = 128
    for fi19 := 0; fi19 < 15; fi19++ {
        var v20 string = "foo bar"
        if (-2.25 < (-2.25 / ((1.5 - 1.5) + 1.5))) {
            if (!false) {
                var v21 uint8 = v18
                var v22 float64 = 1.5
                fmt.Println(v21, v18)
                _ = v21
                _ = v22
            } else {
                for fi23 := 0; fi23 < 20; fi23++ {
                    if len(v4) > 0 {
                        v4[0] = "a"
                    }
                    if false {
                        v18 = (v18 >> ((42) & 63))
                        var v24 string = v20
                        _ = v24
                    }
                    v2 -= (v2 >> ((346) & 63))
                    _ = fi23
                }
                v2 += (v2 >> ((65535) & 63))
            }
            v3 = append(v3, "")
        }
        v2 = (v2 & 266)
        _ = fi19
        _ = v20
    }
    fmt.Println(v2, v2)
    var v25 uint = 652
    v26 := GenS4{F0: 2, F1: -9223372036854775808, F2: (v1 & (v1 << ((255) & 63)))}
    var v27 uint = v25
    if ((v25 & 0) > 18446744073709551615) {
        if true {
            v28 := GenS5{F0: (v18 << ((v27) & 63)), F1: 0}
            _ = v28
        } else {
            switch -742 {
            case 9223372036854775807:
                fmt.Println(v25, v25, v2)
            default:
                switch 4294967295 {
                case 722:
                    var v29 uint8 = 0
                    _ = v29
                case 4294967295:
                    switch ("a" + "hello") {
                    case "foo bar":
                        for fi30 := 0; fi30 < 7; fi30++ {
                            for fi31 := 0; fi31 < 1; fi31++ {
                                v25 = (v25 << (((^(v25))) & 63))
                                if (!(int(v18) < fi30)) {
                                    var v32 float32 = (1.5 / 1e10)
                                    var v33 uint8 = (v18 ^ v18)
                                    _ = v32
                                    _ = v33
                                }
                                var v34 uint32 = 2147483648
                                _ = fi31
                                _ = v34
                            }
                            var v35 float64 = 1e10
                            _ = fi30
                            _ = v35
                        }
                    }
                default:
                    var v36 int16 = 1
                    _ = v36
                }
            }
            var v37 uint16 = 255
            _ = v37
        }
        if false {
            var v38 int16 = (v2 | (v2 - 0))
            _ = v38
        }
        var v39 float32 = 0.0
        _ = v39
    }
    v40 := GenS3{F0: v18, F1: 255, F2: 277}
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v40
}
