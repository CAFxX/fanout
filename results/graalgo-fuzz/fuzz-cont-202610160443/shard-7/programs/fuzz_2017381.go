package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 string; F1 bool; F2 int8 }
type GenS4 struct { F0 int8; F1 uint; F2 uint16 }
type GenS5 struct { F0 int }

func main() {
    v1 := []int{-2147483648}
    var v2 uint64 = 18446744073709551615
    var v3 float64 = 0.0
    v4 := make(map[int]int)
    if (true || true) {
        for fi5 := 0; fi5 < 5; fi5++ {
            for fi6 := 0; fi6 < 11; fi6++ {
                var v7 uint64 = v2
                _ = fi6
                _ = v7
            }
            fmt.Println(v3, v3, v3)
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = fi5
        }
    }
    v8 := GenS1{F0: "\n\t"}
    var v9 int16 = -1
    var v10 uint16 = 1
    v2 = (^(v2))
    var v11 uint16 = v10
    var v12 uint8 = 1
    if (1e10 == (float64(0.1) / v3)) {
        v4[1] = 0
        for fi13 := 0; fi13 < 14; fi13++ {
            var v14 int = (fi13 | fi13)
            _ = fi13
            _ = v14
        }
        if (("" != (("\n\t" + "hello") + "hello")) && (!true)) {
            if false {
                v3 -= (v3 + ((-2.25 + ((3.14159 / v3) + v3)) / ((v3 / v3) / ((v3 + v3) - 1e10))))
            }
            for fi15 := 0; fi15 < 1; fi15++ {
                v16 := GenS2{F0: 92}
                _ = fi15
                _ = v16
            }
        }
    } else {
        var v17 int64 = 0
        if (951 < 508) {
            if (-67 == 100) {
                var v18 int16 = (v9 ^ 32767)
                _ = v18
            }
            delete(v4, 2147483647)
        }
        _ = v17
    }
    v19 := GenS1{F0: ("foo bar" + "\n\t")}
    v20 := GenS3{F0: "", F1: false, F2: 2}
    v12 -= 2
    switch ("" + "") {
    case "foo bar":
        var v21 string = "a"
        _ = v21
    default:
        if (57 != 42) {
            for fi22 := 0; fi22 < 19; fi22++ {
                for fi23 := 0; fi23 < 20; fi23++ {
                    if ((!true) || (-569 < -9223372036854775808)) {
                        switch ("hello" + ("" + "foo bar")) {
                        case "a":
                            var v24 int = fi22
                            _ = v24
                        case "":
                            var v25 uint64 = v2
                            var v26 int16 = (v9 % ((v9) | 1))
                            _ = v25
                            _ = v26
                        default:
                            if true {
                                v27 := GenS4{F0: 1, F1: 361, F2: 324}
                                var v28 bool = (-43 < -304)
                                switch (("" + (("hello" + "hello") + ("hello" + "a"))) + "\n\t") {
                                case "\n\t":
                                    v10 = v11
                                case "":
                                    fmt.Println(fi22, v2, fi23)
                                }
                                _ = v27
                                _ = v28
                            }
                        }
                        if ((v9 >= 42) && false) {
                            switch "foo bar" {
                            case "\n\t":
                                if true {
                                    if len(v1) > 0 {
                                        v1[0] = -30
                                    }
                                    v9 = 785
                                    var v29 uint64 = 0
                                    _ = v29
                                }
                            }
                            if (!false) {
                                v30 := GenS5{F0: fi23}
                                v12 -= (v12 & v12)
                                _ = v30
                            } else {
                                var v31 uint16 = (v11 % ((v11) | 1))
                                _ = v31
                            }
                        } else {
                            var v32 int32 = 1
                            v11 -= (v10 % (((v11 - (v10 << ((65535) & 63)))) | 1))
                            v32 -= v32
                            _ = v32
                        }
                        var v33 bool = true
                        var v34 uint8 = uint8(v11)
                        switch (-(fi23)) {
                        case 4294967295:
                            var v35 int64 = 9223372036854775807
                            _ = v35
                        default:
                            fmt.Println(v34)
                        }
                        _ = v33
                        _ = v34
                    }
                    var v36 int16 = 1
                    _ = len(v4)
                    _ = fi23
                    _ = v36
                }
                for fi37 := 0; fi37 < 19; fi37++ {
                    v10 = (v11 / (((v11 << ((356) & 63))) | 1))
                    if len(v1) > 0 {
                        v1[0] = int(v12)
                    }
                    _ = fi37
                }
                _ = fi22
            }
            var v38 int8 = 0
            v38 = v38
            _ = v38
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
