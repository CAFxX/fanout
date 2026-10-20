package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint; F2 int16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16; F1 int; F2 string }

func main() {
    var v1 int8 = -1
    var v2 int8 = v1
    v3 := make(map[int]uint32)
    var v4 uint8 = 2
    if false {
        for fi5 := 0; fi5 < 10; fi5++ {
            v6 := []int{fi5}
            switch (-(fi5)) {
            case -319:
                v7 := GenS1{F0: -128, F1: 42, F2: -364}
                _ = v7
            }
            v1 += v1
            _ = fi5
            _ = v6
        }
    }
    switch -1 {
    case -2147483648:
        v1 -= -1
    }
    if false {
        var v8 uint64 = 0
        var v9 uint8 = v4
        var v10 uint32 = 2
        _ = v10
        _ = v8
        _ = v9
    } else {
        for fi11 := 0; fi11 < 6; fi11++ {
            v12 := []int{(-(fi11)), -9223372036854775808, 866}
            _ = fi11
            _ = v12
        }
    }
    delete(v3, 9223372036854775807)
    v1 -= v2
    for fi13 := 0; fi13 < 4; fi13++ {
        fmt.Println(v4)
        var v14 float32 = 1e10
        for fi15 := 0; fi15 < 12; fi15++ {
            v16 := GenS1{F0: v1, F1: 4294967295, F2: 1000}
            var v17 bool = true
            _ = fi15
            _ = v16
            _ = v17
        }
        _ = fi13
        _ = v14
    }
    if false {
        var v18 bool = false
        v2 = v2
        _ = v18
    } else {
        v2 += (v1 | (v1 << ((4294967295) & 63)))
        for fi19 := 0; fi19 < 19; fi19++ {
            v3[fi19] = 65535
            if (1 >= 273) {
                v20 := []int{(fi19 + 427), int(v1), (fi19 + fi19)}
                _ = v20
            } else {
                if (43 < 0) {
                    fmt.Println(v2, fi19)
                    if (!(!(!(713 != 42)))) {
                        var v21 uint64 = 113
                        _ = v21
                    } else {
                        fmt.Println(v2, v4)
                    }
                    fmt.Println(v4, fi19)
                } else {
                    v4 -= v4
                    v22 := GenS1{F0: v2, F1: 255, F2: 1}
                    v23 := []int{}
                    _ = v22
                    _ = v23
                }
                switch "hello" {
                case "foo bar":
                    v2 = v1
                case "a":
                    var v24 uint16 = 0
                    _ = v24
                case "":
                    for fi25 := 0; fi25 < 19; fi25++ {
                        v26 := GenS1{F0: 127, F1: 65535, F2: -763}
                        _ = fi25
                        _ = v26
                    }
                }
            }
            if (!(-9223372036854775808 <= 567)) {
                switch "\n\t" {
                case "foo bar":
                    fmt.Println(v2, fi19, v4)
                case "":
                    if false {
                        v27 := GenS1{F0: 11, F1: 328, F2: -1}
                        v28 := GenS2{F0: 1000}
                        if true {
                            fmt.Println(v1, v1, fi19)
                            v1 = v1
                            v29 := []int{fi19, (^(fi19)), fi19}
                            _ = v29
                        }
                        _ = v27
                        _ = v28
                    }
                case "a":
                    var v30 int = fi19
                    _ = v30
                }
                v4 -= v4
                v1 += 127
            } else {
                v4 -= v4
                fmt.Println(v1, v1)
            }
            _ = fi19
        }
    }
    v1 -= -128
    var v31 uint16 = 255
    switch -2147483648 {
    case -2147483648:
        var v32 uint = 714
        _ = v32
    case 1:
        v33 := GenS3{F0: -241, F1: 32767, F2: "\n\t"}
        _ = v33
    default:
        v34 := []int{674, -9223372036854775808}
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v4
}
