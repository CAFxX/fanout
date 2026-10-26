package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint8 }
type GenS2 struct { F0 int8; F1 int64 }
func genFn0() int64 {
    if true {
        var v7 uint8 = 127
        var v8 float32 = (3.14159 + 1.5)
        _ = v7
        _ = v8
    } else {
        v9 := []int{42, 2147483647, 32767, 2147483647}
        _ = v9
    }
    var v10 string = "hello"
    for fi11 := 0; fi11 < 3; fi11++ {
        var v12 bool = ((v10 + (("\n\t" + ("\n\t" + v10)) + v10)) < v10)
        _ = fi11
        _ = v12
    }
    if false {
        var v13 bool = false
        _ = v13
    } else {
        var v14 uint8 = 127
        var v15 int64 = 0
        _ = v14
        _ = v15
    }
    _ = v10
    return 306
}
func genFn1(p0 uint32) int64 {
    if ((2 > 0) && true) {
        if false {
            switch ("a" + "") {
            case "\n\t":
                p0 = p0
            }
            v16 := GenS1{F0: 0, F1: 19}
            p0 += (p0 - p0)
            _ = v16
        } else {
            if true {
                var v17 float32 = (-2.25 + -2.25)
                v18 := GenS1{F0: -1, F1: 129}
                var v19 uint8 = 127
                _ = v17
                _ = v18
                _ = v19
            } else {
                var v20 int16 = -1
                var v21 uint32 = (p0 & p0)
                _ = v20
                _ = v21
            }
        }
        for fi22 := 0; fi22 < 5; fi22++ {
            for fi23 := 0; fi23 < 2; fi23++ {
                v24 := []int{4294967295, fi23}
                p0 = p0
                var v25 int16 = -1
                switch (fi23 | (fi22 + fi23)) {
                case -1:
                    v25 -= (v25 / ((v25) | 1))
                case 215:
                    if false {
                        switch "\n\t" {
                        case "hello":
                            var v26 uint32 = 956
                            _ = v26
                        }
                    } else {
                        v25 = (v25 ^ 42)
                        var v27 float32 = float32(0.1)
                        var v28 uint = 0
                        _ = v27
                        _ = v28
                    }
                    if len(v24) >= 2 {
                        v24 = v24[1:len(v24)]
                    }
                default:
                    if (genFn0() > -9223372036854775808) {
                        if false {
                            var v29 int32 = 293
                            v30 := GenS1{F0: 1, F1: 0}
                            v31 := GenS1{F0: (v29 & (v29 - 0)), F1: 82}
                            _ = v29
                            _ = v30
                            _ = v31
                        } else {
                            var v32 uint16 = 0
                            _ = v32
                        }
                        var v33 float32 = ((1.5 * 1e10) / 1e10)
                        _ = v33
                    }
                }
                _ = fi23
                _ = v24
                _ = v25
            }
            for fi34 := 0; fi34 < 11; fi34++ {
                var v35 uint64 = 186
                _ = fi34
                _ = v35
            }
            _ = fi22
        }
    } else {
        if true {
            fmt.Println(p0)
            p0 += p0
        }
        var v36 uint = 351
        _ = v36
    }
    p0 = p0
    _ = p0
    return 9223372036854775807
}

func main() {
    v1 := []uint32{}
    v2 := make(map[int]uint32)
    v3 := make(map[string]string)
    v4 := []string{(("foo bar" + "hello") + "hello"), ("hello" + "a")}
    v5 := make(map[int]string)
    var v6 bool = (42 < 0)
    switch "hello" {
    case "\n\t":
        if (!v6) {
            for fi37 := 0; fi37 < 17; fi37++ {
                var v38 string = "foo bar"
                for fi39 := 0; fi39 < 15; fi39++ {
                    if (488 == 2) {
                        var v40 bool = (65535 >= 12)
                        _ = v40
                    }
                    v6 = (55 == 127)
                    _ = fi39
                }
                var v41 string = ("hello" + v38)
                _ = fi37
                _ = v38
                _ = v41
            }
            for fi42 := 0; fi42 < 20; fi42++ {
                for fi43 := 0; fi43 < 12; fi43++ {
                    if v6 {
                        var v44 uint16 = 42
                        _ = v44
                    }
                    var v45 int32 = 2147483647
                    v6 = (-1 >= 1000)
                    _ = fi43
                    _ = v45
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
                var v46 uint8 = 8
                _ = fi42
                _ = v46
            }
        }
    case "foo bar":
        v47 := v3[("" + ("" + ("a" + "a")))]
        _ = v47
    default:
        v48 := GenS1{F0: -3, F1: 195}
        _ = v48
    }
    var v49 uint64 = 2
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v50 int64 = genFn1(2147483648)
    v50 += genFn0()
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v51 := GenS2{F0: -83, F1: genFn1(42)}
    if (-1 == -9223372036854775808) {
        for fi52 := 0; fi52 < 4; fi52++ {
            if (42 < -549) {
                v53 := GenS2{F0: 1, F1: (v50 + (v50 ^ v50))}
                for fi54 := 0; fi54 < 6; fi54++ {
                    _ = len(v5)
                    _ = fi54
                }
                var v55 int = 295
                _ = v53
                _ = v55
            } else {
                v6 = (v50 != v50)
            }
            _ = fi52
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v49:", v49)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v50
    _ = v51
    _ = v6
}
