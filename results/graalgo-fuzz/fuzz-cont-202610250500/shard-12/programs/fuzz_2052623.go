package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 uint64; F1 uint64 }
func genFn0(p0 int16, p1 uint16) bool {
    for fi4 := 0; fi4 < 19; fi4++ {
        for fi5 := 0; fi5 < 7; fi5++ {
            var v6 int16 = p0
            _ = fi5
            _ = v6
        }
        if (!true) {
            v7 := GenS1{F0: 0, F1: 4294967295}
            switch fi4 {
            case 127:
                p1 = 2
            }
            if (0.1 < -2.25) {
                for fi8 := 0; fi8 < 5; fi8++ {
                    p1 = (p1 << ((385) & 63))
                    if false {
                        fmt.Println(p0, p0)
                    } else {
                        v9 := GenS1{F0: 1, F1: 1}
                        _ = v9
                    }
                    switch fi8 {
                    case 2147483647:
                        fmt.Println(p0)
                    case -1:
                        fmt.Println(p0)
                    default:
                        switch "foo bar" {
                        case "foo bar":
                            v10 := GenS1{F0: -9223372036854775808, F1: 156}
                            _ = v10
                        case "":
                            if true {
                                v11 := GenS1{F0: 9223372036854775807, F1: 1}
                                var v12 uint16 = (^(p1))
                                if ((-1 > 100) || (1 >= 0)) {
                                    fmt.Println(p1, v12)
                                    v13 := GenS2{F0: 0, F1: 585}
                                    _ = v13
                                }
                                _ = v11
                                _ = v12
                            } else {
                                p0 = (p0 % (((p0 << ((191) & 63))) | 1))
                                v14 := []int{(fi8 / (((fi4 | (fi4 << ((1) & 63)))) | 1)), fi4, fi8}
                                var v15 uint32 = 320
                                _ = v14
                                _ = v15
                            }
                        }
                    }
                    _ = fi8
                }
                p1 += p1
                for fi16 := 0; fi16 < 4; fi16++ {
                    v17 := GenS1{F0: int64(fi16), F1: 2}
                    _ = fi16
                    _ = v17
                }
            }
            _ = v7
        }
        _ = fi4
    }
    var v18 uint8 = 100
    _ = p0
    _ = p1
    _ = v18
    return (((("\n\t" + "\n\t") + "foo bar") + "\n\t") >= "hello")
}

func main() {
    v1 := []string{"a"}
    var v2 int16 = -1
    var v3 string = "foo bar"
    _ = len(v1)
    for fi19 := 0; fi19 < 3; fi19++ {
        for fi20 := 0; fi20 < 11; fi20++ {
            v3 = v3
            var v21 float32 = (0.1 / 1.0)
            fmt.Println(fi19)
            _ = fi20
            _ = v21
        }
        if (42 < -222) {
            v3 = "foo bar"
            if false {
                fmt.Println(v3)
            }
        }
        v3 = (v3 + (v3 + v3))
        _ = fi19
    }
    v1 = append(v1, "a")
    _ = len(v1)
    if (938 == 255) {
        fmt.Println(v3, v3)
        fmt.Println(v2, v3)
        v2 += -32768
    }
    if true {
        v2 = -1
        v22 := GenS1{F0: 9223372036854775807, F1: 4294967295}
        v23 := GenS2{F0: 42, F1: 2}
        _ = v22
        _ = v23
    } else {
        v3 = (("foo bar" + "a") + "foo bar")
        switch "a" {
        case "a":
            v24 := GenS2{F0: 4294967295, F1: 4294967295}
            _ = v24
        }
    }
    fmt.Println(v3, v3)
    switch "a" {
    case "a":
        for fi25 := 0; fi25 < 2; fi25++ {
            fmt.Println(v3)
            _ = fi25
        }
    }
    _ = len(v1)
    fmt.Println(v2)
    v26 := GenS1{F0: -1, F1: 1}
    var v27 uint16 = 633
    var v28 uint8 = 1
    fmt.Println(v27, v28)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
}
