package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 int64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 uint8; F1 int32; F2 uint64 }
type GenS4 struct { F0 uint32; F1 int8; F2 int32 }
func genFn0(p0 uint32) int8 {
    v5 := []int{}
    v6 := GenS1{F0: 775, F1: 57, F2: 54}
    var v7 int16 = -1
    v8 := GenS1{F0: 32767, F1: 0, F2: -221}
    switch -9223372036854775808 {
    case 2147483647:
        var v9 int64 = 630
        _ = v9
    case 154:
        for fi10 := 0; fi10 < 11; fi10++ {
            v7 += (v7 % (((^(v7))) | 1))
            _ = fi10
        }
    }
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 127
}
func genFn1(p0 int16, p1 uint16, p2 bool) uint64 {
    v11 := GenS1{F0: 42, F1: -1, F2: 175}
    if false {
        v12 := GenS1{F0: (p1 >> ((0) & 63)), F1: 2147483647, F2: int64(p1)}
        _ = v12
    } else {
        v13 := GenS2{F0: p1}
        var v14 uint = 385
        p2 = (!((!(2147483647 > 9223372036854775807)) || ((1e10 >= (-2.25 + 0.1)) && p2)))
        _ = v13
        _ = v14
    }
    var v15 int8 = 81
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v15
    return 18446744073709551615
}

func main() {
    v1 := make(map[int]int)
    var v2 int32 = -2147483648
    v3 := make(map[string]int)
    var v4 uint64 = 1
    v16 := GenS2{F0: 346}
    v4 += (v4 * v4)
    var v17 int8 = 46
    fmt.Println(v4, v17, v2)
    fmt.Println(v2)
    fmt.Println(v17, v4, v17)
    if (genFn0(829) == (v17 + int8(v4))) {
        for fi18 := 0; fi18 < 16; fi18++ {
            v2 -= v2
            for fi19 := 0; fi19 < 8; fi19++ {
                var v20 int64 = -738
                _ = len(v1)
                switch (fi19 | (fi18 - int(v2))) {
                case 0:
                    v21 := GenS2{F0: 255}
                    _ = v21
                case -2147483648:
                    if true {
                        if (v17 == int8(v4)) {
                            v4 = 0
                            var v22 uint = 4294967295
                            var v23 int8 = v17
                            _ = v22
                            _ = v23
                        }
                        var v24 uint = 255
                        _ = len(v3)
                        _ = v24
                    }
                case 4294967295:
                    switch ("a" + (("\n\t" + ("foo bar" + "")) + "foo bar")) {
                    case "foo bar":
                        var v25 float64 = (1.5 - float64(-2.25))
                        _ = v25
                    }
                default:
                    if true {
                        v26 := GenS3{F0: 251, F1: (v2 ^ int32(v17)), F2: v4}
                        fmt.Println(v17, fi19, v4)
                        var v27 int64 = v20
                        _ = v26
                        _ = v27
                    } else {
                        var v28 bool = (255 <= 127)
                        _ = v28
                    }
                }
                fmt.Println(v17, v2, v4)
                v2 += v2
                _ = fi19
                _ = v20
            }
            _ = fi18
        }
    }
    for fi29 := 0; fi29 < 13; fi29++ {
        fmt.Println(v4, v4)
        var v30 uint16 = 260
        _ = fi29
        _ = v30
    }
    var v31 uint = 878
    var v32 int32 = 480
    v33 := GenS4{F0: 257, F1: genFn0(4294967295), F2: (v32 ^ v2)}
    fmt.Println(v2)
    for fi34 := 0; fi34 < 7; fi34++ {
        switch (("hello" + ("a" + ((("a" + "\n\t") + (("a" + "\n\t") + "hello")) + ("foo bar" + "")))) + "a") {
        case "foo bar":
            var v35 uint8 = 3
            _ = v35
        case "":
            v31 = (v31 % ((v31) | 1))
        }
        _ = fi34
    }
    v31 = v31
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v31:", v31)
    fmt.Println("v32:", v32)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
}
