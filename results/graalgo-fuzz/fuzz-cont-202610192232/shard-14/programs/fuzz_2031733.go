package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 int }
type GenS2 struct { F0 uint64; F1 uint16 }
type GenS3 struct { F0 uint64; F1 int16; F2 int16 }
type GenS4 struct { F0 string; F1 uint }

func main() {
    var v1 int8 = 1
    var v2 float32 = 0.0
    var v3 float64 = 1.5
    v4 := make(map[uint32]uint32)
    var v5 int8 = (v1 >> ((255) & 63))
    if false {
        v5 = (v1 - v1)
        switch 32767 {
        case 0:
            if (0 >= 0) {
                v6 := GenS1{F0: ("foo bar" + ""), F1: 917, F2: 32767}
                switch ("\n\t" + "foo bar") {
                case "a":
                    var v7 string = ("hello" + "\n\t")
                    _ = v7
                default:
                    if false {
                        var v8 float32 = v2
                        _ = v8
                    } else {
                        v9 := GenS1{F0: "hello", F1: 2, F2: -2147483648}
                        if true {
                            switch 128 {
                            case -981:
                                for fi10 := 0; fi10 < 12; fi10++ {
                                    if (18446744073709551615 != 1) {
                                        v11 := GenS1{F0: "a", F1: 0, F2: fi10}
                                        v12 := GenS2{F0: 4294967295, F1: 0}
                                        _ = v11
                                        _ = v12
                                    }
                                    v13 := GenS2{F0: 415, F1: 812}
                                    if true {
                                        var v14 float32 = (3.14159 / v2)
                                        for fi15 := 0; fi15 < 5; fi15++ {
                                            v16 := []int{(^(fi15)), fi15, -685}
                                            v17 := GenS3{F0: 2, F1: 0, F2: -32768}
                                            var v18 bool = false
                                            if (!(2 >= 579)) {
                                                var v19 int8 = v1
                                                _ = v19
                                            }
                                            _ = fi15
                                            _ = v16
                                            _ = v17
                                            _ = v18
                                        }
                                        _ = v14
                                    }
                                    _ = fi10
                                    _ = v13
                                }
                            default:
                                for fi20 := 0; fi20 < 16; fi20++ {
                                    var v21 int64 = -241
                                    v5 -= 127
                                    v21 = 158
                                    _ = fi20
                                    _ = v21
                                }
                            }
                            v22 := []int{560, 128, 32768, -9223372036854775808}
                            _ = v22
                        } else {
                            v2 = ((float32(v3) * float32(v3)) / (v2 - (v2 + float32(v3))))
                        }
                        delete(v4, 0)
                        _ = v9
                    }
                }
                var v23 float32 = 1e10
                _ = v23
                _ = v6
            } else {
                v2 = 0.1
                v2 = v2
            }
        case 4294967295:
            v24 := GenS1{F0: "", F1: 2147483647, F2: 32767}
            _ = v24
        case -2147483648:
            for fi25 := 0; fi25 < 19; fi25++ {
                switch fi25 {
                case 495:
                    for fi26 := 0; fi26 < 10; fi26++ {
                        if false {
                            v3 += v3
                        }
                        _ = fi26
                    }
                case 32767:
                    var v27 int64 = 0
                    _ = v27
                }
                v3 += ((float64(v2) - v3) + (v3 - (v3 - 0.0)))
                _ = fi25
            }
        default:
            if true {
                delete(v4, 91)
                fmt.Println(v3)
                for fi28 := 0; fi28 < 6; fi28++ {
                    v3 = (0.1 - v3)
                    if true {
                        v4[0] = 4294967295
                        if (3.14159 != float64(v2)) {
                            v5 = -56
                        }
                        delete(v4, 0)
                    } else {
                        v29 := GenS3{F0: 1, F1: 132, F2: 1}
                        _ = v29
                    }
                    if false {
                        v5 += 104
                        var v30 uint = 384
                        _ = v30
                    }
                    _ = fi28
                }
            }
        }
    } else {
        fmt.Println(v5)
    }
    v3 = v3
    var v31 string = ("\n\t" + ((("hello" + "\n\t") + ((("hello" + ("a" + "\n\t")) + ("" + ("\n\t" + "\n\t"))) + "")) + "hello"))
    switch ("a" + v31) {
    case "a":
        v32 := v4[4294967295]
        _ = v32
    case "hello":
        var v33 uint8 = 2
        _ = v33
    }
    v5 -= v5
    for fi34 := 0; fi34 < 5; fi34++ {
        if (v1 == (v1 % ((v1) | 1))) {
            var v35 uint64 = 733
            v31 = v31
            _ = v35
        }
        var v36 uint16 = 767
        var v37 uint64 = 417
        _ = fi34
        _ = v36
        _ = v37
    }
    if (32 < -748) {
        v38 := GenS3{F0: 18446744073709551615, F1: 42, F2: 42}
        v2 = v2
        _ = v38
    } else {
        v4[557] = 2147483647
        v2 = 3.14159
        var v39 uint = 18446744073709551615
        _ = v39
    }
    v40 := []int{}
    delete(v4, 65535)
    v31 = v31
    v41 := GenS4{F0: v31, F1: 65535}
    fmt.Println(v5, v31)
    v5 += (v1 << ((65535) & 63))
    var v42 float32 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v31:", v31)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v5
}
