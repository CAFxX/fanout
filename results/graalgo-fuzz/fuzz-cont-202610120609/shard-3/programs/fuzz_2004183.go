package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint64; F1 int32 }
type GenS5 struct { F0 int16; F1 uint; F2 int32 }
func genFn0() bool {
    var v6 float32 = 1e10
    v6 += v6
    if true {
        var v7 bool = true
        v7 = false
        v8 := GenS1{F0: ("" + "\n\t")}
        _ = v7
        _ = v8
    } else {
        for fi9 := 0; fi9 < 10; fi9++ {
            var v10 int16 = -154
            if false {
                var v11 uint16 = 432
                v10 -= (v10 | v10)
                _ = v11
            }
            _ = fi9
            _ = v10
        }
        if (685 > 1) {
            v12 := GenS1{F0: (((((("" + "\n\t") + ("hello" + "a")) + (("foo bar" + "foo bar") + "foo bar")) + "\n\t") + ("a" + "foo bar")) + ("foo bar" + ""))}
            v13 := []int{}
            if (!(32767 <= 796)) {
                switch 32768 {
                case 133:
                    v6 -= v6
                default:
                    switch "a" {
                    case "hello":
                        for fi14 := 0; fi14 < 18; fi14++ {
                            if (!false) {
                                if len(v13) >= 2 {
                                    v13 = v13[1:len(v13)]
                                }
                                if len(v13) > 0 {
                                    v13[0] = fi14
                                }
                            }
                            if (28 < 42) {
                                v6 = v6
                            } else {
                                if true {
                                    v15 := GenS2{F0: -1}
                                    _ = v15
                                }
                                var v16 uint8 = 127
                                _ = v16
                            }
                            _ = fi14
                        }
                    }
                }
                if len(v13) >= 2 {
                    v13 = v13[1:len(v13)]
                }
                if (true && false) {
                    for fi17 := 0; fi17 < 14; fi17++ {
                        switch (fi17 - (fi17 % ((32767) | 1))) {
                        case -391:
                            var v18 int32 = -2147483648
                            _ = v18
                        case 109:
                            var v19 int = fi17
                            _ = v19
                        }
                        v6 = -2.25
                        if len(v13) >= 2 {
                            v13 = v13[1:len(v13)]
                        }
                        _ = fi17
                    }
                    fmt.Println(v6)
                }
            }
            _ = v12
            _ = v13
        }
    }
    var v20 uint = 4294967295
    fmt.Println(v20)
    _ = v20
    _ = v6
    return (-1 < -9223372036854775808)
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []int{}
    v3 := []string{"a"}
    v4 := []string{"foo bar"}
    var v5 int64 = -409
    switch 42 {
    case -572:
        _ = len(v4)
    }
    for fi21 := 0; fi21 < 19; fi21++ {
        v5 += v5
        v1[4294967295] = 2
        var v22 uint8 = 2
        _ = fi21
        _ = v22
    }
    var v23 int = 127
    if (!((32767 < 32767) && true)) {
        if (4294967295 < 2) {
            switch ("a" + ("foo bar" + ("hello" + "\n\t"))) {
            case "foo bar":
                fmt.Println(v5, v23)
            default:
                var v24 int16 = 1000
                _ = v24
            }
            if (647 >= 2) {
                for fi25 := 0; fi25 < 18; fi25++ {
                    v26 := GenS3{F0: 42}
                    v27 := v1[42]
                    _ = fi25
                    _ = v26
                    _ = v27
                }
                v5 = v5
                var v28 float64 = 1.5
                _ = v28
            } else {
                var v29 int = 9223372036854775807
                v30 := GenS4{F0: 779, F1: 1}
                _ = v29
                _ = v30
            }
            for fi31 := 0; fi31 < 17; fi31++ {
                _ = len(v1)
                _ = fi31
            }
        }
    }
    var v32 float32 = 1e10
    v5 = v5
    for fi33 := 0; fi33 < 5; fi33++ {
        var v34 float64 = 3.14159
        v23 -= 9223372036854775807
        v32 += (((1.5 + 3.14159) + v32) + v32)
        _ = fi33
        _ = v34
    }
    v35 := GenS3{F0: 0}
    v36 := GenS3{F0: 1}
    v23 = -564
    v37 := GenS5{F0: 1, F1: 65535, F2: 96}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v32:", v32)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v3
    _ = v32
    _ = v35
    _ = v36
    _ = v37
    _ = v4
    _ = v5
}
