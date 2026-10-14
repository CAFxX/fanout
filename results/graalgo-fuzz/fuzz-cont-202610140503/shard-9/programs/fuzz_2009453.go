package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int32; F2 int }
type GenS2 struct { F0 uint8; F1 uint16; F2 int64 }
func genFn0(p0 uint) int {
    p0 = (p0 | p0)
    switch -142 {
    case 128:
        v5 := []int{-1, 4294967295, 4294967295}
        _ = v5
    }
    p0 -= (p0 & p0)
    var v6 uint = (^(p0))
    switch -630 {
    case 127:
        v6 += p0
    case 42:
        v7 := GenS1{F0: "foo bar", F1: -583, F2: 32768}
        _ = v7
    case -2147483648:
        v8 := []int{}
        _ = v8
    }
    _ = p0
    _ = v6
    return -296
}
func genFn1(p0 bool, p1 int16) int {
    v9 := []int{genFn0(2), 350, 127}
    _ = len(v9)
    v10 := GenS2{F0: 159, F1: 2, F2: 9223372036854775807}
    v11 := GenS2{F0: 127, F1: 491, F2: 615}
    switch 42 {
    case -748:
        fmt.Println(p1)
    case 32768:
        if ("hello" == ("a" + "foo bar")) {
            switch 127 {
            case -574:
                if p0 {
                    switch "foo bar" {
                    case "hello":
                        if p0 {
                            fmt.Println(p1, p0)
                            switch "\n\t" {
                            case "":
                                fmt.Println(p1, p1)
                            }
                        }
                    case "a":
                        p1 -= (-(p1))
                    default:
                        for fi12 := 0; fi12 < 2; fi12++ {
                            switch int(p1) {
                            case -9223372036854775808:
                                fmt.Println(p0)
                            case -596:
                                p0 = (p0 || p0)
                            default:
                                for fi13 := 0; fi13 < 13; fi13++ {
                                    var v14 int16 = (p1 | p1)
                                    var v15 uint16 = 32767
                                    v9 = append(v9, fi13)
                                    _ = fi13
                                    _ = v14
                                    _ = v15
                                }
                            }
                            _ = fi12
                        }
                    }
                    v16 := GenS2{F0: 255, F1: 32767, F2: 761}
                    p0 = (-1 <= 1)
                    _ = v16
                } else {
                    p0 = true
                    var v17 uint64 = 4294967295
                    if (1 == 42) {
                        switch 2147483647 {
                        case 127:
                            if len(v9) >= 2 {
                                v9 = v9[1:len(v9)]
                            }
                        }
                    }
                    _ = v17
                }
            default:
                if p0 {
                    var v18 int16 = 1000
                    v19 := GenS2{F0: 146, F1: 65535, F2: 0}
                    _ = v18
                    _ = v19
                }
            }
            var v20 string = ""
            for fi21 := 0; fi21 < 13; fi21++ {
                p0 = true
                var v22 int32 = -828
                p0 = ((-2.25 != (1.5 * 1e10)) || (p0 && p0))
                _ = fi21
                _ = v22
            }
            _ = v20
        } else {
            var v23 int16 = (p1 % ((p1) | 1))
            if false {
                var v24 uint16 = 65535
                p1 = v23
                _ = v24
            }
            _ = v23
        }
    }
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v9
    return 9223372036854775807
}

func main() {
    v1 := make(map[uint32]int)
    var v2 string = (("foo bar" + "\n\t") + ("foo bar" + ("hello" + ("" + "hello"))))
    var v3 int16 = -402
    var v4 int64 = 9223372036854775807
    var v25 int16 = v3
    var v26 bool = (!((false && true) && true))
    for fi27 := 0; fi27 < 1; fi27++ {
        var v28 uint = 18446744073709551615
        for fi29 := 0; fi29 < 12; fi29++ {
            var v30 int16 = (^(v25))
            var v31 int8 = -1
            if (!(v2 >= ((((v2 + ((("a" + v2) + "") + "foo bar")) + "") + (v2 + ((v2 + ((v2 + "hello") + "foo bar")) + v2))) + (v2 + v2)))) {
                if v26 {
                    v31 = v31
                }
                _ = len(v1)
            }
            _ = fi29
            _ = v30
            _ = v31
        }
        var v32 uint = v28
        _ = fi27
        _ = v28
        _ = v32
    }
    var v33 uint16 = 2
    v34 := []int{-555}
    for fi35 := 0; fi35 < 15; fi35++ {
        v36 := GenS1{F0: "", F1: -136, F2: fi35}
        _ = fi35
        _ = v36
    }
    var v37 int8 = 42
    var v38 int16 = (v3 & v3)
    v33 += 0
    var v39 uint32 = 0
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("slicelen v34:", len(v34), v34)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v33
    _ = v34
    _ = v37
    _ = v38
    _ = v39
    _ = v4
}
