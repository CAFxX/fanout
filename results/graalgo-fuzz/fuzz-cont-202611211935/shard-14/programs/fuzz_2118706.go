package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 uint16 }
type GenS2 struct { F0 uint64; F1 int32; F2 uint8 }
type GenS3 struct { F0 int16; F1 int16; F2 uint8 }
type GenS4 struct { F0 int16 }

func main() {
    v1 := make(map[uint32]int)
    var v2 float64 = 3.14159
    v3 := make(map[string]bool)
    var v4 float32 = (float32(v2) / (0.1 / ((3.14159 / ((1.5 + (3.14159 / 1.0)) / 0.1)) / (0.1 / 0.1))))
    var v5 uint16 = 255
    switch 1 {
    case 32768:
        var v6 int16 = 42
        _ = v6
    case -1:
        switch 4294967295 {
        case 128:
            var v7 uint64 = 744
            _ = v7
        default:
            fmt.Println(v2, v5)
        }
    }
    var v8 float32 = (v4 / -2.25)
    if true {
        _ = len(v3)
    }
    switch "foo bar" {
    case "foo bar":
        if false {
            if ((^(v5)) != 919) {
                if false {
                    var v9 uint64 = 1
                    _ = v9
                }
                for fi10 := 0; fi10 < 6; fi10++ {
                    v11 := GenS1{F0: 911, F1: 9223372036854775807, F2: v5}
                    if true {
                        v8 = float32(float64(1e10))
                        switch fi10 {
                        case 9223372036854775807:
                            v12 := v3["foo bar"]
                            _ = v12
                        case 127:
                            v13 := GenS1{F0: 9223372036854775807, F1: 9223372036854775807, F2: (v5 & (v5 >> ((uint(v5)) & 63)))}
                            _ = v13
                        case 32768:
                            fmt.Println(fi10, v2, v2)
                        }
                        v4 -= (float32(v2) + v4)
                    }
                    var v14 uint64 = 1
                    _ = fi10
                    _ = v11
                    _ = v14
                }
                if (v5 != v5) {
                    var v15 uint32 = 2147483647
                    _ = v15
                }
            } else {
                var v16 string = "foo bar"
                _ = v16
            }
            var v17 bool = false
            _ = v17
        }
    }
    v2 = (v2 + float64(-2.25))
    for fi18 := 0; fi18 < 17; fi18++ {
        v2 = v2
        _ = fi18
    }
    for fi19 := 0; fi19 < 16; fi19++ {
        v20 := GenS2{F0: 961, F1: -2147483648, F2: 128}
        _ = fi19
        _ = v20
    }
    fmt.Println(v4, v5)
    var v21 bool = (true || false)
    if false {
        var v22 uint64 = 597
        var v23 int = 295
        fmt.Println(v5, v2, v8)
        _ = v22
        _ = v23
    }
    if false {
        v5 += (v5 / ((v5) | 1))
        v21 = (!v21)
        delete(v3, ("a" + "hello"))
    }
    for fi24 := 0; fi24 < 18; fi24++ {
        if ((255 <= 65535) && (!v21)) {
            var v25 float64 = v2
            var v26 bool = ((4294967295 > 42) || (!(0 != 0)))
            delete(v1, 2147483648)
            _ = v25
            _ = v26
        } else {
            for fi27 := 0; fi27 < 4; fi27++ {
                v28 := GenS1{F0: 0, F1: 1, F2: v5}
                _ = fi27
                _ = v28
            }
        }
        v29 := []int{fi24, fi24}
        var v30 uint8 = 42
        _ = fi24
        _ = v29
        _ = v30
    }
    if false {
        delete(v3, "\n\t")
        var v31 bool = v21
        _ = v31
    } else {
        if ((-370 != 1) || (!(1 < -1))) {
            for fi32 := 0; fi32 < 20; fi32++ {
                var v33 string = ((("foo bar" + ((("a" + "\n\t") + ("a" + "")) + "")) + ("\n\t" + "")) + "")
                fmt.Println(v4)
                delete(v1, 4294967295)
                _ = fi32
                _ = v33
            }
            if (v21 && true) {
                switch 603 {
                case -2147483648:
                    var v34 int16 = 388
                    _ = v34
                case 4294967295:
                    if false {
                        var v35 string = ""
                        _ = v35
                    } else {
                        v36 := GenS3{F0: 42, F1: 42, F2: 144}
                        _ = v36
                    }
                case -642:
                    if (int16(v5) != 1000) {
                        v5 = (v5 ^ v5)
                    }
                }
                v8 = float32(v2)
            } else {
                switch -538 {
                case 223:
                    v5 += v5
                case 127:
                    switch 127 {
                    case 127:
                        v21 = true
                    default:
                        v5 += v5
                    }
                case -252:
                    v37 := v3["a"]
                    _ = v37
                default:
                    switch 246 {
                    case 9223372036854775807:
                        var v38 bool = v21
                        _ = v38
                    }
                }
            }
        } else {
            fmt.Println(v2, v2)
            if v21 {
                v39 := GenS4{F0: 303}
                v21 = v21
                _ = v39
            } else {
                switch -147 {
                case -2147483648:
                    var v40 uint32 = 471
                    _ = v40
                default:
                    v41 := []int{1, -479}
                    _ = v41
                }
                for fi42 := 0; fi42 < 6; fi42++ {
                    if ("a" >= "") {
                        v43 := GenS1{F0: 42, F1: -1, F2: (^(v5))}
                        var v44 int32 = -1
                        _ = v43
                        _ = v44
                    }
                    v45 := []int{(fi42 << ((4294967295) & 63)), (fi42 % ((430) | 1))}
                    _ = fi42
                    _ = v45
                }
                var v46 int16 = -104
                _ = v46
            }
        }
        var v47 int32 = 42
        _ = v47
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
