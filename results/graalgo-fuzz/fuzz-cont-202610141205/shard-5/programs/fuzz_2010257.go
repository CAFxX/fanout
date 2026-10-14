package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int8 }

func main() {
    v1 := make(map[int]bool)
    var v2 string = "\n\t"
    var v3 float32 = 1e10
    var v4 float64 = 1e10
    v5 := make(map[int]bool)
    var v6 int64 = 0
    switch 720 {
    case 343:
        fmt.Println(v6)
    }
    v7 := GenS1{F0: 211}
    var v8 uint = 65535
    v9 := []int{425, 4294967295, 219, 128}
    _ = len(v5)
    for fi10 := 0; fi10 < 7; fi10++ {
        for fi11 := 0; fi11 < 2; fi11++ {
            var v12 float32 = (1e10 + 0.0)
            _ = fi11
            _ = v12
        }
        v8 -= 570
        _ = fi10
    }
    var v13 int16 = 1000
    if false {
        if true {
            v2 = "hello"
        }
        v9 = append(v9, 103)
    }
    if true {
        var v14 float64 = float64(0.0)
        _ = v14
    } else {
        switch 2147483647 {
        case 1:
            var v15 string = (v2 + "a")
            _ = v15
        case 493:
            if len(v9) > 0 {
                v9[0] = 127
            }
        default:
            v6 = (v6 - (v6 << ((42) & 63)))
        }
    }
    switch 32767 {
    case -9223372036854775808:
        for fi16 := 0; fi16 < 18; fi16++ {
            for fi17 := 0; fi17 < 12; fi17++ {
                fmt.Println(v3, fi17, v3)
                _ = fi17
            }
            v13 -= (v13 - v13)
            v3 += v3
            _ = fi16
        }
    case 229:
        for fi18 := 0; fi18 < 20; fi18++ {
            var v19 bool = (true || (false || ((!(1e10 <= v4)) && (("a" + "hello") < "a"))))
            _ = fi18
            _ = v19
        }
    }
    var v20 uint16 = uint16(v8)
    var v21 uint64 = 1
    v21 = v21
    if false {
        v22 := GenS1{F0: (v20 * 42)}
        for fi23 := 0; fi23 < 4; fi23++ {
            fmt.Println(v8, v6, fi23)
            var v24 uint = (v8 * (v8 / ((255) | 1)))
            _ = fi23
            _ = v24
        }
        v25 := GenS2{F0: 42}
        _ = v22
        _ = v25
    }
    if true {
        switch 42 {
        case 128:
            v8 -= v8
        default:
            for fi26 := 0; fi26 < 18; fi26++ {
                fmt.Println(v6)
                v27 := GenS3{F0: 42}
                if false {
                    var v28 int16 = (v13 >> ((v8) & 63))
                    var v29 int64 = v6
                    _ = v28
                    _ = v29
                } else {
                    var v30 uint32 = 65535
                    if false {
                        switch int(v8) {
                        case 370:
                            switch fi26 {
                            case 128:
                                v30 = 68
                            case 1:
                                v13 -= v13
                            case 4294967295:
                                delete(v5, 2147483647)
                            default:
                                switch 32767 {
                                case 348:
                                    fmt.Println(v13, v30)
                                case 628:
                                    v6 -= v6
                                }
                            }
                        case 226:
                            var v31 float32 = (v3 * v3)
                            _ = v31
                        case 127:
                            var v32 bool = (((v4 - v4) > v4) && false)
                            _ = v32
                        default:
                            if false {
                                for fi33 := 0; fi33 < 10; fi33++ {
                                    switch v2 {
                                    case "\n\t":
                                        fmt.Println(v8, v30)
                                    }
                                    delete(v5, fi33)
                                    _ = fi33
                                }
                                var v34 string = v2
                                _ = v34
                            } else {
                                delete(v1, 198)
                                var v35 float32 = v3
                                _ = v35
                            }
                        }
                        var v36 int64 = int64(v13)
                        _ = v36
                    } else {
                        var v37 uint = 255
                        var v38 uint = (^(v8))
                        var v39 float32 = v3
                        _ = v37
                        _ = v38
                        _ = v39
                    }
                    v21 = 140
                    _ = v30
                }
                _ = fi26
                _ = v27
            }
        }
        for fi40 := 0; fi40 < 4; fi40++ {
            v41 := GenS2{F0: 0}
            _ = fi40
            _ = v41
        }
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
    } else {
        if (true || false) {
            if false {
                var v42 uint = 394
                _ = v42
            }
            var v43 float64 = -2.25
            _ = v43
        } else {
            if (true || false) {
                v44 := GenS1{F0: (^(v20))}
                _ = v44
            } else {
                var v45 uint16 = v20
                fmt.Println(v4, v45, v45)
                var v46 float64 = ((float64(v3) - 3.14159) - v4)
                _ = v45
                _ = v46
            }
            fmt.Println(v8)
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
