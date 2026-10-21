package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8 }
type GenS2 struct { F0 string; F1 int; F2 string }
func genFn0(p0 uint64, p1 int64, p2 uint8) int64 {
    p1 = p1
    var v5 uint32 = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return p1
}

func main() {
    var v1 int8 = 42
    v2 := make(map[uint32]bool)
    var v3 int = -182
    var v4 bool = true
    switch ((("\n\t" + "a") + ("\n\t" + "foo bar")) + ("foo bar" + ("" + "\n\t"))) {
    case "hello":
        for fi6 := 0; fi6 < 6; fi6++ {
            fmt.Println(fi6)
            v4 = v4
            _ = fi6
        }
    case "a":
        v7 := GenS1{F0: 752, F1: v1}
        _ = v7
    case "\n\t":
        var v8 uint32 = 0
        _ = v8
    default:
        switch ("\n\t" + (("hello" + "hello") + "\n\t")) {
        case "foo bar":
            _ = len(v2)
        case "a":
            for fi9 := 0; fi9 < 11; fi9++ {
                if true {
                    for fi10 := 0; fi10 < 20; fi10++ {
                        v3 = 24
                        v11 := []int{v3, v3, (fi10 >> ((65535) & 63)), v3}
                        if len(v11) > 0 {
                            v11[0] = (-(fi9))
                        }
                        _ = fi10
                        _ = v11
                    }
                    if v4 {
                        for fi12 := 0; fi12 < 5; fi12++ {
                            switch ("" + "") {
                            case "a":
                                v1 = -104
                            default:
                                var v13 uint16 = 215
                                _ = v13
                            }
                            _ = fi12
                        }
                        if true {
                            var v14 uint32 = 42
                            _ = v14
                        } else {
                            var v15 uint64 = 556
                            if v4 {
                                var v16 float32 = -2.25
                                _ = v16
                            } else {
                                if false {
                                    v17 := v2[1]
                                    var v18 float64 = float64(0.0)
                                    _ = v17
                                    _ = v18
                                }
                                v3 = v3
                            }
                            _ = v15
                        }
                    }
                    v3 -= -2147483648
                } else {
                    if v4 {
                        v1 = v1
                    } else {
                        for fi19 := 0; fi19 < 18; fi19++ {
                            switch (fi19 | (fi19 + fi9)) {
                            case 1:
                                v20 := v2[1]
                                _ = v20
                            default:
                                v1 = (v1 / ((v1) | 1))
                            }
                            if false {
                                switch "a" {
                                case "hello":
                                    fmt.Println(fi19)
                                }
                                v1 = (v1 << ((42) & 63))
                            } else {
                                var v21 int64 = 1
                                var v22 uint64 = 923
                                var v23 string = (("\n\t" + "\n\t") + "\n\t")
                                _ = v21
                                _ = v22
                                _ = v23
                            }
                            v24 := GenS2{F0: "hello", F1: (v3 - (^(fi19))), F2: (("a" + "foo bar") + "hello")}
                            v25 := GenS2{F0: "foo bar", F1: (fi19 * fi19), F2: "\n\t"}
                            v26 := GenS1{F0: 255, F1: 127}
                            _ = fi19
                            _ = v24
                            _ = v25
                            _ = v26
                        }
                        v3 += v3
                        for fi27 := 0; fi27 < 7; fi27++ {
                            v28 := []int{fi9, 0, fi27}
                            _ = fi27
                            _ = v28
                        }
                    }
                    if (!v4) {
                        var v29 uint = 65535
                        _ = v29
                    } else {
                        v4 = ((2 >= 42) || false)
                        var v30 string = "hello"
                        _ = v30
                    }
                }
                _ = fi9
            }
        }
    }
    v1 += -1
    fmt.Println(v3, v3)
    v31 := GenS1{F0: 646, F1: (v1 - v1)}
    var v32 bool = ((float32((0.1 - (0.1 + 3.14159))) * (0.1 * (-2.25 + float32(1e10)))) >= 1.5)
    if v32 {
        fmt.Println(v1)
        fmt.Println(v32)
        var v33 float32 = float32(3.14159)
        _ = v33
    } else {
        v34 := []int{(v3 + v3)}
        fmt.Println(v3, v3)
        for fi35 := 0; fi35 < 11; fi35++ {
            v36 := GenS2{F0: "foo bar", F1: 9223372036854775807, F2: ""}
            var v37 int32 = -2147483648
            _ = fi35
            _ = v36
            _ = v37
        }
        _ = v34
    }
    _ = len(v2)
    v38 := []int{}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v38:", len(v38), v38)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v38
    _ = v4
}
