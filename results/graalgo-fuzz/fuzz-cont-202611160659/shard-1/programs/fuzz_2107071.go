package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 uint32 }
type GenS3 struct { F0 uint32; F1 int8 }
type GenS4 struct { F0 uint16; F1 uint32; F2 int16 }
type GenS5 struct { F0 uint; F1 int64 }
type GenS6 struct { F0 int; F1 int8 }
func genFn0() uint8 {
    v7 := GenS1{F0: 0}
    if (-2147483648 > 0) {
        if true {
            var v8 uint16 = 0
            _ = v8
        } else {
            v9 := []int{-604}
            _ = v9
        }
        var v10 bool = (-696 != -9223372036854775808)
        _ = v10
    } else {
        var v11 int8 = -6
        for fi12 := 0; fi12 < 19; fi12++ {
            v11 += v11
            _ = fi12
        }
        _ = v11
    }
    for fi13 := 0; fi13 < 20; fi13++ {
        v14 := GenS1{F0: 42}
        v15 := GenS1{F0: 247}
        if true {
            for fi16 := 0; fi16 < 19; fi16++ {
                var v17 int = (fi13 * -274)
                var v18 int = fi16
                _ = fi16
                _ = v17
                _ = v18
            }
        }
        _ = fi13
        _ = v14
        _ = v15
    }
    v19 := GenS2{F0: 128, F1: 985}
    switch 42 {
    case -9223372036854775808:
        switch -12 {
        case -9223372036854775808:
            var v20 int8 = 1
            _ = v20
        case 9223372036854775807:
            v21 := []int{}
            _ = v21
        default:
            v22 := []int{128, 42, 86, 32767}
            _ = v22
        }
    case 32768:
        var v23 int8 = 127
        _ = v23
    }
    _ = v19
    _ = v7
    return 231
}
func genFn1(p0 int8) string {
    var v24 uint8 = genFn0()
    fmt.Println(v24, v24)
    _ = p0
    _ = v24
    return "\n\t"
}

func main() {
    var v1 string = "\n\t"
    v2 := make(map[int]bool)
    v3 := make(map[uint32]uint32)
    var v4 bool = (-540 < 2147483647)
    v5 := []uint32{}
    var v6 int16 = -1
    if v4 {
        v4 = (!false)
        if v4 {
            v6 = (-(v6))
            v4 = true
            if v4 {
                switch "foo bar" {
                case "hello":
                    fmt.Println(v4, v1)
                case "\n\t":
                    _ = len(v5)
                default:
                    v4 = v4
                }
                for fi25 := 0; fi25 < 2; fi25++ {
                    v4 = (!((v6 != (v6 >> ((634) & 63))) && false))
                    _ = fi25
                }
            }
        } else {
            v6 -= (^(v6))
            var v26 float64 = (3.14159 - 1.5)
            _ = v26
        }
    } else {
        for fi27 := 0; fi27 < 16; fi27++ {
            var v28 uint64 = 2
            fmt.Println(fi27, v6)
            _ = fi27
            _ = v28
        }
        if false {
            if (-2.25 > (((1e10 + 0.1) - 3.14159) + ((1.5 * 1.5) + (0.0 * (0.0 + (0.0 + (0.1 + (float64(1.5) - 0.1)))))))) {
                v29 := GenS1{F0: 2}
                v2[-539] = v4
                _ = v29
            }
            fmt.Println(v6, v6)
        }
        if ((2 != 42) && (genFn0() > 1)) {
            v6 = (v6 * v6)
            switch ((("hello" + ("a" + v1)) + ("\n\t" + ("hello" + (v1 + v1)))) + "") {
            case "hello":
                if v4 {
                    fmt.Println(v1, v1)
                    if true {
                        fmt.Println(v1)
                        switch 4294967295 {
                        case 714:
                            for fi30 := 0; fi30 < 15; fi30++ {
                                v4 = false
                                var v31 uint = 65535
                                _ = fi30
                                _ = v31
                            }
                        case 264:
                            v32 := GenS3{F0: 65535, F1: 1}
                            _ = v32
                        case 696:
                            var v33 string = ("" + v1)
                            _ = v33
                        default:
                            _ = len(v5)
                        }
                        for fi34 := 0; fi34 < 2; fi34++ {
                            v1 = (v1 + (v1 + "foo bar"))
                            v35 := GenS4{F0: 0, F1: 20, F2: 42}
                            _ = fi34
                            _ = v35
                        }
                    }
                } else {
                    v1 = (v1 + "")
                }
            case "\n\t":
                fmt.Println(v4)
            case "":
                if v4 {
                    v1 = "foo bar"
                    v36 := GenS5{F0: 2, F1: -134}
                    v37 := GenS3{F0: 42, F1: 1}
                    _ = v36
                    _ = v37
                } else {
                    v38 := GenS6{F0: 128, F1: 0}
                    if (!false) {
                        v1 = v1
                        fmt.Println(v4, v1)
                        for fi39 := 0; fi39 < 7; fi39++ {
                            v40 := GenS3{F0: 473, F1: -1}
                            var v41 int16 = v6
                            _ = fi39
                            _ = v40
                            _ = v41
                        }
                    }
                    switch 0 {
                    case 2147483647:
                        var v42 float32 = ((0.1 / 0.1) / (float32(-2.25) + ((((float32(float64((((0.0 + 0.0) - 3.14159) - ((1e10 / 0.1) + (1e10 - (1e10 - 1.5)))))) / (1.5 / (1.5 - 0.0))) / 0.1) - (float32(0.1) + ((0.0 - float32(0.1)) + 3.14159))) * ((0.1 * 1e10) * float32((0.0 - 0.1))))))
                        _ = v42
                    case 32768:
                        fmt.Println(v4, v4, v1)
                    default:
                        var v43 int = -1
                        _ = v43
                    }
                    _ = v38
                }
            }
            v4 = (323 > 2)
        }
    }
    var v44 int16 = (v6 | 42)
    var v45 string = v1
    v4 = v4
    v4 = (genFn0() > 117)
    for fi46 := 0; fi46 < 3; fi46++ {
        var v47 float32 = 1e10
        _ = fi46
        _ = v47
    }
    if v4 {
        for fi48 := 0; fi48 < 17; fi48++ {
            v6 -= -32768
            for fi49 := 0; fi49 < 19; fi49++ {
                v45 = v45
                var v50 int16 = 32767
                _ = fi49
                _ = v50
            }
            for fi51 := 0; fi51 < 14; fi51++ {
                fmt.Println(fi48, v45, v1)
                var v52 uint8 = genFn0()
                v2[(fi48 << ((1) & 63))] = v4
                _ = fi51
                _ = v52
            }
            _ = fi48
        }
    } else {
        v4 = (408 > 4294967295)
        switch v1 {
        case "hello":
            var v53 string = v45
            _ = v53
        }
        v1 = (v45 + "foo bar")
    }
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v45
    _ = v5
    _ = v6
}
