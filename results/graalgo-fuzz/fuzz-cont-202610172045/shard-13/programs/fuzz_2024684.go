package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 bool; F1 uint64; F2 int16 }
type GenS3 struct { F0 int; F1 uint; F2 bool }
type GenS4 struct { F0 uint32; F1 bool }
type GenS5 struct { F0 uint16; F1 uint32 }
func genFn0() bool {
    v7 := GenS1{F0: "foo bar", F1: 0}
    v8 := []int{32767, 32767, 1, -9223372036854775808}
    if len(v8) > 0 {
        v8[0] = 32768
    }
    _ = v7
    _ = v8
    return false
}
func genFn1(p0 int16, p1 bool) uint {
    var v9 int8 = 100
    if p1 {
        var v10 bool = ((((1.5 / 1.5) + ((-2.25 + (((-2.25 + ((-2.25 * (-2.25 / (-2.25 / ((3.14159 / (1e10 - float64((3.14159 + 0.0)))) / 0.1)))) - -2.25)) + float64(float32(3.14159))) + 1e10)) + 1e10)) + float64((0.0 - (1.5 + 3.14159)))) < 0.1)
        if (!v10) {
            fmt.Println(v10)
            var v11 int8 = v9
            switch 0 {
            case 1:
                p0 += 42
            default:
                var v12 int8 = v9
                _ = v12
            }
            _ = v11
        }
        fmt.Println(v10)
        _ = v10
    } else {
        for fi13 := 0; fi13 < 5; fi13++ {
            for fi14 := 0; fi14 < 8; fi14++ {
                switch "" {
                case "foo bar":
                    var v15 float64 = float64(1e10)
                    _ = v15
                default:
                    p1 = p1
                }
                _ = fi14
            }
            for fi16 := 0; fi16 < 16; fi16++ {
                var v17 float64 = 1.5
                v9 -= -78
                var v18 int64 = 1
                _ = fi16
                _ = v17
                _ = v18
            }
            _ = fi13
        }
        var v19 int64 = -23
        _ = v19
    }
    if false {
        if (!false) {
            var v20 bool = p1
            v21 := GenS1{F0: "a", F1: 2}
            _ = v20
            _ = v21
        }
    } else {
        var v22 uint16 = 255
        v23 := []int{2147483647, 97, -482, 128}
        _ = v22
        _ = v23
    }
    var v24 uint64 = 861
    switch 747 {
    case 9223372036854775807:
        for fi25 := 0; fi25 < 3; fi25++ {
            if p1 {
                fmt.Println(fi25, v24, p0)
                switch 2147483647 {
                case -9223372036854775808:
                    v26 := []int{(fi25 >> ((781) & 63)), 32768, -439}
                    _ = v26
                default:
                    v27 := GenS1{F0: "hello", F1: v24}
                    _ = v27
                }
                v9 += (v9 << ((791) & 63))
            } else {
                v28 := GenS1{F0: (("" + "") + ("\n\t" + "foo bar")), F1: v24}
                for fi29 := 0; fi29 < 2; fi29++ {
                    var v30 string = (("foo bar" + "\n\t") + ("a" + "foo bar"))
                    v31 := GenS2{F0: false, F1: 79, F2: p0}
                    _ = fi29
                    _ = v30
                    _ = v31
                }
                _ = v28
            }
            for fi32 := 0; fi32 < 4; fi32++ {
                var v33 uint32 = 1
                _ = fi32
                _ = v33
            }
            v34 := []int{}
            _ = fi25
            _ = v34
        }
    }
    _ = p0
    _ = p1
    _ = v24
    _ = v9
    return 1
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint = 255
    v3 := []string{"\n\t", "\n\t"}
    var v4 int64 = 9223372036854775807
    v5 := []string{"hello", "foo bar"}
    var v6 int16 = 32767
    var v35 uint8 = 127
    v35 = v35
    var v36 float32 = float32(((0.1 + -2.25) / (3.14159 + 1e10)))
    var v37 int = 1
    v38 := GenS3{F0: int(v35), F1: 0, F2: (-2147483648 > -2147483648)}
    var v39 int64 = (v4 - v4)
    if false {
        v40 := GenS4{F0: 897, F1: false}
        var v41 uint64 = 885
        if ((false || false) || false) {
            switch ("\n\t" + "hello") {
            case "hello":
                for fi42 := 0; fi42 < 18; fi42++ {
                    if (true && true) {
                        switch v37 {
                        case 1:
                            var v43 float32 = v36
                            _ = v43
                        }
                        switch ("hello" + "hello") {
                        case "hello":
                            switch (("\n\t" + "a") + "\n\t") {
                            case "a":
                                v2 = v2
                            default:
                                v44 := GenS5{F0: 1, F1: 42}
                                _ = v44
                            }
                        case "\n\t":
                            switch int(v41) {
                            case 42:
                                if false {
                                    _ = len(v5)
                                }
                            case 711:
                                v5 = append(v5, "")
                            case 128:
                                v36 += float32((-2.25 * 1e10))
                            }
                        case "":
                            v1[2147483648] = false
                        default:
                            for fi45 := 0; fi45 < 5; fi45++ {
                                v35 += (v35 | 255)
                                _ = fi45
                            }
                        }
                    }
                    _ = fi42
                }
            case "\n\t":
                fmt.Println(v6)
            default:
                v4 = (v39 % ((v4) | 1))
            }
        } else {
            v2 -= genFn1((v6 / ((v6) | 1)), (633 <= 441))
            fmt.Println(v39, v2, v39)
        }
        _ = v40
        _ = v41
    }
    if ((!false) && false) {
        for fi46 := 0; fi46 < 1; fi46++ {
            for fi47 := 0; fi47 < 6; fi47++ {
                v2 -= (v2 | 169)
                _ = fi47
            }
            var v48 int32 = 42
            delete(v1, 832)
            _ = fi46
            _ = v48
        }
    } else {
        if false {
            var v49 int64 = (-(v4))
            _ = v49
        }
    }
    var v50 bool = false
    _ = len(v3)
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v37:", v37)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v50:", v50)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v5
    _ = v50
    _ = v6
}
