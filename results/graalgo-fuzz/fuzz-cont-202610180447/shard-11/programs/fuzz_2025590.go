package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint16 }
func genFn0(p0 uint64, p1 int8) bool {
    fmt.Println(p1, p0)
    for fi4 := 0; fi4 < 17; fi4++ {
        p0 -= (p0 % ((776) | 1))
        var v5 int8 = (p1 - 0)
        var v6 int8 = (p1 << ((18446744073709551615) & 63))
        _ = fi4
        _ = v5
        _ = v6
    }
    p0 = p0
    p0 -= (^(p0))
    _ = p0
    _ = p1
    return false
}
func genFn1(p0 uint32, p1 bool, p2 uint32) uint32 {
    p0 -= p2
    if (79 <= 643) {
        v7 := []int{447}
        _ = v7
    }
    v8 := []int{175}
    for fi9 := 0; fi9 < 6; fi9++ {
        for fi10 := 0; fi10 < 2; fi10++ {
            var v11 string = ("hello" + (("\n\t" + ("" + "foo bar")) + "a"))
            if true {
                p0 -= (p0 - p0)
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                if ((!true) && p1) {
                    var v12 float64 = (-2.25 + 1.5)
                    p2 -= (p2 - p0)
                    _ = len(v8)
                    _ = v12
                }
            } else {
                if true {
                    v13 := GenS1{F0: 393}
                    p1 = false
                    _ = v13
                } else {
                    v8 = append(v8, 9223372036854775807)
                    var v14 bool = true
                    v15 := GenS1{F0: 255}
                    _ = v14
                    _ = v15
                }
                v11 = (v11 + v11)
            }
            p1 = p1
            _ = fi10
            _ = v11
        }
        var v16 int32 = -1
        p2 = p0
        _ = fi9
        _ = v16
    }
    var v17 uint16 = 519
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v8
    return (p2 / (((p0 << ((18446744073709551615) & 63))) | 1))
}

func main() {
    var v1 uint32 = 972
    v2 := []int{9223372036854775807, 4294967295, 281}
    v3 := make(map[int]int)
    fmt.Println(v1)
    for fi18 := 0; fi18 < 9; fi18++ {
        v1 = (v1 & (^(v1)))
        if (((("a" + ("hello" + "foo bar")) + "") + "a") != "a") {
            if (false || true) {
                var v19 uint16 = 2
                switch (("" + "hello") + "hello") {
                case "":
                    v1 = 4294967295
                case "foo bar":
                    v19 = 255
                default:
                    if false {
                        var v20 int8 = 127
                        _ = len(v2)
                        _ = v20
                    } else {
                        switch fi18 {
                        case 42:
                            var v21 float32 = 1e10
                            _ = v21
                        case -2147483648:
                            v19 += v19
                        case 2147483647:
                            v19 += 1
                        }
                        v2 = append(v2, 9223372036854775807)
                    }
                }
                var v22 bool = false
                _ = v19
                _ = v22
            } else {
                v1 += (v1 / (((v1 + (v1 & 2147483648))) | 1))
                v1 += v1
            }
            var v23 uint = 42
            var v24 int8 = 42
            _ = v23
            _ = v24
        } else {
            var v25 int8 = 127
            switch "foo bar" {
            case "":
                var v26 bool = ((true && true) || (42 > 18446744073709551615))
                _ = v26
            case "hello":
                var v27 uint = 904
                _ = v27
            case "\n\t":
                for fi28 := 0; fi28 < 17; fi28++ {
                    switch "a" {
                    case "hello":
                        _ = len(v3)
                    case "a":
                        v25 -= v25
                    case "":
                        _ = len(v2)
                    default:
                        var v29 int64 = -310
                        _ = v29
                    }
                    var v30 uint32 = (^(v1))
                    var v31 float32 = -2.25
                    _ = fi28
                    _ = v30
                    _ = v31
                }
            default:
                var v32 uint16 = 65535
                _ = v32
            }
            _ = v25
        }
        _ = fi18
    }
    fmt.Println(v1)
    switch -511 {
    case 42:
        v33 := GenS1{F0: 65535}
        _ = v33
    case 32767:
        var v34 uint32 = (v1 / ((2147483648) | 1))
        _ = v34
    case -1:
        for fi35 := 0; fi35 < 15; fi35++ {
            _ = len(v3)
            var v36 uint8 = 0
            _ = fi35
            _ = v36
        }
    }
    v2 = append(v2, int(v1))
    var v37 float32 = 1e10
    delete(v3, 656)
    v38 := GenS1{F0: 907}
    _ = len(v2)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if true {
        var v39 uint8 = 128
        v37 = 0.1
        _ = v39
    } else {
        v1 -= v1
        v1 = (v1 ^ (^(v1)))
        v40 := GenS2{F0: 42}
        _ = v40
    }
    var v41 bool = false
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v37:", v37)
    fmt.Println("v41:", v41)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v41
}
