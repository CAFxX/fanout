package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16; F1 uint }
type GenS4 struct { F0 uint32; F1 int16; F2 uint16 }
type GenS5 struct { F0 uint8 }

func main() {
    v1 := make(map[int]uint32)
    var v2 float64 = -2.25
    var v3 int8 = 100
    fmt.Println(v3, v3)
    if false {
        if false {
            v4 := GenS1{F0: 2}
            _ = v4
        }
    }
    delete(v1, -2147483648)
    var v5 bool = (!true)
    if v5 {
        var v6 int = 1
        switch ("a" + "hello") {
        case "hello":
            for fi7 := 0; fi7 < 13; fi7++ {
                if v5 {
                    if false {
                        v8 := []int{(v6 - (fi7 / ((fi7) | 1))), fi7}
                        fmt.Println(v2)
                        for fi9 := 0; fi9 < 2; fi9++ {
                            fmt.Println(v6)
                            _ = fi9
                        }
                        _ = v8
                    }
                    if ((127 <= 42) && v5) {
                        if v5 {
                            fmt.Println(v2, v3, v3)
                            v3 = 100
                        } else {
                            switch 1 {
                            case 128:
                                var v10 float32 = (0.1 / (float32(-2.25) / (0.1 / 1.0)))
                                _ = v10
                            case 2147483647:
                                var v11 uint8 = 195
                                _ = v11
                            case -1:
                                _ = len(v1)
                            }
                            fmt.Println(fi7)
                            if (!v5) {
                                var v12 int64 = -1
                                switch v6 {
                                case 128:
                                    var v13 uint16 = 32767
                                    _ = v13
                                default:
                                    var v14 string = "hello"
                                    _ = v14
                                }
                                _ = v12
                            }
                        }
                        var v15 uint8 = 128
                        _ = v15
                    } else {
                        v16 := GenS2{F0: 0}
                        _ = v16
                    }
                }
                v17 := GenS1{F0: 1}
                var v18 float32 = 1e10
                _ = fi7
                _ = v17
                _ = v18
            }
        default:
            for fi19 := 0; fi19 < 18; fi19++ {
                v20 := GenS2{F0: -1}
                v21 := GenS3{F0: 487, F1: 1}
                _ = fi19
                _ = v20
                _ = v21
            }
        }
        _ = v6
    } else {
        if true {
            v22 := GenS4{F0: 0, F1: -686, F2: 2}
            if v5 {
                var v23 uint64 = 1
                _ = v23
            }
            _ = v22
        }
    }
    for fi24 := 0; fi24 < 13; fi24++ {
        if false {
            switch (((("hello" + ("foo bar" + "")) + ("foo bar" + ("a" + "a"))) + "foo bar") + "hello") {
            case "foo bar":
                v3 = (^(v3))
            case "":
                for fi25 := 0; fi25 < 20; fi25++ {
                    switch 4294967295 {
                    case 32768:
                        var v26 int64 = -757
                        _ = v26
                    case 42:
                        v27 := []int{fi24}
                        _ = v27
                    case -721:
                        var v28 string = "foo bar"
                        _ = v28
                    default:
                        v29 := GenS5{F0: 128}
                        _ = v29
                    }
                    _ = fi25
                }
            }
            var v30 bool = v5
            _ = v30
        } else {
            v31 := []int{(-(fi24)), fi24, fi24}
            var v32 uint16 = 255
            v5 = true
            _ = v31
            _ = v32
        }
        if v5 {
            v3 += v3
            for fi33 := 0; fi33 < 20; fi33++ {
                var v34 int = (fi24 - -1)
                if v5 {
                    var v35 float32 = 0.1
                    var v36 int8 = v3
                    _ = v35
                    _ = v36
                }
                _ = fi33
                _ = v34
            }
            v3 -= 1
        }
        var v37 uint8 = 255
        _ = fi24
        _ = v37
    }
    fmt.Println(v2, v2)
    v3 = 42
    v3 += v3
    v3 -= (v3 * v3)
    v3 -= v3
    var v38 int16 = -117
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v5
}
