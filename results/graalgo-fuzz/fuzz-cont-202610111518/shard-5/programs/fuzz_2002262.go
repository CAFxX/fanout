package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16 }
type GenS2 struct { F0 int32; F1 string; F2 uint64 }
type GenS3 struct { F0 int64; F1 uint64 }

func main() {
    v1 := []string{("\n\t" + ("foo bar" + "")), ("\n\t" + ""), (("" + "a") + "hello")}
    var v2 int16 = 812
    var v3 uint32 = 576
    var v4 uint64 = 0
    v5 := make(map[string]bool)
    var v6 int8 = 9
    v7 := GenS1{F0: 0, F1: (^(v2))}
    fmt.Println(v4, v2)
    v4 = v4
    for fi8 := 0; fi8 < 15; fi8++ {
        var v9 uint = 65535
        var v10 uint64 = 872
        _ = fi8
        _ = v10
        _ = v9
    }
    v3 = v3
    if false {
        if false {
            var v11 int32 = 396
            v6 -= v6
            var v12 int = 4294967295
            _ = v11
            _ = v12
        } else {
            var v13 bool = ("" <= "a")
            for fi14 := 0; fi14 < 10; fi14++ {
                if (0 < 65535) {
                    for fi15 := 0; fi15 < 1; fi15++ {
                        var v16 int8 = (v6 & v6)
                        v5["a"] = ((("foo bar" > (("foo bar" + "a") + "a")) || (v13 || (v13 || (589 >= 65535)))) || (v4 == (^(v4))))
                        _ = fi15
                        _ = v16
                    }
                } else {
                    switch "a" {
                    case "":
                        fmt.Println(v3, v6)
                    case "hello":
                        for fi17 := 0; fi17 < 3; fi17++ {
                            switch "" {
                            case "a":
                                _ = len(v5)
                            default:
                                v3 -= v3
                            }
                            v3 += v3
                            _ = fi17
                        }
                    default:
                        if v13 {
                            var v18 float32 = 1e10
                            _ = v18
                        } else {
                            v2 = (v2 ^ v2)
                            if v13 {
                                v4 -= (v4 % (((v4 ^ (v4 ^ 42))) | 1))
                            }
                            v6 -= v6
                        }
                    }
                }
                for fi19 := 0; fi19 < 20; fi19++ {
                    var v20 int16 = -900
                    v20 -= v20
                    _ = fi19
                    _ = v20
                }
                _ = fi14
            }
            var v21 uint16 = 625
            _ = v13
            _ = v21
        }
    }
    var v22 int8 = v6
    switch "" {
    case "hello":
        if len(v1) > 0 {
            v1[0] = "foo bar"
        }
    case "":
        switch 32767 {
        case -2147483648:
            v6 += 48
        case 1:
            if true {
                for fi23 := 0; fi23 < 13; fi23++ {
                    var v24 uint32 = (v3 ^ (v3 & uint32(v4)))
                    var v25 uint16 = 378
                    if (-2.25 >= 3.14159) {
                        fmt.Println(v4, v3, v3)
                        fmt.Println(v25)
                        v3 = v3
                    }
                    _ = fi23
                    _ = v24
                    _ = v25
                }
            } else {
                fmt.Println(v2)
                switch 0 {
                case 2147483647:
                    for fi26 := 0; fi26 < 13; fi26++ {
                        v27 := GenS1{F0: -352, F1: v2}
                        for fi28 := 0; fi28 < 6; fi28++ {
                            if true {
                                v29 := GenS2{F0: 1, F1: ("\n\t" + "\n\t"), F2: v4}
                                v30 := GenS2{F0: 2147483647, F1: "\n\t", F2: v4}
                                _ = v29
                                _ = v30
                            }
                            _ = fi28
                        }
                        for fi31 := 0; fi31 < 13; fi31++ {
                            var v32 uint64 = 42
                            _ = fi31
                            _ = v32
                        }
                        _ = fi26
                        _ = v27
                    }
                default:
                    if (!true) {
                        var v33 string = "\n\t"
                        for fi34 := 0; fi34 < 5; fi34++ {
                            if len(v1) > 0 {
                                v1[0] = v33
                            }
                            _ = fi34
                        }
                        _ = v33
                    }
                }
                var v35 int = 42
                _ = v35
            }
        case 32767:
            for fi36 := 0; fi36 < 5; fi36++ {
                var v37 bool = (!false)
                v37 = true
                _ = fi36
                _ = v37
            }
        }
    case "a":
        v38 := GenS2{F0: -310, F1: "foo bar", F2: (^(v4))}
        _ = v38
    default:
        if true {
            v5["a"] = ((!true) && (-625 > (v2 - (v2 / ((0) | 1)))))
        }
    }
    var v39 int32 = 42
    v2 -= -187
    if ((v2 ^ (v2 % ((v2) | 1))) >= v2) {
        v40 := GenS3{F0: -1, F1: (v4 % ((v4) | 1))}
        fmt.Println(v39, v39)
        _ = v40
    }
    delete(v5, ("" + "foo bar"))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
