package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint16 }
func genFn0() uint32 {
    if (-1 < -1) {
        v5 := []int{32767, -1, 32767, 32768}
        var v6 int32 = 1
        _ = v5
        _ = v6
    } else {
        var v7 int16 = 1
        _ = v7
    }
    var v8 float32 = 1e10
    _ = v8
    return 392
}

func main() {
    v1 := make(map[uint32]bool)
    v2 := []int{-9223372036854775808}
    var v3 float64 = 3.14159
    var v4 bool = (true || true)
    v4 = true
    v4 = (42 >= 42)
    for fi9 := 0; fi9 < 15; fi9++ {
        v4 = true
        if true {
            v4 = (!(!(1 < 375)))
        }
        var v10 uint8 = 127
        _ = fi9
        _ = v10
    }
    v11 := GenS1{F0: "foo bar"}
    if (v4 && (!(!(-1 == 42)))) {
        for fi12 := 0; fi12 < 3; fi12++ {
            var v13 uint64 = 2
            v4 = v4
            _ = fi12
            _ = v13
        }
        var v14 string = ""
        _ = v14
    } else {
        v3 = (3.14159 + (v3 + (v3 + (v3 + v3))))
    }
    v15 := GenS1{F0: "a"}
    switch 957 {
    case 0:
        if (42 >= 32767) {
            if false {
                delete(v1, 2)
                v2 = append(v2, -979)
            } else {
                v4 = (genFn0() <= genFn0())
                v3 += 0.0
                v3 += 3.14159
            }
            v1[2147483648] = v4
        }
    }
    v4 = true
    v3 = v3
    v4 = ((9223372036854775807 >= 554) || true)
    v16 := GenS2{F0: 32767}
    v17 := GenS2{F0: 2}
    switch (("hello" + ("a" + ("\n\t" + ("a" + "")))) + "") {
    case "\n\t":
        var v18 uint = 2
        _ = v18
    default:
        for fi19 := 0; fi19 < 1; fi19++ {
            v4 = ((2 != 7) && (!(true || v4)))
            _ = fi19
        }
    }
    switch -906 {
    case -2147483648:
        switch 1 {
        case -1:
            if (v4 && true) {
                fmt.Println(v4, v4)
            } else {
                fmt.Println(v3)
                for fi20 := 0; fi20 < 13; fi20++ {
                    var v21 uint8 = 0
                    v2 = append(v2, 1)
                    for fi22 := 0; fi22 < 14; fi22++ {
                        var v23 bool = ((882 >= 42) || ((fi22 / (((fi20 % ((fi22) | 1))) | 1)) < (fi22 / ((32768) | 1))))
                        switch fi20 {
                        case 748:
                            var v24 int32 = 2147483647
                            _ = v24
                        default:
                            if (v3 >= v3) {
                                v25 := GenS2{F0: 1}
                                delete(v1, 4294967295)
                                _ = len(v2)
                                _ = v25
                            }
                        }
                        v3 = (v3 + v3)
                        var v26 int32 = 2147483647
                        _ = len(v1)
                        switch (-(fi20)) {
                        case 9223372036854775807:
                            if ("" > "a") {
                                var v27 int64 = 9223372036854775807
                                _ = v27
                            }
                        default:
                            if (32767 > 1) {
                                switch "" {
                                case "":
                                    v3 += 1.5
                                case "a":
                                    v26 = v26
                                case "hello":
                                    v4 = (!false)
                                }
                                var v28 int = (fi20 ^ -1)
                                _ = v28
                            }
                        }
                        _ = fi22
                        _ = v23
                        _ = v26
                    }
                    _ = fi20
                    _ = v21
                }
                v3 = v3
            }
        default:
            var v29 uint64 = 4294967295
            _ = v29
        }
    case 32767:
        v30 := GenS2{F0: 735}
        _ = v30
    case 128:
        var v31 int32 = 550
        _ = v31
    default:
        var v32 uint64 = 2
        _ = v32
    }
    v33 := GenS2{F0: 255}
    v34 := GenS1{F0: ("" + "a")}
    if true {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v33
    _ = v34
    _ = v4
}
