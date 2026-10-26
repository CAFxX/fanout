package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint32; F1 int16 }
type GenS4 struct { F0 uint16 }
type GenS5 struct { F0 string }
func genFn0() uint16 {
    var v5 uint64 = 42
    if false {
        switch "hello" {
        case "a":
            var v6 uint16 = 2
            _ = v6
        case "foo bar":
            for fi7 := 0; fi7 < 1; fi7++ {
                v8 := GenS1{F0: -32768, F1: 2}
                switch ("\n\t" + "hello") {
                case "hello":
                    fmt.Println(v5, fi7)
                case "foo bar":
                    v9 := []int{fi7, 32767, fi7}
                    _ = v9
                }
                _ = fi7
                _ = v8
            }
        default:
            var v10 int16 = -1
            _ = v10
        }
    }
    v5 = (v5 << ((255) & 63))
    _ = v5
    return 506
}

func main() {
    var v1 int64 = 1
    v2 := make(map[string]bool)
    v3 := []int{-492, 954}
    var v4 uint32 = 2147483647
    var v11 int64 = -978
    delete(v2, "")
    v11 = v1
    var v12 string = "a"
    v13 := GenS1{F0: -32768, F1: 32767}
    for fi14 := 0; fi14 < 12; fi14++ {
        v15 := GenS2{F0: -1}
        _ = fi14
        _ = v15
    }
    v2["a"] = true
    for fi16 := 0; fi16 < 5; fi16++ {
        v12 = "\n\t"
        if true {
            var v17 int = fi16
            _ = v17
        }
        var v18 uint64 = 4294967295
        _ = fi16
        _ = v18
    }
    var v19 int32 = 2147483647
    var v20 int = 2147483647
    if (!false) {
        fmt.Println(v19)
    } else {
        v3 = append(v3, v20)
        switch 124 {
        case 220:
            v21 := GenS3{F0: (v4 >> ((65535) & 63)), F1: 36}
            _ = v21
        case 1:
            switch (v20 * 4294967295) {
            case 32768:
                for fi22 := 0; fi22 < 7; fi22++ {
                    var v23 int16 = 1000
                    for fi24 := 0; fi24 < 19; fi24++ {
                        v2[v12] = false
                        _ = fi24
                    }
                    _ = len(v2)
                    _ = fi22
                    _ = v23
                }
            case 2147483647:
                if (1e10 == 1.5) {
                    switch "a" {
                    case "a":
                        fmt.Println(v20)
                    case "":
                        if true {
                            if (!(0.1 >= (1e10 * (float32(0.1) + (0.1 * 0.1))))) {
                                v25 := v2[v12]
                                for fi26 := 0; fi26 < 17; fi26++ {
                                    v20 -= (v20 ^ (v20 >> ((0) & 63)))
                                    _ = fi26
                                }
                                _ = v25
                            } else {
                                _ = len(v3)
                            }
                            var v27 string = v12
                            switch v12 {
                            case "a":
                                var v28 uint = 1
                                _ = v28
                            default:
                                for fi29 := 0; fi29 < 4; fi29++ {
                                    v30 := GenS1{F0: 32767, F1: 729}
                                    switch (fi29 * (fi29 + (fi29 - (v20 | v20)))) {
                                    case 4294967295:
                                        if true {
                                            fmt.Println(v27, v19)
                                        }
                                    case -9223372036854775808:
                                        switch 798 {
                                        case 0:
                                            delete(v2, v27)
                                        case 32767:
                                            v31 := GenS1{F0: -32768, F1: 1}
                                            _ = v31
                                        default:
                                            var v32 uint8 = 190
                                            _ = v32
                                        }
                                    }
                                    fmt.Println(v12, v1, v19)
                                    _ = fi29
                                    _ = v30
                                }
                            }
                            _ = v27
                        } else {
                            if true {
                                var v33 float64 = float64(1.5)
                                _ = v33
                            } else {
                                for fi34 := 0; fi34 < 16; fi34++ {
                                    v4 = 852
                                    v35 := GenS1{F0: 32767, F1: 65535}
                                    _ = fi34
                                    _ = v35
                                }
                                switch v20 {
                                case 610:
                                    var v36 int8 = -1
                                    _ = v36
                                case 127:
                                    for fi37 := 0; fi37 < 17; fi37++ {
                                        if false {
                                            var v38 bool = false
                                            v39 := GenS4{F0: 65535}
                                            switch "hello" {
                                            case "hello":
                                                var v40 uint16 = 369
                                                _ = v40
                                            }
                                            _ = v38
                                            _ = v39
                                        } else {
                                            v19 = (v19 | v19)
                                            var v41 string = (v12 + v12)
                                            v12 = (v41 + "\n\t")
                                            _ = v41
                                        }
                                        _ = fi37
                                    }
                                case 826:
                                    for fi42 := 0; fi42 < 6; fi42++ {
                                        if (!false) {
                                            v19 = (v19 << ((814) & 63))
                                            v2[("\n\t" + "hello")] = ((!false) || ((255 == 2) && true))
                                            if len(v3) > 0 {
                                                v3[0] = -1
                                            }
                                        } else {
                                            var v43 bool = (1.5 < 0.0)
                                            v44 := v2[v12]
                                            _ = v43
                                            _ = v44
                                        }
                                        _ = len(v2)
                                        _ = fi42
                                    }
                                }
                                v11 = v11
                            }
                        }
                    }
                    v45 := GenS5{F0: (v12 + (v12 + "foo bar"))}
                    for fi46 := 0; fi46 < 7; fi46++ {
                        var v47 int32 = (v19 * v19)
                        if (!((42 > 2) && (867 <= 255))) {
                            v48 := GenS4{F0: 1}
                            _ = v48
                        } else {
                            if ((-128 <= 100) && true) {
                                switch v12 {
                                case "hello":
                                    switch v12 {
                                    case "":
                                        var v49 string = "a"
                                        _ = v49
                                    }
                                case "foo bar":
                                    _ = len(v2)
                                }
                            }
                            v19 += (v47 << ((0) & 63))
                        }
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        _ = fi46
                        _ = v47
                    }
                    _ = v45
                }
            }
        case -16:
            v50 := GenS2{F0: v11}
            _ = v50
        }
        if len(v3) > 0 {
            v3[0] = v20
        }
    }
    v1 += (^(v11))
    v51 := GenS4{F0: 201}
    fmt.Println(v20, v12, v4)
    var v52 int64 = v11
    v53 := v2[v12]
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v51
    _ = v52
    _ = v53
}
