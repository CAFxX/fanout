package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 bool }
type GenS2 struct { F0 bool; F1 bool; F2 int32 }

func main() {
    v1 := make(map[uint32]bool)
    var v2 string = ""
    v3 := make(map[int]bool)
    var v4 bool = (!true)
    var v5 int = -2147483648
    for fi6 := 0; fi6 < 6; fi6++ {
        for fi7 := 0; fi7 < 2; fi7++ {
            v2 = ""
            fmt.Println(fi6, v4)
            _ = fi7
        }
        v8 := v1[2147483648]
        _ = fi6
        _ = v8
    }
    if v4 {
        if v4 {
            var v9 int32 = 42
            v2 = (v2 + (v2 + (("" + v2) + v2)))
            _ = v9
        } else {
            var v10 int32 = 0
            var v11 uint32 = 879
            fmt.Println(v11)
            _ = v10
            _ = v11
        }
    } else {
        v3[-658] = v4
    }
    v4 = (255 == 744)
    v5 += v5
    var v12 int16 = 0
    v13 := []int{}
    v14 := GenS1{F0: -63, F1: "foo bar", F2: (!true)}
    fmt.Println(v12, v5, v2)
    for fi15 := 0; fi15 < 6; fi15++ {
        var v16 uint64 = 0
        v12 -= v12
        _ = fi15
        _ = v16
    }
    for fi17 := 0; fi17 < 19; fi17++ {
        v18 := GenS1{F0: -128, F1: v2, F2: v4}
        v19 := GenS1{F0: -128, F1: v2, F2: v4}
        _ = fi17
        _ = v18
        _ = v19
    }
    var v20 int8 = 44
    var v21 int = v5
    if (false && false) {
        for fi22 := 0; fi22 < 18; fi22++ {
            v23 := GenS2{F0: v4, F1: v4, F2: -2147483648}
            _ = fi22
            _ = v23
        }
        if true {
            if (4294967295 >= 1) {
                if v4 {
                    if len(v13) >= 2 {
                        v13 = v13[1:len(v13)]
                    }
                    var v24 uint64 = 42
                    _ = v24
                } else {
                    for fi25 := 0; fi25 < 4; fi25++ {
                        if v4 {
                            _ = len(v13)
                            var v26 uint16 = 42
                            _ = v26
                        } else {
                            for fi27 := 0; fi27 < 17; fi27++ {
                                var v28 uint = 42
                                _ = fi27
                                _ = v28
                            }
                        }
                        if v4 {
                            switch "a" {
                            case "\n\t":
                                switch (v21 - fi25) {
                                case -84:
                                    v4 = (2147483648 > 106)
                                }
                            case "foo bar":
                                v29 := GenS1{F0: -1, F1: (v2 + "foo bar"), F2: (v20 < 0)}
                                _ = v29
                            case "":
                                v2 = v2
                            default:
                                v30 := v3[v5]
                                _ = v30
                            }
                            var v31 bool = v4
                            if false {
                                var v32 int32 = 42
                                _ = v32
                            }
                            _ = v31
                        } else {
                            if ((v12 | 1) >= v12) {
                                v12 += (v12 ^ 1000)
                            } else {
                                var v33 uint64 = 659
                                _ = v33
                            }
                        }
                        _ = fi25
                    }
                }
                for fi34 := 0; fi34 < 19; fi34++ {
                    v35 := GenS1{F0: v20, F1: "\n\t", F2: v4}
                    var v36 string = ("" + "")
                    if (-2.25 <= 1.5) {
                        v4 = (!v4)
                    } else {
                        v2 = (v2 + ("foo bar" + ((v2 + ((v2 + (v36 + v36)) + "foo bar")) + v2)))
                        var v37 int16 = (v12 * 1000)
                        _ = v37
                    }
                    _ = fi34
                    _ = v35
                    _ = v36
                }
            } else {
                switch (-(v21)) {
                case -2147483648:
                    fmt.Println(v2, v2)
                default:
                    v38 := GenS2{F0: true, F1: false, F2: 0}
                    _ = v38
                }
                for fi39 := 0; fi39 < 15; fi39++ {
                    _ = len(v1)
                    _ = fi39
                }
            }
            var v40 int32 = 2147483647
            _ = v40
        } else {
            v21 = v5
            if (!true) {
                fmt.Println(v2, v4, v2)
            }
            v12 -= v12
        }
        var v41 int8 = v20
        _ = v41
    }
    v20 += (-(v20))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
}
