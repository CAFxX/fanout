package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 int64; F1 int8 }
type GenS3 struct { F0 uint32; F1 int8; F2 int16 }
type GenS4 struct { F0 int16; F1 uint32 }
type GenS5 struct { F0 uint32; F1 uint; F2 int32 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint64 }

func main() {
    v1 := make(map[string]bool)
    v2 := make(map[int]uint32)
    v3 := make(map[int]uint32)
    var v4 uint16 = 65535
    var v5 uint16 = v4
    for fi6 := 0; fi6 < 7; fi6++ {
        for fi7 := 0; fi7 < 8; fi7++ {
            if true {
                var v8 uint8 = 0
                v8 -= v8
                fmt.Println(v8)
                _ = v8
            } else {
                var v9 float32 = ((1e10 - float32(1.5)) + (1.5 * 1.5))
                delete(v1, (("a" + "\n\t") + ("foo bar" + ((("hello" + "hello") + "") + ("hello" + "")))))
                _ = v9
            }
            fmt.Println(fi7)
            fmt.Println(v5, fi6)
            var v10 bool = true
            v10 = false
            fmt.Println(v5)
            _ = fi7
            _ = v10
        }
        var v11 float64 = ((float64((float32(((1.5 * 0.1) * 0.1)) / ((((-2.25 * ((-2.25 * 1.5) - (((-2.25 - 0.0) / (3.14159 / 1.5)) + 1.5))) * 0.1) + 1.5) * -2.25))) - (float64(float32(float64(0.0))) - -2.25)) / 0.1)
        _ = fi6
        _ = v11
    }
    for fi12 := 0; fi12 < 18; fi12++ {
        for fi13 := 0; fi13 < 18; fi13++ {
            v4 += v4
            switch ("foo bar" + "\n\t") {
            case "\n\t":
                v14 := GenS1{F0: 127, F1: 856}
                _ = v14
            case "a":
                v4 = v5
            case "hello":
                fmt.Println(v4, fi13, v5)
            default:
                _ = len(v2)
            }
            _ = fi13
        }
        _ = fi12
    }
    for fi15 := 0; fi15 < 15; fi15++ {
        var v16 int16 = 0
        var v17 int32 = 0
        if (1 >= -110) {
            v18 := []int{int(v16), -9223372036854775808, fi15, (fi15 << ((210) & 63))}
            _ = v18
        } else {
            if false {
                var v19 int16 = v16
                _ = v19
            }
            delete(v3, -847)
        }
        _ = fi15
        _ = v16
        _ = v17
    }
    _ = len(v2)
    v5 = (v4 << ((42) & 63))
    if true {
        v20 := GenS2{F0: 812, F1: 1}
        _ = v20
    }
    if (267 != 2) {
        for fi21 := 0; fi21 < 12; fi21++ {
            for fi22 := 0; fi22 < 13; fi22++ {
                v5 += (v4 * 255)
                v4 = 2
                v23 := GenS3{F0: 2147483648, F1: 0, F2: 0}
                fmt.Println(v4, fi22, v5)
                if true {
                    if ((!true) && false) {
                        v5 -= 42
                        var v24 uint = 18446744073709551615
                        _ = v24
                    }
                    var v25 uint32 = 4294967295
                    if true {
                        fmt.Println(fi21)
                    } else {
                        if false {
                            if true {
                                if false {
                                    v2[9223372036854775807] = v25
                                    v5 = v5
                                } else {
                                    v26 := GenS3{F0: 4294967295, F1: -1, F2: -32768}
                                    _ = v26
                                }
                            }
                            v27 := GenS3{F0: (v25 | (v25 / ((v25) | 1))), F1: -1, F2: 393}
                            _ = v27
                        }
                    }
                    _ = v25
                }
                _ = fi22
                _ = v23
            }
            _ = fi21
        }
    }
    delete(v1, "")
    if true {
        for fi28 := 0; fi28 < 7; fi28++ {
            v29 := v2[fi28]
            if false {
                for fi30 := 0; fi30 < 12; fi30++ {
                    v31 := GenS4{F0: 1, F1: v29}
                    var v32 uint = 606
                    fmt.Println(v29, v5)
                    _ = fi30
                    _ = v31
                    _ = v32
                }
                v33 := []int{fi28, (fi28 & fi28)}
                _ = v33
            }
            var v34 int64 = 530
            _ = fi28
            _ = v29
            _ = v34
        }
        var v35 string = "\n\t"
        for fi36 := 0; fi36 < 9; fi36++ {
            v4 = v4
            if true {
                switch v35 {
                case "\n\t":
                    v5 -= 0
                case "":
                    v37 := GenS1{F0: 101, F1: uint32(v4)}
                    _ = v37
                case "foo bar":
                    v35 = ("foo bar" + (v35 + (v35 + "a")))
                }
                var v38 int16 = 1000
                _ = v38
            } else {
                switch (fi36 - fi36) {
                case -1:
                    var v39 bool = true
                    _ = v39
                case -289:
                    var v40 float32 = float32(1e10)
                    _ = v40
                case 4294967295:
                    fmt.Println(v5)
                }
                var v41 string = v35
                _ = v41
            }
            v4 = 565
            _ = fi36
        }
        _ = v35
    } else {
        if (float32(-2.25) > float32((0.1 / 3.14159))) {
            v5 -= v4
            v4 = 682
            delete(v2, 127)
        }
        switch 32768 {
        case 54:
            var v42 bool = true
            _ = v42
        case 1:
            for fi43 := 0; fi43 < 10; fi43++ {
                switch (-(fi43)) {
                case 4294967295:
                    v4 -= v5
                default:
                    if true {
                        v44 := GenS5{F0: 2, F1: 2, F2: -2147483648}
                        _ = v44
                    }
                }
                _ = fi43
            }
        case 0:
            switch 346 {
            case -947:
                v45 := GenS6{F0: v4}
                _ = v45
            case -2147483648:
                for fi46 := 0; fi46 < 7; fi46++ {
                    var v47 int8 = 100
                    var v48 string = "hello"
                    _ = fi46
                    _ = v47
                    _ = v48
                }
            case 128:
                if (uint8(v4) >= 1) {
                    if false {
                        v49 := GenS6{F0: 0}
                        _ = v49
                    } else {
                        if false {
                            if (-2147483648 == -298) {
                                switch ("a" + ("a" + "foo bar")) {
                                case "":
                                    var v50 bool = (!true)
                                    _ = v50
                                }
                                if false {
                                    v5 = (v4 | v5)
                                }
                            }
                            fmt.Println(v4, v4)
                        }
                        var v51 uint32 = 2147483648
                        v52 := GenS7{F0: 115}
                        _ = v51
                        _ = v52
                    }
                    v4 = v4
                    switch 127 {
                    case -414:
                        fmt.Println(v4, v4)
                    case 32767:
                        switch 32768 {
                        case -391:
                            fmt.Println(v5, v4)
                        case 9223372036854775807:
                            var v53 float32 = (3.14159 * 1.5)
                            _ = v53
                        }
                    default:
                        var v54 uint = 255
                        _ = v54
                    }
                } else {
                    v5 = (v5 | v5)
                    _ = len(v2)
                }
            }
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
