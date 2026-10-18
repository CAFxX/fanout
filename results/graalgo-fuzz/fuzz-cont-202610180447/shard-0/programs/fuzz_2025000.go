package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 bool; F1 int32; F2 uint8 }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]int)
    v3 := make(map[uint32]string)
    var v4 int16 = -641
    var v5 int64 = -223
    var v6 int16 = int16(v5)
    if false {
        for fi7 := 0; fi7 < 12; fi7++ {
            v5 = 1
            v6 += (v4 + v4)
            _ = fi7
        }
        v8 := GenS1{F0: 100, F1: (1e10 > (1e10 + -2.25))}
        for fi9 := 0; fi9 < 16; fi9++ {
            var v10 string = "foo bar"
            var v11 string = v10
            var v12 int32 = -661
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
        }
        _ = v8
    } else {
        var v13 bool = false
        v4 = (v6 | v4)
        v14 := []int{127, 42}
        _ = v13
        _ = v14
    }
    var v15 float64 = 1e10
    switch 32767 {
    case 32768:
        var v16 int64 = v5
        _ = v16
    }
    var v17 string = "\n\t"
    var v18 float32 = (float32(1e10) - 3.14159)
    v17 = (v17 + v17)
    v17 = v17
    fmt.Println(v17, v15)
    switch 128 {
    case 2147483647:
        for fi19 := 0; fi19 < 10; fi19++ {
            v20 := []int{fi19, fi19}
            if (42 != 239) {
                for fi21 := 0; fi21 < 7; fi21++ {
                    _ = len(v1)
                    switch v17 {
                    case "a":
                        if len(v20) >= 2 {
                            v20 = v20[1:len(v20)]
                        }
                    case "":
                        if len(v20) > 0 {
                            v20[0] = (fi21 & (fi21 - fi19))
                        }
                    default:
                        v22 := GenS1{F0: 127, F1: (true || (-128 == 92))}
                        v5 = (v5 ^ (-(v5)))
                        v18 -= 1e10
                        fmt.Println(v6)
                        var v23 uint16 = 255
                        v24 := GenS2{F0: false}
                        if false {
                            v3[0] = "hello"
                        } else {
                            v23 = (v23 % ((v23) | 1))
                        }
                        _ = v22
                        _ = v23
                        _ = v24
                    }
                    _ = fi21
                }
                var v25 uint32 = 2147483648
                if (865 > 42) {
                    var v26 uint8 = 55
                    v25 -= (v25 ^ (v25 << ((0) & 63)))
                    switch (fi19 ^ 2147483647) {
                    case -9223372036854775808:
                        for fi27 := 0; fi27 < 3; fi27++ {
                            var v28 uint64 = 42
                            switch v17 {
                            case "hello":
                                fmt.Println(v4)
                            case "":
                                fmt.Println(v26, v17)
                            case "foo bar":
                                var v29 int64 = int64(v6)
                                _ = v29
                            }
                            if false {
                                fmt.Println(v4)
                            }
                            _ = fi27
                            _ = v28
                        }
                    default:
                        for fi30 := 0; fi30 < 9; fi30++ {
                            fmt.Println(v5)
                            var v31 int64 = (v5 + v5)
                            delete(v2, fi30)
                            _ = fi30
                            _ = v31
                        }
                    }
                    _ = v26
                } else {
                    var v32 uint = 42
                    switch int(v6) {
                    case 397:
                        var v33 int64 = (v5 + v5)
                        _ = v33
                    default:
                        var v34 int8 = 42
                        _ = v34
                    }
                    for fi35 := 0; fi35 < 13; fi35++ {
                        fmt.Println(v25, v15, v5)
                        if false {
                            switch (fi19 & fi35) {
                            case 148:
                                _ = len(v1)
                            }
                            switch (fi19 ^ 0) {
                            case 127:
                                var v36 uint = v32
                                _ = v36
                            default:
                                v4 += 42
                            }
                        } else {
                            var v37 uint16 = 410
                            _ = v37
                        }
                        v38 := GenS2{F0: (-2147483648 >= 964)}
                        _ = fi35
                        _ = v38
                    }
                    _ = v32
                }
                _ = v25
            }
            _ = fi19
            _ = v20
        }
    case -9223372036854775808:
        switch v17 {
        case "\n\t":
            v2[1] = 4294967295
        default:
            v39 := GenS3{F0: (v18 > (v18 * float32(float64(float32(3.14159))))), F1: 1, F2: 1}
            _ = v39
        }
    default:
        delete(v2, 32768)
    }
    if false {
        v40 := GenS3{F0: true, F1: 2147483647, F2: 0}
        var v41 bool = (false || true)
        var v42 int32 = 0
        _ = v40
        _ = v41
        _ = v42
    } else {
        for fi43 := 0; fi43 < 15; fi43++ {
            v44 := []int{fi43}
            _ = fi43
            _ = v44
        }
    }
    v1[127] = v17
    var v45 bool = (v5 == (v5 >> ((18446744073709551615) & 63)))
    v6 = v4
    fmt.Println(v4, v15, v17)
    var v46 uint8 = 1
    if true {
        fmt.Println(v6)
        v15 -= 1.5
        v47 := GenS3{F0: (!((-(v6)) == v4)), F1: -2147483648, F2: (v46 * 109)}
        _ = v47
    }
    var v48 uint64 = 42
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v45:", v45)
    fmt.Println("v46:", v46)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v15
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v45
    _ = v46
    _ = v48
    _ = v5
    _ = v6
}
