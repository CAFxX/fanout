package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16; F2 int8 }
type GenS2 struct { F0 int32 }
func genFn0(p0 int64, p1 string, p2 bool) int16 {
    fmt.Println(p2, p1)
    v5 := []int{}
    if (p1 <= (("a" + "hello") + p1)) {
        var v6 bool = (42 <= 42)
        if v6 {
            if v6 {
                if (112 >= 0) {
                    _ = len(v5)
                    var v7 uint16 = 0
                    _ = v7
                }
                _ = len(v5)
            }
            if v6 {
                switch ((p1 + p1) + "a") {
                case "foo bar":
                    v8 := GenS1{F0: (p1 + ("" + p1)), F1: -1, F2: -70}
                    _ = v8
                default:
                    v9 := GenS1{F0: ("a" + (("" + ("" + "\n\t")) + ("a" + p1))), F1: 1000, F2: 1}
                    _ = v9
                }
                if ((p2 && (397 != 2)) && p2) {
                    var v10 int16 = 32767
                    for fi11 := 0; fi11 < 12; fi11++ {
                        for fi12 := 0; fi12 < 10; fi12++ {
                            switch ((p1 + ("" + p1)) + p1) {
                            case "a":
                                p1 = p1
                            }
                            _ = fi12
                        }
                        p1 = p1
                        v6 = p2
                        _ = fi11
                    }
                    var v13 int = 1
                    _ = v10
                    _ = v13
                }
            } else {
                var v14 uint8 = 2
                _ = v14
            }
        } else {
            if (v6 || ((p1 + "") >= "a")) {
                v5 = append(v5, 32768)
                var v15 float32 = (0.0 - (1e10 - 1.5))
                _ = v15
            } else {
                switch "hello" {
                case "hello":
                    if p2 {
                        v6 = v6
                        for fi16 := 0; fi16 < 17; fi16++ {
                            v17 := GenS2{F0: 1}
                            _ = fi16
                            _ = v17
                        }
                    }
                case "\n\t":
                    v18 := GenS1{F0: "a", F1: 32767, F2: -101}
                    _ = v18
                default:
                    var v19 uint16 = 42
                    _ = v19
                }
                var v20 uint = 42
                var v21 uint32 = 594
                _ = v20
                _ = v21
            }
            p2 = true
            fmt.Println(p0, p2)
        }
        fmt.Println(p1)
        _ = v6
    } else {
        for fi22 := 0; fi22 < 18; fi22++ {
            var v23 uint16 = 42
            _ = fi22
            _ = v23
        }
    }
    v24 := GenS1{F0: p1, F1: 42, F2: 0}
    _ = p0
    _ = p1
    _ = p2
    _ = v24
    _ = v5
    return 42
}
func genFn1() uint64 {
    v25 := []int{2147483647}
    v25 = append(v25, 908)
    if len(v25) > 0 {
        v25[0] = 32768
    }
    var v26 string = ("" + "foo bar")
    var v27 bool = false
    _ = v25
    _ = v26
    _ = v27
    return 653
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int64 = 1
    var v3 uint16 = 2
    var v4 uint = 4294967295
    v4 -= (^(v4))
    v4 = v4
    v28 := GenS1{F0: ("hello" + "\n\t"), F1: 104, F2: 73}
    if ((v2 >> ((v4) & 63)) != v2) {
        var v29 float64 = 0.0
        _ = v29
    }
    switch (("a" + "\n\t") + "hello") {
    case "":
        var v30 int8 = -1
        _ = v30
    default:
        fmt.Println(v4)
    }
    v31 := []int{0, -536, int(v3)}
    v3 += (v3 - 198)
    var v32 int8 = -45
    for fi33 := 0; fi33 < 12; fi33++ {
        v34 := GenS2{F0: 1}
        if len(v31) > 0 {
            v31[0] = 128
        }
        v32 = (-(v32))
        _ = fi33
        _ = v34
    }
    if len(v31) >= 2 {
        v31 = v31[1:len(v31)]
    }
    for fi35 := 0; fi35 < 16; fi35++ {
        v4 = v4
        _ = fi35
    }
    var v36 uint8 = 0
    _ = len(v1)
    fmt.Println(v36)
    fmt.Println(v36, v32)
    v32 -= 100
    v37 := GenS2{F0: 222}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v32:", v32)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v31
    _ = v32
    _ = v36
    _ = v37
    _ = v4
}
