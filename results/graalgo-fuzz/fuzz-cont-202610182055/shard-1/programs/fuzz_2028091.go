package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint; F1 int8 }

func main() {
    var v1 float64 = -2.25
    var v2 int32 = -2147483648
    var v3 int = -435
    v1 -= 1.5
    if (-9223372036854775808 > (v3 << ((705) & 63))) {
        var v4 int8 = 0
        fmt.Println(v4, v3, v2)
        _ = v4
    }
    if true {
        if true {
            var v5 string = "a"
            _ = v5
        }
    } else {
        var v6 uint64 = 1
        var v7 int16 = -32768
        if false {
            for fi8 := 0; fi8 < 6; fi8++ {
                var v9 uint32 = 65535
                v1 += float64(3.14159)
                _ = fi8
                _ = v9
            }
            fmt.Println(v3, v3)
        } else {
            switch "a" {
            case "\n\t":
                if true {
                    for fi10 := 0; fi10 < 10; fi10++ {
                        if true {
                            for fi11 := 0; fi11 < 8; fi11++ {
                                v7 -= (v7 >> ((37) & 63))
                                v2 = (v2 + (^(v2)))
                                _ = fi11
                            }
                            v6 += v6
                        }
                        _ = fi10
                    }
                    fmt.Println(v2, v2, v2)
                } else {
                    var v12 bool = (v6 == (v6 | v6))
                    v13 := []int{v3, (v3 / ((v3) | 1)), int(v2)}
                    _ = v12
                    _ = v13
                }
            }
        }
        _ = v6
        _ = v7
    }
    v14 := []int{99, (^(v3)), v3, (v3 * (v3 - 1))}
    var v15 uint64 = 2
    if false {
        v3 = v3
        v16 := GenS1{F0: 7}
        if len(v14) >= 2 {
            v14 = v14[1:len(v14)]
        }
        _ = v16
    } else {
        var v17 uint64 = v15
        var v18 float32 = 1.5
        if false {
            v14 = append(v14, v3)
        }
        _ = v17
        _ = v18
    }
    for fi19 := 0; fi19 < 11; fi19++ {
        for fi20 := 0; fi20 < 13; fi20++ {
            if true {
                var v21 float64 = -2.25
                _ = v21
            }
            var v22 uint16 = 32767
            switch "\n\t" {
            case "\n\t":
                switch fi20 {
                case 4294967295:
                    if len(v14) > 0 {
                        v14[0] = (fi19 ^ (v3 - v3))
                    }
                }
            case "hello":
                fmt.Println(fi19, v3, fi19)
            case "":
                v22 += (^(v22))
            }
            _ = fi20
            _ = v22
        }
        v1 -= v1
        if len(v14) > 0 {
            v14[0] = -1
        }
        _ = fi19
    }
    v15 += v15
    var v23 uint = 4294967295
    var v24 string = ""
    v23 += (v23 * 284)
    if true {
        v25 := GenS1{F0: 76}
        var v26 uint64 = 939
        if len(v14) >= 2 {
            v14 = v14[1:len(v14)]
        }
        _ = v25
        _ = v26
    }
    var v27 uint32 = 4294967295
    v28 := GenS2{F0: (v23 * (v23 - v23)), F1: -25}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v23
    _ = v24
    _ = v27
    _ = v28
    _ = v3
}
