package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string }
func genFn0(p0 bool, p1 uint64) uint16 {
    var v7 int = 32768
    for fi8 := 0; fi8 < 20; fi8++ {
        v9 := []int{}
        var v10 uint64 = p1
        _ = fi8
        _ = v10
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v7
    return 210
}
func genFn1() int8 {
    if true {
        v11 := []int{127, 128, -9223372036854775808}
        _ = v11
    }
    v12 := []int{-9223372036854775808, -114}
    for fi13 := 0; fi13 < 16; fi13++ {
        switch fi13 {
        case 4294967295:
            for fi14 := 0; fi14 < 8; fi14++ {
                var v15 int8 = 42
                v15 += v15
                v15 -= 42
                _ = fi14
                _ = v15
            }
        case -1:
            var v16 int64 = 42
            _ = v16
        }
        _ = fi13
    }
    for fi17 := 0; fi17 < 12; fi17++ {
        var v18 int8 = 127
        _ = fi17
        _ = v18
    }
    _ = v12
    return 77
}

func main() {
    v1 := make(map[int]uint32)
    var v2 int8 = 42
    var v3 bool = true
    v4 := make(map[int]int)
    var v5 int8 = (v2 >> ((767) & 63))
    var v6 uint = 255
    fmt.Println(v6, v2, v3)
    v2 = (v5 + genFn1())
    switch "" {
    case "a":
        v19 := GenS1{F0: -662}
        _ = v19
    case "":
        if true {
            switch 180 {
            case -9223372036854775808:
                v20 := []int{}
                _ = v20
            case 133:
                if ((!true) && (float64(((-2.25 - float32(-2.25)) / (0.0 + (3.14159 - (0.0 - 0.0))))) > 1e10)) {
                    v21 := []int{-1, 971, 42, -215}
                    if ((1 != int32(v2)) || v3) {
                        var v22 bool = false
                        for fi23 := 0; fi23 < 19; fi23++ {
                            switch "hello" {
                            case "\n\t":
                                v24 := GenS1{F0: 42}
                                _ = v24
                            case "foo bar":
                                v25 := GenS2{F0: (("\n\t" + "foo bar") + "\n\t")}
                                _ = v25
                            }
                            var v26 int16 = -329
                            var v27 uint32 = 65535
                            _ = fi23
                            _ = v26
                            _ = v27
                        }
                        _ = v22
                    }
                    _ = v21
                } else {
                    var v28 uint64 = 683
                    v29 := GenS1{F0: 42}
                    for fi30 := 0; fi30 < 14; fi30++ {
                        for fi31 := 0; fi31 < 7; fi31++ {
                            v2 = v5
                            _ = fi31
                        }
                        for fi32 := 0; fi32 < 12; fi32++ {
                            v28 += 4294967295
                            var v33 int64 = 4
                            fmt.Println(fi30, v3)
                            _ = fi32
                            _ = v33
                        }
                        var v34 uint64 = v28
                        _ = fi30
                        _ = v34
                    }
                    _ = v28
                    _ = v29
                }
            case 236:
                if false {
                    v35 := []int{32767, -466, 843, -1}
                    for fi36 := 0; fi36 < 10; fi36++ {
                        v6 -= (v6 / (((v6 & (v6 * v6))) | 1))
                        _ = fi36
                    }
                    if len(v35) > 0 {
                        v35[0] = -9223372036854775808
                    }
                    _ = v35
                }
            }
            fmt.Println(v2)
        } else {
            v4[-9223372036854775808] = 32767
        }
    default:
        v3 = (("hello" + "foo bar") == "\n\t")
    }
    v37 := GenS2{F0: "foo bar"}
    var v38 int = 9223372036854775807
    v2 += v2
    var v39 uint16 = uint16(v6)
    var v40 int32 = -2147483648
    var v41 uint8 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v39:", v39)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v5
    _ = v6
}
