package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 string; F1 uint }
type GenS3 struct { F0 uint; F1 int16 }
type GenS4 struct { F0 uint16 }

func main() {
    v1 := []string{}
    v2 := []string{"foo bar", "a"}
    v3 := make(map[uint32]bool)
    var v4 int16 = -451
    v5 := make(map[string]uint32)
    var v6 float32 = ((-2.25 + 3.14159) + 1.5)
    for fi7 := 0; fi7 < 10; fi7++ {
        v8 := GenS1{F0: 969}
        v3[0] = true
        switch 316 {
        case 4294967295:
            if true {
                var v9 uint32 = 0
                _ = v9
            } else {
                fmt.Println(fi7, fi7)
            }
        case -2147483648:
            var v10 string = ("foo bar" + ("foo bar" + "\n\t"))
            _ = v10
        }
        _ = fi7
        _ = v8
    }
    if true {
        v11 := GenS1{F0: 0}
        _ = v11
    } else {
        if (487 > -677) {
            switch -2147483648 {
            case 4294967295:
                if (!false) {
                    switch 1 {
                    case 385:
                        _ = len(v2)
                    case 32768:
                        v4 = v4
                    }
                    var v12 uint32 = 2
                    _ = v12
                }
            case 1:
                v13 := GenS2{F0: (("\n\t" + "a") + ("a" + "hello")), F1: 2}
                _ = v13
            }
            v1 = append(v1, "\n\t")
            var v14 uint32 = 865
            _ = v14
        } else {
            v3[2147483648] = (32767 < 891)
            var v15 uint8 = 127
            _ = v15
        }
    }
    var v16 uint32 = 2
    var v17 uint32 = v16
    fmt.Println(v6, v16, v6)
    for fi18 := 0; fi18 < 14; fi18++ {
        _ = len(v2)
        v19 := GenS3{F0: 42, F1: 668}
        var v20 float32 = v6
        _ = fi18
        _ = v19
        _ = v20
    }
    if true {
        var v21 int64 = 42
        delete(v5, ("" + ("" + ("hello" + "hello"))))
        var v22 bool = false
        _ = v21
        _ = v22
    }
    fmt.Println(v6, v16, v16)
    for fi23 := 0; fi23 < 6; fi23++ {
        v4 = (v4 * (v4 | v4))
        _ = fi23
    }
    v5["\n\t"] = v17
    if false {
        switch 42 {
        case -352:
            for fi24 := 0; fi24 < 5; fi24++ {
                if false {
                    fmt.Println(v4, v16, v16)
                } else {
                    if (!(!(128 == 4294967295))) {
                        if (260 >= 800) {
                            v25 := GenS4{F0: 272}
                            v26 := GenS2{F0: "", F1: 1}
                            _ = v25
                            _ = v26
                        }
                        for fi27 := 0; fi27 < 14; fi27++ {
                            v4 = -394
                            var v28 uint = 42
                            _ = fi27
                            _ = v28
                        }
                        v17 = 2147483648
                    }
                }
                var v29 uint64 = 2
                for fi30 := 0; fi30 < 8; fi30++ {
                    fmt.Println(fi30, v17, v16)
                    _ = fi30
                }
                _ = fi24
                _ = v29
            }
        default:
            fmt.Println(v6, v16, v6)
        }
        for fi31 := 0; fi31 < 14; fi31++ {
            v6 += v6
            v16 += 2147483648
            fmt.Println(v17, v17)
            _ = fi31
        }
        v4 -= v4
    } else {
        for fi32 := 0; fi32 < 9; fi32++ {
            for fi33 := 0; fi33 < 1; fi33++ {
                v4 = (v4 * v4)
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                _ = fi33
            }
            _ = fi32
        }
    }
    if true {
        if (false && ((1 <= 2) && (!(42 == 32768)))) {
            _ = len(v3)
        }
    }
    for fi34 := 0; fi34 < 16; fi34++ {
        fmt.Println(v16, v6)
        _ = fi34
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
