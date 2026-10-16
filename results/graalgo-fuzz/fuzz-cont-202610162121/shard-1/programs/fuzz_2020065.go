package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 string; F2 uint64 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint8; F1 uint8 }

func main() {
    var v1 float64 = 0.1
    v2 := make(map[uint32]string)
    var v3 int16 = -106
    v4 := []string{""}
    v5 := GenS1{F0: 4294967295, F1: ("\n\t" + ("\n\t" + "a")), F2: 42}
    delete(v2, 2)
    fmt.Println(v3)
    switch "foo bar" {
    case "a":
        if ((!false) && ((v1 * (v1 + float64(1.5))) != 3.14159)) {
            v3 = (v3 + (v3 ^ -252))
            v3 -= (v3 & (v3 >> ((0) & 63)))
            fmt.Println(v3, v3)
        } else {
            var v6 uint8 = 201
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = v6
        }
    }
    var v7 uint32 = 803
    v8 := GenS1{F0: 42, F1: (((("\n\t" + "hello") + "hello") + "a") + ""), F2: 42}
    v3 = v3
    var v9 uint = 732
    if true {
        v3 = (v3 ^ (v3 & v3))
        if (1 > -1) {
            if ("foo bar" < "a") {
                for fi10 := 0; fi10 < 18; fi10++ {
                    if (!true) {
                        if true {
                            v3 -= (v3 % ((-208) | 1))
                        }
                        v9 = v9
                    } else {
                        switch (fi10 & fi10) {
                        case 0:
                            fmt.Println(v7)
                        case 806:
                            fmt.Println(v7)
                        default:
                            v11 := GenS1{F0: 18446744073709551615, F1: "a", F2: 2}
                            _ = v11
                        }
                        fmt.Println(v1)
                        v3 = v3
                    }
                    v12 := GenS1{F0: 2, F1: "", F2: 2}
                    for fi13 := 0; fi13 < 6; fi13++ {
                        v14 := GenS1{F0: 758, F1: (("a" + "a") + ""), F2: 2}
                        v9 += v9
                        _ = fi13
                        _ = v14
                    }
                    _ = fi10
                    _ = v12
                }
                for fi15 := 0; fi15 < 20; fi15++ {
                    if false {
                        v9 += v9
                        if true {
                            switch "" {
                            case "\n\t":
                                v16 := GenS2{F0: 253}
                                _ = v16
                            default:
                                var v17 uint8 = 127
                                _ = v17
                            }
                            var v18 uint32 = (^(v7))
                            var v19 uint64 = 2
                            _ = v18
                            _ = v19
                        } else {
                            if (-1 < -742) {
                                v20 := GenS1{F0: 42, F1: "hello", F2: 0}
                                if (((!(true || ((((float32(1.5) - 1.5) + -2.25) / (0.0 + (float32((v1 * (v1 - 1e10))) / 1e10))) <= (float32((v1 * (-2.25 / (0.1 * 1e10)))) + (float32(v1) / 0.1))))) && false) && true) {
                                    for fi21 := 0; fi21 < 10; fi21++ {
                                        v7 = (^(v7))
                                        _ = fi21
                                    }
                                } else {
                                    v7 += (v7 & (v7 & v7))
                                }
                                _ = v20
                            }
                            for fi22 := 0; fi22 < 11; fi22++ {
                                var v23 float32 = 1e10
                                fmt.Println(v3, fi15, fi22)
                                _ = fi22
                                _ = v23
                            }
                        }
                    }
                    var v24 uint = (v9 + v9)
                    _ = fi15
                    _ = v24
                }
            } else {
                if len(v4) > 0 {
                    v4[0] = "hello"
                }
                var v25 int = 0
                _ = v25
            }
        } else {
            var v26 uint8 = 50
            for fi27 := 0; fi27 < 13; fi27++ {
                v7 = uint32(v9)
                var v28 uint16 = 42
                switch (fi27 | (^(fi27))) {
                case 128:
                    if ((!(!(42 >= 42))) || true) {
                        if (!(false && (!false))) {
                            for fi29 := 0; fi29 < 9; fi29++ {
                                var v30 int = fi27
                                v31 := GenS3{F0: (^(v26)), F1: 255}
                                _ = fi29
                                _ = v30
                                _ = v31
                            }
                            v1 -= 3.14159
                            fmt.Println(v7, v26, v28)
                        } else {
                            if (1 <= 127) {
                                for fi32 := 0; fi32 < 17; fi32++ {
                                    v33 := GenS2{F0: (v26 << ((18446744073709551615) & 63))}
                                    _ = fi32
                                    _ = v33
                                }
                            } else {
                                if true {
                                    var v34 int = fi27
                                    v35 := GenS3{F0: (v26 ^ uint8(v9)), F1: v26}
                                    var v36 uint = (v9 + v9)
                                    _ = v34
                                    _ = v35
                                    _ = v36
                                } else {
                                    var v37 bool = (v28 <= v28)
                                    var v38 float64 = v1
                                    v28 += v28
                                    _ = v37
                                    _ = v38
                                }
                            }
                        }
                    }
                case 32767:
                    v26 -= (v26 * v26)
                case -1:
                    v3 += (v3 + v3)
                default:
                    var v39 int16 = (^(v3))
                    _ = v39
                }
                _ = fi27
                _ = v28
            }
            _ = v26
        }
        fmt.Println(v1, v3, v1)
    }
    v4 = append(v4, "foo bar")
    switch ((("hello" + "hello") + "foo bar") + "\n\t") {
    case "a":
        v40 := v2[(v7 + v7)]
        _ = v40
    case "\n\t":
        var v41 int32 = -733
        _ = v41
    }
    if false {
        if true {
            var v42 int32 = -380
            var v43 int32 = (v42 % ((v42) | 1))
            _ = v42
            _ = v43
        }
        var v44 float64 = float64(-2.25)
        _ = v44
    } else {
        v45 := GenS3{F0: 83, F1: 128}
        _ = v45
    }
    v46 := GenS3{F0: 127, F1: 159}
    v3 += (v3 ^ 42)
    v3 = (v3 & -32768)
    if (true && (1 > -1)) {
        fmt.Println(v9)
    }
    for fi47 := 0; fi47 < 6; fi47++ {
        fmt.Println(v1, v3, v3)
        if (!true) {
            v48 := v2[(v7 & v7)]
            var v49 bool = ((0 != 18446744073709551615) && (v3 < -32768))
            _ = v48
            _ = v49
        } else {
            var v50 string = (("hello" + "") + "a")
            switch (fi47 - 9223372036854775807) {
            case 589:
                delete(v2, v7)
            case 0:
                var v51 float32 = (0.0 / float32(((0.0 - v1) / 3.14159)))
                _ = v51
            }
            _ = v50
        }
        v3 = int16(v7)
        _ = fi47
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
