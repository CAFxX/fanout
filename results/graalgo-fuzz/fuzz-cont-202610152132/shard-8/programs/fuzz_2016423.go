package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }

func main() {
    var v1 uint8 = 255
    v2 := make(map[string]bool)
    v3 := make(map[string]bool)
    v4 := make(map[int]int)
    var v5 uint64 = 2
    v6 := v3[("hello" + "a")]
    v6 = ((false && false) && (!v6))
    switch -1 {
    case 32767:
        v5 = 75
    case 439:
        v7 := v3[("\n\t" + "hello")]
        _ = v7
    case -2147483648:
        if ((100 >= 1) && (!v6)) {
            v8 := GenS1{F0: 0}
            var v9 float64 = 0.1
            _ = v8
            _ = v9
        }
    default:
        switch -514 {
        case 9223372036854775807:
            v1 = 105
        case 0:
            fmt.Println(v5)
        case 2147483647:
            switch ("hello" + "") {
            case "":
                var v10 uint8 = v1
                _ = v10
            case "hello":
                for fi11 := 0; fi11 < 3; fi11++ {
                    switch fi11 {
                    case 42:
                        var v12 uint16 = 29
                        _ = v12
                    }
                    _ = fi11
                }
            default:
                v5 = 275
            }
        }
    }
    v5 -= v5
    for fi13 := 0; fi13 < 14; fi13++ {
        var v14 int8 = -128
        v15 := []int{fi13, fi13, fi13}
        v3["hello"] = (2 == 0)
        _ = fi13
        _ = v14
        _ = v15
    }
    v16 := GenS1{F0: 685}
    v1 -= (^(v1))
    if v6 {
        switch "a" {
        case "hello":
            for fi17 := 0; fi17 < 17; fi17++ {
                for fi18 := 0; fi18 < 2; fi18++ {
                    fmt.Println(v5, fi17, v6)
                    _ = fi18
                }
                fmt.Println(v5, v6, v5)
                switch (fi17 / (((-(fi17))) | 1)) {
                case 126:
                    if (!(1 <= -1)) {
                        if v6 {
                            v19 := GenS1{F0: -9223372036854775808}
                            _ = v19
                        }
                    } else {
                        switch fi17 {
                        case -866:
                            switch (^(fi17)) {
                            case -2147483648:
                                for fi20 := 0; fi20 < 7; fi20++ {
                                    v6 = v6
                                    v6 = v6
                                    var v21 int64 = 0
                                    switch 9223372036854775807 {
                                    case -18:
                                        switch (-(fi17)) {
                                        case 138:
                                            v21 = (^(v21))
                                        case -462:
                                            if true {
                                                if (((("\n\t" + "\n\t") + "foo bar") + "foo bar") < "foo bar") {
                                                    var v22 uint16 = 255
                                                    _ = v22
                                                }
                                                fmt.Println(fi20, v21, v1)
                                                v21 = 0
                                            }
                                        }
                                    }
                                    v2["foo bar"] = ((fi20 - fi20) < -2147483648)
                                    _ = fi20
                                    _ = v21
                                }
                            case -302:
                                v5 -= 2
                            case 9223372036854775807:
                                v6 = (v6 && (fi17 == fi17))
                            default:
                                var v23 int64 = 42
                                _ = v23
                            }
                        case 330:
                            var v24 int32 = -2147483648
                            _ = v24
                        case 709:
                            var v25 uint = 287
                            _ = v25
                        }
                        for fi26 := 0; fi26 < 15; fi26++ {
                            var v27 int32 = 2147483647
                            _ = fi26
                            _ = v27
                        }
                    }
                default:
                    for fi28 := 0; fi28 < 6; fi28++ {
                        _ = len(v3)
                        v5 -= v5
                        _ = fi28
                    }
                }
                _ = fi17
            }
        default:
            v6 = v6
        }
    }
    v29 := []int{615}
    v30 := GenS1{F0: -66}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v29:", len(v29), v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
