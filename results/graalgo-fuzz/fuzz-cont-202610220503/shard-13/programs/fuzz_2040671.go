package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 uint; F1 string }
func genFn0(p0 bool) int8 {
    switch 32768 {
    case 1:
        for fi6 := 0; fi6 < 19; fi6++ {
            fmt.Println(p0)
            var v7 uint16 = 42
            _ = fi6
            _ = v7
        }
    default:
        var v8 bool = true
        _ = v8
    }
    v9 := GenS1{F0: "\n\t", F1: 1}
    p0 = ((!p0) || true)
    _ = p0
    _ = v9
    return 25
}

func main() {
    var v1 int = -496
    v2 := make(map[int]bool)
    v3 := []int{v1, v1}
    var v4 uint8 = 100
    var v5 float32 = 0.0
    for fi10 := 0; fi10 < 3; fi10++ {
        var v11 uint16 = 221
        var v12 string = ("a" + "a")
        v4 = (^(v4))
        _ = fi10
        _ = v11
        _ = v12
    }
    for fi13 := 0; fi13 < 3; fi13++ {
        var v14 int64 = 1
        _ = fi13
        _ = v14
    }
    var v15 float32 = v5
    if false {
        for fi16 := 0; fi16 < 11; fi16++ {
            for fi17 := 0; fi17 < 9; fi17++ {
                var v18 float32 = (v5 / v15)
                _ = fi17
                _ = v18
            }
            if false {
                var v19 int16 = 1000
                fmt.Println(v1)
                var v20 uint16 = 2
                _ = v19
                _ = v20
            } else {
                if true {
                    fmt.Println(v1, fi16)
                }
                v4 = (^(v4))
                if (!((0.0 - ((0.1 - 1.5) * 3.14159)) == 1e10)) {
                    v1 = (-(fi16))
                    v21 := GenS2{F0: 4294967295, F1: "foo bar"}
                    _ = v21
                }
            }
            _ = fi16
        }
    }
    v2[407] = false
    fmt.Println(v5, v1, v15)
    _ = len(v2)
    v22 := GenS2{F0: 255, F1: "hello"}
    var v23 float32 = 1e10
    switch 494 {
    case -577:
        if (2 == 967) {
            v24 := GenS2{F0: 1, F1: "\n\t"}
            _ = v24
        } else {
            if true {
                for fi25 := 0; fi25 < 15; fi25++ {
                    var v26 float64 = ((3.14159 - -2.25) - 1.5)
                    _ = fi25
                    _ = v26
                }
                v27 := GenS2{F0: 1, F1: ("" + (("foo bar" + ("\n\t" + "a")) + "foo bar"))}
                switch v1 {
                case 2147483647:
                    v28 := GenS2{F0: 0, F1: ((("hello" + "") + "hello") + "a")}
                    _ = v28
                case 511:
                    switch "hello" {
                    case "hello":
                        for fi29 := 0; fi29 < 14; fi29++ {
                            v30 := GenS2{F0: 2, F1: "a"}
                            var v31 float32 = v5
                            _ = fi29
                            _ = v30
                            _ = v31
                        }
                    case "a":
                        v32 := GenS1{F0: "hello", F1: -289}
                        _ = v32
                    }
                default:
                    for fi33 := 0; fi33 < 15; fi33++ {
                        for fi34 := 0; fi34 < 18; fi34++ {
                            if false {
                                var v35 uint = 243
                                v36 := GenS1{F0: ("\n\t" + ("hello" + "\n\t")), F1: -769}
                                v2[32767] = true
                                _ = v35
                                _ = v36
                            } else {
                                if (!true) {
                                    var v37 uint64 = 315
                                    v4 -= (^(v4))
                                    fmt.Println(v4, v15, v23)
                                    _ = v37
                                }
                            }
                            _ = fi34
                        }
                        switch v1 {
                        case 520:
                            v5 += float32(float64((float32((0.1 - 0.1)) + float32((1e10 + 1.5)))))
                        default:
                            v38 := GenS1{F0: "foo bar", F1: -32768}
                            _ = v38
                        }
                        for fi39 := 0; fi39 < 4; fi39++ {
                            var v40 bool = false
                            var v41 int8 = -75
                            var v42 uint32 = 0
                            _ = fi39
                            _ = v40
                            _ = v41
                            _ = v42
                        }
                        _ = fi33
                    }
                }
                _ = v27
            }
            v43 := GenS1{F0: "hello", F1: 1}
            _ = v43
        }
    default:
        v4 = (v4 & 2)
    }
    for fi44 := 0; fi44 < 20; fi44++ {
        var v45 float32 = ((v5 + v5) - (((float32(1e10) * v5) - v5) - (1e10 * (1.5 * ((v5 * 0.1) + v5)))))
        _ = fi44
        _ = v45
    }
    fmt.Println(v5, v1)
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
