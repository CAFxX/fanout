package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }

func main() {
    var v1 string = ""
    v2 := make(map[uint32]int)
    var v3 uint16 = 42
    v4 := []int{752, 32767, 937}
    v1 = v1
    v1 = "hello"
    if ((65535 < 255) && false) {
        fmt.Println(v1)
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        v5 := v2[65535]
        _ = v5
    }
    fmt.Println(v1, v1)
    if false {
        fmt.Println(v3)
        v1 = v1
    } else {
        if false {
            _ = len(v4)
            switch 437 {
            case 32767:
                v1 = v1
            default:
                if (0 < 1) {
                    var v6 uint64 = 176
                    if ((v3 + v3) == v3) {
                        switch v1 {
                        case "a":
                            v7 := GenS1{F0: 42, F1: -1}
                            _ = v7
                        case "":
                            for fi8 := 0; fi8 < 10; fi8++ {
                                switch (("\n\t" + ("hello" + v1)) + "\n\t") {
                                case "hello":
                                    if len(v4) > 0 {
                                        v4[0] = -2147483648
                                    }
                                }
                                var v9 bool = false
                                if (-9223372036854775808 > 9223372036854775807) {
                                    v9 = (v9 || (v6 != v6))
                                    if v9 {
                                        for fi10 := 0; fi10 < 15; fi10++ {
                                            v3 = 1
                                            fmt.Println(fi8)
                                            var v11 int = fi8
                                            _ = fi10
                                            _ = v11
                                        }
                                    }
                                    for fi12 := 0; fi12 < 6; fi12++ {
                                        if v9 {
                                            switch (v1 + ((v1 + "a") + "foo bar")) {
                                            case "foo bar":
                                                v1 = v1
                                            case "hello":
                                                fmt.Println(v6, fi8)
                                            case "a":
                                                var v13 int32 = 1
                                                _ = v13
                                            }
                                            fmt.Println(v1)
                                            _ = len(v4)
                                        } else {
                                            _ = len(v2)
                                        }
                                        _ = fi12
                                    }
                                } else {
                                    v14 := GenS1{F0: 0, F1: 801}
                                    _ = v14
                                }
                                _ = fi8
                                _ = v9
                            }
                        default:
                            var v15 uint32 = 42
                            _ = v15
                        }
                        v6 = v6
                    }
                    _ = len(v2)
                    _ = v6
                } else {
                    _ = len(v4)
                }
            }
        } else {
            switch -2147483648 {
            case -710:
                for fi16 := 0; fi16 < 1; fi16++ {
                    var v17 uint32 = 2147483648
                    _ = fi16
                    _ = v17
                }
            case -9223372036854775808:
                fmt.Println(v1, v3)
            case 2147483647:
                v2[2] = 9223372036854775807
            }
        }
    }
    v3 += v3
    v1 = "a"
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
