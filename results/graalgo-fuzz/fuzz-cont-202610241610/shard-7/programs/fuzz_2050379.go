package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }

func main() {
    v1 := make(map[int]bool)
    var v2 uint8 = 42
    v3 := make(map[int]uint32)
    v4 := make(map[string]uint32)
    v2 -= 255
    v2 -= v2
    v2 = (v2 ^ 128)
    fmt.Println(v2)
    if (42 == 0) {
        v1[32767] = false
        switch 32767 {
        case 0:
            if (1 > 599) {
                v2 = 42
                for fi5 := 0; fi5 < 17; fi5++ {
                    switch ("foo bar" + "foo bar") {
                    case "a":
                        v6 := v3[9223372036854775807]
                        _ = v6
                    case "foo bar":
                        var v7 string = "\n\t"
                        _ = v7
                    }
                    _ = fi5
                }
                var v8 uint16 = 255
                _ = v8
            }
        default:
            if (1.5 != 0.1) {
                v2 = 24
                var v9 uint = 236
                var v10 bool = true
                _ = v10
                _ = v9
            }
        }
    } else {
        switch "hello" {
        case "a":
            var v11 uint8 = v2
            _ = v11
        case "hello":
            if true {
                v12 := []int{-632, 1}
                switch 619 {
                case -2147483648:
                    switch -925 {
                    case -46:
                        var v13 uint64 = 564
                        _ = v13
                    case 4294967295:
                        switch ("\n\t" + "hello") {
                        case "":
                            if ((!true) || true) {
                                if (!false) {
                                    switch 9223372036854775807 {
                                    case 32767:
                                        v14 := GenS1{F0: 31}
                                        _ = v14
                                    case -947:
                                        switch 9223372036854775807 {
                                        case 800:
                                            for fi15 := 0; fi15 < 6; fi15++ {
                                                v3[(fi15 | (^(fi15)))] = 2147483647
                                                _ = fi15
                                            }
                                        case -809:
                                            var v16 int64 = 0
                                            _ = v16
                                        case 315:
                                            v17 := GenS1{F0: 71}
                                            _ = v17
                                        default:
                                            if len(v12) >= 2 {
                                                v12 = v12[1:len(v12)]
                                            }
                                        }
                                    }
                                } else {
                                    switch -1 {
                                    case -923:
                                        var v18 uint64 = 512
                                        _ = v18
                                    }
                                }
                            } else {
                                for fi19 := 0; fi19 < 9; fi19++ {
                                    var v20 uint8 = (^(v2))
                                    switch (fi19 >> ((255) & 63)) {
                                    case 128:
                                        v2 = (v2 & v20)
                                    case -690:
                                        if len(v12) > 0 {
                                            v12[0] = fi19
                                        }
                                    }
                                    _ = fi19
                                    _ = v20
                                }
                                var v21 uint32 = 0
                                switch "a" {
                                case "a":
                                    v22 := GenS1{F0: 127}
                                    _ = v22
                                case "foo bar":
                                    switch "" {
                                    case "foo bar":
                                        v23 := v3[32767]
                                        _ = v23
                                    default:
                                        delete(v4, ("foo bar" + ("\n\t" + "hello")))
                                    }
                                }
                                _ = v21
                            }
                        default:
                            var v24 uint8 = (v2 + 1)
                            _ = v24
                        }
                    case -2147483648:
                        switch ("hello" + "a") {
                        case "hello":
                            if len(v12) >= 2 {
                                v12 = v12[1:len(v12)]
                            }
                        case "a":
                            v2 += v2
                        default:
                            if (false && (-2147483648 != -2147483648)) {
                                v2 = (v2 / ((42) | 1))
                            }
                        }
                    }
                default:
                    v2 = v2
                }
                var v25 float64 = ((-2.25 * float64((1.5 * 0.1))) - 0.1)
                _ = v12
                _ = v25
            } else {
                v2 = 128
            }
        case "\n\t":
            v26 := GenS1{F0: 42}
            _ = v26
        default:
            fmt.Println(v2)
        }
        v2 = (^(v2))
        if false {
            var v27 string = "a"
            if false {
                v28 := []int{1, 32768, 32768}
                _ = v28
            } else {
                for fi29 := 0; fi29 < 11; fi29++ {
                    v2 -= 42
                    v27 = ("hello" + (v27 + ("a" + (("hello" + v27) + "foo bar"))))
                    _ = fi29
                }
                v30 := v4[v27]
                _ = v30
            }
            _ = v27
        }
    }
    v31 := GenS1{F0: 29}
    fmt.Println(v2)
    var v32 bool = false
    var v33 string = "\n\t"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
}
