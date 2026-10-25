package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool; F1 int; F2 int64 }
type GenS3 struct { F0 uint; F1 string }
func genFn0(p0 int32) int64 {
    v4 := []int{128, 42, 32767, -9223372036854775808}
    switch "a" {
    case "hello":
        var v5 int64 = -1
        _ = v5
    default:
        for fi6 := 0; fi6 < 9; fi6++ {
            for fi7 := 0; fi7 < 5; fi7++ {
                p0 += p0
                p0 -= 2147483647
                v4 = append(v4, fi6)
                _ = fi7
            }
            p0 = (p0 >> ((887) & 63))
            switch fi6 {
            case 1:
                if true {
                    if len(v4) > 0 {
                        v4[0] = fi6
                    }
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    switch 0 {
                    case 32767:
                        p0 = (p0 & (p0 ^ (^(p0))))
                    default:
                        var v8 int16 = 32767
                        _ = v8
                    }
                } else {
                    p0 = p0
                    v9 := GenS1{F0: 2147483648}
                    var v10 string = ""
                    _ = v10
                    _ = v9
                }
            case 885:
                v11 := GenS1{F0: 313}
                _ = v11
            case 9223372036854775807:
                for fi12 := 0; fi12 < 6; fi12++ {
                    if true {
                        v13 := GenS1{F0: 2147483648}
                        var v14 uint32 = 908
                        switch ("" + "\n\t") {
                        case "":
                            switch "" {
                            case "":
                                var v15 uint16 = 0
                                _ = v15
                            default:
                                var v16 uint16 = 2
                                _ = v16
                            }
                        default:
                            var v17 uint = 2
                            _ = v17
                        }
                        if (!true) {
                            if (!false) {
                                v18 := GenS2{F0: (!(fi6 >= (^(fi12)))), F1: fi6, F2: 1}
                                var v19 int64 = -1
                                _ = v18
                                _ = v19
                            }
                        } else {
                            p0 += -2147483648
                        }
                        var v20 int32 = (p0 << ((2) & 63))
                        _ = v13
                        _ = v14
                        _ = v20
                    } else {
                        if false {
                            v21 := GenS1{F0: 1}
                            v22 := GenS3{F0: 2, F1: ""}
                            _ = v21
                            _ = v22
                        }
                    }
                    var v23 uint8 = 128
                    v23 = (v23 & 234)
                    var v24 int = (fi6 / (((fi12 % ((fi12) | 1))) | 1))
                    _ = fi12
                    _ = v23
                    _ = v24
                }
            }
            _ = fi6
        }
    }
    for fi25 := 0; fi25 < 3; fi25++ {
        var v26 uint64 = 0
        _ = fi25
        _ = v26
    }
    _ = p0
    _ = v4
    return 531
}

func main() {
    var v1 uint16 = 555
    v2 := make(map[uint32]uint32)
    var v3 float32 = 0.1
    v3 = v3
    for fi27 := 0; fi27 < 6; fi27++ {
        var v28 int32 = 42
        _ = fi27
        _ = v28
    }
    for fi29 := 0; fi29 < 8; fi29++ {
        fmt.Println(v1, fi29)
        _ = fi29
    }
    var v30 uint = 2
    v1 = v1
    for fi31 := 0; fi31 < 18; fi31++ {
        fmt.Println(v30)
        var v32 bool = (true || false)
        _ = fi31
        _ = v32
    }
    for fi33 := 0; fi33 < 19; fi33++ {
        v1 = 412
        _ = fi33
    }
    fmt.Println(v3, v3)
    fmt.Println(v30, v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
}
