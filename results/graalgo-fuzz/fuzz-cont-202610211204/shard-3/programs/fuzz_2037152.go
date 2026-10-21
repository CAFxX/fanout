package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 bool }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]bool)
    v3 := make(map[uint32]uint32)
    v4 := make(map[string]bool)
    var v5 uint = 65535
    var v6 string = "a"
    switch v6 {
    case "hello":
        var v7 int32 = 42
        _ = v7
    case "\n\t":
        var v8 float64 = 0.0
        _ = v8
    }
    for fi9 := 0; fi9 < 1; fi9++ {
        if true {
            v10 := []int{}
            v5 = v5
            _ = v10
        }
        if true {
            v11 := []int{(-(fi9)), fi9, (fi9 % ((0) | 1))}
            _ = v11
        }
        _ = fi9
    }
    v6 = v6
    var v12 int16 = -32768
    fmt.Println(v12, v12, v6)
    var v13 uint64 = 887
    v5 -= (v5 | (v5 | 255))
    v6 = ((v6 + (v6 + "foo bar")) + ("foo bar" + v6))
    fmt.Println(v6)
    v14 := GenS1{F0: -487, F1: v6, F2: true}
    v15 := v1[-9223372036854775808]
    if false {
        delete(v3, 293)
    } else {
        v16 := GenS1{F0: 959, F1: v6, F2: v15}
        if v15 {
            var v17 uint64 = (v13 / ((v13) | 1))
            var v18 float64 = 0.0
            for fi19 := 0; fi19 < 17; fi19++ {
                if (uint32(v13) > 42) {
                    v20 := GenS1{F0: 1, F1: "hello", F2: (((("hello" + v6) + (("" + "") + (v6 + ""))) != "a") && ((!true) && v15))}
                    switch "foo bar" {
                    case "hello":
                        var v21 uint64 = uint64(v5)
                        _ = v21
                    case "a":
                        var v22 float64 = 0.0
                        _ = v22
                    default:
                        v23 := []int{fi19}
                        _ = v23
                    }
                    _ = v20
                } else {
                    for fi24 := 0; fi24 < 4; fi24++ {
                        v25 := GenS1{F0: 1, F1: (v6 + ((v6 + (v6 + ("hello" + ""))) + v6)), F2: true}
                        v5 += (v5 | 431)
                        if (false && (42 > uint32(v5))) {
                            v26 := []int{fi19, fi19, 1}
                            var v27 uint16 = 255
                            fmt.Println(v12)
                            _ = v26
                            _ = v27
                        } else {
                            if false {
                                var v28 uint32 = 938
                                _ = v28
                            } else {
                                var v29 int64 = -804
                                var v30 uint8 = 127
                                v31 := GenS1{F0: -2147483648, F1: "a", F2: (42 < uint8(v17))}
                                _ = v29
                                _ = v30
                                _ = v31
                            }
                        }
                        _ = fi24
                        _ = v25
                    }
                    v32 := GenS1{F0: 0, F1: v6, F2: (!false)}
                    for fi33 := 0; fi33 < 9; fi33++ {
                        var v34 float32 = -2.25
                        v35 := v2["a"]
                        var v36 float64 = v18
                        _ = fi33
                        _ = v34
                        _ = v35
                        _ = v36
                    }
                    _ = v32
                }
                _ = fi19
            }
            _ = v17
            _ = v18
        } else {
            v13 = (v13 | (v13 << ((18446744073709551615) & 63)))
            for fi37 := 0; fi37 < 13; fi37++ {
                var v38 int64 = 9223372036854775807
                for fi39 := 0; fi39 < 1; fi39++ {
                    var v40 bool = (42 <= 192)
                    _ = fi39
                    _ = v40
                }
                var v41 float64 = float64((1.5 + (0.0 * 1.5)))
                _ = fi37
                _ = v38
                _ = v41
            }
        }
        for fi42 := 0; fi42 < 1; fi42++ {
            var v43 uint64 = v13
            var v44 int16 = 42
            switch v6 {
            case "\n\t":
                v45 := GenS1{F0: -2147483648, F1: "foo bar", F2: (int32(fi42) <= -698)}
                _ = v45
            case "foo bar":
                switch fi42 {
                case 32767:
                    var v46 bool = (-1 <= 4)
                    _ = v46
                default:
                    v43 = 66
                }
            case "hello":
                var v47 float64 = float64(0.1)
                _ = v47
            default:
                var v48 string = (v6 + "")
                _ = v48
            }
            _ = fi42
            _ = v43
            _ = v44
        }
        _ = v16
    }
    var v49 uint16 = 32767
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v49
    _ = v5
    _ = v6
}
