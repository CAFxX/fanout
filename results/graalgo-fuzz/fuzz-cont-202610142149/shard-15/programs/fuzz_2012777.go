package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 uint16 }
type GenS3 struct { F0 int; F1 uint16 }

func main() {
    var v1 int = -617
    v2 := make(map[string]int)
    var v3 int32 = -286
    var v4 uint16 = 989
    var v5 bool = false
    var v6 uint64 = 546
    v7 := GenS1{F0: (-(v3))}
    v3 += (v3 >> ((2) & 63))
    v6 -= 1
    var v8 int32 = 873
    var v9 bool = v5
    v10 := GenS1{F0: v8}
    v11 := GenS1{F0: 2147483647}
    v12 := GenS2{F0: (v4 * (v4 >> ((18446744073709551615) & 63))), F1: 42}
    v13 := GenS3{F0: (v1 / ((v1) | 1)), F1: 42}
    for fi14 := 0; fi14 < 6; fi14++ {
        v5 = (!v5)
        _ = fi14
    }
    v1 -= (^(v1))
    for fi15 := 0; fi15 < 5; fi15++ {
        v3 = (v3 % ((v8) | 1))
        v16 := GenS1{F0: -675}
        switch fi15 {
        case 42:
            fmt.Println(v6)
        case -388:
            v17 := []int{fi15, (-(v1)), (v1 << ((2) & 63)), (fi15 / ((v1) | 1))}
            _ = v17
        case 32768:
            if false {
                var v18 uint64 = (v6 + v6)
                if false {
                    for fi19 := 0; fi19 < 15; fi19++ {
                        if v9 {
                            v2[("a" + "\n\t")] = (fi19 / ((v1) | 1))
                            v20 := GenS2{F0: v4, F1: v4}
                            fmt.Println(v5)
                            _ = v20
                        } else {
                            var v21 string = "hello"
                            v22 := GenS3{F0: fi19, F1: 458}
                            _ = v21
                            _ = v22
                        }
                        var v23 int = int(v4)
                        _ = fi19
                        _ = v23
                    }
                    if false {
                        if true {
                            fmt.Println(v3)
                            v24 := GenS1{F0: (v3 | v8)}
                            var v25 int16 = 32767
                            _ = v24
                            _ = v25
                        } else {
                            if (uint32(v18) >= 632) {
                                v8 -= -1
                                v4 = v4
                                var v26 int = (v1 ^ 9223372036854775807)
                                _ = v26
                            } else {
                                for fi27 := 0; fi27 < 2; fi27++ {
                                    var v28 bool = ((!v9) || (5 != 147))
                                    var v29 uint64 = (v18 << ((4294967295) & 63))
                                    _ = fi27
                                    _ = v28
                                    _ = v29
                                }
                                fmt.Println(v8, v1, v1)
                            }
                        }
                        if (1.5 == (3.14159 * (1e10 + 0.0))) {
                            v30 := GenS1{F0: v8}
                            _ = v30
                        }
                        if ((v3 / (((-(v8))) | 1)) <= 1) {
                            _ = len(v2)
                            for fi31 := 0; fi31 < 17; fi31++ {
                                v32 := GenS1{F0: (-(v8))}
                                v8 = v3
                                v33 := []int{-747, (fi15 * (v1 | fi15)), 1}
                                _ = fi31
                                _ = v32
                                _ = v33
                            }
                        }
                    }
                } else {
                    v18 -= v6
                }
                _ = v18
            } else {
                var v34 bool = v5
                _ = v34
            }
        default:
            var v35 int32 = v8
            _ = v35
        }
        _ = fi15
        _ = v16
    }
    fmt.Println(v6)
    var v36 uint16 = 2
    var v37 int16 = -368
    v38 := GenS3{F0: int(v37), F1: (v36 / (((v4 + 1)) | 1))}
    switch "hello" {
    case "\n\t":
        switch "" {
        case "a":
            switch "foo bar" {
            case "\n\t":
                var v39 uint64 = (v6 % ((4294967295) | 1))
                _ = v39
            case "a":
                for fi40 := 0; fi40 < 8; fi40++ {
                    var v41 uint8 = 42
                    v2[((("foo bar" + "") + ("" + "foo bar")) + ("" + "foo bar"))] = fi40
                    _ = fi40
                    _ = v41
                }
            case "hello":
                v3 = 2147483647
            }
        case "foo bar":
            for fi42 := 0; fi42 < 20; fi42++ {
                for fi43 := 0; fi43 < 2; fi43++ {
                    if ((v4 * (v4 << ((18446744073709551615) & 63))) > v4) {
                        v36 -= 393
                        v2["foo bar"] = 4294967295
                        v9 = (v6 > (v6 - (^(v6))))
                    }
                    switch fi42 {
                    case 32768:
                        if v5 {
                            v9 = false
                            v37 += v37
                        }
                    }
                    _ = fi43
                }
                v44 := []int{32767, 522}
                fmt.Println(v8)
                _ = fi42
                _ = v44
            }
        }
    }
    delete(v2, "a")
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
