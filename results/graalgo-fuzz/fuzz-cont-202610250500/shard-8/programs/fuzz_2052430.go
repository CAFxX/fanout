package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 int32) uint64 {
    var v5 uint16 = 2
    var v6 string = "hello"
    v7 := []int{32767, 1, -9223372036854775808}
    for fi8 := 0; fi8 < 13; fi8++ {
        v6 = v6
        var v9 uint = 291
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    return 690
}
func genFn1(p0 int32, p1 int16) string {
    var v10 string = ("a" + ("foo bar" + "hello"))
    p1 = -516
    for fi11 := 0; fi11 < 3; fi11++ {
        fmt.Println(p0)
        _ = fi11
    }
    _ = p0
    _ = p1
    _ = v10
    return ((v10 + ("foo bar" + ((v10 + v10) + (("a" + "foo bar") + v10)))) + ("hello" + ((v10 + ("" + ("foo bar" + v10))) + v10)))
}

func main() {
    var v1 string = "hello"
    v2 := make(map[int]string)
    var v3 uint32 = 2
    var v4 uint32 = v3
    v3 = v3
    v12 := v2[-385]
    fmt.Println(v12)
    v13 := GenS1{F0: v1, F1: 133}
    v14 := GenS1{F0: "", F1: 32767}
    switch (v1 + "hello") {
    case "hello":
        v1 = v1
    case "foo bar":
        switch -1 {
        case -577:
            var v15 int16 = -1
            _ = v15
        }
    }
    for fi16 := 0; fi16 < 13; fi16++ {
        fmt.Println(v12, v3)
        v2[32768] = (v1 + (((v1 + (v12 + "hello")) + "foo bar") + v1))
        _ = fi16
    }
    var v17 string = v12
    for fi18 := 0; fi18 < 3; fi18++ {
        v1 = v1
        fmt.Println(v4, v3, v1)
        fmt.Println(v17, fi18)
        _ = fi18
    }
    v19 := []int{32768, 0, -80}
    v12 = v1
    v4 -= 0
    v1 = v17
    for fi20 := 0; fi20 < 15; fi20++ {
        if (false || (!false)) {
            switch fi20 {
            case -140:
                if (!true) {
                    fmt.Println(v1, v1, v3)
                    v4 = (v4 - v4)
                }
            case 1:
                for fi21 := 0; fi21 < 8; fi21++ {
                    if len(v19) >= 2 {
                        v19 = v19[1:len(v19)]
                    }
                    switch -1 {
                    case 9223372036854775807:
                        v22 := v2[fi21]
                        _ = v22
                    case 128:
                        v23 := GenS2{F0: 689}
                        _ = v23
                    case -318:
                        switch 32767 {
                        case 4294967295:
                            var v24 uint64 = 614
                            _ = v24
                        case 637:
                            v17 = ((("a" + ((((("hello" + ((v17 + (v17 + "")) + ("\n\t" + ("a" + "hello")))) + "") + "hello") + (v1 + v17)) + "a")) + v1) + "")
                        default:
                            v25 := GenS1{F0: v1, F1: 42}
                            v26 := GenS1{F0: v12, F1: 32767}
                            _ = v25
                            _ = v26
                        }
                    }
                    if (-371 == fi20) {
                        v12 = ((v1 + ("a" + v1)) + ("" + v1))
                    } else {
                        var v27 int16 = -32768
                        var v28 bool = true
                        fmt.Println(v17, v3, v3)
                        _ = v27
                        _ = v28
                    }
                    if true {
                        switch v17 {
                        case "foo bar":
                            var v29 int64 = 0
                            _ = v29
                        case "\n\t":
                            var v30 uint32 = (v4 - v4)
                            _ = v30
                        }
                        var v31 bool = (!((4294967295 <= 531) || (9223372036854775807 >= 42)))
                        switch "foo bar" {
                        case "foo bar":
                            v32 := GenS2{F0: 277}
                            _ = v32
                        case "a":
                            var v33 uint8 = 128
                            _ = v33
                        }
                        _ = v31
                    } else {
                        v17 = "hello"
                        switch (fi20 / ((fi21) | 1)) {
                        case -386:
                            var v34 uint16 = 32767
                            _ = v34
                        case 355:
                            var v35 uint16 = 652
                            v12 = v12
                            _ = v35
                        case 4294967295:
                            switch 127 {
                            case 0:
                                v36 := GenS2{F0: 32767}
                                _ = v36
                            default:
                                _ = len(v2)
                            }
                        }
                        v4 += v4
                    }
                    _ = fi21
                }
            case 42:
                if len(v19) > 0 {
                    v19[0] = (fi20 * (fi20 % ((fi20) | 1)))
                }
            }
        }
        _ = fi20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
