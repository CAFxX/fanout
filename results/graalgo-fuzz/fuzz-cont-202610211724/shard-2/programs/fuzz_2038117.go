package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8; F2 uint16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint64 }
func genFn0() string {
    var v5 string = ("hello" + "foo bar")
    var v6 int8 = 51
    if ((2147483647 == 42) || (false && (18446744073709551615 == 0))) {
        v7 := GenS1{F0: 956, F1: (-(v6)), F2: 255}
        var v8 bool = (true && false)
        _ = v7
        _ = v8
    } else {
        for fi9 := 0; fi9 < 10; fi9++ {
            for fi10 := 0; fi10 < 19; fi10++ {
                if (255 != 2) {
                    v11 := GenS2{F0: (v5 + (v5 + v5))}
                    if false {
                        var v12 int64 = -240
                        _ = v12
                    } else {
                        v6 -= v6
                    }
                    _ = v11
                } else {
                    var v13 float64 = 1.5
                    _ = v13
                }
                _ = fi10
            }
            v14 := GenS3{F0: 320}
            var v15 uint64 = 4294967295
            _ = fi9
            _ = v14
            _ = v15
        }
        var v16 int16 = 447
        var v17 uint32 = 990
        _ = v16
        _ = v17
    }
    _ = v5
    _ = v6
    return v5
}
func genFn1(p0 uint8, p1 int32) bool {
    if true {
        if (!true) {
            if true {
                for fi18 := 0; fi18 < 7; fi18++ {
                    p1 -= (p1 | p1)
                    v19 := []int{(-(fi18)), fi18, (fi18 ^ (fi18 & fi18)), (^(fi18))}
                    if true {
                        v19 = append(v19, 0)
                        var v20 uint64 = 1
                        _ = v20
                    }
                    _ = fi18
                    _ = v19
                }
                if true {
                    p1 = (p1 ^ p1)
                    if (!true) {
                        fmt.Println(p0)
                        p0 = (p0 | p0)
                    }
                }
                if (-1 >= 42) {
                    if (-2147483648 >= -172) {
                        if (4294967295 != 0) {
                            p0 -= p0
                            p1 -= (-(p1))
                            switch ((("foo bar" + "\n\t") + ("" + "\n\t")) + ((((("\n\t" + "hello") + ("foo bar" + "\n\t")) + "foo bar") + "a") + "foo bar")) {
                            case "foo bar":
                                for fi21 := 0; fi21 < 2; fi21++ {
                                    fmt.Println(fi21, p0, p0)
                                    for fi22 := 0; fi22 < 17; fi22++ {
                                        p0 -= p0
                                        _ = fi22
                                    }
                                    fmt.Println(p0, p1, p1)
                                    _ = fi21
                                }
                            }
                        }
                    }
                    p1 = p1
                    v23 := GenS2{F0: ""}
                    _ = v23
                }
            } else {
                v24 := []int{333, 32768, 141}
                _ = v24
            }
        }
        p0 -= (p0 - (p0 & p0))
        var v25 int64 = 0
        _ = v25
    }
    p0 += 128
    v26 := []int{-888, 1, 128, 127}
    _ = p0
    _ = p1
    _ = v26
    return true
}

func main() {
    var v1 uint8 = 1
    var v2 int = -9223372036854775808
    var v3 uint8 = v1
    v4 := make(map[string]bool)
    v3 -= v3
    if true {
        for fi27 := 0; fi27 < 2; fi27++ {
            for fi28 := 0; fi28 < 19; fi28++ {
                if (319 < 1) {
                    v29 := GenS1{F0: (fi28 * (-(v2))), F1: 84, F2: 65535}
                    v30 := v4[("\n\t" + "")]
                    _ = v29
                    _ = v30
                }
                _ = fi28
            }
            switch (fi27 * v2) {
            case 717:
                switch ("hello" + ("a" + ("" + (("foo bar" + "hello") + "a")))) {
                case "":
                    for fi31 := 0; fi31 < 12; fi31++ {
                        var v32 uint = 669
                        _ = fi31
                        _ = v32
                    }
                default:
                    if true {
                        var v33 bool = false
                        fmt.Println(v33)
                        var v34 int16 = 1
                        _ = v33
                        _ = v34
                    }
                }
            case 32768:
                fmt.Println(v1, v3, v3)
            default:
                v35 := GenS3{F0: 0}
                _ = v35
            }
            if (!(!false)) {
                v1 = 1
            }
            _ = fi27
        }
    }
    v1 = (v3 | v1)
    if (32767 == 0) {
        v36 := []int{4294967295}
        for fi37 := 0; fi37 < 8; fi37++ {
            v2 = fi37
            _ = fi37
        }
        if true {
            var v38 string = "foo bar"
            _ = v38
        } else {
            v1 = 54
            if false {
                v1 += 2
            } else {
                var v39 int = (v2 - v2)
                var v40 uint32 = 2
                v40 = 65535
                _ = v39
                _ = v40
            }
            if (((1e10 / (1.5 / 1e10)) + -2.25) != -2.25) {
                var v41 int16 = 1000
                for fi42 := 0; fi42 < 7; fi42++ {
                    v43 := GenS4{F0: 459}
                    _ = fi42
                    _ = v43
                }
                _ = v41
            }
        }
        _ = v36
    }
    var v44 uint = 65535
    delete(v4, ((((("foo bar" + "a") + ("hello" + "foo bar")) + "hello") + "foo bar") + "hello"))
    var v45 string = "foo bar"
    var v46 uint8 = (v3 << ((v44) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v46
}
