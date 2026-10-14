package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int8 }
type GenS2 struct { F0 int; F1 bool }
type GenS3 struct { F0 string; F1 int16; F2 int8 }
func genFn0(p0 int64, p1 bool, p2 int) int16 {
    if p1 {
        var v5 uint16 = 1
        for fi6 := 0; fi6 < 15; fi6++ {
            var v7 uint32 = 2
            _ = fi6
            _ = v7
        }
        _ = v5
    } else {
        for fi8 := 0; fi8 < 8; fi8++ {
            var v9 float32 = ((0.0 / float32((0.1 / -2.25))) - ((0.1 - float32(1e10)) - 3.14159))
            for fi10 := 0; fi10 < 4; fi10++ {
                var v11 bool = p1
                v12 := GenS1{F0: "foo bar", F1: 42}
                _ = fi10
                _ = v11
                _ = v12
            }
            _ = fi8
            _ = v9
        }
        if true {
            v13 := GenS1{F0: "a", F1: 0}
            var v14 float32 = 0.0
            var v15 uint64 = 2
            _ = v13
            _ = v14
            _ = v15
        } else {
            if true {
                fmt.Println(p0, p1, p2)
                p0 = -1
            }
            fmt.Println(p0)
            switch ("a" + (("\n\t" + "a") + ((("a" + "hello") + "\n\t") + "a"))) {
            case "\n\t":
                if (!p1) {
                    var v16 string = "a"
                    _ = v16
                } else {
                    switch p2 {
                    case 681:
                        for fi17 := 0; fi17 < 17; fi17++ {
                            p2 += (-(p2))
                            if true {
                                p0 = p0
                            } else {
                                for fi18 := 0; fi18 < 11; fi18++ {
                                    var v19 uint64 = 0
                                    _ = fi18
                                    _ = v19
                                }
                                p1 = ((fi17 % (((p2 >> ((42) & 63))) | 1)) >= (-(p2)))
                                var v20 int16 = 1000
                                _ = v20
                            }
                            fmt.Println(fi17, fi17, fi17)
                            _ = fi17
                        }
                    }
                    switch (("a" + "foo bar") + "") {
                    case "":
                        for fi21 := 0; fi21 < 18; fi21++ {
                            var v22 uint16 = 2
                            _ = fi21
                            _ = v22
                        }
                    case "foo bar":
                        var v23 uint16 = 65535
                        _ = v23
                    case "\n\t":
                        v24 := GenS1{F0: ("" + "a"), F1: 1}
                        _ = v24
                    default:
                        var v25 float32 = 3.14159
                        _ = v25
                    }
                }
            case "a":
                p1 = ((!(!p1)) || p1)
            }
        }
        if true {
            v26 := []int{(p2 ^ (p2 ^ 42)), (p2 / ((p2) | 1)), p2, (^(p2))}
            _ = v26
        }
    }
    v27 := []int{}
    _ = p0
    _ = p1
    _ = p2
    _ = v27
    return 392
}
func genFn1(p0 bool, p1 int32, p2 uint) int16 {
    for fi28 := 0; fi28 < 16; fi28++ {
        v29 := GenS1{F0: (("\n\t" + "hello") + "foo bar"), F1: 100}
        p2 += (^(p2))
        _ = fi28
        _ = v29
    }
    v30 := []int{1}
    if len(v30) > 0 {
        v30[0] = -9223372036854775808
    }
    for fi31 := 0; fi31 < 16; fi31++ {
        var v32 float32 = 1.5
        v33 := GenS1{F0: "a", F1: 42}
        v34 := GenS2{F0: (fi31 ^ fi31), F1: false}
        _ = fi31
        _ = v32
        _ = v33
        _ = v34
    }
    if ("" >= "\n\t") {
        var v35 uint8 = 255
        p2 += (p2 - 963)
        _ = v35
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v30
    return 0
}

func main() {
    var v1 uint = 42
    v2 := make(map[string]bool)
    var v3 uint16 = 0
    var v4 int64 = 621
    var v36 uint8 = 144
    var v37 int = 738
    if true {
        v4 = v4
    } else {
        v38 := GenS2{F0: 119, F1: (false && (-1 >= 32767))}
        _ = v38
    }
    v36 = 127
    v4 = v4
    var v39 bool = (uint64(v3) != 42)
    if (4 >= 42) {
        v3 = v3
        if ((!((v1 + v1) == uint(v36))) && true) {
            v1 += (v1 ^ (v1 + v1))
            v40 := GenS1{F0: "\n\t", F1: 34}
            for fi41 := 0; fi41 < 16; fi41++ {
                v3 = v3
                _ = fi41
            }
            _ = v40
        } else {
            switch ("" + "hello") {
            case "\n\t":
                for fi42 := 0; fi42 < 17; fi42++ {
                    v43 := []int{(fi42 / (((^(v37))) | 1))}
                    _ = fi42
                    _ = v43
                }
            }
            for fi44 := 0; fi44 < 10; fi44++ {
                if true {
                    var v45 string = ("" + "hello")
                    v46 := GenS3{F0: v45, F1: int16(v37), F2: 100}
                    var v47 uint16 = (v3 + v3)
                    _ = v45
                    _ = v46
                    _ = v47
                } else {
                    fmt.Println(v1, v1, v39)
                }
                _ = fi44
            }
        }
        v37 -= 32768
    }
    v39 = true
    v36 = v36
    var v48 float32 = (0.1 + (3.14159 * -2.25))
    delete(v2, ("\n\t" + "a"))
    switch "hello" {
    case "foo bar":
        var v49 int64 = v4
        _ = v49
    case "hello":
        v48 = v48
    case "a":
        v50 := []int{v37, (v37 % ((128) | 1)), 623}
        _ = v50
    }
    for fi51 := 0; fi51 < 12; fi51++ {
        v52 := v2["foo bar"]
        _ = len(v2)
        for fi53 := 0; fi53 < 4; fi53++ {
            v3 = v3
            _ = fi53
        }
        _ = fi51
        _ = v52
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v39
    _ = v4
    _ = v48
}
