package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 int8; F1 bool; F2 int16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint16; F1 int }
func genFn0(p0 uint16, p1 uint16, p2 int8) int16 {
    fmt.Println(p0, p2)
    p1 = (p0 << ((74) & 63))
    v5 := []int{}
    var v6 uint = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 42
}
func genFn1() uint8 {
    var v7 float32 = 1.5
    switch (("" + "a") + "hello") {
    case "hello":
        var v8 string = "\n\t"
        _ = v8
    case "foo bar":
        v9 := []int{2147483647, 0, 2147483647, 2147483647}
        _ = v9
    }
    v10 := GenS1{F0: -636}
    v11 := GenS2{F0: 1, F1: true, F2: -32768}
    _ = v10
    _ = v11
    _ = v7
    return 148
}

func main() {
    var v1 uint16 = 1
    v2 := make(map[uint32]string)
    var v3 int = -9223372036854775808
    var v4 string = "foo bar"
    fmt.Println(v3, v1)
    var v12 int = -1
    if false {
        delete(v2, 2147483648)
        switch 1 {
        case 32768:
            var v13 int = v12
            _ = v13
        default:
            if (genFn0(65535, (v1 << ((895) & 63)), 0) != 692) {
                var v14 int16 = -1
                v3 += v12
                _ = v14
            } else {
                if false {
                    for fi15 := 0; fi15 < 19; fi15++ {
                        v16 := []int{9223372036854775807, v12, 9223372036854775807}
                        var v17 int16 = 42
                        v18 := GenS3{F0: 5}
                        _ = fi15
                        _ = v16
                        _ = v17
                        _ = v18
                    }
                } else {
                    switch v3 {
                    case 128:
                        if (!true) {
                            switch "\n\t" {
                            case "\n\t":
                                var v19 int = v12
                                _ = v19
                            }
                            var v20 int8 = 0
                            v12 += v3
                            _ = v20
                        }
                    }
                    fmt.Println(v1, v4)
                }
                var v21 uint = 255
                _ = v21
            }
        }
        var v22 uint64 = 4294967295
        _ = v22
    } else {
        if (32767 != 32767) {
            v2[360] = (("" + v4) + v4)
            v23 := []int{(v12 << ((964) & 63))}
            var v24 float32 = 0.1
            _ = v23
            _ = v24
        }
        for fi25 := 0; fi25 < 4; fi25++ {
            for fi26 := 0; fi26 < 13; fi26++ {
                v3 -= fi26
                _ = len(v2)
                _ = fi26
            }
            if (v4 <= ("a" + v4)) {
                v3 += fi25
                v3 = 825
            } else {
                for fi27 := 0; fi27 < 15; fi27++ {
                    var v28 int16 = genFn0(v1, (v1 ^ (v1 >> ((1) & 63))), 0)
                    _ = fi27
                    _ = v28
                }
                v29 := []int{(fi25 ^ fi25)}
                _ = v29
            }
            var v30 float32 = 1.5
            _ = fi25
            _ = v30
        }
        fmt.Println(v4, v4)
    }
    fmt.Println(v12)
    v31 := []int{(-(v12)), 128, v12, (^(v12))}
    v12 = 0
    switch v4 {
    case "foo bar":
        var v32 bool = (true && false)
        _ = v32
    case "a":
        v33 := v2[341]
        _ = v33
    case "":
        v1 = (^(v1))
    }
    v34 := GenS4{F0: (v1 & (^(v1))), F1: v12}
    delete(v2, 4294967295)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v31
    _ = v34
    _ = v4
}
