package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64 }
type GenS2 struct { F0 uint8 }
func genFn0(p0 int, p1 uint64, p2 uint) uint64 {
    if (false && (42 < -1)) {
        p0 = p0
        v7 := GenS1{F0: (("a" + "\n\t") + "foo bar"), F1: -1}
        v8 := GenS1{F0: (("a" + "hello") + ("\n\t" + ("\n\t" + ((("foo bar" + "a") + "a") + "hello")))), F1: 9223372036854775807}
        _ = v7
        _ = v8
    } else {
        v9 := GenS1{F0: ("\n\t" + "hello"), F1: 432}
        _ = v9
    }
    for fi10 := 0; fi10 < 3; fi10++ {
        if (false || (-9223372036854775808 == -9223372036854775808)) {
            p0 += (-(p0))
            var v11 uint64 = 18446744073709551615
            var v12 int64 = -1
            _ = v11
            _ = v12
        }
        _ = fi10
    }
    fmt.Println(p2, p0)
    p2 += p2
    v13 := []int{p0, 32768}
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 861
}
func genFn1() uint {
    v14 := GenS1{F0: "foo bar", F1: 42}
    var v15 bool = false
    _ = v14
    _ = v15
    return 888
}

func main() {
    var v1 int = 96
    var v2 int16 = 1000
    var v3 string = ((("hello" + "hello") + "") + (("foo bar" + "\n\t") + ""))
    v4 := make(map[string]int)
    var v5 bool = true
    v6 := []int{0, 127, v1, (v1 - (v1 >> ((2) & 63)))}
    if true {
        var v16 float64 = -2.25
        for fi17 := 0; fi17 < 15; fi17++ {
            v3 = v3
            _ = fi17
        }
        var v18 int16 = v2
        _ = v16
        _ = v18
    }
    v6 = append(v6, 390)
    v3 = v3
    switch v1 {
    case 127:
        v3 = ""
    default:
        v2 -= (v2 % ((v2) | 1))
    }
    var v19 string = v3
    switch (-(v1)) {
    case -1:
        if v5 {
            var v20 bool = true
            var v21 int32 = 0
            v1 = v1
            _ = v20
            _ = v21
        }
    case 127:
        for fi22 := 0; fi22 < 10; fi22++ {
            var v23 float32 = 0.1
            _ = fi22
            _ = v23
        }
    case 1:
        switch v19 {
        case "\n\t":
            v19 = v19
        case "a":
            var v24 int = v1
            _ = v24
        default:
            fmt.Println(v5)
        }
    default:
        var v25 string = ("\n\t" + v19)
        _ = v25
    }
    v5 = v5
    if (-1 <= -2147483648) {
        v2 = v2
        for fi26 := 0; fi26 < 10; fi26++ {
            var v27 bool = v5
            v3 = v19
            _ = fi26
            _ = v27
        }
        for fi28 := 0; fi28 < 9; fi28++ {
            v6 = append(v6, v1)
            v5 = true
            _ = fi28
        }
    } else {
        for fi29 := 0; fi29 < 1; fi29++ {
            v30 := v4[v3]
            v19 = v3
            _ = fi29
            _ = v30
        }
        if false {
            if v5 {
                v3 = "a"
                for fi31 := 0; fi31 < 11; fi31++ {
                    if v5 {
                        var v32 bool = v5
                        v3 = v3
                        for fi33 := 0; fi33 < 1; fi33++ {
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            var v34 bool = v5
                            if (!((9223372036854775807 >= fi31) || ((genFn0((-(fi33)), 0, 377) != 1) && (!(((true && v32) || true) && (float32(-2.25) > 0.1)))))) {
                                var v35 int16 = v2
                                _ = v35
                            } else {
                                v32 = (2 > 315)
                                v36 := GenS1{F0: v19, F1: 0}
                                _ = v36
                            }
                            _ = fi33
                            _ = v34
                        }
                        _ = v32
                    } else {
                        switch ("foo bar" + "") {
                        case "hello":
                            _ = len(v4)
                        default:
                            delete(v4, v19)
                        }
                        v37 := GenS2{F0: 1}
                        var v38 int16 = int16(fi31)
                        _ = v37
                        _ = v38
                    }
                    var v39 uint32 = 42
                    switch v1 {
                    case 1:
                        v2 = (v2 - (v2 | (v2 * 42)))
                    case -352:
                        v3 = "a"
                    }
                    _ = fi31
                    _ = v39
                }
            } else {
                fmt.Println(v2, v2)
                if (!v5) {
                    v40 := GenS1{F0: v3, F1: -1}
                    _ = v40
                }
                v2 = v2
            }
            fmt.Println(v1)
        } else {
            v3 = v19
            var v41 uint64 = 42
            _ = v41
        }
        if v5 {
            delete(v4, (v19 + v3))
        }
    }
    v2 = -32768
    for fi42 := 0; fi42 < 8; fi42++ {
        if (fi42 > (v1 | fi42)) {
            switch -1 {
            case 127:
                var v43 int32 = 241
                _ = v43
            case 46:
                if true {
                    v44 := GenS2{F0: 127}
                    _ = v44
                } else {
                    var v45 uint = 65535
                    _ = v45
                }
            }
            fmt.Println(v3)
            var v46 int = (v1 << ((1) & 63))
            _ = v46
        }
        v2 = (v2 & -346)
        _ = fi42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
