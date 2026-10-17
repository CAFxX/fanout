package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int32; F1 int32; F2 uint64 }
type GenS4 struct { F0 int16; F1 uint64; F2 bool }
func genFn0() uint16 {
    v4 := []int{-339, 9223372036854775807}
    for fi5 := 0; fi5 < 15; fi5++ {
        var v6 int = (fi5 - fi5)
        _ = fi5
        _ = v6
    }
    var v7 float64 = 0.0
    _ = v4
    _ = v7
    return 589
}

func main() {
    var v1 float32 = 0.1
    v2 := make(map[int]bool)
    v3 := make(map[int]uint32)
    var v8 bool = (2 >= 4294967295)
    for fi9 := 0; fi9 < 14; fi9++ {
        if ((0.0 > 0.1) && v8) {
            for fi10 := 0; fi10 < 17; fi10++ {
                v11 := []int{(fi10 << ((18446744073709551615) & 63))}
                _ = fi10
                _ = v11
            }
        } else {
            fmt.Println(fi9)
            v1 = 0.0
        }
        _ = fi9
    }
    for fi12 := 0; fi12 < 13; fi12++ {
        fmt.Println(v1, v8)
        v13 := []int{fi12, (fi12 * 42)}
        v14 := GenS1{F0: 2147483647, F1: (0 == -1)}
        _ = fi12
        _ = v13
        _ = v14
    }
    for fi15 := 0; fi15 < 18; fi15++ {
        switch fi15 {
        case 9223372036854775807:
            fmt.Println(v8, fi15, fi15)
        case -405:
            for fi16 := 0; fi16 < 3; fi16++ {
                if false {
                    v8 = true
                    var v17 int8 = 127
                    v1 = (v1 - (0.1 + ((float32(1.5) * v1) - v1)))
                    _ = v17
                }
                fmt.Println(fi15, fi15, fi15)
                _ = fi16
            }
        case 127:
            if v8 {
                switch (fi15 & 0) {
                case -9223372036854775808:
                    v8 = v8
                case -890:
                    v8 = v8
                case 9223372036854775807:
                    if (42 != -9223372036854775808) {
                        for fi18 := 0; fi18 < 9; fi18++ {
                            fmt.Println(fi18)
                            switch -1 {
                            case 127:
                                if (!(1 != 90)) {
                                    if v8 {
                                        if (3.14159 != (-2.25 / (-2.25 - 0.1))) {
                                            var v19 float32 = v1
                                            _ = v19
                                        } else {
                                            v20 := GenS2{F0: 0}
                                            _ = v20
                                        }
                                    } else {
                                        var v21 uint16 = 65535
                                        v21 = 638
                                        var v22 int = (-(fi15))
                                        _ = v21
                                        _ = v22
                                    }
                                    v1 = v1
                                } else {
                                    v1 -= 0.1
                                }
                            default:
                                v23 := GenS3{F0: 2147483647, F1: 0, F2: 371}
                                _ = v23
                            }
                            _ = fi18
                        }
                    } else {
                        if ("foo bar" != ("hello" + (("hello" + ("foo bar" + "hello")) + "foo bar"))) {
                            v24 := v3[(fi15 << ((65535) & 63))]
                            _ = v24
                        }
                    }
                }
            }
        }
        var v25 uint32 = 907
        var v26 uint = 703
        _ = fi15
        _ = v25
        _ = v26
    }
    delete(v2, 42)
    var v27 string = "foo bar"
    var v28 uint64 = 18446744073709551615
    v29 := GenS4{F0: -32768, F1: 4294967295, F2: (!(85 >= 49))}
    var v30 uint16 = uint16(v28)
    fmt.Println(v1, v8, v30)
    v27 = "a"
    var v31 uint = 4294967295
    if (v1 == (v1 / (1e10 - v1))) {
        var v32 bool = false
        _ = v32
    }
    for fi33 := 0; fi33 < 5; fi33++ {
        v34 := []int{fi33, 32767, 2147483647, (fi33 * 128)}
        _ = fi33
        _ = v34
    }
    var v35 uint8 = 104
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v35:", v35)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v35
    _ = v8
}
